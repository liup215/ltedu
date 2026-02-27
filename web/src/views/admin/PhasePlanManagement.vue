<template>
  <div class="p-6">
    <header class="mb-6">
      <div class="flex justify-between items-start">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ t('phasePlan.title') }}</h1>
          <p class="mt-2 text-sm text-gray-600">{{ t('phasePlan.subtitle') }}</p>
          <div v-if="parentPlan" class="mt-3 text-sm text-gray-700 font-medium">
            Plan #{{ parentPlan.id }} — {{ t(`learningPlan.type${parentPlan.planType.charAt(0).toUpperCase() + parentPlan.planType.slice(1)}`) }}
          </div>
        </div>
        <div class="flex gap-2">
          <button
            @click="openCreateModal"
            class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none"
          >
            {{ t('phasePlan.createPhasePlan') }}
          </button>
          <button
            @click="goBack"
            class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
          >
            {{ t('phasePlan.backToPlans') }}
          </button>
        </div>
      </div>
    </header>

    <!-- Table -->
    <div class="bg-white shadow rounded-lg overflow-hidden">
      <div v-if="loading" class="text-center py-12 text-gray-500">{{ t('phasePlan.loading') }}</div>
      <div v-else-if="!phasePlans.length" class="text-center py-12 text-gray-500">{{ t('phasePlan.noPlans') }}</div>
      <table v-else class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('phasePlan.phaseTitle') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('phasePlan.examNode') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('phasePlan.startDate') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('phasePlan.endDate') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('phasePlan.sortOrder') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('phasePlan.chapters') }}</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="pp in phasePlans" :key="pp.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ pp.title }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ pp.examNode?.name || pp.examNodeId }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ pp.startDate?.slice(0, 10) || '—' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ pp.endDate?.slice(0, 10) || '—' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ pp.sortOrder }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              <div class="flex flex-wrap gap-1">
                <span v-for="ch in pp.chapters" :key="ch.id" class="inline-flex items-center gap-1 px-1.5 py-0.5 bg-gray-100 rounded text-xs">
                  {{ ch.name }}
                  <button @click="doRemoveChapter(pp, ch.id)" class="text-red-400 hover:text-red-600">×</button>
                </span>
                <button @click="openAddChapterModal(pp)" class="text-xs text-green-600 hover:text-green-800 font-medium">+ {{ t('phasePlan.addChapter') }}</button>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2">
              <button @click="openEditModal(pp)" class="text-indigo-600 hover:text-indigo-900">{{ t('common.edit') }}</button>
              <button @click="confirmDelete(pp)" class="text-red-600 hover:text-red-900">{{ t('common.delete') }}</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Create / Edit Modal -->
    <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-lg p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">
          {{ editingPP ? t('phasePlan.editPhasePlan') : t('phasePlan.createPhasePlan') }}
        </h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('phasePlan.phaseTitle') }}</label>
            <input v-model="form.title" type="text" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
          </div>
          <div v-if="!editingPP">
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('phasePlan.examNode') }}</label>
            <select v-model.number="form.examNodeId" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500">
              <option value="">{{ t('phasePlan.selectExamNode') }}</option>
              <option v-for="n in examNodes" :key="n.id" :value="n.id">{{ n.name }}</option>
            </select>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('phasePlan.startDate') }}</label>
              <input v-model="form.startDate" type="date" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('phasePlan.endDate') }}</label>
              <input v-model="form.endDate" type="date" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('phasePlan.sortOrder') }}</label>
            <input v-model.number="form.sortOrder" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
          </div>
        </div>
        <div class="mt-6 flex justify-end space-x-3">
          <button @click="closeModal" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('phasePlan.cancel') }}</button>
          <button @click="submitForm" :disabled="saving" class="px-4 py-2 text-sm font-medium text-white bg-indigo-600 rounded-md hover:bg-indigo-700 disabled:opacity-50">
            {{ saving ? t('phasePlan.loading') : (editingPP ? t('phasePlan.save') : t('phasePlan.create')) }}
          </button>
        </div>
      </div>
    </div>

    <!-- Add Chapter Modal -->
    <div v-if="showAddChapterModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-lg p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">{{ t('phasePlan.addChapter') }}</h2>
        <p class="text-sm text-gray-500 mb-3">{{ t('phasePlan.selectChapter') }}</p>
        <div class="max-h-80 overflow-y-auto border border-gray-200 rounded-md">
          <div v-if="!chapters.length" class="p-4 text-sm text-gray-400 text-center">{{ t('phasePlan.loading') }}</div>
          <div v-for="ch in chapters" :key="ch.id"
            @click="selectedChapterId = ch.id"
            :class="['px-4 py-2 cursor-pointer text-sm hover:bg-indigo-50', selectedChapterId === ch.id ? 'bg-indigo-100 font-medium text-indigo-900' : 'text-gray-700']">
            {{ ch.name }}
          </div>
        </div>
        <div class="mt-6 flex justify-end space-x-3">
          <button @click="showAddChapterModal = false" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('phasePlan.cancel') }}</button>
          <button @click="doAddChapter" :disabled="!selectedChapterId || saving" class="px-4 py-2 text-sm font-medium text-white bg-green-600 rounded-md hover:bg-green-700 disabled:opacity-50">
            {{ saving ? t('phasePlan.loading') : t('phasePlan.addChapter') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Delete Confirm Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-sm p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-2">{{ t('phasePlan.deletePhasePlan') }}</h2>
        <p class="text-sm text-gray-600 mb-6">{{ t('phasePlan.confirmDelete') }}</p>
        <div class="flex justify-end space-x-3">
          <button @click="showDeleteModal = false" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('phasePlan.cancel') }}</button>
          <button @click="doDelete" :disabled="saving" class="px-4 py-2 text-sm font-medium text-white bg-red-600 rounded-md hover:bg-red-700 disabled:opacity-50">
            {{ saving ? t('phasePlan.loading') : t('phasePlan.confirm') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import learningPlanService from '../../services/learningPlanService'
import examNodeService from '../../services/examNodeService'
import chapterService from '../../services/chapterService'
import type { LearningPhasePlan } from '../../models/learningPlan.model'

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const planId = Number(route.params.planId)

const phasePlans = ref<LearningPhasePlan[]>([])
const parentPlan = ref<any>(null)
const examNodes = ref<any[]>([])
const chapters = ref<any[]>([])
const loading = ref(false)
const saving = ref(false)

const showModal = ref(false)
const showAddChapterModal = ref(false)
const showDeleteModal = ref(false)

const editingPP = ref<LearningPhasePlan | null>(null)
const addingChapterPP = ref<LearningPhasePlan | null>(null)
const deletingPP = ref<LearningPhasePlan | null>(null)
const selectedChapterId = ref<number | null>(null)

const form = ref({ title: '', examNodeId: 0, startDate: '', endDate: '', sortOrder: 0 })

async function loadPhasePlans() {
  loading.value = true
  try {
    const res = await learningPlanService.listPhasePlans(planId)
    if (res.code === 0) phasePlans.value = res.data.list ?? []
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }
}

async function loadParentPlan() {
  try {
    const res = await learningPlanService.getById(planId)
    if (res.code === 0) {
      parentPlan.value = res.data
      // Load exam nodes for this plan's class syllabus
      // We need syllabusId — try loading it from class info via the plan's class
      if (res.data.class?.syllabusId) {
        loadExamNodes(res.data.class.syllabusId)
        loadChapters(res.data.class.syllabusId)
      }
    }
  } catch (e) {
    console.error(e)
  }
}

async function loadExamNodes(syllabusId: number) {
  try {
    const res = await examNodeService.list(syllabusId)
    if (res.code === 0) examNodes.value = res.data.list ?? []
  } catch (e) {
    console.error(e)
  }
}

async function loadChapters(syllabusId: number) {
  try {
    const res = await chapterService.getChapterList({ syllabusId, pageSize: 200, pageIndex: 1 })
    if (res.code === 0) chapters.value = res.data.list ?? res.data
  } catch (e) {
    console.error(e)
  }
}

function goBack() {
  router.back()
}

function openCreateModal() {
  editingPP.value = null
  form.value = { title: '', examNodeId: 0, startDate: '', endDate: '', sortOrder: 0 }
  showModal.value = true
}

function openEditModal(pp: LearningPhasePlan) {
  editingPP.value = pp
  form.value = { title: pp.title, examNodeId: pp.examNodeId, startDate: pp.startDate || '', endDate: pp.endDate || '', sortOrder: pp.sortOrder }
  showModal.value = true
}

function closeModal() {
  showModal.value = false
}

async function submitForm() {
  saving.value = true
  try {
    if (editingPP.value) {
      const res = await learningPlanService.updatePhasePlan({ id: editingPP.value.id, title: form.value.title, startDate: form.value.startDate || undefined, endDate: form.value.endDate || undefined, sortOrder: form.value.sortOrder })
      if (res.code === 0) { await loadPhasePlans(); closeModal() }
    } else {
      const res = await learningPlanService.createPhasePlan({ planId, examNodeId: form.value.examNodeId, title: form.value.title, startDate: form.value.startDate || undefined, endDate: form.value.endDate || undefined, sortOrder: form.value.sortOrder })
      if (res.code === 0) { await loadPhasePlans(); closeModal() }
    }
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

function openAddChapterModal(pp: LearningPhasePlan) {
  addingChapterPP.value = pp
  selectedChapterId.value = null
  showAddChapterModal.value = true
}

async function doAddChapter() {
  if (!addingChapterPP.value || !selectedChapterId.value) return
  saving.value = true
  try {
    const res = await learningPlanService.addChapterToPhasePlan(addingChapterPP.value.id, selectedChapterId.value)
    if (res.code === 0) { await loadPhasePlans(); showAddChapterModal.value = false }
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

async function doRemoveChapter(pp: LearningPhasePlan, chapterId: number) {
  try {
    const res = await learningPlanService.removeChapterFromPhasePlan(pp.id, chapterId)
    if (res.code === 0) await loadPhasePlans()
  } catch (e) {
    console.error(e)
  }
}

function confirmDelete(pp: LearningPhasePlan) {
  deletingPP.value = pp
  showDeleteModal.value = true
}

async function doDelete() {
  if (!deletingPP.value) return
  saving.value = true
  try {
    const res = await learningPlanService.deletePhasePlan(deletingPP.value.id)
    if (res.code === 0) { await loadPhasePlans(); showDeleteModal.value = false }
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

onMounted(() => {
  loadPhasePlans()
  loadParentPlan()
})
</script>
