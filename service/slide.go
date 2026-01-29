package service

import (
"edu/repository"
"edu/model"
"errors"
)

var SlideSvr = &SlideService{baseService: newBaseService()}

type SlideService struct {
baseService
}

func (s *SlideService) SelectSlideById(id uint) (*model.Slide, error) {
if id == 0 {
return nil, errors.New("无效的ID.")
}
return repository.SlideRepo.FindByID(id)
}

func (s *SlideService) SelectSlideList(q model.SlideQueryRequest) ([]*model.Slide, int64, error) {
page := q.CheckPage()
list, total, err := repository.SlideRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
return list, total, err
}

func (s *SlideService) SelectSlideAll(q model.SlideQueryRequest) ([]*model.Slide, error) {
return repository.SlideRepo.FindAll(&q)
}

func (s *SlideService) CreateSlide(vs model.Slide) error {
if vs.Name == "" {
return errors.New("单词集名称不能为空！")
}
if vs.SyllabusId == 0 {
return errors.New("单词集考纲不能为空!")
}
return repository.SlideRepo.Create(&vs)
}

func (s *SlideService) EditSlide(vs model.SlideCreateEditRequest) error {
if vs.ID == 0 {
return errors.New("无效的ID.")
}
if vs.Name == "" {
return errors.New("单词集名称不能为空！")
}
if vs.SyllabusId == 0 {
return errors.New("单词集考纲不能为空!")
}
slide := model.Slide{
Model:       model.Model{ID: vs.ID},
Name:        vs.Name,
Description: vs.Description,
SyllabusId:  vs.SyllabusId,
}
err := repository.SlideRepo.Update(&slide)
return err
}

func (s *SlideService) DeleteSlide(id uint) error {
if id == 0 {
return errors.New("无效的ID.")
}
return repository.SlideRepo.Delete(id)
}
