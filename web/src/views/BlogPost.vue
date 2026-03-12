<template>
  <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
    <!-- Back link -->
    <router-link to="/blog" class="inline-flex items-center text-indigo-600 hover:text-indigo-800 mb-8 text-sm font-medium">
      <svg class="mr-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
      </svg>
      {{ $t('blog.post.backToList') }}
    </router-link>

    <!-- Loading -->
    <div v-if="loading" class="text-center py-20 text-gray-500">{{ $t('common.loading') }}</div>

    <!-- Not found -->
    <div v-else-if="!post" class="text-center py-20 text-gray-500">{{ $t('blog.post.notFound') }}</div>

    <!-- Post Content -->
    <article v-else>
      <!-- Cover image -->
      <div v-if="post.coverImage" class="mb-8 rounded-xl overflow-hidden shadow">
        <img :src="post.coverImage" :alt="post.title" class="w-full object-cover max-h-96" />
      </div>

      <!-- Category & Date -->
      <div class="flex items-center gap-3 mb-4">
        <span class="text-xs font-medium px-3 py-1 rounded-full" :class="categoryClass(post.category)">
          {{ categoryLabel(post.category) }}
        </span>
        <span class="text-sm text-gray-500">{{ formatDate(post.publishedAt || post.createdAt) }}</span>
        <span class="text-sm text-gray-400">· {{ $t('blog.list.views', { count: post.viewCount }) }}</span>
      </div>

      <!-- Title -->
      <h1 class="text-3xl sm:text-4xl font-bold text-gray-900 mb-4 leading-tight">{{ post.title }}</h1>

      <!-- Summary -->
      <p v-if="post.summary" class="text-lg text-gray-600 mb-6 border-l-4 border-indigo-400 pl-4 italic">{{ post.summary }}</p>

      <!-- Author -->
      <div class="flex items-center gap-2 mb-8 pb-6 border-b border-gray-200">
        <div class="w-8 h-8 rounded-full bg-indigo-100 flex items-center justify-center text-indigo-600 font-bold text-sm">
          {{ post.authorName?.[0]?.toUpperCase() || 'A' }}
        </div>
        <span class="text-sm text-gray-700 font-medium">{{ post.authorName }}</span>
      </div>

      <!-- Tags -->
      <div v-if="post.tags" class="mb-6 flex flex-wrap gap-2">
        <span v-for="tag in parseTags(post.tags)" :key="tag" class="text-sm bg-gray-100 text-gray-600 px-3 py-1 rounded-full">
          #{{ tag }}
        </span>
      </div>

      <!-- Content rendered as Markdown -->
      <div class="prose prose-indigo max-w-none text-gray-800 leading-relaxed" v-html="renderedContent"></div>
    </article>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { marked } from 'marked'
import DOMPurify from 'dompurify'
import blogService from '../services/blogService'
import { BLOG_CATEGORIES } from '../models/blog.model'
import type { BlogPost } from '../models/blog.model'

const route = useRoute()
const { locale } = useI18n()
const post = ref<BlogPost | null>(null)
const loading = ref(true)
const categories = BLOG_CATEGORIES

const renderedContent = computed(() => {
  if (!post.value?.content) return ''
  const html = marked.parse(post.value.content, { async: false }) as string
  return DOMPurify.sanitize(html)
})

function updateMetaTags() {
  if (!post.value) return
  document.title = `${post.value.title} | Nerdlet Blog`
  setMeta('og:title', post.value.title)
  setMeta('og:description', post.value.summary || post.value.title)
  setMeta('og:type', 'article')
  if (post.value.coverImage) {
    setMeta('og:image', post.value.coverImage)
  }
  setMeta('og:url', window.location.href)
  setMeta('description', post.value.summary || post.value.title)
}

function setMeta(name: string, content: string) {
  const attr = name.startsWith('og:') ? 'property' : 'name'
  let el = document.querySelector(`meta[${attr}="${name}"]`) as HTMLMetaElement | null
  if (!el) {
    el = document.createElement('meta')
    el.setAttribute(attr, name)
    document.head.appendChild(el)
  }
  el.setAttribute('content', content)
}

onMounted(async () => {
  const slug = route.params.slug as string
  try {
    const res = await blogService.getPostBySlug(slug)
    if (res.code === 0 && res.data) {
      post.value = res.data
    }
  } finally {
    loading.value = false
  }
})

watch(post, updateMetaTags, { immediate: true })

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

function formatDate(dateStr?: string) {
  if (!dateStr) return ''
  return new Date(dateStr).toLocaleDateString()
}
function parseTags(tags: string) {
  return tags.split(',').map(t => t.trim()).filter(Boolean)
}
</script>
