package model

const (
	LearningPlanTypeLong  = "long"  // 长期计划
	LearningPlanTypeMid   = "mid"   // 中期计划
	LearningPlanTypeShort = "short" // 短期计划
)

// StudentLearningPlan 学生个性化学习计划（三层体系）
type StudentLearningPlan struct {
	Model
	ClassId   uint    `json:"classId" gorm:"index"`
	Class     *Class  `json:"class,omitempty" gorm:"foreignKey:ClassId"`
	UserId    uint    `json:"userId" gorm:"index"`
	User      *User   `json:"user,omitempty" gorm:"foreignKey:UserId"`
	PlanType  string  `json:"planType" gorm:"size:10;index"` // "long", "mid", "short"
	Content   string  `json:"content" gorm:"type:text"`      // JSON / markdown 内容
	Version   int     `json:"version" gorm:"default:1"`
	CreatedBy uint    `json:"createdBy" gorm:"index"` // 创建人（通常是教师）
}

// StudentLearningPlanVersion 学习计划历史版本
type StudentLearningPlanVersion struct {
	Model
	PlanId    uint   `json:"planId" gorm:"index"`
	Version   int    `json:"version"`
	Content   string `json:"content" gorm:"type:text"`
	ChangedBy uint   `json:"changedBy" gorm:"index"` // 修改人
	Comment   string `json:"comment" gorm:"size:500"`
}

// StudentLearningPlanQuery 查询条件
type StudentLearningPlanQuery struct {
	ID       uint   `json:"id"`
	ClassId  uint   `json:"classId"`
	UserId   uint   `json:"userId"`
	PlanType string `json:"planType"`
	Page
}

// StudentLearningPlanCreateRequest 创建请求
type StudentLearningPlanCreateRequest struct {
	ClassId  uint   `json:"classId" binding:"required"`
	UserId   uint   `json:"userId" binding:"required"`
	PlanType string `json:"planType" binding:"required"`
	Content  string `json:"content"`
	Comment  string `json:"comment"` // 初始版本备注
}

// StudentLearningPlanUpdateRequest 更新请求
type StudentLearningPlanUpdateRequest struct {
	ID      uint   `json:"id" binding:"required"`
	Content string `json:"content"`
	Comment string `json:"comment"` // 版本备注
}

// StudentLearningPlanVersionQuery 版本历史查询
type StudentLearningPlanVersionQuery struct {
	PlanId uint `json:"planId" binding:"required"`
	Page
}

// StudentLearningPlanRollbackRequest 版本回滚请求
type StudentLearningPlanRollbackRequest struct {
	PlanId  uint   `json:"planId" binding:"required"`
	Version int    `json:"version" binding:"required"`
	Comment string `json:"comment"`
}

// GeneratePlansRequest 批量生成模板学习计划请求
type GeneratePlansRequest struct {
	ClassId      uint   `json:"classId" binding:"required"`
	SyllabusId   uint   `json:"syllabusId" binding:"required"`
	StartMonth   string `json:"startMonth" binding:"required"` // "YYYY-MM"
	EndMonth     string `json:"endMonth" binding:"required"`   // "YYYY-MM"
	PhaseRatios  []int  `json:"phaseRatios" binding:"required"` // e.g. [30,20,20,10]
	ExamNodeMode string `json:"examNodeMode"` // "sequential" | "parallel" (default: "sequential")
	Comment      string `json:"comment"`
}

// GeneratePlansResult 批量生成模板学习计划结果
type GeneratePlansResult struct {
	StudentCount int      `json:"studentCount"`
	Count        int      `json:"count"` // total plans created
	Errors       []string `json:"errors,omitempty"`
}
