package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var TaskCtrl = &TaskController{
	taskSvr: service.TaskSvr,
}

type TaskController struct {
	taskSvr *service.TaskService
}

// CreateTask creates a new task
// POST /api/v1/task/create
// @Summary      创建任务
// @Description  为当前用户创建新学习任务
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  model.TaskCreateRequest  true  "任务信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/task/create [post]
func (ctrl *TaskController) CreateTask(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req model.TaskCreateRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	task, err := ctrl.taskSvr.CreateTask(u.ID, req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Task created successfully!", task)
}

// UpdateTask updates an existing task
// POST /api/v1/task/edit
// @Summary      更新任务
// @Description  更新现有学习任务
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  model.TaskUpdateRequest  true  "任务信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/task/edit [post]
func (ctrl *TaskController) UpdateTask(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req model.TaskUpdateRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	task, err := ctrl.taskSvr.UpdateTask(u.ID, req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Task updated successfully!", task)
}

// GetTaskById gets a task by ID
// POST /api/v1/task/byId
// @Summary      根据ID获取任务
// @Description  获取指定ID的学习任务详情
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "任务ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/task/byId [post]
func (ctrl *TaskController) GetTaskById(c *gin.Context) {
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

	task, err := ctrl.taskSvr.GetTaskByID(u.ID, req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", task)
}

// GetTaskStream gets today's tasks, upcoming tasks, and overdue tasks
// POST /api/v1/task/stream
// @Summary      获取任务流
// @Description  获取今日任务、即将到期任务和逾期任务
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "目标ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/task/stream [post]
func (ctrl *TaskController) GetTaskStream(c *gin.Context) {
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

	stream, err := ctrl.taskSvr.GetTaskStream(u.ID, req.GoalId)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", stream)
}

// CompleteTask marks a task as completed
// POST /api/v1/task/complete
// @Summary      完成任务
// @Description  标记指定任务为已完成
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  model.TaskCompleteRequest  true  "任务ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/task/complete [post]
func (ctrl *TaskController) CompleteTask(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req model.TaskCompleteRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	err = ctrl.taskSvr.CompleteTask(u.ID, req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Task completed successfully!", nil)
}

// GenerateInitialPlan generates initial 7-day task plan for a new goal
// POST /api/v1/task/generate-plan
// @Summary      生成初始学习计划
// @Description  为新目标生成初始7天学习任务计划
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "目标ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/task/generate-plan [post]
func (ctrl *TaskController) GenerateInitialPlan(c *gin.Context) {
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

	err = ctrl.taskSvr.GenerateInitialPlan(u.ID, req.GoalId)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Initial plan generated successfully!", nil)
}

// ListTasks lists tasks with pagination
// POST /api/v1/task/list
// @Summary      获取任务列表
// @Description  分页查询当前用户的学习任务列表
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  model.TaskQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/task/list [post]
func (ctrl *TaskController) ListTasks(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var query model.TaskQuery
	if err := c.BindJSON(&query); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	query.CheckPage()
	tasks, total, err := ctrl.taskSvr.ListTasks(u.ID, query)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", gin.H{
		"list":  tasks,
		"total": total,
	})
}

// DeleteTask deletes a task
// POST /api/v1/task/delete
// @Summary      删除任务
// @Description  删除指定学习任务
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "任务ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/task/delete [post]
func (ctrl *TaskController) DeleteTask(c *gin.Context) {
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

	err = ctrl.taskSvr.DeleteTask(u.ID, req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Task deleted successfully!", nil)
}
