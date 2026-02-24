// ltedu-api/server/api/v1/practice.go

package v1

import (
	"edu/lib/net/http"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var PracticeCtrl = &PracticeController{
	practiceSvr: service.PracticeSvr,
}

type PracticeController struct {
	practiceSvr *service.PracticeService
}

// @Summary      快速练习
// @Description  根据条件生成一组快速练习题目
// @Tags         练习
// @Accept       json
// @Produce      json
// @Param        body  body  model.PracticeQuickRequest  true  "练习条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/practice/quick [post]
// POST /practice/quick
func (ctrl *PracticeController) QuickPractice(c *gin.Context) {
	var req model.PracticeQuickRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	questionIDs, err := ctrl.practiceSvr.GenerateQuickPractice(req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	resp := model.PracticeQuickResponse{
		List:  questionIDs,
		Total: len(questionIDs),
	}
	http.SuccessData(c, "数据获取成功!", resp)
}

// @Summary      试卷练习
// @Description  根据试卷生成练习题目列表
// @Tags         练习
// @Accept       json
// @Produce      json
// @Param        body  body  model.PracticePaperRequest  true  "练习条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/practice/paper [post]
// POST /practice/paper
func (ctrl *PracticeController) PaperPractice(c *gin.Context) {
	var req model.PracticePaperRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	questionIDs, err := ctrl.practiceSvr.GeneratePaperPractice(req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	resp := model.PracticePaperResponse{
		List: questionIDs,
	}
	http.SuccessData(c, "数据获取成功!", resp)
}

// @Summary      批改练习
// @Description  批改练习提交，返回得分和答题详情
// @Tags         练习
// @Accept       json
// @Produce      json
// @Param        body  body  model.PracticeGradeRequest  true  "答题数据"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/practice/grade [post]
// POST /practice/grade
func (ctrl *PracticeController) GradePractice(c *gin.Context) {
	var req model.PracticeGradeRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	result, err := ctrl.practiceSvr.GradePracticeSubmission(req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	resp := model.PracticeGradeResponse{
		Score:   result.Score,
		Total:   result.Total,
		Results: result.Results,
	}
	http.SuccessData(c, "批改成功!", resp)
}
