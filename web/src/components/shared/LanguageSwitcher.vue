<template>
  <div class="relative">
    <span
      @click="toggleDropdown"
      class="flex items-center cursor-pointer select-none text-gray-700 text-sm font-medium hover:text-indigo-600"
    >
      <span class="mr-1">{{ currentLabel }}</span>
      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
      </svg>
    </span>
    <div
      v-if="dropdownOpen"
      class="absolute right-0 mt-2 w-32 bg-white rounded-md shadow-lg py-1 z-50 border border-gray-200"
      @click.away="dropdownOpen = false"
    >
      <button
        v-for="lang in languages"
        :key="lang.value"
        @click="selectLanguage(lang.value)"
        class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
        :class="{ 'font-bold text-indigo-600': currentLocale === lang.value }"
      >
        {{ lang.label }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'

const { locale } = useI18n()

const languages = [
  { value: 'en', label: 'English' },
  { value: 'zh', label: '中文' },
]

const currentLocale = ref(locale.value)
const dropdownOpen = ref(false)

const currentLabel = computed(() => {
  const found = languages.find(l => l.value === currentLocale.value)
  return found ? found.label : ''
})

function toggleDropdown() {
  dropdownOpen.value = !dropdownOpen.value
}

function selectLanguage(lang: string) {
  currentLocale.value = lang
  locale.value = lang
  localStorage.setItem('locale', lang)
  dropdownOpen.value = false
  location.reload()
}
</script>
