<template>
  <div class="p-6">
    <header class="mb-6">
      <div class="flex justify-between items-start">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ t('examNode.title') }}</h1>
          <p class="mt-2 text-sm text-gray-600">{{ t('examNode.subtitle') }}</p>
          <div v-if="syllabus" class="mt-3 text-sm text-gray-700 font-medium">{{ syllabus.name }}</div>
        </div>
        <div class="flex gap-2">
          <button
            @click="openCreateModal"
            class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            {{ t('examNode.createNode') }}
          </button>
          <router-link
            to="/admin/syllabuses"
            class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
          >
            {{ t('examNode.backToSyllabuses') }}
          </router-link>
        </div>
      </div>
    </header>

    <!-- Table -->
    <div class="bg-white shadow rounded-lg overflow-hidden">
      <div v-if="loading" class="text-center py-12 text-gray-500">{{ t('examNode.loading') }}</div>
      <div v-else-if="!nodes.length" class="text-center py-12 text-gray-500">{{ t('examNode.noNodes') }}</div>
      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('examNode.name') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('examNode.description') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('examNode.sortOrder') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('examNode.chapterCount') }}</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <template v-for="node in nodes" :key="node.id">
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ node.name }}</td>
              <td class="px-6 py-4 text-sm text-gray-500 max-w-xs truncate">{{ node.description || '-' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ node.sortOrder }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ node.chapters?.length ?? 0 }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2">
                <button @click="toggleExpand(node.id)" class="text-gray-600 hover:text-gray-900">
                  {{ expandedNodeId === node.id ? '▲' : '▼' }}
                </button>
                <button @click="openAddChapterModal(node)" class="text-green-600 hover:text-green-900">{{ t('examNode.addChapter') }}</button>
                <button @click="openEditModal(node)" class="text-indigo-600 hover:text-indigo-900">{{ t('common.edit') }}</button>
                <button @click="confirmDelete(node)" class="text-red-600 hover:text-red-900">{{ t('common.delete') }}</button>
              </td>
            </tr>
            <tr v-if="expandedNodeId === node.id">
              <td colspan="5" class="px-6 py-4 bg-gray-50">
                <div class="text-sm font-medium text-gray-700 mb-2">{{ t('examNode.chapters') }}</div>
                <div v-if="!node.chapters?.length" class="text-sm text-gray-400 italic mb-3">—</div>
                <div v-else class="flex flex-wrap gap-2 mb-3">
                  <span v-for="ch in node.chapters" :key="ch.id" class="inline-flex items-center gap-1 px-2 py-1 bg-white border border-gray-200 rounded text-xs text-gray-700">
                    {{ ch.name }}
                    <button @click="doRemoveChapter(node, ch.id)" class="text-red-400 hover:text-red-600 ml-1">×</button>
                  </span>
                </div>
                <div class="text-sm font-medium text-gray-700 mb-2">{{ t('examNode.paperCodes') }}</div>
                <div v-if="!node.paperCodes?.length" class="text-sm text-gray-400 italic">—</div>
                <div v-else class="flex flex-wrap gap-2">
                  <span v-for="pc in node.paperCodes" :key="pc.id" class="inline-flex items-center px-2 py-1 bg-blue-50 border border-blue-200 rounded text-xs text-blue-700">
                    {{ pc.code || pc.name }}
                  </span>
                </div>
              </td>
            </tr>
          </template>
        </tbody>
      </table>
      </div>
    </div>

    <!-- Create / Edit Modal -->
    <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-md p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">
          {{ editingNode ? t('examNode.editNode') : t('examNode.createNode') }}
        </h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('examNode.name') }}</label>
            <input v-model="form.name" type="text" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('examNode.description') }}</label>
            <textarea v-model="form.description" rows="3" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('examNode.sortOrder') }}</label>
            <input v-model.number="form.sortOrder" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
          </div>
        </div>
        <div class="mt-6 flex justify-end space-x-3">
          <button @click="closeModal" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('examNode.cancel') }}</button>
          <button @click="submitForm" :disabled="saving" class="px-4 py-2 text-sm font-medium text-white bg-indigo-600 rounded-md hover:bg-indigo-700 disabled:opacity-50">
            {{ saving ? t('examNode.loading') : (editingNode ? t('examNode.save') : t('examNode.create')) }}
          </button>
        </div>
      </div>
    </div>

    <!-- Add Chapter Modal -->
    <div v-if="showAddChapterModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-lg p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">{{ t('examNode.addChapter') }}</h2>
        <p class="text-sm text-gray-500 mb-3">{{ t('examNode.selectChapter') }}</p>
        <div class="max-h-80 overflow-y-auto border border-gray-200 rounded-md">
          <div v-if="!chapters.length" class="p-4 text-sm text-gray-400 text-center">{{ t('examNode.loading') }}</div>
          <div v-for="ch in chapters" :key="ch.id"
            @click="selectedChapterId = ch.id"
            :class="['px-4 py-2 cursor-pointer text-sm hover:bg-indigo-50', selectedChapterId === ch.id ? 'bg-indigo-100 font-medium text-indigo-900' : 'text-gray-700']">
            {{ ch.name }}
          </div>
        </div>
        <div class="mt-6 flex justify-end space-x-3">
          <button @click="showAddChapterModal = false" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('examNode.cancel') }}</button>
          <button @click="doAddChapter" :disabled="!selectedChapterId || saving" class="px-4 py-2 text-sm font-medium text-white bg-green-600 rounded-md hover:bg-green-700 disabled:opacity-50">
            {{ saving ? t('examNode.loading') : t('examNode.addChapter') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Delete Confirm Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-sm p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-2">{{ t('examNode.deleteNode') }}</h2>
        <p class="text-sm text-gray-600 mb-6">{{ t('examNode.confirmDelete') }}</p>
        <div class="flex justify-end space-x-3">
          <button @click="showDeleteModal = false" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('examNode.cancel') }}</button>
          <button @click="doDelete" :disabled="saving" class="px-4 py-2 text-sm font-medium text-white bg-red-600 rounded-md hover:bg-red-700 disabled:opacity-50">
            {{ saving ? t('examNode.loading') : t('examNode.confirm') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import examNodeService from '../../services/examNodeService'
import syllabusService from '../../services/syllabusService'
import chapterService from '../../services/chapterService'
import type { SyllabusExamNode } from '../../models/examNode.model'

const { t } = useI18n()
const route = useRoute()
const syllabusId = Number(route.params.syllabusId)

const nodes = ref<SyllabusExamNode[]>([])
const syllabus = ref<any>(null)
const chapters = ref<any[]>([])
const loading = ref(false)
const saving = ref(false)

const showModal = ref(false)
const showAddChapterModal = ref(false)
const showDeleteModal = ref(false)

const editingNode = ref<SyllabusExamNode | null>(null)
const addingChapterNode = ref<SyllabusExamNode | null>(null)
const deletingNode = ref<SyllabusExamNode | null>(null)
const expandedNodeId = ref<number | null>(null)
const selectedChapterId = ref<number | null>(null)

const form = ref({ name: '', description: '', sortOrder: 0 })

async function loadNodes() {
  loading.value = true
  try {
    const res = await examNodeService.list(syllabusId)
    if (res.code === 0) nodes.value = res.data.list ?? []
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }
}

async function loadSyllabus() {
  try {
    const res = await syllabusService.getSyllabusById(syllabusId)
    if (res.code === 0) syllabus.value = res.data
  } catch (e) {
    console.error(e)
  }
}

async function loadChapters() {
  try {
    const res = await chapterService.getChapterList({ syllabusId, pageSize: 200, pageIndex: 1 })
    if (res.code === 0) chapters.value = res.data.list ?? res.data
  } catch (e) {
    console.error(e)
  }
}

function toggleExpand(id: number) {
  expandedNodeId.value = expandedNodeId.value === id ? null : id
}

function openCreateModal() {
  editingNode.value = null
  form.value = { name: '', description: '', sortOrder: 0 }
  showModal.value = true
}

function openEditModal(node: SyllabusExamNode) {
  editingNode.value = node
  form.value = { name: node.name, description: node.description || '', sortOrder: node.sortOrder }
  showModal.value = true
}

function closeModal() {
  showModal.value = false
}

async function submitForm() {
  saving.value = true
  try {
    if (editingNode.value) {
      const res = await examNodeService.update({ id: editingNode.value.id, name: form.value.name, description: form.value.description, sortOrder: form.value.sortOrder })
      if (res.code === 0) { await loadNodes(); closeModal() }
    } else {
      const res = await examNodeService.create({ syllabusId, name: form.value.name, description: form.value.description, sortOrder: form.value.sortOrder })
      if (res.code === 0) { await loadNodes(); closeModal() }
    }
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

function openAddChapterModal(node: SyllabusExamNode) {
  addingChapterNode.value = node
  selectedChapterId.value = null
  showAddChapterModal.value = true
}

async function doAddChapter() {
  if (!addingChapterNode.value || !selectedChapterId.value) return
  saving.value = true
  try {
    const res = await examNodeService.addChapter(addingChapterNode.value.id, selectedChapterId.value)
    if (res.code === 0) { await loadNodes(); showAddChapterModal.value = false }
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

async function doRemoveChapter(node: SyllabusExamNode, chapterId: number) {
  try {
    const res = await examNodeService.removeChapter(node.id, chapterId)
    if (res.code === 0) await loadNodes()
  } catch (e) {
    console.error(e)
  }
}

function confirmDelete(node: SyllabusExamNode) {
  deletingNode.value = node
  showDeleteModal.value = true
}

async function doDelete() {
  if (!deletingNode.value) return
  saving.value = true
  try {
    const res = await examNodeService.delete(deletingNode.value.id)
    if (res.code === 0) { await loadNodes(); showDeleteModal.value = false }
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

onMounted(() => {
  loadNodes()
  loadSyllabus()
  loadChapters()
})
</script>
