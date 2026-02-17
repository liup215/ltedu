package service

import (
	"edu/model"
	"edu/repository"
	"errors"
	"time"
)

var GoalSvr = &GoalService{baseService: newBaseService()}

type GoalService struct {
	baseService
}

// CreateGoal creates a new learning goal for a user
func (svr *GoalService) CreateGoal(userId uint, req model.GoalCreateRequest) (*model.Goal, error) {
	// Validate request
	if req.SyllabusId == 0 {
		return nil, errors.New("syllabusId is required")
	}
	if req.ExamDate.Before(time.Now()) {
		return nil, errors.New("examDate must be in the future")
	}
	if req.WeeklyHours <= 0 {
		return nil, errors.New("weeklyHours must be greater than 0")
	}
	if req.Mode != "sync" && req.Mode != "self" {
		return nil, errors.New("mode must be 'sync' or 'self'")
	}

	// Check if user already has an active goal for this syllabus
	existing, err := repository.GoalRepo.GetByUserAndSyllabus(userId, req.SyllabusId)
	if err == nil && existing != nil && existing.Status == "active" {
		return nil, errors.New("user already has an active goal for this syllabus")
	}

	// Create new goal
	goal := &model.Goal{
		UserId:         userId,
		SyllabusId:     req.SyllabusId,
		ExamDate:       req.ExamDate,
		TargetScore:    req.TargetScore,
		TargetGrade:    req.TargetGrade,
		WeeklyHours:    req.WeeklyHours,
		Mode:           req.Mode,
		Status:         "active",
		DiagnosticDone: false,
		StartDate:      time.Now(),
	}

	if err := repository.GoalRepo.Create(goal); err != nil {
		return nil, err
	}

	// Initialize knowledge states for all chapters in the syllabus
	syllabus, err := repository.SyllabusRepo.FindByID(req.SyllabusId)
	if err != nil {
		return nil, err
	}

	if err := repository.KnowledgeStateRepo.InitializeForGoal(userId, goal.ID, syllabus.ID); err != nil {
		return nil, err
	}

	// Load complete goal with relations
	return repository.GoalRepo.GetByID(goal.ID)
}

// UpdateGoal updates an existing goal
func (svr *GoalService) UpdateGoal(userId uint, req model.GoalUpdateRequest) (*model.Goal, error) {
	// Get existing goal
	goal, err := repository.GoalRepo.GetByID(req.ID)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to update this goal")
	}

	// Update fields
	if req.ExamDate != nil {
		if req.ExamDate.Before(time.Now()) {
			return nil, errors.New("examDate must be in the future")
		}
		goal.ExamDate = *req.ExamDate
	}
	if req.TargetScore != nil {
		goal.TargetScore = *req.TargetScore
	}
	if req.TargetGrade != nil {
		goal.TargetGrade = *req.TargetGrade
	}
	if req.WeeklyHours != nil {
		if *req.WeeklyHours <= 0 {
			return nil, errors.New("weeklyHours must be greater than 0")
		}
		goal.WeeklyHours = *req.WeeklyHours
	}
	if req.Mode != nil {
		if *req.Mode != "sync" && *req.Mode != "self" {
			return nil, errors.New("mode must be 'sync' or 'self'")
		}
		goal.Mode = *req.Mode
	}
	if req.Status != nil {
		goal.Status = *req.Status
		if *req.Status == "completed" {
			now := time.Now()
			goal.CompletedAt = &now
		}
	}

	if err := repository.GoalRepo.Update(goal); err != nil {
		return nil, err
	}

	return repository.GoalRepo.GetByID(goal.ID)
}

// GetGoalByID gets a goal by ID
func (svr *GoalService) GetGoalByID(userId, goalId uint) (*model.Goal, error) {
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to access this goal")
	}
	return goal, nil
}

// GetActiveGoals gets all active goals for a user
func (svr *GoalService) GetActiveGoals(userId uint) ([]model.Goal, error) {
	return repository.GoalRepo.GetActiveGoalsByUser(userId)
}

// ListGoals lists goals with pagination
func (svr *GoalService) ListGoals(userId uint, query model.GoalQuery) ([]model.Goal, int64, error) {
	query.UserId = userId
	return repository.GoalRepo.FindPage(&query)
}

// DeleteGoal soft deletes a goal
func (svr *GoalService) DeleteGoal(userId, goalId uint) error {
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return err
	}
	if goal.UserId != userId {
		return errors.New("unauthorized to delete this goal")
	}
	return repository.GoalRepo.Delete(goalId)
}

// CompleteDiagnostic marks the diagnostic test as completed for a goal
func (svr *GoalService) CompleteDiagnostic(userId, goalId uint) error {
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return err
	}
	if goal.UserId != userId {
		return errors.New("unauthorized to update this goal")
	}

	goal.DiagnosticDone = true
	return repository.GoalRepo.Update(goal)
}
