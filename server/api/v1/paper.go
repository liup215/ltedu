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
// @Summary      获取试卷系列列表
// @Description  分页查询试卷系列列表
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperSeriesQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/pastPaper/series/list [post]
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

// @Summary      根据ID获取试卷系列
// @Description  根据系列ID获取试卷系列详情
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperSeriesQuery  true  "系列ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/pastPaper/series/getById [post]
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

// @Summary      获取所有试卷系列
// @Description  获取全部试卷系列列表（不分页）
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperSeriesQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/pastPaper/series/all [post]
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

// @Summary      创建试卷系列
// @Description  创建新试卷系列
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperSeries  true  "系列信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/pastPaper/series/create [post]
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

// @Summary      编辑试卷系列
// @Description  修改试卷系列信息
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperSeries  true  "系列信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/pastPaper/series/edit [post]
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

// @Summary      删除试卷系列
// @Description  删除指定试卷系列
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperSeries  true  "系列ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/pastPaper/series/delete [post]
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
// @Summary      获取试卷代码列表
// @Description  分页查询试卷代码列表
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperCodeQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/pastPaper/code/list [post]
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

// @Summary      根据ID获取试卷代码
// @Description  根据代码ID获取试卷代码详情
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperCodeQuery  true  "代码ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/pastPaper/code/getById [post]
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

// @Summary      获取所有试卷代码
// @Description  获取全部试卷代码列表（不分页）
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperCodeQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/pastPaper/code/all [post]
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

// @Summary      创建试卷代码
// @Description  创建新试卷代码
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperCode  true  "代码信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/pastPaper/code/create [post]
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

// @Summary      编辑试卷代码
// @Description  修改试卷代码信息
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperCode  true  "代码信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/pastPaper/code/edit [post]
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

// @Summary      删除试卷代码
// @Description  删除指定试卷代码
// @Tags         试卷管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.PaperCode  true  "代码ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/pastPaper/code/delete [post]
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
