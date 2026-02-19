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

// AutoMigrateSyllabus 批量自动化迁移考纲
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
		"keypoints": kps,
		"count":     len(kps),
	})
}

// GetBySyllabus 根据考纲ID获取知识点列表
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
		"keypoints": kps,
		"count":     len(kps),
	})
}

// List 获取知识点列表（带分页）
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
		"keypoints": kps,
		"total":     total,
	})
}
