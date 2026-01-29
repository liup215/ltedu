package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

// IPastPaperRepository 历年试卷数据访问接口
type IPastPaperRepository interface {
	Create(paper *model.PastPaper) error
	Update(paper *model.PastPaper) error
	Delete(id uint) error
	FindByID(id uint) (*model.PastPaper, error)
	FindPage(query *model.PastPaperQuery, offset, limit int) ([]*model.PastPaper, int64, error)
	FindAll(query *model.PastPaperQuery) ([]*model.PastPaper, error)
}

type pastPaperRepository struct {
	db *gorm.DB
}

// NewPastPaperRepository 创建历年试卷仓储实例
func NewPastPaperRepository(db *gorm.DB) IPastPaperRepository {
	return &pastPaperRepository{db: db}
}

func (r *pastPaperRepository) Create(paper *model.PastPaper) error {
	return r.db.Omit("Syllabus", "PaperSeries", "PaperCode").Create(paper).Error
}

func (r *pastPaperRepository) Update(paper *model.PastPaper) error {
	return r.db.Model(paper).Updates(paper).Error
}

func (r *pastPaperRepository) Delete(id uint) error {
	return r.db.Delete(&model.PastPaper{}, id).Error
}

func (r *pastPaperRepository) FindByID(id uint) (*model.PastPaper, error) {
	var paper model.PastPaper
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

func (r *pastPaperRepository) FindPage(query *model.PastPaperQuery, offset, limit int) ([]*model.PastPaper, int64, error) {
	var papers []*model.PastPaper
	var total int64

	q := r.db.Model(&model.PastPaper{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.Name != "" {
		q = q.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.PaperSeriesId != 0 {
		q = q.Where("paper_series_id = ?", query.PaperSeriesId)
	}
	if query.PaperCodeId != 0 {
		q = q.Where("paper_code_id = ?", query.PaperCodeId)
	}
	if query.Year != 0 {
		q = q.Where("year = ?", query.Year)
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

func (r *pastPaperRepository) FindAll(query *model.PastPaperQuery) ([]*model.PastPaper, error) {
	var papers []*model.PastPaper

	q := r.db.Model(&model.PastPaper{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.Name != "" {
		q = q.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.PaperSeriesId != 0 {
		q = q.Where("paper_series_id = ?", query.PaperSeriesId)
	}
	if query.PaperCodeId != 0 {
		q = q.Where("paper_code_id = ?", query.PaperCodeId)
	}
	if query.Year != 0 {
		q = q.Where("year = ?", query.Year)
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
