<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('feedback.adminTitle') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('feedback.adminSubtitle') }}</p>
    </header>

    <!-- Stats cards -->
    <div v-if="stats" class="grid grid-cols-2 sm:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4 text-center">
        <p class="text-3xl font-bold text-indigo-600">{{ stats.total }}</p>
        <p class="text-sm text-gray-500 mt-1">{{ $t('feedback.statsTotal') }}</p>
      </div>
      <div class="bg-white rounded-lg shadow p-4 text-center">
        <p class="text-3xl font-bold text-yellow-500">{{ stats.avgRating > 0 ? stats.avgRating.toFixed(1) : '—' }}</p>
        <p class="text-sm text-gray-500 mt-1">{{ $t('feedback.statsAvgRating') }}</p>
      </div>
      <div class="bg-white rounded-lg shadow p-4 text-center">
        <p class="text-3xl font-bold text-green-600">{{ stats.bySentiment?.positive || 0 }}</p>
        <p class="text-sm text-gray-500 mt-1">{{ $t('feedback.statsPositive') }}</p>
      </div>
      <div class="bg-white rounded-lg shadow p-4 text-center">
        <p class="text-3xl font-bold text-red-500">{{ stats.bySentiment?.negative || 0 }}</p>
        <p class="text-sm text-gray-500 mt-1">{{ $t('feedback.statsNegative') }}</p>
      </div>
    </div>

    <!-- Filters -->
    <div class="mb-4 flex flex-col sm:flex-row gap-3">
      <select
        v-model="filter.status"
        @change="loadFeedback"
        class="px-3 py-2 border border-gray-300 rounded-md shadow-sm text-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
      >
        <option value="">{{ $t('feedback.filterAllStatuses') }}</option>
        <option value="new">{{ $t('feedback.statusNew') }}</option>
        <option value="reviewed">{{ $t('feedback.statusReviewed') }}</option>
        <option value="resolved">{{ $t('feedback.statusResolved') }}</option>
      </select>
      <select
        v-model="filter.type"
        @change="loadFeedback"
        class="px-3 py-2 border border-gray-300 rounded-md shadow-sm text-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
      >
        <option value="">{{ $t('feedback.filterAllTypes') }}</option>
        <option value="general">{{ $t('feedback.typeGeneral') }}</option>
        <option value="bug">{{ $t('feedback.typeBug') }}</option>
        <option value="feature">{{ $t('feedback.typeFeature') }}</option>
        <option value="praise">{{ $t('feedback.typePraise') }}</option>
      </select>
      <button
        @click="loadFeedback"
        class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md shadow-sm text-gray-700 bg-white hover:bg-gray-50"
      >{{ $t('feedback.refresh') }}</button>
    </div>

    <!-- Table -->
    <div class="bg-white shadow overflow-x-auto sm:rounded-lg">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('feedback.colUser') }}</th>
            <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('feedback.colType') }}</th>
            <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('feedback.colContent') }}</th>
            <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('feedback.colRating') }}</th>
            <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('feedback.colSentiment') }}</th>
            <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('feedback.colStatus') }}</th>
            <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('feedback.colDate') }}</th>
            <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('feedback.colActions') }}</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-if="loading">
            <td colspan="9" class="px-4 py-6 text-center text-sm text-gray-500">{{ $t('feedback.loading') }}</td>
          </tr>
          <tr v-else-if="items.length === 0">
            <td colspan="9" class="px-4 py-6 text-center text-sm text-gray-500">{{ $t('feedback.noData') }}</td>
          </tr>
          <tr v-for="item in items" :key="item.id" class="hover:bg-gray-50">
            <td class="px-4 py-3 text-sm text-gray-900">{{ item.id }}</td>
            <td class="px-4 py-3 text-sm text-gray-500">{{ item.userId }}</td>
            <td class="px-4 py-3">
              <span :class="typeBadgeClass(item.type)" class="px-2 py-0.5 rounded-full text-xs font-medium">
                {{ $t('feedback.type_' + item.type) }}
              </span>
            </td>
            <td class="px-4 py-3 text-sm text-gray-700 max-w-xs truncate" :title="item.content">{{ item.content }}</td>
            <td class="px-4 py-3 text-sm text-yellow-500">{{ item.rating > 0 ? '★'.repeat(item.rating) : '—' }}</td>
            <td class="px-4 py-3">
              <span :class="sentimentBadgeClass(item.sentiment)" class="px-2 py-0.5 rounded-full text-xs font-medium">
                {{ $t('feedback.sentiment_' + item.sentiment) }}
              </span>
            </td>
            <td class="px-4 py-3">
              <span :class="statusBadgeClass(item.status)" class="px-2 py-0.5 rounded-full text-xs font-medium">
                {{ $t('feedback.status_' + item.status) }}
              </span>
            </td>
            <td class="px-4 py-3 text-xs text-gray-500">{{ formatDate(item.createdAt) }}</td>
            <td class="px-4 py-3 text-sm space-x-2">
              <button @click="openDetail(item)" class="text-indigo-600 hover:text-indigo-900 text-xs">{{ $t('feedback.actionView') }}</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div class="mt-4 flex items-center justify-between">
      <p class="text-sm text-gray-500">{{ $t('feedback.paginationTotal', { total }) }}</p>
      <div class="flex gap-2">
        <button
          :disabled="page <= 1"
          @click="changePage(page - 1)"
          class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-50"
        >{{ $t('feedback.prev') }}</button>
        <button
          :disabled="page * pageSize >= total"
          @click="changePage(page + 1)"
          class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-50"
        >{{ $t('feedback.next') }}</button>
      </div>
    </div>

    <!-- Detail modal -->
    <div v-if="detailItem" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50 p-4">
      <div class="bg-white rounded-xl shadow-xl w-full max-w-lg overflow-hidden">
        <div class="flex items-center justify-between bg-indigo-600 px-4 py-3">
          <h2 class="text-white font-semibold text-sm">{{ $t('feedback.detailTitle') }} #{{ detailItem.id }}</h2>
          <button @click="detailItem = null" class="text-indigo-200 hover:text-white">✕</button>
        </div>
        <div class="p-4 space-y-3 text-sm">
          <div><span class="font-medium text-gray-700">{{ $t('feedback.colUser') }}:</span> {{ detailItem.userId }}</div>
          <div><span class="font-medium text-gray-700">{{ $t('feedback.colType') }}:</span> {{ $t('feedback.type_' + detailItem.type) }}</div>
          <div><span class="font-medium text-gray-700">{{ $t('feedback.colRating') }}:</span> {{ detailItem.rating > 0 ? '★'.repeat(detailItem.rating) : '—' }}</div>
          <div><span class="font-medium text-gray-700">{{ $t('feedback.colSentiment') }}:</span> {{ $t('feedback.sentiment_' + detailItem.sentiment) }}</div>
          <div><span class="font-medium text-gray-700">{{ $t('feedback.pageContext') }}:</span> {{ detailItem.pageContext || '—' }}</div>
          <div>
            <span class="font-medium text-gray-700">{{ $t('feedback.colContent') }}:</span>
            <p class="mt-1 text-gray-600 whitespace-pre-wrap bg-gray-50 p-2 rounded">{{ detailItem.content }}</p>
          </div>
          <!-- Status update -->
          <div>
            <label class="block font-medium text-gray-700 mb-1">{{ $t('feedback.colStatus') }}</label>
            <select v-model="statusForm.status" class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm">
              <option value="new">{{ $t('feedback.statusNew') }}</option>
              <option value="reviewed">{{ $t('feedback.statusReviewed') }}</option>
              <option value="resolved">{{ $t('feedback.statusResolved') }}</option>
            </select>
          </div>
          <div>
            <label class="block font-medium text-gray-700 mb-1">{{ $t('feedback.adminNote') }}</label>
            <textarea v-model="statusForm.adminNote" rows="2" class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm" />
          </div>
          <p v-if="detailError" class="text-red-600 text-xs">{{ detailError }}</p>
          <button
            @click="saveStatus"
            :disabled="savingStatus"
            class="w-full py-2 bg-indigo-600 text-white text-sm rounded-md hover:bg-indigo-700 disabled:opacity-50"
          >{{ $t('feedback.saveStatus') }}</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import feedbackService from '../../services/feedbackService'
import type { UserFeedback, FeedbackStats } from '../../models/feedback.model'

const { t } = useI18n()

const loading = ref(false)
const items = ref<UserFeedback[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = 20
const stats = ref<FeedbackStats | null>(null)

const filter = reactive({ status: '', type: '' })

const detailItem = ref<UserFeedback | null>(null)
const statusForm = reactive({ status: '', adminNote: '' })
const savingStatus = ref(false)
const detailError = ref('')

onMounted(() => {
  loadFeedback()
  loadStats()
})

async function loadFeedback() {
  loading.value = true
  try {
    const res = await feedbackService.list({
      pageSize,
      pageIndex: page.value,
      status: filter.status || undefined,
      type: filter.type || undefined
    })
    if (res.code === 0) {
      items.value = res.data.list
      total.value = res.data.total
    }
  } finally {
    loading.value = false
  }
}

async function loadStats() {
  const res = await feedbackService.getStats()
  if (res.code === 0) stats.value = res.data
}

function changePage(p: number) {
  page.value = p
  loadFeedback()
}

function openDetail(item: UserFeedback) {
  detailItem.value = item
  statusForm.status = item.status
  statusForm.adminNote = item.adminNote || ''
  detailError.value = ''
}

async function saveStatus() {
  if (!detailItem.value) return
  savingStatus.value = true
  detailError.value = ''
  try {
    const res = await feedbackService.updateStatus(detailItem.value.id, statusForm.status, statusForm.adminNote)
    if (res.code === 0) {
      detailItem.value = null
      loadFeedback()
      loadStats()
    } else {
      detailError.value = res.message
    }
  } catch {
    detailError.value = t('feedback.saveError')
  } finally {
    savingStatus.value = false
  }
}

function formatDate(dateStr: string) {
  if (!dateStr) return '—'
  return new Date(dateStr).toLocaleString()
}

function typeBadgeClass(type: string) {
  const map: Record<string, string> = {
    bug: 'bg-red-100 text-red-800',
    feature: 'bg-blue-100 text-blue-800',
    praise: 'bg-green-100 text-green-800',
    general: 'bg-gray-100 text-gray-800'
  }
  return map[type] || 'bg-gray-100 text-gray-800'
}

function sentimentBadgeClass(sentiment: string) {
  const map: Record<string, string> = {
    positive: 'bg-green-100 text-green-800',
    negative: 'bg-red-100 text-red-800',
    neutral: 'bg-yellow-100 text-yellow-800'
  }
  return map[sentiment] || 'bg-gray-100 text-gray-800'
}

function statusBadgeClass(status: string) {
  const map: Record<string, string> = {
    new: 'bg-blue-100 text-blue-800',
    reviewed: 'bg-yellow-100 text-yellow-800',
    resolved: 'bg-green-100 text-green-800'
  }
  return map[status] || 'bg-gray-100 text-gray-800'
}
</script>
