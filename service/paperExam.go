package service

import (
	"edu/repository"
	"edu/model"
	"encoding/json"
	"errors"
)

func (s *QuestionPaperService) CreateExamPaper(p model.ExamPaper) (*model.ExamPaper, error) {
	if p.ID != 0 {
		p.ID = uint(0)
	}

	qIdByte, err := json.Marshal(p.QuestionIds)
	if err != nil {
		return nil, errors.New("Failed to parse questions")
	}

	p.QuestionIdsStr = string(qIdByte)

	err = repository.ExamPaperRepo.Create(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (s *QuestionPaperService) EditExamPaper(p model.ExamPaper) error {
	if p.ID == 0 {
		return errors.New("Invalid ID")
	}

	qIdByte, err := json.Marshal(p.QuestionIds)
	if err != nil {
		return errors.New("Failed to parse questions")
	}

	p.QuestionIdsStr = string(qIdByte)

err = repository.ExamPaperRepo.Update(&p)
return err
}

func (s *QuestionPaperService) DeleteExamPaper(id uint) error {
	if id == 0 {
		return errors.New("无效的ID")
	}

	err := repository.ExamPaperRepo.Delete(id)
	return err
}

func (s *QuestionPaperService) SelectExamPaperById(id uint) (*model.ExamPaper, error) {
	if id == 0 {
		return nil, errors.New("Invalid ID")
	}

	p, err := repository.ExamPaperRepo.FindByID(id)
	if err != nil || p == nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(p.QuestionIdsStr), &p.QuestionIds)
	if err != nil {
		return nil, errors.New("Failed to parse questions")
	}

	p.Questions = s.queryQuestionByIds(p.QuestionIds)
	return p, nil
}

func (s *QuestionPaperService) SelectExamPaperList(q model.ExamPaperQuery) ([]*model.ExamPaper, int64, error) {
	page := q.Page.CheckPage()
	offset := (page.PageIndex - 1) * page.PageSize

	pList, total, err := repository.ExamPaperRepo.FindPage(&q, offset, page.PageSize)

	for i, p := range pList {
		err = json.Unmarshal([]byte(p.QuestionIdsStr), &p.QuestionIds)
		if err != nil {
			continue
		}
		pList[i].Questions = s.queryQuestionByIds(p.QuestionIds)
	}

	return pList, total, err
}

func (s *QuestionPaperService) SelectExamPaperAll(q model.ExamPaperQuery) ([]*model.ExamPaper, error) {
	pList, err := repository.ExamPaperRepo.FindAll(&q)

	for i, p := range pList {
		err = json.Unmarshal([]byte(p.QuestionIdsStr), &p.QuestionIds)
		if err != nil {
			continue
		}
		pList[i].Questions = s.queryQuestionByIds(p.QuestionIds)
	}

	return pList, err
}

func (s *QuestionPaperService) queryQuestionByIds(qIds []uint) []*model.Question {
	// TODO: 需要创建QuestionRepository
	// list, _ := repository.QuestionRepo.FindByIds(qIds)
	// for _, q := range list {
	// 	q.Format()
	// }

	// // reorder the list based on the original order of qIds
	// idMap := make(map[uint]*model.Question)
	// for _, q := range list {
	// 	idMap[q.ID] = q
	// }
	// var orderedList []*model.Question
	// for _, id := range qIds {
	// 	if q, exists := idMap[id]; exists {
	// 		orderedList = append(orderedList, q)
	// 	}
	// }

	// return orderedList

	// Placeholder for now
	return []*model.Question{}
}

func (svr *QuestionPaperService) UpdateExamPaperQuestion(p model.ExamPaper) error {
	if p.ID == 0 {
		return errors.New("Invalid ID")
	}

	qIdByte, err := json.Marshal(p.QuestionIds)
	if err != nil {
		return errors.New("Failed to parse questions")
	}

	p.QuestionIdsStr = string(qIdByte)
	return repository.ExamPaperRepo.Update(&p)
}
