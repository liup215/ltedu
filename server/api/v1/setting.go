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

// @Summary      获取图片上传配置
// @Description  获取当前的图片上传存储配置
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Security     BearerAuth
// @Router       /v1/syssetting/imageUpload [get]
func (svr *SettingController) GetImageUploadConfig(c *gin.Context) {
	cf, err := svr.configSvr.GetImageUploadConfig()
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", cf)
}

// @Summary      保存图片上传配置
// @Description  保存图片上传存储配置
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Param        body  body  model.ImageUploadConfig  true  "图片上传配置"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/syssetting/imageUpload [post]
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

// @Summary      获取视频上传配置
// @Description  获取当前的视频上传存储配置
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Security     BearerAuth
// @Router       /v1/syssetting/videoUpload [get]
func (svr *SettingController) GetVideoUploadConfig(c *gin.Context) {
	cf, err := svr.configSvr.GetVideoUploadConfig()
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", cf)
}

// @Summary      保存视频上传配置
// @Description  保存视频上传存储配置
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Param        body  body  model.VideoUploadConfig  true  "视频上传配置"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/syssetting/videoUpload [post]
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

// @Summary      获取网站配置
// @Description  获取当前的网站基本配置
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Security     BearerAuth
// @Router       /v1/syssetting/webSite [get]
func (svr *SettingController) GetWebSiteConfig(c *gin.Context) {
	cf, err := svr.configSvr.GetWebSiteConfig()
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", cf)
}

// @Summary      保存网站配置
// @Description  保存网站基本配置
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Param        body  body  model.WebSiteConfig  true  "网站配置"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/syssetting/webSite [post]
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

// @Summary      迁移Base64图片
// @Description  将题目中的Base64图片迁移到存储服务（仅管理员）
// @Tags         系统设置
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Security     BearerAuth
// @Router       /v1/syssetting/image/migrate [post]
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
