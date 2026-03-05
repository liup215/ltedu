<template>
  <aside class="w-64 bg-gray-50 p-4 border-r border-gray-200 h-full"> <!-- Ensure sidebar takes full height of its container -->
    <nav class="space-y-2">
      <h3 class="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider mb-2">
        {{ $t('accountSidebar.account') }}
      </h3>
      <router-link
        to="/account/profile"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900"
        active-class="bg-gray-200 text-gray-900"
      >
        {{ $t('accountSidebar.profile') }}
      </router-link>
      <router-link
        to="/account/settings"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900"
        active-class="bg-gray-200 text-gray-900"
      >
        {{ $t('accountSidebar.settings') }}
      </router-link>
      <router-link
        to="/account/teacher-application"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900"
        active-class="bg-gray-200 text-gray-900"
      >
        {{ $t('accountSidebar.teacherApplication') }}
      </router-link>
      <router-link
        to="/account/mcp-tokens"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900"
        active-class="bg-gray-200 text-gray-900"
      >
        {{ $t('accountSidebar.mcpTokens') }}
      </router-link>
      <router-link
        to="/account/cli-tokens"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900"
        active-class="bg-gray-200 text-gray-900"
      >
        {{ $t('accountSidebar.cliTokens') }}
      </router-link>
      <router-link
        to="/account/analytics"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900"
        active-class="bg-gray-200 text-gray-900"
      >
        {{ $t('analytics.recommendations') }}
      </router-link>
      <!-- Add more account-related links here as needed -->
      <!-- Example:
      <router-link
        to="/account/security"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900"
        active-class="bg-gray-200 text-gray-900"
      >
        Security
      </router-link>
      <router-link
        to="/account/notifications"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900"
        active-class="bg-gray-200 text-gray-900"
      >
        Notifications
      </router-link>
      -->
    </nav>
  </aside>
</template>

<script setup lang="ts">

</script>

<style scoped>
/* Tailwind handles styling. active-class provides styling for active links. */
/* Ensure the sidebar itself can fill height if its parent is flex */
</style>
