package model

type ExamPaper struct {
	Model
	Name           string      `json:"name"`
	SyllabusId     uint        `json:"syllabusId"`
	Syllabus       Syllabus    `json:"syllabus"`
	Questions      []*Question `json:"questions" gorm:"-"`
	QuestionIds    []uint      `json:"questionIds" gorm:"-"`
	QuestionIdsStr string      `json:"-"`
	UserId         uint        `json:"userId" gorm:"index"`
	User           *User       `json:"user"`
}

type ExamPaperQuery struct {
	ID         uint `json:"id"`
	SyllabusId uint `json:"syllabusId"`
	UserId     uint `json:"userId"`
	Page
}
