package service

import (
	"edu/lib/strings"
	"edu/model"
	"edu/repository"
	"errors"
)

var AuthSvr *AuthService = &AuthService{baseService: newBaseService()}

type AuthService struct {
	baseService
}

// 生成新的 tokenSalt
func (svr *AuthService) GenerateTokenSalt() string {
	return strings.Random(32)
}

// 校验原密码
func (svr *AuthService) VerifyPassword(userID uint, password string) bool {
	user, err := UserSvr.SelectUserById(userID)
	if err != nil || user == nil {
		return false
	}
	return user.Password == strings.Md5(password+user.Salt)
}

// UserAuthenticator authenticates a standard user.
// If successful, it returns the *model.User object.
// The calling code (e.g., login controller) will then check if this user is also an admin.
func (svr *AuthService) UserAuthenticator(username, password, captchaId, captchaValue, ip string) (*model.User, error) {
	// 验证码校验
	if !CaptchaSvr.VerifyCaptcha(captchaId, captchaValue) {
		return nil, errors.New("验证码错误")
	}

	// 登录过程验证
	user, err := repository.UserRepo.FindByUsername(username)
	if err != nil || user == nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	if user.Password != strings.Md5(password+user.Salt) {
		return nil, errors.New("用户名或密码错误")
	}

	if user.Username != "admin" && user.Status != model.UserStatusNormal {
		return nil, errors.New("用户账户状态异常或未激活")
	}

	return user, nil
}

func (svr *AuthService) UserRegister(r model.UserRegisterRequest) error {
	if !strings.CheckUsername(r.Username) {
		return errors.New("无效的用户名")
	}

	if !strings.CheckPassword(r.Password) {
		return errors.New("密码由6-16位字母或数字组成,必须同时包含大小写字母")
	}

	if r.Password != r.PasswordConfirm {
		return errors.New("两次输入密码不一致")
	}

	// 检查用户名是否已存在
	userByUsername, _ := repository.UserRepo.FindByUsername(r.Username)
	if userByUsername != nil {
		return errors.New("用户名已存在，请更换用户名后重新注册")
	}

	// 检查邮箱是否已存在
	userByEmail, _ := repository.UserRepo.FindByEmail(r.Email)
	if userByEmail != nil {
		return errors.New("邮箱已被注册，请更换邮箱后重新注册")
	}

	t := model.User{
		Username:  r.Username,
		Mobile:    r.Mobile,
		Email:     r.Email,
		Salt:      strings.Random(8),
		TokenSalt: svr.GenerateTokenSalt(),
		Status:    model.UserStatusNormal,
	}

	t.Password = strings.Md5(r.Password + t.Salt)
	t.IsActive = model.YES
	t.IsUsernameSet = model.YES
	t.IsPasswordSet = model.YES

	return repository.UserRepo.Create(&t)
}
