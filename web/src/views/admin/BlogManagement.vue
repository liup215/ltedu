<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('blog.management.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('blog.management.subtitle') }}</p>
    </header>

    <!-- Filters and actions -->
    <div class="mb-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0 sm:space-x-4">
      <div class="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 w-full sm:w-auto">
        <input
          type="text"
          v-model="searchKeyword"
          :placeholder="$t('blog.management.searchPlaceholder')"
          class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-sm w-full sm:w-64"
        />
        <select
          v-model="filterCategory"
          class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-sm"
        >
          <option value="">{{ $t('blog.management.allCategories') }}</option>
          <option v-for="cat in categories" :key="cat.value" :value="cat.value">
            {{ $i18n.locale === 'zh' ? cat.labelZh : cat.labelEn }}
          </option>
        </select>
        <select
          v-model="filterStatus"
          class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-sm"
        >
          <option value="">{{ $t('blog.management.allStatuses') }}</option>
          <option value="draft">{{ $t('blog.status.draft') }}</option>
          <option value="published">{{ $t('blog.status.published') }}</option>
        </select>
      </div>
      <button
        @click="openCreateModal"
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 whitespace-nowrap"
      >
        {{ $t('blog.management.createPost') }}
      </button>
    </div>

    <!-- Table -->
    <div class="bg-white shadow overflow-x-auto sm:rounded-lg">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('blog.fields.title') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('blog.fields.category') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('blog.fields.status') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('blog.fields.author') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('blog.fields.createdAt') }}</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{{ $t('common.actions') }}</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-if="loading">
            <td colspan="7" class="px-6 py-4 text-center text-sm text-gray-500">{{ $t('common.loading') }}</td>
          </tr>
          <tr v-else-if="!posts || posts.length === 0">
            <td colspan="7" class="px-6 py-4 text-center text-sm text-gray-500">{{ $t('common.noData') }}</td>
          </tr>
          <tr v-for="post in posts" :key="post.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ post.id }}</td>
            <td class="px-6 py-4 text-sm text-gray-900 max-w-xs truncate">{{ post.title }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              <span class="px-2 py-1 rounded-full text-xs font-medium" :class="categoryClass(post.category)">
                {{ categoryLabel(post.category) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm">
              <span class="px-2 py-1 rounded-full text-xs font-medium" :class="post.status === 'published' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'">
                {{ post.status === 'published' ? $t('blog.status.published') : $t('blog.status.draft') }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ post.authorName }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ formatDate(post.createdAt) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-3">
              <button @click="openEditModal(post)" class="text-indigo-600 hover:text-indigo-900">{{ $t('common.edit') }}</button>
              <button @click="confirmDelete(post)" class="text-red-600 hover:text-red-900">{{ $t('common.delete') }}</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div v-if="total > 0" class="mt-4 flex items-center justify-between">
      <p class="text-sm text-gray-700">
        {{ $t('blog.management.pageInfo', { from: (pageIndex - 1) * pageSize + 1, to: Math.min(pageIndex * pageSize, total), total }) }}
      </p>
      <div class="flex space-x-2">
        <button @click="prevPage" :disabled="pageIndex <= 1" class="px-3 py-1 border rounded text-sm disabled:opacity-50">{{ $t('blog.management.previous') }}</button>
        <button @click="nextPage" :disabled="pageIndex * pageSize >= total" class="px-3 py-1 border rounded text-sm disabled:opacity-50">{{ $t('blog.management.next') }}</button>
      </div>
    </div>

    <!-- Create/Edit Modal -->
    <div v-if="showModal" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-start justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75" @click="closeModal"></div>
        <div class="inline-block align-top bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-3xl sm:w-full">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <h3 class="text-lg font-medium text-gray-900 mb-4">
              {{ editingPost ? $t('blog.management.editPost') : $t('blog.management.createPost') }}
            </h3>
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.title') }} *</label>
                <input v-model="form.title" type="text" class="mt-1 w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-indigo-500 focus:border-indigo-500" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.slug') }}</label>
                <input v-model="form.slug" type="text" class="mt-1 w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-indigo-500 focus:border-indigo-500" :placeholder="$t('blog.management.slugPlaceholder')" />
              </div>
              <div class="flex space-x-4">
                <div class="flex-1">
                  <label class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.category') }} *</label>
                  <select v-model="form.category" class="mt-1 w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="">{{ $t('blog.management.selectCategory') }}</option>
                    <option v-for="cat in categories" :key="cat.value" :value="cat.value">
                      {{ $i18n.locale === 'zh' ? cat.labelZh : cat.labelEn }}
                    </option>
                  </select>
                </div>
                <div class="flex-1">
                  <label class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.status') }}</label>
                  <select v-model="form.status" class="mt-1 w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="draft">{{ $t('blog.status.draft') }}</option>
                    <option value="published">{{ $t('blog.status.published') }}</option>
                  </select>
                </div>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.tags') }}</label>
                <input v-model="form.tags" type="text" class="mt-1 w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-indigo-500 focus:border-indigo-500" :placeholder="$t('blog.management.tagsPlaceholder')" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.coverImage') }}</label>
                <input v-model="form.coverImage" type="text" class="mt-1 w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-indigo-500 focus:border-indigo-500" :placeholder="$t('blog.management.coverImagePlaceholder')" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.summary') }}</label>
                <textarea v-model="form.summary" rows="2" class="mt-1 w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-indigo-500 focus:border-indigo-500"></textarea>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.content') }} *</label>
                <textarea v-model="form.content" rows="10" class="mt-1 w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-indigo-500 focus:border-indigo-500 font-mono"></textarea>
                <p class="mt-1 text-xs text-gray-500">{{ $t('blog.management.contentTip') }}</p>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse gap-2">
            <button @click="submitForm" :disabled="saving" class="w-full sm:w-auto inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-indigo-600 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none disabled:opacity-50 text-sm">
              {{ saving ? $t('blog.management.saving') : $t('common.save') }}
            </button>
            <button @click="closeModal" class="w-full sm:w-auto inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none text-sm">
              {{ $t('common.cancel') }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Delete Confirm Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 z-50 flex items-center justify-center">
      <div class="fixed inset-0 bg-gray-500 bg-opacity-75" @click="showDeleteModal = false"></div>
      <div class="relative bg-white rounded-lg shadow-xl p-6 max-w-md w-full mx-4">
        <h3 class="text-lg font-medium text-gray-900 mb-2">{{ $t('blog.management.confirmDelete') }}</h3>
        <p class="text-sm text-gray-500 mb-4">{{ deletingPost?.title }}</p>
        <div class="flex justify-end gap-2">
          <button @click="showDeleteModal = false" class="px-4 py-2 border border-gray-300 rounded-md text-sm text-gray-700 hover:bg-gray-50">{{ $t('common.cancel') }}</button>
          <button @click="doDelete" :disabled="deleting" class="px-4 py-2 bg-red-600 text-white rounded-md text-sm hover:bg-red-700 disabled:opacity-50">
            {{ deleting ? $t('blog.management.deleting') : $t('common.delete') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import blogService from '../../services/blogService'
import { BLOG_CATEGORIES } from '../../models/blog.model'
import type { BlogPost } from '../../models/blog.model'

const { locale, t } = useI18n()

const posts = ref<BlogPost[]>([])
const loading = ref(false)
const total = ref(0)
const pageIndex = ref(1)
const pageSize = ref(20)
const searchKeyword = ref('')
const filterCategory = ref('')
const filterStatus = ref('')

const categories = BLOG_CATEGORIES

const showModal = ref(false)
const editingPost = ref<BlogPost | null>(null)
const saving = ref(false)
const form = ref({
  id: 0,
  title: '',
  slug: '',
  summary: '',
  content: '',
  category: '',
  tags: '',
  coverImage: '',
  status: 'draft',
})

const showDeleteModal = ref(false)
const deletingPost = ref<BlogPost | null>(null)
const deleting = ref(false)

async function loadPosts() {
  loading.value = true
  try {
    const res = await blogService.listPosts({
      keyword: searchKeyword.value,
      category: filterCategory.value,
      status: filterStatus.value,
      pageIndex: pageIndex.value,
      pageSize: pageSize.value,
    })
    if (res.code === 0 && res.data) {
      posts.value = res.data.list || []
      total.value = res.data.total || 0
    }
  } finally {
    loading.value = false
  }
}

onMounted(loadPosts)
watch([searchKeyword, filterCategory, filterStatus], () => {
  pageIndex.value = 1
  loadPosts()
})

function prevPage() {
  if (pageIndex.value > 1) {
    pageIndex.value--
    loadPosts()
  }
}
function nextPage() {
  if (pageIndex.value * pageSize.value < total.value) {
    pageIndex.value++
    loadPosts()
  }
}

function openCreateModal() {
  editingPost.value = null
  form.value = { id: 0, title: '', slug: '', summary: '', content: '', category: '', tags: '', coverImage: '', status: 'draft' }
  showModal.value = true
}

function openEditModal(post: BlogPost) {
  editingPost.value = post
  form.value = {
    id: post.id,
    title: post.title,
    slug: post.slug,
    summary: post.summary,
    content: post.content,
    category: post.category,
    tags: post.tags,
    coverImage: post.coverImage,
    status: post.status,
  }
  showModal.value = true
}

function closeModal() {
  showModal.value = false
}

async function submitForm() {
  if (!form.value.title || !form.value.category || !form.value.content) {
    alert(t('blog.management.requiredFields'))
    return
  }
  saving.value = true
  try {
    let res
    if (editingPost.value) {
      res = await blogService.updatePost(form.value)
    } else {
      const { id: _, ...createData } = form.value
      res = await blogService.createPost(createData)
    }
    if (res.code === 0) {
      closeModal()
      loadPosts()
    } else {
      alert(res.message || t('blog.management.saveFailed'))
    }
  } finally {
    saving.value = false
  }
}

function confirmDelete(post: BlogPost) {
  deletingPost.value = post
  showDeleteModal.value = true
}

async function doDelete() {
  if (!deletingPost.value) return
  deleting.value = true
  try {
    const res = await blogService.deletePost(deletingPost.value.id)
    if (res.code === 0) {
      showDeleteModal.value = false
      loadPosts()
    } else {
      alert(res.message || t('blog.management.deleteFailed'))
    }
  } finally {
    deleting.value = false
  }
}

function categoryLabel(cat: string) {
  const found = categories.find(c => c.value === cat)
  if (!found) return cat
  return locale.value === 'zh' ? found.labelZh : found.labelEn
}

const categoryColors: Record<string, string> = {
  system_updates: 'bg-blue-100 text-blue-800',
  user_guides: 'bg-purple-100 text-purple-800',
  learning_methods: 'bg-green-100 text-green-800',
  major_events: 'bg-orange-100 text-orange-800',
}
function categoryClass(cat: string) {
  return categoryColors[cat] || 'bg-gray-100 text-gray-800'
}

function formatDate(dateStr: string) {
  if (!dateStr) return '-'
  return new Date(dateStr).toLocaleDateString()
}
</script>
