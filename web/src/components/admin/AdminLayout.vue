<template>
  <div class="flex flex-col min-h-screen bg-gray-100"> <!-- Main background for admin area -->
    <NavBar />
    <div class="flex flex-1 pt-20"> <!-- pt-20 to account for fixed NavBar -->
      <AdminSidebar />
      <main class="flex-1 p-6"> <!-- Main content area for admin pages -->
        <router-view />
      </main>
    </div>
    <!-- No separate footer in this admin layout for now -->
  </div>
</template>

<script setup lang="ts">
import NavBar from '../NavBar.vue'; // Adjust path if NavBar is moved or aliased
import AdminSidebar from './AdminSidebar.vue';
</script>

<style scoped>
/* Tailwind handles styling */
/* Ensure the main content area can scroll if content overflows */
main {
  overflow-y: auto;
}
</style>
