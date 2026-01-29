package model

import (
	"time"
)

type RandomPaper struct {
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
	// InsertedQuestionNumber int         `json:"insertedQuestionNumber"`
	QuestionsInfo []*QuestionRandomPapers `json:"questionsInfo" gorm:"-"`
}

func (p *RandomPaper) GetRandomPaperResponse() *RandomPaperQueryResponse {
	r := &RandomPaperQueryResponse{
		Id:                     p.ID,
		CreatedAt:              p.CreatedAt,
		UpdatedAt:              p.UpdatedAt,
		Name:                   p.Name,
		OrganisationId:         p.Syllabus.Qualification.OrganisationId,
		OrganisationName:       p.Syllabus.Qualification.Organisation.Name,
		QualificationId:        p.Syllabus.QualificationId,
		QualificationName:      p.Syllabus.Qualification.Name,
		SyllabusId:             p.SyllabusId,
		SyllabusName:           p.Syllabus.Name,
		SyllabusCode:           p.Syllabus.Code,
		QuestionNumber:         p.QuestionNumber,
		InsertedQuestionNumber: len(p.QuestionsInfo),
		Questions:              p.QuestionsInfo,
	}

	return r
}

type QuestionRandomPapers struct {
	Model
	QuestionId               uint                     `json:"qustionId"`
	Question                 Question                 `json:"-"`
	QuestionQueryResponse    Question                 `json:"question" gorm:"-"`
	RandomPaperId            uint                     `json:"RandomPaperId"`
	RandomPaper              RandomPaper              `json:"-"`
	RandomPaperQueryResponse RandomPaperQueryResponse `json:"RandomPaper" gorm:"-"`
	Index                    int                      `json:"index"`
}

type RandomPaperQuery struct {
	ID         uint `json:"id"`
	SyllabusId uint `json:"syllabusId"`
	Page
}

type RandomPaperQueryResponse struct {
	Id                     uint                    `json:"id"`
	CreatedAt              time.Time               `json:"createdAt"`
	UpdatedAt              time.Time               `json:"updatedAt"`
	Name                   string                  `json:"name"`
	OrganisationId         uint                    `json:"organisationId"`
	OrganisationName       string                  `json:"organisationName"`
	QualificationId        uint                    `json:"qualificationId"`
	QualificationName      string                  `json:"qualificationName"`
	SyllabusId             uint                    `json:"syllabusId"`
	SyllabusCode           string                  `json:"syllabusCode"`
	SyllabusName           string                  `json:"syllabusName"`
	QuestionNumber         int                     `json:"questionNumber"`
	InsertedQuestionNumber int                     `json:"insertedQuestionNumber"`
	Questions              []*QuestionRandomPapers `json:"questions"`
}
