package v1

import (
	// Use the project's http lib
	"edu/model"
	"edu/service"

	// Import standard library errors

	"github.com/pkg/errors"
)

// initializeDefaultAdmin initializes the super admin role and admin user if they don't exist
func EnsureSuperUserExists() error {
	const (
		superUserRoleSlug = "super-admin"
		adminUsername     = "admin"
		adminPassword     = "123456" // Default password
		adminEmail        = "admin@ltedu.com"
		adminMobile       = "13800000000"
	)

	// Check if admin user exists
	adminUser, err := service.UserSvr.SelectUserByUsername(adminUsername)
	if err == nil && adminUser != nil {
		// Admin user already exists, nothing to do
		return nil
	}

	// Ensure Superuser AdminRole exists or create it
	superAdminRole, err := service.AdminSvr.SelectAdminRoleBySlug(superUserRoleSlug)
	if err != nil {
		// Create super admin role if it doesn't exist
		if err.Error() == "admin role not found with this slug" {
			newRole := model.AdminRole{
				DisplayName: "Super Administrator",
				Slug:        superUserRoleSlug,
				Description: "System superuser with all permissions.",
			}
			createdRole, createRoleErr := service.AdminSvr.CreateAdminRole(newRole)
			if createRoleErr != nil {
				return errors.Wrap(createRoleErr, "failed to create superuser role")
			}
			superAdminRole = createdRole
		} else {
			return errors.Wrap(err, "failed to get superuser role")
		}
	}

	// Create admin user
	user := &model.User{
		Username: adminUsername,
		Password: adminPassword, // Will be hashed by service layer
		Email:    adminEmail,
		Mobile:   adminMobile,
		Status:   model.UserStatusNormal,
	}

	createdUser, err := service.UserSvr.CreateUser(*user)
	if err != nil {
		return errors.Wrap(err, "failed to create admin user")
	}

	// Assign super_admin role to the user via RBAC
	if err := service.AdminSvr.AssignRoleToUser(createdUser.ID, superAdminRole.ID); err != nil {
		return errors.Wrap(err, "failed to assign super_admin role to admin user")
	}

	return nil
}
