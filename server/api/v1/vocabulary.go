package v1

import (
	"edu/lib/net/http"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

func init() {
	VocabularyCtrl = &VocabularyController{
		vocabularySvr: service.VocabularySvr,
	}
}

var VocabularyCtrl *VocabularyController

type VocabularyController struct {
	vocabularySvr *service.VocabularyService
}

// @Summary      根据ID获取词汇集
// @Description  根据词汇集ID获取词汇集详情
// @Tags         词汇
// @Accept       json
// @Produce      json
// @Param        body  body  model.VocabularySetQuery  true  "词汇集ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/vocabularySet/byId [post]
func (ctrl *VocabularyController) SelectVocabularySetById(c *gin.Context) {
	q := model.VocabularySetQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "数据解析失败:"+err.Error(), nil)
		return
	}

	if set, err := ctrl.vocabularySvr.SelectVocabularySetById(q.ID); err != nil {
		http.ErrorData(c, "数据获取失败:"+err.Error(), nil)
		return
	} else {

		http.SuccessData(c, "数据获取成功!", set)
		return
	}
}

// @Summary      获取词汇集列表
// @Description  分页查询词汇集列表
// @Tags         词汇
// @Accept       json
// @Produce      json
// @Param        body  body  model.VocabularySetQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/vocabularySet/list [post]
func (ctrl *VocabularyController) SelectVocabularySetList(c *gin.Context) {
	q := model.VocabularySetQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "数据解析失败:"+err.Error(), nil)
		return
	}

	if set, total, err := ctrl.vocabularySvr.SelectVocabularySetList(q); err != nil {
		http.ErrorData(c, "数据获取失败:"+err.Error(), nil)
		return
	} else {

		http.SuccessData(c, "数据获取成功!", gin.H{
			"list":  set,
			"total": total,
		})
		return
	}
}

// @Summary      创建词汇集
// @Description  创建新词汇集
// @Tags         词汇
// @Accept       json
// @Produce      json
// @Param        body  body  model.VocabularySetCreateEditRequest  true  "词汇集信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/vocabularySet/create [post]
func (ctrl *VocabularyController) CreateVocabularySet(c *gin.Context) {
	vs := model.VocabularySetCreateEditRequest{}

	if err := c.BindJSON(&vs); err != nil {
		http.ErrorData(c, "数据解析失败："+err.Error(), nil)
		return
	}

	if err := ctrl.vocabularySvr.CreateVocabularySet(vs); err != nil {
		http.ErrorData(c, "数据解析失败："+err.Error(), nil)
		return
	} else {
		http.SuccessData(c, "创建成功！", nil)
	}

}

// @Summary      编辑词汇集
// @Description  修改词汇集信息
// @Tags         词汇
// @Accept       json
// @Produce      json
// @Param        body  body  model.VocabularySetCreateEditRequest  true  "词汇集信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/vocabularySet/edit [post]
func (ctrl *VocabularyController) EditVocabularySet(c *gin.Context) {
	vs := model.VocabularySetCreateEditRequest{}

	if err := c.BindJSON(&vs); err != nil {
		http.ErrorData(c, "数据解析失败："+err.Error(), nil)
		return
	}

	if err := ctrl.vocabularySvr.EditVocabularySet(vs); err != nil {
		http.ErrorData(c, "数据解析失败："+err.Error(), nil)
		return
	} else {
		http.SuccessData(c, "编辑成功！", nil)
	}

}

// @Summary      删除词汇集
// @Description  删除指定词汇集
// @Tags         词汇
// @Accept       json
// @Produce      json
// @Param        body  body  model.VocabularySetQuery  true  "词汇集ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/vocabularySet/delete [post]
func (ctrl *VocabularyController) DeleteVocabularySet(c *gin.Context) {
	q := model.VocabularySetQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "数据解析失败："+err.Error(), nil)
		return
	}

	if err := ctrl.vocabularySvr.DeleteVocabularySet(q.ID); err != nil {
		http.ErrorData(c, "数据解析失败："+err.Error(), nil)
		return
	} else {
		http.SuccessData(c, "删除成功！", nil)
	}

}

// @Summary      添加词汇条目
// @Description  向词汇集添加新词汇条目
// @Tags         词汇
// @Accept       json
// @Produce      json
// @Param        body  body  model.VocabularyItemCreateEditRequest  true  "词汇条目信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/vocabularyItem/insert [post]
func (ctrl *VocabularyController) InsertVocabularyItem(c *gin.Context) {
	vi := model.VocabularyItemCreateEditRequest{}

	if err := c.BindJSON(&vi); err != nil {
		http.ErrorData(c, "数据解析失败："+err.Error(), nil)
		return
	}

	if err := ctrl.vocabularySvr.InsertVocabularyItem(vi); err != nil {
		http.ErrorData(c, "添加失败："+err.Error(), nil)
		return
	} else {
		http.SuccessData(c, "添加成功！", nil)
	}

}

// @Summary      更新词汇条目
// @Description  更新词汇集中的词汇条目
// @Tags         词汇
// @Accept       json
// @Produce      json
// @Param        body  body  model.VocabularyItemCreateEditRequest  true  "词汇条目信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/vocabularyItem/update [post]
func (ctrl *VocabularyController) UpdateVocabularyItem(c *gin.Context) {
	vi := model.VocabularyItemCreateEditRequest{}

	if err := c.BindJSON(&vi); err != nil {
		http.ErrorData(c, "数据解析失败："+err.Error(), nil)
		return
	}

	if err := ctrl.vocabularySvr.UpdateVocabularyItem(vi); err != nil {
		http.ErrorData(c, "更新失败："+err.Error(), nil)
		return
	} else {
		http.SuccessData(c, "更新成功！", nil)
	}

}

// @Summary      删除词汇条目
// @Description  删除指定词汇条目
// @Tags         词汇
// @Accept       json
// @Produce      json
// @Param        body  body  model.VocabularyItemQuery  true  "词汇条目ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/vocabularyItem/delete [post]
func (ctrl *VocabularyController) DeleteVocabularyItem(c *gin.Context) {
	q := model.VocabularyItemQuery{}

	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "数据解析失败："+err.Error(), nil)
		return
	}

	if err := ctrl.vocabularySvr.DeleteVocabularyItem(q.ID); err != nil {
		http.ErrorData(c, "删除失败："+err.Error(), nil)
		return
	} else {
		http.SuccessData(c, "删除成功！", nil)
	}

}

// func (ctrl *VocabularyController) SelectLearningVocabularySet(c *gin.Context) {
// 	cu, ok := auth.GetCurrentUser(c)
// 	if !ok {
// 		http.ErrorData(c, "用户数据获取失败", nil)
// 		return
// 	}

// 	if list, err := ctrl.vocabularySvr.SelectLearningVocabularySet(uint(cu.Id)); err != nil {
// 		http.ErrorData(c, "数据获取失败："+err.Error(), nil)
// 		return
// 	} else {
// 		http.SuccessData(c, "数据获取成功！", list)
// 	}
// }

type SummaryQuery struct {
	VocabularySetId uint `json:"vocabularySetId"`
}

// func (ctrl *VocabularyController) GetLearningSummaryForSet(c *gin.Context) {
// 	q := SummaryQuery{}
// 	if err := c.BindJSON(&q); err != nil {
// 		http.ErrorData(c, "数据解析失败:"+err.Error(), nil)
// 		return
// 	}

// 	cu, ok := auth.GetCurrentUser(c)
// 	if !ok {
// 		http.ErrorData(c, "用户数据获取失败！", nil)
// 		return
// 	}

// 	s, err := ctrl.vocabularySvr.GetLearningSummaryForSet(uint(cu.Id), q.VocabularySetId)
// 	if err != nil {
// 		http.ErrorData(c, "数据获取失败:"+err.Error(), nil)
// 		return
// 	}

// 	http.SuccessData(c, "数据获取成功!", s)
// }

// func (ctrl *VocabularyController) LearnVocabularyItem(c *gin.Context) {
// 	q := model.VocabularyItem{}
// 	if err := c.BindJSON(&q); err != nil {
// 		http.ErrorData(c, "数据解析失败:"+err.Error(), nil)
// 		return
// 	}

// 	cu, ok := auth.GetCurrentUser(c)
// 	if !ok {
// 		http.ErrorData(c, "用户数据获取失败！", nil)
// 		return
// 	}

// 	err := ctrl.vocabularySvr.LearnVocabularyItem(uint(cu.Id), q.ID)
// 	if err != nil {
// 		http.ErrorData(c, "学习失败:"+err.Error(), nil)
// 		return
// 	}

// 	http.SuccessData(c, "数据获取成功!", nil)

// }
