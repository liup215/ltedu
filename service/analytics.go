package service

import (
	"edu/model"
	"edu/repository"
	"time"
)

// AnalyticsSvr is the singleton analytics service.
var AnalyticsSvr = &AnalyticsService{baseService: newBaseService()}

// AnalyticsService provides learning analytics for teachers and students.
type AnalyticsService struct {
	baseService
}

// atRiskThresholds defines when a student is considered at-risk.
const (
	atRiskInactiveDays      = 7   // no practice in N days
	atRiskConsecWrong       = 5   // N or more consecutive wrong answers
	atRiskLowMastery        = 0.3 // mastery below 30 %
	atRiskLowCoverage       = 0.2 // coverage below 20 %
)

// GetClassSummary returns a high-level performance summary for all students in a class.
func (svr *AnalyticsService) GetClassSummary(classID uint) (*model.ClassPerformanceSummary, error) {
	db := repository.GetDB()

	// Fetch class info
	class, err := repository.ClassRepo.FindByID(classID)
	if err != nil || class == nil {
		return nil, err
	}

	// Fetch students in this class
	students, err := repository.ClassRepo.FindStudents(classID)
	if err != nil {
		return nil, err
	}

	summary := &model.ClassPerformanceSummary{
		ClassID:       classID,
		ClassName:     class.Name,
		TotalStudents: len(students),
	}

	if len(students) == 0 {
		return summary, nil
	}

	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	var totalMastery, totalCoverage, totalAccuracy float64

	for _, student := range students {
		// Aggregate knowledge states for this student
		var states []model.KnowledgeState
		db.Where("user_id = ?", student.ID).Find(&states)

		if len(states) == 0 {
			continue
		}

		var mastery, coverage float64
		coveredCount := 0
		for _, ks := range states {
			mastery += ks.MasteryLevel
			if ks.IsCovered {
				coveredCount++
			}
		}
		mastery = mastery / float64(len(states)) * 100
		coverage = float64(coveredCount) / float64(len(states)) * 100

		totalMastery += mastery
		totalCoverage += coverage

		// Attempt stats for accuracy
		var correctCount, totalAttempts int64
		db.Model(&model.Attempt{}).Where("user_id = ?", student.ID).Count(&totalAttempts)
		db.Model(&model.Attempt{}).Where("user_id = ? AND is_correct = ?", student.ID, true).Count(&correctCount)
		if totalAttempts > 0 {
			totalAccuracy += float64(correctCount) / float64(totalAttempts) * 100
		}
		summary.TotalAttempts += totalAttempts

		// Weekly attempts
		var weeklyAttempts int64
		db.Model(&model.Attempt{}).Where("user_id = ? AND attempted_at >= ?", student.ID, sevenDaysAgo).Count(&weeklyAttempts)
		summary.WeeklyAttempts += weeklyAttempts
		if weeklyAttempts > 0 {
			summary.ActiveStudents++
		}

		// Early warning check
		reasons := svr.evaluateRiskReasons(states, totalAttempts, sevenDaysAgo, student.ID)
		if len(reasons) > 0 {
			summary.AtRiskCount++
		}
	}

	n := float64(len(students))
	summary.AvgMastery = round2(totalMastery / n)
	summary.AvgCoverage = round2(totalCoverage / n)
	summary.AvgAccuracy = round2(totalAccuracy / n)

	return summary, nil
}

// GetStudentPerformanceList returns per-student analytics for a class.
func (svr *AnalyticsService) GetStudentPerformanceList(classID uint) ([]model.StudentPerformanceSummary, error) {
	db := repository.GetDB()

	students, err := repository.ClassRepo.FindStudents(classID)
	if err != nil {
		return nil, err
	}

	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	var result []model.StudentPerformanceSummary

	for _, student := range students {
		var states []model.KnowledgeState
		db.Where("user_id = ?", student.ID).Find(&states)

		var mastery, coverage float64
		var maxConsecWrong int
		coveredCount := 0
		for _, ks := range states {
			mastery += ks.MasteryLevel
			if ks.IsCovered {
				coveredCount++
			}
			if ks.ConsecutiveWrong > maxConsecWrong {
				maxConsecWrong = ks.ConsecutiveWrong
			}
		}
		if len(states) > 0 {
			mastery = mastery / float64(len(states)) * 100
			coverage = float64(coveredCount) / float64(len(states)) * 100
		}

		var correctCount, totalAttempts int64
		db.Model(&model.Attempt{}).Where("user_id = ?", student.ID).Count(&totalAttempts)
		db.Model(&model.Attempt{}).Where("user_id = ? AND is_correct = ?", student.ID, true).Count(&correctCount)
		var accuracy float64
		if totalAttempts > 0 {
			accuracy = float64(correctCount) / float64(totalAttempts) * 100
		}

		var weeklyAttempts int64
		db.Model(&model.Attempt{}).Where("user_id = ? AND attempted_at >= ?", student.ID, sevenDaysAgo).Count(&weeklyAttempts)

		// Find last active time
		var lastAttempt model.Attempt
		var lastActiveAt *time.Time
		if err := db.Where("user_id = ?", student.ID).Order("attempted_at DESC").First(&lastAttempt).Error; err == nil {
			lastActiveAt = &lastAttempt.AttemptedAt
		}

		reasons := svr.evaluateRiskReasons(states, totalAttempts, sevenDaysAgo, student.ID)

		sp := model.StudentPerformanceSummary{
			UserID:         student.ID,
			UserName:       student.Username,
			MasteryLevel:   round2(mastery),
			CoverageLevel:  round2(coverage),
			AccuracyRate:   round2(accuracy),
			TotalAttempts:  totalAttempts,
			WeeklyAttempts: weeklyAttempts,
			ConsecWrong:    maxConsecWrong,
			LastActiveAt:   lastActiveAt,
			IsAtRisk:       len(reasons) > 0,
			RiskReasons:    reasons,
		}
		result = append(result, sp)
	}

	return result, nil
}

// GetClassHeatmap returns a chapter × student mastery heatmap for a class.
func (svr *AnalyticsService) GetClassHeatmap(classID uint) (*model.ClassHeatmap, error) {
	db := repository.GetDB()

	students, err := repository.ClassRepo.FindStudents(classID)
	if err != nil {
		return nil, err
	}

	if len(students) == 0 {
		return &model.ClassHeatmap{}, nil
	}

	// Collect all chapter IDs that at least one student has a knowledge state for
	studentIDs := make([]uint, len(students))
	studentMap := make(map[uint]string)
	for i, s := range students {
		studentIDs[i] = s.ID
		studentMap[s.ID] = s.Username
	}

	var allStates []model.KnowledgeState
	db.Where("user_id IN (?)", studentIDs).Preload("Chapter").Find(&allStates)

	// Build chapter → student → KS map
	type chapterKey = uint
	chapterStudentMap := make(map[chapterKey]map[uint]*model.KnowledgeState)
	chapterNames := make(map[uint]string)
	for i := range allStates {
		ks := &allStates[i]
		if _, ok := chapterStudentMap[ks.ChapterId]; !ok {
			chapterStudentMap[ks.ChapterId] = make(map[uint]*model.KnowledgeState)
		}
		chapterStudentMap[ks.ChapterId][ks.UserId] = ks
		if ks.Chapter.Name != "" {
			chapterNames[ks.ChapterId] = ks.Chapter.Name
		}
	}

	// Build ordered student list
	studentInfos := make([]model.StudentInfo, len(students))
	for i, s := range students {
		studentInfos[i] = model.StudentInfo{UserID: s.ID, UserName: s.Username}
	}

	// Build heatmap rows per chapter
	var rows []model.ClassHeatmapRow
	for chapterID, studentKSMap := range chapterStudentMap {
		row := model.ClassHeatmapRow{
			ChapterID:   chapterID,
			ChapterName: chapterNames[chapterID],
		}
		var totalMastery float64
		for _, student := range students {
			score := model.StudentChapterScore{UserID: student.ID}
			if ks, ok := studentKSMap[student.ID]; ok {
				score.MasteryLevel = round2(ks.MasteryLevel * 100)
				score.IsCovered = ks.IsCovered
				totalMastery += ks.MasteryLevel
			}
			row.StudentData = append(row.StudentData, score)
		}
		row.AvgMastery = round2(totalMastery / float64(len(students)) * 100)
		rows = append(rows, row)
	}

	return &model.ClassHeatmap{
		Students: studentInfos,
		Chapters: rows,
	}, nil
}

// GetAttemptTrends returns a time-series of attempt counts for a class.
func (svr *AnalyticsService) GetAttemptTrends(classID uint, startDate, endDate time.Time) ([]model.AttemptTrendPoint, error) {
	db := repository.GetDB()

	students, err := repository.ClassRepo.FindStudents(classID)
	if err != nil {
		return nil, err
	}
	if len(students) == 0 {
		return nil, nil
	}

	studentIDs := make([]uint, len(students))
	for i, s := range students {
		studentIDs[i] = s.ID
	}

	var attempts []model.Attempt
	db.Where("user_id IN (?) AND attempted_at BETWEEN ? AND ?", studentIDs, startDate, endDate).Find(&attempts)

	// Bucket by date
	totals := make(map[string]int)
	corrects := make(map[string]int)
	for _, a := range attempts {
		date := a.AttemptedAt.Format("2006-01-02")
		totals[date]++
		if a.IsCorrect != nil && *a.IsCorrect {
			corrects[date]++
		}
	}

	var result []model.AttemptTrendPoint
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		date := d.Format("2006-01-02")
		result = append(result, model.AttemptTrendPoint{
			Date:            date,
			TotalAttempts:   totals[date],
			CorrectAttempts: corrects[date],
		})
	}
	return result, nil
}

// GetEarlyWarnings returns a list of students who may be struggling.
func (svr *AnalyticsService) GetEarlyWarnings(classID uint) ([]model.EarlyWarningStudent, error) {
	db := repository.GetDB()

	students, err := repository.ClassRepo.FindStudents(classID)
	if err != nil {
		return nil, err
	}

	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	var warnings []model.EarlyWarningStudent

	for _, student := range students {
		var states []model.KnowledgeState
		db.Where("user_id = ?", student.ID).Find(&states)

		var totalAttempts int64
		db.Model(&model.Attempt{}).Where("user_id = ?", student.ID).Count(&totalAttempts)

		reasons := svr.evaluateRiskReasons(states, totalAttempts, sevenDaysAgo, student.ID)
		if len(reasons) == 0 {
			continue
		}

		severity := "low"
		if len(reasons) >= 3 {
			severity = "high"
		} else if len(reasons) == 2 {
			severity = "medium"
		}

		warnings = append(warnings, model.EarlyWarningStudent{
			UserID:   student.ID,
			UserName: student.Username,
			Reasons:  reasons,
			Severity: severity,
		})
	}

	return warnings, nil
}

// GetStudentAnalytics returns personal analytics for a single student's goal.
func (svr *AnalyticsService) GetStudentAnalytics(userID, goalID uint) (*model.StudentPerformanceSummary, error) {
	db := repository.GetDB()

	var states []model.KnowledgeState
	db.Where("user_id = ? AND goal_id = ?", userID, goalID).Preload("Chapter").Find(&states)

	var mastery, coverage float64
	var maxConsecWrong int
	coveredCount := 0
	for _, ks := range states {
		mastery += ks.MasteryLevel
		if ks.IsCovered {
			coveredCount++
		}
		if ks.ConsecutiveWrong > maxConsecWrong {
			maxConsecWrong = ks.ConsecutiveWrong
		}
	}
	if len(states) > 0 {
		mastery = mastery / float64(len(states)) * 100
		coverage = float64(coveredCount) / float64(len(states)) * 100
	}

	var correctCount, totalAttempts int64
	db.Model(&model.Attempt{}).Where("user_id = ? AND goal_id = ?", userID, goalID).Count(&totalAttempts)
	db.Model(&model.Attempt{}).Where("user_id = ? AND goal_id = ? AND is_correct = ?", userID, goalID, true).Count(&correctCount)
	var accuracy float64
	if totalAttempts > 0 {
		accuracy = float64(correctCount) / float64(totalAttempts) * 100
	}

	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	var weeklyAttempts int64
	db.Model(&model.Attempt{}).Where("user_id = ? AND goal_id = ? AND attempted_at >= ?", userID, goalID, sevenDaysAgo).Count(&weeklyAttempts)

	var lastAttempt model.Attempt
	var lastActiveAt *time.Time
	if err := db.Where("user_id = ? AND goal_id = ?", userID, goalID).Order("attempted_at DESC").First(&lastAttempt).Error; err == nil {
		lastActiveAt = &lastAttempt.AttemptedAt
	}

	reasons := svr.evaluateRiskReasons(states, totalAttempts, sevenDaysAgo, userID)

	return &model.StudentPerformanceSummary{
		UserID:         userID,
		MasteryLevel:   round2(mastery),
		CoverageLevel:  round2(coverage),
		AccuracyRate:   round2(accuracy),
		TotalAttempts:  totalAttempts,
		WeeklyAttempts: weeklyAttempts,
		ConsecWrong:    maxConsecWrong,
		LastActiveAt:   lastActiveAt,
		IsAtRisk:       len(reasons) > 0,
		RiskReasons:    reasons,
	}, nil
}

// evaluateRiskReasons returns the list of risk flags for a student based on their knowledge states.
func (svr *AnalyticsService) evaluateRiskReasons(states []model.KnowledgeState, totalAttempts int64, sevenDaysAgo time.Time, userID uint) []string {
	db := repository.GetDB()
	var reasons []string

	if len(states) == 0 && totalAttempts == 0 {
		reasons = append(reasons, "No learning activity recorded")
		return reasons
	}

	// Check inactivity
	var recentAttempts int64
	db.Model(&model.Attempt{}).Where("user_id = ? AND attempted_at >= ?", userID, sevenDaysAgo).Count(&recentAttempts)
	if recentAttempts == 0 && totalAttempts > 0 {
		reasons = append(reasons, "No practice in the last 7 days")
	}

	// Check consecutive wrong answers
	for _, ks := range states {
		if ks.ConsecutiveWrong >= atRiskConsecWrong {
			reasons = append(reasons, "Multiple consecutive wrong answers in a chapter")
			break
		}
	}

	// Check average mastery
	if len(states) > 0 {
		var totalMastery float64
		for _, ks := range states {
			totalMastery += ks.MasteryLevel
		}
		avgMastery := totalMastery / float64(len(states))
		if avgMastery < atRiskLowMastery {
			reasons = append(reasons, "Overall mastery below 30%")
		}

		// Check coverage
		covered := 0
		for _, ks := range states {
			if ks.IsCovered {
				covered++
			}
		}
		if float64(covered)/float64(len(states)) < atRiskLowCoverage {
			reasons = append(reasons, "Less than 20% of syllabus covered")
		}
	}

	return reasons
}

// round2 rounds a float64 to 2 decimal places.
func round2(v float64) float64 {
	return float64(int(v*100+0.5)) / 100
}
