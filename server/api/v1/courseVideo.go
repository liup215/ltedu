package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"
	"time"

	"github.com/gin-gonic/gin"
)

var CourseVideoCtrl = &CourseVideoController{
	courseVideoSvr: service.CourseVideoSvr,
}

type CourseVideoController struct {
	courseVideoSvr *service.CourseVideoService
}

func (ctrl *CourseVideoController) CreateCourseVideo(c *gin.Context) {
	req := model.CourseVideoCreateEditRequest{}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败: "+err.Error(), nil)
		return
	}

	courseVideo := model.CourseVideo{
		CourseId:     req.CourseId,
		MediaVideoId: req.MediaVideoId,
		Name:         req.Name,
		Duration:     req.Duration,
		BanDrag:      req.BanDrag,
		IsShow:       req.IsShow,
		ChapterId:    req.ChapterId,
	}

	if t, err := time.Parse("2006-01-02 15:04:05", req.PublishedAt); err != nil {
		http.ErrorData(c, "创建失败"+err.Error(), nil)
		return
	} else {
		courseVideo.PublishedAt = &t
	}

	err := ctrl.courseVideoSvr.CreateCourseVideo(courseVideo)

	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "创建成功!", nil)
}

func (ctrl *CourseVideoController) EditCourseVideo(c *gin.Context) {
	req := model.CourseVideoCreateEditRequest{}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败: "+err.Error(), nil)
		return
	}

	courseVideo := model.CourseVideo{
		Model:        model.Model{ID: uint(req.ID)},
		CourseId:     req.CourseId,
		MediaVideoId: req.MediaVideoId,
		Name:         req.Name,
		Duration:     req.Duration,
		BanDrag:      req.BanDrag,
		IsShow:       req.IsShow,
		ChapterId:    req.ChapterId,
	}

	if t, err := time.Parse("2006-01-02 15:04:05", req.PublishedAt); err != nil {
		http.ErrorData(c, "创建失败"+err.Error(), nil)
		return
	} else {
		courseVideo.PublishedAt = &t
	}

	err := ctrl.courseVideoSvr.EditCourseVideo(courseVideo)

	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "编辑成功!", nil)
}

func (ctrl *CourseVideoController) DeleteCourseVideo(c *gin.Context) {
	courseVideo := model.CourseVideo{}

	if err := c.BindJSON(&courseVideo); err != nil {
		http.ErrorData(c, "参数解析失败: "+err.Error(), nil)
		return
	}

	// 用户校验
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	_, err = service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "用户未登录！", nil)
		return
	}

	err = ctrl.courseVideoSvr.DeleteCourseVideo(courseVideo.ID)

	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "删除成功!", nil)
}

func (ctrl *CourseVideoController) SelectCourseVideoById(c *gin.Context) {
	q := model.CourseVideoQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	o, err := ctrl.courseVideoSvr.SelectCourseVideoById(q.ID)
	if err != nil {
		http.ErrorData(c, "查询失败", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *CourseVideoController) SelectCourseVideoList(c *gin.Context) {
	q := model.CourseVideoQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	list, total, err := ctrl.courseVideoSvr.SelectCourseVideoList(q)
	if err != nil {
		http.ErrorData(c, "查询失败,"+err.Error(), nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}
