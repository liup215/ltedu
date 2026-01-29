package model

type PastPaper struct {
	Model
	Name           string      `json:"name"`
	SyllabusId     uint        `json:"syllabusId"`
	Syllabus       Syllabus    `json:"syllabus"`
	Year           int         `json:"year"`
	PaperCodeId    uint        `json:"paperCodeId" gorm:"index"`
	PaperCode      PaperCode   `json:"paperCode"`
	PaperSeriesId  uint        `json:"paperSeriesId" gorm:"index"`
	PaperSeries    PaperSeries `json:"paperSeries"`
	QuestionNumber int         `json:"questionNumber"`
	Questions      []*Question `json:"questions" gorm:"-"`
}

type PastPaperQuery struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	SyllabusId    uint   `json:"syllabusId"`
	Year          int    `json:"year"`
	PaperCodeId   uint   `json:"paperCodeId"`
	PaperSeriesId uint   `json:"paperSeriesId"`
	Page
}
