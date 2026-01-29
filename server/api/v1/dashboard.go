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

func (dc *DashboardController) Index(c *gin.Context) {

	http.SuccessData(c, "数据获取成功！", dc.dashboardSvr.AdminDashboard())
}
