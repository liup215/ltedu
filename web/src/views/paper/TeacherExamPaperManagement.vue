<template>
  <div class="p-6 bg-gray-50 min-h-screen">
    <!-- Header Section -->
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('teacherExamPaper.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('teacherExamPaper.subtitle') }}</p>
    </header>

    <!-- Filters and Actions -->
    <div class="mb-6 flex flex-col space-y-4">
      <div class="flex flex-col lg:flex-row space-y-4 lg:space-y-0 lg:space-x-4">
        <input 
          v-model="searchQuery.name"
          @input="debouncedFetchExamPapers"
          :placeholder="$t('teacherExamPaper.searchName')"
          class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-1"
        />
        <input 
          type="number"
          v-model.number="searchQuery.year"
          @input="debouncedFetchExamPapers"
          :placeholder="$t('teacherExamPaper.year')"
          class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-32"
        />
        <router-link 
          to="/paper/exam/create"
          class="inline-flex justify-center items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          {{ $t('teacherExamPaper.newExamPaper') }}
        </router-link>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="isLoading" class="flex justify-center items-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
      <span class="ml-3 text-gray-600">{{ $t('teacherExamPaper.loading') }}</span>
    </div>

    <!-- No Data State -->
    <div v-else-if="!examPapers || examPapers.length === 0" class="text-center py-12">
      <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
      </svg>
      <h3 class="mt-2 text-sm font-medium text-gray-900">{{ $t('teacherExamPaper.noExamPapers') }}</h3>
      <p class="mt-1 text-sm text-gray-500">{{ $t('teacherExamPaper.createFirst') }}</p>
    </div>

    <!-- Papers Grid -->
    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="paper in examPapers" 
        :key="paper.id"
        class="bg-white rounded-lg shadow-md border border-gray-200 hover:shadow-lg transition-shadow duration-200"
      >
        <!-- Card Header -->
        <div class="p-4 border-b border-gray-200">
          <div class="flex justify-between items-start">
            <div class="flex-1">
              <div class="flex items-center space-x-2 mb-2">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                  ID: {{ paper.id }}
                </span>
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                  {{ paper.questionIds?.length || 0 }} {{ $t('teacherExamPaper.questions') }}
                </span>
              </div>
              <h3 class="text-lg font-medium text-gray-900 line-clamp-2">{{ paper.name }}</h3>
            </div>
            <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-gray-100 text-gray-800">
              {{ paper.year }}
            </span>
          </div>
        </div>

        <!-- Card Content -->
        <div class="p-4">
          <div class="mb-4">
            <h4 class="text-sm font-medium text-gray-700 mb-2">{{ $t('teacherExamPaper.syllabus') }}:</h4>
            <p class="text-sm text-gray-600">{{ paper.syllabus?.name || '-' }}</p>
            <p class="text-xs text-gray-500">{{ paper.syllabus?.code || '' }}</p>
          </div>

          <div class="text-xs text-gray-500">
            {{ $t('teacherExamPaper.createdAt') }}: {{ paper.createdAt ? formatDate(paper.createdAt) : '-' }}
          </div>
        </div>

        <!-- Card Footer -->
        <div class="px-4 py-3 bg-gray-50 border-t border-gray-200 flex justify-end items-center space-x-3">
          <router-link 
            :to="`/paper/exam/edit/${paper.id}`"
            class="inline-flex items-center px-3 py-1.5 border border-blue-600 text-sm font-medium rounded text-blue-600 bg-white hover:bg-blue-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            {{ $t('teacherExamPaper.edit') }}
          </router-link>
          <router-link 
            :to="`/paper/exam/preview/${paper.id}`"
            class="inline-flex items-center px-3 py-1.5 border border-transparent text-sm font-medium rounded text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          >
            {{ $t('teacherExamPaper.preview') }}
          </router-link>
          <button 
            @click="handleExportPaper(paper)"
            class="inline-flex items-center px-3 py-1.5 border border-transparent text-sm font-medium rounded text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            {{ $t('teacherExamPaper.download') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Pagination -->
    <div v-if="!isLoading && examPapers && examPapers.length > 0 && pagination.totalPages > 1" 
         class="mt-6 flex justify-between items-center">
      <div class="text-sm text-gray-700">
        {{ $t('teacherExamPaper.page') }} {{ pagination.pageIndex }} / {{ pagination.totalPages }} ({{ $t('teacherExamPaper.total') }} {{ pagination.totalItems }})
      </div>
      <div class="flex space-x-2">
        <button
          @click="handlePageChange(pagination.pageIndex - 1)"
          :disabled="pagination.pageIndex <= 1"
          class="px-3 py-1 border border-gray-300 rounded-md text-sm hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ $t('teacherExamPaper.previous') }}
        </button>
        <button
          @click="handlePageChange(pagination.pageIndex + 1)"
          :disabled="pagination.pageIndex >= pagination.totalPages"
          class="px-3 py-1 border border-gray-300 rounded-md text-sm hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ $t('teacherExamPaper.next') }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { examPaperService } from '../../services/examPaperService'
import type { ExamPaper, ExamPaperQuery } from '../../models/examPaper.model'
import { exportExamPaperToDocx } from '../../utils/exportDocx'

const examPapers = ref<ExamPaper[] | null>(null)
const isLoading = ref(false)
const errorMessage = ref<string | null>(null)

const pagination = ref({
  pageIndex: 1,
  pageSize: 12, // 每页显示12份试卷
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
    if (response.data) {
      examPapers.value = response.data.list || []
      pagination.value.totalItems = response.data.total
      pagination.value.totalPages = Math.ceil(response.data.total / pagination.value.pageSize)
      pagination.value.pageIndex = query.pageIndex || 1
    } else {
      examPapers.value = []
      errorMessage.value = response.message || '加载失败'
    }
  } catch (error: any) {
    examPapers.value = []
    errorMessage.value = error.message || '加载失败'
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

const handlePageChange = (newPage: number) => {
  if (newPage > 0 && pagination.value.totalPages && newPage <= pagination.value.totalPages) {
    pagination.value.pageIndex = newPage
    fetchExamPapers()
  }
}

const handleExportPaper = async (paper: ExamPaper) => {
  // Try to get questions from paper object (if not present, fetch them)
  if ((paper as any).questions && Array.isArray((paper as any).questions)) {
    await exportExamPaperToDocx(paper);
  } else if (paper.questionIds && paper.questionIds.length > 0) {
    // Fallback: fetch questions by IDs
    // You may want to optimize this with batch API if available
    const questionService = await import('../../services/questionService');
    const responses = await Promise.all(
      paper.questionIds.map((qId: number) => questionService.default.getQuestionById(qId))
    );
    const questions = responses.filter(res => res.data).map(res => res.data);
    const exportPaper = { ...paper, questions };
    await exportExamPaperToDocx(exportPaper);
  }
}

const formatDate = (dateString: string): string => {
  return new Date(dateString).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

onMounted(fetchExamPapers)
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
