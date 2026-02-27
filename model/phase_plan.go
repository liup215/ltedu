package model

import "time"

// LearningPhasePlan 阶段性计划 — 每个学习计划(StudentLearningPlan)包含若干阶段性计划
// 每个阶段性计划针对某个考试节点，有明确的起止时间，并覆盖多个章节。
type LearningPhasePlan struct {
	Model
	PlanId     uint               `json:"planId" gorm:"index"`
	Plan       *StudentLearningPlan `json:"plan,omitempty" gorm:"foreignKey:PlanId"`
	ExamNodeId uint               `json:"examNodeId" gorm:"index"`
	ExamNode   *SyllabusExamNode  `json:"examNode,omitempty" gorm:"foreignKey:ExamNodeId"`
	Title      string             `json:"title" gorm:"size:255"`
	StartDate  *time.Time         `json:"startDate"`
	EndDate    *time.Time         `json:"endDate"`
	SortOrder  int                `json:"sortOrder" gorm:"default:0"`
	Chapters   []*Chapter         `json:"chapters,omitempty" gorm:"many2many:phase_plan_chapters;"`

	// 以下字段为未来扩展保留（教学资源、学生任务、完成情况追踪）
	// Resources      []*TeachingResource  `json:"-" gorm:"-"`
	// Tasks          []*StudentTask       `json:"-" gorm:"-"`
	// CompletionRate float32              `json:"-" gorm:"-"`
}

// LearningPhasePlanQuery 查询条件
type LearningPhasePlanQuery struct {
	ID         uint `json:"id"`
	PlanId     uint `json:"planId"`
	ExamNodeId uint `json:"examNodeId"`
	Page
}

// LearningPhasePlanCreateRequest 创建请求
type LearningPhasePlanCreateRequest struct {
	PlanId     uint       `json:"planId" binding:"required"`
	ExamNodeId uint       `json:"examNodeId" binding:"required"`
	Title      string     `json:"title"`
	StartDate  *time.Time `json:"startDate"`
	EndDate    *time.Time `json:"endDate"`
	SortOrder  int        `json:"sortOrder"`
}

// LearningPhasePlanUpdateRequest 更新请求
type LearningPhasePlanUpdateRequest struct {
	ID         uint       `json:"id" binding:"required"`
	Title      string     `json:"title"`
	StartDate  *time.Time `json:"startDate"`
	EndDate    *time.Time `json:"endDate"`
	SortOrder  int        `json:"sortOrder"`
}

// LearningPhasePlanAddChapterRequest 为阶段性计划添加章节请求
type LearningPhasePlanAddChapterRequest struct {
	PhasePlanId uint `json:"phasePlanId" binding:"required"`
	ChapterId   uint `json:"chapterId" binding:"required"`
}

// LearningPhasePlanRemoveChapterRequest 移除阶段性计划的章节请求
type LearningPhasePlanRemoveChapterRequest struct {
	PhasePlanId uint `json:"phasePlanId" binding:"required"`
	ChapterId   uint `json:"chapterId" binding:"required"`
}
