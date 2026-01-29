package model

import "time"

// protected $fillable = [
//         'user_id', 'title', 'slug', 'thumb', 'charge',
//         'short_description', 'original_desc', 'render_desc', 'seo_keywords',
//         'seo_description', 'published_at', 'is_show', 'category_id',
//         'is_rec', 'user_count', 'is_free',
//     ];

type Course struct {
	Model
	UserID            int        `json:"userId"`
	UserType          int        `json:"userType"`
	UserTypeName      string     `json:"userTypeName" gorm:"-"`
	Title             string     `json:"title"`
	Slug              string     `json:"slug"`
	Thumb             string     `json:"thumb"`
	Charge            int        `json:"charge"`
	ShortDescription  string     `json:"shortDescription"`
	OriginalDesc      string     `json:"originalDesc"`
	RenderDesc        string     `json:"renderDesc"`
	SeoKeywords       string     `json:"seoKeywords"`
	SeoDescription    string     `json:"seoDescription"`
	PublishedAtString string     `json:"publishedAt" gorm:"-"`
	PublishedAt       *time.Time `json:"-"`
	IsShow            int        `json:"isShow"`
	IsRec             int        `json:"isRec"`
	UserCount         int        `json:"userCount"`
	IsFree            int        `json:"isFree"`
	// MediaVideos       []MediaVideo `json:"mediaVideos"`
	SyllabusID int      `json:"syllabusId"`
	Syllabus   Syllabus `json:"syllabus"`
}

type CourseCreateEditRequest struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Thumb            string `json:"thumb"`
	Charge           int    `json:"charge"`
	ShortDescription string `json:"shortDescription"`
	OriginalDesc     string `json:"originalDesc"`
	PublishedAt      string `json:"publishedAt" gorm:"-"`
	IsShow           int    `json:"isShow"`
	IsFree           int    `json:"isFree"`
	SyllabusID       int    `json:"syllabusId"`
}

type CourseQueryRequest struct {
	Model
	SyllabusID int `json:"syllabusId"`
	Page
}
