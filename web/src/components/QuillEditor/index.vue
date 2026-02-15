<script setup lang="ts">
import { shallowRef, toRaw, ref, watch, onMounted, onBeforeUnmount } from 'vue'
import Quill from 'quill'
import QuillBetterTable from 'quill-better-table'
import { CONTENT_CHANGE_EVENT, DEFAULT_EDITOR_HEIGHT, DEFAULT_EDITOR_MIN_HEIGHT, DEFAULT_PLACEHOLDER } from '../../const/editor'
import 'quill/dist/quill.snow.css'
import 'quill-better-table/dist/quill-better-table.css'
import getApiClient from '../../services/apiClient'
import { showError } from '../../utils/notification'

// Register quill-better-table module
Quill.register({
  'modules/better-table': QuillBetterTable
}, true)

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

// Table Handler - Insert a 3x3 table
const tableHandler = () => {
  if (quill.value) {
    const q = toRaw(quill.value)
    const tableModule = q.getModule('better-table') as QuillBetterTable
    
    // Insert a 3x3 table
    tableModule.insertTable(3, 3)
    
    // Focus the editor
    q.focus()
  }
}

// Initialize Quill editor
onMounted(() => {
  if (!editorEl.value) return

  const option: any = {
    modules: {
      toolbar: false, // Disable toolbar for read-only mode
      'better-table': {
        operationMenu: {
          items: {
            insertColumnRight: { text: 'Insert Column Right' },
            insertColumnLeft: { text: 'Insert Column Left' },
            insertRowUp: { text: 'Insert Row Above' },
            insertRowDown: { text: 'Insert Row Below' },
            mergeCells: { text: 'Merge Cells' },
            unmergeCells: { text: 'Unmerge Cells' },
            deleteColumn: { text: 'Delete Column' },
            deleteRow: { text: 'Delete Row' },
            deleteTable: { text: 'Delete Table' }
          }
        }
      }
    },
    theme: 'snow',
    readOnly: props.readOnly,
    placeholder: props.placeholder,
    bounds: 'self'
  }

  if (!props.readOnly) {
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
        ['table-insert']  // Custom table button
      ],
      handlers: {
        image: imageHandler,
        'table-insert': tableHandler
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

/* Custom table button icon */
:deep(.ql-table-insert::before) {
  content: "📊";
  font-size: 14px;
}

/* Table styles */
:deep(.ql-editor table) {
  border-collapse: collapse;
  width: 100%;
}

:deep(.ql-editor table td) {
  border: 1px solid #ccc;
  padding: 8px;
  min-width: 50px;
}

:deep(.ql-editor table td:focus) {
  outline: 2px solid #007bff;
}
</style>