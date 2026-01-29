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
