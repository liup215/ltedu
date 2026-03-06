package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var RecommendationCtrl = &RecommendationController{
	recommendationSvr: service.RecommendationSvr,
}

type RecommendationController struct {
	recommendationSvr *service.RecommendationService
}

// GetQuestionRecommendations returns personalized question recommendations for a student.
// GET /api/v1/recommendations/questions?studentId=<id>
// @Summary      获取个性化题目推荐
// @Description  根据学生的学习进度和知识薄弱点，推荐个性化练习题。若调用方为学生本人可省略 studentId；管理员或教师可通过 studentId 查询任意学生的推荐。
// @Tags         推荐系统
// @Produce      json
// @Param        studentId  query  int  false  "目标学生ID（管理员/教师专用）"
// @Success      200  {object}  map[string]interface{}  "成功"
// @Failure      400  {object}  map[string]interface{}  "参数错误"
// @Failure      403  {object}  map[string]interface{}  "无权限"
// @Security     BearerAuth
// @Router       /v1/recommendations/questions [get]
func (ctrl *RecommendationController) GetQuestionRecommendations(c *gin.Context) {
	currentUser, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	// Determine which student's recommendations to fetch.
	// By default use the current user; admins/teachers may pass studentId.
	targetStudentID := currentUser.ID
	if raw := c.Query("studentId"); raw != "" {
		parsed, err := strconv.ParseUint(raw, 10, 64)
		if err != nil || parsed == 0 {
			http.ErrorData(c, "Invalid studentId parameter", nil)
			return
		}
		requestedID := uint(parsed)
		if requestedID != currentUser.ID {
			// Only admins may request recommendations for other students.
			user, err := service.UserSvr.SelectUserById(currentUser.ID)
			if err != nil || user == nil || !user.IsAdmin {
				http.ForbiddenData(c, "You do not have permission to view another student's recommendations", nil)
				return
			}
			targetStudentID = requestedID
		}
	}

	resp, err := ctrl.recommendationSvr.GetQuestionRecommendations(targetStudentID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Recommendations retrieved successfully!", resp)
}
