<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ t('analytics.title') }}</h1>
      <p class="mt-2 text-sm text-gray-600">{{ t('analytics.subtitle') }}</p>
    </header>

    <!-- Class Selector -->
    <div class="mb-6 flex items-center gap-4">
      <label class="text-sm font-medium text-gray-700">{{ t('class.title', 'Class') }}</label>
      <select
        v-model="selectedClassId"
        @change="loadClassData"
        class="border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 w-64"
      >
        <option value="">{{ t('analytics.selectClass') }}</option>
        <option v-for="cls in classes" :key="cls.id" :value="cls.id">{{ cls.name }}</option>
      </select>
      <button
        v-if="selectedClassId"
        @click="loadClassData"
        class="inline-flex items-center px-3 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
      >
        ↻ Refresh
      </button>
    </div>

    <div v-if="!selectedClassId" class="text-center py-16 text-gray-400">
      <svg class="mx-auto h-16 w-16 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
      </svg>
      <p class="text-lg">{{ t('analytics.selectClass') }}</p>
    </div>

    <div v-else>
      <!-- Loading indicator -->
      <div v-if="loading" class="text-center py-8 text-gray-500">{{ t('analytics.loading') }}</div>

      <div v-else>
        <!-- Class Summary Cards -->
        <div v-if="summary" class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-8">
          <div class="bg-white rounded-lg shadow p-4 text-center">
            <div class="text-3xl font-bold text-indigo-600">{{ summary.totalStudents }}</div>
            <div class="text-sm text-gray-500 mt-1">{{ t('analytics.totalStudents') }}</div>
          </div>
          <div class="bg-white rounded-lg shadow p-4 text-center">
            <div class="text-3xl font-bold text-green-600">{{ summary.activeStudents }}</div>
            <div class="text-sm text-gray-500 mt-1">{{ t('analytics.activeStudents') }}</div>
          </div>
          <div class="bg-white rounded-lg shadow p-4 text-center">
            <div class="text-3xl font-bold" :class="masteryColor(summary.avgMastery)">{{ summary.avgMastery.toFixed(1) }}%</div>
            <div class="text-sm text-gray-500 mt-1">{{ t('analytics.avgMastery') }}</div>
          </div>
          <div class="bg-white rounded-lg shadow p-4 text-center">
            <div class="text-3xl font-bold" :class="masteryColor(summary.avgAccuracy)">{{ summary.avgAccuracy.toFixed(1) }}%</div>
            <div class="text-sm text-gray-500 mt-1">{{ t('analytics.avgAccuracy') }}</div>
          </div>
          <div class="bg-white rounded-lg shadow p-4 text-center">
            <div class="text-3xl font-bold text-blue-600">{{ summary.avgCoverage.toFixed(1) }}%</div>
            <div class="text-sm text-gray-500 mt-1">{{ t('analytics.avgCoverage') }}</div>
          </div>
          <div class="bg-white rounded-lg shadow p-4 text-center">
            <div class="text-3xl font-bold text-gray-700">{{ summary.totalAttempts.toLocaleString() }}</div>
            <div class="text-sm text-gray-500 mt-1">{{ t('analytics.totalAttempts') }}</div>
          </div>
          <div class="bg-white rounded-lg shadow p-4 text-center">
            <div class="text-3xl font-bold text-purple-600">{{ summary.weeklyAttempts.toLocaleString() }}</div>
            <div class="text-sm text-gray-500 mt-1">{{ t('analytics.weeklyAttempts') }}</div>
          </div>
          <div class="bg-white rounded-lg shadow p-4 text-center">
            <div class="text-3xl font-bold text-red-600">{{ summary.atRiskCount }}</div>
            <div class="text-sm text-gray-500 mt-1">{{ t('analytics.atRiskStudents') }}</div>
          </div>
        </div>

        <!-- Tabs -->
        <div class="mb-4 border-b border-gray-200">
          <nav class="-mb-px flex space-x-6 overflow-x-auto">
            <button
              v-for="tab in tabs"
              :key="tab.key"
              @click="activeTab = tab.key"
              :class="['pb-2 text-sm font-medium border-b-2 transition-colors whitespace-nowrap', activeTab === tab.key ? 'border-indigo-500 text-indigo-600' : 'border-transparent text-gray-500 hover:text-gray-700']"
            >
              {{ tab.label }}
            </button>
          </nav>
        </div>

        <!-- Student Performance Tab -->
        <div v-if="activeTab === 'students'">
          <div v-if="!students.length" class="text-center py-8 text-gray-400">{{ t('analytics.noData') }}</div>
          <div v-else class="bg-white shadow rounded-lg overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">{{ t('analytics.student') }}</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">{{ t('analytics.mastery') }}</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">{{ t('analytics.coverage') }}</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">{{ t('analytics.accuracy') }}</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">{{ t('analytics.totalAttempts') }}</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">{{ t('analytics.lastActive') }}</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">{{ t('analytics.atRisk') }}</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="s in students" :key="s.userId">
                  <td class="px-4 py-3 text-sm font-medium text-gray-900">{{ s.userName }}</td>
                  <td class="px-4 py-3 text-sm">
                    <div class="flex items-center gap-2">
                      <div class="w-24 bg-gray-200 rounded-full h-2">
                        <div class="h-2 rounded-full" :class="barColor(s.masteryLevel)" :style="{ width: s.masteryLevel + '%' }"></div>
                      </div>
                      <span :class="masteryColor(s.masteryLevel)">{{ s.masteryLevel.toFixed(1) }}%</span>
                    </div>
                  </td>
                  <td class="px-4 py-3 text-sm text-gray-700">{{ s.coverageLevel.toFixed(1) }}%</td>
                  <td class="px-4 py-3 text-sm" :class="masteryColor(s.accuracyRate)">{{ s.accuracyRate.toFixed(1) }}%</td>
                  <td class="px-4 py-3 text-sm text-gray-700">{{ s.totalAttempts }}</td>
                  <td class="px-4 py-3 text-sm text-gray-500">{{ s.lastActiveAt ? new Date(s.lastActiveAt).toLocaleDateString() : '—' }}</td>
                  <td class="px-4 py-3 text-sm">
                    <span v-if="s.isAtRisk" class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-red-100 text-red-800">
                      ⚠ {{ t('analytics.atRisk') }}
                    </span>
                    <span v-else class="text-green-600 text-xs">✓</span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- Heatmap Tab -->
        <div v-if="activeTab === 'heatmap'">
          <div v-if="!heatmap || !heatmap.chapters.length" class="text-center py-8 text-gray-400">{{ t('analytics.noData') }}</div>
          <div v-else class="bg-white shadow rounded-lg overflow-x-auto">
            <table class="min-w-full text-xs">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-3 py-2 text-left font-medium text-gray-500 min-w-[160px]">{{ t('analytics.chapter') }}</th>
                  <th class="px-3 py-2 text-center font-medium text-gray-500">Avg</th>
                  <th v-for="student in heatmap.students" :key="student.userId" class="px-2 py-2 text-center font-medium text-gray-500 min-w-[60px] truncate max-w-[80px]" :title="student.userName">
                    {{ student.userName.slice(0, 6) }}
                  </th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-100">
                <tr v-for="row in heatmap.chapters" :key="row.chapterId">
                  <td class="px-3 py-2 font-medium text-gray-800 truncate max-w-[160px]" :title="row.chapterName">{{ row.chapterName }}</td>
                  <td class="px-3 py-2 text-center">
                    <span class="px-1.5 py-0.5 rounded text-xs font-medium" :class="heatmapCellClass(row.avgMastery)">{{ row.avgMastery.toFixed(0) }}%</span>
                  </td>
                  <td v-for="score in row.studentData" :key="score.userId" class="px-2 py-2 text-center">
                    <span v-if="score.isCovered" class="px-1.5 py-0.5 rounded text-xs font-medium" :class="heatmapCellClass(score.masteryLevel)">{{ score.masteryLevel.toFixed(0) }}</span>
                    <span v-else class="text-gray-300 text-xs">—</span>
                  </td>
                </tr>
              </tbody>
            </table>
            <div class="p-3 border-t border-gray-100 flex items-center gap-4 text-xs text-gray-500">
              <span class="flex items-center gap-1"><span class="inline-block w-3 h-3 rounded bg-red-200"></span> &lt;30%</span>
              <span class="flex items-center gap-1"><span class="inline-block w-3 h-3 rounded bg-yellow-200"></span> 30–69%</span>
              <span class="flex items-center gap-1"><span class="inline-block w-3 h-3 rounded bg-green-200"></span> ≥70%</span>
            </div>
          </div>
        </div>

        <!-- Trends Tab -->
        <div v-if="activeTab === 'trends'">
          <div v-if="!trends.length" class="text-center py-8 text-gray-400">{{ t('analytics.noData') }}</div>
          <div v-else class="bg-white shadow rounded-lg p-4">
            <div class="overflow-x-auto">
              <div class="flex items-end gap-1 h-40 min-w-max">
                <template v-for="point in trends" :key="point.date">
                  <div class="flex flex-col items-center gap-0.5 w-8">
                    <div
                      class="w-4 bg-indigo-400 rounded-t"
                      :style="{ height: Math.max(2, (point.totalAttempts / maxTrendValue) * 120) + 'px' }"
                      :title="`${point.date}: ${point.totalAttempts} attempts`"
                    ></div>
                    <div
                      class="w-4 bg-green-400 rounded-t"
                      :style="{ height: Math.max(2, (point.correctAttempts / maxTrendValue) * 120) + 'px' }"
                      :title="`${point.date}: ${point.correctAttempts} correct`"
                    ></div>
                    <span class="text-gray-400 text-xs transform -rotate-45 origin-top-left mt-1 hidden sm:block">{{ point.date.slice(5) }}</span>
                  </div>
                </template>
              </div>
            </div>
            <div class="mt-3 flex items-center gap-4 text-xs text-gray-500">
              <span class="flex items-center gap-1"><span class="inline-block w-3 h-3 rounded bg-indigo-400"></span> {{ t('analytics.total_attempts') }}</span>
              <span class="flex items-center gap-1"><span class="inline-block w-3 h-3 rounded bg-green-400"></span> {{ t('analytics.correct_attempts') }}</span>
            </div>
          </div>
        </div>

        <!-- Early Warnings Tab -->
        <div v-if="activeTab === 'warnings'">
          <div v-if="!warnings.length" class="text-center py-8 text-green-600">
            <svg class="mx-auto h-12 w-12 mb-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            {{ t('analytics.noWarnings') }}
          </div>
          <div v-else class="space-y-3">
            <div v-for="w in warnings" :key="w.userId" :class="['bg-white shadow rounded-lg p-4 border-l-4', severityBorderClass(w.severity)]">
              <div class="flex items-start justify-between">
                <div>
                  <div class="font-medium text-gray-900 flex items-center gap-2">
                    {{ w.userName }}
                    <span :class="['text-xs px-2 py-0.5 rounded-full font-medium', severityBadgeClass(w.severity)]">{{ w.severity }}</span>
                  </div>
                  <ul class="mt-2 space-y-1">
                    <li v-for="(reason, i) in w.reasons" :key="i" class="text-sm text-gray-600 flex items-center gap-1">
                      <span class="text-red-400">•</span> {{ reason }}
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import analyticsService, {
  type ClassPerformanceSummary,
  type StudentPerformanceSummary,
  type ClassHeatmap,
  type AttemptTrendPoint,
  type EarlyWarningStudent
} from '../../services/analyticsService'
import classService from '../../services/classService'

const { t } = useI18n()

const classes = ref<Array<{ id: number; name: string }>>([])
const selectedClassId = ref<number | ''>('')
const loading = ref(false)
const activeTab = ref('students')

const summary = ref<ClassPerformanceSummary | null>(null)
const students = ref<StudentPerformanceSummary[]>([])
const heatmap = ref<ClassHeatmap | null>(null)
const trends = ref<AttemptTrendPoint[]>([])
const warnings = ref<EarlyWarningStudent[]>([])

const tabs = computed(() => [
  { key: 'students', label: t('analytics.studentPerformance') },
  { key: 'heatmap', label: t('analytics.heatmap') },
  { key: 'trends', label: t('analytics.trends') },
  { key: 'warnings', label: t('analytics.earlyWarnings') + (warnings.value.length ? ` (${warnings.value.length})` : '') }
])

const maxTrendValue = computed(() => Math.max(1, ...trends.value.map(p => p.totalAttempts)))

onMounted(async () => {
  try {
    const res = await classService.list({ pageSize: 100, pageIndex: 1 })
    if (res.code === 0 && res.data) {
      classes.value = (res.data.list || []).map((c: any) => ({ id: c.id, name: c.name }))
    }
  } catch {
    // ignore
  }
})

async function loadClassData() {
  if (!selectedClassId.value) return
  loading.value = true
  const id = Number(selectedClassId.value)
  try {
    const [sumRes, studentsRes, heatmapRes, trendsRes, warningsRes] = await Promise.all([
      analyticsService.getClassSummary(id),
      analyticsService.getStudentPerformanceList(id),
      analyticsService.getClassHeatmap(id),
      analyticsService.getAttemptTrends({ classId: id }),
      analyticsService.getEarlyWarnings(id)
    ])
    if (sumRes.code === 0) summary.value = sumRes.data
    if (studentsRes.code === 0) students.value = studentsRes.data?.list || []
    if (heatmapRes.code === 0) heatmap.value = heatmapRes.data
    if (trendsRes.code === 0) trends.value = trendsRes.data || []
    if (warningsRes.code === 0) warnings.value = warningsRes.data?.list || []
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

function barColor(value: number) {
  if (value >= 70) return 'bg-green-500'
  if (value >= 40) return 'bg-yellow-400'
  return 'bg-red-400'
}

function heatmapCellClass(value: number) {
  if (value >= 70) return 'bg-green-100 text-green-800'
  if (value >= 30) return 'bg-yellow-100 text-yellow-800'
  return 'bg-red-100 text-red-800'
}

function severityBorderClass(severity: string) {
  if (severity === 'high') return 'border-red-500'
  if (severity === 'medium') return 'border-yellow-500'
  return 'border-blue-400'
}

function severityBadgeClass(severity: string) {
  if (severity === 'high') return 'bg-red-100 text-red-800'
  if (severity === 'medium') return 'bg-yellow-100 text-yellow-800'
  return 'bg-blue-100 text-blue-800'
}
</script>
