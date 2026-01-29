<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">MCP Token Management</h1>
      <p class="mt-1 text-sm text-gray-500">
        Manage your Model Context Protocol (MCP) tokens for API access.
      </p>
    </header>

    <!-- MCP Endpoint Info -->
    <div class="bg-blue-50 border border-blue-200 rounded-lg p-4 mb-6">
      <h3 class="text-sm font-medium text-blue-800 mb-2">MCP Endpoint</h3>
      <div class="flex items-center space-x-2">
        <code class="flex-1 bg-white px-3 py-2 rounded border border-blue-200 text-sm font-mono text-gray-800 break-all">
          {{ mcpEndpoint }}
        </code>
        <button
          @click="copyMcpEndpoint"
          class="inline-flex items-center px-3 py-2 border border-blue-300 text-sm font-medium rounded-md text-blue-700 bg-white hover:bg-blue-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          Copy
        </button>
      </div>
      <p class="mt-2 text-xs text-blue-600">
        Use this endpoint with your token in AI assistants like Claude, Cursor, etc.
      </p>
    </div>

    <!-- Create Token Form -->
    <div class="bg-white shadow sm:rounded-lg p-6 mb-6">
      <h2 class="text-lg font-medium text-gray-900 mb-4">Create New Token</h2>
      <form @submit.prevent="createToken" class="space-y-4">
        <div>
          <label for="tokenName" class="block text-sm font-medium text-gray-700">
            Token Name
          </label>
          <input
            v-model="newToken.name"
            type="text"
            id="tokenName"
            required
            placeholder="e.g., My Development Token"
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
          <p class="mt-1 text-xs text-gray-500">
            Default: 1 year from now
          </p>
        </div>
        <button
          type="submit"
          :disabled="loading || !newToken.name"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Create Token
        </button>
      </form>
    </div>

    <!-- Token List -->
    <div class="bg-white shadow overflow-hidden sm:rounded-lg">
      <div class="px-4 py-5 sm:px-6 flex justify-between items-center">
        <h2 class="text-lg font-medium text-gray-900">Your MCP Tokens</h2>
        <button
          @click="loadTokens"
          class="inline-flex items-center px-3 py-1 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          Refresh
        </button>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Name
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Token
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Status
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Created
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Expires
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Last Used
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-if="loading">
              <td colspan="7" class="px-6 py-4 text-center text-sm text-gray-500">
                Loading tokens...
              </td>
            </tr>
            <tr v-else-if="!tokens || tokens.length === 0">
              <td colspan="7" class="px-6 py-4 text-center text-sm text-gray-500">
                No tokens found. Create your first token above.
              </td>
            </tr>
            <tr v-for="token in tokens" :key="token.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ token.name }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                <div class="flex items-center space-x-2">
                  <code class="bg-gray-100 px-2 py-1 rounded text-xs font-mono">
                    {{ showFullToken[token.id] ? token.token : maskToken(token.token) }}
                  </code>
                  <button
                    @click="toggleTokenVisibility(token.id)"
                    class="text-indigo-600 hover:text-indigo-900 text-xs"
                  >
                    {{ showFullToken[token.id] ? 'Hide' : 'Show' }}
                  </button>
                  <button
                    @click="copyToken(token.token)"
                    class="text-indigo-600 hover:text-indigo-900 text-xs"
                  >
                    Copy Token
                  </button>
                  <button
                    @click="copyFullMcpUrl(token.token)"
                    class="text-green-600 hover:text-green-900 text-xs"
                  >
                    Copy URL
                  </button>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
                  :class="[
                    token.isActive ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800',
                    'px-2 inline-flex text-xs leading-5 font-semibold rounded-full'
                  ]"
                >
                  {{ token.isActive ? 'Active' : 'Inactive' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(token.createdAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(token.expiresAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(token.lastUsed) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
                <button
                  v-if="token.isActive"
                  @click="deactivateToken(token.id)"
                  class="text-yellow-600 hover:text-yellow-900"
                >
                  Deactivate
                </button>
                <button
                  v-else
                  @click="activateToken(token.id)"
                  class="text-green-600 hover:text-green-900"
                >
                  Activate
                </button>
                <button
                  @click="confirmDeleteToken(token.id)"
                  class="text-red-600 hover:text-red-900"
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div v-if="total > pageSize" class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
        <div class="flex-1 flex justify-between sm:hidden">
          <button
            @click="previousPage"
            :disabled="currentPage === 1"
            class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            Previous
          </button>
          <button
            @click="nextPage"
            :disabled="currentPage >= totalPages"
            class="ml-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            Next
          </button>
        </div>
        <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
          <div>
            <p class="text-sm text-gray-700">
              Showing
              <span class="font-medium">{{ (currentPage - 1) * pageSize + 1 }}</span>
              to
              <span class="font-medium">{{ Math.min(currentPage * pageSize, total) }}</span>
              of
              <span class="font-medium">{{ total }}</span>
              results
            </p>
          </div>
          <div>
            <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
              <button
                @click="previousPage"
                :disabled="currentPage === 1"
                class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                Previous
              </button>
              <button
                @click="nextPage"
                :disabled="currentPage >= totalPages"
                class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                Next
              </button>
            </nav>
          </div>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div
      v-if="deleteConfirmation.show"
      class="fixed z-10 inset-0 overflow-y-auto"
      aria-labelledby="modal-title"
      role="dialog"
      aria-modal="true"
    >
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true"></div>
        <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>
        <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
                <svg class="h-6 w-6 text-red-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                  Delete Token
                </h3>
                <div class="mt-2">
                  <p class="text-sm text-gray-500">
                    Are you sure you want to delete this token? This action cannot be undone.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button
              type="button"
              @click="deleteToken"
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm"
            >
              Delete
            </button>
            <button
              type="button"
              @click="deleteConfirmation.show = false"
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
            >
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import mcpTokenService from '../../services/mcpTokenService';
import type { MCPToken, MCPTokenCreateRequest } from '../../models/mcpToken.model';
import { showSuccess, showError } from '../../utils/notification';
import { API_BASE_URL } from '../../const/config';

const loading = ref(false);
const tokens = ref<MCPToken[]>([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(20);
const mcpEndpoint = ref('');

const newToken = ref<MCPTokenCreateRequest>({
  name: '',
  expiresAt: undefined
});

const showFullToken = ref<Record<number, boolean>>({});
const deleteConfirmation = ref({
  show: false,
  tokenId: null as number | null
});

const totalPages = computed(() => Math.ceil(total.value / pageSize.value));

onMounted(async () => {
  // Build MCP endpoint URL
  const baseUrl = await API_BASE_URL();
  mcpEndpoint.value = `${baseUrl}/api/mcp?token=YOUR_TOKEN`;
  loadTokens();
});

async function loadTokens() {
  loading.value = true;
  try {
    const response = await mcpTokenService.listTokens({
      pageIndex: currentPage.value,
      pageSize: pageSize.value
    });
    
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
  loading.value = true;
  try {
    // Convert datetime-local to ISO 8601 format if provided
    const data: MCPTokenCreateRequest = {
      name: newToken.value.name
    };
    
    if (newToken.value.expiresAt) {
      const date = new Date(newToken.value.expiresAt);
      data.expiresAt = date.toISOString();
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
    loading.value = false;
  }
}

function confirmDeleteToken(tokenId: number) {
  deleteConfirmation.value = {
    show: true,
    tokenId
  };
}

async function deleteToken() {
  if (!deleteConfirmation.value.tokenId) return;
  
  loading.value = true;
  try {
    const response = await mcpTokenService.deleteToken(deleteConfirmation.value.tokenId);
    
    if (response.code === 0) {
      showSuccess('Token deleted successfully');
      deleteConfirmation.value = { show: false, tokenId: null };
      await loadTokens();
    } else {
      showError(response.message || 'Failed to delete token');
    }
  } catch (error: any) {
    showError(error.message || 'Failed to delete token');
  } finally {
    loading.value = false;
  }
}

async function deactivateToken(tokenId: number) {
  loading.value = true;
  try {
    const response = await mcpTokenService.deactivateToken(tokenId);
    
    if (response.code === 0) {
      showSuccess('Token deactivated successfully');
      await loadTokens();
    } else {
      showError(response.message || 'Failed to deactivate token');
    }
  } catch (error: any) {
    showError(error.message || 'Failed to deactivate token');
  } finally {
    loading.value = false;
  }
}

async function activateToken(tokenId: number) {
  loading.value = true;
  try {
    const response = await mcpTokenService.activateToken(tokenId);
    
    if (response.code === 0) {
      showSuccess('Token activated successfully');
      await loadTokens();
    } else {
      showError(response.message || 'Failed to activate token');
    }
  } catch (error: any) {
    showError(error.message || 'Failed to activate token');
  } finally {
    loading.value = false;
  }
}

function maskToken(token: string): string {
  if (token.length <= 8) return token;
  return token.substring(0, 4) + '...' + token.substring(token.length - 4);
}

function toggleTokenVisibility(tokenId: number) {
  showFullToken.value[tokenId] = !showFullToken.value[tokenId];
}

async function copyToken(token: string) {
  try {
    await navigator.clipboard.writeText(token);
    showSuccess('Token copied to clipboard');
  } catch (error) {
    showError('Failed to copy token');
  }
}

async function copyMcpEndpoint() {
  try {
    await navigator.clipboard.writeText(mcpEndpoint.value);
    showSuccess('MCP Endpoint copied to clipboard');
  } catch (error) {
    showError('Failed to copy endpoint');
  }
}

async function copyFullMcpUrl(token: string) {
  try {
    const baseUrl = await API_BASE_URL();
    const fullUrl = `${baseUrl}/api/mcp?token=${token}`;
    await navigator.clipboard.writeText(fullUrl);
    showSuccess('Full MCP URL copied to clipboard');
  } catch (error) {
    showError('Failed to copy URL');
  }
}

function formatDate(dateString: string): string {
  if (!dateString) return 'N/A';
  const date = new Date(dateString);
  return date.toLocaleString();
}

function previousPage() {
  if (currentPage.value > 1) {
    currentPage.value--;
    loadTokens();
  }
}

function nextPage() {
  if (currentPage.value < totalPages.value) {
    currentPage.value++;
    loadTokens();
  }
}
</script>
