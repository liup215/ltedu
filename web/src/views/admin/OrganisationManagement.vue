<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('organisationManagement.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('organisationManagement.subtitle') }}</p>
    </header>

    <div class="mb-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0 sm:space-x-4">
      <div class="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 w-full sm:w-auto">
        <input 
          type="text" 
          v-model="searchQuery"
          :placeholder="$t('organisationManagement.searchByName')" 
          class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full sm:w-auto"
        />
      </div>
      <router-link 
        to="/admin/organisations/create"
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 w-full sm:w-auto"
      >
        {{ $t('organisationManagement.addOrganisation') }}
      </router-link>
    </div>

    <div class="bg-white shadow overflow-x-auto sm:rounded-lg">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('organisationManagement.name') }}</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider min-w-[120px]">{{ $t('organisationManagement.actions') }}</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-if="loading">
            <td colspan="3" class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center">{{ $t('organisationManagement.loading') }}</td>
          </tr>
          <tr v-else-if="!organisations || organisations.length === 0">
            <td colspan="3" class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center">{{ $t('organisationManagement.noData') }}</td>
          </tr>
          <tr v-for="org in organisations" :key="org.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ org.id }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ org.name }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
              <router-link 
                :to="`/admin/organisations/${org.id}/edit`" 
                class="text-indigo-600 hover:text-indigo-900"
              >
                {{ $t('organisationManagement.edit') }}
              </router-link>
              <button 
                @click="deleteOrganisation(org.id)"
                class="text-red-600 hover:text-red-900"
              >
                {{ $t('organisationManagement.delete') }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    
    <div v-if="!loading && totalOrganisations > 0" class="mt-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0">
      <p class="text-sm text-gray-700">
        {{ $t('organisationManagement.pageInfo', { from: (currentPage - 1) * pageSize + 1, to: Math.min(currentPage * pageSize, totalOrganisations), total: totalOrganisations }) }}
      </p>
      <nav v-if="totalPages > 1" class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
        <button
          @click="goToPage(currentPage - 1)"
          :disabled="currentPage === 1"
          class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ $t('organisationManagement.previous') }}
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
          {{ $t('organisationManagement.next') }}
        </button>
      </nav>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import organisationService from '../../services/organisationService';
import type { Organisation } from '../../models/organisation.model';

// Data
const organisations = ref<Organisation[]>([]);
const loading = ref(true);
const totalOrganisations = ref(0);
const currentPage = ref(1);
const pageSize = 10;
const searchQuery = ref('');

// Fetch organisations
const fetchOrganisations = async () => {
  loading.value = true;
  try {
    const response = await organisationService.getOrganisations({
      pageIndex: currentPage.value,
      pageSize,
      name: searchQuery.value.trim() || undefined,
    });
    console.log('Fetched organisations:', response);
    organisations.value = response.data.list;
    totalOrganisations.value = response.data.total;
  } catch (error) {
    console.error('Failed to fetch organisations:', error);
    // TODO: Show error message to user
  } finally {
    loading.value = false;
  }
};

const deleteOrganisation = async (id: number) => {
  if (confirm('Are you sure you want to delete this organisation? This action cannot be undone.')) {
    try {
      await organisationService.deleteOrganisation(id);
      if (organisations.value.length === 1 && currentPage.value > 1) {
        currentPage.value--;
      }
      fetchOrganisations(); // Refresh list
    } catch (error) {
      console.error('Failed to delete organisation:', error);
      // TODO: Show error message to user
    }
  }
};

// Pagination
const totalPages = computed(() => {
  return Math.ceil(totalOrganisations.value / pageSize);
});

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
    currentPage.value = page;
    fetchOrganisations();
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
watch(searchQuery, () => {
  clearTimeout(searchDebounceTimer);
  searchDebounceTimer = window.setTimeout(() => {
    currentPage.value = 1;
    fetchOrganisations();
  }, 500);
});

onMounted(() => {
  fetchOrganisations();
});
</script>
