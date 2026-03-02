package repository

import (
	"edu/model"
	stdErrors "errors"
	"gorm.io/gorm"
)

// IAdminRoleRepository 管理员角色数据访问接口
type IAdminRoleRepository interface {
	Create(role *model.AdminRole) error
	Update(role *model.AdminRole) error
	Delete(id uint) error
	FindByID(id uint) (*model.AdminRole, error)
	FindByIDWithPermissions(id uint) (*model.AdminRole, error)
	FindBySlug(slug string) (*model.AdminRole, error)
	FindAll() ([]*model.AdminRole, error)
	FindAllWithPermissions() ([]*model.AdminRole, error)
	AddPermission(roleID, permissionID uint) error
	RemovePermission(roleID, permissionID uint) error
	GetPermissions(roleID uint) ([]*model.AdminPermission, error)
}

// IAdminPermissionRepository 管理员权限数据访问接口
type IAdminPermissionRepository interface {
	Create(permission *model.AdminPermission) error
	Update(permission *model.AdminPermission) error
	Delete(id uint) error
	FindByID(id uint) (*model.AdminPermission, error)
	FindBySlug(slug string) (*model.AdminPermission, error)
	FindAll() ([]*model.AdminPermission, error)
}

// IUserRoleRepository 用户-角色关联数据访问接口
type IUserRoleRepository interface {
	AssignRole(userID, roleID uint) error
	RemoveRole(userID, roleID uint) error
	GetUserRoles(userID uint) ([]*model.AdminRole, error)
	GetUserRolesWithPermissions(userID uint) ([]*model.AdminRole, error)
	HasRole(userID, roleID uint) (bool, error)
}

type adminRoleRepository struct {
	db *gorm.DB
}

type adminPermissionRepository struct {
	db *gorm.DB
}

type userRoleRepository struct {
	db *gorm.DB
}

// NewAdminRoleRepository 创建管理员角色仓储实例
func NewAdminRoleRepository(db *gorm.DB) IAdminRoleRepository {
	return &adminRoleRepository{db: db}
}

// NewAdminPermissionRepository 创建管理员权限仓储实例
func NewAdminPermissionRepository(db *gorm.DB) IAdminPermissionRepository {
	return &adminPermissionRepository{db: db}
}

// NewUserRoleRepository 创建用户-角色关联仓储实例
func NewUserRoleRepository(db *gorm.DB) IUserRoleRepository {
	return &userRoleRepository{db: db}
}

// ============ AdminRole 实现 ============

// Create 创建管理员角色
func (r *adminRoleRepository) Create(role *model.AdminRole) error {
	return r.db.Create(role).Error
}

// Update 更新管理员角色
func (r *adminRoleRepository) Update(role *model.AdminRole) error {
	return r.db.Save(role).Error
}

// Delete 删除管理员角色
func (r *adminRoleRepository) Delete(id uint) error {
	return r.db.Delete(&model.AdminRole{}, id).Error
}

// FindByID 根据ID查询管理员角色
func (r *adminRoleRepository) FindByID(id uint) (*model.AdminRole, error) {
	var role model.AdminRole
	err := r.db.Where("id = ?", id).First(&role).Error

	if stdErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &role, err
}

// FindBySlug 根据slug查询管理员角色
func (r *adminRoleRepository) FindBySlug(slug string) (*model.AdminRole, error) {
	var role model.AdminRole
	err := r.db.Where("slug = ?", slug).First(&role).Error

	if stdErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &role, err
}

// FindAll 查询所有管理员角色
func (r *adminRoleRepository) FindAll() ([]*model.AdminRole, error) {
	var roles []*model.AdminRole
	err := r.db.Find(&roles).Error
	return roles, err
}

// FindByIDWithPermissions 根据ID查询角色（包含权限）
func (r *adminRoleRepository) FindByIDWithPermissions(id uint) (*model.AdminRole, error) {
	var role model.AdminRole
	err := r.db.Preload("Permissions").Where("id = ?", id).First(&role).Error
	if stdErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &role, err
}

// FindAllWithPermissions 查询所有管理员角色（包含权限）
func (r *adminRoleRepository) FindAllWithPermissions() ([]*model.AdminRole, error) {
	var roles []*model.AdminRole
	err := r.db.Preload("Permissions").Find(&roles).Error
	return roles, err
}

// AddPermission 给角色添加权限
func (r *adminRoleRepository) AddPermission(roleID, permissionID uint) error {
	role := model.AdminRole{}
	role.ID = roleID
	perm := model.AdminPermission{}
	perm.ID = permissionID
	return r.db.Model(&role).Association("Permissions").Append(&perm)
}

// RemovePermission 从角色移除权限
func (r *adminRoleRepository) RemovePermission(roleID, permissionID uint) error {
	role := model.AdminRole{}
	role.ID = roleID
	perm := model.AdminPermission{}
	perm.ID = permissionID
	return r.db.Model(&role).Association("Permissions").Delete(&perm)
}

// GetPermissions 获取角色的所有权限
func (r *adminRoleRepository) GetPermissions(roleID uint) ([]*model.AdminPermission, error) {
	role, err := r.FindByIDWithPermissions(roleID)
	if err != nil || role == nil {
		return nil, err
	}
	return role.Permissions, nil
}

// ============ AdminPermission 实现 ============

// Create 创建管理员权限
func (r *adminPermissionRepository) Create(permission *model.AdminPermission) error {
	return r.db.Create(permission).Error
}

// Update 更新管理员权限
func (r *adminPermissionRepository) Update(permission *model.AdminPermission) error {
	return r.db.Save(permission).Error
}

// Delete 删除管理员权限
func (r *adminPermissionRepository) Delete(id uint) error {
	return r.db.Delete(&model.AdminPermission{}, id).Error
}

// FindByID 根据ID查询管理员权限
func (r *adminPermissionRepository) FindByID(id uint) (*model.AdminPermission, error) {
	var permission model.AdminPermission
	err := r.db.Where("id = ?", id).First(&permission).Error

	if stdErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &permission, err
}

// FindAll 查询所有管理员权限
func (r *adminPermissionRepository) FindAll() ([]*model.AdminPermission, error) {
	var permissions []*model.AdminPermission
	err := r.db.Find(&permissions).Error
	return permissions, err
}

// FindBySlug 根据slug查询管理员权限
func (r *adminPermissionRepository) FindBySlug(slug string) (*model.AdminPermission, error) {
	var permission model.AdminPermission
	err := r.db.Where("slug = ?", slug).First(&permission).Error
	if stdErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &permission, err
}

// ============ UserRole 实现 ============

// AssignRole 给用户分配角色
func (r *userRoleRepository) AssignRole(userID, roleID uint) error {
	user := model.User{}
	user.ID = userID
	role := model.AdminRole{}
	role.ID = roleID
	return r.db.Model(&user).Association("Roles").Append(&role)
}

// RemoveRole 从用户移除角色
func (r *userRoleRepository) RemoveRole(userID, roleID uint) error {
	user := model.User{}
	user.ID = userID
	role := model.AdminRole{}
	role.ID = roleID
	return r.db.Model(&user).Association("Roles").Delete(&role)
}

// GetUserRoles 获取用户的所有角色
func (r *userRoleRepository) GetUserRoles(userID uint) ([]*model.AdminRole, error) {
	var user model.User
	err := r.db.Preload("Roles").Where("id = ?", userID).First(&user).Error
	if stdErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user.Roles, nil
}

// GetUserRolesWithPermissions 获取用户的所有角色（含权限）
func (r *userRoleRepository) GetUserRolesWithPermissions(userID uint) ([]*model.AdminRole, error) {
	var user model.User
	err := r.db.Preload("Roles.Permissions").Where("id = ?", userID).First(&user).Error
	if stdErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user.Roles, nil
}

// HasRole 检查用户是否有指定角色
func (r *userRoleRepository) HasRole(userID, roleID uint) (bool, error) {
	roles, err := r.GetUserRoles(userID)
	if err != nil {
		return false, err
	}
	for _, role := range roles {
		if role.ID == roleID {
			return true, nil
		}
	}
	return false, nil
}
