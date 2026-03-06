package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

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
