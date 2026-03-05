package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var ConversationCtrl = &ConversationController{
	svr: service.ConversationSvr,
}

type ConversationController struct {
	svr *service.ConversationService
}

// StartSession creates a new conversation session
// POST /api/v1/ai/conversation/start
func (ctrl *ConversationController) StartSession(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req model.ConversationStartRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	session, err := ctrl.svr.StartSession(u.ID, req.Title)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Session started successfully!", session)
}

// SendMessage sends a message and gets AI reply
// POST /api/v1/ai/conversation/message
func (ctrl *ConversationController) SendMessage(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req model.ConversationMessageRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	reply, err := ctrl.svr.SendMessage(u.ID, req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Message sent successfully!", reply)
}

// GetHistory returns conversation message history
// POST /api/v1/ai/conversation/history
func (ctrl *ConversationController) GetHistory(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req struct {
		SessionId uint `json:"sessionId" binding:"required"`
		PageIndex int  `json:"pageIndex"`
		PageSize  int  `json:"pageSize"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	msgs, total, err := ctrl.svr.GetHistory(u.ID, req.SessionId, req.PageIndex, req.PageSize)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", gin.H{
		"list":  msgs,
		"total": total,
	})
}

// ListSessions lists all conversation sessions for current user
// POST /api/v1/ai/conversation/sessions
func (ctrl *ConversationController) ListSessions(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req struct {
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	sessions, total, err := ctrl.svr.ListSessions(u.ID, req.PageIndex, req.PageSize)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", gin.H{
		"list":  sessions,
		"total": total,
	})
}

// ResetSession clears all messages in a session
// POST /api/v1/ai/conversation/reset
func (ctrl *ConversationController) ResetSession(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req struct {
		SessionId uint `json:"sessionId" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	if err := ctrl.svr.ResetSession(u.ID, req.SessionId); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Session reset successfully!", nil)
}

// CloseSession marks a session as inactive
// POST /api/v1/ai/conversation/close
func (ctrl *ConversationController) CloseSession(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req struct {
		SessionId uint `json:"sessionId" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	if err := ctrl.svr.CloseSession(u.ID, req.SessionId); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Session closed successfully!", nil)
}
