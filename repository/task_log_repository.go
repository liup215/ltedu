package repository

import (
	"edu/model"
	"errors"
	"gorm.io/gorm"
)

type ITaskLogRepository interface {
	Create(log *model.TaskLog) error
	Update(log *model.TaskLog) error
	Delete(id uint) error
	GetByID(id uint) (*model.TaskLog, error)
	FindPage(query *model.TaskLogQuery) ([]model.TaskLog, int64, error)
	GetByTask(taskId uint) (*model.TaskLog, error)
	GetByUserAndGoal(userId, goalId uint) ([]model.TaskLog, error)
}

type taskLogRepository struct {
	db *gorm.DB
}

func NewTaskLogRepository(db *gorm.DB) ITaskLogRepository {
	return &taskLogRepository{db: db}
}

func (r *taskLogRepository) Create(log *model.TaskLog) error {
	return r.db.Create(log).Error
}

func (r *taskLogRepository) Update(log *model.TaskLog) error {
	return r.db.Save(log).Error
}

func (r *taskLogRepository) Delete(id uint) error {
	return r.db.Delete(&model.TaskLog{}, id).Error
}

func (r *taskLogRepository) GetByID(id uint) (*model.TaskLog, error) {
	var log model.TaskLog
	err := r.db.Where("id = ?", id).
		Preload("Task").
		First(&log).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &log, err
}

func (r *taskLogRepository) FindPage(query *model.TaskLogQuery) ([]model.TaskLog, int64, error) {
	var logs []model.TaskLog
	var total int64

	q := r.db.Model(&model.TaskLog{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.UserId != 0 {
		q = q.Where("user_id = ?", query.UserId)
	}
	if query.TaskId != 0 {
		q = q.Where("task_id = ?", query.TaskId)
	}
	if query.GoalId != 0 {
		q = q.Where("goal_id = ?", query.GoalId)
	}
	if !query.DateFrom.IsZero() {
		q = q.Where("completed_at >= ?", query.DateFrom)
	}
	if !query.DateTo.IsZero() {
		q = q.Where("completed_at <= ?", query.DateTo)
	}

	q.Count(&total)

	offset := (query.PageIndex - 1) * query.PageSize
	err := q.Preload("Task").
		Order("completed_at DESC").
		Offset(offset).
		Limit(query.PageSize).
		Find(&logs).Error

	return logs, total, err
}

func (r *taskLogRepository) GetByTask(taskId uint) (*model.TaskLog, error) {
	var log model.TaskLog
	err := r.db.Where("task_id = ?", taskId).
		Preload("Task").
		First(&log).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &log, err
}

func (r *taskLogRepository) GetByUserAndGoal(userId, goalId uint) ([]model.TaskLog, error) {
	var logs []model.TaskLog
	err := r.db.Where("user_id = ? AND goal_id = ?", userId, goalId).
		Preload("Task").
		Order("completed_at DESC").
		Find(&logs).Error
	return logs, err
}
