package storage

import (
	"context"
	"edu/model"
	"strings"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiniuKodo struct {
	config model.ImageUploadConfig
}

func NewQiniuKodo(config model.ImageUploadConfig) (*QiniuKodo, error) {
	return &QiniuKodo{config: config}, nil
}

func (s *QiniuKodo) Upload(objectName string, localPath string) (url string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: s.config.QiniuBucket,
	}
	mac := qbox.NewMac(s.config.QiniuAccessKey, s.config.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// Zone automatically determined
	cfg.UseHTTPS = true
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	err = formUploader.PutFile(context.Background(), &ret, upToken, objectName, localPath, &putExtra)
	if err != nil {
		return "", err
	}

	// Return URL
	if s.config.QiniuCDNUrl != "" {
		baseUrl := s.config.QiniuCDNUrl
		if !strings.HasSuffix(baseUrl, "/") {
			baseUrl += "/"
		}
		return baseUrl + ret.Key, nil
	}

	// If no CDN URL is configured, we can't easily guess the domain.
	// But usually Qiniu requires a bound domain.
	return ret.Key, nil
}
