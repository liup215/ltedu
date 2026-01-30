package repository

import (
	"edu/model"

	"gorm.io/gorm"
)

// ISyllabusRepository 考纲数据访问接口
type ISyllabusRepository interface {
	Create(s *model.Syllabus) error
	Update(s *model.Syllabus) error
	Delete(id uint) error
	FindByID(id uint) (*model.Syllabus, error)
	FindPage(query *model.SyllabusQuery, offset, limit int) ([]*model.Syllabus, int64, error)
	FindAll(query *model.SyllabusQuery) ([]*model.Syllabus, error)
}

type syllabusRepository struct {
	db *gorm.DB
}

func NewSyllabusRepository(db *gorm.DB) ISyllabusRepository {
	return &syllabusRepository{db: db}
}

func (r *syllabusRepository) Create(s *model.Syllabus) error {
	return r.db.Create(s).Error
}

func (r *syllabusRepository) Update(s *model.Syllabus) error {
	return r.db.Model(s).Updates(s).Error
}

func (r *syllabusRepository) Delete(id uint) error {
	return r.db.Delete(&model.Syllabus{}, id).Error
}

func (r *syllabusRepository) FindByID(id uint) (*model.Syllabus, error) {
	var syl model.Syllabus
	err := r.db.Where("id = ?", id).
		Preload("Qualification").
		Preload("Qualification.Organisation").
		First(&syl).Error
	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return &syl, err
}

func (r *syllabusRepository) FindPage(query *model.SyllabusQuery, offset, limit int) ([]*model.Syllabus, int64, error) {
	var syllabi []*model.Syllabus
	var total int64

	q := r.db.Model(&model.Syllabus{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.Code != "" {
		q = q.Where("code = ?", query.Code)
	}
	if query.QualificationId != 0 {
		q = q.Where("qualification_id = ?", query.QualificationId)
	}

	q.Count(&total)
	err := q.
		Preload("Qualification").
		Preload("Qualification.Organisation").
		Order("id DESC").
		Offset(offset).
		Limit(limit).
		Find(&syllabi).Error
	return syllabi, total, err
}

func (r *syllabusRepository) FindAll(query *model.SyllabusQuery) ([]*model.Syllabus, error) {
	var syllabi []*model.Syllabus
	q := r.db.Model(&model.Syllabus{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.Code != "" {
		q = q.Where("code = ?", query.Code)
	}
	if query.QualificationId != 0 {
		q = q.Where("qualification_id = ?", query.QualificationId)
	}

	err := q.
		Preload("Qualification").
		Preload("Qualification.Organisation").
		Order("id DESC").
		Find(&syllabi).Error
	return syllabi, err
}
