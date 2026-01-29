package service

import (
	"edu/model"
	"edu/repository"
	"errors"
	"fmt"
	"time"
)

var CourseSvr = &CourseService{baseService: newBaseService()}

type CourseService struct {
	baseService
}

// GetCourseList
func (s *CourseService) SelectCourseList(req model.CourseQueryRequest) ([]*model.Course, int64, error) {
	list, total, err := repository.CourseRepo.FindList(req)
	if err != nil {
		return nil, 0, err
	}

	// 设置显示格式
	for _, c := range list {
		if userType, ok := model.UserTypeMap[c.UserType]; ok {
			c.UserTypeName = userType
		}
		if c.PublishedAt != nil {
			c.PublishedAtString = c.PublishedAt.Format("2006-01-02 15:04:05")
		}
	}

	return list, total, nil
}

func (s *CourseService) SelectCourseById(id uint) (*model.Course, error) {
	if id == 0 {
		return nil, errors.New("id is required")
	}

	c, err := repository.CourseRepo.FindByID(id)
	if err != nil || c == nil {
		return nil, err
	}

	// 设置显示格式
	if userType, ok := model.UserTypeMap[c.UserType]; ok {
		c.UserTypeName = userType
	}
	if c.PublishedAt != nil {
		c.PublishedAtString = c.PublishedAt.Format("2006-01-02 15:04:05")
	}

	return c, nil
}

func (s *CourseService) CreateCourse(course model.Course) error {
	return repository.CourseRepo.Create(&course)
}

func (s *CourseService) EditCourse(course model.Course) error {
	if course.ID == 0 {
		return errors.New("id is required")
	}

	existing, err := repository.CourseRepo.FindByID(course.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("id can not be found")
	}

	fmt.Println("---------------\n\t", course.PublishedAtString)

	t, err := time.Parse("2006-01-02 15:04:05", course.PublishedAtString)
	if err != nil {
		return errors.New("time parse error: " + course.PublishedAtString + ":" + err.Error())
	}

	// 更新字段
	existing.Title = course.Title
	existing.Thumb = course.Thumb
	existing.IsFree = course.IsFree
	existing.PublishedAt = &t
	existing.IsShow = course.IsShow
	existing.ShortDescription = course.ShortDescription
	existing.OriginalDesc = course.OriginalDesc
	existing.Charge = course.Charge
	existing.RenderDesc = course.RenderDesc

	return repository.CourseRepo.Update(existing)
}

func (s *CourseService) DeleteCourse(id uint) error {
	if id == 0 {
		return errors.New("id is required")
	}
	return repository.CourseRepo.Delete(id)
}
