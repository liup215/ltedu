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
// @Summary      获取机构列表
// @Description  分页查询机构列表
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.OrganisationQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/organisation/list [post]
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

// @Summary      根据ID获取机构
// @Description  根据机构ID获取机构详情
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.OrganisationQuery  true  "机构ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/organisation/byId [post]
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

// @Summary      获取所有机构
// @Description  获取全部机构列表（不分页）
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.OrganisationQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/organisation/all [post]
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

// @Summary      创建机构
// @Description  创建新机构
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Organisation  true  "机构信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/organisation/create [post]
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

// @Summary      编辑机构
// @Description  修改机构信息
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Organisation  true  "机构信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/organisation/edit [post]
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

// @Summary      删除机构
// @Description  删除指定机构
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Organisation  true  "机构ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/organisation/delete [post]
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
// @Summary      获取资质列表
// @Description  分页查询资质（考试）列表
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.QualificationQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/qualification/list [post]
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

// @Summary      根据ID获取资质
// @Description  根据资质ID获取详情
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.QualificationQuery  true  "资质ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/qualification/byId [post]
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

// @Summary      获取所有资质
// @Description  获取全部资质列表（不分页）
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.QualificationQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/qualification/all [post]
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

// @Summary      创建资质
// @Description  创建新资质（考试）
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Qualification  true  "资质信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/qualification/create [post]
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

// @Summary      编辑资质
// @Description  修改资质信息
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Qualification  true  "资质信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/qualification/edit [post]
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

// @Summary      删除资质
// @Description  删除指定资质
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Qualification  true  "资质ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/qualification/delete [post]
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
// @Summary      获取考纲列表
// @Description  分页查询考纲列表
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.SyllabusQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/syllabus/list [post]
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

// @Summary      根据ID获取考纲
// @Description  根据考纲ID获取考纲详情
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.SyllabusQuery  true  "考纲ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/syllabus/byId [post]
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

// @Summary      获取所有考纲
// @Description  获取全部考纲列表（不分页）
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.SyllabusQuery  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/syllabus/all [post]
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

// @Summary      创建考纲
// @Description  创建新考纲
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Syllabus  true  "考纲信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/syllabus/create [post]
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

// @Summary      编辑考纲
// @Description  修改考纲信息
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Syllabus  true  "考纲信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/syllabus/edit [post]
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

// @Summary      删除考纲
// @Description  删除指定考纲
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Syllabus  true  "考纲ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/syllabus/delete [post]
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
// @Summary      获取章节树
// @Description  根据考纲ID获取章节树形结构
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ChapterQuery  true  "考纲ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/chapter/tree [post]
func (ctrl *QualificationController) GetChapterTree(c *gin.Context) {
	s := model.ChapterQuery{}
	if err := c.BindJSON(&s); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	tree := ctrl.qualificationSvr.GetChapterTree(s.SyllabusId)
	http.SuccessData(c, "获取成功！", tree)
}

// @Summary      根据ID获取章节
// @Description  根据章节ID获取章节详情
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.ChapterQuery  true  "章节ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/chapter/byId [post]
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

// @Summary      创建章节
// @Description  创建新章节
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Chapter  true  "章节信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/chapter/create [post]
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

// @Summary      编辑章节
// @Description  修改章节信息
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Chapter  true  "章节信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/chapter/edit [post]
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

// @Summary      删除章节
// @Description  删除指定章节
// @Tags         考纲管理
// @Accept       json
// @Produce      json
// @Param        body  body  model.Chapter  true  "章节ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/chapter/delete [post]
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
