<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('teacherApplication.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('teacherApplication.subtitle') }}</p>
    </header>

    <!-- Application Status -->
    <div v-if="existingApplication" class="bg-white shadow sm:rounded-lg mb-6">
      <div class="px-8 py-6">
        <div class="flex justify-between items-start mb-4">
          <h2 class="text-lg font-medium leading-6 text-gray-900">{{ $t('teacherApplication.statusTitle') }}</h2>
          <div v-if="existingApplication.status === TeacherApplicationStatus.Rejected">
            <button
              @click="startNewApplication"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              {{ $t('teacherApplication.reapply') }}
            </button>
          </div>
        </div>
        
        <div class="space-y-4">
          <div class="flex items-center">
            <span class="font-medium text-gray-700 w-24">{{ $t('teacherApplication.status') }}</span>
            <span :class="{
              'px-3 py-1 rounded-full text-sm font-medium': true,
              'bg-yellow-100 text-yellow-700': existingApplication.status === TeacherApplicationStatus.Pending,
              'bg-green-100 text-green-700': existingApplication.status === TeacherApplicationStatus.Approved,
              'bg-red-100 text-red-700': existingApplication.status === TeacherApplicationStatus.Rejected,
            }">
              {{ getStatusText(existingApplication.status) }}
            </span>
          </div>
          <div v-if="existingApplication.adminNotes" class="flex">
            <span class="font-medium text-gray-700 w-24">{{ $t('teacherApplication.adminNotes') }}</span>
            <p class="text-gray-900">{{ existingApplication.adminNotes }}</p>
          </div>
          <div class="flex">
            <span class="font-medium text-gray-700 w-24">{{ $t('teacherApplication.appliedAt') }}</span>
            <p class="text-gray-900">{{ formatDate(existingApplication.appliedAt) }}</p>
          </div>
          <div v-if="existingApplication.reviewedAt" class="flex">
            <span class="font-medium text-gray-700 w-24">{{ $t('teacherApplication.reviewedAt') }}</span>
            <p class="text-gray-900">{{ formatDate(existingApplication.reviewedAt) }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Application Form -->
    <div v-if="!existingApplication || (existingApplication.status === TeacherApplicationStatus.Rejected && showNewApplicationForm)" class="bg-white shadow sm:rounded-lg">
      <form @submit.prevent="onSubmit" class="space-y-8 divide-y divide-gray-200 px-8 py-6">
        <div class="space-y-6">
          <div>
            <h3 class="text-lg font-medium leading-6 text-gray-900">{{ $t('teacherApplication.formTitle') }}</h3>
            <p class="mt-1 text-sm text-gray-500">
              {{ $t('teacherApplication.formTip') }}
            </p>
          </div>

          <div class="space-y-6">
            <div class="sm:w-3/4">
              <label for="motivation" class="block text-sm font-medium text-gray-700 mb-1">{{ $t('teacherApplication.motivation') }}</label>
              <textarea
                id="motivation"
                v-model="form.motivation"
                rows="4"
                class="block w-full rounded-md border-gray-300 shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm resize-none"
                :placeholder="$t('teacherApplication.motivationPlaceholder')"
                maxlength="1000"
                :disabled="loading"
                required
              ></textarea>
              <p class="mt-1 text-sm text-gray-500 text-right">
                {{ form.motivation.length }}/1000
              </p>
            </div>

            <div class="sm:w-3/4">
              <label for="experience" class="block text-sm font-medium text-gray-700 mb-1">{{ $t('teacherApplication.experience') }}</label>
              <textarea
                id="experience"
                v-model="form.experience"
                rows="4"
                class="block w-full rounded-md border-gray-300 shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm resize-none"
                :placeholder="$t('teacherApplication.experiencePlaceholder')"
                maxlength="1000"
                :disabled="loading"
                required
              ></textarea>
              <p class="mt-1 text-sm text-gray-500 text-right">
                {{ form.experience.length }}/1000
              </p>
            </div>
          </div>
        </div>

        <div class="pt-6 flex justify-end space-x-3">
          <button
            type="button"
            @click="router.back()"
            class="inline-flex justify-center rounded-md border border-gray-300 bg-white py-2 px-4 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
          >
            {{ $t('teacherApplication.cancel') }}
          </button>
          <button
            type="submit"
            :disabled="loading"
            class="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 disabled:opacity-50"
          >
            {{ loading ? $t('teacherApplication.submitting') : $t('teacherApplication.submit') }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';
import { teacherApplicationService } from '../../services/teacherApplicationService';
import type { TeacherApplication, TeacherApplicationCreateRequest, TeacherApplicationStatusType } from '../../models/teacher-application.model';
import { TeacherApplicationStatus } from '../../models/teacher-application.model';
import { useUserStore } from '../../stores/userStore';
import { authService } from '../../services/authService';

const router = useRouter();
const { t } = useI18n();
const loading = ref(false);
const existingApplication = ref<TeacherApplication | null>(null);
const showNewApplicationForm = ref(false);

const form = ref<TeacherApplicationCreateRequest>({
  motivation: '',
  experience: '',
});

// Load existing application if any
const loadApplication = async () => {
  try {
    const response = await teacherApplicationService.getCurrentApplication();
    existingApplication.value = response.data;

    // 对比当前储存用户数据和审核数据，如果审核通过且用户数据未更新，则更新userStore
    if (existingApplication.value && existingApplication.value.status === TeacherApplicationStatus.Approved) {
      const userStore = useUserStore();
      if (userStore.user && !userStore.user.isTeacher) {
        let res = await authService.getCurrentUserProfile();
        if (res.code === 0) {
          userStore.setUser(res.data);
        } else {
          console.error('更新用户信息失败:', res.message);
        }
      }
    }

    showNewApplicationForm.value = false;

  } catch (error: any) {
    if (error.response?.status !== 404) {
      // Use browser's native alert since we're not using Element Plus
      alert('加载申请信息失败');
    }
  }
};

const startNewApplication = () => {
  showNewApplicationForm.value = true;
  form.value = {
    motivation: '',
    experience: ''
  };
};

const onSubmit = async () => {
  loading.value = true;
  try {
    let res = await teacherApplicationService.apply(form.value);
    console.log('申请结果:', res.message);
    if (res.code !== 0) {
      throw new Error(res.message);
    }
    await loadApplication();
  } catch (error: any) {
    alert(error.message || '申请提交失败');
  } finally {
    loading.value = false;
  }
};

const getStatusText = (status: TeacherApplicationStatusType) => {
  switch (status) {
    case TeacherApplicationStatus.Pending:
      return t('teacherApplication.pending');
    case TeacherApplicationStatus.Approved:
      return t('teacherApplication.approved');
    case TeacherApplicationStatus.Rejected:
      return t('teacherApplication.rejected');
    default:
      return t('teacherApplication.unknownStatus');
  }
};

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
  });
};

onMounted(() => {
  loadApplication();
});
</script>
