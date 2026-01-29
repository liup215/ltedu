package model

type AiVocabularyCheckRequest struct {
	Stem          string `json:"stem"`
	Answer        string `json:"answer"`
	StudentAnswer string `json:"studentAnswer"`
}

type AiQuestionCheckRequest struct {
	QuestionId    uint   `json:"questionId"`
	StudentAnswer string `json:"studentAnswer"`
}

type CheckResponse struct {
	Text   string `json:"text"`
	Score  int    `json:"score"`
	Answer string `json:"answer"`
}
