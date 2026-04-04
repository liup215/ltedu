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
// @Summary      获取题目列表
// @Description  分页查询题目列表
// @Tags         题目管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.QuestionQueryRequest  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/question/list [post]
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

// @Summary      根据ID获取题目
// @Description  根据题目ID获取题目详情
// @Tags         题目管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperSeriesQuery  true  "题目ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/question/byId [post]
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

// @Summary      获取所有题目
// @Description  获取全部题目列表（不分页）
// @Tags         题目管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.QuestionQueryRequest  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/question/all [post]
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

// @Summary      创建题目
// @Description  创建新题目
// @Tags         题目管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Question  true  "题目信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/question/create [post]
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

	newID, err := ctrl.questionSvr.CreateQuestion(o, user.ID)
	if err != nil {
		http.ErrorData(c, "添加失败:"+err.Error(), nil)
		return
	}
	http.SuccessData(c, "添加成功!", gin.H{"id": newID})
}

// @Summary      关联知识点到题目
// @Description  为题目添加知识点关联
// @Tags         题目管理
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "题目ID和知识点ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/question/link-knowledge-point [post]
func (ctrl *QuestionController) LinkKnowledgePoint(c *gin.Context) {
	var req struct {
		QuestionId       uint `json:"questionId" binding:"required"`
		KnowledgePointId uint `json:"knowledgePointId" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.questionSvr.LinkKnowledgePoint(req.QuestionId, req.KnowledgePointId); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "知识点关联成功!", nil)
}

// @Summary      取消题目与知识点的关联
// @Description  删除题目与知识点的关联关系
// @Tags         题目管理
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "题目ID和知识点ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/question/unlink-knowledge-point [post]
func (ctrl *QuestionController) UnlinkKnowledgePoint(c *gin.Context) {
	var req struct {
		QuestionId       uint `json:"questionId" binding:"required"`
		KnowledgePointId uint `json:"knowledgePointId" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.questionSvr.UnlinkKnowledgePoint(req.QuestionId, req.KnowledgePointId); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "知识点取消关联成功!", nil)
}

// @Summary      编辑题目
// @Description  修改题目信息
// @Tags         题目管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Question  true  "题目信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/question/edit [post]
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

// @Summary      删除题目
// @Description  删除指定题目
// @Tags         题目管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.QuestionQueryRequest  true  "题目ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/question/delete [post]
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
