<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '../stores/authStore'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const navItems = [
  { name: 'Chat', icon: '💬', path: '/' },
  { name: 'Settings', icon: '⚙', path: '/settings' },
]

function isActive(path: string) {
  return route.path === path
}

async function handleLogout() {
  await authStore.logout()
  router.push({ name: 'Login' })
}
</script>

<template>
  <aside class="sidebar">
    <div class="sidebar-brand">
      <span class="brand-icon">🎓</span>
      <span class="brand-name">LTEdu</span>
    </div>

    <nav class="sidebar-nav">
      <RouterLink
        v-for="item in navItems"
        :key="item.path"
        :to="item.path"
        :class="['nav-item', { active: isActive(item.path) }]"
      >
        <span class="nav-icon">{{ item.icon }}</span>
        <span class="nav-label">{{ item.name }}</span>
      </RouterLink>
    </nav>

    <div class="sidebar-footer">
      <div class="user-badge" v-if="authStore.user">
        <div class="avatar">{{ (authStore.user.username ?? '?')[0].toUpperCase() }}</div>
        <div class="user-info">
          <span class="user-name">{{ authStore.user.username }}</span>
          <span class="user-role">Student</span>
        </div>
      </div>
      <button class="logout-btn" @click="handleLogout" title="Sign out">⏻</button>
    </div>
  </aside>
</template>

<style scoped>
.sidebar {
  width: 220px;
  background: #1a202c;
  display: flex;
  flex-direction: column;
  padding: 0;
  flex-shrink: 0;
}

.sidebar-brand {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 20px 18px;
  border-bottom: 1px solid #2d3748;
}

.brand-icon {
  font-size: 24px;
}

.brand-name {
  font-size: 18px;
  font-weight: 700;
  color: white;
  letter-spacing: 0.5px;
}

.sidebar-nav {
  flex: 1;
  padding: 12px 10px;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 12px;
  border-radius: 8px;
  text-decoration: none;
  color: #a0aec0;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.15s;
}

.nav-item:hover {
  background: #2d3748;
  color: white;
}

.nav-item.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.nav-icon {
  font-size: 18px;
  width: 24px;
  text-align: center;
}

.sidebar-footer {
  padding: 14px 14px 18px;
  border-top: 1px solid #2d3748;
  display: flex;
  align-items: center;
  gap: 8px;
}

.user-badge {
  display: flex;
  align-items: center;
  gap: 8px;
  flex: 1;
  min-width: 0;
}

.avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: 700;
  flex-shrink: 0;
}

.user-info {
  display: flex;
  flex-direction: column;
  min-width: 0;
}

.user-name {
  font-size: 13px;
  font-weight: 600;
  color: white;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.user-role {
  font-size: 11px;
  color: #718096;
}

.logout-btn {
  background: none;
  border: none;
  color: #718096;
  font-size: 18px;
  cursor: pointer;
  padding: 4px;
  border-radius: 6px;
  transition: color 0.2s;
  flex-shrink: 0;
}

.logout-btn:hover {
  color: #fc8181;
}
</style>
