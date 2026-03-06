package repository

import (
	"edu/model"
	"time"

	"gorm.io/gorm"
)

// IConversationRepository defines data-access operations for AI conversation sessions and messages.
type IConversationRepository interface {
	// Session operations
	CreateSession(session *model.ConversationSession) error
	GetSessionByKey(sessionKey string) (*model.ConversationSession, error)
	GetSessionsByUserID(userID uint) ([]model.ConversationSession, error)
	UpdateSession(session *model.ConversationSession) error
	DeactivateSession(sessionKey string, userID uint) error
	DeleteExpiredSessions() (int64, error)

	// Message operations
	CreateMessage(msg *model.ConversationMessage) error
	GetMessagesBySessionID(sessionID uint) ([]model.ConversationMessage, error)
	DeleteMessagesBySessionID(sessionID uint) error

	// Combined helpers
	GetActiveSessionWithMessages(sessionKey string, userID uint) (*model.ConversationSession, []model.ConversationMessage, error)
}

type conversationRepository struct {
	db *gorm.DB
}

// NewConversationRepository constructs a new conversation repository.
func NewConversationRepository(db *gorm.DB) IConversationRepository {
	return &conversationRepository{db: db}
}

func (r *conversationRepository) CreateSession(session *model.ConversationSession) error {
	return r.db.Create(session).Error
}

func (r *conversationRepository) GetSessionByKey(sessionKey string) (*model.ConversationSession, error) {
	var session model.ConversationSession
	err := r.db.Where("session_key = ?", sessionKey).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *conversationRepository) GetSessionsByUserID(userID uint) ([]model.ConversationSession, error) {
	var sessions []model.ConversationSession
	err := r.db.Where("user_id = ? AND is_active = ?", userID, true).
		Order("last_active_at DESC").
		Find(&sessions).Error
	return sessions, err
}

func (r *conversationRepository) UpdateSession(session *model.ConversationSession) error {
	return r.db.Save(session).Error
}

func (r *conversationRepository) DeactivateSession(sessionKey string, userID uint) error {
	return r.db.Model(&model.ConversationSession{}).
		Where("session_key = ? AND user_id = ?", sessionKey, userID).
		Update("is_active", false).Error
}

func (r *conversationRepository) DeleteExpiredSessions() (int64, error) {
	result := r.db.Where("expires_at < ?", time.Now()).Delete(&model.ConversationSession{})
	return result.RowsAffected, result.Error
}

func (r *conversationRepository) CreateMessage(msg *model.ConversationMessage) error {
	return r.db.Create(msg).Error
}

func (r *conversationRepository) GetMessagesBySessionID(sessionID uint) ([]model.ConversationMessage, error) {
	var messages []model.ConversationMessage
	err := r.db.Where("session_id = ?", sessionID).
		Order("order_index ASC").
		Find(&messages).Error
	return messages, err
}

func (r *conversationRepository) DeleteMessagesBySessionID(sessionID uint) error {
	return r.db.Where("session_id = ?", sessionID).Delete(&model.ConversationMessage{}).Error
}

func (r *conversationRepository) GetActiveSessionWithMessages(sessionKey string, userID uint) (*model.ConversationSession, []model.ConversationMessage, error) {
	session, err := r.GetSessionByKey(sessionKey)
	if err != nil {
		return nil, nil, err
	}

	if session.UserID != userID {
		return nil, nil, gorm.ErrRecordNotFound
	}

	if !session.Active || time.Now().After(session.ExpiresAt) {
		return nil, nil, gorm.ErrRecordNotFound
	}

	messages, err := r.GetMessagesBySessionID(session.ID)
	if err != nil {
		return nil, nil, err
	}

	return session, messages, nil
}

// IConversationSessionRepository manages ConversationSession persistence.
type IConversationSessionRepository interface {
	Create(session *model.ConversationSession) error
	GetByID(id uint) (*model.ConversationSession, error)
	GetByUserID(userID uint) ([]model.ConversationSession, error)
	Update(session *model.ConversationSession) error
}

type conversationSessionRepository struct {
	db *gorm.DB
}

func NewConversationSessionRepository(db *gorm.DB) IConversationSessionRepository {
	return &conversationSessionRepository{db: db}
}

func (r *conversationSessionRepository) Create(session *model.ConversationSession) error {
	return r.db.Create(session).Error
}

func (r *conversationSessionRepository) GetByID(id uint) (*model.ConversationSession, error) {
	var session model.ConversationSession
	err := r.db.First(&session, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &session, err
}

func (r *conversationSessionRepository) GetByUserID(userID uint) ([]model.ConversationSession, error) {
	var sessions []model.ConversationSession
	err := r.db.Where("user_id = ? AND active = ?", userID, true).
		Order("created_at DESC").Find(&sessions).Error
	return sessions, err
}

func (r *conversationSessionRepository) Update(session *model.ConversationSession) error {
	return r.db.Save(session).Error
}

// IConversationMessageRepository manages ConversationMessage persistence.
type IConversationMessageRepository interface {
	Create(msg *model.ConversationMessage) error
	GetBySessionID(sessionID uint) ([]model.ConversationMessage, error)
	DeleteBySessionID(sessionID uint) error
}

type conversationMessageRepository struct {
	db *gorm.DB
}

func NewConversationMessageRepository(db *gorm.DB) IConversationMessageRepository {
	return &conversationMessageRepository{db: db}
}

func (r *conversationMessageRepository) Create(msg *model.ConversationMessage) error {
	return r.db.Create(msg).Error
}

func (r *conversationMessageRepository) GetBySessionID(sessionID uint) ([]model.ConversationMessage, error) {
	var msgs []model.ConversationMessage
	err := r.db.Where("session_id = ?", sessionID).Order("created_at ASC").Find(&msgs).Error
	return msgs, err
}

func (r *conversationMessageRepository) DeleteBySessionID(sessionID uint) error {
	return r.db.Where("session_id = ?", sessionID).Delete(&model.ConversationMessage{}).Error
}
