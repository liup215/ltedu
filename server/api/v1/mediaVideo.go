package v1

import (
	"edu/lib/net/http"
	"edu/model"
	"edu/service"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

var MediaVideoCtrl = &MediaVideoController{mediaVideoSvr: service.MediaVideoSvr, attachmentSvr: service.AttachmentSvr}

type MediaVideoController struct {
	mediaVideoSvr *service.MediaVideoService
	attachmentSvr *service.AttachmentService
}

func (ctrl *MediaVideoController) UploadVideoToDisk(c *gin.Context) {
	name := "file"
	fileheader, err := c.FormFile(name)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	attachment, err := ctrl.saveVideo(fileheader)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}
	attachment.Type = model.AttachmentTypeVideo

	err = ctrl.attachmentSvr.CreateAttachment(attachment)
	if err != nil {
		http.ErrorData(c, err.Error(), nil)
		return
	}

	// ctx.JSON(http.StatusOK, ginResponse{Code: http.StatusOK, Message: "ok", Data: attachment})
	http.SuccessData(c, "上传成功！", attachment)
}

func (ctrl *MediaVideoController) saveVideo(fileHeader *multipart.FileHeader) (attachment *model.Attachment, err error) {

	video, err := ctrl.mediaVideoSvr.SaveVideo(fileHeader)

	attachment = &model.Attachment{
		Size:   fileHeader.Size,
		Name:   fileHeader.Filename,
		Ext:    video.Ext,
		Enable: true, // 默认都是合法的
		Hash:   video.Hash,
		Path:   video.Url,
		Width:  video.Width,
		Height: video.Height,
	}

	err = ctrl.attachmentSvr.CreateAttachment(attachment)
	if err != nil {
		return
	}

	// err = dao.MediaVideo.Omit(field.AssociationFields).Create(video)
	err = ctrl.mediaVideoSvr.CreateVideo(*video)

	return
}

func (ctrl *MediaVideoController) CreateVideo(c *gin.Context) {
	att := model.Attachment{}
	if err := c.BindJSON(&att); err != nil {
		http.ErrorData(c, "参数绑定失败: "+err.Error(), nil)
		return
	}

	att.Type = model.AttachmentTypeVideo
	ctrl.attachmentSvr.CreateAttachment(&att)

	video := model.MediaVideo{
		AttachmentId: att.ID,
		Name:         att.Name,
		Hash:         att.Hash,
		Path:         att.Path,
		Ext:          att.Ext,
		Width:        att.Width,
		Height:       att.Height,
		Duration:     att.Duration,
	}

	if err := ctrl.mediaVideoSvr.CreateVideo(video); err != nil {
		http.ErrorData(c, "创建失败: "+err.Error(), nil)
		return
	}

	http.SuccessData(c, "创建成功！", nil)
}

func (ctrl *MediaVideoController) SelectVideoList(c *gin.Context) {
	req := model.VideoQueryRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数绑定失败:"+err.Error(), nil)
		return
	}

	if list, total, err := ctrl.mediaVideoSvr.SelectVideoList(req); err != nil {
		http.ErrorData(c, "数据获取失败:"+err.Error(), nil)
		return
	} else {
		http.SuccessData(c, "数据获取成功!", gin.H{"list": list, "total": total})
	}
}

func (ctrl *MediaVideoController) SelectVideoById(c *gin.Context) {
	req := model.VideoQueryRequest{}
	if err := c.BindJSON(&req); err != nil {
		http.ErrorData(c, "参数绑定失败:"+err.Error(), nil)
		return
	}

	if video, err := ctrl.mediaVideoSvr.SelectVideoById(req.ID); err != nil {
		http.ErrorData(c, "数据获取失败:"+err.Error(), nil)
		return
	} else {
		http.SuccessData(c, "数据获取成功!", video)
	}
}

func (ctrl *MediaVideoController) UploadDisk(c *gin.Context) {
	disk := ctrl.mediaVideoSvr.GetVideoUploadDisk()

	http.SuccessData(c, "数据获取成功!", gin.H{"disk": disk})
}

func (ctrl *MediaVideoController) QiniuUploadToken(c *gin.Context) {

	token, key, e1 := ctrl.mediaVideoSvr.QiniuUploadToken()
	if e1 != nil {
		http.ErrorData(c, "token获取失败!", nil)
		return
	}
	http.SuccessData(c, "数据获取成功!", gin.H{
		"token": token,
		"key":   key,
	})
}
