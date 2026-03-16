package repository

import (
	"edu/model"
	"errors"
	"gorm.io/gorm"
)

type IMigrationJobRepository interface {
	Create(job *model.MigrationJob) error
	Update(job *model.MigrationJob) error
	GetByID(id uint) (*model.MigrationJob, error)
	FindPage(query *model.MigrationJobQuery) ([]model.MigrationJob, int64, error)
}

type migrationJobRepository struct {
	db *gorm.DB
}

func NewMigrationJobRepository(db *gorm.DB) IMigrationJobRepository {
	return &migrationJobRepository{db: db}
}

func (r *migrationJobRepository) Create(job *model.MigrationJob) error {
	return r.db.Create(job).Error
}

func (r *migrationJobRepository) Update(job *model.MigrationJob) error {
	return r.db.Save(job).Error
}

func (r *migrationJobRepository) GetByID(id uint) (*model.MigrationJob, error) {
	var job model.MigrationJob
	err := r.db.First(&job, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &job, err
}

func (r *migrationJobRepository) FindPage(query *model.MigrationJobQuery) ([]model.MigrationJob, int64, error) {
	var jobs []model.MigrationJob
	var total int64

	q := r.db.Model(&model.MigrationJob{})

	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.Status != "" {
		q = q.Where("`status` = ?", query.Status)
	}
	if query.CreatedBy != 0 {
		q = q.Where("created_by = ?", query.CreatedBy)
	}

	q.Count(&total)

	p := query.Page.CheckPage()
	offset := (p.PageIndex - 1) * p.PageSize
	err := q.Order("id DESC").
		Offset(offset).
		Limit(p.PageSize).
		Find(&jobs).Error

	return jobs, total, err
}
