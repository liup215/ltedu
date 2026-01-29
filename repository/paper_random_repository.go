package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

// IRandomPaperRepository 随机试卷数据访问接口
type IRandomPaperRepository interface {
	Create(paper *model.RandomPaper) error
	Update(paper *model.RandomPaper) error
	Delete(id uint) error
	FindByID(id uint) (*model.RandomPaper, error)
	FindPage(query *model.RandomPaperQuery, offset, limit int) ([]*model.RandomPaper, int64, error)
	FindAll(query *model.RandomPaperQuery) ([]*model.RandomPaper, error)
}

type randomPaperRepository struct {
	db *gorm.DB
}

// NewRandomPaperRepository 创建随机试卷仓储实例
func NewRandomPaperRepository(db *gorm.DB) IRandomPaperRepository {
	return &randomPaperRepository{db: db}
}

func (r *randomPaperRepository) Create(paper *model.RandomPaper) error {
	return r.db.Omit("Syllabus", "PaperSeries", "PaperCode").Create(paper).Error
}

func (r *randomPaperRepository) Update(paper *model.RandomPaper) error {
	return r.db.Model(paper).Updates(paper).Error
}

func (r *randomPaperRepository) Delete(id uint) error {
	return r.db.Delete(&model.RandomPaper{}, id).Error
}

func (r *randomPaperRepository) FindByID(id uint) (*model.RandomPaper, error) {
	var paper model.RandomPaper
	err := r.db.Where("id = ?", id).
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Preload("PaperSeries").
		Preload("PaperCode").
		First(&paper).Error
	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return &paper, err
}

func (r *randomPaperRepository) FindPage(query *model.RandomPaperQuery, offset, limit int) ([]*model.RandomPaper, int64, error) {
	var papers []*model.RandomPaper
	var total int64

	q := r.db.Model(&model.RandomPaper{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}

	q.Count(&total)

	err := q.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Preload("PaperSeries").
		Preload("PaperCode").
		Order("id DESC").
		Offset(offset).
		Limit(limit).
		Find(&papers).Error

	return papers, total, err
}

func (r *randomPaperRepository) FindAll(query *model.RandomPaperQuery) ([]*model.RandomPaper, error) {
	var papers []*model.RandomPaper

	q := r.db.Model(&model.RandomPaper{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}

	err := q.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Preload("PaperSeries").
		Preload("PaperCode").
		Order("id DESC").
		Find(&papers).Error

	return papers, err
}
