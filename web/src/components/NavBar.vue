<template>
  <nav
    class="w-full bg-white border-b border-gray-200 shadow-none py-4 px-8 flex items-center justify-between fixed top-0 left-0 z-50">
    <div class="flex items-center gap-2">
      <router-link to="/" class="flex items-center gap-2">
        <img src="/nerdlet_logo_blue_only.png" alt="Nerdlet Logo" class="h-10 w-auto" />
      </router-link>
    </div>

    <!-- Desktop nav links -->
    <div class="hidden md:flex gap-4 items-center">
      <router-link to="/"
        class="text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" custom
        v-slot="{ navigate, href }">
        <button :href="href" @click="navigate" type="button">{{ $t('navbar.home') }}</button>
      </router-link>
      <button
        class="text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-green-100 hover:text-green-900"
        type="button"
        @click="handleQuickPracticeClick"
      >
        {{ $t('navbar.quickPractice') }}
      </button>
      <button
        class="text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-purple-100 hover:text-purple-900"
        type="button"
        @click="handlePaperPracticeClick"
      >
        {{ $t('navbar.pastPaperPractice') }}
      </button>

      <!-- Exam Paper Links - always visible, navigation logic handled in click -->
      <button
        class="text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-blue-100 hover:text-blue-900"
        type="button"
        @click="handleExamPaperClick('teacher')"
      >
        {{ $t('navbar.myExamPapers') }}
      </button>
      <button
        class="text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-blue-100 hover:text-blue-900"
        type="button"
        @click="handleExamPaperClick('builder')"
      >
        {{ $t('navbar.examPaperBuilder') }}
      </button>

      <!-- Admin Link - visible only to admins -->
      <router-link v-if="userStore.user?.isAdmin" to="/admin"
        class="text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" custom
        v-slot="{ navigate, href }">
        <button :href="href" @click="navigate" type="button">{{ $t('navbar.systemManagement') }}</button>
      </router-link>

      <!-- Help Link -->
      <router-link to="/help"
        class="text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" custom
        v-slot="{ navigate, href }">
        <button :href="href" @click="navigate" type="button">{{ $t('navbar.help') }}</button>
      </router-link>

      <button
        class="px-4 py-2 rounded font-normal shadow transition bg-yellow-400 text-white hover:bg-yellow-500 hover:scale-105 border-2 border-yellow-500"
        type="button"
        @click="router.push('/donate')"
      >
        {{ $t('navbar.donate') }}
      </button>

      <router-link
        to="/blog"
        class="px-4 py-2 rounded font-normal text-gray-700 hover:text-indigo-600 hover:bg-gray-100 transition text-sm"
      >
        {{ $t('navbar.blog') }}
      </router-link>

      <!-- Language Switcher Dropdown -->
      <div class="relative">
        <span
          @click="toggleLangDropdown"
          class="flex items-center cursor-pointer select-none text-gray-700 text-sm font-normal hover:text-indigo-600"
        >
          <span class="mr-1">{{ currentLangLabel }}</span>
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </span>
        <div
          v-if="langDropdownOpen"
          class="absolute right-0 mt-2 w-32 bg-white rounded-md shadow-lg py-1 z-50 border border-gray-200"
          @click.away="langDropdownOpen = false"
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

      <!-- Conditional rendering for Sign In / User Avatar -->
      <div v-if="!userStore.isAuthenticated" class="flex items-center">
        <router-link to="/login"
          class="text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" custom
          v-slot="{ navigate, href }">
          <button :href="href" @click="navigate" type="button">{{ $t('navbar.signIn') }}</button>
        </router-link>
      </div>
      <div v-else class="relative flex items-center">
        <button @click="toggleDropdown"
          class="flex items-center justify-center w-10 h-10 bg-indigo-600 rounded-full text-white text-lg font-normal focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
          {{ userStore.user?.username?.charAt(0).toUpperCase() || 'U' }}
        </button>
        <!-- Dropdown Menu -->
        <div v-if="isDropdownOpen"
          class="absolute right-0 mt-2 top-full w-48 bg-white rounded-md shadow-lg py-1 z-50 border border-gray-200">
          <div class="px-4 py-3">
            <p class="text-sm text-gray-700">{{ $t('navbar.signedInAs') }}</p>
            <p class="text-sm font-normal text-gray-900 truncate">{{ userStore.user?.username }}</p>
          </div>
          <hr class="border-gray-200">
          <router-link to="/account/profile" @click="closeDropdown"
            class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left">{{ $t('navbar.yourProfile') }}</router-link>
          <router-link to="/account/settings" @click="closeDropdown"
            class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left">{{ $t('navbar.settings') }}</router-link>
          <hr class="border-gray-200">
          <button @click="handleLogout"
            class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
            {{ $t('navbar.signOut') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Mobile: right side controls (lang switcher + user avatar + hamburger) -->
    <div class="flex md:hidden items-center gap-2">
      <!-- Language Switcher (mobile) -->
      <div class="relative">
        <span
          @click="toggleLangDropdown"
          class="flex items-center cursor-pointer select-none text-gray-700 text-sm font-normal hover:text-indigo-600"
        >
          <span class="mr-1">{{ currentLangLabel }}</span>
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </span>
        <div
          v-if="langDropdownOpen"
          class="absolute right-0 mt-2 w-32 bg-white rounded-md shadow-lg py-1 z-50 border border-gray-200"
          @click.away="langDropdownOpen = false"
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

      <!-- User Avatar (mobile) -->
      <div v-if="userStore.isAuthenticated" class="relative flex items-center">
        <button @click="toggleDropdown"
          class="flex items-center justify-center w-10 h-10 bg-indigo-600 rounded-full text-white text-lg font-normal focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
          {{ userStore.user?.username?.charAt(0).toUpperCase() || 'U' }}
        </button>
        <div v-if="isDropdownOpen"
          class="absolute right-0 mt-2 top-full w-48 bg-white rounded-md shadow-lg py-1 z-50 border border-gray-200">
          <div class="px-4 py-3">
            <p class="text-sm text-gray-700">{{ $t('navbar.signedInAs') }}</p>
            <p class="text-sm font-normal text-gray-900 truncate">{{ userStore.user?.username }}</p>
          </div>
          <hr class="border-gray-200">
          <router-link to="/account/profile" @click="closeDropdown"
            class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left">{{ $t('navbar.yourProfile') }}</router-link>
          <router-link to="/account/settings" @click="closeDropdown"
            class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left">{{ $t('navbar.settings') }}</router-link>
          <hr class="border-gray-200">
          <button @click="handleLogout"
            class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
            {{ $t('navbar.signOut') }}
          </button>
        </div>
      </div>

      <!-- Hamburger button -->
      <button
        @click="toggleMobileMenu"
        class="inline-flex items-center justify-center w-10 h-10 rounded-md text-gray-700 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-indigo-500"
        aria-label="Toggle menu"
      >
        <svg v-if="!isMobileMenuOpen" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
        </svg>
        <svg v-else class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>
  </nav>

  <!-- Mobile menu dropdown -->
  <div
    v-if="isMobileMenuOpen"
    class="md:hidden fixed top-16 left-0 right-0 bg-white border-b border-gray-200 shadow-lg z-40 py-2"
  >
    <div class="flex flex-col">
      <router-link to="/" @click="closeMobileMenu"
        class="text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center">
        {{ $t('navbar.home') }}
      </router-link>
      <button
        class="text-left text-gray-700 px-6 py-3 font-normal transition hover:bg-green-50 hover:text-green-900 min-h-[48px] flex items-center"
        type="button"
        @click="handleQuickPracticeClick(); closeMobileMenu()"
      >
        {{ $t('navbar.quickPractice') }}
      </button>
      <button
        class="text-left text-gray-700 px-6 py-3 font-normal transition hover:bg-purple-50 hover:text-purple-900 min-h-[48px] flex items-center"
        type="button"
        @click="handlePaperPracticeClick(); closeMobileMenu()"
      >
        {{ $t('navbar.pastPaperPractice') }}
      </button>
      <button
        class="text-left text-gray-700 px-6 py-3 font-normal transition hover:bg-blue-50 hover:text-blue-900 min-h-[48px] flex items-center"
        type="button"
        @click="handleExamPaperClick('teacher'); closeMobileMenu()"
      >
        {{ $t('navbar.myExamPapers') }}
      </button>
      <button
        class="text-left text-gray-700 px-6 py-3 font-normal transition hover:bg-blue-50 hover:text-blue-900 min-h-[48px] flex items-center"
        type="button"
        @click="handleExamPaperClick('builder'); closeMobileMenu()"
      >
        {{ $t('navbar.examPaperBuilder') }}
      </button>
      <router-link v-if="userStore.user?.isAdmin" to="/admin" @click="closeMobileMenu"
        class="text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center">
        {{ $t('navbar.systemManagement') }}
      </router-link>
      <button
        class="text-left px-6 py-3 font-normal transition text-yellow-700 hover:bg-yellow-50 min-h-[48px] flex items-center"
        type="button"
        @click="router.push('/donate'); closeMobileMenu()"
      >
        {{ $t('navbar.donate') }}
      </button>
      <div v-if="!userStore.isAuthenticated" class="border-t border-gray-100 mt-1 pt-1">
        <router-link to="/login" @click="closeMobileMenu"
          class="text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center">
          {{ $t('navbar.signIn') }}
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { useUserStore } from '../stores/userStore';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n'

const { locale } = useI18n()
const userStore = useUserStore();
const router = useRouter();

const isDropdownOpen = ref(false);
const isMobileMenuOpen = ref(false);

const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value;
};

const closeDropdown = () => {
  isDropdownOpen.value = false;
};

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value;
  // Close other dropdowns when toggling mobile menu
  isDropdownOpen.value = false;
  langDropdownOpen.value = false;
};

const closeMobileMenu = () => {
  isMobileMenuOpen.value = false;
};

const handleLogout = async () => {
  try {
    await userStore.logout(); // userStore.logout should handle token removal and state reset
    closeDropdown();
    closeMobileMenu();
    router.push('/login');
  } catch (error) {
    console.error('Logout failed:', error);
    // Optionally show an error message to the user
  }
};

// Exam Paper navigation logic
const handleExamPaperClick = (type: 'teacher' | 'builder') => {
  if (type === 'teacher') {
    router.push('/paper/exam/teacher');
  } else {
    router.push('/paper/exam/create');
  }
};

const handleQuickPracticeClick = () => {
  router.push('/practice/quick');
};

const handlePaperPracticeClick = () => {
  router.push('/practice/paper');
};

// Language switcher logic
const languages = [
  { value: 'en', label: 'English' },
  { value: 'zh', label: '中文' },
]
const currentLocale = ref(locale.value)
const langDropdownOpen = ref(false)
const currentLangLabel = computed(() => {
  const found = languages.find(l => l.value === currentLocale.value)
  return found ? found.label : ''
})
function toggleLangDropdown() {
  langDropdownOpen.value = !langDropdownOpen.value
}
function selectLanguage(lang: string) {
  currentLocale.value = lang
  locale.value = lang
  localStorage.setItem('locale', lang)
  langDropdownOpen.value = false
  location.reload()
}
</script>

<style scoped>
/* Tailwind handles all styling */
.top-full {
  top: 100%;
}
</style>
