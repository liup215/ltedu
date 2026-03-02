package service

import (
"edu/lib/strings"
"edu/model"
"edu/repository"
"errors"
)

func init() {
	UserSvr = &UserService{
		baseService: newBaseService(),
	}

}

var UserSvr *UserService = &UserService{baseService: newBaseService()}

type UserService struct {
	baseService
}

// SetAdmin: 设置用户为管理员（通过RBAC角色分配）
func (svr *UserService) SetAdmin(userID uint) error {
	if userID == 0 {
		return errors.New("无效的ID")
	}
	// Find the "admin" role
	adminRole, err := repository.AdminRoleRepo.FindBySlug("admin")
	if err != nil {
		return errors.New("查找admin角色失败: " + err.Error())
	}
	if adminRole == nil {
		return errors.New("admin角色不存在，请先初始化权限数据")
	}
	return repository.UserRoleRepo.AssignRole(userID, adminRole.ID)
}

func (svr *UserService) SelectUserList(q model.UserQuery) ([]*model.User, int64, error) {
	return repository.UserRepo.FindList(q)
}

func (svr *UserService) GetUserCount() (int64, error) {
	return repository.UserRepo.Count()
}

func (svr *UserService) SelectUserById(id uint) (*model.User, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}

	return repository.UserRepo.FindByID(id)
}

func (svr *UserService) SelectUserByUsername(username string) (*model.User, error) {
	if username == "" {
		return nil, errors.New("无效的用户名")
	}

	return repository.UserRepo.FindByUsername(username)
}

func (svr *UserService) SelectUserAll(q model.UserQuery) ([]*model.User, error) {
	return repository.UserRepo.FindAll(q)
}

func (svr *UserService) CreateUser(o model.User) (*model.User, error) {
	if o.ID != 0 {
		o.ID = 0
	}
	if o.Mobile == "" || o.Password == "" {
		return nil, errors.New("手机号和密码不能为空！")
	}

	if o.IsActive == 0 {
		o.IsActive = model.YES
	}

	if o.Username != "" {
		record, _ := repository.UserRepo.FindByUsername(o.Username)
		if record != nil {
			return nil, errors.New("该用户名已存在！")
		}
	} else {
		o.IsUsernameSet = model.NO
		for {
			o.Username = strings.Random(10)
			record, _ := repository.UserRepo.FindByUsername(o.Username)
			if record == nil {
				break
			}
		}
	}

	o.Salt = strings.Random(8)
	o.Password = strings.Md5(o.Password + o.Salt)
	o.IsPasswordSet = model.YES

	if o.Status == 0 {
		o.Status = model.UserStatusPendingActivation
	}

	err := repository.UserRepo.Create(&o)
	return &o, err
}

func (svr *UserService) EditUser(o model.UserEditRequest) error {
	if o.ID == 0 {
		return errors.New("无效的ID")
	}

	user, err := repository.UserRepo.FindByID(o.ID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("用户不存在")
	}

	// 更新字段
	user.Nickname = o.Nickname
	user.Realname = o.Realname
	user.Engname = o.Engname
	user.Sex = o.Sex
	if o.Status != 0 {
		user.Status = o.Status
	}

	return repository.UserRepo.Update(user)
}

// UpdateOwnAccount updates only the personal fields of the current user.
// It does not allow changing status (admin-only field).
func (svr *UserService) UpdateOwnAccount(id uint, o model.AccountUpdateRequest) error {
	if id == 0 {
		return errors.New("无效的ID")
	}

	if o.Sex != 0 && o.Sex != 1 && o.Sex != 2 {
		return errors.New("性别值无效，请使用 1（男）或 2（女）")
	}

	user, err := repository.UserRepo.FindByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("用户不存在")
	}

	user.Realname = o.Realname
	user.Nickname = o.Nickname
	user.Engname = o.Engname
	user.Sex = o.Sex
	user.Mobile = o.Mobile

	return repository.UserRepo.Update(user)
}

func (svr *UserService) DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("无效的ID")
	}
	return repository.UserRepo.Delete(id)
}

// UpdateUserLoginStats updates login-related statistics for a user.
func (svr *UserService) UpdateUserLoginStats(userID uint, ip string) error {
	if userID == 0 {
		return errors.New("invalid user ID")
	}
	return repository.UserRepo.UpdateLoginStats(userID, ip)
}

// RevokeUserAdminRole removes admin role from a user (RBAC-based).
func (svr *UserService) RevokeUserAdminRole(userID uint) error {
	if userID == 0 {
		return errors.New("invalid user ID")
	}
	var errs []error
	// Remove "admin" role
	adminRole, err := repository.AdminRoleRepo.FindBySlug("admin")
	if err == nil && adminRole != nil {
		if removeErr := repository.UserRoleRepo.RemoveRole(userID, adminRole.ID); removeErr != nil {
			errs = append(errs, removeErr)
		}
	}
	// Also remove "super_admin" role
	superRole, err := repository.AdminRoleRepo.FindBySlug("super_admin")
	if err == nil && superRole != nil {
		if removeErr := repository.UserRoleRepo.RemoveRole(userID, superRole.ID); removeErr != nil {
			errs = append(errs, removeErr)
		}
	}
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}

// Grant one month VIP to user
func (svr *UserService) GrantVipMonth(userID uint) error {
	if userID == 0 {
		return errors.New("invalid user ID")
	}
	return repository.UserRepo.GrantVipMonth(userID)
}

// 修改密码并更新 tokenSalt
func (svr *UserService) UpdatePasswordAndSalt(userID uint, newPassword string, newSalt string) error {
	if userID == 0 || newPassword == "" || newSalt == "" {
		return errors.New("参数无效")
	}
	user, err := repository.UserRepo.FindByID(userID)
	if err != nil || user == nil {
		return errors.New("用户不存在")
	}
	hashedPwd := strings.Md5(newPassword + user.Salt)
	return repository.UserRepo.UpdatePasswordAndSalt(userID, hashedPwd, newSalt)
}
