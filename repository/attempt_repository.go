package repository

import (
	"edu/model"
	"errors"
	"gorm.io/gorm"
)

type IAttemptRepository interface {
	Create(attempt *model.Attempt) error
	Update(attempt *model.Attempt) error
	Delete(id uint) error
	GetByID(id uint) (*model.Attempt, error)
	FindPage(query *model.AttemptQuery) ([]model.Attempt, int64, error)
	GetByUserGoalAndChapter(userId, goalId, chapterId uint) ([]model.Attempt, error)
	GetRecentAttempts(userId, goalId uint, limit int) ([]model.Attempt, error)
	GetAttemptStats(userId, goalId uint) (*model.AttemptStatsResponse, error)
	GetAttemptStatsByChapter(userId, goalId, chapterId uint) (*model.AttemptStatsResponse, error)
}

type attemptRepository struct {
	db *gorm.DB
}

func NewAttemptRepository(db *gorm.DB) IAttemptRepository {
	return &attemptRepository{db: db}
}

func (r *attemptRepository) Create(attempt *model.Attempt) error {
	return r.db.Create(attempt).Error
}

func (r *attemptRepository) Update(attempt *model.Attempt) error {
	return r.db.Save(attempt).Error
}

func (r *attemptRepository) Delete(id uint) error {
	return r.db.Delete(&model.Attempt{}, id).Error
}

func (r *attemptRepository) GetByID(id uint) (*model.Attempt, error) {
	var attempt model.Attempt
	err := r.db.Where("id = ?", id).
		Preload("Question").
		Preload("Chapter").
		First(&attempt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &attempt, err
}

func (r *attemptRepository) FindPage(query *model.AttemptQuery) ([]model.Attempt, int64, error) {
	var attempts []model.Attempt
	var total int64

	q := r.db.Model(&model.Attempt{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.UserId != 0 {
		q = q.Where("user_id = ?", query.UserId)
	}
	if query.GoalId != 0 {
		q = q.Where("goal_id = ?", query.GoalId)
	}
	if query.TaskId != 0 {
		q = q.Where("task_id = ?", query.TaskId)
	}
	if query.QuestionId != 0 {
		q = q.Where("question_id = ?", query.QuestionId)
	}
	if query.ChapterId != 0 {
		q = q.Where("chapter_id = ?", query.ChapterId)
	}
	if !query.DateFrom.IsZero() {
		q = q.Where("attempted_at >= ?", query.DateFrom)
	}
	if !query.DateTo.IsZero() {
		q = q.Where("attempted_at <= ?", query.DateTo)
	}

	q.Count(&total)

	offset := (query.PageIndex - 1) * query.PageSize
	err := q.Preload("Question").
		Preload("Chapter").
		Order("attempted_at DESC").
		Offset(offset).
		Limit(query.PageSize).
		Find(&attempts).Error

	return attempts, total, err
}

func (r *attemptRepository) GetByUserGoalAndChapter(userId, goalId, chapterId uint) ([]model.Attempt, error) {
	var attempts []model.Attempt
	err := r.db.Where("user_id = ? AND goal_id = ? AND chapter_id = ?", userId, goalId, chapterId).
		Preload("Question").
		Order("attempted_at DESC").
		Find(&attempts).Error
	return attempts, err
}

func (r *attemptRepository) GetRecentAttempts(userId, goalId uint, limit int) ([]model.Attempt, error) {
	var attempts []model.Attempt
	err := r.db.Where("user_id = ? AND goal_id = ?", userId, goalId).
		Preload("Question").
		Preload("Chapter").
		Order("attempted_at DESC").
		Limit(limit).
		Find(&attempts).Error
	return attempts, err
}

func (r *attemptRepository) GetAttemptStats(userId, goalId uint) (*model.AttemptStatsResponse, error) {
	var stats model.AttemptStatsResponse
	
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

func (r *attemptRepository) GetAttemptStatsByChapter(userId, goalId, chapterId uint) (*model.AttemptStatsResponse, error) {
	var stats model.AttemptStatsResponse
	
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
