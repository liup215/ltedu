package model

import "time"

// TaskLog represents the completion log of a task
type TaskLog struct {
	Model
	UserId         uint      `json:"userId" gorm:"index:idx_user_task"`
	User           User      `json:"user" gorm:"foreignKey:UserId"`
	TaskId         uint      `json:"taskId" gorm:"index:idx_user_task"`
	Task           Task      `json:"task" gorm:"foreignKey:TaskId"`
	GoalId         uint      `json:"goalId" gorm:"index"`
	Goal           Goal      `json:"goal" gorm:"foreignKey:GoalId"`
	StartedAt      time.Time `json:"startedAt"`                      // When the user started the task
	CompletedAt    time.Time `json:"completedAt"`                    // When the user completed the task
	TimeSpentSeconds int     `json:"timeSpentSeconds"`               // Total time spent on task
	Score          int       `json:"score"`                          // Score earned (for test/drill tasks)
	MaxScore       int       `json:"maxScore"`                       // Maximum possible score
	Status         string    `json:"status"`                         // completed, failed, skipped
	Notes          string    `json:"notes" gorm:"type:text"`         // Optional notes/feedback
}

// TaskLogQuery for filtering task logs
type TaskLogQuery struct {
	ID         uint      `json:"id"`
	UserId     uint      `json:"userId"`
	TaskId     uint      `json:"taskId"`
	GoalId     uint      `json:"goalId"`
	DateFrom   time.Time `json:"dateFrom"`
	DateTo     time.Time `json:"dateTo"`
	Page
}

// TaskLogCreateRequest for creating a new task log
type TaskLogCreateRequest struct {
	TaskId           uint      `json:"taskId" binding:"required"`
	StartedAt        time.Time `json:"startedAt" binding:"required"`
	CompletedAt      time.Time `json:"completedAt" binding:"required"`
	TimeSpentSeconds int       `json:"timeSpentSeconds"`
	Score            int       `json:"score"`
	MaxScore         int       `json:"maxScore"`
	Status           string    `json:"status" binding:"required"`
	Notes            string    `json:"notes"`
}
