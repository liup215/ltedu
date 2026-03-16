<template>
  <div class="p-6 max-w-5xl mx-auto">
    <header class="mb-6">
      <div class="flex justify-between items-start">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ t('knowledgePoint.migration.title') }}</h1>
          <p class="mt-2 text-sm text-gray-600">{{ t('knowledgePoint.migration.subtitle') }}</p>
          <span class="mt-2 inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-purple-100 text-purple-800">
            {{ t('knowledgePoint.migration.adminOnly') }}
          </span>
        </div>
        <router-link 
          to="/admin/syllabuses"
          class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
        >
          {{ t('knowledgePoint.backToSyllabuses') }}
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
          <h3 class="text-sm font-medium text-yellow-800">{{ t('knowledgePoint.migration.warning') }}</h3>
          <p class="mt-2 text-sm text-yellow-700">{{ t('knowledgePoint.migration.warningMessage') }}</p>
        </div>
      </div>
    </div>

    <!-- Main Form -->
    <div class="bg-white shadow rounded-lg mb-8">
      <div class="px-6 py-4 border-b border-gray-200">
        <h2 class="text-lg font-medium text-gray-900">{{ t('knowledgePoint.migration.options') }}</h2>
      </div>
      
      <form @submit.prevent="startMigration" class="px-6 py-4 space-y-6">
        <!-- Syllabus Selection -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            {{ t('knowledgePoint.migration.selectSyllabus') }} *
          </label>
          <select
            v-model="selectedSyllabusId"
            required
            class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 py-3"
          >
            <option value="">{{ t('knowledgePoint.migration.syllabusPlaceholder') }}</option>
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
                {{ t('knowledgePoint.migration.generateKeypoints') }}
              </label>
              <p class="text-gray-500">{{ t('knowledgePoint.migration.generateKeypointsTip') }}</p>
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
                {{ t('knowledgePoint.migration.linkQuestions') }}
              </label>
              <p class="text-gray-500">{{ t('knowledgePoint.migration.linkQuestionsTip') }}</p>
            </div>
          </div>

          <div v-if="options.linkQuestions">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              {{ t('knowledgePoint.migration.batchSize') }}
            </label>
            <input
              type="number"
              v-model.number="options.batchSize"
              min="1"
              max="100"
              class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 py-3"
            />
            <p class="mt-1 text-sm text-gray-500">{{ t('knowledgePoint.migration.batchSizeTip') }}</p>
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
              {{ t('knowledgePoint.migration.confirm') }}
            </label>
          </div>
        </div>

        <!-- Submit Button -->
        <div class="flex justify-end">
          <button
            type="submit"
            :disabled="submitting || !confirmed || !selectedSyllabusId"
            class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <svg v-if="submitting" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ t('knowledgePoint.migration.startMigration') }}
          </button>
        </div>
      </form>
    </div>

    <!-- Migration Jobs List -->
    <div class="bg-white shadow rounded-lg">
      <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
        <h2 class="text-lg font-medium text-gray-900">{{ t('knowledgePoint.migration.jobList') }}</h2>
        <button
          @click="loadJobs"
          class="inline-flex items-center px-3 py-1.5 border border-gray-300 rounded-md text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
        >
          <svg class="h-4 w-4 mr-1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z" clip-rule="evenodd" />
          </svg>
          {{ t('knowledgePoint.migration.refreshJobs') }}
        </button>
      </div>

      <div v-if="loadingJobs" class="px-6 py-8 text-center text-gray-500">
        <svg class="animate-spin h-6 w-6 mx-auto mb-2 text-indigo-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
      </div>

      <div v-else-if="jobs.length === 0" class="px-6 py-8 text-center text-gray-500">
        {{ t('knowledgePoint.migration.noJobs') }}
      </div>

      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('knowledgePoint.migration.jobId') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('knowledgePoint.migration.jobSyllabus') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('knowledgePoint.migration.jobStatus') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('knowledgePoint.migration.jobProgress') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('knowledgePoint.migration.jobCreatedAt') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('knowledgePoint.migration.jobActions') }}</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="job in jobs" :key="job.id">
              <td class="px-4 py-3 text-sm text-gray-900">{{ job.id }}</td>
              <td class="px-4 py-3 text-sm text-gray-900">{{ getSyllabusName(job.syllabusId) }}</td>
              <td class="px-4 py-3 text-sm">
                <span :class="statusBadgeClass(job.status)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                  <span v-if="job.status === 'running'" class="mr-1">
                    <svg class="animate-spin h-3 w-3" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                  </span>
                  {{ statusLabel(job.status) }}
                </span>
              </td>
              <td class="px-4 py-3 text-sm text-gray-900">
                <div class="flex items-center space-x-2">
                  <div class="flex-1 bg-gray-200 rounded-full h-2 w-24">
                    <div
                      class="h-2 rounded-full transition-all duration-300"
                      :class="job.status === 'completed' ? 'bg-green-500' : job.status === 'failed' ? 'bg-red-500' : 'bg-indigo-500'"
                      :style="{ width: job.progress + '%' }"
                    ></div>
                  </div>
                  <span class="text-xs text-gray-500 w-10">{{ job.progress }}%</span>
                </div>
              </td>
              <td class="px-4 py-3 text-sm text-gray-500">{{ formatDate(job.createdAt) }}</td>
              <td class="px-4 py-3 text-sm space-x-2">
                <button
                  @click="viewJobDetail(job)"
                  class="text-indigo-600 hover:text-indigo-900 text-xs font-medium"
                >
                  {{ t('knowledgePoint.migration.jobViewDetail') }}
                </button>
                <button
                  v-if="job.status === 'failed'"
                  @click="retryJob(job.id)"
                  class="text-green-600 hover:text-green-900 text-xs font-medium"
                >
                  {{ t('knowledgePoint.migration.jobRetry') }}
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Job Detail Modal -->
    <div v-if="selectedJob" class="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center z-50" @click.self="selectedJob = null">
      <div class="bg-white rounded-lg shadow-xl max-w-lg w-full mx-4 p-6">
        <div class="flex justify-between items-start mb-4">
          <h3 class="text-lg font-medium text-gray-900">{{ t('knowledgePoint.migration.jobDetail') }} #{{ selectedJob.id }}</h3>
          <button @click="selectedJob = null" class="text-gray-400 hover:text-gray-600">
            <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
              <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
            </svg>
          </button>
        </div>

        <dl class="space-y-3 text-sm">
          <div class="flex justify-between">
            <dt class="text-gray-500">{{ t('knowledgePoint.migration.jobStatus') }}</dt>
            <dd><span :class="statusBadgeClass(selectedJob.status)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">{{ statusLabel(selectedJob.status) }}</span></dd>
          </div>
          <div class="flex justify-between">
            <dt class="text-gray-500">{{ t('knowledgePoint.migration.jobProgress') }}</dt>
            <dd class="text-gray-900">{{ selectedJob.progress }}% ({{ selectedJob.doneItems }}/{{ selectedJob.totalItems }})</dd>
          </div>
          <div v-if="selectedJob.startedAt" class="flex justify-between">
            <dt class="text-gray-500">{{ t('knowledgePoint.migration.jobStartedAt') }}</dt>
            <dd class="text-gray-900">{{ formatDate(selectedJob.startedAt) }}</dd>
          </div>
          <div v-if="selectedJob.completedAt" class="flex justify-between">
            <dt class="text-gray-500">{{ t('knowledgePoint.migration.jobCompletedAt') }}</dt>
            <dd class="text-gray-900">{{ formatDate(selectedJob.completedAt) }}</dd>
          </div>
          <div v-if="selectedJob.errorMessage" class="mt-2">
            <dt class="text-gray-500 mb-1">{{ t('knowledgePoint.migration.jobErrorMessage') }}</dt>
            <dd class="bg-red-50 border-l-4 border-red-400 p-3 text-red-700 text-xs">{{ selectedJob.errorMessage }}</dd>
          </div>
          <div v-if="selectedJobReport" class="mt-2">
            <dt class="text-gray-500 mb-2">{{ t('knowledgePoint.migration.jobReport') }}</dt>
            <dd>
              <div class="grid grid-cols-3 gap-3">
                <div class="bg-green-50 rounded p-2 text-center">
                  <div class="text-lg font-semibold text-green-600">{{ selectedJobReport.generatedKeypoints }}</div>
                  <div class="text-xs text-gray-500">{{ t('knowledgePoint.migration.jobDetailGeneratedKeypoints') }}</div>
                </div>
                <div class="bg-blue-50 rounded p-2 text-center">
                  <div class="text-lg font-semibold text-blue-600">{{ selectedJobReport.linkedQuestions }}</div>
                  <div class="text-xs text-gray-500">{{ t('knowledgePoint.migration.jobDetailLinkedQuestions') }}</div>
                </div>
                <div class="bg-purple-50 rounded p-2 text-center">
                  <div class="text-lg font-semibold text-purple-600">{{ selectedJobReport.totalLinks }}</div>
                  <div class="text-xs text-gray-500">{{ t('knowledgePoint.migration.jobDetailTotalLinks') }}</div>
                </div>
              </div>
              <div v-if="selectedJobReport.errors?.length" class="mt-3">
                <div class="text-xs font-medium text-gray-500 mb-1">{{ t('knowledgePoint.migration.jobDetailErrors') }}</div>
                <div class="bg-red-50 border-l-4 border-red-400 p-3 max-h-40 overflow-y-auto">
                  <p v-for="(err, i) in selectedJobReport.errors" :key="i" class="text-xs text-red-700">{{ err }}</p>
                </div>
              </div>
            </dd>
          </div>
        </dl>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { push } from 'notivue'
import knowledgePointService from '../../services/knowledgePointService'
import syllabusService from '../../services/syllabusService'
import type { MigrationJob, MigrateReport } from '../../models/knowledgePoint.model'
import type { Syllabus } from '../../models/syllabus.model'

const route = useRoute()
const { t } = useI18n()
const syllabuses = ref<Syllabus[]>([])
const selectedSyllabusId = ref<number | null>(null)
const submitting = ref(false)
const confirmed = ref(false)
const jobs = ref<MigrationJob[]>([])
const loadingJobs = ref(false)
const selectedJob = ref<MigrationJob | null>(null)

const options = ref({
  generateKeypoints: true,
  linkQuestions: false,
  batchSize: 50
})

let pollTimer: ReturnType<typeof setInterval> | null = null

onMounted(async () => {
  await Promise.all([loadSyllabuses(), loadJobs()])

  // Check if syllabusId is provided in query parameters
  const querySyllabusId = route.query.syllabusId
  if (querySyllabusId && !isNaN(Number(querySyllabusId))) {
    selectedSyllabusId.value = Number(querySyllabusId)
  }

  // Poll for job updates every 5 seconds if there are running/pending jobs
  pollTimer = setInterval(pollActiveJobs, 5000)
})

onUnmounted(() => {
  if (pollTimer) clearInterval(pollTimer)
})

async function pollActiveJobs() {
  const hasActive = jobs.value.some(j => j.status === 'pending' || j.status === 'running')
  if (hasActive) {
    await loadJobs()
  }
}

async function loadSyllabuses() {
  try {
    const response = await syllabusService.getSyllabuses({ pageIndex: 1, pageSize: 100 })
    if (response.code === 0) {
      syllabuses.value = response.data.list
    }
  } catch (error) {
    console.error('Failed to load syllabuses:', error)
  }
}

async function loadJobs() {
  loadingJobs.value = true
  try {
    const response = await knowledgePointService.listMigrationJobs({ pageIndex: 1, pageSize: 50 })
    if (response.code === 0) {
      jobs.value = response.data.list
    }
  } catch (error) {
    console.error('Failed to load migration jobs:', error)
    push.error(t('knowledgePoint.migration.loadJobsError'))
  } finally {
    loadingJobs.value = false
  }
}

async function startMigration() {
  if (!selectedSyllabusId.value) return

  try {
    submitting.value = true

    const response = await knowledgePointService.createMigrationJob({
      syllabusId: selectedSyllabusId.value,
      options: options.value
    })

    if (response.code === 0) {
      push.success(t('knowledgePoint.migration.migrationSuccess'))
      confirmed.value = false
      await loadJobs()
    } else {
      throw new Error(response.msg)
    }
  } catch (error: any) {
    console.error('Failed to create migration job:', error)
    push.error(error.response?.data?.msg || t('knowledgePoint.migration.migrationError'))
  } finally {
    submitting.value = false
  }
}

async function retryJob(id: number) {
  try {
    const response = await knowledgePointService.retryMigrationJob(id)
    if (response.code === 0) {
      push.success(t('knowledgePoint.migration.retrySuccess'))
      await loadJobs()
    } else {
      throw new Error(response.msg)
    }
  } catch (error: any) {
    push.error(error.response?.data?.msg || t('knowledgePoint.migration.retryError'))
  }
}

function viewJobDetail(job: MigrationJob) {
  selectedJob.value = job
}

function getSyllabusName(syllabusId: number): string {
  const s = syllabuses.value.find(s => s.id === syllabusId)
  return s ? `${s.name} (${s.code})` : `#${syllabusId}`
}

function statusLabel(status: string): string {
  const map: Record<string, string> = {
    pending: t('knowledgePoint.migration.statusPending'),
    running: t('knowledgePoint.migration.statusRunning'),
    completed: t('knowledgePoint.migration.statusCompleted'),
    failed: t('knowledgePoint.migration.statusFailed')
  }
  return map[status] ?? status
}

function statusBadgeClass(status: string): string {
  const map: Record<string, string> = {
    pending: 'bg-yellow-100 text-yellow-800',
    running: 'bg-blue-100 text-blue-800',
    completed: 'bg-green-100 text-green-800',
    failed: 'bg-red-100 text-red-800'
  }
  return map[status] ?? 'bg-gray-100 text-gray-800'
}

function formatDate(dateStr?: string): string {
  if (!dateStr) return '-'
  return new Date(dateStr).toLocaleString()
}

function parsedReport(job: MigrationJob): MigrateReport | null {
  if (!job.report) return null
  try {
    return JSON.parse(job.report) as MigrateReport
  } catch {
    return null
  }
}

const selectedJobReport = computed<MigrateReport | null>(() => {
  if (!selectedJob.value) return null
  return parsedReport(selectedJob.value)
})
</script>

