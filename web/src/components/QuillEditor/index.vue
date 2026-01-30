<script setup lang="ts">
import { shallowRef, toRaw, ref, watch, onMounted, onBeforeUnmount } from 'vue'
import Quill from 'quill'
import { CONTENT_CHANGE_EVENT, DEFAULT_EDITOR_HEIGHT, DEFAULT_EDITOR_MIN_HEIGHT, DEFAULT_PLACEHOLDER } from '../../const/editor'
import 'quill/dist/quill.snow.css'
import getApiClient from '../../services/apiClient'
import { showError } from '../../utils/notification'

interface Props {
  modelValue: string
  height?: string
  minHeight?: string
  placeholder?: string
  readOnly?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  height: DEFAULT_EDITOR_HEIGHT,
  minHeight: DEFAULT_EDITOR_MIN_HEIGHT,
  placeholder: DEFAULT_PLACEHOLDER,
  readOnly: false
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
  (e: 'change', value: string): void
}>()

const editorEl = ref<HTMLElement>()
const quill = shallowRef<Quill>()

// Image Handler
const imageHandler = () => {
  const input = document.createElement('input')
  input.setAttribute('type', 'file')
  input.setAttribute('accept', 'image/*')
  input.click()

  input.onchange = async () => {
    const file = input.files ? input.files[0] : null
    if (!file) return

    const formData = new FormData()
    formData.append('file', file)

    try {
      const client = await getApiClient()
      const response = await client.post('/api/v1/upload/image', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      })
      
      const attachment = response.data.data
      const url = attachment.path
      
      if (quill.value) {
        // Use toRaw to avoid Vue Proxy interfering with Quill's internal state
        const q = toRaw(quill.value)
        
        // Focus the editor first to ensure internal selection state is valid
        q.focus()

        // Use setTimeout to allow the focus event to process and selection to update
        setTimeout(() => {
          if (!quill.value) return
          // Re-acquire raw instance inside timeout
          const q = toRaw(quill.value)
          
          let index = 0
          
          try {
            // getSelection can throw "reading 'offset'" error if DOM selection is invalid
            // even after focus(). We must try-catch it.
            const range = q.getSelection()
            if (range) {
              index = range.index
            } else {
              index = q.getLength()
            }
          } catch (e) {
            // Fallback to inserting at the end if selection retrieval fails
            index = q.getLength()
          }
          
          q.insertEmbed(index, 'image', url)
          q.setSelection(index + 1, 0)
        }, 0)
      }
    } catch (error: any) {
      console.error('Image upload failed', error)
      // Show specific error if available to help debugging
      showError(error.message || 'Image upload failed')
    }
  }
}

// Initialize Quill editor
onMounted(() => {
  if (!editorEl.value) return

  const option = {
    modules: {
      toolbar: null as any, // Disable toolbar for read-only mode
    },
    theme: 'snow',
    readOnly: props.readOnly,
    placeholder: props.placeholder,
    bounds: 'self'
  }

  if (!props.readOnly) {
    option.modules.toolbar = {
      container: [
        ['bold', 'italic', 'underline', 
        // 'strike'
        ],
        // ['blockquote', 'code-block'],
        // [{ 'header': 1 }, { 'header': 2 }],
        [{ 'list': 'ordered' }, { 'list': 'bullet' }],
        [{ 'script': 'sub'}, { 'script': 'super' }],
        ['table'],
        [{ 'indent': '-1' }, { 'indent': '+1' }],
        [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
        [{ 'color': [] }, { 'background': [] }],
        [{ 'align': [] }],
        ['link', 'image', 'formula']
      ],
      handlers: {
        image: imageHandler
      }
    }
  }

  quill.value = new Quill(editorEl.value, option)

  // Set initial content
  quill.value.root.innerHTML = props.modelValue

  // Listen for content changes
  quill.value.on(CONTENT_CHANGE_EVENT, () => {
    const html = quill.value?.root.innerHTML || ''
    if (html !== props.modelValue) {
      emit('update:modelValue', html)
      emit('change', html)
    }
  })
})

// Watch for external changes to modelValue
watch(() => props.modelValue, (newValue) => {
  if (quill.value && quill.value.root.innerHTML !== newValue) {
    quill.value.root.innerHTML = newValue
  }
})

// Clean up when component is destroyed
onBeforeUnmount(() => {
  quill.value?.off('text-change')
})
</script>

<template>
  <div class="quill-editor-container">
    <div ref="editorEl" :style="{ height, minHeight }" />
  </div>
</template>

<style scoped>
.quill-editor-container {
  width: 100%;
}

:deep(.ql-editor) {
  line-height: 1.6;
  font-size: 14px;
  min-height: v-bind(minHeight);
  height: v-bind(height);
}
</style>
