package v1

import (
"edu/lib/net/http"
"edu/lib/net/http/middleware/auth"
"edu/model"
"edu/service"

"github.com/gin-gonic/gin"
)

func init() {
ConversationCtrl = &ConversationController{}
}

var ConversationCtrl *ConversationController

// ConversationController handles multi-turn AI conversation endpoints.
type ConversationController struct{}

// StartSession begins a new AI conversation session for the authenticated user.
// @Summary      Start a new AI conversation session
// @Description  Creates a new conversation session with optional user context (role, preferences, recent actions).
// @Tags         AI Conversation
// @Accept       json
// @Produce      json
// @Param        body  body  model.StartSessionRequest  false  "Optional initial context"
// @Success      200   {object}  map[string]interface{}  "Session created"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Security     BearerAuth
// @Router       /v1/ai/conversation/start [post]
func (ctrl *ConversationController) StartSession(c *gin.Context) {
currentUser, err := auth.GetCurrentUser(c)
if err != nil {
http.ErrorData(c, "未登录", err.Error())
return
}

var req model.StartSessionRequest
_ = c.ShouldBindJSON(&req) // optional body

// Determine role from user record (best-effort)
user, _ := service.UserSvr.SelectUserById(currentUser.ID)
userRole := "student"
userName := "User"
if user != nil {
userName = user.Username
if user.HasAdminRole() {
userRole = "admin"
} else if user.HasTeacherRole() {
userRole = "teacher"
}
}

// Override with caller-supplied role if provided
if req.Context != nil && req.Context.UserRole != "" {
userRole = req.Context.UserRole
}

session, err := service.ConversationSvr.StartSession(currentUser.ID, userName, userRole, req.Context)
if err != nil {
http.ErrorData(c, "创建会话失败", err.Error())
return
}

resp := model.ConversationSessionResponse{
SessionKey:   session.SessionKey,
UserRole:     session.UserRole,
MessageCount: session.MessageCount,
LastActiveAt: session.LastActiveAt,
ExpiresAt:    session.ExpiresAt,
IsActive:     session.IsActive,
CreatedAt:    session.CreatedAt,
}
http.SuccessData(c, "会话创建成功", resp)
}

// SendMessage sends a message in an existing conversation session.
// @Summary      Send a message in a conversation
// @Description  Sends a user message to the AI and returns the assistant's reply. Maintains full conversation context.
// @Tags         AI Conversation
// @Accept       json
// @Produce      json
// @Param        body  body  model.SendMessageRequest  true  "Message payload"
// @Success      200   {object}  map[string]interface{}  "Message sent"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Security     BearerAuth
// @Router       /v1/ai/conversation/message [post]
func (ctrl *ConversationController) SendMessage(c *gin.Context) {
currentUser, err := auth.GetCurrentUser(c)
if err != nil {
http.ErrorData(c, "未登录", err.Error())
return
}

var req model.SendMessageRequest
if err := c.ShouldBindJSON(&req); err != nil {
http.ErrorData(c, "参数错误", err.Error())
return
}

// Resolve user's display name for the system prompt
user, _ := service.UserSvr.SelectUserById(currentUser.ID)
userName := "User"
if user != nil {
userName = user.Username
}

resp, err := service.ConversationSvr.SendMessage(currentUser.ID, userName, req.SessionKey, req.Message)
if err != nil {
http.ErrorData(c, "发送消息失败", err.Error())
return
}

http.SuccessData(c, "消息发送成功", resp)
}

// GetHistory retrieves all messages in a conversation session.
// @Summary      Get conversation history
// @Description  Returns the full message history for the specified conversation session.
// @Tags         AI Conversation
// @Accept       json
// @Produce      json
// @Param        body  body  model.ConversationHistoryRequest  true  "Session key"
// @Success      200   {object}  map[string]interface{}  "History retrieved"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Security     BearerAuth
// @Router       /v1/ai/conversation/history [post]
func (ctrl *ConversationController) GetHistory(c *gin.Context) {
currentUser, err := auth.GetCurrentUser(c)
if err != nil {
http.ErrorData(c, "未登录", err.Error())
return
}

var req model.ConversationHistoryRequest
if err := c.ShouldBindJSON(&req); err != nil {
http.ErrorData(c, "参数错误", err.Error())
return
}

messages, err := service.ConversationSvr.GetHistory(currentUser.ID, req.SessionKey)
if err != nil {
http.ErrorData(c, "获取历史记录失败", err.Error())
return
}

http.SuccessData(c, "数据获取成功", messages)
}

// GetSessions returns all active conversation sessions for the current user.
// @Summary      List conversation sessions
// @Description  Returns a list of all active AI conversation sessions for the authenticated user.
// @Tags         AI Conversation
// @Produce      json
// @Success      200   {object}  map[string]interface{}  "Sessions listed"
// @Security     BearerAuth
// @Router       /v1/ai/conversation/sessions [post]
func (ctrl *ConversationController) GetSessions(c *gin.Context) {
currentUser, err := auth.GetCurrentUser(c)
if err != nil {
http.ErrorData(c, "未登录", err.Error())
return
}

sessions, err := service.ConversationSvr.ListSessions(currentUser.ID)
if err != nil {
http.ErrorData(c, "获取会话列表失败", err.Error())
return
}

var responses []model.ConversationSessionResponse
for _, s := range sessions {
responses = append(responses, model.ConversationSessionResponse{
SessionKey:   s.SessionKey,
UserRole:     s.UserRole,
MessageCount: s.MessageCount,
LastActiveAt: s.LastActiveAt,
ExpiresAt:    s.ExpiresAt,
IsActive:     s.IsActive,
CreatedAt:    s.CreatedAt,
})
}

http.SuccessData(c, "数据获取成功", responses)
}

// ResetSession clears the history and optionally updates context for a session.
// @Summary      Reset conversation context
// @Description  Clears the message history for the specified session. Optionally accepts new context data.
// @Tags         AI Conversation
// @Accept       json
// @Produce      json
// @Param        body  body  model.ResetContextRequest  true  "Session key and optional new context"
// @Success      200   {object}  map[string]interface{}  "Context reset"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Security     BearerAuth
// @Router       /v1/ai/conversation/reset [post]
func (ctrl *ConversationController) ResetSession(c *gin.Context) {
currentUser, err := auth.GetCurrentUser(c)
if err != nil {
http.ErrorData(c, "未登录", err.Error())
return
}

var req model.ResetContextRequest
if err := c.ShouldBindJSON(&req); err != nil {
http.ErrorData(c, "参数错误", err.Error())
return
}

if err := service.ConversationSvr.ResetContext(currentUser.ID, req.SessionKey, req.Context); err != nil {
http.ErrorData(c, "重置上下文失败", err.Error())
return
}

http.SuccessData(c, "上下文已重置", nil)
}

// CloseSession deactivates a conversation session.
// @Summary      Close a conversation session
// @Description  Deactivates the specified conversation session for the current user.
// @Tags         AI Conversation
// @Accept       json
// @Produce      json
// @Param        body  body  model.ConversationHistoryRequest  true  "Session key"
// @Success      200   {object}  map[string]interface{}  "Session closed"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Security     BearerAuth
// @Router       /v1/ai/conversation/close [post]
func (ctrl *ConversationController) CloseSession(c *gin.Context) {
currentUser, err := auth.GetCurrentUser(c)
if err != nil {
http.ErrorData(c, "未登录", err.Error())
return
}

var req model.ConversationHistoryRequest
if err := c.ShouldBindJSON(&req); err != nil {
http.ErrorData(c, "参数错误", err.Error())
return
}

if err := service.ConversationSvr.CloseSession(currentUser.ID, req.SessionKey); err != nil {
http.ErrorData(c, "关闭会话失败", err.Error())
return
}

http.SuccessData(c, "会话已关闭", nil)
}
