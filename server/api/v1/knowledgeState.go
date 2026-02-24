package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
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
// @Summary      获取章节知识状态
// @Description  获取指定目标和章节的知识掌握状态
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "目标ID和章节ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-state/byChapter [post]
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
// @Summary      获取目标知识状态列表
// @Description  获取指定目标的所有章节知识状态
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "目标ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-state/list [post]
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
// @Summary      获取学习进度
// @Description  获取用户在指定目标上的学习进度
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "目标ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-state/progress [post]
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
// @Summary      获取待复习章节
// @Description  获取指定目标中需要复习的章节列表
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "目标ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-state/due-review [post]
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
