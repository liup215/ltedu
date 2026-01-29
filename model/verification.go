package model

import "time"

const (
	VERIFICATION_TYPE_EMAIL  = 1 // 邮箱验证码
	VERIFICATION_TYPE_MOBILE = 2 // 手机验证码
)

const (
	VERIFICATION_STATUS_UNUSED  = 1 // 未使用
	VERIFICATION_STATUS_USED    = 2 // 已使用
	VERIFICATION_STATUS_EXPIRED = 3 // 已过期
)

type Verification struct {
	Model
	Type   int       `json:"type"`   // 1 邮箱验证码，2 手机验证码
	Target string    `json:"target"` // 目标，如邮箱或手机号
	Code   string    `json:"Code"`   // 验证码
	Expire time.Time `json:"expire"` // 过期时间
	Status int       `json:"status"` //
}

type VerificationQuery struct {
	ID     uint   `json:"id"`
	Type   int    `json:"type"`   // 1 图形验证码，2 邮箱验证码，3 手机验证码
	Target string `json:"target"` // 目标，如邮箱或手机号
}

type VerificationGenerateRequest struct {
	Type   int    `json:"type" binding:"required"`   // 1 图形验证码，2 邮箱验证码，3 手机验证码
	Target string `json:"target" binding:"required"` // 目标，如邮箱或手机号
}

type VerificationGenerateResponse struct {
	ID uint `json:"id"`
}
