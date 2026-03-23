package service

import (
	"edu/model"
	"edu/repository"
	"sort"
	"time"
)

// RecommendationSvr is the singleton recommendation service.
var RecommendationSvr = &RecommendationService{baseService: newBaseService()}

// RecommendationService provides personalized question recommendations for students.
type RecommendationService struct {
	baseService
}

const (
	// masteryThreshold is the mastery level below which a chapter needs reinforcement.
	masteryThreshold = 0.7
	// maxRecommendations is the maximum number of recommended questions to return.
	maxRecommendations = 10
	// consecutiveWrongThreshold triggers high-priority recommendations.
	consecutiveWrongThreshold = 3
	// collaborativeNeighborLimit is the max number of peer students used in collaborative filtering.
	collaborativeNeighborLimit = 20
)

// GetQuestionRecommendations returns personalized question recommendations for a student.
//
// Algorithm (four-signal ranking):
//  1. SRS review due          → highest priority (signal 4.0)
//  2. Consecutive wrong ≥ 3   → high priority    (signal 3.0)
//  3. Low mastery (< 70 %)    → medium priority  (signal 2.0)
//  4. Collaborative filtering  → additional boost for questions that helped similar students
//
// For each prioritised chapter the service selects questions linked to its
// KnowledgePoints, adapts the difficulty to the student's mastery level, and
// decorates every recommendation with a human-readable explanation.
func (svr *RecommendationService) GetQuestionRecommendations(studentID uint) (*model.QuestionRecommendationResponse, error) {
	db := repository.GetDB()

	// ── 1. Load all active goals for the student ──────────────────────────
	var goals []model.Goal
	if err := db.Where("user_id = ? AND status = ?", studentID, "active").Find(&goals).Error; err != nil {
		return nil, err
	}

	// ── 2. Collect all knowledge states across goals ──────────────────────
	var allStates []model.KnowledgeState
	for _, g := range goals {
		var states []model.KnowledgeState
		if err := db.Where("user_id = ? AND goal_id = ?", studentID, g.ID).
			Preload("Chapter").
			Find(&states).Error; err != nil {
			return nil, err
		}
		allStates = append(allStates, states...)
	}

	now := time.Now()
	weakAreaCount := 0
	reviewDueCount := 0

	// chapterPriority maps chapterId → (priority, reason, mastery)
	type chapterInfo struct {
		priority    float64
		reason      string
		mastery     float64
		chapterName string
		goalID      uint
		syllabusID  uint
	}
	priorityMap := make(map[uint]chapterInfo)

	for _, ks := range allStates {
		if ks.Chapter.ID == 0 {
			continue
		}
		info := chapterInfo{
			mastery:     ks.MasteryLevel,
			chapterName: ks.Chapter.Name,
			goalID:      ks.GoalId,
		}

		// Signal 1: SRS review overdue
		if ks.NextReviewAt != nil && !ks.NextReviewAt.After(now) {
			info.priority = 4.0
			info.reason = "This chapter is due for review based on your spaced-repetition schedule"
			reviewDueCount++
		} else if ks.ConsecutiveWrong >= consecutiveWrongThreshold {
			// Signal 2: streak of wrong answers
			info.priority = 3.0
			info.reason = "You have had several consecutive wrong answers in this chapter"
		} else if ks.IsCovered && ks.MasteryLevel < masteryThreshold {
			// Signal 3: low mastery
			info.priority = 2.0
			info.reason = "Your mastery level in this chapter is below the recommended threshold"
			weakAreaCount++
		} else {
			// Skip chapters that are in good shape
			continue
		}

		// Keep only the highest-priority signal per chapter
		if existing, ok := priorityMap[ks.ChapterId]; !ok || info.priority > existing.priority {
			priorityMap[ks.ChapterId] = info
		}
	}

	// ── 3. Collaborative filtering boost ─────────────────────────────────
	// Find peer students who share the same goals (same syllabus) and look at
	// which questions they got right after previously struggling in the same
	// chapters. Boost those questions' ranking score.
	collaborativeBoost := svr.collaborativeFilteringBoost(studentID, goals)

	// ── 4. For each prioritised chapter, find matching questions ──────────
	type candidate struct {
		rec   model.QuestionRecommendation
		score float64
	}
	var candidates []candidate

	for chapterID, info := range priorityMap {
		// Determine adapted difficulty range based on mastery
		minDiff, maxDiff := adaptedDifficultyRange(info.mastery)

		// Load KnowledgePoints for this chapter
		kps, err := repository.KnowledgePointRepo.FindByChapterId(chapterID)
		if err != nil || len(kps) == 0 {
			continue
		}

		// Collect question IDs linked to these knowledge points
		var kpIDs []uint
		kpNameByID := make(map[uint]string)
		for _, kp := range kps {
			kpIDs = append(kpIDs, kp.ID)
			kpNameByID[kp.ID] = kp.Name
		}

		// Fetch questions via join table question_keypoints
		type qkRow struct {
			QuestionID       uint
			KnowledgePointID uint
		}
		var rows []qkRow
		if err := db.Table(db.NamingStrategy.JoinTableName("question_keypoints")).
			Select("question_id, knowledge_point_id").
			Where("knowledge_point_id IN ?", kpIDs).
			Find(&rows).Error; err != nil {
			continue
		}

		// Track which questions the student has already attempted recently
		alreadyAttempted := svr.recentlyAttemptedQuestions(studentID, chapterID)

		seen := make(map[uint]bool)
		for _, row := range rows {
			if seen[row.QuestionID] {
				continue
			}
			seen[row.QuestionID] = true

			// Skip questions attempted very recently (within last 24 h)
			if alreadyAttempted[row.QuestionID] {
				continue
			}

			q, err := repository.QuestionRepo.FindByID(row.QuestionID)
			if err != nil || q == nil {
				continue
			}
			if q.Status != model.QUESTION_STATE_NORMAL {
				continue
			}
			// Adaptive difficulty filter
			if q.Difficult < minDiff || q.Difficult > maxDiff {
				continue
			}

			// Base score from chapter priority + difficulty match
			score := info.priority + difficultyMatchScore(q.Difficult, info.mastery)
			// Apply collaborative boost if available
			score += collaborativeBoost[row.QuestionID]

			candidates = append(candidates, candidate{
				rec: model.QuestionRecommendation{
					QuestionID:       q.ID,
					Question:         q,
					KnowledgePointID: row.KnowledgePointID,
					KnowledgePoint:   kpNameByID[row.KnowledgePointID],
					ChapterID:        chapterID,
					ChapterName:      info.chapterName,
					Reason:           info.reason,
					Priority:         priorityToInt(info.priority),
					Difficulty:       q.Difficult,
					MasteryLevel:     info.mastery,
				},
				score: score,
			})
		}
	}

	// ── 5. Sort by score desc, deduplicate, cap at maxRecommendations ─────
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score > candidates[j].score
	})

	recommendations := make([]model.QuestionRecommendation, 0, maxRecommendations)
	seen := make(map[uint]bool)
	for _, c := range candidates {
		if seen[c.rec.QuestionID] {
			continue
		}
		seen[c.rec.QuestionID] = true
		c.rec.Score = c.score
		recommendations = append(recommendations, c.rec)
		if len(recommendations) >= maxRecommendations {
			break
		}
	}

	return &model.QuestionRecommendationResponse{
		StudentID:       studentID,
		Recommendations: recommendations,
		TotalCount:      len(recommendations),
		WeakAreaCount:   weakAreaCount,
		ReviewDueCount:  reviewDueCount,
	}, nil
}

// collaborativeFilteringBoost computes a score boost for questions that helped
// similar students (same syllabus, same weak chapters) improve their mastery.
// It returns a map of questionID → boost value.
func (svr *RecommendationService) collaborativeFilteringBoost(
	studentID uint,
	goals []model.Goal,
) map[uint]float64 {
	boost := make(map[uint]float64)

	gormDB := repository.GetDB()
	if gormDB == nil {
		return boost
	}

	for _, goal := range goals {
		// Find peer students who have a goal on the same syllabus
		var peerIDs []uint
		if err := gormDB.Model(&model.Goal{}).
			Select("user_id").
			Where("syllabus_id = ? AND user_id != ? AND status = ?", goal.SyllabusId, studentID, "active").
			Limit(collaborativeNeighborLimit).
			Pluck("user_id", &peerIDs).Error; err != nil || len(peerIDs) == 0 {
			continue
		}

		// Find questions that peers answered correctly AND in which the student
		// still has low mastery chapters.  We use a simple proxy: questions
		// that peers got right and the student hasn't attempted recently.
		type boostRow struct {
			QuestionID uint
			CorrectRate float64
		}
		var boostRows []boostRow
		if err := gormDB.
			Table("lt_attempt").
			Select("question_id, AVG(CASE WHEN is_correct = 1 THEN 1.0 ELSE 0.0 END) AS correct_rate").
			Where("user_id IN ? AND goal_id IN (?)",
				peerIDs,
				gormDB.Model(&model.Goal{}).
					Select("id").
					Where("user_id IN ? AND syllabus_id = ?", peerIDs, goal.SyllabusId),
			).
			Group("question_id").
			Having("correct_rate >= 0.6").
			Scan(&boostRows).Error; err != nil {
			continue
		}

		for _, br := range boostRows {
			// Boost up to 0.5 based on peer correct rate
			if br.CorrectRate > boost[br.QuestionID] {
				boost[br.QuestionID] = br.CorrectRate * 0.5
			}
		}
	}
	return boost
}

// recentlyAttemptedQuestions returns the set of question IDs that a student
// has attempted in the given chapter within the last 24 hours.
func (svr *RecommendationService) recentlyAttemptedQuestions(studentID, chapterID uint) map[uint]bool {
	result := make(map[uint]bool)
	gormDB := repository.GetDB()
	if gormDB == nil {
		return result
	}
	cutoff := time.Now().Add(-24 * time.Hour)
	var ids []uint
	gormDB.Model(&model.Attempt{}).
		Select("question_id").
		Where("user_id = ? AND chapter_id = ? AND attempted_at > ?", studentID, chapterID, cutoff).
		Pluck("question_id", &ids)
	for _, id := range ids {
		result[id] = true
	}
	return result
}

// adaptedDifficultyRange returns a [min, max] difficulty range suited to the
// student's current mastery level using an adaptive difficulty strategy.
//
// Mastery < 0.3  → easy questions (1–2)
// Mastery 0.3–0.6 → medium questions (2–3)
// Mastery 0.6–0.8 → medium-hard (2–4)
// Mastery ≥ 0.8  → challenging (3–5)
func adaptedDifficultyRange(mastery float64) (min, max int) {
	switch {
	case mastery < 0.3:
		return 1, 2
	case mastery < 0.6:
		return 2, 3
	case mastery < 0.8:
		return 2, 4
	default:
		return 3, 5
	}
}

// difficultyMatchScore returns a small positive bonus when the question's
// difficulty is well-matched to the student's mastery level.
func difficultyMatchScore(difficult int, mastery float64) float64 {
	_, max := adaptedDifficultyRange(mastery)
	// Best match is a question at the upper end of the adaptive range (slight challenge)
	if difficult == max {
		return 0.3
	}
	return 0.0
}

// priorityToInt converts a float priority to its integer equivalent (1 = highest).
func priorityToInt(p float64) int {
	switch {
	case p >= 4.0:
		return 1
	case p >= 3.0:
		return 2
	default:
		return 3
	}
}

// GetRecommendations returns a personalized list of question recommendations for a student's goal.
// The algorithm:
//  1. Chapters due for SRS review   → highest priority
//  2. Chapters with consecutive wrong answers  → high priority
//  3. Chapters with low mastery (< 70 %)       → medium priority
//  4. Uncovered chapters that the student hasn't touched yet → low priority
func (svr *RecommendationService) GetRecommendations(userID, goalID uint) (*model.RecommendationResponse, error) {
	db := repository.GetDB()

	// Verify goal belongs to the user
	goal, err := repository.GoalRepo.GetByID(goalID)
	if err != nil || goal == nil || goal.UserId != userID {
		return nil, err
	}

	// Load all knowledge states for this goal
	var states []model.KnowledgeState
	if err := db.Where("user_id = ? AND goal_id = ?", userID, goalID).
		Preload("Chapter").
		Find(&states).Error; err != nil {
		return nil, err
	}

	type candidate struct {
		chapterID   uint
		chapterName string
		mastery     float64
		reason      string
		priority    int // lower = more urgent
	}

	var candidates []candidate
	now := time.Now()
	reviewCount := 0
	gapCount := 0

	for _, ks := range states {
		// 1. SRS due review
		if ks.NextReviewAt != nil && !ks.NextReviewAt.After(now) {
			reviewCount++
			candidates = append(candidates, candidate{
				chapterID:   ks.ChapterId,
				chapterName: ks.Chapter.Name,
				mastery:     ks.MasteryLevel,
				reason:      "Due for spaced-repetition review",
				priority:    1,
			})
			continue
		}

		// 2. Consecutive wrong answers
		if ks.ConsecutiveWrong >= consecutiveWrongThreshold {
			candidates = append(candidates, candidate{
				chapterID:   ks.ChapterId,
				chapterName: ks.Chapter.Name,
				mastery:     ks.MasteryLevel,
				reason:      "Recent consecutive wrong answers – reinforcement needed",
				priority:    2,
			})
			continue
		}

		// 3. Low mastery
		if ks.IsCovered && ks.MasteryLevel < masteryThreshold {
			gapCount++
			candidates = append(candidates, candidate{
				chapterID:   ks.ChapterId,
				chapterName: ks.Chapter.Name,
				mastery:     ks.MasteryLevel,
				reason:      "Mastery below 70% – practice more",
				priority:    3,
			})
			continue
		}

		// 4. Not yet covered
		if !ks.IsCovered {
			candidates = append(candidates, candidate{
				chapterID:   ks.ChapterId,
				chapterName: ks.Chapter.Name,
				mastery:     ks.MasteryLevel,
				reason:      "Not yet covered – start practicing",
				priority:    4,
			})
		}
	}

	// Sort by priority, then by mastery ascending (lowest mastery first within same priority)
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].priority != candidates[j].priority {
			return candidates[i].priority < candidates[j].priority
		}
		return candidates[i].mastery < candidates[j].mastery
	})

	// For each top candidate chapter, pick a random un-attempted question
	var recommendations []model.RecommendedQuestion
	seen := make(map[uint]bool)

	for _, c := range candidates {
		if len(recommendations) >= maxRecommendations {
			break
		}
		if seen[c.chapterID] {
			continue
		}
		seen[c.chapterID] = true

		recommendations = append(recommendations, model.RecommendedQuestion{
			ChapterID:    c.chapterID,
			ChapterName:  c.chapterName,
			Reason:       c.reason,
			Priority:     c.priority,
			MasteryLevel: round2(c.mastery * 100),
		})
	}

	return &model.RecommendationResponse{
		GoalID:      goalID,
		Questions:   recommendations,
		ReviewCount: reviewCount,
		GapCount:    gapCount,
	}, nil
}
