import { defineStore } from 'pinia'
import { invoke } from '@tauri-apps/api/core'
import { ref, computed } from 'vue'

export interface User {
  id: number
  username: string
  email: string
  nickname?: string
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(null)
  const user = ref<User | null>(null)
  const apiBaseUrl = ref<string>('http://localhost:8080')

  const isAuthenticated = computed(() => !!token.value)

  async function restoreSession() {
    try {
      const result = await invoke<{ token: string | null; user: User | null; apiBaseUrl: string | null }>(
        'cmd_get_current_user'
      )
      if (result.token) {
        token.value = result.token
        user.value = result.user
      }
      if (result.apiBaseUrl) {
        apiBaseUrl.value = result.apiBaseUrl
      }
    } catch (err) {
      console.error('Failed to restore session:', err)
    }
  }

  async function login(username: string, password: string, baseUrl?: string) {
    const url = baseUrl ?? apiBaseUrl.value
    const result = await invoke<{ token: string; user: User }>('cmd_login', {
      request: { username, password, api_base_url: url },
    })
    token.value = result.token
    user.value = result.user
    apiBaseUrl.value = url
  }

  async function logout() {
    await invoke('cmd_logout')
    token.value = null
    user.value = null
  }

  return { token, user, apiBaseUrl, isAuthenticated, restoreSession, login, logout }
})
