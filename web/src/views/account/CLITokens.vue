<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">CLI Tokens</h1>
      <p class="mt-1 text-sm text-gray-500">
        Generate tokens to authenticate <code class="font-mono bg-gray-100 px-1 rounded">edu-cli</code> with the backend API.
      </p>
    </header>

    <!-- Quick Setup Guide -->
    <div class="bg-gray-50 border border-gray-200 rounded-lg p-5 mb-6">
      <h2 class="text-base font-semibold text-gray-800 mb-3 flex items-center gap-2">
        <!-- terminal icon -->
        <svg class="h-5 w-5 text-gray-500" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
        </svg>
        Quick Setup
      </h2>

      <ol class="space-y-3 text-sm text-gray-700">
        <li>
          <span class="font-medium">1. Download <code class="font-mono bg-gray-200 px-1 rounded">edu-cli</code></span> from the
          <a href="https://github.com/liup215/ltedu/releases" target="_blank" rel="noopener noreferrer" class="text-indigo-600 hover:text-indigo-800 underline">GitHub Releases page</a>.
        </li>
        <li>
          <span class="font-medium">2. Configure the backend URL:</span>
          <div class="mt-1 relative">
            <code class="block bg-gray-800 text-green-300 text-xs font-mono px-4 py-2 rounded">
              edu-cli config set-url {{ baseUrl }}
            </code>
            <button
              @click="copyText(`edu-cli config set-url ${baseUrl}`)"
              class="absolute top-1 right-2 text-gray-400 hover:text-white text-xs px-1"
              aria-label="Copy set-url command"
              title="Copy"
            >⎘</button>
          </div>
        </li>
        <li>
          <span class="font-medium">3. Create a token below, then configure it:</span>
          <div class="mt-1 relative">
            <code class="block bg-gray-800 text-green-300 text-xs font-mono px-4 py-2 rounded">
              edu-cli config set-token &lt;your-token&gt;
            </code>
          </div>
        </li>
        <li>
          <span class="font-medium">4. Verify your setup:</span>
          <div class="mt-1 relative">
            <code class="block bg-gray-800 text-green-300 text-xs font-mono px-4 py-2 rounded">
              edu-cli syllabus list
            </code>
          </div>
        </li>
      </ol>

      <p class="mt-4 text-xs text-gray-500">
        Alternatively, set environment variables:
        <code class="font-mono bg-gray-200 px-1 rounded">EDU_BASE_URL={{ baseUrl }}</code>
        and <code class="font-mono bg-gray-200 px-1 rounded">EDU_TOKEN=&lt;your-token&gt;</code>
      </p>
    </div>

    <!-- Create Token Form -->
    <div class="bg-white shadow sm:rounded-lg p-6 mb-6">
      <h2 class="text-lg font-medium text-gray-900 mb-4">Create New CLI Token</h2>
      <form @submit.prevent="createToken" class="space-y-4">
        <div>
          <label for="tokenName" class="block text-sm font-medium text-gray-700">Token Name</label>
          <input
            v-model="newToken.name"
            type="text"
            id="tokenName"
            required
            placeholder="e.g., My Laptop CLI"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
        </div>
        <div>
          <label for="expiresAt" class="block text-sm font-medium text-gray-700">
            Expiration Date (Optional)
          </label>
          <input
            v-model="newToken.expiresAt"
            type="datetime-local"
            id="expiresAt"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
          <p class="mt-1 text-xs text-gray-500">Default: 1 year from now</p>
        </div>
        <button
          type="submit"
          :disabled="creating || !newToken.name"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ creating ? 'Creating...' : 'Create Token' }}
        </button>
      </form>
    </div>

    <!-- Token List -->
    <div class="bg-white shadow overflow-hidden sm:rounded-lg">
      <div class="px-4 py-5 sm:px-6 flex justify-between items-center">
        <h2 class="text-lg font-medium text-gray-900">Your CLI Tokens</h2>
        <button
          @click="loadTokens"
          class="inline-flex items-center px-3 py-1 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none"
        >
          Refresh
        </button>
      </div>

      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Token</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Created</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Expires</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-if="loading">
              <td colspan="6" class="px-6 py-4 text-center text-sm text-gray-500">Loading tokens...</td>
            </tr>
            <tr v-else-if="!tokens || tokens.length === 0">
              <td colspan="6" class="px-6 py-4 text-center text-sm text-gray-500">
                No tokens yet. Create your first token above.
              </td>
            </tr>
            <tr v-for="token in tokens" :key="token.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ token.name }}</td>
              <td class="px-6 py-4 text-sm text-gray-500">
                <div class="flex items-center space-x-2">
                  <code class="bg-gray-100 px-2 py-1 rounded text-xs font-mono max-w-xs truncate">
                    {{ showFull[token.id] ? token.token : maskToken(token.token) }}
                  </code>
                  <button @click="showFull[token.id] = !showFull[token.id]" class="text-indigo-600 hover:text-indigo-900 text-xs whitespace-nowrap">
                    {{ showFull[token.id] ? 'Hide' : 'Show' }}
                  </button>
                  <button @click="copyToken(token.token)" class="text-indigo-600 hover:text-indigo-900 text-xs whitespace-nowrap">Copy</button>
                  <button @click="copySetTokenCmd(token.token)" class="text-green-700 hover:text-green-900 text-xs whitespace-nowrap font-medium">
                    Copy CLI cmd
                  </button>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="token.isActive ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'"
                  class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full">
                  {{ token.isActive ? 'Active' : 'Inactive' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ formatDate(token.createdAt) }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ formatDate(token.expiresAt) }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
                <button v-if="token.isActive" @click="deactivateToken(token.id)" class="text-yellow-600 hover:text-yellow-900">Deactivate</button>
                <button v-else @click="activateToken(token.id)" class="text-green-600 hover:text-green-900">Activate</button>
                <button @click="pendingDelete = token.id" class="text-red-600 hover:text-red-900">Delete</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="pendingDelete !== null" class="fixed z-10 inset-0 overflow-y-auto" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75" @click="pendingDelete = null"></div>
        <span class="hidden sm:inline-block sm:align-middle sm:h-screen">&#8203;</span>
        <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
                <svg class="h-6 w-6 text-red-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                <h3 class="text-lg leading-6 font-medium text-gray-900">Delete Token</h3>
                <p class="mt-2 text-sm text-gray-500">Are you sure you want to delete this token? This action cannot be undone.</p>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button @click="deleteToken" type="button" class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 sm:ml-3 sm:w-auto sm:text-sm">
              Delete
            </button>
            <button @click="pendingDelete = null" type="button" class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm">
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import mcpTokenService from '../../services/mcpTokenService';
import type { MCPToken, MCPTokenCreateRequest } from '../../models/mcpToken.model';
import { showSuccess, showError } from '../../utils/notification';
import { API_BASE_URL } from '../../const/config';

const loading = ref(false);
const creating = ref(false);
const tokens = ref<MCPToken[]>([]);
const total = ref(0);
const baseUrl = ref('');
const showFull = ref<Record<number, boolean>>({});
const pendingDelete = ref<number | null>(null);

const newToken = ref<MCPTokenCreateRequest>({ name: '', expiresAt: undefined });

onMounted(async () => {
  baseUrl.value = await API_BASE_URL();
  loadTokens();
});

async function loadTokens() {
  loading.value = true;
  try {
    const response = await mcpTokenService.listTokens({ pageIndex: 1, pageSize: 100 });
    if (response.code === 0 && response.data) {
      tokens.value = response.data.list || [];
      total.value = response.data.total || 0;
    } else {
      showError(response.message || 'Failed to load tokens');
    }
  } catch (error: any) {
    showError(error.message || 'Failed to load tokens');
  } finally {
    loading.value = false;
  }
}

async function createToken() {
  if (!newToken.value.name) return;
  creating.value = true;
  try {
    const data: MCPTokenCreateRequest = { name: newToken.value.name };
    if (newToken.value.expiresAt) {
      data.expiresAt = new Date(newToken.value.expiresAt).toISOString();
    }
    const response = await mcpTokenService.createToken(data);
    if (response.code === 0) {
      showSuccess('Token created successfully');
      newToken.value = { name: '', expiresAt: undefined };
      await loadTokens();
    } else {
      showError(response.message || 'Failed to create token');
    }
  } catch (error: any) {
    showError(error.message || 'Failed to create token');
  } finally {
    creating.value = false;
  }
}

async function deleteToken() {
  if (pendingDelete.value === null) return;
  try {
    const response = await mcpTokenService.deleteToken(pendingDelete.value);
    if (response.code === 0) {
      showSuccess('Token deleted');
      pendingDelete.value = null;
      await loadTokens();
    } else {
      showError(response.message || 'Failed to delete token');
    }
  } catch (error: any) {
    showError(error.message || 'Failed to delete token');
  }
}

async function deactivateToken(id: number) {
  try {
    const response = await mcpTokenService.deactivateToken(id);
    if (response.code === 0) {
      showSuccess('Token deactivated');
      await loadTokens();
    } else {
      showError(response.message || 'Failed to deactivate token');
    }
  } catch (error: any) {
    showError(error.message || 'Failed to deactivate token');
  }
}

async function activateToken(id: number) {
  try {
    const response = await mcpTokenService.activateToken(id);
    if (response.code === 0) {
      showSuccess('Token activated');
      await loadTokens();
    } else {
      showError(response.message || 'Failed to activate token');
    }
  } catch (error: any) {
    showError(error.message || 'Failed to activate token');
  }
}

function maskToken(token: string): string {
  if (token.length <= 8) return token;
  return token.substring(0, 4) + '...' + token.substring(token.length - 4);
}

async function copyText(text: string) {
  try {
    await navigator.clipboard.writeText(text);
    showSuccess('Copied to clipboard');
  } catch {
    showError('Failed to copy');
  }
}

async function copyToken(token: string) {
  await copyText(token);
}

async function copySetTokenCmd(token: string) {
  await copyText(`edu-cli config set-token ${token}`);
  showSuccess('CLI command copied — paste it in your terminal');
}

function formatDate(dateString: string): string {
  if (!dateString) return 'N/A';
  return new Date(dateString).toLocaleString();
}
</script>
