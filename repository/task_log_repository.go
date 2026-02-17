package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

type ITaskLogRepository interface {
	IRepository[model.TaskLog, model.TaskLogQuery]
	GetByTask(taskId uint) (*model.TaskLog, error)
	GetByUserAndGoal(userId, goalId uint) ([]model.TaskLog, error)
}

type TaskLogRepository struct {
	*Repository[model.TaskLog, model.TaskLogQuery]
}

func NewTaskLogRepository(db *gorm.DB) ITaskLogRepository {
	return &TaskLogRepository{
		Repository: NewRepository[model.TaskLog, model.TaskLogQuery](db),
	}
}

func (r *TaskLogRepository) ApplyFilters(query *gorm.DB, filter model.TaskLogQuery) *gorm.DB {
	if filter.ID != 0 {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.UserId != 0 {
		query = query.Where("user_id = ?", filter.UserId)
	}
	if filter.TaskId != 0 {
		query = query.Where("task_id = ?", filter.TaskId)
	}
	if filter.GoalId != 0 {
		query = query.Where("goal_id = ?", filter.GoalId)
	}
	if !filter.DateFrom.IsZero() {
		query = query.Where("completed_at >= ?", filter.DateFrom)
	}
	if !filter.DateTo.IsZero() {
		query = query.Where("completed_at <= ?", filter.DateTo)
	}
	return query
}

func (r *TaskLogRepository) GetByTask(taskId uint) (*model.TaskLog, error) {
	var log model.TaskLog
	err := r.db.Where("task_id = ?", taskId).
		Preload("Task").
		First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *TaskLogRepository) GetByUserAndGoal(userId, goalId uint) ([]model.TaskLog, error) {
	var logs []model.TaskLog
	err := r.db.Where("user_id = ? AND goal_id = ?", userId, goalId).
		Preload("Task").
		Order("completed_at DESC").
		Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}
