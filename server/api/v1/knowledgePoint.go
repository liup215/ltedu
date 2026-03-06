package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var KnowledgePointCtrl = &KnowledgePointController{}

type KnowledgePointController struct{}

// GenerateKeypoints 为章节生成知识点
// @Summary      生成章节知识点
// @Description  为指定章节自动生成知识点
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "章节ID和模式"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/chapter/generate-keypoints [post]
func (ctrl *KnowledgePointController) GenerateKeypoints(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)
	_ = u

	var req struct {
		ChapterId uint   `json:"chapterId" binding:"required"`
		Mode      string `json:"mode"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	keypoints, err := service.KnowledgePointSvr.AutoGenerateFromChapter(req.ChapterId)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Knowledge points generated successfully!", gin.H{
		"keypoints": keypoints,
		"count":     len(keypoints),
	})
}

// AutoLinkQuestion 自动关联题目到知识点
// @Summary      自动关联题目到知识点
// @Description  将题目自动关联到匹配的知识点
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "题目ID、章节ID和考纲ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/question/auto-link-keypoints [post]
func (ctrl *KnowledgePointController) AutoLinkQuestion(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)
	_ = u

	var req struct {
		QuestionId uint `json:"questionId" binding:"required"`
		ChapterId  uint `json:"chapterId"`
		SyllabusId uint `json:"syllabusId"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	// 必须指定范围
	if req.ChapterId == 0 && req.SyllabusId == 0 {
		http.ErrorData(c, "chapterId or syllabusId is required", nil)
		return
	}

	linkedIds, err := service.KnowledgePointSvr.AutoLinkQuestionToKeypoints(
		req.QuestionId, req.ChapterId, req.SyllabusId)

	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Question linked to knowledge points!", gin.H{
		"linkedKeypoints": linkedIds,
		"count":           len(linkedIds),
	})
}

// AutoLinkQuestionIntelligent 智能关联题目到知识点（两阶段方法）
// @Summary      智能关联题目到知识点
// @Description  使用两阶段智能方法将题目关联到知识点
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "题目ID和考纲ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/question/auto-link-keypoints-intelligent [post]
func (ctrl *KnowledgePointController) AutoLinkQuestionIntelligent(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)
	_ = u

	var req struct {
		QuestionId uint `json:"questionId" binding:"required"`
		SyllabusId uint `json:"syllabusId" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	linkedIds, err := service.KnowledgePointSvr.AutoLinkQuestionToKeypointsIntelligent(
		req.QuestionId, req.SyllabusId)

	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Question intelligently linked to knowledge points!", gin.H{
		"linkedKeypoints": linkedIds,
		"count":           len(linkedIds),
	})
}

// AutoMigrateSyllabus 批量自动化迁移考纲
// @Summary      批量自动迁移考纲
// @Description  批量自动化迁移考纲知识点（仅管理员）
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "考纲ID和迁移选项"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/syllabus/auto-migrate-keypoints [post]
func (ctrl *KnowledgePointController) AutoMigrateSyllabus(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)

	// 仅管理员可执行
	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || !user.IsAdmin {
		http.ErrorData(c, "Only admin can perform migration", nil)
		return
	}

	var req struct {
		SyllabusId uint                   `json:"syllabusId" binding:"required"`
		Options    service.MigrateOptions `json:"options"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	report, err := service.KnowledgePointSvr.AutoMigrateSyllabus(req.SyllabusId, req.Options)
	if err != nil {
		http.ErrorData(c, err.Error(), report)
		return
	}

	http.SuccessData(c, "Migration completed!", report)
}

// Create 创建知识点
// @Summary      创建知识点
// @Description  创建新知识点
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  model.KnowledgePoint  true  "知识点信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-point/create [post]
func (ctrl *KnowledgePointController) Create(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)
	_ = u

	var kp model.KnowledgePoint
	if err := c.BindJSON(&kp); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	err := service.KnowledgePointSvr.Create(&kp)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Knowledge point created!", kp)
}

// Update 更新知识点
// @Summary      更新知识点
// @Description  更新现有知识点信息
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  model.KnowledgePoint  true  "知识点信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-point/edit [post]
func (ctrl *KnowledgePointController) Update(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)
	_ = u

	var kp model.KnowledgePoint
	if err := c.BindJSON(&kp); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	err := service.KnowledgePointSvr.Update(&kp)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Knowledge point updated!", kp)
}

// Delete 删除知识点
// @Summary      删除知识点
// @Description  删除指定知识点
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "知识点ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-point/delete [post]
func (ctrl *KnowledgePointController) Delete(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)
	_ = u

	var req struct {
		ID uint `json:"id" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	err := service.KnowledgePointSvr.Delete(req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Knowledge point deleted!", nil)
}

// GetByID 根据ID获取知识点
// @Summary      根据ID获取知识点
// @Description  根据知识点ID获取知识点详情
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "知识点ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-point/byId [post]
func (ctrl *KnowledgePointController) GetByID(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)
	_ = u

	var req struct {
		ID uint `json:"id" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	kp, err := service.KnowledgePointSvr.GetByID(req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Success", kp)
}

// GetByChapter 根据章节ID获取知识点列表
// @Summary      根据章节获取知识点
// @Description  根据章节ID获取该章节下的所有知识点
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "章节ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-point/byChapter [post]
func (ctrl *KnowledgePointController) GetByChapter(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)
	_ = u

	var req struct {
		ChapterId uint `json:"chapterId" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	kps, err := service.KnowledgePointSvr.GetByChapterId(req.ChapterId)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Success", gin.H{
		"list":      kps,
		"total":     len(kps),
		"page":      1,
		"pageSize":  len(kps),
	})
}

// GetBySyllabus 根据考纲ID获取知识点列表
// @Summary      根据考纲获取知识点
// @Description  根据考纲ID获取该考纲下的所有知识点
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "考纲ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-point/bySyllabus [post]
func (ctrl *KnowledgePointController) GetBySyllabus(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)
	_ = u

	var req struct {
		SyllabusId uint `json:"syllabusId" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	kps, err := service.KnowledgePointSvr.GetBySyllabusId(req.SyllabusId)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Success", gin.H{
		"list":      kps,
		"total":     len(kps),
		"page":      1,
		"pageSize":  len(kps),
	})
}

// List 获取知识点列表（带分页）
// @Summary      获取知识点列表
// @Description  分页查询知识点列表
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  model.KnowledgePointQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-point/list [post]
func (ctrl *KnowledgePointController) List(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)
	_ = u

	var query model.KnowledgePointQuery
	if err := c.BindJSON(&query); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	kps, total, err := service.KnowledgePointSvr.GetAll(&query)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Success", gin.H{
		"list":      kps,
		"total":     total,
		"page":      query.PageIndex,
		"pageSize":  query.PageSize,
	})
}

// LinkQuestion 手动关联题目到知识点
// @Summary      关联题目到知识点
// @Description  手动将题目关联到指定知识点
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "知识点ID和题目ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-point/link-question [post]
func (ctrl *KnowledgePointController) LinkQuestion(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)
	_ = u

	var req struct {
		KnowledgePointId uint `json:"knowledgePointId" binding:"required"`
		QuestionId       uint `json:"questionId" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	if err := service.KnowledgePointSvr.LinkQuestion(req.KnowledgePointId, req.QuestionId); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Question linked to knowledge point!", nil)
}

// UnlinkQuestion 取消题目与知识点的关联
// @Summary      取消题目与知识点的关联
// @Description  手动取消题目与知识点的关联关系
// @Tags         知识点
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "知识点ID和题目ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/knowledge-point/unlink-question [post]
func (ctrl *KnowledgePointController) UnlinkQuestion(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)
	_ = u

	var req struct {
		KnowledgePointId uint `json:"knowledgePointId" binding:"required"`
		QuestionId       uint `json:"questionId" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	if err := service.KnowledgePointSvr.UnlinkQuestion(req.KnowledgePointId, req.QuestionId); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Question unlinked from knowledge point!", nil)
}
