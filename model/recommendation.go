package model

// QuestionRecommendation represents a single recommended question for a student.
type QuestionRecommendation struct {
	QuestionID      uint    `json:"questionId"`
	Question        *Question `json:"question,omitempty"`
	KnowledgePointID uint   `json:"knowledgePointId"`
	KnowledgePoint  string  `json:"knowledgePoint"`
	ChapterID       uint    `json:"chapterId"`
	ChapterName     string  `json:"chapterName"`
	Reason          string  `json:"reason"`
	Priority        int     `json:"priority"`    // 1 = highest priority
	Difficulty      int     `json:"difficulty"`  // 1-5, adapted to student level
	MasteryLevel    float64 `json:"masteryLevel"` // current mastery of the chapter (0.0 – 1.0)
	Score           float64 `json:"score"`        // internal ranking score (higher = more recommended)
}

// QuestionRecommendationResponse is the full recommendation result for a student.
type QuestionRecommendationResponse struct {
	StudentID       uint                     `json:"studentId"`
	Recommendations []QuestionRecommendation `json:"recommendations"`
	TotalCount      int                      `json:"totalCount"`
	WeakAreaCount   int                      `json:"weakAreaCount"`   // number of chapters below mastery threshold
	ReviewDueCount  int                      `json:"reviewDueCount"`  // chapters due for SRS review
}
