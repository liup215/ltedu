<template>
  <div class="p-6">
    <!-- Header Section -->
    <header class="mb-6">
      <div class="flex justify-between items-center">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">Question Details</h1>
          <p class="mt-1 text-sm text-gray-500">View question information</p>
        </div>
        <div class="space-x-4">
          <router-link 
            :to="`/admin/questions/${question?.id}/edit`"
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Edit Question
          </router-link>
          <button 
            @click="goBack"
            class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Back to List
          </button>
        </div>
      </div>
    </header>

    <!-- Loading State -->
    <div v-if="loading" class="bg-white shadow sm:rounded-lg p-6">
      <div class="text-center">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600 mx-auto"></div>
        <p class="mt-4 text-gray-500">Loading question...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="bg-white shadow sm:rounded-lg p-6">
      <div class="text-center">
        <p class="text-red-600">{{ error }}</p>
        <button 
          @click="fetchQuestion"
          class="mt-4 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700"
        >
          Retry
        </button>
      </div>
    </div>

    <!-- Question Details -->
    <div v-else-if="question" class="bg-white shadow sm:rounded-lg">
      <div class="px-6 py-6">
        <!-- Basic Information -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-8">
          <div class="lg:col-span-2">
            <h2 class="text-xl font-semibold text-gray-900 mb-4">Question Stem</h2>
            <div class="prose max-w-none">
              <QuillEditor v-model="question.stem" readOnly height="100%"></QuillEditor>
            </div>
          </div>
          
          <div class="space-y-6">
            <!-- Question Metadata -->
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-3">Question Information</h3>
              <dl class="space-y-3">
                <div>
                  <dt class="text-sm font-medium text-gray-500">Question ID</dt>
                  <dd class="text-sm text-gray-900">{{ question.id }}</dd>
                </div>
                <div>
                  <dt class="text-sm font-medium text-gray-500">Status</dt>
                  <dd class="text-sm text-gray-900">
                    <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                          :class="getStatusColor(question.status)">
                      {{ getStatusName(question.status) }}
                    </span>
                  </dd>
                </div>
                <div>
                  <dt class="text-sm font-medium text-gray-500">Difficulty</dt>
                  <dd class="text-sm text-gray-900">
                    <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                          :class="getDifficultyColor(question.difficult)">
                      {{ getDifficultyName(question.difficult) }}
                    </span>
                  </dd>
                </div>
                <div v-if="question.syllabus">
                  <dt class="text-sm font-medium text-gray-500">Syllabus</dt>
                  <dd class="text-sm text-gray-900">{{ question.syllabus.name }}</dd>
                </div>
                <div v-if="question.totalScore">
                  <dt class="text-sm font-medium text-gray-500">Total Score</dt>
                  <dd class="text-sm text-gray-900">{{ question.totalScore }}</dd>
                </div>
                <div v-if="question.pastPaper">
                  <dt class="text-sm font-medium text-gray-500">Past Paper</dt>
                  <dd class="text-sm text-gray-900">{{ question.pastPaper.name }} ({{ question.pastPaper.year }} {{ question.pastPaper.paperSeries.name }})</dd>
                </div>
                <div v-if="question.indexInPastPaper">
                  <dt class="text-sm font-medium text-gray-500">Question Index</dt>
                  <dd class="text-sm text-gray-900">{{ question.indexInPastPaper }}</dd>
                </div>
              </dl>
            </div>

            <!-- Chapters -->
            <div v-if="question.chapters && question.chapters.length > 0">
              <h3 class="text-lg font-medium text-gray-900 mb-3">Chapters</h3>
              <div class="flex flex-wrap gap-2">
                <span 
                  v-for="chapter in question.chapters" 
                  :key="chapter.id"
                  class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-gray-100 text-gray-800"
                >
                  {{ chapter.name }}
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Question Contents -->
        <div v-if="question.questionContents && question.questionContents.length > 0" class="mb-8">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Question Parts</h3>
          <div class="space-y-6">
            <div 
              v-for="(content, index) in question.questionContents" 
              :key="index"
              class="border rounded-lg p-6 bg-gray-50"
            >
              <!-- Part Header -->
              <div class="flex justify-between items-start mb-4">
                <div>
                  <h4 class="text-md font-medium text-gray-900">
                    Part {{ content.partLabel || (index + 1) }}
                    <span v-if="content.subpartLabel" class="text-gray-600">.{{ content.subpartLabel }}</span>
                  </h4>
                  <p v-if="content.questionTypeId" class="text-sm text-gray-600 mt-1">
                    Type: {{ getQuestionTypeName(content.questionTypeId) }}
                  </p>
                </div>
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                  {{ content.score }} marks
                </span>
              </div>

              <!-- Single Choice Options -->
              <div v-if="content.questionTypeId == QUESTION_TYPE_SINGLE_CHOICE" class="mb-4">
                <h5 class="font-medium text-gray-900 mb-2">Options:</h5>
                <div class="space-y-2">
                  <div 
                    v-for="option in content.singleChoice?.options" 
                    :key="option.prefix"
                    class="flex items-start p-3 border rounded"
                    :class="option.prefix === content.singleChoice?.answer ? 'border-green-500 bg-green-50' : 'border-gray-200 bg-white'"
                  >
                    <span class="font-medium text-sm mr-3">{{ option.prefix }}.</span>
                    <span class="text-sm">{{ option.content }}</span>
                    <svg v-if="option.prefix === content.singleChoice?.answer" class="h-4 w-4 text-green-500 ml-auto" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                    </svg>
                  </div>
                </div>
              </div>

              <!-- Multiple Choice Options -->
              <div v-if="content.questionTypeId === QUESTION_TYPE_MULTIPLE_CHOICE" class="mb-4">
                <h5 class="font-medium text-gray-900 mb-2">Options:</h5>
                <div class="space-y-2">
                  <div 
                    v-for="option in content.multipleChoice?.options" 
                    :key="option.prefix"
                    class="flex items-start p-3 border rounded"
                    :class="content.multipleChoice?.answer.includes(option.prefix) ? 'border-green-500 bg-green-50' : 'border-gray-200 bg-white'"
                  >
                    <span class="font-medium text-sm mr-3">{{ option.prefix }}.</span>
                    <span class="text-sm">{{ option.content }}</span>
                    <svg v-if="content.multipleChoice?.answer.includes(option.prefix)" class="h-4 w-4 text-green-500 ml-auto" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                    </svg>
                  </div>
                </div>
              </div>

              <!-- True/False Answer -->
              <div v-if="content.questionTypeId === QUESTION_TYPE_TRUE_FALSE" class="mb-4">
                <h5 class="font-medium text-gray-900 mb-2">Correct Answer:</h5>
                <div class="p-3 border rounded-lg bg-green-50 border-green-200">
                  <span class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-green-100 text-green-800">
                    {{ content.trueOrFalse?.answer === 1 ? 'True' : 'False' }}
                  </span>
                </div>
              </div>

              <!-- Gap Filling Answer -->
              <div v-if="content.questionTypeId === QUESTION_TYPE_GAP_FILLING" class="mb-4">
                <h5 class="font-medium text-gray-900 mb-2">Answers:</h5>
                <div class="space-y-1">
                  <span 
                    v-for="(answer, idx) in content.gapFilling?.answer" 
                    :key="idx"
                    class="inline-block px-3 py-1 mr-2 mb-1 text-sm bg-blue-100 text-blue-800 rounded"
                  >
                    {{ answer }}
                  </span>
                </div>
              </div>

              <!-- Short Answer -->
              <div v-if="content.questionTypeId === QUESTION_TYPE_SHORT_ANSWER" class="mb-4">
                <h5 class="font-medium text-gray-900 mb-2">Answer:</h5>
                <div class="p-3 bg-gray-100 rounded border">
                  <div v-html="content.shortAnswer?.answer"></div>
                </div>
              </div>

              <!-- Analysis/Explanation -->
              <div v-if="content.analyze" class="mb-4">
                <h5 class="font-medium text-gray-900 mb-2">Analysis:</h5>
                <div class="p-3 bg-yellow-50 rounded border border-yellow-200">
                  <div v-html="content.analyze"></div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Metadata -->
        <div class="border-t pt-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Metadata</h3>
          <dl class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
            <div>
              <dt class="text-sm font-medium text-gray-500">Created</dt>
              <dd class="text-sm text-gray-900">{{ formatDate(question.createdAt) }}</dd>
            </div>
            <div>
              <dt class="text-sm font-medium text-gray-500">Updated</dt>
              <dd class="text-sm text-gray-900">{{ formatDate(question.updatedAt) }}</dd>
            </div>
            <div>
              <dt class="text-sm font-medium text-gray-500">Syllabus ID</dt>
              <dd class="text-sm text-gray-900">{{ question.syllabusId }}</dd>
            </div>
          </dl>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import questionService from '../../services/questionService';
import type { Question } from '../../models/question.model';
import { 
  QUESTION_TYPE_SINGLE_CHOICE, 
  QUESTION_TYPE_MULTIPLE_CHOICE, 
  QUESTION_TYPE_TRUE_FALSE, 
  QUESTION_TYPE_GAP_FILLING, 
  QUESTION_TYPE_SHORT_ANSWER 
} from '../../models/question.model';
import { 
  QUESTION_STATUS_NAMES, 
  DIFFICULTY_NAMES,
  QUESTION_STATE_NORMAL,
  QUESTION_STATE_FORBIDDEN,
  QUESTION_TYPE_NAMES
} from '../../models/question.model';
import QuillEditor from '../../components/QuillEditor/index.vue';

const route = useRoute();
const router = useRouter();

const question = ref<Question | null>(null);
const loading = ref(true);
const error = ref<string | null>(null);

const fetchQuestion = async () => {
  const questionId = Number(route.params.id);
  if (!questionId) {
    error.value = 'Invalid question ID';
    loading.value = false;
    return;
  }

  loading.value = true;
  error.value = null;
  
  try {
    const response = await questionService.getQuestionById(questionId);
    question.value = response.data;
  } catch (err) {
    console.error('Failed to fetch question:', err);
    error.value = 'Failed to load question details';
  } finally {
    loading.value = false;
  }
};

const getStatusName = (status: number): string => {
  return (QUESTION_STATUS_NAMES as any)[status] || 'Unknown';
};

const getStatusColor = (status: number): string => {
  switch (status) {
    case QUESTION_STATE_NORMAL:
      return 'bg-green-100 text-green-800';
    case QUESTION_STATE_FORBIDDEN:
      return 'bg-red-100 text-red-800';
    default:
      return 'bg-gray-100 text-gray-800';
  }
};

const getDifficultyName = (difficulty: number): string => {
  return (DIFFICULTY_NAMES as any)[difficulty] || 'Unknown';
};

const getDifficultyColor = (difficulty: number): string => {
  switch (difficulty) {
    case 1:
      return 'bg-green-100 text-green-800';
    case 2:
      return 'bg-yellow-100 text-yellow-800';
    case 3:
      return 'bg-orange-100 text-orange-800';
    case 4:
      return 'bg-red-100 text-red-800';
    case 5:
      return 'bg-purple-100 text-purple-800';
    default:
      return 'bg-gray-100 text-gray-800';
  }
};

const formatDate = (dateString?: string): string => {
  if (!dateString) return 'N/A';
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  });
};

const getQuestionTypeName = (typeId: number): string => {
  return (QUESTION_TYPE_NAMES as any)[typeId] || 'Unknown';
};

const goBack = () => {
  router.push('/admin/questions');
};

onMounted(() => {
  fetchQuestion();
});
</script>
