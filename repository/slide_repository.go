package repository

import (
"edu/model"
"gorm.io/gorm"
)

// ISlideRepository 幻灯片数据访问接口
type ISlideRepository interface {
Create(s *model.Slide) error
Update(s *model.Slide) error
Delete(id uint) error
FindByID(id uint) (*model.Slide, error)
FindPage(query *model.SlideQueryRequest, offset, limit int) ([]*model.Slide, int64, error)
FindAll(query *model.SlideQueryRequest) ([]*model.Slide, error)
}

type slideRepository struct {
db *gorm.DB
}

func NewSlideRepository(db *gorm.DB) ISlideRepository {
return &slideRepository{db: db}
}

func (r *slideRepository) Create(s *model.Slide) error {
return r.db.Create(s).Error
}

func (r *slideRepository) Update(s *model.Slide) error {
return r.db.Model(s).Updates(s).Error
}

func (r *slideRepository) Delete(id uint) error {
return r.db.Delete(&model.Slide{}, id).Error
}

func (r *slideRepository) FindByID(id uint) (*model.Slide, error) {
var slide model.Slide
err := r.db.Where("id = ?", id).First(&slide).Error
if gorm.ErrRecordNotFound == err {
return nil, nil
}
return &slide, err
}

func (r *slideRepository) FindPage(query *model.SlideQueryRequest, offset, limit int) ([]*model.Slide, int64, error) {
var slides []*model.Slide
var total int64

q := r.db.Model(&model.Slide{})

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.SyllabusId != 0 {
q = q.Where("syllabus_id = ?", query.SyllabusId)
}

q.Count(&total)
err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&slides).Error
return slides, total, err
}

func (r *slideRepository) FindAll(query *model.SlideQueryRequest) ([]*model.Slide, error) {
var slides []*model.Slide
q := r.db.Model(&model.Slide{})

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.SyllabusId != 0 {
q = q.Where("syllabus_id = ?", query.SyllabusId)
}

err := q.Order("id DESC").Find(&slides).Error
return slides, err
}
