package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var QuestionCtrl = &QuestionController{
	questionSvr: service.QuestionSvr,
}

type QuestionController struct {
	questionSvr *service.QuestionService
}

// question管理
func (ctrl *QuestionController) SelectQuestionList(c *gin.Context) {
	q := model.QuestionQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.questionSvr.SelectQuestionList(q)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *QuestionController) SelectQuestionById(c *gin.Context) {
	q := model.PaperSeriesQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.questionSvr.SelectQuestionById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *QuestionController) SelectQuestionAll(c *gin.Context) {
	oq := model.QuestionQueryRequest{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.questionSvr.SelectQuestionAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

func (ctrl *QuestionController) CreateQuestion(c *gin.Context) {
	o := model.Question{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败:"+err.Error(), nil)
		return
	}

	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "用户获取失败", nil)
		return
	}

	_, err = ctrl.questionSvr.CreateQuestion(o, user.ID)
	if err != nil {
		http.ErrorData(c, "添加失败:"+err.Error(), nil)
		return
	}
	http.SuccessData(c, "添加成功!", nil)
}

func (ctrl *QuestionController) EditQuestion(c *gin.Context) {
	o := model.Question{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.questionSvr.EditQuestion(o)
	if err != nil {
		http.ErrorData(c, "编辑失败: "+err.Error(), nil)
		return
	}
	http.SuccessData(c, "编辑成功!", nil)
}

func (ctrl *QuestionController) DeleteQuestion(c *gin.Context) {
	o := model.QuestionQueryRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.questionSvr.DeleteQuestion(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *QuestionController) AddQuestionChapter(c *gin.Context) {
	o := model.QuestionChapterUpdateRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.questionSvr.AddQuestionChapter(o)
	if err != nil {
		http.ErrorData(c, "添加失败："+err.Error(), nil)
		return
	}
	http.SuccessData(c, "添加成功!", nil)
}

func (ctrl *QuestionController) DeleteQuestionChapter(c *gin.Context) {
	o := model.QuestionChapterUpdateRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.questionSvr.DeleteQuestionChapter(o)
	if err != nil {
		http.ErrorData(c, "删除失败: "+err.Error(), nil)
		return
	}
	http.SuccessData(c, "删除成功!", nil)
}
