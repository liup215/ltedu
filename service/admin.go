package service

import (
	"edu/model"
	"edu/repository"
	stdErrors "errors"

	"github.com/pkg/errors"
)

/*
// const (
// 	SUPERUSER_NAME     = "admin"
// 	SUPERUSER_PASSWORD = "123456"
// )

// func init() {
// 	// This init logic is problematic after Admin model refactor.
// 	// Superuser creation should be handled by the InitSuperUser endpoint or a dedicated setup script.
// 	// It would require creating a User first, then an Admin.
// 	// Commenting out for now.

// 	// _, err := AdminSvr.SelectAdminByUsername(SUPERUSER_NAME) // This method itself needs refactor
// 	// if err == gorm.ErrRecordNotFound {
// 	// 	// Create User first
// 	// 	// user := model.User{Username: SUPERUSER_NAME, Password: SUPERUSER_PASSWORD, ...}
// 	// 	// createdUser, _ := UserSvr.CreateUser(user)

// 	// 	// admin := model.Admin{
// 	// 	// 	UserID: createdUser.ID,
// 	// 	// 	Status:   model.ADMIN_STATUS_OK,
// 	// 	// }
// 	// 	// AdminSvr.CreateAdmin(admin)
// 	// }
// }
*/

var AdminSvr *AdminService = &AdminService{
	baseService: newBaseService(),
}

// SuperuserUsername is the reserved username that always has full system permissions.
const SuperuserUsername = "admin"

type AdminService struct {
	baseService
}

// CreateAdminRole creates a new admin role.
func (svr *AdminService) CreateAdminRole(role model.AdminRole) (*model.AdminRole, error) {
	// Ensure role.ID is 0 for creation.
	role.ID = 0
	// Check if role with the same slug already exists
	existingRole, err := repository.AdminRoleRepo.FindBySlug(role.Slug)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check for existing admin role by slug")
	}
	if existingRole != nil {
		return nil, stdErrors.New("admin role with this slug already exists")
	}

	err = repository.AdminRoleRepo.Create(&role)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create admin role")
	}
	return &role, nil
}

// UpdateAdminRole updates an existing admin role.
func (svr *AdminService) UpdateAdminRole(role model.AdminRole) (*model.AdminRole, error) {
	if role.ID == 0 {
		return nil, stdErrors.New("invalid admin role ID")
	}
	err := repository.AdminRoleRepo.Update(&role)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update admin role")
	}
	return &role, nil
}

// DeleteAdminRole deletes an admin role by ID.
func (svr *AdminService) DeleteAdminRole(id uint) error {
	if id == 0 {
		return stdErrors.New("invalid admin role ID")
	}
	return repository.AdminRoleRepo.Delete(id)
}

// SelectAdminRoleById retrieves an admin role by its ID.
func (svr *AdminService) SelectAdminRoleById(id uint) (*model.AdminRole, error) {
	if id == 0 {
		return nil, stdErrors.New("invalid admin role ID")
	}
	role, err := repository.AdminRoleRepo.FindByIDWithPermissions(id)
	if err != nil || role == nil {
		return nil, stdErrors.New("admin role not found")
	}
	return role, nil
}

// SelectAdminRoleBySlug retrieves an admin role by its slug.
func (svr *AdminService) SelectAdminRoleBySlug(slug string) (*model.AdminRole, error) {
	if slug == "" {
		return nil, stdErrors.New("slug cannot be empty")
	}
	role, err := repository.AdminRoleRepo.FindBySlug(slug)
	if err != nil || role == nil {
		return nil, stdErrors.New("admin role not found with this slug")
	}
	return role, nil
}

// ListAdminRoles returns all admin roles with their permissions.
func (svr *AdminService) ListAdminRoles() ([]*model.AdminRole, error) {
	return repository.AdminRoleRepo.FindAllWithPermissions()
}

// GetAdminRolePermissions retrieves all permissions associated with a given AdminRoleID.
func (svr *AdminService) GetAdminRolePermissions(roleID uint) ([]*model.AdminPermission, error) {
	return repository.AdminRoleRepo.GetPermissions(roleID)
}

// GetUserPermissions retrieves all permissions for a given userID.
func (svr *AdminService) GetUserPermissions(userID uint) ([]*model.AdminPermission, error) {
	user, err := UserSvr.SelectUserById(userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user for permissions check")
	}
	if user == nil {
		return nil, stdErrors.New("user not found for permissions check")
	}

	if !user.IsAdmin {
		return []*model.AdminPermission{}, nil // Not an admin, no admin permissions
	}

	// Superuser (e.g., username "admin") gets all permissions
	if user.Username == SuperuserUsername {
		permissions, err := repository.AdminPermRepo.FindAll()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get all permissions for superuser")
		}
		return permissions, nil
	}

	// Collect permissions from all RBAC roles assigned to this user
	roles, err := repository.UserRoleRepo.GetUserRolesWithPermissions(userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user roles")
	}

	permMap := make(map[uint]*model.AdminPermission)
	for _, role := range roles {
		for _, perm := range role.Permissions {
			permMap[perm.ID] = perm
		}
	}

	// Also check the legacy single AdminRoleID field for backward compatibility
	if user.AdminRoleID != nil && *user.AdminRoleID != 0 {
		legacyPerms, err := svr.GetAdminRolePermissions(*user.AdminRoleID)
		if err == nil {
			for _, perm := range legacyPerms {
				permMap[perm.ID] = perm
			}
		}
	}

	perms := make([]*model.AdminPermission, 0, len(permMap))
	for _, p := range permMap {
		perms = append(perms, p)
	}
	return perms, nil
}

// HasPermission checks if a user has a specific permission by slug.
func (svr *AdminService) HasPermission(userID uint, permSlug string) (bool, error) {
	perms, err := svr.GetUserPermissions(userID)
	if err != nil {
		return false, err
	}
	for _, p := range perms {
		if p.Slug == permSlug {
			return true, nil
		}
	}
	return false, nil
}

// ============ Permission CRUD ============

// CreatePermission creates a new permission.
func (svr *AdminService) CreatePermission(perm model.AdminPermission) (*model.AdminPermission, error) {
	perm.ID = 0
	existing, err := repository.AdminPermRepo.FindBySlug(perm.Slug)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check for existing permission")
	}
	if existing != nil {
		return nil, stdErrors.New("permission with this slug already exists")
	}
	if err := repository.AdminPermRepo.Create(&perm); err != nil {
		return nil, errors.Wrap(err, "failed to create permission")
	}
	return &perm, nil
}

// UpdatePermission updates an existing permission.
func (svr *AdminService) UpdatePermission(perm model.AdminPermission) (*model.AdminPermission, error) {
	if perm.ID == 0 {
		return nil, stdErrors.New("invalid permission ID")
	}
	if err := repository.AdminPermRepo.Update(&perm); err != nil {
		return nil, errors.Wrap(err, "failed to update permission")
	}
	return &perm, nil
}

// DeletePermission deletes a permission by ID.
func (svr *AdminService) DeletePermission(id uint) error {
	if id == 0 {
		return stdErrors.New("invalid permission ID")
	}
	return repository.AdminPermRepo.Delete(id)
}

// ListPermissions returns all permissions.
func (svr *AdminService) ListPermissions() ([]*model.AdminPermission, error) {
	return repository.AdminPermRepo.FindAll()
}

// ============ Role-Permission Assignment ============

// AssignPermissionToRole assigns a permission to a role.
func (svr *AdminService) AssignPermissionToRole(roleID, permissionID uint) error {
	return repository.AdminRoleRepo.AddPermission(roleID, permissionID)
}

// RemovePermissionFromRole removes a permission from a role.
func (svr *AdminService) RemovePermissionFromRole(roleID, permissionID uint) error {
	return repository.AdminRoleRepo.RemovePermission(roleID, permissionID)
}

// ============ User-Role Assignment ============

// AssignRoleToUser assigns a role to a user.
func (svr *AdminService) AssignRoleToUser(userID, roleID uint) error {
	return repository.UserRoleRepo.AssignRole(userID, roleID)
}

// RemoveRoleFromUser removes a role from a user.
func (svr *AdminService) RemoveRoleFromUser(userID, roleID uint) error {
	return repository.UserRoleRepo.RemoveRole(userID, roleID)
}

// GetUserRoles returns all roles assigned to a user.
func (svr *AdminService) GetUserRoles(userID uint) ([]*model.AdminRole, error) {
	return repository.UserRoleRepo.GetUserRoles(userID)
}

// ============ Default Seeding ============

// DefaultPermissions lists the canonical permissions for the system.
var DefaultPermissions = []model.AdminPermission{
	{Slug: "user:read", DisplayName: "Read Users", GroupName: "user"},
	{Slug: "user:create", DisplayName: "Create Users", GroupName: "user"},
	{Slug: "user:edit", DisplayName: "Edit Users", GroupName: "user"},
	{Slug: "user:delete", DisplayName: "Delete Users", GroupName: "user"},
	{Slug: "question:read", DisplayName: "Read Questions", GroupName: "question"},
	{Slug: "question:create", DisplayName: "Create Questions", GroupName: "question"},
	{Slug: "question:edit", DisplayName: "Edit Questions", GroupName: "question"},
	{Slug: "question:delete", DisplayName: "Delete Questions", GroupName: "question"},
	{Slug: "syllabus:read", DisplayName: "Read Syllabuses", GroupName: "syllabus"},
	{Slug: "syllabus:create", DisplayName: "Create Syllabuses", GroupName: "syllabus"},
	{Slug: "syllabus:edit", DisplayName: "Edit Syllabuses", GroupName: "syllabus"},
	{Slug: "syllabus:delete", DisplayName: "Delete Syllabuses", GroupName: "syllabus"},
	{Slug: "syllabus:manage", DisplayName: "Manage Syllabuses (full)", GroupName: "syllabus"},
	{Slug: "paper:read", DisplayName: "Read Papers", GroupName: "paper"},
	{Slug: "paper:create", DisplayName: "Create Papers", GroupName: "paper"},
	{Slug: "paper:edit", DisplayName: "Edit Papers", GroupName: "paper"},
	{Slug: "paper:delete", DisplayName: "Delete Papers", GroupName: "paper"},
	{Slug: "class:read", DisplayName: "Read Classes", GroupName: "class"},
	{Slug: "class:create", DisplayName: "Create Classes", GroupName: "class"},
	{Slug: "class:edit", DisplayName: "Edit Classes", GroupName: "class"},
	{Slug: "class:delete", DisplayName: "Delete Classes", GroupName: "class"},
	{Slug: "class:manage", DisplayName: "Manage Classes (full)", GroupName: "class"},
	{Slug: "role:read", DisplayName: "Read Roles", GroupName: "role"},
	{Slug: "role:create", DisplayName: "Create Roles", GroupName: "role"},
	{Slug: "role:edit", DisplayName: "Edit Roles", GroupName: "role"},
	{Slug: "role:delete", DisplayName: "Delete Roles", GroupName: "role"},
	{Slug: "permission:read", DisplayName: "Read Permissions", GroupName: "permission"},
	{Slug: "permission:create", DisplayName: "Create Permissions", GroupName: "permission"},
	{Slug: "permission:edit", DisplayName: "Edit Permissions", GroupName: "permission"},
	{Slug: "permission:delete", DisplayName: "Delete Permissions", GroupName: "permission"},
}

// allDefaultPermissionSlugs returns all slugs from DefaultPermissions.
func allDefaultPermissionSlugs() []string {
	slugs := make([]string, len(DefaultPermissions))
	for i, p := range DefaultPermissions {
		slugs[i] = p.Slug
	}
	return slugs
}

// DefaultRoles defines the standard roles and their permission slugs.
var DefaultRoles = []struct {
	Slug        string
	DisplayName string
	Description string
	Permissions []string
}{
	{
		Slug:        "super_admin",
		DisplayName: "Super Admin",
		Description: "Full access to all resources",
		Permissions: allDefaultPermissionSlugs(),
	},
	{
		Slug:        "admin",
		DisplayName: "Admin",
		Description: "Administrative access",
		Permissions: []string{
			"user:read", "user:edit",
			"question:read", "question:create", "question:edit", "question:delete",
			"syllabus:read", "syllabus:manage",
			"paper:read", "paper:create", "paper:edit", "paper:delete",
			"class:read", "class:manage",
			"role:read", "permission:read",
		},
	},
	{
		Slug:        "teacher",
		DisplayName: "Teacher",
		Description: "Teacher access: create and manage own content",
		Permissions: []string{
			"question:read", "question:create", "question:edit",
			"paper:read", "paper:create", "paper:edit",
			"class:read", "class:manage",
			"syllabus:read",
		},
	},
}

// SeedDefaultRolesAndPermissions seeds the database with default roles and permissions.
// It is idempotent: existing records are not duplicated.
func (svr *AdminService) SeedDefaultRolesAndPermissions() error {
	// Ensure all default permissions exist
	for _, dp := range DefaultPermissions {
		existing, err := repository.AdminPermRepo.FindBySlug(dp.Slug)
		if err != nil {
			return errors.Wrapf(err, "failed to check permission %s", dp.Slug)
		}
		if existing == nil {
			p := dp
			if err := repository.AdminPermRepo.Create(&p); err != nil {
				return errors.Wrapf(err, "failed to create permission %s", dp.Slug)
			}
		}
	}

	// Ensure all default roles exist with their permissions
	for _, dr := range DefaultRoles {
		role, err := repository.AdminRoleRepo.FindBySlug(dr.Slug)
		if err != nil {
			return errors.Wrapf(err, "failed to check role %s", dr.Slug)
		}
		if role == nil {
			newRole := model.AdminRole{
				Slug:        dr.Slug,
				DisplayName: dr.DisplayName,
				Description: dr.Description,
			}
			if err := repository.AdminRoleRepo.Create(&newRole); err != nil {
				return errors.Wrapf(err, "failed to create role %s", dr.Slug)
			}
			role = &newRole
		}
		// Assign permissions to role
		for _, permSlug := range dr.Permissions {
			perm, err := repository.AdminPermRepo.FindBySlug(permSlug)
			if err != nil || perm == nil {
				continue
			}
			// Try to add; ignore duplicate-assignment errors silently
			_ = repository.AdminRoleRepo.AddPermission(role.ID, perm.ID)
		}
	}
	return nil
}

func (svr *AdminService) AdminMenu() []model.Menu {
	return model.AdminMenu
}

func (svr *AdminService) AdminPermission() []string {
	return []string{"sys:user:add", "sys:user:edit", "sys:user:delete", "sys:user:import", "sys:user:export"}
}
