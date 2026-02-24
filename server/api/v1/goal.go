package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var GoalCtrl = &GoalController{
	goalSvr: service.GoalSvr,
}

type GoalController struct {
	goalSvr *service.GoalService
}

// CreateGoal creates a new learning goal
// POST /api/v1/goal/create
// @Summary      创建学习目标
// @Description  为当前用户创建新的学习目标
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  model.GoalCreateRequest  true  "目标信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/goal/create [post]
func (ctrl *GoalController) CreateGoal(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req model.GoalCreateRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	goal, err := ctrl.goalSvr.CreateGoal(u.ID, req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Goal created successfully!", goal)
}

// UpdateGoal updates an existing goal
// POST /api/v1/goal/edit
// @Summary      更新学习目标
// @Description  更新现有学习目标
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  model.GoalUpdateRequest  true  "目标信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/goal/edit [post]
func (ctrl *GoalController) UpdateGoal(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req model.GoalUpdateRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	goal, err := ctrl.goalSvr.UpdateGoal(u.ID, req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Goal updated successfully!", goal)
}

// GetGoalById gets a goal by ID
// POST /api/v1/goal/byId
// @Summary      根据ID获取学习目标
// @Description  获取指定ID的学习目标详情
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "目标ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/goal/byId [post]
func (ctrl *GoalController) GetGoalById(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	goal, err := ctrl.goalSvr.GetGoalByID(u.ID, req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", goal)
}

// ListGoals lists goals with pagination
// POST /api/v1/goal/list
// @Summary      获取学习目标列表
// @Description  分页查询当前用户的学习目标列表
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  model.GoalQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/goal/list [post]
func (ctrl *GoalController) ListGoals(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var query model.GoalQuery
	if err := c.BindJSON(&query); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	query.CheckPage()
	goals, total, err := ctrl.goalSvr.ListGoals(u.ID, query)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", gin.H{
		"list":  goals,
		"total": total,
	})
}

// GetActiveGoals gets all active goals for the current user
// POST /api/v1/goal/active
// @Summary      获取活跃学习目标
// @Description  获取当前用户所有活跃状态的学习目标
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Failure      400  {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/goal/active [post]
func (ctrl *GoalController) GetActiveGoals(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	goals, err := ctrl.goalSvr.GetActiveGoals(u.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", goals)
}

// DeleteGoal deletes a goal
// POST /api/v1/goal/delete
// @Summary      删除学习目标
// @Description  删除指定学习目标
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "目标ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/goal/delete [post]
func (ctrl *GoalController) DeleteGoal(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	err = ctrl.goalSvr.DeleteGoal(u.ID, req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Goal deleted successfully!", nil)
}

// CompleteDiagnostic marks the diagnostic test as completed
// POST /api/v1/goal/diagnostic/complete
// @Summary      完成诊断测试
// @Description  标记指定目标的诊断测试已完成
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "目标ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/goal/diagnostic/complete [post]
func (ctrl *GoalController) CompleteDiagnostic(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req struct {
		GoalId uint `json:"goalId" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	err = ctrl.goalSvr.CompleteDiagnostic(u.ID, req.GoalId)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Diagnostic completed successfully!", nil)
}
