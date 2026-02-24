// Task Model - Daily task cards for learning activities
export type TaskType = 'Learn' | 'Drill' | 'Review' | 'Test' | 'Mock';
export type TaskStatus = 'pending' | 'in_progress' | 'completed' | 'failed' | 'skipped';

export interface Task {
  id: number;
  goalId: number;
  chapterId?: number;
  chapter?: any;
  type: TaskType;
  title: string;
  description?: string;
  estimatedMinutes?: number;
  priority: number;
  scheduledDate: string;
  dueDate?: string;
  completedAt?: string;
  status: TaskStatus;
  metadata?: Record<string, any>;
  createdAt: string;
  updatedAt: string;
}

export interface CreateTaskRequest {
  goalId: number;
  chapterId?: number;
  type: TaskType;
  title: string;
  description?: string;
  estimatedMinutes?: number;
  priority?: number;
  scheduledDate: string;
  dueDate?: string;
}

export interface UpdateTaskRequest {
  id: number;
  title?: string;
  description?: string;
  estimatedMinutes?: number;
  priority?: number;
  scheduledDate?: string;
  dueDate?: string;
  status?: TaskStatus;
}

export interface CompleteTaskRequest {
  id: number;
  success: boolean;
  notes?: string;
}

export interface TaskStreamResponse {
  today: Task[];
  upcoming: Task[];
  overdue: Task[];
}

export interface GeneratePlanRequest {
  goalId: number;
  days?: number;
}
