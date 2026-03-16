package repository

import (
	"edu/model"

	"gorm.io/gorm"
)

type IKnowledgePointRepository interface {
	Create(kp *model.KnowledgePoint) error
	Update(kp *model.KnowledgePoint) error
	Delete(id uint) error
	DeleteByChapterId(chapterId uint) error
	FindByID(id uint) (*model.KnowledgePoint, error)
	FindByChapterId(chapterId uint) ([]model.KnowledgePoint, error)
	FindBySyllabusId(syllabusId uint) ([]model.KnowledgePoint, error)
	FindAll(query *model.KnowledgePointQuery) ([]model.KnowledgePoint, int64, error)
	BatchCreate(kps []model.KnowledgePoint) error
	LinkQuestion(knowledgePointId uint, questionId uint) error
	UnlinkQuestion(knowledgePointId uint, questionId uint) error
}

type KnowledgePointRepository struct {
	db *gorm.DB
}

func NewKnowledgePointRepository(db *gorm.DB) IKnowledgePointRepository {
	return &KnowledgePointRepository{db: db}
}

func (r *KnowledgePointRepository) Create(kp *model.KnowledgePoint) error {
	return r.db.Create(kp).Error
}

func (r *KnowledgePointRepository) Update(kp *model.KnowledgePoint) error {
	return r.db.Save(kp).Error
}

func (r *KnowledgePointRepository) Delete(id uint) error {
	return r.db.Delete(&model.KnowledgePoint{}, id).Error
}

func (r *KnowledgePointRepository) DeleteByChapterId(chapterId uint) error {
	return r.db.Where("chapter_id = ?", chapterId).Delete(&model.KnowledgePoint{}).Error
}

func (r *KnowledgePointRepository) FindByID(id uint) (*model.KnowledgePoint, error) {
	var kp model.KnowledgePoint
	err := r.db.Preload("Chapter").Preload("Chapter.Syllabus").First(&kp, id).Error
	return &kp, err
}

func (r *KnowledgePointRepository) FindByChapterId(chapterId uint) ([]model.KnowledgePoint, error) {
	var kps []model.KnowledgePoint
	err := r.db.Where("chapter_id = ?", chapterId).
		Preload("Chapter").
		Order("order_index ASC").
		Find(&kps).Error
	return kps, err
}

func (r *KnowledgePointRepository) FindBySyllabusId(syllabusId uint) ([]model.KnowledgePoint, error) {
	var kps []model.KnowledgePoint
	err := r.db.Joins("JOIN chapters ON chapters.id = knowledge_points.chapter_id").
		Where("chapters.syllabus_id = ?", syllabusId).
		Preload("Chapter").
		Order("knowledge_points.chapter_id ASC, knowledge_points.order_index ASC").
		Find(&kps).Error
	return kps, err
}

func (r *KnowledgePointRepository) FindAll(query *model.KnowledgePointQuery) ([]model.KnowledgePoint, int64, error) {
	var kps []model.KnowledgePoint
	var total int64

	db := r.db.Model(&model.KnowledgePoint{})

	if query.ChapterId != 0 {
		db = db.Where("chapter_id = ?", query.ChapterId)
	}

	if query.SyllabusId != 0 {
		db = db.Joins("JOIN chapters ON chapters.id = knowledge_points.chapter_id").
			Where("chapters.syllabus_id = ?", query.SyllabusId)
	}

	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}

	if query.Difficulty != "" {
		db = db.Where("difficulty = ?", query.Difficulty)
	}

	db.Count(&total)

	if query.PageSize > 0 {
		offset := (query.PageIndex - 1) * query.PageSize
		db = db.Offset(offset).Limit(query.PageSize)
	}

	err := db.Preload("Chapter").Order("chapter_id ASC, order_index ASC").Find(&kps).Error
	return kps, total, err
}

func (r *KnowledgePointRepository) BatchCreate(kps []model.KnowledgePoint) error {
	return r.db.Create(&kps).Error
}

func (r *KnowledgePointRepository) LinkQuestion(knowledgePointId uint, questionId uint) error {
	return r.db.Model(&model.KnowledgePoint{Model: model.Model{ID: knowledgePointId}}).
		Association("Questions").
		Append(&model.Question{Model: model.Model{ID: questionId}})
}

func (r *KnowledgePointRepository) UnlinkQuestion(knowledgePointId uint, questionId uint) error {
	return r.db.Model(&model.KnowledgePoint{Model: model.Model{ID: knowledgePointId}}).
		Association("Questions").
		Delete(&model.Question{Model: model.Model{ID: questionId}})
}
