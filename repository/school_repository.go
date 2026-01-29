package repository

import (
"edu/model"
"gorm.io/gorm"
)

// IGradeRepository 年级数据访问接口
type IGradeRepository interface {
Create(g *model.Grade) error
Update(g *model.Grade) error
Delete(id uint) error
FindByID(id uint) (*model.Grade, error)
FindPage(query *model.GradeQuery, offset, limit int) ([]*model.Grade, int64, error)
FindAll(query *model.GradeQuery) ([]*model.Grade, error)
}

type gradeRepository struct {
db *gorm.DB
}

func NewGradeRepository(db *gorm.DB) IGradeRepository {
return &gradeRepository{db: db}
}

func (r *gradeRepository) Create(g *model.Grade) error {
return r.db.Create(g).Error
}

func (r *gradeRepository) Update(g *model.Grade) error {
return r.db.Model(g).Updates(g).Error
}

func (r *gradeRepository) Delete(id uint) error {
return r.db.Delete(&model.Grade{}, id).Error
}

func (r *gradeRepository) FindByID(id uint) (*model.Grade, error) {
var g model.Grade
err := r.db.Where("id = ?", id).
Preload("GradeLeadTeacher").
First(&g).Error
if gorm.ErrRecordNotFound == err {
return nil, nil
}
return &g, err
}

func (r *gradeRepository) FindPage(query *model.GradeQuery, offset, limit int) ([]*model.Grade, int64, error) {
var grades []*model.Grade
var total int64

q := r.db.Model(&model.Grade{}).Preload("GradeLeadTeacher")

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.GradeTeacherId != 0 {
q = q.Where("grade_lead_teacher_id = ?", query.GradeTeacherId)
}

q.Count(&total)
err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&grades).Error
return grades, total, err
}

func (r *gradeRepository) FindAll(query *model.GradeQuery) ([]*model.Grade, error) {
var grades []*model.Grade
q := r.db.Model(&model.Grade{}).Preload("GradeLeadTeacher")

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.GradeTeacherId != 0 {
q = q.Where("grade_lead_teacher_id = ?", query.GradeTeacherId)
}

err := q.Order("id DESC").Find(&grades).Error
return grades, err
}
