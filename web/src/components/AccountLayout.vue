<template>
  <div class="flex flex-col min-h-screen bg-white">
    <NavBar />
    <div class="flex flex-1 pt-20"> <!-- pt-20 to account for fixed NavBar -->
      <AccountSidebar />
      <main class="flex-1 p-6 bg-gray-100"> <!-- Main content area with some padding and a different bg for distinction -->
        <router-view />
      </main>
    </div>
    <!-- Footer could be here if needed for this specific layout, or rely on a global one if App.vue has one -->
    <!-- For now, no separate footer in this layout -->
  </div>
</template>

<script setup lang="ts">
import NavBar from './NavBar.vue';
import AccountSidebar from './AccountSidebar.vue';
</script>

<style scoped>
/* Tailwind handles styling */
/* Ensure the main content area can scroll if content overflows */
main {
  overflow-y: auto;
}
</style>
