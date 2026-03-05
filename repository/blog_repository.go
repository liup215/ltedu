package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

type IBlogPostRepository interface {
	Create(post *model.BlogPost) error
	Update(post *model.BlogPost) error
	Delete(id uint) error
	GetByID(id uint) (*model.BlogPost, error)
	GetBySlug(slug string) (*model.BlogPost, error)
	List(query *model.BlogPostQuery) ([]model.BlogPost, int64, error)
	IncrementViewCount(id uint) error
}

type blogPostRepository struct {
	db *gorm.DB
}

func NewBlogPostRepository(db *gorm.DB) IBlogPostRepository {
	return &blogPostRepository{db: db}
}

func (r *blogPostRepository) Create(post *model.BlogPost) error {
	return r.db.Create(post).Error
}

func (r *blogPostRepository) Update(post *model.BlogPost) error {
	return r.db.Save(post).Error
}

func (r *blogPostRepository) Delete(id uint) error {
	return r.db.Delete(&model.BlogPost{}, id).Error
}

func (r *blogPostRepository) GetByID(id uint) (*model.BlogPost, error) {
	var post model.BlogPost
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *blogPostRepository) GetBySlug(slug string) (*model.BlogPost, error) {
	var post model.BlogPost
	err := r.db.Where("slug = ? AND deleted_at IS NULL", slug).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *blogPostRepository) List(query *model.BlogPostQuery) ([]model.BlogPost, int64, error) {
	var posts []model.BlogPost
	var total int64

	q := r.db.Model(&model.BlogPost{}).Where("deleted_at IS NULL")

	if query.Status != "" {
		q = q.Where("status = ?", query.Status)
	} else {
		q = q.Where("status = 'published'")
	}

	if query.Category != "" {
		q = q.Where("category = ?", query.Category)
	}

	if query.Keyword != "" {
		like := "%" + query.Keyword + "%"
		q = q.Where("title LIKE ? OR summary LIKE ?", like, like)
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	page := query.PageIndex
	if page <= 0 {
		page = 1
	}
	pageSize := query.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	err := q.Order("is_top DESC, created_at DESC").
		Offset(offset).Limit(pageSize).
		Find(&posts).Error
	return posts, total, err
}

func (r *blogPostRepository) IncrementViewCount(id uint) error {
	return r.db.Model(&model.BlogPost{}).Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}
