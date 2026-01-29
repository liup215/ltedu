<template>
<div class="p-6 w-full">
<PracticeSidebar
    :visible="sidebarVisible"
    :question-ids="questionIds"
    :questions="questions"
    :current-index="currentIndex"
    :answers="answers"
    :result="result"
    :jump-to="jumpTo"
    :on-close="() => sidebarVisible = false"
/>
<button
  v-if="questionIds.length > 0 && !sidebarVisible"
  @click="sidebarVisible = true"
  class="fixed right-4 bottom-4 z-50 bg-indigo-600 text-white px-4 py-2 rounded shadow-lg hover:bg-indigo-700"
>
  {{ $t('quickPractice.showSidebar') }}
</button>
  <div v-if="!userStore.isAuthenticated" class="mb-6 bg-yellow-100 border-l-4 border-yellow-500 text-yellow-700 p-4 rounded">
    <span>{{ $t('quickPractice.notLoggedIn') }}</span>
    <router-link to="/login" class="ml-2 text-blue-600 underline">{{ $t('quickPractice.goToLogin') }}</router-link>
  </div>
    <header class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">{{ $t('quickPractice.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('quickPractice.subtitle') }}</p>
      <div class="mt-2 bg-yellow-50 border-l-4 border-yellow-400 text-yellow-700 p-3 rounded">
        <span>
          {{ $t('quickPractice.knowledgePointNotice') }}
        </span>
      </div>
    </header>
<form @submit.prevent="startPractice" class="mb-6 flex flex-wrap gap-4 items-end w-full">
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('quickPractice.organisation') }}</label>
<select v-model="selectedOrganisationId" @change="onOrganisationChange" required class="px-3 py-2 border border-gray-300 rounded-md min-w-[180px] w-full">
          <option value="" disabled>{{ $t('quickPractice.selectOrganisation') }}</option>
          <option v-for="org in organisations" :key="org.id" :value="org.id">
            {{ org.name }}
          </option>
        </select>
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('quickPractice.qualification') }}</label>
<select v-model="selectedQualificationId" @change="onQualificationChange" :disabled="!selectedOrganisationId" required class="px-3 py-2 border border-gray-300 rounded-md min-w-[180px] w-full">
          <option value="" disabled>{{ $t('quickPractice.selectQualification') }}</option>
          <option v-for="q in qualifications" :key="q.id" :value="q.id">
            {{ q.name }}
          </option>
        </select>
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('quickPractice.syllabus') }}</label>
<select v-model="syllabusId" :disabled="!selectedQualificationId" required class="px-3 py-2 border border-gray-300 rounded-md min-w-[180px] w-full">
          <option value="" disabled>{{ $t('quickPractice.selectSyllabus') }}</option>
          <option v-for="syllabus in syllabuses" :key="syllabus.id" :value="syllabus.id">
            {{ syllabus.name }} ({{ syllabus.code }})
          </option>
        </select>
      </div>
      <!-- 章节筛选控件 -->
      <div v-if="syllabusId">
        <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('quickPractice.chapters') }}</label>
        <div class="relative min-w-[200px] max-w-xl">
          <button
            type="button"
            @click="showChapterSelector = !showChapterSelector"
            :disabled="!chapterTree.length"
            class="flex justify-between items-center w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 disabled:bg-gray-100 disabled:cursor-not-allowed"
          >
            <span class="text-sm truncate">
              {{ selectedChapterIds.length ? $t('quickPractice.chaptersSelected', { count: selectedChapterIds.length }) : $t('quickPractice.allChapters') }}
            </span>
            <svg class="w-5 h-5 text-gray-400" viewBox="0 0 20 20" fill="currentColor">
              <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
            </svg>
          </button>
          <div v-if="showChapterSelector" class="absolute z-10 w-full lg:w-96 mt-1 bg-white rounded-md shadow-lg border border-gray-200">
            <div class="max-h-96 overflow-y-auto p-2">
              <div class="flex items-center justify-between p-2 border-b">
                <span class="text-sm font-medium text-gray-900">{{ $t('quickPractice.selectChapters') }}</span>
                <button 
                  type="button"
                  @click="selectedChapterIds = []; showChapterSelector = false"
                  class="text-sm text-gray-500 hover:text-gray-700"
                >
                  {{ $t('quickPractice.clear') }}
                </button>
              </div>
              <div class="mt-2">
                <ChapterOption
                  v-for="(chapter, index) in chapterTree"
                  :key="chapter.id"
                  :chapter="chapter"
                  :level="0"
                  :is-last="index === chapterTree.length - 1"
                  :selected-chapters="selectedChapterIds"
                  @update:selected="val => { selectedChapterIds = val }"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('quickPractice.questionCount') }}</label>
<input v-model.number="questionCount" type="number" min="1" max="20" required class="px-3 py-2 border border-gray-300 rounded-md min-w-[100px] w-full" />
      </div>
      <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700">{{ $t('quickPractice.start') }}</button>
      <button type="button" @click="resetAll" class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300">{{ $t('quickPractice.reset') }}</button>
    </form>

    <div v-if="questionIds.length > 0" class="mb-6">
      <form @submit.prevent="submitAnswers">
        <div class="mb-6 bg-white rounded shadow p-4">
          <template v-if="questions[questionIds[currentIndex]]">
            <div class="mb-2 font-semibold text-gray-800">
              Q{{ currentIndex + 1 }}: 
              <QuillEditor v-model="currentQuestion.stem" readOnly height="100%"></QuillEditor>
            </div>
            <div v-if="currentQuestionContents.length > 0">
              <div v-for="(content, cidx) in currentQuestionContents" :key="cidx" class="mb-4 p-3 bg-white rounded border">
                <div class="mb-2 flex items-center gap-2">
                  <span class="text-xs font-medium text-gray-500">
                    Part: {{ content.partLabel }}{{ content.subpartLabel ? '.' + content.subpartLabel : '' }}
                  </span>
                  <span class="text-xs text-gray-400">|</span>
                  <span class="text-xs font-medium text-gray-500">
                    Type: {{ QUESTION_TYPE_NAMES[content.questionTypeId as keyof typeof QUESTION_TYPE_NAMES] || content.questionTypeId || '' }}
                  </span>
                  <span class="text-xs text-gray-400">|</span>
                  <span class="text-xs text-gray-500">
                    Score: {{ content.score }} pts
                  </span>
                </div>
                <div class="mt-3 text-sm text-gray-600">
                  <div v-if="content.questionTypeId === QUESTION_TYPE_SINGLE_CHOICE">
                    <label v-for="opt in content.singleChoice?.options" :key="opt.prefix" class="block">
                      <input type="radio" :name="'single-' + questionIds[currentIndex] + '-' + cidx" :value="opt.prefix" v-model="answers[questionIds[currentIndex] + '-' + cidx]" />
                      {{ opt.prefix }}. <span v-html="opt.content"></span>
                    </label>
                  </div>
                  <div v-else-if="content.questionTypeId === QUESTION_TYPE_MULTIPLE_CHOICE">
                    <label v-for="opt in content.multipleChoice?.options" :key="opt.prefix" class="block">
                      <input type="checkbox" :value="opt.prefix" v-model="answersMulti[questionIds[currentIndex] + '-' + cidx]" />
                      {{ opt.prefix }}. <span v-html="opt.content"></span>
                    </label>
                  </div>
                  <div v-else-if="content.questionTypeId === QUESTION_TYPE_TRUE_FALSE">
                    <label>
                      <input type="radio" :name="'tf-' + questionIds[currentIndex] + '-' + cidx" value="true" v-model="answers[questionIds[currentIndex] + '-' + cidx]" /> True
                    </label>
                    <label>
                      <input type="radio" :name="'tf-' + questionIds[currentIndex] + '-' + cidx" value="false" v-model="answers[questionIds[currentIndex] + '-' + cidx]" /> False
                    </label>
                  </div>
                  <div v-else-if="content.questionTypeId === QUESTION_TYPE_GAP_FILLING">
                    <input v-model="answers[questionIds[currentIndex] + '-' + cidx]" type="text" class="px-3 py-2 border border-gray-300 rounded-md w-full" placeholder="Fill the gap" />
                  </div>
                  <div v-else-if="content.questionTypeId === QUESTION_TYPE_SHORT_ANSWER">
                    <QuillEditor :model-value="answers[questionIds[currentIndex] + '-' + cidx] ?? ''" :readOnly="false" height="120px" :placeholder="'Your answer'" @update:modelValue="val => { answers[questionIds[currentIndex] + '-' + cidx] = val ?? ''; }" />
                  </div>
                </div>
              </div>
            </div>
          </template>
        </div>
        <div class="flex justify-center items-center gap-4 mb-4">
          <button type="button" class="px-4 py-2 bg-gray-200 text-gray-800 rounded hover:bg-gray-300" :disabled="currentIndex === 0" @click="jumpTo(currentIndex - 1)">{{ $t('quickPractice.previous') }}</button>
          <span class="font-bold text-lg">Q{{ currentIndex + 1 }} / {{ questionIds.length }}</span>
          <button type="button" class="px-4 py-2 bg-gray-200 text-gray-800 rounded hover:bg-gray-300" :disabled="currentIndex === questionIds.length - 1" @click="jumpTo(currentIndex + 1)">{{ $t('quickPractice.next') }}</button>
        </div>
        <button type="submit" class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700">{{ $t('quickPractice.submitAnswers') }}</button>
      </form>
    </div>

    <div v-if="result && questionIds.length > 0" class="mt-8 bg-white rounded shadow p-6">
      <h2 class="text-xl font-bold mb-4">{{ $t('quickPractice.practiceResult') }}</h2>
      <div class="mb-2">{{ $t('quickPractice.score') }}: {{ result.score }} / {{ result.total }}</div>
      <template v-if="resultItem">
        <div class="mb-4 border-b pb-2">
          <div class="font-semibold">Q{{ currentIndex + 1 }}</div>
          <div v-for="(sub, pidx) in resultItem.subResults" :key="pidx" class="ml-4 mb-2">
            <div>{{ $t('quickPractice.part') }} {{ pidx + 1 }}</div>
            <div v-if="sub.questionType === QUESTION_TYPE_SHORT_ANSWER">
              <div class="mb-1">{{ $t('quickPractice.yourAnswer') }}</div>
              <QuillEditor :modelValue="sub.studentAnswer" :readOnly="true" height="120px" />
              <div v-if="sub.correctAnswer" class="mb-1 mt-2">{{ $t('quickPractice.correctAnswer') }}</div>
              <QuillEditor v-if="sub.correctAnswer" :modelValue="sub.correctAnswer" :readOnly="true" height="120px" />
              <div v-if="sub.modelAnswer" class="mb-1 mt-2">{{ $t('quickPractice.modelAnswer') }}</div>
              <QuillEditor v-if="sub.modelAnswer" :modelValue="sub.modelAnswer" :readOnly="true" height="120px" />
            </div>
            <div v-else>
              <div>{{ $t('quickPractice.yourAnswer') }}: {{ sub.studentAnswer }}</div>
              <div>{{ $t('quickPractice.correctAnswer') }}: {{ sub.correctAnswer }}</div>
              <div v-if="sub.isCorrect === true" class="text-green-600">{{ $t('quickPractice.correct') }}</div>
              <div v-else-if="sub.isCorrect === false" class="text-red-600">{{ $t('quickPractice.incorrect') }}</div>
              <div v-else class="text-gray-600">{{ $t('quickPractice.subjective') }}</div>
              <div v-if="sub.modelAnswer">{{ $t('quickPractice.modelAnswer') }}: {{ sub.modelAnswer }}</div>
            </div>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, watch, computed } from 'vue'
import { useUserStore } from '../stores/userStore'
import QuillEditor from '../components/QuillEditor/index.vue'
import PracticeSidebar from '../components/PracticeSidebar.vue'
import { practiceService } from '../services/practiceService'
import syllabusService from '../services/syllabusService'
import organisationService from '../services/organisationService'
import qualificationService from '../services/qualificationService'
import chapterService from '../services/chapterService'
import ChapterOption from '../components/admin/ChapterOption.vue'
import type { Syllabus } from '../models/syllabus.model'
import type { Organisation } from '../models/organisation.model'
import type { Qualification } from '../models/qualification.model'
import type { Chapter } from '../models/chapter.model'
import type { Question } from '../models/question.model'
import type { PracticeQuickRequest, PracticeGradeRequest, PracticeGradeResponse, PracticeSubmissionAnswer, PracticePartAnswer } from '../models/practice.model'
import {
  QUESTION_TYPE_SINGLE_CHOICE,
  QUESTION_TYPE_MULTIPLE_CHOICE,
  QUESTION_TYPE_TRUE_FALSE,
  QUESTION_TYPE_GAP_FILLING,
  QUESTION_TYPE_SHORT_ANSWER,
  QUESTION_TYPE_NAMES
} from '../models/question.model'

const userStore = useUserStore()
const organisations = ref<Organisation[]>([])
const qualifications = ref<Qualification[]>([])
const syllabuses = ref<Syllabus[]>([])
const selectedOrganisationId = ref<number | ''>('')
const selectedQualificationId = ref<number | ''>('')
const syllabusId = ref<number | ''>('')
const questionCount = ref<number>(5)
const questionIds = ref<number[]>([])
const questions = ref<Record<number, Question>>({})
const currentIndex = ref(0)
const currentQuestion = computed(() => questions.value[questionIds.value[currentIndex.value]] || { stem: '', questionContents: [] })
const answers = reactive<Record<string, string>>({})
const answersMulti = reactive<Record<string, string[]>>({})
const result = ref<PracticeGradeResponse | null>(null)
const sidebarVisible = ref(true)
const resultItem = computed(() => {
  if (!result.value || !result.value.results) return null
  return result.value.results.find(r => r.questionId === questionIds.value[currentIndex.value]) || null
})
const currentQuestionContents = computed(() => {
  const q = questions.value[questionIds.value[currentIndex.value]]
  return q && Array.isArray(q.questionContents) ? q.questionContents : []
})

// 新增章节筛选相关
const chapterTree = ref<Chapter[]>([])
const selectedChapterIds = ref<number[]>([])
const showChapterSelector = ref(false)

const resetAll = () => {
  questions.value = []
  result.value = null
  Object.keys(answers).forEach(k => delete answers[k])
  Object.keys(answersMulti).forEach(k => delete answersMulti[k])
  selectedOrganisationId.value = ''
  selectedQualificationId.value = ''
  syllabusId.value = ''
  questionCount.value = 5
  qualifications.value = []
  syllabuses.value = []
  chapterTree.value = []
  selectedChapterIds.value = []
  showChapterSelector.value = false
}

const fetchOrganisations = async () => {
  const res = await organisationService.getOrganisations({ pageIndex: 1, pageSize: 100 })
  organisations.value = res.data?.list || []
}

const fetchQualifications = async () => {
  if (!selectedOrganisationId.value) {
    qualifications.value = []
    return
  }
  const res = await qualificationService.getQualifications({ organisationId: selectedOrganisationId.value, pageIndex: 1, pageSize: 100 })
  qualifications.value = res.data?.list || []
}

const fetchSyllabuses = async () => {
  if (!selectedQualificationId.value) {
    syllabuses.value = []
    return
  }
  const res = await syllabusService.getSyllabuses({ qualificationId: selectedQualificationId.value, pageIndex: 1, pageSize: 100 })
  syllabuses.value = res.data?.list || []
}

const fetchChapterTree = async () => {
  chapterTree.value = []
  selectedChapterIds.value = []
  if (!syllabusId.value) return
  try {
    const res = await chapterService.getChapterTree(Number(syllabusId.value))
    chapterTree.value = res.data || []
  } catch (e) {
    chapterTree.value = []
  }
}

const onOrganisationChange = async () => {
  selectedQualificationId.value = ''
  syllabusId.value = ''
  qualifications.value = []
  syllabuses.value = []
  chapterTree.value = []
  selectedChapterIds.value = []
  showChapterSelector.value = false
  await fetchQualifications()
}

const onQualificationChange = async () => {
  syllabusId.value = ''
  syllabuses.value = []
  chapterTree.value = []
  selectedChapterIds.value = []
  showChapterSelector.value = false
  await fetchSyllabuses()
}

watch(syllabusId, async (newVal) => {
  chapterTree.value = []
  selectedChapterIds.value = []
  showChapterSelector.value = false
  if (newVal) {
    await fetchChapterTree()
  }
})

onMounted(async () => {
  await fetchOrganisations()
})

const startPractice = async () => {
  if (!syllabusId.value || !questionCount.value) return
  const req: PracticeQuickRequest = {
    syllabusId: Number(syllabusId.value),
    questionCount: questionCount.value,
    chapterIds: selectedChapterIds.value.length > 0 ? selectedChapterIds.value : undefined
  }
  const res = await practiceService.quickPractice(req)
  questionIds.value = res.data?.list || []
  questions.value = {}
  currentIndex.value = 0
  result.value = null
  Object.keys(answers).forEach(k => delete answers[k])
  Object.keys(answersMulti).forEach(k => delete answersMulti[k])
  // 加载第一页题目
  if (questionIds.value.length > 0) {
    await loadQuestionByIndex(0)
  }
}

const loadQuestionByIndex = async (idx: number) => {
  const id = questionIds.value[idx]
  if (!id) return
  if (!questions.value[id]) {
    const res = await (await import('../services/questionService')).default.getQuestionById(id)
    questions.value[id] = res.data
  }
  currentIndex.value = idx
}

const jumpTo = async (idx: number) => {
  await loadQuestionByIndex(idx)
}

const submitAnswers = async () => {
  // Build answers grouped by question, each with answers per part
  const submissionAnswers: PracticeSubmissionAnswer[] = [];
  questionIds.value.forEach(id => {
    const q = questions.value[id]
    if (q && q.questionContents && q.questionContents.length > 0) {
      const partAnswers: PracticePartAnswer[] = [];
      q.questionContents.forEach((content, cidx) => {
        let answer: string | string[] = '';
        if (content.questionTypeId === QUESTION_TYPE_MULTIPLE_CHOICE) {
          answer = answersMulti[id + '-' + cidx] || [];
        } else {
          answer = answers[id + '-' + cidx] || '';
        }
        partAnswers.push({
          questionContentId: cidx, // If content.id exists, use content.id instead
          answer: Array.isArray(answer) ? JSON.stringify(answer) : answer
        });
      });
      submissionAnswers.push({
        questionId: id,
        answers: partAnswers
      });
    }
  });
  const req: PracticeGradeRequest = submissionAnswers;
  const res = await practiceService.gradePractice(req);
  result.value = res.data;
}
</script>
