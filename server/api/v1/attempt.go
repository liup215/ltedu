package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var AttemptCtrl = &AttemptController{
	attemptSvr: service.AttemptSvr,
}

type AttemptController struct {
	attemptSvr *service.AttemptService
}

// CreateAttempt creates a new learning attempt
// POST /api/v1/attempt/create
// @Summary      创建学习记录
// @Description  记录一次学习练习尝试
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  model.AttemptCreateRequest  true  "学习记录信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/attempt/create [post]
func (ctrl *AttemptController) CreateAttempt(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req model.AttemptCreateRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	attempt, err := ctrl.attemptSvr.CreateAttempt(u.ID, req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Attempt created successfully!", attempt)
}

// GetRecentAttempts gets recent attempts for a goal
// POST /api/v1/attempt/recent
// @Summary      获取最近学习记录
// @Description  获取指定目标最近的学习记录
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "目标ID和数量限制"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/attempt/recent [post]
func (ctrl *AttemptController) GetRecentAttempts(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req struct {
		GoalId uint `json:"goalId" binding:"required"`
		Limit  int  `json:"limit"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	if req.Limit == 0 {
		req.Limit = 20
	}

	attempts, err := ctrl.attemptSvr.GetRecentAttempts(u.ID, req.GoalId, req.Limit)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", attempts)
}

// GetAttemptStats gets attempt statistics for a goal
// POST /api/v1/attempt/stats
// @Summary      获取学习统计
// @Description  获取指定目标的学习统计数据
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "目标ID和可选章节ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/attempt/stats [post]
func (ctrl *AttemptController) GetAttemptStats(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req struct {
		GoalId    uint `json:"goalId" binding:"required"`
		ChapterId uint `json:"chapterId"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	var stats *model.AttemptStatsResponse
	var err2 error
	if req.ChapterId != 0 {
		stats, err2 = ctrl.attemptSvr.GetAttemptStatsByChapter(u.ID, req.GoalId, req.ChapterId)
	} else {
		stats, err2 = ctrl.attemptSvr.GetAttemptStats(u.ID, req.GoalId)
	}

	if err2 != nil {
		http.ErrorData(c, err2.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", stats)
}

// ListAttempts lists attempts with pagination
// POST /api/v1/attempt/list
// @Summary      获取学习记录列表
// @Description  分页查询当前用户的学习记录
// @Tags         学习导航
// @Accept       json
// @Produce      json
// @Param        body  body  model.AttemptQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/attempt/list [post]
func (ctrl *AttemptController) ListAttempts(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var query model.AttemptQuery
	if err := c.BindJSON(&query); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	query.CheckPage()
	attempts, total, err := ctrl.attemptSvr.ListAttempts(u.ID, query)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", gin.H{
		"list":  attempts,
		"total": total,
	})
}
