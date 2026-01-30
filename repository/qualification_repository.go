package repository

import (
	"edu/model"

	"gorm.io/gorm"
)

// IQualificationRepository 资格证书数据访问接口
type IQualificationRepository interface {
	Create(q *model.Qualification) error
	Update(q *model.Qualification) error
	Delete(id uint) error
	FindByID(id uint) (*model.Qualification, error)
	FindPage(query *model.QualificationQuery, offset, limit int) ([]*model.Qualification, int64, error)
	FindAll(query *model.QualificationQuery) ([]*model.Qualification, error)
}

type qualificationRepository struct {
	db *gorm.DB
}

func NewQualificationRepository(db *gorm.DB) IQualificationRepository {
	return &qualificationRepository{db: db}
}

func (r *qualificationRepository) Create(q *model.Qualification) error {
	return r.db.Create(q).Error
}

func (r *qualificationRepository) Update(q *model.Qualification) error {
	return r.db.Model(q).Updates(q).Error
}

func (r *qualificationRepository) Delete(id uint) error {
	return r.db.Delete(&model.Qualification{}, id).Error
}

func (r *qualificationRepository) FindByID(id uint) (*model.Qualification, error) {
	var qual model.Qualification
	err := r.db.Where("id = ?", id).Preload("Organisation").First(&qual).Error
	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return &qual, err
}

func (r *qualificationRepository) FindPage(query *model.QualificationQuery, offset, limit int) ([]*model.Qualification, int64, error) {
	var quals []*model.Qualification
	var total int64

	q := r.db.Model(&model.Qualification{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.OrganisationId != 0 {
		q = q.Where("organisation_id = ?", query.OrganisationId)
	}

	q.Count(&total)
	err := q.Preload("Organisation").Order("id DESC").Offset(offset).Limit(limit).Find(&quals).Error
	return quals, total, err
}

func (r *qualificationRepository) FindAll(query *model.QualificationQuery) ([]*model.Qualification, error) {
	var quals []*model.Qualification
	q := r.db.Model(&model.Qualification{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.OrganisationId != 0 {
		q = q.Where("organisation_id = ?", query.OrganisationId)
	}

	err := q.Preload("Organisation").Order("id DESC").Find(&quals).Error
	return quals, err
}
