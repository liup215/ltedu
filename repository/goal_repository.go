package repository

import (
	"edu/model"
	"errors"
	"gorm.io/gorm"
)

type IGoalRepository interface {
	Create(goal *model.Goal) error
	Update(goal *model.Goal) error
	Delete(id uint) error
	GetByID(id uint) (*model.Goal, error)
	FindPage(query *model.GoalQuery) ([]model.Goal, int64, error)
	GetByUserAndSyllabus(userId, syllabusId uint) (*model.Goal, error)
	GetActiveGoalsByUser(userId uint) ([]model.Goal, error)
}

type goalRepository struct {
	db *gorm.DB
}

func NewGoalRepository(db *gorm.DB) IGoalRepository {
	return &goalRepository{db: db}
}

func (r *goalRepository) Create(goal *model.Goal) error {
	return r.db.Create(goal).Error
}

func (r *goalRepository) Update(goal *model.Goal) error {
	return r.db.Save(goal).Error
}

func (r *goalRepository) Delete(id uint) error {
	return r.db.Delete(&model.Goal{}, id).Error
}

func (r *goalRepository) GetByID(id uint) (*model.Goal, error) {
	var goal model.Goal
	err := r.db.Where("id = ?", id).
		Preload("User").
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		First(&goal).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &goal, err
}

func (r *goalRepository) FindPage(query *model.GoalQuery) ([]model.Goal, int64, error) {
	var goals []model.Goal
	var total int64

	q := r.db.Model(&model.Goal{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.UserId != 0 {
		q = q.Where("user_id = ?", query.UserId)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.Status != "" {
		q = q.Where("status = ?", query.Status)
	}

	q.Count(&total)

	offset := (query.PageIndex - 1) * query.PageSize
	err := q.Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Order("id DESC").
		Offset(offset).
		Limit(query.PageSize).
		Find(&goals).Error

	return goals, total, err
}

func (r *goalRepository) GetByUserAndSyllabus(userId, syllabusId uint) (*model.Goal, error) {
	var goal model.Goal
	err := r.db.Where("user_id = ? AND syllabus_id = ?", userId, syllabusId).
		Preload("User").
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		First(&goal).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &goal, err
}

func (r *goalRepository) GetActiveGoalsByUser(userId uint) ([]model.Goal, error) {
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
