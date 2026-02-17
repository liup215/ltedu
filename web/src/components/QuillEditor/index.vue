<script setup lang="ts">
import { shallowRef, toRaw, ref, watch, onMounted, onBeforeUnmount } from 'vue'
import Quill from 'quill'
import QuillTableBetter from 'quill-table-better'
import { CONTENT_CHANGE_EVENT, DEFAULT_EDITOR_HEIGHT, DEFAULT_EDITOR_MIN_HEIGHT, DEFAULT_PLACEHOLDER } from '../../const/editor'
import 'quill/dist/quill.snow.css'
import 'quill-table-better/dist/quill-table-better.css'
import getApiClient from '../../services/apiClient'
import { showError } from '../../utils/notification'

// Register quill-table-better module
Quill.register({
  'modules/table-better': QuillTableBetter
}, true)

interface Props {
  modelValue: string
  height?: string
  minHeight?: string
  placeholder?: string
  readOnly?: boolean
  tableWidth?: 'auto' | 'fixed' | 'full'
}

const props = withDefaults(defineProps<Props>(), {
  height: DEFAULT_EDITOR_HEIGHT,
  minHeight: DEFAULT_EDITOR_MIN_HEIGHT,
  placeholder: DEFAULT_PLACEHOLDER,
  readOnly: false,
  tableWidth: 'auto'
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
  (e: 'change', value: string): void
}>()

const editorEl = ref<HTMLElement>()
const quill = shallowRef<Quill>()

// Store the paste handler reference for cleanup
const pasteHandlerRef = ref<((e: ClipboardEvent) => Promise<void>) | null>(null)

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

  const option: any = {
    modules: {
      toolbar: false, // Disable toolbar for read-only mode
      table: false // Disable default table module
    },
    theme: 'snow',
    readOnly: props.readOnly,
    placeholder: props.placeholder,
    bounds: 'self'
  }

  // Only configure table-better and keyboard bindings for editable mode
  if (!props.readOnly) {
    option.modules['table-better'] = {
      language: 'en_US',
      menus: ['column', 'row', 'merge', 'table', 'cell', 'wrap', 'delete'],
      toolbarTable: true // Let the module handle toolbar integration
    }
    option.modules.keyboard = {
      bindings: QuillTableBetter.keyboardBindings
    }
    option.modules.toolbar = {
      container: [
        ['bold', 'italic', 'underline'],
        [{ 'list': 'ordered' }, { 'list': 'bullet' }],
        [{ 'script': 'sub'}, { 'script': 'super' }],
        [{ 'indent': '-1' }, { 'indent': '+1' }],
        [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
        [{ 'color': [] }, { 'background': [] }],
        [{ 'align': [] }],
        ['link', 'image', 'formula'],
        ['table-better']  // Built-in table button from quill-table-better
      ],
      handlers: {
        image: imageHandler
      }
    }
  }

  quill.value = new Quill(editorEl.value, option)

  // Handle paste event to intercept image paste and upload to server
  const handlePaste = async (e: ClipboardEvent) => {
    const clipboardData = e.clipboardData
    if (!clipboardData) return

    const items = clipboardData.items
    for (let i = 0; i < items.length; i++) {
      const item = items[i]
      if (item.type.startsWith('image/')) {
        e.preventDefault() // Prevent default base64 insertion

        const file = item.getAsFile()
        if (file && quill.value) {
          const q = toRaw(quill.value)
          
          try {
            // Upload the image file
            const formData = new FormData()
            formData.append('file', file, 'pasted-image.png')

            const client = await getApiClient()
            const response = await client.post('/api/v1/upload/image', formData, {
              headers: {
                'Content-Type': 'multipart/form-data'
              }
            })

            const url = response.data.data.path

            // Insert the server URL as image
            q.focus()
            const range = q.getSelection()
            const index = range ? range.index : q.getLength()
            q.insertEmbed(index, 'image', url)
            q.setSelection(index + 1, 0)
          } catch (error: any) {
            console.error('Pasted image upload failed', error)
            showError(error.message || 'Image upload failed')
          }
        }
        break // Only handle the first image
      }
    }
  }

  // Add paste event listener with capture phase to intercept before Quill
  pasteHandlerRef.value = handlePaste
  quill.value.root.addEventListener('paste', handlePaste, true)

  // Set initial content
  quill.value.root.innerHTML = props.modelValue

  // Apply table width class
  updateTableWidthClass()

  // Listen for content changes
  quill.value.on(CONTENT_CHANGE_EVENT, () => {
    const html = quill.value?.root.innerHTML || ''
    if (html !== props.modelValue) {
      emit('update:modelValue', html)
      emit('change', html)
    }
  })
})

// Update table width class on editor root
const updateTableWidthClass = () => {
  if (!quill.value) return
  
  const root = quill.value.root
  // Remove existing table width classes
  root.classList.remove('table-width-auto', 'table-width-fixed', 'table-width-full')
  // Add the new class
  root.classList.add(`table-width-${props.tableWidth}`)
}

// Watch for external changes to modelValue
watch(() => props.modelValue, (newValue) => {
  if (quill.value && quill.value.root.innerHTML !== newValue) {
    quill.value.root.innerHTML = newValue
  }
})

// Watch for tableWidth prop changes
watch(() => props.tableWidth, () => {
  updateTableWidthClass()
})

// Clean up when component is destroyed
onBeforeUnmount(() => {
  quill.value?.off('text-change')
  
  // Remove paste event listener
  if (pasteHandlerRef.value && quill.value) {
    quill.value.root.removeEventListener('paste', pasteHandlerRef.value, true)
  }
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

/* Table width variations */
:deep(.ql-editor.table-width-full table) {
  width: 100%;
}

:deep(.ql-editor.table-width-fixed table) {
  table-layout: fixed;
}
</style>