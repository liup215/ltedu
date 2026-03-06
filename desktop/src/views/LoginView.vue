<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/authStore'

const authStore = useAuthStore()
const router = useRouter()

const form = reactive({ username: '', password: '' })
const advancedVisible = ref(false)
const serverUrl = ref(authStore.apiBaseUrl)
const loading = ref(false)
const errorMsg = ref('')

async function handleLogin() {
  errorMsg.value = ''
  loading.value = true
  try {
    await authStore.login(form.username, form.password, serverUrl.value || undefined)
    router.push({ name: 'Chat' })
  } catch (err: unknown) {
    errorMsg.value = err instanceof Error ? err.message : String(err)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="login-wrapper">
    <div class="login-card">
      <div class="brand">
        <span class="brand-icon">🎓</span>
        <h1>LTEdu Desktop</h1>
        <p class="subtitle">AI-native learning platform</p>
      </div>

      <form @submit.prevent="handleLogin" class="login-form">
        <div class="field">
          <label for="username">Username</label>
          <input
            id="username"
            v-model="form.username"
            type="text"
            placeholder="Enter your username"
            autocomplete="username"
            required
          />
        </div>

        <div class="field">
          <label for="password">Password</label>
          <input
            id="password"
            v-model="form.password"
            type="password"
            placeholder="Enter your password"
            autocomplete="current-password"
            required
          />
        </div>

        <div class="advanced-toggle" @click="advancedVisible = !advancedVisible">
          <span>⚙ Advanced (Server URL)</span>
          <span>{{ advancedVisible ? '▲' : '▼' }}</span>
        </div>

        <div v-if="advancedVisible" class="field">
          <label for="serverUrl">Server URL</label>
          <input
            id="serverUrl"
            v-model="serverUrl"
            type="url"
            placeholder="http://localhost:8080"
          />
        </div>

        <div v-if="errorMsg" class="error">{{ errorMsg }}</div>

        <button type="submit" :disabled="loading" class="btn-primary">
          <span v-if="loading">Signing in…</span>
          <span v-else>Sign In</span>
        </button>
      </form>
    </div>
  </div>
</template>

<style scoped>
.login-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.login-card {
  background: white;
  border-radius: 16px;
  padding: 40px;
  width: 380px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
}

.brand {
  text-align: center;
  margin-bottom: 32px;
}

.brand-icon {
  font-size: 48px;
}

.brand h1 {
  font-size: 24px;
  font-weight: 700;
  color: #1a202c;
  margin-top: 8px;
}

.subtitle {
  color: #718096;
  font-size: 14px;
  margin-top: 4px;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.field label {
  font-size: 13px;
  font-weight: 600;
  color: #4a5568;
}

.field input {
  padding: 10px 14px;
  border: 1.5px solid #e2e8f0;
  border-radius: 8px;
  font-size: 14px;
  outline: none;
  transition: border-color 0.2s;
}

.field input:focus {
  border-color: #667eea;
}

.advanced-toggle {
  display: flex;
  justify-content: space-between;
  cursor: pointer;
  font-size: 12px;
  color: #667eea;
  padding: 4px 0;
  user-select: none;
}

.error {
  background: #fff5f5;
  border: 1px solid #fed7d7;
  border-radius: 8px;
  padding: 10px 14px;
  color: #c53030;
  font-size: 13px;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 8px;
  padding: 12px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: opacity 0.2s;
}

.btn-primary:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
</style>
