package v1

import (
	"edu/lib/net/http"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var SchoolCtrl = &SchoolController{
	schoolSvr: service.SchoolSvr,
}

type SchoolController struct {
	schoolSvr *service.SchoolService
}

// Grade管理
func (ctrl *SchoolController) SelectGradeList(c *gin.Context) {
	q := model.GradeQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.schoolSvr.SelectGradeList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *SchoolController) SelectGradeById(c *gin.Context) {
	q := model.GradeQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.schoolSvr.SelectGradeById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *SchoolController) SelectGradeAll(c *gin.Context) {
	oq := model.GradeQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.schoolSvr.SelectGradeAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

func (ctrl *SchoolController) CreateGrade(c *gin.Context) {
	o := model.GradeCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败: "+err.Error(), nil)
		return
	}
	r, err := ctrl.schoolSvr.CreateGrade(o)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", r)
}

func (ctrl *SchoolController) EditGrade(c *gin.Context) {
	o := model.GradeCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败: "+err.Error(), nil)
		return
	}
	err := ctrl.schoolSvr.EditGrade(o)
	if err != nil {
		http.ErrorData(c, "更新失败："+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *SchoolController) DeleteGrade(c *gin.Context) {
	o := model.Grade{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.DeleteGrade(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// ClassType管理
func (ctrl *SchoolController) SelectClassTypeList(c *gin.Context) {
	q := model.ClassTypeQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.schoolSvr.SelectClassTypeList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *SchoolController) SelectClassTypeById(c *gin.Context) {
	q := model.ClassTypeQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.schoolSvr.SelectClassTypeById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *SchoolController) SelectClassTypeAll(c *gin.Context) {
	oq := model.ClassTypeQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.schoolSvr.SelectClassTypeAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

func (ctrl *SchoolController) CreateClassType(c *gin.Context) {
	o := model.ClassTypeCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.CreateClassType(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *SchoolController) EditClassType(c *gin.Context) {
	o := model.ClassTypeCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.EditClassType(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *SchoolController) DeleteClassType(c *gin.Context) {
	o := model.ClassType{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.DeleteClassType(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

// Class管理
func (ctrl *SchoolController) SelectClassList(c *gin.Context) {
	q := model.ClassQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, total, err := ctrl.schoolSvr.SelectClassList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *SchoolController) SelectClassById(c *gin.Context) {
	q := model.ClassQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	o, err := ctrl.schoolSvr.SelectClassById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *SchoolController) SelectClassAll(c *gin.Context) {
	oq := model.ClassQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.schoolSvr.SelectClassAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

func (ctrl *SchoolController) CreateClass(c *gin.Context) {
	o := model.ClassCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.CreateClass(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *SchoolController) EditClass(c *gin.Context) {
	o := model.ClassCreateEditRequest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.EditClass(o)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}

func (ctrl *SchoolController) DeleteClass(c *gin.Context) {
	o := model.Class{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.DeleteClass(o.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	http.SuccessData(c, "数据删除成功!", nil)
}

func (ctrl *SchoolController) GetStudentsByClassId(c *gin.Context) {
	o := model.Class{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	list, err := ctrl.schoolSvr.GetStudentsByClassId(o.ID)
	if err != nil {
		http.ErrorData(c, "数据获取失败:"+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", list)
}

type AddStudentToClassRepuest struct {
	ClassId   uint `json:"classId"`
	StudentId uint `json:"studentId"`
}

func (ctrl *SchoolController) AddStudentToClass(c *gin.Context) {
	o := AddStudentToClassRepuest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.AddStudentToClass(model.Class{Model: model.Model{ID: o.ClassId}}, model.User{Model: model.Model{ID: o.StudentId}})
	if err != nil {
		http.ErrorData(c, "数据获取失败:"+err.Error(), nil)
		return
	}
	http.SuccessData(c, "添加成功!", nil)
}

func (ctrl *SchoolController) DeleteStudentFromClass(c *gin.Context) {
	o := AddStudentToClassRepuest{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}
	err := ctrl.schoolSvr.DeleteStudentFromClass(model.Class{Model: model.Model{ID: o.ClassId}}, model.User{Model: model.Model{ID: o.StudentId}})
	if err != nil {
		http.ErrorData(c, "数据获取失败:"+err.Error(), nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", nil)
}
