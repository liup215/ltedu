package v1

import (
	"edu/lib/net/http"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var PaperCtrl = &ExamPaperController{
	questionPaperSvr: service.QuestionPaperSvr,
	questionSvr:      service.QuestionSvr,
	userSvr:          service.UserSvr,
}

type ExamPaperController struct {
	questionPaperSvr *service.QuestionPaperService
	questionSvr      *service.QuestionService
	userSvr          *service.UserService
}

// PaperSeries管理
func (ctrl *ExamPaperController) SelectSeriesList(c *gin.Context) {
	q := model.PaperSeriesQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.questionPaperSvr.SelectSeriesList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *ExamPaperController) SelectSeriesById(c *gin.Context) {
	q := model.PaperSeriesQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.questionPaperSvr.SelectSeriesById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *ExamPaperController) SelectSeriesAll(c *gin.Context) {
	oq := model.PaperSeriesQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.questionPaperSvr.SelectSeriesAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

func (ctrl *ExamPaperController) CreateSeries(c *gin.Context) {
	o := model.PaperSeries{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败:"+err.Error(), nil)
		return
	}
	err := ctrl.questionPaperSvr.CreateSeries(o)
	if err != nil {
		http.ErrorData(c, "添加失败:"+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *ExamPaperController) EditSeries(c *gin.Context) {
	o := model.PaperSeries{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.questionPaperSvr.EditSeries(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *ExamPaperController) DeleteSeries(c *gin.Context) {
	o := model.PaperSeries{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.questionPaperSvr.DeleteSeries(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// PaperCode管理
func (ctrl *ExamPaperController) SelectCodeList(c *gin.Context) {
	q := model.PaperCodeQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.questionPaperSvr.SelectCodeList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *ExamPaperController) SelectCodeById(c *gin.Context) {
	q := model.PaperCodeQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.questionPaperSvr.SelectCodeById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *ExamPaperController) SelectCodeAll(c *gin.Context) {
	oq := model.PaperCodeQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.questionPaperSvr.SelectCodeAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

func (ctrl *ExamPaperController) CreateCode(c *gin.Context) {
	o := model.PaperCode{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.questionPaperSvr.CreateCode(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *ExamPaperController) EditCode(c *gin.Context) {
	o := model.PaperCode{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.questionPaperSvr.EditCode(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *ExamPaperController) DeleteCode(c *gin.Context) {
	o := model.PaperCode{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.questionPaperSvr.DeleteCode(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}
