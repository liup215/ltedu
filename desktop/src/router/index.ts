import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/authStore'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/LoginView.vue'),
    },
    {
      path: '/',
      name: 'Main',
      component: () => import('../views/MainLayout.vue'),
      meta: { requiresAuth: true },
      children: [
        {
          path: '',
          name: 'Chat',
          component: () => import('../views/ChatView.vue'),
        },
        {
          path: 'settings',
          name: 'Settings',
          component: () => import('../views/SettingsView.vue'),
        },
      ],
    },
  ],
})

router.beforeEach((to, _from, next) => {
  const authStore = useAuthStore()
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next({ name: 'Login' })
  } else if (to.name === 'Login' && authStore.isAuthenticated) {
    next({ name: 'Chat' })
  } else {
    next()
  }
})

export default router
