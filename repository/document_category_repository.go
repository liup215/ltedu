package repository

import (
"edu/model"
"gorm.io/gorm"
)

// IDocumentCategoryRepository 文档分类数据访问接口
type IDocumentCategoryRepository interface {
Create(cat *model.DocumentCategory) error
Update(cat *model.DocumentCategory) error
Delete(id uint) error
FindByID(id uint) (*model.DocumentCategory, error)
FindByParentID(pid uint) ([]*model.DocumentCategory, error)
CountByParentID(pid uint) (int64, error)
FindPage(offset, limit int) ([]*model.DocumentCategory, int64, error)
FindAll() ([]*model.DocumentCategory, error)
}

type documentCategoryRepository struct {
db *gorm.DB
}

func NewDocumentCategoryRepository(db *gorm.DB) IDocumentCategoryRepository {
return &documentCategoryRepository{db: db}
}

func (r *documentCategoryRepository) Create(cat *model.DocumentCategory) error {
return r.db.Create(cat).Error
}

func (r *documentCategoryRepository) Update(cat *model.DocumentCategory) error {
return r.db.Model(cat).Updates(cat).Error
}

func (r *documentCategoryRepository) Delete(id uint) error {
return r.db.Delete(&model.DocumentCategory{}, id).Error
}

func (r *documentCategoryRepository) FindByID(id uint) (*model.DocumentCategory, error) {
var cat model.DocumentCategory
err := r.db.Where("id = ?", id).First(&cat).Error
if gorm.ErrRecordNotFound == err {
return nil, nil
}
return &cat, err
}

func (r *documentCategoryRepository) FindByParentID(pid uint) ([]*model.DocumentCategory, error) {
var cats []*model.DocumentCategory
err := r.db.Where("parent_id = ?", pid).Find(&cats).Error
return cats, err
}

func (r *documentCategoryRepository) CountByParentID(pid uint) (int64, error) {
var count int64
err := r.db.Model(&model.DocumentCategory{}).Where("parent_id = ?", pid).Count(&count).Error
return count, err
}

func (r *documentCategoryRepository) FindPage(offset, limit int) ([]*model.DocumentCategory, int64, error) {
var cats []*model.DocumentCategory
var total int64
q := r.db.Model(&model.DocumentCategory{})
q.Count(&total)
err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&cats).Error
return cats, total, err
}

func (r *documentCategoryRepository) FindAll() ([]*model.DocumentCategory, error) {
var cats []*model.DocumentCategory
err := r.db.Order("id DESC").Find(&cats).Error
return cats, err
}
