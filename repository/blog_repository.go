package repository

import (
	"edu/model"
	"strings"

	"gorm.io/gorm"
)

type IBlogPostRepository interface {
	Create(post *model.BlogPost) error
	Update(post *model.BlogPost) error
	Delete(id uint) error
	FindByID(id uint) (*model.BlogPost, error)
	FindBySlug(slug string) (*model.BlogPost, error)
	FindByPage(query *model.BlogPostQueryRequest, offset, limit int) ([]*model.BlogPost, int64, error)
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
	return r.db.Model(post).Updates(post).Error
}

func (r *blogPostRepository) Delete(id uint) error {
	return r.db.Delete(&model.BlogPost{}, id).Error
}

func (r *blogPostRepository) FindByID(id uint) (*model.BlogPost, error) {
	var post model.BlogPost
	err := r.db.Where("id = ?", id).First(&post).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &post, err
}

func (r *blogPostRepository) FindBySlug(slug string) (*model.BlogPost, error) {
	var post model.BlogPost
	err := r.db.Where("slug = ?", slug).First(&post).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &post, err
}

func (r *blogPostRepository) FindByPage(query *model.BlogPostQueryRequest, offset, limit int) ([]*model.BlogPost, int64, error) {
	var posts []*model.BlogPost
	var total int64

	q := r.db.Model(&model.BlogPost{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.Category != "" {
		q = q.Where("category = ?", query.Category)
	}
	if query.Status != "" {
		q = q.Where("status = ?", query.Status)
	}
	if query.Keyword != "" {
		like := "%" + strings.TrimSpace(query.Keyword) + "%"
		q = q.Where("title LIKE ? OR summary LIKE ? OR tags LIKE ?", like, like, like)
	}

	q.Count(&total)

	err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&posts).Error
	return posts, total, err
}

func (r *blogPostRepository) IncrementViewCount(id uint) error {
	return r.db.Model(&model.BlogPost{}).Where("id = ?", id).UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}
