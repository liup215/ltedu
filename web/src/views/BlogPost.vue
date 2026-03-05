<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-3xl mx-auto px-4 py-12">
      <!-- Back button -->
      <button
        @click="router.back()"
        class="flex items-center gap-2 text-gray-500 hover:text-indigo-600 transition mb-6 text-sm"
      >
        ← {{ $t('blog.back') }}
      </button>

      <!-- Loading -->
      <div v-if="loading" class="text-center py-16 text-gray-400">
        <div class="animate-spin w-8 h-8 border-4 border-indigo-600 border-t-transparent rounded-full mx-auto mb-4"></div>
        {{ $t('common.loading') }}
      </div>

      <!-- Not found -->
      <div v-else-if="!post" class="text-center py-16">
        <div class="text-5xl mb-4">📄</div>
        <p class="text-gray-500">{{ $t('blog.notFound') }}</p>
      </div>

      <!-- Post content -->
      <article v-else class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
        <!-- Cover image -->
        <div v-if="post.coverImage" class="w-full h-64 bg-gray-100">
          <img :src="post.coverImage" :alt="post.title" class="w-full h-full object-cover" />
        </div>

        <div class="p-8">
          <!-- Category + badges -->
          <div class="flex flex-wrap items-center gap-2 mb-4">
            <span
              v-if="post.isTop"
              class="px-2 py-0.5 bg-yellow-100 text-yellow-700 text-xs rounded-full"
            >
              📌 {{ $t('blog.pinned') }}
            </span>
            <span
              v-if="post.category"
              class="px-3 py-1 bg-indigo-100 text-indigo-700 text-sm rounded-full"
            >
              {{ post.category }}
            </span>
          </div>

          <!-- Title -->
          <h1 class="text-3xl font-bold text-gray-900 mb-4 leading-tight">{{ post.title }}</h1>

          <!-- Meta -->
          <div class="flex items-center gap-4 text-sm text-gray-400 mb-6 pb-6 border-b border-gray-100">
            <span>{{ formatDate(post.createdAt) }}</span>
            <span>👁 {{ post.viewCount }} {{ $t('blog.views') }}</span>
          </div>

          <!-- Tags -->
          <div v-if="post.tags" class="flex flex-wrap gap-2 mb-6">
            <span
              v-for="tag in post.tags.split(',').filter(t => t.trim())"
              :key="tag"
              class="px-2 py-0.5 bg-gray-100 text-gray-600 text-xs rounded"
            >
              # {{ tag.trim() }}
            </span>
          </div>

          <!-- Summary -->
          <div v-if="post.summary" class="bg-indigo-50 border-l-4 border-indigo-400 px-4 py-3 rounded mb-6 text-gray-700 text-sm italic">
            {{ post.summary }}
          </div>

          <!-- Content -->
          <div
            class="prose max-w-none text-gray-800 leading-relaxed"
            v-html="renderedContent"
          ></div>
        </div>
      </article>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import blogService from '../services/blogService'
import type { BlogPost } from '../models/blog.model'

const route = useRoute()
const router = useRouter()
const { t } = useI18n()

const post = ref<BlogPost | null>(null)
const loading = ref(false)

const renderedContent = computed(() => {
  if (!post.value?.content) return ''
  // Sanitize HTML to prevent XSS: strip dangerous tags and attributes
  return sanitizeHtml(post.value.content)
})

/**
 * Basic HTML sanitizer that removes dangerous tags and event handler attributes.
 * This guards against stored XSS in blog content.
 */
function sanitizeHtml(html: string): string {
  const div = document.createElement('div')
  div.innerHTML = html
  const dangerousTags = ['script', 'iframe', 'object', 'embed', 'form', 'input', 'button', 'link', 'meta', 'style']
  dangerousTags.forEach((tag) => {
    div.querySelectorAll(tag).forEach((el) => el.remove())
  })

  // URL schemes that can execute scripts
  const dangerousSchemes = /^(javascript|vbscript|data):/i

  // Remove event handler attributes (on*) and dangerous URL attributes
  const urlAttributes = new Set(['href', 'src', 'action', 'formaction', 'xlink:href'])
  div.querySelectorAll('*').forEach((el) => {
    Array.from(el.attributes).forEach((attr) => {
      if (attr.name.startsWith('on')) {
        el.removeAttribute(attr.name)
      } else if (urlAttributes.has(attr.name.toLowerCase()) && dangerousSchemes.test(attr.value.trim())) {
        el.removeAttribute(attr.name)
      }
    })
  })
  return div.innerHTML
}

async function loadPost() {
  const id = Number(route.params.id)
  if (!id) return

  loading.value = true
  try {
    const res = await blogService.getPost(id)
    post.value = res.data || null
  } catch {
    post.value = null
  } finally {
    loading.value = false
  }
}

function formatDate(dateStr: string): string {
  const d = new Date(dateStr)
  return d.toLocaleDateString(undefined, {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}

onMounted(() => {
  loadPost()
})
</script>
