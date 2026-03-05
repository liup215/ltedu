<template>
  <div class="flex h-screen pt-16 bg-gray-50">
    <!-- Sidebar: session list -->
    <div class="w-64 bg-white border-r border-gray-200 flex flex-col flex-shrink-0">
      <div class="p-4 border-b border-gray-200">
        <button
          @click="startNewSession"
          class="w-full px-4 py-2 bg-indigo-600 text-white rounded-lg font-normal hover:bg-indigo-700 transition text-sm"
        >
          + {{ $t('aiChat.newChat') }}
        </button>
      </div>
      <div class="flex-1 overflow-y-auto">
        <div v-if="sessions.length === 0" class="p-4 text-gray-500 text-sm text-center">
          {{ $t('aiChat.noSessions') }}
        </div>
        <div
          v-for="session in sessions"
          :key="session.id"
          @click="selectSession(session.id)"
          class="p-3 border-b border-gray-100 cursor-pointer hover:bg-gray-50 transition"
          :class="{ 'bg-indigo-50 border-l-4 border-l-indigo-500': currentSessionId === session.id }"
        >
          <div class="text-sm font-medium text-gray-900 truncate">{{ session.title }}</div>
          <div class="text-xs text-gray-500 mt-1">{{ formatDate(session.updatedAt) }}</div>
          <div class="text-xs text-gray-400">
            {{ session.messageCount }} {{ $t('aiChat.messages') }}
            <span v-if="!session.isActive" class="ml-1 text-red-400">({{ $t('aiChat.closed') }})</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Main chat area -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <!-- Chat header -->
      <div class="bg-white border-b border-gray-200 px-6 py-4 flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">
            {{ currentSession?.title || $t('aiChat.title') }}
          </h2>
          <p class="text-sm text-gray-500">{{ $t('aiChat.subtitle') }}</p>
        </div>
        <div v-if="currentSessionId" class="flex gap-2">
          <button
            @click="resetSession"
            class="px-3 py-1 text-sm text-gray-600 border border-gray-300 rounded hover:bg-gray-100 transition"
          >
            {{ $t('aiChat.clearHistory') }}
          </button>
        </div>
      </div>

      <!-- Messages area -->
      <div ref="messagesContainer" class="flex-1 overflow-y-auto p-6 space-y-4">
        <!-- Welcome message when no session -->
        <div v-if="!currentSessionId" class="flex flex-col items-center justify-center h-full text-center">
          <div class="text-6xl mb-4">🤖</div>
          <h3 class="text-xl font-semibold text-gray-900 mb-2">{{ $t('aiChat.welcomeTitle') }}</h3>
          <p class="text-gray-500 max-w-md">{{ $t('aiChat.welcomeDesc') }}</p>
          <button
            @click="startNewSession"
            class="mt-6 px-6 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 transition"
          >
            {{ $t('aiChat.startChat') }}
          </button>
        </div>

        <!-- Messages -->
        <template v-if="currentSessionId">
          <div v-if="messages.length === 0 && !loading" class="text-center text-gray-400 py-8">
            {{ $t('aiChat.startTyping') }}
          </div>
          <div
            v-for="msg in messages"
            :key="msg.id"
            class="flex"
            :class="msg.role === 'user' ? 'justify-end' : 'justify-start'"
          >
            <!-- Assistant avatar -->
            <div v-if="msg.role === 'assistant'" class="flex items-start gap-3 max-w-3xl">
              <div class="w-8 h-8 rounded-full bg-indigo-100 flex items-center justify-center flex-shrink-0 mt-1">
                🤖
              </div>
              <div class="bg-white border border-gray-200 rounded-2xl rounded-tl-none px-4 py-3 shadow-sm">
                <div class="text-gray-800 text-sm whitespace-pre-wrap leading-relaxed">{{ msg.content }}</div>
                <div class="text-xs text-gray-400 mt-1">{{ formatTime(msg.createdAt) }}</div>
              </div>
            </div>

            <!-- User message -->
            <div v-else class="flex items-start gap-3 max-w-3xl flex-row-reverse">
              <div class="w-8 h-8 rounded-full bg-indigo-600 flex items-center justify-center flex-shrink-0 mt-1 text-white text-sm font-semibold">
                {{ userInitial }}
              </div>
              <div class="bg-indigo-600 text-white rounded-2xl rounded-tr-none px-4 py-3 shadow-sm">
                <div class="text-sm whitespace-pre-wrap leading-relaxed">{{ msg.content }}</div>
                <div class="text-xs text-indigo-200 mt-1">{{ formatTime(msg.createdAt) }}</div>
              </div>
            </div>
          </div>

          <!-- Loading indicator -->
          <div v-if="sending" class="flex justify-start">
            <div class="flex items-start gap-3">
              <div class="w-8 h-8 rounded-full bg-indigo-100 flex items-center justify-center flex-shrink-0">
                🤖
              </div>
              <div class="bg-white border border-gray-200 rounded-2xl rounded-tl-none px-4 py-3 shadow-sm">
                <div class="flex gap-1">
                  <span class="w-2 h-2 bg-gray-400 rounded-full animate-bounce" style="animation-delay: 0ms"></span>
                  <span class="w-2 h-2 bg-gray-400 rounded-full animate-bounce" style="animation-delay: 150ms"></span>
                  <span class="w-2 h-2 bg-gray-400 rounded-full animate-bounce" style="animation-delay: 300ms"></span>
                </div>
              </div>
            </div>
          </div>
        </template>
      </div>

      <!-- Input area -->
      <div v-if="currentSessionId" class="bg-white border-t border-gray-200 p-4">
        <div class="flex gap-3 items-end">
          <textarea
            v-model="inputMessage"
            @keydown.enter.exact.prevent="sendMessage"
            :placeholder="$t('aiChat.inputPlaceholder')"
            :disabled="sending || !currentSession?.isActive"
            rows="1"
            class="flex-1 px-4 py-3 border border-gray-300 rounded-xl resize-none focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent disabled:bg-gray-100 text-sm"
            style="max-height: 120px; overflow-y: auto"
            @input="autoResize"
            ref="textareaRef"
          ></textarea>
          <button
            @click="sendMessage"
            :disabled="!inputMessage.trim() || sending || !currentSession?.isActive"
            class="px-4 py-3 bg-indigo-600 text-white rounded-xl hover:bg-indigo-700 disabled:opacity-50 disabled:cursor-not-allowed transition font-normal text-sm flex-shrink-0"
          >
            {{ $t('aiChat.send') }}
          </button>
        </div>
        <p v-if="currentSession && !currentSession.isActive" class="text-xs text-red-400 mt-2">
          {{ $t('aiChat.sessionClosed') }}
        </p>
        <p class="text-xs text-gray-400 mt-2">{{ $t('aiChat.enterHint') }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, nextTick, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { useUserStore } from '../stores/userStore'
import conversationService from '../services/conversationService'
import type { ConversationSession, ConversationMessage } from '../models/conversation.model'

const { t } = useI18n()
const userStore = useUserStore()

const sessions = ref<ConversationSession[]>([])
const messages = ref<ConversationMessage[]>([])
const currentSessionId = ref<number | null>(null)
const inputMessage = ref('')
const loading = ref(false)
const sending = ref(false)
const messagesContainer = ref<HTMLElement | null>(null)
const textareaRef = ref<HTMLTextAreaElement | null>(null)

const currentSession = computed(() =>
  sessions.value.find((s) => s.id === currentSessionId.value) || null,
)

const userInitial = computed(() =>
  userStore.user?.username?.charAt(0).toUpperCase() || 'U',
)

function formatDate(dateStr: string): string {
  const d = new Date(dateStr)
  return d.toLocaleDateString()
}

function formatTime(dateStr: string): string {
  const d = new Date(dateStr)
  return d.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
}

async function loadSessions() {
  try {
    const res = await conversationService.listSessions()
    sessions.value = res.data?.list || []
  } catch {
    sessions.value = []
  }
}

async function startNewSession() {
  try {
    const res = await conversationService.startSession({ title: t('aiChat.newChat') })
    const session = res.data
    sessions.value.unshift(session)
    await selectSession(session.id)
  } catch {
    // ignore
  }
}

async function selectSession(id: number) {
  currentSessionId.value = id
  messages.value = []
  loading.value = true
  try {
    const res = await conversationService.getHistory(id)
    messages.value = res.data?.list || []
    await scrollToBottom()
  } finally {
    loading.value = false
  }
}

async function sendMessage() {
  if (!inputMessage.value.trim() || sending.value || !currentSessionId.value) return
  const content = inputMessage.value.trim()
  inputMessage.value = ''
  resetTextareaHeight()

  // Optimistically add user message
  const optimisticUserMsg: ConversationMessage = {
    id: Date.now(),
    sessionId: currentSessionId.value,
    role: 'user',
    content,
    createdAt: new Date().toISOString(),
  }
  messages.value.push(optimisticUserMsg)
  await scrollToBottom()

  sending.value = true
  try {
    const res = await conversationService.sendMessage({
      sessionId: currentSessionId.value,
      content,
    })
    messages.value.push(res.data)
    // Update session title/count in sidebar
    await loadSessions()
    await scrollToBottom()
  } catch {
    // Remove optimistic message on error
    messages.value.pop()
  } finally {
    sending.value = false
  }
}

async function resetSession() {
  if (!currentSessionId.value) return
  if (!confirm(t('aiChat.confirmClear'))) return
  try {
    await conversationService.resetSession(currentSessionId.value)
    messages.value = []
    await loadSessions()
  } catch {
    // ignore
  }
}

async function scrollToBottom() {
  await nextTick()
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}

function autoResize(e: Event) {
  const el = e.target as HTMLTextAreaElement
  el.style.height = 'auto'
  el.style.height = Math.min(el.scrollHeight, 120) + 'px'
}

function resetTextareaHeight() {
  if (textareaRef.value) {
    textareaRef.value.style.height = 'auto'
  }
}

watch(currentSessionId, async () => {
  await scrollToBottom()
})

onMounted(async () => {
  await loadSessions()
})
</script>
