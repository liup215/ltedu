package service

import (
	"edu/model"
	"edu/repository"
	"errors"
	"math"
	"time"
)

var KnowledgeStateSvr = &KnowledgeStateService{baseService: newBaseService()}

type KnowledgeStateService struct {
	baseService
}

// GetKnowledgeState gets knowledge state for a specific chapter
func (svr *KnowledgeStateService) GetKnowledgeState(userId, goalId, chapterId uint) (*model.KnowledgeState, error) {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to access this goal")
	}

	return repository.KnowledgeStateRepo.GetByUserGoalChapter(userId, goalId, chapterId)
}

// GetKnowledgeStates gets all knowledge states for a goal
func (svr *KnowledgeStateService) GetKnowledgeStates(userId, goalId uint) ([]model.KnowledgeState, error) {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to access this goal")
	}

	return repository.KnowledgeStateRepo.GetByUserAndGoal(userId, goalId)
}

// GetProgress calculates and returns user's progress for a goal
func (svr *KnowledgeStateService) GetProgress(userId, goalId uint) (*model.KnowledgeStateProgressResponse, error) {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to access this goal")
	}

	states, err := repository.KnowledgeStateRepo.GetByUserAndGoal(userId, goalId)
	if err != nil {
		return nil, err
	}

	if len(states) == 0 {
		return &model.KnowledgeStateProgressResponse{
			Coverage:      0,
			Mastery:       0,
			Stability:     0,
			ChapterStates: []model.KnowledgeStateChapterInfo{},
		}, nil
	}

	// Calculate metrics
	coveredCount := 0
	totalMastery := 0.0
	totalStability := 0.0
	chapterStates := make([]model.KnowledgeStateChapterInfo, len(states))

	for i, state := range states {
		if state.IsCovered {
			coveredCount++
		}
		totalMastery += state.MasteryLevel
		totalStability += state.StabilityScore

		chapterStates[i] = model.KnowledgeStateChapterInfo{
			ChapterId:      state.ChapterId,
			ChapterName:    state.Chapter.Name,
			MasteryLevel:   state.MasteryLevel,
			StabilityScore: state.StabilityScore,
			IsCovered:      state.IsCovered,
		}
	}

	coverage := float64(coveredCount) / float64(len(states)) * 100
	mastery := totalMastery / float64(len(states)) * 100
	stability := totalStability / float64(len(states)) * 100

	return &model.KnowledgeStateProgressResponse{
		Coverage:      math.Round(coverage*100) / 100,
		Mastery:       math.Round(mastery*100) / 100,
		Stability:     math.Round(stability*100) / 100,
		ChapterStates: chapterStates,
	}, nil
}

// UpdateKnowledgeStateFromAttempt updates knowledge state based on an attempt
func (svr *KnowledgeStateService) UpdateKnowledgeStateFromAttempt(attempt *model.Attempt) error {
	// Get or create knowledge state
	ks, err := repository.KnowledgeStateRepo.GetByUserGoalChapter(attempt.UserId, attempt.GoalId, attempt.ChapterId)
	if err != nil {
		return err
	}

	// Update counters
	ks.TotalCount++
	if attempt.IsCorrect != nil && *attempt.IsCorrect {
		ks.CorrectCount++
		ks.ConsecutiveCorrect++
		ks.ConsecutiveWrong = 0
	} else if attempt.IsCorrect != nil {
		ks.ConsecutiveWrong++
		ks.ConsecutiveCorrect = 0
	}

	// Mark as covered
	ks.IsCovered = true

	// Update mastery level (0.0 to 1.0)
	ks.MasteryLevel = float64(ks.CorrectCount) / float64(ks.TotalCount)

	// Factor in stability (0.0 to 1.0)
	// Low stability = short interval, high stability = longer interval
	const stabilityIncrement = 0.1  // Increment per consecutive correct
	const stabilityDecrement = 0.15 // Decrement per consecutive wrong
	
	// Update stability score based on consecutive correct answers
	// Stability increases with consecutive correct, decreases with consecutive wrong
	if ks.ConsecutiveCorrect > 0 {
		// Increase stability: cap at 1.0
		ks.StabilityScore = math.Min(1.0, ks.StabilityScore+stabilityIncrement*float64(ks.ConsecutiveCorrect))
	} else if ks.ConsecutiveWrong > 0 {
		// Decrease stability: floor at 0.0
		ks.StabilityScore = math.Max(0.0, ks.StabilityScore-stabilityDecrement*float64(ks.ConsecutiveWrong))
	}

	// Update last practice time
	now := time.Now()
	ks.LastPracticeAt = &now

	// Schedule next review using SRS
	nextReview := svr.calculateNextReview(ks)
	ks.NextReviewAt = &nextReview

	return repository.KnowledgeStateRepo.Update(ks)
}

// calculateNextReview calculates the next review date using a simple SRS algorithm
func (svr *KnowledgeStateService) calculateNextReview(ks *model.KnowledgeState) time.Time {
	// Simple SRS: interval based on mastery and stability
	// Higher mastery + higher stability = longer interval
	
	baseInterval := 1.0 // Base interval in days
	
	// Factor in mastery level (0.0 to 1.0)
	// Low mastery = short interval, high mastery = longer interval
	masteryFactor := 1.0 + ks.MasteryLevel*5.0 // 1x to 6x
	
	// Factor in stability (0.0 to 1.0)
	// Low stability = short interval, high stability = longer interval
	stabilityFactor := 1.0 + ks.StabilityScore*3.0 // 1x to 4x
	
	// Calculate interval in days
	intervalDays := baseInterval * masteryFactor * stabilityFactor
	
	// Apply consecutive correct bonus (up to 2x multiplier)
	if ks.ConsecutiveCorrect > 0 {
		consecutiveBonus := math.Min(2.0, 1.0+float64(ks.ConsecutiveCorrect)*0.2)
		intervalDays *= consecutiveBonus
	}
	
	// Apply consecutive wrong penalty (down to 0.5x multiplier)
	if ks.ConsecutiveWrong > 0 {
		consecutivePenalty := math.Max(0.5, 1.0-float64(ks.ConsecutiveWrong)*0.15)
		intervalDays *= consecutivePenalty
	}
	
	// Cap interval at 30 days max, 1 day min
	intervalDays = math.Min(30.0, math.Max(1.0, intervalDays))
	
	// Add interval to current time
	return time.Now().Add(time.Duration(intervalDays*24) * time.Hour)
}

// GetDueForReview gets chapters that are due for review
func (svr *KnowledgeStateService) GetDueForReview(userId, goalId uint) ([]model.KnowledgeState, error) {
	// Verify user owns the goal
	goal, err := repository.GoalRepo.GetByID(goalId)
	if err != nil {
		return nil, err
	}
	if goal.UserId != userId {
		return nil, errors.New("unauthorized to access this goal")
	}

	return repository.KnowledgeStateRepo.GetDueForReview(userId, goalId)
}
