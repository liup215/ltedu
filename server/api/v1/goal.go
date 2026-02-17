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

	query = query.CheckPage()
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
