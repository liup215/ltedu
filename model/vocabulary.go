package model

type VocabularySet struct {
	Model
	Name        string            `json:"name"`
	Description string            `json:"description"`
	SyllabusId  uint              `json:"syllabusId"`
	Syllabus    *Syllabus         `json:"syllabus"`
	Words       []*VocabularyItem `json:"words"`
}

func (vs *VocabularySet) ToView() *VocabularySetView {
	return &VocabularySetView{
		Model:             vs.Model,
		Name:              vs.Name,
		Description:       vs.Description,
		OrganisationId:    vs.Syllabus.Qualification.OrganisationId,
		OrganisationName:  vs.Syllabus.Qualification.Organisation.Name,
		QualificationId:   vs.Syllabus.QualificationId,
		QualificationName: vs.Syllabus.Qualification.Name,
		SyllabusId:        vs.SyllabusId,
		SyllabusName:      vs.Syllabus.Name,
		SyllabusCode:      vs.Syllabus.Code,
		WordNumber:        len(vs.Words),
		Words:             vs.Words,
	}
}

type VocabularySetCreateEditRequest struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	SyllabusId  uint              `json:"syllabusId"`
	Words       []*VocabularyItem `json:"words"`
}

type VocabularySetQuery struct {
ID              uint   `json:"id"`
OrganisationId  uint   `json:"organisationId"`
QualificationId uint   `json:"qualificationId"`
SyllabusId      uint   `json:"syllabusId"`
Name            string `json:"name"`
Page
}

type VocabularySetView struct {
	Model
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	OrganisationId    uint              `json:"organisationId"`
	OrganisationName  string            `json:"organisationName"`
	QualificationId   uint              `json:"qualificationId"`
	QualificationName string            `json:"qualificationName"`
	SyllabusId        uint              `json:"syllabusId"`
	SyllabusName      string            `json:"syllabusName"`
	SyllabusCode      string            `json:"syllabusCode"`
	WordNumber        int               `json:"wordNumber"`
	Words             []*VocabularyItem `json:"words"`
}

type VocabularyItem struct {
	Model
	VocabularySetId uint   `json:"vocabularySetId"`
	Key             string `json:"key"`
	Value           string `json:"value"`
	Image           string `json:"image"`
	Order           int    `json:"order"`
}

type VocabularyItemCreateEditRequest struct {
	ID              uint   `json:"id"`
	VocabularySetId uint   `json:"vocabularySetId"`
	Key             string `json:"key"`
	Value           string `json:"value"`
	Image           string `json:"image"`
	Order           int    `json:"order"`
}

type VocabularyItemQuery struct {
Model
VocabularySetId uint   `json:"vocabularySetId"`
Key             string `json:"key"`
}

type VocabularySetTestQuestion struct {
	Word                *VocabularyItem   `json:"word"`
	Options             []*VocabularyItem `json:"options"`
	LearnerAnswer       uint              `json:"learnerAnswer"`
	LearnerAnswerString string            `json:"learnerAnswerString"`
	AiComment           string            `json:"aiComment"`
	Score               int               `json:"score"`
}

type VocabularySetTestRequest struct {
	Id uint `json:"id"`
}

type VocabularySetTestResponse struct {
	VocabularySetView
	Questions []*VocabularySetTestQuestion `json:"questions"`
}

type VocabularySetTest struct {
	Model
	VocabularySetId uint                         `json:"vocabularySetId"`
	VocabularySet   VocabularySet                `json:"vocabularySet"`
	UserId          uint                         `json:"userId"`
	User            User                         `json:"user"`
	Questions       []*VocabularySetTestQuestion `json:"questions" gorm:"-"`
	QuestionText    string                       `json:"-" gorm:"type:text;default:[]"`
}

func (vst *VocabularySetTest) QuestionToText() string {
	return ""
}
