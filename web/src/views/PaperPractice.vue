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
      {{ $t('paperPractice.showSidebar') }}
    </button>
    <header class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">{{ $t('paperPractice.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('paperPractice.subtitle') }}</p>
    </header>
<form @submit.prevent="startPractice" class="mb-6 flex flex-wrap gap-4 items-end w-full">
  <div>
    <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('paperPractice.organisation') }}</label>
    <select v-model="selectedOrganisationId" required class="px-3 py-2 border border-gray-300 rounded-md min-w-[180px] w-full" @change="onOrganisationChange">
      <option value="" disabled>{{ $t('paperPractice.selectOrganisation') }}</option>
      <option v-for="org in organisations" :key="org.id" :value="org.id">
        {{ org.name }}
      </option>
    </select>
  </div>
  <div>
    <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('paperPractice.qualification') }}</label>
    <select v-model="selectedQualificationId" required class="px-3 py-2 border border-gray-300 rounded-md min-w-[180px] w-full" @change="onQualificationChange">
      <option value="" disabled>{{ $t('paperPractice.selectQualification') }}</option>
      <option v-for="qual in qualifications" :key="qual.id" :value="qual.id">
        {{ qual.name }}
      </option>
    </select>
  </div>
  <div>
    <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('paperPractice.syllabus') }}</label>
    <select v-model="selectedSyllabusId" required class="px-3 py-2 border border-gray-300 rounded-md min-w-[180px] w-full" @change="onSyllabusChange">
      <option value="" disabled>{{ $t('paperPractice.selectSyllabus') }}</option>
      <option v-for="syllabus in syllabuses" :key="syllabus.id" :value="syllabus.id">
        {{ syllabus.name }} ({{ syllabus.code }})
      </option>
    </select>
  </div>
  <div>
    <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('paperPractice.year') }}</label>
    <select v-model="year" class="px-3 py-2 border border-gray-300 rounded-md min-w-[120px] w-full" @change="onYearChange">
      <option value="" disabled>{{ $t('paperPractice.selectYear') }}</option>
      <option v-for="y in years" :key="y" :value="y">{{ y }}</option>
    </select>
  </div>
  <div>
    <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('paperPractice.paperCode') }}</label>
    <select v-model="paperCode" class="px-3 py-2 border border-gray-300 rounded-md min-w-[120px] w-full" @change="onPaperCodeChange">
      <option value="" disabled>{{ $t('paperPractice.selectPaperCode') }}</option>
      <option v-for="code in paperCodes" :key="code.id" :value="code.id">{{ code.name }}</option>
    </select>
  </div>
  <div>
    <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('paperPractice.paperSeries') }}</label>
    <select v-model="paperSeries" class="px-3 py-2 border border-gray-300 rounded-md min-w-[120px] w-full" @change="onPaperSeriesChange">
      <option value="" disabled>{{ $t('paperPractice.selectPaperSeries') }}</option>
      <option v-for="series in paperSeriesList" :key="series.id" :value="series.id">{{ series.name }}</option>
    </select>
  </div>
  <div>
    <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('paperPractice.pastPaper') }}</label>
    <select v-model="selectedPastPaperId" class="px-3 py-2 border border-gray-300 rounded-md min-w-[180px] w-full">
      <option value="" disabled>{{ $t('paperPractice.selectPastPaper') }}</option>
      <option v-for="paper in pastPapers" :key="paper.id" :value="paper.id">
        {{ paper.name }} ({{ paper.year }})
      </option>
    </select>
  </div>
  <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700">{{ $t('paperPractice.start') }}</button>
  <button type="button" @click="resetAll" class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300">{{ $t('paperPractice.reset') }}</button>
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
          <button type="button" class="px-4 py-2 bg-gray-200 text-gray-800 rounded hover:bg-gray-300" :disabled="currentIndex === 0" @click="jumpTo(currentIndex - 1)">{{ $t('paperPractice.previous') }}</button>
          <span class="font-bold text-lg">Q{{ currentIndex + 1 }} / {{ questionIds.length }}</span>
          <button type="button" class="px-4 py-2 bg-gray-200 text-gray-800 rounded hover:bg-gray-300" :disabled="currentIndex === questionIds.length - 1" @click="jumpTo(currentIndex + 1)">{{ $t('paperPractice.next') }}</button>
        </div>
        <button type="submit" class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700">{{ $t('paperPractice.submitAnswers') }}</button>
      </form>
    </div>

    <div v-if="result && questionIds.length > 0" class="mt-8 bg-white rounded shadow p-6">
      <h2 class="text-xl font-bold mb-4">{{ $t('paperPractice.practiceResult') }}</h2>
      <div class="mb-2">{{ $t('paperPractice.score') }}: {{ result.score }} / {{ result.total }}</div>
      <template v-if="resultItem">
        <div class="mb-4 border-b pb-2">
          <div class="font-semibold">Q{{ currentIndex + 1 }}</div>
          <div v-for="(sub, pidx) in resultItem.subResults" :key="pidx" class="ml-4 mb-2">
            <div>{{ $t('paperPractice.part') }} {{ pidx + 1 }}</div>
            <div v-if="sub.questionType === QUESTION_TYPE_SHORT_ANSWER">
              <div class="mb-1">{{ $t('paperPractice.yourAnswer') }}</div>
              <QuillEditor :modelValue="sub.studentAnswer" :readOnly="true" height="120px" />
              <div v-if="sub.correctAnswer" class="mb-1 mt-2">{{ $t('paperPractice.correctAnswer') }}</div>
              <QuillEditor v-if="sub.correctAnswer" :modelValue="sub.correctAnswer" :readOnly="true" height="120px" />
              <div v-if="sub.modelAnswer" class="mb-1 mt-2">{{ $t('paperPractice.modelAnswer') }}</div>
              <QuillEditor v-if="sub.modelAnswer" :modelValue="sub.modelAnswer" :readOnly="true" height="120px" />
            </div>
            <div v-else>
              <div>{{ $t('paperPractice.yourAnswer') }}: {{ sub.studentAnswer }}</div>
              <div>{{ $t('paperPractice.correctAnswer') }}: {{ sub.correctAnswer }}</div>
              <div v-if="sub.isCorrect === true" class="text-green-600">{{ $t('paperPractice.correct') }}</div>
              <div v-else-if="sub.isCorrect === false" class="text-red-600">{{ $t('paperPractice.incorrect') }}</div>
              <div v-else class="text-gray-600">{{ $t('paperPractice.subjective') }}</div>
              <div v-if="sub.modelAnswer">{{ $t('paperPractice.modelAnswer') }}: {{ sub.modelAnswer }}</div>
            </div>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'

import pastPaperService from '../services/pastPaperService'
import type { PastPaperQuery, PastPaper } from '../models/pastPaper.model'

const pastPapers = ref<PastPaper[]>([])
const selectedPastPaperId = ref<number | ''>('')

function fetchPastPapers() {
  const query: PastPaperQuery = {
    syllabusId: selectedSyllabusId.value ? Number(selectedSyllabusId.value) : undefined,
    year: year.value ? Number(year.value) : undefined,
    paperCodeId: paperCode.value ? Number(paperCode.value) : undefined,
    paperSeriesId: paperSeries.value ? Number(paperSeries.value) : undefined,
    pageIndex: 1,
    pageSize: 50
  }
  pastPaperService.getPastPapers(query).then(res => {
    pastPapers.value = res.data?.list || []
  })
}

function onOrganisationChange() {
  selectedQualificationId.value = ''
  selectedSyllabusId.value = ''
  paperCode.value = ''
  paperSeries.value = ''
  qualifications.value = []
  syllabuses.value = []
  paperCodes.value = []
  paperSeriesList.value = []
  fetchQualifications()
}

function onQualificationChange() {
  selectedSyllabusId.value = ''
  paperCode.value = ''
  paperSeries.value = ''
  syllabuses.value = []
  paperCodes.value = []
  paperSeriesList.value = []
  fetchSyllabuses()
}

function onSyllabusChange() {
  paperCode.value = ''
  paperSeries.value = ''
  paperCodes.value = []
  paperSeriesList.value = []
  fetchPaperCodes()
  fetchPaperSeries()
  fetchPastPapers()
}

function onYearChange() {
  fetchPastPapers()
}

function onPaperCodeChange() {
  fetchPastPapers()
}

function onPaperSeriesChange() {
  fetchPastPapers()
}

import QuillEditor from '../components/QuillEditor/index.vue'
import PracticeSidebar from '../components/PracticeSidebar.vue'
import { practiceService } from '../services/practiceService'
import questionService from '../services/questionService'
import syllabusService from '../services/syllabusService'
import organisationService from '../services/organisationService'
import qualificationService from '../services/qualificationService'
import paperCodeService from '../services/paperCodeService'
import paperSeriesService from '../services/paperSeriesService'
import type { Syllabus } from '../models/syllabus.model'
import type { Organisation } from '../models/organisation.model'
import type { Qualification } from '../models/qualification.model'
import type { PaperCode } from '../models/paperCode.model'
import type { PaperSeries } from '../models/paperSeries.model'
import type { Question } from '../models/question.model'
import type { PracticeGradeResponse, PracticeSubmissionAnswer, PracticePartAnswer } from '../models/practice.model'
import {
  QUESTION_TYPE_SINGLE_CHOICE,
  QUESTION_TYPE_MULTIPLE_CHOICE,
  QUESTION_TYPE_TRUE_FALSE,
  QUESTION_TYPE_GAP_FILLING,
  QUESTION_TYPE_SHORT_ANSWER,
  QUESTION_TYPE_NAMES
} from '../models/question.model'

const organisations = ref<Organisation[]>([])
const qualifications = ref<Qualification[]>([])
const syllabuses = ref<Syllabus[]>([])
const paperCodes = ref<PaperCode[]>([])
const paperSeriesList = ref<PaperSeries[]>([])
const years = ref<number[]>([])

const selectedOrganisationId = ref<number | ''>('')
const selectedQualificationId = ref<number | ''>('')
const selectedSyllabusId = ref<number | ''>('')
const paperCode = ref<number | ''>('')
const paperSeries = ref<number | ''>('')
const year = ref<number | ''>('')
const questionIds = ref<number[]>([])
const questions = ref<Record<number, Question>>({})
const currentIndex = ref(0)
const currentQuestion = computed(() => questions.value[questionIds.value[currentIndex.value]] || { stem: '', questionContents: [] })
const currentQuestionContents = computed(() => {
  const q = questions.value[questionIds.value[currentIndex.value]]
  return q && Array.isArray(q.questionContents) ? q.questionContents : []
})
const answers = reactive<Record<string, string>>({})
const answersMulti = reactive<Record<string, string[]>>({})
const result = ref<PracticeGradeResponse | null>(null)
const sidebarVisible = ref(true)
const resultItem = computed(() => {
  if (!result.value || !result.value.results) return null
  return result.value.results.find(r => r.questionId === questionIds.value[currentIndex.value]) || null
})

const resetAll = () => {
  questions.value = {}
  result.value = null
  Object.keys(answers).forEach(k => delete answers[k])
  Object.keys(answersMulti).forEach(k => delete answersMulti[k])
  selectedOrganisationId.value = ''
  selectedQualificationId.value = ''
  selectedSyllabusId.value = ''
  year.value = ''
  paperCode.value = ''
  paperSeries.value = ''
  questionIds.value = []
  currentIndex.value = 0
  qualifications.value = []
  syllabuses.value = []
  paperCodes.value = []
  paperSeriesList.value = []
}

const fetchQualifications = async () => {
  if (!selectedOrganisationId.value) {
    qualifications.value = []
    return
  }
  const res = await qualificationService.getQualifications({ organisationId: Number(selectedOrganisationId.value), pageIndex: 1, pageSize: 100 })
  qualifications.value = res.data?.list || []
}

const fetchSyllabuses = async () => {
  if (!selectedQualificationId.value) {
    syllabuses.value = []
    return
  }
  const res = await syllabusService.getSyllabuses({ qualificationId: Number(selectedQualificationId.value), pageIndex: 1, pageSize: 100 })
  syllabuses.value = res.data?.list || []
}

const fetchYears = async () => {
  // Replace with actual API call if available
  years.value = [2025, 2024, 2023, 2022, 2021, 2020]
}

const fetchPaperCodes = async () => {
  if (!selectedSyllabusId.value) {
    paperCodes.value = []
    return
  }
  const res = await paperCodeService.getPaperCodeList({ syllabusId: Number(selectedSyllabusId.value), pageIndex: 1, pageSize: 100 })
  paperCodes.value = res.data?.list || []
}

const fetchPaperSeries = async () => {
  if (!selectedSyllabusId.value) {
    paperSeriesList.value = []
    return
  }
  const res = await paperSeriesService.getPaperSeriesList({ syllabusId: Number(selectedSyllabusId.value), pageIndex: 1, pageSize: 100 })
  paperSeriesList.value = res.data?.list || []
}

onMounted(async () => {
  await fetchOrganisations()
  await fetchYears()
})

const fetchOrganisations = async () => {
  const res = await organisationService.getOrganisations({ pageIndex: 1, pageSize: 100 })
  organisations.value = res.data?.list || []
}

// --- FIX: Remove getPaperId and use selectedPastPaperId directly ---
const startPractice = async () => {
  if (!selectedPastPaperId.value) return
  const res = await practiceService.paperPractice({ paperId: Number(selectedPastPaperId.value) })
  questionIds.value = res.data?.list || []
  questions.value = {}
  currentIndex.value = 0
  result.value = null
  Object.keys(answers).forEach(k => delete answers[k])
  Object.keys(answersMulti).forEach(k => delete answersMulti[k])
  if (questionIds.value.length > 0) {
    await loadQuestionByIndex(0)
  }
}

const loadQuestionByIndex = async (idx: number) => {
  const id = questionIds.value[idx]
  if (!id) return
  if (!questions.value[id]) {
    const res = await questionService.getQuestionById(id)
    questions.value[id] = res.data
  }
  currentIndex.value = idx
}

const jumpTo = async (idx: number) => {
  await loadQuestionByIndex(idx)
}

const submitAnswers = async () => {
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
          questionContentId: cidx,
          answer: Array.isArray(answer) ? JSON.stringify(answer) : answer
        });
      });
      submissionAnswers.push({
        questionId: id,
        answers: partAnswers
      });
    }
  });
  const res = await practiceService.gradePractice(submissionAnswers);
  result.value = res.data;
}
</script>
