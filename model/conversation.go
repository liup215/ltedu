package model

import "time"

// ConversationSession represents a multi-turn AI chat session
type ConversationSession struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	UserId    uint   `gorm:"not null;index" json:"userId"`
	Title     string `gorm:"size:255" json:"title"`
	IsActive  bool   `gorm:"default:true" json:"isActive"`
	MessageCount int `gorm:"default:0" json:"messageCount"`
}

// ConversationMessage represents a single message in a conversation
type ConversationMessage struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	SessionId uint   `gorm:"not null;index" json:"sessionId"`
	Role      string `gorm:"size:16;not null" json:"role"` // "user" or "assistant"
	Content   string `gorm:"type:text;not null" json:"content"`
}

// ConversationStartRequest is used to start a new conversation session
type ConversationStartRequest struct {
	Title string `json:"title"`
}

// ConversationMessageRequest is used to send a message
type ConversationMessageRequest struct {
	SessionId uint   `json:"sessionId" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

// ConversationHistoryRequest is used to fetch message history
type ConversationHistoryRequest struct {
	SessionId uint `json:"sessionId" binding:"required"`
	Page      Page `json:"page"`
}

// ConversationSessionsRequest is used to list sessions
type ConversationSessionsRequest struct {
	Page Page `json:"page"`
}
