<template>
  <!-- Floating feedback button -->
  <div class="fixed bottom-6 right-6 z-50">
    <button
      v-if="!isOpen"
      @click="open"
      class="flex items-center justify-center w-14 h-14 rounded-full bg-indigo-600 text-white shadow-lg hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition-colors"
      :title="$t('feedback.buttonTitle')"
    >
      <!-- Chat bubble icon -->
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
          d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
      </svg>
    </button>

    <!-- Modal -->
    <transition name="slide-up">
      <div
        v-if="isOpen"
        class="bg-white rounded-2xl shadow-2xl w-80 sm:w-96 overflow-hidden"
      >
        <!-- Header -->
        <div class="flex items-center justify-between bg-indigo-600 px-4 py-3">
          <h2 class="text-white font-semibold text-sm">{{ $t('feedback.title') }}</h2>
          <button @click="close" class="text-indigo-200 hover:text-white focus:outline-none">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Success state -->
        <div v-if="submitted" class="px-4 py-8 text-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-green-500 mx-auto mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p class="text-gray-700 font-medium">{{ $t('feedback.thankYou') }}</p>
          <p class="text-gray-500 text-sm mt-1">{{ $t('feedback.successMessage') }}</p>
          <button @click="reset" class="mt-4 text-sm text-indigo-600 hover:underline">
            {{ $t('feedback.submitAnother') }}
          </button>
        </div>

        <!-- Form -->
        <form v-else @submit.prevent="submit" class="px-4 py-4 space-y-3">
          <!-- Type -->
          <div>
            <label class="block text-xs font-medium text-gray-700 mb-1">{{ $t('feedback.typeLabel') }}</label>
            <div class="flex flex-wrap gap-2">
              <button
                v-for="t in feedbackTypes"
                :key="t.value"
                type="button"
                @click="form.type = t.value"
                :class="[
                  'px-3 py-1 rounded-full text-xs font-medium border transition-colors focus:outline-none',
                  form.type === t.value
                    ? 'bg-indigo-600 text-white border-indigo-600'
                    : 'bg-white text-gray-700 border-gray-300 hover:border-indigo-400'
                ]"
              >{{ t.label }}</button>
            </div>
          </div>

          <!-- Rating -->
          <div>
            <label class="block text-xs font-medium text-gray-700 mb-1">{{ $t('feedback.ratingLabel') }}</label>
            <div class="flex gap-1">
              <button
                v-for="star in 5"
                :key="star"
                type="button"
                @click="form.rating = star"
                :class="[
                  'text-xl focus:outline-none transition-colors',
                  star <= form.rating ? 'text-yellow-400' : 'text-gray-300'
                ]"
              >★</button>
            </div>
          </div>

          <!-- Content -->
          <div>
            <label class="block text-xs font-medium text-gray-700 mb-1">{{ $t('feedback.contentLabel') }}</label>
            <textarea
              v-model="form.content"
              rows="3"
              required
              :placeholder="$t('feedback.contentPlaceholder')"
              class="w-full px-3 py-2 text-sm border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-indigo-500 focus:border-indigo-500 resize-none"
            />
          </div>

          <!-- Consent -->
          <div class="flex items-start gap-2">
            <input
              id="fb-consent"
              type="checkbox"
              v-model="form.consentGiven"
              class="mt-0.5 h-4 w-4 text-indigo-600 border-gray-300 rounded"
            />
            <label for="fb-consent" class="text-xs text-gray-600">{{ $t('feedback.consentText') }}</label>
          </div>

          <!-- Error -->
          <p v-if="error" class="text-red-600 text-xs">{{ error }}</p>

          <!-- Submit -->
          <button
            type="submit"
            :disabled="!form.consentGiven || submitting"
            class="w-full py-2 bg-indigo-600 text-white text-sm font-medium rounded-md hover:bg-indigo-700 disabled:opacity-50 disabled:cursor-not-allowed focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition-colors"
          >
            <span v-if="submitting">{{ $t('feedback.submitting') }}</span>
            <span v-else>{{ $t('feedback.submitButton') }}</span>
          </button>
        </form>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import feedbackService from '../services/feedbackService'
import { FEEDBACK_TYPE_GENERAL, FEEDBACK_TYPE_BUG, FEEDBACK_TYPE_FEATURE, FEEDBACK_TYPE_PRAISE } from '../models/feedback.model'

const { t } = useI18n()

const isOpen = ref(false)
const submitted = ref(false)
const submitting = ref(false)
const error = ref('')

const form = reactive({
  type: FEEDBACK_TYPE_GENERAL,
  content: '',
  rating: 0,
  consentGiven: false
})

const feedbackTypes = computed(() => [
  { value: FEEDBACK_TYPE_GENERAL, label: t('feedback.typeGeneral') },
  { value: FEEDBACK_TYPE_BUG, label: t('feedback.typeBug') },
  { value: FEEDBACK_TYPE_FEATURE, label: t('feedback.typeFeature') },
  { value: FEEDBACK_TYPE_PRAISE, label: t('feedback.typePraise') }
])

function open() {
  isOpen.value = true
  error.value = ''
}

function close() {
  isOpen.value = false
}

function reset() {
  submitted.value = false
  form.type = FEEDBACK_TYPE_GENERAL
  form.content = ''
  form.rating = 0
  form.consentGiven = false
  error.value = ''
}

async function submit() {
  if (!form.consentGiven) return
  error.value = ''
  submitting.value = true
  try {
    const res = await feedbackService.submit({
      type: form.type,
      content: form.content,
      rating: form.rating,
      pageContext: window.location.pathname,
      consentGiven: form.consentGiven
    })
    if (res.code === 0) {
      submitted.value = true
    } else {
      error.value = res.message || t('feedback.submitError')
    }
  } catch {
    error.value = t('feedback.submitError')
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
.slide-up-enter-active,
.slide-up-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}
.slide-up-enter-from,
.slide-up-leave-to {
  opacity: 0;
  transform: translateY(16px);
}
</style>
