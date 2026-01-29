<template>
  <div class="flex flex-col min-h-screen bg-white">
    <NavBar />
    <main class="flex-1 flex flex-col pt-20"> <!-- Added pt-20 for NavBar height, changed div to main for semantics -->
      <router-view />
    </main>
    <Footer />
  </div>
</template>

<script setup lang="ts">
import NavBar from './NavBar.vue'
import Footer from './Footer.vue'
</script>

<style scoped>
/* Tailwind handles all styling */
</style>
