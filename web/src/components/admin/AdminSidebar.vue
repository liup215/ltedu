<template>
  <aside class="w-64 bg-gray-50 p-4 border-r border-gray-200 h-full overflow-y-auto">
    <nav class="space-y-2">
      <h3 class="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider mb-2">
        {{ $t('adminSidebar.menu') }}
      </h3>
      <router-link
        to="/admin/dashboard"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/home -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
        </svg>
        {{ $t('adminSidebar.dashboard') }}
      </router-link>
      <router-link
        to="/admin/users"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/users -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
        </svg>
        {{ $t('adminSidebar.users') }}
      </router-link>
      <router-link
        to="/admin/teacher-applications"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/academic-cap -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
        </svg>
        {{ $t('adminSidebar.teacherApplications') }}
      </router-link>
      <router-link
        to="/admin/organisations"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/academic-cap -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path d="M12 14l9-5-9-5-9 5 9 5z" />
          <path d="M12 14l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14z" />
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 14l9-5-9-5-9 5 9 5z" />
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20v-6" />
        </svg>
        {{ $t('adminSidebar.organisations') }}
      </router-link>
      <router-link
        to="/admin/qualifications"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/clipboard-document-list -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
        </svg>
        {{ $t('adminSidebar.qualifications') }}
      </router-link>
      <router-link
        to="/admin/syllabuses"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/bookmark -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
        </svg>
        {{ $t('adminSidebar.syllabuses') }}
      </router-link>
      <router-link
        to="/admin/classes"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/office-building -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
        </svg>
        {{ $t('adminSidebar.classes') }}
      </router-link>
      <router-link
        to="/admin/analytics"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900"
        active-class="bg-gray-200 text-gray-900"
      >
        <!-- Heroicon name: outline/chart-bar -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
        </svg>
        {{ $t('adminSidebar.analytics') }}
      </router-link>
      <router-link
        to="/admin/paper-series"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/document -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
        </svg>
        {{ $t('adminSidebar.paperSeries') }}
      </router-link>
      <router-link
        to="/admin/paper-codes"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/document-duplicate -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7v8a2 2 0 002 2h6M8 7V5a2 2 0 012-2h4.586a1 1 0 01.707.293l4.414 4.414a1 1 0 01.293.707V15a2 2 0 01-2 2h-2M8 7H6a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2v-2" />
        </svg>
        {{ $t('adminSidebar.paperCodes') }}
      </router-link>
      <router-link
        to="/admin/past-papers"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/document-text -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
        </svg>
        {{ $t('adminSidebar.pastPapers') }}
      </router-link>
      <router-link
        to="/admin/exam-papers"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/document-report -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17v-2a2 2 0 012-2h2a2 2 0 012 2v2m-6 4h6a2 2 0 002-2V7a2 2 0 00-2-2H7a2 2 0 00-2 2v12a2 2 0 002 2z" />
        </svg>
        {{ $t('adminSidebar.examPapers') }}
      </router-link>
      <router-link
        to="/admin/questions"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/question-mark-circle -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        {{ $t('adminSidebar.questions') }}
      </router-link>
      <router-link
        to="/admin/system-settings"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/cog -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
        </svg>
        {{ $t('adminSidebar.systemSettings') }}
      </router-link>
      <router-link
        to="/admin/roles"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/shield-check -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
        </svg>
        {{ $t('adminSidebar.roles') }}
      </router-link>
      <router-link
        to="/admin/permissions"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900 min-h-[44px]"
        active-class="bg-gray-200 text-gray-900"
        @click="emit('navigate')"
      >
        <!-- Heroicon name: outline/key -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500 shrink-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z" />
        </svg>
        {{ $t('adminSidebar.permissions') }}
      </router-link>
      <router-link
        to="/admin/feedback"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900"
        active-class="bg-gray-200 text-gray-900"
      >
        <!-- Heroicon name: outline/chat -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
        </svg>
        {{ $t('adminSidebar.feedback') }}
      </router-link>
      <router-link
        to="/admin/blog"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md text-gray-700 hover:bg-gray-200 hover:text-gray-900"
        active-class="bg-gray-200 text-gray-900"
      >
        <!-- Heroicon name: outline/newspaper -->
        <svg class="mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-500" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9M7 16h6M7 8h6v4H7V8z" />
        </svg>
        {{ $t('adminSidebar.blog') }}
      </router-link>
    </nav>
  </aside>
</template>

<script setup lang="ts">
const emit = defineEmits<{ navigate: [] }>()
</script>

<style scoped>
/* Tailwind handles styling */
</style>
