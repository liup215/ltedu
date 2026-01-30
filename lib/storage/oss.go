package storage

import (
	"edu/model"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliyunOSS struct {
	config model.ImageUploadConfig
}

func NewAliyunOSS(config model.ImageUploadConfig) (*AliyunOSS, error) {
	return &AliyunOSS{config: config}, nil
}

func (s *AliyunOSS) Upload(objectName string, localPath string) (url string, err error) {
	client, err := oss.New(s.config.OssEndpoint, s.config.OssAccessKeyId, s.config.OssAccessKeySecret)
	if err != nil {
		return "", err
	}

	bucket, err := client.Bucket(s.config.OssBucket)
	if err != nil {
		return "", err
	}

	// Explicitly set the object ACL to Public Read to ensure it can be accessed
	// This requires the RAM user to have proper permissions (e.g., AliyunOSSFullAccess)
	err = bucket.PutObjectFromFile(objectName, localPath, oss.ACL(oss.ACLPublicRead))
	if err != nil {
		return "", err
	}

	// Calculate and return the URL
	if s.config.OssCDNUrl != "" {
		baseUrl := s.config.OssCDNUrl
		if !strings.HasSuffix(baseUrl, "/") {
			baseUrl += "/"
		}
		return baseUrl + objectName, nil
	}

	// Default to bucket domain
	// https://bucket.endpoint/object
	scheme := "https://"
	endpoint := s.config.OssEndpoint
	if strings.HasPrefix(endpoint, "http://") {
		scheme = "http://"
		endpoint = strings.TrimPrefix(endpoint, "http://")
	} else if strings.HasPrefix(endpoint, "https://") {
		endpoint = strings.TrimPrefix(endpoint, "https://")
	}

	return scheme + s.config.OssBucket + "." + endpoint + "/" + objectName, nil
}
