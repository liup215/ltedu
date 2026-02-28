package model

import "time"

const (
	UserStatusNormal            = 1 // Active and normal user
	UserStatusPendingActivation = 2 // e.g. waiting for email verification
	UserStatusSuspended         = 3 // Temporarily suspended
	UserStatusBanned            = 4 // Permanently banned
)

type User struct {
	Model
	Username            string     `json:"username"`
	Email               string     `json:"email" gorm:"unique,length:100"` // Added Email field
	Nickname            string     `json:"nickname"`
	Realname            string     `json:"realname"`
	Engname             string     `json:"engname"`
	Sex                 uint       `json:"sex"` // 1 男，2 女
	Password            string     `json:"-"`
	Salt                string     `json:"-"`
	TokenSalt           string     `json:"-" gorm:"size:32"` // 用于Token签名的动态密钥，仅后端使用
	Mobile              string     `json:"mobile"`
	Avatar              string     `json:"avatar"`
	Status              int        `json:"status"`
	IsActive            int        `json:"isActive"`
	InviteUserId        uint       `json:"inviteUserId"`
	InviteBalance       int        `json:"inviteBalance"`
	InviteUserExpiredAt *time.Time `json:"inviteUserExpiredAt"`
	IsPasswordSet       int        `json:"isPasswordSet"`
	IsUsernameSet       int        `json:"isUsernameSet"`
	RegisterIp          string     `json:"registerIp"`
	RegisterArea        string     `json:"registerArea"`
	LastLoginId         uint       `json:"lastLoginId"`
	Classes             []*Class   `json:"classes" gorm:"many2many:user_class_relation;"`
	FinalPerformMark    int        `json:"finalPerformMark"`
	FinalActivityMark   int        `json:"finalActivityMark"`
	// Admin-related fields (moved from old Admin model or new)
	IsAdmin       bool       `json:"isAdmin" gorm:"default:false"`
	AdminRoleID   *uint      `json:"adminRoleId,omitempty" gorm:"index"`                // Pointer to allow null
	AdminRole     *AdminRole `json:"adminRole,omitempty" gorm:"foreignKey:AdminRoleID"` // Pointer to allow null
	AdminStatus   *int       `json:"adminStatus,omitempty"`                             // Pointer to allow null, uses ADMIN_STATUS_OK etc. from model/admin.go
	LastLoginIp   string     `json:"lastLoginIp,omitempty" gorm:"size:50"`
	LastLoginDate *time.Time `json:"lastLoginDate,omitempty"`
	LoginTimes    uint       `json:"loginTimes,omitempty" gorm:"default:0"`
	// Teacher-related fields
	IsTeacher            bool  `json:"isTeacher" gorm:"default:false"`
	TeacherApplyStatus   int   `json:"teacherApplyStatus" gorm:"default:0"`         // 0: Not Applied, 1: Pending, 2: Approved, 3: Rejected
	TeacherApplicationID *uint `json:"teacherApplicationId,omitempty" gorm:"index"` // Link to TeacherApplication

	// VIP expire field
	VipExpireAt *time.Time `json:"vipExpireAt"`
}

type UserQuery struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Realname string `json:"realname"`
	Engname  string `json:"engname"`
	Mobile   string `json:"mobile"`
	Status   int    `json:"status"`
	ClassId  uint   `json:"classId"`
	Page
}

type UserRegisterRequest struct {
	Username         string `json:"username"`
	Mobile           string `json:"mobile"`
	Password         string `json:"password"`
	PasswordConfirm  string `json:"passwordConfirm"`
	Captcha          string `json:"captcha"`
	Realname         string `json:"realname"`
	Email            string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verificationCode" binding:"required"`
}

type UserEditRequest struct {
	ID       uint   `json:"id"`
	Realname string `json:"realname"`
	Nickname string `json:"nickname"`
	Engname  string `json:"engname"`
	Sex      uint   `json:"sex"`
	Status   int    `json:"status"` // 1: Active, 2: Inactive, 3: Suspended, 4: Banned
}

// AccountUpdateRequest contains fields a user can update on their own account.
// Unlike UserEditRequest, it does not include ID (derived from JWT) or Status (admin-only).
type AccountUpdateRequest struct {
	Realname string `json:"realname"`
	Nickname string `json:"nickname"`
	Engname  string `json:"engname"`
	Sex      uint   `json:"sex"`
	Mobile   string `json:"mobile"`
}
