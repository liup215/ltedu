package service

import (
"edu/repository"
"edu/model"
"errors"
)

var CourseVideoSvr = &CourseVideoService{baseService: newBaseService()}

type CourseVideoService struct {
	baseService
}

func (s *CourseVideoService) SelectCourseVideoList(req model.CourseVideoQueryRequest) ([]*model.CourseVideo, int64, error) {
	offset := (req.PageIndex - 1) * req.PageSize
list, total, err := repository.CourseVideoRepo.FindPage(&req, offset, req.PageSize)

	if err != nil {
		return nil, 0, err
	}

	for _, c := range list {
		if c.PublishedAt != nil {
			c.PublishedAtString = c.PublishedAt.Format("2006-01-02 15:04:05")
		}
	}

	return list, total, nil
}

func (s *CourseVideoService) SelectCourseVideoById(id uint) (*model.CourseVideo, error) {
	if id == 0 {
		return nil, errors.New("id is required")
	}

	c, err := repository.CourseVideoRepo.FindByID(id)
	if err != nil || c == nil {
		return nil, err
	}

	if c.PublishedAt != nil {
		c.PublishedAtString = c.PublishedAt.Format("2006-01-02 15:04:05")
	}
	return c, nil
}

func (s *CourseVideoService) CreateCourseVideo(course model.CourseVideo) error {
	return repository.CourseVideoRepo.Create(&course)
}

func (s *CourseVideoService) EditCourseVideo(courseVideo model.CourseVideo) error {
	if courseVideo.ID == 0 {
		return errors.New("id is required")
	}

_, err := repository.CourseVideoRepo.FindByID(courseVideo.ID)
if err != nil {
return err
}

err = repository.CourseVideoRepo.Update(&courseVideo)
return err
}

func (s *CourseVideoService) DeleteCourseVideo(id uint) error {
	if id == 0 {
		return errors.New("id is required")
	}

	return repository.CourseVideoRepo.Delete(id)
}
