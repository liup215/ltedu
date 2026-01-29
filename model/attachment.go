package model

const (
	AttachmentTypeUnknown       = iota // 未知
	AttachmentTypeAvatar               // 用户头像
	AttachmentTypeDocument             // 文档
	AttachmentTypeArticle              // 文章
	AttachmentTypeComment              // 评论
	AttachmentTypeBanner               // 横幅
	AttachmentTypeCategoryCover        // 分类封面
	AttachmentTypeConfig               // 配置
	AttachmentTypeImage                // 图片
	AttachmentTypeVideo                // 视频
)

var AttachmentTypeName = map[int]string{
	AttachmentTypeAvatar:        "头像",
	AttachmentTypeArticle:       "文章",
	AttachmentTypeBanner:        "横幅",
	AttachmentTypeCategoryCover: "分类封面",
	AttachmentTypeComment:       "评论",
	AttachmentTypeDocument:      "文档",
	AttachmentTypeConfig:        "配置",
	AttachmentTypeImage:         "图片",
	AttachmentTypeVideo:         "视频",
}

type Attachment struct {
	Model
	Hash        string `json:"hash"`
	Type        int    `json:"type"`
	Enable      bool   `json:"enable"`
	Path        string `json:"path"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Ext         string `json:"ext"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
}
