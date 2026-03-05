<template>
  <div class="flex h-screen bg-gray-50">
    <!-- Sidebar: session list -->
    <aside class="w-64 bg-white border-r border-gray-200 flex flex-col">
      <div class="p-4 border-b border-gray-200">
        <h2 class="text-lg font-semibold text-gray-800">AI Assistant</h2>
        <p class="text-xs text-gray-500 mt-1">Multi-turn conversation</p>
      </div>

      <div class="p-3">
        <button
          @click="createNewSession"
          :disabled="loading"
          class="w-full py-2 px-4 bg-indigo-600 text-white rounded-lg text-sm hover:bg-indigo-700 disabled:opacity-50 transition"
        >
          + New Conversation
        </button>
      </div>

      <div class="flex-1 overflow-y-auto">
        <div v-if="sessions.length === 0" class="p-4 text-sm text-gray-400 text-center">
          No conversations yet
        </div>
        <button
          v-for="session in sessions"
          :key="session.sessionKey"
          @click="loadSession(session.sessionKey)"
          :class="[
            'w-full text-left px-4 py-3 border-b border-gray-100 hover:bg-gray-50 transition',
            activeSessionKey === session.sessionKey ? 'bg-indigo-50 border-l-4 border-l-indigo-500' : ''
          ]"
        >
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-700 truncate">
              {{ sessionLabel(session) }}
            </span>
            <span class="text-xs text-gray-400 ml-2 shrink-0">
              {{ session.messageCount }} msgs
            </span>
          </div>
          <div class="text-xs text-gray-400 mt-1">
            {{ formatDate(session.lastActiveAt) }}
          </div>
        </button>
      </div>
    </aside>

    <!-- Main chat area -->
    <div class="flex-1 flex flex-col">
      <!-- Header -->
      <div class="bg-white border-b border-gray-200 px-6 py-3 flex items-center justify-between">
        <div>
          <h3 class="font-semibold text-gray-800">
            {{ activeSession ? sessionLabel(activeSession) : 'Select or start a conversation' }}
          </h3>
          <p v-if="activeSession" class="text-xs text-gray-400">
            Role: {{ activeSession.userRole }} · {{ activeSession.messageCount }} messages ·
            Expires {{ formatDate(activeSession.expiresAt) }}
          </p>
        </div>
        <div v-if="activeSessionKey" class="flex gap-2">
          <button
            @click="resetCurrentContext"
            title="Reset context (start over)"
            class="px-3 py-1 text-xs bg-yellow-50 text-yellow-700 border border-yellow-200 rounded hover:bg-yellow-100 transition"
          >
            Reset Context
          </button>
          <button
            @click="closeCurrentSession"
            title="Close session"
            class="px-3 py-1 text-xs bg-red-50 text-red-700 border border-red-200 rounded hover:bg-red-100 transition"
          >
            Close
          </button>
        </div>
      </div>

      <!-- Messages -->
      <div ref="messagesContainer" class="flex-1 overflow-y-auto p-6 space-y-4">
        <div v-if="!activeSessionKey" class="h-full flex items-center justify-center text-gray-400">
          <div class="text-center">
            <div class="text-5xl mb-4">🤖</div>
            <p class="text-lg">Start a new conversation or select one from the sidebar.</p>
          </div>
        </div>

        <template v-else>
          <div v-if="messages.length === 0" class="text-center text-gray-400 py-12">
            <p>No messages yet. Type something below to start!</p>
          </div>

          <div
            v-for="msg in messages"
            :key="msg.id"
            :class="[
              'flex',
              msg.role === 'user' ? 'justify-end' : 'justify-start'
            ]"
          >
            <div
              :class="[
                'max-w-2xl rounded-2xl px-4 py-3 text-sm',
                msg.role === 'user'
                  ? 'bg-indigo-600 text-white rounded-br-sm'
                  : 'bg-white border border-gray-200 text-gray-800 rounded-bl-sm shadow-sm'
              ]"
            >
              <div class="whitespace-pre-wrap">{{ msg.content }}</div>
              <div
                :class="[
                  'text-xs mt-1 text-right',
                  msg.role === 'user' ? 'text-indigo-200' : 'text-gray-400'
                ]"
              >
                {{ formatTime(msg.createdAt) }}
              </div>
            </div>
          </div>

          <!-- Typing indicator -->
          <div v-if="sending" class="flex justify-start">
            <div class="bg-white border border-gray-200 rounded-2xl rounded-bl-sm px-4 py-3 shadow-sm">
              <div class="flex space-x-1">
                <span class="w-2 h-2 bg-gray-400 rounded-full animate-bounce" style="animation-delay:0ms"></span>
                <span class="w-2 h-2 bg-gray-400 rounded-full animate-bounce" style="animation-delay:150ms"></span>
                <span class="w-2 h-2 bg-gray-400 rounded-full animate-bounce" style="animation-delay:300ms"></span>
              </div>
            </div>
          </div>
        </template>
      </div>

      <!-- Error banner -->
      <div v-if="errorMsg" class="mx-6 mb-2 p-3 bg-red-50 text-red-700 text-sm rounded-lg border border-red-200 flex justify-between items-center">
        <span>{{ errorMsg }}</span>
        <button @click="errorMsg = ''" class="ml-2 text-red-400 hover:text-red-600">✕</button>
      </div>

      <!-- Input area -->
      <div class="bg-white border-t border-gray-200 px-6 py-4">
        <div class="flex gap-3 items-end">
          <textarea
            v-model="inputMessage"
            @keydown.enter.prevent="handleEnter"
            :disabled="!activeSessionKey || sending"
            placeholder="Type a message... (Enter to send, Shift+Enter for new line)"
            rows="2"
            class="flex-1 resize-none rounded-xl border border-gray-300 px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 disabled:bg-gray-50 disabled:text-gray-400"
          />
          <button
            @click="sendCurrentMessage"
            :disabled="!inputMessage.trim() || !activeSessionKey || sending"
            class="px-5 py-3 bg-indigo-600 text-white rounded-xl text-sm hover:bg-indigo-700 disabled:opacity-50 disabled:cursor-not-allowed transition"
          >
            Send
          </button>
        </div>
        <p class="text-xs text-gray-400 mt-2">
          Enter to send · Shift+Enter for new line · Type "reset" to clear context
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, nextTick } from 'vue';
import { conversationService } from '../services/conversationService';
import type { ConversationSession, ConversationMessage } from '../models/conversation.model';
import { useUserStore } from '../stores/userStore';

// Special command strings that trigger a context reset instead of sending a message.
const RESET_COMMANDS = ['reset', 'start over'];

const userStore = useUserStore();

const sessions = ref<ConversationSession[]>([]);
const activeSessionKey = ref<string | null>(null);
const messages = ref<ConversationMessage[]>([]);
const inputMessage = ref('');
const loading = ref(false);
const sending = ref(false);
const errorMsg = ref('');
const messagesContainer = ref<HTMLElement | null>(null);

const activeSession = computed(() =>
  sessions.value.find(s => s.sessionKey === activeSessionKey.value) ?? null,
);

function sessionLabel(session: ConversationSession): string {
  const role = session.userRole.charAt(0).toUpperCase() + session.userRole.slice(1);
  return `${role} Conversation`;
}

function formatDate(dateStr: string): string {
  if (!dateStr) return '';
  const d = new Date(dateStr);
  return d.toLocaleDateString(undefined, { month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' });
}

function formatTime(dateStr: string): string {
  if (!dateStr) return '';
  const d = new Date(dateStr);
  return d.toLocaleTimeString(undefined, { hour: '2-digit', minute: '2-digit' });
}

async function scrollToBottom() {
  await nextTick();
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight;
  }
}

async function loadSessionList() {
  try {
    const resp = await conversationService.listSessions();
    if (resp.code === 0) {
      sessions.value = resp.data ?? [];
    }
  } catch {
    // silently fail
  }
}

async function createNewSession() {
  loading.value = true;
  errorMsg.value = '';
  try {
    // Build initial context from current user info
    const user = userStore.user;
    const userRole = user?.isAdmin ? 'admin' : user?.isTeacher ? 'teacher' : 'student';
    const resp = await conversationService.startSession({ context: { userRole } });
    if (resp.code === 0 && resp.data) {
      await loadSessionList();
      await loadSession(resp.data.sessionKey);
    } else {
      errorMsg.value = resp.message || 'Failed to start session';
    }
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : 'Failed to start session';
  } finally {
    loading.value = false;
  }
}

async function loadSession(sessionKey: string) {
  activeSessionKey.value = sessionKey;
  messages.value = [];
  errorMsg.value = '';
  try {
    const resp = await conversationService.getHistory({ sessionKey });
    if (resp.code === 0) {
      messages.value = resp.data ?? [];
    }
  } catch {
    // silently fail
  }
  scrollToBottom();
}

async function sendCurrentMessage() {
  const text = inputMessage.value.trim();
  if (!text || !activeSessionKey.value || sending.value) return;

  // Handle explicit reset command
  if (RESET_COMMANDS.includes(text.toLowerCase())) {
    await resetCurrentContext();
    inputMessage.value = '';
    return;
  }

  sending.value = true;
  errorMsg.value = '';
  const optimisticUser: ConversationMessage = {
    id: Date.now(),
    sessionId: 0,
    role: 'user',
    content: text,
    orderIndex: messages.value.length,
    createdAt: new Date().toISOString(),
  };
  messages.value.push(optimisticUser);
  inputMessage.value = '';
  await scrollToBottom();

  try {
    const resp = await conversationService.sendMessage({
      sessionKey: activeSessionKey.value,
      message: text,
    });
    if (resp.code === 0 && resp.data) {
      // Replace optimistic message with the persisted one
      messages.value.pop();
      messages.value.push(resp.data.userMessage);
      messages.value.push(resp.data.assistantMessage);
      // Update session message count
      const idx = sessions.value.findIndex(s => s.sessionKey === activeSessionKey.value);
      if (idx !== -1) {
        sessions.value[idx] = {
          ...sessions.value[idx],
          messageCount: resp.data.messageCount,
          lastActiveAt: new Date().toISOString(),
        };
      }
    } else {
      errorMsg.value = resp.message || 'Failed to send message';
      messages.value.pop(); // remove optimistic
    }
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : 'Failed to send message';
    messages.value.pop(); // remove optimistic
  } finally {
    sending.value = false;
    await scrollToBottom();
  }
}

function handleEnter(event: KeyboardEvent) {
  if (event.shiftKey) return; // allow newline
  sendCurrentMessage();
}

async function resetCurrentContext() {
  if (!activeSessionKey.value) return;
  errorMsg.value = '';
  try {
    const resp = await conversationService.resetContext({ sessionKey: activeSessionKey.value });
    if (resp.code === 0) {
      messages.value = [];
      const idx = sessions.value.findIndex(s => s.sessionKey === activeSessionKey.value);
      if (idx !== -1) {
        sessions.value[idx] = { ...sessions.value[idx], messageCount: 0 };
      }
    } else {
      errorMsg.value = resp.message || 'Failed to reset context';
    }
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : 'Failed to reset context';
  }
}

async function closeCurrentSession() {
  if (!activeSessionKey.value) return;
  try {
    await conversationService.closeSession(activeSessionKey.value);
    activeSessionKey.value = null;
    messages.value = [];
    await loadSessionList();
  } catch {
    // silently fail
  }
}

onMounted(() => {
  loadSessionList();
});
</script>
