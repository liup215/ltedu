package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

// ICourseVideoRepository 课程视频数据访问接口
type ICourseVideoRepository interface {
	Create(video *model.CourseVideo) error
	Update(video *model.CourseVideo) error
	Delete(id uint) error
	FindByID(id uint) (*model.CourseVideo, error)
	FindPage(query *model.CourseVideoQueryRequest, offset, limit int) ([]*model.CourseVideo, int64, error)
}

type courseVideoRepository struct {
	db *gorm.DB
}

// NewCourseVideoRepository 创建课程视频仓储实例
func NewCourseVideoRepository(db *gorm.DB) ICourseVideoRepository {
	return &courseVideoRepository{db: db}
}

func (r *courseVideoRepository) Create(video *model.CourseVideo) error {
	return r.db.Create(video).Error
}

func (r *courseVideoRepository) Update(video *model.CourseVideo) error {
	return r.db.Save(video).Error
}

func (r *courseVideoRepository) Delete(id uint) error {
	return r.db.Delete(&model.CourseVideo{}, id).Error
}

func (r *courseVideoRepository) FindByID(id uint) (*model.CourseVideo, error) {
	var video model.CourseVideo
	err := r.db.Where("id = ?", id).First(&video).Error
	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return &video, err
}

func (r *courseVideoRepository) FindPage(query *model.CourseVideoQueryRequest, offset, limit int) ([]*model.CourseVideo, int64, error) {
	var videos []*model.CourseVideo
	var total int64

	q := r.db.Model(&model.CourseVideo{})

	if query.CourseID != 0 {
		q = q.Where("course_id = ?", query.CourseID)
	}

	var totalInt int64
	q.Count(&totalInt)
	total = totalInt

	err := q.Order("id DESC").Offset(offset).Limit(limit).Preload("Course").Find(&videos).Error
	return videos, total, err
}
