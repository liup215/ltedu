package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

// FeedbackCtrl is the singleton controller for feedback endpoints.
var FeedbackCtrl = &FeedbackController{
	feedbackSvr: service.FeedbackSvr,
}

// FeedbackController handles all feedback-related HTTP requests.
type FeedbackController struct {
	feedbackSvr *service.FeedbackService
}

// Submit handles POST /v1/feedback/submit — authenticated users submit feedback.
//
// @Summary      Submit user feedback
// @Description  Allows authenticated users to submit feedback with optional rating and type.
// @Tags         Feedback
// @Accept       json
// @Produce      json
// @Param        body  body  object  true  "Feedback payload"
// @Success      200   {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /v1/feedback/submit [post]
func (ctrl *FeedbackController) Submit(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req struct {
		Type         string `json:"type"`
		Content      string `json:"content" binding:"required"`
		Rating       int    `json:"rating"`
		PageContext  string `json:"pageContext"`
		ConsentGiven bool   `json:"consentGiven"`
		UserAgent    string `json:"userAgent"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "Invalid parameters", nil)
		return
	}

	feedbackType := req.Type
	if feedbackType == "" {
		feedbackType = model.FeedbackTypeGeneral
	}

	feedback := &model.UserFeedback{
		UserID:       u.ID,
		Type:         feedbackType,
		Content:      req.Content,
		Rating:       req.Rating,
		PageContext:  req.PageContext,
		ConsentGiven: req.ConsentGiven,
		UserAgent:    req.UserAgent,
	}

	if err := ctrl.feedbackSvr.Submit(feedback); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "Feedback submitted successfully", feedback)
}

// List handles POST /v1/feedback/list — admin list of all feedback entries.
//
// @Summary      List all feedback (admin)
// @Description  Returns paginated feedback with optional filters for status and type.
// @Tags         Feedback
// @Accept       json
// @Produce      json
// @Param        body  body  model.FeedbackListRequest  true  "List request"
// @Success      200   {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /v1/feedback/list [post]
func (ctrl *FeedbackController) List(c *gin.Context) {
	var req model.FeedbackListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "Invalid parameters", nil)
		return
	}

	items, total, err := ctrl.feedbackSvr.List(req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "OK", gin.H{"list": items, "total": total})
}

// MyFeedback handles POST /v1/feedback/my — list feedback submitted by current user.
//
// @Summary      List own feedback
// @Description  Returns paginated feedback submitted by the currently authenticated user.
// @Tags         Feedback
// @Accept       json
// @Produce      json
// @Param        body  body  model.Page  true  "Page parameters"
// @Success      200   {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /v1/feedback/my [post]
func (ctrl *FeedbackController) MyFeedback(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var page model.Page
	_ = c.ShouldBindJSON(&page)

	items, total, err := ctrl.feedbackSvr.MyFeedback(u.ID, page)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "OK", gin.H{"list": items, "total": total})
}

// GetByID handles POST /v1/feedback/byId — get single feedback record (admin).
//
// @Summary      Get feedback by ID (admin)
// @Tags         Feedback
// @Accept       json
// @Produce      json
// @Param        body  body  object  true  "{ \"id\": 1 }"
// @Success      200   {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /v1/feedback/byId [post]
func (ctrl *FeedbackController) GetByID(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "Invalid parameters", nil)
		return
	}

	item, err := ctrl.feedbackSvr.GetByID(req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "OK", item)
}

// UpdateStatus handles POST /v1/feedback/updateStatus — update status and admin note (admin).
//
// @Summary      Update feedback status (admin)
// @Tags         Feedback
// @Accept       json
// @Produce      json
// @Param        body  body  object  true  "{ \"id\": 1, \"status\": \"reviewed\", \"adminNote\": \"...\" }"
// @Success      200   {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /v1/feedback/updateStatus [post]
func (ctrl *FeedbackController) UpdateStatus(c *gin.Context) {
	var req struct {
		ID        uint   `json:"id" binding:"required"`
		Status    string `json:"status" binding:"required"`
		AdminNote string `json:"adminNote"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "Invalid parameters", nil)
		return
	}

	if err := ctrl.feedbackSvr.UpdateStatus(req.ID, req.Status, req.AdminNote); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "Status updated", nil)
}

// GetStats handles GET /v1/feedback/stats — aggregate statistics (admin).
//
// @Summary      Get feedback statistics (admin)
// @Tags         Feedback
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /v1/feedback/stats [get]
func (ctrl *FeedbackController) GetStats(c *gin.Context) {
	stats, err := ctrl.feedbackSvr.GetStats()
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "OK", stats)
}
