package repository

import (
	"edu/model"

	"gorm.io/gorm"
)

// IQuestionRepository 题目数据访问接口
type IQuestionRepository interface {
	Create(q *model.Question) error
	Update(q *model.Question) error
	Delete(id uint) error
	FindByID(id uint) (*model.Question, error)
	FindPage(query *model.QuestionQueryRequest, offset, limit int) ([]*model.Question, int64, error)
	FindAll(query *model.QuestionQueryRequest) ([]*model.Question, error)
	Count(query *model.QuestionQueryRequest) (int64, error)
	FindByIDs(ids []uint) ([]*model.Question, error)
}

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) IQuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) Create(q *model.Question) error {
	return r.db.Create(q).Error
}

func (r *questionRepository) Update(q *model.Question) error {
	return r.db.Model(q).Updates(q).Error
}

func (r *questionRepository) Delete(id uint) error {
	return r.db.Delete(&model.Question{}, id).Error
}

func (r *questionRepository) FindByID(id uint) (*model.Question, error) {
	var q model.Question
	err := r.db.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Preload("PastPaper").
		Preload("PastPaper.Syllabus").
		Preload("PastPaper.Syllabus.Qualification").
		Preload("PastPaper.Syllabus.Qualification.Organisation").
		Preload("PastPaper.PaperCode").
		Preload("PastPaper.PaperCode.Syllabus").
		Preload("PastPaper.PaperCode.Syllabus.Qualification").
		Preload("PastPaper.PaperCode.Syllabus.Qualification.Organisation").
		Preload("PastPaper.PaperSeries").
		Preload("PastPaper.PaperSeries.Syllabus").
		Preload("PastPaper.PaperSeries.Syllabus.Qualification").
		Preload("PastPaper.PaperSeries.Syllabus.Qualification.Organisation").
		Preload("Chapters").
		Preload("Chapters.Syllabus").
		Preload("Chapters.Syllabus.Qualification").
		Preload("Chapters.Syllabus.Qualification.Organisation").
		Where("id = ?", id).First(&q).Error
	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	if err == nil {
		_ = q.Format()
	}
	return &q, err
}

func (r *questionRepository) FindPage(query *model.QuestionQueryRequest, offset, limit int) ([]*model.Question, int64, error) {
	var questions []*model.Question
	var total int64

	q := r.db.Model(&model.Question{}).
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Preload("PastPaper").
		Preload("PastPaper.Syllabus").
		Preload("PastPaper.Syllabus.Qualification").
		Preload("PastPaper.Syllabus.Qualification.Organisation").
		Preload("PastPaper.PaperCode").
		Preload("PastPaper.PaperCode.Syllabus").
		Preload("PastPaper.PaperCode.Syllabus.Qualification").
		Preload("PastPaper.PaperCode.Syllabus.Qualification.Organisation").
		Preload("PastPaper.PaperSeries").
		Preload("PastPaper.PaperSeries.Syllabus").
		Preload("PastPaper.PaperSeries.Syllabus.Qualification").
		Preload("PastPaper.PaperSeries.Syllabus.Qualification.Organisation").
		Preload("Chapters").
		Preload("Chapters.Syllabus").
		Preload("Chapters.Syllabus.Qualification").
		Preload("Chapters.Syllabus.Qualification.Organisation")

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.Stem != "" {
		q = q.Where("stem LIKE ?", "%"+query.Stem+"%")
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.Difficult != 0 {
		q = q.Where("difficult = ?", query.Difficult)
	}
	if query.Status != 0 {
		q = q.Where("status = ?", query.Status)
	}
	if query.PastPaperId != 0 {
		q = q.Where("past_paper_id = ?", query.PastPaperId)
	}

	q.Count(&total)
	err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&questions).Error
	for _, q := range questions {
		_ = q.Format()
	}
	return questions, total, err
}

func (r *questionRepository) FindAll(query *model.QuestionQueryRequest) ([]*model.Question, error) {
	var questions []*model.Question
	q := r.db.Model(&model.Question{}).
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Preload("PastPaper").
		Preload("PastPaper.Syllabus").
		Preload("PastPaper.Syllabus.Qualification").
		Preload("PastPaper.Syllabus.Qualification.Organisation").
		Preload("PastPaper.PaperCode").
		Preload("PastPaper.PaperCode.Syllabus").
		Preload("PastPaper.PaperCode.Syllabus.Qualification").
		Preload("PastPaper.PaperCode.Syllabus.Qualification.Organisation").
		Preload("PastPaper.PaperSeries").
		Preload("PastPaper.PaperSeries.Syllabus").
		Preload("PastPaper.PaperSeries.Syllabus.Qualification").
		Preload("PastPaper.PaperSeries.Syllabus.Qualification.Organisation").
		Preload("Chapters").
		Preload("Chapters.Syllabus").
		Preload("Chapters.Syllabus.Qualification").
		Preload("Chapters.Syllabus.Qualification.Organisation")

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.Stem != "" {
		q = q.Where("stem LIKE ?", "%"+query.Stem+"%")
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.Difficult != 0 {
		q = q.Where("difficult = ?", query.Difficult)
	}
	if query.Status != 0 {
		q = q.Where("status = ?", query.Status)
	}
	if query.PastPaperId != 0 {
		q = q.Where("past_paper_id = ?", query.PastPaperId)
	}

	err := q.Order("id DESC").Find(&questions).Error
	for _, q := range questions {
		_ = q.Format()
	}
	return questions, err
}

func (r *questionRepository) Count(query *model.QuestionQueryRequest) (int64, error) {
	var total int64
	q := r.db.Model(&model.Question{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.Stem != "" {
		q = q.Where("stem LIKE ?", "%"+query.Stem+"%")
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.Difficult != 0 {
		q = q.Where("difficult = ?", query.Difficult)
	}
	if query.Status != 0 {
		q = q.Where("status = ?", query.Status)
	}
	if query.PastPaperId != 0 {
		q = q.Where("past_paper_id = ?", query.PastPaperId)
	}

	err := q.Count(&total).Error
	return total, err
}

func (r *questionRepository) FindByIDs(ids []uint) ([]*model.Question, error) {
	var questions []*model.Question
	err := r.db.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Preload("PastPaper").
		Preload("PastPaper.Syllabus").
		Preload("PastPaper.Syllabus.Qualification").
		Preload("PastPaper.Syllabus.Qualification.Organisation").
		Preload("PastPaper.PaperCode").
		Preload("PastPaper.PaperCode.Syllabus").
		Preload("PastPaper.PaperCode.Syllabus.Qualification").
		Preload("PastPaper.PaperCode.Syllabus.Qualification.Organisation").
		Preload("PastPaper.PaperSeries").
		Preload("PastPaper.PaperSeries.Syllabus").
		Preload("PastPaper.PaperSeries.Syllabus.Qualification").
		Preload("PastPaper.PaperSeries.Syllabus.Qualification.Organisation").
		Preload("Chapters").
		Preload("Chapters.Syllabus").
		Preload("Chapters.Syllabus.Qualification").
		Preload("Chapters.Syllabus.Qualification.Organisation").
		Where("id IN ?", ids).Find(&questions).Error
	for _, q := range questions {
		_ = q.Format()
	}
	return questions, err
}
