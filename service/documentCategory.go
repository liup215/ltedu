package service

import (
"edu/repository"
"edu/model"
"errors"
)

var DocumentCategorySvr = &DocumentCategoryService{baseService: newBaseService()}

type DocumentCategoryService struct {
baseService
}

func (s *DocumentCategoryService) CreateCategory(ce model.DocumentCategoryCreateEditRequest) error {
if ce.ID != 0 {
ce.ID = uint(0)
}
c := (&ce).GetCategory()
return repository.DocumentCategoryRepo.Create(&c)
}

func (s *DocumentCategoryService) EditCategory(ce model.DocumentCategoryCreateEditRequest) error {
if ce.ID == 0 {
return errors.New("id不能为空")
}
c := (&ce).GetCategory()
err := repository.DocumentCategoryRepo.Update(&c)
return err
}

func (s *DocumentCategoryService) DeleteCategory(id uint) error {
if id == 0 {
return errors.New("id不能为空")
}
count, err := repository.DocumentCategoryRepo.CountByParentID(id)
if err != nil {
return errors.New("查询子分类失败, " + err.Error())
}
if count > 0 {
return errors.New("该分类下有子分类，不能删除")
}
return repository.DocumentCategoryRepo.Delete(id)
}

func (s *DocumentCategoryService) SelectCategoryList(q interface{}) ([]*model.DocumentCategory, int64, error) {
return nil, 0, nil
}

func (s *DocumentCategoryService) SelectCategoryAll(q interface{}) ([]*model.DocumentCategory, error) {
return nil, nil
}

func (s *DocumentCategoryService) SelectCategoryById(id uint) (*model.DocumentCategory, error) {
if id == 0 {
return nil, errors.New("id不能为空")
}
return repository.DocumentCategoryRepo.FindByID(id)
}
