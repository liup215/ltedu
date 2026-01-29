<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ isEdit ? $t('syllabusForm.editTitle') : $t('syllabusForm.createTitle') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ isEdit ? $t('syllabusForm.editSubtitle') : $t('syllabusForm.createSubtitle') }}</p>
    </header>

    <div class="bg-white shadow sm:rounded-lg">
      <form @submit.prevent="onSubmit" class="space-y-8 divide-y divide-gray-200 px-8 py-6">
        <div class="space-y-6">
          <div>
            <h3 class="text-lg font-medium leading-6 text-gray-900">{{ $t('syllabusForm.sectionTitle') }}</h3>
            <p class="mt-1 text-sm text-gray-500">
              {{ $t('syllabusForm.sectionTip') }}
            </p>
          </div>

          <div class="space-y-6">
            <div class="sm:w-3/4">
              <label for="organisation" class="block text-sm font-medium text-gray-700 mb-1">{{ $t('syllabusForm.organisation') }}</label>
              <select
                id="organisation"
                v-model="selectedOrganisationId"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-base py-3"
                required
              >
                <option value="">{{ $t('syllabusForm.selectOrganisation') }}</option>
                <option v-for="org in organisations" :key="org.id" :value="org.id">{{ org.name }}</option>
              </select>
            </div>

            <div class="sm:w-3/4">
              <label for="qualification" class="block text-sm font-medium text-gray-700 mb-1">{{ $t('syllabusForm.qualification') }}</label>
              <select
                id="qualification"
                v-model="form.qualificationId"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-base py-3"
                required
                :disabled="!selectedOrganisationId"
              >
                <option value="">{{ $t('syllabusForm.selectQualification') }}</option>
                <option v-for="qual in filteredQualifications" :key="qual.id" :value="qual.id">{{ qual.name }}</option>
              </select>
            </div>

            <div class="sm:w-3/4">
              <label for="name" class="block text-sm font-medium text-gray-700 mb-1">{{ $t('syllabusForm.name') }}</label>
              <input
                type="text"
                id="name"
                v-model="form.name"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-base py-3"
                required
              />
            </div>

            <div class="sm:w-3/4">
              <label for="code" class="block text-sm font-medium text-gray-700 mb-1">{{ $t('syllabusForm.code') }}</label>
              <input
                type="text"
                id="code"
                v-model="form.code"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-base py-3"
                required
              />
            </div>
          </div>
        </div>

        <div class="pt-6 flex justify-end space-x-3">
          <router-link
            to="/admin/syllabuses"
            class="inline-flex justify-center rounded-md border border-gray-300 bg-white py-3 px-6 text-base font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
          >
            {{ $t('syllabusForm.cancel') }}
          </router-link>
          <button
            type="submit"
            class="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-3 px-6 text-base font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
            :disabled="loading"
          >
            {{ loading ? $t('syllabusForm.saving') : $t('syllabusForm.save') }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import syllabusService from '../../services/syllabusService';
import qualificationService from '../../services/qualificationService';
import organisationService from '../../services/organisationService';
import type { Qualification } from '../../models/qualification.model';
import type { Organisation } from '../../models/organisation.model';
import type { SyllabusCreateRequest, SyllabusUpdateRequest } from '../../models/syllabus.model';

const route = useRoute();
const router = useRouter();

const loading = ref(false);
const organisations = ref<Organisation[]>([]);
const qualifications = ref<Qualification[]>([]);
const selectedOrganisationId = ref<number | ''>('');

const isEdit = computed(() => route.params.id !== undefined);

const filteredQualifications = computed(() => {
  if (!selectedOrganisationId.value) return [];
  return qualifications.value.filter(q => q.organisationId === selectedOrganisationId.value);
});

const form = ref<SyllabusCreateRequest | SyllabusUpdateRequest>({
  name: '',
  code: '',
  qualificationId: 0,
});

// Load organisations for dropdown
const loadOrganisations = async () => {
  try {
    const response = await organisationService.getAllOrganisations({});
    organisations.value = response.data.list;
  } catch (error) {
    console.error('Failed to load organisations:', error);
  }
};

// Load qualifications
const loadQualifications = async () => {
  try {
    const response = await qualificationService.getAllQualifications();
    qualifications.value = response.data.list;
  } catch (error) {
    console.error('Failed to load qualifications:', error);
  }
};

// Reset qualification when organisation changes
watch(selectedOrganisationId, () => {
  form.value.qualificationId = 0;
});

// Load syllabus data if editing
const loadSyllabus = async () => {
  if (!isEdit.value) return;
  
  const id = Number(route.params.id);
  if (!id) return;

  loading.value = true;
  try {
    const response = await syllabusService.getSyllabusById(id);
    const syllabus = response.data;
    form.value = {
      id,
      name: syllabus.name,
      code: syllabus.code,
      qualificationId: syllabus.qualificationId,
    };
    
    // Set organisation if editing
    if (syllabus.qualification?.organisationId) {
      selectedOrganisationId.value = syllabus.qualification.organisationId;
    }
  } catch (error) {
    console.error('Failed to load syllabus:', error);
  } finally {
    loading.value = false;
  }
};

const onSubmit = async () => {
  loading.value = true;
  try {
    if (isEdit.value) {
      await syllabusService.updateSyllabus(form.value as SyllabusUpdateRequest);
    } else {
      await syllabusService.createSyllabus(form.value as SyllabusCreateRequest);
    }
    router.push('/admin/syllabuses');
  } catch (error) {
    console.error('Failed to save syllabus:', error);
    // TODO: Show error message
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  loadOrganisations();
  loadQualifications();
  loadSyllabus();
});
</script>
