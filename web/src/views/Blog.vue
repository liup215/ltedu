<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Hero section -->
    <div class="bg-white border-b border-gray-200 py-12 px-8 text-center">
      <h1 class="text-3xl font-bold text-gray-900 mb-3">{{ $t('blog.title') }}</h1>
      <p class="text-gray-600 max-w-xl mx-auto">{{ $t('blog.subtitle') }}</p>
    </div>

    <div class="max-w-5xl mx-auto px-4 py-8">
      <!-- Filter bar -->
      <div class="flex flex-wrap gap-3 mb-8 items-center justify-between">
        <div class="flex flex-wrap gap-2">
          <button
            @click="filterCategory('')"
            class="px-4 py-1.5 rounded-full text-sm font-normal transition"
            :class="selectedCategory === '' ? 'bg-indigo-600 text-white' : 'bg-white text-gray-600 border border-gray-300 hover:bg-gray-50'"
          >
            {{ $t('blog.allCategories') }}
          </button>
          <button
            v-for="cat in categories"
            :key="cat"
            @click="filterCategory(cat)"
            class="px-4 py-1.5 rounded-full text-sm font-normal transition"
            :class="selectedCategory === cat ? 'bg-indigo-600 text-white' : 'bg-white text-gray-600 border border-gray-300 hover:bg-gray-50'"
          >
            {{ cat }}
          </button>
        </div>

        <div class="flex items-center gap-2">
          <input
            v-model="searchKeyword"
            @keydown.enter="search"
            :placeholder="$t('blog.search')"
            class="px-3 py-1.5 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 w-48"
          />
          <button
            @click="search"
            class="px-3 py-1.5 bg-indigo-600 text-white rounded-lg text-sm hover:bg-indigo-700 transition"
          >
            {{ $t('blog.searchBtn') }}
          </button>
        </div>
      </div>

      <!-- Loading state -->
      <div v-if="loading" class="text-center py-16 text-gray-400">
        <div class="animate-spin w-8 h-8 border-4 border-indigo-600 border-t-transparent rounded-full mx-auto mb-4"></div>
        {{ $t('common.loading') }}
      </div>

      <!-- Post list -->
      <div v-else-if="posts.length > 0" class="space-y-6">
        <!-- Pinned post -->
        <div
          v-for="post in posts"
          :key="post.id"
          @click="viewPost(post.id)"
          class="bg-white rounded-xl shadow-sm border border-gray-200 hover:shadow-md transition cursor-pointer overflow-hidden"
        >
          <div class="flex gap-0">
            <!-- Cover image -->
            <div
              v-if="post.coverImage"
              class="w-48 flex-shrink-0 bg-gray-100"
            >
              <img :src="post.coverImage" :alt="post.title" class="w-full h-full object-cover" />
            </div>

            <div class="flex-1 p-6">
              <!-- Top badge + category -->
              <div class="flex items-center gap-2 mb-2">
                <span
                  v-if="post.isTop"
                  class="px-2 py-0.5 bg-yellow-100 text-yellow-700 text-xs rounded-full font-normal"
                >
                  📌 {{ $t('blog.pinned') }}
                </span>
                <span
                  v-if="post.category"
                  class="px-2 py-0.5 bg-indigo-100 text-indigo-700 text-xs rounded-full"
                >
                  {{ post.category }}
                </span>
              </div>

              <h2 class="text-xl font-semibold text-gray-900 mb-2 hover:text-indigo-600 transition">
                {{ post.title }}
              </h2>

              <p v-if="post.summary" class="text-gray-600 text-sm leading-relaxed mb-4 line-clamp-2">
                {{ post.summary }}
              </p>

              <!-- Tags -->
              <div v-if="post.tags" class="flex flex-wrap gap-1 mb-3">
                <span
                  v-for="tag in post.tags.split(',').filter(t => t.trim())"
                  :key="tag"
                  class="px-2 py-0.5 bg-gray-100 text-gray-600 text-xs rounded"
                >
                  # {{ tag.trim() }}
                </span>
              </div>

              <div class="flex items-center justify-between text-xs text-gray-400">
                <span>{{ formatDate(post.createdAt) }}</span>
                <span>👁 {{ post.viewCount }} {{ $t('blog.views') }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Pagination -->
        <div v-if="total > pageSize" class="flex justify-center gap-2 mt-8">
          <button
            @click="prevPage"
            :disabled="page <= 1"
            class="px-4 py-2 border border-gray-300 rounded-lg text-sm disabled:opacity-50 hover:bg-gray-50 transition"
          >
            {{ $t('blog.prev') }}
          </button>
          <span class="px-4 py-2 text-sm text-gray-600">
            {{ page }} / {{ Math.ceil(total / pageSize) }}
          </span>
          <button
            @click="nextPage"
            :disabled="page >= Math.ceil(total / pageSize)"
            class="px-4 py-2 border border-gray-300 rounded-lg text-sm disabled:opacity-50 hover:bg-gray-50 transition"
          >
            {{ $t('blog.next') }}
          </button>
        </div>
      </div>

      <!-- Empty state -->
      <div v-else class="text-center py-16 text-gray-400">
        <div class="text-5xl mb-4">📝</div>
        <p>{{ $t('blog.noPosts') }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import blogService from '../services/blogService'
import type { BlogPost } from '../models/blog.model'

const router = useRouter()
const { t } = useI18n()

const posts = ref<BlogPost[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = 10
const loading = ref(false)
const selectedCategory = ref('')
const searchKeyword = ref('')
const categories = ref<string[]>([])

async function loadPosts() {
  loading.value = true
  try {
    const res = await blogService.listPosts({
      pageIndex: page.value,
      pageSize,
      category: selectedCategory.value,
      keyword: searchKeyword.value,
    })
    posts.value = res.data?.list || []
    total.value = res.data?.total || 0

    // Update categories from all posts (load all without category filter for category list)
    if (searchKeyword.value === '' && selectedCategory.value === '') {
      const catSet = new Set<string>()
      posts.value.forEach((p) => {
        if (p.category) catSet.add(p.category)
      })
      categories.value = Array.from(catSet)
    }
  } catch {
    posts.value = []
    total.value = 0
  } finally {
    loading.value = false
  }
}

function filterCategory(cat: string) {
  selectedCategory.value = cat
  page.value = 1
  loadPosts()
}

function search() {
  page.value = 1
  loadPosts()
}

function prevPage() {
  if (page.value > 1) {
    page.value--
    loadPosts()
  }
}

function nextPage() {
  if (page.value < Math.ceil(total.value / pageSize)) {
    page.value++
    loadPosts()
  }
}

function viewPost(id: number) {
  router.push(`/blog/${id}`)
}

function formatDate(dateStr: string): string {
  const d = new Date(dateStr)
  return d.toLocaleDateString()
}

onMounted(() => {
  loadPosts()
})
</script>
