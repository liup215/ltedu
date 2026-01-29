package service

import (
"bytes"
"context"
"edu/conf"
"edu/repository"
myStrings "edu/lib/strings"
"edu/lib/utils"
"edu/model"
"encoding/json"
"errors"
"fmt"
"io"
"mime/multipart"
"os"
"path/filepath"
"strings"
"time"

"github.com/gofrs/uuid"
"github.com/qiniu/go-sdk/v7/auth/qbox"
"github.com/qiniu/go-sdk/v7/storage"
"gorm.io/gorm"
)

var MediaImageSvr = &MediaImageService{baseService: newBaseService()}

type MediaImageService struct {
baseService
}

func (svr *MediaImageService) CreateImage(img model.MediaImage) (err error) {
if img.AttachmentId == 0 {
err = errors.New("attachmentId is required")
return
}

if img.Path == "" {
err = errors.New("path is required")
return
}

if img.Url == "" {
cfg, e := repository.AppConfigRepo.FindByKey(model.LTEDU_CONFIG_IMAGE_UPLOAD_KEY)
if e != nil {
err = e
return
}

imgCfg := model.ImageUploadConfig{}

err = json.Unmarshal([]byte(cfg.Value), &imgCfg)
if err != nil {
return
}

if imgCfg.Disk == "" {
imgCfg.Disk = model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC
}

img.Disk = imgCfg.Disk

if imgCfg.Disk == model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC {
img.Url = img.Path
} else if imgCfg.Disk == model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_QINIU {
if imgCfg.QiniuCDNUrls != "" {
img.Url = imgCfg.QiniuCDNUrls + "/" + img.Path
} else {
img.Url = imgCfg.QiniuCDNUrl + "/" + img.Path
}

}
}
err = repository.MediaImageRepo.Create(&img)
return
}

func (svr *MediaImageService) GetImageUploadDisk() string {
cfg, e := repository.AppConfigRepo.FindByKey(model.LTEDU_CONFIG_IMAGE_UPLOAD_KEY)
if e != nil {
return ""
}

imgCfg := model.ImageUploadConfig{}

e = json.Unmarshal([]byte(cfg.Value), &imgCfg)
if e != nil {
return ""
}

if imgCfg.Disk == "" {
imgCfg.Disk = model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC
}

return imgCfg.Disk
}

func (svr *MediaImageService) QiniuUploadToken() (token string, key string, err error) {
key, err = svr.generateImageName()
if err != nil {
return
}

// get qiniu token
cfg, err := repository.AppConfigRepo.FindByKey(model.LTEDU_CONFIG_IMAGE_UPLOAD_KEY)
if err != nil {
return
}

imgCfg := model.ImageUploadConfig{}
err = json.Unmarshal([]byte(cfg.Value), &imgCfg)
if err != nil {
return
}

putPolicy := storage.PutPolicy{
Scope:      imgCfg.QiniuBucket + ":" + key,
ReturnBody: `{"path": $(key), "size": $(fsize), "width": $(imageInfo.width), "height": $(imageInfo.height), "name": $(fname), "ext": $(ext), "hash": $(etag)}`,
}

mac := qbox.NewMac(imgCfg.QiniuAccessKey, imgCfg.QiniuSecretKey)
token = putPolicy.UploadToken(mac)
return
}

func (svr *MediaImageService) SaveImage(fileHeader *multipart.FileHeader) (img *model.MediaImage, err error) {
cfg, err := repository.AppConfigRepo.FindByKey(model.LTEDU_CONFIG_IMAGE_UPLOAD_KEY)
if err != nil {
return
}

imgCfg := model.ImageUploadConfig{}

err = json.Unmarshal([]byte(cfg.Value), &imgCfg)
if err != nil {
return
}

if imgCfg.Disk == "" {
imgCfg.Disk = model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC

}

cacheDir := fmt.Sprintf("cache/uploads/%s", time.Now().Format("2006/01/02"))
os.MkdirAll(cacheDir, os.ModePerm)
ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
cachePath := fmt.Sprintf("%s/%s%s", cacheDir, uuid.Must(uuid.NewV4()).String(), ext)
defer func() {
os.Remove(cachePath)
}()

// 保存到临时文件

src, err := fileHeader.Open()
if err != nil {
return
}
defer src.Close()

out, err := os.Create(cachePath)
if err != nil {
err = errors.New("fold create failed: " + err.Error())
return
}
defer out.Close()

_, err = io.Copy(out, src)

if err != nil {
return
}

// 获取文件md5值
md5hash, errHash := utils.GetFileMD5(cachePath)
if errHash != nil {
err = errHash
return
}

fileName := md5hash
fileHeader.Filename = fileName
img = &model.MediaImage{}

if imgCfg.Disk == model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC {
url := fmt.Sprintf("images/%s/%s%s", strings.Join(strings.Split(fileName, "")[0:5], "/"), fileName, ext)
savePath := filepath.Join(conf.Conf.Public, url)
err = os.MkdirAll(filepath.Dir(savePath), os.ModePerm)

if err != nil {
err = errors.New("fold create failed: " + err.Error())
return
}
img.Url = savePath
err = utils.CopyFile(cachePath, savePath)
} else if imgCfg.Disk == model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_OSS {
// img.Url = imgCfg.OssCDNUrl + "/" + fileName
img, err = svr.saveImageOss(fileHeader)
} else if imgCfg.Disk == model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_COS {
// img.Url = imgCfg.CosCDNUrl + "/" + fileName
img, err = svr.saveImageCos(fileHeader)
} else if imgCfg.Disk == model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_QINIU {
img.Url = imgCfg.QiniuCDNUrl + "/" + fileName
err = svr.saveImageQiniu(fileHeader, imgCfg)
} else {
err = errors.New("invalid disk")
}

if err != nil {
return
}

img.Name = fileName
img.Disk = imgCfg.Disk
img.Hash = md5hash
img.Ext = ext
img.Width, img.Height, _ = utils.GetImageSize(cachePath)

return
}

func (svr *MediaImageService) saveImageOss(f *multipart.FileHeader) (img *model.MediaImage, err error) {
err = errors.New("Oss save failed: not implemented")
return
}

func (svr *MediaImageService) saveImageCos(f *multipart.FileHeader) (img *model.MediaImage, err error) {
err = errors.New("Oss save failed: not implemented")
return
}

func (svr *MediaImageService) saveImageQiniu(f *multipart.FileHeader, imgCfg model.ImageUploadConfig) (err error) {
putPolicy := storage.PutPolicy{
Scope: imgCfg.QiniuBucket,
}

mac := qbox.NewMac(imgCfg.QiniuAccessKey, imgCfg.QiniuSecretKey)
upToken := putPolicy.UploadToken(mac)

cfg := storage.Config{}
cfg.Zone = &storage.ZoneHuanan
cfg.UseHTTPS = false
cfg.UseCdnDomains = false

formUploader := storage.NewFormUploader(&cfg)
ret := storage.PutRet{}
putExtra := storage.PutExtra{
Params: map[string]string{},
}

file, err := f.Open()
if err != nil {
return
}

defer file.Close()
data, err := io.ReadAll(file)
if err != nil {
return
}
dataLen := int64(len(data))
err = formUploader.Put(context.Background(), &ret, upToken, f.Filename, bytes.NewReader(data), dataLen, &putExtra)

if err != nil {
err = errors.New("qiniu save failed:" + err.Error())
}
return
}

func (s *MediaImageService) generateImageName() (string, error) {
prefix := "IMG"
name := prefix + myStrings.Random(20)

for {
img, err := repository.MediaImageRepo.FindByName(name)
if err == gorm.ErrRecordNotFound || img == nil {
break
} else if err != nil {
return "", err
} else {
name = prefix + myStrings.Random(20)
}
}

return name, nil
}

func (s *MediaImageService) SelectImageById(id uint) (*model.MediaImage, error) {
if id == 0 {
return nil, errors.New("id is required")
}
return repository.MediaImageRepo.FindByID(id)
}

func (s *MediaImageService) SelectImageList(req model.ImageQueryRequest) ([]*model.MediaImage, int64, error) {
list, total, err := repository.MediaImageRepo.FindPage(&req, (req.PageIndex-1)*req.PageSize, req.PageSize)
return list, total, err
}
