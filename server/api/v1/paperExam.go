package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

// ExamPaper 管理
// @Summary      获取考试试卷列表
// @Description  分页查询考试试卷列表（管理员可查全部，教师只查自己的）
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ExamPaperQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/paper/exam/list [post]
func (ctrl *ExamPaperController) SelectExamPaperList(c *gin.Context) {
	q := model.ExamPaperQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "Failed to parse parameters", nil)
		return
	}

	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Unable to get current user info", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "User not logged in", nil)
		return
	}

	if !user.IsAdmin && !user.IsTeacher {
		http.ErrorData(c, "No permission to access exam paper list", nil)
		return
	}

	// 如果不是管理员，只能查看自己创建的试卷
	if !user.IsAdmin {
		q.UserId = user.ID
	}

	list, total, err := ctrl.questionPaperSvr.SelectExamPaperList(q)
	if err != nil {
		http.ErrorData(c, "Failed to retrieve data", nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", gin.H{
		"list":  list,
		"total": total,
	})
}

// @Summary      根据ID获取考试试卷
// @Description  根据试卷ID获取考试试卷详情
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ExamPaperQuery  true  "试卷ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/paper/exam/byId [post]
func (ctrl *ExamPaperController) SelectExamPaperById(c *gin.Context) {
	q := model.ExamPaperQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "Failed to parse parameters", nil)
		return
	}
	o, err := ctrl.questionPaperSvr.SelectExamPaperById(q.ID)
	if err != nil {
		http.ErrorData(c, "Failed to parse parameters", nil)
		return
	}
	http.SuccessData(c, "Data retrieved successfully!", o)
}

// @Summary      获取所有考试试卷
// @Description  获取全部考试试卷列表（不分页）
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ExamPaperQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/paper/exam/all [post]
func (ctrl *ExamPaperController) SelectExamPaperAll(c *gin.Context) {
	oq := model.ExamPaperQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "Failed to retrieve data!", nil)
		return
	}
	list, err := ctrl.questionPaperSvr.SelectExamPaperAll(oq)
	if err != nil {
		http.ErrorData(c, "Failed to retrieve data!", nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

// @Summary      创建考试试卷
// @Description  创建新考试试卷（仅教师可操作）
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ExamPaper  true  "试卷信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/paper/exam/create [post]
func (ctrl *ExamPaperController) CreateExamPaper(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Unable to get current user info", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "User not logged in", nil)
		return
	}

	if user.ID == 0 {
		http.ErrorData(c, "Invalid user ID", nil)
		return
	}

	if !user.IsTeacher {
		http.ErrorData(c, "No permission to operate!", nil)
	}

	o := model.ExamPaper{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "Failed to parse parameters: "+err.Error(), nil)
		return
	}

	o.UserId = user.ID

	_, err = ctrl.questionPaperSvr.CreateExamPaper(o)
	if err != nil {
		http.ErrorData(c, "Failed to add", nil)
		return
	}
	http.SuccessData(c, "Data retrieved successfully!", nil)
}

// @Summary      编辑考试试卷
// @Description  修改考试试卷信息
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ExamPaper  true  "试卷信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/paper/exam/edit [post]
func (ctrl *ExamPaperController) EditExamPaper(c *gin.Context) {
	o := model.ExamPaper{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "Failed to parse parameters", nil)
		return
	}

	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Unable to get current user info", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "User not logged in", nil)
		return
	}

	if !user.IsAdmin && !user.IsTeacher {
		http.ErrorData(c, "No permission to access exam paper list", nil)
		return
	}

	// 获取原试卷数据
	existing, err := ctrl.questionPaperSvr.SelectExamPaperById(o.ID)
	if err != nil {
		http.ErrorData(c, "Exam paper does not exist", nil)
		return
	}

	// 如果不是管理员，只能修改自己的试卷
	if !user.IsAdmin && existing.UserId != user.ID {
		http.ErrorData(c, "No permission to edit this exam paper", nil)
		return
	}

	err = ctrl.questionPaperSvr.EditExamPaper(o)
	if err != nil {
		http.ErrorData(c, "Failed to edit", nil)
		return
	}
	http.SuccessData(c, "Edited successfully!", nil)
}

// @Summary      删除考试试卷
// @Description  删除指定考试试卷
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ExamPaper  true  "试卷ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/paper/exam/delete [post]
func (ctrl *ExamPaperController) DeleteExamPaper(c *gin.Context) {
	o := model.ExamPaper{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "Failed to parse parameters", nil)
		return
	}

	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Unable to get current user info", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "User not logged in", nil)
		return
	}

	if !user.IsAdmin && !user.IsTeacher {
		http.ErrorData(c, "No permission to access exam paper list", nil)
		return
	}

	// 获取原试卷数据
	existing, err := ctrl.questionPaperSvr.SelectExamPaperById(o.ID)
	if err != nil {
		http.ErrorData(c, "Exam paper does not exist", nil)
		return
	}

	// 如果不是管理员，只能删除自己的试卷
	if !user.IsAdmin && existing.UserId != user.ID {
		http.ErrorData(c, "No permission to delete this exam paper", nil)
		return
	}

	err = ctrl.questionPaperSvr.DeleteExamPaper(o.ID)
	if err != nil {
		http.ErrorData(c, "Failed to delete", nil)
		return
	}
	http.SuccessData(c, "Deleted successfully!", nil)
}

// @Summary      上传考试试卷文件
// @Description  上传考试试卷PDF文件
// @Tags         试卷管理
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "试卷文件"
// @Success      200   {object}  map[string]interface{}  "上传成功"
// @Security     BearerAuth
// @Router       /v1/paper/exam/upload [post]
func (ctrl *ExamPaperController) UploadExamPaper(c *gin.Context) {
	// file, err := c.FormFile("file")
	http.SuccessData(c, "Uploaded successfully!", nil)
}

// @Summary      更新考试试卷题目
// @Description  更新考试试卷中的题目列表
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ExamPaper  true  "试卷与题目信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/paper/exam/question/update [post]
func (ctrl *ExamPaperController) UpdateExamPaperQuestion(c *gin.Context) {
	q := model.ExamPaper{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "Failed to parse parameters!", nil)
		return
	}
	err := ctrl.questionPaperSvr.UpdateExamPaperQuestion(q)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", nil)
}
