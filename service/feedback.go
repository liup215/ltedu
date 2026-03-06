package service

import (
	"edu/model"
	"edu/repository"
	"strings"
)

// FeedbackSvr is the singleton instance of FeedbackService.
var FeedbackSvr = &FeedbackService{}

// FeedbackService handles user feedback submission and management.
type FeedbackService struct{}

// positiveKeywords and negativeKeywords drive the simple keyword-based sentiment analysis.
var positiveKeywords = []string{
	"great", "excellent", "love", "amazing", "awesome", "helpful", "good", "nice",
	"fantastic", "wonderful", "perfect", "easy", "useful", "intuitive", "clear",
	"thank", "happy", "pleased", "impressed", "brilliant", "superb", "best",
	"非常好", "很棒", "喜欢", "满意", "优秀", "感谢", "好用", "方便", "清晰",
}

var negativeKeywords = []string{
	"bad", "bug", "error", "crash", "broken", "slow", "terrible", "awful",
	"horrible", "difficult", "confusing", "unclear", "useless", "fail", "wrong",
	"problem", "issue", "fix", "missing", "frustrating", "disappointed", "hate",
	"not working", "doesn't work", "can't", "cannot", "unable",
	"很差", "问题", "错误", "崩溃", "慢", "难用", "不好", "失败", "找不到", "无法",
}

// AnalyseSentiment returns positive, neutral, or negative based on content keywords.
// When a star rating is provided it also influences the result.
func AnalyseSentiment(content string, rating int) string {
	lower := strings.ToLower(content)

	posScore := 0
	for _, kw := range positiveKeywords {
		if strings.Contains(lower, kw) {
			posScore++
		}
	}
	negScore := 0
	for _, kw := range negativeKeywords {
		if strings.Contains(lower, kw) {
			negScore++
		}
	}

	// Incorporate star rating: ratings 4-5 → positive, 1-2 → negative.
	if rating >= 4 {
		posScore += 2
	} else if rating == 3 {
		// neutral — no adjustment
	} else if rating > 0 && rating <= 2 {
		negScore += 2
	}

	switch {
	case posScore > negScore:
		return model.FeedbackSentimentPositive
	case negScore > posScore:
		return model.FeedbackSentimentNegative
	default:
		return model.FeedbackSentimentNeutral
	}
}

// Submit validates and persists a new UserFeedback record.
func (s *FeedbackService) Submit(feedback *model.UserFeedback) error {
	if !feedback.ConsentGiven {
		return errFeedbackConsentRequired
	}

	// Compute sentiment automatically.
	feedback.Sentiment = AnalyseSentiment(feedback.Content, feedback.Rating)
	feedback.Status = model.FeedbackStatusNew

	return repository.FeedbackRepo.Create(feedback)
}

// List returns paginated feedback visible to admins.
func (s *FeedbackService) List(req model.FeedbackListRequest) ([]*model.UserFeedback, int64, error) {
	return repository.FeedbackRepo.FindAll(req)
}

// MyFeedback returns the feedback submitted by a specific user.
func (s *FeedbackService) MyFeedback(userID uint, page model.Page) ([]*model.UserFeedback, int64, error) {
	return repository.FeedbackRepo.FindByUserID(userID, page)
}

// GetByID returns a single feedback record.
func (s *FeedbackService) GetByID(id uint) (*model.UserFeedback, error) {
	return repository.FeedbackRepo.FindByID(id)
}

// UpdateStatus allows an admin to change the status and add a note.
func (s *FeedbackService) UpdateStatus(id uint, status, adminNote string) error {
	validStatuses := map[string]bool{
		model.FeedbackStatusNew:      true,
		model.FeedbackStatusReviewed: true,
		model.FeedbackStatusResolved: true,
	}
	if !validStatuses[status] {
		return errFeedbackInvalidStatus
	}
	return repository.FeedbackRepo.UpdateStatus(id, status, adminNote)
}

// GetStats returns aggregate statistics for the admin dashboard.
func (s *FeedbackService) GetStats() (*model.FeedbackStats, error) {
	return repository.FeedbackRepo.GetStats()
}
