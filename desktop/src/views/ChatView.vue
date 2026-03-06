<script setup lang="ts">
import AIChatInterface from '../components/AIChatInterface.vue'
</script>

<template>
  <div class="chat-view">
    <AIChatInterface />
  </div>
</template>

<style scoped>
.chat-view {
  display: flex;
  flex-direction: column;
  height: 100%;
  overflow: hidden;
}
</style>
