package model

import "time"

// MigrationJob status constants
const (
	MigrationJobStatusPending   = "pending"
	MigrationJobStatusRunning   = "running"
	MigrationJobStatusCompleted = "completed"
	MigrationJobStatusFailed    = "failed"
)

// MigrationJob represents a background knowledge-point migration task
type MigrationJob struct {
	Model
	SyllabusId          uint       `json:"syllabusId" gorm:"index"`
	Status              string     `json:"status" gorm:"default:pending"`
	Options             string     `json:"options" gorm:"type:text"`             // JSON-encoded MigrateOptions
	Progress            int        `json:"progress" gorm:"default:0"`            // 0-100
	TotalItems          int        `json:"totalItems" gorm:"default:0"`          // total leaf chapters (or questions)
	DoneItems           int        `json:"doneItems" gorm:"default:0"`           // processed so far
	Report              string     `json:"report" gorm:"type:text"`              // JSON-encoded MigrateReport
	ErrorMessage        string     `json:"errorMessage" gorm:"type:text"`
	ProcessedChapterIds string     `json:"processedChapterIds" gorm:"type:text"` // JSON array of successfully processed chapter IDs
	CreatedBy           uint       `json:"createdBy" gorm:"index"`
	StartedAt           *time.Time `json:"startedAt,omitempty"`
	CompletedAt         *time.Time `json:"completedAt,omitempty"`
}

// MigrationJobQuery for list/filter
type MigrationJobQuery struct {
	SyllabusId uint   `json:"syllabusId"`
	Status     string `json:"status"`
	CreatedBy  uint   `json:"createdBy"`
	Page
}
