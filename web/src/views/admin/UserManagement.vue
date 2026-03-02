<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">User Management</h1>
      <p class="mt-1 text-sm text-gray-500">Manage all users, roles, and permissions within the system.</p>
    </header>

    <div class="mb-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0 sm:space-x-4">
      <div class="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 w-full sm:w-auto">
        <input 
          type="text" 
          v-model="searchQuery"
          placeholder="Search by username..." 
          class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full sm:w-auto"
        />
        <select 
          v-model="statusFilter"
          class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full sm:w-auto"
        >
          <option :value="0">All Statuses</option>
          <option :value="1">Active</option>
          <option :value="2">Inactive</option>
        </select>
      </div>
      <!-- <button 
        type="button" 
        @click="openAddUserModal"
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 w-full sm:w-auto"
      >
        Add New User
      </button> -->
    </div>

    <div class="bg-white shadow overflow-x-auto sm:rounded-lg">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Username</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Email</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Role</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Registered</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">VIP Expiry</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider min-w-[120px]">Actions</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-if="loading">
            <td colspan="7" class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center">Loading users...</td>
          </tr>
          <tr v-else-if="!users || users.length === 0">
            <td colspan="7" class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center">No users found.</td>
          </tr>
          <tr v-for="user in users" :key="user.id ?? Math.random()"> <!-- Fallback key if id is null, though API should provide non-null IDs -->
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ user.id ?? 'N/A' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ user.username ?? 'N/A' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ user.email ?? 'N/A' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              <div class="flex flex-wrap gap-1">
                <span
                  v-for="role in (user.roles || [])"
                  :key="role.id"
                  class="px-2 py-0.5 text-xs rounded-full bg-indigo-100 text-indigo-700"
                >{{ role.displayName }}</span>
                <span v-if="!user.roles || user.roles.length === 0" class="text-gray-400">—</span>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="[
                getUserStatusInfo(user.status).class,
                'px-2 inline-flex text-xs leading-5 font-semibold rounded-full'
              ]">
                {{ getUserStatusInfo(user.status).text }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ formatDate(user.createdAt) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ user.vipExpireAt ? formatDate(user.vipExpireAt) : 'None' }}
            </td>
<td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2">
              <!-- Edit and Delete buttons removed as per request -->
              <button 
                v-if="user.status === 2" 
                @click="setUserStatus(user.id, 1)" 
                :disabled="user.id === null"
                class="text-green-600 hover:text-green-900 disabled:opacity-50 disabled:cursor-not-allowed px-3 py-1 text-xs rounded"
              >
                Activate
              </button>
              <button 
                v-if="user.status === 1" 
                @click="setUserStatus(user.id, 2)" 
                :disabled="user.id === null"
                class="text-yellow-600 hover:text-yellow-900 disabled:opacity-50 disabled:cursor-not-allowed px-3 py-1 text-xs rounded"
              >
                Deactivate
              </button>
              <button
                v-if="!user.isAdmin"
                @click="setUserAsAdmin(user.id)"
                :disabled="user.id === null"
                class="text-indigo-600 hover:text-indigo-900 disabled:opacity-50 disabled:cursor-not-allowed px-3 py-1 text-xs rounded"
              >
                Set as Admin
              </button>
<button
                v-if="user.isAdmin"
                @click="removeAdmin(user.id)"
                :disabled="user.id === null"
                class="text-red-600 hover:text-red-900 disabled:opacity-50 disabled:cursor-not-allowed px-3 py-1 text-xs rounded"
              >
                Remove Admin
              </button>
              <button
                @click="grantVip(user.id)"
                :disabled="user.id === null"
                class="text-yellow-600 hover:text-yellow-900 disabled:opacity-50 disabled:cursor-not-allowed px-3 py-1 text-xs rounded"
              >
                Grant VIP
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    
    <div v-if="!loading && totalUsers > 0" class="mt-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0">
      <p class="text-sm text-gray-700">
        Showing <span class="font-medium">{{ (currentPage - 1) * pageSize + 1 }}</span>
        to <span class="font-medium">{{ Math.min(currentPage * pageSize, totalUsers) }}</span>
        of <span class="font-medium">{{ totalUsers }}</span> results
      </p>
      <nav v-if="totalPages > 1" class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
        <button
          @click="goToPage(currentPage - 1)"
          :disabled="currentPage === 1"
          class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Previous
        </button>
        <button
          v-for="pageNumber in paginationRange"
          :key="pageNumber"
          @click="goToPage(pageNumber)"
          :class="[
            'relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium',
            currentPage === pageNumber ? 'z-10 bg-indigo-50 border-indigo-500 text-indigo-600' : 'bg-white text-gray-700 hover:bg-gray-50'
          ]"
        >
          {{ pageNumber }}
        </button>
        <button
          @click="goToPage(currentPage + 1)"
          :disabled="currentPage === totalPages || totalPages === 0"
          class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Next
        </button>
      </nav>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import userService from '../../services/userService';
import type { User, UserQueryCriteria } from '../../models/user.model';

const users = ref<User[]>([]);
const loading = ref(true);
const totalUsers = ref(0);
const currentPage = ref(1);
const pageSize = 10; 
const searchQuery = ref('');
const statusFilter = ref<number | undefined>(0);

const fetchUsers = async () => {
  loading.value = true;
  try {
    const query: UserQueryCriteria = {
      pageIndex: currentPage.value,
      pageSize: pageSize,
    };
    if (searchQuery.value.trim() !== '') {
      query.username = searchQuery.value.trim();
    }
    if (statusFilter.value !== undefined) {
      query.status = statusFilter.value;
    }
    const response = await userService.getUsers(query);

    console.log('Fetched users:', response);
    users.value = response.data.list;
    totalUsers.value = response.data.total;
  } catch (error) {
    console.error('Failed to fetch users:', error);
    users.value = [];
    totalUsers.value = 0;
    // TODO: Show error message to user
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchUsers();
});

// const openAddUserModal = () => {
//   console.log('Open add user modal/navigate to create user page (Placeholder)');
//   // Example: router.push('/admin/users/create');
// };

// Edit user function removed as per request
// const editUser = (userId: number | null) => {
//   if (userId === null) {
//     console.warn('Edit action on user with null ID');
//     return;
//   }
//   console.log('Edit user (placeholder):', userId);
//   // Example: router.push(`/admin/users/edit/${userId}`);
// };

// Delete user function removed as per request
// const deleteUser = async (userId: number | null) => {
//   if (userId === null) {
//     console.warn('Delete action on user with null ID');
//     return;
//   }
//   if (confirm(`Are you sure you want to delete user with ID ${userId}? This action cannot be undone.`)) {
//     try {
//       loading.value = true;
//       await userService.deleteUser(userId);
//       if (users.value.length === 1 && currentPage.value > 1 && totalUsers.value > 1) { // check totalUsers to prevent going to page 0
//         currentPage.value--;
//       }
//       await fetchUsers(); // Refresh list
//     } catch (error) {
//       console.error(`Failed to delete user ${userId}:`, error);
//       // TODO: Show error message to user
//     } finally {
//       loading.value = false;
//     }
//   }
// };

const setUserStatus = async (userId: number | null, newStatus: number) => {
  if (userId === null) {
    console.warn('Set status action on user with null ID');
    // TODO: Show error message to user
    return;
  }
  const actionText = newStatus === 1 ? 'activate' : 'deactivate';
  if (confirm(`Are you sure you want to ${actionText} user with ID ${userId}?`)) {
    try {
      loading.value = true; // Indicate an operation is in progress
      await userService.updateUser(userId, { status: newStatus });
      // Optionally, find the user in the list and update their status locally
      // for a more responsive UI before refetching, or just refetch.
      // const userIndex = users.value.findIndex(u => u.id === userId);
      // if (userIndex !== -1) {
      //   users.value[userIndex].status = newStatus;
      // }
      await fetchUsers(); // Refresh the entire list to ensure data consistency
    } catch (error) {
      console.error(`Failed to ${actionText} user ${userId}:`, error);
      // TODO: Show error message to user (e.g., using a toast notification)
    } finally {
      loading.value = false;
    }
  }
};

const formatDate = (dateValue?: string | number) => {
  if (!dateValue) return 'N/A';
  try {
    const dateObj = typeof dateValue === 'number' ? new Date(dateValue) : new Date(dateValue);
    return dateObj.toLocaleDateString();
  } catch (e) {
    console.error('Invalid date value for formatting:', dateValue, e);
    return 'Invalid Date';
  }
};

const getUserStatusInfo = (statusValue?: number): { text: string; class: string } => {
  switch (statusValue) {
    case 1:
      return { text: 'Active', class: 'bg-green-100 text-green-800' };
    case 2:
      return { text: 'Inactive', class: 'bg-yellow-100 text-yellow-800' };
    default:
      return { text: 'Unknown', class: 'bg-gray-100 text-gray-800' };
  }
};

const totalPages = computed(() => {
  if (totalUsers.value === 0) return 0; // Return 0 if no users, to prevent NaN or Infinity
  return Math.ceil(totalUsers.value / pageSize);
});

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
    currentPage.value = page;
    fetchUsers();
  }
};

const paginationRange = computed(() => {
  const range = [];
  const maxPagesToShow = 5;
  let start = Math.max(1, currentPage.value - Math.floor(maxPagesToShow / 2));
  let end = Math.min(totalPages.value, start + maxPagesToShow - 1);

  if (totalPages.value > 0 && end - start + 1 < maxPagesToShow) { // Ensure totalPages > 0
    if (currentPage.value <= Math.floor(maxPagesToShow / 2) ) { // If near the beginning
        end = Math.min(totalPages.value, maxPagesToShow);
    } else { // If near the end
        start = Math.max(1, totalPages.value - maxPagesToShow + 1);
    }
  }
   // Recalculate end if start was adjusted
  if (totalPages.value > 0 && end - start + 1 < maxPagesToShow && start === 1 && totalPages.value < maxPagesToShow) {
    end = totalPages.value;
  }


  for (let i = start; i <= end; i++) {
    if (i > 0) range.push(i); // Ensure only positive page numbers
  }
  return range;
});


let searchDebounceTimer: number | undefined;
watch(searchQuery, () => { // No need for newValue if not used
  clearTimeout(searchDebounceTimer);
  searchDebounceTimer = window.setTimeout(() => { // Use window.setTimeout for NodeJS compatibility if ever needed, or just setTimeout
    currentPage.value = 1; 
    fetchUsers();
  }, 500); 
});

watch(statusFilter, () => {
  currentPage.value = 1;
  fetchUsers();
});

const setUserAsAdmin = async (userId: number | null) => {
  if (userId === null) {
    console.warn('Set admin action on user with null ID');
    return;
  }
  if (confirm(`Are you sure you want to set user with ID ${userId} as admin?`)) {
    try {
      loading.value = true;
      await userService.setAdmin(userId);
      await fetchUsers();
    } catch (error) {
      console.error(`Failed to set user ${userId} as admin:`, error);
    } finally {
      loading.value = false;
    }
  }
};

const removeAdmin = async (userId: number | null) => {
  if (userId === null) {
    console.warn('Remove admin action on user with null ID');
    return;
  }
  if (confirm(`Are you sure you want to remove admin privileges from user ID ${userId}?`)) {
    try {
      loading.value = true;
      await userService.removeAdmin(userId);
      await fetchUsers();
    } catch (error) {
      console.error(`Failed to remove admin from user ${userId}:`, error);
    } finally {
      loading.value = false;
    }
  }
};

const grantVip = async (userId: number | null) => {
  if (userId === null) {
    console.warn('Grant VIP action on user with null ID');
    return;
  }
  if (confirm(`Are you sure you want to grant 1 month VIP to user ID ${userId}?`)) {
    try {
      loading.value = true;
      await userService.grantVip(userId);
      alert('VIP granted for 1 month!');
      await fetchUsers();
    } catch (error) {
      console.error(`Failed to grant VIP to user ${userId}:`, error);
      alert('Failed to grant VIP.');
    } finally {
      loading.value = false;
    }
  }
};

</script>

<style scoped>
/* Tailwind handles styling, additional custom styles can go here if needed */
</style>
