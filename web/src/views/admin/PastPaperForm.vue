<template>
  <div class="p-6 bg-gray-50 min-h-screen">
    <div class="bg-white p-8 rounded-lg shadow-md max-w-2xl mx-auto">
      <h1 class="text-2xl font-semibold text-gray-700 mb-6">{{ isEditMode ? $t('pastPaper.edit') : $t('pastPaper.add') }}</h1>

      <div v-if="isLoading && isEditMode && !initialDataLoaded" class="text-center py-4">
        <p class="text-gray-600">{{ $t('common.loading') }}</p>
      </div>
      <div v-if="successMessage" class="mb-4 p-3 bg-green-100 text-green-700 rounded-md">
        {{ successMessage }}
      </div>
      <div v-if="errorMessage" class="mb-4 p-3 bg-red-100 text-red-700 rounded-md">
        {{ errorMessage }}
      </div>

      <form @submit.prevent="handleSubmit" class="space-y-6">
        <!-- Name input at top -->
        <div>
          <label for="name" class="block text-sm font-medium text-gray-700">{{ $t('pastPaper.name') }} <span class="text-red-500">*</span></label>
          <input
            type="text"
            id="name"
            v-model="form.name"
            required
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
          />
        </div>

        <!-- Organisation Dropdown -->
        <div>
          <label for="organisation" class="block text-sm font-medium text-gray-700">{{ $t('pastPaper.organisation') }} <span class="text-red-500">*</span></label>
          <select
            id="organisation"
            v-model="selectedOrganisationId"
            required
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
            :disabled="isLoadingDropdowns.organisations"
          >
            <option :value="null" disabled>{{ isLoadingDropdowns.organisations ? $t('common.loading') : $t('pastPaper.organisationPlaceholder') }}</option>
            <option v-for="org in organisations" :key="org.id" :value="org.id">
              {{ org.name }}
            </option>
          </select>
        </div>

        <!-- Qualification Dropdown -->
        <div>
          <label for="qualification" class="block text-sm font-medium text-gray-700">{{ $t('pastPaper.qualification') }} <span class="text-red-500">*</span></label>
          <select
            id="qualification"
            v-model="selectedQualificationId"
            required
            :disabled="!selectedOrganisationId || isLoadingDropdowns.qualifications"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm disabled:bg-gray-100"
          >
            <option :value="null" disabled>{{ isLoadingDropdowns.qualifications ? $t('common.loading') : (selectedOrganisationId ? $t('pastPaper.qualificationPlaceholder') : $t('pastPaper.organisationPlaceholder')) }}</option>
            <option v-for="qual in qualifications" :key="qual.id" :value="qual.id">
              {{ qual.name }}
            </option>
          </select>
        </div>

        <!-- Syllabus Dropdown -->
        <div>
          <label for="syllabus" class="block text-sm font-medium text-gray-700">{{ $t('pastPaper.syllabus') }} <span class="text-red-500">*</span></label>
          <select
            id="syllabus"
            v-model="selectedSyllabusId"
            required
            :disabled="!selectedQualificationId || isLoadingDropdowns.syllabuses"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm disabled:bg-gray-100"
          >
            <option :value="null" disabled>{{ isLoadingDropdowns.syllabuses ? $t('common.loading') : (selectedQualificationId ? $t('pastPaper.selectSyllabus') : $t('pastPaper.selectQualification')) }}</option>
            <option v-for="syl in syllabuses" :key="syl.id" :value="syl.id">
              {{ syl.name }} ({{ syl.code }})
            </option>
          </select>
        </div>

        <!-- Paper Series Dropdown -->
        <div>
          <label for="paperSeries" class="block text-sm font-medium text-gray-700">{{ $t('pastPaper.series') }} <span class="text-red-500">*</span></label>
          <select
            id="paperSeries"
            v-model="selectedPaperSeriesId"
            required
            :disabled="!selectedSyllabusId || isLoadingDropdowns.paperSeries"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm disabled:bg-gray-100"
          >
            <option :value="null" disabled>{{ isLoadingDropdowns.paperSeries ? $t('common.loading') : (selectedSyllabusId ? $t('pastPaper.selectSeries') : $t('pastPaper.selectSyllabus')) }}</option>
            <option v-for="series in paperSeries" :key="series.id" :value="series.id">
              {{ series.name }}
            </option>
          </select>
        </div>

        <!-- Paper Code Dropdown -->
        <div>
          <label for="paperCode" class="block text-sm font-medium text-gray-700">{{ $t('pastPaper.code') }} <span class="text-red-500">*</span></label>
          <select
            id="paperCode"
            v-model="selectedPaperCodeId"
            required
            :disabled="!selectedSyllabusId || isLoadingDropdowns.paperCodes"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm disabled:bg-gray-100"
          >
            <option :value="null" disabled>{{ isLoadingDropdowns.paperCodes ? $t('common.loading') : (selectedSyllabusId ? $t('pastPaper.selectCode') : $t('pastPaper.selectSyllabus')) }}</option>
            <option v-for="code in paperCodes" :key="code.id" :value="code.id">
              {{ code.name }}
            </option>
          </select>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label for="year" class="block text-sm font-medium text-gray-700">{{ $t('pastPaper.year') }} <span class="text-red-500">*</span></label>
            <input
              type="number"
              id="year"
              v-model.number="form.year"
              required
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
            />
          </div>
          <div>
            <label for="questionNumber" class="block text-sm font-medium text-gray-700">{{ $t('examPaperManagement.questionCount') }} <span class="text-red-500">*</span></label>
            <input
              type="number"
              id="questionNumber"
              v-model.number="form.questionNumber"
              required
              min="0"
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
            />
          </div>
        </div>

        <div class="flex justify-end space-x-3 pt-4">
          <button
            type="button"
            @click="router.push({ name: 'AdminPastPaperManagement' })"
            class="px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            :disabled="isLoading"
          >
            {{ $t('common.cancel') }}
          </button>
          <button
            type="submit"
            class="px-4 py-2 bg-blue-600 text-white rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
            :disabled="isLoading || isLoadingDropdowns.organisations || isLoadingDropdowns.qualifications || isLoadingDropdowns.syllabuses"
          >
            <span v-if="isLoading">{{ $t('pastPaper.saving') }}</span>
            <span v-else>{{ isEditMode ? $t('pastPaper.edit') : $t('pastPaper.add') }}</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import pastPaperService from '../../services/pastPaperService'
import organisationService from '../../services/organisationService'
import qualificationService from '../../services/qualificationService'
import syllabusService from '../../services/syllabusService'
import paperSeriesService from '../../services/paperSeriesService'
import paperCodeService from '../../services/paperCodeService'
import type { PastPaperCreateRequest, PastPaperUpdateRequest } from '../../models/pastPaper.model'
import type { Organisation } from '../../models/organisation.model'
import type { Qualification } from '../../models/qualification.model'
import type { Syllabus } from '../../models/syllabus.model'
import type { PaperSeries } from '../../models/paperSeries.model'
import type { PaperCode } from '../../models/paperCode.model'

const route = useRoute()
const router = useRouter()

const pastPaperId = computed(() => {
  const id = route.params.id
  return typeof id === 'string' ? parseInt(id, 10) : undefined
})
const isEditMode = computed(() => !!pastPaperId.value)
const initialDataLoaded = ref(false)

const form = ref<PastPaperCreateRequest | PastPaperUpdateRequest>({
  name: '',
  year: new Date().getFullYear(),
  syllabusId: 0,
  paperSeriesId: 0,
  paperCodeId: 0,
  questionNumber: 1
})

const isLoading = ref(false)
const errorMessage = ref<string | null>(null)
const successMessage = ref<string | null>(null)

// Dropdown state
const organisations = ref<Organisation[]>([])
const qualifications = ref<Qualification[]>([])
const syllabuses = ref<Syllabus[]>([])
const paperSeries = ref<PaperSeries[]>([])
const paperCodes = ref<PaperCode[]>([])

const selectedOrganisationId = ref<number | null>(null)
const selectedQualificationId = ref<number | null>(null)
const selectedSyllabusId = ref<number | null>(null)
const selectedPaperSeriesId = ref<number | null>(null)
const selectedPaperCodeId = ref<number | null>(null)

const isLoadingDropdowns = ref({
  organisations: false,
  qualifications: false,
  syllabuses: false,
  paperSeries: false,
  paperCodes: false,
})

// Loading functions
const loadOrganisations = async () => {
  isLoadingDropdowns.value.organisations = true
  errorMessage.value = null
  try {
    const response = await organisationService.getAllOrganisations({})
    if (response.code === 0 && response.data) {
      organisations.value = response.data.list
    } else {
      errorMessage.value = response.message || 'Failed to load organisations.'
    }
  } catch (error: any) {
    errorMessage.value = 'Error loading organisations: ' + error.message
  } finally {
    isLoadingDropdowns.value.organisations = false
  }
}

const loadQualifications = async (organisationId: number | null) => {
  if (!organisationId) {
    qualifications.value = []
    selectedQualificationId.value = null
    return
  }
  isLoadingDropdowns.value.qualifications = true
  errorMessage.value = null
  try {
    const response = await qualificationService.getAllQualifications({ organisationId })
    if (response.code === 0 && response.data) {
      qualifications.value = response.data.list
    } else {
      qualifications.value = []
      errorMessage.value = response.message || 'Failed to load qualifications.'
    }
  } catch (error: any) {
    qualifications.value = []
    errorMessage.value = 'Error loading qualifications: ' + error.message
  } finally {
    isLoadingDropdowns.value.qualifications = false
  }
}

const loadSyllabuses = async (qualificationId: number | null) => {
  if (!qualificationId) {
    syllabuses.value = []
    selectedSyllabusId.value = null
    return
  }
  isLoadingDropdowns.value.syllabuses = true
  errorMessage.value = null
  try {
    const response = await syllabusService.getAllSyllabuses({ qualificationId })
    if (response.code === 0 && response.data) {
      syllabuses.value = response.data.list
    } else {
      syllabuses.value = []
      errorMessage.value = response.message || 'Failed to load syllabuses.'
    }
  } catch (error: any) {
    syllabuses.value = []
    errorMessage.value = 'Error loading syllabuses: ' + error.message
  } finally {
    isLoadingDropdowns.value.syllabuses = false
  }
}

const loadPaperSeries = async (syllabusId: number | null) => {
  if (!syllabusId) {
    paperSeries.value = []
    selectedPaperSeriesId.value = null
    return
  }
  isLoadingDropdowns.value.paperSeries = true
  errorMessage.value = null
  try {
    const response = await paperSeriesService.getAllPaperSeries({ syllabusId })
    if (response.code === 0 && response.data) {
      paperSeries.value = response.data.list
    } else {
      paperSeries.value = []
      errorMessage.value = response.message || 'Failed to load paper series.'
    }
  } catch (error: any) {
    paperSeries.value = []
    errorMessage.value = 'Error loading paper series: ' + error.message
  } finally {
    isLoadingDropdowns.value.paperSeries = false
  }
}

const loadPaperCodes = async (syllabusId: number | null) => {
  if (!syllabusId) {
    paperCodes.value = []
    selectedPaperCodeId.value = null
    return
  }
  isLoadingDropdowns.value.paperCodes = true
  errorMessage.value = null
  try {
    const response = await paperCodeService.getAllPaperCodes({ syllabusId })
    if (response.code === 0 && response.data) {
      paperCodes.value = response.data.list
    } else {
      paperCodes.value = []
      errorMessage.value = response.message || 'Failed to load paper codes.'
    }
  } catch (error: any) {
    paperCodes.value = []
    errorMessage.value = 'Error loading paper codes: ' + error.message
  } finally {
    isLoadingDropdowns.value.paperCodes = false
  }
}

// Watchers
watch(selectedOrganisationId, async (newOrgId) => {
  selectedQualificationId.value = null
  qualifications.value = []
  selectedSyllabusId.value = null
  syllabuses.value = []
  selectedPaperSeriesId.value = null
  paperSeries.value = []
  selectedPaperCodeId.value = null
  paperCodes.value = []
  form.value.syllabusId = 0
  form.value.paperSeriesId = 0
  form.value.paperCodeId = 0
  if (newOrgId) {
    await loadQualifications(newOrgId)
  }
})

watch(selectedQualificationId, async (newQualId) => {
  selectedSyllabusId.value = null
  syllabuses.value = []
  selectedPaperSeriesId.value = null
  paperSeries.value = []
  selectedPaperCodeId.value = null
  paperCodes.value = []
  form.value.syllabusId = 0
  form.value.paperSeriesId = 0
  form.value.paperCodeId = 0
  if (newQualId) {
    await loadSyllabuses(newQualId)
  }
})

watch(selectedSyllabusId, async (newSyllabusId) => {
  selectedPaperSeriesId.value = null
  paperSeries.value = []
  selectedPaperCodeId.value = null
  paperCodes.value = []
  form.value.paperSeriesId = 0
  form.value.paperCodeId = 0
  if (newSyllabusId) {
    form.value.syllabusId = newSyllabusId
    await Promise.all([
      loadPaperSeries(newSyllabusId),
      loadPaperCodes(newSyllabusId)
    ])
  } else {
    form.value.syllabusId = 0
  }
})

watch(selectedPaperSeriesId, (newSeriesId) => {
  if (newSeriesId) {
    const selected = paperSeries.value.find(s => s.id === newSeriesId)
    if (selected) {
      form.value.paperSeriesId = selected.id || 0
    }
  } else {
    form.value.paperSeriesId = 0
  }
})

watch(selectedPaperCodeId, (newCodeId) => {
  if (newCodeId) {
    const selected = paperCodes.value.find(c => c.id === newCodeId)
    if (selected) {
      form.value.paperCodeId = selected.id || 0
    }
  } else {
    form.value.paperCodeId = 0
  }
})

const loadPastPaperDetails = async () => {
  if (isEditMode.value && pastPaperId.value) {
    isLoading.value = true
    errorMessage.value = null
    try {
      const response = await pastPaperService.getPastPaperById(pastPaperId.value)
      if (response.code === 0 && response.data) {
        const paperData = response.data
        form.value = {
          id: paperData.id,
          name: paperData.name,
          year: paperData.year,
          syllabusId: paperData.syllabusId,
          paperSeriesId: paperData.paperSeriesId,
          paperCodeId: 0, // Will need to be updated based on your data structure
          questionNumber: 1 // Will need to be updated based on your data structure
        }
        
        // Load the dropdowns and set their values
        if (paperData.syllabus.qualificationId) {
          const initialQualId = paperData.syllabus.qualificationId

          await loadOrganisations()

          const qualResponse = await qualificationService.getQualificationById(initialQualId)
          if (qualResponse.code === 0 && qualResponse.data) {
            selectedOrganisationId.value = qualResponse.data.organisationId
            await loadQualifications(selectedOrganisationId.value)
            selectedQualificationId.value = initialQualId
            
            await loadSyllabuses(selectedQualificationId.value)
            const targetSyllabus = syllabuses.value.find(s => 
              s.qualificationId === initialQualId
            )
            if (targetSyllabus) {
              selectedSyllabusId.value = targetSyllabus.id
              
              await Promise.all([
                loadPaperSeries(targetSyllabus.id),
                loadPaperCodes(targetSyllabus.id)
              ])

              // Set paper series ID
              if (paperData.paperSeriesId) {
                const seriesId = paperData.paperSeriesId
                if (!isNaN(seriesId)) {
                  selectedPaperSeriesId.value = seriesId
                }
              }

              // Set paper code ID
              if (paperData.paperCodeId) {
                const codeId = paperData.paperCodeId
                if (!isNaN(codeId)) {
                  selectedPaperCodeId.value = codeId
                }
              }

              // Set question number
              if (paperData.questionNumber != null) {
                form.value.questionNumber = paperData.questionNumber
              } else {
                form.value.questionNumber = 1 // Default to 1 if not set
              }
            }
          }
        }
      }
    } catch (error: any) {
      errorMessage.value = 'Error loading past paper details: ' + error.message
    } finally {
      isLoading.value = false
      initialDataLoaded.value = true
    }
  } else {
    await loadOrganisations()
    initialDataLoaded.value = true
  }
}

onMounted(async () => {
  await loadPastPaperDetails()
})

const handleSubmit = async () => {
  isLoading.value = true
  errorMessage.value = null
  successMessage.value = null

  if (!selectedSyllabusId.value || !selectedPaperSeriesId.value || !selectedPaperCodeId.value) {
    errorMessage.value = "Please select all required fields."
    isLoading.value = false
    return
  }

  if (!form.value.name || !form.value.year || form.value.questionNumber == null) {
    errorMessage.value = "Please fill in all required fields."
    isLoading.value = false
    return
  }

  try {
    let response
    if (isEditMode.value && pastPaperId.value) {
      const updatePayload: PastPaperUpdateRequest = {
        ...form.value as PastPaperCreateRequest,
        id: pastPaperId.value
      }
      response = await pastPaperService.updatePastPaper(updatePayload)
    } else {
      response = await pastPaperService.createPastPaper(form.value as PastPaperCreateRequest)
    }

    if (response.code === 0) {
      successMessage.value = `Past paper ${isEditMode.value ? 'updated' : 'created'} successfully!`
      setTimeout(() => {
        router.push({ name: 'AdminPastPaperManagement' })
      }, 1500)
    } else {
      errorMessage.value = response.message || `Failed to ${isEditMode.value ? 'update' : 'create'} past paper.`
    }
  } catch (error: any) {
    errorMessage.value = 'An unexpected error occurred: ' + error.message
  } finally {
    isLoading.value = false
  }
}

// const pageTitle = computed(() => (isEditMode.value ? 'Edit Past Paper' : 'Create New Past Paper'))
</script>
