package service

import (
	"edu/model"
	"edu/repository"
	"errors"
	"strconv"
)

var LearningPlanSvr = &LearningPlanService{baseService: newBaseService()}

type LearningPlanService struct {
	baseService
}

// CreatePlan 为学生创建学习计划并记录初始版本
func (svr *LearningPlanService) CreatePlan(req model.StudentLearningPlanCreateRequest, creatorId uint) (*model.StudentLearningPlan, error) {
	if req.ClassId == 0 || req.UserId == 0 {
		return nil, errors.New("班级ID和用户ID不能为空")
	}
	if req.PlanType != model.LearningPlanTypeLong &&
		req.PlanType != model.LearningPlanTypeMid &&
		req.PlanType != model.LearningPlanTypeShort {
		return nil, errors.New("无效的计划类型，必须为 long/mid/short")
	}

	// 验证班级存在
	class, err := repository.ClassRepo.FindByID(req.ClassId)
	if err != nil || class == nil {
		return nil, errors.New("班级不存在")
	}
	if class.ClassType != model.ClassTypeTeaching {
		return nil, errors.New("只有教学班可以管理学习计划")
	}
	if class.SyllabusId == nil {
		return nil, errors.New("该教学班尚未绑定syllabus，请先绑定syllabus")
	}

	// 验证学生存在
	user, err := repository.UserRepo.FindByID(req.UserId)
	if err != nil || user == nil {
		return nil, errors.New("学生不存在")
	}

	plan := &model.StudentLearningPlan{
		ClassId:   req.ClassId,
		UserId:    req.UserId,
		PlanType:  req.PlanType,
		Content:   req.Content,
		Version:   1,
		CreatedBy: creatorId,
	}
	if err := repository.StudentLearningPlanRepo.Create(plan); err != nil {
		return nil, err
	}

	// 记录初始版本
	version := &model.StudentLearningPlanVersion{
		PlanId:    plan.ID,
		Version:   1,
		Content:   req.Content,
		ChangedBy: creatorId,
		Comment:   req.Comment,
	}
	if err := repository.StudentLearningPlanRepo.CreateVersion(version); err != nil {
		return nil, err
	}

	return plan, nil
}

// UpdatePlan 更新学习计划并记录新版本
func (svr *LearningPlanService) UpdatePlan(req model.StudentLearningPlanUpdateRequest, updaterId uint) (*model.StudentLearningPlan, error) {
	if req.ID == 0 {
		return nil, errors.New("无效的ID")
	}
	plan, err := repository.StudentLearningPlanRepo.FindByID(req.ID)
	if err != nil || plan == nil {
		return nil, errors.New("学习计划不存在")
	}

	newVersion := plan.Version + 1
	updated := &model.StudentLearningPlan{
		Model:   model.Model{ID: req.ID},
		Content: req.Content,
		Version: newVersion,
	}
	if err := repository.StudentLearningPlanRepo.Update(updated); err != nil {
		return nil, err
	}

	// 记录新版本
	v := &model.StudentLearningPlanVersion{
		PlanId:    plan.ID,
		Version:   newVersion,
		Content:   req.Content,
		ChangedBy: updaterId,
		Comment:   req.Comment,
	}
	if err := repository.StudentLearningPlanRepo.CreateVersion(v); err != nil {
		return nil, err
	}

	plan.Content = req.Content
	plan.Version = newVersion
	return plan, nil
}

// DeletePlan 删除学习计划
func (svr *LearningPlanService) DeletePlan(id uint) error {
	if id == 0 {
		return errors.New("无效的ID")
	}
	return repository.StudentLearningPlanRepo.Delete(id)
}

// GetPlanById 根据ID获取学习计划
func (svr *LearningPlanService) GetPlanById(id uint) (*model.StudentLearningPlan, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}
	return repository.StudentLearningPlanRepo.FindByID(id)
}

// ListPlans 分页查询学习计划
func (svr *LearningPlanService) ListPlans(q model.StudentLearningPlanQuery) ([]*model.StudentLearningPlan, int64, error) {
	page := q.Page.CheckPage()
	return repository.StudentLearningPlanRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
}

// GetAllPlans 获取全部学习计划（不分页）
func (svr *LearningPlanService) GetAllPlans(q model.StudentLearningPlanQuery) ([]*model.StudentLearningPlan, error) {
	return repository.StudentLearningPlanRepo.FindAll(&q)
}

// ListPlanVersions 获取学习计划的历史版本列表
func (svr *LearningPlanService) ListPlanVersions(q model.StudentLearningPlanVersionQuery) ([]*model.StudentLearningPlanVersion, int64, error) {
	if q.PlanId == 0 {
		return nil, 0, errors.New("planId不能为空")
	}
	page := q.Page.CheckPage()
	return repository.StudentLearningPlanRepo.FindVersionsByPlanId(q.PlanId, (page.PageIndex-1)*page.PageSize, page.PageSize)
}

// RollbackPlan 回滚学习计划到指定版本
func (svr *LearningPlanService) RollbackPlan(req model.StudentLearningPlanRollbackRequest, operatorId uint) (*model.StudentLearningPlan, error) {
	if req.PlanId == 0 || req.Version == 0 {
		return nil, errors.New("planId和version不能为空")
	}
	plan, err := repository.StudentLearningPlanRepo.FindByID(req.PlanId)
	if err != nil || plan == nil {
		return nil, errors.New("学习计划不存在")
	}
	targetVersion, err := repository.StudentLearningPlanRepo.FindVersionByPlanAndVersion(req.PlanId, req.Version)
	if err != nil || targetVersion == nil {
		return nil, errors.New("目标版本不存在")
	}

	newVersion := plan.Version + 1
	comment := req.Comment
	if comment == "" {
		comment = "回滚到版本 " + strconv.Itoa(req.Version)
	}

	updated := &model.StudentLearningPlan{
		Model:   model.Model{ID: req.PlanId},
		Content: targetVersion.Content,
		Version: newVersion,
	}
	if err := repository.StudentLearningPlanRepo.Update(updated); err != nil {
		return nil, err
	}

	// 记录回滚版本
	v := &model.StudentLearningPlanVersion{
		PlanId:    req.PlanId,
		Version:   newVersion,
		Content:   targetVersion.Content,
		ChangedBy: operatorId,
		Comment:   comment,
	}
	if err := repository.StudentLearningPlanRepo.CreateVersion(v); err != nil {
		return nil, err
	}

	plan.Content = targetVersion.Content
	plan.Version = newVersion
	return plan, nil
}
