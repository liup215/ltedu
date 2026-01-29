<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('adminDashboard.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('adminDashboard.subtitle') }}</p>
    </header>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <!-- Placeholder cards for admin sections -->
      <div class="bg-white shadow rounded-lg p-4">
        <h2 class="text-lg font-medium text-gray-900">{{ $t('adminDashboard.userManagement') }}</h2>
        <p class="mt-1 text-sm text-gray-500">{{ $t('adminDashboard.userManagementTip') }}</p>
        <router-link to="/admin/users" class="mt-3 inline-block text-indigo-600 hover:text-indigo-800">{{ $t('adminDashboard.goToUsers') }}</router-link>
      </div>
      <div class="bg-white shadow rounded-lg p-4">
        <h2 class="text-lg font-medium text-gray-900">{{ $t('adminDashboard.systemSettings') }}</h2>
        <p class="mt-1 text-sm text-gray-500">{{ $t('adminDashboard.systemSettingsTip') }}</p>
        <router-link to="/admin/system-settings" class="mt-3 inline-block text-indigo-600 hover:text-indigo-800">{{ $t('adminDashboard.goToSettings') }}</router-link>
      </div>
      <!-- <div class="bg-white shadow rounded-lg p-4">
        <h2 class="text-lg font-medium text-gray-900">Content Management</h2>
        <p class="mt-1 text-sm text-gray-500">Manage courses, lessons, and other content.</p>
        <router-link to="/admin/content" class="mt-3 inline-block text-indigo-600 hover:text-indigo-800">Go to Content &raquo;</router-link>
      </div> -->
    </div>
  </div>
</template>

<script setup lang="ts">
// Future logic for dashboard data
</script>

<style scoped>
/* Tailwind handles styling */
</style>
