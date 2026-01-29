package v1

import (
	"edu/conf"
	"edu/lib/logger"
	"edu/lib/net/http"
	"edu/lib/utils"
	"edu/model"
	"edu/service"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

var AttachmentCtrl = &AttachmentController{
	attachmentSvr: service.AttachmentSvr,
	mediaImageSvr: service.MediaImageSvr,
}

type AttachmentController struct {
	attachmentSvr *service.AttachmentService
	mediaImageSvr *service.MediaImageService
}

func (s *AttachmentController) UploadDocument(c *gin.Context) {

	name := "file[0]"
	fileheader, err := c.FormFile(name)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	unsuportedExt := "不支持的文档类型"
	ext := strings.ToLower(filepath.Ext(fileheader.Filename))
	if !utils.IsDocument(ext) {
		http.ErrorData(c, unsuportedExt, nil)
		return
	}

	attachment, err := s.saveFile(c, fileheader, true)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	attachment.Type = model.AttachmentTypeDocument

	err = s.attachmentSvr.CreateAttachment(attachment)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	// ctx.JSON(http.StatusOK, ginResponse{Code: http.StatusOK, Message: "ok", Data: attachment})
	http.SuccessData(c, "上传成功！", attachment)
}

// saveFile 保存文件。文件以md5值命名以及存储
// 同时，返回附件信息
func (s *AttachmentController) saveFile(ctx *gin.Context, fileHeader *multipart.FileHeader, isDocument ...bool) (attachment *model.Attachment, err error) {
	cacheDir := fmt.Sprintf("cache/uploads/%s", time.Now().Format("2006/01/02"))
	os.MkdirAll(cacheDir, os.ModePerm)
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	cachePath := fmt.Sprintf("%s/%s%s", cacheDir, uuid.Must(uuid.NewV4()).String(), ext)
	defer func() {
		os.Remove(cachePath)
	}()

	// 保存到临时文件
	err = ctx.SaveUploadedFile(fileHeader, cachePath)
	if err != nil {
		return
	}

	// 获取文件md5值
	md5hash, errHash := utils.GetFileMD5(cachePath)
	if errHash != nil {
		err = errHash
		return
	}

	savePathFormat := "uploads/%s/%s%s"
	if len(isDocument) > 0 && isDocument[0] {
		savePathFormat = "documents/%s/%s%s"
	}
	savePath := fmt.Sprintf(savePathFormat, strings.Join(strings.Split(md5hash, "")[0:5], "/"), md5hash, ext)
	os.MkdirAll(filepath.Dir(savePath), os.ModePerm)
	err = utils.CopyFile(cachePath, savePath)
	if err != nil {
		return
	}

	attachment = &model.Attachment{
		Size:   fileHeader.Size,
		Name:   fileHeader.Filename,
		Ext:    ext,
		Enable: true, // 默认都是合法的
		Hash:   md5hash,
		Path:   "/" + savePath,
	}

	// 对于图片，直接获取图片的宽高
	if utils.IsImage(ext) {
		attachment.Width, attachment.Height, _ = utils.GetImageSize(cachePath)
	}

	return
}

func (s *AttachmentController) ViewDocumentCover(c *gin.Context) {
	hash := c.Param("hash")
	if len(hash) != 32 {
		http.ErrorData(c, "无效文件", nil)
		return
	}

	file := fmt.Sprintf("documents/%s/%s/cover.png", strings.Join(strings.Split(hash, "")[:5], "/"), hash)
	if len(hash) != 32 {
		http.ErrorData(c, "文件不存在", nil)
		return
	}
	c.File(file)
}

func (s *AttachmentController) ViewDocumentPages(c *gin.Context) {
	hash := c.Param("hash")
	fmt.Println(hash)
	if len(hash) != 32 {
		http.ErrorData(c, "无效文件", nil)
		return
	}
	page := strings.TrimLeft(c.Param("page"), "./")
	if strings.HasSuffix(page, ".svg") {
		if strings.HasSuffix(page, ".gzip.svg") {
			c.Header("Content-Encoding", "gzip")
		}
		c.Header("Content-Type", "image/svg+xml")
	}

	file := fmt.Sprintf("documents/%s/%s/%s", strings.Join(strings.Split(hash, "")[:5], "/"), hash, page)
	logger.Logger.Debug("ViewDocumentPages", zap.String("hash", hash), zap.String("page", page), zap.String("file", file))
	c.File(file)
}

func (ctrl *AttachmentController) DownloadDocument(c *gin.Context) {
	claims := jwt.RegisteredClaims{}
	token := c.Param("jwt")
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Conf.SecretKey), nil
	})
	if err != nil {
		http.ErrorData(c, "无效的下载链接", nil)
		return
	}

	filename := c.Query("filename")
	file := fmt.Sprintf("documents/%s/%s%s", strings.Join(strings.Split(claims.ID, "")[:5], "/"), claims.ID, filepath.Ext(filename))
	c.FileAttachment(file, filename)
}
