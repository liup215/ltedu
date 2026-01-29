<template>
  <div class="space-y-6">
    <header>
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('accountSettings.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('accountSettings.subtitle') }}</p>
    </header>

    <div class="space-y-8">
      <!-- Notification Settings Section -->
      <section class="bg-white shadow sm:rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">{{ $t('accountSettings.notificationTitle') }}</h3>
          <p class="mt-1 max-w-2xl text-sm text-gray-500">
            {{ $t('accountSettings.notificationTip') }}
          </p>
          <div class="mt-5 space-y-4">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-gray-900">{{ $t('accountSettings.emailNotifications') }}</p>
                <p class="text-sm text-gray-500">{{ $t('accountSettings.emailNotificationsTip') }}</p>
              </div>
              <button type="button" :class="[emailNotificationsEnabled ? 'bg-indigo-600' : 'bg-gray-200', 'relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500']" role="switch" :aria-checked="emailNotificationsEnabled" @click="toggleEmailNotifications">
                <span class="sr-only">{{ $t('accountSettings.useSetting') }}</span>
                <span aria-hidden="true" :class="[emailNotificationsEnabled ? 'translate-x-5' : 'translate-x-0', 'pointer-events-none inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200']"></span>
              </button>
            </div>
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-gray-900">{{ $t('accountSettings.inAppNotifications') }}</p>
                <p class="text-sm text-gray-500">{{ $t('accountSettings.inAppNotificationsTip') }}</p>
              </div>
              <button type="button" :class="[inAppNotificationsEnabled ? 'bg-indigo-600' : 'bg-gray-200', 'relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500']" role="switch" :aria-checked="inAppNotificationsEnabled" @click="toggleInAppNotifications">
                <span class="sr-only">{{ $t('accountSettings.useSetting') }}</span>
                <span aria-hidden="true" :class="[inAppNotificationsEnabled ? 'translate-x-5' : 'translate-x-0', 'pointer-events-none inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200']"></span>
              </button>
            </div>
          </div>
          <div class="mt-6">
            <button type="button" class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              {{ $t('accountSettings.saveNotification') }}
            </button>
          </div>
        </div>
      </section>

      <!-- Language & Region Section -->
      <section class="bg-white shadow sm:rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">{{ $t('accountSettings.languageRegionTitle') }}</h3>
          <p class="mt-1 max-w-2xl text-sm text-gray-500">
            {{ $t('accountSettings.languageRegionTip') }}
          </p>
          <div class="mt-5 grid grid-cols-1 gap-y-6 sm:grid-cols-2 sm:gap-x-4">
            <div>
              <label for="language" class="block text-sm font-medium text-gray-700">{{ $t('accountSettings.language') }}</label>
              <select id="language" name="language" class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md">
                <option>English</option>
                <option>中文 (简体)</option>
                <option>Español</option>
              </select>
            </div>
            <div>
              <label for="timezone" class="block text-sm font-medium text-gray-700">{{ $t('accountSettings.timezone') }}</label>
              <select id="timezone" name="timezone" class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md">
                <option>UTC-8:00 Pacific Time (US & Canada)</option>
                <option>UTC+0:00 Greenwich Mean Time</option>
                <option>UTC+8:00 Beijing, Perth, Singapore</option>
              </select>
            </div>
          </div>
           <div class="mt-6">
            <button type="button" class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              {{ $t('accountSettings.saveLanguageRegion') }}
            </button>
          </div>
        </div>
      </section>

      <!-- Account Deletion Section (Optional) -->
      <section class="bg-white shadow sm:rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium leading-6 text-red-700">{{ $t('accountSettings.deleteTitle') }}</h3>
          <p class="mt-1 max-w-2xl text-sm text-gray-500">
            {{ $t('accountSettings.deleteTip') }}
          </p>
          <div class="mt-5">
            <button type="button" class="inline-flex items-center justify-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-red-700 bg-red-100 hover:bg-red-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500">
              {{ $t('accountSettings.requestDelete') }}
            </button>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const emailNotificationsEnabled = ref(true);
const inAppNotificationsEnabled = ref(true);

const toggleEmailNotifications = () => {
  emailNotificationsEnabled.value = !emailNotificationsEnabled.value;
};

const toggleInAppNotifications = () => {
  inAppNotificationsEnabled.value = !inAppNotificationsEnabled.value;
};
</script>

<style scoped>
/* Tailwind handles most styling. Add specific overrides if needed. */
</style>
