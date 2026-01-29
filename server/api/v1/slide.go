package v1

import (
	"edu/lib/net/http"
	stringUtils "edu/lib/strings"
	"edu/model"
	"edu/service"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var SlideCtrl = &SlideController{
	slideSvr: service.SlideSvr,
}

type SlideController struct {
	slideSvr *service.SlideService
}

// Slide管理
func (ctrl *SlideController) SelectSlideList(c *gin.Context) {
	q := model.SlideQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	slides, total, err := ctrl.slideSvr.SelectSlideList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	list := []*model.SlideQueryResponse{}
	for _, v := range slides {
		list = append(list, v.GetResponse())
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *SlideController) SelectSlideById(c *gin.Context) {
	q := model.SlideQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.slideSvr.SelectSlideById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o.GetResponse())
}

func (ctrl *SlideController) SelectSlideAll(c *gin.Context) {
	oq := model.SlideQueryRequest{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	slides, err := ctrl.slideSvr.SelectSlideAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	list := []*model.SlideQueryResponse{}
	for _, v := range slides {
		list = append(list, v.GetResponse())
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

func (ctrl *SlideController) CreateSlide(c *gin.Context) {
	o := model.SlideCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败:"+err.Error(), nil)
		return
	}

	s := model.Slide{
		Name:        o.Name,
		Description: o.Description,
		SyllabusId:  o.SyllabusId,
	}

	// get random name
	now := time.Now().Nanosecond()
	md5hash := stringUtils.Md5(fmt.Sprintf("%d%d", now, o.SyllabusId))
	s.Hash = md5hash

	// savePathFormat := "slides/%s/"

	// savePath := fmt.Sprintf(savePathFormat, md5hash)
	// os.MkdirAll(filepath.Dir(savePath), os.ModePerm)

	// // utils.CopyFile("slide-template.html", savePath+"index.html")
	// // write html to file
	// html := `
	// <html>
	// <head>
	// 		<title>Online Slide - Alevel.ICU</title>
	// </head>
	// <body>
	// 		<h1>Online Slide - Alevel.ICU</h1>
	// 		<p>The slide is coming soon...</p>
	// </body>
	// </html>
	// `
	// os.WriteFile(savePath+"index.html", []byte(html), 0644)

	err := ctrl.slideSvr.CreateSlide(s)
	if err != nil {
		http.ErrorData(c, "添加失败:"+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *SlideController) EditSlide(c *gin.Context) {
	o := model.SlideCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.slideSvr.EditSlide(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *SlideController) DeleteSlide(c *gin.Context) {
	o := model.Slide{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.slideSvr.DeleteSlide(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *SlideController) ViewSlide(c *gin.Context) {

	q := model.SlideQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	slide, err := ctrl.slideSvr.SelectSlideById(q.ID)
	if err != nil {
		http.ErrorData(c, "课件查询失败！", nil)
		return
	}

	if slide.Hash == "" {
		http.ErrorData(c, "课件不存在！", nil)
		return
	}

	http.SuccessData(c, "获取下载地址成功！", gin.H{
		"hash": slide.Hash,
	})
}
