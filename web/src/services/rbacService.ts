import apiClient from './apiClient'

export interface Permission {
  id: number
  displayName: string
  slug: string
  description?: string
  method?: string
  url?: string
  groupName?: string
}

export interface Role {
  id: number
  displayName: string
  slug: string
  description?: string
  permissions?: Permission[]
}

export interface AssignPermissionRequest {
  roleId: number
  permissionId: number
}

export interface AssignRoleRequest {
  userId: number
  roleId: number
}

const rbacService = {
  // ============ Roles ============

  listRoles(): Promise<any> {
    return apiClient.post('/v1/rbac/roles/list', {})
  },

  getRoleById(id: number): Promise<any> {
    return apiClient.post('/v1/rbac/roles/byId', { id })
  },

  createRole(role: Partial<Role>): Promise<any> {
    return apiClient.post('/v1/rbac/roles/create', role)
  },

  updateRole(role: Partial<Role>): Promise<any> {
    return apiClient.post('/v1/rbac/roles/edit', role)
  },

  deleteRole(id: number): Promise<any> {
    return apiClient.post('/v1/rbac/roles/delete', { id })
  },

  // ============ Permissions ============

  listPermissions(): Promise<any> {
    return apiClient.post('/v1/rbac/permissions/list', {})
  },

  createPermission(perm: Partial<Permission>): Promise<any> {
    return apiClient.post('/v1/rbac/permissions/create', perm)
  },

  updatePermission(perm: Partial<Permission>): Promise<any> {
    return apiClient.post('/v1/rbac/permissions/edit', perm)
  },

  deletePermission(id: number): Promise<any> {
    return apiClient.post('/v1/rbac/permissions/delete', { id })
  },

  // ============ Role-Permission Assignment ============

  assignPermissionToRole(roleId: number, permissionId: number): Promise<any> {
    return apiClient.post('/v1/rbac/roles/permission/assign', { roleId, permissionId })
  },

  removePermissionFromRole(roleId: number, permissionId: number): Promise<any> {
    return apiClient.post('/v1/rbac/roles/permission/remove', { roleId, permissionId })
  },

  // ============ User-Role Assignment ============

  getUserRoles(userId: number): Promise<any> {
    return apiClient.post('/v1/rbac/user/roles/list', { userId })
  },

  assignRoleToUser(userId: number, roleId: number): Promise<any> {
    return apiClient.post('/v1/rbac/user/roles/assign', { userId, roleId })
  },

  removeRoleFromUser(userId: number, roleId: number): Promise<any> {
    return apiClient.post('/v1/rbac/user/roles/remove', { userId, roleId })
  },

  // ============ Current User ============

  getMyPermissions(): Promise<any> {
    return apiClient.post('/v1/rbac/me/permissions', {})
  },

  checkPermission(permission: string): Promise<any> {
    return apiClient.post('/v1/rbac/me/check-permission', { permission })
  }
}

export default rbacService
