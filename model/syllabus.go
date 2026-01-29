package model

type Syllabus struct {
	Model
	Name            string        `json:"name"`
	Code            string        `json:"code"`
	QualificationId uint          `json:"qualificationId"`
	Qualification   Qualification `json:"qualification"`
	// Chapters        []Chapter     `json:"chapters"`
}

type SyllabusQuery struct {
	ID              uint   `json:"id"`
	Code            string `json:"code"`
	QualificationId uint   `json:"qualificationId"`
	Page
}
