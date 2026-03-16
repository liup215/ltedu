package service

import (
	"edu/model"
	"edu/repository"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var MigrationJobSvr = &MigrationJobService{baseService: newBaseService()}

type MigrationJobService struct {
	baseService
}

// CreateJob creates a new migration job and starts it asynchronously.
func (s *MigrationJobService) CreateJob(createdBy uint, syllabusId uint, options MigrateOptions) (*model.MigrationJob, error) {
	optJSON, err := json.Marshal(options)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize options: %w", err)
	}

	job := &model.MigrationJob{
		SyllabusId: syllabusId,
		Status:     model.MigrationJobStatusPending,
		Options:    string(optJSON),
		CreatedBy:  createdBy,
	}

	if err := repository.MigrationJobRepo.Create(job); err != nil {
		return nil, fmt.Errorf("failed to create migration job: %w", err)
	}

	// Start the job in background
	go s.runJob(job.ID)

	return job, nil
}

// GetJob retrieves a migration job by ID.
func (s *MigrationJobService) GetJob(id uint) (*model.MigrationJob, error) {
	job, err := repository.MigrationJobRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if job == nil {
		return nil, errors.New("migration job not found")
	}
	return job, nil
}

// ListJobs lists migration jobs with pagination.
func (s *MigrationJobService) ListJobs(query model.MigrationJobQuery) ([]model.MigrationJob, int64, error) {
	return repository.MigrationJobRepo.FindPage(&query)
}

// RetryJob re-queues a failed job.
func (s *MigrationJobService) RetryJob(id uint) (*model.MigrationJob, error) {
	job, err := repository.MigrationJobRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if job == nil {
		return nil, errors.New("migration job not found")
	}
	if job.Status != model.MigrationJobStatusFailed {
		return nil, errors.New("only failed jobs can be retried")
	}

	job.Status = model.MigrationJobStatusPending
	job.Progress = 0
	job.DoneItems = 0
	job.TotalItems = 0
	job.ErrorMessage = ""
	job.Report = ""
	job.StartedAt = nil
	job.CompletedAt = nil

	if err := repository.MigrationJobRepo.Update(job); err != nil {
		return nil, fmt.Errorf("failed to reset job: %w", err)
	}

	go s.runJob(job.ID)

	return job, nil
}

// runJob executes a migration job in the background.
func (s *MigrationJobService) runJob(jobID uint) {
	job, err := repository.MigrationJobRepo.GetByID(jobID)
	if err != nil || job == nil {
		return
	}

	// Mark as running
	now := time.Now()
	job.Status = model.MigrationJobStatusRunning
	job.StartedAt = &now
	_ = repository.MigrationJobRepo.Update(job)

	// Decode options
	var options MigrateOptions
	if err := json.Unmarshal([]byte(job.Options), &options); err != nil {
		s.failJob(job, fmt.Sprintf("invalid options: %v", err))
		return
	}

	// Progress callback: updates job in DB
	onProgress := func(done, total int) {
		j, err := repository.MigrationJobRepo.GetByID(jobID)
		if err != nil || j == nil {
			return
		}
		j.DoneItems = done
		j.TotalItems = total
		if total > 0 {
			j.Progress = done * 100 / total
		}
		_ = repository.MigrationJobRepo.Update(j)
	}

	report, err := KnowledgePointSvr.AutoMigrateSyllabusWithProgress(job.SyllabusId, options, onProgress)

	// Reload job to avoid overwriting concurrent updates
	latest, dbErr := repository.MigrationJobRepo.GetByID(jobID)
	if dbErr != nil || latest == nil {
		return
	}

	completedAt := time.Now()
	latest.CompletedAt = &completedAt
	latest.Progress = 100

	if err != nil {
		latest.Status = model.MigrationJobStatusFailed
		latest.ErrorMessage = err.Error()
	} else {
		latest.Status = model.MigrationJobStatusCompleted
	}

	if report != nil {
		reportJSON, jsonErr := json.Marshal(report)
		if jsonErr == nil {
			latest.Report = string(reportJSON)
		}
		if len(report.Errors) > 0 && latest.ErrorMessage == "" {
			latest.ErrorMessage = fmt.Sprintf("%d error(s) during migration", len(report.Errors))
		}
	}

	_ = repository.MigrationJobRepo.Update(latest)
}

// failJob marks a job as failed with an error message.
func (s *MigrationJobService) failJob(job *model.MigrationJob, msg string) {
	completedAt := time.Now()
	job.Status = model.MigrationJobStatusFailed
	job.ErrorMessage = msg
	job.CompletedAt = &completedAt
	_ = repository.MigrationJobRepo.Update(job)
}
