package v1

import (
	"edu/lib/net/http"
	"edu/lib/utils"
	"edu/model"
	"edu/service"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var MediaImageCtrl = &MediaImageController{mediaImageSvr: service.MediaImageSvr, attachmentSvr: service.AttachmentSvr}

type MediaImageController struct {
	mediaImageSvr *service.MediaImageService
	attachmentSvr *service.AttachmentService
}

func (ctrl *MediaImageController) UploadImageToDisk(c *gin.Context) {
	name := "file"
	fileheader, err := c.FormFile(name)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	unsuportedExt := "不支持的图片类型"
	ext := strings.ToLower(filepath.Ext(fileheader.Filename))
	if !utils.IsImage(ext) {
		http.ErrorData(c, unsuportedExt, nil)
		return
	}

	attachment, err := ctrl.saveImage(fileheader)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	attachment.Type = model.AttachmentTypeImage

	err = ctrl.attachmentSvr.CreateAttachment(attachment)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	// ctx.JSON(http.StatusOK, ginResponse{Code: http.StatusOK, Message: "ok", Data: attachment})
	http.SuccessData(c, "上传成功！", attachment)
}

func (ctrl *MediaImageController) saveImage(fileHeader *multipart.FileHeader) (attachment *model.Attachment, err error) {

	img, err := ctrl.mediaImageSvr.SaveImage(fileHeader)

	attachment = &model.Attachment{
		Size:   fileHeader.Size,
		Name:   fileHeader.Filename,
		Ext:    img.Ext,
		Enable: true, // 默认都是合法的
		Hash:   img.Hash,
		Path:   img.Url,
		Width:  img.Width,
		Height: img.Height,
	}

	err = ctrl.attachmentSvr.CreateAttachment(attachment)
	if err != nil {
		return
	}

	// err = dao.MediaImage.Omit(field.AssociationFields).Create(img)
	err = ctrl.mediaImageSvr.CreateImage(*img)

	return
}

func (ctrl *MediaImageController) CreateImage(c *gin.Context) {
	att := model.Attachment{}
	if err := c.BindJSON(&att); err != nil {
		http.ErrorData(c, "参数绑定失败: "+err.Error(), nil)
		return
	}

	att.Type = model.AttachmentTypeImage
	ctrl.attachmentSvr.CreateAttachment(&att)

	img := model.MediaImage{
		AttachmentId: att.ID,
		Name:         att.Name,
		Hash:         att.Hash,
		Path:         att.Path,
		Ext:          att.Ext,
		Width:        att.Width,
		Height:       att.Height,
	}

	if err := ctrl.mediaImageSvr.CreateImage(img); err != nil {
		http.ErrorData(c, "创建失败: "+err.Error(), nil)
		return
	}

	http.SuccessData(c, "创建成功！", nil)
}

func (ctrl *MediaImageController) SelectImageList(c *gin.Context) {
	req := model.ImageQueryRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数绑定失败:"+err.Error(), nil)
		return
	}

	if list, total, err := ctrl.mediaImageSvr.SelectImageList(req); err != nil {
		http.ErrorData(c, "数据获取失败:"+err.Error(), nil)
		return
	} else {
		http.SuccessData(c, "数据获取成功!", gin.H{"list": list, "total": total})
	}
}

func (ctrl *MediaImageController) SelectImageById(c *gin.Context) {
	req := model.ImageQueryRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数绑定失败:"+err.Error(), nil)
		return
	}

	if image, err := ctrl.mediaImageSvr.SelectImageById(req.ID); err != nil {
		http.ErrorData(c, "数据删除失败:"+err.Error(), nil)
		return
	} else {
		http.SuccessData(c, "数据删除成功!", image)
	}
}

func (ctrl *MediaImageController) UploadDisk(c *gin.Context) {
	disk := ctrl.mediaImageSvr.GetImageUploadDisk()

	http.SuccessData(c, "数据获取成功!", gin.H{"disk": disk})
}

func (ctrl *MediaImageController) QiniuUploadToken(c *gin.Context) {

	token, key, e1 := ctrl.mediaImageSvr.QiniuUploadToken()
	if e1 != nil {
		http.ErrorData(c, "token获取失败!", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"token": token,
		"key":   key,
	})
}
