package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

type IGoalRepository interface {
	IRepository[model.Goal, model.GoalQuery]
	GetByUserAndSyllabus(userId, syllabusId uint) (*model.Goal, error)
	GetActiveGoalsByUser(userId uint) ([]model.Goal, error)
}

type GoalRepository struct {
	*Repository[model.Goal, model.GoalQuery]
}

func NewGoalRepository(db *gorm.DB) IGoalRepository {
	return &GoalRepository{
		Repository: NewRepository[model.Goal, model.GoalQuery](db),
	}
}

func (r *GoalRepository) ApplyFilters(query *gorm.DB, filter model.GoalQuery) *gorm.DB {
	if filter.ID != 0 {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.UserId != 0 {
		query = query.Where("user_id = ?", filter.UserId)
	}
	if filter.SyllabusId != 0 {
		query = query.Where("syllabus_id = ?", filter.SyllabusId)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	return query
}

func (r *GoalRepository) GetByUserAndSyllabus(userId, syllabusId uint) (*model.Goal, error) {
	var goal model.Goal
	err := r.db.Where("user_id = ? AND syllabus_id = ?", userId, syllabusId).
		Preload("User").
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		First(&goal).Error
	if err != nil {
		return nil, err
	}
	return &goal, nil
}

func (r *GoalRepository) GetActiveGoalsByUser(userId uint) ([]model.Goal, error) {
	var goals []model.Goal
	err := r.db.Where("user_id = ? AND status = ?", userId, "active").
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Find(&goals).Error
	if err != nil {
		return nil, err
	}
	return goals, nil
}
