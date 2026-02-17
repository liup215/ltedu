package repository

import (
	"edu/model"
	"time"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	IRepository[model.Task, model.TaskQuery]
	GetByUserAndDate(userId uint, targetDate time.Time) ([]model.Task, error)
	GetByUserGoalAndDateRange(userId, goalId uint, dateFrom, dateTo time.Time) ([]model.Task, error)
	GetTodayTasks(userId, goalId uint) ([]model.Task, error)
	GetUpcomingTasks(userId, goalId uint, days int) ([]model.Task, error)
	GetOverdueTasks(userId, goalId uint) ([]model.Task, error)
	GetPendingTasksByGoal(goalId uint) ([]model.Task, error)
}

type TaskRepository struct {
	*Repository[model.Task, model.TaskQuery]
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{
		Repository: NewRepository[model.Task, model.TaskQuery](db),
	}
}

func (r *TaskRepository) ApplyFilters(query *gorm.DB, filter model.TaskQuery) *gorm.DB {
	if filter.ID != 0 {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.UserId != 0 {
		query = query.Where("user_id = ?", filter.UserId)
	}
	if filter.GoalId != 0 {
		query = query.Where("goal_id = ?", filter.GoalId)
	}
	if filter.Type != "" {
		query = query.Where("type = ?", filter.Type)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.ChapterId != 0 {
		query = query.Where("chapter_id = ?", filter.ChapterId)
	}
	if !filter.TargetDate.IsZero() {
		query = query.Where("target_date = ?", filter.TargetDate)
	}
	if !filter.DateFrom.IsZero() {
		query = query.Where("target_date >= ?", filter.DateFrom)
	}
	if !filter.DateTo.IsZero() {
		query = query.Where("target_date <= ?", filter.DateTo)
	}
	return query
}

func (r *TaskRepository) GetByUserAndDate(userId uint, targetDate time.Time) ([]model.Task, error) {
	var tasks []model.Task
	dateStr := targetDate.Format("2006-01-02")
	err := r.db.Where("user_id = ? AND DATE(target_date) = ?", userId, dateStr).
		Preload("Chapter").
		Preload("Paper").
		Order("priority DESC, id ASC").
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetByUserGoalAndDateRange(userId, goalId uint, dateFrom, dateTo time.Time) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Where("user_id = ? AND goal_id = ? AND target_date >= ? AND target_date <= ?", 
		userId, goalId, dateFrom, dateTo).
		Preload("Chapter").
		Preload("Paper").
		Order("target_date ASC, priority DESC").
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetTodayTasks(userId, goalId uint) ([]model.Task, error) {
	today := time.Now().Truncate(24 * time.Hour)
	return r.GetByUserAndDate(userId, today)
}

func (r *TaskRepository) GetUpcomingTasks(userId, goalId uint, days int) ([]model.Task, error) {
	today := time.Now().Truncate(24 * time.Hour)
	endDate := today.Add(time.Duration(days) * 24 * time.Hour)
	return r.GetByUserGoalAndDateRange(userId, goalId, today.Add(24*time.Hour), endDate)
}

func (r *TaskRepository) GetOverdueTasks(userId, goalId uint) ([]model.Task, error) {
	var tasks []model.Task
	today := time.Now().Truncate(24 * time.Hour)
	err := r.db.Where("user_id = ? AND goal_id = ? AND target_date < ? AND status IN (?)", 
		userId, goalId, today, []string{model.TaskStatusPending, model.TaskStatusInProgress}).
		Preload("Chapter").
		Preload("Paper").
		Order("target_date ASC, priority DESC").
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetPendingTasksByGoal(goalId uint) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Where("goal_id = ? AND status = ?", goalId, model.TaskStatusPending).
		Preload("Chapter").
		Preload("Paper").
		Order("target_date ASC, priority DESC").
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
