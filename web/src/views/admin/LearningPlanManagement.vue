<template>
  <div class="p-6">
    <header class="mb-6">
      <div class="flex justify-between items-start">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ t('learningPlan.title') }}</h1>
          <p class="mt-2 text-sm text-gray-600">{{ t('learningPlan.subtitle') }}</p>
          <div v-if="classInfo" class="mt-3 text-sm text-gray-700 font-medium">
            {{ classInfo.name }}
          </div>
        </div>
        <div class="flex gap-2">
          <button
            @click="openGenerateModal"
            class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-green-600 hover:bg-green-700 focus:outline-none"
          >
            {{ t('learningPlan.generateTemplate') }}
          </button>
          <button
            @click="openCreateModal"
            class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none"
          >
            {{ t('learningPlan.createPlan') }}
          </button>
          <router-link
            to="/admin/classes"
            class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
          >
            {{ t('learningPlan.backToClasses') }}
          </router-link>
        </div>
      </div>
    </header>

    <!-- Filters -->
    <div class="mb-4 flex gap-4 items-center">
      <div>
        <label class="text-sm font-medium text-gray-700 mr-2">{{ t('learningPlan.userId') }}</label>
        <input v-model.number="filterUserId" type="number" placeholder="User ID" @change="loadPlans"
          class="border border-gray-300 rounded-md px-3 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 w-32" />
      </div>
      <div class="flex gap-1">
        <button v-for="type in ['', 'long', 'mid', 'short']" :key="type"
          @click="filterPlanType = type; loadPlans()"
          :class="['px-3 py-1 text-sm rounded-md font-medium border', filterPlanType === type ? 'bg-indigo-600 text-white border-indigo-600' : 'bg-white text-gray-700 border-gray-300 hover:bg-gray-50']">
          {{ type === '' ? 'All' : t(`learningPlan.type${type.charAt(0).toUpperCase() + type.slice(1)}`) }}
        </button>
      </div>
    </div>

    <!-- Table -->
    <div class="bg-white shadow rounded-lg overflow-hidden">
      <div v-if="loading" class="text-center py-12 text-gray-500">{{ t('learningPlan.loading') }}</div>
      <div v-else-if="!plans.length" class="text-center py-12 text-gray-500">{{ t('learningPlan.noPlans') }}</div>
      <table v-else class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('learningPlan.userId') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('learningPlan.planType') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('learningPlan.version') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Updated</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="plan in plans" :key="plan.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ plan.id }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ plan.userId }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm">
              <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                :class="{
                  'bg-blue-100 text-blue-800': plan.planType === 'long',
                  'bg-yellow-100 text-yellow-800': plan.planType === 'mid',
                  'bg-green-100 text-green-800': plan.planType === 'short'
                }">
                {{ t(`learningPlan.type${plan.planType.charAt(0).toUpperCase() + plan.planType.slice(1)}`) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">v{{ plan.version }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ plan.updatedAt?.slice(0, 10) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2">
              <button @click="openEditModal(plan)" class="text-indigo-600 hover:text-indigo-900">{{ t('learningPlan.editPlan') }}</button>
              <button @click="openVersionsModal(plan)" class="text-gray-600 hover:text-gray-900">{{ t('learningPlan.viewVersions') }}</button>
              <router-link :to="`/admin/learning-plans/${plan.id}/phase-plans`" class="text-purple-600 hover:text-purple-900">{{ t('learningPlan.phaseManage') }}</router-link>
              <button @click="confirmDelete(plan)" class="text-red-600 hover:text-red-900">{{ t('common.delete') }}</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Create / Edit Modal -->
    <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-lg p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">
          {{ editingPlan ? t('learningPlan.editPlan') : t('learningPlan.createPlan') }}
        </h2>
        <div class="space-y-4">
          <div v-if="!editingPlan">
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('learningPlan.userId') }}</label>
            <input v-model.number="form.userId" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
          </div>
          <div v-if="!editingPlan">
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('learningPlan.planType') }}</label>
            <select v-model="form.planType" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500">
              <option value="long">{{ t('learningPlan.typeLong') }}</option>
              <option value="mid">{{ t('learningPlan.typeMid') }}</option>
              <option value="short">{{ t('learningPlan.typeShort') }}</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('learningPlan.content') }}</label>
            <textarea v-model="form.content" rows="6" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('learningPlan.comment') }}</label>
            <input v-model="form.comment" type="text" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
          </div>
        </div>
        <div class="mt-6 flex justify-end space-x-3">
          <button @click="closeModal" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('learningPlan.cancel') }}</button>
          <button @click="submitForm" :disabled="saving" class="px-4 py-2 text-sm font-medium text-white bg-indigo-600 rounded-md hover:bg-indigo-700 disabled:opacity-50">
            {{ saving ? t('learningPlan.loading') : (editingPlan ? t('learningPlan.save') : t('learningPlan.create')) }}
          </button>
        </div>
      </div>
    </div>

    <!-- Version History Modal -->
    <div v-if="showVersionsModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-2xl p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">{{ t('learningPlan.viewVersions') }}</h2>
        <div v-if="loadingVersions" class="text-center py-8 text-gray-500">{{ t('learningPlan.loading') }}</div>
        <div v-else-if="!versions.length" class="text-center py-8 text-gray-400">—</div>
        <div v-else class="space-y-3 max-h-96 overflow-y-auto">
          <div v-for="v in versions" :key="v.id" class="border border-gray-200 rounded-lg p-4">
            <div class="flex justify-between items-start">
              <div>
                <span class="font-medium text-gray-900">v{{ v.version }}</span>
                <span class="ml-3 text-sm text-gray-500">{{ v.createdAt?.slice(0, 10) }}</span>
                <p v-if="v.comment" class="text-sm text-gray-600 mt-1">{{ v.comment }}</p>
              </div>
              <button @click="doRollback(v)" :disabled="saving" class="text-sm text-indigo-600 hover:text-indigo-900 disabled:opacity-50">
                {{ t('learningPlan.rollback') }}
              </button>
            </div>
            <p class="mt-2 text-xs text-gray-400 line-clamp-2">{{ v.content }}</p>
          </div>
        </div>
        <div class="mt-6 flex justify-end">
          <button @click="showVersionsModal = false" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('learningPlan.cancel') }}</button>
        </div>
      </div>
    </div>

    <!-- Generate Template Plans Modal -->
    <div v-if="showGenerateModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-lg p-6 max-h-[90vh] overflow-y-auto">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">{{ t('learningPlan.generateTemplate') }}</h2>
        <p class="text-sm text-gray-600 mb-4">{{ t('learningPlan.generateTemplateDesc') }}</p>
        <div class="space-y-4">
          <!-- Per-exam-node time range pickers -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">{{ t('learningPlan.examNodeSchedules') }}</label>
            <p class="text-xs text-gray-500 mb-2">{{ t('learningPlan.examNodeSchedulesDesc') }}</p>
            <div v-if="genForm.examNodes.length === 0" class="text-sm text-gray-400 italic">{{ t('learningPlan.noExamNodesFound') }}</div>
            <div v-for="(node, idx) in genForm.examNodes" :key="node.examNodeId" class="mb-3 p-3 border border-gray-200 rounded-md">
              <p class="text-sm font-medium text-gray-700 mb-2">{{ node.name }}</p>
              <div class="grid grid-cols-2 gap-3">
                <div>
                  <label class="block text-xs text-gray-500 mb-1">{{ t('learningPlan.startMonth') }}</label>
                  <input v-model="genForm.examNodes[idx].startMonth" type="month" class="w-full border border-gray-300 rounded-md px-2 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
                </div>
                <div>
                  <label class="block text-xs text-gray-500 mb-1">{{ t('learningPlan.endMonth') }}</label>
                  <input v-model="genForm.examNodes[idx].endMonth" type="month" class="w-full border border-gray-300 rounded-md px-2 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
                </div>
              </div>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('learningPlan.phaseRatios') }}</label>
            <p class="text-xs text-gray-500 mb-1">{{ t('learningPlan.phaseRatiosDesc') }}</p>
            <div class="grid grid-cols-4 gap-2">
              <div v-for="(label, i) in phaseLabels" :key="i">
                <label class="block text-xs text-gray-500 mb-1">{{ label }}</label>
                <input v-model.number="genForm.phaseRatios[i]" type="number" min="0" max="100"
                  class="w-full border border-gray-300 rounded-md px-2 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
              </div>
            </div>
            <p class="text-xs text-gray-400 mt-1">{{ t('learningPlan.phaseRatiosSum') }}: {{ genForm.phaseRatios.reduce((a, b) => a + b, 0) }}%</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('learningPlan.comment') }}</label>
            <input v-model="genForm.comment" type="text" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
          </div>
        </div>
        <div v-if="generateResult" class="mt-4 p-3 rounded-md" :class="generateResult.errors?.length ? 'bg-yellow-50 border border-yellow-200' : 'bg-green-50 border border-green-200'">
          <p class="text-sm font-medium">{{ t('learningPlan.generateSuccess') }}: {{ generateResult.count }} {{ t('learningPlan.plans') }} ({{ generateResult.studentCount }} {{ t('learningPlan.students') }})</p>
          <ul v-if="generateResult.errors?.length" class="mt-2 text-xs text-yellow-700 space-y-1">
            <li v-for="e in generateResult.errors" :key="e">{{ e }}</li>
          </ul>
        </div>
        <div class="mt-6 flex justify-end space-x-3">
          <button @click="closeGenerateModal" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('learningPlan.cancel') }}</button>
          <button @click="submitGenerate" :disabled="generating" class="px-4 py-2 text-sm font-medium text-white bg-green-600 rounded-md hover:bg-green-700 disabled:opacity-50">
            {{ generating ? t('learningPlan.generating') : t('learningPlan.generate') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Delete Confirm Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-sm p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-2">{{ t('learningPlan.deletePlan') }}</h2>
        <p class="text-sm text-gray-600 mb-6">{{ t('learningPlan.confirmDelete') }}</p>
        <div class="flex justify-end space-x-3">
          <button @click="showDeleteModal = false" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('learningPlan.cancel') }}</button>
          <button @click="doDelete" :disabled="saving" class="px-4 py-2 text-sm font-medium text-white bg-red-600 rounded-md hover:bg-red-700 disabled:opacity-50">
            {{ saving ? t('learningPlan.loading') : t('learningPlan.confirm') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import learningPlanService from '../../services/learningPlanService'
import classService from '../../services/classService'
import examNodeService from '../../services/examNodeService'
import type { StudentLearningPlan, StudentLearningPlanVersion } from '../../models/learningPlan.model'

const { t } = useI18n()
const route = useRoute()
const classId = Number(route.params.classId)

const plans = ref<StudentLearningPlan[]>([])
const classInfo = ref<any>(null)
const versions = ref<StudentLearningPlanVersion[]>([])
const loading = ref(false)
const saving = ref(false)
const loadingVersions = ref(false)

const filterUserId = ref<number | ''>('')
const filterPlanType = ref('')

const showModal = ref(false)
const showVersionsModal = ref(false)
const showDeleteModal = ref(false)

const editingPlan = ref<StudentLearningPlan | null>(null)
const versionsPlan = ref<StudentLearningPlan | null>(null)
const deletingPlan = ref<StudentLearningPlan | null>(null)

const form = ref({ userId: 0, planType: 'long', content: '', comment: '' })

async function loadPlans() {
  loading.value = true
  try {
    const query: any = { classId, pageSize: 100, pageIndex: 1 }
    if (filterUserId.value) query.userId = filterUserId.value
    if (filterPlanType.value) query.planType = filterPlanType.value
    const res = await learningPlanService.list(query)
    if (res.code === 0) plans.value = res.data.list ?? []
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }
}

async function loadClassInfo() {
  try {
    const res = await classService.getById(classId)
    if (res.code === 0) classInfo.value = res.data
  } catch (e) {
    console.error(e)
  }
}

function openCreateModal() {
  editingPlan.value = null
  form.value = { userId: 0, planType: 'long', content: '', comment: '' }
  showModal.value = true
}

function openEditModal(plan: StudentLearningPlan) {
  editingPlan.value = plan
  form.value = { userId: plan.userId, planType: plan.planType, content: plan.content, comment: '' }
  showModal.value = true
}

function closeModal() {
  showModal.value = false
}

async function submitForm() {
  saving.value = true
  try {
    if (editingPlan.value) {
      const res = await learningPlanService.update({ id: editingPlan.value.id, content: form.value.content, comment: form.value.comment })
      if (res.code === 0) { await loadPlans(); closeModal() }
    } else {
      const res = await learningPlanService.create({ classId, userId: form.value.userId, planType: form.value.planType, content: form.value.content, comment: form.value.comment })
      if (res.code === 0) { await loadPlans(); closeModal() }
    }
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

async function openVersionsModal(plan: StudentLearningPlan) {
  versionsPlan.value = plan
  showVersionsModal.value = true
  loadingVersions.value = true
  versions.value = []
  try {
    const res = await learningPlanService.versions(plan.id)
    if (res.code === 0) versions.value = res.data.list ?? []
  } catch (e) {
    console.error(e)
  } finally {
    loadingVersions.value = false
  }
}

async function doRollback(v: StudentLearningPlanVersion) {
  if (!versionsPlan.value) return
  saving.value = true
  try {
    const res = await learningPlanService.rollback(versionsPlan.value.id, v.version, `Rollback to v${v.version}`)
    if (res.code === 0) { await loadPlans(); showVersionsModal.value = false }
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

function confirmDelete(plan: StudentLearningPlan) {
  deletingPlan.value = plan
  showDeleteModal.value = true
}

async function doDelete() {
  if (!deletingPlan.value) return
  saving.value = true
  try {
    const res = await learningPlanService.delete(deletingPlan.value.id)
    if (res.code === 0) { await loadPlans(); showDeleteModal.value = false }
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

const showGenerateModal = ref(false)
const generating = ref(false)
const generateResult = ref<any>(null)
const syllabusExamNodes = ref<{ id: number; name: string }[]>([])
const genForm = ref({
  phaseRatios: [30, 20, 20, 10],
  examNodes: [] as { examNodeId: number; name: string; startMonth: string; endMonth: string }[],
  comment: ''
})
const phaseLabels = ['新课学习(%)', '一轮复习(%)', '专题综合(%)', '集中刷题(%)']

async function openGenerateModal() {
  generateResult.value = null
  // Load exam nodes for the bound syllabus.
  syllabusExamNodes.value = []
  if (classInfo.value?.syllabusId) {
    try {
      const res = await examNodeService.list(classInfo.value.syllabusId)
      if (res.code === 0) {
        const nodes: { id: number; name: string }[] = (res.data?.list ?? []).map(
          (n: { id: number; name: string }) => ({ id: n.id, name: n.name })
        )
        syllabusExamNodes.value = nodes
        // Pre-fill examNodes with empty time ranges.
        genForm.value.examNodes = nodes.map(n => ({
          examNodeId: n.id,
          name: n.name,
          startMonth: '',
          endMonth: ''
        }))
      }
    } catch (e) {
      console.error(e)
    }
  }
  showGenerateModal.value = true
}

function closeGenerateModal() {
  showGenerateModal.value = false
  if (generateResult.value) loadPlans()
}

async function submitGenerate() {
  const hasAllTimes = genForm.value.examNodes.length > 0 &&
    genForm.value.examNodes.every(n => n.startMonth && n.endMonth)
  if (!hasAllTimes) {
    alert(t('learningPlan.examNodesTimeRequired'))
    return
  }
  if (!classInfo.value?.syllabusId) {
    alert(t('learningPlan.noSyllabus'))
    return
  }
  generating.value = true
  generateResult.value = null
  try {
    const res = await learningPlanService.generateTemplate({
      classId,
      syllabusId: classInfo.value.syllabusId,
      phaseRatios: genForm.value.phaseRatios,
      examNodes: genForm.value.examNodes.map(n => ({
        examNodeId: n.examNodeId,
        startMonth: n.startMonth,
        endMonth: n.endMonth
      })),
      comment: genForm.value.comment
    })
    if (res.code === 0) generateResult.value = res.data
  } catch (e) {
    console.error(e)
  } finally {
    generating.value = false
  }
}

onMounted(() => {
  loadPlans()
  loadClassInfo()
})
</script>
