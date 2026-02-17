package repository

import (
	"edu/model"
	"time"
	"gorm.io/gorm"
)

type IAttemptRepository interface {
	IRepository[model.Attempt, model.AttemptQuery]
	GetByUserGoalAndChapter(userId, goalId, chapterId uint) ([]model.Attempt, error)
	GetRecentAttempts(userId, goalId uint, limit int) ([]model.Attempt, error)
	GetAttemptStats(userId, goalId uint) (*model.AttemptStatsResponse, error)
	GetAttemptStatsByChapter(userId, goalId, chapterId uint) (*model.AttemptStatsResponse, error)
}

type AttemptRepository struct {
	*Repository[model.Attempt, model.AttemptQuery]
}

func NewAttemptRepository(db *gorm.DB) IAttemptRepository {
	return &AttemptRepository{
		Repository: NewRepository[model.Attempt, model.AttemptQuery](db),
	}
}

func (r *AttemptRepository) ApplyFilters(query *gorm.DB, filter model.AttemptQuery) *gorm.DB {
	if filter.ID != 0 {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.UserId != 0 {
		query = query.Where("user_id = ?", filter.UserId)
	}
	if filter.GoalId != 0 {
		query = query.Where("goal_id = ?", filter.GoalId)
	}
	if filter.TaskId != 0 {
		query = query.Where("task_id = ?", filter.TaskId)
	}
	if filter.QuestionId != 0 {
		query = query.Where("question_id = ?", filter.QuestionId)
	}
	if filter.ChapterId != 0 {
		query = query.Where("chapter_id = ?", filter.ChapterId)
	}
	if !filter.DateFrom.IsZero() {
		query = query.Where("attempted_at >= ?", filter.DateFrom)
	}
	if !filter.DateTo.IsZero() {
		query = query.Where("attempted_at <= ?", filter.DateTo)
	}
	return query
}

func (r *AttemptRepository) GetByUserGoalAndChapter(userId, goalId, chapterId uint) ([]model.Attempt, error) {
	var attempts []model.Attempt
	err := r.db.Where("user_id = ? AND goal_id = ? AND chapter_id = ?", userId, goalId, chapterId).
		Preload("Question").
		Order("attempted_at DESC").
		Find(&attempts).Error
	if err != nil {
		return nil, err
	}
	return attempts, nil
}

func (r *AttemptRepository) GetRecentAttempts(userId, goalId uint, limit int) ([]model.Attempt, error) {
	var attempts []model.Attempt
	err := r.db.Where("user_id = ? AND goal_id = ?", userId, goalId).
		Preload("Question").
		Preload("Chapter").
		Order("attempted_at DESC").
		Limit(limit).
		Find(&attempts).Error
	if err != nil {
		return nil, err
	}
	return attempts, nil
}

func (r *AttemptRepository) GetAttemptStats(userId, goalId uint) (*model.AttemptStatsResponse, error) {
	var stats model.AttemptStatsResponse
	
	// Get all attempts
	var attempts []model.Attempt
	err := r.db.Where("user_id = ? AND goal_id = ?", userId, goalId).Find(&attempts).Error
	if err != nil {
		return nil, err
	}
	
	stats.TotalAttempts = len(attempts)
	if stats.TotalAttempts == 0 {
		return &stats, nil
	}
	
	for _, attempt := range attempts {
		if attempt.IsCorrect != nil && *attempt.IsCorrect {
			stats.CorrectAttempts++
		} else if attempt.IsCorrect != nil {
			stats.WrongAttempts++
		}
		stats.TotalTimeSpent += attempt.TimeSpentSeconds
		stats.TotalScore += attempt.Score
		stats.TotalMaxScore += attempt.MaxScore
	}
	
	if stats.TotalAttempts > 0 {
		stats.AccuracyRate = float64(stats.CorrectAttempts) / float64(stats.TotalAttempts) * 100
		stats.AverageTimePerQ = stats.TotalTimeSpent / stats.TotalAttempts
	}
	
	return &stats, nil
}

func (r *AttemptRepository) GetAttemptStatsByChapter(userId, goalId, chapterId uint) (*model.AttemptStatsResponse, error) {
	var stats model.AttemptStatsResponse
	
	// Get attempts for specific chapter
	var attempts []model.Attempt
	err := r.db.Where("user_id = ? AND goal_id = ? AND chapter_id = ?", userId, goalId, chapterId).
		Find(&attempts).Error
	if err != nil {
		return nil, err
	}
	
	stats.TotalAttempts = len(attempts)
	if stats.TotalAttempts == 0 {
		return &stats, nil
	}
	
	for _, attempt := range attempts {
		if attempt.IsCorrect != nil && *attempt.IsCorrect {
			stats.CorrectAttempts++
		} else if attempt.IsCorrect != nil {
			stats.WrongAttempts++
		}
		stats.TotalTimeSpent += attempt.TimeSpentSeconds
		stats.TotalScore += attempt.Score
		stats.TotalMaxScore += attempt.MaxScore
	}
	
	if stats.TotalAttempts > 0 {
		stats.AccuracyRate = float64(stats.CorrectAttempts) / float64(stats.TotalAttempts) * 100
		stats.AverageTimePerQ = stats.TotalTimeSpent / stats.TotalAttempts
	}
	
	return &stats, nil
}
