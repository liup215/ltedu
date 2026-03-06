<template>
  <div class="flex-1 bg-gray-50 min-h-screen">
    <!-- Header -->
    <section class="bg-gradient-to-r from-indigo-600 to-purple-600 text-white py-16">
      <div class="container mx-auto px-6 text-center">
        <h1 class="text-4xl font-normal mb-4">{{ $t('help.title') }}</h1>
        <p class="text-xl text-indigo-100 mb-8">{{ $t('help.subtitle') }}</p>
        <!-- Search bar -->
        <div class="max-w-xl mx-auto relative">
          <input
            v-model="searchQuery"
            type="text"
            :placeholder="$t('help.searchPlaceholder')"
            class="w-full px-5 py-3 rounded-lg text-gray-800 text-base focus:outline-none focus:ring-2 focus:ring-indigo-300 shadow"
          />
          <svg class="absolute right-4 top-3.5 w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-4.35-4.35M17 11A6 6 0 1 1 5 11a6 6 0 0 1 12 0z" />
          </svg>
        </div>
      </div>
    </section>

    <!-- Tab Navigation -->
    <div class="bg-white border-b border-gray-200 sticky top-16 z-10">
      <div class="container mx-auto px-6">
        <nav class="flex gap-1 overflow-x-auto">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="activeTab = tab.id"
            class="px-5 py-4 text-sm font-normal whitespace-nowrap border-b-2 transition-colors"
            :class="activeTab === tab.id
              ? 'border-indigo-600 text-indigo-600'
              : 'border-transparent text-gray-600 hover:text-indigo-600 hover:border-indigo-300'"
          >
            {{ tab.label }}
          </button>
        </nav>
      </div>
    </div>

    <!-- Content -->
    <div class="container mx-auto px-6 py-10 max-w-5xl">

      <!-- Getting Started Tab -->
      <div v-if="activeTab === 'getting-started'">
        <h2 class="text-2xl font-normal text-gray-800 mb-8">{{ $t('help.gettingStarted.title') }}</h2>

        <!-- Role selector -->
        <div class="flex gap-4 mb-8">
          <button
            @click="activeRole = 'student'"
            class="px-6 py-3 rounded-lg font-normal transition"
            :class="activeRole === 'student'
              ? 'bg-indigo-600 text-white shadow'
              : 'bg-white text-gray-700 border border-gray-300 hover:bg-indigo-50'"
          >
            {{ $t('help.gettingStarted.studentGuide') }}
          </button>
          <button
            @click="activeRole = 'teacher'"
            class="px-6 py-3 rounded-lg font-normal transition"
            :class="activeRole === 'teacher'
              ? 'bg-purple-600 text-white shadow'
              : 'bg-white text-gray-700 border border-gray-300 hover:bg-purple-50'"
          >
            {{ $t('help.gettingStarted.teacherGuide') }}
          </button>
        </div>

        <!-- Student Steps -->
        <div v-if="activeRole === 'student'" class="space-y-4">
          <div
            v-for="(step, index) in studentSteps"
            :key="index"
            class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 flex gap-4"
          >
            <div class="flex-shrink-0 w-10 h-10 bg-indigo-100 rounded-full flex items-center justify-center text-indigo-600 font-normal text-lg">
              {{ index + 1 }}
            </div>
            <div>
              <h3 class="text-lg font-normal text-gray-800 mb-1">{{ step.title }}</h3>
              <p class="text-gray-600">{{ step.desc }}</p>
              <router-link
                v-if="step.link"
                :to="step.link"
                class="mt-2 inline-flex items-center text-indigo-600 hover:text-indigo-800 text-sm"
              >
                {{ step.linkLabel }} →
              </router-link>
            </div>
          </div>
        </div>

        <!-- Teacher Steps -->
        <div v-if="activeRole === 'teacher'" class="space-y-4">
          <div
            v-for="(step, index) in teacherSteps"
            :key="index"
            class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 flex gap-4"
          >
            <div class="flex-shrink-0 w-10 h-10 bg-purple-100 rounded-full flex items-center justify-center text-purple-600 font-normal text-lg">
              {{ index + 1 }}
            </div>
            <div>
              <h3 class="text-lg font-normal text-gray-800 mb-1">{{ step.title }}</h3>
              <p class="text-gray-600">{{ step.desc }}</p>
              <router-link
                v-if="step.link"
                :to="step.link"
                class="mt-2 inline-flex items-center text-purple-600 hover:text-purple-800 text-sm"
              >
                {{ step.linkLabel }} →
              </router-link>
            </div>
          </div>
        </div>
      </div>

      <!-- Features Tab -->
      <div v-if="activeTab === 'features'">
        <h2 class="text-2xl font-normal text-gray-800 mb-8">{{ $t('help.features.title') }}</h2>
        <div class="grid md:grid-cols-2 gap-6">
          <div
            v-for="feature in features"
            :key="feature.id"
            class="bg-white rounded-lg shadow-sm border border-gray-200 p-6"
          >
            <div class="flex items-start gap-4">
              <div class="flex-shrink-0 w-12 h-12 rounded-lg flex items-center justify-center text-2xl" :class="feature.bgClass">
                {{ feature.icon }}
              </div>
              <div>
                <h3 class="text-lg font-normal text-gray-800 mb-2">{{ feature.title }}</h3>
                <p class="text-gray-600 text-sm leading-relaxed">{{ feature.desc }}</p>
                <ul class="mt-3 space-y-1">
                  <li v-for="bullet in feature.bullets" :key="bullet" class="text-sm text-gray-600 flex items-start gap-2">
                    <span class="text-green-500 mt-0.5">✓</span>
                    {{ bullet }}
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- FAQ Tab -->
      <div v-if="activeTab === 'faq'">
        <h2 class="text-2xl font-normal text-gray-800 mb-8">{{ $t('help.faq.title') }}</h2>
        <div class="space-y-3">
          <div
            v-for="(item, index) in filteredFaqs"
            :key="index"
            class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden"
          >
            <button
              @click="toggleFaq(index)"
              class="w-full flex items-center justify-between px-6 py-4 text-left hover:bg-gray-50 transition"
            >
              <span class="font-normal text-gray-800">{{ item.question }}</span>
              <svg
                class="w-5 h-5 text-gray-400 transition-transform flex-shrink-0 ml-2"
                :class="{ 'rotate-180': openFaqs.has(index) }"
                fill="none" stroke="currentColor" viewBox="0 0 24 24"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
              </svg>
            </button>
            <div v-if="openFaqs.has(index)" class="px-6 pb-4 text-gray-600 text-sm leading-relaxed border-t border-gray-100 pt-3">
              {{ item.answer }}
            </div>
          </div>
          <div v-if="filteredFaqs.length === 0" class="text-center py-10 text-gray-500">
            {{ $t('help.faq.noResults') }}
          </div>
        </div>
      </div>

      <!-- API & Developer Tab -->
      <div v-if="activeTab === 'api'">
        <h2 class="text-2xl font-normal text-gray-800 mb-4">{{ $t('help.api.title') }}</h2>
        <p class="text-gray-600 mb-8">{{ $t('help.api.subtitle') }}</p>

        <div class="grid md:grid-cols-2 gap-6 mb-8">
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
            <h3 class="text-lg font-normal text-gray-800 mb-3">{{ $t('help.api.restTitle') }}</h3>
            <p class="text-gray-600 text-sm mb-4">{{ $t('help.api.restDesc') }}</p>
            <div class="bg-gray-50 rounded p-3 text-sm font-mono text-gray-700">
              <div>GET /api/v1/syllabus</div>
              <div>POST /api/v1/practice</div>
              <div>GET /api/v1/question</div>
              <div>POST /api/v1/paper</div>
            </div>
          </div>
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
            <h3 class="text-lg font-normal text-gray-800 mb-3">{{ $t('help.api.mcpTitle') }}</h3>
            <p class="text-gray-600 text-sm mb-4">{{ $t('help.api.mcpDesc') }}</p>
            <router-link
              to="/account/mcp-tokens"
              class="inline-flex items-center px-4 py-2 bg-indigo-600 text-white rounded-lg text-sm hover:bg-indigo-700 transition"
            >
              {{ $t('help.api.getMcpToken') }} →
            </router-link>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
          <h3 class="text-lg font-normal text-gray-800 mb-3">{{ $t('help.api.cliTitle') }}</h3>
          <p class="text-gray-600 text-sm mb-4">{{ $t('help.api.cliDesc') }}</p>
          <div class="bg-gray-900 rounded-lg p-4 text-sm font-mono text-green-400 space-y-1">
            <div><span class="text-gray-500"># {{ $t('help.api.cliInstall') }}</span></div>
            <div>go install github.com/liup215/ltedu/cli@latest</div>
            <div class="mt-2"><span class="text-gray-500"># {{ $t('help.api.cliLogin') }}</span></div>
            <div>ltedu login --token &lt;your-cli-token&gt;</div>
            <div class="mt-2"><span class="text-gray-500"># {{ $t('help.api.cliSync') }}</span></div>
            <div>ltedu question sync --file questions.csv</div>
          </div>
          <router-link
            to="/account/cli-tokens"
            class="mt-4 inline-flex items-center px-4 py-2 bg-gray-800 text-white rounded-lg text-sm hover:bg-gray-900 transition"
          >
            {{ $t('help.api.getCliToken') }} →
          </router-link>
        </div>

        <div class="bg-indigo-50 rounded-lg border border-indigo-200 p-6">
          <h3 class="text-lg font-normal text-indigo-800 mb-2">{{ $t('help.api.authTitle') }}</h3>
          <p class="text-indigo-700 text-sm leading-relaxed">{{ $t('help.api.authDesc') }}</p>
          <div class="mt-3 bg-white rounded p-3 text-sm font-mono text-gray-700">
            Authorization: Bearer &lt;your-token&gt;
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const activeTab = ref('getting-started')
const activeRole = ref<'student' | 'teacher'>('student')
const searchQuery = ref('')
const openFaqs = ref<Set<number>>(new Set())

const tabs = computed(() => [
  { id: 'getting-started', label: t('help.tabs.gettingStarted') },
  { id: 'features', label: t('help.tabs.features') },
  { id: 'faq', label: t('help.tabs.faq') },
  { id: 'api', label: t('help.tabs.api') },
])

const studentSteps = computed(() => [
  {
    title: t('help.gettingStarted.student.step1Title'),
    desc: t('help.gettingStarted.student.step1Desc'),
    link: '/register',
    linkLabel: t('help.gettingStarted.student.step1Link'),
  },
  {
    title: t('help.gettingStarted.student.step2Title'),
    desc: t('help.gettingStarted.student.step2Desc'),
    link: '/practice/quick',
    linkLabel: t('help.gettingStarted.student.step2Link'),
  },
  {
    title: t('help.gettingStarted.student.step3Title'),
    desc: t('help.gettingStarted.student.step3Desc'),
    link: '/practice/paper',
    linkLabel: t('help.gettingStarted.student.step3Link'),
  },
  {
    title: t('help.gettingStarted.student.step4Title'),
    desc: t('help.gettingStarted.student.step4Desc'),
    link: '',
    linkLabel: '',
  },
])

const teacherSteps = computed(() => [
  {
    title: t('help.gettingStarted.teacher.step1Title'),
    desc: t('help.gettingStarted.teacher.step1Desc'),
    link: '/register',
    linkLabel: t('help.gettingStarted.teacher.step1Link'),
  },
  {
    title: t('help.gettingStarted.teacher.step2Title'),
    desc: t('help.gettingStarted.teacher.step2Desc'),
    link: '/account/teacher-application',
    linkLabel: t('help.gettingStarted.teacher.step2Link'),
  },
  {
    title: t('help.gettingStarted.teacher.step3Title'),
    desc: t('help.gettingStarted.teacher.step3Desc'),
    link: '/paper/exam/create',
    linkLabel: t('help.gettingStarted.teacher.step3Link'),
  },
  {
    title: t('help.gettingStarted.teacher.step4Title'),
    desc: t('help.gettingStarted.teacher.step4Desc'),
    link: '/paper/exam/teacher',
    linkLabel: t('help.gettingStarted.teacher.step4Link'),
  },
])

const features = computed(() => [
  {
    id: 'quick-practice',
    icon: '⚡',
    bgClass: 'bg-green-100',
    title: t('help.features.quickPractice.title'),
    desc: t('help.features.quickPractice.desc'),
    bullets: [
      t('help.features.quickPractice.bullet1'),
      t('help.features.quickPractice.bullet2'),
      t('help.features.quickPractice.bullet3'),
    ],
  },
  {
    id: 'past-papers',
    icon: '📄',
    bgClass: 'bg-blue-100',
    title: t('help.features.pastPapers.title'),
    desc: t('help.features.pastPapers.desc'),
    bullets: [
      t('help.features.pastPapers.bullet1'),
      t('help.features.pastPapers.bullet2'),
      t('help.features.pastPapers.bullet3'),
    ],
  },
  {
    id: 'exam-builder',
    icon: '🔧',
    bgClass: 'bg-purple-100',
    title: t('help.features.examBuilder.title'),
    desc: t('help.features.examBuilder.desc'),
    bullets: [
      t('help.features.examBuilder.bullet1'),
      t('help.features.examBuilder.bullet2'),
      t('help.features.examBuilder.bullet3'),
    ],
  },
  {
    id: 'ai-analysis',
    icon: '🤖',
    bgClass: 'bg-indigo-100',
    title: t('help.features.aiAnalysis.title'),
    desc: t('help.features.aiAnalysis.desc'),
    bullets: [
      t('help.features.aiAnalysis.bullet1'),
      t('help.features.aiAnalysis.bullet2'),
      t('help.features.aiAnalysis.bullet3'),
    ],
  },
  {
    id: 'syllabus',
    icon: '📚',
    bgClass: 'bg-yellow-100',
    title: t('help.features.syllabus.title'),
    desc: t('help.features.syllabus.desc'),
    bullets: [
      t('help.features.syllabus.bullet1'),
      t('help.features.syllabus.bullet2'),
      t('help.features.syllabus.bullet3'),
    ],
  },
  {
    id: 'learning-plans',
    icon: '🗓️',
    bgClass: 'bg-pink-100',
    title: t('help.features.learningPlans.title'),
    desc: t('help.features.learningPlans.desc'),
    bullets: [
      t('help.features.learningPlans.bullet1'),
      t('help.features.learningPlans.bullet2'),
      t('help.features.learningPlans.bullet3'),
    ],
  },
])

const allFaqs = computed(() => [
  { question: t('help.faq.q1'), answer: t('help.faq.a1') },
  { question: t('help.faq.q2'), answer: t('help.faq.a2') },
  { question: t('help.faq.q3'), answer: t('help.faq.a3') },
  { question: t('help.faq.q4'), answer: t('help.faq.a4') },
  { question: t('help.faq.q5'), answer: t('help.faq.a5') },
  { question: t('help.faq.q6'), answer: t('help.faq.a6') },
  { question: t('help.faq.q7'), answer: t('help.faq.a7') },
  { question: t('help.faq.q8'), answer: t('help.faq.a8') },
  { question: t('help.faq.q9'), answer: t('help.faq.a9') },
  { question: t('help.faq.q10'), answer: t('help.faq.a10') },
])

const filteredFaqs = computed(() => {
  if (!searchQuery.value.trim()) return allFaqs.value
  const q = searchQuery.value.toLowerCase()
  return allFaqs.value.filter(
    item => item.question.toLowerCase().includes(q) || item.answer.toLowerCase().includes(q)
  )
})

function toggleFaq(index: number) {
  if (openFaqs.value.has(index)) {
    openFaqs.value.delete(index)
  } else {
    openFaqs.value.add(index)
  }
}
</script>
