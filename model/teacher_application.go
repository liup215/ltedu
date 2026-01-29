package model

import "time"

const (
	TeacherApplyStatusPending  = 1
	TeacherApplyStatusApproved = 2
	TeacherApplyStatusRejected = 3
)

// TeacherApplication represents an application from a user to become a teacher
type TeacherApplication struct {
	Model
	UserID         uint       `json:"userId" gorm:"index"`
	User           *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Motivation     string     `json:"motivation" gorm:"type:text"`
	Experience     string     `json:"experience" gorm:"type:text"`
	Status         int        `json:"status" gorm:"default:1"` // Default to Pending
	AdminNotes     string     `json:"adminNotes,omitempty" gorm:"type:text"`
	AppliedAt      time.Time  `json:"appliedAt"`
	ReviewedAt     *time.Time `json:"reviewedAt,omitempty"`
	ReviewedByID   *uint      `json:"reviewedById,omitempty" gorm:"column:reviewed_by_user_id"`
	ReviewedByUser *User      `json:"reviewedByUser,omitempty" gorm:"foreignKey:ReviewedByID"`
}

// TeacherApplicationQuery represents the query parameters for teacher applications
type TeacherApplicationQuery struct {
	UserID    uint   `json:"userId,omitempty"`
	Status    *int   `json:"status,omitempty"`
	StartDate string `json:"startDate,omitempty"` // Format: YYYY-MM-DD
	EndDate   string `json:"endDate,omitempty"`   // Format: YYYY-MM-DD
	Page
}

// TeacherApplicationCreateRequest represents the request body for creating a teacher application
type TeacherApplicationCreateRequest struct {
	Motivation string `json:"motivation" binding:"required"`
	Experience string `json:"experience" binding:"required"`
}

// TeacherApplicationReviewRequest represents the request body for reviewing a teacher application
type TeacherApplicationReviewRequest struct {
	AdminNotes string `json:"adminNotes"`
}
