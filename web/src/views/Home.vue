<template>
  <div class="flex-1 bg-white">
    <!-- Hero Section -->
    <section class="min-h-screen flex flex-col items-center justify-center bg-gradient-to-r from-indigo-600 to-purple-600 text-white">
      <div class="container mx-auto px-6 text-center">
        <h1 class="text-3xl md:text-5xl font-normal mb-6">{{ $t('home.heroTitle') }}</h1>
        <p class="text-lg md:text-xl mb-8">
          {{ $t('home.heroSubtitle') }}
        </p>
        <div class="flex flex-col md:flex-row gap-6 justify-center mb-10">
          <router-link to="/practice/quick">
            <button class="bg-white text-indigo-700 font-normal py-4 px-10 rounded-lg shadow-lg hover:bg-gray-100 transition text-xl">
              {{ $t('home.studentPractice') }}
            </button>
          </router-link>
          <router-link to="/paper/exam/create">
            <button class="bg-white text-purple-700 font-normal py-4 px-10 rounded-lg shadow-lg hover:bg-gray-100 transition text-xl">
              {{ $t('home.teacherBuildExam') }}
            </button>
          </router-link>
        </div>
        <div class="mt-6 text-lg">
          <span class="font-normal">{{ $t('home.supports') }}</span>
          <span class="ml-2">{{ $t('home.supportedSyllabuses') }}</span>
        </div>
      </div>
    </section>

    <!-- Features Section -->
    <section class="py-16">
      <div class="container mx-auto px-6">
        <h2 class="text-3xl font-normal text-center text-gray-800 mb-12">{{ $t('home.platformHighlights') }}</h2>
        <div class="grid md:grid-cols-3 gap-8">
          <!-- Feature 1: Adaptive Practice -->
          <div class="bg-white p-6 rounded-lg shadow-lg">
            <div class="text-indigo-600 mb-4">
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
          <div class="bg-white p-6 rounded-lg shadow-lg">
            <div class="text-purple-600 mb-4">
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
          <div class="bg-white p-6 rounded-lg shadow-lg">
            <div class="text-green-600 mb-4">
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
