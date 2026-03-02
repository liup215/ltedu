<template>
  <div class="p-6">
    <header class="mb-6 flex justify-between items-center">
      <div>
        <h1 class="text-3xl font-bold text-gray-900">{{ $t('rbac.permissionsTitle') }}</h1>
        <p class="mt-1 text-sm text-gray-500">{{ $t('rbac.permissionsDescription') }}</p>
      </div>
      <button
        @click="openCreateModal"
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none"
      >
        {{ $t('rbac.createPermission') }}
      </button>
    </header>

    <!-- Group filter -->
    <div class="mb-4 flex flex-wrap gap-2">
      <button
        @click="selectedGroup = ''"
        :class="['px-3 py-1 text-sm rounded-full border', selectedGroup === '' ? 'bg-indigo-600 text-white border-indigo-600' : 'bg-white text-gray-700 border-gray-300 hover:bg-gray-50']"
      >
        All
      </button>
      <button
        v-for="group in groups"
        :key="group"
        @click="selectedGroup = group"
        :class="['px-3 py-1 text-sm rounded-full border', selectedGroup === group ? 'bg-indigo-600 text-white border-indigo-600' : 'bg-white text-gray-700 border-gray-300 hover:bg-gray-50']"
      >
        {{ group }}
      </button>
    </div>

    <div class="bg-white shadow overflow-x-auto sm:rounded-lg">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Slug</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('rbac.displayName') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('rbac.group') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('rbac.description') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('common.actions') }}</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-if="loading">
            <td colspan="6" class="px-6 py-4 text-center text-sm text-gray-500">{{ $t('common.loading') }}</td>
          </tr>
          <tr v-else-if="filteredPermissions.length === 0">
            <td colspan="6" class="px-6 py-4 text-center text-sm text-gray-500">{{ $t('common.noData') }}</td>
          </tr>
          <tr v-for="perm in filteredPermissions" :key="perm.id">
            <td class="px-6 py-4 text-sm text-gray-900">{{ perm.id }}</td>
            <td class="px-6 py-4 text-sm font-mono text-indigo-700">{{ perm.slug }}</td>
            <td class="px-6 py-4 text-sm text-gray-900">{{ perm.displayName }}</td>
            <td class="px-6 py-4 text-sm">
              <span class="px-2 py-0.5 text-xs rounded-full bg-gray-100 text-gray-600">{{ perm.groupName || '-' }}</span>
            </td>
            <td class="px-6 py-4 text-sm text-gray-500">{{ perm.description || '-' }}</td>
            <td class="px-6 py-4 text-sm space-x-2">
              <button @click="openEditModal(perm)" class="text-indigo-600 hover:text-indigo-900">{{ $t('common.edit') }}</button>
              <button @click="confirmDelete(perm.id)" class="text-red-600 hover:text-red-900">{{ $t('common.delete') }}</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Create/Edit Permission Modal -->
    <div v-if="showModal" class="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center z-10">
      <div class="bg-white rounded-lg p-6 w-full max-w-md shadow-xl">
        <h2 class="text-lg font-semibold mb-4">{{ editingPerm?.id ? $t('rbac.editPermission') : $t('rbac.createPermission') }}</h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">Slug <span class="text-red-500">*</span></label>
            <input v-model="form.slug" type="text" placeholder="resource:action" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 text-sm font-mono focus:outline-none focus:ring-indigo-500 focus:border-indigo-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">{{ $t('rbac.displayName') }}</label>
            <input v-model="form.displayName" type="text" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">{{ $t('rbac.group') }}</label>
            <input v-model="form.groupName" type="text" placeholder="e.g. question" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">{{ $t('rbac.description') }}</label>
            <textarea v-model="form.description" rows="2" class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"></textarea>
          </div>
        </div>
        <div class="mt-6 flex justify-end space-x-3">
          <button @click="showModal = false" class="px-4 py-2 border border-gray-300 rounded-md text-sm text-gray-700 hover:bg-gray-50">{{ $t('common.cancel') }}</button>
          <button @click="save" class="px-4 py-2 bg-indigo-600 text-white rounded-md text-sm hover:bg-indigo-700">{{ $t('common.save') }}</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import rbacService from '../../services/rbacService'
import type { Permission } from '../../services/rbacService'

const { t } = useI18n()

const permissions = ref<Permission[]>([])
const loading = ref(true)
const selectedGroup = ref('')

const showModal = ref(false)
const editingPerm = ref<Permission | null>(null)
const form = ref({ slug: '', displayName: '', groupName: '', description: '' })

const fetchPermissions = async () => {
  loading.value = true
  try {
    const res = await rbacService.listPermissions()
    permissions.value = res.data.list || []
  } catch (e) {
    console.error('Failed to fetch permissions', e)
  } finally {
    loading.value = false
  }
}

onMounted(fetchPermissions)

const groups = computed(() => {
  const set = new Set<string>()
  for (const p of permissions.value) {
    if (p.groupName) set.add(p.groupName)
  }
  return Array.from(set).sort()
})

const filteredPermissions = computed(() => {
  if (!selectedGroup.value) return permissions.value
  return permissions.value.filter(p => p.groupName === selectedGroup.value)
})

const openCreateModal = () => {
  editingPerm.value = null
  form.value = { slug: '', displayName: '', groupName: '', description: '' }
  showModal.value = true
}

const openEditModal = (perm: Permission) => {
  editingPerm.value = perm
  form.value = {
    slug: perm.slug,
    displayName: perm.displayName,
    groupName: perm.groupName || '',
    description: perm.description || ''
  }
  showModal.value = true
}

const save = async () => {
  try {
    if (editingPerm.value?.id) {
      await rbacService.updatePermission({ ...form.value, id: editingPerm.value.id })
    } else {
      await rbacService.createPermission(form.value)
    }
    showModal.value = false
    await fetchPermissions()
  } catch (e) {
    console.error('Failed to save permission', e)
  }
}

const confirmDelete = async (id: number) => {
  if (confirm(t('rbac.confirmDeletePermission'))) {
    try {
      await rbacService.deletePermission(id)
      await fetchPermissions()
    } catch (e) {
      console.error('Failed to delete permission', e)
    }
  }
}
</script>
