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
	AddKnowledgePoint(questionId, knowledgePointId uint) error
	RemoveKnowledgePoint(questionId, knowledgePointId uint) error
	ClearKnowledgePoints(questionId uint) error
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
	return r.db.Model(q).Omit("Syllabus", "PastPaper", "Chapters").Updates(q).Error
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
		Where("id = ?", id).First(&q).Error
	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	if err == nil {
		_ = q.Format()
	}
	return &q, err
}

func (r *questionRepository) applyQuestionFilters(q *gorm.DB, query *model.QuestionQueryRequest, tableName string) *gorm.DB {
	if query.ID != 0 {
		q = q.Where(tableName+".id = ?", query.ID)
	}
	if query.Stem != "" {
		q = q.Where("stem LIKE ?", "%"+query.Stem+"%")
	}
	if query.SyllabusId != 0 {
		q = q.Where(tableName+".syllabus_id = ?", query.SyllabusId)
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
	if query.PaperName != "" {
		q = q.Joins("PastPaper").
			Where("PastPaper.name LIKE ?", "%"+query.PaperName+"%")
	}
	if query.ExamNodeId != 0 {
		q = q.Joins("JOIN past_papers ON past_papers.id = "+tableName+".past_paper_id").
			Joins("JOIN paper_codes ON paper_codes.id = past_papers.paper_code_id").
			Where("paper_codes.exam_node_id = ?", query.ExamNodeId)
	}
	if query.KnowledgePointId != 0 {
		q = q.Joins("JOIN question_keypoints ON question_keypoints.question_id = "+tableName+".id").
			Where("question_keypoints.knowledge_point_id = ?", query.KnowledgePointId)
	}
	return q
}

func (r *questionRepository) FindPage(query *model.QuestionQueryRequest, offset, limit int) ([]*model.Question, int64, error) {
	var questions []*model.Question
	var total int64

	tableName := GetTableName(r.db, &model.Question{})
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
		Preload("KnowledgePoints")

	q = r.applyQuestionFilters(q, query, tableName)

	q.Count(&total)
	err := q.Order(tableName + ".id DESC").Offset(offset).Limit(limit).Find(&questions).Error
	for _, q := range questions {
		_ = q.Format()
	}
	return questions, total, err
}

func (r *questionRepository) FindAll(query *model.QuestionQueryRequest) ([]*model.Question, error) {
	var questions []*model.Question
	tableName := GetTableName(r.db, &model.Question{})
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
		Preload("PastPaper.PaperSeries.Syllabus.Qualification.Organisation")

	q = r.applyQuestionFilters(q, query, tableName)

	err := q.Order(tableName + ".id DESC").Find(&questions).Error
	for _, q := range questions {
		_ = q.Format()
	}
	return questions, err
}

func (r *questionRepository) Count(query *model.QuestionQueryRequest) (int64, error) {
	var total int64
	tableName := GetTableName(r.db, &model.Question{})
	q := r.db.Model(&model.Question{})

	q = r.applyQuestionFilters(q, query, tableName)

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
		Where("id IN ?", ids).Find(&questions).Error
	for _, q := range questions {
		_ = q.Format()
	}
	return questions, err
}

func (r *questionRepository) AddKnowledgePoint(questionId, knowledgePointId uint) error {
	return r.db.Exec("INSERT IGNORE INTO question_keypoints (question_id, knowledge_point_id) VALUES (?, ?)",
		questionId, knowledgePointId).Error
}

func (r *questionRepository) RemoveKnowledgePoint(questionId, knowledgePointId uint) error {
	return r.db.Exec("DELETE FROM question_keypoints WHERE question_id = ? AND knowledge_point_id = ?",
		questionId, knowledgePointId).Error
}

func (r *questionRepository) ClearKnowledgePoints(questionId uint) error {
	return r.db.Exec("DELETE FROM question_keypoints WHERE question_id = ?", questionId).Error
}
