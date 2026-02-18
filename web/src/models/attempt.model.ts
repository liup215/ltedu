// Attempt Model - Records of user's practice attempts
export interface Attempt {
  id: number;
  userId: number;
  goalId: number;
  chapterId?: number;
  questionId: number;
  question?: any;
  answer: string;
  correct: boolean;
  timeSpent: number; // seconds
  score?: number;
  createdAt: string;
}

export interface CreateAttemptRequest {
  goalId: number;
  chapterId?: number;
  questionId: number;
  answer: string;
  correct: boolean;
  timeSpent: number;
  score?: number;
}

export interface BatchCreateAttemptsRequest {
  goalId: number;
  chapterId?: number;
  attempts: {
    questionId: number;
    answer: string;
    correct: boolean;
    timeSpent: number;
    score?: number;
  }[];
}

export interface AttemptStatsResponse {
  totalAttempts: number;
  correctAttempts: number;
  accuracy: number;
  totalTimeSpent: number;
  avgTimePerQuestion: number;
  byChapter: {
    chapterId: number;
    chapterName: string;
    attempts: number;
    correct: number;
    accuracy: number;
  }[];
}

export interface RecentAttemptsResponse {
  attempts: Attempt[];
  total: number;
}
