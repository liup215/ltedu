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
