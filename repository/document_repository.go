package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

// IDocumentRepository 文档数据访问接口
type IDocumentRepository interface {
	Create(doc *model.Document) error
	Update(doc *model.Document) error
	Delete(id uint) error
	FindByID(id uint) (*model.Document, error)
	FindByPage(query *model.DocumentQueryRequest, offset, limit int) ([]*model.Document, int64, error)
	FindAll(query *model.DocumentQueryRequest) ([]*model.Document, error)
	FindFirst(query *model.DocumentQueryRequest) (*model.Document, error)
	UpdateStatus(id uint, status int) error
	UpdateFields(id uint, fields map[string]interface{}) error
}

type documentRepository struct {
	db *gorm.DB
}

// NewDocumentRepository 创建文档仓储实例
func NewDocumentRepository(db *gorm.DB) IDocumentRepository {
	return &documentRepository{db: db}
}

func (r *documentRepository) Create(doc *model.Document) error {
	return r.db.Omit("Syllabus", "DocumentCategory", "Attachment").Create(doc).Error
}

func (r *documentRepository) Update(doc *model.Document) error {
	return r.db.Model(doc).Updates(doc).Error
}

func (r *documentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Document{}, id).Error
}

func (r *documentRepository) FindByID(id uint) (*model.Document, error) {
	var doc model.Document
	err := r.db.Where("id = ?", id).
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		First(&doc).Error
	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return &doc, err
}

func (r *documentRepository) FindByPage(query *model.DocumentQueryRequest, offset, limit int) ([]*model.Document, int64, error) {
	var docs []*model.Document
	var total int64

	q := r.db.Model(&model.Document{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.DocumentCategoryId != 0 {
		q = q.Where("document_category_id = ?", query.DocumentCategoryId)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}

	q.Count(&total)

	err := q.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Order("id DESC").
		Offset(offset).
		Limit(limit).
		Find(&docs).Error

	return docs, total, err
}

func (r *documentRepository) FindAll(query *model.DocumentQueryRequest) ([]*model.Document, error) {
	var docs []*model.Document

	q := r.db.Model(&model.Document{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.DocumentCategoryId != 0 {
		q = q.Where("document_category_id = ?", query.DocumentCategoryId)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}

	err := q.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Order("id DESC").
		Find(&docs).Error

	return docs, err
}

func (r *documentRepository) FindFirst(query *model.DocumentQueryRequest) (*model.Document, error) {
	var doc model.Document

	q := r.db.Model(&model.Document{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.DocumentCategoryId != 0 {
		q = q.Where("document_category_id = ?", query.DocumentCategoryId)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}

	err := q.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Order("id DESC").
		First(&doc).Error

	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return &doc, err
}

func (r *documentRepository) UpdateStatus(id uint, status int) error {
	return r.db.Model(&model.Document{}).Where("id = ?", id).Update("status", status).Error
}

func (r *documentRepository) UpdateFields(id uint, fields map[string]interface{}) error {
	return r.db.Model(&model.Document{}).Where("id = ?", id).Updates(fields).Error
}
