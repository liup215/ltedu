package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"
	"time"

	"github.com/gin-gonic/gin"
)

// AnalyticsCtrl is the singleton analytics controller.
var AnalyticsCtrl = &AnalyticsController{
	analyticsSvr:     service.AnalyticsSvr,
	recommendSvr:     service.RecommendationSvr,
}

// AnalyticsController handles learning analytics and recommendation endpoints.
type AnalyticsController struct {
	analyticsSvr     *service.AnalyticsService
	recommendSvr     *service.RecommendationService
}

// GetClassSummary returns a high-level performance summary for a class.
// POST /api/v1/analytics/class/summary
func (ctrl *AnalyticsController) GetClassSummary(c *gin.Context) {
	var req model.ClassAnalyticsQuery
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	summary, err := ctrl.analyticsSvr.GetClassSummary(req.ClassID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Class summary fetched successfully", summary)
}

// GetStudentPerformanceList returns per-student analytics for a class.
// POST /api/v1/analytics/class/students
func (ctrl *AnalyticsController) GetStudentPerformanceList(c *gin.Context) {
	var req model.ClassAnalyticsQuery
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	list, err := ctrl.analyticsSvr.GetStudentPerformanceList(req.ClassID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	if list == nil {
		list = []model.StudentPerformanceSummary{}
	}

	http.SuccessData(c, "Student performance list fetched successfully", gin.H{
		"list":  list,
		"total": len(list),
	})
}

// GetClassHeatmap returns a chapter × student mastery heatmap for a class.
// POST /api/v1/analytics/class/heatmap
func (ctrl *AnalyticsController) GetClassHeatmap(c *gin.Context) {
	var req model.ClassAnalyticsQuery
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	heatmap, err := ctrl.analyticsSvr.GetClassHeatmap(req.ClassID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Class heatmap fetched successfully", heatmap)
}

// GetAttemptTrends returns a time-series of attempt counts for a class.
// POST /api/v1/analytics/class/trends
func (ctrl *AnalyticsController) GetAttemptTrends(c *gin.Context) {
	var req model.TrendQuery
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	// Default to last 30 days if no dates provided
	if req.EndDate.IsZero() {
		req.EndDate = time.Now()
	}
	if req.StartDate.IsZero() {
		req.StartDate = req.EndDate.AddDate(0, 0, -30)
	}

	trends, err := ctrl.analyticsSvr.GetAttemptTrends(req.ClassID, req.StartDate, req.EndDate)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	if trends == nil {
		trends = []model.AttemptTrendPoint{}
	}

	http.SuccessData(c, "Trends fetched successfully", trends)
}

// GetEarlyWarnings returns students flagged as at-risk in a class.
// POST /api/v1/analytics/class/earlyWarning
func (ctrl *AnalyticsController) GetEarlyWarnings(c *gin.Context) {
	var req model.ClassAnalyticsQuery
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	warnings, err := ctrl.analyticsSvr.GetEarlyWarnings(req.ClassID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	if warnings == nil {
		warnings = []model.EarlyWarningStudent{}
	}

	http.SuccessData(c, "Early warnings fetched successfully", gin.H{
		"list":  warnings,
		"total": len(warnings),
	})
}

// GetStudentAnalytics returns personal analytics for the current student's goal.
// POST /api/v1/analytics/student/summary
func (ctrl *AnalyticsController) GetStudentAnalytics(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req model.StudentAnalyticsQuery
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	summary, err := ctrl.analyticsSvr.GetStudentAnalytics(u.ID, req.GoalID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Student analytics fetched successfully", summary)
}

// GetRecommendations returns personalized chapter recommendations for the current student.
// POST /api/v1/analytics/recommend
func (ctrl *AnalyticsController) GetRecommendations(c *gin.Context) {
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}

	var req model.StudentAnalyticsQuery
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	recs, err := ctrl.recommendSvr.GetRecommendations(u.ID, req.GoalID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Recommendations fetched successfully", recs)
}
