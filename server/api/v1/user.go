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

func (ctrl *UserController) EditUser(c *gin.Context) {
	o := model.UserEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	err := ctrl.userSvr.EditUser(o)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

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
