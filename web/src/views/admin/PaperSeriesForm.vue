<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">
        {{ isEditMode ? $t('paperSeries.edit') : $t('paperSeries.add') }}
      </h1>
    </header>

    <form @submit.prevent="handleSubmit" class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
      <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
          {{ $t('paperSeries.name') }}
        </label>
        <input
          id="name"
          type="text"
          v-model="form.name"
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          :class="{ 'border-red-500': errors.name }"
          required
        />
        <p v-if="errors.name" class="text-red-500 text-xs italic">{{ errors.name }}</p>
      </div>

      <!-- Organisation and Qualification filters for Syllabus dropdown -->
       <div class="mb-4 grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label class="block text-gray-700 text-sm font-bold mb-2" for="filter-organisation">
            {{ $t('syllabusForm.organisation') }}
          </label>
          <select
            id="filter-organisation"
            v-model.number="filterOrganisationId"
            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          >
            <option :value="0">{{ $t('syllabusManagement.allOrganisations') }}</option>
            <option v-for="org in organisations" :key="org.id" :value="org.id">{{ org.name }}</option>
          </select>
        </div>
        <div>
          <label class="block text-gray-700 text-sm font-bold mb-2" for="filter-qualification">
            {{ $t('syllabusForm.qualification') }}
          </label>
          <select
            id="filter-qualification"
            v-model.number="filterQualificationId"
            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            :disabled="!filterOrganisationId && !qualificationsForFilter.length"
          >
            <option :value="0">{{ $t('syllabusManagement.allQualifications') }}</option>
            <option v-for="q in qualificationsForFilter" :key="q.id" :value="q.id">{{ q.name }}</option>
          </select>
        </div>
      </div>

      <div class="mb-6">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="syllabus">
          {{ $t('syllabusManagement.name') }}
        </label>
        <select
          id="syllabus"
          v-model.number="form.syllabusId"
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          :class="{ 'border-red-500': errors.syllabusId }"
          required
          :disabled="!availableSyllabuses.length"
        >
          <option :value="0" disabled>
            {{ availableSyllabuses.length ? $t('paperSeries.selectSyllabus') : $t('paperSeries.noSyllabusMatch') }}
          </option>
          <option v-for="syllabus in availableSyllabuses" :key="syllabus.id" :value="syllabus.id">
            {{ syllabus.name }}
          </option>
        </select>
        <p v-if="errors.syllabusId" class="text-red-500 text-xs italic">{{ errors.syllabusId }}</p>
      </div>

      <div class="flex items-center justify-between">
        <button
          type="submit"
          class="bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
          :disabled="isSubmitting"
        >
          {{ isEditMode ? $t('paperSeries.edit') : $t('paperSeries.add') }}
        </button>
        <router-link
          to="/admin/paper-series"
          class="inline-block align-baseline font-bold text-sm text-indigo-600 hover:text-indigo-800"
        >
          {{ $t('common.cancel') }}
        </router-link>
      </div>
      <p v-if="submitError" class="text-red-500 text-xs italic mt-4">{{ submitError }}</p>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import paperService from '../../services/paperSeriesService'
import syllabusService from '../../services/syllabusService'
import organisationService from '../../services/organisationService'
import qualificationService from '../../services/qualificationService'
import type { PaperSeries } from '../../models/paperSeries.model'
import type { Syllabus } from '../../models/syllabus.model'
import type { Organisation } from '../../models/organisation.model'
import type { Qualification } from '../../models/qualification.model'

const route = useRoute()
const router = useRouter()

const form = ref<Partial<PaperSeries>>({
  name: '',
  syllabusId: 0,
})
const errors = ref<{ name?: string; syllabusId?: string }>({})
const submitError = ref<string | null>(null)
const isSubmitting = ref(false)

const isEditMode = computed(() => !!route.params.id)
const paperSeriesId = computed(() => Number(route.params.id) || null)

const organisations = ref<Organisation[]>([])
const allQualifications = ref<Qualification[]>([]) // Store all qualifications
const allSyllabuses = ref<Syllabus[]>([]) // Store all syllabuses

const filterOrganisationId = ref<number>(0)
const filterQualificationId = ref<number>(0)

// Fetch initial data for dropdowns and form (if editing)
const fetchOrganisations = async () => {
  try {
    const response = await organisationService.getAllOrganisations({})
    organisations.value = response.data.list
  } catch (error) {
    console.error('Failed to fetch organisations:', error)
  }
}

const fetchAllQualifications = async () => {
  try {
    const response = await qualificationService.getAllQualifications({ pageIndex: 1, pageSize: 10000 })
    allQualifications.value = response.data.list
  } catch (error) {
    console.error('Failed to fetch qualifications:', error)
  }
}

const fetchAllSyllabuses = async () => {
  try {
    const response = await syllabusService.getAllSyllabuses({ pageIndex: 1, pageSize: 10000 })
    allSyllabuses.value = response.data.list
  } catch (error) {
    console.error('Failed to fetch syllabuses:', error)
  }
}

const fetchPaperSeriesDetails = async (id: number) => {
  try {
    const response = await paperService.getPaperSeriesById(id)
    const seriesData = response.data
    form.value = {
      id: seriesData.id,
      name: seriesData.name,
      syllabusId: seriesData.syllabusId,
    }
    // Pre-fill filters if syllabus is set
    if (seriesData.syllabus && seriesData.syllabus.qualification) {
      filterOrganisationId.value = seriesData.syllabus.qualification.organisationId || 0
      // Wait for qualifications to be filtered by organisation before setting qualification filter
      await new Promise(resolve => setTimeout(resolve, 0)); // Allow computed prop to update
      filterQualificationId.value = seriesData.syllabus.qualificationId || 0
    }
  } catch (error) {
    console.error('Failed to fetch paper series details:', error)
    submitError.value = 'Failed to load paper series data.'
  }
}

// Computed properties for filtering dropdowns
const qualificationsForFilter = computed(() => {
  if (!filterOrganisationId.value) {
    return allQualifications.value // Or an empty array if you prefer no qualifications shown without org
  }
  return allQualifications.value.filter(q => q.organisationId === filterOrganisationId.value)
})

const availableSyllabuses = computed(() => {
  let syllabusesToFilter = allSyllabuses.value
  if (filterOrganisationId.value) {
    syllabusesToFilter = syllabusesToFilter.filter(s => s.qualification?.organisationId === filterOrganisationId.value)
  }
  if (filterQualificationId.value) {
    syllabusesToFilter = syllabusesToFilter.filter(s => s.qualificationId === filterQualificationId.value)
  }
  return syllabusesToFilter
})

// Watchers to reset dependent filters
watch(filterOrganisationId, () => {
  filterQualificationId.value = 0
  // If the current form.syllabusId is not in the new availableSyllabuses, reset it
  if (form.value.syllabusId && !availableSyllabuses.value.find(s => s.id === form.value.syllabusId)) {
    form.value.syllabusId = 0
  }
})

watch(filterQualificationId, () => {
 // If the current form.syllabusId is not in the new availableSyllabuses, reset it
  if (form.value.syllabusId && !availableSyllabuses.value.find(s => s.id === form.value.syllabusId)) {
    form.value.syllabusId = 0
  }
})


const validateForm = (): boolean => {
  errors.value = {}
  if (!form.value.name?.trim()) {
    errors.value.name = 'Name is required.'
  }
  if (!form.value.syllabusId || form.value.syllabusId === 0) {
    errors.value.syllabusId = 'Syllabus is required.'
  }
  return Object.keys(errors.value).length === 0
}

const handleSubmit = async () => {
  if (!validateForm()) {
    return
  }
  isSubmitting.value = true
  submitError.value = null
  try {
    const payload = {
      name: form.value.name!,
      syllabusId: form.value.syllabusId!,
    }
    if (isEditMode.value && paperSeriesId.value) {
      await paperService.updatePaperSeries({ ...payload, id: paperSeriesId.value })
    } else {
      await paperService.createPaperSeries(payload)
    }
    router.push('/admin/paper-series')
  } catch (error: any) {
    console.error('Failed to submit paper series:', error)
    submitError.value = error.response?.data?.message || 'An unexpected error occurred.'
  } finally {
    isSubmitting.value = false
  }
}

onMounted(async () => {
  await Promise.all([
    fetchOrganisations(),
    fetchAllQualifications(),
    fetchAllSyllabuses(),
  ])
  if (isEditMode.value && paperSeriesId.value) {
    await fetchPaperSeriesDetails(paperSeriesId.value)
  }
})
</script>
