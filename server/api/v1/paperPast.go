package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

// PastPaper 管理
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
