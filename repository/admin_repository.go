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
	FindBySlug(slug string) (*model.AdminRole, error)
	FindAll() ([]*model.AdminRole, error)
}

// IAdminPermissionRepository 管理员权限数据访问接口
type IAdminPermissionRepository interface {
	Create(permission *model.AdminPermission) error
	Update(permission *model.AdminPermission) error
	Delete(id uint) error
	FindByID(id uint) (*model.AdminPermission, error)
	FindAll() ([]*model.AdminPermission, error)
}

type adminRoleRepository struct {
	db *gorm.DB
}

type adminPermissionRepository struct {
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
