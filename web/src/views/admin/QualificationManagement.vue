<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('qualificationManagement.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('qualificationManagement.subtitle') }}</p>
    </header>

    <div class="mb-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0 sm:space-x-4">
      <div class="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 w-full sm:w-auto">
        <input 
          type="text" 
          v-model="searchQuery"
          :placeholder="$t('qualificationManagement.searchByName')" 
          class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full sm:w-auto"
        />
        <select 
          v-model="selectedOrganisationId"
          class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full sm:w-auto"
        >
          <option value="">{{ $t('qualificationManagement.allOrganisations') }}</option>
          <option v-for="org in organisations" :key="org.id" :value="org.id">{{ org.name }}</option>
        </select>
      </div>
      <router-link 
        to="/admin/qualifications/create"
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 w-full sm:w-auto"
      >
        {{ $t('qualificationManagement.addQualification') }}
      </router-link>
    </div>

    <div class="bg-white shadow overflow-x-auto sm:rounded-lg">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('qualificationManagement.name') }}</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('qualificationManagement.organisation') }}</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider min-w-[120px]">{{ $t('qualificationManagement.actions') }}</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-if="loading">
            <td colspan="4" class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center">{{ $t('qualificationManagement.loading') }}</td>
          </tr>
          <tr v-else-if="!qualifications || qualifications.length === 0">
            <td colspan="4" class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center">{{ $t('qualificationManagement.noData') }}</td>
          </tr>
          <tr v-for="qual in qualifications" :key="qual.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ qual.id }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ qual.name }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ qual.organisation?.name || '-' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2">
              <router-link 
                :to="`/admin/qualifications/${qual.id}/edit`" 
                class="text-indigo-600 hover:text-indigo-900"
              >
                {{ $t('qualificationManagement.edit') }}
              </router-link>
              <button 
                @click="deleteQualification(qual.id)"
                class="text-red-600 hover:text-red-900"
              >
                {{ $t('qualificationManagement.delete') }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    
    <div v-if="!loading && totalQualifications > 0" class="mt-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0">
      <p class="text-sm text-gray-700">
        {{ $t('qualificationManagement.pageInfo', { from: (currentPage - 1) * pageSize + 1, to: Math.min(currentPage * pageSize, totalQualifications), total: totalQualifications }) }}
      </p>
      <nav v-if="totalPages > 1" class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
        <button
          @click="goToPage(currentPage - 1)"
          :disabled="currentPage === 1"
          class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ $t('qualificationManagement.previous') }}
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
          {{ $t('qualificationManagement.next') }}
        </button>
      </nav>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import qualificationService from '../../services/qualificationService';
import organisationService from '../../services/organisationService';
import type { Qualification } from '../../models/qualification.model';
import type { Organisation } from '../../models/organisation.model';

// Data
const qualifications = ref<Qualification[]>([]);
const organisations = ref<Organisation[]>([]);
const loading = ref(true);
const totalQualifications = ref(0);
const currentPage = ref(1);
const pageSize = 10;
const searchQuery = ref('');
const selectedOrganisationId = ref<number | ''>('');

// Fetch qualifications
const fetchQualifications = async () => {
  loading.value = true;
  try {
    const response = await qualificationService.getQualifications({
      pageIndex: currentPage.value,
      pageSize,
      organisationId: selectedOrganisationId.value || undefined,
    });
    qualifications.value = response.data.list;
    totalQualifications.value = response.data.total;
  } catch (error) {
    console.error('Failed to fetch qualifications:', error);
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

const deleteQualification = async (id: number) => {
  if (confirm('Are you sure you want to delete this qualification? This action cannot be undone.')) {
    try {
      await qualificationService.deleteQualification(id);
      if (qualifications.value.length === 1 && currentPage.value > 1) {
        currentPage.value--;
      }
      fetchQualifications(); // Refresh list
    } catch (error) {
      console.error('Failed to delete qualification:', error);
      // TODO: Show error message to user
    }
  }
};

// Pagination
const totalPages = computed(() => {
  return Math.ceil(totalQualifications.value / pageSize);
});

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
    currentPage.value = page;
    fetchQualifications();
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
watch([searchQuery, selectedOrganisationId], () => {
  clearTimeout(searchDebounceTimer);
  searchDebounceTimer = window.setTimeout(() => {
    currentPage.value = 1;
    fetchQualifications();
  }, 500);
});

onMounted(() => {
  fetchQualifications();
  fetchOrganisations();
});
</script>
