package model

// NLU intent constants represent the primary goal of a user's educational query.
const (
	IntentFindQuestion   = "find_question"    // User wants to find a question
	IntentCheckAnswer    = "check_answer"     // User wants to check their answer
	IntentExplainConcept = "explain_concept"  // User wants a concept explained
	IntentPractice       = "practice"         // User wants to start practice
	IntentLearningPlan   = "get_learning_plan" // User wants a learning plan
	IntentGetProgress    = "get_progress"     // User wants to see their progress
	IntentGetSyllabus    = "get_syllabus"     // User wants syllabus information
	IntentGeneralQuery   = "general_query"    // General educational query
)

// NLU language constants for detected input language.
const (
	LanguageChinese = "zh"
	LanguageEnglish = "en"
	LanguageMixed   = "mixed"
)

// NLUEntity holds a single extracted entity from a user query.
type NLUEntity struct {
	Type  string `json:"type"`  // e.g. "subject", "chapter", "question_type", "difficulty", "topic"
	Value string `json:"value"` // extracted value
}

// NLUResult is the output of the NLU analysis for a single query.
type NLUResult struct {
	Intent     string      `json:"intent"`
	Confidence float64     `json:"confidence"`
	Entities   []NLUEntity `json:"entities"`
	Language   string      `json:"language"`
	Normalized string      `json:"normalized"` // cleaned/normalized query text
}

// NLURequest is the API request payload for NLU analysis.
type NLURequest struct {
	Query   string `json:"query"`
	Context string `json:"context,omitempty"` // optional surrounding context
}

// NLUFeedback is used to submit user corrections and improve future accuracy.
type NLUFeedback struct {
	Model
	UserID          uint   `json:"userId"   gorm:"column:user_id"`
	Query           string `json:"query"    gorm:"type:text"`
	PredictedIntent string `json:"predictedIntent"`
	CorrectIntent   string `json:"correctIntent"`
	Feedback        string `json:"feedback"  gorm:"type:text"`
}

// NLUFeedbackRequest is the API request for submitting NLU feedback.
type NLUFeedbackRequest struct {
	Query           string `json:"query"`
	PredictedIntent string `json:"predictedIntent"`
	CorrectIntent   string `json:"correctIntent"`
	Feedback        string `json:"feedback"`
}
