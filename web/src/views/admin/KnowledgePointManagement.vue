<template>
  <div class="p-6">
    <header class="mb-6">
      <div class="flex justify-between items-start">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ $t('knowledgePoint.title') }}</h1>
          <p class="mt-2 text-sm text-gray-600">{{ $t('knowledgePoint.subtitle') }}</p>
          
          <div class="mt-4 bg-white rounded-lg shadow p-4 border-l-4 border-indigo-500">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <h3 class="text-sm font-medium text-gray-500">{{ $t('knowledgePoint.syllabusName') }}</h3>
                <p class="mt-1 text-lg font-semibold text-gray-900">{{ syllabus?.name || '-' }}</p>
              </div>
              <div>
                <h3 class="text-sm font-medium text-gray-500">{{ $t('knowledgePoint.syllabusCode') }}</h3>
                <p class="mt-1 text-lg font-semibold text-gray-900">{{ syllabus?.code || '-' }}</p>
              </div>
            </div>
          </div>
        </div>
        <router-link 
          :to="`/admin/syllabuses`"
          class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          {{ $t('knowledgePoint.backToSyllabuses') }}
        </router-link>
      </div>
    </header>

    <div class="flex gap-6">
      <!-- Left side: Chapter list (leaf chapters only) -->
      <div class="w-1/3 bg-white rounded-lg shadow">
        <div class="p-4 border-b border-gray-200">
          <h2 class="text-lg font-medium text-gray-900">{{ $t('knowledgePoint.leafChaptersOnly') }}</h2>
          <p class="text-sm text-gray-500 mt-1">{{ $t('knowledgePoint.leafChapterRequired') }}</p>
        </div>
        <div class="p-4">
          <div v-if="loading" class="text-center py-4">
            <div class="text-sm text-gray-500">{{ $t('knowledgePoint.loading') }}</div>
          </div>
          <div v-else-if="!leafChapters.length" class="text-center py-4">
            <div class="text-sm text-gray-500">{{ $t('knowledgePoint.noChapters') }}</div>
          </div>
          <div v-else class="space-y-2">
            <button
              v-for="chapter in leafChapters"
              :key="chapter.id"
              @click="selectChapter(chapter)"
              :class="[
                'w-full text-left px-4 py-3 rounded-md transition-colors',
                selectedChapterId === chapter.id
                  ? 'bg-indigo-100 text-indigo-900 border-2 border-indigo-500'
                  : 'bg-gray-50 hover:bg-gray-100 border-2 border-transparent'
              ]"
            >
              <div class="font-medium">{{ chapter.name }}</div>
              <div class="text-sm text-gray-500 mt-1">
                {{ chapter.knowledgePointCount || 0 }} {{ $t('knowledgePoint.knowledgePoints') }}
              </div>
            </button>
          </div>
        </div>
      </div>

      <!-- Right side: Knowledge points for selected chapter -->
      <div class="flex-1 bg-white rounded-lg shadow">
        <div class="p-4 border-b border-gray-200 flex justify-between items-center">
          <h2 class="text-lg font-medium text-gray-900">
            {{ selectedChapter ? selectedChapter.name : $t('knowledgePoint.selectChapter') }}
          </h2>
          <div v-if="selectedChapter" class="flex gap-2">
            <button
              @click="generateKnowledgePoints"
              :disabled="generating"
              class="inline-flex items-center px-3 py-1 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-purple-600 hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500 disabled:opacity-50"
            >
              <svg v-if="generating" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              {{ generating ? $t('knowledgePoint.generating') : $t('knowledgePoint.generateKnowledgePoints') }}
            </button>
            <button
              @click="showCreateModal"
              class="inline-flex items-center px-3 py-1 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              {{ $t('knowledgePoint.addKnowledgePoint') }}
            </button>
          </div>
        </div>
        <div class="p-4">
          <div v-if="!selectedChapter" class="text-center py-8 text-gray-500">
            {{ $t('knowledgePoint.selectChapter') }}
          </div>
          <div v-else-if="loadingKnowledgePoints" class="text-center py-8 text-gray-500">
            {{ $t('knowledgePoint.loading') }}
          </div>
          <div v-else-if="!knowledgePoints.length" class="text-center py-8">
            <div class="text-gray-500 mb-4">{{ $t('knowledgePoint.noKnowledgePoints') }}</div>
            <button
              @click="generateKnowledgePoints"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-purple-600 hover:bg-purple-700"
            >
              {{ $t('knowledgePoint.generateKnowledgePoints') }}
            </button>
          </div>
          <div v-else class="space-y-4">
            <div
              v-for="kp in knowledgePoints"
              :key="kp.id"
              class="border border-gray-200 rounded-lg p-4 hover:border-indigo-300 transition-colors"
            >
              <div class="flex justify-between items-start">
                <div class="flex-1">
                  <h3 class="text-lg font-medium text-gray-900">{{ kp.name }}</h3>
                  <p class="text-sm text-gray-600 mt-1">{{ kp.description }}</p>
                  <div class="mt-2 flex items-center gap-4 text-sm text-gray-500">
                    <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                      :class="{
                        'bg-green-100 text-green-800': kp.difficulty === 'basic',
                        'bg-yellow-100 text-yellow-800': kp.difficulty === 'medium',
                        'bg-red-100 text-red-800': kp.difficulty === 'hard'
                      }">
                      {{ $t(`knowledgePoint.${kp.difficulty}`) }}
                    </span>
                    <span>{{ kp.estimatedMinutes }} {{ $t('knowledgePoint.estimatedMinutes') }}</span>
                  </div>
                </div>
                <div class="ml-4 flex gap-2">
                  <button
                    @click="editKnowledgePoint(kp)"
                    class="text-indigo-600 hover:text-indigo-900"
                  >
                    {{ $t('common.edit') }}
                  </button>
                  <button
                    @click="deleteKnowledgePointConfirm(kp)"
                    class="text-red-600 hover:text-red-900"
                  >
                    {{ $t('common.delete') }}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Create/Edit Modal -->
    <div v-if="showModal" class="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl max-w-2xl w-full mx-4">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-medium text-gray-900">
            {{ isEditing ? $t('knowledgePoint.editKnowledgePoint') : $t('knowledgePoint.addKnowledgePoint') }}
          </h3>
        </div>
        <form @submit.prevent="saveKnowledgePoint" class="px-6 py-4 space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">{{ $t('knowledgePoint.name') }}</label>
            <input
              type="text"
              v-model="form.name"
              required
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm py-3"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">{{ $t('knowledgePoint.description') }}</label>
            <textarea
              v-model="form.description"
              rows="3"
              required
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
            ></textarea>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">{{ $t('knowledgePoint.difficulty') }}</label>
              <select
                v-model="form.difficulty"
                required
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm py-3"
              >
                <option value="basic">{{ $t('knowledgePoint.basic') }}</option>
                <option value="medium">{{ $t('knowledgePoint.medium') }}</option>
                <option value="hard">{{ $t('knowledgePoint.hard') }}</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">{{ $t('knowledgePoint.estimatedMinutes') }}</label>
              <input
                type="number"
                v-model.number="form.estimatedMinutes"
                required
                min="1"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm py-3"
              />
            </div>
          </div>
          <div class="flex justify-end gap-2 pt-4">
            <button
              type="button"
              @click="showModal = false"
              class="px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
            >
              {{ $t('knowledgePoint.cancel') }}
            </button>
            <button
              type="submit"
              :disabled="saving"
              class="px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50"
            >
              {{ saving ? $t('common.loading') : $t('knowledgePoint.save') }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { push } from 'notivue'
import knowledgePointService from '../../services/knowledgePointService'
import syllabusService from '../../services/syllabusService'
import chapterService from '../../services/chapterService'
import type { KnowledgePoint } from '../../models/knowledgePoint.model'
import type { Syllabus } from '../../models/syllabus.model'

const route = useRoute()
const syllabusId = computed(() => Number(route.params.id))

const syllabus = ref<Syllabus | null>(null)
const loading = ref(true)
const loadingKnowledgePoints = ref(false)
const leafChapters = ref<any[]>([])
const selectedChapterId = ref<number | null>(null)
const selectedChapter = ref<any | null>(null)
const knowledgePoints = ref<KnowledgePoint[]>([])

const showModal = ref(false)
const isEditing = ref(false)
const saving = ref(false)
const generating = ref(false)
const form = ref({
  id: 0,
  name: '',
  description: '',
  difficulty: 'basic' as 'basic' | 'medium' | 'hard',
  estimatedMinutes: 30
})

onMounted(async () => {
  await loadSyllabus()
  await loadChapters()
})

async function loadSyllabus() {
  try {
    const response = await syllabusService.getById(syllabusId.value)
    if (response.code === 200) {
      syllabus.value = response.data
    }
  } catch (error) {
    console.error('Failed to load syllabus:', error)
    push.error('Failed to load syllabus')
  }
}

async function loadChapters() {
  try {
    loading.value = true
    const response = await chapterService.getChapterTree(syllabusId.value)
    if (response.code === 200) {
      // Filter to get only leaf chapters (chapters without children)
      leafChapters.value = filterLeafChapters(response.data.chapters || [])
    }
  } catch (error) {
    console.error('Failed to load chapters:', error)
    push.error('Failed to load chapters')
  } finally {
    loading.value = false
  }
}

function filterLeafChapters(chapters: any[]): any[] {
  const result: any[] = []
  
  function traverse(chaps: any[]) {
    for (const chapter of chaps) {
      if (!chapter.children || chapter.children.length === 0) {
        result.push(chapter)
      } else {
        traverse(chapter.children)
      }
    }
  }
  
  traverse(chapters)
  return result
}

async function selectChapter(chapter: any) {
  selectedChapterId.value = chapter.id
  selectedChapter.value = chapter
  await loadKnowledgePoints()
}

async function loadKnowledgePoints() {
  if (!selectedChapterId.value) return
  
  try {
    loadingKnowledgePoints.value = true
    const response = await knowledgePointService.getByChapter(selectedChapterId.value)
    if (response.code === 200) {
      knowledgePoints.value = response.data.list || []
    }
  } catch (error) {
    console.error('Failed to load knowledge points:', error)
    push.error('Failed to load knowledge points')
  } finally {
    loadingKnowledgePoints.value = false
  }
}

async function generateKnowledgePoints() {
  if (!selectedChapterId.value) return
  
  try {
    generating.value = true
    const response = await knowledgePointService.generateKeypoints({
      chapterId: selectedChapterId.value,
      mode: 'auto'
    })
    if (response.code === 200) {
      push.success($t('knowledgePoint.generateSuccess'))
      await loadKnowledgePoints()
    } else {
      throw new Error(response.msg)
    }
  } catch (error: any) {
    console.error('Failed to generate knowledge points:', error)
    push.error(error.response?.data?.msg || $t('knowledgePoint.generateError'))
  } finally {
    generating.value = false
  }
}

function showCreateModal() {
  isEditing.value = false
  form.value = {
    id: 0,
    name: '',
    description: '',
    difficulty: 'basic',
    estimatedMinutes: 30
  }
  showModal.value = true
}

function editKnowledgePoint(kp: KnowledgePoint) {
  isEditing.value = true
  form.value = {
    id: kp.id,
    name: kp.name,
    description: kp.description,
    difficulty: kp.difficulty,
    estimatedMinutes: kp.estimatedMinutes
  }
  showModal.value = true
}

async function saveKnowledgePoint() {
  if (!selectedChapterId.value) return
  
  try {
    saving.value = true
    let response
    
    if (isEditing.value) {
      response = await knowledgePointService.update({
        id: form.value.id,
        name: form.value.name,
        description: form.value.description,
        difficulty: form.value.difficulty,
        estimatedMinutes: form.value.estimatedMinutes
      })
      if (response.code === 200) {
        push.success($t('knowledgePoint.updateSuccess'))
      }
    } else {
      response = await knowledgePointService.create({
        chapterId: selectedChapterId.value,
        name: form.value.name,
        description: form.value.description,
        difficulty: form.value.difficulty,
        estimatedMinutes: form.value.estimatedMinutes
      })
      if (response.code === 200) {
        push.success($t('knowledgePoint.createSuccess'))
      }
    }
    
    showModal.value = false
    await loadKnowledgePoints()
  } catch (error: any) {
    console.error('Failed to save knowledge point:', error)
    push.error(error.response?.data?.msg || 'Failed to save knowledge point')
  } finally {
    saving.value = false
  }
}

async function deleteKnowledgePointConfirm(kp: KnowledgePoint) {
  if (!confirm($t('knowledgePoint.confirmDelete'))) return
  
  try {
    const response = await knowledgePointService.delete(kp.id)
    if (response.code === 200) {
      push.success($t('knowledgePoint.deleteSuccess'))
      await loadKnowledgePoints()
    }
  } catch (error: any) {
    console.error('Failed to delete knowledge point:', error)
    push.error(error.response?.data?.msg || 'Failed to delete knowledge point')
  }
}

function $t(key: string): string {
  // Placeholder for i18n translation
  return key
}
</script>
