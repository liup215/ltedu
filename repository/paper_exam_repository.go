package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

// IExamPaperRepository 考试试卷数据访问接口
type IExamPaperRepository interface {
	Create(paper *model.ExamPaper) error
	Update(paper *model.ExamPaper) error
	Delete(id uint) error
	FindByID(id uint) (*model.ExamPaper, error)
	FindAll(query *model.ExamPaperQuery) ([]*model.ExamPaper, error)
	FindPage(query *model.ExamPaperQuery, offset, limit int) ([]*model.ExamPaper, int64, error)
}

type examPaperRepository struct {
	db *gorm.DB
}

// NewExamPaperRepository 创建考试试卷仓储实例
func NewExamPaperRepository(db *gorm.DB) IExamPaperRepository {
	return &examPaperRepository{db: db}
}

func (r *examPaperRepository) Create(paper *model.ExamPaper) error {
	return r.db.Omit("Syllabus", "User").Create(paper).Error
}

func (r *examPaperRepository) Update(paper *model.ExamPaper) error {
	return r.db.Model(paper).Updates(paper).Error
}

func (r *examPaperRepository) Delete(id uint) error {
	return r.db.Delete(&model.ExamPaper{}, id).Error
}

func (r *examPaperRepository) FindByID(id uint) (*model.ExamPaper, error) {
	var paper model.ExamPaper
	err := r.db.Where("id = ?", id).
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		First(&paper).Error
	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return &paper, err
}

func (r *examPaperRepository) FindPage(query *model.ExamPaperQuery, offset, limit int) ([]*model.ExamPaper, int64, error) {
	var papers []*model.ExamPaper
	var total int64

	q := r.db.Model(&model.ExamPaper{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.UserId != 0 {
		q = q.Where("user_id = ?", query.UserId)
	}

	q.Count(&total)

	err := q.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Order("id DESC").
		Offset(offset).
		Limit(limit).
		Find(&papers).Error

	return papers, total, err
}

func (r *examPaperRepository) FindAll(query *model.ExamPaperQuery) ([]*model.ExamPaper, error) {
	var papers []*model.ExamPaper

	q := r.db.Model(&model.ExamPaper{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.UserId != 0 {
		q = q.Where("user_id = ?", query.UserId)
	}

	err := q.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Order("id DESC").
		Find(&papers).Error

	return papers, err
}
