package storage

import (
	"edu/model"
	"errors"
)

type Storage interface {
	// Upload uploads a file from local path to the storage.
	// objectName is the destination path/name in the storage (e.g. "images/abc.jpg")
	// localPath is the path to the local file to upload.
	Upload(objectName string, localPath string) (url string, err error)
}

func NewStorage(config model.ImageUploadConfig) (Storage, error) {
	switch config.Disk {
	case model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC:
		return NewLocalStorage(), nil
	case model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_OSS:
		return NewAliyunOSS(config)
	case model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_COS:
		return NewTencentCOS(config)
	case model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_QINIU:
		return NewQiniuKodo(config)
	default:
		return nil, errors.New("unsupported storage type: " + config.Disk)
	}
}
