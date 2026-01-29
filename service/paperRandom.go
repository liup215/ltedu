package service

import (
	"edu/repository"
	"edu/model"
	"errors"
)

func (s *QuestionPaperService) CreateRandomPaper(p model.RandomPaper) error {
	if p.ID != 0 {
		p.ID = uint(0)
	}

	return repository.RandomPaperRepo.Create(&p)
}

func (s *QuestionPaperService) SelectRandomPaperById(id uint) (*model.RandomPaperQueryResponse, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}

	p, err := repository.RandomPaperRepo.FindByID(id)
	if err != nil || p == nil {
		return nil, err
	}

	p.QuestionsInfo = s.queryQuestionByPaperId(p.ID)
	return p.GetRandomPaperResponse(), nil
}

func (s *QuestionPaperService) SelectRandomPaperList(q model.RandomPaperQuery) ([]*model.RandomPaperQueryResponse, int64, error) {
	page := q.Page.CheckPage()
	offset := (page.PageIndex - 1) * page.PageSize

pList, total, err := repository.RandomPaperRepo.FindPage(&q, offset, page.PageSize)

	list := []*model.RandomPaperQueryResponse{}
	for _, p := range pList {
		p.QuestionsInfo = s.queryQuestionByPaperId(p.ID)
		list = append(list, p.GetRandomPaperResponse())
	}

	return list, total, err
}

func (s *QuestionPaperService) queryQuestionByPaperId(paperId uint) []*model.QuestionRandomPapers {
	// TODO: 需要创建QuestionRandomPapersRepository
	return []*model.QuestionRandomPapers{}
}
