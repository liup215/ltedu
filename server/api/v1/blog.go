package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var BlogCtrl = &BlogController{
	svr: service.BlogSvr,
}

type BlogController struct {
	svr *service.BlogService
}

// ListPosts lists published blog posts (public)
// POST /api/v1/blog/list
func (ctrl *BlogController) ListPosts(c *gin.Context) {
	var query model.BlogPostQuery
	if err := c.BindJSON(&query); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	posts, total, err := ctrl.svr.ListPublished(query)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", gin.H{
		"list":  posts,
		"total": total,
	})
}

// GetPost gets a blog post by ID (public)
// POST /api/v1/blog/byId
func (ctrl *BlogController) GetPost(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	post, err := ctrl.svr.GetByID(req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	if post.Status != "published" {
		http.ErrorData(c, "Post not found", nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", post)
}

// GetPostBySlug gets a blog post by slug (public)
// POST /api/v1/blog/bySlug
func (ctrl *BlogController) GetPostBySlug(c *gin.Context) {
	var req struct {
		Slug string `json:"slug" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	post, err := ctrl.svr.GetBySlug(req.Slug)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	if post.Status != "published" {
		http.ErrorData(c, "Post not found", nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", post)
}

// CreatePost creates a new blog post (requires auth)
// POST /api/v1/blog/create
func (ctrl *BlogController) CreatePost(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req model.BlogPostCreateRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	post, err := ctrl.svr.CreatePost(u.ID, req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Post created successfully!", post)
}

// UpdatePost updates a blog post (requires auth, author only)
// POST /api/v1/blog/edit
func (ctrl *BlogController) UpdatePost(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req model.BlogPostUpdateRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	post, err := ctrl.svr.UpdatePost(u.ID, req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Post updated successfully!", post)
}

// DeletePost deletes a blog post (requires auth, author only)
// POST /api/v1/blog/delete
func (ctrl *BlogController) DeletePost(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	if err := ctrl.svr.DeletePost(u.ID, req.ID); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Post deleted successfully!", nil)
}

// AdminListPosts lists all posts including drafts (admin only)
// POST /api/v1/blog/admin/list
func (ctrl *BlogController) AdminListPosts(c *gin.Context) {
	var query model.BlogPostQuery
	if err := c.BindJSON(&query); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	posts, total, err := ctrl.svr.ListAll(query)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", gin.H{
		"list":  posts,
		"total": total,
	})
}

// AdminGetPost gets any post by ID regardless of status (admin only)
// POST /api/v1/blog/admin/byId
func (ctrl *BlogController) AdminGetPost(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	post, err := ctrl.svr.GetByID(req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", post)
}

// AdminUpdatePost allows admin to update any post
// POST /api/v1/blog/admin/edit
func (ctrl *BlogController) AdminUpdatePost(c *gin.Context) {
	var req model.BlogPostUpdateRequest
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	post, err := ctrl.svr.AdminUpdatePost(req)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Post updated successfully!", post)
}

// AdminDeletePost allows admin to delete any post
// POST /api/v1/blog/admin/delete
func (ctrl *BlogController) AdminDeletePost(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	if err := ctrl.svr.AdminDeletePost(req.ID); err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Post deleted successfully!", nil)
}
