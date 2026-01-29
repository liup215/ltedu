package v1

import (
	"edu/lib/net/http"
	"edu/model"

	"github.com/gin-gonic/gin"
)

// RandomPaper 管理
func (ctrl *ExamPaperController) SelectRandomPaperList(c *gin.Context) {
	q := model.RandomPaperQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.questionPaperSvr.SelectRandomPaperList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *ExamPaperController) SelectRandomPaperById(c *gin.Context) {
	q := model.RandomPaperQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.questionPaperSvr.SelectRandomPaperById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *ExamPaperController) CreateRandomPaper(c *gin.Context) {
	o := model.RandomPaper{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.questionPaperSvr.CreateRandomPaper(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}
