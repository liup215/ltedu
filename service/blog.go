package service

import (
	"edu/model"
	"edu/repository"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var BlogSvr = &BlogService{baseService: newBaseService()}

type BlogService struct {
	baseService
}

// generateSlug creates a URL-friendly slug from a title
func generateSlug(title string) string {
	slug := strings.ToLower(title)
	// Replace spaces and special chars with hyphens
	re := regexp.MustCompile(`[^a-z0-9\u4e00-\u9fff]+`)
	slug = re.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	if slug == "" {
		slug = fmt.Sprintf("post-%d", time.Now().UnixNano())
	}
	// Append timestamp to ensure uniqueness
	slug = fmt.Sprintf("%s-%d", slug, time.Now().Unix())
	return slug
}

// CreatePost creates a new blog post
func (s *BlogService) CreatePost(authorId uint, req model.BlogPostCreateRequest) (*model.BlogPost, error) {
	if req.Status == "" {
		req.Status = "draft"
	}
	if req.Status != "draft" && req.Status != "published" {
		return nil, errors.New("status must be 'draft' or 'published'")
	}

	post := &model.BlogPost{
		Title:      req.Title,
		Slug:       generateSlug(req.Title),
		Summary:    req.Summary,
		Content:    req.Content,
		CoverImage: req.CoverImage,
		Category:   req.Category,
		Tags:       req.Tags,
		AuthorId:   authorId,
		Status:     req.Status,
	}

	if err := repository.BlogPostRepo.Create(post); err != nil {
		return nil, err
	}
	return post, nil
}

// UpdatePost updates an existing blog post
func (s *BlogService) UpdatePost(authorId uint, req model.BlogPostUpdateRequest) (*model.BlogPost, error) {
	post, err := repository.BlogPostRepo.GetByID(req.ID)
	if err != nil {
		return nil, errors.New("post not found")
	}

	// Only author or admin can edit (caller should verify admin separately if needed)
	if post.AuthorId != authorId {
		return nil, errors.New("unauthorized to edit this post")
	}

	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Summary != "" {
		post.Summary = req.Summary
	}
	if req.Content != "" {
		post.Content = req.Content
	}
	if req.CoverImage != "" {
		post.CoverImage = req.CoverImage
	}
	if req.Category != "" {
		post.Category = req.Category
	}
	if req.Tags != "" {
		post.Tags = req.Tags
	}
	if req.Status != "" {
		if req.Status != "draft" && req.Status != "published" {
			return nil, errors.New("status must be 'draft' or 'published'")
		}
		post.Status = req.Status
	}
	if req.IsTop != nil {
		post.IsTop = *req.IsTop
	}

	if err := repository.BlogPostRepo.Update(post); err != nil {
		return nil, err
	}
	return post, nil
}

// DeletePost soft-deletes a blog post
func (s *BlogService) DeletePost(authorId uint, postId uint) error {
	post, err := repository.BlogPostRepo.GetByID(postId)
	if err != nil {
		return errors.New("post not found")
	}
	if post.AuthorId != authorId {
		return errors.New("unauthorized to delete this post")
	}
	return repository.BlogPostRepo.Delete(postId)
}

// GetByID retrieves a post by ID and increments view count for published posts
func (s *BlogService) GetByID(id uint) (*model.BlogPost, error) {
	post, err := repository.BlogPostRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("post not found")
	}
	if post.Status == "published" {
		_ = repository.BlogPostRepo.IncrementViewCount(id)
	}
	return post, nil
}

// GetBySlug retrieves a post by its URL slug
func (s *BlogService) GetBySlug(slug string) (*model.BlogPost, error) {
	post, err := repository.BlogPostRepo.GetBySlug(slug)
	if err != nil {
		return nil, errors.New("post not found")
	}
	if post.Status == "published" {
		_ = repository.BlogPostRepo.IncrementViewCount(post.ID)
	}
	return post, nil
}

// ListPublished lists published blog posts (public)
func (s *BlogService) ListPublished(query model.BlogPostQuery) ([]model.BlogPost, int64, error) {
	query.Status = "published"
	query.CheckPage()
	return repository.BlogPostRepo.List(&query)
}

// ListAll lists all posts for admin (including drafts)
func (s *BlogService) ListAll(query model.BlogPostQuery) ([]model.BlogPost, int64, error) {
	query.CheckPage()
	return repository.BlogPostRepo.List(&query)
}

// AdminUpdatePost allows admin to update any post
func (s *BlogService) AdminUpdatePost(req model.BlogPostUpdateRequest) (*model.BlogPost, error) {
	post, err := repository.BlogPostRepo.GetByID(req.ID)
	if err != nil {
		return nil, errors.New("post not found")
	}

	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Summary != "" {
		post.Summary = req.Summary
	}
	if req.Content != "" {
		post.Content = req.Content
	}
	if req.CoverImage != "" {
		post.CoverImage = req.CoverImage
	}
	if req.Category != "" {
		post.Category = req.Category
	}
	if req.Tags != "" {
		post.Tags = req.Tags
	}
	if req.Status != "" {
		if req.Status != "draft" && req.Status != "published" {
			return nil, errors.New("status must be 'draft' or 'published'")
		}
		post.Status = req.Status
	}
	if req.IsTop != nil {
		post.IsTop = *req.IsTop
	}

	if err := repository.BlogPostRepo.Update(post); err != nil {
		return nil, err
	}
	return post, nil
}

// AdminDeletePost allows admin to delete any post
func (s *BlogService) AdminDeletePost(postId uint) error {
	_, err := repository.BlogPostRepo.GetByID(postId)
	if err != nil {
		return errors.New("post not found")
	}
	return repository.BlogPostRepo.Delete(postId)
}
