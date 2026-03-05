package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

// INLUFeedbackRepository manages NLU feedback persistence.
type INLUFeedbackRepository interface {
	Create(feedback *model.NLUFeedback) error
	ListByUserID(userID uint) ([]model.NLUFeedback, error)
}

type nluFeedbackRepository struct {
	db *gorm.DB
}

func NewNLUFeedbackRepository(db *gorm.DB) INLUFeedbackRepository {
	return &nluFeedbackRepository{db: db}
}

func (r *nluFeedbackRepository) Create(feedback *model.NLUFeedback) error {
	return r.db.Create(feedback).Error
}

func (r *nluFeedbackRepository) ListByUserID(userID uint) ([]model.NLUFeedback, error) {
	var feedbacks []model.NLUFeedback
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&feedbacks).Error
	return feedbacks, err
}
