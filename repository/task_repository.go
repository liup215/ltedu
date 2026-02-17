package repository

import (
	"edu/model"
	"errors"
	"time"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	Create(task *model.Task) error
	Update(task *model.Task) error
	Delete(id uint) error
	GetByID(id uint) (*model.Task, error)
	FindPage(query *model.TaskQuery) ([]model.Task, int64, error)
	GetByUserAndDate(userId uint, targetDate time.Time) ([]model.Task, error)
	GetByUserGoalAndDateRange(userId, goalId uint, dateFrom, dateTo time.Time) ([]model.Task, error)
	GetTodayTasks(userId, goalId uint) ([]model.Task, error)
	GetUpcomingTasks(userId, goalId uint, days int) ([]model.Task, error)
	GetOverdueTasks(userId, goalId uint) ([]model.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&model.Task{}, id).Error
}

func (r *taskRepository) GetByID(id uint) (*model.Task, error) {
	var task model.Task
	err := r.db.Where("id = ?", id).
		Preload("Chapter").
		Preload("Paper").
		First(&task).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &task, err
}

func (r *taskRepository) FindPage(query *model.TaskQuery) ([]model.Task, int64, error) {
	var tasks []model.Task
	var total int64

	q := r.db.Model(&model.Task{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.UserId != 0 {
		q = q.Where("user_id = ?", query.UserId)
	}
	if query.GoalId != 0 {
		q = q.Where("goal_id = ?", query.GoalId)
	}
	if query.Type != "" {
		q = q.Where("type = ?", query.Type)
	}
	if query.Status != "" {
		q = q.Where("status = ?", query.Status)
	}
	if query.ChapterId != 0 {
		q = q.Where("chapter_id = ?", query.ChapterId)
	}
	if !query.TargetDate.IsZero() {
		q = q.Where("target_date = ?", query.TargetDate)
	}
	if !query.DateFrom.IsZero() {
		q = q.Where("target_date >= ?", query.DateFrom)
	}
	if !query.DateTo.IsZero() {
		q = q.Where("target_date <= ?", query.DateTo)
	}

	q.Count(&total)

	offset := (query.PageIndex - 1) * query.PageSize
	err := q.Preload("Chapter").
		Preload("Paper").
		Order("target_date ASC, priority DESC").
		Offset(offset).
		Limit(query.PageSize).
		Find(&tasks).Error

	return tasks, total, err
}

func (r *taskRepository) GetByUserAndDate(userId uint, targetDate time.Time) ([]model.Task, error) {
	var tasks []model.Task
	dateStr := targetDate.Format("2006-01-02")
	err := r.db.Where("user_id = ? AND DATE(target_date) = ?", userId, dateStr).
		Preload("Chapter").
		Preload("Paper").
		Order("priority DESC, id ASC").
		Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetByUserGoalAndDateRange(userId, goalId uint, dateFrom, dateTo time.Time) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Where("user_id = ? AND goal_id = ? AND target_date >= ? AND target_date <= ?", 
		userId, goalId, dateFrom, dateTo).
		Preload("Chapter").
		Preload("Paper").
		Order("target_date ASC, priority DESC").
		Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTodayTasks(userId, goalId uint) ([]model.Task, error) {
	today := time.Now().Truncate(24 * time.Hour)
	return r.GetByUserAndDate(userId, today)
}

func (r *taskRepository) GetUpcomingTasks(userId, goalId uint, days int) ([]model.Task, error) {
	today := time.Now().Truncate(24 * time.Hour)
	endDate := today.Add(time.Duration(days) * 24 * time.Hour)
	return r.GetByUserGoalAndDateRange(userId, goalId, today.Add(24*time.Hour), endDate)
}

func (r *taskRepository) GetOverdueTasks(userId, goalId uint) ([]model.Task, error) {
	var tasks []model.Task
	today := time.Now().Truncate(24 * time.Hour)
	err := r.db.Where("user_id = ? AND goal_id = ? AND target_date < ? AND status IN (?)", 
		userId, goalId, today, []string{model.TaskStatusPending, model.TaskStatusInProgress}).
		Preload("Chapter").
		Preload("Paper").
		Order("target_date ASC, priority DESC").
		Find(&tasks).Error
	return tasks, err
}
