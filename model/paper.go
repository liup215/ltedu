package model

type PaperCode struct {
	Model
	Name       string   `json:"name"`
	SyllabusId uint     `json:"syllabusId"`
	Syllabus   Syllabus `json:"syllabus"`
}

type PaperCodeQuery struct {
	ID         uint   `json:"id"`
	SyllabusId uint   `json:"syllabusId"`
	Name       string `json:"name"`
	Page
}

type PaperSeries struct {
	Model
	Name       string   `json:"name"`
	SyllabusId uint     `json:"syllabusId"`
	Syllabus   Syllabus `json:"syllabus"`
}

type PaperSeriesQuery struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	SyllabusId uint   `json:"syllabusId"`
	Page
}

const (
	IS_PASTPAPER_YES = 1
	IS_PASTPAPER_NO  = 2
)
