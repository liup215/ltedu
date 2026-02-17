package service

import (
	"edu/model"
	"edu/repository"
	"errors"
)

var TaskLogSvr = &TaskLogService{baseService: newBaseService()}

type TaskLogService struct {
	baseService
}

// CreateTaskLog creates a new task log
func (svr *TaskLogService) CreateTaskLog(userId uint, req model.TaskLogCreateRequest) (*model.TaskLog, error) {
	// Get task
	task, err := repository.TaskRepo.GetByID(req.TaskId)
	if err != nil {
		return nil, err
	}
	if task.UserId != userId {
		return nil, errors.New("unauthorized to create log for this task")
	}

	// Create task log
	log := &model.TaskLog{
		UserId:           userId,
		TaskId:           req.TaskId,
		GoalId:           task.GoalId,
		StartedAt:        req.StartedAt,
		CompletedAt:      req.CompletedAt,
		TimeSpentSeconds: req.TimeSpentSeconds,
		Score:            req.Score,
		MaxScore:         req.MaxScore,
		Status:           req.Status,
		Notes:            req.Notes,
	}

	if err := repository.TaskLogRepo.Create(log); err != nil {
		return nil, err
	}

	// Update task status based on log
	task.Status = req.Status
	task.CompletedAt = &req.CompletedAt
	if err := repository.TaskRepo.Update(task); err != nil {
		return nil, err
	}

	return repository.TaskLogRepo.GetByID(log.ID)
}

// GetTaskLogByID gets a task log by ID
func (svr *TaskLogService) GetTaskLogByID(userId, logId uint) (*model.TaskLog, error) {
	log, err := repository.TaskLogRepo.GetByID(logId)
	if err != nil {
		return nil, err
	}
	if log.UserId != userId {
		return nil, errors.New("unauthorized to access this task log")
	}
	return log, nil
}

// GetTaskLogByTask gets task log for a specific task
func (svr *TaskLogService) GetTaskLogByTask(userId, taskId uint) (*model.TaskLog, error) {
	// Verify user owns the task
	task, err := repository.TaskRepo.GetByID(taskId)
	if err != nil {
		return nil, err
	}
	if task.UserId != userId {
		return nil, errors.New("unauthorized to access this task")
	}

	return repository.TaskLogRepo.GetByTask(taskId)
}

// GetTaskLogsByGoal gets all task logs for a goal
func (svr *TaskLogService) GetTaskLogsByGoal(userId, goalId uint) ([]model.TaskLog, error) {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to access this goal")
	}

	return repository.TaskLogRepo.GetByUserAndGoal(userId, goalId)
}

// ListTaskLogs lists task logs with pagination
func (svr *TaskLogService) ListTaskLogs(userId uint, query model.TaskLogQuery) ([]model.TaskLog, int64, error) {
	query.UserId = userId
	return repository.TaskLogRepo.FindPage(&query)
}
