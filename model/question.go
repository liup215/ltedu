package model

import "encoding/json"

const (
	QUESTION_TYPE_SINGLE_CHOICE   = 1 // 单选
	QUESTION_TYPE_MULTIPLE_CHOICE = 2 // 多选
	QUESTION_TYPE_TRUE_FALSE      = 3 // 判断
	QUESTION_TYPE_GAP_FILLING     = 4 // 填空
	QUESTION_TYPE_SHORT_ANSWER    = 5 // 简答
	QUESTION_TYPE_STRUCTURED      = 6 // 结构化问题
)

const (
	QUESTION_STATE_NORMAL    = 1 // 正常
	QUESTION_STATE_FORBIDDEN = 2 // 禁用
	QUESTION_STATE_DELETE    = 3 // 删除
)

type SingleChoiceQuestion struct {
	Options []ChoiceQuestionOption `json:"options"`
	Answer  string                 `json:"answer"`
}

type MultipleChoiceQuestion struct {
	Options []ChoiceQuestionOption `json:"options"`
	Answer  []string               `json:"answer"`
}

type ChoiceQuestionOption struct {
	Prefix  string `json:"prefix"`
	Content string `json:"content"`
}

type TrueOrFalseQuestion struct {
	Answer int `json:"answer"` // 1 正确 2 错误
}

type GapFillingQuestion struct {
	Answer []string `json:"answer"`
}

type ShortAnswerQuestion struct {
	Answer string `json:"answer"`
}

type StructuredQuestion struct {
	Answer string `json:"answer"`
}

type QuestionContent struct {
	Score            int                    `json:"score"`
	PartLabel        string                 `json:"partLabel"` // 部分标签
	SubpartLabel     string                 `json:"subpartLabel"`
	Answer           string                 `gorm:"type:text" json:"answer"`
	Analyze          string                 `gorm:"type:text" json:"analyze"`
	QuestionTypeId   uint                   `json:"questionTypeId" gorm:"-"`
	QuestionTypeName string                 `json:"questionTypeName" gorm:"-"`
	SingleChoice     SingleChoiceQuestion   `json:"singleChoice" gorm:"-"`
	MultipleChoice   MultipleChoiceQuestion `json:"multipleChoice" gorm:"-"`
	TrueOrFalse      TrueOrFalseQuestion    `json:"trueOrFalse" gorm:"-"`
	GapFilling       GapFillingQuestion     `json:"gapFilling" gorm:"-"`
	ShortAnswer      ShortAnswerQuestion    `json:"shortAnswer" gorm:"-"`
}

type Question struct {
	Model
	TotalScore             int               `json:"totalScore" gorm:"-"`
	Difficult              int               `json:"difficult"`
	SyllabusId             uint              `gorm:"syllabusId" json:"syllabusId"`
	Syllabus               Syllabus          `json:"syllabus"`
	Status                 int               `json:"status"`
	Stem                   string            `json:"stem" gorm:"type:longtext"` // 原问题，接受后存入QuestionString后储存
	QuestionContents       []QuestionContent `gorm:"-" json:"questionContents"`
	QuestionContentsString string            `json:"-" gorm:"type:longtext;"` // JSON中没有此信息，只用于存储原问题
	KnowledgePoints        []*KnowledgePoint `gorm:"many2many:question_keypoints" json:"knowledgePoints,omitempty"`
	PastPaperId            uint              `json:"pastPaperId"`
	PastPaper              PastPaper         `json:"pastPaper"`
	IndexInPastPaper       int               `json:"indexInPastPaper"`
}

func (q *Question) Format() error {
	qcs := []QuestionContent{}
	err := json.Unmarshal([]byte(q.QuestionContentsString), &qcs)
	if err != nil {
		return err
	}

	q.QuestionContents = qcs

	// 计算总分
	total := 0
	for _, qc := range qcs {
		total += qc.Score
	}
	q.TotalScore = total

	// 保证难度最小为1
	if q.Difficult == 0 {
		q.Difficult = 1
	}

	return nil
}

type QuestionQueryRequest struct {
	ID               uint   `json:"id"`
	Stem             string `json:"stem"`
	SyllabusId       uint   `json:"syllabusId"`
	Difficult        int    `json:"difficult"`
	Status           int    `json:"Status"`
	PastPaperId      uint   `json:"pastPaperId"`
	PaperName        string `json:"paperName"`
	ExamNodeId       uint   `json:"examNodeId"`
	KnowledgePointId uint   `json:"knowledgePointId"`
	Page
}
