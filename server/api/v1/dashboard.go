package v1

import (
	"edu/lib/net/http"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var DashboardCtrl = &DashboardController{
	dashboardSvr: service.DashboardSvr,
}

type DashboardController struct {
	dashboardSvr *service.DashboardService
}

// @Summary      获取仪表盘数据
// @Description  获取管理后台仪表盘统计数据
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Security     BearerAuth
// @Router       /v1/dashboard [get]
func (dc *DashboardController) Index(c *gin.Context) {

	http.SuccessData(c, "数据获取成功！", dc.dashboardSvr.AdminDashboard())
}
