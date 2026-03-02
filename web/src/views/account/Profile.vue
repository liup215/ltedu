<template>
  <div class="space-y-6">
    <header>
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('accountProfile.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('accountProfile.subtitle') }}</p>
    </header>

    <div v-if="userStore.user" class="space-y-8">
      <!-- Personal Information Section -->
      <section class="bg-white shadow sm:rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">{{ $t('accountProfile.personalInfo') }}</h3>
          <p class="mt-1 max-w-2xl text-sm text-gray-500">
            {{ $t('accountProfile.personalInfoTip') }}
          </p>
          <dl class="mt-5 grid grid-cols-1 gap-x-4 gap-y-8 sm:grid-cols-2">
            <div class="sm:col-span-1">
              <dt class="text-sm font-medium text-gray-500">{{ $t('accountProfile.username') }}</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ userStore.user.username }}</dd>
            </div>
            <div class="sm:col-span-1">
              <dt class="text-sm font-medium text-gray-500">{{ $t('accountProfile.email') }}</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ userStore.user.email || $t('accountProfile.na') }}</dd>
            </div>
            <div class="sm:col-span-1">
              <dt class="text-sm font-medium text-gray-500">{{ $t('accountProfile.role') }}</dt>
              <dd class="mt-1 text-sm text-gray-900">
                {{ userStore.user.isAdmin ? $t('accountProfile.admin') : (userStore.user.roles?.[0]?.displayName || $t('accountProfile.user')) }}
              </dd>
            </div>
          </dl>
          <div class="mt-6">
            <button type="button"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              {{ $t('accountProfile.editProfile') }}
            </button>
          </div>
        </div>
      </section>

      <!-- Change Password Section -->
      <section class="bg-white shadow sm:rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">{{ $t('accountProfile.changePassword') }}</h3>
          <p class="mt-1 max-w-2xl text-sm text-gray-500">
            {{ $t('accountProfile.changePasswordTip') }}
          </p>
          <div class="mt-5">
            <form class="space-y-4" @submit="handleChangePassword">
              <div>
                <label for="current_password" class="block text-sm font-medium text-gray-700">{{ $t('accountProfile.currentPassword') }}</label>
                <input type="password" name="current_password" id="current_password" autocomplete="current-password"
                  class="mt-1 block w-full sm:max-w-xs border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  :placeholder="$t('accountProfile.passwordPlaceholder')"
                  v-model="changePasswordForm.oldPassword"
                  :disabled="loading"
                >
              </div>
              <div>
                <label for="new_password" class="block text-sm font-medium text-gray-700">{{ $t('accountProfile.newPassword') }}</label>
                <input type="password" name="new_password" id="new_password" autocomplete="new-password"
                  class="mt-1 block w-full sm:max-w-xs border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  :placeholder="$t('accountProfile.passwordPlaceholder')"
                  v-model="changePasswordForm.newPassword"
                  :disabled="loading"
                >
              </div>
              <div>
                <label for="confirm_password" class="block text-sm font-medium text-gray-700">{{ $t('accountProfile.confirmPassword') }}</label>
                <input type="password" name="confirm_password" id="confirm_password" autocomplete="new-password"
                  class="mt-1 block w-full sm:max-w-xs border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  :placeholder="$t('accountProfile.passwordPlaceholder')"
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
                {{ loading ? $t('accountProfile.changing') : $t('accountProfile.changePasswordBtn') }}
              </button>
            </form>
          </div>
        </div>
      </section>
    </div>
    <div v-else class="text-center py-10">
      <p class="text-gray-500">{{ $t('accountProfile.loading') }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useUserStore } from '../../stores/userStore';
import { authService } from '../../services/authService';
import type { ChangePasswordRequest } from '../../models/auth.model';

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
