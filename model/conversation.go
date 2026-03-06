package model

import "time"

// ConvRole constants for message roles.
const (
	ConvRoleUser      = "user"
	ConvRoleAssistant = "assistant"
	ConvRoleSystem    = "system"
)

// ConversationSession holds metadata for an ongoing AI conversation.
type ConversationSession struct {
	Model
	UserID    uint      `json:"userId"   gorm:"column:user_id;index"`
	Title     string    `json:"title"    gorm:"type:varchar(255)"`
	Subject   string    `json:"subject"  gorm:"type:varchar(100)"` // educational subject context
	Active    bool      `json:"active"   gorm:"default:true"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// ConversationMessage is a single message in a conversation session.
type ConversationMessage struct {
	Model
	SessionID uint   `json:"sessionId" gorm:"column:session_id;index"`
	Role      string `json:"role"      gorm:"type:varchar(50)"`
	Content   string `json:"content"   gorm:"type:text"`
}

// ConversationStartRequest is the API payload to begin a new session.
type ConversationStartRequest struct {
	Subject string `json:"subject"` // optional subject context (e.g. "Biology")
	Title   string `json:"title"`
}

// ConversationMessageRequest is the API payload to send a message.
type ConversationMessageRequest struct {
	SessionID uint   `json:"sessionId"`
	Message   string `json:"message"`
}

// ConversationHistoryRequest is the API payload to fetch message history.
type ConversationHistoryRequest struct {
	SessionID uint `json:"sessionId"`
}

// ConversationResetRequest is the API payload to clear a session's history.
type ConversationResetRequest struct {
	SessionID uint `json:"sessionId"`
}

// ConversationCloseRequest is the API payload to deactivate a session.
type ConversationCloseRequest struct {
	SessionID uint `json:"sessionId"`
}

// ConversationSessionsRequest is the API payload to list user sessions.
type ConversationSessionsRequest struct {
	UserID uint `json:"userId"`
}
