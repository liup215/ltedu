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
