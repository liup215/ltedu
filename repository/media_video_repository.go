package repository

import (
"edu/model"
"gorm.io/gorm"
)

// IMediaVideoRepository 视频数据访问接口
type IMediaVideoRepository interface {
Create(v *model.MediaVideo) error
Update(v *model.MediaVideo) error
Delete(id uint) error
FindByID(id uint) (*model.MediaVideo, error)
FindPage(query *model.VideoQueryRequest, offset, limit int) ([]*model.MediaVideo, int64, error)
FindAll(query *model.VideoQueryRequest) ([]*model.MediaVideo, error)
FindByName(name string) (*model.MediaVideo, error)
}

type mediaVideoRepository struct {
db *gorm.DB
}

func NewMediaVideoRepository(db *gorm.DB) IMediaVideoRepository {
return &mediaVideoRepository{db: db}
}

func (r *mediaVideoRepository) Create(v *model.MediaVideo) error {
return r.db.Create(v).Error
}

func (r *mediaVideoRepository) Update(v *model.MediaVideo) error {
return r.db.Model(v).Updates(v).Error
}

func (r *mediaVideoRepository) Delete(id uint) error {
return r.db.Delete(&model.MediaVideo{}, id).Error
}

func (r *mediaVideoRepository) FindByID(id uint) (*model.MediaVideo, error) {
var video model.MediaVideo
err := r.db.Where("id = ?", id).First(&video).Error
if gorm.ErrRecordNotFound == err {
return nil, nil
}
return &video, err
}

func (r *mediaVideoRepository) FindPage(query *model.VideoQueryRequest, offset, limit int) ([]*model.MediaVideo, int64, error) {
var videos []*model.MediaVideo
var total int64

q := r.db.Model(&model.MediaVideo{})

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
err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&videos).Error
return videos, total, err
}

func (r *mediaVideoRepository) FindAll(query *model.VideoQueryRequest) ([]*model.MediaVideo, error) {
var videos []*model.MediaVideo
q := r.db.Model(&model.MediaVideo{})

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

err := q.Order("id DESC").Find(&videos).Error
return videos, err
}

func (r *mediaVideoRepository) FindByName(name string) (*model.MediaVideo, error) {
var video model.MediaVideo
err := r.db.Where("name = ?", name).First(&video).Error
if gorm.ErrRecordNotFound == err {
return nil, nil
}
return &video, err
}
