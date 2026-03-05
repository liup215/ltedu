<template>
  <div class="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
    <header class="mb-10 text-center">
      <h1 class="text-4xl font-bold text-gray-900">{{ $t('blog.list.title') }}</h1>
      <p class="mt-2 text-lg text-gray-600">{{ $t('blog.list.subtitle') }}</p>
    </header>

    <!-- Category Filter -->
    <div class="flex flex-wrap gap-2 mb-8 justify-center">
      <button
        @click="selectedCategory = ''"
        :class="['px-4 py-2 rounded-full text-sm font-medium transition-colors', selectedCategory === '' ? 'bg-indigo-600 text-white' : 'bg-gray-100 text-gray-700 hover:bg-gray-200']"
      >
        {{ $t('blog.list.all') }}
      </button>
      <button
        v-for="cat in categories"
        :key="cat.value"
        @click="selectedCategory = cat.value"
        :class="['px-4 py-2 rounded-full text-sm font-medium transition-colors', selectedCategory === cat.value ? 'bg-indigo-600 text-white' : 'bg-gray-100 text-gray-700 hover:bg-gray-200']"
      >
        {{ $i18n.locale === 'zh' ? cat.labelZh : cat.labelEn }}
      </button>
    </div>

    <!-- Search -->
    <div class="mb-8 max-w-lg mx-auto">
      <input
        v-model="searchKeyword"
        type="text"
        :placeholder="$t('blog.list.searchPlaceholder')"
        class="w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
      />
    </div>

    <!-- Loading -->
    <div v-if="loading" class="text-center py-20 text-gray-500">{{ $t('common.loading') }}</div>

    <!-- Empty -->
    <div v-else-if="!posts || posts.length === 0" class="text-center py-20 text-gray-500">{{ $t('blog.list.noPosts') }}</div>

    <!-- Posts Grid -->
    <div v-else class="grid gap-8 md:grid-cols-2 lg:grid-cols-3">
      <article
        v-for="post in posts"
        :key="post.id"
        class="bg-white rounded-xl shadow hover:shadow-md transition-shadow overflow-hidden flex flex-col cursor-pointer"
        @click="goToPost(post.slug)"
      >
        <div v-if="post.coverImage" class="h-48 bg-gray-100 overflow-hidden">
          <img :src="post.coverImage" :alt="post.title" class="w-full h-full object-cover" />
        </div>
        <div v-else class="h-2 w-full" :class="categoryBgClass(post.category)"></div>
        <div class="p-5 flex flex-col flex-1">
          <div class="flex items-center justify-between mb-3">
            <span class="text-xs font-medium px-2 py-1 rounded-full" :class="categoryClass(post.category)">
              {{ categoryLabel(post.category) }}
            </span>
            <span class="text-xs text-gray-400">{{ formatDate(post.publishedAt || post.createdAt) }}</span>
          </div>
          <h2 class="text-lg font-bold text-gray-900 mb-2 line-clamp-2 hover:text-indigo-600 transition-colors">{{ post.title }}</h2>
          <p class="text-sm text-gray-600 flex-1 line-clamp-3">{{ post.summary }}</p>
          <div class="mt-4 flex items-center justify-between">
            <span class="text-xs text-gray-500">{{ post.authorName }}</span>
            <span class="text-xs text-gray-400">{{ $t('blog.list.views', { count: post.viewCount }) }}</span>
          </div>
          <div v-if="post.tags" class="mt-3 flex flex-wrap gap-1">
            <span v-for="tag in parseTags(post.tags)" :key="tag" class="text-xs bg-gray-100 text-gray-600 px-2 py-0.5 rounded">{{ tag }}</span>
          </div>
        </div>
      </article>
    </div>

    <!-- Pagination -->
    <div v-if="total > pageSize" class="mt-10 flex justify-center gap-4">
      <button @click="prevPage" :disabled="pageIndex <= 1" class="px-4 py-2 border rounded-lg text-sm disabled:opacity-50 hover:bg-gray-50">
        {{ $t('blog.list.previous') }}
      </button>
      <span class="px-4 py-2 text-sm text-gray-600">{{ pageIndex }} / {{ Math.ceil(total / pageSize) }}</span>
      <button @click="nextPage" :disabled="pageIndex * pageSize >= total" class="px-4 py-2 border rounded-lg text-sm disabled:opacity-50 hover:bg-gray-50">
        {{ $t('blog.list.next') }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import blogService from '../services/blogService'
import { BLOG_CATEGORIES } from '../models/blog.model'
import type { BlogPost } from '../models/blog.model'

const router = useRouter()
const { locale } = useI18n()

const posts = ref<BlogPost[]>([])
const loading = ref(false)
const total = ref(0)
const pageIndex = ref(1)
const pageSize = 9
const selectedCategory = ref('')
const searchKeyword = ref('')
const categories = BLOG_CATEGORIES

async function loadPosts() {
  loading.value = true
  try {
    const res = await blogService.getPublishedPosts({
      category: selectedCategory.value,
      keyword: searchKeyword.value,
      pageIndex: pageIndex.value,
      pageSize,
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
watch([selectedCategory, searchKeyword], () => {
  pageIndex.value = 1
  loadPosts()
})

function prevPage() {
  if (pageIndex.value > 1) { pageIndex.value--; loadPosts() }
}
function nextPage() {
  if (pageIndex.value * pageSize < total.value) { pageIndex.value++; loadPosts() }
}

function goToPost(slug: string) {
  router.push({ name: 'BlogPost', params: { slug } })
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
const categoryBgColors: Record<string, string> = {
  system_updates: 'bg-blue-500',
  user_guides: 'bg-purple-500',
  learning_methods: 'bg-green-500',
  major_events: 'bg-orange-500',
}
function categoryClass(cat: string) {
  return categoryColors[cat] || 'bg-gray-100 text-gray-800'
}
function categoryBgClass(cat: string) {
  return categoryBgColors[cat] || 'bg-gray-300'
}

function formatDate(dateStr?: string) {
  if (!dateStr) return ''
  return new Date(dateStr).toLocaleDateString()
}

function parseTags(tags: string) {
  return tags.split(',').map(t => t.trim()).filter(Boolean)
}
</script>
