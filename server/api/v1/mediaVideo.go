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

// @Summary      上传视频到磁盘
// @Description  上传视频文件并保存到本地存储
// @Tags         媒体
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "视频文件"
// @Success      200   {object}  map[string]interface{}  "上传成功"
// @Failure      400   {object}  map[string]interface{}  "上传失败"
// @Security     BearerAuth
// @Router       /v1/mediaVideo/uploadToDisk [post]
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

// @Summary      创建视频记录
// @Description  根据已有附件信息创建视频媒体记录
// @Tags         媒体
// @Accept       json
// @Produce      json
// @Param        body  body  model.Attachment  true  "附件信息"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Security     BearerAuth
// @Router       /v1/mediaVideo/create [post]
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

// @Summary      获取视频列表
// @Description  分页查询视频列表
// @Tags         媒体
// @Accept       json
// @Produce      json
// @Param        body  body  model.VideoQueryRequest  true  "查询条件"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/mediaVideo/list [post]
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

// @Summary      根据ID获取视频
// @Description  根据视频ID获取视频详情
// @Tags         媒体
// @Accept       json
// @Produce      json
// @Param        body  body  model.VideoQueryRequest  true  "视频ID"
// @Success      200   {object}  map[string]interface{}  "成功"
// @Failure      400   {object}  map[string]interface{}  "参数错误"
// @Router       /v1/mediaVideo/byId [post]
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

// @Summary      获取视频存储类型
// @Description  获取当前视频存储磁盘类型配置
// @Tags         媒体
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Security     BearerAuth
// @Router       /v1/mediaVideo/disk [get]
func (ctrl *MediaVideoController) UploadDisk(c *gin.Context) {
	disk := ctrl.mediaVideoSvr.GetVideoUploadDisk()

	http.SuccessData(c, "数据获取成功!", gin.H{"disk": disk})
}

// @Summary      获取七牛视频上传Token
// @Description  获取七牛云存储视频上传令牌
// @Tags         媒体
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "成功"
// @Failure      400  {object}  map[string]interface{}  "获取失败"
// @Security     BearerAuth
// @Router       /v1/mediaVideo/token/qiniu [post]
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
