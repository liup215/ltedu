import { createApp } from 'vue'
import { createPinia } from 'pinia'
import './style.css'
import App from './App.vue'
import router from './router'
import { useUserStore } from './stores/userStore'
import { useAppStore } from './stores/appStore'
import { createNotivue } from 'notivue'
import 'notivue/notification.css'
import 'notivue/animations.css'
import i18n from './i18n'



const app = createApp(App)

const pinia = createPinia()
const notivue = createNotivue()

app.use(pinia)
app.use(i18n)
app.use(router)
app.use(notivue)

// Initialize stores by loading data from localStorage
const userStore = useUserStore()
userStore.loadUserFromStorage()

const appStore = useAppStore()
appStore.loadThemeFromStorage()

// Register service worker for offline support and API response caching
if ('serviceWorker' in navigator) {
  window.addEventListener('load', () => {
    navigator.serviceWorker.register('/sw.js').catch(() => {
      // Service worker registration failure is non-fatal
    })
  })
}

app.mount('#app');
