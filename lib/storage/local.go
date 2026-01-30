package storage

import (
	"edu/lib/utils"
	"os"
	"path/filepath"
)

type LocalStorage struct{}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

func (s *LocalStorage) Upload(objectName string, localPath string) (url string, err error) {
	// objectName is like "uploads/a/b/c/abcde.jpg"
	// Ensure directory exists
	dir := filepath.Dir(objectName)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Copy file
	err = utils.CopyFile(localPath, objectName)
	if err != nil {
		return "", err
	}

	return "/" + objectName, nil
}
