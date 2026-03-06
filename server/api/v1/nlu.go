package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var NLUCtrl = &NLUController{
	nluSvr: service.NLUSvr,
}

// NLUController handles NLU (Natural Language Understanding) API endpoints.
type NLUController struct {
	nluSvr *service.NLUService
}

// AnalyzeQuery performs intent classification and entity extraction on a user query.
// @Summary      Analyze educational query intent and entities
// @Description  Performs NLU analysis on an educational query, returning detected intent, entities, and language
// @Tags         NLU
// @Accept       json
// @Produce      json
// @Param        body  body  model.NLURequest  true  "Query to analyze"
// @Success      200   {object}  map[string]interface{}  "NLU result"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Router       /v1/ai/nlu/analyze [post]
func (ctrl *NLUController) AnalyzeQuery(c *gin.Context) {
	var req model.NLURequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if req.Query == "" {
		http.ErrorData(c, "query不能为空", nil)
		return
	}

	result, err := ctrl.nluSvr.AnalyzeQuery(req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "分析成功", result)
}

// SubmitFeedback stores user-provided NLU correction feedback.
// @Summary      Submit NLU correction feedback
// @Description  Allows users to correct an NLU prediction to improve future accuracy
// @Tags         NLU
// @Accept       json
// @Produce      json
// @Param        body  body  model.NLUFeedbackRequest  true  "Feedback data"
// @Success      200   {object}  map[string]interface{}  "Feedback saved"
// @Failure      400   {object}  map[string]interface{}  "Bad request"
// @Router       /v1/ai/nlu/feedback [post]
func (ctrl *NLUController) SubmitFeedback(c *gin.Context) {
	var req model.NLUFeedbackRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if req.Query == "" || req.CorrectIntent == "" {
		http.ErrorData(c, "query和correctIntent不能为空", nil)
		return
	}

	user, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "未登录", nil)
		return
	}
	if err := ctrl.nluSvr.SaveFeedback(uint(user.ID), req); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "反馈已保存", nil)
}
