package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var LearningPlanCtrl = &LearningPlanController{
	planSvr: service.LearningPlanSvr,
}

type LearningPlanController struct {
	planSvr *service.LearningPlanService
}

// CreatePlan 创建学习计划
func (ctrl *LearningPlanController) CreatePlan(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", nil)
		return
	}
	req := model.StudentLearningPlanCreateRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	plan, err := ctrl.planSvr.CreatePlan(req, u.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "学习计划创建成功!", plan)
}

// UpdatePlan 更新学习计划
func (ctrl *LearningPlanController) UpdatePlan(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", nil)
		return
	}
	req := model.StudentLearningPlanUpdateRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	plan, err := ctrl.planSvr.UpdatePlan(req, u.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "学习计划更新成功!", plan)
}

// DeletePlan 删除学习计划
func (ctrl *LearningPlanController) DeletePlan(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.planSvr.DeletePlan(req.ID); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "学习计划删除成功!", nil)
}

// GetPlanById 根据ID获取学习计划
func (ctrl *LearningPlanController) GetPlanById(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	plan, err := ctrl.planSvr.GetPlanById(req.ID)
	if err != nil {
		http.ErrorData(c, "数据获取失败: "+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", plan)
}

// ListPlans 分页查询学习计划
func (ctrl *LearningPlanController) ListPlans(c *gin.Context) {
	q := model.StudentLearningPlanQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.planSvr.ListPlans(q)
	if err != nil {
		http.ErrorData(c, "数据获取失败: "+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

// ListPlanVersions 获取学习计划的历史版本
func (ctrl *LearningPlanController) ListPlanVersions(c *gin.Context) {
	q := model.StudentLearningPlanVersionQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	versions, total, err := ctrl.planSvr.ListPlanVersions(q)
	if err != nil {
		http.ErrorData(c, "数据获取失败: "+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  versions,
		"total": total,
	})
}

// RollbackPlan 回滚学习计划到历史版本
func (ctrl *LearningPlanController) RollbackPlan(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", nil)
		return
	}
	req := model.StudentLearningPlanRollbackRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	plan, err := ctrl.planSvr.RollbackPlan(req, u.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "学习计划回滚成功!", plan)
}
