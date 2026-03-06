package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var BlogCtrl = &BlogController{
	blogSvr: service.BlogSvr,
}

type BlogController struct {
	blogSvr *service.BlogService
}

// @Summary      Create blog post
// @Description  Admin creates a new blog post
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        body  body  model.BlogPostCreateEditRequest  true  "Blog post"
// @Success      200   {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /v1/blog/create [post]
func (ctrl *BlogController) CreateBlogPost(c *gin.Context) {
	req := model.BlogPostCreateEditRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "获取用户信息失败", nil)
		return
	}
	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || user == nil {
		http.ErrorData(c, "用户不存在", nil)
		return
	}

	if err := ctrl.blogSvr.CreateBlogPost(req, user.ID, user.Username); err != nil {
		http.ErrorData(c, "创建失败: "+err.Error(), nil)
		return
	}

	http.SuccessData(c, "创建成功", nil)
}

// @Summary      Edit blog post
// @Description  Admin edits an existing blog post
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        body  body  model.BlogPostCreateEditRequest  true  "Blog post"
// @Success      200   {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /v1/blog/edit [post]
func (ctrl *BlogController) EditBlogPost(c *gin.Context) {
	req := model.BlogPostCreateEditRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.blogSvr.EditBlogPost(req); err != nil {
		http.ErrorData(c, "编辑失败: "+err.Error(), nil)
		return
	}

	http.SuccessData(c, "编辑成功", nil)
}

// @Summary      Delete blog post
// @Description  Admin deletes a blog post
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        body  body  model.BlogPostQueryRequest  true  "Blog post ID"
// @Success      200   {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /v1/blog/delete [post]
func (ctrl *BlogController) DeleteBlogPost(c *gin.Context) {
	req := model.BlogPostQueryRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.blogSvr.DeleteBlogPost(req.ID); err != nil {
		http.ErrorData(c, "删除失败: "+err.Error(), nil)
		return
	}

	http.SuccessData(c, "删除成功", nil)
}

// @Summary      Get blog post by ID (admin)
// @Description  Get blog post details by ID (admin, any status)
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        body  body  model.BlogPostQueryRequest  true  "Blog post ID"
// @Success      200   {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /v1/blog/byId [post]
func (ctrl *BlogController) GetBlogPostById(c *gin.Context) {
	req := model.BlogPostQueryRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	post, err := ctrl.blogSvr.GetBlogPostById(req.ID)
	if err != nil || post == nil {
		http.ErrorData(c, "未找到", nil)
		return
	}

	http.SuccessData(c, "获取成功", post)
}

// @Summary      List blog posts (admin)
// @Description  Paginated list of blog posts for admin (all statuses)
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        body  body  model.BlogPostQueryRequest  true  "Query"
// @Success      200   {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /v1/blog/list [post]
func (ctrl *BlogController) ListBlogPosts(c *gin.Context) {
	req := model.BlogPostQueryRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	list, total, err := ctrl.blogSvr.ListBlogPosts(req)
	if err != nil {
		http.ErrorData(c, "查询失败: "+err.Error(), nil)
		return
	}

	http.SuccessData(c, "获取成功", gin.H{
		"list":  list,
		"total": total,
	})
}

// @Summary      Public blog post list
// @Description  List published blog posts for public consumption
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        body  body  model.BlogPostQueryRequest  true  "Query"
// @Success      200   {object}  map[string]interface{}
// @Router       /v1/blog/public/list [post]
func (ctrl *BlogController) PublicListBlogPosts(c *gin.Context) {
	req := model.BlogPostQueryRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	// Force published only for public endpoint
	req.Status = model.BlogStatusPublished

	list, total, err := ctrl.blogSvr.ListBlogPosts(req)
	if err != nil {
		http.ErrorData(c, "查询失败: "+err.Error(), nil)
		return
	}

	http.SuccessData(c, "获取成功", gin.H{
		"list":  list,
		"total": total,
	})
}

// @Summary      Get published blog post by slug
// @Description  Get a single published blog post by its slug (public)
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        body  body  object  true  "slug"
// @Success      200   {object}  map[string]interface{}
// @Router       /v1/blog/public/bySlug [post]
func (ctrl *BlogController) PublicGetBlogPostBySlug(c *gin.Context) {
	var req struct {
		Slug string `json:"slug"`
	}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	post, err := ctrl.blogSvr.GetBlogPostBySlug(req.Slug)
	if err != nil || post == nil || post.Status != model.BlogStatusPublished {
		http.ErrorData(c, "未找到", nil)
		return
	}

	http.SuccessData(c, "获取成功", post)
}
