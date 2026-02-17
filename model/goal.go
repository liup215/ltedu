package model

import "time"

// Goal represents a learning project with a target syllabus and exam date
type Goal struct {
	Model
	UserId         uint      `json:"userId" gorm:"index"`
	User           User      `json:"user" gorm:"foreignKey:UserId"`
	SyllabusId     uint      `json:"syllabusId" gorm:"index"`
	Syllabus       Syllabus  `json:"syllabus" gorm:"foreignKey:SyllabusId"`
	ExamDate       time.Time `json:"examDate"`
	TargetScore    int       `json:"targetScore"`           // Optional target score
	TargetGrade    string    `json:"targetGrade"`           // Optional target grade (A*, A, B, etc.)
	WeeklyHours    int       `json:"weeklyHours"`           // Available hours per week
	Mode           string    `json:"mode"`                  // "sync" (school sync) or "self" (self-paced/cram)
	Status         string    `json:"status"`                // "active", "completed", "paused", "cancelled"
	DiagnosticDone bool      `json:"diagnosticDone"`        // Whether initial diagnostic test is completed
	StartDate      time.Time `json:"startDate"`             // When the user started this goal
	CompletedAt    *time.Time `json:"completedAt,omitempty"` // When the goal was completed
}

// GoalQuery for filtering goals
type GoalQuery struct {
	ID         uint   `json:"id"`
	UserId     uint   `json:"userId"`
	SyllabusId uint   `json:"syllabusId"`
	Status     string `json:"status"`
	Page
}

// GoalCreateRequest for creating a new goal
type GoalCreateRequest struct {
	SyllabusId  uint      `json:"syllabusId" binding:"required"`
	ExamDate    time.Time `json:"examDate" binding:"required"`
	TargetScore int       `json:"targetScore"`
	TargetGrade string    `json:"targetGrade"`
	WeeklyHours int       `json:"weeklyHours" binding:"required"`
	Mode        string    `json:"mode" binding:"required"` // "sync" or "self"
}

// GoalUpdateRequest for updating an existing goal
type GoalUpdateRequest struct {
	ID          uint       `json:"id" binding:"required"`
	ExamDate    *time.Time `json:"examDate,omitempty"`
	TargetScore *int       `json:"targetScore,omitempty"`
	TargetGrade *string    `json:"targetGrade,omitempty"`
	WeeklyHours *int       `json:"weeklyHours,omitempty"`
	Mode        *string    `json:"mode,omitempty"`
	Status      *string    `json:"status,omitempty"`
}
