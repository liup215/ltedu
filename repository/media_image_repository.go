package repository

import (
"edu/model"
"gorm.io/gorm"
)

// IMediaImageRepository 图片数据访问接口
type IMediaImageRepository interface {
Create(img *model.MediaImage) error
Update(img *model.MediaImage) error
Delete(id uint) error
FindByID(id uint) (*model.MediaImage, error)
FindPage(query *model.ImageQueryRequest, offset, limit int) ([]*model.MediaImage, int64, error)
FindAll(query *model.ImageQueryRequest) ([]*model.MediaImage, error)
FindByName(name string) (*model.MediaImage, error)
}

type mediaImageRepository struct {
db *gorm.DB
}

func NewMediaImageRepository(db *gorm.DB) IMediaImageRepository {
return &mediaImageRepository{db: db}
}

func (r *mediaImageRepository) Create(img *model.MediaImage) error {
return r.db.Create(img).Error
}

func (r *mediaImageRepository) Update(img *model.MediaImage) error {
return r.db.Model(img).Updates(img).Error
}

func (r *mediaImageRepository) Delete(id uint) error {
return r.db.Delete(&model.MediaImage{}, id).Error
}

func (r *mediaImageRepository) FindByID(id uint) (*model.MediaImage, error) {
var img model.MediaImage
err := r.db.Where("id = ?", id).
Preload("Attachment").
First(&img).Error
if gorm.ErrRecordNotFound == err {
return nil, nil
}
return &img, err
}

func (r *mediaImageRepository) FindPage(query *model.ImageQueryRequest, offset, limit int) ([]*model.MediaImage, int64, error) {
var imgs []*model.MediaImage
var total int64

q := r.db.Model(&model.MediaImage{}).Preload("Attachment")

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.Name != "" {
q = q.Where("name LIKE ?", "%"+query.Name+"%")
}
if query.Hash != "" {
q = q.Where("hash = ?", query.Hash)
}
if query.Disk != "" {
q = q.Where("disk = ?", query.Disk)
}

q.Count(&total)
err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&imgs).Error
return imgs, total, err
}

func (r *mediaImageRepository) FindAll(query *model.ImageQueryRequest) ([]*model.MediaImage, error) {
var imgs []*model.MediaImage
q := r.db.Model(&model.MediaImage{}).Preload("Attachment")

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.Name != "" {
q = q.Where("name LIKE ?", "%"+query.Name+"%")
}
if query.Hash != "" {
q = q.Where("hash = ?", query.Hash)
}
if query.Disk != "" {
q = q.Where("disk = ?", query.Disk)
}

err := q.Order("id DESC").Find(&imgs).Error
return imgs, err
}

func (r *mediaImageRepository) FindByName(name string) (*model.MediaImage, error) {
var img model.MediaImage
err := r.db.Where("name = ?", name).
Preload("Attachment").
First(&img).Error
if gorm.ErrRecordNotFound == err {
return nil, nil
}
return &img, err
}
