<template>
  <div class="p-6">
    <header class="mb-6 flex justify-between items-center">
      <div>
        <h1 class="text-3xl font-bold text-gray-900">{{ $t('rbac.rolesTitle') }}</h1>
        <p class="mt-1 text-sm text-gray-500">{{ $t('rbac.rolesDescription') }}</p>
      </div>
      <button
        @click="openCreateModal"
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none"
      >
        {{ $t('rbac.createRole') }}
      </button>
    </header>

    <div class="bg-white shadow overflow-x-auto sm:rounded-lg mb-8">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('rbac.displayName') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Slug</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('rbac.description') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('rbac.permissions') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('common.actions') }}</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-if="loading">
            <td colspan="6" class="px-6 py-4 text-center text-sm text-gray-500">{{ $t('common.loading') }}</td>
          </tr>
          <tr v-else-if="roles.length === 0">
            <td colspan="6" class="px-6 py-4 text-center text-sm text-gray-500">{{ $t('common.noData') }}</td>
          </tr>
          <tr v-for="role in roles" :key="role.id">
            <td class="px-6 py-4 text-sm text-gray-900">{{ role.id }}</td>
            <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ role.displayName }}</td>
            <td class="px-6 py-4 text-sm text-gray-500 font-mono">{{ role.slug }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ role.description }}</td>
            <td class="px-6 py-4 text-sm text-gray-500">
              <div class="flex flex-wrap gap-1">
                <span
                  v-for="perm in (role.permissions || []).slice(0, 5)"
                  :key="perm.id"
                  class="px-2 py-0.5 text-xs rounded-full bg-indigo-100 text-indigo-700 font-mono"
                >{{ perm.slug }}</span>
                <span
                  v-if="(role.permissions || []).length > 5"
                  class="px-2 py-0.5 text-xs rounded-full bg-gray-100 text-gray-500"
                >+{{ (role.permissions || []).length - 5 }} more</span>
              </div>
            </td>
            <td class="px-6 py-4 text-sm space-x-2">
              <button @click="openEditModal(role)" class="text-indigo-600 hover:text-indigo-900">{{ $t('common.edit') }}</button>
              <button @click="openPermissionsModal(role)" class="text-green-600 hover:text-green-900">{{ $t('rbac.managePermissions') }}</button>
              <button @click="confirmDeleteRole(role.id)" class="text-red-600 hover:text-red-900">{{ $t('common.delete') }}</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Create/Edit Role Modal -->
    <div v-if="showRoleModal" class="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center z-10">
      <div class="bg-white rounded-lg p-6 w-full max-w-md shadow-xl">
        <h2 class="text-lg font-semibold mb-4">{{ editingRole?.id ? $t('rbac.editRole') : $t('rbac.createRole') }}</h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">{{ $t('rbac.displayName') }}</label>
            <input v-model="roleForm.displayName" type="text" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Slug</label>
            <input v-model="roleForm.slug" type="text" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 text-sm font-mono focus:outline-none focus:ring-indigo-500 focus:border-indigo-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">{{ $t('rbac.description') }}</label>
            <textarea v-model="roleForm.description" rows="2" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"></textarea>
          </div>
        </div>
        <div class="mt-6 flex justify-end space-x-3">
          <button @click="showRoleModal = false" class="px-4 py-2 border border-gray-300 rounded-md text-sm text-gray-700 hover:bg-gray-50">{{ $t('common.cancel') }}</button>
          <button @click="saveRole" class="px-4 py-2 bg-indigo-600 text-white rounded-md text-sm hover:bg-indigo-700">{{ $t('common.save') }}</button>
        </div>
      </div>
    </div>

    <!-- Permissions Assignment Modal -->
    <div v-if="showPermissionsModal" class="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center z-10">
      <div class="bg-white rounded-lg p-6 w-full max-w-2xl shadow-xl max-h-screen overflow-y-auto">
        <h2 class="text-lg font-semibold mb-2">{{ $t('rbac.managePermissionsFor') }}: <span class="text-indigo-600">{{ selectedRole?.displayName }}</span></h2>
        <p class="text-sm text-gray-500 mb-4">{{ $t('rbac.managePermissionsDescription') }}</p>

        <div class="grid grid-cols-1 gap-2 max-h-96 overflow-y-auto border border-gray-200 rounded p-3">
          <div v-for="group in permissionGroups" :key="group.name" class="mb-3">
            <h3 class="text-xs font-semibold text-gray-500 uppercase tracking-wider mb-1">{{ group.name }}</h3>
            <div class="grid grid-cols-2 gap-1">
              <label
                v-for="perm in group.permissions"
                :key="perm.id"
                class="flex items-center space-x-2 text-sm cursor-pointer hover:bg-gray-50 p-1 rounded"
              >
                <input
                  type="checkbox"
                  :checked="isPermissionAssigned(perm.id)"
                  @change="togglePermission(perm)"
                  class="h-4 w-4 text-indigo-600 border-gray-300 rounded"
                />
                <span class="font-mono text-xs text-gray-700">{{ perm.slug }}</span>
              </label>
            </div>
          </div>
        </div>

        <div class="mt-4 flex justify-end">
          <button @click="showPermissionsModal = false" class="px-4 py-2 bg-gray-100 text-gray-700 rounded-md text-sm hover:bg-gray-200">{{ $t('common.cancel') }}</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import rbacService from '../../services/rbacService'
import type { Role, Permission } from '../../services/rbacService'

const { t } = useI18n()

const roles = ref<Role[]>([])
const allPermissions = ref<Permission[]>([])
const loading = ref(true)

const showRoleModal = ref(false)
const showPermissionsModal = ref(false)
const editingRole = ref<Role | null>(null)
const selectedRole = ref<Role | null>(null)
const assignedPermissionIds = ref<Set<number>>(new Set())

const roleForm = ref({ displayName: '', slug: '', description: '' })

const fetchRoles = async () => {
  loading.value = true
  try {
    const res = await rbacService.listRoles()
    roles.value = res.data.list || []
  } catch (e) {
    console.error('Failed to fetch roles', e)
  } finally {
    loading.value = false
  }
}

const fetchPermissions = async () => {
  try {
    const res = await rbacService.listPermissions()
    allPermissions.value = res.data.list || []
  } catch (e) {
    console.error('Failed to fetch permissions', e)
  }
}

onMounted(() => {
  fetchRoles()
  fetchPermissions()
})

const permissionGroups = computed(() => {
  const groups: Record<string, { name: string; permissions: Permission[] }> = {}
  for (const perm of allPermissions.value) {
    const groupName = perm.groupName || 'other'
    if (!groups[groupName]) {
      groups[groupName] = { name: groupName, permissions: [] }
    }
    groups[groupName].permissions.push(perm)
  }
  return Object.values(groups)
})

const openCreateModal = () => {
  editingRole.value = null
  roleForm.value = { displayName: '', slug: '', description: '' }
  showRoleModal.value = true
}

const openEditModal = (role: Role) => {
  editingRole.value = role
  roleForm.value = { displayName: role.displayName, slug: role.slug, description: role.description || '' }
  showRoleModal.value = true
}

const saveRole = async () => {
  try {
    if (editingRole.value?.id) {
      await rbacService.updateRole({ ...roleForm.value, id: editingRole.value.id })
    } else {
      await rbacService.createRole(roleForm.value)
    }
    showRoleModal.value = false
    await fetchRoles()
  } catch (e) {
    console.error('Failed to save role', e)
  }
}

const confirmDeleteRole = async (id: number) => {
  if (confirm(t('rbac.confirmDeleteRole'))) {
    try {
      await rbacService.deleteRole(id)
      await fetchRoles()
    } catch (e) {
      console.error('Failed to delete role', e)
    }
  }
}

const openPermissionsModal = (role: Role) => {
  selectedRole.value = role
  assignedPermissionIds.value = new Set((role.permissions || []).map(p => p.id))
  showPermissionsModal.value = true
}

const isPermissionAssigned = (permId: number) => assignedPermissionIds.value.has(permId)

const togglePermission = async (perm: Permission) => {
  if (!selectedRole.value) return
  try {
    if (isPermissionAssigned(perm.id)) {
      await rbacService.removePermissionFromRole(selectedRole.value.id, perm.id)
      assignedPermissionIds.value.delete(perm.id)
    } else {
      await rbacService.assignPermissionToRole(selectedRole.value.id, perm.id)
      assignedPermissionIds.value.add(perm.id)
    }
    // Refresh roles list silently to update the permission chips
    await fetchRoles()
    // Update selected role's permissions
    const updatedRole = roles.value.find(r => r.id === selectedRole.value?.id)
    if (updatedRole) {
      selectedRole.value = updatedRole
    }
  } catch (e) {
    console.error('Failed to toggle permission', e)
  }
}
</script>
