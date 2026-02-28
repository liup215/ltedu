package model

// SyllabusExamNode 考试节点 — 每个Syllabus可绑定多个考试节点
// 每个节点包含若干章节（设置章节时自动批量加入所有子章节）和试卷代码
type SyllabusExamNode struct {
	Model
	SyllabusId  uint         `json:"syllabusId" gorm:"index"`
	Syllabus    *Syllabus    `json:"syllabus,omitempty"`
	Name        string       `json:"name"`
	Description string       `json:"description" gorm:"type:text"`
	SortOrder   int          `json:"sortOrder" gorm:"default:0"`
	Chapters    []*Chapter   `json:"chapters,omitempty" gorm:"many2many:exam_node_chapters;"`
	PaperCodes  []*PaperCode `json:"paperCodes,omitempty" gorm:"foreignKey:ExamNodeId"`
}

// SyllabusExamNodeQuery 考试节点查询条件
type SyllabusExamNodeQuery struct {
	ID         uint `json:"id"`
	SyllabusId uint `json:"syllabusId"`
	Page
}

// SyllabusExamNodeCreateRequest 创建考试节点请求
type SyllabusExamNodeCreateRequest struct {
	SyllabusId  uint   `json:"syllabusId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	SortOrder   int    `json:"sortOrder"`
}

// SyllabusExamNodeUpdateRequest 更新考试节点请求
type SyllabusExamNodeUpdateRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SortOrder   int    `json:"sortOrder"`
}

// SyllabusExamNodeAddChapterRequest 为考试节点添加章节请求（自动递归添加所有子章节）
type SyllabusExamNodeAddChapterRequest struct {
	ExamNodeId uint `json:"examNodeId" binding:"required"`
	ChapterId  uint `json:"chapterId" binding:"required"`
}

// SyllabusExamNodeRemoveChapterRequest 移除考试节点的章节请求
type SyllabusExamNodeRemoveChapterRequest struct {
	ExamNodeId uint `json:"examNodeId" binding:"required"`
	ChapterId  uint `json:"chapterId" binding:"required"`
}

// SyllabusExamNodeAddPaperCodeRequest 为考试节点添加试卷代码请求
type SyllabusExamNodeAddPaperCodeRequest struct {
	ExamNodeId  uint `json:"examNodeId" binding:"required"`
	PaperCodeId uint `json:"paperCodeId" binding:"required"`
}

// SyllabusExamNodeRemovePaperCodeRequest 移除考试节点的试卷代码请求
type SyllabusExamNodeRemovePaperCodeRequest struct {
	ExamNodeId  uint `json:"examNodeId" binding:"required"`
	PaperCodeId uint `json:"paperCodeId" binding:"required"`
}
