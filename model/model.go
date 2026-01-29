package model

import "time"

const (
	YES = 1
	NO  = 2
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

func (m Model) GetID() uint {
	return m.ID
}

type Page struct {
	PageSize  int `json:"pageSize"`
	PageIndex int `json:"pageIndex"`
}

func (p Page) CheckPage() Page {
	if p.PageIndex <= 0 {
		p.PageIndex = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 20
	}
	return p
}
