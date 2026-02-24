package v1

import (
	"edu/lib/net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	AddonCtrl = &AddonController{}
}

var AddonCtrl *AddonController

type AddonController struct{}

// @Summary      获取插件列表
// @Description  返回已安装的插件列表
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Security     BearerAuth
// @Router       /v1/addons [get]
func (ac *AddonController) Index(c *gin.Context) {
	http.SuccessData(c, "暂无插件！", []string{})
}
