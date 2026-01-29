package repository

import (
	"errors"
	"time"

	"edu/model"
	"gorm.io/gorm"
)

// IUserRepository 用户数据访问接口
type IUserRepository interface {
	// 基础CRUD
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uint) error
	FindByID(id uint) (*model.User, error)

	// 唯一查询
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByMobile(mobile string) (*model.User, error)

	// 列表查询
	FindList(query model.UserQuery) ([]*model.User, int64, error)
	FindAll(query model.UserQuery) ([]*model.User, error)
	Count() (int64, error)

	// 批量操作
	BatchCreate(users []*model.User) error
	BatchDelete(ids []uint) error

	// 业务相关查询
	FindByAdminStatus(status int) ([]*model.User, error)
	FindVipUsers() ([]*model.User, error)
	UpdateLoginStats(userID uint, ip string) error
	GrantVipMonth(userID uint) error
	UpdatePasswordAndSalt(userID uint, password, salt string) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建用户仓储实例
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

// Create 创建用户
func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// Update 更新用户
func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete 删除用户
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

// FindByID 根据ID查询用户
func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).
		Preload("Classes").
		Preload("AdminRole").
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// FindByUsername 根据用户名查询
func (r *userRepository) FindByUsername(username string) (*model.User, error) {
var user model.User
err := r.db.Where("username = ?", username).
Preload("Classes").
Preload("AdminRole").
First(&user).Error

if errors.Is(err, gorm.ErrRecordNotFound) {
return nil, nil
}
return &user, err
}

// FindByEmail 根据邮箱查询
func (r *userRepository) FindByEmail(email string) (*model.User, error) {
var user model.User
err := r.db.Where("email = ?", email).
Preload("Classes").
Preload("AdminRole").
First(&user).Error

if errors.Is(err, gorm.ErrRecordNotFound) {
return nil, nil
}
return &user, err
}

// FindByMobile 根据手机号查询
func (r *userRepository) FindByMobile(mobile string) (*model.User, error) {
var user model.User
err := r.db.Where("mobile = ?", mobile).
Preload("Classes").
Preload("AdminRole").
First(&user).Error

if errors.Is(err, gorm.ErrRecordNotFound) {
return nil, nil
}
return &user, err
}

// FindList 分页查询
func (r *userRepository) FindList(query model.UserQuery) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	q := r.buildQuery(query)

	// 统计总数
	q = q.Model(&model.User{})
	q.Count(&total)

	// 分页查询
	page := query.Page.CheckPage()
	err := q.
		Order("id DESC").
		Offset((page.PageIndex - 1) * page.PageSize).
		Limit(page.PageSize).
		Preload("Classes").
		Preload("AdminRole").
		Find(&users).Error

	return users, total, err
}

// FindAll 查询所有符合条件的用户
func (r *userRepository) FindAll(query model.UserQuery) ([]*model.User, error) {
	var users []*model.User
	err := r.buildQuery(query).
		Order("id DESC").
		Preload("Classes").
		Preload("AdminRole").
		Find(&users).Error

	return users, err
}

// Count 统计用户总数
func (r *userRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&model.User{}).Count(&count).Error
	return count, err
}

// BatchCreate 批量创建用户
func (r *userRepository) BatchCreate(users []*model.User) error {
	if len(users) == 0 {
		return nil
	}
	return r.db.Create(&users).Error
}

// BatchDelete 批量删除用户
func (r *userRepository) BatchDelete(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}
	return r.db.Delete(&model.User{}, ids).Error
}

// FindByAdminStatus 根据管理员状态查询
func (r *userRepository) FindByAdminStatus(status int) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Where("is_admin = ? AND admin_status = ?", true, status).
		Preload("AdminRole").
		Find(&users).Error
	return users, err
}

// FindVipUsers 查询VIP用户
func (r *userRepository) FindVipUsers() ([]*model.User, error) {
now := time.Now()
var users []*model.User
err := r.db.Where("vip_expire_at > ?", now).
Preload("Classes").
Preload("AdminRole").
Find(&users).Error
return users, err
}

// UpdateLoginStats 更新登录统计
func (r *userRepository) UpdateLoginStats(userID uint, ip string) error {
	now := time.Now()
	return r.db.Model(&model.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"last_login_ip":   ip,
			"last_login_date": &now,
			"login_times":     gorm.Expr("login_times + ?", 1),
		}).Error
}

// GrantVipMonth 授予一个月VIP
func (r *userRepository) GrantVipMonth(userID uint) error {
	var user model.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}

	now := time.Now()
	var newExpire time.Time
	if user.VipExpireAt != nil && user.VipExpireAt.After(now) {
		newExpire = user.VipExpireAt.AddDate(0, 1, 0)
	} else {
		newExpire = now.AddDate(0, 1, 0)
	}

	return r.db.Model(&model.User{}).
		Where("id = ?", userID).
		Update("vip_expire_at", &newExpire).Error
}

// UpdatePasswordAndSalt 更新密码和salt
func (r *userRepository) UpdatePasswordAndSalt(userID uint, password, salt string) error {
	return r.db.Model(&model.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"password":   password,
			"token_salt": salt,
		}).Error
}

// buildQuery 构建查询条件
func (r *userRepository) buildQuery(query model.UserQuery) *gorm.DB {
	q := r.db.Model(&model.User{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.Username != "" {
		q = q.Where("username = ?", query.Username)
	}
	if query.Realname != "" {
		q = q.Where("realname = ?", query.Realname)
	}
	if query.Engname != "" {
		q = q.Where("engname = ?", query.Engname)
	}
	if query.Mobile != "" {
		q = q.Where("mobile = ?", query.Mobile)
	}
	if query.Status != 0 {
		q = q.Where("status = ?", query.Status)
	}
	if query.ClassId != 0 {
		q = q.Joins("JOIN user_class_relation ON user_class_relation.user_id = users.id").
			Where("user_class_relation.class_id = ?", query.ClassId)
	}

	return q
}
