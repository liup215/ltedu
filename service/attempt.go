package service

import (
	"edu/model"
	"edu/repository"
	"errors"
	"time"
)

var AttemptSvr = &AttemptService{baseService: newBaseService()}

type AttemptService struct {
	baseService
}

// CreateAttempt creates a new learning attempt and updates knowledge state
func (svr *AttemptService) CreateAttempt(userId uint, req model.AttemptCreateRequest) (*model.Attempt, error) {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(req.GoalId)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to create attempt for this goal")
	}

	// Create attempt
	attempt := &model.Attempt{
		UserId:            userId,
		GoalId:            req.GoalId,
		TaskId:            req.TaskId,
		QuestionId:        req.QuestionId,
		ChapterId:         req.ChapterId,
		QuestionContentId: req.QuestionContentId,
		IsCorrect:         req.IsCorrect,
		StudentAnswer:     req.StudentAnswer,
		CorrectAnswer:     req.CorrectAnswer,
		TimeSpentSeconds:  req.TimeSpentSeconds,
		Score:             req.Score,
		MaxScore:          req.MaxScore,
		AttemptedAt:       time.Now(),
	}

	if err := repository.AttemptRepo.Create(attempt); err != nil {
		return nil, err
	}

	// Update knowledge state based on this attempt
	if err := KnowledgeStateSvr.UpdateKnowledgeStateFromAttempt(attempt); err != nil {
		// Log error but don't fail the attempt creation
		// This ensures data consistency even if knowledge state update fails
		return attempt, err
	}

	return repository.AttemptRepo.GetByID(attempt.ID)
}

// GetAttemptByID gets an attempt by ID
func (svr *AttemptService) GetAttemptByID(userId, attemptId uint) (*model.Attempt, error) {
	attempt, err := repository.AttemptRepo.GetByID(attemptId)
	if err != nil {
		return nil, err
	}
	if attempt.UserId != userId {
		return nil, errors.New("unauthorized to access this attempt")
	}
	return attempt, nil
}

// GetRecentAttempts gets recent attempts for a goal
func (svr *AttemptService) GetRecentAttempts(userId, goalId uint, limit int) ([]model.Attempt, error) {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to access this goal")
	}

	return repository.AttemptRepo.GetRecentAttempts(userId, goalId, limit)
}

// GetAttemptStats gets attempt statistics for a goal
func (svr *AttemptService) GetAttemptStats(userId, goalId uint) (*model.AttemptStatsResponse, error) {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to access this goal")
	}

	return repository.AttemptRepo.GetAttemptStats(userId, goalId)
}

// GetAttemptStatsByChapter gets attempt statistics for a specific chapter
func (svr *AttemptService) GetAttemptStatsByChapter(userId, goalId, chapterId uint) (*model.AttemptStatsResponse, error) {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to access this goal")
	}

	return repository.AttemptRepo.GetAttemptStatsByChapter(userId, goalId, chapterId)
}

// ListAttempts lists attempts with pagination
func (svr *AttemptService) ListAttempts(userId uint, query model.AttemptQuery) ([]model.Attempt, int64, error) {
	query.UserId = userId
	return repository.AttemptRepo.FindPage(&query)
}

// BatchCreateAttempts creates multiple attempts from a practice submission
func (svr *AttemptService) BatchCreateAttempts(userId, goalId uint, taskId *uint, submissions []model.AttemptCreateRequest) error {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return err
	}
	if goal.UserId != userId {
		return errors.New("unauthorized to create attempts for this goal")
	}

	for _, req := range submissions {
		req.GoalId = goalId
		req.TaskId = taskId
		_, err := svr.CreateAttempt(userId, req)
		if err != nil {
			// Log error but continue processing other attempts
			// This ensures partial success doesn't fail the entire batch
			continue
		}
	}

	return nil
}
