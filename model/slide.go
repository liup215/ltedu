package model

type Slide struct {
	Model
	Name        string    `json:"name"`
	Hash        string    `json:"hash"`
	Description string    `json:"description"`
	SyllabusId  uint      `json:"syllabusId"`
	Syllabus    *Syllabus `json:"syllabus"`
}

func (s *Slide) GetResponse() *SlideQueryResponse {
	return &SlideQueryResponse{
		ID:                s.ID,
		Name:              s.Name,
		Hash:              s.Hash,
		Description:       s.Description,
		OrganisationId:    s.Syllabus.Qualification.Organisation.ID,
		OrganisationName:  s.Syllabus.Qualification.Organisation.Name,
		QualificationId:   s.Syllabus.Qualification.ID,
		QualificationName: s.Syllabus.Qualification.Name,
		SyllabusId:        s.Syllabus.ID,
		SyllabusName:      s.Syllabus.Name,
		SyllabusCode:      s.Syllabus.Code,
	}
}

type SlideCreateEditRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SyllabusId  uint   `json:"syllabusId"`
}

type SlideQueryRequest struct {
	ID         uint `json:"id"`
	SyllabusId uint `json:"syllabusId"`
	Page
}

type SlideQueryResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	OrganisationId    uint   `json:"organisationId"`
	OrganisationName  string `json:"organisationName"`
	QualificationId   uint   `json:"qualificationId"`
	QualificationName string `json:"qualificationName"`
	SyllabusId        uint   `json:"syllabusId"`
	SyllabusName      string `json:"syllabusName"`
	SyllabusCode      string `json:"syllabusCode"`
	Hash              string `json:"hash"`
}
