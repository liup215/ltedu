<template>
  <div class="p-6 w-full min-h-screen bg-gray-50 flex flex-col md:flex-row gap-6">
    <!-- Side Bar: Organisation, Qualification, Syllabus, Chapters -->
    <aside class="w-full md:w-72 flex-shrink-0 mb-6 md:mb-0">
      <h2 class="text-lg font-semibold text-gray-700 mb-4">{{ $t('examPaperForm.paperStructure') }}</h2>
      <!-- Organisation Select -->
      <div class="mb-4">
        <label class="block text-gray-600 mb-1">{{ $t('examPaperForm.organisation') }}</label>
        <select v-model="selectedOrganisationId" @change="onOrganisationChange"
          class="w-full p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500">
          <option value="" disabled>{{ $t('examPaperForm.selectOrganisation') }}</option>
          <option v-for="org in organisations" :key="org.id" :value="org.id">
            {{ org.name }}
          </option>
        </select>
      </div>
      <!-- Qualification Select -->
      <div class="mb-4">
        <label class="block text-gray-600 mb-1">{{ $t('examPaperForm.qualification') }}</label>
        <select v-model="selectedQualificationId" @change="onQualificationChange"
          class="w-full p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500">
          <option value="" disabled>{{ $t('examPaperForm.selectQualification') }}</option>
          <option v-for="q in qualifications" :key="q.id" :value="q.id">
            {{ q.name }}
          </option>
        </select>
      </div>
      <!-- Syllabus Select -->
      <div class="mb-4">
        <label class="block text-gray-600 mb-1">{{ $t('examPaperForm.syllabus') }}</label>
        <select v-model="form.syllabusId" @change="onSyllabusChange"
          class="w-full p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500">
          <option value="" disabled>{{ $t('examPaperForm.selectSyllabus') }}</option>
          <option v-for="syllabus in syllabuses" :key="syllabus.id" :value="syllabus.id">
            {{ syllabus.name }} ({{ syllabus.code }})
          </option>
        </select>
      </div>
      <!-- Chapters Selector (QuickPractice style) -->
      <div v-if="form.syllabusId">
        <label class="block text-gray-600 mb-1">{{ $t('examPaperForm.chapters') }}</label>
<div class="relative min-w-[200px] max-w-xl">
  <button
    type="button"
    @click="showChapterSelector = !showChapterSelector"
    :disabled="!chapterTree.length"
    class="flex justify-between items-center w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 disabled:bg-gray-100 disabled:cursor-not-allowed"
  >
    <span class="text-sm truncate">
      {{ selectedChapterIds.length ? $t('examPaperForm.chaptersSelected', { count: selectedChapterIds.length }) : $t('examPaperForm.allChapters') }}
    </span>
    <svg class="w-5 h-5 text-gray-400" viewBox="0 0 20 20" fill="currentColor">
      <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
    </svg>
  </button>
  <div v-if="showChapterSelector" class="absolute z-10 w-full lg:w-96 mt-1 bg-white rounded-md shadow-lg border border-gray-200">
    <div class="max-h-96 overflow-y-auto p-2">
      <div class="flex items-center justify-between p-2 border-b">
        <span class="text-sm font-medium text-gray-900">{{ $t('examPaperForm.selectChapters') }}</span>
        <button 
          type="button"
          @click="selectedChapterIds = []; showChapterSelector = false"
          class="text-sm text-gray-500 hover:text-gray-700"
        >
          {{ $t('examPaperForm.clear') }}
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
          @update:selected="val => { selectedChapterIds = val; fetchQuestions(); }"
        />
      </div>
    </div>
  </div>
</div>
      </div>
    </aside>

    <!-- Main Content -->
    <div class="flex-1">
      <!-- Header -->
      <header class="mb-6">
        <h1 class="text-3xl font-bold text-gray-900">{{ isEdit ? $t('examPaperForm.editTitle') : $t('examPaperForm.createTitle') }}</h1>
        <p class="mt-1 text-sm text-gray-500">{{ $t('examPaperForm.subtitle') }}</p>
        <div class="mt-2 bg-yellow-50 border-l-4 border-yellow-400 text-yellow-700 p-3 rounded">
          <span>
            {{ $t('examPaperForm.knowledgePointNotice') }}
          </span>
        </div>
      </header>

      <!-- Paper Form -->
      <form @submit.prevent="handleSubmit" class="mb-6 space-y-4 max-w-2xl">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('examPaperForm.paperName') }}</label>
          <input v-model="form.name" type="text" required :placeholder="$t('examPaperForm.paperNamePlaceholder')" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500 sm:text-sm" />
        </div>
        <div class="flex flex-row gap-2 items-center">
          <div class="flex-1"></div>
          <div class="flex gap-2">
            <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition">{{ $t('examPaperForm.save') }}</button>
            <button type="button" class="px-4 py-2 bg-gray-200 text-gray-700 rounded-md hover:bg-gray-300 transition" @click="handleCancel">{{ $t('examPaperForm.cancel') }}</button>
          </div>
        </div>
      </form>

      <!-- Filters -->
      <div class="mb-6 flex flex-col lg:flex-row gap-4">
        <input v-model="questionFilter.stem" @input="fetchQuestions" :placeholder="$t('examPaperForm.searchStem')" class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full lg:w-64" />
        <input v-model="questionFilter.paperName" @input="onPaperNameInput" :placeholder="$t('examPaperForm.searchByPaperName')" class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full lg:w-64" />
        <select v-model="questionFilter.difficult" @change="fetchQuestions"
          class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full lg:w-32">
          <option value="">{{ $t('examPaperForm.allDifficulty') }}</option>
          <option value="1">{{ $t('examPaperForm.easy') }}</option>
          <option value="2">{{ $t('examPaperForm.medium') }}</option>
          <option value="3">{{ $t('examPaperForm.hard') }}</option>
        </select>
      </div>

      <!-- Questions Card (full width, like QuestionManagement) -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-4">
        <div class="flex justify-between items-center mb-4">
          <div class="text-sm text-gray-600">
            {{ $t('examPaperForm.showingQuestions', { count: questions.length }) }}
          </div>
          <!-- Floating selected questions icon (bottom right on mobile/desktop) -->
          <button @click="showSelected = !showSelected" class="relative flex items-center px-3 py-2 bg-blue-600 text-white rounded-full shadow hover:bg-blue-700 transition">
            <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
            </svg>
            {{ $t('examPaperForm.selectedQuestions', { count: form.questionIds.length }) }}
          </button>
        </div>
        <div v-if="questions.length === 0" class="text-center text-gray-400 py-8">
          {{ $t('examPaperForm.noQuestions') }}
        </div>
        <div v-else class="grid grid-cols-1 gap-4">
          <div v-for="q in questions" :key="q.id" class="bg-white rounded-lg shadow-md border border-gray-200 hover:shadow-lg transition-shadow duration-200">
            <!-- Card Header -->
            <div class="p-4 border-b border-gray-200">
              <div class="flex justify-between items-start">
                <div class="flex-1">
                  <div class="flex items-center space-x-2 mb-2">
                    <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                      {{ $t('examPaperForm.id') }}: {{ q.id }}
                    </span>
                    <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                      {{ $t('examPaperForm.totalScore') }}: {{ q.totalScore || 0 }}
                    </span>
                    <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-yellow-100 text-yellow-800">
                      {{ DIFFICULTY_NAMES[q.difficult as keyof typeof DIFFICULTY_NAMES] || $t('examPaperForm.unknown') }}
                    </span>
                  </div>
                  
                </div>
              </div>
            </div>

            <!-- Card Content -->
            <div class="p-4">
              <!-- Question Stem -->
              <div class="mb-3">
                <h4 class="text-sm font-medium text-gray-700 mb-1">{{ $t('examPaperForm.question') }}</h4>
                <!-- <div class="text-sm font-medium" v-html="q.stem"></div> -->
                <QuillEditor v-model="q.stem" :readOnly="true" height="100%" placeholder=""></QuillEditor>
              </div>

              <!-- Question Parts Tabs (read-only, like QuestionManagement.vue) -->
              <div v-if="q.questionContents && q.questionContents.length > 0" class="mb-3">
                <h4 class="text-sm font-medium text-gray-700 mb-2">
                  {{ $t('examPaperForm.parts') }} ({{ $t('examPaperForm.totalScore') }}: {{ q.totalScore }})
                </h4>
                <div class="border border-gray-300 rounded-lg overflow-hidden">
                  <div class="flex overflow-x-auto bg-gray-50 border-b border-gray-300">
                    <button
                      v-for="(content, index) in q.questionContents"
                      :key="index"
                      @click="() => setActiveTab(q.id, index)"
                      :class="[
                        'px-4 py-2 text-sm font-medium border-r border-gray-200 whitespace-nowrap transition',
                        getActiveTab(q.id) === index
                          ? 'text-blue-600 bg-white border-b-2 border-b-blue-500 -mb-px'
                          : 'text-gray-700 hover:text-blue-600 hover:bg-gray-100'
                      ]"
                      type="button"
                    >
                      {{ content.partLabel }}{{ content.subpartLabel ? '.' + content.subpartLabel : '' }}
                      <span class="ml-1 text-xs text-gray-500">({{ content.score }} pts)</span>
                    </button>
                  </div>
                  <div v-if="q.questionContents[getActiveTab(q.id) || 0]" class="p-3 bg-white">
                    <div class="mb-2 flex items-center gap-2">
                      <span class="text-xs font-medium text-gray-500">
                        {{ $t('examPaperForm.type') }}: {{ QUESTION_TYPE_NAMES[
                          q.questionContents[getActiveTab(q.id) || 0].questionTypeId as keyof typeof QUESTION_TYPE_NAMES
                        ] || $t('examPaperForm.unknown') }}
                      </span>
                      <span class="text-xs text-gray-400">|</span>
                      <span class="text-xs text-gray-500">
                        {{ $t('examPaperForm.score') }}: {{ q.questionContents[getActiveTab(q.id) || 0].score }} pts
                      </span>
                    </div>
                    <div class="mt-3 text-sm text-gray-600">
<!-- Single Choice -->
<div v-if="q.questionContents[getActiveTab(q.id) || 0].questionTypeId === QUESTION_TYPE_SINGLE_CHOICE" class="space-y-2">
  <div v-for="option in q.questionContents[getActiveTab(q.id) || 0].singleChoice?.options" 
       :key="option.prefix" 
       class="flex gap-2">
    <span class="text-gray-500">{{ option.prefix }}.</span>
    <span v-html="option.content"></span>
  </div>
  <div class="mt-2">
    <span class="font-semibold">{{ $t('examPaperForm.correctAnswer') }}</span>
    <span class="ml-2 text-blue-600">
      {{ q.questionContents[getActiveTab(q.id) || 0].singleChoice?.answer }}
    </span>
  </div>
</div>
<!-- Multiple Choice -->
<div v-else-if="q.questionContents[getActiveTab(q.id) || 0].questionTypeId === QUESTION_TYPE_MULTIPLE_CHOICE" class="space-y-2">
  <div v-for="option in q.questionContents[getActiveTab(q.id) || 0].multipleChoice?.options" 
       :key="option.prefix" 
       class="flex gap-2">
    <span class="text-gray-500">{{ option.prefix }}.</span>
    <span v-html="option.content"></span>
  </div>
  <div class="mt-2">
    <span class="font-semibold">Correct Answer:</span>
    <span class="ml-2 text-blue-600">
      {{ (q.questionContents[getActiveTab(q.id) || 0].multipleChoice?.answer || []).join(', ') }}
    </span>
  </div>
</div>
<!-- True/False -->
<div v-else-if="q.questionContents[getActiveTab(q.id) || 0].questionTypeId === QUESTION_TYPE_TRUE_FALSE" class="space-y-2">
  <div>
    <span class="font-semibold">Correct Answer:</span>
    <span class="ml-2">
      {{
        String(q.questionContents[getActiveTab(q.id) || 0].trueOrFalse?.answer).toLowerCase() === 'true'
          ? 'True'
          : 'False'
      }}
    </span>
  </div>
</div>
<!-- Gap Filling -->
<div v-else-if="q.questionContents[getActiveTab(q.id) || 0].questionTypeId === QUESTION_TYPE_GAP_FILLING" class="space-y-2">
  <div v-for="(ans, gidx) in q.questionContents[getActiveTab(q.id) || 0].gapFilling?.answer" :key="gidx">
    <span class="font-semibold">{{ $t('examPaperForm.gap') }} {{ gidx + 1 }}:</span>
    <span class="ml-2 font-semibold">{{ $t('examPaperForm.correctAnswer') }}</span>
    <span class="ml-2">{{ ans }}</span>
  </div>
</div>
<!-- Short Answer -->
<div v-else-if="q.questionContents[getActiveTab(q.id) || 0].questionTypeId === QUESTION_TYPE_SHORT_ANSWER" class="space-y-2">
  <div>
    <span class="font-semibold">Correct Answer:</span>
    <QuillEditor
      :model-value="q.questionContents[getActiveTab(q.id) || 0].shortAnswer?.answer || ''"
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

              <div class="mb-3">
                <h4 class="text-sm font-medium text-gray-700 mb-1">Syllabus:</h4>
                <p class="text-sm text-gray-600">{{ q.syllabus?.name || '-' }}</p>
                <p class="text-xs text-gray-500">{{ q.syllabus?.code || '' }}</p>
              </div>


            </div>

            <!-- Card Footer -->
            <div class="px-4 py-3 bg-gray-50 border-t border-gray-200 flex justify-end">
              <button v-if="!form.questionIds.includes(q.id)"
                @click="addQuestion(q.id)"
                class="inline-flex items-center px-3 py-1.5 border border-transparent text-sm font-medium rounded text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              >{{ $t('examPaperForm.addToPaper') }}</button>
              <span v-else class="inline-flex items-center px-3 py-1.5 text-sm font-medium text-green-600">
                <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                {{ $t('examPaperForm.added') }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Pagination -->
      <div v-if="totalQuestions > 0" class="mt-8 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0">
        <div class="text-sm text-gray-700">
          {{ $t('examPaperForm.paginationInfo', { from: ((currentPage - 1) * pageSize) + 1, to: Math.min(currentPage * pageSize, totalQuestions), total: totalQuestions }) }}
        </div>
        <nav class="flex items-center space-x-2" v-if="totalPages > 1">
          <button 
            @click="goToPage(currentPage - 1)"
            :disabled="currentPage === 1"
            class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {{ $t('examPaperForm.previous') }}
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
            {{ $t('examPaperForm.next') }}
          </button>
        </nav>
      </div>

      <!-- Floating Selected Questions Drawer -->
      <transition name="fade">
        <div v-if="showSelected" class="fixed bottom-6 right-6 z-50 max-w-full bg-white border border-gray-200 rounded-lg shadow-lg p-4">
          <div class="flex justify-between items-center mb-2">
            <h3 class="font-semibold text-gray-700">Selected Questions ({{ form.questionIds.length }})</h3>
            <button @click="showSelected = false" class="text-gray-400 hover:text-gray-700">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
          <ul class="divide-y divide-gray-100 overflow-y-auto max-h-50%">
            <li v-for="qid in form.questionIds" :key="qid" 
                class="border-b border-gray-100 px-2 py-3 cursor-move"
                draggable="true"
                @dragstart="onDragStart($event, qid)"
                @dragover.prevent
                @dragenter.prevent
                @drop="onDrop($event, qid)">
              <div class="flex gap-2 items-start">
                <div class="flex-shrink-0 text-gray-500">#{{ qid }}</div>
                <div class="flex-1">
                  <QuillEditor
                    v-if="selectedQuestionsMap[qid]"
                    :model-value="selectedQuestionsMap[qid].stem"
                    :readOnly="true"
                    height="100%"
                    placeholder=""
                  />
                  <div v-else class="text-gray-400 text-xs">Loading...</div>
                </div>
                <button type="button" class="text-red-500 hover:underline ml-2 flex-shrink-0" @click="removeQuestion(qid)">Remove</button>
              </div>
            </li>
          </ul>
        </div>
      </transition>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { examPaperService } from '../../services/examPaperService'
import syllabusService from '../../services/syllabusService'
import questionService from '../../services/questionService'
import organisationService from '../../services/organisationService'
import qualificationService from '../../services/qualificationService'
import chapterService from '../../services/chapterService'
import type { ExamPaperCreateRequest, ExamPaperUpdateRequest } from '../../models/examPaper.model'
import type { Syllabus } from '../../models/syllabus.model'
import { 
  QUESTION_STATE_NORMAL, 
  type Question, 
  DIFFICULTY_NAMES, 
  QUESTION_TYPE_NAMES, 
  QUESTION_TYPE_SINGLE_CHOICE, 
  QUESTION_TYPE_MULTIPLE_CHOICE, 
  QUESTION_TYPE_TRUE_FALSE, 
  QUESTION_TYPE_GAP_FILLING, 
  QUESTION_TYPE_SHORT_ANSWER 
} from '../../models/question.model'
import type { Organisation } from '../../models/organisation.model'
import type { Qualification } from '../../models/qualification.model'
import type { Chapter } from '../../models/chapter.model'
import QuillEditor from '../../components/QuillEditor/index.vue'
import ChapterOption from '../../components/admin/ChapterOption.vue'

const router = useRouter()
const route = useRoute()

const isEdit = !!route.params.id
const form = reactive<ExamPaperCreateRequest & { id?: number }>({
  name: '',
  syllabusId: 0,
  questionIds: []
})

const organisations = ref<Organisation[]>([])
const qualifications = ref<Qualification[]>([])
const syllabuses = ref<Syllabus[]>([])
const chapters = ref<Chapter[]>([])
const chapterTree = ref<Chapter[]>([])
const showChapterSelector = ref(false)

const selectedOrganisationId = ref<number | ''>('')
const selectedQualificationId = ref<number | ''>('')

const questions = ref<Question[]>([])
const totalQuestions = ref(0)
const currentPage = ref(1)
const pageSize = 12
const totalPages = computed(() => Math.ceil(totalQuestions.value / pageSize))
const paginationRange = computed(() => {
  const range = []
  const maxPagesToShow = 5
  let start = Math.max(1, currentPage.value - Math.floor(maxPagesToShow / 2))
  let end = Math.min(totalPages.value, start + maxPagesToShow - 1)
  if (totalPages.value > 0 && end - start + 1 < maxPagesToShow) {
    if (currentPage.value <= Math.floor(maxPagesToShow / 2)) {
      end = Math.min(totalPages.value, maxPagesToShow)
    } else {
      start = Math.max(1, totalPages.value - maxPagesToShow + 1)
    }
  }
  for (let i = start; i <= end; i++) {
    if (i > 0) range.push(i)
  }
  return range
})
// For read-only question part tabs
const questionActiveTabMap = ref(new Map<number, number>())
const setActiveTab = (qid: number, idx: number) => {
  questionActiveTabMap.value.set(qid, idx)
}
const getActiveTab = (qid: number) => {
  return questionActiveTabMap.value.get(qid) ?? 0
}
const questionFilter = reactive({
  stem: '',
  difficult: '',
  chapterId: '',
  paperName: '',
})

const showSelected = ref(false)
const selectedChapterIds = ref<number[]>([])

// 存储所有已选qid对应的题目对象，分页切换不丢失
const selectedQuestionsMap = reactive<Record<number, Question>>({})

// 监听form.questionIds变化，自动补全selectedQuestionsMap
watch(
  () => [...form.questionIds],
  async (qids: number[]) => {
    for (const qid of qids) {
      if (!selectedQuestionsMap[qid]) {
        try {
          const res = await questionService.getQuestionById(qid)
          if (res.data) selectedQuestionsMap[qid] = res.data
        } catch (e) {
          // 填充所有Question必需字段，防止类型报错
          selectedQuestionsMap[qid] = {
            id: qid,
stem: 'Load failed',
            questionContents: [],
            syllabusId: 0,
            totalScore: 0,
            difficult: 1,
            status: 1,
            indexInPastPaper: 0
          }
        }
      }
    }
    // 清理已移除的qid
    Object.keys(selectedQuestionsMap).forEach(k => {
      if (!qids.includes(Number(k))) delete selectedQuestionsMap[Number(k)]
    })
  },
  { immediate: true }
)

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

const fetchChapters = async () => {
  if (!form.syllabusId) {
    chapters.value = []
    return
  }
  const res = await chapterService.getChapterTree(form.syllabusId)
  chapters.value = res.data || []
}

const fetchQuestions = async () => {
  const res = await questionService.getQuestions({
    stem: questionFilter.stem,
    paperName: questionFilter.paperName?.trim() || undefined,
    difficult: questionFilter.difficult ? Number(questionFilter.difficult) : undefined,
    syllabusId: form.syllabusId || undefined,
    status: QUESTION_STATE_NORMAL, // Only fetch normal questions
    pageIndex: currentPage.value,
    pageSize
  })
  questions.value = res.data?.list || []
  totalQuestions.value = res.data?.total || 0
}

const addQuestion = (id: number) => {
  if (!form.questionIds.includes(id)) {
    form.questionIds.push(id)
  }
}
const removeQuestion = (id: number) => {
  form.questionIds = form.questionIds.filter(qid => qid !== id)
}

const onDragStart = (event: DragEvent, id: number) => {
  if (event.dataTransfer) {
    event.dataTransfer.effectAllowed = 'move'
    event.dataTransfer.setData('text/plain', id.toString())
  }
}

const onDrop = (event: DragEvent, targetId: number) => {
  if (event.dataTransfer) {
    const sourceId = Number(event.dataTransfer.getData('text/plain'))
    const sourceIndex = form.questionIds.indexOf(sourceId)
    const targetIndex = form.questionIds.indexOf(targetId)
    if (sourceIndex !== -1 && targetIndex !== -1) {
      form.questionIds.splice(sourceIndex, 1)
      form.questionIds.splice(targetIndex, 0, sourceId)
    }
  }
}

const handleSubmit = async () => {
  console.log('Submitting form:', form)
  if (!form.name || !form.syllabusId || form.questionIds.length === 0) {
    window.alert('Name, syllabus and at least one question are needed!')
    return
  }
  if (isEdit && form.id) {
    const req: ExamPaperUpdateRequest = {
      id: form.id,
      name: form.name,
      syllabusId: form.syllabusId,
      questionIds: form.questionIds
    }
    await examPaperService.updateExamPaper(req)
    window.alert('Exam paper updated!')
  } else {
    const req: ExamPaperCreateRequest = {
      name: form.name,
      syllabusId: form.syllabusId,
      questionIds: form.questionIds
    }
    await examPaperService.createExamPaper(req)
    window.alert('Exam paper created!')
  }
  router.back()
}

const handleCancel = () => {
  router.back()
}

const onOrganisationChange = async () => {
  selectedQualificationId.value = ''
  form.syllabusId = 0
  chapters.value = []
  await fetchQualifications()
  syllabuses.value = []
}

const onQualificationChange = async () => {
  form.syllabusId = 0
  chapters.value = []
  await fetchSyllabuses()
}

const onSyllabusChange = async () => {
  await fetchChapters()
  // 同步章节树
  chapterTree.value = chapters.value
  await fetchQuestions()
}

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
    currentPage.value = page
    fetchQuestions()
  }
}

let paperNameDebounceTimer: number | undefined;
const onPaperNameInput = () => {
  clearTimeout(paperNameDebounceTimer);
  paperNameDebounceTimer = window.setTimeout(() => {
    currentPage.value = 1;
    fetchQuestions();
  }, 500);
};

onMounted(async () => {
  await fetchOrganisations()
  if (isEdit && route.params.id) {
    // 编辑模式下，需回填所有级联选择
    const res = await examPaperService.getExamPaperById({ id: Number(route.params.id) })
    const paper = res.data
    form.id = paper.id
    form.name = paper.name
    form.syllabusId = paper.syllabusId
    form.questionIds = Array.isArray(paper.questionIds) ? [...paper.questionIds] : []

    // 反查 syllabus -> qualification -> organisation
    const syllabusRes = await syllabusService.getSyllabusById(form.syllabusId)
    const syllabus = syllabusRes.data
    selectedQualificationId.value = syllabus.qualificationId
    const qualRes = await qualificationService.getQualifications({ id: syllabus.qualificationId })
    selectedOrganisationId.value = qualRes.data?.list?.[0]?.organisationId || ''
    await fetchQualifications()
    await fetchSyllabuses()
    await fetchChapters()
  }
  await fetchQuestions()
})
</script>
