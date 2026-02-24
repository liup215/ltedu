// Knowledge State Model - Tracks user mastery per chapter
export interface KnowledgeState {
  id: number;
  goalId: number;
  chapterId: number;
  chapter?: any;
  masteryLevel: number; // 0-1
  stabilityScore: number; // 0-1
  totalCount: number;
  correctCount: number;
  consecutiveCorrect: number;
  consecutiveWrong: number;
  lastPracticeAt?: string;
  nextReviewAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface KnowledgeStateByChapterRequest {
  goalId: number;
  chapterId: number;
}

export interface KnowledgeStateListRequest {
  goalId: number;
  page?: number;
  perPage?: number;
}

export interface KnowledgeStateListResponse {
  states: KnowledgeState[];
  total: number;
}

export interface ProgressResponse {
  goalId: number;
  coverage: number; // 0-1, % of chapters touched
  mastery: number; // 0-1, average mastery level
  stability: number; // 0-1, average stability score
  totalChapters: number;
  coveredChapters: number;
  byChapter: {
    chapterId: number;
    chapterName: string;
    masteryLevel: number;
    stabilityScore: number;
    totalAttempts: number;
    lastPracticeAt?: string;
  }[];
}

export interface DueReviewRequest {
  goalId: number;
}

export interface DueReviewResponse {
  chapters: {
    chapterId: number;
    chapterName: string;
    masteryLevel: number;
    stabilityScore: number;
    nextReviewAt: string;
    daysSinceLastPractice: number;
  }[];
  total: number;
}
