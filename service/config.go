package service

import (
	"edu/model"
	"edu/repository"
	"encoding/json"
	"errors"
)

func init() {
	ConfigSvr = &ConfigService{baseService: newBaseService()}
}

var ConfigSvr *ConfigService

type ConfigService struct {
	baseService
}

// 图片上传设置
// GetImageUploadConfigRaw 获取图片上传配置（包含明文密钥），仅供内部使用
func (svr *ConfigService) GetImageUploadConfigRaw() (icfg model.ImageUploadConfig, err error) {
	cfg, err := repository.AppConfigRepo.FirstOrCreateByKey(model.LTEDU_CONFIG_IMAGE_UPLOAD_KEY)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(cfg.Value), &icfg)
	if err != nil {
		return
	}
	return
}

// GetImageUploadConfig 获取图片上传配置（密钥已脱敏），供前端展示使用
func (svr *ConfigService) GetImageUploadConfig() (icfg model.ImageUploadConfig, err error) {
	icfg, err = svr.GetImageUploadConfigRaw()
	if err != nil {
		return
	}

	if icfg.CosSecretKey != "" {
		icfg.CosSecretKey = "**********************"
	}
	if icfg.OssAccessKeySecret != "" {
		icfg.OssAccessKeySecret = "**********************"
	}
	if icfg.QiniuSecretKey != "" {
		icfg.QiniuSecretKey = "**********************"
	}

	return
}

func (svr *ConfigService) SaveImageUploadConfig(c model.ImageUploadConfig) error {
	// 本地存储
	if c.Disk == "" {
		return errors.New("存储类型不能为空")
	}

	// 获取旧配置（使用 Raw 方法获取真实密钥）
	oldCfg, _ := svr.GetImageUploadConfigRaw()

	// 如果密钥是掩码，则使用旧配置的密钥
	if c.CosSecretKey == "**********************" {
		c.CosSecretKey = oldCfg.CosSecretKey
	}
	if c.OssAccessKeySecret == "**********************" {
		c.OssAccessKeySecret = oldCfg.OssAccessKeySecret
	}
	if c.QiniuSecretKey == "**********************" {
		c.QiniuSecretKey = oldCfg.QiniuSecretKey
	}

	b, err := json.Marshal(&c)
	if err != nil {
		return err
	}

	cfg, _ := repository.AppConfigRepo.FirstOrCreateByKey(model.LTEDU_CONFIG_IMAGE_UPLOAD_KEY)

	cfg.Value = string(b)
	return repository.AppConfigRepo.Update(cfg)
}

// 视频上传设置
// GetVideoUploadConfigRaw 获取视频上传配置（包含明文密钥），仅供内部使用
func (svr *ConfigService) GetVideoUploadConfigRaw() (icfg model.VideoUploadConfig, err error) {
	cfg, err := repository.AppConfigRepo.FirstOrCreateByKey(model.LTEDU_CONFIG_VIDEO_UPLOAD_KEY)
	if err != nil {
		return icfg, err
	}

	err = json.Unmarshal([]byte(cfg.Value), &icfg)
	if err != nil {
		return
	}
	return
}

// GetVideoUploadConfig 获取视频上传配置（密钥已脱敏），供前端展示使用
func (svr *ConfigService) GetVideoUploadConfig() (icfg model.VideoUploadConfig, err error) {
	icfg, err = svr.GetVideoUploadConfigRaw()
	if err != nil {
		return
	}

	if icfg.CosSecretKey != "" {
		icfg.CosSecretKey = "**********************"
	}
	if icfg.OssAccessKeySecret != "" {
		icfg.OssAccessKeySecret = "**********************"
	}
	if icfg.QiniuSecretKey != "" {
		icfg.QiniuSecretKey = "**********************"
	}
	return icfg, nil
}

func (svr *ConfigService) SaveVideoUploadConfig(c model.VideoUploadConfig) error {
	// 本地存储
	if c.Disk == "" {
		return errors.New("存储类型不能为空")
	}

	// 获取旧配置（使用 Raw 方法获取真实密钥）
	oldCfg, _ := svr.GetVideoUploadConfigRaw()

	// 如果密钥是掩码，则使用旧配置的密钥
	if c.CosSecretKey == "**********************" {
		c.CosSecretKey = oldCfg.CosSecretKey
	}
	if c.OssAccessKeySecret == "**********************" {
		c.OssAccessKeySecret = oldCfg.OssAccessKeySecret
	}
	if c.QiniuSecretKey == "**********************" {
		c.QiniuSecretKey = oldCfg.QiniuSecretKey
	}

	b, err := json.Marshal(&c)
	if err != nil {
		return err
	}

	cfg, _ := repository.AppConfigRepo.FirstOrCreateByKey(model.LTEDU_CONFIG_VIDEO_UPLOAD_KEY)
	cfg.Value = string(b)
	return repository.AppConfigRepo.Update(cfg)
}

// 站点设置
func (svr *ConfigService) GetWebSiteConfig() (model.WebSiteConfig, error) {
	wsCfg := model.WebSiteConfig{}
	cfg, err := repository.AppConfigRepo.FirstOrCreateByKey(model.LTEDU_CONFIG_WEB_SITE_KEY)
	if err != nil {
		return wsCfg, err
	}
	json.Unmarshal([]byte(cfg.Value), &wsCfg)

	return wsCfg, nil
}

func (svr *ConfigService) SaveWebSiteConfig(c model.WebSiteConfig) error {

	b, err := json.Marshal(&c)
	if err != nil {
		return err
	}

	cfg, _ := repository.AppConfigRepo.FirstOrCreateByKey(model.LTEDU_CONFIG_WEB_SITE_KEY)
	cfg.Value = string(b)
	return repository.AppConfigRepo.Update(cfg)
}
