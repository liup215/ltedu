package model

const (
	ADMIN_STATUS_OK       = 1
	ADMIN_STATUS_DISABLED = 2
)

// The Admin struct is now removed. Admin-related fields are moved to the User model.
// Admin-specific status constants (ADMIN_STATUS_OK, ADMIN_STATUS_DISABLED) are kept here
// as they define the states for User.AdminStatus.

const (
	ADMINLOG_OPT_VIEW    = "VIEW"
	ADMINLOG_OPT_STORE   = "STORE"
	ADMINLOG_OPT_IMPORT  = "IMPORT"
	ADMINLOG_OPT_UPDATE  = "UPDATE"
	ADMINLOG_OPT_DESTROY = "DESTROY"
	ADMINLOG_OPT_LOGIN   = "LOGIN"
	ADMINLOG_OPT_LOGOUT  = "LOGOUT"
	ADMINLOG_OPT_REFUND  = "REFUND"

	ADMINLOG_MODULE_VOD               = "vod"
	ADMINLOG_MODULE_VOD_VIDEO         = "vod-video"
	ADMINLOG_MODULE_VOD_VIDEO_COMMENT = "vod-video-comment"
	ADMINLOG_MODULE_VOD_ATTACH        = "vod-attach"
	ADMINLOG_MODULE_VOD_CATEGORY      = "vod-category"
	ADMINLOG_MODULE_VOD_CHAPTER       = "vod-chapter"
	ADMINLOG_MODULE_VOD_COMMENT       = "vod-comment"

	ADMINLOG_MODULE_ADMIN_DASHBOARD   = "admin-dashboard"
	ADMINLOG_MODULE_ADMIN_LOGIN       = "admin-login"
	ADMINLOG_MODULE_ADMIN_MEDIA_IMAGE = "admin-media-image"
	ADMINLOG_MODULE_ADMIN_MEDIA_VIDEO = "admin-media-video"

	ADMINLOG_MODULE_SYSTEM_CONFIG = "system-config"

	ADMINLOG_MODULE_STATS = "stats"

	ADMINLOG_MODULE_MEMBER     = "member"
	ADMINLOG_MODULE_MEMBER_TAG = "member-tag"

	ADMINLOG_MODULE_MP_MENU    = "mp-menu"
	ADMINLOG_MODULE_MP_MESSAGE = "mp-message"

	ADMINLOG_MODULE_VIP          = "vip"
	ADMINLOG_MODULE_ADDONS       = "addons"
	ADMINLOG_MODULE_AD_FROM      = "ad-from"
	ADMINLOG_MODULE_LINK         = "link"
	ADMINLOG_MODULE_NAV          = "nav"
	ADMINLOG_MODULE_SLIDER       = "slider"
	ADMINLOG_MODULE_ORDER        = "order"
	ADMINLOG_MODULE_VIEW_BLOCK   = "view-block"
	ADMINLOG_MODULE_PROMO_CODE   = "promo-code"
	ADMINLOG_MODULE_ANNOUNCEMENT = "announcement"

	ADMINLOG_MODULE_ADMINISTRATOR      = "administrator"
	ADMINLOG_MODULE_ADMINISTRATOR_ROLE = "administrator-role"
)

type AdminLog struct {
	Model
	AdminId uint   `json:"adminId"`
	Module  string `json:"module"`
	Opt     string `json:"opt"`
	Remark  string `json:"remark"`
	Ip      string `json:"ip"`
}

// AdminRolePermission is the explicit join table for AdminRole <-> AdminPermission many2many.
type AdminRolePermission struct {
	AdminRoleID       uint `gorm:"primaryKey"`
	AdminPermissionID uint `gorm:"primaryKey"`
}

// UserRole is the explicit join table for User <-> AdminRole many2many.
type UserRole struct {
	UserID      uint `gorm:"primaryKey"`
	AdminRoleID uint `gorm:"primaryKey"`
}

type AdminPermission struct {
	// 'display_name', 'slug', 'description',
	//     'method', 'url', 'route', 'group_name',
	Model
	DisplayName string `json:"displayName"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Method      string `json:"method"`
	Url         string `json:"url"`
	Route       string `json:"route"`
	GroupName   string `json:"groupName"`
}

type AdminRole struct {
	Model
	DisplayName string             `json:"displayName"`
	Slug        string             `json:"slug"`
	Description string             `json:"description"`
	Permissions []*AdminPermission `json:"permissions,omitempty" gorm:"many2many:admin_role_permissions;"`
}
