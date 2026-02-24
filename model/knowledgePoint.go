package model

type KnowledgePoint struct {
	Model
	ChapterId        uint        `json:"chapterId" gorm:"index"`
	Chapter          Chapter     `json:"chapter" gorm:"foreignKey:ChapterId"`
	Name             string      `json:"name"`
	Description      string      `json:"description" gorm:"type:text"`
	Difficulty       string      `json:"difficulty"` // basic/medium/hard
	EstimatedMinutes int         `json:"estimatedMinutes"`
	OrderIndex       int         `json:"orderIndex"`
	Questions        []*Question `gorm:"many2many:question_keypoints" json:"questions,omitempty"`
}

type KnowledgePointQuery struct {
	Model
	ChapterId  uint   `json:"chapterId"`
	SyllabusId uint   `json:"syllabusId"`
	Name       string `json:"name"`
	Difficulty string `json:"difficulty"`
	Page
}

// AI生成知识点的响应结构
type AIKnowledgePointData struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Difficulty       string `json:"difficulty"`
	EstimatedMinutes int    `json:"estimatedMinutes"`
}

// AI关联题目的响应结构
type AILinkQuestionResponse struct {
	Indices []int `json:"indices"`
}
