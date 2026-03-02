package v1

import (
	"edu/lib/net/http/middleware/auth"
	http2 "edu/lib/net/http"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

func init() {
	RBACCtrl = &RBACController{
		adminSvr: service.AdminSvr,
	}
}

var RBACCtrl *RBACController

type RBACController struct {
	adminSvr *service.AdminService
}

// requireAdmin checks that the current user is a system admin.
func requireAdmin(c *gin.Context) (uint, bool) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http2.ErrorData(c, "无法获取当前用户信息", err.Error())
		return 0, false
	}
	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || user == nil || !user.IsAdmin {
		http2.ErrorData(c, "需要管理员权限", nil)
		return 0, false
	}
	return u.ID, true
}

// ============ Roles ============

// ListRoles lists all roles with their permissions.
func (ctrl *RBACController) ListRoles(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	roles, err := ctrl.adminSvr.ListAdminRoles()
	if err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "获取成功", gin.H{"list": roles, "total": len(roles)})
}

// GetRole gets a role by ID.
func (ctrl *RBACController) GetRole(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	role, err := ctrl.adminSvr.SelectAdminRoleById(req.ID)
	if err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "获取成功", role)
}

// CreateRole creates a new role.
func (ctrl *RBACController) CreateRole(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	var role model.AdminRole
	if err := c.BindJSON(&role); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	created, err := ctrl.adminSvr.CreateAdminRole(role)
	if err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "创建成功", created)
}

// UpdateRole updates a role.
func (ctrl *RBACController) UpdateRole(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	var role model.AdminRole
	if err := c.BindJSON(&role); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	updated, err := ctrl.adminSvr.UpdateAdminRole(role)
	if err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "更新成功", updated)
}

// DeleteRole deletes a role by ID.
func (ctrl *RBACController) DeleteRole(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.adminSvr.DeleteAdminRole(req.ID); err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "删除成功", nil)
}

// ============ Permissions ============

// ListPermissions lists all permissions.
func (ctrl *RBACController) ListPermissions(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	perms, err := ctrl.adminSvr.ListPermissions()
	if err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "获取成功", gin.H{"list": perms, "total": len(perms)})
}

// CreatePermission creates a new permission.
func (ctrl *RBACController) CreatePermission(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	var perm model.AdminPermission
	if err := c.BindJSON(&perm); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	created, err := ctrl.adminSvr.CreatePermission(perm)
	if err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "创建成功", created)
}

// UpdatePermission updates a permission.
func (ctrl *RBACController) UpdatePermission(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	var perm model.AdminPermission
	if err := c.BindJSON(&perm); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	updated, err := ctrl.adminSvr.UpdatePermission(perm)
	if err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "更新成功", updated)
}

// DeletePermission deletes a permission by ID.
func (ctrl *RBACController) DeletePermission(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.adminSvr.DeletePermission(req.ID); err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "删除成功", nil)
}

// ============ Role-Permission Assignment ============

// AssignPermissionToRole assigns a permission to a role.
func (ctrl *RBACController) AssignPermissionToRole(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	var req struct {
		RoleID       uint `json:"roleId" binding:"required"`
		PermissionID uint `json:"permissionId" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.adminSvr.AssignPermissionToRole(req.RoleID, req.PermissionID); err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "权限分配成功", nil)
}

// RemovePermissionFromRole removes a permission from a role.
func (ctrl *RBACController) RemovePermissionFromRole(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	var req struct {
		RoleID       uint `json:"roleId" binding:"required"`
		PermissionID uint `json:"permissionId" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.adminSvr.RemovePermissionFromRole(req.RoleID, req.PermissionID); err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "权限移除成功", nil)
}

// ============ User-Role Assignment ============

// GetUserRoles gets the roles assigned to a user.
func (ctrl *RBACController) GetUserRoles(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	var req struct {
		UserID uint `json:"userId" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	roles, err := ctrl.adminSvr.GetUserRoles(req.UserID)
	if err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "获取成功", gin.H{"list": roles, "total": len(roles)})
}

// AssignRoleToUser assigns a role to a user.
func (ctrl *RBACController) AssignRoleToUser(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	var req struct {
		UserID uint `json:"userId" binding:"required"`
		RoleID uint `json:"roleId" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.adminSvr.AssignRoleToUser(req.UserID, req.RoleID); err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "角色分配成功", nil)
}

// RemoveRoleFromUser removes a role from a user.
func (ctrl *RBACController) RemoveRoleFromUser(c *gin.Context) {
	if _, ok := requireAdmin(c); !ok {
		return
	}
	var req struct {
		UserID uint `json:"userId" binding:"required"`
		RoleID uint `json:"roleId" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.adminSvr.RemoveRoleFromUser(req.UserID, req.RoleID); err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "角色移除成功", nil)
}

// GetMyPermissions returns the permissions of the currently authenticated user.
func (ctrl *RBACController) GetMyPermissions(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http2.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}
	perms, err := ctrl.adminSvr.GetUserPermissions(u.ID)
	if err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "获取成功", gin.H{"list": perms, "total": len(perms)})
}

// CheckPermission checks whether the current user has a given permission.
func (ctrl *RBACController) CheckPermission(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http2.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}
	var req struct {
		Permission string `json:"permission" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http2.ErrorData(c, "参数解析失败", nil)
		return
	}
	has, err := ctrl.adminSvr.HasPermission(u.ID, req.Permission)
	if err != nil {
		http2.ErrorData(c, err.Error(), nil)
		return
	}
	http2.SuccessData(c, "检查完成", gin.H{"hasPermission": has})
}
