package storage

import (
	"context"
	"edu/model"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentCOS struct {
	config model.ImageUploadConfig
}

func NewTencentCOS(config model.ImageUploadConfig) (*TencentCOS, error) {
	return &TencentCOS{config: config}, nil
}

func (s *TencentCOS) Upload(objectName string, localPath string) (urlStr string, err error) {
	// Bucket URL format: https://<bucket>-<appid>.cos.<region>.myqcloud.com
	bucketName := s.config.CosBucket
	if s.config.CosAppId != "" && !strings.HasSuffix(bucketName, "-"+s.config.CosAppId) {
		bucketName = bucketName + "-" + s.config.CosAppId
	}

	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucketName, s.config.CosRegion))
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  s.config.CosSecretId,
			SecretKey: s.config.CosSecretKey,
		},
	})

	// Upload file
	_, _, err = client.Object.Upload(
		context.Background(),
		objectName,
		localPath,
		nil,
	)
	if err != nil {
		return "", err
	}

	// Return URL
	if s.config.CosCDNUrl != "" {
		baseUrl := s.config.CosCDNUrl
		if !strings.HasSuffix(baseUrl, "/") {
			baseUrl += "/"
		}
		return baseUrl + objectName, nil
	}

	return u.String() + "/" + objectName, nil
}
