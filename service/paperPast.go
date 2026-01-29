package service

import (
	"edu/repository"
	"edu/model"
	"errors"
)

func (s *QuestionPaperService) CreatePastPaper(p model.PastPaper) (*model.PastPaper, error) {
	if p.ID != 0 {
		p.ID = uint(0)
	}

	err := repository.PastPaperRepo.Create(&p)
	return &p, err
}

func (s *QuestionPaperService) EditPastPaper(p model.PastPaper) error {
	if p.ID == 0 {
		return errors.New("无效的ID")
	}

err := repository.PastPaperRepo.Update(&p)
return err
}

func (s *QuestionPaperService) DeletePastPaper(id uint) error {
	if id == 0 {
		return errors.New("无效的ID")
	}

	err := repository.PastPaperRepo.Delete(id)
	return err
}

func (s *QuestionPaperService) SelectPastPaperById(id uint) (*model.PastPaper, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}

	p, err := repository.PastPaperRepo.FindByID(id)
	if p == nil || err != nil {
		return nil, err
	}

	// p.Questions = s.queryQuestionByPastPaperId(p.ID)
	return p, nil
}

func (s *QuestionPaperService) SelectPastPaperList(q model.PastPaperQuery) ([]*model.PastPaper, int64, error) {
	page := q.Page.CheckPage()
	offset := (page.PageIndex - 1) * page.PageSize

	return repository.PastPaperRepo.FindPage(&q, offset, page.PageSize)
}

func (s *QuestionPaperService) SelectPastPaperAll(q model.PastPaperQuery) ([]*model.PastPaper, error) {
	return repository.PastPaperRepo.FindAll(&q)
}

func (s *QuestionPaperService) queryQuestionByPastPaperId(paperId uint) []*model.Question {
	// TODO: 需要创建QuestionRepository
	list := []*model.Question{}
	// list, _ = repository.QuestionRepo.FindByPastPaperId(paperId)
	// for _, q := range list {
	// 	q.Format()
	// }
	return list
}
