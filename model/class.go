package model

type ClassDetailView struct {
	Class
	VocabularySetList []*VocabularySet `json:"vocabularySetList"`
	UserList          []*User          `json:"userList"`
}
