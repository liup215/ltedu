package v1

import (
	"edu/lib/net/http"
	"edu/model"
	"edu/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

var QualificationCtrl = &QualificationController{
	qualificationSvr: service.QualificationSvr,
}

type QualificationController struct {
	qualificationSvr *service.QualificationService
}

// Organisation管理
func (ctrl *QualificationController) SelectOrganisationList(c *gin.Context) {
	q := model.OrganisationQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.qualificationSvr.SelectOrganisationList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *QualificationController) SelectOrganisationById(c *gin.Context) {
	q := model.OrganisationQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.qualificationSvr.SelectOrganisationById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *QualificationController) SelectOrganisationAll(c *gin.Context) {
	oq := model.OrganisationQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.qualificationSvr.SelectOrganisationAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

func (ctrl *QualificationController) CreateOrganisation(c *gin.Context) {
	o := model.Organisation{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	r, err := ctrl.qualificationSvr.CreateOrganisation(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", r)
}

func (ctrl *QualificationController) EditOrganisation(c *gin.Context) {
	o := model.Organisation{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	r, err := ctrl.qualificationSvr.EditOrganisation(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", r)
}

func (ctrl *QualificationController) DeleteOrganisation(c *gin.Context) {
	o := model.Organisation{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.qualificationSvr.DeleteOrganisation(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// Qualification管理
func (ctrl *QualificationController) SelectQualificationList(c *gin.Context) {
	q := model.QualificationQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.qualificationSvr.SelectQualificationList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *QualificationController) SelectQualificationById(c *gin.Context) {
	q := model.QualificationQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.qualificationSvr.SelectQualificationById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *QualificationController) SelectQualificationAll(c *gin.Context) {
	oq := model.QualificationQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.qualificationSvr.SelectQualificationAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

func (ctrl *QualificationController) CreateQualification(c *gin.Context) {
	o := model.Qualification{}
	if err := c.BindJSON(&o); err != nil {
		fmt.Println(err.Error())
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	r, err := ctrl.qualificationSvr.CreateQualification(o)
	if err != nil {
		http.ErrorData(c, "创建失败"+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", r)
}

func (ctrl *QualificationController) EditQualification(c *gin.Context) {
	o := model.Qualification{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败"+err.Error(), nil)
		return
	}
	r, err := ctrl.qualificationSvr.EditQualification(o)
	if err != nil {
		http.ErrorData(c, "编辑失败"+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", r)
}

func (ctrl *QualificationController) DeleteQualification(c *gin.Context) {
	o := model.Qualification{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.qualificationSvr.DeleteQualification(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// Subject管理
// func (ctrl *QualificationController) SelectSubjectList(c *gin.Context) {
// 	q := model.SubjectQuery{}
// 	if err := c.BindJSON(&q); err != nil {
// 		http.ErrorData(c, "参数解析失败", nil)
// 		return
// 	}
// 	list, total, err := ctrl.qualificationSvr.SelectSubjectList(q)
// 	if err != nil {
// 		http.ErrorData(c, "参数解析失败", nil)
// 		return
// 	}
// 	http.SuccessData(c, "数据获取成功!", gin.H{
// 		"list":  list,
// 		"total": total,
// 	})
// }

// func (ctrl *QualificationController) SelectSubjectById(c *gin.Context) {
// 	q := model.SubjectQuery{}
// 	if err := c.BindJSON(&q); err != nil {
// 		http.ErrorData(c, "参数解析失败", nil)
// 		return
// 	}
// 	o, err := ctrl.qualificationSvr.SelectSubjectById(q.ID)
// 	if err != nil {
// 		http.ErrorData(c, "参数解析失败", nil)
// 		return
// 	}
// 	http.SuccessData(c, "数据获取成功!", o)
// }

// func (ctrl *QualificationController) SelectSubjectAll(c *gin.Context) {
// 	sq := model.SubjectQuery{}
// 	if err := c.BindJSON(&sq); err != nil {
// 		http.ErrorData(c, "数据获取失败", nil)
// 		return
// 	}
// 	list, err := ctrl.qualificationSvr.SelectSubjectAll(sq)
// 	if err != nil {
// 		http.ErrorData(c, "数据获取失败!", nil)
// 		return
// 	}

// 	http.SuccessData(c, "数据获取成功!", gin.H{
// 		"list":  list,
// 		"total": len(list),
// 	})
// }

// func (ctrl *QualificationController) CreateSubject(c *gin.Context) {
// 	o := model.Subject{}
// 	if err := c.BindJSON(&o); err != nil {
// 		http.ErrorData(c, "参数解析失败", nil)
// 		return
// 	}
// 	r, err := ctrl.qualificationSvr.CreateSubject(o)
// 	if err != nil {
// 		http.ErrorData(c, "参数解析失败", nil)
// 		return
// 	}
// 	http.SuccessData(c, "数据获取成功!", r)
// }

// func (ctrl *QualificationController) EditSubject(c *gin.Context) {
// 	o := model.Subject{}
// 	if err := c.BindJSON(&o); err != nil {
// 		http.ErrorData(c, "参数解析失败", nil)
// 		return
// 	}
// 	r, err := ctrl.qualificationSvr.EditSubject(o)
// 	if err != nil {
// 		http.ErrorData(c, err.Error(), nil)
// 		return
// 	}
// 	http.SuccessData(c, "数据获取成功!", r)
// }

// func (ctrl *QualificationController) DeleteSubject(c *gin.Context) {
// 	o := model.Subject{}
// 	if err := c.BindJSON(&o); err != nil {
// 		http.ErrorData(c, "参数解析失败", nil)
// 		return
// 	}
// 	err := ctrl.qualificationSvr.DeleteSubject(o.ID)
// 	if err != nil {
// 		http.ErrorData(c, "参数解析失败", nil)
// 		return
// 	}
// 	http.SuccessData(c, "数据获取成功!", nil)
// }

// Syllabus管理
func (ctrl *QualificationController) SelectSyllabusList(c *gin.Context) {
	q := model.SyllabusQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.qualificationSvr.SelectSyllabusList(q)
	if err != nil {
		http.ErrorData(c, "数据获取失败:"+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *QualificationController) SelectSyllabusById(c *gin.Context) {
	q := model.SyllabusQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.qualificationSvr.SelectSyllabusById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *QualificationController) SelectSyllabusAll(c *gin.Context) {
	oq := model.SyllabusQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "参数解析失败: "+err.Error(), nil)
		return
	}
	list, err := ctrl.qualificationSvr.SelectSyllabusAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败: "+err.Error(), nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
		"query": oq,
	})
}

func (ctrl *QualificationController) CreateSyllabus(c *gin.Context) {
	o := model.Syllabus{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	r, err := ctrl.qualificationSvr.CreateSyllabus(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", r)
}

func (ctrl *QualificationController) EditSyllabus(c *gin.Context) {
	o := model.Syllabus{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	r, err := ctrl.qualificationSvr.EditSyllabus(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", r)
}

func (ctrl *QualificationController) DeleteSyllabus(c *gin.Context) {
	o := model.Syllabus{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.qualificationSvr.DeleteSyllabus(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// 章节管理
func (ctrl *QualificationController) GetChapterTree(c *gin.Context) {
	s := model.ChapterQuery{}
	if err := c.BindJSON(&s); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	tree := ctrl.qualificationSvr.GetChapterTree(s.SyllabusId)
	http.SuccessData(c, "获取成功！", tree)
}

func (ctrl *QualificationController) SelectChapterById(c *gin.Context) {
	q := model.ChapterQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.qualificationSvr.SelectChapterById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *QualificationController) CreateChapter(c *gin.Context) {
	o := model.Chapter{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	r, err := ctrl.qualificationSvr.CreateChapter(o)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", r)
}

func (ctrl *QualificationController) EditChapter(c *gin.Context) {
	o := model.Chapter{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	r, err := ctrl.qualificationSvr.EditChapter(o)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", r)
}

func (ctrl *QualificationController) DeleteChapter(c *gin.Context) {
	o := model.Chapter{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.qualificationSvr.DeleteChapter(o.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}
