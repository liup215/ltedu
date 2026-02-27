package service

import (
	"edu/model"
	"edu/repository"
	"errors"
)

var PhasePlanSvr = &PhasePlanService{baseService: newBaseService()}

type PhasePlanService struct {
	baseService
}

// CreatePhasePlan 为学习计划创建阶段性计划
func (svr *PhasePlanService) CreatePhasePlan(req model.LearningPhasePlanCreateRequest) (*model.LearningPhasePlan, error) {
	plan, err := repository.StudentLearningPlanRepo.FindByID(req.PlanId)
	if err != nil || plan == nil {
		return nil, errors.New("学习计划不存在")
	}
	node, err := repository.ExamNodeRepo.FindByID(req.ExamNodeId)
	if err != nil || node == nil {
		return nil, errors.New("考试节点不存在")
	}

	pp := &model.LearningPhasePlan{
		PlanId:     req.PlanId,
		ExamNodeId: req.ExamNodeId,
		Title:      req.Title,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		SortOrder:  req.SortOrder,
	}
	if err := repository.PhasePlanRepo.Create(pp); err != nil {
		return nil, err
	}
	return pp, nil
}

// UpdatePhasePlan 更新阶段性计划基本信息
func (svr *PhasePlanService) UpdatePhasePlan(req model.LearningPhasePlanUpdateRequest) (*model.LearningPhasePlan, error) {
	if req.ID == 0 {
		return nil, errors.New("无效的阶段性计划ID")
	}
	pp := &model.LearningPhasePlan{
		Model:     model.Model{ID: req.ID},
		Title:     req.Title,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		SortOrder: req.SortOrder,
	}
	if err := repository.PhasePlanRepo.Update(pp); err != nil {
		return nil, err
	}
	return repository.PhasePlanRepo.FindByID(req.ID)
}

// DeletePhasePlan 删除阶段性计划
func (svr *PhasePlanService) DeletePhasePlan(id uint) error {
	if id == 0 {
		return errors.New("无效的阶段性计划ID")
	}
	return repository.PhasePlanRepo.Delete(id)
}

// GetPhasePlan 根据ID获取阶段性计划
func (svr *PhasePlanService) GetPhasePlan(id uint) (*model.LearningPhasePlan, error) {
	if id == 0 {
		return nil, errors.New("无效的阶段性计划ID")
	}
	return repository.PhasePlanRepo.FindByID(id)
}

// ListPhasePlans 获取某个学习计划下的所有阶段性计划
func (svr *PhasePlanService) ListPhasePlans(planId uint) ([]*model.LearningPhasePlan, error) {
	if planId == 0 {
		return nil, errors.New("planId不能为空")
	}
	return repository.PhasePlanRepo.FindByPlanID(planId)
}

// AddChapterToPhasePlan 为阶段性计划添加章节
func (svr *PhasePlanService) AddChapterToPhasePlan(req model.LearningPhasePlanAddChapterRequest) error {
	if req.PhasePlanId == 0 || req.ChapterId == 0 {
		return errors.New("phasePlanId和chapterId不能为空")
	}
	pp, err := repository.PhasePlanRepo.FindByID(req.PhasePlanId)
	if err != nil || pp == nil {
		return errors.New("阶段性计划不存在")
	}
	ch, err := repository.ChapterRepo.FindByID(req.ChapterId)
	if err != nil || ch == nil {
		return errors.New("章节不存在")
	}
	return repository.PhasePlanRepo.AddChapters(req.PhasePlanId, []uint{req.ChapterId})
}

// RemoveChapterFromPhasePlan 从阶段性计划移除章节
func (svr *PhasePlanService) RemoveChapterFromPhasePlan(req model.LearningPhasePlanRemoveChapterRequest) error {
	if req.PhasePlanId == 0 || req.ChapterId == 0 {
		return errors.New("phasePlanId和chapterId不能为空")
	}
	return repository.PhasePlanRepo.RemoveChapter(req.PhasePlanId, req.ChapterId)
}
