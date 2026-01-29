package service

import (
	"edu/repository"
	"edu/model"
	"errors"
)

var QualificationSvr = &QualificationService{baseService: newBaseService()}

type QualificationService struct {
	baseService
}

// Organisation 管理
func (svr *QualificationService) SelectOrganisationList(q model.OrganisationQuery) ([]*model.Organisation, int64, error) {
	page := q.Page.CheckPage()
list, total, err := repository.OrganisationRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
return list, total, err
}

func (svr *QualificationService) SelectOrganisationById(id uint) (*model.Organisation, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}
	return repository.OrganisationRepo.FindByID(id)
}

func (svr *QualificationService) SelectOrganisationAll(q model.OrganisationQuery) ([]*model.Organisation, error) {
	return repository.OrganisationRepo.FindAll(&q)
}

func (svr *QualificationService) CreateOrganisation(o model.Organisation) (*model.Organisation, error) {
	if o.ID != 0 {
		o.ID = 0
	}
	err := repository.OrganisationRepo.Create(&o)
	return &o, err
}

func (svr *QualificationService) EditOrganisation(o model.Organisation) (*model.Organisation, error) {
	if o.ID == 0 {
		return nil, errors.New("无效的ID")
	}
err := repository.OrganisationRepo.Update(&o)
return &o, err
}

func (svr *QualificationService) DeleteOrganisation(id uint) error {
	if id == 0 {
		return errors.New("无效的ID")
	}
	// 这里可加业务校验
	return repository.OrganisationRepo.Delete(id)
}

// Qualification 管理
func (svr *QualificationService) SelectQualificationList(q model.QualificationQuery) ([]*model.Qualification, int64, error) {
	page := q.Page.CheckPage()
list, total, err := repository.QualificationRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
return list, total, err
}

func (svr *QualificationService) SelectQualificationById(id uint) (*model.Qualification, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}
	return repository.QualificationRepo.FindByID(id)
}

func (svr *QualificationService) SelectQualificationAll(q model.QualificationQuery) ([]*model.Qualification, error) {
	return repository.QualificationRepo.FindAll(&q)
}

func (svr *QualificationService) CreateQualification(o model.Qualification) (*model.Qualification, error) {
	if o.ID != 0 {
		o.ID = 0
	}
	err := repository.QualificationRepo.Create(&o)
	return &o, err
}

func (svr *QualificationService) EditQualification(o model.Qualification) (*model.Qualification, error) {
	if o.ID == 0 {
		return nil, errors.New("无效的ID")
	}
err := repository.QualificationRepo.Update(&o)
return &o, err
}

func (svr *QualificationService) DeleteQualification(id uint) error {
	if id == 0 {
		return errors.New("无效的ID")
	}
	// 这里可加业务校验
	return repository.QualificationRepo.Delete(id)
}

// Syllabus 管理
func (svr *QualificationService) SelectSyllabusList(q model.SyllabusQuery) ([]*model.Syllabus, int64, error) {
	page := q.Page.CheckPage()
list, total, err := repository.SyllabusRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
return list, total, err
}

func (svr *QualificationService) SelectSyllabusById(id uint) (*model.Syllabus, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}
	return repository.SyllabusRepo.FindByID(id)
}

func (svr *QualificationService) SelectSyllabusAll(q model.SyllabusQuery) ([]*model.Syllabus, error) {
	return repository.SyllabusRepo.FindAll(&q)
}

func (svr *QualificationService) CreateSyllabus(o model.Syllabus) (*model.Syllabus, error) {
	if o.ID != 0 {
		o.ID = 0
	}
	err := repository.SyllabusRepo.Create(&o)
	return &o, err
}

func (svr *QualificationService) EditSyllabus(o model.Syllabus) (*model.Syllabus, error) {
	if o.ID == 0 {
		return nil, errors.New("无效的ID")
	}
err := repository.SyllabusRepo.Update(&o)
return &o, err
}

func (svr *QualificationService) DeleteSyllabus(id uint) error {
	if id == 0 {
		return errors.New("无效的ID")
	}
	// 这里可加业务校验
	return repository.SyllabusRepo.Delete(id)
}

// Chapter 管理
func (s *QualificationService) ChapterList(q model.ChapterQuery) ([]*model.Chapter, int64, error) {
	page := q.PageIndex
	if page <= 0 {
		page = 1
	}
	size := q.PageSize
	if size <= 0 {
		size = 20
	}
list, total, err := repository.ChapterRepo.FindPage(&q, (page-1)*size, size)
return list, total, err
}

func (s *QualificationService) GetChapterTree(sId uint) []*model.Chapter {
	tree := []*model.Chapter{}
	list, _ := repository.ChapterRepo.FindBySyllabusID(sId)
	for _, c := range list {
		if c.ParentId == 0 {
			children := s.getChapterChildrenFromList(c.ID, list)
			c.Children = children
			tree = append(tree, c)
		}
	}
	return tree
}

func (s *QualificationService) getChapterChildrenFromList(cid uint, list []*model.Chapter) []*model.Chapter {
	tree := []*model.Chapter{}
	for _, c := range list {
		if c.ParentId == cid {
			children := s.getChapterChildrenFromList(c.ID, list)
			c.Children = children
			tree = append(tree, c)
		}
	}
	return tree
}

func (s *QualificationService) buildChapterTree(parentId, syllabusId uint) []*model.Chapter {
	list, _ := repository.ChapterRepo.FindByParentID(parentId)
	for i, chapter := range list {
		children := s.buildChapterTree(chapter.ID, syllabusId)
		if len(children) == 0 {
			list[i].IsLeaf = 1
		}
		list[i].Children = children
	}
	return list
}

func (svr *QualificationService) SelectChapterById(id uint) (*model.Chapter, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}
	return repository.ChapterRepo.FindByID(id)
}

func (s *QualificationService) CreateChapter(chapter model.Chapter) (*model.Chapter, error) {
	if chapter.ID != uint(0) {
		chapter.ID = uint(0)
	}
	if chapter.SyllabusId == 0 {
		return nil, errors.New("没有考纲信息！")
	}
	// 校验考纲存在
	if _, e := repository.SyllabusRepo.FindByID(chapter.SyllabusId); e != nil {
		return nil, errors.New("考纲查询失败！")
	}
	e := repository.ChapterRepo.Create(&chapter)
	return &chapter, e
}

func (s *QualificationService) EditChapter(chapter model.Chapter) (*model.Chapter, error) {
	if chapter.ID == 0 {
		return nil, errors.New("无效的chapterID")
	}
e := repository.ChapterRepo.Update(&chapter)
return &chapter, e
}

func (s *QualificationService) DeleteChapter(id uint) error {
	if id == uint(0) {
		return errors.New("无效的ID")
	}
	// 校验是否有子章节
	children, _ := repository.ChapterRepo.FindByParentID(id)
	if len(children) > 0 {
		return errors.New("该章节存在子章节，不能删除！")
	}
	return repository.ChapterRepo.Delete(id)
}

func (s *QualificationService) ChapterById(id uint) (*model.Chapter, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}
	return repository.ChapterRepo.FindByID(id)
}
