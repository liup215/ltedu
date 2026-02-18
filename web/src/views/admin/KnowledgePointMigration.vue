<template>
  <div class="p-6 max-w-4xl mx-auto">
    <header class="mb-6">
      <div class="flex justify-between items-start">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ $t('knowledgePoint.migration.title') }}</h1>
          <p class="mt-2 text-sm text-gray-600">{{ $t('knowledgePoint.migration.subtitle') }}</p>
          <span class="mt-2 inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-purple-100 text-purple-800">
            {{ $t('knowledgePoint.migration.adminOnly') }}
          </span>
        </div>
        <router-link 
          to="/admin/syllabuses"
          class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
        >
          {{ $t('knowledgePoint.backToSyllabuses') }}
        </router-link>
      </div>
    </header>

    <!-- Warning Banner -->
    <div class="bg-yellow-50 border-l-4 border-yellow-400 p-4 mb-6">
      <div class="flex">
        <div class="flex-shrink-0">
          <svg class="h-5 w-5 text-yellow-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
          </svg>
        </div>
        <div class="ml-3">
          <h3 class="text-sm font-medium text-yellow-800">{{ $t('knowledgePoint.migration.warning') }}</h3>
          <p class="mt-2 text-sm text-yellow-700">{{ $t('knowledgePoint.migration.warningMessage') }}</p>
        </div>
      </div>
    </div>

    <!-- Main Form -->
    <div class="bg-white shadow rounded-lg">
      <div class="px-6 py-4 border-b border-gray-200">
        <h2 class="text-lg font-medium text-gray-900">{{ $t('knowledgePoint.migration.options') }}</h2>
      </div>
      
      <form @submit.prevent="startMigration" class="px-6 py-4 space-y-6">
        <!-- Syllabus Selection -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            {{ $t('knowledgePoint.migration.selectSyllabus') }} *
          </label>
          <select
            v-model="selectedSyllabusId"
            required
            class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 py-3"
          >
            <option value="">{{ $t('knowledgePoint.migration.syllabusPlaceholder') }}</option>
            <option v-for="syllabus in syllabuses" :key="syllabus.id" :value="syllabus.id">
              {{ syllabus.name }} ({{ syllabus.code }})
            </option>
          </select>
        </div>

        <!-- Migration Options -->
        <div class="space-y-4">
          <div class="relative flex items-start">
            <div class="flex items-center h-5">
              <input
                id="generateKeypoints"
                v-model="options.generateKeypoints"
                type="checkbox"
                class="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 rounded"
              />
            </div>
            <div class="ml-3 text-sm">
              <label for="generateKeypoints" class="font-medium text-gray-700">
                {{ $t('knowledgePoint.migration.generateKeypoints') }}
              </label>
              <p class="text-gray-500">{{ $t('knowledgePoint.migration.generateKeypointsTip') }}</p>
            </div>
          </div>

          <div class="relative flex items-start">
            <div class="flex items-center h-5">
              <input
                id="linkQuestions"
                v-model="options.linkQuestions"
                type="checkbox"
                class="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 rounded"
              />
            </div>
            <div class="ml-3 text-sm">
              <label for="linkQuestions" class="font-medium text-gray-700">
                {{ $t('knowledgePoint.migration.linkQuestions') }}
              </label>
              <p class="text-gray-500">{{ $t('knowledgePoint.migration.linkQuestionsTip') }}</p>
            </div>
          </div>

          <div v-if="options.linkQuestions">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              {{ $t('knowledgePoint.migration.batchSize') }}
            </label>
            <input
              type="number"
              v-model.number="options.batchSize"
              min="1"
              max="100"
              class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 py-3"
            />
            <p class="mt-1 text-sm text-gray-500">{{ $t('knowledgePoint.migration.batchSizeTip') }}</p>
          </div>
        </div>

        <!-- Confirmation Checkbox -->
        <div class="relative flex items-start">
          <div class="flex items-center h-5">
            <input
              id="confirmed"
              v-model="confirmed"
              type="checkbox"
              required
              class="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 rounded"
            />
          </div>
          <div class="ml-3 text-sm">
            <label for="confirmed" class="font-medium text-red-700">
              {{ $t('knowledgePoint.migration.confirm') }}
            </label>
          </div>
        </div>

        <!-- Submit Button -->
        <div class="flex justify-end">
          <button
            type="submit"
            :disabled="migrating || !confirmed || !selectedSyllabusId"
            class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <svg v-if="migrating" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ migrating ? $t('knowledgePoint.migration.migrating') : $t('knowledgePoint.migration.startMigration') }}
          </button>
        </div>
      </form>
    </div>

    <!-- Migration Report -->
    <div v-if="report" class="mt-6 bg-white shadow rounded-lg">
      <div class="px-6 py-4 border-b border-gray-200">
        <h2 class="text-lg font-medium text-gray-900">{{ $t('knowledgePoint.migration.migrationReport') }}</h2>
      </div>
      <div class="px-6 py-4">
        <div class="grid grid-cols-1 gap-4 sm:grid-cols-3">
          <div class="bg-green-50 rounded-lg p-4">
            <dt class="text-sm font-medium text-gray-500">{{ $t('knowledgePoint.migration.generatedKeypoints') }}</dt>
            <dd class="mt-1 text-3xl font-semibold text-green-600">{{ report.generatedKeypoints }}</dd>
          </div>
          <div class="bg-blue-50 rounded-lg p-4">
            <dt class="text-sm font-medium text-gray-500">{{ $t('knowledgePoint.migration.linkedQuestions') }}</dt>
            <dd class="mt-1 text-3xl font-semibold text-blue-600">{{ report.linkedQuestions }}</dd>
          </div>
          <div class="bg-purple-50 rounded-lg p-4">
            <dt class="text-sm font-medium text-gray-500">{{ $t('knowledgePoint.migration.totalLinks') }}</dt>
            <dd class="mt-1 text-3xl font-semibold text-purple-600">{{ report.totalLinks }}</dd>
          </div>
        </div>

        <!-- Errors Section -->
        <div v-if="report.errors && report.errors.length > 0" class="mt-4">
          <h3 class="text-sm font-medium text-gray-700 mb-2">{{ $t('knowledgePoint.migration.errors') }}</h3>
          <div class="bg-red-50 border-l-4 border-red-400 p-4">
            <div class="text-sm text-red-700 space-y-1">
              <p v-for="(error, index) in report.errors" :key="index">{{ error }}</p>
            </div>
          </div>
        </div>
        <div v-else class="mt-4">
          <div class="bg-green-50 border-l-4 border-green-400 p-4">
            <p class="text-sm text-green-700">{{ $t('knowledgePoint.migration.noErrors') }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { push } from 'notivue'
import knowledgePointService from '../../services/knowledgePointService'
import syllabusService from '../../services/syllabusService'
import type { MigrateReport } from '../../models/knowledgePoint.model'
import type { Syllabus } from '../../models/syllabus.model'

const syllabuses = ref<Syllabus[]>([])
const selectedSyllabusId = ref<number | null>(null)
const migrating = ref(false)
const confirmed = ref(false)
const report = ref<MigrateReport | null>(null)

const options = ref({
  generateKeypoints: true,
  linkQuestions: false,
  batchSize: 50
})

onMounted(async () => {
  await loadSyllabuses()
})

async function loadSyllabuses() {
  try {
    const response = await syllabusService.list({ page: 1, pageSize: 100 })
    if (response.code === 200) {
      syllabuses.value = response.data.list
    }
  } catch (error) {
    console.error('Failed to load syllabuses:', error)
    push.error('Failed to load syllabuses')
  }
}

async function startMigration() {
  if (!selectedSyllabusId.value) return
  
  try {
    migrating.value = true
    report.value = null
    
    const response = await knowledgePointService.migrateSyllabus({
      syllabusId: selectedSyllabusId.value,
      options: options.value
    })
    
    if (response.code === 200) {
      report.value = response.data
      push.success($t('knowledgePoint.migration.migrationSuccess'))
    } else {
      throw new Error(response.msg)
    }
  } catch (error: any) {
    console.error('Migration failed:', error)
    push.error(error.response?.data?.msg || $t('knowledgePoint.migration.migrationError'))
  } finally {
    migrating.value = false
  }
}

function $t(key: string): string {
  // Placeholder for i18n translation
  return key
}
</script>
