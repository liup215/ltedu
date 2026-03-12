package model

const (
	BlogCategorySystemUpdates  = "system_updates"
	BlogCategoryUserGuides     = "user_guides"
	BlogCategoryLearningMethods = "learning_methods"
	BlogCategoryMajorEvents    = "major_events"
)

const (
	BlogStatusDraft     = "draft"
	BlogStatusPublished = "published"
)

type BlogPost struct {
	Model
	Title       string `json:"title"`
	Slug        string `json:"slug" gorm:"type:varchar(255);uniqueIndex"`
	Summary     string `json:"summary"`
	Content     string `json:"content" gorm:"type:text"`
	Category    string `json:"category"`
	Tags        string `json:"tags"`
	CoverImage  string `json:"coverImage"`
	Status      string `json:"status"`
	AuthorId    uint   `json:"authorId"`
	AuthorName  string `json:"authorName"`
	ViewCount   int    `json:"viewCount"`
	PublishedAt *string `json:"publishedAt,omitempty"`
}

type BlogPostCreateEditRequest struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Summary     string `json:"summary"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	Tags        string `json:"tags"`
	CoverImage  string `json:"coverImage"`
	Status      string `json:"status"`
}

type BlogPostQueryRequest struct {
	ID       uint   `json:"id"`
	Category string `json:"category"`
	Status   string `json:"status"`
	Keyword  string `json:"keyword"`
	Page
}
