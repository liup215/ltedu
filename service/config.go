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
func (svr *ConfigService) GetImageUploadConfig() (icfg model.ImageUploadConfig, err error) {
	cfg, err := repository.AppConfigRepo.FirstOrCreateByKey(model.LTEDU_CONFIG_IMAGE_UPLOAD_KEY)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(cfg.Value), &icfg)
	if err != nil {
		return
	}

	icfg.CosSecretKey = "**********************"
	icfg.OssAccessKeySecret = "**********************"
	icfg.QiniuSecretKey = "**********************"

	return
}

func (svr *ConfigService) SaveImageUploadConfig(c model.ImageUploadConfig) error {
	// 本地存储
	if c.Disk == "" {
		return errors.New("存储类型不能为空")
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
func (svr *ConfigService) GetVideoUploadConfig() (icfg model.VideoUploadConfig, err error) {
	// icfg := model.VideoUploadConfig{}

	cfg, err := repository.AppConfigRepo.FirstOrCreateByKey(model.LTEDU_CONFIG_VIDEO_UPLOAD_KEY)
	if err != nil {
		return icfg, err
	}

	err = json.Unmarshal([]byte(cfg.Value), &icfg)
	if err != nil {
		return
	}

	icfg.CosSecretKey = "**********************"
	icfg.OssAccessKeySecret = "**********************"
	icfg.QiniuSecretKey = "**********************"
	return icfg, nil
}

func (svr *ConfigService) SaveVideoUploadConfig(c model.VideoUploadConfig) error {
	// 本地存储
	if c.Disk == "" {
		return errors.New("存储类型不能为空")
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
