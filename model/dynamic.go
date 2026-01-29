package model

const (
	DynamicTypeComment        = 1  // 发表评论
	DynamicTypeFavorite       = 2  // 收藏文档
	DynamicTypeUpload         = 3  // 上传文档
	DynamicTypeDownload       = 4  // 下载文档
	DynamicTypeLogin          = 5  // 登录
	DynamicTypeRegister       = 6  // 注册
	DynamicTypeAvatar         = 7  // 更新了头像
	DynamicTypePassword       = 8  // 修改密码
	DynamicTypeInfo           = 9  // 修改个人信息
	DynamicTypeVerify         = 10 // 实名认证
	DynamicTypeSign           = 11 // 签到
	DynamicTypeShare          = 12 // 分享文档
	DynamicTypeFollow         = 13 // 关注用户
	DynamicTypeDeleteDocument = 14 // 删除文档
)

type Dynamic struct {
	Model
	UserId   int64  `json:"userId"`
	Username string `json:"userName"`
	UserType int    `json:"userType"`
	Content  string `json:"content"`
	Type     int    `json:"type"`
}
