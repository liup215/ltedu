package v1

import (
	"edu/lib/net/http"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var DocumentCategoryCtrl = &DocumentCategoryController{
	categorySvr: service.DocumentCategorySvr,
}

type DocumentCategoryController struct {
	categorySvr *service.DocumentCategoryService
}

// @Summary      创建文档分类
// @Description  创建新文档分类
// @Tags         文档
// @Accept       json
// @Produce      json
// @Param        body  body  model.DocumentCategoryCreateEditRequest  true  "分类信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/documentCategory/create [post]
func (ctrl *DocumentCategoryController) CreateCategory(c *gin.Context) {
	dc := model.DocumentCategoryCreateEditRequest{}
	if err := c.BindJSON(&dc); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.categorySvr.CreateCategory(dc); err != nil {
		http.ErrorData(c, "创建失败", nil)
		return
	}

	http.SuccessData(c, "创建成功!", nil)
}

// @Summary      编辑文档分类
// @Description  修改文档分类信息
// @Tags         文档
// @Accept       json
// @Produce      json
// @Param        body  body  model.DocumentCategoryCreateEditRequest  true  "分类信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/documentCategory/edit [post]
func (ctrl *DocumentCategoryController) EditCategory(c *gin.Context) {
	dc := model.DocumentCategoryCreateEditRequest{}
	if err := c.BindJSON(&dc); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.categorySvr.EditCategory(dc); err != nil {
		http.ErrorData(c, "编辑失败", nil)
		return
	}

	http.SuccessData(c, "编辑成功!", nil)
}

// @Summary      删除文档分类
// @Description  删除指定文档分类
// @Tags         文档
// @Accept       json
// @Produce      json
// @Param        body  body  model.DocumentCategoryQueryRequest  true  "分类ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/documentCategory/delete [post]
func (ctrl *DocumentCategoryController) DeleteCategory(c *gin.Context) {
	q := model.DocumentCategoryQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.categorySvr.DeleteCategory(q.ID); err != nil {
		http.ErrorData(c, "删除失败", nil)
		return
	}

	http.SuccessData(c, "删除成功!", nil)
}

// @Summary      根据ID获取文档分类
// @Description  根据分类ID获取文档分类详情
// @Tags         文档
// @Accept       json
// @Produce      json
// @Param        body  body  model.DocumentCategoryQueryRequest  true  "分类ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/documentCategory/byId [post]
func (ctrl *DocumentCategoryController) SelectCategoryById(c *gin.Context) {
	q := model.DocumentCategoryQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	o, err := ctrl.categorySvr.SelectCategoryById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", o)
}

// @Summary      获取文档分类列表
// @Description  分页查询文档分类列表
// @Tags         文档
// @Accept       json
// @Produce      json
// @Param        body  body  model.DocumentCategoryQueryRequest  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/documentCategory/list [post]
func (ctrl *DocumentCategoryController) SelectCategoryList(c *gin.Context) {
	q := model.DocumentCategoryQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	list, total, err := ctrl.categorySvr.SelectCategoryList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

// @Summary      获取所有文档分类
// @Description  获取全部文档分类列表（不分页）
// @Tags         文档
// @Accept       json
// @Produce      json
// @Param        body  body  model.DocumentCategoryQueryRequest  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/documentCategory/all [post]
func (ctrl *DocumentCategoryController) SelectCategoryAll(c *gin.Context) {
	oq := model.DocumentCategoryQueryRequest{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.categorySvr.SelectCategoryAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}
