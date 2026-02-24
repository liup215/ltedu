package repository

import (
	"edu/model"
	"errors"
	"time"
	"gorm.io/gorm"
)

type IKnowledgeStateRepository interface {
	Create(ks *model.KnowledgeState) error
	Update(ks *model.KnowledgeState) error
	Delete(id uint) error
	GetByID(id uint) (*model.KnowledgeState, error)
	FindPage(query *model.KnowledgeStateQuery) ([]model.KnowledgeState, int64, error)
	GetByUserGoalChapter(userId, goalId, chapterId uint) (*model.KnowledgeState, error)
	GetByUserAndGoal(userId, goalId uint) ([]model.KnowledgeState, error)
	GetDueForReview(userId, goalId uint) ([]model.KnowledgeState, error)
	InitializeForGoal(userId, goalId, syllabusId uint) error
}

type knowledgeStateRepository struct {
	db          *gorm.DB
	chapterRepo IChapterRepository
}

func NewKnowledgeStateRepository(db *gorm.DB, chapterRepo IChapterRepository) IKnowledgeStateRepository {
	return &knowledgeStateRepository{
		db:          db,
		chapterRepo: chapterRepo,
	}
}

func (r *knowledgeStateRepository) Create(ks *model.KnowledgeState) error {
	return r.db.Create(ks).Error
}

func (r *knowledgeStateRepository) Update(ks *model.KnowledgeState) error {
	return r.db.Save(ks).Error
}

func (r *knowledgeStateRepository) Delete(id uint) error {
	return r.db.Delete(&model.KnowledgeState{}, id).Error
}

func (r *knowledgeStateRepository) GetByID(id uint) (*model.KnowledgeState, error) {
	var ks model.KnowledgeState
	err := r.db.Where("id = ?", id).
		Preload("Chapter").
		First(&ks).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ks, err
}

func (r *knowledgeStateRepository) FindPage(query *model.KnowledgeStateQuery) ([]model.KnowledgeState, int64, error) {
	var states []model.KnowledgeState
	var total int64

	q := r.db.Model(&model.KnowledgeState{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.UserId != 0 {
		q = q.Where("user_id = ?", query.UserId)
	}
	if query.GoalId != 0 {
		q = q.Where("goal_id = ?", query.GoalId)
	}
	if query.ChapterId != 0 {
		q = q.Where("chapter_id = ?", query.ChapterId)
	}

	q.Count(&total)

	offset := (query.PageIndex - 1) * query.PageSize
	err := q.Preload("Chapter").
		Order("chapter_id").
		Offset(offset).
		Limit(query.PageSize).
		Find(&states).Error

	return states, total, err
}

func (r *knowledgeStateRepository) GetByUserGoalChapter(userId, goalId, chapterId uint) (*model.KnowledgeState, error) {
	var ks model.KnowledgeState
	err := r.db.Where("user_id = ? AND goal_id = ? AND chapter_id = ?", userId, goalId, chapterId).
		Preload("Chapter").
		First(&ks).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ks, err
}

func (r *knowledgeStateRepository) GetByUserAndGoal(userId, goalId uint) ([]model.KnowledgeState, error) {
	var states []model.KnowledgeState
	err := r.db.Where("user_id = ? AND goal_id = ?", userId, goalId).
		Preload("Chapter").
		Order("chapter_id").
		Find(&states).Error
	return states, err
}

func (r *knowledgeStateRepository) GetDueForReview(userId, goalId uint) ([]model.KnowledgeState, error) {
	var states []model.KnowledgeState
	now := time.Now()
	err := r.db.Where("user_id = ? AND goal_id = ? AND next_review_at IS NOT NULL AND next_review_at <= ?", userId, goalId, now).
		Preload("Chapter").
		Order("next_review_at").
		Find(&states).Error
	return states, err
}

func (r *knowledgeStateRepository) InitializeForGoal(userId, goalId, syllabusId uint) error {
	// Get all chapters for the syllabus
	chapters, err := r.chapterRepo.FindBySyllabusID(syllabusId)
	if err != nil {
		return err
	}

	// Create initial knowledge states for all chapters
	for _, chapter := range chapters {
		// Check if knowledge state already exists
		var existingState model.KnowledgeState
		err := r.db.Where("user_id = ? AND goal_id = ? AND chapter_id = ?", userId, goalId, chapter.ID).
			First(&existingState).Error
		
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create new knowledge state
			newState := model.KnowledgeState{
				UserId:             userId,
				GoalId:             goalId,
				ChapterId:          chapter.ID,
				MasteryLevel:       0.0,
				StabilityScore:     0.0,
				CorrectCount:       0,
				TotalCount:         0,
				ConsecutiveCorrect: 0,
				ConsecutiveWrong:   0,
				IsCovered:          false,
			}
			if err := r.db.Create(&newState).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
