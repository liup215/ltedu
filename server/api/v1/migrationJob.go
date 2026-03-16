package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var MigrationJobCtrl = &MigrationJobController{}

type MigrationJobController struct{}

// Create creates a new migration job and starts it asynchronously.
// @Summary      创建迁移任务
// @Description  创建考纲知识点迁移后台任务（仅管理员）
// @Tags         迁移任务
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /v1/migration-job/create [post]
func (ctrl *MigrationJobController) Create(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || !user.IsAdmin {
		http.ErrorData(c, "Only admin can create migration jobs", nil)
		return
	}

	var req struct {
		SyllabusId uint                   `json:"syllabusId" binding:"required"`
		Options    service.MigrateOptions `json:"options"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	job, err := service.MigrationJobSvr.CreateJob(u.ID, req.SyllabusId, req.Options)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Migration job created", job)
}

// GetByID gets a migration job by ID.
// @Summary      获取迁移任务
// @Description  根据ID获取迁移任务详情（仅管理员）
// @Tags         迁移任务
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /v1/migration-job/byId [post]
func (ctrl *MigrationJobController) GetByID(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || !user.IsAdmin {
		http.ErrorData(c, "Only admin can view migration jobs", nil)
		return
	}

	var req struct {
		ID uint `json:"id" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	job, err := service.MigrationJobSvr.GetJob(req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Success", job)
}

// List lists migration jobs with pagination.
// @Summary      获取迁移任务列表
// @Description  获取迁移任务列表（仅管理员）
// @Tags         迁移任务
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /v1/migration-job/list [post]
func (ctrl *MigrationJobController) List(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || !user.IsAdmin {
		http.ErrorData(c, "Only admin can list migration jobs", nil)
		return
	}

	var query model.MigrationJobQuery
	if err := c.BindJSON(&query); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	jobs, total, err := service.MigrationJobSvr.ListJobs(query)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Success", gin.H{
		"list":      jobs,
		"total":     total,
		"page":      query.PageIndex,
		"pageSize":  query.PageSize,
	})
}

// Retry retries a failed migration job.
// @Summary      重试迁移任务
// @Description  重新执行失败的迁移任务（仅管理员）
// @Tags         迁移任务
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /v1/migration-job/retry [post]
func (ctrl *MigrationJobController) Retry(c *gin.Context) {
	u, _ := auth.GetCurrentUser(c)

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || !user.IsAdmin {
		http.ErrorData(c, "Only admin can retry migration jobs", nil)
		return
	}

	var req struct {
		ID uint `json:"id" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "Parameter parsing failed", nil)
		return
	}

	job, err := service.MigrationJobSvr.RetryJob(req.ID)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	http.SuccessData(c, "Migration job queued for retry", job)
}
