package model

import "time"

// ConversationSession represents an ongoing AI conversation session for a user.
type ConversationSession struct {
	Model
	UserID       uint      `json:"userId" gorm:"index;not null"`
	User         User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	SessionKey   string    `json:"sessionKey" gorm:"uniqueIndex;size:64;not null"`
	UserRole     string    `json:"userRole" gorm:"size:20"`      // "student", "teacher", "admin"
	ContextData  string    `json:"contextData" gorm:"type:text"` // JSON-encoded user context
	Summary      string    `json:"summary" gorm:"type:text"`     // AI-generated summary for long conversations
	MessageCount int       `json:"messageCount" gorm:"default:0"`
	LastActiveAt time.Time `json:"lastActiveAt"`
	ExpiresAt    time.Time `json:"expiresAt" gorm:"index"`
	IsActive     bool      `json:"isActive" gorm:"default:true"`
}

// ConversationMessage stores an individual message within a conversation session.
type ConversationMessage struct {
	Model
	SessionID  uint                `json:"sessionId" gorm:"index;not null"`
	Session    ConversationSession `json:"-" gorm:"foreignKey:SessionID"`
	Role       string              `json:"role" gorm:"size:20;not null"` // "user", "assistant", "system"
	Content    string              `json:"content" gorm:"type:text;not null"`
	OrderIndex int                 `json:"orderIndex" gorm:"index"`
}

// ConversationContext holds structured context about a user's current state.
type ConversationContext struct {
	UserRole        string            `json:"userRole"`
	Preferences     map[string]string `json:"preferences,omitempty"`
	RecentActions   []string          `json:"recentActions,omitempty"`
	CurrentSelection map[string]interface{} `json:"currentSelection,omitempty"`
}

// StartSessionRequest is the payload for starting a new conversation session.
type StartSessionRequest struct {
	Context *ConversationContext `json:"context,omitempty"`
}

// SendMessageRequest is the payload for sending a message in a session.
type SendMessageRequest struct {
	SessionKey string `json:"sessionKey" binding:"required"`
	Message    string `json:"message" binding:"required"`
}

// ConversationHistoryRequest is the payload for retrieving conversation history.
type ConversationHistoryRequest struct {
	SessionKey string `json:"sessionKey" binding:"required"`
}

// ResetContextRequest is the payload for resetting the context of a session.
type ResetContextRequest struct {
	SessionKey string               `json:"sessionKey" binding:"required"`
	Context    *ConversationContext `json:"context,omitempty"`
}

// ConversationSessionResponse is the API response for a conversation session.
type ConversationSessionResponse struct {
	SessionKey   string      `json:"sessionKey"`
	UserRole     string      `json:"userRole"`
	MessageCount int         `json:"messageCount"`
	LastActiveAt time.Time   `json:"lastActiveAt"`
	ExpiresAt    time.Time   `json:"expiresAt"`
	IsActive     bool        `json:"isActive"`
	CreatedAt    time.Time   `json:"createdAt"`
}

// SendMessageResponse is the API response after sending a message.
type SendMessageResponse struct {
	UserMessage      ConversationMessage `json:"userMessage"`
	AssistantMessage ConversationMessage `json:"assistantMessage"`
	SessionKey       string              `json:"sessionKey"`
	MessageCount     int                 `json:"messageCount"`
}

// ConversationRole constants for message roles.
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
