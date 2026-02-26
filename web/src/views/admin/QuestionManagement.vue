<template>
  <div class="p-6">
    <!-- Header Section -->
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('question.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('question.noData') }}</p>
    </header>

    <!-- Filters and Actions -->
    <div class="mb-6 flex flex-col space-y-4">
      <!-- Search and Filters Row -->
<div class="flex flex-wrap gap-4">
  <input 
    type="text" 
    v-model="searchQuery"
    :placeholder="$t('question.searchByStem')" 
    class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_260px] min-w-[200px] max-w-lg"
  />
  <input 
    type="text" 
    v-model="paperNameQuery"
    :placeholder="$t('question.searchByPaperName')"
    class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_260px] min-w-[200px] max-w-lg"
  />

  <select 
    v-model="selectedOrganisationId"
    class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_220px] min-w-[180px] max-w-md"
  >
    <option :value="null">{{ $t('question.selectOrganisation') }}</option>
    <option v-for="org in organisations" :key="org.id" :value="org.id">{{ org.name }}</option>
  </select>

  <select 
    v-model="selectedQualificationId"
    class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_220px] min-w-[180px] max-w-md"
  >
    <option :value="null">{{ $t('question.selectQualification') }}</option>
    <option v-for="qual in qualifications" :key="qual.id" :value="qual.id">{{ qual.name }}</option>
  </select>
  <select 
    v-model="selectedSyllabusId"
    class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_220px] min-w-[180px] max-w-md"
  >
    <option :value="null">{{ $t('question.selectSyllabus') }}</option>
    <option v-for="syllabus in syllabi" :key="syllabus.id" :value="syllabus.id">{{ syllabus.name }}</option>
  </select>
  
  <select 
    v-model="selectedDifficulty"
    class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_180px] min-w-[140px] max-w-md"
  >
    <option value="">{{ $t('examPaperForm.allDifficulty') }}</option>
    <option value="1">{{ $t('examPaperForm.easy') }}</option>
    <option value="2">{{ $t('examPaperForm.medium') }}</option>
    <option value="3">{{ $t('examPaperForm.hard') }}</option>
    <option value="4">Very Hard</option>
    <option value="5">Extremely Hard</option>
  </select>

  <select 
    v-model="selectedStatus"
    class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_180px] min-w-[140px] max-w-md"
  >
    <option value="">{{ $t('common.actions') }}</option>
    <option value="1">Active</option>
    <option value="2">Inactive</option>
  </select>
</div>

      <!-- Action Buttons Row -->
      <div class="flex justify-between items-center">
        <button
          @click="$router.push('/admin/questions/create')"
          class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition duration-150"
        >
          {{ $t('question.add') }}
        </button>
        <div class="text-sm text-gray-600">
          {{ $t('question.pageInfo', { from: ((currentPage - 1) * pageSize) + 1, to: Math.min(currentPage * pageSize, totalQuestions), total: totalQuestions }) }}
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center items-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600"></div>
      <span class="ml-3 text-gray-600">Loading questions...</span>
    </div>

    <!-- No Data State -->
    <div v-else-if="!questions || questions.length === 0" class="text-center py-12">
      <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
      </svg>
      <h3 class="mt-2 text-sm font-medium text-gray-900">{{ $t('question.noData') }}</h3>
      <p class="mt-1 text-sm text-gray-500">{{ $t('question.add') }}</p>
    </div>

    <!-- Questions Grid -->
    <div v-else class="grid grid-cols-1 gap-6">
      <div 
        v-for="question in questions" 
        :key="question.id"
        class="bg-white rounded-lg shadow-md border border-gray-200 hover:shadow-lg transition-shadow duration-200"
      >
        <!-- Card Header -->
        <div class="p-4 border-b border-gray-200">
          <div class="flex justify-between items-start">
            <div class="flex-1">
              <div class="flex items-center space-x-2 mb-2">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                  {{ $t('common.id') }}: {{ question.id }}
                </span>
                <span 
                  class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                  :class="getDifficultyClass(question.difficult)"
                >
                  {{ $t('question.difficulty') }}: {{ getDifficultyName(question.difficult) }}
                </span>
                <span 
                  class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                  :class="getStatusClass(question.status)"
                >
                  {{ getStatusName(question.status) }}
                </span>
              </div>
              <h3 class="text-sm font-medium text-gray-900 mb-1">
                {{ question.syllabus?.name || $t('examPaperForm.unknown') }}
              </h3>
              <p class="text-xs text-gray-500">
                {{ $t('examPaperForm.totalScore') }}: {{ question.totalScore || $t('examPaperForm.unknown') }}
              </p>
            </div>
          </div>
        </div>

        <!-- Card Content -->
        <div class="p-4">
          <!-- Question Stem -->
          <div class="mb-4">
            <!-- <h4 class="text-sm font-medium text-gray-700 mb-2">Question:</h4> -->
            <!-- <div 
              class="text-sm text-gray-900 bg-gray-50 p-3 rounded-md whitespace-pre-wrap"
              v-html="question.stem"
            >
            </div> -->
            <QuillEditor v-model="question.stem" readOnly height="100%"></QuillEditor>
          </div>

          <!-- Question Contents Tabs -->
          <div v-if="question.questionContents && question.questionContents.length > 0" class="mb-4">
            <h4 class="text-sm font-medium text-gray-700 mb-2">
              {{ $t('examPaperForm.parts') }} ({{ $t('examPaperForm.totalScore') }}: {{ question.totalScore }})
            </h4>
            
            <!-- Tabs Navigation -->
            <div class="border border-gray-300 rounded-lg overflow-hidden">
              <div class="flex overflow-x-auto bg-gray-50 border-b border-gray-300">
                <button
                  v-for="(content, index) in question.questionContents"
                  :key="index"
                  @click="() => activeTabMap.set(question.id, index)"
                  :class="[
                    'px-4 py-2 text-sm font-medium border-r border-gray-200 whitespace-nowrap transition',
                    activeTabMap.get(question.id) === index
                      ? 'text-blue-600 bg-white border-b-2 border-b-blue-500 -mb-px'
                      : 'text-gray-700 hover:text-blue-600 hover:bg-gray-100'
                  ]"
                >
                  {{ content.partLabel }}{{ content.subpartLabel ? '.' + content.subpartLabel : '' }}
                  <span class="ml-1 text-xs text-gray-500">({{ content.score }} pts)</span>
                </button>
              </div>

              <!-- Tab Content -->
              <div v-if="question.questionContents[activeTabMap.get(question.id) || 0]" class="p-3 bg-white">
                <div class="mb-2 flex items-center gap-2">
                  <span class="text-xs font-medium text-gray-500">
                    {{ $t('examPaperForm.type') }}: {{ QUESTION_TYPE_NAMES[
                      question.questionContents[activeTabMap.get(question.id) || 0].questionTypeId as keyof typeof QUESTION_TYPE_NAMES
                    ] || $t('examPaperForm.unknown') }}
                  </span>
                  <span class="text-xs text-gray-400">|</span>
                  <span class="text-xs text-gray-500">
                    {{ $t('examPaperForm.score') }}: {{ question.questionContents[activeTabMap.get(question.id) || 0].score }}
                  </span>
                </div>

                <!-- Question Content by Type -->
                <div class="mt-3 text-sm text-gray-600">
                  <!-- Single Choice -->
<div v-if="question.questionContents[activeTabMap.get(question.id) || 0].questionTypeId === QUESTION_TYPE_SINGLE_CHOICE" class="space-y-2">
  <div v-for="option in question.questionContents[activeTabMap.get(question.id) || 0].singleChoice?.options" 
       :key="option.prefix" 
       class="flex gap-2">
    <span class="text-gray-500">{{ option.prefix }}.</span>
    <span v-html="option.content"></span>
  </div>
  <div class="mt-2">
    <span class="font-semibold">{{ $t('examPaperForm.correctAnswer') }}</span>
    <span class="ml-2 text-blue-600">
      {{
        question.questionContents[activeTabMap.get(question.id) || 0].singleChoice?.answer
      }}
    </span>
  </div>
</div>

                  <!-- Multiple Choice -->
<div v-else-if="question.questionContents[activeTabMap.get(question.id) || 0].questionTypeId === QUESTION_TYPE_MULTIPLE_CHOICE" class="space-y-2">
  <div v-for="option in question.questionContents[activeTabMap.get(question.id) || 0].multipleChoice?.options" 
       :key="option.prefix" 
       class="flex gap-2">
    <span class="text-gray-500">{{ option.prefix }}.</span>
    <span v-html="option.content"></span>
  </div>
  <div class="mt-2">
    <span class="font-semibold">Correct Answer:</span>
    <span class="ml-2 text-blue-600">
      {{
        (question.questionContents[activeTabMap.get(question.id) || 0].multipleChoice?.answer || []).join(', ')
      }}
    </span>
  </div>
</div>

                  <!-- True/False -->
                  <div v-else-if="question.questionContents[activeTabMap.get(question.id) || 0].questionTypeId === QUESTION_TYPE_TRUE_FALSE" class="space-y-2">
                    <div>
                      <span class="font-semibold">Correct Answer:</span>
                      <span class="ml-2">
                        {{
                          String(question.questionContents[activeTabMap.get(question.id) || 0].trueOrFalse?.answer).toLowerCase() === 'true'
                            ? 'True'
                            : 'False'
                        }}
                      </span>
                    </div>
                  </div>

                  <!-- Gap Filling -->
                  <div v-else-if="question.questionContents[activeTabMap.get(question.id) || 0].questionTypeId === QUESTION_TYPE_GAP_FILLING" class="space-y-2">
                    <div v-for="(ans, gidx) in question.questionContents[activeTabMap.get(question.id) || 0].gapFilling?.answer" :key="gidx">
                      <span class="font-semibold">{{ $t('examPaperForm.gap') }} {{ gidx + 1 }}:</span>
                      <span class="ml-2 font-semibold">{{ $t('examPaperForm.correctAnswer') }}</span>
                      <span class="ml-2">{{ ans }}</span>
                    </div>
                  </div>

                  <!-- Short Answer -->
                  <div v-else-if="question.questionContents[activeTabMap.get(question.id) || 0].questionTypeId === QUESTION_TYPE_SHORT_ANSWER" class="space-y-2">
                    <div>
                      <span class="font-semibold">Correct Answer:</span>
<QuillEditor
  :model-value="question.questionContents[activeTabMap.get(question.id) || 0].shortAnswer?.answer || ''"
  readOnly
  height="100%"
></QuillEditor>
                    </div>
                  </div>

                  <!-- Other types -->
                  <div v-else class="italic">
                    {{ $t('examPaperForm.pleaseViewFull') }}
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Past Paper Info -->
          <div v-if="question.pastPaper" class="mb-4">
            <h4 class="text-sm font-medium text-gray-700 mb-2">{{ $t('pastPaper.title') }}:</h4>
            <div class="text-xs text-gray-600 bg-yellow-50 p-2 rounded">
              {{ question.pastPaper.name }} ({{ question.pastPaper.year }})
              <span v-if="question.indexInPastPaper"> - Q{{ question.indexInPastPaper }}</span>
            </div>
          </div>
        </div>

        <!-- Card Footer -->
        <div class="px-4 py-3 bg-gray-50 border-t border-gray-200 flex justify-between items-center">
          <div class="text-xs text-gray-500">
            {{ formatDate(question.updatedAt) }}
          </div>
          <div class="flex space-x-2">
            <div class="flex space-x-2">
              <router-link 
                :to="`/admin/questions/${question.id}`"
                class="text-indigo-600 hover:text-indigo-900 text-sm font-medium"
              >
                {{ $t('common.view') }}
              </router-link>
              <router-link 
                :to="`/admin/questions/${question.id}/edit`" 
                class="text-indigo-600 hover:text-indigo-900 text-sm font-medium"
              >
                {{ $t('common.edit') }}
              </router-link>
              <button 
                @click="deleteQuestion(question.id)"
                class="text-red-600 hover:text-red-900 text-sm font-medium"
              >
                {{ $t('common.delete') }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Pagination -->
    <div v-if="!loading && totalQuestions > 0" class="mt-8 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0">
      <div class="text-sm text-gray-700">
        Showing {{ ((currentPage - 1) * pageSize) + 1 }} to {{ Math.min(currentPage * pageSize, totalQuestions) }} of {{ totalQuestions }} results
      </div>
      
      <nav class="flex items-center space-x-2" v-if="totalPages > 1">
        <button 
          @click="goToPage(currentPage - 1)"
          :disabled="currentPage === 1"
          class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Previous
        </button>
        
        <button 
          v-for="page in paginationRange" 
          :key="page"
          @click="goToPage(page)"
          :class="[
            'px-3 py-2 text-sm font-medium rounded-md',
            page === currentPage 
              ? 'text-white bg-indigo-600 border border-indigo-600' 
              : 'text-gray-700 bg-white border border-gray-300 hover:bg-gray-50'
          ]"
        >
          {{ page }}
        </button>
        
        <button 
          @click="goToPage(currentPage + 1)"
          :disabled="currentPage === totalPages"
          class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Next
        </button>
      </nav>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import questionService from '../../services/questionService';
import type { Question } from '../../models/question.model';
import organisationService from '../../services/organisationService';
import type { Organisation } from '../../models/organisation.model';
import qualificationService from '../../services/qualificationService';
import type { Qualification } from '../../models/qualification.model';
import syllabusService from '../../services/syllabusService';
import type { Syllabus } from '../../models/syllabus.model';
import { 
  DIFFICULTY_NAMES, 
  QUESTION_STATUS_NAMES,
  QUESTION_STATE_NORMAL,
  QUESTION_STATE_FORBIDDEN,
  QUESTION_TYPE_NAMES,
  QUESTION_TYPE_SINGLE_CHOICE,
  QUESTION_TYPE_MULTIPLE_CHOICE,
  QUESTION_TYPE_TRUE_FALSE,
  QUESTION_TYPE_GAP_FILLING,
  QUESTION_TYPE_SHORT_ANSWER
} from '../../models/question.model';
import QuillEditor from '../../components/QuillEditor/index.vue';

// Reactive data
const questions = ref<Question[]>([]);
const loading = ref(true);
const activeTabMap = ref(new Map<number, number>()); // Map to store active tab index for each question
const totalQuestions = ref(0);
const currentPage = ref(1);
const pageSize = 12; // Cards per page
const searchQuery = ref('');
const paperNameQuery = ref('');
const selectedDifficulty = ref('');
const selectedStatus = ref('');

// Filter state
const organisations = ref<Organisation[]>([]);
const qualifications = ref<Qualification[]>([]);
const syllabi = ref<Syllabus[]>([]);
const selectedOrganisationId = ref<number | null>(null);
const selectedQualificationId = ref<number | null>(null);
const selectedSyllabusId = ref<number | null>(null);

// Computed properties
const totalPages = computed(() => {
  return Math.ceil(totalQuestions.value / pageSize);
});

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

  for (let i = start; i <= end; i++) {
    if (i > 0) range.push(i);
  }
  return range;
});

// Methods for cascading dropdowns
const fetchOrganisations = async () => {
  try {
    const response = await organisationService.getOrganisations({ pageIndex: 1, pageSize: 1000 }); // Assuming a large enough page size
    organisations.value = response.data.list;
  } catch (error) {
    console.error('Failed to fetch organisations:', error);
    organisations.value = [];
  }
};

const fetchQualifications = async () => {
  if (!selectedOrganisationId.value) {
    qualifications.value = [];
    syllabi.value = [];
    selectedQualificationId.value = null;
    selectedSyllabusId.value = null;
    return;
  }
  try {
    // Assuming qualificationService.getQualifications can filter by organisationId
    // You might need to adjust this based on your actual service implementation
    const response = await qualificationService.getQualifications({ 
      pageIndex: 1, 
      pageSize: 1000, 
      organisationId: Number(selectedOrganisationId.value) 
    });
    qualifications.value = response.data.list;
  } catch (error) {
    console.error('Failed to fetch qualifications:', error);
    qualifications.value = [];
  }
  syllabi.value = [];
    selectedSyllabusId.value = null;
};

const fetchSyllabi = async () => {
  if (!selectedQualificationId.value) {
    syllabi.value = [];
    selectedSyllabusId.value = null;
    return;
  }
  try {
    // Assuming syllabusService.getSyllabuses can filter by qualificationId
    // You might need to adjust this based on your actual service implementation
    const response = await syllabusService.getSyllabuses({ 
      pageIndex: 1, 
      pageSize: 1000, 
      qualificationId: Number(selectedQualificationId.value) 
    });
    syllabi.value = response.data.list;
  } catch (error) {
    console.error('Failed to fetch syllabi:', error);
    syllabi.value = [];
  }
  selectedSyllabusId.value = null; // Reset syllabus selection when new list is fetched
};


// Methods
const fetchQuestions = async () => {
  loading.value = true;
  try {
    const response = await questionService.getQuestions({
      pageIndex: currentPage.value,
      pageSize,
      syllabusId: selectedSyllabusId.value ? Number(selectedSyllabusId.value) : undefined,
      difficult: selectedDifficulty.value ? Number(selectedDifficulty.value) : undefined,
      status: selectedStatus.value ? Number(selectedStatus.value) : undefined,
      stem: searchQuery.value.trim() || undefined,
      paperName: paperNameQuery.value.trim() || undefined
    });
    questions.value = response.data.list;
    totalQuestions.value = response.data.total;
  } catch (error) {
    console.error('Failed to fetch questions:', error);
    // TODO: Show error message to user
  } finally {
    loading.value = false;
  }
};

const deleteQuestion = async (id: number) => {
  if (confirm('Are you sure you want to delete this question? This action cannot be undone.')) {
    try {
      await questionService.deleteQuestion(id);
      if (questions.value.length === 1 && currentPage.value > 1) {
        currentPage.value--;
      }
      fetchQuestions();
    } catch (error) {
      console.error('Failed to delete question:', error);
      // TODO: Show error message to user
    }
  }
};

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
    currentPage.value = page;
    fetchQuestions();
  }
};

// Utility methods
const getDifficultyName = (difficult: number): string => {
  return DIFFICULTY_NAMES[difficult as keyof typeof DIFFICULTY_NAMES] || 'Unknown';
};

const getDifficultyClass = (difficult: number): string => {
  const classes = {
    1: 'bg-green-100 text-green-800',
    2: 'bg-yellow-100 text-yellow-800',
    3: 'bg-orange-100 text-orange-800',
    4: 'bg-red-100 text-red-800',
    5: 'bg-purple-100 text-purple-800'
  };
  return classes[difficult as keyof typeof classes] || 'bg-gray-100 text-gray-800';
};

const getStatusName = (status: number): string => {
  return QUESTION_STATUS_NAMES[status as keyof typeof QUESTION_STATUS_NAMES] || 'Unknown';
};

const getStatusClass = (status: number): string => {
  if (status === QUESTION_STATE_NORMAL) {
    return 'bg-green-100 text-green-800';
  } else if (status === QUESTION_STATE_FORBIDDEN) {
    return 'bg-red-100 text-red-800';
  }
  return 'bg-gray-100 text-gray-800';
};

const formatDate = (dateString?: string): string => {
  if (!dateString) return 'N/A';
  return new Date(dateString).toLocaleDateString();
};

// Search debounce
let searchDebounceTimer: number | undefined;
watch([searchQuery, paperNameQuery, selectedSyllabusId, selectedDifficulty, selectedStatus], () => {
  clearTimeout(searchDebounceTimer);
  searchDebounceTimer = window.setTimeout(() => {
    currentPage.value = 1;
    fetchQuestions();
  }, 500);
});

watch(selectedSyllabusId, (newValue) => {
  if (!newValue) { // If syllabus is cleared
    questions.value = [];
    totalQuestions.value = 0;
  }
});

watch(selectedOrganisationId, (newValue) => {
    selectedQualificationId.value = null;
    selectedSyllabusId.value = null;
  qualifications.value = [];
  syllabi.value = [];
  
  if (newValue) {
    fetchQualifications();
    // Don't fetch questions here, let the main filter watcher do it if syllabus changes
    // or if no syllabus is selected, questions should be empty or for the whole org (not implemented yet)
    // For now, if org is selected but no syllabus, clear questions.
     questions.value = []; // Clear questions as syllabus is now reset
     totalQuestions.value = 0;
  } else {
    // Organisation filter cleared, fetch all questions (or based on other active filters)
    fetchQuestions();
  }
});

watch(selectedQualificationId, (newValue) => {
  selectedSyllabusId.value = null;
  syllabi.value = [];

  if (newValue) {
    fetchSyllabi();
    // Similar to above, clear questions as syllabus is now reset
    questions.value = []; 
    totalQuestions.value = 0;
  } else if (selectedOrganisationId.value) {
    // Qualification cleared, but organisation is still selected.
    // Fetch questions for the organisation (if such a feature is desired) or clear.
    // For now, clear questions as no specific syllabus path is selected.
    questions.value = [];
    totalQuestions.value = 0;
    // Potentially fetch questions for the selectedOrganisationId if that's a requirement
    // fetchQuestions(); // This would fetch based on current filters (likely no syllabusId)
  }
  // If selectedQualificationId is cleared and selectedOrganisationId is also cleared,
  // the selectedOrganisationId watcher would have already triggered a fetchQuestions for all.
});

// The main watch for [searchQuery, selectedSyllabusId, ...] handles fetching when selectedSyllabusId is set.

onMounted(() => {
  fetchOrganisations(); // Load organisations first
  fetchQuestions();     // Then load initial questions (likely all, as filters are empty)
});
</script>

<style scoped>
.line-clamp-4 {
  display: -webkit-box;
  -webkit-line-clamp: 4;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
