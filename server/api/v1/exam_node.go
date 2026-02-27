package v1

import (
	"edu/lib/net/http"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var ExamNodeCtrl = &ExamNodeController{
	svr: service.ExamNodeSvr,
}

type ExamNodeController struct {
	svr *service.ExamNodeService
}

// CreateExamNode 创建考试节点
func (ctrl *ExamNodeController) CreateExamNode(c *gin.Context) {
	req := model.SyllabusExamNodeCreateRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	node, err := ctrl.svr.CreateExamNode(req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "考试节点创建成功!", node)
}

// UpdateExamNode 更新考试节点基本信息
func (ctrl *ExamNodeController) UpdateExamNode(c *gin.Context) {
	req := model.SyllabusExamNodeUpdateRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	node, err := ctrl.svr.UpdateExamNode(req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "考试节点更新成功!", node)
}

// DeleteExamNode 删除考试节点
func (ctrl *ExamNodeController) DeleteExamNode(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.svr.DeleteExamNode(req.ID); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "考试节点删除成功!", nil)
}

// GetExamNodeById 根据ID获取考试节点
func (ctrl *ExamNodeController) GetExamNodeById(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	node, err := ctrl.svr.GetExamNode(req.ID)
	if err != nil {
		http.ErrorData(c, "数据获取失败: "+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", node)
}

// ListExamNodes 获取某Syllabus下的所有考试节点
func (ctrl *ExamNodeController) ListExamNodes(c *gin.Context) {
	var req struct {
		SyllabusId uint `json:"syllabusId" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	nodes, err := ctrl.svr.ListExamNodes(req.SyllabusId)
	if err != nil {
		http.ErrorData(c, "数据获取失败: "+err.Error(), nil)
		return
	}
	if nodes == nil {
		nodes = []*model.SyllabusExamNode{}
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  nodes,
		"total": len(nodes),
	})
}

// AddChapter 为考试节点添加章节（自动递归包含所有子章节）
func (ctrl *ExamNodeController) AddChapter(c *gin.Context) {
	req := model.SyllabusExamNodeAddChapterRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.svr.AddChapterToExamNode(req); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "章节添加成功!", nil)
}

// RemoveChapter 从考试节点移除章节
func (ctrl *ExamNodeController) RemoveChapter(c *gin.Context) {
	req := model.SyllabusExamNodeRemoveChapterRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.svr.RemoveChapterFromExamNode(req); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "章节移除成功!", nil)
}

// AddPaperCode 为考试节点添加试卷代码
func (ctrl *ExamNodeController) AddPaperCode(c *gin.Context) {
	req := model.SyllabusExamNodeAddPaperCodeRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.svr.AddPaperCodeToExamNode(req); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "试卷代码添加成功!", nil)
}

// RemovePaperCode 从考试节点移除试卷代码
func (ctrl *ExamNodeController) RemovePaperCode(c *gin.Context) {
	req := model.SyllabusExamNodeRemovePaperCodeRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	if err := ctrl.svr.RemovePaperCodeFromExamNode(req); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "试卷代码移除成功!", nil)
}
