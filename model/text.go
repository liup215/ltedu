package model

type TextContent struct {
	Model
	Content string `gorm:"type:TEXT" json:"content"`
}
