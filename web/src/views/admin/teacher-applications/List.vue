<template>
  <div>
    <div class="p-6">
      <header class="mb-6">
        <h1 class="text-3xl font-bold text-gray-900">教师申请管理</h1>
        <p class="mt-1 text-sm text-gray-500">审核教师资格申请</p>
      </header>

      <!-- Search Form -->
      <div class="bg-white p-6 rounded-lg shadow mb-6">
        <form @submit.prevent="loadData(1)" class="flex flex-wrap gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">状态</label>
            <select
              v-model="query.status"
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
            >
              <option :value="undefined">全部</option>
              <option :value="TeacherApplicationStatus.Pending">审核中</option>
              <option :value="TeacherApplicationStatus.Approved">已通过</option>
              <option :value="TeacherApplicationStatus.Rejected">已拒绝</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">申请时间</label>
            <div class="flex gap-2 items-center">
              <input
                type="date"
                v-model="query.startDate"
                class="block rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              />
              <span>至</span>
              <input
                type="date"
                v-model="query.endDate"
                class="block rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              />
            </div>
          </div>
          <div class="flex items-end">
            <button
              type="submit"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              搜索
            </button>
            <button
              type="button"
              @click="resetSearch"
              class="ml-3 inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              重置
            </button>
          </div>
        </form>
      </div>

      <!-- Applications Table -->
      <div class="bg-white shadow rounded-lg">
        <div class="flex flex-col">
          <div class="overflow-x-auto">
            <div class="align-middle inline-block min-w-full">
              <div class="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
                <table class="min-w-full divide-y divide-gray-200">
                  <thead class="bg-gray-50">
                    <tr>
                      <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        申请人
                      </th>
                      <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        申请理由
                      </th>
                      <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        教学经验
                      </th>
                      <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        状态
                      </th>
                      <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        申请时间
                      </th>
                      <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        操作
                      </th>
                    </tr>
                  </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <tr v-for="item in applications" :key="item.id" class="hover:bg-gray-50">
                      <td class="px-6 py-4 whitespace-nowrap">
                        <div class="text-sm font-medium text-gray-900">{{ item.user?.username }}</div>
                        <div class="text-sm text-gray-500">{{ item.user?.email }}</div>
                      </td>
                      <td class="px-6 py-4">
                        <div class="text-sm text-gray-900 line-clamp-2">{{ item.motivation }}</div>
                      </td>
                      <td class="px-6 py-4">
                        <div class="text-sm text-gray-900 line-clamp-2">{{ item.experience }}</div>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span :class="{
                          'px-2 inline-flex text-xs leading-5 font-semibold rounded-full': true,
                          'bg-yellow-100 text-yellow-800': item.status === TeacherApplicationStatus.Pending,
                          'bg-green-100 text-green-800': item.status === TeacherApplicationStatus.Approved,
                          'bg-red-100 text-red-800': item.status === TeacherApplicationStatus.Rejected,
                        }">
                          {{ getStatusText(item.status) }}
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        {{ formatDate(item.appliedAt) }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
                        <button
                          v-if="item.status === TeacherApplicationStatus.Pending"
                          @click="openApproveDialog(item)"
                          class="text-indigo-600 hover:text-indigo-900"
                        >
                          通过
                        </button>
                        <button
                          v-if="item.status === TeacherApplicationStatus.Pending"
                          @click="openRejectDialog(item)"
                          class="text-red-600 hover:text-red-900"
                        >
                          拒绝
                        </button>
                        <button
                          @click="openDetailsDialog(item)"
                          class="text-gray-600 hover:text-gray-900"
                        >
                          详情
                        </button>
                      </td>
                    </tr>
                    <tr v-if="applications.length === 0">
                      <td colspan="6" class="px-6 py-4 text-center text-sm text-gray-500">
                        暂无数据
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>

        <!-- Pagination -->
        <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
          <div class="flex-1 flex justify-between sm:hidden">
            <button
              @click="handlePageChange(currentPage - 1)"
              :disabled="currentPage === 1"
              class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
            >
              上一页
            </button>
            <button
              @click="handlePageChange(currentPage + 1)"
              :disabled="currentPage * pageSize >= total"
              class="ml-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
            >
              下一页
            </button>
          </div>
          <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
            <div>
              <p class="text-sm text-gray-700">
                共 <span class="font-medium">{{ total }}</span> 条记录
              </p>
            </div>
            <div>
              <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
                <button
                  v-for="page in totalPages"
                  :key="page"
                  @click="handlePageChange(page)"
                  :class="[
                    'relative inline-flex items-center px-4 py-2 border text-sm font-medium',
                    currentPage === page
                      ? 'z-10 bg-indigo-50 border-indigo-500 text-indigo-600'
                      : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50'
                  ]"
                >
                  {{ page }}
                </button>
              </nav>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Modals Container -->
    <div class="modals-container">
      <!-- Review Dialog -->
      <div v-if="showReviewDialog" class="fixed inset-0 z-[100]" role="dialog" aria-modal="true">
        <div class="absolute inset-0 bg-black/50"></div>
        <div class="absolute inset-0 z-10 overflow-y-auto">
          <div class="flex min-h-full items-center justify-center p-4">
            <div class="relative w-full max-w-lg rounded-lg bg-white p-6 shadow-xl">
              <h3 class="text-lg font-medium text-gray-900">
                {{ reviewType === 'approve' ? '通过申请' : '拒绝申请' }}
              </h3>
              <div class="mt-4">
                <label for="notes" class="block text-sm font-medium text-gray-700">审核意见</label>
                <textarea
                  id="notes"
                  v-model="adminNotes"
                  rows="4"
                  class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  :placeholder="reviewType === 'approve' ? '请输入通过理由（可选）' : '请输入拒绝理由（必填）'"
                  :required="reviewType === 'reject'"
                ></textarea>
              </div>
              <div class="mt-6 grid grid-cols-2 gap-3">
                <button
                  type="button"
                  class="w-full rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50"
                  @click="closeReviewDialog"
                >
                  取消
                </button>
                <button
                  type="button"
                  :class="[
                    'w-full rounded-md border border-transparent px-4 py-2 text-sm font-medium text-white shadow-sm',
                    reviewType === 'approve'
                      ? 'bg-green-600 hover:bg-green-700'
                      : 'bg-red-600 hover:bg-red-700'
                  ]"
                  @click="handleReview"
                  :disabled="loading || (reviewType === 'reject' && !adminNotes.trim())"
                >
                  确认
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Details Dialog -->
      <div v-if="showDetailsDialog" class="fixed inset-0 z-[100]" role="dialog" aria-modal="true">
        <div class="absolute inset-0 bg-black/50"></div>
        <div class="absolute inset-0 z-10 overflow-y-auto">
          <div class="flex min-h-full items-center justify-center p-4">
            <div class="relative w-full max-w-2xl rounded-lg bg-white p-6 shadow-xl">
              <h3 class="text-lg font-medium text-gray-900">申请详情</h3>
              <div class="mt-4 space-y-4" v-if="selectedApplication">
                <div>
                  <h4 class="text-sm font-medium text-gray-500">申请人信息</h4>
                  <div class="mt-2 grid grid-cols-2 gap-4">
                    <div>
                      <span class="text-sm text-gray-500">用户名：</span>
                      <span class="text-sm text-gray-900">{{ selectedApplication.user?.username }}</span>
                    </div>
                    <div>
                      <span class="text-sm text-gray-500">邮箱：</span>
                      <span class="text-sm text-gray-900">{{ selectedApplication.user?.email || '未设置' }}</span>
                    </div>
                  </div>
                </div>
                <div>
                  <h4 class="text-sm font-medium text-gray-500">申请理由</h4>
                  <p class="mt-1 text-sm text-gray-900 whitespace-pre-wrap">{{ selectedApplication.motivation }}</p>
                </div>
                <div>
                  <h4 class="text-sm font-medium text-gray-500">教学经验</h4>
                  <p class="mt-1 text-sm text-gray-900 whitespace-pre-wrap">{{ selectedApplication.experience }}</p>
                </div>
                <div>
                  <h4 class="text-sm font-medium text-gray-500">申请时间</h4>
                  <p class="mt-1 text-sm text-gray-900">{{ formatDate(selectedApplication.appliedAt) }}</p>
                </div>
                <template v-if="selectedApplication.reviewedAt">
                  <div>
                    <h4 class="text-sm font-medium text-gray-500">审核时间</h4>
                    <p class="mt-1 text-sm text-gray-900">{{ formatDate(selectedApplication.reviewedAt) }}</p>
                  </div>
                  <div>
                    <h4 class="text-sm font-medium text-gray-500">审核意见</h4>
                    <p class="mt-1 text-sm text-gray-900">{{ selectedApplication.adminNotes || '无' }}</p>
                  </div>
                </template>
              </div>
              <div class="mt-6">
                <button
                  type="button"
                  class="w-full rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50"
                  @click="closeDetailsDialog"
                >
                  关闭
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { teacherApplicationService } from '../../../services/teacherApplicationService';
import type { TeacherApplication, TeacherApplicationQuery, TeacherApplicationStatusType } from '../../../models/teacher-application.model';
import { TeacherApplicationStatus } from '../../../models/teacher-application.model';

const loading = ref(false);
const applications = ref<TeacherApplication[]>([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);

const query = ref<TeacherApplicationQuery>({
  status: undefined,
  startDate: '',
  endDate: '',
  pageIndex: 1,
  pageSize: 10,
});

// Review Dialog
const showReviewDialog = ref(false);
const reviewType = ref<'approve' | 'reject'>('approve');
const adminNotes = ref('');
const selectedApplication = ref<TeacherApplication | null>(null);

// Details Dialog
const showDetailsDialog = ref(false);

// Load applications
const loadData = async (page?: number) => {
  if (page) {
    currentPage.value = page;
  }
  query.value.pageIndex = currentPage.value;

  loading.value = true;
  try {
    const response = await teacherApplicationService.list(query.value);
    applications.value = response.data.list;
    total.value = response.data.total;
  } catch (error) {
    console.log(error)
  } finally {
    loading.value = false;
  }
};

// Computed
const totalPages = computed(() => Math.ceil(total.value / pageSize.value));

// Pagination handlers
const handlePageChange = (page: number) => {
  if (page < 1 || page > totalPages.value) return;
  loadData(page);
};

// Review handlers
const openApproveDialog = (application: TeacherApplication) => {
  reviewType.value = 'approve';
  selectedApplication.value = application;
  adminNotes.value = '';
  showReviewDialog.value = true;
};

const openRejectDialog = (application: TeacherApplication) => {
  reviewType.value = 'reject';
  selectedApplication.value = application;
  adminNotes.value = '';
  showReviewDialog.value = true;
};

const handleReview = async () => {
  if (!selectedApplication.value) return;

  if (reviewType.value === 'reject' && !adminNotes.value.trim()) {
    alert('请输入拒绝理由');
    return;
  }

  loading.value = true;
  try {
    if (reviewType.value === 'approve') {
      await teacherApplicationService.approve(selectedApplication.value.id, adminNotes.value);
      alert('申请已通过');
    } else {
      await teacherApplicationService.reject(selectedApplication.value.id, adminNotes.value);
      alert('申请已拒绝');
    }
    closeReviewDialog();
    loadData();
  } catch (error: any) {
    alert(error.response?.data?.message || '操作失败');
  } finally {
    loading.value = false;
  }
};

const closeReviewDialog = () => {
  showReviewDialog.value = false;
  selectedApplication.value = null;
  adminNotes.value = '';
};

// Details handlers
const openDetailsDialog = (application: TeacherApplication) => {
  console.log('Opening details dialog', application);
  selectedApplication.value = application;
  showDetailsDialog.value = true;
  console.log('Details dialog state:', showDetailsDialog.value);
};

const closeDetailsDialog = () => {
  showDetailsDialog.value = false;
  selectedApplication.value = null;
};

// Search handlers
const resetSearch = () => {
  query.value.status = undefined;
  query.value.startDate = '';
  query.value.endDate = '';
  loadData(1);
};

// Utils
const getStatusText = (status: TeacherApplicationStatusType) => {
  switch (status) {
    case TeacherApplicationStatus.Pending:
      return '审核中';
    case TeacherApplicationStatus.Approved:
      return '已通过';
    case TeacherApplicationStatus.Rejected:
      return '已拒绝';
    default:
      return '未知状态';
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

// Initial load
loadData();
</script>
