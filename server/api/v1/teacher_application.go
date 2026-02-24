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
// @Summary      申请成为教师
// @Description  提交教师资格申请
// @Tags         教师申请
// @Accept       json
// @Produce      json
// @Param        body  body  model.TeacherApplicationCreateRequest  true  "申请信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/user/teacher/apply [post]
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
// @Summary      获取我的教师申请
// @Description  获取当前用户的教师资格申请详情
// @Tags         教师申请
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Failure      400  {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/user/teacher/application [post]
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
// @Summary      获取教师申请列表
// @Description  获取所有教师申请列表（仅管理员）
// @Tags         教师申请
// @Accept       json
// @Produce      json
// @Param        body  body  model.TeacherApplicationQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/admin/teacher-applications/list [post]
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
// @Summary      获取教师申请详情
// @Description  获取指定教师申请的详细信息（仅管理员）
// @Tags         教师申请
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "申请ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/admin/teacher-applications/detail [post]
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
// @Summary      审批通过教师申请
// @Description  通过指定的教师资格申请（仅管理员）
// @Tags         教师申请
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "申请ID和备注"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/admin/teacher-applications/approve [post]
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
// @Summary      拒绝教师申请
// @Description  拒绝指定的教师资格申请（仅管理员）
// @Tags         教师申请
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "申请ID和拒绝理由"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/admin/teacher-applications/reject [post]
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
