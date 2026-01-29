package service

import (
	"edu/model"
	"edu/repository"
	"errors"
)

var QuestionPaperSvr = &QuestionPaperService{baseService: newBaseService()}

type QuestionPaperService struct {
	baseService
}

func (s *QuestionPaperService) CreateSeries(se model.PaperSeries) error {
	if se.ID != uint(0) {
		se.ID = uint(0)
	}
	return repository.PaperSeriesRepo.Create(&se)
}

func (s *QuestionPaperService) EditSeries(se model.PaperSeries) error {
	if se.ID == uint(0) {
		return errors.New("无效的ID")
	}
	return repository.PaperSeriesRepo.Update(&se)
}

func (s *QuestionPaperService) DeleteSeries(id uint) error {
	if id == 0 {
		return errors.New("id不能为空!")
	}
	return repository.PaperSeriesRepo.Delete(id)
}

func (s *QuestionPaperService) SelectSeriesById(id uint) (*model.PaperSeries, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}
	return repository.PaperSeriesRepo.FindByID(id)
}

func (s *QuestionPaperService) SelectSeriesList(q model.PaperSeriesQuery) ([]*model.PaperSeries, int64, error) {
	return repository.PaperSeriesRepo.FindList(q)
}

func (s *QuestionPaperService) SelectSeriesAll(q model.PaperSeriesQuery) ([]*model.PaperSeries, error) {
	return repository.PaperSeriesRepo.FindAll(q)
}

func (s *QuestionPaperService) CreateCode(c model.PaperCode) error {
	if c.ID != uint(0) {
		c.ID = uint(0)
	}
	return repository.PaperCodeRepo.Create(&c)
}

func (s *QuestionPaperService) EditCode(c model.PaperCode) error {
	if c.ID == uint(0) {
		return errors.New("无效的ID")
	}
	return repository.PaperCodeRepo.Update(&c)
}

func (s *QuestionPaperService) SelectCodeById(id uint) (*model.PaperCode, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}
	return repository.PaperCodeRepo.FindByID(id)
}

func (s *QuestionPaperService) SelectCodeList(q model.PaperCodeQuery) ([]*model.PaperCode, int64, error) {
	return repository.PaperCodeRepo.FindList(q)
}

func (s *QuestionPaperService) SelectCodeAll(q model.PaperCodeQuery) ([]*model.PaperCode, error) {
	return repository.PaperCodeRepo.FindAll(q)
}

func (s *QuestionPaperService) DeleteCode(id uint) error {
	if id == 0 {
		return errors.New("id不能为空!")
	}
	return repository.PaperCodeRepo.Delete(id)
}
