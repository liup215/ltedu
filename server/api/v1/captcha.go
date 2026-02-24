package v1

import (
	"edu/lib/net/http"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var CaptchaCtrl = &CaptchaController{
	CaptchaService: service.CaptchaSvr,
}

type CaptchaController struct {
	CaptchaService *service.CaptchaService
}

// @Summary      获取图形验证码
// @Description  生成一个新的图形验证码
// @Tags         认证
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Failure      400  {object}  map[string]interface{}  "生成失败"
// @Router       /v1/captcha [post]
func (ctrl *CaptchaController) GetImage(c *gin.Context) {
	captcha, err := ctrl.CaptchaService.GenerateCaptcha()
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "Success", captcha)
}
