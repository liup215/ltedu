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

      <router-link to="/download"
        class="text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" custom
        v-slot="{ navigate, href }">
        <button :href="href" @click="navigate" type="button">{{ $t('navbar.download') }}</button>
      </router-link>

      <router-link to="/help"
        class="text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" custom
        v-slot="{ navigate, href }">
        <button :href="href" @click="navigate" type="button">{{ $t('navbar.docs') }}</button>
      </router-link>

      <router-link
        to="/blog"
        class="px-4 py-2 rounded font-normal text-gray-700 hover:text-indigo-600 hover:bg-gray-100 transition text-sm"
      >
        {{ $t('navbar.blog') }}
      </router-link>

      <!-- Admin Link - visible only to admins -->
      <router-link v-if="userStore.user?.isAdmin" to="/admin"
        class="text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" custom
        v-slot="{ navigate, href }">
        <button :href="href" @click="navigate" type="button">{{ $t('navbar.systemManagement') }}</button>
      </router-link>

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
          <router-link to="/account/analytics" @click="closeDropdown"
            class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left">{{ $t('navbar.usageStats') }}</router-link>
          <router-link to="/account/cli-tokens" @click="closeDropdown"
            class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left">{{ $t('navbar.apiTokens') }}</router-link>
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

    <!-- Mobile: right side controls (user avatar + hamburger) -->
    <div class="flex md:hidden items-center gap-2">
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
          <router-link to="/account/analytics" @click="closeDropdown"
            class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left">{{ $t('navbar.usageStats') }}</router-link>
          <router-link to="/account/cli-tokens" @click="closeDropdown"
            class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left">{{ $t('navbar.apiTokens') }}</router-link>
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
      <router-link to="/download" @click="closeMobileMenu"
        class="text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center">
        {{ $t('navbar.download') }}
      </router-link>
      <router-link to="/help" @click="closeMobileMenu"
        class="text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center">
        {{ $t('navbar.docs') }}
      </router-link>
      <router-link to="/blog" @click="closeMobileMenu"
        class="text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center">
        {{ $t('navbar.blog') }}
      </router-link>
      <router-link v-if="userStore.user?.isAdmin" to="/admin" @click="closeMobileMenu"
        class="text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center">
        {{ $t('navbar.systemManagement') }}
      </router-link>
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
import { ref } from 'vue';
import { useUserStore } from '../stores/userStore';
import { useRouter } from 'vue-router';

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
  isDropdownOpen.value = false;
};

const closeMobileMenu = () => {
  isMobileMenuOpen.value = false;
};

const handleLogout = async () => {
  try {
    await userStore.logout();
    closeDropdown();
    closeMobileMenu();
    router.push('/login');
  } catch (error) {
    console.error('Logout failed:', error);
  }
};
</script>

<style scoped>
/* Tailwind handles all styling */
.top-full {
  top: 100%;
}
</style>

