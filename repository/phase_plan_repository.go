package repository

import (
	"edu/model"

	"gorm.io/gorm"
)

// IPhasePlanRepository 阶段性计划数据访问接口
type IPhasePlanRepository interface {
	Create(p *model.LearningPhasePlan) error
	Update(p *model.LearningPhasePlan) error
	Delete(id uint) error
	FindByID(id uint) (*model.LearningPhasePlan, error)
	FindByPlanID(planId uint) ([]*model.LearningPhasePlan, error)
	FindPage(query *model.LearningPhasePlanQuery, offset, limit int) ([]*model.LearningPhasePlan, int64, error)
	// Chapter management
	AddChapters(phasePlanId uint, chapterIds []uint) error
	RemoveChapter(phasePlanId uint, chapterId uint) error
}

type phasePlanRepository struct {
	db *gorm.DB
}

func NewPhasePlanRepository(db *gorm.DB) IPhasePlanRepository {
	return &phasePlanRepository{db: db}
}

func (r *phasePlanRepository) Create(p *model.LearningPhasePlan) error {
	return r.db.Create(p).Error
}

func (r *phasePlanRepository) Update(p *model.LearningPhasePlan) error {
	return r.db.Model(p).Select("title", "start_date", "end_date", "sort_order").Updates(p).Error
}

func (r *phasePlanRepository) Delete(id uint) error {
	return r.db.Delete(&model.LearningPhasePlan{}, id).Error
}

func (r *phasePlanRepository) FindByID(id uint) (*model.LearningPhasePlan, error) {
	var p model.LearningPhasePlan
	err := r.db.Where("id = ?", id).
		Preload("ExamNode").
		Preload("Chapters").
		First(&p).Error
	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return &p, err
}

func (r *phasePlanRepository) FindByPlanID(planId uint) ([]*model.LearningPhasePlan, error) {
	var plans []*model.LearningPhasePlan
	err := r.db.Where("plan_id = ?", planId).
		Preload("ExamNode").
		Preload("Chapters").
		Order("sort_order ASC, id ASC").
		Find(&plans).Error
	return plans, err
}

func (r *phasePlanRepository) FindPage(query *model.LearningPhasePlanQuery, offset, limit int) ([]*model.LearningPhasePlan, int64, error) {
	var plans []*model.LearningPhasePlan
	var total int64

	q := r.db.Model(&model.LearningPhasePlan{})
	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.PlanId != 0 {
		q = q.Where("plan_id = ?", query.PlanId)
	}
	if query.ExamNodeId != 0 {
		q = q.Where("exam_node_id = ?", query.ExamNodeId)
	}

	q.Count(&total)
	err := q.
		Preload("ExamNode").
		Preload("Chapters").
		Order("sort_order ASC, id ASC").
		Offset(offset).
		Limit(limit).
		Find(&plans).Error
	return plans, total, err
}

func (r *phasePlanRepository) AddChapters(phasePlanId uint, chapterIds []uint) error {
	if len(chapterIds) == 0 {
		return nil
	}
	// Load existing to avoid duplicates in the join table
	var pp model.LearningPhasePlan
	if err := r.db.Where("id = ?", phasePlanId).Preload("Chapters").First(&pp).Error; err != nil {
		return err
	}
	existing := make(map[uint]struct{}, len(pp.Chapters))
	for _, ch := range pp.Chapters {
		existing[ch.ID] = struct{}{}
	}
	var toAdd []*model.Chapter
	for _, id := range chapterIds {
		if _, ok := existing[id]; !ok {
			toAdd = append(toAdd, &model.Chapter{Model: model.Model{ID: id}})
		}
	}
	if len(toAdd) == 0 {
		return nil
	}
	return r.db.Model(&model.LearningPhasePlan{Model: model.Model{ID: phasePlanId}}).
		Association("Chapters").
		Append(toAdd)
}

func (r *phasePlanRepository) RemoveChapter(phasePlanId uint, chapterId uint) error {
	return r.db.Model(&model.LearningPhasePlan{Model: model.Model{ID: phasePlanId}}).
		Association("Chapters").
		Delete(&model.Chapter{Model: model.Model{ID: chapterId}})
}
