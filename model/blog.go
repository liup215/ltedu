package model

import "time"

// BlogPost represents a blog post for knowledge sharing
type BlogPost struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	Title      string `gorm:"size:255;not null" json:"title"`
	Slug       string `gorm:"size:255;uniqueIndex" json:"slug"`
	Summary    string `gorm:"size:500" json:"summary"`
	Content    string `gorm:"type:longtext" json:"content"`
	CoverImage string `gorm:"size:500" json:"coverImage"`
	Category   string `gorm:"size:100;index" json:"category"`
	Tags       string `gorm:"size:500" json:"tags"` // comma-separated
	AuthorId   uint   `gorm:"index" json:"authorId"`
	Status     string `gorm:"size:20;default:'draft'" json:"status"` // draft, published
	ViewCount  int    `gorm:"default:0" json:"viewCount"`
	IsTop      bool   `gorm:"default:false" json:"isTop"`
}

// BlogPostCreateRequest is used to create a blog post
type BlogPostCreateRequest struct {
	Title      string `json:"title" binding:"required"`
	Summary    string `json:"summary"`
	Content    string `json:"content" binding:"required"`
	CoverImage string `json:"coverImage"`
	Category   string `json:"category"`
	Tags       string `json:"tags"`
	Status     string `json:"status"` // draft or published
}

// BlogPostUpdateRequest is used to update a blog post
type BlogPostUpdateRequest struct {
	ID         uint   `json:"id" binding:"required"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Content    string `json:"content"`
	CoverImage string `json:"coverImage"`
	Category   string `json:"category"`
	Tags       string `json:"tags"`
	Status     string `json:"status"`
	IsTop      *bool  `json:"isTop"`
}

// BlogPostQuery is used to query/filter blog posts
type BlogPostQuery struct {
	Page
	Category string `json:"category"`
	Status   string `json:"status"` // empty = published only for public
	Keyword  string `json:"keyword"`
}
