<template>
  <div class="p-6">
    <!-- Header Section -->
    <header class="mb-6">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">
            {{ isEdit ? 'Edit Question' : 'Create Question' }}
          </h1>
          <p class="mt-1 text-sm text-gray-500">
            {{ isEdit ? 'Update question information' : 'Add a new question to the system' }}
          </p>
          <div v-if="isEdit" class="mt-3 space-y-1">
            <div v-if="syllabusName" class="text-sm text-gray-700">
              <span class="font-semibold">Syllabus:</span> {{ syllabusName }}
            </div>
            <div v-if="pastPaperName" class="text-sm text-gray-700">
              <span class="font-semibold">Past Paper:</span> {{ pastPaperName }}
            </div>
            <div v-if="form.indexInPastPaper" class="text-sm text-gray-700">
              <span class="font-semibold">Question Number:</span> {{ form.indexInPastPaper }}
            </div>
          </div>
        </div>
        <router-link 
          to="/admin/questions"
          class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
          </svg>
          Back to Questions
        </router-link>
      </div>
    </header>

    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center items-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600"></div>
      <span class="ml-3 text-gray-600">Loading question data...</span>
    </div>

    <!-- Form -->
    <form v-else @submit.prevent="handleSubmit" class="space-y-8">
      <!-- Basic Information -->
      <div class="bg-white shadow-sm ring-1 ring-gray-900/5 sm:rounded-xl">
        <div class="px-8 py-6 sm:p-10">
          <div class="grid max-w-5xl grid-cols-1 gap-x-10 gap-y-8 sm:grid-cols-6">
            <div class="sm:col-span-6">
              <h2 class="text-base font-semibold leading-7 text-gray-900">Basic Information</h2>
              <p class="mt-1 text-sm leading-6 text-gray-600">Essential question details and metadata.</p>
            </div>
            <!-- Outstanding Info: Organisation, Qualification, Syllabus, Past Paper, Paper Series, Paper Code, Question Number -->
            <div class="sm:col-span-6">
              <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-8">
                <div v-if="organisationName">
                  <div class="text-xs font-semibold text-gray-500">Organisation</div>
                  <div class="text-lg font-bold text-indigo-700">{{ organisationName }}</div>
                </div>
                <div v-if="qualificationName">
                  <div class="text-xs font-semibold text-gray-500">Qualification</div>
                  <div class="text-lg font-bold text-indigo-700">{{ qualificationName }}</div>
                </div>
                <div v-if="syllabusName">
                  <div class="text-xs font-semibold text-gray-500">Syllabus</div>
                  <div class="text-lg font-bold text-indigo-700">
                    {{ syllabusName }}
                    <span v-if="syllabusCode" class="ml-2 text-base font-semibold text-gray-400">({{ syllabusCode }})</span>
                  </div>
                </div>
                <div v-if="pastPaperName">
                  <div class="text-xs font-semibold text-gray-500">Past Paper</div>
                  <div class="text-lg font-bold text-indigo-700">{{ pastPaperName }}</div>
                </div>
                <div v-if="paperSeriesName">
                  <div class="text-xs font-semibold text-gray-500">Paper Series</div>
                  <div class="text-lg font-bold text-indigo-700">{{ paperSeriesName }}</div>
                </div>
                <div v-if="paperCodeName">
                  <div class="text-xs font-semibold text-gray-500">Paper Code</div>
                  <div class="text-lg font-bold text-indigo-700">{{ paperCodeName }}</div>
                </div>
                <div v-if="form.indexInPastPaper">
                  <div class="text-xs font-semibold text-gray-500">Question Number</div>
                  <div class="text-lg font-bold text-indigo-700">{{ form.indexInPastPaper }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Difficulty, Score, Status -->
      <div class="bg-white shadow-sm ring-1 ring-gray-900/5 sm:rounded-xl">
        <div class="px-4 py-6 sm:p-8">
          <div class="grid max-w-2xl grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
            <div class="sm:col-span-2">
              <label for="difficult" class="block text-sm font-medium leading-6 text-gray-900">
                Difficulty <span class="text-red-500">*</span>
              </label>
              <select
                id="difficult"
                v-model="form.difficult"
                required
                class="mt-2 block w-full rounded-md border-0 py-2 px-3 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
              >
                <option value="0">Select difficulty</option>
                <option value="1">easy</option>
                <option value="2">medium</option>
                <option value="3">hard</option>
                <option value="4">very hard</option>
                <option value="5">extremely hard</option>
              </select>
            </div>
            <div class="sm:col-span-2">
              <label for="totalScore" class="block text-sm font-medium leading-6 text-gray-900">
                Total Score
              </label>
              <input
                type="number"
                id="totalScore"
                v-model.number="form.totalScore"
                min="0"
                step="0.5"
                class="mt-2 block w-full rounded-md border-0 py-1.5 px-3 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
              />
            </div>
            <div class="sm:col-span-2">
              <label for="status" class="block text-sm font-medium leading-6 text-gray-900">
                Status <span class="text-red-500">*</span>
              </label>
              <select
                id="status"
                v-model="form.status"
                required
                class="mt-2 block w-full rounded-md border-0 px-3 py-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
              >
                <option value="1">Active</option>
                <option value="2">Inactive</option>
              </select>
            </div>
          </div>
        </div>
      </div>

      <!-- Question Stem -->
      <div class="bg-white shadow-sm ring-1 ring-gray-900/5 sm:rounded-xl">
        <div class="px-4 py-6 sm:p-8">
          <div class="grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
            <div class="sm:col-span-6 w-full">
              <label for="stem" class="block text-sm font-medium leading-6 text-gray-900">
                Question Stem <span class="text-red-500">*</span>
              </label>
              <QuillEditor
                id="stem"
                v-model="form.stem"
                :placeholder="'Enter the main question text...'"
                height="120px"
              />
            </div>
          </div>
        </div>
      </div>


      <!-- Question Contents -->
      <div class="bg-white shadow-sm ring-1 ring-gray-900/5 sm:rounded-xl">
        <div class="px-4 py-6 sm:p-8">
          <div class="space-y-6">
            <div>
              <div class="flex items-center justify-between mb-4">
                <div>
                  <h2 class="text-base font-semibold leading-7 text-gray-900">Question Parts</h2>
                  <p class="mt-1 text-sm leading-6 text-gray-600">Break down the question into individual parts.</p>
                </div>
                <button
                  type="button"
                  @click="addQuestionContent"
                  class="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-indigo-700 bg-indigo-100 hover:bg-indigo-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                >
                  <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                  </svg>
                  Add Part
                </button>
              </div>

              <div v-if="form.questionContents.length === 0" class="text-center py-8 text-gray-500">
                No question parts added yet. Click "Add Part" to get started.
              </div>

              <div v-else class="space-y-4">
                <div
                  v-for="(content, index) in form.questionContents"
                  :key="index"
                  class="border border-gray-200 rounded-lg p-4"
                >
                  <div class="flex items-center justify-between mb-4">
                    <h3 class="text-sm font-medium text-gray-900">Part {{ index + 1 }}</h3>
                    <button
                      type="button"
                      @click="removeQuestionContent(index)"
                      class="text-red-600 hover:text-red-800 text-sm"
                    >
                      Remove
                    </button>
                  </div>

                  <div class="grid grid-cols-1 gap-4 sm:grid-cols-6">
                    <div class="sm:col-span-1">
                      <label class="block text-sm font-medium text-gray-700">Part Label</label>
                      <input
                        type="text"
                        v-model="content.partLabel"
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-1.5"
                        placeholder="a"
                      />
                    </div>

                    <div class="sm:col-span-1">
                      <label class="block text-sm font-medium text-gray-700">Subpart Label</label>
                      <input
                        type="text"
                        v-model="content.subpartLabel"
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-1.5"
                        placeholder="i"
                      />
                    </div>

                    <div class="sm:col-span-1">
                      <label class="block text-sm font-medium text-gray-700">Score</label>
                      <input
                        type="number"
                        v-model.number="content.score"
                        min="0"
                        step="0.5"
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-1.5"
                      />
                    </div>

                    <!-- Question Type -->
                    <div class="sm:col-span-2">
                      <label class="block text-sm font-medium text-gray-700">Question Type</label>
                      <select
                        v-model="content.questionTypeId"
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-1.5"
                      >
                        <option :value="undefined">请选择题型</option>
                        <option :value="1">单选题</option>
                        <option :value="2">多选题</option>
                        <option :value="3">判断题</option>
                        <option :value="4">填空题</option>
                        <option :value="5">简答题</option>
                      </select>
                    </div>
                  </div>

                  

                  <!-- Dynamic fields based on questionTypeId -->
                  <div v-if="content.questionTypeId === 1" class="mt-4">
                    <label class="block text-sm font-medium text-gray-700 mb-2">单选题选项</label>
                    <div class="grid grid-cols-1 gap-2">
                      <div v-for="(option, idx) in content.singleChoice?.options || []" :key="idx" class="flex items-center gap-2">
                        <input
                          type="text"
                          v-model="option.content"
                          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2"
                          :placeholder="`选项${option.prefix}`"
                        />
                        <label class="flex items-center text-xs text-gray-600 w-30">
                          <input
                            type="radio"
                            v-model="content.singleChoice!.answer"
                            :value="option.prefix"
                            class="mr-1 w-10"
                          />
                          正确
                        </label>
                        <button type="button" @click="content.singleChoice!.options.splice(idx, 1)" class="text-red-500 text-xs ml-2 w-20">移除</button>
                      </div>
                    </div>
                    <button
                      type="button"
                      class="mt-2 px-3 py-1 bg-indigo-100 text-indigo-700 rounded text-xs"
                      @click="() => {
                        if (!content.singleChoice) content.singleChoice = { options: [], answer: '' };
                        const nextPrefix = String.fromCharCode(65 + (content.singleChoice.options.length || 0));
                        content.singleChoice.options.push({ prefix: nextPrefix, content: '' });
                      }"
                    >
                      添加选项
                    </button>
                  </div>
                  <div v-else-if="content.questionTypeId === 2" class="mt-4">
                    <label class="block text-sm font-medium text-gray-700 mb-2">多选题选项</label>
                    <div class="grid grid-cols-1 gap-2">
                      <div v-for="(option, idx) in content.multipleChoice?.options || []" :key="idx" class="flex items-center gap-2">
                        <input
                          type="text"
                          v-model="option.content"
                          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2"
                          :placeholder="`选项${option.prefix}`"
                        />
                        <label class="flex items-center text-xs text-gray-600 w-20">
                          <input
                            type="checkbox"
                            :value="option.prefix"
                            v-model="content.multipleChoice!.answer"
                            class="mr-1"
                          />
                          正确
                        </label>
                        <button type="button" @click="content.multipleChoice!.options.splice(idx, 1)" class="text-red-500 text-xs ml-2 w-10">移除</button>
                      </div>
                    </div>
                    <button
                      type="button"
                      class="mt-2 px-3 py-1 bg-indigo-100 text-indigo-700 rounded text-xs"
                      @click="() => {
                        if (!content.multipleChoice) content.multipleChoice = { options: [], answer: [] };
                        const nextPrefix = String.fromCharCode(65 + (content.multipleChoice.options.length || 0));
                        content.multipleChoice.options.push({ prefix: nextPrefix, content: '' });
                      }"
                    >
                      添加选项
                    </button>
                  </div>
                  <div v-else-if="content.questionTypeId === 3" class="mt-4">
                    <label class="block text-sm font-medium text-gray-700 mb-2">判断答案</label>
                    <select
                      :value="content.trueOrFalse?.answer ?? 1"
                      @change="e => {
                        if (!content.trueOrFalse) content.trueOrFalse = { answer: 1 };
                        content.trueOrFalse.answer = Number((e.target as HTMLSelectElement).value);
                      }"
                      class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2"
                    >
                      <option :value="1">正确</option>
                      <option :value="2">错误</option>
                    </select>
                  </div>
                  <div v-else-if="content.questionTypeId === 4" class="mt-4">
                    <div class="space-y-4">
                      <label class="block text-sm font-medium text-gray-700">填空题答案</label>
                      <div class="space-y-2">
                        <div v-for="(_, idx) in (content.gapFilling?.answer || [])" :key="idx" class="flex items-center gap-2">
                          <div class="flex-grow flex items-center gap-2">
                            <span class="text-sm text-gray-500 w-32">填空答案 {{ idx + 1 }}</span>
                            <input
                              type="text"
                              v-model="content.gapFilling!.answer[idx]"
                              class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2"
                              :placeholder="`答案 ${idx + 1}`"
                            />
                          </div>
                          <button 
                            type="button" 
                            @click="() => content.gapFilling!.answer.splice(idx, 1)"
                            class="text-red-500 hover:text-red-700"
                          >
                            删除
                          </button>
                        </div>
                      </div>
                      <button
                        type="button"
                        @click="() => {
                          if (!content.gapFilling) content.gapFilling = { answer: [] };
                          content.gapFilling.answer.push('');
                        }"
                        class="inline-flex items-center px-3 py-1 border border-transparent text-sm leading-4 font-medium rounded-md text-indigo-700 bg-indigo-100 hover:bg-indigo-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                      >
                        <span class="mr-1">+</span> 添加答案
                      </button>
                    </div>
                  </div>
                  <div v-else-if="content.questionTypeId === 5" class="mt-4">
                    <label class="block text-sm font-medium text-gray-700 mb-2">简答题答案</label>
                    <QuillEditor
                      v-model="(content.shortAnswer ?? (content.shortAnswer = { answer: '' })).answer"
                    />
                  </div>
                  <!-- Answer Analysis -->
                  <div class="sm:col-span-6 mt-4">
                    <label class="block text-sm font-medium text-gray-700">Answer Analysis</label>
                    <QuillEditor
                      v-model="content.analyze as string"
                      height="200px"
                      placeholder="Enter a detailed analysis and explanation of the answer..."
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Form Actions -->
      <div class="flex items-center justify-end gap-x-6">
        <router-link 
          to="/admin/questions"
          class="text-sm font-semibold leading-6 text-gray-900"
        >
          Cancel
        </router-link>
        <button
          type="submit"
          :disabled="submitting"
          class="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          <span v-if="submitting">{{ isEdit ? 'Updating...' : 'Creating...' }}</span>
          <span v-else>{{ isEdit ? 'Update Question' : 'Create Question' }}</span>
        </button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import questionService from '../../services/questionService';
import type { QuestionContent } from '../../models/question.model';
import QuillEditor from '../../components/QuillEditor/index.vue'

// Outstanding info refs
const organisationName = ref('');
const qualificationName = ref('');
const syllabusName = ref('');
const syllabusCode = ref('');
const pastPaperName = ref('');
const paperSeriesName = ref('');
const paperCodeName = ref('');

const route = useRoute();
const router = useRouter();

// Component state
const loading = ref(false);
const submitting = ref(false);

// Form data
const form = ref({
  id: 0,
  syllabusId: '',
  difficult: '',
  totalScore: 0,
  status: '1',
  stem: '',
  pastPaperId: '',
  indexInPastPaper: undefined as number | undefined,
  questionContents: [] as QuestionContent[]
});

// Computed properties
const isEdit = computed(() => {
  console.log('Checking if in edit mode:', route.name, route.params);
  return route.name === 'AdminQuestionEdit' && !!route.params.id;
});

const questionId = computed(() => {
  return isEdit.value ? Number(route.params.id) : null;
});

// Methods
const fetchQuestion = async () => {
  if (!isEdit.value || !questionId.value) return;

  loading.value = true;
  try {
    const response = await questionService.getQuestionById(questionId.value);
    const question = response.data;
    
    form.value = {
      id: question.id,
      syllabusId: question.syllabusId?.toString() || '',
      difficult: question.difficult?.toString() || '',
      totalScore: question.totalScore || 0,
      status: question.status?.toString() || '1',
      stem: question.stem || '',
      pastPaperId: question.pastPaperId?.toString() || '',
      indexInPastPaper: question.indexInPastPaper,
      questionContents: question.questionContents || []
    };

    // Outstanding info assignment (simulate or use actual API if available)
    organisationName.value = question.syllabus?.qualification?.organisation?.name || '';
    qualificationName.value = question.syllabus?.qualification?.name || '';
    syllabusName.value = question.syllabus?.name || '';
    syllabusCode.value = question.syllabus?.code || '';
    pastPaperName.value = question.pastPaper?.name
      ? `${question.pastPaper.name}${question.pastPaper.year ? ' (' + question.pastPaper.year + ')' : ''}`
      : '';
    paperSeriesName.value = question.pastPaper?.paperSeries.name || '';
    paperCodeName.value = question.pastPaper?.paperCode.name || '';
  } catch (error) {
    console.error('Failed to fetch question:', error);
    // TODO: Show error message to user
  } finally {
    loading.value = false;
  }
};

const addQuestionContent = () => {
  form.value.questionContents.push({ analyze: '' } as QuestionContent);
};

const removeQuestionContent = (index: number) => {
  form.value.questionContents.splice(index, 1);
};

const handleSubmit = async () => {
  submitting.value = true;
  try {
    const formData = {
      ...form.value,
      syllabusId: form.value.syllabusId ? Number(form.value.syllabusId) : 0,
      difficult: form.value.difficult ? Number(form.value.difficult) : 0,
      status: Number(form.value.status),
      pastPaperId: form.value.pastPaperId ? Number(form.value.pastPaperId) : undefined
    };

    if (isEdit.value) {
      await questionService.updateQuestion(formData);
    } else {
      await questionService.createQuestion(formData);
    }

    router.push('/admin/questions');
  } catch (error) {
    console.error('Failed to save question:', error);
    // TODO: Show error message to user
  } finally {
    submitting.value = false;
  }
};

onMounted(() => {
  if (isEdit.value) {
    fetchQuestion();
  }
});
</script>
