<template>
  <div class="flex-1 bg-white">
    <!-- Hero Section -->
    <section class="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-sky-400 to-blue-600 text-white">
      <div class="container mx-auto px-6 text-center">
        <!-- AI Agent Badge -->
        <div class="inline-flex items-center gap-2 bg-white/20 backdrop-blur-sm border border-white/30 rounded-full px-4 py-2 mb-6 text-sm font-medium">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17H3a2 2 0 01-2-2V5a2 2 0 012-2h14a2 2 0 012 2v3"></path>
          </svg>
          {{ $t('home.aiAgentBadge') }}
        </div>
        <h1 class="text-3xl md:text-5xl font-normal mb-6">{{ $t('home.heroTitle') }}</h1>
        <p class="text-lg md:text-xl mb-8 max-w-2xl mx-auto opacity-90">
          {{ $t('home.heroSubtitle') }}
        </p>
        <div class="flex flex-col md:flex-row gap-6 justify-center mb-10">
          <router-link to="/practice/quick">
            <button class="bg-white text-blue-700 font-normal py-4 px-10 rounded-lg shadow-lg hover:bg-blue-50 transition text-xl">
              {{ $t('home.studentPractice') }}
            </button>
          </router-link>
          <router-link to="/paper/exam/create">
            <button class="bg-white text-blue-700 font-normal py-4 px-10 rounded-lg shadow-lg hover:bg-blue-50 transition text-xl">
              {{ $t('home.teacherBuildExam') }}
            </button>
          </router-link>
        </div>
        <div class="mt-6 text-lg opacity-80">
          <span class="font-normal">{{ $t('home.supports') }}</span>
          <span class="ml-2">{{ $t('home.supportedSyllabuses') }}</span>
        </div>
      </div>
    </section>

    <!-- AI Agent Section -->
    <section class="py-16 bg-sky-50">
      <div class="container mx-auto px-6">
        <div class="text-center mb-12">
          <h2 class="text-3xl font-normal text-gray-800 mb-4">{{ $t('home.aiAgentTitle') }}</h2>
          <p class="text-lg text-gray-600 max-w-2xl mx-auto">{{ $t('home.aiAgentSubtitle') }}</p>
        </div>
        <div class="grid md:grid-cols-3 gap-8">
          <!-- AI Feature 1: API First -->
          <div class="bg-white p-6 rounded-xl shadow-md border border-sky-100">
            <div class="text-sky-500 mb-4">
              <svg class="w-12 h-12 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"></path>
              </svg>
            </div>
            <h3 class="text-xl font-normal text-gray-800 mb-3">{{ $t('home.aiFeatureApiFirst') }}</h3>
            <p class="text-gray-600">{{ $t('home.aiFeatureApiFirstDesc') }}</p>
          </div>
          <!-- AI Feature 2: CLI Tool -->
          <div class="bg-white p-6 rounded-xl shadow-md border border-sky-100">
            <div class="text-sky-500 mb-4">
              <svg class="w-12 h-12 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </div>
            <h3 class="text-xl font-normal text-gray-800 mb-3">{{ $t('home.aiFeatureCli') }}</h3>
            <p class="text-gray-600">{{ $t('home.aiFeatureCliDesc') }}</p>
          </div>
          <!-- AI Feature 3: Conversational AI -->
          <div class="bg-white p-6 rounded-xl shadow-md border border-sky-100">
            <div class="text-sky-500 mb-4">
              <svg class="w-12 h-12 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z"></path>
              </svg>
            </div>
            <h3 class="text-xl font-normal text-gray-800 mb-3">{{ $t('home.aiFeatureChat') }}</h3>
            <p class="text-gray-600">{{ $t('home.aiFeatureChatDesc') }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- edu-cli Section -->
    <section class="py-16 bg-gray-900 text-white">
      <div class="container mx-auto px-6">
        <div class="flex flex-col md:flex-row items-center gap-12">
          <div class="flex-1">
            <div class="inline-flex items-center gap-2 bg-sky-500/20 border border-sky-400/30 rounded-full px-4 py-1.5 mb-6 text-sky-300 text-sm font-medium">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
              edu-cli
            </div>
            <h2 class="text-3xl font-normal mb-4">{{ $t('home.cliTitle') }}</h2>
            <p class="text-gray-300 text-lg mb-6">{{ $t('home.cliSubtitle') }}</p>
            <ul class="space-y-3 mb-8">
              <li class="flex items-center gap-3 text-gray-300">
                <svg class="w-5 h-5 text-sky-400 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                </svg>
                {{ $t('home.cliFeature1') }}
              </li>
              <li class="flex items-center gap-3 text-gray-300">
                <svg class="w-5 h-5 text-sky-400 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                </svg>
                {{ $t('home.cliFeature2') }}
              </li>
              <li class="flex items-center gap-3 text-gray-300">
                <svg class="w-5 h-5 text-sky-400 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                </svg>
                {{ $t('home.cliFeature3') }}
              </li>
            </ul>
            <router-link to="/download">
              <button class="bg-sky-500 hover:bg-sky-400 text-white font-normal py-3 px-8 rounded-lg shadow-lg transition text-lg">
                {{ $t('home.cliDownload') }}
              </button>
            </router-link>
          </div>
          <div class="flex-1">
            <div class="bg-gray-800 rounded-xl p-6 font-mono text-sm border border-gray-700 shadow-2xl">
              <div class="flex items-center gap-2 mb-4">
                <span class="w-3 h-3 rounded-full bg-red-500"></span>
                <span class="w-3 h-3 rounded-full bg-yellow-500"></span>
                <span class="w-3 h-3 rounded-full bg-green-500"></span>
                <span class="ml-2 text-gray-500 text-xs">edu-cli</span>
              </div>
              <div class="space-y-2">
                <div><span class="text-sky-400">$</span> <span class="text-green-300">edu</span> <span class="text-white">login</span></div>
                <div class="text-gray-400 pl-4">✓ {{ $t('home.cliDemo1') }}</div>
                <div class="mt-3"><span class="text-sky-400">$</span> <span class="text-green-300">edu</span> <span class="text-white">practice --syllabus cambridge-a-level</span></div>
                <div class="text-gray-400 pl-4">✓ {{ $t('home.cliDemo2') }}</div>
                <div class="mt-3"><span class="text-sky-400">$</span> <span class="text-green-300">edu</span> <span class="text-white">paper create --topic algebra</span></div>
                <div class="text-gray-400 pl-4">✓ {{ $t('home.cliDemo3') }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Platform Features Section -->
    <section class="py-16">
      <div class="container mx-auto px-6">
        <h2 class="text-3xl font-normal text-center text-gray-800 mb-12">{{ $t('home.platformHighlights') }}</h2>
        <div class="grid md:grid-cols-3 gap-8">
          <!-- Feature 1: Adaptive Practice -->
          <div class="bg-white p-6 rounded-xl shadow-md border border-gray-100">
            <div class="text-blue-500 mb-4">
              <svg class="w-12 h-12 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
              </svg>
            </div>
            <h3 class="text-xl font-normal text-gray-800 mb-3">{{ $t('home.adaptivePractice') }}</h3>
            <p class="text-gray-600">
              {{ $t('home.adaptivePracticeDesc') }}
            </p>
          </div>
          <!-- Feature 2: Exam Paper Builder -->
          <div class="bg-white p-6 rounded-xl shadow-md border border-gray-100">
            <div class="text-blue-500 mb-4">
              <svg class="w-12 h-12 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M9 12h6m-6 4h6m-6-8h6"></path>
              </svg>
            </div>
            <h3 class="text-xl font-normal text-gray-800 mb-3">{{ $t('home.flexibleExamBuilder') }}</h3>
            <p class="text-gray-600">
              {{ $t('home.flexibleExamBuilderDesc') }}
            </p>
          </div>
          <!-- Feature 3: Syllabus Coverage -->
          <div class="bg-white p-6 rounded-xl shadow-md border border-gray-100">
            <div class="text-green-500 mb-4">
              <svg class="w-12 h-12 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M5 13l4 4L19 7"></path>
              </svg>
            </div>
            <h3 class="text-xl font-normal text-gray-800 mb-3">{{ $t('home.internationalSyllabusSupport') }}</h3>
            <p class="text-gray-600">
              {{ $t('home.internationalSyllabusSupportDesc') }}
            </p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// No logic needed for static home page
</script>

<style scoped>
/* Tailwind handles all styling */
</style>
