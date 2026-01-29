// ltedu-api/model/practice.go

package model

type PracticeQuickRequest struct {
	SyllabusId    uint   `json:"syllabusId"`
	QuestionCount int    `json:"questionCount"`
	ChapterIds    []uint `json:"chapterIds,omitempty"`
}

type PracticeQuickResponse struct {
	List  []uint `json:"list"`
	Total int    `json:"total"`
}

type PracticePaperRequest struct {
	PaperId uint `json:"paperId"`
}

type PracticePaperResponse struct {
	List []uint `json:"list"`
}

type PracticePartSubmission struct {
	QuestionContentId uint   `json:"questionContentId"`
	Answer            string `json:"answer"`
}

type PracticeSubmission struct {
	QuestionID uint                      `json:"questionId"`
	Answers    []*PracticePartSubmission `json:"answers"`
}

type PracticeGradeRequest []*PracticeSubmission

// GradePracticeResponse represents the response after grading a practice submission for each question.
type PracticeGradeResponse struct {
	Score   int                  `json:"score"`
	Total   int                  `json:"total"`
	Results []PracticeResultItem `json:"results"`
}

type PracticeSubResultItem struct {
	QuestionContentId uint   `json:"questionContentId"`
	QuestionType      int    `json:"questionType"`
	CorrectAnswer     string `json:"correctAnswer"`
	StudentAnswer     string `json:"studentAnswer"`
	IsCorrect         *bool  `json:"isCorrect"` // nil for subjective
	ModelAnswer       string `json:"modelAnswer,omitempty"`
}

type PracticeResultItem struct {
	QuestionID uint                    `json:"questionId"`
	SubResults []PracticeSubResultItem `json:"subResults"`
}

type PracticeResult struct {
	Score   int                  `json:"score"`
	Total   int                  `json:"total"`
	Results []PracticeResultItem `json:"results"`
}
