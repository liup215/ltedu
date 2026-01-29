<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('paperSeries.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('paperSeries.subtitle') || '' }}</p>
    </header>

    <div class="mb-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0 sm:space-x-4">
      <div class="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 w-full sm:w-auto">
        <select
          v-model.number="selectedOrganisationId"
          class="px-3 py-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-base w-full sm:w-auto"
          :disabled="dropdownLoading.organisations"
        >
          <option :value="0">{{ $t('syllabusManagement.allOrganisations') }}</option>
          <option v-for="org in organisations" :key="org.id" :value="org.id">{{ org.name }}</option>
        </select>
        <select
          v-model.number="selectedQualificationId"
          class="px-3 py-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-base w-full sm:w-auto"
          :disabled="dropdownLoading.qualifications || !selectedOrganisationId"
        >
          <option :value="0">{{ $t('syllabusManagement.allQualifications') }}</option>
          <option v-for="q in qualifications" :key="q.id" :value="q.id">{{ q.name }}</option>
        </select>
        <select 
          v-model.number="selectedSyllabusId"
          class="px-3 py-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-base w-full sm:w-auto"
          :disabled="dropdownLoading.syllabi || !selectedQualificationId"
        >
          <option :value="0">{{ $t('syllabusManagement.addSyllabus') }}</option>
          <option v-for="syllabus in syllabi" :key="syllabus.id" :value="syllabus.id">
            {{ syllabus.name }}
          </option>
        </select>
        <input 
          type="text" 
          v-model="searchQuery"
          :placeholder="$t('paperSeries.searchByName')"
          class="px-3 py-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-base w-full sm:w-auto"
        />
      </div>
      <button 
        @click="goToCreatePage"
        class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 w-full sm:w-auto"
      >
        {{ $t('paperSeries.add') }}
      </button>
    </div>

    <div class="bg-white shadow overflow-x-auto sm:rounded-lg">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('paperSeries.name') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('syllabusManagement.name') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('qualificationManagement.name') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('organisationManagement.name') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider min-w-[180px]">{{ $t('paperSeries.actions') }}</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-if="loading">
            <td colspan="6" class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center">{{ $t('paperSeries.loading') }}</td>
          </tr>
          <tr v-else-if="!paperSeries.length">
            <td colspan="6" class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center">{{ $t('paperSeries.noData') }}</td>
          </tr>
          <tr v-for="series in paperSeries" :key="series.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ series.id }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ series.name }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ series.syllabus?.name || '-' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ series.syllabus?.qualification?.name || '-' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ series.syllabus?.qualification?.organisation?.name || '-' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-3">
              <router-link
                :to="`/admin/paper-series/${series.id}/edit`"
                class="text-indigo-600 hover:text-indigo-900"
              >
                {{ $t('paperSeries.edit') }}
              </router-link>
              <button 
                @click="deletePaperSeries(series.id!)"
                class="text-red-600 hover:text-red-900"
              >
                {{ $t('paperSeries.delete') }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    
    <div v-if="!loading && totalItems > 0" class="mt-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0">
      <p class="text-sm text-gray-700">
        {{ $t('paperSeries.pageInfo', { from: (currentPage - 1) * pageSize + 1, to: Math.min(currentPage * pageSize, totalItems), total: totalItems }) }}
      </p>
      <nav v-if="totalPages > 1" class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
        <button
          @click="goToPage(currentPage - 1)"
          :disabled="currentPage === 1"
          class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ $t('syllabusManagement.previous') }}
        </button>
        <button
          v-for="pageNumber in paginationRange"
          :key="pageNumber"
          @click="goToPage(pageNumber)"
          :class="[
            'relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium',
            currentPage === pageNumber ? 'z-10 bg-indigo-50 border-indigo-500 text-indigo-600' : 'bg-white text-gray-700 hover:bg-gray-50'
          ]"
        >
          {{ pageNumber }}
        </button>
        <button
          @click="goToPage(currentPage + 1)"
          :disabled="currentPage === totalPages || totalPages === 0"
          class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ $t('syllabusManagement.next') }}
        </button>
      </nav>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed, watch } from 'vue'
import { useRouter } from 'vue-router'
import paperService from '../../services/paperSeriesService'
import syllabusService from '../../services/syllabusService'
import organisationService from '../../services/organisationService'
import qualificationService from '../../services/qualificationService'
import type { PaperSeries, PaperSeriesQuery } from '../../models/paperSeries.model'
import type { Syllabus, SyllabusQuery } from '../../models/syllabus.model'
import type { Organisation } from '../../models/organisation.model'
import type { Qualification, QualificationQuery } from '../../models/qualification.model'

const router = useRouter()

const paperSeries = ref<PaperSeries[]>([])
const organisations = ref<Organisation[]>([])
const qualifications = ref<Qualification[]>([])
const syllabi = ref<Syllabus[]>([])

const loading = ref(false) // Overall loading state for the table
const dropdownLoading = ref({
  organisations: false,
  qualifications: false,
  syllabi: false,
})

const totalItems = ref(0)
const currentPage = ref(1)
const pageSize = 10
const searchQuery = ref('')
const selectedOrganisationId = ref<number>(0)
const selectedQualificationId = ref<number>(0)
const selectedSyllabusId = ref<number>(0)

// Fetch data for dropdowns
const fetchOrganisationsForDropdown = async () => {
  dropdownLoading.value.organisations = true
  try {
    // Assuming pageSize 1000 is enough to get all, adjust if necessary
    const response = await organisationService.getAllOrganisations({ pageIndex: 1, pageSize: 1000 }) 
    organisations.value = response.data.list
  } catch (error) {
    console.error('Failed to fetch organisations:', error)
    organisations.value = []
  } finally {
    dropdownLoading.value.organisations = false
  }
}

const fetchQualificationsForDropdown = async (organisationId?: number) => {
  dropdownLoading.value.qualifications = true
  qualifications.value = [] // Clear previous before fetching
  try {
    const query: QualificationQuery = { pageIndex: 1, pageSize: 1000 }
    if (organisationId) {
      query.organisationId = organisationId
    }
    const response = await qualificationService.getAllQualifications(query)
    qualifications.value = response.data.list
  } catch (error) {
    console.error('Failed to fetch qualifications:', error)
    qualifications.value = []
  } finally {
    dropdownLoading.value.qualifications = false
  }
}

const fetchSyllabiForDropdown = async (qualificationId?: number) => {
  dropdownLoading.value.syllabi = true
  syllabi.value = [] // Clear previous before fetching
  try {
    const query: SyllabusQuery = { pageIndex: 1, pageSize: 1000 }
    // Removed: if (organisationId) { query.organisationId = organisationId }
    if (qualificationId) { // Filter by qual if provided
      query.qualificationId = qualificationId
    }
    const response = await syllabusService.getAllSyllabuses(query)
    syllabi.value = response.data.list
  } catch (error) {
    console.error('Failed to fetch syllabi:', error)
    syllabi.value = []
  } finally {
    dropdownLoading.value.syllabi = false
  }
}

const fetchPaperSeriesList = async () => {
  loading.value = true
  try {
    const query: PaperSeriesQuery = {
      pageIndex: currentPage.value,
      pageSize,
    }
    // if (selectedOrganisationId.value) { // Removed: This parameter is not supported by PaperSeriesQuery
    //   query.organisationId = selectedOrganisationId.value;
    // }
    // if (selectedQualificationId.value) { // Removed: This parameter is not supported by PaperSeriesQuery
    //   query.qualificationId = selectedQualificationId.value;
    // }
    if (selectedSyllabusId.value) {
      query.syllabusId = selectedSyllabusId.value
    }
    if (searchQuery.value.trim()) { // Removed: This parameter is not supported by PaperSeriesQuery
      query.name = searchQuery.value.trim();
    }
    
    const response = await paperService.getPaperSeriesList(query)
    paperSeries.value = response.data.list
    totalItems.value = response.data.total
  } catch (error) {
    console.error('Failed to fetch paper series:', error)
    paperSeries.value = []
    totalItems.value = 0
  } finally {
    loading.value = false
  }
}

// Pagination
const totalPages = computed(() => Math.ceil(totalItems.value / pageSize))
const paginationRange = computed(() => {
  const range = []
  const maxPagesToShow = 5
  let start = Math.max(1, currentPage.value - Math.floor(maxPagesToShow / 2))
  let end = Math.min(totalPages.value, start + maxPagesToShow - 1)
  if (totalPages.value > maxPagesToShow && end - start + 1 < maxPagesToShow) {
    start = Math.max(1, end - maxPagesToShow + 1)
  }
  for (let i = start; i <= end; i++) {
    range.push(i)
  }
  return range
})

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
    currentPage.value = page
    fetchPaperSeriesList()
  }
}

// Actions
const goToCreatePage = () => {
  router.push('/admin/paper-series/create')
}

const deletePaperSeries = async (id: number) => {
  if (confirm('Are you sure you want to delete this paper series? This action cannot be undone.')) {
    try {
      await paperService.deletePaperSeries(id)
      if (paperSeries.value.length === 1 && currentPage.value > 1) {
        currentPage.value--
      }
      fetchPaperSeriesList()
    } catch (error) {
      console.error('Failed to delete paper series:', error)
      // TODO: Show error notification to user
    }
  }
}

// Watchers
let searchDebounceTimer: number | undefined
watch(searchQuery, () => {
  clearTimeout(searchDebounceTimer)
  searchDebounceTimer = window.setTimeout(() => {
    currentPage.value = 1
    // Client-side filtering if API doesn't support name search for PaperSeries
    // For now, we assume the API handles filtering or we re-fetch and rely on pagination
    // If the API doesn't support search by name for PaperSeries, this will just reload the current view
    // based on dropdowns.
    // To implement client-side search on the current page's data:
    // 1. Fetch all data if feasible, or
    // 2. Filter the 'paperSeries.value' array directly if API doesn't support search.
    // The current setup re-calls fetchPaperSeriesList, which doesn't send 'name' in query.
    // This means the search input currently doesn't filter by name via API.
    // If client-side search is desired on the fetched page:
    // You would need a computed property that filters `paperSeries.value` based on `searchQuery.value`.
    // However, this is often combined with server-side pagination/filtering for consistency.
    // For now, leaving as is, which means search input doesn't filter server-side.
    fetchPaperSeriesList() 
  }, 500)
})

watch(selectedOrganisationId, async (newOrgId) => {
  selectedQualificationId.value = 0 
  selectedSyllabusId.value = 0    
  qualifications.value = []       
  syllabi.value = []              
  currentPage.value = 1
  
  await fetchQualificationsForDropdown(newOrgId || undefined) 
  // When Org changes, fetch syllabi for that org (or all if "All Orgs")
  // If "All Orgs", Quals will be all, Syllabi will be all.
  // If specific Org, Quals for that Org, Syllabi for that Org (indirectly via Quals).
  // The fetchSyllabiForDropdown needs to be smart or be called after Quals are set.
  // Current logic: fetchSyllabiForDropdown is called with newOrgId and undefined QualId.
  // This might not be what we want if we want syllabi filtered by the new org directly.
  // However, PaperSeriesQuery only takes syllabusId. So dropdowns are for narrowing down syllabusId.
  await fetchSyllabiForDropdown(undefined) 
  
  fetchPaperSeriesList() 
})

watch(selectedQualificationId, async (newQualId) => {
  selectedSyllabusId.value = 0 
  syllabi.value = []           
  currentPage.value = 1
  
  await fetchSyllabiForDropdown(newQualId || undefined)
  fetchPaperSeriesList() 
})

watch(selectedSyllabusId, () => {
  currentPage.value = 1
  fetchPaperSeriesList() 
})

// Lifecycle
const fetchInitialDropdownData = async () => {
  await fetchOrganisationsForDropdown()
  // Fetch all qualifications and syllabi initially if no org/qual is pre-selected
  await fetchQualificationsForDropdown() 
  await fetchSyllabiForDropdown() 
}

onMounted(async () => {
  loading.value = true 
  dropdownLoading.value.organisations = true
  dropdownLoading.value.qualifications = true
  dropdownLoading.value.syllabi = true

  await fetchInitialDropdownData()
  await fetchPaperSeriesList() 
  
  // Set loading to false after all initial data is fetched
  // loading.value = false; // This was here, but fetchPaperSeriesList has its own finally block.
  // Dropdown loading should be false after their respective fetches.
})

onUnmounted(() => {
  clearTimeout(searchDebounceTimer)
})
</script>
