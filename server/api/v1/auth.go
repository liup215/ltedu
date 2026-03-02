package v1

import (
"edu/lib/net/http"
"edu/lib/net/http/middleware/auth"
"edu/model"
"edu/service"
"encoding/json"
"errors"
"fmt"
"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func init() {
	AuthCtrl = &AuthController{
		authSvr: service.AuthSvr,
	}
}

var AuthCtrl *AuthController

type AuthController struct {
	authSvr *service.AuthService
}

// JWT Payload 方法：返回完整 User 数据
func (lc *AuthController) PayloadFunc(data interface{}) jwt.MapClaims {
	js, err := json.Marshal(data)
	if err != nil {
		return jwt.MapClaims{}
	}

	claims := jwt.MapClaims{}
	if json.Unmarshal(js, &claims) != nil {
		return jwt.MapClaims{}
	}

	fmt.Println("PayloadFunc claims:", claims)

	return claims
}

func (lc *AuthController) Authorizator(data interface{}, c *gin.Context) bool {
	js, err := json.Marshal(jwt.ExtractClaims(c))
	if err != nil {
		return false
	}

	authUser := &auth.User{}
	if json.Unmarshal(js, authUser) != nil {
		return false
	}

	if authUser.ID == 0 {
		return false
	}

	record, err := service.UserSvr.SelectUserById(authUser.ID)
	if err != nil {
		return false
	}

	if record.TokenSalt != authUser.TokenSalt {
		return false
	}

	return true
}

type Login struct {
	Username     string `json:"username" form:"username" binding:"required"`
	Password     string `json:"password" form:"password" binding:"required"`
	CaptchaId    string `json:"captchaId" form:"captchaId" binding:"required"`
	CaptchaValue string `json:"captchaValue" form:"captchaValue" binding:"required"`
}

type RegistrationRequest struct {
	Username         string `json:"username" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	Mobile           string `json:"mobile"` // Optional
	Password         string `json:"password" binding:"required,min=6"`
	PasswordConfirm  string `json:"passwordConfirm" binding:"required,eqfield=Password"`
	VerificationCode string `json:"verificationCode" binding:"required"`
}

// UserResponse defines the user data returned upon successful registration or profile fetch
// This should align with what the frontend UserProfileData expects
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"isAdmin"`
}

func (lc *AuthController) Authenticator(c *gin.Context) (interface{}, error) {
	var loginReq Login

	if err := c.Bind(&loginReq); err != nil {
		return nil, errors.New("登录参数错误")
	}

	// Step 1: Authenticate as a general user. UserAuthenticator now returns *model.User
	authUser, err := lc.authSvr.UserAuthenticator(loginReq.Username, loginReq.Password, loginReq.CaptchaId, loginReq.CaptchaValue, c.ClientIP())
	if err != nil {
		return nil, err // UserAuthenticator already returns appropriate error messages
	}
	// authUser is now *model.User, which includes IsAdmin (computed from roles), etc.

	// Update User login stats
	err = service.UserSvr.UpdateUserLoginStats(authUser.ID, c.ClientIP())
	if err != nil {
		// Log this error but don't necessarily fail the login, as authentication was successful
		// log.Printf("Failed to update user login stats for UserID %d: %v", authUser.ID, err)
		// Depending on policy, you might choose to return an error here.
		// For now, we proceed with login.
	}

	// If all checks pass, return the authenticated *model.User.
	// The auth middleware's IdentityHandler will need to extract User.ID from this.

	return &auth.User{ID: authUser.ID, TokenSalt: authUser.TokenSalt}, nil
}

func (lc *AuthController) LoginResponse(c *gin.Context, code int, message string, time time.Time) {

	http.SuccessData(c, "登录成功！", gin.H{
		"token":  message,
		"expire": time,
	})
}

// @Summary      修改密码
// @Description  修改当前用户密码（需要原密码验证，成功后强制重新登录）
// @Tags         认证
// @Accept       json
// @Produce      json
// @Param        body  body  map[string]interface{}  true  "旧密码与新密码"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/change-password [post]
func (lc *AuthController) ChangePassword(c *gin.Context) {
	type ChangePasswordRequest struct {
		OldPassword string `json:"oldPassword" binding:"required"`
		NewPassword string `json:"newPassword" binding:"required,min=6"`
	}
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "参数错误", err.Error())
		return
	}

	// 获取当前用户
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "无法获取当前用户信息", err.Error())
		return
	}

	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || user.ID == 0 {
		http.ErrorData(c, "用户未登录", nil)
		return
	}

	// 校验原密码
	if !service.AuthSvr.VerifyPassword(user.ID, req.OldPassword) {
		http.ErrorData(c, "原密码错误", nil)
		return
	}

	// 生成新 tokenSalt
	newSalt := service.AuthSvr.GenerateTokenSalt()

	// 更新密码和 tokenSalt
	err = service.UserSvr.UpdatePasswordAndSalt(user.ID, req.NewPassword, newSalt)
	if err != nil {
		http.ErrorData(c, "密码修改失败", err.Error())
		return
	}

	http.SuccessData(c, "密码修改成功，已强制重新登录", nil)
}

// @Summary      用户注册
// @Description  注册新用户账号（需要邮箱验证码）
// @Tags         认证
// @Accept       json
// @Produce      json
// @Param        body  body  RegistrationRequest  true  "注册信息"
// @Success      200   {object}  map[string]interface{}  "注册成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/register [post]
func (lc *AuthController) Register(c *gin.Context) {
	var req RegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ErrorData(c, "参数错误", err.Error())
		return
	}

	// Validate email verification code
	if !service.VerificationSvr.VerifyCode(req.Email, req.VerificationCode) {
		http.ErrorData(c, "邮箱验证码错误或已过期", nil)
		return
	}

	err := lc.authSvr.UserRegister(model.UserRegisterRequest{
		Username:        req.Username,
		Email:           req.Email,
		Mobile:          req.Mobile,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		http.ErrorData(c, "注册失败", err.Error())
		return
	}

	http.SuccessData(c, "注册成功", nil)
}
