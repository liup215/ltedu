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

func (ac *AddonController) Index(c *gin.Context) {
	http.SuccessData(c, "暂无插件！", []string{})
}
