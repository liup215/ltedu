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

// SetAdmin: 设置用户为管理员（分配默认角色和状态）
func (svr *UserService) SetAdmin(userID uint) error {
	if userID == 0 {
		return errors.New("无效的ID")
	}
	// 默认角色ID和状态，可根据实际调整
	const defaultAdminRoleID = 1
	const defaultAdminStatus = model.ADMIN_STATUS_OK
	return svr.AssignUserAdminRole(userID, defaultAdminRoleID, defaultAdminStatus)
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
	user.Status = o.Status

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

// SetUserAdminStatus updates the admin-specific status of a user.
// It does not change IsAdmin or AdminRoleID.
func (svr *UserService) SetUserAdminStatus(userID uint, adminStatus int) error {
	if userID == 0 {
		return errors.New("invalid user ID")
	}

	user, err := repository.UserRepo.FindByID(userID)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	user.AdminStatus = &adminStatus
	return repository.UserRepo.Update(user)
}

// AssignUserAdminRole makes a user an admin, assigns a role, and sets their admin status.
func (svr *UserService) AssignUserAdminRole(userID uint, adminRoleID uint, adminStatus int) error {
	if userID == 0 {
		return errors.New("invalid user ID")
	}
	if adminRoleID == 0 {
		return errors.New("invalid admin role ID for assignment")
	}

	// Check if user exists
	user, err := repository.UserRepo.FindByID(userID)
	if err != nil || user == nil {
		return errors.New("user not found for admin assignment")
	}

	// Check if admin role exists
	// TODO: 需要实现 AdminSvr.SelectAdminRoleById 或者 AdminRepository
	// _, err = AdminSvr.SelectAdminRoleById(adminRoleID)
	// if err != nil {
	// 	return errors.New("admin role not found for assignment")
	// }

	// Update user admin info
	user.IsAdmin = true
	user.AdminRoleID = &adminRoleID
	user.AdminStatus = &adminStatus

	return repository.UserRepo.Update(user)
}

// RevokeUserAdminRole removes admin privileges from a user.
func (svr *UserService) RevokeUserAdminRole(userID uint) error {
	if userID == 0 {
		return errors.New("invalid user ID")
	}

	user, err := repository.UserRepo.FindByID(userID)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	user.IsAdmin = false
	user.AdminRoleID = nil
	user.AdminStatus = nil

	return repository.UserRepo.Update(user)
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
