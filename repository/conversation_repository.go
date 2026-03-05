package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

type IConversationSessionRepository interface {
	Create(session *model.ConversationSession) error
	Update(session *model.ConversationSession) error
	GetByID(id uint) (*model.ConversationSession, error)
	ListByUser(userId uint, page, pageSize int) ([]model.ConversationSession, int64, error)
	DeleteByID(id uint) error
}

type IConversationMessageRepository interface {
	Create(msg *model.ConversationMessage) error
	ListBySession(sessionId uint, page, pageSize int) ([]model.ConversationMessage, int64, error)
	AllBySession(sessionId uint) ([]model.ConversationMessage, error)
	DeleteBySession(sessionId uint) error
}

// --- Session repository ---

type conversationSessionRepository struct {
	db *gorm.DB
}

func NewConversationSessionRepository(db *gorm.DB) IConversationSessionRepository {
	return &conversationSessionRepository{db: db}
}

func (r *conversationSessionRepository) Create(session *model.ConversationSession) error {
	return r.db.Create(session).Error
}

func (r *conversationSessionRepository) Update(session *model.ConversationSession) error {
	return r.db.Save(session).Error
}

func (r *conversationSessionRepository) GetByID(id uint) (*model.ConversationSession, error) {
	var session model.ConversationSession
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *conversationSessionRepository) ListByUser(userId uint, page, pageSize int) ([]model.ConversationSession, int64, error) {
	var sessions []model.ConversationSession
	var total int64

	query := r.db.Model(&model.ConversationSession{}).
		Where("user_id = ? AND deleted_at IS NULL", userId)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.Order("updated_at DESC").
		Offset(offset).Limit(pageSize).
		Find(&sessions).Error
	return sessions, total, err
}

func (r *conversationSessionRepository) DeleteByID(id uint) error {
	return r.db.Delete(&model.ConversationSession{}, id).Error
}

// --- Message repository ---

type conversationMessageRepository struct {
	db *gorm.DB
}

func NewConversationMessageRepository(db *gorm.DB) IConversationMessageRepository {
	return &conversationMessageRepository{db: db}
}

func (r *conversationMessageRepository) Create(msg *model.ConversationMessage) error {
	return r.db.Create(msg).Error
}

func (r *conversationMessageRepository) ListBySession(sessionId uint, page, pageSize int) ([]model.ConversationMessage, int64, error) {
	var msgs []model.ConversationMessage
	var total int64

	query := r.db.Model(&model.ConversationMessage{}).
		Where("session_id = ? AND deleted_at IS NULL", sessionId)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.Order("created_at ASC").
		Offset(offset).Limit(pageSize).
		Find(&msgs).Error
	return msgs, total, err
}

func (r *conversationMessageRepository) AllBySession(sessionId uint) ([]model.ConversationMessage, error) {
	var msgs []model.ConversationMessage
	err := r.db.Where("session_id = ? AND deleted_at IS NULL", sessionId).
		Order("created_at ASC").
		Find(&msgs).Error
	return msgs, err
}

func (r *conversationMessageRepository) DeleteBySession(sessionId uint) error {
	return r.db.Where("session_id = ?", sessionId).Delete(&model.ConversationMessage{}).Error
}
