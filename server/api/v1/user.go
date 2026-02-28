package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var UserCtrl = &UserController{
	userSvr: service.UserSvr,
	// adminSvr: service.AdminSvr, // Removed as it's no longer used
}

type UserController struct {
	userSvr *service.UserService
	// adminSvr *service.AdminService, // Removed as it's no longer used
}

// @Summary      设置管理员
// @Description  将指定用户设置为管理员
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "用户ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/user/setAdmin [post]
// 设置为管理员
func (ctrl *UserController) SetAdmin(c *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}
	err := ctrl.userSvr.SetAdmin(req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "设置管理员成功!", nil)
}

// @Summary      取消管理员
// @Description  取消指定用户的管理员权限
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "用户ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/user/removeAdmin [post]
// 取消管理员
func (ctrl *UserController) RemoveAdmin(c *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.userSvr.RevokeUserAdminRole(req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "取消管理员成功!", nil)
}

// User管理
// @Summary      获取用户列表
// @Description  分页查询用户列表
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.UserQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/user/list [post]
func (ctrl *UserController) SelectUserList(c *gin.Context) {
	q := model.UserQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.userSvr.SelectUserList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

// @Summary      根据ID获取用户
// @Description  根据用户ID获取用户详情
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.UserQuery  true  "用户ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/user/byId [post]
func (ctrl *UserController) SelectUserById(c *gin.Context) {
	q := model.UserQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.userSvr.SelectUserById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

// @Summary      根据用户名获取用户
// @Description  根据用户名查询用户信息
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.UserQuery  true  "用户名"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/user/byUsername [post]
func (ctrl *UserController) SelectUserByUsername(c *gin.Context) {
	q := model.UserQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.userSvr.SelectUserByUsername(q.Username)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

// @Summary      获取所有用户
// @Description  获取全部用户列表（不分页）
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.UserQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/user/all [post]
func (ctrl *UserController) SelectUserAll(c *gin.Context) {
	oq := model.UserQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.userSvr.SelectUserAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

// @Summary      创建用户
// @Description  创建新用户（默认密码123456）
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.User  true  "用户信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/user/create [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	o := model.User{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o.Password = "123456"
	r, err := ctrl.userSvr.CreateUser(o)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", r)
}

// @Summary      编辑用户
// @Description  修改用户信息
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.UserEditRequest  true  "用户信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/user/edit [post]
func (ctrl *UserController) EditUser(c *gin.Context) {
	o := model.UserEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	// Restrict account status modification to super admin only
	if o.Status != 0 {
		u, err := auth.GetCurrentUser(c)
		if err != nil {
			http.ErrorData(c, "无法获取当前用户信息", nil)
			return
		}
		currentUser, err := service.UserSvr.SelectUserById(u.ID)
		if err != nil {
			http.ErrorData(c, "无法验证管理员权限", nil)
			return
		}
		if currentUser == nil || !currentUser.IsAdmin {
			http.ErrorData(c, "只有超级管理员可以修改用户账户状态", nil)
			return
		}
	}

	err := ctrl.userSvr.EditUser(o)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// @Summary      删除用户
// @Description  删除指定用户
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.User  true  "用户ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/user/delete [post]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	o := model.User{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.userSvr.DeleteUser(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// Grant one month VIP to user
// @Summary      授予VIP
// @Description  为指定用户授予一个月VIP（仅管理员）
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "用户ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/user/vip [post]
func (ctrl *UserController) GrantVipMonth(c *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	// 权限验证：仅管理员可调用
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}
	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || user == nil || !user.IsAdmin {
		http.ErrorData(c, "No permission, only admin can grant VIP", nil)
		return
	}
	err = ctrl.userSvr.GrantVipMonth(req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "VIP granted for 1 month!", nil)
}

// @Summary      获取当前用户信息
// @Description  获取当前已登录用户的详细信息
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Failure      400  {object}  map[string]interface{}  "未登录"
// @Security     BearerAuth
// @Router       /v1/user [get]
func (ac *UserController) User(c *gin.Context) {
	// Get UserID from JWT claims (auth.CurrentUser contains the ID)
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "获取用户信息失败: "+err.Error(), nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", user)
}

// @Summary      更新当前用户账户信息
// @Description  更新当前已登录用户的个人信息（不包含状态等管理员字段）
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.AccountUpdateRequest  true  "账户更新信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/account/update [post]
func (ac *UserController) UpdateOwnAccount(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	var req model.AccountUpdateRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ac.userSvr.UpdateOwnAccount(u.ID, req); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "账户信息更新成功", nil)
}
