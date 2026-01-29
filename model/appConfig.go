package model

const (
	APPCONFIG_PRIVATE_IS  = 1
	APPCONFIT_PRIVATE_NOT = 2

	APPCONFIG_SHOW_IS  = 1
	APPCONFIG_SHOW_NOT = 2
)

type AppConfig struct {
	Model
	Name         string `json:"name"`
	Group        string `json:"group"`
	FieldType    string `json:"fieldType"`
	Sort         int    `json:"sort"`
	DefaultValue string `json:"defaultValue"`
	Key          string `json:"key"`
	Value        string `json:"value" gorm:"type:longtext;"`
	IsPrivate    int    `json:"isPrivate"`
	OptionValue  string `json:"optionValue"`
	Help         string `json:"help"`
	IsShow       int    `json:"isShow"`
}

const (
	LTEDU_CONFIG_IMAGE_UPLOAD_KEY         = "ltedu.upload.image.upload"
	LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC = "public"
	LTEDU_CONFIG_IMAGE_UPLOAD_DISK_OSS    = "oss"
	LTEDU_CONFIG_IMAGE_UPLOAD_DISK_COS    = "cos"
	LTEDU_CONFIG_IMAGE_UPLOAD_DISK_QINIU  = "qiniu"
)

type ImageUploadConfig struct {
	Disk               string `json:"disk"`
	OssAccessKeyId     string `json:"ossAccessKeyId"`
	OssAccessKeySecret string `json:"ossAccessKeySecret"`
	OssBucket          string `json:"ossBucket"`
	OssEndpoint        string `json:"ossEndpoint"`
	OssCDNUrl          string `json:"ossCDNUrl"`
	CosRegion          string `json:"cosRegion"`
	CosAppId           string `json:"cosAppId"`
	CosSecretId        string `json:"cosSecretId"`
	CosSecretKey       string `json:"cosSecretKey"`
	CosBucket          string `json:"cosBucket"`
	CosCDNUrl          string `json:"cosCDNUrl"`
	QiniuAccessKey     string `json:"qiniuAccessKey"`
	QiniuSecretKey     string `json:"qiniuSecretKey"`
	QiniuCDNUrl        string `json:"qiniuCDNUrl"`
	QiniuCDNUrls       string `json:"qiniuCDNUrls"`
	QiniuBucket        string `json:"qiniuBucket"`
}

const (
	LTEDU_CONFIG_VIDEO_UPLOAD_KEY         = "ltedu.upload.video.upload"
	LTEDU_CONFIG_VIDEO_UPLOAD_DISK_PUBLIC = "public"
	LTEDU_CONFIG_VIDEO_UPLOAD_DISK_OSS    = "oss"
	LTEDU_CONFIG_VIDEO_UPLOAD_DISK_COS    = "cos"
	LTEDU_CONFIG_VIDEO_UPLOAD_DISK_QINIU  = "qiniu"
)

type VideoUploadConfig struct {
	Disk               string `json:"disk"`
	OssAccessKeyId     string `json:"ossAccessKeyId"`
	OssAccessKeySecret string `json:"ossAccessKeySecret"`
	OssBucket          string `json:"ossBucket"`
	OssEndpoint        string `json:"ossEndpoint"`
	OssCDNUrl          string `json:"ossCDNUrl"`
	CosRegion          string `json:"cosRegion"`
	CosAppId           string `json:"cosAppId"`
	CosSecretId        string `json:"cosSecretId"`
	CosSecretKey       string `json:"cosSecretKey"`
	CosBucket          string `json:"cosBucket"`
	CosCDNUrl          string `json:"cosCDNUrl"`
	QiniuAccessKey     string `json:"qiniuAccessKey"`
	QiniuSecretKey     string `json:"qiniuSecretKey"`
	QiniuCDNUrl        string `json:"qiniuCDNUrl"`
	QiniuCDNUrls       string `json:"qiniuCDNUrls"`
	QiniuBucket        string `json:"qiniuBucket"`
}

const (
	LTEDU_CONFIG_WEB_SITE_KEY = "website"
)

type WebSiteConfig struct {
	WebSiteName string `json:"webSiteName"`
	WebSiteUrl  string `json:"webSiteUrl"`
	WebSiteLogo string `json:"webSiteLogo"`
	ICPCode     string `json:"iCPCode"`
}
