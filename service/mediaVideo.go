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

var MediaVideoSvr = &MediaVideoService{baseService: newBaseService()}

type MediaVideoService struct {
baseService
}

func (svr *MediaVideoService) CreateVideo(video model.MediaVideo) (err error) {
if video.AttachmentId == 0 {
err = errors.New("attachmentId is required")
return
}

if video.Path == "" {
err = errors.New("path is required")
return
}

if video.Url == "" {
cfg, e := repository.AppConfigRepo.FindByKey(model.LTEDU_CONFIG_IMAGE_UPLOAD_KEY)
if e != nil {
err = e
return
}

videoCfg := model.VideoUploadConfig{}

err = json.Unmarshal([]byte(cfg.Value), &videoCfg)
if err != nil {
return
}

if videoCfg.Disk == "" {
videoCfg.Disk = model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC
}

video.Disk = videoCfg.Disk

if videoCfg.Disk == model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC {
video.Url = video.Path
} else if videoCfg.Disk == model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_QINIU {
if videoCfg.QiniuCDNUrls != "" {
video.Url = videoCfg.QiniuCDNUrls + "/" + video.Path
} else {
video.Url = videoCfg.QiniuCDNUrl + "/" + video.Path
}

}
}
err = repository.MediaVideoRepo.Create(&video)
return
}

func (svr *MediaVideoService) GetVideoUploadDisk() string {
cfg, e := repository.AppConfigRepo.FindByKey(model.LTEDU_CONFIG_IMAGE_UPLOAD_KEY)
if e != nil {
return ""
}

videoCfg := model.VideoUploadConfig{}

e = json.Unmarshal([]byte(cfg.Value), &videoCfg)
if e != nil {
return ""
}

if videoCfg.Disk == "" {
videoCfg.Disk = model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC
}

return videoCfg.Disk
}

func (svr *MediaVideoService) QiniuUploadToken() (token string, key string, err error) {
key, err = svr.generateVideoName()
if err != nil {
return
}

// get qiniu token
cfg, err := repository.AppConfigRepo.FindByKey(model.LTEDU_CONFIG_IMAGE_UPLOAD_KEY)
if err != nil {
return
}

videoCfg := model.VideoUploadConfig{}
err = json.Unmarshal([]byte(cfg.Value), &videoCfg)
if err != nil {
return
}

putPolicy := storage.PutPolicy{
Scope:      videoCfg.QiniuBucket + ":" + key,
ReturnBody: `{"path": $(key), "size": $(fsize), "name": $(fname), "ext": $(ext), "hash": $(etag), "duration": $(avinfo.format.duration)}`,
}

mac := qbox.NewMac(videoCfg.QiniuAccessKey, videoCfg.QiniuSecretKey)
token = putPolicy.UploadToken(mac)
return
}

func (svr *MediaVideoService) SaveVideo(fileHeader *multipart.FileHeader) (video *model.MediaVideo, err error) {
cfg, err := repository.AppConfigRepo.FindByKey(model.LTEDU_CONFIG_IMAGE_UPLOAD_KEY)
if err != nil {
return
}

videoCfg := model.VideoUploadConfig{}

err = json.Unmarshal([]byte(cfg.Value), &videoCfg)
if err != nil {
return
}

if videoCfg.Disk == "" {
videoCfg.Disk = model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC

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
video = &model.MediaVideo{}

if videoCfg.Disk == model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC {
url := fmt.Sprintf("images/%s/%s%s", strings.Join(strings.Split(fileName, "")[0:5], "/"), fileName, ext)
savePath := filepath.Join(conf.Conf.Public, url)
err = os.MkdirAll(filepath.Dir(savePath), os.ModePerm)

if err != nil {
err = errors.New("fold create failed: " + err.Error())
return
}
video.Url = savePath
err = utils.CopyFile(cachePath, savePath)
} else if videoCfg.Disk == model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_OSS {
// video.Url = videoCfg.OssCDNUrl + "/" + fileName
video, err = svr.saveVideoOss(fileHeader)
} else if videoCfg.Disk == model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_COS {
// video.Url = videoCfg.CosCDNUrl + "/" + fileName
video, err = svr.saveVideoCos(fileHeader)
} else if videoCfg.Disk == model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_QINIU {
video.Url = videoCfg.QiniuCDNUrl + "/" + fileName
err = svr.saveVideoQiniu(fileHeader, videoCfg)
} else {
err = errors.New("invalid disk")
}

if err != nil {
return
}

video.Name = fileName
video.Disk = videoCfg.Disk
video.Hash = md5hash
video.Ext = ext

return
}

func (svr *MediaVideoService) saveVideoOss(f *multipart.FileHeader) (video *model.MediaVideo, err error) {
err = errors.New("Oss save failed: not implemented")
return
}

func (svr *MediaVideoService) saveVideoCos(f *multipart.FileHeader) (video *model.MediaVideo, err error) {
err = errors.New("Oss save failed: not implemented")
return
}

func (svr *MediaVideoService) saveVideoQiniu(f *multipart.FileHeader, videoCfg model.VideoUploadConfig) (err error) {
putPolicy := storage.PutPolicy{
Scope: videoCfg.QiniuBucket,
}

mac := qbox.NewMac(videoCfg.QiniuAccessKey, videoCfg.QiniuSecretKey)
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

func (s *MediaVideoService) generateVideoName() (string, error) {
prefix := "VID"
name := prefix + myStrings.Random(20)

for {
video, err := repository.MediaVideoRepo.FindByName(name)
if err == gorm.ErrRecordNotFound || video == nil {
break
} else if err != nil {
return "", err
} else {
name = prefix + myStrings.Random(20)
}
}

return name, nil
}

func (s *MediaVideoService) SelectVideoById(id uint) (*model.MediaVideo, error) {
if id == 0 {
return nil, errors.New("id is required")
}
return repository.MediaVideoRepo.FindByID(id)
}

func (s *MediaVideoService) SelectVideoList(req model.VideoQueryRequest) ([]*model.MediaVideo, int64, error) {
list, total, err := repository.MediaVideoRepo.FindPage(&req, (req.PageIndex-1)*req.PageSize, req.PageSize)
return list, total, err
}
