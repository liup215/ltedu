package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	MCPTokenCtrl = &MCPTokenController{}
}

var MCPTokenCtrl *MCPTokenController

type MCPTokenController struct{}

// CreateToken creates a new MCP token for the current user
// @Summary      创建MCP令牌
// @Description  为当前用户创建新的MCP访问令牌
// @Tags         MCP
// @Accept       json
// @Produce      json
// @Param        body  body  model.MCPTokenCreateRequest  true  "令牌信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/mcp/token/create [post]
func (ctrl *MCPTokenController) CreateToken(c *gin.Context) {
	var req model.MCPTokenCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "参数错误", err.Error())
		return
	}

	// Get current user
	currentUser, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "未登录", err.Error())
		return
	}

	// Parse expiration date if provided
	var expiresAt time.Time
	if req.ExpiresAt != "" {
		expiresAt, err = time.Parse(time.RFC3339, req.ExpiresAt)
		if err != nil {
			http.ErrorData(c, "过期时间格式错误", "请使用ISO 8601格式（如：2025-12-31T23:59:59Z）")
			return
		}
	}

	// Create token
	token, err := service.MCPTokenSvr.CreateToken(currentUser.ID, req.Name, expiresAt)
	if err != nil {
		http.ErrorData(c, "创建MCP令牌失败", err.Error())
		return
	}

	// Return response
	response := model.MCPTokenResponse{
		ID:        token.ID,
		Token:     token.Token,
		Name:      token.Name,
		ExpiresAt: token.ExpiresAt,
		IsActive:  token.IsActive,
		CreatedAt: token.CreatedAt,
		LastUsed:  token.LastUsed,
	}

	http.SuccessData(c, "MCP令牌创建成功", response)
}

// ListTokens lists all MCP tokens for the current user
// @Summary      获取MCP令牌列表
// @Description  获取当前用户的所有MCP令牌
// @Tags         MCP
// @Accept       json
// @Produce      json
// @Param        body  body  model.Page  false  "分页参数"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/mcp/token/list [post]
func (ctrl *MCPTokenController) ListTokens(c *gin.Context) {
	var req model.Page
	if err := c.ShouldBindJSON(&req); err != nil {
		req = model.Page{PageIndex: 1, PageSize: 20}
	}

	// Get current user
	currentUser, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "未登录", err.Error())
		return
	}

	// Get tokens
	tokens, total, err := service.MCPTokenSvr.ListUserTokens(currentUser.ID, req)
	if err != nil {
		http.ErrorData(c, "获取MCP令牌列表失败", err.Error())
		return
	}

	// Convert to response format
	var responses []model.MCPTokenResponse
	for _, token := range tokens {
		responses = append(responses, model.MCPTokenResponse{
			ID:        token.ID,
			Token:     token.Token,
			Name:      token.Name,
			ExpiresAt: token.ExpiresAt,
			IsActive:  token.IsActive,
			CreatedAt: token.CreatedAt,
			LastUsed:  token.LastUsed,
		})
	}

	http.SuccessData(c, "获取MCP令牌列表成功", gin.H{
		"total": total,
		"list":  responses,
	})
}

// DeleteToken deletes an MCP token
// @Summary      删除MCP令牌
// @Description  删除当前用户的指定MCP令牌
// @Tags         MCP
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "令牌ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/mcp/token/delete [post]
func (ctrl *MCPTokenController) DeleteToken(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "参数错误", err.Error())
		return
	}

	// Get current user
	currentUser, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "未登录", err.Error())
		return
	}

	// Delete token
	err = service.MCPTokenSvr.DeleteToken(req.ID, currentUser.ID)
	if err != nil {
		http.ErrorData(c, "删除MCP令牌失败", err.Error())
		return
	}

	http.SuccessData(c, "MCP令牌删除成功", nil)
}

// DeactivateToken deactivates an MCP token
// @Summary      停用MCP令牌
// @Description  停用当前用户的指定MCP令牌
// @Tags         MCP
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "令牌ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/mcp/token/deactivate [post]
func (ctrl *MCPTokenController) DeactivateToken(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "参数错误", err.Error())
		return
	}

	// Get current user
	currentUser, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "未登录", err.Error())
		return
	}

	// Deactivate token
	err = service.MCPTokenSvr.DeactivateToken(req.ID, currentUser.ID)
	if err != nil {
		http.ErrorData(c, "停用MCP令牌失败", err.Error())
		return
	}

	http.SuccessData(c, "MCP令牌已停用", nil)
}

// ActivateToken activates an MCP token
// @Summary      激活MCP令牌
// @Description  激活当前用户的指定MCP令牌
// @Tags         MCP
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "令牌ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/mcp/token/activate [post]
func (ctrl *MCPTokenController) ActivateToken(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "参数错误", err.Error())
		return
	}

	// Get current user
	currentUser, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "未登录", err.Error())
		return
	}

	// Activate token
	err = service.MCPTokenSvr.ActivateToken(req.ID, currentUser.ID)
	if err != nil {
		http.ErrorData(c, "激活MCP令牌失败", err.Error())
		return
	}

	http.SuccessData(c, "MCP令牌已激活", nil)
}

// AdminListTokens lists all MCP tokens (admin only)
// @Summary      管理员获取所有MCP令牌
// @Description  管理员获取全平台所有用户的MCP令牌列表
// @Tags         MCP
// @Accept       json
// @Produce      json
// @Param        body  body  model.MCPTokenQuery  false  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/mcp/token/admin/list [post]
func (ctrl *MCPTokenController) AdminListTokens(c *gin.Context) {
	// Check admin permission
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "未登录", err.Error())
		return
	}
	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || user == nil || !user.IsAdmin {
		http.ForbiddenData(c, "无权限访问此资源", nil)
		return
	}

	var req model.MCPTokenQuery
	if err := c.ShouldBindJSON(&req); err != nil {
		req = model.MCPTokenQuery{Page: model.Page{PageIndex: 1, PageSize: 20}}
	}

	// Get tokens
	tokens, total, err := service.MCPTokenSvr.ListAllTokens(req)
	if err != nil {
		http.ErrorData(c, "获取MCP令牌列表失败", err.Error())
		return
	}

	// Convert to response format with user info
	type AdminTokenResponse struct {
		ID        uint      `json:"id"`
		UserID    uint      `json:"userId"`
		Username  string    `json:"username"`
		Token     string    `json:"token"`
		Name      string    `json:"name"`
		ExpiresAt time.Time `json:"expiresAt"`
		IsActive  bool      `json:"isActive"`
		CreatedAt time.Time `json:"createdAt"`
		LastUsed  time.Time `json:"lastUsed"`
	}

	var responses []AdminTokenResponse
	for _, token := range tokens {
		responses = append(responses, AdminTokenResponse{
			ID:        token.ID,
			UserID:    token.UserID,
			Username:  token.User.Username,
			Token:     token.Token,
			Name:      token.Name,
			ExpiresAt: token.ExpiresAt,
			IsActive:  token.IsActive,
			CreatedAt: token.CreatedAt,
			LastUsed:  token.LastUsed,
		})
	}

	http.SuccessData(c, "获取MCP令牌列表成功", gin.H{
		"total": total,
		"list":  responses,
	})
}

// AdminDeleteToken deletes any MCP token (admin only)
// @Summary      管理员删除MCP令牌
// @Description  管理员删除任意用户的MCP令牌
// @Tags         MCP
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "令牌ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/mcp/token/admin/delete [post]
func (ctrl *MCPTokenController) AdminDeleteToken(c *gin.Context) {
	// Check admin permission
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "未登录", err.Error())
		return
	}
	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || user == nil || !user.IsAdmin {
		http.ForbiddenData(c, "无权限访问此资源", nil)
		return
	}

	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "参数错误", err.Error())
		return
	}

	err = service.MCPTokenSvr.AdminDeleteToken(req.ID)
	if err != nil {
		http.ErrorData(c, "删除MCP令牌失败", err.Error())
		return
	}

	http.SuccessData(c, "MCP令牌删除成功", nil)
}

// AdminDeactivateToken deactivates any MCP token (admin only)
// @Summary      管理员停用MCP令牌
// @Description  管理员停用任意用户的MCP令牌
// @Tags         MCP
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "令牌ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/mcp/token/admin/deactivate [post]
func (ctrl *MCPTokenController) AdminDeactivateToken(c *gin.Context) {
	// Check admin permission
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "未登录", err.Error())
		return
	}
	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || user == nil || !user.IsAdmin {
		http.ForbiddenData(c, "无权限访问此资源", nil)
		return
	}

	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "参数错误", err.Error())
		return
	}

	err = service.MCPTokenSvr.AdminDeactivateToken(req.ID)
	if err != nil {
		http.ErrorData(c, "停用MCP令牌失败", err.Error())
		return
	}

	http.SuccessData(c, "MCP令牌已停用", nil)
}

// AdminActivateToken activates any MCP token (admin only)
// @Summary      管理员激活MCP令牌
// @Description  管理员激活任意用户的MCP令牌
// @Tags         MCP
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "令牌ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/mcp/token/admin/activate [post]
func (ctrl *MCPTokenController) AdminActivateToken(c *gin.Context) {
	// Check admin permission
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "未登录", err.Error())
		return
	}
	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || user == nil || !user.IsAdmin {
		http.ForbiddenData(c, "无权限访问此资源", nil)
		return
	}

	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "参数错误", err.Error())
		return
	}

	err = service.MCPTokenSvr.AdminActivateToken(req.ID)
	if err != nil {
		http.ErrorData(c, "激活MCP令牌失败", err.Error())
		return
	}

	http.SuccessData(c, "MCP令牌已激活", nil)
}
