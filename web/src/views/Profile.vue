<template>
  <div class="space-y-6">
    <header>
      <h1 class="text-3xl font-bold text-gray-900">User Profile</h1>
      <p class="mt-1 text-sm text-gray-500">Manage your personal information and account settings.</p>
    </header>

    <div v-if="userStore.user" class="space-y-8">
      <!-- Personal Information Section -->
      <section class="bg-white shadow sm:rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">Personal Information</h3>
          <p class="mt-1 max-w-2xl text-sm text-gray-500">
            This information will be displayed publicly so be careful what you share.
          </p>
          <dl class="mt-5 grid grid-cols-1 gap-x-4 gap-y-8 sm:grid-cols-2">
            <div class="sm:col-span-1">
              <dt class="text-sm font-medium text-gray-500">Username</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ userStore.user.username }}</dd>
            </div>
            <div class="sm:col-span-1">
              <dt class="text-sm font-medium text-gray-500">Email address</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ userStore.user.email || 'N/A' }}</dd>
            </div>
            <!-- Mobile field removed as it's not in the User store type -->
            <div class="sm:col-span-1">
              <dt class="text-sm font-medium text-gray-500">Role</dt>
              <dd class="mt-1 text-sm text-gray-900">
                {{ userStore.user.isAdmin ? 'Administrator' : (userStore.user.roles?.[0]?.displayName || 'User') }}
              </dd>
            </div>
            <!-- Add more fields like Full Name, Bio, etc. if available in userStore.user -->
          </dl>
          <div class="mt-6">
            <button type="button"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              Edit Profile (Placeholder)
            </button>
          </div>
        </div>
      </section>

      <!-- Change Password Section -->
      <section class="bg-white shadow sm:rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">Change Password</h3>
          <p class="mt-1 max-w-2xl text-sm text-gray-500">
            Ensure your account is using a long, random password to stay secure.
          </p>
          <div class="mt-5">
            <!-- Placeholder for change password form -->
<form class="space-y-4" @submit="handleChangePassword">
  <div>
    <label for="current_password" class="block text-sm font-medium text-gray-700">Current Password</label>
    <input type="password" name="current_password" id="current_password" autocomplete="current-password"
      class="mt-1 block w-full sm:max-w-xs border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
      placeholder="••••••••"
      v-model="changePasswordForm.oldPassword"
      :disabled="loading"
    >
  </div>
  <div>
    <label for="new_password" class="block text-sm font-medium text-gray-700">New Password</label>
    <input type="password" name="new_password" id="new_password" autocomplete="new-password"
      class="mt-1 block w-full sm:max-w-xs border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
      placeholder="••••••••"
      v-model="changePasswordForm.newPassword"
      :disabled="loading"
    >
  </div>
  <div>
    <label for="confirm_password" class="block text-sm font-medium text-gray-700">Confirm New Password</label>
    <input type="password" name="confirm_password" id="confirm_password" autocomplete="new-password"
      class="mt-1 block w-full sm:max-w-xs border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
      placeholder="••••••••"
      v-model="confirmPassword"
      :disabled="loading"
    >
  </div>
  <div v-if="errorMessage" class="text-sm text-red-600">{{ errorMessage }}</div>
  <div v-if="successMessage" class="text-sm text-green-600">{{ successMessage }}</div>
  <button type="submit"
    class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
    :disabled="loading"
  >
    {{ loading ? 'Changing...' : 'Change Password' }}
  </button>
</form>
          </div>
        </div>
      </section>

      <!-- Avatar Section (Optional) -->
      <!--
      <section class="bg-white shadow sm:rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">Profile Picture</h3>
          <div class="mt-4 flex items-center space-x-4">
            <span class="inline-block h-12 w-12 rounded-full overflow-hidden bg-gray-100">
              <svg v-if="!userStore.user.avatarUrl" class="h-full w-full text-gray-300" fill="currentColor" viewBox="0 0 24 24">
                <path d="M24 20.993V24H0v-2.996A14.977 14.977 0 0112.004 15c4.904 0 9.26 2.354 11.996 5.993zM16.002 8.999a4 4 0 11-8 0 4 4 0 018 0z" />
              </svg>
              <img v-else :src="userStore.user.avatarUrl" alt="User Avatar" class="h-full w-full object-cover">
            </span>
            <button type="button" class="ml-5 bg-white py-2 px-3 border border-gray-300 rounded-md shadow-sm text-sm leading-4 font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              Change (Placeholder)
            </button>
          </div>
        </div>
      </section>
      -->

    </div>
    <div v-else class="text-center py-10">
      <p class="text-gray-500">Loading user profile...</p>
      <!-- Or a spinner component -->
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useUserStore } from '../stores/userStore';
import { authService } from '../services/authService';
import type { ChangePasswordRequest } from '../models/auth.model';

const userStore = useUserStore();

const changePasswordForm = ref<ChangePasswordRequest>({
  oldPassword: '',
  newPassword: ''
});
const confirmPassword = ref('');
const loading = ref(false);
const errorMessage = ref('');
const successMessage = ref('');

const handleChangePassword = async (e: Event) => {
  e.preventDefault();
  errorMessage.value = '';
  successMessage.value = '';

  if (!changePasswordForm.value.oldPassword || !changePasswordForm.value.newPassword || !confirmPassword.value) {
    errorMessage.value = 'All fields are required.';
    return;
  }
  if (changePasswordForm.value.newPassword !== confirmPassword.value) {
    errorMessage.value = 'New passwords do not match.';
    return;
  }
  loading.value = true;
  try {
    const res = await authService.changePassword(changePasswordForm.value);
    if (res.code === 0) {
      successMessage.value = 'Password changed successfully. Please log in again on all devices.';
      changePasswordForm.value.oldPassword = '';
      changePasswordForm.value.newPassword = '';
      confirmPassword.value = '';
    } else {
      errorMessage.value = res.message || 'Failed to change password.';
    }
  } catch (err: any) {
    errorMessage.value = err?.message || 'Failed to change password.';
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
/* Tailwind handles most styling. Add specific overrides if needed. */
</style>
