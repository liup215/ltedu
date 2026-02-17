package model

import "time"

// Attempt represents a single learning activity log (e.g., answering a question)
type Attempt struct {
	Model
	UserId            uint       `json:"userId" gorm:"index:idx_user_goal_question"`
	User              User       `json:"user" gorm:"foreignKey:UserId"`
	GoalId            uint       `json:"goalId" gorm:"index:idx_user_goal_question"`
	Goal              Goal       `json:"goal" gorm:"foreignKey:GoalId"`
	TaskId            *uint      `json:"taskId,omitempty" gorm:"index"`           // Optional link to task
	Task              *Task      `json:"task,omitempty" gorm:"foreignKey:TaskId"`
	QuestionId        uint       `json:"questionId" gorm:"index:idx_user_goal_question"`
	Question          Question   `json:"question" gorm:"foreignKey:QuestionId"`
	ChapterId         uint       `json:"chapterId" gorm:"index"`                  // Denormalized for easier queries
	Chapter           Chapter    `json:"chapter" gorm:"foreignKey:ChapterId"`
	QuestionContentId uint       `json:"questionContentId"`                       // Specific question part
	IsCorrect         *bool      `json:"isCorrect"`                               // nil for subjective questions
	StudentAnswer     string     `json:"studentAnswer" gorm:"type:text"`          // User's answer
	CorrectAnswer     string     `json:"correctAnswer" gorm:"type:text"`          // Correct answer
	TimeSpentSeconds  int        `json:"timeSpentSeconds"`                        // Time spent on this question
	Score             int        `json:"score"`                                   // Points earned
	MaxScore          int        `json:"maxScore"`                                // Maximum possible points
	AttemptedAt       time.Time  `json:"attemptedAt" gorm:"index"`                // When the attempt was made
}

// AttemptQuery for filtering attempts
type AttemptQuery struct {
	ID         uint      `json:"id"`
	UserId     uint      `json:"userId"`
	GoalId     uint      `json:"goalId"`
	TaskId     uint      `json:"taskId"`
	QuestionId uint      `json:"questionId"`
	ChapterId  uint      `json:"chapterId"`
	DateFrom   time.Time `json:"dateFrom"`
	DateTo     time.Time `json:"dateTo"`
	Page
}

// AttemptCreateRequest for creating a new attempt (from practice submission)
type AttemptCreateRequest struct {
	GoalId            uint   `json:"goalId" binding:"required"`
	TaskId            *uint  `json:"taskId,omitempty"`
	QuestionId        uint   `json:"questionId" binding:"required"`
	ChapterId         uint   `json:"chapterId" binding:"required"`
	QuestionContentId uint   `json:"questionContentId" binding:"required"`
	IsCorrect         *bool  `json:"isCorrect"`
	StudentAnswer     string `json:"studentAnswer"`
	CorrectAnswer     string `json:"correctAnswer"`
	TimeSpentSeconds  int    `json:"timeSpentSeconds"`
	Score             int    `json:"score"`
	MaxScore          int    `json:"maxScore"`
}

// AttemptStatsResponse for showing user's attempt statistics
type AttemptStatsResponse struct {
	TotalAttempts    int     `json:"totalAttempts"`
	CorrectAttempts  int     `json:"correctAttempts"`
	WrongAttempts    int     `json:"wrongAttempts"`
	AccuracyRate     float64 `json:"accuracyRate"`
	TotalTimeSpent   int     `json:"totalTimeSpent"`   // In seconds
	AverageTimePerQ  int     `json:"averageTimePerQ"`  // In seconds
	TotalScore       int     `json:"totalScore"`
	TotalMaxScore    int     `json:"totalMaxScore"`
}
