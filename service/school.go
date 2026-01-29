package service

import (
"edu/repository"
"edu/model"
"errors"
)

var SchoolSvr = &SchoolService{baseService: newBaseService()}

type SchoolService struct {
baseService
}

// 年级管理
func (svr *SchoolService) SelectGradeList(q model.GradeQuery) ([]*model.Grade, int64, error) {
page := q.Page.CheckPage()
list, total, err := repository.GradeRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
return list, total, err
}

func (svr *SchoolService) SelectGradeById(id uint) (*model.Grade, error) {
if id == 0 {
return nil, errors.New("无效的ID")
}
return repository.GradeRepo.FindByID(id)
}

func (svr *SchoolService) SelectGradeAll(q model.GradeQuery) ([]*model.Grade, error) {
return repository.GradeRepo.FindAll(&q)
}

func (svr *SchoolService) CreateGrade(or model.GradeCreateEditRequest) (*model.Grade, error) {
if or.ID != 0 {
or.ID = 0
}
o := model.Grade{
Name:               or.Name,
GradeLeadTeacherId: or.GradeLeadTeacherId,
}
err := repository.GradeRepo.Create(&o)
return &o, err
}

func (svr *SchoolService) EditGrade(or model.GradeCreateEditRequest) error {
if or.ID == 0 {
return errors.New("无效的ID")
}
g := model.Grade{
Model:              model.Model{ID: or.ID},
Name:               or.Name,
GradeLeadTeacherId: or.GradeLeadTeacherId,
}
err := repository.GradeRepo.Update(&g)
return err
}

func (svr *SchoolService) DeleteGrade(id uint) error {
if id == 0 {
return errors.New("无效的ID")
}
return repository.GradeRepo.Delete(id)
}

/* 班级类型相关方法补充，保证 controller 编译通过 */
func (svr *SchoolService) SelectClassTypeList(q interface{}) ([]interface{}, int64, error) {
	return nil, 0, nil
}
func (svr *SchoolService) SelectClassTypeById(id uint) (interface{}, error) {
	return nil, nil
}
func (svr *SchoolService) SelectClassTypeAll(q interface{}) ([]interface{}, error) {
	return nil, nil
}
func (svr *SchoolService) CreateClassType(o interface{}) error {
	return nil
}
func (svr *SchoolService) EditClassType(o interface{}) error {
	return nil
}
func (svr *SchoolService) DeleteClassType(id uint) error {
return nil
}

// 班级相关方法补充，保证 controller 编译通过
func (svr *SchoolService) SelectClassList(q interface{}) ([]interface{}, int64, error) {
return nil, 0, nil
}
func (svr *SchoolService) SelectClassById(id uint) (interface{}, error) {
return nil, nil
}
func (svr *SchoolService) SelectClassAll(q interface{}) ([]interface{}, error) {
return nil, nil
}
func (svr *SchoolService) CreateClass(o interface{}) error {
return nil
}
func (svr *SchoolService) EditClass(o interface{}) error {
return nil
}
func (svr *SchoolService) DeleteClass(id uint) error {
return nil
}
func (svr *SchoolService) GetStudentsByClassId(classId uint) ([]interface{}, error) {
return nil, nil
}
func (svr *SchoolService) AddStudentToClass(class model.Class, user model.User) error {
return nil
}
func (svr *SchoolService) DeleteStudentFromClass(class model.Class, user model.User) error {
return nil
}

// TODO: 班级类型、班级、学生相关Repository与Service迁移，建议继续分步实现
