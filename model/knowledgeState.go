package model

import "time"

// KnowledgeState represents user's mastery and retention for a chapter/knowledge point
type KnowledgeState struct {
	Model
	UserId            uint       `json:"userId" gorm:"index:idx_user_goal_chapter"`
	User              User       `json:"user" gorm:"foreignKey:UserId"`
	GoalId            uint       `json:"goalId" gorm:"index:idx_user_goal_chapter"`
	Goal              Goal       `json:"goal" gorm:"foreignKey:GoalId"`
	ChapterId         uint       `json:"chapterId" gorm:"index:idx_user_goal_chapter"`
	Chapter           Chapter    `json:"chapter" gorm:"foreignKey:ChapterId"`
	MasteryLevel      float64    `json:"masteryLevel"`      // 0.0 to 1.0, represents how well the user knows this
	StabilityScore    float64    `json:"stabilityScore"`    // 0.0 to 1.0, represents memory retention
	CorrectCount      int        `json:"correctCount"`      // Number of correct attempts
	TotalCount        int        `json:"totalCount"`        // Total number of attempts
	LastPracticeAt    *time.Time `json:"lastPracticeAt"`    // Last practice time
	NextReviewAt      *time.Time `json:"nextReviewAt"`      // When to review next (SRS)
	ConsecutiveCorrect int        `json:"consecutiveCorrect"` // Consecutive correct answers
	ConsecutiveWrong  int        `json:"consecutiveWrong"`  // Consecutive wrong answers
	IsCovered         bool       `json:"isCovered"`         // Whether this chapter has been covered/touched
}

// KnowledgeStateQuery for filtering knowledge states
type KnowledgeStateQuery struct {
	ID        uint `json:"id"`
	UserId    uint `json:"userId"`
	GoalId    uint `json:"goalId"`
	ChapterId uint `json:"chapterId"`
	Page
}

// KnowledgeStateUpdateRequest for updating knowledge state (usually internal use)
type KnowledgeStateUpdateRequest struct {
	ID                 uint       `json:"id" binding:"required"`
	MasteryLevel       *float64   `json:"masteryLevel,omitempty"`
	StabilityScore     *float64   `json:"stabilityScore,omitempty"`
	CorrectCount       *int       `json:"correctCount,omitempty"`
	TotalCount         *int       `json:"totalCount,omitempty"`
	LastPracticeAt     *time.Time `json:"lastPracticeAt,omitempty"`
	NextReviewAt       *time.Time `json:"nextReviewAt,omitempty"`
	ConsecutiveCorrect *int       `json:"consecutiveCorrect,omitempty"`
	ConsecutiveWrong   *int       `json:"consecutiveWrong,omitempty"`
	IsCovered          *bool      `json:"isCovered,omitempty"`
}

// KnowledgeStateProgressResponse for showing user progress
type KnowledgeStateProgressResponse struct {
	Coverage   float64                       `json:"coverage"`   // Percentage of chapters covered
	Mastery    float64                       `json:"mastery"`    // Average mastery level
	Stability  float64                       `json:"stability"`  // Average stability score
	ChapterStates []KnowledgeStateChapterInfo `json:"chapterStates"` // Individual chapter states
}

// KnowledgeStateChapterInfo for individual chapter state in progress response
type KnowledgeStateChapterInfo struct {
	ChapterId      uint    `json:"chapterId"`
	ChapterName    string  `json:"chapterName"`
	MasteryLevel   float64 `json:"masteryLevel"`
	StabilityScore float64 `json:"stabilityScore"`
	IsCovered      bool    `json:"isCovered"`
}
