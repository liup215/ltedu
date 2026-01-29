<template>
  <div class="p-6 bg-gray-50 min-h-screen">
    <div class="bg-white p-6 rounded-lg shadow-md">
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-2xl font-semibold text-gray-700">{{ $t('pastPaper.title') }}</h1>
        <button
          @click="goToCreate"
          class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition duration-150"
        >
          {{ $t('pastPaper.add') }}
        </button>
      </div>

      <!-- Search and Filters -->
      <div class="mb-4 grid grid-cols-1 md:grid-cols-2 gap-4">
        <input
          type="text"
          v-model="searchQuery.name"
          @input="debouncedFetchPastPapers"
          :placeholder="$t('pastPaper.searchByName')"
          class="p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
        />
        <input
          type="number"
          v-model.number="searchQuery.year"
          @input="debouncedFetchPastPapers"
          :placeholder="$t('pastPaper.selectYear')"
          class="p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
        />
      </div>

      <div v-if="isLoading" class="text-center py-4">
        <p class="text-gray-600">{{ $t('pastPaper.loading') }}</p>
      </div>
      <div v-else-if="errorMessage" class="text-center py-4 text-red-600 bg-red-100 p-3 rounded-md">
        <p>{{ errorMessage }}</p>
      </div>
      <div v-else-if="pastPapers.length === 0" class="text-center py-4">
        <p class="text-gray-600">{{ $t('pastPaper.noData') }}</p>
      </div>

      <div v-else class="overflow-x-auto">
        <table class="min-w-full bg-white">
          <thead class="bg-gray-100">
            <tr>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('pastPaper.name') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('pastPaper.year') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('pastPaper.organisation') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('pastPaper.qualification') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('pastPaper.syllabus') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('pastPaper.series') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('pastPaper.code') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('examPaperManagement.questionCount') }}</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('pastPaper.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="paper in pastPapers" :key="paper.id" class="hover:bg-gray-50">
              <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-700">{{ paper.name }}</td>
              <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-700">{{ paper.year }}</td>
              <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-700">
                {{ paper.syllabus?.qualification?.organisation?.name || '-' }}
              </td>
              <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-700">
                {{ paper.syllabus?.qualification?.name || '-' }}
              </td>
              <td class="px-4 py-4 whitespace-normal text-sm text-gray-700">
                <div>{{ paper.syllabus?.name || '-' }}</div>
                <div class="text-gray-500 text-xs">{{ paper.syllabus?.code || '' }}</div>
              </td>
              <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-700">
                {{ paper.paperSeries?.name || '-' }}
              </td>
              <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-700">
                {{ paper.paperCode?.name || '-' }}
              </td>
              <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-700">{{ paper.questionNumber }}</td>
              <td class="px-4 py-4 whitespace-nowrap text-sm font-medium">
                <button
                  @click="goToEdit(paper.id)"
                  class="text-blue-600 hover:text-blue-800 mr-3 transition duration-150"
                >
                  {{ $t('pastPaper.edit') }}
                </button>
                <button
                  @click="confirmDelete(paper.id)"
                  class="text-red-600 hover:text-red-800 transition duration-150"
                >
                  {{ $t('pastPaper.delete') }}
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div v-if="!isLoading && pastPapers.length > 0 && pagination.totalPages > 1" class="mt-6 flex justify-between items-center">
        <div>
          <p class="text-sm text-gray-700">
            {{ $t('pastPaper.pageInfo', { from: (pagination.pageIndex - 1) * pagination.pageSize + 1, to: Math.min(pagination.pageIndex * pagination.pageSize, pagination.totalItems), total: pagination.totalItems }) }}
          </p>
        </div>
        <div class="flex space-x-2">
          <button
            @click="handlePageChange(pagination.pageIndex - 1)"
            :disabled="pagination.pageIndex <= 1"
            class="px-3 py-1 border border-gray-300 rounded-md text-sm hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {{ $t('syllabusManagement.previous') }}
          </button>
          <button
            @click="handlePageChange(pagination.pageIndex + 1)"
            :disabled="pagination.pageIndex >= pagination.totalPages"
            class="px-3 py-1 border border-gray-300 rounded-md text-sm hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {{ $t('syllabusManagement.next') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import pastPaperService from '../../services/pastPaperService'
import type { PastPaper, PastPaperQuery } from '../../models/pastPaper.model'

const router = useRouter()

const pastPapers = ref<PastPaper[]>([])
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

const fetchPastPapers = async () => {
  isLoading.value = true
  errorMessage.value = null
  try {
    const query: PastPaperQuery = {
      pageIndex: pagination.value.pageIndex,
      pageSize: pagination.value.pageSize,
      name: searchQuery.value.name || undefined,
      year: searchQuery.value.year || undefined,
    }
    const response = await pastPaperService.getPastPapers(query)
    if (response.code === 0 && response.data) {
      pastPapers.value = response.data.list
      pagination.value.totalItems = response.data.total
      pagination.value.totalPages = Math.ceil(response.data.total / pagination.value.pageSize)
      pagination.value.pageIndex = query.pageIndex || 1
    } else {
      errorMessage.value = response.message || 'Failed to load past papers.'
    }
  } catch (error: any) {
    errorMessage.value = error.message || 'An unexpected error occurred.'
  } finally {
    isLoading.value = false
  }
}

let debounceTimer: number | undefined
const debouncedFetchPastPapers = () => {
  clearTimeout(debounceTimer)
  debounceTimer = window.setTimeout(() => {
    pagination.value.pageIndex = 1
    fetchPastPapers()
  }, 500)
}

onMounted(fetchPastPapers)

const goToCreate = () => {
  router.push({ name: 'AdminPastPaperCreate' })
}

const goToEdit = (id: number) => {
  router.push({ name: 'AdminPastPaperEdit', params: { id: id.toString() } })
}

const confirmDelete = async (id: number) => {
  if (window.confirm('Are you sure you want to delete this past paper?')) {
    isLoading.value = true
    try {
      const response = await pastPaperService.deletePastPaper(id.toString())
      if (response.code === 0) {
        fetchPastPapers()
      } else {
        errorMessage.value = response.message || 'Failed to delete past paper.'
      }
    } catch (error: any) {
      errorMessage.value = error.message || 'An unexpected error occurred during deletion.'
    } finally {
      isLoading.value = false
    }
  }
}

const handlePageChange = (newPage: number) => {
  if (newPage > 0 && newPage <= pagination.value.totalPages) {
    pagination.value.pageIndex = newPage
    fetchPastPapers()
  }
}
</script>
