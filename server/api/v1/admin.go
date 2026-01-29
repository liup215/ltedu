package v1

import (
	"edu/model"
	"edu/service"
)

func init() {
	AdminCtrl = &AdminController{
		adminSvr: service.AdminSvr,
	}
}

var AdminCtrl *AdminController

type AdminController struct {
	// loginSvr *service.LoginService
	adminSvr *service.AdminService
}

type AdminUserInfoResponse struct {
	ID          uint                     `json:"id"`
	Username    string                   `json:"username"`
	Email       string                   `json:"email"`
	IsAdmin     bool                     `json:"isAdmin"`
	AdminRoleID *uint                    `json:"adminRoleId,omitempty"`
	AdminRole   *model.AdminRole         `json:"adminRole,omitempty"` // Include AdminRole details
	Permissions []*model.AdminPermission `json:"permissions,omitempty"`
}
