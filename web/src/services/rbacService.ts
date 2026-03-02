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

  async listRoles(): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/roles/list', {})
  },

  async getRoleById(id: number): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/roles/byId', { id })
  },

  async createRole(role: Partial<Role>): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/roles/create', role)
  },

  async updateRole(role: Partial<Role>): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/roles/edit', role)
  },

  async deleteRole(id: number): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/roles/delete', { id })
  },

  // ============ Permissions ============

  async listPermissions(): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/permissions/list', {})
  },

  async createPermission(perm: Partial<Permission>): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/permissions/create', perm)
  },

  async updatePermission(perm: Partial<Permission>): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/permissions/edit', perm)
  },

  async deletePermission(id: number): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/permissions/delete', { id })
  },

  // ============ Role-Permission Assignment ============

  async assignPermissionToRole(roleId: number, permissionId: number): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/roles/permission/assign', { roleId, permissionId })
  },

  async removePermissionFromRole(roleId: number, permissionId: number): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/roles/permission/remove', { roleId, permissionId })
  },

  // ============ User-Role Assignment ============

  async getUserRoles(userId: number): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/user/roles/list', { userId })
  },

  async assignRoleToUser(userId: number, roleId: number): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/user/roles/assign', { userId, roleId })
  },

  async removeRoleFromUser(userId: number, roleId: number): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/user/roles/remove', { userId, roleId })
  },

  // ============ Current User ============

  async getMyPermissions(): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/me/permissions', {})
  },

  async checkPermission(permission: string): Promise<any> {
    const client = await apiClient()
    return client.post('/api/v1/rbac/me/check-permission', { permission })
  }
}

export default rbacService
