package model

type MediaVideo struct {
	Model
	Name         string     `json:"name"`
	Disk         string     `json:"disk"`
	From         string     `json:"from"`
	Path         string     `json:"path"`
	Url          string     `json:"url"`
	Ext          string     `json:"ext"`
	Hash         string     `json:"hash"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	AttachmentId uint       `json:"attachment_id"`
	Attachment   Attachment `json:"attachment"`
}

type VideoQueryRequest struct {
Model
Page
Name string `json:"name"`
Hash string `json:"hash"`
Disk string `json:"disk"`
}

type MediaVideoUploadToken struct {
	Disk       string `json:"disk"`
	Qiniutoken string `json:"qiniutoken"`
}

type MediaVideoUploadTokenQiniu struct {
	Token string `json:"token"`
	Key   string `json:"key"`
}
