import { createI18n } from 'vue-i18n'
import en from '../locales/en'
import zh from '../locales/zh'

const getDefaultLocale = () => {
  const saved = localStorage.getItem('locale')
  if (saved) return saved
  const browser = navigator.language.toLowerCase()
  if (browser === 'zh-cn' || browser === 'zh') return 'zh'
  return 'en'
}

const i18n = createI18n({
  legacy: false,
  locale: getDefaultLocale(),
  fallbackLocale: 'en',
  messages: {
    en,
    zh,
  },
  globalInjection: true,
})

export default i18n
