<script setup lang="ts">
import { ref, nextTick, computed } from 'vue'
import { apiRequest, cacheGet, cacheSet } from '../services/tauriClient'
import { useAuthStore } from '../stores/authStore'

interface Message {
  id: string
  role: 'user' | 'assistant' | 'system'
  content: string
  timestamp: Date
  loading?: boolean
}

const authStore = useAuthStore()
const messages = ref<Message[]>([
  {
    id: 'welcome',
    role: 'assistant',
    content: `👋 Hello${authStore.user ? ', ' + authStore.user.username : ''}! I'm your LTEdu AI assistant.\n\nI can help you with:\n• 📚 **Practice questions** – find questions by topic or chapter\n• 📝 **Exam papers** – browse past papers and build custom exams\n• 📊 **Analytics** – view your class performance and trends\n• 🗺 **Learning plans** – manage and generate personalised study plans\n\nWhat would you like to do today?`,
    timestamp: new Date(),
  },
])
const inputText = ref('')
const chatContainer = ref<HTMLElement>()
const isTyping = ref(false)

const formattedMessages = computed(() =>
  messages.value.map(m => ({
    ...m,
    html: m.content
      .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
      .replace(/\n/g, '<br>'),
  }))
)

async function scrollToBottom() {
  await nextTick()
  if (chatContainer.value) {
    chatContainer.value.scrollTop = chatContainer.value.scrollHeight
  }
}

function addMessage(role: Message['role'], content: string, loading = false): Message {
  const msg: Message = {
    id: crypto.randomUUID(),
    role,
    content,
    timestamp: new Date(),
    loading,
  }
  messages.value.push(msg)
  scrollToBottom()
  return msg
}

/** Simple intent → API route resolution */
async function resolveIntent(text: string): Promise<string> {
  const lower = text.toLowerCase()

  if (/practice|question|quiz|exercise/.test(lower)) {
    // Try to retrieve from cache first
    const cacheKey = 'quick_practice_subjects'
    const cached = await cacheGet(cacheKey)
    if (cached) {
      return `Here are the cached subjects for quick practice:\n\n${cached}\n\n*Data served from local cache.*`
    }
    const result = await apiRequest({ path: '/v1/syllabus', method: 'GET' })
    const body = result.body as { data?: { list?: Array<{ name: string }> } }
    const items: Array<{ name: string }> = body?.data?.list ?? []
    const names = items.map((s) => `• ${s.name}`).join('\n') || 'No syllabuses found.'
    await cacheSet(cacheKey, names)
    return `Available syllabuses for practice:\n\n${names}\n\nTell me which subject you'd like to practise and I'll find questions for you!`
  }

  if (/paper|exam/.test(lower)) {
    const result = await apiRequest({ path: '/v1/paper/series', method: 'GET' })
    const body = result.body as { data?: { list?: Array<{ name: string }> } }
    const items: Array<{ name: string }> = body?.data?.list ?? []
    const names = items.map((s) => `• ${s.name}`).join('\n') || 'No paper series found.'
    return `Available exam paper series:\n\n${names}`
  }

  if (/analytic|performance|trend|warning/.test(lower)) {
    return `📊 **Analytics Features**\n\nI can pull analytics data for your classes. Please specify:\n• A class ID, or\n• "class summary", "student list", "heatmap", "trends", or "early warning"\n\nExample: *"Show me the heatmap for class 42"*`
  }

  if (/learning plan|study plan|phase plan/.test(lower)) {
    return `🗺 **Learning Plans**\n\nI can help you manage learning plans. Try:\n• *"List learning plans for class 5"*\n• *"Generate a template plan"*`
  }

  if (/help|what can you do/.test(lower)) {
    return `Here's what I can help you with:\n\n📚 **Practice** – find & start practice sessions\n📝 **Exam Papers** – browse past paper series\n📊 **Analytics** – class performance insights\n🗺 **Learning Plans** – generate and manage plans\n⚙ **Settings** – configure server URL and cache`
  }

  // Default: forward as a general query
  return `I understood your message. To fetch live data, please be more specific — for example:\n• *"Show me available syllabuses"*\n• *"List past papers"*\n• *"Show analytics for class 5"*`
}

async function sendMessage() {
  const text = inputText.value.trim()
  if (!text || isTyping.value) return

  inputText.value = ''
  addMessage('user', text)

  isTyping.value = true
  const assistantMsg = addMessage('assistant', '', true)

  try {
    const reply = await resolveIntent(text)
    const idx = messages.value.findIndex(m => m.id === assistantMsg.id)
    if (idx !== -1) {
      messages.value[idx].content = reply
      messages.value[idx].loading = false
    }
  } catch (err: unknown) {
    const idx = messages.value.findIndex(m => m.id === assistantMsg.id)
    if (idx !== -1) {
      messages.value[idx].content = `⚠ Error: ${err instanceof Error ? err.message : String(err)}`
      messages.value[idx].loading = false
    }
  } finally {
    isTyping.value = false
    scrollToBottom()
  }
}

function handleKeydown(event: KeyboardEvent) {
  if (event.key === 'Enter' && !event.shiftKey) {
    event.preventDefault()
    sendMessage()
  }
}

function clearChat() {
  messages.value = []
  addMessage('assistant', '🧹 Chat cleared. How can I help you?')
}
</script>

<template>
  <div class="chat-interface">
    <header class="chat-header">
      <div class="header-title">
        <span class="ai-icon">🤖</span>
        <div>
          <h2>AI Assistant</h2>
          <span class="status">{{ isTyping ? 'Typing…' : 'Online' }}</span>
        </div>
      </div>
      <button class="btn-icon" title="Clear chat" @click="clearChat">🗑</button>
    </header>

    <div ref="chatContainer" class="messages-container">
      <div
        v-for="msg in formattedMessages"
        :key="msg.id"
        :class="['message', msg.role]"
      >
        <div class="bubble">
          <div v-if="msg.loading" class="typing-dots">
            <span></span><span></span><span></span>
          </div>
          <!-- eslint-disable vue/no-v-html -->
          <span v-else v-html="msg.html"></span>
        </div>
        <span class="time">{{ msg.timestamp.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) }}</span>
      </div>
    </div>

    <div class="suggestions">
      <button
        v-for="hint in ['Practice questions', 'Browse past papers', 'View analytics', 'Learning plans']"
        :key="hint"
        class="suggestion-chip"
        @click="inputText = hint; sendMessage()"
      >
        {{ hint }}
      </button>
    </div>

    <div class="input-area">
      <textarea
        v-model="inputText"
        placeholder="Ask anything about your studies…"
        rows="1"
        @keydown="handleKeydown"
      ></textarea>
      <button class="send-btn" :disabled="!inputText.trim() || isTyping" @click="sendMessage">
        ➤
      </button>
    </div>
  </div>
</template>

<style scoped>
.chat-interface {
  display: flex;
  flex-direction: column;
  height: 100%;
  background: #f7fafc;
}

.chat-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px;
  background: white;
  border-bottom: 1px solid #e2e8f0;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}

.header-title {
  display: flex;
  align-items: center;
  gap: 12px;
}

.ai-icon {
  font-size: 28px;
}

.chat-header h2 {
  font-size: 16px;
  font-weight: 700;
  color: #1a202c;
}

.status {
  font-size: 12px;
  color: #48bb78;
  font-weight: 500;
}

.btn-icon {
  background: none;
  border: none;
  font-size: 18px;
  cursor: pointer;
  opacity: 0.5;
  transition: opacity 0.2s;
  padding: 4px 8px;
  border-radius: 6px;
}

.btn-icon:hover {
  opacity: 1;
  background: #f1f5f9;
}

.messages-container {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.message {
  display: flex;
  flex-direction: column;
  max-width: 75%;
}

.message.user {
  align-self: flex-end;
  align-items: flex-end;
}

.message.assistant {
  align-self: flex-start;
  align-items: flex-start;
}

.bubble {
  padding: 12px 16px;
  border-radius: 18px;
  font-size: 14px;
  line-height: 1.6;
  word-wrap: break-word;
}

.message.user .bubble {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-bottom-right-radius: 4px;
}

.message.assistant .bubble {
  background: white;
  color: #1a202c;
  border: 1px solid #e2e8f0;
  border-bottom-left-radius: 4px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}

.time {
  font-size: 11px;
  color: #a0aec0;
  margin-top: 4px;
  padding: 0 4px;
}

/* Typing indicator */
.typing-dots {
  display: flex;
  gap: 4px;
  align-items: center;
  padding: 4px 0;
}

.typing-dots span {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #a0aec0;
  animation: bounce 1.2s ease-in-out infinite;
}

.typing-dots span:nth-child(2) { animation-delay: 0.2s; }
.typing-dots span:nth-child(3) { animation-delay: 0.4s; }

@keyframes bounce {
  0%, 80%, 100% { transform: scale(0.8); opacity: 0.5; }
  40% { transform: scale(1.1); opacity: 1; }
}

/* Suggestions */
.suggestions {
  display: flex;
  gap: 8px;
  padding: 8px 20px;
  flex-wrap: wrap;
}

.suggestion-chip {
  background: white;
  border: 1px solid #e2e8f0;
  border-radius: 20px;
  padding: 6px 14px;
  font-size: 12px;
  color: #4a5568;
  cursor: pointer;
  transition: all 0.2s;
}

.suggestion-chip:hover {
  background: #667eea;
  color: white;
  border-color: #667eea;
}

/* Input */
.input-area {
  display: flex;
  align-items: flex-end;
  gap: 10px;
  padding: 12px 20px 16px;
  background: white;
  border-top: 1px solid #e2e8f0;
}

.input-area textarea {
  flex: 1;
  border: 1.5px solid #e2e8f0;
  border-radius: 12px;
  padding: 10px 14px;
  font-size: 14px;
  resize: none;
  outline: none;
  font-family: inherit;
  transition: border-color 0.2s;
  max-height: 120px;
  overflow-y: auto;
}

.input-area textarea:focus {
  border-color: #667eea;
}

.send-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 10px;
  width: 42px;
  height: 42px;
  font-size: 16px;
  cursor: pointer;
  transition: opacity 0.2s;
  flex-shrink: 0;
}

.send-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}
</style>
