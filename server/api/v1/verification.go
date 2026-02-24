package v1

import (
	"edu/lib/net/http"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var VerificationCtrl = &VerificationController{
	VerificationService: service.VerificationSvr,
	UserService:         service.UserSvr,
}

type VerificationController struct {
	VerificationService *service.VerificationService
	UserService         *service.UserService
}

type SendCodeRequest struct {
	Email      string `json:"email" binding:"required,email"`
	CaptchaId  string `json:"captchaId" binding:"required"`
	CaptchaVal string `json:"captchaValue" binding:"required"`
}

// SendRegistrationCode sends a verification code to the user's email.
// @Summary      发送邮箱验证码
// @Description  向指定邮箱发送验证码（需要图形验证码校验）
// @Tags         认证
// @Accept       json
// @Produce      json
// @Param        body  body  SendCodeRequest  true  "邮箱和图形验证码"
// @Success      200   {object}  map[string]interface{}  "发送成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/verification/send-code [post]
func (ctrl *VerificationController) SendCode(c *gin.Context) {
	var req SendCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "Invalid request parameters", nil)
		return
	}

	// Validate captcha
	if !service.CaptchaSvr.VerifyCaptcha(req.CaptchaId, req.CaptchaVal) {
		http.ErrorData(c, "Invalid or expired captcha", nil)
		return
	}

	r := model.VerificationGenerateRequest{}
	if req.Email != "" {
		r.Type = model.VERIFICATION_TYPE_EMAIL
		r.Target = req.Email
	}
	err := ctrl.VerificationService.GenerateAndSendCode(r)
	if err != nil {
		http.ErrorData(c, "Failed to send verification code: "+err.Error(), nil)
		return
	}

	http.SuccessData(c, "Verification code sent successfully", nil)
}
