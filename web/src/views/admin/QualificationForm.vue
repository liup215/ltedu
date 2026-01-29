<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ isEdit ? $t('qualificationForm.editTitle') : $t('qualificationForm.createTitle') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ isEdit ? $t('qualificationForm.editSubtitle') : $t('qualificationForm.createSubtitle') }}</p>
    </header>

    <div class="bg-white shadow sm:rounded-lg p-6">
      <form @submit.prevent="handleSubmit">
        <div class="space-y-6">
          <!-- Name -->
          <div>
            <label for="name" class="block text-sm font-medium text-gray-700">{{ $t('qualificationForm.name') }}</label>
            <div class="mt-1">
              <input 
                type="text"
                id="name"
                v-model="formData.name"
                class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                :placeholder="$t('qualificationForm.namePlaceholder')"
                required
              />
            </div>
          </div>

          <!-- Organisation -->
          <div>
            <label for="organisationId" class="block text-sm font-medium text-gray-700">{{ $t('qualificationForm.organisation') }}</label>
            <div class="mt-1">
              <select
                id="organisationId"
                v-model="formData.organisationId"
                class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                required
              >
                <option value="">{{ $t('qualificationForm.selectOrganisation') }}</option>
                <option v-for="org in organisations" :key="org.id" :value="org.id">{{ org.name }}</option>
              </select>
            </div>
          </div>

          <!-- Buttons -->
          <div class="flex justify-end space-x-4">
            <button 
              type="button"
              @click="goBack"
              class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              {{ $t('qualificationForm.cancel') }}
            </button>
            <button 
              type="submit"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              :disabled="loading"
            >
              <svg v-if="loading" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              {{ loading ? $t('qualificationForm.saving') : $t('qualificationForm.save') }}
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import qualificationService from '../../services/qualificationService';
import organisationService from '../../services/organisationService';
import type { Organisation } from '../../models/organisation.model';

const route = useRoute();
const router = useRouter();
const loading = ref(false);

// Check if we're in edit mode
const isEdit = computed(() => !!route.params.id);

// Form data
const formData = ref({
  name: '',
  organisationId: '' as number | ''
});

// Organisations for dropdown
const organisations = ref<Organisation[]>([]);

// Load organisations
const fetchOrganisations = async () => {
  try {
    const response = await organisationService.getAllOrganisations({});
    organisations.value = response.data.list;
  } catch (error) {
    console.error('Failed to load organisations:', error);
    // TODO: Show error message
  }
};

// Load qualification data if in edit mode
onMounted(async () => {
  await fetchOrganisations();
  
  if (isEdit.value) {
    loading.value = true;
    try {
      const response = await qualificationService.getQualificationById(Number(route.params.id));
      formData.value.name = response.data.name;
      formData.value.organisationId = response.data.organisationId;
    } catch (error) {
      console.error('Failed to load qualification:', error);
      // TODO: Show error message
    } finally {
      loading.value = false;
    }
  }
});

// Form submission
const handleSubmit = async () => {
  if (!formData.value.name.trim() || !formData.value.organisationId) {
    // TODO: Show validation error
    return;
  }

  loading.value = true;
  try {
    if (isEdit.value) {
      await qualificationService.updateQualification({
        id: Number(route.params.id),
        name: formData.value.name.trim(),
        organisationId: Number(formData.value.organisationId),
      });
    } else {
      await qualificationService.createQualification({
        name: formData.value.name.trim(),
        organisationId: Number(formData.value.organisationId),
      });
    }
    router.push('/admin/qualifications');
  } catch (error) {
    console.error('Failed to save qualification:', error);
    // TODO: Show error message
  } finally {
    loading.value = false;
  }
};

// Navigation
const goBack = () => {
  router.push('/admin/qualifications');
};
</script>
