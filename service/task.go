package service

import (
	"edu/model"
	"edu/repository"
	"errors"
	"fmt"
	"time"
)

var TaskSvr = &TaskService{baseService: newBaseService()}

type TaskService struct {
	baseService
}

// CreateTask creates a new task
func (svr *TaskService) CreateTask(userId uint, req model.TaskCreateRequest) (*model.Task, error) {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(req.GoalId)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to create task for this goal")
	}

	// Validate task type
	validTypes := map[string]bool{
		model.TaskTypeLearn:  true,
		model.TaskTypeDrill:  true,
		model.TaskTypeReview: true,
		model.TaskTypeTest:   true,
		model.TaskTypeMock:   true,
	}
	if !validTypes[req.Type] {
		return nil, errors.New("invalid task type")
	}

	task := &model.Task{
		UserId:           userId,
		GoalId:           req.GoalId,
		Type:             req.Type,
		Status:           model.TaskStatusPending,
		TargetDate:       req.TargetDate,
		ChapterId:        req.ChapterId,
		PastPaperId:      req.PastPaperId,
		Title:            req.Title,
		Description:      req.Description,
		EstimatedMinutes: req.EstimatedMinutes,
		QuestionCount:    req.QuestionCount,
		Priority:         req.Priority,
	}

	if err := repository.TaskRepo.Create(task); err != nil {
		return nil, err
	}

	return repository.TaskRepo.GetByID(task.ID)
}

// UpdateTask updates a task
func (svr *TaskService) UpdateTask(userId uint, req model.TaskUpdateRequest) (*model.Task, error) {
	task, err := repository.TaskRepo.GetByID(req.ID)
	if err != nil {
		return nil, err
	}
	if task.UserId != userId {
		return nil, errors.New("unauthorized to update this task")
	}

	// Update fields
	if req.Status != nil {
		validStatuses := map[string]bool{
			model.TaskStatusPending:    true,
			model.TaskStatusInProgress: true,
			model.TaskStatusCompleted:  true,
			model.TaskStatusSkipped:    true,
			model.TaskStatusFailed:     true,
		}
		if !validStatuses[*req.Status] {
			return nil, errors.New("invalid task status")
		}
		task.Status = *req.Status
	}
	if req.TargetDate != nil {
		task.TargetDate = *req.TargetDate
	}
	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.EstimatedMinutes != nil {
		task.EstimatedMinutes = *req.EstimatedMinutes
	}
	if req.QuestionCount != nil {
		task.QuestionCount = *req.QuestionCount
	}
	if req.Priority != nil {
		task.Priority = *req.Priority
	}
	if req.IsLocked != nil {
		task.IsLocked = *req.IsLocked
	}
	if req.CompletedAt != nil {
		task.CompletedAt = req.CompletedAt
	}

	if err := repository.TaskRepo.Update(task); err != nil {
		return nil, err
	}

	return repository.TaskRepo.GetByID(task.ID)
}

// GetTaskByID gets a task by ID
func (svr *TaskService) GetTaskByID(userId, taskId uint) (*model.Task, error) {
	task, err := repository.TaskRepo.GetByID(taskId)
	if err != nil {
		return nil, err
	}
	if task.UserId != userId {
		return nil, errors.New("unauthorized to access this task")
	}
	return task, nil
}

// GetTaskStream gets today's tasks, upcoming tasks, and overdue tasks
func (svr *TaskService) GetTaskStream(userId, goalId uint) (*model.TaskStreamResponse, error) {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to access this goal")
	}

	todayTasks, err := repository.TaskRepo.GetTodayTasks(userId, goalId)
	if err != nil {
		return nil, err
	}

	upcomingTasks, err := repository.TaskRepo.GetUpcomingTasks(userId, goalId, 7)
	if err != nil {
		return nil, err
	}

	overdueTasks, err := repository.TaskRepo.GetOverdueTasks(userId, goalId)
	if err != nil {
		return nil, err
	}

	return &model.TaskStreamResponse{
		TodayTasks:    todayTasks,
		UpcomingTasks: upcomingTasks,
		OverdueTasks:  overdueTasks,
	}, nil
}

// CompleteTask marks a task as completed
func (svr *TaskService) CompleteTask(userId, taskId uint) error {
	task, err := repository.TaskRepo.GetByID(taskId)
	if err != nil {
		return err
	}
	if task.UserId != userId {
		return errors.New("unauthorized to complete this task")
	}

	now := time.Now()
	task.Status = model.TaskStatusCompleted
	task.CompletedAt = &now

	return repository.TaskRepo.Update(task)
}

// GenerateInitialPlan generates initial 7-day task plan for a new goal
func (svr *TaskService) GenerateInitialPlan(userId, goalId uint) error {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return err
	}
	if goal.UserId != userId {
		return errors.New("unauthorized to generate plan for this goal")
	}

	// Get all chapters for the syllabus
	chapters, err := repository.ChapterRepo.FindBySyllabusID(goal.SyllabusId)
	if err != nil {
		return err
	}

	if len(chapters) == 0 {
		return errors.New("no chapters found for this syllabus")
	}

	// Generate tasks for the next 7 days
	today := time.Now().Truncate(24 * time.Hour)

	// Simple initial plan: alternate between learn and drill tasks
	// Day 1-2: Learn first chapter
	// Day 3: Drill first chapter
	// Day 4-5: Learn second chapter
	// Day 6: Drill second chapter
	// Day 7: Review all covered chapters

	tasks := []model.Task{
		{
			UserId:           userId,
			GoalId:           goalId,
			Type:             model.TaskTypeLearn,
			Status:           model.TaskStatusPending,
			TargetDate:       today,
			ChapterId:        &chapters[0].ID,
			Title:            fmt.Sprintf("Learn: %s", chapters[0].Name),
			Description:      fmt.Sprintf("Study the basics of %s", chapters[0].Name),
			EstimatedMinutes: 30,
			Priority:         10,
			PlanVersion:      1,
		},
		{
			UserId:           userId,
			GoalId:           goalId,
			Type:             model.TaskTypeDrill,
			Status:           model.TaskStatusPending,
			TargetDate:       today.Add(24 * time.Hour),
			ChapterId:        &chapters[0].ID,
			Title:            fmt.Sprintf("Practice: %s", chapters[0].Name),
			Description:      "Complete practice questions",
			EstimatedMinutes: 30,
			QuestionCount:    10,
			Priority:         10,
			PlanVersion:      1,
		},
	}

	// Add more chapters if available
	if len(chapters) > 1 {
		tasks = append(tasks,
			model.Task{
				UserId:           userId,
				GoalId:           goalId,
				Type:             model.TaskTypeLearn,
				Status:           model.TaskStatusPending,
				TargetDate:       today.Add(2 * 24 * time.Hour),
				ChapterId:        &chapters[1].ID,
				Title:            fmt.Sprintf("Learn: %s", chapters[1].Name),
				Description:      fmt.Sprintf("Study the basics of %s", chapters[1].Name),
				EstimatedMinutes: 30,
				Priority:         10,
				PlanVersion:      1,
			},
			model.Task{
				UserId:           userId,
				GoalId:           goalId,
				Type:             model.TaskTypeDrill,
				Status:           model.TaskStatusPending,
				TargetDate:       today.Add(3 * 24 * time.Hour),
				ChapterId:        &chapters[1].ID,
				Title:            fmt.Sprintf("Practice: %s", chapters[1].Name),
				Description:      "Complete practice questions",
				EstimatedMinutes: 30,
				QuestionCount:    10,
				Priority:         10,
				PlanVersion:      1,
			},
		)
	}

	// Add review task for day 7
	tasks = append(tasks, model.Task{
		UserId:           userId,
		GoalId:           goalId,
		Type:             model.TaskTypeReview,
		Status:           model.TaskStatusPending,
		TargetDate:       today.Add(6 * 24 * time.Hour),
		Title:            "Review: Weekly Recap",
		Description:      "Review all chapters covered this week",
		EstimatedMinutes: 45,
		QuestionCount:    15,
		Priority:         10,
		PlanVersion:      1,
	})

	// Create all tasks
	for _, task := range tasks {
		if err := repository.TaskRepo.Create(&task); err != nil {
			return err
		}
	}

	return nil
}

// DeleteTask deletes a task
func (svr *TaskService) DeleteTask(userId, taskId uint) error {
	task, err := repository.TaskRepo.GetByID(taskId)
	if err != nil {
		return err
	}
	if task.UserId != userId {
		return errors.New("unauthorized to delete this task")
	}
	return repository.TaskRepo.Delete(taskId)
}

// ListTasks lists tasks with pagination
func (svr *TaskService) ListTasks(userId uint, query model.TaskQuery) ([]model.Task, int64, error) {
	query.UserId = userId
	return repository.TaskRepo.FindPage(&query)
}
