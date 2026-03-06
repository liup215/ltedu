/**
 * API response model matching the LTEdu backend response format.
 * Backend always returns HTTP 200 with code 0 for success, 1 for error.
 */
export interface ApiResponse<T> {
  code: number;
  msg: string;
  data: T;
}

export interface PaginatedList<T> {
  list: T[];
  total: number;
}

export interface User {
  id: number;
  username: string;
  email: string;
  phone?: string;
  avatar?: string;
  isAdmin?: boolean;
}

export interface LoginCredentials {
  username: string;
  password: string;
}

export interface LoginResponseData {
  token: string;
  user: User;
}

export interface ChatMessage {
  id: string;
  role: 'user' | 'assistant';
  content: string;
  timestamp: number;
}

export interface Question {
  id: number;
  type: number;
  stem: string;
  analyze?: string;
  answer?: string;
  options?: QuestionOption[];
}

export interface QuestionOption {
  label: string;
  content: string;
}

export interface PracticeResult {
  questions: Question[];
  total: number;
}

export interface LearningPlan {
  id: number;
  planType: string;
  content: string;
  comment?: string;
  createdAt: string;
  updatedAt: string;
}

export interface LearningPlanListResponse {
  list: LearningPlan[];
  total: number;
}

export interface QuickPracticeRequest {
  qualificationId?: number;
  syllabusId?: number;
  count?: number;
  questionType?: number;
}

export interface GradeRequest {
  answers: Record<string, string>;
  practiceId?: number;
}
