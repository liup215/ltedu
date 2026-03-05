package service

import (
	"edu/model"
	"edu/repository"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var BlogSvr = &BlogService{baseService: newBaseService()}

type BlogService struct {
	baseService
}

func (s *BlogService) CreateBlogPost(req model.BlogPostCreateEditRequest, authorId uint, authorName string) error {
	if strings.TrimSpace(req.Title) == "" {
		return errors.New("title cannot be empty")
	}
	if strings.TrimSpace(req.Content) == "" {
		return errors.New("content cannot be empty")
	}
	if !isValidCategory(req.Category) {
		return errors.New("invalid category")
	}

	slug := req.Slug
	if slug == "" {
		slug = generateSlug(req.Title)
	}
	// Ensure slug uniqueness
	slug = s.ensureUniqueSlug(slug, 0)

	status := req.Status
	if status == "" {
		status = model.BlogStatusDraft
	}

	var publishedAt *string
	if status == model.BlogStatusPublished {
		now := time.Now().Format(time.RFC3339)
		publishedAt = &now
	}

	post := &model.BlogPost{
		Title:       req.Title,
		Slug:        slug,
		Summary:     req.Summary,
		Content:     req.Content,
		Category:    req.Category,
		Tags:        req.Tags,
		CoverImage:  req.CoverImage,
		Status:      status,
		AuthorId:    authorId,
		AuthorName:  authorName,
		PublishedAt: publishedAt,
	}

	return repository.BlogPostRepo.Create(post)
}

func (s *BlogService) EditBlogPost(req model.BlogPostCreateEditRequest) error {
	if req.ID == 0 {
		return errors.New("id cannot be empty")
	}
	if strings.TrimSpace(req.Title) == "" {
		return errors.New("title cannot be empty")
	}
	if !isValidCategory(req.Category) {
		return errors.New("invalid category")
	}

	existing, err := repository.BlogPostRepo.FindByID(req.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("blog post not found")
	}

	slug := req.Slug
	if slug == "" {
		slug = generateSlug(req.Title)
	}
	slug = s.ensureUniqueSlug(slug, req.ID)

	var publishedAt *string
	if req.Status == model.BlogStatusPublished && existing.Status != model.BlogStatusPublished {
		now := time.Now().Format(time.RFC3339)
		publishedAt = &now
	} else {
		publishedAt = existing.PublishedAt
	}

	post := &model.BlogPost{
		Model:       model.Model{ID: req.ID},
		Title:       req.Title,
		Slug:        slug,
		Summary:     req.Summary,
		Content:     req.Content,
		Category:    req.Category,
		Tags:        req.Tags,
		CoverImage:  req.CoverImage,
		Status:      req.Status,
		PublishedAt: publishedAt,
	}

	return repository.BlogPostRepo.Update(post)
}

func (s *BlogService) DeleteBlogPost(id uint) error {
	if id == 0 {
		return errors.New("id cannot be empty")
	}
	return repository.BlogPostRepo.Delete(id)
}

func (s *BlogService) GetBlogPostById(id uint) (*model.BlogPost, error) {
	if id == 0 {
		return nil, errors.New("id cannot be empty")
	}
	return repository.BlogPostRepo.FindByID(id)
}

func (s *BlogService) GetBlogPostBySlug(slug string) (*model.BlogPost, error) {
	if slug == "" {
		return nil, errors.New("slug cannot be empty")
	}
	post, err := repository.BlogPostRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}
	if post != nil {
		_ = repository.BlogPostRepo.IncrementViewCount(post.ID)
	}
	return post, nil
}

func (s *BlogService) ListBlogPosts(q model.BlogPostQueryRequest) ([]*model.BlogPost, int64, error) {
	q.Page = q.Page.CheckPage()
	offset := (q.PageIndex - 1) * q.PageSize
	return repository.BlogPostRepo.FindByPage(&q, offset, q.PageSize)
}

func isValidCategory(cat string) bool {
	switch cat {
	case model.BlogCategorySystemUpdates,
		model.BlogCategoryUserGuides,
		model.BlogCategoryLearningMethods,
		model.BlogCategoryMajorEvents:
		return true
	}
	return false
}

func generateSlug(title string) string {
	// Lowercase
	s := strings.ToLower(title)
	// Replace non-ASCII with empty (Chinese chars become empty after lower)
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		} else if unicode.IsSpace(r) || r == '-' || r == '_' {
			b.WriteRune('-')
		}
	}
	slug := b.String()
	// Collapse multiple dashes
	re := regexp.MustCompile(`-+`)
	slug = re.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	if slug == "" {
		slug = fmt.Sprintf("post-%d", time.Now().Unix())
	}
	return slug
}

func (s *BlogService) ensureUniqueSlug(slug string, excludeID uint) string {
	base := slug
	for i := 1; ; i++ {
		existing, err := repository.BlogPostRepo.FindBySlug(slug)
		if err != nil || existing == nil {
			return slug
		}
		if excludeID != 0 && existing.ID == excludeID {
			return slug
		}
		slug = fmt.Sprintf("%s-%d", base, i)
	}
}
