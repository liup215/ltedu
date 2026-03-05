package service

import (
	"edu/model"
	"edu/repository"
	"sort"
	"time"
)

// RecommendationSvr is the singleton recommendation service.
var RecommendationSvr = &RecommendationService{baseService: newBaseService()}

// RecommendationService provides AI-driven practice question recommendations.
type RecommendationService struct {
	baseService
}

const (
	// masteryThreshold is the mastery level below which a chapter needs reinforcement.
	masteryThreshold = 0.7
	// maxRecommendations is the maximum number of recommended questions to return.
	maxRecommendations = 10
)

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
		if ks.ConsecutiveWrong >= 3 {
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
