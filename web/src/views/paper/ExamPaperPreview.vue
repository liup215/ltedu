<template>
  <!-- Previous template content remains the same until paper header -->
  <div class="p-6 bg-gray-50 min-h-screen">
    <!-- Header and Loading States remain the same -->
    <header class="mb-6">
      <div class="flex justify-between items-center">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ $t('examPaperPreview.title') }}</h1>
          <p class="mt-1 text-sm text-gray-500">{{ $t('examPaperPreview.subtitle') }}</p>
        </div>
        <div class="flex gap-3">
          <button 
            @click="handleExportPaper"
            class="px-4 py-2 text-white bg-blue-600 rounded-md hover:bg-blue-700 flex items-center gap-2"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
            </svg>
            {{ $t('examPaperPreview.exportPaper') }}
          </button>
          <router-link 
            :to="`/paper/exam/teacher`"
            class="px-4 py-2 text-gray-700 bg-white border rounded-md hover:bg-gray-50"
          >
            {{ $t('examPaperPreview.backToList') }}
          </router-link>
        </div>
      </div>
    </header>

    <!-- Loading State -->
    <div v-if="isLoading" class="flex justify-center items-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
      <span class="ml-3 text-gray-600">{{ $t('examPaperPreview.loading') }}</span>
    </div>

    <!-- Error State -->
    <div v-else-if="errorMessage" class="text-center py-12">
      <div class="text-red-600 bg-red-50 p-4 rounded-md">
        {{ errorMessage }}
      </div>
    </div>

    <!-- Content -->
    <div v-else-if="paper" class="bg-white rounded-lg shadow-md">
      <!-- Paper Header -->
      <div class="p-6 border-b border-gray-200">
        <h2 class="text-2xl font-bold text-center mb-4">{{ paper.name }}</h2>
        <div class="flex justify-center items-center gap-4 text-sm text-gray-500">
          <span>{{ $t('examPaperPreview.year') }}: {{ paper.year }}</span>
          <span>{{ $t('examPaperPreview.totalScore') }}: {{ totalPaperScore }}</span>
          <span>{{ $t('examPaperPreview.questions') }}: {{ paper.questionIds?.length || 0 }}</span>
        </div>
        <div v-if="syllabusInfo" class="mt-3 text-sm text-gray-600 text-center">
          {{ syllabusInfo }}
        </div>
      </div>

      <!-- Questions -->
      <div class="p-6">
        <div v-for="(q, index) in questions" :key="q.id" class="mb-8">
          <div class="bg-white rounded-lg border border-gray-300 p-4">
            <!-- Question Header -->
            <div class="flex justify-between items-start mb-4">
              <span class="text-lg font-semibold">{{ $t('examPaperPreview.question') }} {{ index + 1 }}</span>
              <div class="flex gap-2">
                <span class="px-2 py-1 bg-blue-100 text-blue-800 rounded-full text-xs font-medium">
                  {{ $t('examPaperPreview.score') }}: {{ q.totalScore || getQuestionTotalScore(q) || '-' }}
                </span>
                <span class="px-2 py-1 bg-yellow-100 text-yellow-800 rounded-full text-xs font-medium">
                  {{ DIFFICULTY_NAMES[q.difficult as keyof typeof DIFFICULTY_NAMES] || $t('examPaperPreview.unknown') }}
                </span>
              </div>
            </div>

            <!-- Question Content -->
            <div class="prose max-w-none">
              <QuillEditor v-model="q.stem" readOnly height="100%"></QuillEditor>
            </div>

            <!-- Question Parts -->
            <div v-if="q.questionContents && q.questionContents.length > 0" class="mt-4 space-y-4">
              <div v-for="content in q.questionContents" :key="content.id" class="ml-4">
                <div class="font-medium mb-2">
                  {{ content.partLabel }}{{ content.subpartLabel ? '.' + content.subpartLabel : '' }}
                  <span class="text-sm text-gray-500">({{ content.score }} pts)</span>
                </div>
                <!-- Display different question types -->
                <div class="prose max-w-none mt-2">
                  <div v-if="content.questionTypeId === 1 && content.singleChoice" class="space-y-2">
                    <div 
                      v-for="option in content.singleChoice.options" 
                      :key="option.prefix" 
                      :class="[
                        'pl-4 flex items-start',
                        content.singleChoice?.answer === option.prefix ? 'text-blue-600' : ''
                      ]"
                    >
                      <span class="font-medium mr-2">{{ option.prefix }}.</span>
                      <span v-html="option.content"></span>
                      <svg v-if="content.singleChoice?.answer === option.prefix" class="w-4 h-4 ml-2 mt-1 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                      </svg>
                    </div>
                  </div>

                  <div v-else-if="content.questionTypeId === 2 && content.multipleChoice" class="space-y-2">
                    <div 
                      v-for="option in content.multipleChoice.options" 
                      :key="option.prefix" 
                      :class="[
                        'pl-4 flex items-start',
                        content.multipleChoice?.answer?.includes(option.prefix) ? 'text-blue-600' : ''
                      ]"
                    >
                      <span class="font-medium mr-2">{{ option.prefix }}.</span>
                      <span v-html="option.content"></span>
                      <svg v-if="content.multipleChoice?.answer?.includes(option.prefix)" class="w-4 h-4 ml-2 mt-1 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                      </svg>
                    </div>
                  </div>

                  <div v-else-if="content.questionTypeId === 3 && content.trueOrFalse" class="pl-4">
                    <div class="space-y-2">
                      <div 
                        :class="[
                          'pl-4 flex items-start',
                          Boolean(content.trueOrFalse?.answer) ? 'text-blue-600' : ''
                        ]"
                      >
                        <span class="mr-2">True</span>
                        <svg v-if="Boolean(content.trueOrFalse?.answer)" class="w-4 h-4 ml-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                        </svg>
                      </div>
                      <div 
                        :class="[
                          'pl-4 flex items-start',
                          content.trueOrFalse?.answer === 0 ? 'text-blue-600' : ''
                        ]"
                      >
                        <span class="mr-2">False</span>
                        <svg v-if="content.trueOrFalse?.answer === 0" class="w-4 h-4 ml-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                        </svg>
                      </div>
                    </div>
                  </div>

                  <div v-else-if="content.questionTypeId === 4" class="pl-4">
                    <div class="text-gray-600">{{ $t('examPaperPreview.gapFilling') }}</div>
                    <div class="mt-2 text-blue-600 font-medium">
                      <div v-if="Array.isArray(content.gapFilling?.answer)">
                        <div>{{ $t('examPaperPreview.answers') }}</div>
                        <ul class="list-decimal pl-5 mt-1 space-y-1">
                          <li v-for="(ans, idx) in content.gapFilling.answer" :key="idx">{{ ans }}</li>
                        </ul>
                      </div>
                      <div v-else>
                        {{ $t('examPaperPreview.answer') }}: {{ content.gapFilling?.answer }}
                      </div>
                    </div>
                  </div>

<div v-else-if="content.questionTypeId === 5" class="pl-4">
  <div class="text-gray-600">{{ $t('examPaperPreview.shortAnswer') }}</div>
  <div class="mt-2 space-y-2">
    <div class="text-blue-600 font-medium">{{ $t('examPaperPreview.answer') }}</div>
    <div class="pl-4 border-l-2 border-blue-200">
<QuillEditor :model-value="content.shortAnswer?.answer ?? ''" readOnly height="100%"></QuillEditor>
    </div>
  </div>
</div>

                  <div v-else class="pl-4">
                    <div class="text-gray-600">{{ $t('examPaperPreview.otherType') }}</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { examPaperService } from '../../services/examPaperService'
import questionService from '../../services/questionService'
import type { ExamPaper } from '../../models/examPaper.model'
import type { Question } from '../../models/question.model'
import { DIFFICULTY_NAMES } from '../../models/question.model'
import { exportExamPaperToDocx } from '../../utils/exportDocx'
import QuillEditor from '../../components/QuillEditor/index.vue'

const route = useRoute()

const paper = ref<ExamPaper | null>(null)
const questions = ref<Question[]>([])
const isLoading = ref(true)
const errorMessage = ref<string | null>(null)

const getQuestionTotalScore = (question: Question) => {
  if (!question.questionContents?.length) return 0;
  return question.questionContents.reduce((sum, content) => sum + (content.score || 0), 0);
}

const totalPaperScore = computed(() => 
  questions.value.reduce((sum, q) => sum + (q.totalScore || getQuestionTotalScore(q)), 0)
)

const syllabusInfo = computed(() => {
  const syl = paper.value?.syllabus;
  if (!syl) return '';
  const parts = [
    syl.qualification?.organisation?.name,
    syl.qualification?.name,
    syl.name
  ].filter(Boolean);
  return parts.join(' - ') + (syl.code ? ` (${syl.code})` : '');
})

const fetchPaper = async () => {
  isLoading.value = true
  errorMessage.value = null

  try {
    const paperId = Number(route.params.id)
    if (!paperId) {
      throw new Error('Invalid paper ID')
    }

    // 获取试卷信息
    const paperRes = await examPaperService.getExamPaperById({ id: paperId })
    if (!paperRes.data) {
      throw new Error('Failed to load paper')
    }
    paper.value = paperRes.data

    // 获取试卷中的所有题目
    if (paper.value.questionIds && paper.value.questionIds.length > 0) {
      const questionsPromises = paper.value.questionIds.map(qId => 
        questionService.getQuestionById(qId)
      )
      const responses = await Promise.all(questionsPromises)
      questions.value = responses
        .filter(res => res.data)
        .map(res => res.data)
        .sort((a, b) => {
          const indexA = paper.value?.questionIds.indexOf(a.id) ?? 0
          const indexB = paper.value?.questionIds.indexOf(b.id) ?? 0
          return indexA - indexB
        })
    }

  } catch (error: any) {
    errorMessage.value = error.message || '加载试卷失败'
  } finally {
    isLoading.value = false
  }
}

const handleExportPaper = async () => {
  if (paper.value && Array.isArray(questions.value) && questions.value.length > 0) {
    // Attach questions to paper for export utility
    const exportPaper = { ...paper.value, questions: questions.value };
    await exportExamPaperToDocx(exportPaper);
  }
}

onMounted(fetchPaper)
</script>
