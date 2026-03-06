<template>
  <div class="flex flex-col min-h-screen bg-white">
    <NavBar />
    <div class="flex flex-1 pt-16"> <!-- pt-16 to account for fixed NavBar -->
      <!-- Mobile sidebar toggle button -->
      <button
        @click="sidebarOpen = !sidebarOpen"
        class="md:hidden fixed bottom-4 right-4 z-50 w-14 h-14 bg-indigo-600 text-white rounded-full shadow-lg flex items-center justify-center focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        aria-label="Toggle account menu"
      >
        <svg v-if="!sidebarOpen" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
        </svg>
        <svg v-else class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>

      <!-- Sidebar overlay for mobile -->
      <div
        v-if="sidebarOpen"
        class="md:hidden fixed inset-0 bg-gray-600 bg-opacity-50 z-30"
        @click="sidebarOpen = false"
      ></div>

      <!-- Sidebar -->
      <div
        :class="[
          'fixed md:static inset-y-0 left-0 z-40 transform transition-transform duration-300 ease-in-out md:transform-none',
          sidebarOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0',
          'pt-16 md:pt-0'
        ]"
      >
        <AccountSidebar @navigate="sidebarOpen = false" />
      </div>

      <main class="flex-1 p-4 md:p-6 bg-gray-100 min-w-0"> <!-- Main content area with some padding and a different bg for distinction -->
        <router-view />
      </main>
    </div>
    <!-- Footer could be here if needed for this specific layout, or rely on a global one if App.vue has one -->
    <!-- For now, no separate footer in this layout -->
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import NavBar from './NavBar.vue';
import AccountSidebar from './AccountSidebar.vue';

const sidebarOpen = ref(false);
</script>

<style scoped>
/* Tailwind handles styling */
/* Ensure the main content area can scroll if content overflows */
main {
  overflow-y: auto;
}
</style>
