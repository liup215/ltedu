package service

import (
"edu/repository"
"edu/model"
"errors"
"math/rand"
"time"
)

var VocabularySvr = &VocabularyService{baseService: newBaseService()}

type VocabularyService struct {
baseService
}

func (s *VocabularyService) SelectVocabularySetById(id uint) (*model.VocabularySet, error) {
if id == 0 {
return nil, errors.New("无效的ID.")
}
return repository.VocabularySetRepo.FindByID(id)
}

func (s *VocabularyService) SelectVocabularySetList(q model.VocabularySetQuery) ([]*model.VocabularySet, int64, error) {
page := q.CheckPage()
list, total, err := repository.VocabularySetRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
return list, total, err
}

func (s *VocabularyService) SelectVocabularySetAll(q model.VocabularySetQuery) ([]*model.VocabularySet, error) {
return repository.VocabularySetRepo.FindAll(&q)
}

func (s *VocabularyService) CreateVocabularySet(vs model.VocabularySetCreateEditRequest) error {
if vs.Name == "" {
return errors.New("单词集名称不能为空！")
}
if vs.SyllabusId == 0 {
return errors.New("单词集考纲不能为空!")
}
set := model.VocabularySet{
Name:        vs.Name,
Description: vs.Description,
SyllabusId:  vs.SyllabusId,
}
return repository.VocabularySetRepo.Create(&set)
}

func (s *VocabularyService) EditVocabularySet(vs model.VocabularySetCreateEditRequest) error {
if vs.ID == 0 {
return errors.New("无效的ID.")
}
if vs.Name == "" {
return errors.New("单词集名称不能为空！")
}
if vs.SyllabusId == 0 {
return errors.New("单词集考纲不能为空!")
}
set := model.VocabularySet{
Model:       model.Model{ID: vs.ID},
Name:        vs.Name,
Description: vs.Description,
SyllabusId:  vs.SyllabusId,
}
err := repository.VocabularySetRepo.Update(&set)
return err
}

func (s *VocabularyService) DeleteVocabularySet(id uint) error {
if id == 0 {
return errors.New("无效的ID.")
}
return repository.VocabularySetRepo.Delete(id)
}

func (s *VocabularyService) InsertVocabularyItem(vc model.VocabularyItemCreateEditRequest) error {
if vc.VocabularySetId == 0 {
return errors.New("单词集ID不能为空！")
}
item := model.VocabularyItem{
VocabularySetId: vc.VocabularySetId,
Key:             vc.Key,
Value:           vc.Value,
Image:           vc.Image,
Order:           vc.Order,
}
return repository.VocabularyItemRepo.Create(&item)
}

func (s *VocabularyService) UpdateVocabularyItem(vc model.VocabularyItemCreateEditRequest) error {
if vc.ID == 0 {
return errors.New("单词ID不能为空！")
}
item := model.VocabularyItem{
Model: model.Model{ID: vc.ID},
Key:   vc.Key,
Value: vc.Value,
Image: vc.Image,
Order: vc.Order,
}
err := repository.VocabularyItemRepo.Update(&item)
return err
}

func (s *VocabularyService) DeleteVocabularyItem(id uint) error {
if id == 0 {
return errors.New("单词ID不能为空！")
}
return repository.VocabularyItemRepo.Delete(id)
}

// 仅保留题目生成接口，其他学习相关接口建议后续补充
func (s *VocabularyService) GetVocabularySetTestQuestions(id uint) ([]*model.VocabularySetTestQuestion, error) {
if id == 0 {
return nil, errors.New("无效的ID.")
}
query := model.VocabularyItemQuery{VocabularySetId: id}
total, err := repository.VocabularyItemRepo.Count(&query)
if err != nil {
return nil, errors.New("查询失败.")
}
if total < 4 {
return nil, errors.New("单词数量不足.")
}
items, err := repository.VocabularyItemRepo.FindAll(&query)
if err != nil {
return nil, err
}
if total <= 20 {
return s.buildTestQuestions(items), nil
}
rand.Seed(time.Now().UnixNano())
randNums := rand.Perm(int(total))
randNums = randNums[:20]
selected := []*model.VocabularyItem{}
for _, num := range randNums {
selected = append(selected, items[num])
}
return s.buildTestQuestions(selected), nil
}

func (s *VocabularyService) buildTestQuestions(items []*model.VocabularyItem) []*model.VocabularySetTestQuestion {
questions := []*model.VocabularySetTestQuestion{}
total := len(items)
for _, item := range items {
options := []*model.VocabularyItem{}
rand.Seed(time.Now().UnixNano())
randNums := rand.Perm(total)
randNums = randNums[:4]
for _, n := range randNums {
options = append(options, items[n])
}
iexisted := false
for _, op := range options {
if op.ID == item.ID {
iexisted = true
break
}
}
if !iexisted {
options[0] = item
}
rand.Seed(time.Now().UnixNano())
rand.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
question := &model.VocabularySetTestQuestion{
Word:    item,
Options: options,
}
questions = append(questions, question)
}
rand.Seed(time.Now().UnixNano())
rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })
return questions
}
