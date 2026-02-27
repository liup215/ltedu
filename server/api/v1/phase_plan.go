package v1

import (
	"edu/lib/net/http"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var PhasePlanCtrl = &PhasePlanController{
	svr: service.PhasePlanSvr,
}

type PhasePlanController struct {
	svr *service.PhasePlanService
}

// CreatePhasePlan 创建阶段性计划
func (ctrl *PhasePlanController) CreatePhasePlan(c *gin.Context) {
	req := model.LearningPhasePlanCreateRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	pp, err := ctrl.svr.CreatePhasePlan(req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "阶段性计划创建成功!", pp)
}

// UpdatePhasePlan 更新阶段性计划
func (ctrl *PhasePlanController) UpdatePhasePlan(c *gin.Context) {
	req := model.LearningPhasePlanUpdateRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	pp, err := ctrl.svr.UpdatePhasePlan(req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "阶段性计划更新成功!", pp)
}

// DeletePhasePlan 删除阶段性计划
func (ctrl *PhasePlanController) DeletePhasePlan(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.svr.DeletePhasePlan(req.ID); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "阶段性计划删除成功!", nil)
}

// GetPhasePlanById 根据ID获取阶段性计划
func (ctrl *PhasePlanController) GetPhasePlanById(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	pp, err := ctrl.svr.GetPhasePlan(req.ID)
	if err != nil {
		http.ErrorData(c, "数据获取失败: "+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", pp)
}

// ListPhasePlans 获取某学习计划的所有阶段性计划
func (ctrl *PhasePlanController) ListPhasePlans(c *gin.Context) {
	var req struct {
		PlanId uint `json:"planId" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, err := ctrl.svr.ListPhasePlans(req.PlanId)
	if err != nil {
		http.ErrorData(c, "数据获取失败: "+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

// AddChapter 为阶段性计划添加章节
func (ctrl *PhasePlanController) AddChapter(c *gin.Context) {
	req := model.LearningPhasePlanAddChapterRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.svr.AddChapterToPhasePlan(req); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "章节添加成功!", nil)
}

// RemoveChapter 从阶段性计划移除章节
func (ctrl *PhasePlanController) RemoveChapter(c *gin.Context) {
	req := model.LearningPhasePlanRemoveChapterRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.svr.RemoveChapterFromPhasePlan(req); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "章节移除成功!", nil)
}
