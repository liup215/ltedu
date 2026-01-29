package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

// ExamPaper 管理
func (ctrl *ExamPaperController) SelectExamPaperList(c *gin.Context) {
	q := model.ExamPaperQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "Failed to parse parameters", nil)
		return
	}

	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Unable to get current user info", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "User not logged in", nil)
		return
	}

	if !user.IsAdmin && !user.IsTeacher {
		http.ErrorData(c, "No permission to access exam paper list", nil)
		return
	}

	// 如果不是管理员，只能查看自己创建的试卷
	if !user.IsAdmin {
		q.UserId = user.ID
	}

	list, total, err := ctrl.questionPaperSvr.SelectExamPaperList(q)
	if err != nil {
		http.ErrorData(c, "Failed to retrieve data", nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *ExamPaperController) SelectExamPaperById(c *gin.Context) {
	q := model.ExamPaperQuery{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "Failed to parse parameters", nil)
		return
	}
	o, err := ctrl.questionPaperSvr.SelectExamPaperById(q.ID)
	if err != nil {
		http.ErrorData(c, "Failed to parse parameters", nil)
		return
	}
	http.SuccessData(c, "Data retrieved successfully!", o)
}

func (ctrl *ExamPaperController) SelectExamPaperAll(c *gin.Context) {
	oq := model.ExamPaperQuery{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "Failed to retrieve data!", nil)
		return
	}
	list, err := ctrl.questionPaperSvr.SelectExamPaperAll(oq)
	if err != nil {
		http.ErrorData(c, "Failed to retrieve data!", nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

func (ctrl *ExamPaperController) CreateExamPaper(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Unable to get current user info", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "User not logged in", nil)
		return
	}

	if user.ID == 0 {
		http.ErrorData(c, "Invalid user ID", nil)
		return
	}

	if !user.IsTeacher {
		http.ErrorData(c, "No permission to operate!", nil)
	}

	o := model.ExamPaper{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "Failed to parse parameters: "+err.Error(), nil)
		return
	}

	o.UserId = user.ID

	_, err = ctrl.questionPaperSvr.CreateExamPaper(o)
	if err != nil {
		http.ErrorData(c, "Failed to add", nil)
		return
	}
	http.SuccessData(c, "Data retrieved successfully!", nil)
}

func (ctrl *ExamPaperController) EditExamPaper(c *gin.Context) {
	o := model.ExamPaper{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "Failed to parse parameters", nil)
		return
	}

	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Unable to get current user info", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "User not logged in", nil)
		return
	}

	if !user.IsAdmin && !user.IsTeacher {
		http.ErrorData(c, "No permission to access exam paper list", nil)
		return
	}

	// 获取原试卷数据
	existing, err := ctrl.questionPaperSvr.SelectExamPaperById(o.ID)
	if err != nil {
		http.ErrorData(c, "Exam paper does not exist", nil)
		return
	}

	// 如果不是管理员，只能修改自己的试卷
	if !user.IsAdmin && existing.UserId != user.ID {
		http.ErrorData(c, "No permission to edit this exam paper", nil)
		return
	}

	err = ctrl.questionPaperSvr.EditExamPaper(o)
	if err != nil {
		http.ErrorData(c, "Failed to edit", nil)
		return
	}
	http.SuccessData(c, "Edited successfully!", nil)
}

func (ctrl *ExamPaperController) DeleteExamPaper(c *gin.Context) {
	o := model.ExamPaper{}
	if err := c.BindJSON(&o); err != nil {
		http.ErrorData(c, "Failed to parse parameters", nil)
		return
	}

	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Unable to get current user info", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		http.ErrorData(c, "User not logged in", nil)
		return
	}

	if !user.IsAdmin && !user.IsTeacher {
		http.ErrorData(c, "No permission to access exam paper list", nil)
		return
	}

	// 获取原试卷数据
	existing, err := ctrl.questionPaperSvr.SelectExamPaperById(o.ID)
	if err != nil {
		http.ErrorData(c, "Exam paper does not exist", nil)
		return
	}

	// 如果不是管理员，只能删除自己的试卷
	if !user.IsAdmin && existing.UserId != user.ID {
		http.ErrorData(c, "No permission to delete this exam paper", nil)
		return
	}

	err = ctrl.questionPaperSvr.DeleteExamPaper(o.ID)
	if err != nil {
		http.ErrorData(c, "Failed to delete", nil)
		return
	}
	http.SuccessData(c, "Deleted successfully!", nil)
}

func (ctrl *ExamPaperController) UploadExamPaper(c *gin.Context) {
	// file, err := c.FormFile("file")
	http.SuccessData(c, "Uploaded successfully!", nil)
}

func (ctrl *ExamPaperController) UpdateExamPaperQuestion(c *gin.Context) {
	q := model.ExamPaper{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "Failed to parse parameters!", nil)
		return
	}
	err := ctrl.questionPaperSvr.UpdateExamPaperQuestion(q)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Data retrieved successfully!", nil)
}
