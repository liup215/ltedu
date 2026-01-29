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

// Image generates a new captcha image.
// @Summary Generate captcha
// @Description Generate a new captcha image
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {object} service.CaptchaImage
// @Router /captcha [post]
func (ctrl *CaptchaController) GetImage(c *gin.Context) {
	captcha, err := ctrl.CaptchaService.GenerateCaptcha()
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	http.SuccessData(c, "Success", captcha)
}
