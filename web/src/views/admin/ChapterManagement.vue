<template>
  <div class="p-6">
    <header class="mb-6">
      <div class="flex flex-col sm:flex-row sm:justify-between sm:items-start gap-4">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ $t('chapterManagement.title') }}</h1>
          <div class="mt-4 bg-white rounded-lg shadow p-4 border-l-4 border-indigo-500">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <h3 class="text-sm font-medium text-gray-500">{{ $t('chapterManagement.syllabusName') }}</h3>
                <p class="mt-1 text-lg font-semibold text-gray-900">{{ syllabus?.name || '-' }}</p>
              </div>
              <div>
                <h3 class="text-sm font-medium text-gray-500">{{ $t('chapterManagement.syllabusCode') }}</h3>
                <p class="mt-1 text-lg font-semibold text-gray-900">{{ syllabus?.code || '-' }}</p>
              </div>
              <div>
                <h3 class="text-sm font-medium text-gray-500">{{ $t('chapterManagement.organisation') }}</h3>
                <p class="mt-1 text-lg font-semibold text-gray-900">
                  {{ syllabus?.qualification?.organisation?.name || '-' }}
                </p>
              </div>
              <div>
                <h3 class="text-sm font-medium text-gray-500">{{ $t('chapterManagement.qualification') }}</h3>
                <p class="mt-1 text-lg font-semibold text-gray-900">
                  {{ syllabus?.qualification?.name || '-' }}
                </p>
              </div>
            </div>
          </div>
        </div>
        <router-link 
          :to="`/admin/syllabuses`"
          class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          {{ $t('chapterManagement.backToSyllabuses') }}
        </router-link>
      </div>
    </header>

    <div class="flex flex-col md:flex-row gap-6">
      <!-- Left side: Tree view -->
      <div class="w-full md:w-1/3 bg-white rounded-lg shadow">
        <div class="p-4 border-b border-gray-200 flex justify-between items-center">
          <h2 class="text-lg font-medium text-gray-900">{{ $t('chapterManagement.chapters') }}</h2>
          <button
            @click="showCreateModal"
            class="inline-flex items-center px-3 py-1 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            {{ $t('chapterManagement.addChapter') }}
          </button>
        </div>
        <div class="p-4">
          <div v-if="loading" class="text-center py-4">
            <div class="text-sm text-gray-500">{{ $t('chapterManagement.loading') }}</div>
          </div>
          <div v-else-if="!chapters.length" class="text-center py-4">
            <div class="text-sm text-gray-500">{{ $t('chapterManagement.noChapters') }}</div>
            <button
              @click="showCreateModal"
              class="mt-2 inline-flex items-center px-3 py-1 border border-transparent text-sm font-medium rounded-md text-indigo-700 bg-indigo-100 hover:bg-indigo-200"
            >
              {{ $t('chapterManagement.createFirstChapter') }}
            </button>
          </div>
          <div v-else class="space-y-2">
            <ChapterTree
              :chapters="chapters"
              :selected-id="selectedChapterId"
              :expanded-chapters="expandedChapters"
              @select="selectChapter"
              @toggle-expand="toggleExpand"
            />
          </div>
        </div>
      </div>

      <!-- Right side: Details/Editor -->
      <div class="flex-1 bg-white rounded-lg shadow">
        <div class="p-4 border-b border-gray-200">
          <h2 class="text-lg font-medium text-gray-900">
            {{ selectedChapter ? $t('chapterManagement.detailsTitle') : $t('chapterManagement.selectTitle') }}
          </h2>
        </div>
        <div class="p-4">
          <div v-if="!selectedChapter" class="text-center py-8 text-gray-500">
            {{ $t('chapterManagement.selectTip') }}
          </div>
          <form v-else @submit.prevent="saveChapter" class="space-y-6">
            <div>
              <label for="chapterName" class="block text-sm font-medium text-gray-700">{{ $t('chapterManagement.chapterName') }}</label>
              <input
                type="text"
                id="chapterName"
                v-model="form.name"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm py-3"
                required
              />
            </div>
            <div class="flex justify-between">
              <button
                type="button"
                @click="deleteChapter"
                class="inline-flex items-center px-4 py-2 border border-red-300 text-sm font-medium rounded-md text-red-700 bg-white hover:bg-red-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
              >
                {{ $t('chapterManagement.deleteChapter') }}
              </button>
              <div class="space-x-3">
                <button
                  type="button"
                  @click="cancelEdit"
                  class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                >
                  {{ $t('common.cancel') }}
                </button>
                <button
                  type="submit"
                  class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                >
                  {{ $t('chapterManagement.saveChanges') }}
                </button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Create Chapter Modal -->
    <div v-if="showModal" class="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-md">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">{{ $t('chapterManagement.createModalTitle') }}</h3>
          <form @submit.prevent="createChapter" class="mt-4 space-y-4">
            <div>
              <label for="newChapterName" class="block text-sm font-medium text-gray-700">{{ $t('chapterManagement.chapterName') }}</label>
              <input
                type="text"
                id="newChapterName"
                v-model="newChapterForm.name"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm py-3"
                required
              />
            </div>
            <div>
              <label for="parentChapter" class="block text-sm font-medium text-gray-700">{{ $t('chapterManagement.parentChapterOptional') }}</label>
              <select
                id="parentChapter"
                v-model="newChapterForm.parentId"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm py-3"
              >
                <option value="">{{ $t('chapterManagement.noParent') }}</option>
                <option v-for="chapter in flatChapters" :key="chapter.id" :value="chapter.id">
                  {{ chapter.name }}
                </option>
              </select>
            </div>
            <div class="mt-5 sm:mt-6 flex justify-end space-x-3">
              <button
                type="button"
                @click="hideModal"
                class="inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:text-sm"
              >
                {{ $t('common.cancel') }}
              </button>
              <button
                type="submit"
                class="inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-indigo-600 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:text-sm"
              >
                {{ $t('chapterManagement.create') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import ChapterTree from '../../components/admin/ChapterTree.vue';
import syllabusService from '../../services/syllabusService';
import chapterService from '../../services/chapterService';
import type { 
  Chapter,
  ChapterCreateRequest,
  ChapterUpdateRequest 
} from '../../models/chapter.model';
import type { Syllabus } from '../../models/syllabus.model';

const route = useRoute();
const syllabusId = computed(() => Number(route.params.id));

const syllabus = ref<Syllabus | null>(null);
const chapters = ref<Chapter[]>([]);
const loading = ref(true);
const selectedChapterId = ref<number | null>(null);
const showModal = ref(false);
const expandedChapters = ref<number[]>([]);

const form = ref<ChapterUpdateRequest>({
  id: 0,
  name: '',
});

const newChapterForm = ref<ChapterCreateRequest>({
  name: '',
  syllabusId: syllabusId.value,
  parentId: 0,
});

// Helper function to find chapter by ID in nested structure
const findChapterById = (chapters: Chapter[], id: number | null): Chapter | undefined => {
  if (!id) return undefined;
  
  for (const chapter of chapters) {
    if (chapter.id === id) return chapter;
    if (chapter.children?.length) {
      const found = findChapterById(chapter.children, id);
      if (found) return found;
    }
  }
  return undefined;
};

// Computed
const selectedChapter = computed(() => {
  return findChapterById(chapters.value, selectedChapterId.value);
});

const flatChapters = computed(() => {
  const flat: Chapter[] = [];
  const flatten = (chapters: Chapter[]) => {
    chapters.forEach(chapter => {
      flat.push(chapter);
      if (chapter.children?.length) {
        flatten(chapter.children);
      }
    });
  };
  flatten(chapters.value);
  return flat;
});

// Methods
const loadSyllabus = async () => {
  try {
    const response = await syllabusService.getSyllabusById(syllabusId.value);
    syllabus.value = response.data;
  } catch (error) {
    console.error('Failed to load syllabus:', error);
  }
};

const loadChapters = async () => {
  loading.value = true;
  try {
    const response = await chapterService.getChapterTree(syllabusId.value);
    chapters.value = response.data;
  } catch (error) {
    console.error('Failed to load chapters:', error);
  } finally {
    loading.value = false;
  }
};

const toggleExpand = (id: number) => {
  const index = expandedChapters.value.indexOf(id);
  if (index === -1) {
    expandedChapters.value.push(id);
    
    // Find chapter to ensure parent nodes are expanded
    const chapter = findChapterById(chapters.value, id);
    if (chapter) {
      let parentId = chapter.parentId;
      while (parentId > 0) {
        if (!expandedChapters.value.includes(parentId)) {
          expandedChapters.value.push(parentId);
        }
        const parent = findChapterById(chapters.value, parentId);
        if (!parent) break;
        parentId = parent.parentId;
      }
    }
  } else {
    expandedChapters.value.splice(index, 1);
  }
};

const selectChapter = (id: number) => {
  selectedChapterId.value = id;
  // Find chapter to ensure it exists in the tree
  const chapter = findChapterById(chapters.value, id);
  if (chapter) {
    form.value = {
      id: chapter.id,
      name: chapter.name,
    };
    
    // Ensure all parent chapters are expanded
    let parentId = chapter.parentId;
    while (parentId > 0) {
      if (!expandedChapters.value.includes(parentId)) {
        expandedChapters.value.push(parentId);
      }
      const parent = findChapterById(chapters.value, parentId);
      if (!parent) break;
      parentId = parent.parentId;
    }
  }
};

const showCreateModal = () => {
  newChapterForm.value = {
    name: '',
    syllabusId: syllabusId.value,
    parentId: 0,
  };
  showModal.value = true;
};

const hideModal = () => {
  showModal.value = false;
};

const createChapter = async () => {
  try {
    await chapterService.createChapter({
      name: newChapterForm.value.name,
      syllabusId: syllabusId.value,
      parentId: newChapterForm.value.parentId ? Number(newChapterForm.value.parentId) : 0,
    });
    hideModal();
    loadChapters();
  } catch (error) {
    console.error('Failed to create chapter:', error);
  }
};

const saveChapter = async () => {
  if (!selectedChapter.value) return;
  
  try {
    await chapterService.updateChapter({
      id: selectedChapter.value.id,
      name: form.value.name,
    });
    loadChapters();
  } catch (error) {
    console.error('Failed to update chapter:', error);
  }
};

const deleteChapter = async () => {
  if (!selectedChapter.value || !confirm('Are you sure you want to delete this chapter?')) {
    return;
  }

  try {
    await chapterService.deleteChapter(selectedChapter.value.id);
    selectedChapterId.value = null;
    loadChapters();
  } catch (error) {
    console.error('Failed to delete chapter:', error);
  }
};

const cancelEdit = () => {
  selectedChapterId.value = null;
  form.value = {
    id: 0,
    name: '',
  };
};

// Lifecycle
onMounted(() => {
  loadSyllabus();
  loadChapters();
});
</script>
