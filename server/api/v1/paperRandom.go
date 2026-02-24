package v1

import (
	"edu/lib/net/http"
	"edu/model"

	"github.com/gin-gonic/gin"
)

// RandomPaper 管理
// @Summary      获取随机试卷列表
// @Description  分页查询随机试卷列表
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.RandomPaperQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/paper/random/list [post]
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

// @Summary      根据ID获取随机试卷
// @Description  根据试卷ID获取随机试卷详情
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.RandomPaperQuery  true  "试卷ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/paper/random/getById [post]
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

// @Summary      创建随机试卷
// @Description  创建新随机试卷
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.RandomPaper  true  "试卷信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/paper/random/create [post]
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
