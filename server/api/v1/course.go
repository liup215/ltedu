package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"
	"time"

	"github.com/gin-gonic/gin"
)

var CourseCtrl = &CourseController{
	courseSvr: service.CourseSvr,
}

type CourseController struct {
	courseSvr *service.CourseService
}

func (ctrl *CourseController) CreateCourse(c *gin.Context) {
	req := model.CourseCreateEditRequest{}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败: "+err.Error(), nil)
		return
	}

	course := model.Course{
		Title:            req.Title,
		IsFree:           req.IsFree,
		Charge:           req.Charge,
		ShortDescription: req.ShortDescription,
		OriginalDesc:     req.OriginalDesc,
		SyllabusID:       req.SyllabusID,
	}

	if t, err := time.Parse("2006-01-02 15:04:05", req.PublishedAt); err != nil {
		http.ErrorData(c, "创建失败"+err.Error(), nil)
		return
	} else {
		course.PublishedAt = &t
	}

	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "用户未登录！", nil)
		return
	}

	course.UserType = model.USER_TYPE_ADMIN
	course.UserID = int(user.ID)

	err = ctrl.courseSvr.CreateCourse(course)

	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "创建成功!", nil)
}

func (ctrl *CourseController) EditCourse(c *gin.Context) {
	req := model.CourseCreateEditRequest{}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数解析失败: "+err.Error(), nil)
		return
	}

	course := model.Course{
		Model:             model.Model{ID: uint(req.ID)},
		Title:             req.Title,
		IsFree:            req.IsFree,
		Charge:            req.Charge,
		ShortDescription:  req.ShortDescription,
		OriginalDesc:      req.OriginalDesc,
		SyllabusID:        req.SyllabusID,
		PublishedAtString: req.PublishedAt,
		Thumb:             req.Thumb,
	}

	if t, err := time.Parse("2006-01-02 15:04:05", req.PublishedAt); err != nil {
		http.ErrorData(c, "创建失败"+err.Error(), nil)
		return
	} else {
		course.PublishedAt = &t
	}

	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "用户未登录！", nil)
		return
	}

	course.UserType = model.USER_TYPE_ADMIN
	course.UserID = int(user.ID)

	err = ctrl.courseSvr.EditCourse(course)

	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "编辑成功!", nil)
}

func (ctrl *CourseController) DeleteCourse(c *gin.Context) {
	course := model.Course{}

	if err := c.BindJSON(&course); err != nil {
		http.ErrorData(c, "参数解析失败: "+err.Error(), nil)
		return
	}

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

	err = ctrl.courseSvr.DeleteCourse(course.ID)

	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "创建成功!", nil)
}

func (ctrl *CourseController) SelectCourseById(c *gin.Context) {
	q := model.CourseQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	o, err := ctrl.courseSvr.SelectCourseById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *CourseController) SelectCourseList(c *gin.Context) {
	q := model.CourseQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	list, total, err := ctrl.courseSvr.SelectCourseList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败,"+err.Error(), nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}
