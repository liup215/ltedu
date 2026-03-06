<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '../stores/authStore'
import { cacheClear } from '../services/tauriClient'

const authStore = useAuthStore()
const serverUrl = ref(authStore.apiBaseUrl)
const saving = ref(false)
const saved = ref(false)
const clearingCache = ref(false)
const cacheCleared = ref(false)

async function saveSettings() {
  saving.value = true
  authStore.apiBaseUrl = serverUrl.value
  await new Promise(resolve => setTimeout(resolve, 400))
  saving.value = false
  saved.value = true
  setTimeout(() => { saved.value = false }, 2000)
}

async function handleClearCache() {
  clearingCache.value = true
  await cacheClear()
  clearingCache.value = false
  cacheCleared.value = true
  setTimeout(() => { cacheCleared.value = false }, 2000)
}

async function handleLogout() {
  await authStore.logout()
}
</script>

<template>
  <div class="settings-view">
    <h2>Settings</h2>

    <section class="section">
      <h3>Connection</h3>
      <div class="field">
        <label>LTEdu Server URL</label>
        <input v-model="serverUrl" type="url" placeholder="http://localhost:8080" />
      </div>
      <button class="btn-primary" :disabled="saving" @click="saveSettings">
        {{ saving ? 'Saving…' : saved ? '✓ Saved' : 'Save' }}
      </button>
    </section>

    <section class="section">
      <h3>Offline Cache</h3>
      <p class="desc">
        The desktop client caches API responses in a local SQLite database for
        offline access. Use this to wipe the cache if you encounter stale data.
      </p>
      <button class="btn-danger" :disabled="clearingCache" @click="handleClearCache">
        {{ clearingCache ? 'Clearing…' : cacheCleared ? '✓ Cache cleared' : 'Clear Cache' }}
      </button>
    </section>

    <section class="section">
      <h3>Account</h3>
      <div class="user-info" v-if="authStore.user">
        <span>Signed in as <strong>{{ authStore.user.username }}</strong></span>
      </div>
      <button class="btn-danger" @click="handleLogout">Sign Out</button>
    </section>
  </div>
</template>

<style scoped>
.settings-view {
  padding: 32px;
  max-width: 600px;
  overflow-y: auto;
  height: 100%;
}

h2 {
  font-size: 22px;
  font-weight: 700;
  color: #1a202c;
  margin-bottom: 24px;
}

.section {
  background: white;
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 20px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.08);
}

.section h3 {
  font-size: 15px;
  font-weight: 600;
  color: #4a5568;
  margin-bottom: 16px;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 6px;
  margin-bottom: 16px;
}

.field label {
  font-size: 13px;
  font-weight: 600;
  color: #718096;
}

.field input {
  padding: 10px 14px;
  border: 1.5px solid #e2e8f0;
  border-radius: 8px;
  font-size: 14px;
  outline: none;
}

.field input:focus {
  border-color: #667eea;
}

.desc {
  font-size: 13px;
  color: #718096;
  margin-bottom: 16px;
  line-height: 1.5;
}

.user-info {
  font-size: 14px;
  color: #4a5568;
  margin-bottom: 16px;
}

.btn-primary {
  background: #667eea;
  color: white;
  border: none;
  border-radius: 8px;
  padding: 10px 20px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
}

.btn-primary:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.btn-danger {
  background: #e53e3e;
  color: white;
  border: none;
  border-radius: 8px;
  padding: 10px 20px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
}

.btn-danger:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
</style>
