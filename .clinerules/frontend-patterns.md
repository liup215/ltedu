# LTEDU-Web Frontend Development Patterns

## Overview

This document defines the frontend development patterns, standards, and conventions used in the LTEDU web application. All frontend development should follow these established patterns for consistency, maintainability, and scalability.

## Project Structure

### Directory Organization

```
src/
├── components/          # Reusable Vue components
│   ├── admin/          # Admin-specific components
│   └── shared/         # Shared components across roles
├── models/             # TypeScript interfaces and types
├── services/           # API service classes
├── stores/             # Pinia state management stores
├── views/              # Vue page components
│   ├── admin/          # Admin management pages
│   ├── teacher/        # Teacher dashboard pages
│   └── student/        # Student learning pages
├── router/             # Vue Router configuration
└── assets/             # Static assets
```

## Model Patterns

### Base Entity Interface

All data models should follow this pattern:

```typescript
// Base entity with ID
export interface EntityName {
  id: number;
  field1: string;
  field2?: string;          // Optional fields use ?
  relationId?: number;      // Foreign key references
  relation?: RelatedEntity; // Optional populated relation
}

// Query interface for filtering/pagination
export interface EntityNameQuery {
  pageIndex?: number;
  pageSize?: number;
  id?: number;
  searchField?: string;
  filterField?: string;
}

// Paginated response interface
export interface PaginatedEntityNames {
  list: EntityName[];
  total: number;
}

// Create request interface (omits id and computed fields)
export interface EntityNameCreateRequest {
  field1: string;
  field2?: string;
  relationId: number;
}

// Update request interface (includes id)
export interface EntityNameUpdateRequest {
  id: number;
  field1: string;
  field2?: string;
  relationId: number;
}

// List response interface
export type EntityNameListResponse = AipResponse<PaginatedEntityNames>;
export type EntityNameResponse = ApiResponse<EntityName>;
```

### Model Naming Conventions

- **Entity Models**: Use singular PascalCase (e.g., `Organisation`, `Qualification`)
- **Query Models**: Entity name + `Query` (e.g., `OrganisationQuery`)
- **Paginated Models**: `Paginated` + plural entity name (e.g., `PaginatedOrganisations`)
- **Request Models**: Entity name + `CreateRequest` or `UpdateRequest`

### Example Implementation

```typescript
// organisation.model.ts
export interface Organisation {
  id: number;
  name: string;
}

export interface OrganisationQuery {
  pageIndex?: number;
  pageSize?: number;
  id?: number;
  name?: string;
}

export interface PaginatedOrganisations {
  list: Organisation[];
  total: number;
}

export type OrganisationListResponse = ApiResponse<PaginatedOrganisations>;
export type OrganisationResponse = ApiResponse<Organisation>;
```

## Service Patterns

### Service Class Structure

All API services should follow this standardized pattern:

```typescript
import apiClient from './apiClient';
import type { 
  EntityName, 
  EntityNameQuery, 
  PaginatedEntityNames,
  EntityNameCreateRequest,
  EntityNameUpdateRequest 
} from '../models/entityName.model';
import type { ApiResponse } from '../models/api.model';

class EntityNameService {
  private baseUrl = '/api/v1/entityname';

  // GET LIST - Paginated list with filtering
  async getEntityNames(query: EntityNameQuery = {}): Promise<ApiResponse<PaginatedEntityNames>> {
    const response = await apiClient.post(`${this.baseUrl}/list`, query);
    return response.data;
  }

  // GET BY ID - Single entity by ID
  async getEntityNameById(id: number): Promise<ApiResponse<EntityName>> {
    const response = await apiClient.post(`${this.baseUrl}/byId`, { id });
    return response.data;
  }

  // GET ALL - All entities (for dropdowns/selects)
  async getAllEntityNames(query: EntityNameQuery = {}): Promise<ApiResponse<PaginatedEntityNames>> {
    const response = await apiClient.post(`${this.baseUrl}/all`, query);
    return response.data;
  }

  // CREATE - Create new entity
  async createEntityName(entity: EntityNameCreateRequest): Promise<ApiResponse<EntityName>> {
    const response = await apiClient.post(`${this.baseUrl}/create`, entity);
    return response.data;
  }

  // UPDATE - Update existing entity
  async updateEntityName(entity: EntityNameUpdateRequest): Promise<ApiResponse<void>> {
    const response = await apiClient.post(`${this.baseUrl}/edit`, entity);
    return response.data;
  }

  // DELETE - Delete entity
  async deleteEntityName(id: number): Promise<ApiResponse<void>> {
    const response = await apiClient.post(`${this.baseUrl}/delete`, { id });
    return response.data;
  }
}

export default new EntityNameService();
```

### Service Method Patterns

1. **Consistent Naming**: 
   - `getEntityNames()` - List with pagination
   - `getEntityNameById()` - Single item
   - `getAllEntityNames()` - All items (for dropdowns)
   - `createEntityName()` - Create new
   - `updateEntityName()` - Update existing
   - `deleteEntityName()` - Delete item

2. **Error Handling**: Let apiClient interceptors handle errors
3. **Return Types**: Always return `ApiResponse<T>` wrapper
4. **Request Body**: All data sent via POST body, not URL params

## Vue Component Patterns

### Management Page Structure

Management pages should follow this template:

```vue
<template>
  <div class="p-6">
    <!-- Header Section -->
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">Entity Management</h1>
      <p class="mt-1 text-sm text-gray-500">Manage system entities</p>
    </header>

    <!-- Filters and Actions -->
    <div class="mb-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0 sm:space-x-4">
      <div class="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 w-full sm:w-auto">
        <input 
          type="text" 
          v-model="searchQuery"
          placeholder="Search by name..." 
          class="px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full sm:w-auto"
        />
      </div>
      <router-link 
        to="/admin/entities/create"
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 w-full sm:w-auto"
      >
        Add Entity
      </router-link>
    </div>

    <!-- Data Table -->
    <div class="bg-white shadow overflow-x-auto sm:rounded-lg">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Field Name
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider min-w-[120px]">
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-if="loading">
            <td colspan="2" class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center">
              Loading...
            </td>
          </tr>
          <tr v-else-if="!entities || entities.length === 0">
            <td colspan="2" class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center">
              No data found
            </td>
          </tr>
          <tr v-for="entity in entities" :key="entity.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ entity.name }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
              <router-link 
                :to="`/admin/entities/${entity.id}/edit`" 
                class="text-indigo-600 hover:text-indigo-900"
              >
                Edit
              </router-link>
              <button 
                @click="deleteEntity(entity.id)"
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
    <div v-if="!loading && totalEntities > 0" class="mt-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0">
      <!-- Pagination info and controls here -->
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import entityService from '../../services/entityService';
import type { EntityName } from '../../models/entityName.model';

// Reactive data
const entities = ref<EntityName[]>([]);
const loading = ref(true);
const totalEntities = ref(0);
const currentPage = ref(1);
const pageSize = 10;
const searchQuery = ref('');

// Methods
const fetchEntities = async () => {
  loading.value = true;
  try {
    const response = await entityService.getEntityNames({
      pageIndex: currentPage.value,
      pageSize,
      name: searchQuery.value.trim() || undefined,
    });
    entities.value = response.data.list;
    totalEntities.value = response.data.total;
  } catch (error) {
    console.error('Failed to fetch entities:', error);
    // TODO: Show error message to user
  } finally {
    loading.value = false;
  }
};

const deleteEntity = async (id: number) => {
  if (confirm('Are you sure you want to delete this entity? This action cannot be undone.')) {
    try {
      await entityService.deleteEntityName(id);
      if (entities.value.length === 1 && currentPage.value > 1) {
        currentPage.value--;
      }
      fetchEntities();
    } catch (error) {
      console.error('Failed to delete entity:', error);
      // TODO: Show error message to user
    }
  }
};

// Search debounce
let searchDebounceTimer: number | undefined;
watch(searchQuery, () => {
  clearTimeout(searchDebounceTimer);
  searchDebounceTimer = window.setTimeout(() => {
    currentPage.value = 1;
    fetchEntities();
  }, 500);
});

onMounted(() => {
  fetchEntities();
});
</script>
```

### Form Component Structure

Form components should follow this pattern:

```vue
<template>
  <div class="p-6">
    <!-- Header -->
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">
        {{ isEdit ? 'Edit Entity' : 'Create Entity' }}
      </h1>
      <p class="mt-1 text-sm text-gray-500">
        {{ isEdit ? 'Update entity information' : 'Add a new entity' }}
      </p>
    </header>

    <!-- Form -->
    <div class="bg-white shadow sm:rounded-lg p-6">
      <form @submit.prevent="handleSubmit">
        <div class="space-y-6">
          <!-- Form fields here -->
          
          <!-- Action buttons -->
          <div class="flex justify-end space-x-4">
            <button 
              type="button"
              @click="goBack"
              class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Cancel
            </button>
            <button 
              type="submit"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              :disabled="loading"
            >
              {{ loading ? 'Saving...' : 'Save' }}
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import entityService from '../../services/entityService';

const route = useRoute();
const router = useRouter();
const loading = ref(false);
const isEdit = computed(() => !!route.params.id);

// Form data
const formData = ref({
  name: '',
  // other fields
});

// Load data for edit mode
onMounted(async () => {
  if (isEdit.value) {
    loading.value = true;
    try {
      const response = await entityService.getEntityNameById(Number(route.params.id));
      formData.value.name = response.data.name;
    } catch (error) {
      console.error('Failed to load entity:', error);
    } finally {
      loading.value = false;
    }
  }
});

// Form submission
const handleSubmit = async () => {
  if (!formData.value.name.trim()) {
    // TODO: Show validation error
    return;
  }

  loading.value = true;
  try {
    if (isEdit.value) {
      await entityService.updateEntityName({
        id: Number(route.params.id),
        name: formData.value.name.trim(),
      });
    } else {
      await entityService.createEntityName({
        name: formData.value.name.trim(),
      });
    }
    router.push('/admin/entities');
  } catch (error) {
    console.error('Failed to save entity:', error);
  } finally {
    loading.value = false;
  }
};

const goBack = () => {
  router.push('/admin/entities');
};
</script>
```

## State Management Patterns

### Using Pinia Stores

```typescript
// stores/entityStore.ts
import { defineStore } from 'pinia';
import { ref } from 'vue';
import type { EntityName } from '../models/entityName.model';
import entityService from '../services/entityService';

export const useEntityStore = defineStore('entity', () => {
  // State
  const entities = ref<EntityName[]>([]);
  const loading = ref(false);
  const error = ref<string | null>(null);

  // Actions
  const fetchEntities = async (query = {}) => {
    loading.value = true;
    error.value = null;
    try {
      const response = await entityService.getEntityNames(query);
      entities.value = response.data.list;
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch entities';
    } finally {
      loading.value = false;
    }
  };

  const createEntity = async (entity: EntityNameCreateRequest) => {
    try {
      const response = await entityService.createEntityName(entity);
      entities.value.push(response.data);
      return response.data;
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to create entity';
      throw err;
    }
  };

  // Getters
  const getEntityById = (id: number) => {
    return entities.value.find(entity => entity.id === id);
  };

  return {
    // State
    entities,
    loading,
    error,
    // Actions
    fetchEntities,
    createEntity,
    // Getters
    getEntityById,
  };
});
```

## Router Patterns

### Route Configuration

```typescript
// router/entityRoutes.ts
import type { RouteRecordRaw } from 'vue-router';

export const entityRoutes: RouteRecordRaw[] = [
  {
    path: '/admin/entities',
    name: 'EntityManagement',
    component: () => import('../views/admin/EntityManagement.vue'),
    meta: {
      requiresAuth: true,
      roles: ['admin']
    }
  },
  {
    path: '/admin/entities/create',
    name: 'EntityCreate',
    component: () => import('../views/admin/EntityForm.vue'),
    meta: {
      requiresAuth: true,
      roles: ['admin']
    }
  },
  {
    path: '/admin/entities/:id/edit',
    name: 'EntityEdit',
    component: () => import('../views/admin/EntityForm.vue'),
    meta: {
      requiresAuth: true,
      roles: ['admin']
    }
  }
];
```

## Styling Patterns

### Tailwind CSS Classes

#### Common Button Styles
```html
<!-- Primary button -->
<button class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
  Primary Action
</button>

<!-- Secondary button -->
<button class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
  Secondary Action
</button>

<!-- Danger button -->
<button class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500">
  Delete
</button>
```

#### Form Input Styles
```html
<!-- Text input -->
<input 
  type="text"
  class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
  placeholder="Enter value"
/>

<!-- Select dropdown -->
<select class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
  <option value="">Select option</option>
</select>

<!-- Textarea -->
<textarea 
  rows="3"
  class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
  placeholder="Enter description"
></textarea>
```

#### Table Styles
```html
<div class="bg-white shadow overflow-x-auto sm:rounded-lg">
  <table class="min-w-full divide-y divide-gray-200">
    <thead class="bg-gray-50">
      <tr>
        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
          Header
        </th>
      </tr>
    </thead>
    <tbody class="bg-white divide-y divide-gray-200">
      <tr>
        <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
          Data
        </td>
      </tr>
    </tbody>
  </table>
</div>
```

## Error Handling Patterns

### Service Error Handling

```typescript
// In service methods
try {
  const response = await apiClient.post(endpoint, data);
  return response.data;
} catch (error) {
  console.error('Service operation failed:', error);
  // Error is already handled by apiClient interceptors
  throw error; // Re-throw to let component handle UI feedback
}
```

### Component Error Handling

```typescript
// In Vue components
const handleOperation = async () => {
  loading.value = true;
  try {
    await someService.performOperation();
    // Show success message
  } catch (error) {
    console.error('Operation failed:', error);
    // TODO: Show error message to user
    // errorMessage.value = error.message || 'Operation failed';
  } finally {
    loading.value = false;
  }
};
```

## Form Validation Patterns

### Basic Validation

```typescript
// Form validation in components
const validateForm = () => {
  const errors: string[] = [];
  
  if (!formData.value.name.trim()) {
    errors.push('Name is required');
  }
  
  if (formData.value.name.length > 100) {
    errors.push('Name must be less than 100 characters');
  }
  
  return errors;
};

const handleSubmit = async () => {
  const validationErrors = validateForm();
  if (validationErrors.length > 0) {
    // Show validation errors
    return;
  }
  
  // Proceed with form submission
};
```

## Pagination Patterns

### Standard Pagination Implementation

```typescript
// Pagination computed properties
const totalPages = computed(() => {
  return Math.ceil(totalItems.value / pageSize);
});

const paginationRange = computed(() => {
  const range = [];
  const maxPagesToShow = 5;
  let start = Math.max(1, currentPage.value - Math.floor(maxPagesToShow / 2));
  let end = Math.min(totalPages.value, start + maxPagesToShow - 1);

  if (totalPages.value > 0 && end - start + 1 < maxPagesToShow) {
    if (currentPage.value <= Math.floor(maxPagesToShow / 2)) {
      end = Math.min(totalPages.value, maxPagesToShow);
    } else {
      start = Math.max(1, totalPages.value - maxPagesToShow + 1);
    }
  }

  for (let i = start; i <= end; i++) {
    if (i > 0) range.push(i);
  }
  return range;
});

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
    currentPage.value = page;
    fetchData();
  }
};
```

## Search and Filtering Patterns

### Debounced Search

```typescript
// Search with debounce
const searchQuery = ref('');
let searchDebounceTimer: number | undefined;

watch(searchQuery, () => {
  clearTimeout(searchDebounceTimer);
  searchDebounceTimer = window.setTimeout(() => {
    currentPage.value = 1; // Reset to first page
    fetchData();
  }, 500);
});
```

## Best Practices

### Code Organization
1. **File Naming**: Use kebab-case for files, PascalCase for components
2. **Import Order**: External packages → Internal services → Components → Types
3. **Component Structure**: Template → Script → Style (if needed)

### Performance
1. **Lazy Loading**: Use dynamic imports for routes
2. **Debouncing**: Implement search debouncing (500ms)
3. **Pagination**: Always paginate large datasets
4. **Loading States**: Show loading indicators for async operations

### Accessibility
1. **Semantic HTML**: Use proper HTML elements
2. **ARIA Labels**: Add appropriate ARIA attributes
3. **Keyboard Navigation**: Ensure keyboard accessibility
4. **Focus Management**: Manage focus states properly

### Security
1. **Input Sanitization**: Sanitize user inputs
2. **XSS Prevention**: Use Vue's built-in escaping
3. **Route Protection**: Implement proper route guards
4. **Token Management**: Handle JWT tokens securely

### Testing
1. **Unit Tests**: Test service methods and utilities
2. **Component Tests**: Test component behavior
3. **E2E Tests**: Test complete user workflows
4. **API Mocking**: Mock API responses for testing

## Migration and Updates

### Component Updates
1. **Incremental Changes**: Make small, focused changes
2. **Backward Compatibility**: Maintain API compatibility
3. **Documentation**: Update documentation with changes
4. **Testing**: Thoroughly test all changes

### Breaking Changes
1. **Version Planning**: Plan breaking changes carefully
2. **Migration Guides**: Provide clear migration instructions
3. **Deprecation Warnings**: Add warnings before removing features
4. **Communication**: Communicate changes to team

## Examples

### Complete CRUD Implementation

Refer to the Organisation management implementation as the canonical example:
- `src/models/organisation.model.ts` - Model definitions
- `src/services/organisationService.ts` - Service implementation
- `src/views/admin/OrganisationManagement.vue` - List management page
- `src/views/admin/OrganisationForm.vue` - Create/edit form

This pattern should be replicated for all entity management features in the LTEDU system.
