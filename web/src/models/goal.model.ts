// Goal Model - Learning objectives for Syllabus Navigator
export interface Goal {
  id: number;
  userId: number;
  syllabusId: number;
  syllabus?: any; // Syllabus object
  name: string;
  description?: string;
  examDate?: string;
  targetScore?: number;
  weeklyHours?: number;
  mode?: 'self-paced' | 'school-sync' | 'cram';
  status: 'active' | 'completed' | 'paused' | 'abandoned';
  createdAt: string;
  updatedAt: string;
}

export interface CreateGoalRequest {
  syllabusId: number;
  name: string;
  description?: string;
  examDate?: string;
  targetScore?: number;
  weeklyHours?: number;
  mode?: string;
}

export interface UpdateGoalRequest {
  id: number;
  name?: string;
  description?: string;
  examDate?: string;
  targetScore?: number;
  weeklyHours?: number;
  mode?: string;
  status?: string;
}

export interface GoalListResponse {
  total: number;
  goals: Goal[];
}

export interface GoalProgressResponse {
  goalId: number;
  coverage: number; // 0-1
  mastery: number; // 0-1
  stability: number; // 0-1
  totalChapters: number;
  coveredChapters: number;
  estimatedCompletionDate?: string;
}
