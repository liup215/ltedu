<template>
  <div class="p-6 bg-gray-50 min-h-screen">
    <div class="bg-white p-6 rounded-lg shadow-md">
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-2xl font-semibold text-gray-700">{{ $t('examPaperManagement.title') }}</h1>
      </div>

      <!-- Search and Filters -->
      <div class="mb-4 grid grid-cols-1 md:grid-cols-2 gap-4">
        <input
          type="text"
          v-model="searchQuery.name"
          @input="debouncedFetchExamPapers"
          :placeholder="$t('examPaperManagement.searchByName')"
          class="p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
        />
        <input
          type="number"
          v-model.number="searchQuery.year"
          @input="debouncedFetchExamPapers"
          :placeholder="$t('examPaperManagement.searchByYear')"
          class="p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
        />
      </div>

      <div v-if="isLoading" class="text-center py-4">
        <p class="text-gray-600">{{ $t('examPaperManagement.loading') }}</p>
      </div>
      <div v-else-if="errorMessage" class="text-center py-4 text-red-600 bg-red-100 p-3 rounded-md">
        <p>{{ errorMessage }}</p>
      </div>
      <div v-else-if="examPapers && examPapers.length === 0" class="text-center py-4">
        <p class="text-gray-600">{{ $t('examPaperManagement.noData') }}</p>
      </div>

      <div v-else-if="examPapers && examPapers.length > 0" class="overflow-x-auto">
        <table class="min-w-full bg-white">
          <thead class="bg-gray-100">
            <tr>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('examPaperManagement.name') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('examPaperManagement.creator') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('examPaperManagement.year') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('examPaperManagement.syllabus') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('examPaperManagement.questionCount') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('examPaperManagement.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="paper in examPapers" :key="paper.id" class="hover:bg-gray-50">
              <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-700">{{ paper.name }}</td>
              <td class="px-4 py-4 whitespace-normal text-sm text-gray-700">
                <div>{{ paper.user?.realname || paper.user?.nickname || paper.user?.username || '-' }}</div>
                <div class="text-xs text-gray-500" v-if="paper.user?.email">{{ paper.user.email }}</div>
              </td>
              <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-700">{{ paper.year }} </td>
              <td class="px-4 py-4 whitespace-normal text-sm text-gray-700">
                <div>{{ paper.syllabus?.name || '-' }}</div>
                <div class="text-gray-500 text-xs">{{ paper.syllabus?.code || '' }}</div>
              </td>
              <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-700">{{ paper.questionIds?.length || paper.questions?.length || '-' }}</td>
              <td class="px-4 py-4 whitespace-nowrap text-sm font-medium">
                <button
                  @click="goToView()"
                  class="text-blue-600 hover:text-blue-800 mr-3 transition duration-150"
                >
                  {{ $t('examPaperManagement.view') }}
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div v-if="!isLoading && examPapers && examPapers.length > 0 && pagination.totalPages > 1" class="mt-6 flex justify-between items-center">
        <div>
          <p class="text-sm text-gray-700">
            {{ $t('examPaperManagement.pageInfo', { page: pagination.pageIndex, totalPages: pagination.totalPages, total: pagination.totalItems }) }}
          </p>
        </div>
        <div class="flex space-x-2">
          <button
            @click="handlePageChange(pagination.pageIndex - 1)"
            :disabled="pagination.pageIndex <= 1"
            class="px-3 py-1 border border-gray-300 rounded-md text-sm hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {{ $t('examPaperManagement.previous') }}
          </button>
          <button
            @click="handlePageChange(pagination.pageIndex + 1)"
            :disabled="pagination.pageIndex >= pagination.totalPages"
            class="px-3 py-1 border border-gray-300 rounded-md text-sm hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {{ $t('examPaperManagement.next') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { examPaperService } from '../../services/examPaperService'
import type { ExamPaper, ExamPaperQuery } from '../../models/examPaper.model'

const examPapers = ref<ExamPaper[] | null>(null)
const isLoading = ref(false)
const errorMessage = ref<string | null>(null)

const pagination = ref({
  pageIndex: 1,
  pageSize: 10,
  totalItems: 0,
  totalPages: 0
})

const searchQuery = ref({
  name: '',
  year: undefined as number | undefined,
})

const fetchExamPapers = async () => {
  isLoading.value = true
  errorMessage.value = null
  try {
    const query: ExamPaperQuery = {
      pageIndex: pagination.value.pageIndex,
      pageSize: pagination.value.pageSize,
      name: searchQuery.value.name || undefined,
      year: searchQuery.value.year || undefined,
    }
    const response = await examPaperService.getExamPaperList(query)
    console.log('Exam Papers Response:', response)
    if (response.data) {
      examPapers.value = response.data.list || []
      pagination.value.totalItems = response.data.total
      pagination.value.totalPages = Math.ceil(response.data.total / pagination.value.pageSize)
      pagination.value.pageIndex = query.pageIndex || 1
    } else {
      examPapers.value = []
      errorMessage.value = response.message || 'Failed to load exam papers.'
    }
  } catch (error: any) {
    examPapers.value = []
    errorMessage.value = error.message || 'An unexpected error occurred.'
  } finally {
    isLoading.value = false
  }
}

let debounceTimer: number | undefined
const debouncedFetchExamPapers = () => {
  clearTimeout(debounceTimer)
  debounceTimer = window.setTimeout(() => {
    pagination.value.pageIndex = 1
    fetchExamPapers()
  }, 500)
}

onMounted(fetchExamPapers)

const goToView = () => {
  // TODO: Replace with actual route name for view
  // router.push({ name: 'AdminExamPaperView', params: { id: id.toString() } })
  window.alert('View Exam Paper (待实现)')
}

const handlePageChange = (newPage: number) => {
  if (newPage > 0 && pagination.value.totalPages && newPage <= pagination.value.totalPages) {
    pagination.value.pageIndex = newPage
    fetchExamPapers()
  }
}
</script>
