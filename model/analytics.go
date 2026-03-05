package model

import "time"

// ClassPerformanceSummary aggregates class-wide learning analytics
type ClassPerformanceSummary struct {
	ClassID         uint    `json:"classId"`
	ClassName       string  `json:"className"`
	TotalStudents   int     `json:"totalStudents"`
	ActiveStudents  int     `json:"activeStudents"` // practiced in last 7 days
	AvgMastery      float64 `json:"avgMastery"`     // 0-100
	AvgCoverage     float64 `json:"avgCoverage"`    // 0-100
	AvgAccuracy     float64 `json:"avgAccuracy"`    // 0-100
	TotalAttempts   int64   `json:"totalAttempts"`
	WeeklyAttempts  int64   `json:"weeklyAttempts"` // attempts in last 7 days
	AtRiskCount     int     `json:"atRiskCount"`    // students flagged as at-risk
}

// StudentPerformanceSummary is the per-student analytics view used in teacher dashboards
type StudentPerformanceSummary struct {
	UserID          uint    `json:"userId"`
	UserName        string  `json:"userName"`
	MasteryLevel    float64 `json:"masteryLevel"`    // 0-100
	CoverageLevel   float64 `json:"coverageLevel"`   // 0-100
	AccuracyRate    float64 `json:"accuracyRate"`    // 0-100
	TotalAttempts   int64   `json:"totalAttempts"`
	WeeklyAttempts  int64   `json:"weeklyAttempts"`
	ConsecWrong     int     `json:"consecWrong"`     // max consecutive wrong across chapters
	LastActiveAt    *time.Time `json:"lastActiveAt"`
	IsAtRisk        bool    `json:"isAtRisk"`
	RiskReasons     []string `json:"riskReasons,omitempty"`
}

// ClassHeatmapRow is one row in the chapter-vs-student heatmap
type ClassHeatmapRow struct {
	ChapterID   uint    `json:"chapterId"`
	ChapterName string  `json:"chapterName"`
	AvgMastery  float64 `json:"avgMastery"`  // 0-100 average across all students
	StudentData []StudentChapterScore `json:"studentData"`
}

// StudentChapterScore is the mastery score for one student in one chapter
type StudentChapterScore struct {
	UserID       uint    `json:"userId"`
	MasteryLevel float64 `json:"masteryLevel"` // 0-100
	IsCovered    bool    `json:"isCovered"`
}

// ClassHeatmap is the full heatmap response
type ClassHeatmap struct {
	Students []StudentInfo      `json:"students"`
	Chapters []ClassHeatmapRow  `json:"chapters"`
}

// StudentInfo is a minimal student record used in heatmap headers
type StudentInfo struct {
	UserID   uint   `json:"userId"`
	UserName string `json:"userName"`
}

// AttemptTrendPoint is a single data point in a time-series trend
type AttemptTrendPoint struct {
	Date           string `json:"date"`
	TotalAttempts  int    `json:"totalAttempts"`
	CorrectAttempts int   `json:"correctAttempts"`
}

// EarlyWarningStudent is a student flagged by the early warning system
type EarlyWarningStudent struct {
	UserID      uint     `json:"userId"`
	UserName    string   `json:"userName"`
	Reasons     []string `json:"reasons"`
	Severity    string   `json:"severity"` // "low", "medium", "high"
}

// RecommendedQuestion is a chapter/knowledge area recommended for a student to practice.
// Despite the name keeping backward compat, recommendations are at the chapter level.
type RecommendedQuestion struct {
	QuestionID    uint    `json:"questionId"`   // zero when no specific question; kept for future use
	ChapterID     uint    `json:"chapterId"`
	ChapterName   string  `json:"chapterName"`
	Reason        string  `json:"reason"`
	Priority      int     `json:"priority"`  // 1 = highest
	MasteryLevel  float64 `json:"masteryLevel"`
}

// RecommendationResponse is the full recommendation result for a student/goal
type RecommendationResponse struct {
	GoalID       uint                  `json:"goalId"`
	Questions    []RecommendedQuestion `json:"questions"`
	ReviewCount  int                   `json:"reviewCount"`  // chapters due for SRS review
	GapCount     int                   `json:"gapCount"`     // chapters with mastery < threshold
}

// --- Request types ---

// ClassAnalyticsQuery is the request payload for class analytics endpoints
type ClassAnalyticsQuery struct {
	ClassID   uint `json:"classId" binding:"required"`
	GoalID    uint `json:"goalId"`  // optional – filter by a specific goal/syllabus
}

// StudentAnalyticsQuery is the request payload for student-level analytics
type StudentAnalyticsQuery struct {
	GoalID uint `json:"goalId" binding:"required"`
}

// TrendQuery adds a date range to class analytics
type TrendQuery struct {
	ClassID   uint      `json:"classId" binding:"required"`
	GoalID    uint      `json:"goalId"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
