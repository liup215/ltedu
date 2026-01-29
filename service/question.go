package service

import (
"edu/repository"
"edu/model"
"encoding/json"
"errors"
"math/rand"
"time"
)

var QuestionSvr = &QuestionService{baseService: newBaseService()}

type QuestionService struct {
baseService
}

func (svr *QuestionService) SelectQuestionAllCount(q model.QuestionQueryRequest) (int64, error) {
	return repository.QuestionRepo.Count(&q)
}

func (svr *QuestionService) SelectQuestionAll(q model.QuestionQueryRequest) ([]*model.Question, error) {
	list, err := repository.QuestionRepo.FindAll(&q)
	for _, q := range list {
		q.Format()
	}
	return list, err
}

func (svr *QuestionService) SelectQuestionList(q model.QuestionQueryRequest) ([]*model.Question, int64, error) {
	page := q.Page.CheckPage()
	list, total, err := repository.QuestionRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
	for _, q := range list {
		q.Format()
	}
	return list, total, err
}

func (svr *QuestionService) SelectQuestionById(id uint) (*model.Question, error) {
	if id == 0 {
		return nil, errors.New("ID不能为空")
	}
	q, err := repository.QuestionRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	q.Format()
	return q, nil
}

func (svr *QuestionService) CreateQuestion(q model.Question, uid uint) (uint, error) {
	if uid == 0 {
		return 0, errors.New("用户ID不能为空")
	}
	if q.ID != 0 {
		q.ID = 0
	}
	if q.Status == 0 {
		q.Status = model.QUESTION_STATE_FORBIDDEN
	}
	contentsStringBytes, err := json.Marshal(q.QuestionContents)
	if err != nil {
		return 0, err
	}
	q.QuestionContentsString = string(contentsStringBytes)
	e := repository.QuestionRepo.Create(&q)
	if e != nil {
		return 0, e
	}
	return q.ID, nil
}

func (svr *QuestionService) EditQuestion(q model.Question) error {
	if q.ID == 0 {
		return errors.New("无效的ID")
	}
	if q.Status == 0 {
		q.Status = model.QUESTION_STATE_FORBIDDEN
	}
	contentsStringBytes, err := json.Marshal(q.QuestionContents)
	if err != nil {
		return err
	}
	q.QuestionContentsString = string(contentsStringBytes)
	return repository.QuestionRepo.Update(&q)
}

func (svr *QuestionService) DeleteQuestion(id uint) error {
	if id == 0 {
		return errors.New("无效的ID")
	}
	return repository.QuestionRepo.Delete(id)
}

// 生成练习题
func (svr *QuestionService) GenerateQuestionExercise(query model.QuestionQueryRequest) (list []*model.Question, err error) {
total, err := repository.QuestionRepo.Count(&query)
if err != nil {
return
}
if total < 20 {
list, err = repository.QuestionRepo.FindAll(&query)
return
}
rand.New(rand.NewSource(time.Now().UnixNano()))
randNums := rand.Perm(int(total))
randNums = randNums[:20]
for _, num := range randNums {
questions, _, err := repository.QuestionRepo.FindPage(&query, num, 1)
if err != nil || len(questions) == 0 {
continue
}
q := questions[0]
err = q.Format()
if err != nil {
continue
}
list = append(list, q)
}
return
}

// 兼容 controller 空实现
func (svr *QuestionService) AddQuestionChapter(o interface{}) error {
return nil
}
func (svr *QuestionService) DeleteQuestionChapter(o interface{}) error {
return nil
}
