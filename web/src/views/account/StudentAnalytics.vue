<template>
  <div class="max-w-4xl mx-auto py-8 px-4">
    <header class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">{{ t('analytics.recommendations') }}</h1>
      <p class="mt-1 text-sm text-gray-600">{{ t('analytics.subtitle') }}</p>
    </header>

    <!-- Goal Selector -->
    <div class="mb-6 bg-white rounded-lg shadow p-4 flex items-center gap-4">
      <label class="text-sm font-medium text-gray-700 whitespace-nowrap">Goal</label>
      <select
        v-model="selectedGoalId"
        @change="loadData"
        class="border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 flex-1 max-w-xs"
      >
        <option value="">-- Select a goal --</option>
        <option v-for="goal in goals" :key="goal.id" :value="goal.id">{{ goal.syllabusName || `Goal #${goal.id}` }}</option>
      </select>
    </div>

    <div v-if="!selectedGoalId" class="text-center py-12 text-gray-400">
      <p>{{ t('analytics.goalRequired') }}</p>
    </div>

    <div v-else-if="loading" class="text-center py-8 text-gray-500">{{ t('analytics.loading') }}</div>

    <div v-else class="space-y-6">
      <!-- Personal Progress Summary -->
      <div v-if="studentSummary" class="bg-white rounded-lg shadow p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">Your Progress</h2>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div class="text-center">
            <div class="text-2xl font-bold" :class="masteryColor(studentSummary.masteryLevel)">{{ studentSummary.masteryLevel.toFixed(1) }}%</div>
            <div class="text-xs text-gray-500 mt-1">{{ t('analytics.mastery') }}</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-blue-600">{{ studentSummary.coverageLevel.toFixed(1) }}%</div>
            <div class="text-xs text-gray-500 mt-1">{{ t('analytics.coverage') }}</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold" :class="masteryColor(studentSummary.accuracyRate)">{{ studentSummary.accuracyRate.toFixed(1) }}%</div>
            <div class="text-xs text-gray-500 mt-1">{{ t('analytics.accuracy') }}</div>
          </div>
          <div class="text-center">
            <div class="text-2xl font-bold text-gray-700">{{ studentSummary.totalAttempts }}</div>
            <div class="text-xs text-gray-500 mt-1">{{ t('analytics.totalAttempts') }}</div>
          </div>
        </div>

        <!-- At-risk warning -->
        <div v-if="studentSummary.isAtRisk" class="mt-4 p-3 bg-yellow-50 border border-yellow-200 rounded-md">
          <p class="text-sm font-medium text-yellow-800 mb-1">⚠ Areas to improve:</p>
          <ul class="text-sm text-yellow-700 space-y-0.5">
            <li v-for="(r, i) in studentSummary.riskReasons" :key="i">• {{ r }}</li>
          </ul>
        </div>
      </div>

      <!-- AI Recommendations -->
      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-semibold text-gray-900">{{ t('analytics.recommendations') }}</h2>
          <div class="flex items-center gap-3 text-sm text-gray-500">
            <span v-if="recommendations" class="flex items-center gap-1 text-orange-600">
              <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
              {{ recommendations.reviewCount }} {{ t('analytics.reviewDue') }}
            </span>
            <span v-if="recommendations" class="flex items-center gap-1 text-red-500">
              <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
              {{ recommendations.gapCount }} {{ t('analytics.knowledgeGaps') }}
            </span>
          </div>
        </div>

        <div v-if="!recommendations || !recommendations.questions.length" class="text-center py-6 text-gray-400">
          {{ t('analytics.noRecommendations') }}
        </div>

        <div v-else class="space-y-3">
          <div
            v-for="rec in recommendations.questions"
            :key="rec.chapterId"
            class="flex items-center justify-between p-3 rounded-lg border"
            :class="priorityBorderClass(rec.priority)"
          >
            <div class="flex-1 min-w-0 mr-3">
              <div class="flex items-center gap-2">
                <span :class="['text-xs px-2 py-0.5 rounded-full font-medium', priorityBadgeClass(rec.priority)]">
                  P{{ rec.priority }}
                </span>
                <span class="font-medium text-gray-900 truncate">{{ rec.chapterName }}</span>
              </div>
              <p class="mt-0.5 text-sm text-gray-500">{{ rec.reason }}</p>
              <div class="mt-1 flex items-center gap-2">
                <div class="w-20 bg-gray-200 rounded-full h-1.5">
                  <div class="h-1.5 rounded-full bg-indigo-500" :style="{ width: rec.masteryLevel + '%' }"></div>
                </div>
                <span class="text-xs text-gray-400">{{ rec.masteryLevel.toFixed(0) }}% mastery</span>
              </div>
            </div>
            <router-link
              :to="`/practice/quick?chapterId=${rec.chapterId}`"
              class="flex-shrink-0 inline-flex items-center px-3 py-1.5 border border-transparent rounded-md text-xs font-medium text-white bg-indigo-600 hover:bg-indigo-700"
            >
              {{ t('analytics.studyNow') }}
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import analyticsService, {
  type StudentPerformanceSummary,
  type RecommendationResponse
} from '../../services/analyticsService'
import goalService from '../../services/goalService'

const { t } = useI18n()

const goals = ref<Array<{ id: number; syllabusName: string }>>([])
const selectedGoalId = ref<number | ''>('')
const loading = ref(false)
const studentSummary = ref<StudentPerformanceSummary | null>(null)
const recommendations = ref<RecommendationResponse | null>(null)

onMounted(async () => {
  try {
    const res = await goalService.getActiveGoals()
    if (res.code === 0 && res.data) {
      goals.value = (Array.isArray(res.data) ? res.data : [res.data]).map((g: any) => ({
        id: g.id,
        syllabusName: g.syllabus?.name || `Goal #${g.id}`
      }))
      if (goals.value.length === 1) {
        selectedGoalId.value = goals.value[0].id
        await loadData()
      }
    }
  } catch {
    // ignore
  }
})

async function loadData() {
  if (!selectedGoalId.value) return
  loading.value = true
  const id = Number(selectedGoalId.value)
  try {
    const [summaryRes, recsRes] = await Promise.all([
      analyticsService.getStudentSummary(id),
      analyticsService.getRecommendations(id)
    ])
    if (summaryRes.code === 0) studentSummary.value = summaryRes.data
    if (recsRes.code === 0) recommendations.value = recsRes.data
  } catch {
    // ignore
  } finally {
    loading.value = false
  }
}

function masteryColor(value: number) {
  if (value >= 70) return 'text-green-600'
  if (value >= 40) return 'text-yellow-600'
  return 'text-red-600'
}

function priorityBorderClass(priority: number) {
  if (priority === 1) return 'border-red-300 bg-red-50'
  if (priority === 2) return 'border-orange-300 bg-orange-50'
  if (priority === 3) return 'border-yellow-200 bg-yellow-50'
  return 'border-gray-200 bg-white'
}

function priorityBadgeClass(priority: number) {
  if (priority === 1) return 'bg-red-100 text-red-800'
  if (priority === 2) return 'bg-orange-100 text-orange-800'
  if (priority === 3) return 'bg-yellow-100 text-yellow-800'
  return 'bg-gray-100 text-gray-600'
}
</script>
