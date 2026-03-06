package model

const (
	FeedbackTypeGeneral = "general"
	FeedbackTypeBug     = "bug"
	FeedbackTypeFeature = "feature"
	FeedbackTypePraise  = "praise"

	FeedbackSentimentPositive = "positive"
	FeedbackSentimentNeutral  = "neutral"
	FeedbackSentimentNegative = "negative"

	FeedbackStatusNew      = "new"
	FeedbackStatusReviewed = "reviewed"
	FeedbackStatusResolved = "resolved"
)

// UserFeedback stores user-submitted feedback with optional consent tracking (GDPR).
type UserFeedback struct {
	Model
	UserID       uint   `json:"userId" gorm:"index"`
	Type         string `json:"type"`        // general, bug, feature, praise
	Content      string `json:"content"`
	Rating       int    `json:"rating"`      // 1-5 stars (0 means not rated)
	Sentiment    string `json:"sentiment"`   // positive, neutral, negative (auto-computed)
	PageContext  string `json:"pageContext"` // page or feature the feedback refers to
	ConsentGiven bool   `json:"consentGiven"`
	Status       string `json:"status"`     // new, reviewed, resolved
	AdminNote    string `json:"adminNote"`
	UserAgent    string `json:"userAgent"`
}

// FeedbackListRequest is used to query feedback with optional filters.
type FeedbackListRequest struct {
	Page
	Status string `json:"status"`
	Type   string `json:"type"`
	UserID uint   `json:"userId"`
}

// FeedbackStats summarises feedback data for the admin dashboard.
type FeedbackStats struct {
	Total     int64          `json:"total"`
	ByType    map[string]int `json:"byType"`
	BySentiment map[string]int `json:"bySentiment"`
	ByStatus  map[string]int `json:"byStatus"`
	AvgRating float64        `json:"avgRating"`
}
