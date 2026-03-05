<template>
  <div v-if="visible" class="w-full sm:w-72 bg-white shadow-lg rounded-lg p-4 fixed right-0 bottom-0 h-1/2 z-50 flex flex-col">
    <div class="flex justify-between items-center mb-2">
      <h2 class="text-lg font-bold">{{ $t('practiceSidebar.questions') }}</h2>
      <button @click="onClose" class="text-gray-500 hover:text-gray-700 text-xl font-bold px-2 py-1 rounded">&times;</button>
    </div>
    <div class="flex-1 overflow-y-auto">
      <ul>
        <li
          v-for="(id, idx) in questionIds"
          :key="id"
          class="mb-2 flex items-center cursor-pointer"
          :class="{
            'bg-indigo-100': idx === currentIndex,
            'border-l-4 border-indigo-500': idx === currentIndex,
            'bg-green-50': isAnswered(id),
            'bg-red-50': isWrong(id),
          }"
          @click="jumpTo(idx)"
        >
          <span class="w-8 h-8 flex items-center justify-center rounded-full font-bold mr-2"
            :class="{
              'bg-indigo-500 text-white': idx === currentIndex,
              'bg-green-500 text-white': isAnswered(id) && !isWrong(id),
              'bg-red-500 text-white': isWrong(id),
              'bg-gray-200': !isAnswered(id),
            }"
          >{{ idx + 1 }}</span>
          <span class="flex-1">{{ $t('practiceSidebar.q') }}{{ idx + 1 }}</span>
          <span v-if="isWrong(id)" class="ml-2 text-red-600 font-bold">✗</span>
          <span v-else-if="isAnswered(id)" class="ml-2 text-green-600 font-bold">✓</span>
        </li>
      </ul>
    </div>
    <div v-if="result" class="mt-4">
      <div class="font-bold">{{ $t('practiceSidebar.score') }}: {{ result.score }} / {{ result.total }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Question } from '../models/question.model'
import type { PracticeGradeResponse } from '../models/practice.model'

const props = defineProps<{
  visible: boolean
  questionIds: number[]
  questions: Record<number, Question>
  currentIndex: number
  answers: Record<string, string>
  result?: PracticeGradeResponse | null
  jumpTo: (idx: number) => void
  onClose: () => void
}>()

const isAnswered = (id: number) => {
  const q = props.questions[id]
  if (!q || !q.questionContents) return false
  return q.questionContents.some((_, cidx) => {
    const key = id + '-' + cidx
    return props.answers[key] && props.answers[key].trim() !== ''
  })
}

const isWrong = (id: number) => {
  if (!props.result) return false
  const item = props.result.results?.find((r: any) => r.questionId === id)
  if (!item) return false
  return item.subResults.some((sub: any) => sub.isCorrect === false)
}
</script>

<style scoped>
/* Custom scrollbar for sidebar */
.w-72::-webkit-scrollbar {
  width: 8px;
}
.w-72::-webkit-scrollbar-thumb {
  background: #e5e7eb;
  border-radius: 4px;
}
</style>
