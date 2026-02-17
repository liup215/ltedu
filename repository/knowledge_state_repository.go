package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

type IKnowledgeStateRepository interface {
	IRepository[model.KnowledgeState, model.KnowledgeStateQuery]
	GetByUserGoalChapter(userId, goalId, chapterId uint) (*model.KnowledgeState, error)
	GetByUserAndGoal(userId, goalId uint) ([]model.KnowledgeState, error)
	GetDueForReview(userId, goalId uint) ([]model.KnowledgeState, error)
	InitializeForGoal(userId, goalId, syllabusId uint) error
}

type KnowledgeStateRepository struct {
	*Repository[model.KnowledgeState, model.KnowledgeStateQuery]
	chapterRepo IChapterRepository
}

func NewKnowledgeStateRepository(db *gorm.DB, chapterRepo IChapterRepository) IKnowledgeStateRepository {
	return &KnowledgeStateRepository{
		Repository:  NewRepository[model.KnowledgeState, model.KnowledgeStateQuery](db),
		chapterRepo: chapterRepo,
	}
}

func (r *KnowledgeStateRepository) ApplyFilters(query *gorm.DB, filter model.KnowledgeStateQuery) *gorm.DB {
	if filter.ID != 0 {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.UserId != 0 {
		query = query.Where("user_id = ?", filter.UserId)
	}
	if filter.GoalId != 0 {
		query = query.Where("goal_id = ?", filter.GoalId)
	}
	if filter.ChapterId != 0 {
		query = query.Where("chapter_id = ?", filter.ChapterId)
	}
	return query
}

func (r *KnowledgeStateRepository) GetByUserGoalChapter(userId, goalId, chapterId uint) (*model.KnowledgeState, error) {
	var ks model.KnowledgeState
	err := r.db.Where("user_id = ? AND goal_id = ? AND chapter_id = ?", userId, goalId, chapterId).
		Preload("Chapter").
		First(&ks).Error
	if err != nil {
		return nil, err
	}
	return &ks, nil
}

func (r *KnowledgeStateRepository) GetByUserAndGoal(userId, goalId uint) ([]model.KnowledgeState, error) {
	var states []model.KnowledgeState
	err := r.db.Where("user_id = ? AND goal_id = ?", userId, goalId).
		Preload("Chapter").
		Order("chapter_id").
		Find(&states).Error
	if err != nil {
		return nil, err
	}
	return states, nil
}

func (r *KnowledgeStateRepository) GetDueForReview(userId, goalId uint) ([]model.KnowledgeState, error) {
	var states []model.KnowledgeState
	err := r.db.Where("user_id = ? AND goal_id = ? AND next_review_at IS NOT NULL AND next_review_at <= NOW()", userId, goalId).
		Preload("Chapter").
		Order("next_review_at").
		Find(&states).Error
	if err != nil {
		return nil, err
	}
	return states, nil
}

func (r *KnowledgeStateRepository) InitializeForGoal(userId, goalId, syllabusId uint) error {
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
		
		if err == gorm.ErrRecordNotFound {
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
