package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

// PastPaper 管理
// @Summary      获取历年试卷列表
// @Description  分页查询历年试卷列表
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PastPaperQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/paper/past/list [post]
func (ctrl *ExamPaperController) SelectPastPaperList(c *gin.Context) {
	q := model.PastPaperQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.questionPaperSvr.SelectPastPaperList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

// @Summary      根据ID获取历年试卷
// @Description  根据试卷ID获取历年试卷详情
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PastPaperQuery  true  "试卷ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/paper/past/getById [post]
func (ctrl *ExamPaperController) SelectPastPaperById(c *gin.Context) {
	q := model.PastPaperQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.questionPaperSvr.SelectPastPaperById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

// @Summary      获取所有历年试卷
// @Description  获取全部历年试卷列表（不分页）
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PastPaperQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/paper/past/all [post]
func (ctrl *ExamPaperController) SelectPastPaperAll(c *gin.Context) {
	oq := model.PastPaperQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.questionPaperSvr.SelectPastPaperAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

// @Summary      创建历年试卷
// @Description  创建新历年试卷
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PastPaper  true  "试卷信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/paper/past/create [post]
func (ctrl *ExamPaperController) CreatePastPaper(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	_, err = service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "用户未登录！", nil)
		return
	}

	o := model.PastPaper{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败: "+err.Error(), nil)
		return
	}

	_, err = ctrl.questionPaperSvr.CreatePastPaper(o)
	if err != nil {
		http.ErrorData(c, "试题创建失败："+err.Error(), nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", nil)
}

// @Summary      编辑历年试卷
// @Description  修改历年试卷信息
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PastPaper  true  "试卷信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/paper/past/edit [post]
func (ctrl *ExamPaperController) EditPastPaper(c *gin.Context) {
	o := model.PastPaper{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.questionPaperSvr.EditPastPaper(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// @Summary      删除历年试卷
// @Description  删除指定历年试卷
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PastPaper  true  "试卷ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/paper/past/delete [post]
func (ctrl *ExamPaperController) DeletePastPaper(c *gin.Context) {
	o := model.PastPaper{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.questionPaperSvr.DeletePastPaper(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// @Summary      上传历年试卷文件
// @Description  上传历年试卷PDF文件
// @Tags         试卷管理
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "试卷文件"
// @Success      200   {object}  map[string]interface{}  "上传成功"
// @Security     BearerAuth
// @Router       /v1/paper/past/upload [post]
func (ctrl *ExamPaperController) UploadPastPaper(c *gin.Context) {
	// file, err := c.FormFile("file")
	http.SuccessData(c, "上传成功!", nil)
}

// func (ctrl *ExamPaperController) InitiateQuestion(c *gin.Context) {
// 	o := model.PastPaperQuery{}
// 	if err := c.BindJSON(&o); err != nil {
// 		http.ErrorData(c, "参数解析失败!", nil)
// 		return
// 	}

// 	p, err := ctrl.questionPaperSvr.SelectPastPaperById(o.ID)
// 	if err != nil {
// 		http.ErrorData(c, "查无此试卷！", nil)
// 		return
// 	}

// 	err = ctrl.questionSvr.InitiateQuestion(p.Id, p.QuestionNumber)
// 	if err != nil {
// 		http.ErrorData(c, err.Error(), nil)
// 		return
// 	}
// 	http.SuccessData(c, "数据获取成功!", nil)
// }

// func (ctrl *ExamPaperController) UpdatePastPaperQuestion(c *gin.Context) {
// 	q := model.PaperQuestionUpdateRequest{}
// 	if err := c.BindJSON(&q); err != nil {
// 		http.ErrorData(c, "参数解析失败!", nil)
// 		return
// 	}
// 	err := ctrl.questionPaperSvr.UpdatePastPaperQuestion(q)
// 	if err != nil {
// 		http.ErrorData(c, err.Error(), nil)
// 		return
// 	}

// 	http.SuccessData(c, "数据获取成功!", nil)
// }

// func (ctrl *ExamPaperController) AddPastPaperQuestion(c *gin.Context) {
// 	q := model.PaperQuestionUpdateRequest{}
// 	if err := c.BindJSON(&q); err != nil {
// 		http.ErrorData(c, "参数解析失败!", nil)
// 		return
// 	}
// 	err := ctrl.questionPaperSvr.AddPastPaperQuestion(q)
// 	if err != nil {
// 		http.ErrorData(c, err.Error(), nil)
// 		return
// 	}

// 	http.SuccessData(c, "数据获取成功!", nil)
// }

// func (ctrl *ExamPaperController) DeletePastPaperQuestion(c *gin.Context) {
// 	q := model.PaperQuestionUpdateRequest{}
// 	if err := c.BindJSON(&q); err != nil {
// 		http.ErrorData(c, "参数解析失败!", nil)
// 		return
// 	}
// 	err := ctrl.questionPaperSvr.DeletePastPaperQuestion(q)
// 	if err != nil {
// 		http.ErrorData(c, err.Error(), nil)
// 		return
// 	}

// 	http.SuccessData(c, "数据获取成功!", nil)
// }
