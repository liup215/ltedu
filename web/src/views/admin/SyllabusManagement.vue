<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('syllabusManagement.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('syllabusManagement.subtitle') }}</p>
    </header>

    <div class="mb-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0 sm:space-x-4">
      <div class="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 w-full sm:w-auto">
        <input 
          type="text" 
          v-model="searchQuery"
          :placeholder="$t('syllabusManagement.searchByNameOrCode')" 
          class="px-3 py-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-base w-full sm:w-auto"
        />
        <select 
          v-model="selectedOrganisationId"
          class="px-3 py-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-base w-full sm:w-auto"
        >
          <option value="">{{ $t('syllabusManagement.allOrganisations') }}</option>
          <option v-for="org in organisations" :key="org.id" :value="org.id">{{ org.name }}</option>
        </select>
        <select 
          v-model="selectedQualificationId"
          class="px-3 py-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-base w-full sm:w-auto"
          :disabled="!selectedOrganisationId"
        >
          <option value="">{{ $t('syllabusManagement.allQualifications') }}</option>
          <option v-for="qual in filteredQualifications" :key="qual.id" :value="qual.id">{{ qual.name }}</option>
        </select>
      </div>
      <router-link 
        to="/admin/syllabuses/create"
        class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 w-full sm:w-auto"
      >
        {{ $t('syllabusManagement.addSyllabus') }}
      </router-link>
    </div>

    <div class="bg-white shadow overflow-x-auto sm:rounded-lg">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('syllabusManagement.name') }}</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('syllabusManagement.code') }}</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('syllabusManagement.organisation') }}</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('syllabusManagement.qualification') }}</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider min-w-[180px]">{{ $t('syllabusManagement.actions') }}</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-if="loading">
            <td colspan="6" class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center">{{ $t('syllabusManagement.loading') }}</td>
          </tr>
          <tr v-else-if="!syllabuses || syllabuses.length === 0">
            <td colspan="6" class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center">{{ $t('syllabusManagement.noData') }}</td>
          </tr>
          <tr v-for="syllabus in syllabuses" :key="syllabus.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ syllabus.id }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ syllabus.name }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ syllabus.code }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ syllabus.qualification?.organisation?.name || '-' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ syllabus.qualification?.name || '-' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-3">
              <router-link 
                :to="`/admin/syllabuses/${syllabus.id}/chapters`" 
                class="text-indigo-600 hover:text-indigo-900"
              >
                {{ $t('syllabusManagement.manageChapters') }}
              </router-link>
              <router-link 
                :to="`/admin/syllabuses/${syllabus.id}/exam-nodes`" 
                class="text-orange-600 hover:text-orange-900"
              >
                {{ $t('examNode.title') }}
              </router-link>
              <router-link 
                :to="`/admin/syllabuses/${syllabus.id}/knowledge-points`" 
                class="text-purple-600 hover:text-purple-900"
              >
                {{ $t('syllabusManagement.knowledgePoints') }}
              </router-link>
              <router-link 
                :to="`/admin/syllabuses/${syllabus.id}/edit`" 
                class="text-indigo-600 hover:text-indigo-900"
              >
                {{ $t('syllabusManagement.edit') }}
              </router-link>
              <button 
                @click="deleteSyllabus(syllabus.id)"
                class="text-red-600 hover:text-red-900"
              >
                {{ $t('syllabusManagement.delete') }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    
    <div v-if="!loading && totalSyllabuses > 0" class="mt-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0">
      <p class="text-sm text-gray-700">
        {{ $t('syllabusManagement.pageInfo', { from: (currentPage - 1) * pageSize + 1, to: Math.min(currentPage * pageSize, totalSyllabuses), total: totalSyllabuses }) }}
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
import { ref, onMounted, computed, watch } from 'vue';
import syllabusService from '../../services/syllabusService';
import qualificationService from '../../services/qualificationService';
import organisationService from '../../services/organisationService';
import type { Syllabus } from '../../models/syllabus.model';
import type { Qualification } from '../../models/qualification.model';
import type { Organisation } from '../../models/organisation.model';

// Data
const syllabuses = ref<Syllabus[]>([]);
const organisations = ref<Organisation[]>([]);
const qualifications = ref<Qualification[]>([]);
const loading = ref(true);
const totalSyllabuses = ref(0);
const currentPage = ref(1);
const pageSize = 10;
const searchQuery = ref('');
const selectedOrganisationId = ref<number | ''>('');
const selectedQualificationId = ref<number | ''>('');

// Computed
const filteredQualifications = computed(() => {
  if (!selectedOrganisationId.value) return [];
  return qualifications.value.filter(q => q.organisationId === selectedOrganisationId.value);
});

// Fetch syllabuses
const fetchSyllabuses = async () => {
  loading.value = true;
  try {
    const response = await syllabusService.getSyllabuses({
      pageIndex: currentPage.value,
      pageSize,
      qualificationId: selectedQualificationId.value || undefined,
    });
    syllabuses.value = response.data.list;
    totalSyllabuses.value = response.data.total;
  } catch (error) {
    console.error('Failed to fetch syllabuses:', error);
    // TODO: Show error message to user
  } finally {
    loading.value = false;
  }
};

// Fetch organisations for filter dropdown
const fetchOrganisations = async () => {
  try {
    const response = await organisationService.getAllOrganisations({});
    organisations.value = response.data.list;
  } catch (error) {
    console.error('Failed to fetch organisations:', error);
  }
};

// Fetch qualifications
const fetchQualifications = async () => {
  try {
    const response = await qualificationService.getAllQualifications();
    qualifications.value = response.data.list;
  } catch (error) {
    console.error('Failed to fetch qualifications:', error);
  }
};

const deleteSyllabus = async (id: number) => {
  if (confirm('Are you sure you want to delete this syllabus? This action cannot be undone.')) {
    try {
      await syllabusService.deleteSyllabus(id);
      if (syllabuses.value.length === 1 && currentPage.value > 1) {
        currentPage.value--;
      }
      fetchSyllabuses(); // Refresh list
    } catch (error) {
      console.error('Failed to delete syllabus:', error);
      // TODO: Show error message to user
    }
  }
};

// Reset qualification when organisation changes
watch(selectedOrganisationId, () => {
  selectedQualificationId.value = '';
});

// Pagination
const totalPages = computed(() => {
  return Math.ceil(totalSyllabuses.value / pageSize);
});

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
    currentPage.value = page;
    fetchSyllabuses();
  }
};

const paginationRange = computed(() => {
  const range = [];
  const maxPagesToShow = 5;
  let start = Math.max(1, currentPage.value - Math.floor(maxPagesToShow / 2));
  let end = Math.min(totalPages.value, start + maxPagesToShow - 1);

  if (totalPages.value > 0 && end - start + 1 < maxPagesToShow) {
    if (currentPage.value <= Math.floor(maxPagesToShow / 2)) {
      end = Math.min(totalPages.value, maxPagesToShow);
    } else {
      start = Math.max(1, totalPages.value - maxPagesToShow + 1);
    }
  }

  if (totalPages.value > 0 && end - start + 1 < maxPagesToShow && start === 1 && totalPages.value < maxPagesToShow) {
    end = totalPages.value;
  }

  for (let i = start; i <= end; i++) {
    if (i > 0) range.push(i);
  }
  return range;
});

// Search debounce
let searchDebounceTimer: number | undefined;
watch([searchQuery, selectedOrganisationId, selectedQualificationId], () => {
  clearTimeout(searchDebounceTimer);
  searchDebounceTimer = window.setTimeout(() => {
    currentPage.value = 1;
    fetchSyllabuses();
  }, 500);
});

onMounted(() => {
  fetchOrganisations();
  fetchQualifications();
  fetchSyllabuses();
});
</script>
