package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var SchoolCtrl = &SchoolController{
	schoolSvr: service.SchoolSvr,
}

type SchoolController struct {
	schoolSvr *service.SchoolService
}

// Grade管理

// @Summary      获取年级列表
// @Description  分页查询年级列表
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.GradeQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/grade/list [post]
func (ctrl *SchoolController) SelectGradeList(c *gin.Context) {
	q := model.GradeQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.schoolSvr.SelectGradeList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

// @Summary      根据ID获取年级
// @Description  根据年级ID获取年级详情
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.GradeQuery  true  "年级ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/grade/byId [post]
func (ctrl *SchoolController) SelectGradeById(c *gin.Context) {
	q := model.GradeQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.schoolSvr.SelectGradeById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

// @Summary      获取所有年级
// @Description  获取全部年级列表（不分页）
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.GradeQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/grade/all [post]
func (ctrl *SchoolController) SelectGradeAll(c *gin.Context) {
	oq := model.GradeQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.schoolSvr.SelectGradeAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

// @Summary      创建年级
// @Description  创建新年级
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.GradeCreateEditRequest  true  "年级信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/grade/create [post]
func (ctrl *SchoolController) CreateGrade(c *gin.Context) {
	o := model.GradeCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败: "+err.Error(), nil)
		return
	}
	r, err := ctrl.schoolSvr.CreateGrade(o)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", r)
}

// @Summary      编辑年级
// @Description  修改年级信息
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.GradeCreateEditRequest  true  "年级信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/grade/edit [post]
func (ctrl *SchoolController) EditGrade(c *gin.Context) {
	o := model.GradeCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败: "+err.Error(), nil)
		return
	}
	err := ctrl.schoolSvr.EditGrade(o)
	if err != nil {
		http.ErrorData(c, "更新失败："+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// @Summary      删除年级
// @Description  删除指定年级
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Grade  true  "年级ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/grade/delete [post]
func (ctrl *SchoolController) DeleteGrade(c *gin.Context) {
	o := model.Grade{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.DeleteGrade(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// ClassType管理

// @Summary      获取班型列表
// @Description  分页查询班型列表
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ClassTypeQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/classType/list [post]
func (ctrl *SchoolController) SelectClassTypeList(c *gin.Context) {
	q := model.ClassTypeQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.schoolSvr.SelectClassTypeList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

// @Summary      根据ID获取班型
// @Description  根据班型ID获取班型详情
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ClassTypeQuery  true  "班型ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/classType/byId [post]
func (ctrl *SchoolController) SelectClassTypeById(c *gin.Context) {
	q := model.ClassTypeQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.schoolSvr.SelectClassTypeById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

// @Summary      获取所有班型
// @Description  获取全部班型列表（不分页）
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ClassTypeQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/classType/all [post]
func (ctrl *SchoolController) SelectClassTypeAll(c *gin.Context) {
	oq := model.ClassTypeQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.schoolSvr.SelectClassTypeAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

// @Summary      创建班型
// @Description  创建新班型
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ClassTypeCreateEditRequest  true  "班型信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/classType/create [post]
func (ctrl *SchoolController) CreateClassType(c *gin.Context) {
	o := model.ClassTypeCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.CreateClassType(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// @Summary      编辑班型
// @Description  修改班型信息
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ClassTypeCreateEditRequest  true  "班型信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/classType/edit [post]
func (ctrl *SchoolController) EditClassType(c *gin.Context) {
	o := model.ClassTypeCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.EditClassType(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// @Summary      删除班型
// @Description  删除指定班型
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ClassType  true  "班型ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/classType/delete [post]
func (ctrl *SchoolController) DeleteClassType(c *gin.Context) {
	o := model.ClassType{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.DeleteClassType(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// Class管理

// @Summary      获取班级列表
// @Description  分页查询班级列表
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ClassQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/class/list [post]
func (ctrl *SchoolController) SelectClassList(c *gin.Context) {
	q := model.ClassQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.schoolSvr.SelectClassList(q)
	if err != nil {
		http.ErrorData(c, "数据获取失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

// @Summary      根据ID获取班级
// @Description  根据班级ID获取班级详情
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ClassQuery  true  "班级ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/class/byId [post]
func (ctrl *SchoolController) SelectClassById(c *gin.Context) {
	q := model.ClassQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.schoolSvr.SelectClassById(q.ID)
	if err != nil {
		http.ErrorData(c, "数据获取失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

// @Summary      获取所有班级
// @Description  获取全部班级列表（不分页）
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ClassQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/class/all [post]
func (ctrl *SchoolController) SelectClassAll(c *gin.Context) {
	oq := model.ClassQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.schoolSvr.SelectClassAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

// @Summary      创建班级
// @Description  创建新班级（仅教师可创建，创建者自动成为管理员，自动生成邀请码）
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ClassCreateEditRequest  true  "班级信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/class/create [post]
func (ctrl *SchoolController) CreateClass(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", nil)
		return
	}
	o := model.ClassCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	class, err := ctrl.schoolSvr.CreateClass(o, u.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "班级创建成功!", class)
}

// @Summary      编辑班级
// @Description  修改班级信息
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ClassCreateEditRequest  true  "班级信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/class/edit [post]
func (ctrl *SchoolController) EditClass(c *gin.Context) {
	o := model.ClassCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.EditClass(o)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// @Summary      删除班级
// @Description  删除指定班级
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Class  true  "班级ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/class/delete [post]
func (ctrl *SchoolController) DeleteClass(c *gin.Context) {
	o := model.Class{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.DeleteClass(o.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据删除成功!", nil)
}

// @Summary      获取班级学生列表
// @Description  获取指定班级的所有学生
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Class  true  "班级ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/class/studentList [post]
func (ctrl *SchoolController) GetStudentsByClassId(c *gin.Context) {
	o := model.Class{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, err := ctrl.schoolSvr.GetStudentsByClassId(o.ID)
	if err != nil {
		http.ErrorData(c, "数据获取失败:"+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

type RemoveStudentFromClassRequest struct {
	ClassId uint `json:"classId"`
	UserId  uint `json:"userId"`
}

// @Summary      从班级移除学生
// @Description  将学生从指定班级移除（管理员操作）
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  RemoveStudentFromClassRequest  true  "班级ID和学生ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/class/removeStudent [post]
func (ctrl *SchoolController) DeleteStudentFromClass(c *gin.Context) {
	o := RemoveStudentFromClassRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.DeleteStudentFromClass(o.ClassId, o.UserId)
	if err != nil {
		http.ErrorData(c, "数据获取失败:"+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// @Summary      申请加入班级
// @Description  学生使用邀请码申请加入班级
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "邀请码和消息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/class/apply [post]
func (ctrl *SchoolController) ApplyToJoinClass(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", nil)
		return
	}
	var req struct {
		InviteCode string `json:"inviteCode" binding:"required"`
		Message    string `json:"message"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	joinReq, err := ctrl.schoolSvr.ApplyToJoinClass(req.InviteCode, u.ID, req.Message)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "申请已提交，等待管理员审核", joinReq)
}

// @Summary      获取加入申请列表
// @Description  获取班级的加入申请列表（管理员）
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ClassJoinRequestQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/class/joinRequest/list [post]
func (ctrl *SchoolController) ListJoinRequests(c *gin.Context) {
	q := model.ClassJoinRequestQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.schoolSvr.ListClassJoinRequests(q)
	if err != nil {
		http.ErrorData(c, "数据获取失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

// @Summary      审核通过加入申请
// @Description  管理员审核通过学生的加入申请
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "申请ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/class/joinRequest/approve [post]
func (ctrl *SchoolController) ApproveJoinRequest(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", nil)
		return
	}
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.schoolSvr.ApproveJoinRequest(req.ID, u.ID); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "申请已通过", nil)
}

// @Summary      拒绝加入申请
// @Description  管理员拒绝学生的加入申请
// @Tags         学校管理
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "申请ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/school/class/joinRequest/reject [post]
func (ctrl *SchoolController) RejectJoinRequest(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", nil)
		return
	}
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.schoolSvr.RejectJoinRequest(req.ID, u.ID); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "申请已拒绝", nil)
}
