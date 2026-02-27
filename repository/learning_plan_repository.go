package repository

import (
	"edu/model"

	"gorm.io/gorm"
)

// IStudentLearningPlanRepository 学习计划数据访问接口
type IStudentLearningPlanRepository interface {
	Create(p *model.StudentLearningPlan) error
	Update(p *model.StudentLearningPlan) error
	Delete(id uint) error
	FindByID(id uint) (*model.StudentLearningPlan, error)
	FindPage(query *model.StudentLearningPlanQuery, offset, limit int) ([]*model.StudentLearningPlan, int64, error)
	FindAll(query *model.StudentLearningPlanQuery) ([]*model.StudentLearningPlan, error)
	// Version management
	CreateVersion(v *model.StudentLearningPlanVersion) error
	FindVersionsByPlanId(planId uint, offset, limit int) ([]*model.StudentLearningPlanVersion, int64, error)
	FindVersionByPlanAndVersion(planId uint, version int) (*model.StudentLearningPlanVersion, error)
}

type studentLearningPlanRepository struct {
	db *gorm.DB
}

func NewStudentLearningPlanRepository(db *gorm.DB) IStudentLearningPlanRepository {
	return &studentLearningPlanRepository{db: db}
}

func (r *studentLearningPlanRepository) Create(p *model.StudentLearningPlan) error {
	return r.db.Create(p).Error
}

func (r *studentLearningPlanRepository) Update(p *model.StudentLearningPlan) error {
	return r.db.Model(p).Updates(p).Error
}

func (r *studentLearningPlanRepository) Delete(id uint) error {
	return r.db.Delete(&model.StudentLearningPlan{}, id).Error
}

func (r *studentLearningPlanRepository) FindByID(id uint) (*model.StudentLearningPlan, error) {
	var p model.StudentLearningPlan
	err := r.db.Where("id = ?", id).
		Preload("User").
		Preload("Class").
		First(&p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &p, err
}

func (r *studentLearningPlanRepository) FindPage(query *model.StudentLearningPlanQuery, offset, limit int) ([]*model.StudentLearningPlan, int64, error) {
	var plans []*model.StudentLearningPlan
	var total int64

	q := r.db.Model(&model.StudentLearningPlan{}).Preload("User").Preload("Class")

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.ClassId != 0 {
		q = q.Where("class_id = ?", query.ClassId)
	}
	if query.UserId != 0 {
		q = q.Where("user_id = ?", query.UserId)
	}
	if query.PlanType != "" {
		q = q.Where("plan_type = ?", query.PlanType)
	}

	q.Count(&total)
	err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&plans).Error
	return plans, total, err
}

func (r *studentLearningPlanRepository) FindAll(query *model.StudentLearningPlanQuery) ([]*model.StudentLearningPlan, error) {
	var plans []*model.StudentLearningPlan

	q := r.db.Model(&model.StudentLearningPlan{}).Preload("User").Preload("Class")

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.ClassId != 0 {
		q = q.Where("class_id = ?", query.ClassId)
	}
	if query.UserId != 0 {
		q = q.Where("user_id = ?", query.UserId)
	}
	if query.PlanType != "" {
		q = q.Where("plan_type = ?", query.PlanType)
	}

	err := q.Order("id DESC").Find(&plans).Error
	return plans, err
}

func (r *studentLearningPlanRepository) CreateVersion(v *model.StudentLearningPlanVersion) error {
	return r.db.Create(v).Error
}

func (r *studentLearningPlanRepository) FindVersionsByPlanId(planId uint, offset, limit int) ([]*model.StudentLearningPlanVersion, int64, error) {
	var versions []*model.StudentLearningPlanVersion
	var total int64

	q := r.db.Model(&model.StudentLearningPlanVersion{}).Where("plan_id = ?", planId)
	q.Count(&total)
	err := q.Order("version DESC").Offset(offset).Limit(limit).Find(&versions).Error
	return versions, total, err
}

func (r *studentLearningPlanRepository) FindVersionByPlanAndVersion(planId uint, version int) (*model.StudentLearningPlanVersion, error) {
	var v model.StudentLearningPlanVersion
	err := r.db.Where("plan_id = ? AND version = ?", planId, version).First(&v).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &v, err
}
