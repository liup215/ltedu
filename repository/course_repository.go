package repository

import (
	"errors"

	"edu/model"
	"gorm.io/gorm"
)

// ICourseRepository 课程数据访问接口
type ICourseRepository interface {
	Create(course *model.Course) error
	Update(course *model.Course) error
	Delete(id uint) error
	FindByID(id uint) (*model.Course, error)
	FindList(query model.CourseQueryRequest) ([]*model.Course, int64, error)
	FindBySyllabusID(syllabusID int) ([]*model.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

// NewCourseRepository 创建课程仓储实例
func NewCourseRepository(db *gorm.DB) ICourseRepository {
	return &courseRepository{db: db}
}

// Create 创建课程
func (r *courseRepository) Create(course *model.Course) error {
	return r.db.Create(course).Error
}

// Update 更新课程
func (r *courseRepository) Update(course *model.Course) error {
	return r.db.Save(course).Error
}

// Delete 删除课程
func (r *courseRepository) Delete(id uint) error {
	return r.db.Delete(&model.Course{}, id).Error
}

// FindByID 根据ID查询课程
func (r *courseRepository) FindByID(id uint) (*model.Course, error) {
	var course model.Course
	err := r.db.Where("id = ?", id).
		Preload("Syllabus").
		First(&course).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &course, err
}

// FindList 分页查询课程
func (r *courseRepository) FindList(query model.CourseQueryRequest) ([]*model.Course, int64, error) {
	var courses []*model.Course
	var total int64

	q := r.db.Model(&model.Course{})

	if query.SyllabusID != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusID)
	}

	// 统计总数
	q.Count(&total)

	// 分页查询
	offset := (query.PageIndex - 1) * query.PageSize
	err := q.
		Order("id DESC").
		Offset(offset).
		Limit(query.PageSize).
		Preload("Syllabus").
		Find(&courses).Error

	// 设置显示格式（移除，由Service层处理）

	return courses, total, err
}

// FindBySyllabusID 根据大纲ID查询所有课程
func (r *courseRepository) FindBySyllabusID(syllabusID int) ([]*model.Course, error) {
	var courses []*model.Course
	err := r.db.Where("syllabus_id = ?", syllabusID).
		Order("id DESC").
		Preload("Syllabus").
		Find(&courses).Error

	return courses, err
}
