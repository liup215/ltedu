package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var KnowledgeStateCtrl = &KnowledgeStateController{
	knowledgeStateSvr: service.KnowledgeStateSvr,
}

type KnowledgeStateController struct {
	knowledgeStateSvr *service.KnowledgeStateService
}

// GetKnowledgeState gets knowledge state for a specific chapter
// POST /api/v1/knowledge-state/byChapter
func (ctrl *KnowledgeStateController) GetKnowledgeState(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req struct {
		GoalId    uint `json:"goalId" binding:"required"`
		ChapterId uint `json:"chapterId" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	state, err := ctrl.knowledgeStateSvr.GetKnowledgeState(u.ID, req.GoalId, req.ChapterId)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", state)
}

// GetKnowledgeStates gets all knowledge states for a goal
// POST /api/v1/knowledge-state/list
func (ctrl *KnowledgeStateController) GetKnowledgeStates(c *gin.Context) {
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

	states, err := ctrl.knowledgeStateSvr.GetKnowledgeStates(u.ID, req.GoalId)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", states)
}

// GetProgress gets user's progress for a goal
// POST /api/v1/knowledge-state/progress
func (ctrl *KnowledgeStateController) GetProgress(c *gin.Context) {
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

	progress, err := ctrl.knowledgeStateSvr.GetProgress(u.ID, req.GoalId)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", progress)
}

// GetDueForReview gets chapters that are due for review
// POST /api/v1/knowledge-state/due-review
func (ctrl *KnowledgeStateController) GetDueForReview(c *gin.Context) {
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

	states, err := ctrl.knowledgeStateSvr.GetDueForReview(u.ID, req.GoalId)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", states)
}
