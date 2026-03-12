<template>
  <div class="p-6">
    <header class="mb-6 flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold text-gray-900">{{ isEdit ? $t('blog.management.editPost') : $t('blog.management.createPost') }}</h1>
        <p class="mt-1 text-sm text-gray-500">{{ $t('blog.management.subtitle') }}</p>
      </div>
      <button @click="goBack" class="text-sm text-indigo-600 hover:text-indigo-800">
        &larr; {{ $t('common.back') }}
      </button>
    </header>

    <div v-if="loadingPost" class="text-center py-20 text-gray-500">{{ $t('common.loading') }}</div>

    <div v-else class="bg-white shadow sm:rounded-lg">
      <form @submit.prevent="handleSubmit" class="px-6 py-6 space-y-6">
        <!-- Error message -->
        <div v-if="errorMessage" class="p-3 bg-red-100 text-red-700 rounded-md text-sm">{{ errorMessage }}</div>

        <!-- Title -->
        <div>
          <label for="title" class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.title') }} <span class="text-red-500">*</span></label>
          <input
            id="title"
            v-model="form.title"
            type="text"
            required
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
        </div>

        <!-- Slug -->
        <div>
          <label for="slug" class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.slug') }}</label>
          <input
            id="slug"
            v-model="form.slug"
            type="text"
            :placeholder="$t('blog.management.slugPlaceholder')"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
        </div>

        <!-- Category + Status row -->
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <div>
            <label for="category" class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.category') }} <span class="text-red-500">*</span></label>
            <select
              id="category"
              v-model="form.category"
              required
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            >
              <option value="">{{ $t('blog.management.selectCategory') }}</option>
              <option v-for="cat in categories" :key="cat.value" :value="cat.value">
                {{ $i18n.locale === 'zh' ? cat.labelZh : cat.labelEn }}
              </option>
            </select>
          </div>
          <div>
            <label for="status" class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.status') }}</label>
            <select
              id="status"
              v-model="form.status"
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            >
              <option value="draft">{{ $t('blog.status.draft') }}</option>
              <option value="published">{{ $t('blog.status.published') }}</option>
            </select>
          </div>
        </div>

        <!-- Tags -->
        <div>
          <label for="tags" class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.tags') }}</label>
          <input
            id="tags"
            v-model="form.tags"
            type="text"
            :placeholder="$t('blog.management.tagsPlaceholder')"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
        </div>

        <!-- Cover Image -->
        <div>
          <label for="coverImage" class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.coverImage') }}</label>
          <input
            id="coverImage"
            v-model="form.coverImage"
            type="text"
            :placeholder="$t('blog.management.coverImagePlaceholder')"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
        </div>

        <!-- Summary -->
        <div>
          <label for="summary" class="block text-sm font-medium text-gray-700">{{ $t('blog.fields.summary') }}</label>
          <textarea
            id="summary"
            v-model="form.summary"
            rows="3"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          ></textarea>
        </div>

        <!-- Content (rich text editor) -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">{{ $t('blog.fields.content') }} <span class="text-red-500">*</span></label>
          <div class="border border-gray-300 rounded-md overflow-hidden">
            <QuillEditor v-model="form.content" height="400px" minHeight="300px" />
          </div>
        </div>

        <!-- Actions -->
        <div class="flex justify-end gap-3 pt-4 border-t border-gray-200">
          <button
            type="button"
            @click="goBack"
            class="px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none"
          >
            {{ $t('common.cancel') }}
          </button>
          <button
            type="submit"
            :disabled="saving"
            class="px-4 py-2 border border-transparent rounded-md text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none disabled:opacity-50"
          >
            {{ saving ? $t('blog.management.saving') : $t('common.save') }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import QuillEditor from '../../components/QuillEditor/index.vue'
import blogService from '../../services/blogService'
import { BLOG_CATEGORIES } from '../../models/blog.model'

const route = useRoute()
const router = useRouter()
const { t } = useI18n()

const isEdit = computed(() => !!route.params.id)
const categories = BLOG_CATEGORIES

const loadingPost = ref(false)
const saving = ref(false)
const errorMessage = ref('')

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

onMounted(async () => {
  if (isEdit.value) {
    loadingPost.value = true
    try {
      const res = await blogService.getPostById(Number(route.params.id))
      if (res.code === 0 && res.data) {
        const post = res.data
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
      } else {
        errorMessage.value = res.message || t('common.error')
      }
    } catch {
      errorMessage.value = t('common.error')
    } finally {
      loadingPost.value = false
    }
  }
})

async function handleSubmit() {
  if (!form.value.title || !form.value.category || !form.value.content) {
    errorMessage.value = t('blog.management.requiredFields')
    return
  }
  saving.value = true
  errorMessage.value = ''
  try {
    let res
    if (isEdit.value) {
      res = await blogService.updatePost(form.value)
    } else {
      const { id: _, ...createData } = form.value
      res = await blogService.createPost(createData)
    }
    if (res.code === 0) {
      router.push('/admin/blog')
    } else {
      errorMessage.value = res.message || t('blog.management.saveFailed')
    }
  } catch {
    errorMessage.value = t('blog.management.saveFailed')
  } finally {
    saving.value = false
  }
}

function goBack() {
  router.push('/admin/blog')
}
</script>
