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

// @Summary      上传图片到磁盘
// @Description  上传图片文件并保存到本地存储
// @Tags         媒体
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "图片文件"
// @Success      200   {object}  map[string]interface{}  "上传成功"
// @Failure      400   {object}  map[string]interface{}  "上传失败"
// @Security     BearerAuth
// @Router       /v1/mediaImage/uploadToDisk [post]
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

// @Summary      创建图片记录
// @Description  根据已有附件信息创建图片媒体记录
// @Tags         媒体
// @Accept       json
// @Produce      json
// @Param        body  body  model.Attachment  true  "附件信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/mediaImage/create [post]
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

// @Summary      获取图片列表
// @Description  分页查询图片列表
// @Tags         媒体
// @Accept       json
// @Produce      json
// @Param        body  body  model.ImageQueryRequest  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/mediaImage/list [post]
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

// @Summary      根据ID获取图片
// @Description  根据图片ID获取图片详情
// @Tags         媒体
// @Accept       json
// @Produce      json
// @Param        body  body  model.ImageQueryRequest  true  "图片ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/mediaImage/byId [post]
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

// @Summary      获取图片存储类型
// @Description  获取当前图片存储磁盘类型配置
// @Tags         媒体
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Security     BearerAuth
// @Router       /v1/mediaImage/disk [get]
func (ctrl *MediaImageController) UploadDisk(c *gin.Context) {
	disk := ctrl.mediaImageSvr.GetImageUploadDisk()

	http.SuccessData(c, "数据获取成功!", gin.H{"disk": disk})
}

// @Summary      获取七牛上传Token
// @Description  获取七牛云存储上传令牌
// @Tags         媒体
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Failure      400  {object}  map[string]interface{}  "获取失败"
// @Security     BearerAuth
// @Router       /v1/mediaImage/token/qiniu [post]
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
