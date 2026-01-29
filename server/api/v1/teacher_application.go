package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var TeacherApplicationCtrl = &TeacherApplicationController{
	service: service.TeacherApplicationSvr,
}

type TeacherApplicationController struct {
	service *service.TeacherApplicationService
}

// Apply handles the submission of a teacher application
func (ctrl *TeacherApplicationController) Apply(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "用户未登录或认证信息无效", nil)
		return
	}

	var req model.TeacherApplicationCreateRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.service.Apply(user.ID, &req); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "申请提交成功！", nil)
}

// GetByUser gets the current user's teacher application
func (ctrl *TeacherApplicationController) GetByUser(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "用户未登录或认证信息无效", nil)
		return
	}

	application, err := ctrl.service.GetByUserID(user.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "数据获取成功！", application)
}

// List lists teacher applications (admin only)
func (ctrl *TeacherApplicationController) List(c *gin.Context) {
	// Check admin access
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "用户未登录或认证信息无效", nil)
		return
	}

	if !user.IsAdmin {
		http.ErrorData(c, "非管理员账户", nil)
		return
	}

	var query model.TeacherApplicationQuery
	if err := c.BindJSON(&query); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	applications, total, err := ctrl.service.List(&query)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "数据获取成功！", gin.H{
		"list":  applications,
		"total": total,
	})
}

// Get gets a specific teacher application by ID (admin only)
func (ctrl *TeacherApplicationController) Get(c *gin.Context) {
	// Check admin access
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "用户未登录或认证信息无效", nil)
		return
	}

	if !user.IsAdmin {
		http.ErrorData(c, "非管理员账户", nil)
		return
	}

	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	application, err := ctrl.service.GetByID(req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "数据获取成功！", application)
}

// Approve approves a teacher application (admin only)
func (ctrl *TeacherApplicationController) Approve(c *gin.Context) {
	// Check admin access
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "用户未登录或认证信息无效", nil)
		return
	}

	if !user.IsAdmin {
		http.ErrorData(c, "非管理员账户", nil)
		return
	}

	var req struct {
		ID         uint   `json:"id" binding:"required"`
		AdminNotes string `json:"adminNotes"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.service.Approve(req.ID, user.ID, req.AdminNotes); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "申请已通过！", nil)
}

// Reject rejects a teacher application (admin only)
func (ctrl *TeacherApplicationController) Reject(c *gin.Context) {
	// Check admin access
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "用户未登录或认证信息无效", nil)
		return
	}

	if !user.IsAdmin {
		http.ErrorData(c, "非管理员账户", nil)
		return
	}

	var req struct {
		ID         uint   `json:"id" binding:"required"`
		AdminNotes string `json:"adminNotes" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.service.Reject(req.ID, user.ID, req.AdminNotes); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "申请已拒绝！", nil)
}
