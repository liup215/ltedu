<template>
  <div class="p-6">
    <header class="mb-6 flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold text-gray-900">{{ $t('adminBlog.title') }}</h1>
        <p class="mt-1 text-sm text-gray-600">{{ $t('adminBlog.subtitle') }}</p>
      </div>
      <button
        @click="openCreateModal"
        class="px-4 py-2 bg-indigo-600 text-white rounded-lg text-sm hover:bg-indigo-700 transition"
      >
        + {{ $t('adminBlog.createPost') }}
      </button>
    </header>

    <!-- Filter -->
    <div class="flex gap-3 mb-4 items-center">
      <select v-model="filterStatus" @change="loadPosts" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
        <option value="">{{ $t('adminBlog.allStatus') }}</option>
        <option value="published">{{ $t('adminBlog.published') }}</option>
        <option value="draft">{{ $t('adminBlog.draft') }}</option>
      </select>
      <input
        v-model="searchKeyword"
        @keydown.enter="loadPosts"
        :placeholder="$t('blog.search')"
        class="px-3 py-2 border border-gray-300 rounded-lg text-sm"
      />
      <button @click="loadPosts" class="px-3 py-2 bg-indigo-600 text-white rounded-lg text-sm hover:bg-indigo-700 transition">
        {{ $t('blog.searchBtn') }}
      </button>
    </div>

    <!-- Table -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
      <table class="w-full text-sm">
        <thead class="bg-gray-50 border-b border-gray-200">
          <tr>
            <th class="px-4 py-3 text-left text-gray-600 font-medium">{{ $t('adminBlog.colTitle') }}</th>
            <th class="px-4 py-3 text-left text-gray-600 font-medium">{{ $t('adminBlog.colCategory') }}</th>
            <th class="px-4 py-3 text-left text-gray-600 font-medium">{{ $t('adminBlog.colStatus') }}</th>
            <th class="px-4 py-3 text-left text-gray-600 font-medium">{{ $t('blog.views') }}</th>
            <th class="px-4 py-3 text-left text-gray-600 font-medium">{{ $t('adminBlog.colDate') }}</th>
            <th class="px-4 py-3 text-left text-gray-600 font-medium">{{ $t('common.actions') }}</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="loading"><td colspan="6" class="text-center py-8 text-gray-400">{{ $t('common.loading') }}</td></tr>
          <tr v-else-if="posts.length === 0"><td colspan="6" class="text-center py-8 text-gray-400">{{ $t('common.noData') }}</td></tr>
          <tr v-else v-for="post in posts" :key="post.id" class="border-b border-gray-100 hover:bg-gray-50">
            <td class="px-4 py-3">
              <div class="font-medium text-gray-900">{{ post.title }}</div>
              <div v-if="post.isTop" class="text-xs text-yellow-600">📌 {{ $t('blog.pinned') }}</div>
            </td>
            <td class="px-4 py-3 text-gray-600">{{ post.category || '-' }}</td>
            <td class="px-4 py-3">
              <span
                class="px-2 py-0.5 rounded-full text-xs"
                :class="post.status === 'published' ? 'bg-green-100 text-green-700' : 'bg-yellow-100 text-yellow-700'"
              >
                {{ post.status === 'published' ? $t('adminBlog.published') : $t('adminBlog.draft') }}
              </span>
            </td>
            <td class="px-4 py-3 text-gray-600">{{ post.viewCount }}</td>
            <td class="px-4 py-3 text-gray-500 text-xs">{{ formatDate(post.createdAt) }}</td>
            <td class="px-4 py-3 flex gap-2">
              <button @click="openEditModal(post)" class="px-2 py-1 text-xs text-indigo-600 border border-indigo-300 rounded hover:bg-indigo-50">
                {{ $t('common.edit') }}
              </button>
              <button
                @click="togglePublish(post)"
                class="px-2 py-1 text-xs border rounded"
                :class="post.status === 'published' ? 'text-yellow-600 border-yellow-300 hover:bg-yellow-50' : 'text-green-600 border-green-300 hover:bg-green-50'"
              >
                {{ post.status === 'published' ? $t('adminBlog.unpublish') : $t('adminBlog.publish') }}
              </button>
              <button @click="confirmDelete(post.id)" class="px-2 py-1 text-xs text-red-600 border border-red-300 rounded hover:bg-red-50">
                {{ $t('common.delete') }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div v-if="total > pageSize" class="flex justify-center gap-2 mt-4">
      <button @click="prevPage" :disabled="page <= 1" class="px-3 py-1.5 border rounded text-sm disabled:opacity-50">{{ $t('blog.prev') }}</button>
      <span class="px-3 py-1.5 text-sm">{{ page }} / {{ Math.ceil(total / pageSize) }}</span>
      <button @click="nextPage" :disabled="page >= Math.ceil(total / pageSize)" class="px-3 py-1.5 border rounded text-sm disabled:opacity-50">{{ $t('blog.next') }}</button>
    </div>

    <!-- Create/Edit Modal -->
    <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-xl shadow-xl w-full max-w-2xl max-h-screen overflow-y-auto">
        <div class="p-6 border-b border-gray-200 flex items-center justify-between">
          <h3 class="text-lg font-semibold">{{ editingId ? $t('adminBlog.editPost') : $t('adminBlog.createPost') }}</h3>
          <button @click="closeModal" class="text-gray-400 hover:text-gray-600">✕</button>
        </div>
        <div class="p-6 space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('adminBlog.fieldTitle') }} *</label>
            <input v-model="form.title" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('adminBlog.fieldSummary') }}</label>
            <textarea v-model="form.summary" rows="2" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm resize-none"></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('adminBlog.fieldContent') }} *</label>
            <textarea v-model="form.content" rows="8" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm font-mono"></textarea>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('adminBlog.fieldCategory') }}</label>
              <input v-model="form.category" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('adminBlog.fieldTags') }}</label>
              <input v-model="form.tags" :placeholder="$t('adminBlog.tagsPlaceholder')" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" />
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('adminBlog.fieldCoverImage') }}</label>
            <input v-model="form.coverImage" :placeholder="$t('adminBlog.coverImagePlaceholder')" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" />
          </div>
          <div class="flex items-center gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('adminBlog.fieldStatus') }}</label>
              <select v-model="form.status" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
                <option value="draft">{{ $t('adminBlog.draft') }}</option>
                <option value="published">{{ $t('adminBlog.published') }}</option>
              </select>
            </div>
            <div class="flex items-center gap-2 mt-5">
              <input type="checkbox" v-model="form.isTop" id="isTop" class="w-4 h-4" />
              <label for="isTop" class="text-sm text-gray-700">{{ $t('adminBlog.pinPost') }}</label>
            </div>
          </div>
        </div>
        <div class="p-6 border-t border-gray-200 flex justify-end gap-3">
          <button @click="closeModal" class="px-4 py-2 text-sm text-gray-600 border border-gray-300 rounded-lg hover:bg-gray-50">{{ $t('common.cancel') }}</button>
          <button @click="savePost" :disabled="saving" class="px-4 py-2 text-sm bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 disabled:opacity-50">
            {{ saving ? $t('common.loading') : $t('common.save') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import blogService from '../../services/blogService'
import type { BlogPost } from '../../models/blog.model'

const { t } = useI18n()

const posts = ref<BlogPost[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = 20
const loading = ref(false)
const filterStatus = ref('')
const searchKeyword = ref('')

const showModal = ref(false)
const saving = ref(false)
const editingId = ref<number | null>(null)

interface BlogForm {
  title: string
  summary: string
  content: string
  coverImage: string
  category: string
  tags: string
  status: 'draft' | 'published'
  isTop: boolean
}

const form = ref<BlogForm>({
  title: '',
  summary: '',
  content: '',
  coverImage: '',
  category: '',
  tags: '',
  status: 'draft',
  isTop: false,
})

async function loadPosts() {
  loading.value = true
  try {
    const res = await blogService.adminListPosts({
      pageIndex: page.value,
      pageSize,
      status: filterStatus.value,
      keyword: searchKeyword.value,
    })
    posts.value = res.data?.list || []
    total.value = res.data?.total || 0
  } catch {
    posts.value = []
  } finally {
    loading.value = false
  }
}

function openCreateModal() {
  editingId.value = null
  form.value = { title: '', summary: '', content: '', coverImage: '', category: '', tags: '', status: 'draft', isTop: false }
  showModal.value = true
}

function openEditModal(post: BlogPost) {
  editingId.value = post.id
  form.value = {
    title: post.title,
    summary: post.summary,
    content: post.content,
    coverImage: post.coverImage,
    category: post.category,
    tags: post.tags,
    status: post.status,
    isTop: post.isTop,
  }
  showModal.value = true
}

function closeModal() {
  showModal.value = false
}

async function savePost() {
  if (!form.value.title || !form.value.content) return
  saving.value = true
  try {
    if (editingId.value) {
      await blogService.adminUpdatePost({ id: editingId.value, ...form.value })
    } else {
      await blogService.createPost(form.value)
    }
    closeModal()
    await loadPosts()
  } catch {
    // ignore
  } finally {
    saving.value = false
  }
}

async function togglePublish(post: BlogPost) {
  const newStatus = post.status === 'published' ? 'draft' : 'published'
  try {
    await blogService.adminUpdatePost({ id: post.id, status: newStatus })
    await loadPosts()
  } catch {
    // ignore
  }
}

async function confirmDelete(id: number) {
  if (!confirm(t('adminBlog.confirmDelete'))) return
  try {
    await blogService.adminDeletePost(id)
    await loadPosts()
  } catch {
    // ignore
  }
}

function prevPage() {
  if (page.value > 1) { page.value--; loadPosts() }
}

function nextPage() {
  if (page.value < Math.ceil(total.value / pageSize)) { page.value++; loadPosts() }
}

function formatDate(dateStr: string): string {
  return new Date(dateStr).toLocaleDateString()
}

onMounted(() => loadPosts())
</script>
