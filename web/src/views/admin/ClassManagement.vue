<template>
  <div class="p-6">
    <header class="mb-6">
      <div class="flex justify-between items-start">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">{{ t('class.title') }}</h1>
          <p class="mt-2 text-sm text-gray-600">{{ t('class.subtitle') }}</p>
        </div>
        <button
          @click="openCreateModal"
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          {{ t('class.createClass') }}
        </button>
      </div>
    </header>

    <!-- Table -->
    <div class="bg-white shadow rounded-lg overflow-hidden">
      <div v-if="loading" class="text-center py-12 text-gray-500">{{ t('class.loading') }}</div>
      <div v-else-if="!classes.length" class="text-center py-12 text-gray-500">{{ t('class.noClasses') }}</div>
      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('class.name') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('class.type') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('class.syllabus') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ t('class.inviteCode') }}</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="cls in classes" :key="cls.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ cls.name }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                :class="cls.classType === 1 ? 'bg-blue-100 text-blue-800' : 'bg-purple-100 text-purple-800'">
                {{ cls.classType === 1 ? t('class.typeTeaching') : t('class.typeAdmin') }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              <span v-if="cls.syllabus">{{ cls.syllabus.name }}</span>
              <span v-else class="text-gray-400 italic">{{ t('class.syllabusNotBound') }}</span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 font-mono">{{ cls.inviteCode }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2">
              <button @click="openEditModal(cls)" class="text-indigo-600 hover:text-indigo-900">{{ t('common.edit') }}</button>
              <button v-if="!cls.syllabus" @click="openBindSyllabusModal(cls)" class="text-green-600 hover:text-green-900">{{ t('class.bindSyllabus') }}</button>
              <button v-if="cls.syllabus" @click="doUnbindSyllabus(cls)" class="text-yellow-600 hover:text-yellow-900">{{ t('class.unbindSyllabus') }}</button>
              <router-link :to="`/admin/classes/${cls.id}/learning-plans`" class="text-purple-600 hover:text-purple-900">{{ t('learningPlan.title') }}</router-link>
              <button @click="openStudentsModal(cls)" class="text-teal-600 hover:text-teal-900">{{ t('class.students') }}</button>
              <button @click="confirmDelete(cls)" class="text-red-600 hover:text-red-900">{{ t('common.delete') }}</button>
            </td>
          </tr>
        </tbody>
      </table>
      </div>
    </div>

    <!-- Create / Edit Modal -->
    <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-md p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">
          {{ editingClass ? t('class.editClass') : t('class.createClass') }}
        </h2>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('class.name') }}</label>
            <input v-model="form.name" type="text" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
          </div>
          <div v-if="!editingClass">
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('class.type') }}</label>
            <select v-model="form.classType" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500">
              <option :value="1">{{ t('class.typeTeaching') }}</option>
              <option :value="2">{{ t('class.typeAdmin') }}</option>
            </select>
          </div>
        </div>
        <div class="mt-6 flex justify-end space-x-3">
          <button @click="closeModal" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('class.cancel') }}</button>
          <button @click="submitForm" :disabled="saving" class="px-4 py-2 text-sm font-medium text-white bg-indigo-600 rounded-md hover:bg-indigo-700 disabled:opacity-50">
            {{ saving ? t('class.loading') : (editingClass ? t('class.save') : t('class.create')) }}
          </button>
        </div>
      </div>
    </div>

    <!-- Bind Syllabus Modal -->
    <div v-if="showBindModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-md p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">{{ t('class.bindSyllabus') }}</h2>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">{{ t('class.selectSyllabus') }}</label>
          <select v-model="selectedSyllabusId" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500">
            <option value="">{{ t('class.selectSyllabus') }}</option>
            <option v-for="s in syllabuses" :key="s.id" :value="s.id">{{ s.name }}</option>
          </select>
        </div>
        <div class="mt-6 flex justify-end space-x-3">
          <button @click="showBindModal = false" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('class.cancel') }}</button>
          <button @click="doBindSyllabus" :disabled="!selectedSyllabusId || saving" class="px-4 py-2 text-sm font-medium text-white bg-green-600 rounded-md hover:bg-green-700 disabled:opacity-50">
            {{ saving ? t('class.loading') : t('class.confirm') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Students Modal -->
    <div v-if="showStudentsModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-2xl p-6">
        <div class="flex justify-between items-center mb-4">
          <h2 class="text-lg font-semibold text-gray-900">{{ t('class.students') }} — {{ studentsClass?.name }}</h2>
          <button @click="showStudentsModal = false" class="text-gray-400 hover:text-gray-600 text-xl leading-none">&times;</button>
        </div>
        <div v-if="studentsLoading" class="text-center py-8 text-gray-500">{{ t('class.loading') }}</div>
        <div v-else-if="!students.length" class="text-center py-8 text-gray-500">{{ t('class.noStudents') }}</div>
        <div v-else class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">{{ t('class.name') }}</th>
              <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">{{ t('class.studentStatus') }}</th>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase">{{ t('common.actions') }}</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="stu in students" :key="stu.id">
              <td class="px-4 py-2 text-sm text-gray-900">
                {{ stu.realname || stu.nickname || stu.username }}
              </td>
              <td class="px-4 py-2 text-sm">
                <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium"
                  :class="studentStatusClass(stu.studentStatus)">
                  {{ studentStatusLabel(stu.studentStatus) }}
                </span>
              </td>
              <td class="px-4 py-2 text-right text-sm">
                <select
                  :value="stu.studentStatus"
                  @change="onStatusChange(stu, ($event.target as HTMLSelectElement).value)"
                  class="border border-gray-300 rounded px-2 py-1 text-sm focus:outline-none focus:ring-1 focus:ring-indigo-500"
                >
                  <option :value="1">{{ t('class.statusStudying') }}</option>
                  <option :value="2">{{ t('class.statusGraduated') }}</option>
                  <option :value="3">{{ t('class.statusTransferred') }}</option>
                  <option :value="4">{{ t('class.statusDropped') }}</option>
                </select>
              </td>
            </tr>
          </tbody>
        </table>
        </div>
      </div>
    </div>

    <!-- Delete Confirm Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-sm p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-2">{{ t('class.deleteClass') }}</h2>
        <p class="text-sm text-gray-600 mb-6">{{ t('class.confirmDelete') }}</p>
        <div class="flex justify-end space-x-3">
          <button @click="showDeleteModal = false" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50">{{ t('class.cancel') }}</button>
          <button @click="doDelete" :disabled="saving" class="px-4 py-2 text-sm font-medium text-white bg-red-600 rounded-md hover:bg-red-700 disabled:opacity-50">
            {{ saving ? t('class.loading') : t('class.confirm') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import classService from '../../services/classService'
import syllabusService from '../../services/syllabusService'
import type { Class, ClassStudent } from '../../models/class.model'
import {
  CLASS_STUDENT_STATUS_STUDYING,
  CLASS_STUDENT_STATUS_GRADUATED,
  CLASS_STUDENT_STATUS_TRANSFERRED,
  CLASS_STUDENT_STATUS_DROPPED,
} from '../../models/class.model'

const { t } = useI18n()

const classes = ref<Class[]>([])
const syllabuses = ref<any[]>([])
const loading = ref(false)
const saving = ref(false)

const showModal = ref(false)
const showBindModal = ref(false)
const showDeleteModal = ref(false)

const editingClass = ref<Class | null>(null)
const bindingClass = ref<Class | null>(null)
const deletingClass = ref<Class | null>(null)
const selectedSyllabusId = ref<number | ''>('')

const form = ref({ name: '', classType: 1 })

async function loadClasses() {
  loading.value = true
  try {
    const res = await classService.list({ pageSize: 100, pageIndex: 1 })
    if (res.code === 0) classes.value = res.data.list
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }
}

async function loadSyllabuses() {
  try {
    const res = await syllabusService.getSyllabuses({ pageSize: 100, pageIndex: 1 })
    if (res.code === 0) syllabuses.value = res.data.list
  } catch (e) {
    console.error(e)
  }
}

function openCreateModal() {
  editingClass.value = null
  form.value = { name: '', classType: 1 }
  showModal.value = true
}

function openEditModal(cls: Class) {
  editingClass.value = cls
  form.value = { name: cls.name, classType: cls.classType }
  showModal.value = true
}

function closeModal() {
  showModal.value = false
}

async function submitForm() {
  saving.value = true
  try {
    if (editingClass.value) {
      const res = await classService.update({ id: editingClass.value.id, name: form.value.name })
      if (res.code === 0) {
        await loadClasses()
        closeModal()
      }
    } else {
      const res = await classService.create({ name: form.value.name, classType: form.value.classType })
      if (res.code === 0) {
        await loadClasses()
        closeModal()
      }
    }
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

function openBindSyllabusModal(cls: Class) {
  bindingClass.value = cls
  selectedSyllabusId.value = ''
  showBindModal.value = true
}

async function doBindSyllabus() {
  if (!bindingClass.value || !selectedSyllabusId.value) return
  saving.value = true
  try {
    const res = await classService.bindSyllabus({ classId: bindingClass.value.id, syllabusId: Number(selectedSyllabusId.value) })
    if (res.code === 0) {
      await loadClasses()
      showBindModal.value = false
    }
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

async function doUnbindSyllabus(cls: Class) {
  saving.value = true
  try {
    const res = await classService.unbindSyllabus({ classId: cls.id })
    if (res.code === 0) await loadClasses()
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

function confirmDelete(cls: Class) {
  deletingClass.value = cls
  showDeleteModal.value = true
}

async function doDelete() {
  if (!deletingClass.value) return
  saving.value = true
  try {
    const res = await classService.delete(deletingClass.value.id)
    if (res.code === 0) {
      await loadClasses()
      showDeleteModal.value = false
    }
  } catch (e) {
    console.error(e)
  } finally {
    saving.value = false
  }
}

// --- Students management ---
const showStudentsModal = ref(false)
const studentsClass = ref<Class | null>(null)
const students = ref<ClassStudent[]>([])
const studentsLoading = ref(false)

function studentStatusLabel(status: number): string {
  switch (status) {
    case CLASS_STUDENT_STATUS_STUDYING:    return t('class.statusStudying')
    case CLASS_STUDENT_STATUS_GRADUATED:   return t('class.statusGraduated')
    case CLASS_STUDENT_STATUS_TRANSFERRED: return t('class.statusTransferred')
    case CLASS_STUDENT_STATUS_DROPPED:     return t('class.statusDropped')
    default: return t('class.statusStudying')
  }
}

function studentStatusClass(status: number): string {
  switch (status) {
    case CLASS_STUDENT_STATUS_STUDYING:    return 'bg-green-100 text-green-800'
    case CLASS_STUDENT_STATUS_GRADUATED:   return 'bg-blue-100 text-blue-800'
    case CLASS_STUDENT_STATUS_TRANSFERRED: return 'bg-yellow-100 text-yellow-800'
    case CLASS_STUDENT_STATUS_DROPPED:     return 'bg-red-100 text-red-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

async function openStudentsModal(cls: Class) {
  studentsClass.value = cls
  showStudentsModal.value = true
  studentsLoading.value = true
  try {
    const res = await classService.getStudents(cls.id)
    if (res.code === 0) students.value = res.data.list
  } catch (e) {
    console.error(e)
  } finally {
    studentsLoading.value = false
  }
}

async function onStatusChange(stu: ClassStudent, value: string) {
  if (!studentsClass.value) return
  const newStatus = Number(value)
  try {
    const res = await classService.updateStudentStatus(studentsClass.value.id, stu.id, newStatus)
    if (res.code === 0) {
      stu.studentStatus = newStatus
    } else {
      alert(t('class.updateStatusFailed'))
    }
  } catch (e) {
    console.error(e)
    alert(t('class.updateStatusFailed'))
  }
}

onMounted(() => {
  loadClasses()
  loadSyllabuses()
})
</script>
