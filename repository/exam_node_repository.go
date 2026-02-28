package repository

import (
	"edu/model"

	"gorm.io/gorm"
)

// IExamNodeRepository 考试节点数据访问接口
type IExamNodeRepository interface {
	Create(node *model.SyllabusExamNode) error
	Update(node *model.SyllabusExamNode) error
	Delete(id uint) error
	FindByID(id uint) (*model.SyllabusExamNode, error)
	FindBySyllabusID(syllabusId uint) ([]*model.SyllabusExamNode, error)
	FindPage(query *model.SyllabusExamNodeQuery, offset, limit int) ([]*model.SyllabusExamNode, int64, error)
	// Chapter management
	AddChapters(examNodeId uint, chapterIds []uint) error
	RemoveChapter(examNodeId uint, chapterId uint) error
	// Paper code management
	AddPaperCode(examNodeId uint, paperCodeId uint) error
	RemovePaperCode(examNodeId uint, paperCodeId uint) error
}

type examNodeRepository struct {
	db *gorm.DB
}

func NewExamNodeRepository(db *gorm.DB) IExamNodeRepository {
	return &examNodeRepository{db: db}
}

func (r *examNodeRepository) Create(node *model.SyllabusExamNode) error {
	return r.db.Create(node).Error
}

func (r *examNodeRepository) Update(node *model.SyllabusExamNode) error {
	return r.db.Model(node).Select("name", "description", "sort_order").Updates(node).Error
}

func (r *examNodeRepository) Delete(id uint) error {
	return r.db.Delete(&model.SyllabusExamNode{}, id).Error
}

func (r *examNodeRepository) FindByID(id uint) (*model.SyllabusExamNode, error) {
	var node model.SyllabusExamNode
	err := r.db.Where("id = ?", id).
		Preload("Syllabus").
		Preload("Chapters").
		Preload("PaperCodes").
		First(&node).Error
	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return &node, err
}

func (r *examNodeRepository) FindBySyllabusID(syllabusId uint) ([]*model.SyllabusExamNode, error) {
	var nodes []*model.SyllabusExamNode
	err := r.db.Where("syllabus_id = ?", syllabusId).
		Preload("Chapters").
		Preload("PaperCodes").
		Order("sort_order ASC, id ASC").
		Find(&nodes).Error
	return nodes, err
}

func (r *examNodeRepository) FindPage(query *model.SyllabusExamNodeQuery, offset, limit int) ([]*model.SyllabusExamNode, int64, error) {
	var nodes []*model.SyllabusExamNode
	var total int64

	q := r.db.Model(&model.SyllabusExamNode{})
	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}

	q.Count(&total)
	err := q.
		Preload("Chapters").
		Preload("PaperCodes").
		Order("sort_order ASC, id ASC").
		Offset(offset).
		Limit(limit).
		Find(&nodes).Error
	return nodes, total, err
}

func (r *examNodeRepository) AddChapters(examNodeId uint, chapterIds []uint) error {
	if len(chapterIds) == 0 {
		return nil
	}
	// Load existing associations to avoid duplicate entries in the join table
	var node model.SyllabusExamNode
	if err := r.db.Where("id = ?", examNodeId).Preload("Chapters").First(&node).Error; err != nil {
		return err
	}
	existing := make(map[uint]struct{}, len(node.Chapters))
	for _, ch := range node.Chapters {
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
	return r.db.Model(&model.SyllabusExamNode{Model: model.Model{ID: examNodeId}}).
		Association("Chapters").
		Append(toAdd)
}

func (r *examNodeRepository) RemoveChapter(examNodeId uint, chapterId uint) error {
	return r.db.Model(&model.SyllabusExamNode{Model: model.Model{ID: examNodeId}}).
		Association("Chapters").
		Delete(&model.Chapter{Model: model.Model{ID: chapterId}})
}

func (r *examNodeRepository) AddPaperCode(examNodeId uint, paperCodeId uint) error {
	if err := r.db.Model(&model.SyllabusExamNode{Model: model.Model{ID: examNodeId}}).
		Association("PaperCodes").
		Append(&model.PaperCode{Model: model.Model{ID: paperCodeId}}); err != nil {
		return err
	}
	return r.db.Model(&model.PaperCode{}).
		Where("id = ?", paperCodeId).
		Update("exam_node_id", examNodeId).Error
}

func (r *examNodeRepository) RemovePaperCode(examNodeId uint, paperCodeId uint) error {
	if err := r.db.Model(&model.SyllabusExamNode{Model: model.Model{ID: examNodeId}}).
		Association("PaperCodes").
		Delete(&model.PaperCode{Model: model.Model{ID: paperCodeId}}); err != nil {
		return err
	}
	return r.db.Model(&model.PaperCode{}).
		Where("id = ?", paperCodeId).
		Update("exam_node_id", 0).Error
}
