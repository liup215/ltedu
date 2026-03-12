package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"
	"encoding/xml"
	"fmt"
	nethttp "net/http"
	"time"

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

// RSS feed types

type rssItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	GUID        string   `xml:"guid"`
}

type rssChannel struct {
	XMLName       xml.Name  `xml:"channel"`
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	Description   string    `xml:"description"`
	LastBuildDate string    `xml:"lastBuildDate"`
	Items         []rssItem `xml:"item"`
}

type rssFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Channel rssChannel `xml:"channel"`
}

// @Summary      Blog RSS feed
// @Description  Returns an RSS 2.0 feed of recent published blog posts
// @Tags         Blog
// @Produce      xml
// @Success      200  {string}  string  "RSS XML"
// @Router       /v1/blog/rss.xml [get]
func (ctrl *BlogController) GetBlogRSS(c *gin.Context) {
	req := model.BlogPostQueryRequest{
		Status: model.BlogStatusPublished,
	}
	req.PageIndex = 1
	req.PageSize = 20

	posts, _, err := ctrl.blogSvr.ListBlogPosts(req)
	if err != nil {
		c.Status(nethttp.StatusInternalServerError)
		return
	}

	scheme := "https"
	if c.Request.TLS == nil {
		scheme = "http"
	}
	baseURL := fmt.Sprintf("%s://%s", scheme, c.Request.Host)
	blogURL := baseURL + "/blog"

	items := make([]rssItem, 0, len(posts))
	for _, p := range posts {
		// Use PublishedAt if set, otherwise fall back to CreatedAt
		pubDate := p.CreatedAt.Format(time.RFC1123Z)
		if p.PublishedAt != nil && *p.PublishedAt != "" {
			if t, err := time.Parse(time.RFC3339, *p.PublishedAt); err == nil {
				pubDate = t.Format(time.RFC1123Z)
			}
		}
		link := fmt.Sprintf("%s/%s", blogURL, p.Slug)
		items = append(items, rssItem{
			Title:       p.Title,
			Link:        link,
			Description: p.Summary,
			PubDate:     pubDate,
			GUID:        link,
		})
	}

	feed := rssFeed{
		Version: "2.0",
		Channel: rssChannel{
			Title:         "Nerdlet Blog",
			Link:          blogURL,
			Description:   "Latest posts from Nerdlet – the smart, fun learning platform",
			LastBuildDate: time.Now().Format(time.RFC1123Z),
			Items:         items,
		},
	}

	output, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		c.Status(nethttp.StatusInternalServerError)
		return
	}

	c.Header("Content-Type", "application/rss+xml; charset=utf-8")
	c.String(nethttp.StatusOK, xml.Header+string(output))
}
