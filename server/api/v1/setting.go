package v1

import (
	"edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"

	"github.com/gin-gonic/gin"
)

var SettingCtrl = &SettingController{configSvr: service.ConfigSvr}

type SettingController struct {
	configSvr *service.ConfigService
}

func (svr *SettingController) GetImageUploadConfig(c *gin.Context) {
	cf, err := svr.configSvr.GetImageUploadConfig()
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", cf)
}

func (svr *SettingController) SaveImageUploadConfig(c *gin.Context) {
	cf := model.ImageUploadConfig{}
	if err := c.BindJSON(&cf); err != nil {
		http.ErrorData(c, "设置失败!", nil)
		return
	}

	if err := svr.configSvr.SaveImageUploadConfig(cf); err != nil {
		http.ErrorData(c, "设置失败!", nil)
		return
	}

	http.SuccessData(c, "设置成功!", nil)
}

func (svr *SettingController) GetVideoUploadConfig(c *gin.Context) {
	cf, err := svr.configSvr.GetVideoUploadConfig()
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", cf)
}

func (svr *SettingController) SaveVideoUploadConfig(c *gin.Context) {
	cf := model.VideoUploadConfig{}
	if err := c.BindJSON(&cf); err != nil {
		http.ErrorData(c, "设置失败!", nil)
		return
	}

	if err := svr.configSvr.SaveVideoUploadConfig(cf); err != nil {
		http.ErrorData(c, "设置失败!", nil)
		return
	}

	http.SuccessData(c, "设置成功!", nil)
}

func (svr *SettingController) GetWebSiteConfig(c *gin.Context) {
	cf, err := svr.configSvr.GetWebSiteConfig()
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", cf)
}

func (svr *SettingController) SaveWebSiteConfig(c *gin.Context) {
	cf := model.WebSiteConfig{}
	if err := c.BindJSON(&cf); err != nil {
		http.ErrorData(c, "设置失败!", nil)
		return
	}

	if err := svr.configSvr.SaveWebSiteConfig(cf); err != nil {
		http.ErrorData(c, "设置失败!", nil)
		return
	}

	http.SuccessData(c, "设置成功!", nil)
}

func (svr *SettingController) MigrateBase64Images(c *gin.Context) {
	// 权限验证：仅管理员可调用
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		http.ErrorData(c, "Failed to get current user info", nil)
		return
	}
	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil || user == nil || !user.IsAdmin {
		http.ErrorData(c, "No permission, only admin can start migration", nil)
		return
	}

	if err := service.QuestionSvr.MigrateBase64Images(); err != nil {
		http.ErrorData(c, "迁移失败: "+err.Error(), nil)
		return
	}
	http.SuccessData(c, "迁移成功!", nil)
}
