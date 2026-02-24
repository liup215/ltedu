# Syllabus Navigator Frontend Implementation Guide

## Overview

This document provides a complete implementation guide for the Syllabus Navigator frontend. The backend APIs are fully implemented (23 endpoints across Goal, Task, Attempt, KnowledgeState modules). TypeScript models are complete. This guide details the remaining frontend work.

---

## Architecture

```
Syllabus Navigator Frontend
├── Models (✅ DONE - 4 files)
│   ├── goal.model.ts
│   ├── task.model.ts
│   ├── attempt.model.ts
│   └── knowledgeState.model.ts
│
├── Services (TODO - 4 files)
│   ├── goalService.ts
│   ├── taskService.ts
│   ├── attemptService.ts
│   └── knowledgeStateService.ts
│
├── Views/Student (TODO - 8-10 files)
│   ├── GoalCreation.vue - Create learning goals
│   ├── GoalDashboard.vue - View all goals
│   ├── TaskStream.vue - Daily task cards
│   ├── ProgressDashboard.vue - Main progress view
│   ├── KnowledgeHeatmap.vue - Mastery visualization
│   ├── ReviewDue.vue - SRS review reminders
│   ├── PracticeSession.vue - Integrated practice
│   └── PlanGenerator.vue - Generate study plans
│
└── Components (TODO - 5-7 files)
    ├── TaskCard.vue - Individual task display
    ├── ProgressChart.vue - Charts & metrics
    ├── MasteryHeatmap.vue - Chapter mastery grid
    ├── ReviewReminder.vue - SRS notifications
    └── GoalCard.vue - Goal display card
```

---

## Implementation Priority

### Phase 1: Core Student Experience (MVP)

#### 1.1 Goal Service (`goalService.ts`)
```typescript
import { apiClient } from './apiClient';
import type {
  Goal,
  CreateGoalRequest,
  UpdateGoalRequest,
  GoalListResponse,
  GoalProgressResponse
} from '@/models/goal.model';

class GoalService {
  private readonly basePath = '/api/v1/goal';

  async create(data: CreateGoalRequest): Promise<Goal> {
    const response = await apiClient.post(`${this.basePath}/create`, data);
    return response.data;
  }

  async update(data: UpdateGoalRequest): Promise<Goal> {
    const response = await apiClient.post(`${this.basePath}/edit`, data);
    return response.data;
  }

  async getById(id: number): Promise<Goal> {
    const response = await apiClient.post(`${this.basePath}/byId`, { id });
    return response.data;
  }

  async list(page = 1, perPage = 20): Promise<GoalListResponse> {
    const response = await apiClient.post(`${this.basePath}/list`, {
      page,
      perPage
    });
    return response.data;
  }

  async getActive(): Promise<Goal[]> {
    const response = await apiClient.post(`${this.basePath}/active`);
    return response.data;
  }

  async delete(id: number): Promise<void> {
    await apiClient.post(`${this.basePath}/delete`, { id });
  }

  async getProgress(goalId: number): Promise<GoalProgressResponse> {
    const response = await apiClient.post(
      `/api/v1/knowledge-state/progress`,
      { goalId }
    );
    return response.data;
  }
}

export default new GoalService();
```

#### 1.2 Task Service (`taskService.ts`)
```typescript
import { apiClient } from './apiClient';
import type {
  Task,
  CreateTaskRequest,
  UpdateTaskRequest,
  CompleteTaskRequest,
  TaskStreamResponse,
  GeneratePlanRequest
} from '@/models/task.model';

class TaskService {
  private readonly basePath = '/api/v1/task';

  async create(data: CreateTaskRequest): Promise<Task> {
    const response = await apiClient.post(`${this.basePath}/create`, data);
    return response.data;
  }

  async update(data: UpdateTaskRequest): Promise<Task> {
    const response = await apiClient.post(`${this.basePath}/edit`, data);
    return response.data;
  }

  async getById(id: number): Promise<Task> {
    const response = await apiClient.post(`${this.basePath}/byId`, { id });
    return response.data;
  }

  async list(goalId: number, page = 1, perPage = 50): Promise<{ tasks: Task[], total: number }> {
    const response = await apiClient.post(`${this.basePath}/list`, {
      goalId,
      page,
      perPage
    });
    return response.data;
  }

  async getStream(goalId: number): Promise<TaskStreamResponse> {
    const response = await apiClient.post(`${this.basePath}/stream`, { goalId });
    return response.data;
  }

  async complete(data: CompleteTaskRequest): Promise<Task> {
    const response = await apiClient.post(`${this.basePath}/complete`, data);
    return response.data;
  }

  async generatePlan(data: GeneratePlanRequest): Promise<{ tasks: Task[], count: number }> {
    const response = await apiClient.post(`${this.basePath}/generate-plan`, data);
    return response.data;
  }

  async delete(id: number): Promise<void> {
    await apiClient.post(`${this.basePath}/delete`, { id });
  }
}

export default new TaskService();
```

#### 1.3 Goal Creation View (`GoalCreation.vue`)
```vue
<template>
  <div class="max-w-4xl mx-auto p-6">
    <h1 class="text-3xl font-bold mb-8">{{ t('syllabus_navigator.create_goal') }}</h1>
    
    <form @submit.prevent="handleSubmit" class="space-y-6">
      <!-- Syllabus Selection -->
      <div>
        <label class="block text-sm font-medium mb-2">
          {{ t('syllabus_navigator.select_syllabus') }}
        </label>
        <select 
          v-model="form.syllabusId" 
          required
          class="w-full px-4 py-2 border rounded-lg"
        >
          <option value="">{{ t('common.select') }}</option>
          <option v-for="s in syllabuses" :key="s.id" :value="s.id">
            {{ s.name }}
          </option>
        </select>
      </div>

      <!-- Goal Name -->
      <div>
        <label class="block text-sm font-medium mb-2">
          {{ t('syllabus_navigator.goal_name') }}
        </label>
        <input
          v-model="form.name"
          type="text"
          required
          class="w-full px-4 py-2 border rounded-lg"
          :placeholder="t('syllabus_navigator.goal_name_placeholder')"
        />
      </div>

      <!-- Exam Date -->
      <div>
        <label class="block text-sm font-medium mb-2">
          {{ t('syllabus_navigator.exam_date') }}
        </label>
        <input
          v-model="form.examDate"
          type="date"
          class="w-full px-4 py-2 border rounded-lg"
        />
      </div>

      <!-- Weekly Hours -->
      <div>
        <label class="block text-sm font-medium mb-2">
          {{ t('syllabus_navigator.weekly_hours') }}
        </label>
        <input
          v-model.number="form.weeklyHours"
          type="number"
          min="1"
          max="168"
          class="w-full px-4 py-2 border rounded-lg"
          :placeholder="t('syllabus_navigator.weekly_hours_placeholder')"
        />
      </div>

      <!-- Learning Mode -->
      <div>
        <label class="block text-sm font-medium mb-2">
          {{ t('syllabus_navigator.learning_mode') }}
        </label>
        <select v-model="form.mode" class="w-full px-4 py-2 border rounded-lg">
          <option value="self-paced">{{ t('syllabus_navigator.mode_self_paced') }}</option>
          <option value="school-sync">{{ t('syllabus_navigator.mode_school_sync') }}</option>
          <option value="cram">{{ t('syllabus_navigator.mode_cram') }}</option>
        </select>
      </div>

      <!-- Description -->
      <div>
        <label class="block text-sm font-medium mb-2">
          {{ t('syllabus_navigator.description') }}
        </label>
        <textarea
          v-model="form.description"
          rows="3"
          class="w-full px-4 py-2 border rounded-lg"
          :placeholder="t('syllabus_navigator.description_placeholder')"
        />
      </div>

      <!-- Submit Button -->
      <div class="flex gap-4">
        <button
          type="submit"
          :disabled="loading"
          class="flex-1 px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
        >
          {{ loading ? t('common.creating') : t('syllabus_navigator.create_goal') }}
        </button>
        <button
          type="button"
          @click="$router.back()"
          class="px-6 py-3 border border-gray-300 rounded-lg hover:bg-gray-50"
        >
          {{ t('common.cancel') }}
        </button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { push } from 'notivue';
import goalService from '@/services/goalService';
import syllabusService from '@/services/syllabusService';
import type { CreateGoalRequest } from '@/models/goal.model';

const { t } = useI18n();
const router = useRouter();

const syllabuses = ref([]);
const loading = ref(false);
const form = ref<CreateGoalRequest>({
  syllabusId: 0,
  name: '',
  description: '',
  examDate: '',
  targetScore: undefined,
  weeklyHours: 10,
  mode: 'self-paced'
});

onMounted(async () => {
  try {
    const response = await syllabusService.list(1, 100);
    syllabuses.value = response.syllabuses;
  } catch (error) {
    push.error({ message: t('errors.load_syllabuses_failed') });
  }
});

const handleSubmit = async () => {
  loading.value = true;
  try {
    const goal = await goalService.create(form.value);
    push.success({ message: t('syllabus_navigator.goal_created') });
    router.push(`/student/goals/${goal.id}`);
  } catch (error: any) {
    push.error({ message: error.response?.data?.message || t('errors.create_goal_failed') });
  } finally {
    loading.value = false;
  }
};
</script>
```

#### 1.4 Task Stream View (`TaskStream.vue`)
```vue
<template>
  <div class="max-w-7xl mx-auto p-6">
    <!-- Header -->
    <div class="flex justify-between items-center mb-8">
      <div>
        <h1 class="text-3xl font-bold">{{ t('syllabus_navigator.my_tasks') }}</h1>
        <p class="text-gray-600 mt-2">
          {{ activeGoal?.name }} - {{ formatDate(new Date()) }}
        </p>
      </div>
      <button
        @click="loadTasks"
        :disabled="loading"
        class="px-4 py-2 border rounded-lg hover:bg-gray-50"
      >
        🔄 {{ t('common.refresh') }}
      </button>
    </div>

    <!-- Today's Tasks -->
    <section class="mb-8">
      <h2 class="text-2xl font-semibold mb-4 flex items-center gap-2">
        📅 {{ t('syllabus_navigator.today_tasks') }}
        <span class="text-sm font-normal text-gray-500">
          ({{ taskStream.today.length }})
        </span>
      </h2>
      <div v-if="taskStream.today.length === 0" class="text-gray-500 text-center py-8">
        {{ t('syllabus_navigator.no_tasks_today') }}
      </div>
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <TaskCard
          v-for="task in taskStream.today"
          :key="task.id"
          :task="task"
          @complete="handleComplete"
        />
      </div>
    </section>

    <!-- Overdue Tasks (if any) -->
    <section v-if="taskStream.overdue.length > 0" class="mb-8">
      <h2 class="text-2xl font-semibold mb-4 flex items-center gap-2 text-red-600">
        ⚠️ {{ t('syllabus_navigator.overdue_tasks') }}
        <span class="text-sm font-normal">
          ({{ taskStream.overdue.length }})
        </span>
      </h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <TaskCard
          v-for="task in taskStream.overdue"
          :key="task.id"
          :task="task"
          :overdue="true"
          @complete="handleComplete"
        />
      </div>
    </section>

    <!-- Upcoming Tasks -->
    <section>
      <h2 class="text-2xl font-semibold mb-4 flex items-center gap-2">
        📆 {{ t('syllabus_navigator.upcoming_tasks') }}
        <span class="text-sm font-normal text-gray-500">
          ({{ t('syllabus_navigator.next_7_days') }})
        </span>
      </h2>
      <div v-if="taskStream.upcoming.length === 0" class="text-gray-500 text-center py-8">
        {{ t('syllabus_navigator.no_upcoming_tasks') }}
      </div>
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <TaskCard
          v-for="task in taskStream.upcoming"
          :key="task.id"
          :task="task"
          :upcoming="true"
        />
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import { push } from 'notivue';
import taskService from '@/services/taskService';
import goalService from '@/services/goalService';
import TaskCard from '@/components/TaskCard.vue';
import type { TaskStreamResponse } from '@/models/task.model';
import type { Goal } from '@/models/goal.model';

const { t } = useI18n();
const loading = ref(false);
const activeGoal = ref<Goal | null>(null);
const taskStream = ref<TaskStreamResponse>({
  today: [],
  upcoming: [],
  overdue: []
});

onMounted(async () => {
  await loadActiveGoal();
  await loadTasks();
});

const loadActiveGoal = async () => {
  try {
    const goals = await goalService.getActive();
    if (goals.length > 0) {
      activeGoal.value = goals[0];
    }
  } catch (error) {
    push.error({ message: t('errors.load_goal_failed') });
  }
};

const loadTasks = async () => {
  if (!activeGoal.value) return;
  
  loading.value = true;
  try {
    taskStream.value = await taskService.getStream(activeGoal.value.id);
  } catch (error) {
    push.error({ message: t('errors.load_tasks_failed') });
  } finally {
    loading.value = false;
  }
};

const handleComplete = async (taskId: number, success: boolean) => {
  try {
    await taskService.complete({ id: taskId, success });
    push.success({ message: t('syllabus_navigator.task_completed') });
    await loadTasks(); // Refresh
  } catch (error) {
    push.error({ message: t('errors.complete_task_failed') });
  }
};

const formatDate = (date: Date) => {
  return date.toLocaleDateString(undefined, { 
    weekday: 'long', 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric' 
  });
};
</script>
```

---

## Implementation Checklist

### Services (4 files - 2-3 hours)
- [ ] goalService.ts
- [ ] taskService.ts
- [ ] attemptService.ts
- [ ] knowledgeStateService.ts

### Student Views (8 files - 10-12 hours)
- [ ] GoalCreation.vue
- [ ] GoalDashboard.vue
- [ ] TaskStream.vue
- [ ] ProgressDashboard.vue
- [ ] KnowledgeHeatmap.vue
- [ ] ReviewDue.vue
- [ ] PracticeSession.vue
- [ ] PlanGenerator.vue

### Components (5 files - 5-6 hours)
- [ ] TaskCard.vue
- [ ] ProgressChart.vue
- [ ] MasteryHeatmap.vue
- [ ] ReviewReminder.vue
- [ ] GoalCard.vue

### Routing (1 hour)
- [ ] Add student routes
- [ ] Add navigation links
- [ ] Protect routes with auth

### I18n (2 hours)
- [ ] Add 100+ translation keys for en.ts
- [ ] Add 100+ translation keys for zh.ts
- [ ] Task types, statuses, metrics

### Integration (3-4 hours)
- [ ] Connect practice flow with attempt logging
- [ ] Integrate task completion with knowledge state updates
- [ ] Add notifications for review reminders
- [ ] Test complete user flow

**Total Estimated Time**: 25-30 hours for complete implementation

---

## Backend API Reference

All APIs are POST-based, authenticated with JWT:

### Goal APIs
- `POST /api/v1/goal/create`
- `POST /api/v1/goal/edit`
- `POST /api/v1/goal/byId`
- `POST /api/v1/goal/list`
- `POST /api/v1/goal/active`
- `POST /api/v1/goal/delete`
- `POST /api/v1/goal/diagnostic` (optional)
- `POST /api/v1/goal/complete` (optional)

### Task APIs
- `POST /api/v1/task/create`
- `POST /api/v1/task/edit`
- `POST /api/v1/task/byId`
- `POST /api/v1/task/list`
- `POST /api/v1/task/stream` ⭐ Main API
- `POST /api/v1/task/complete`
- `POST /api/v1/task/generate-plan` ⭐
- `POST /api/v1/task/delete`

### Attempt APIs
- `POST /api/v1/attempt/create`
- `POST /api/v1/attempt/batch-create` ⭐
- `POST /api/v1/attempt/recent`
- `POST /api/v1/attempt/stats`

### Knowledge State APIs
- `POST /api/v1/knowledge-state/byChapter`
- `POST /api/v1/knowledge-state/list`
- `POST /api/v1/knowledge-state/progress` ⭐
- `POST /api/v1/knowledge-state/due-review` ⭐

---

## Key Design Patterns

### 1. Service Layer Pattern
```typescript
class ServiceName {
  private readonly basePath = '/api/v1/resource';
  
  async method(params): Promise<ReturnType> {
    const response = await apiClient.post(`${this.basePath}/endpoint`, params);
    return response.data;
  }
}

export default new ServiceName();
```

### 2. Vue Composition API Pattern
```vue
<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import { push } from 'notivue';
import service from '@/services/service';

const { t } = useI18n();
const loading = ref(false);
const data = ref([]);

onMounted(async () => {
  await loadData();
});

const loadData = async () => {
  loading.value = true;
  try {
    data.value = await service.list();
  } catch (error) {
    push.error({ message: t('errors.load_failed') });
  } finally {
    loading.value = false;
  }
};
</script>
```

### 3. Task Card Component Pattern
```vue
<template>
  <div class="border rounded-lg p-4 hover:shadow-lg transition cursor-pointer"
       :class="cardClass"
       @click="handleClick">
    <!-- Task Type Icon -->
    <div class="text-3xl mb-2">{{ taskIcon }}</div>
    
    <!-- Task Title -->
    <h3 class="font-semibold text-lg mb-1">{{ task.title }}</h3>
    
    <!-- Task Meta -->
    <div class="text-sm text-gray-600 space-y-1">
      <div>📖 {{ task.chapter?.name }}</div>
      <div>⏱️ {{ task.estimatedMinutes }} {{ t('common.minutes') }}</div>
      <div>📅 {{ formatDate(task.scheduledDate) }}</div>
    </div>
    
    <!-- Actions -->
    <div class="mt-4 flex gap-2">
      <button @click.stop="$emit('complete', task.id, true)"
              class="flex-1 px-3 py-2 bg-green-600 text-white rounded hover:bg-green-700">
        ✓ {{ t('common.complete') }}
      </button>
      <button @click.stop="$emit('skip', task.id)"
              class="px-3 py-2 border rounded hover:bg-gray-50">
        {{ t('common.skip') }}
      </button>
    </div>
  </div>
</template>
```

---

## Progress Metrics Visualization

### Coverage (0-1)
```
Covered Chapters / Total Chapters
Visual: Progress bar with percentage
```

### Mastery (0-1)
```
Average(Chapter Mastery Levels)
Visual: Circular progress or gauge chart
Color: Red (0-0.5), Yellow (0.5-0.75), Green (0.75-1.0)
```

### Stability (0-1)
```
Average(Chapter Stability Scores)
Visual: Line chart over time
Shows memory retention trend
```

---

## SRS (Spaced Repetition System) UI

### Review Due Indicator
```vue
<div v-if="dueReviewCount > 0" class="bg-yellow-100 border-l-4 border-yellow-500 p-4">
  <div class="flex items-center">
    <span class="text-2xl mr-3">🔄</span>
    <div>
      <p class="font-semibold">{{ dueReviewCount }} {{ t('syllabus_navigator.chapters_need_review') }}</p>
      <button @click="goToReview" class="text-blue-600 hover:underline">
        {{ t('syllabus_navigator.review_now') }} →
      </button>
    </div>
  </div>
</div>
```

### Next Review Schedule
```vue
<div class="space-y-2">
  <div v-for="chapter in upcomingReviews" :key="chapter.id"
       class="flex justify-between items-center p-3 border rounded">
    <span>{{ chapter.name }}</span>
    <span class="text-sm text-gray-600">
      {{ t('syllabus_navigator.review_in') }} {{ chapter.daysUntilReview }} {{ t('common.days') }}
    </span>
  </div>
</div>
```

---

## Testing Recommendations

### Unit Tests
- Service methods (mock axios)
- Component rendering
- Event emissions
- Computed properties

### Integration Tests
- Complete user flows
- Goal creation → Task generation → Practice → Progress update
- Error handling
- Loading states

### E2E Tests
- Full learning cycle
- Multi-day workflow
- SRS scheduling
- Progress tracking accuracy

---

## Performance Optimizations

### 1. Lazy Loading
```typescript
const ProgressDashboard = () => import('@/views/student/ProgressDashboard.vue');
```

### 2. Caching
```typescript
const cachedGoal = ref<Goal | null>(null);
const cacheTimestamp = ref(0);
const CACHE_TTL = 5 * 60 * 1000; // 5 minutes

const getGoal = async (id: number, force = false) => {
  if (!force && cachedGoal.value && Date.now() - cacheTimestamp.value < CACHE_TTL) {
    return cachedGoal.value;
  }
  cachedGoal.value = await goalService.getById(id);
  cacheTimestamp.value = Date.now();
  return cachedGoal.value;
};
```

### 3. Debouncing
```typescript
import { debounce } from 'lodash-es';

const handleSearch = debounce(async (query: string) => {
  // Search logic
}, 300);
```

---

## Accessibility (A11y)

### ARIA Labels
```vue
<button 
  aria-label="Complete task"
  :aria-pressed="task.status === 'completed'"
>
  ✓
</button>
```

### Keyboard Navigation
```vue
<div 
  tabindex="0"
  @keydown.enter="handleClick"
  @keydown.space.prevent="handleClick"
>
  <!-- Content -->
</div>
```

### Screen Reader Support
```vue
<span class="sr-only">
  {{ t('syllabus_navigator.task_description_for_screen_reader') }}
</span>
```

---

## Mobile Responsive Design

All views should be mobile-first with breakpoints:
- `sm`: 640px
- `md`: 768px
- `lg`: 1024px
- `xl`: 1280px

Example:
```vue
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
  <!-- Cards adapt to screen size -->
</div>
```

---

## Deployment Checklist

- [ ] All services implemented
- [ ] All views created
- [ ] All components functional
- [ ] I18n complete (en + zh)
- [ ] Routing configured
- [ ] Error handling comprehensive
- [ ] Loading states everywhere
- [ ] Mobile responsive
- [ ] Accessibility tested
- [ ] Performance optimized
- [ ] E2E tests pass
- [ ] Documentation updated

---

## Estimated Completion

**Phase 1 (MVP)**: 15-20 hours
- Goal creation, Task stream, Basic progress

**Phase 2 (Intelligence)**: 10-12 hours  
- Knowledge state viz, SRS, Advanced metrics

**Phase 3 (Polish)**: 5-8 hours
- Mobile responsive, A11y, Performance

**Total**: 30-40 hours for complete, production-ready implementation

---

## Support & References

- Backend API Documentation: `docs/SYLLABUS_NAVIGATOR_API.md`
- Business Logic: `docs/BUSINESS_LOGIC_USER_PERSPECTIVE.md`
- Architecture: `docs/IMPLEMENTATION_SUMMARY.md`
- Existing Pattern: Knowledge Point Management (reference implementation)

---

**Status**: Implementation guide complete. Ready for development team to proceed with frontend implementation following this comprehensive plan.
