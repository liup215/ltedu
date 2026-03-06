package v1

import (
"edu/lib/net/http"
"edu/lib/net/http/middleware/auth"
"edu/model"
"edu/service"

"github.com/gin-gonic/gin"
)

var ConversationCtrl = &ConversationController{
conversationSvr: service.ConversationSvr,
}

// ConversationController handles AI conversation session API endpoints.
type ConversationController struct {
conversationSvr *service.ConversationService
}

// StartSession creates a new conversation session for the authenticated user.
// @Summary      Start a new AI conversation session
// @Description  Creates a new educational AI conversation session, optionally scoped to a subject
// @Tags         Conversation
// @Accept       json
// @Produce      json
// @Param        body  body  model.ConversationStartRequest  true  "Session config"
// @Success      200   {object}  map[string]interface{}  "Session created"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Router       /v1/ai/conversation/start [post]
func (ctrl *ConversationController) StartSession(c *gin.Context) {
var req model.ConversationStartRequest
if err := c.BindJSON(&req); err != nil {
http.ErrorData(c, "参数解析失败", nil)
return
}

user, err := auth.GetCurrentUser(c)
if err != nil {
http.ErrorData(c, "未登录", nil)
return
}

session, err := ctrl.conversationSvr.StartSession(user.ID, req)
if err != nil {
http.ErrorData(c, err.Error(), nil)
return
}
http.SuccessData(c, "会话已创建", session)
}

// SendMessage sends a user message in an existing session and returns the AI reply.
// @Summary      Send a message in a conversation session
// @Description  Sends a user message and returns the AI assistant reply within the educational context
// @Tags         Conversation
// @Accept       json
// @Produce      json
// @Param        body  body  model.ConversationMessageRequest  true  "Message"
// @Success      200   {object}  map[string]interface{}  "Reply"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Router       /v1/ai/conversation/message [post]
func (ctrl *ConversationController) SendMessage(c *gin.Context) {
var req model.ConversationMessageRequest
if err := c.BindJSON(&req); err != nil {
http.ErrorData(c, "参数解析失败", nil)
return
}
if req.SessionID == 0 || req.Message == "" {
http.ErrorData(c, "sessionId和message不能为空", nil)
return
}

user, err := auth.GetCurrentUser(c)
if err != nil {
http.ErrorData(c, "未登录", nil)
return
}

reply, err := ctrl.conversationSvr.SendMessage(user.ID, req)
if err != nil {
http.ErrorData(c, err.Error(), nil)
return
}
http.SuccessData(c, "ok", reply)
}

// GetHistory returns the message history of a conversation session.
// @Summary      Get conversation history
// @Description  Returns all messages in the specified conversation session
// @Tags         Conversation
// @Accept       json
// @Produce      json
// @Param        body  body  model.ConversationHistoryRequest  true  "Session ID"
// @Success      200   {object}  map[string]interface{}  "History"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Router       /v1/ai/conversation/history [post]
func (ctrl *ConversationController) GetHistory(c *gin.Context) {
var req model.ConversationHistoryRequest
if err := c.BindJSON(&req); err != nil {
http.ErrorData(c, "参数解析失败", nil)
return
}
if req.SessionID == 0 {
http.ErrorData(c, "sessionId不能为空", nil)
return
}

user, err := auth.GetCurrentUser(c)
if err != nil {
http.ErrorData(c, "未登录", nil)
return
}

messages, err := ctrl.conversationSvr.GetHistory(user.ID, req.SessionID)
if err != nil {
http.ErrorData(c, err.Error(), nil)
return
}
http.SuccessData(c, "ok", messages)
}

// GetSessions returns all active conversation sessions for the current user.
// @Summary      List user conversation sessions
// @Description  Returns all active AI conversation sessions for the authenticated user
// @Tags         Conversation
// @Accept       json
// @Produce      json
// @Success      200   {object}  map[string]interface{}  "Sessions"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Router       /v1/ai/conversation/sessions [post]
func (ctrl *ConversationController) GetSessions(c *gin.Context) {
user, err := auth.GetCurrentUser(c)
if err != nil {
http.ErrorData(c, "未登录", nil)
return
}

sessions, err := ctrl.conversationSvr.GetSessions(user.ID)
if err != nil {
http.ErrorData(c, err.Error(), nil)
return
}
http.SuccessData(c, "ok", sessions)
}

// ResetSession clears all messages in a conversation session.
// @Summary      Reset a conversation session
// @Description  Clears all messages in the specified conversation session
// @Tags         Conversation
// @Accept       json
// @Produce      json
// @Param        body  body  model.ConversationResetRequest  true  "Session ID"
// @Success      200   {object}  map[string]interface{}  "Session reset"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Router       /v1/ai/conversation/reset [post]
func (ctrl *ConversationController) ResetSession(c *gin.Context) {
var req model.ConversationResetRequest
if err := c.BindJSON(&req); err != nil {
http.ErrorData(c, "参数解析失败", nil)
return
}
if req.SessionID == 0 {
http.ErrorData(c, "sessionId不能为空", nil)
return
}

user, err := auth.GetCurrentUser(c)
if err != nil {
http.ErrorData(c, "未登录", nil)
return
}

if err := ctrl.conversationSvr.ResetSession(user.ID, req.SessionID); err != nil {
http.ErrorData(c, err.Error(), nil)
return
}
http.SuccessData(c, "会话已重置", nil)
}

// CloseSession deactivates a conversation session.
// @Summary      Close a conversation session
// @Description  Marks a conversation session as inactive
// @Tags         Conversation
// @Accept       json
// @Produce      json
// @Param        body  body  model.ConversationCloseRequest  true  "Session ID"
// @Success      200   {object}  map[string]interface{}  "Session closed"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Router       /v1/ai/conversation/close [post]
func (ctrl *ConversationController) CloseSession(c *gin.Context) {
var req model.ConversationCloseRequest
if err := c.BindJSON(&req); err != nil {
http.ErrorData(c, "参数解析失败", nil)
return
}
if req.SessionID == 0 {
http.ErrorData(c, "sessionId不能为空", nil)
return
}

user, err := auth.GetCurrentUser(c)
if err != nil {
http.ErrorData(c, "未登录", nil)
return
}

if err := ctrl.conversationSvr.CloseSession(user.ID, req.SessionID); err != nil {
http.ErrorData(c, err.Error(), nil)
return
}
http.SuccessData(c, "会话已关闭", nil)
}
