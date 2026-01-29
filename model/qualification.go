package model

type Qualification struct {
	Model
	Name           string       `json:"name"`
	OrganisationId uint         `json:"organisationId"`
	Organisation   Organisation `json:"organisation"`
	// Syllabuses     []Syllabus   `json:"syllabuses"`
}

type QualificationQuery struct {
	ID             uint `json:"id"`
	OrganisationId uint `json:"organisationId"`
	Page
}

// type Subject struct {
// 	Model
// 	Name            string        `json:"name"` // 语文 数学 英语 等
// 	QualificationId uint          `json:"qualificationId"`
// 	Qualification   Qualification `json:"qualification"`
// 	Syllabuses      []Syllabus    `json:"syllabuses"`
// }

// type SubjectQuery struct {
// 	ID              uint `json:"id"`
// 	OrganisationId  uint `json:"organisationId"`
// 	QualificationId uint `json:"qualificationId"`
// 	Page
// }
