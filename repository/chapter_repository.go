package repository

import (
	"edu/model"

	"gorm.io/gorm"
)

// IChapterRepository 章节数据访问接口
type IChapterRepository interface {
	Create(c *model.Chapter) error
	Update(c *model.Chapter) error
	Delete(id uint) error
	FindByID(id uint) (*model.Chapter, error)
	FindByParentID(parentId uint) ([]*model.Chapter, error)
	FindBySyllabusID(syllabusId uint) ([]*model.Chapter, error)
	FindPage(query *model.ChapterQuery, offset, limit int) ([]*model.Chapter, int64, error)
	FindAll(query *model.ChapterQuery) ([]*model.Chapter, error)
}

type chapterRepository struct {
	db *gorm.DB
}

func NewChapterRepository(db *gorm.DB) IChapterRepository {
	return &chapterRepository{db: db}
}

func (r *chapterRepository) Create(c *model.Chapter) error {
	return r.db.Create(c).Error
}

func (r *chapterRepository) Update(c *model.Chapter) error {
	return r.db.Model(c).Updates(c).Error
}

func (r *chapterRepository) Delete(id uint) error {
	return r.db.Delete(&model.Chapter{}, id).Error
}

func (r *chapterRepository) FindByID(id uint) (*model.Chapter, error) {
	var ch model.Chapter
	err := r.db.Where("id = ?", id).
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		First(&ch).Error
	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return &ch, err
}

func (r *chapterRepository) FindByParentID(parentId uint) ([]*model.Chapter, error) {
	var chapters []*model.Chapter
	err := r.db.Where("parent_id = ?", parentId).
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Find(&chapters).Error
	return chapters, err
}

func (r *chapterRepository) FindBySyllabusID(syllabusId uint) ([]*model.Chapter, error) {
	var chapters []*model.Chapter
	err := r.db.Where("syllabus_id = ?", syllabusId).
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Find(&chapters).Error
	return chapters, err
}

func (r *chapterRepository) FindPage(query *model.ChapterQuery, offset, limit int) ([]*model.Chapter, int64, error) {
	var chapters []*model.Chapter
	var total int64

	q := r.db.Model(&model.Chapter{})

	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.FilterRoot {
		q = q.Where("parent_id = ?", 0)
	} else if query.ParentId != 0 {
		q = q.Where("parent_id = ?", query.ParentId)
	}

	q.Count(&total)
	err := q.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Order("id DESC").
		Offset(offset).
		Limit(limit).
		Find(&chapters).Error
	return chapters, total, err
}

func (r *chapterRepository) FindAll(query *model.ChapterQuery) ([]*model.Chapter, error) {
	var chapters []*model.Chapter
	q := r.db.Model(&model.Chapter{})

	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.ParentId != 0 {
		q = q.Where("parent_id = ?", query.ParentId)
	}

	err := q.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Order("id DESC").
		Find(&chapters).Error
	return chapters, err
}
