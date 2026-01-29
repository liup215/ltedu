package model

type Organisation struct {
	Model
	Name string `json:"name"`
}

type OrganisationQuery struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Page
}
