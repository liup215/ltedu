// ltedu-web/src/models/practice.model.ts

export interface PracticeQuickRequest {
  syllabusId: number;
  questionCount: number;
  chapterIds?: number[];
}

export interface PracticeQuickResponse {
  list: number[];
  total: number;
}

export interface PracticePaperRequest {
  paperId: number;
}

export interface PracticePaperResponse {
  list: number[];
}

export type PracticeQuickListResponse = ApiResponse<PracticeQuickResponse>;
export type PracticePaperListResponse = ApiResponse<PracticePaperResponse>;

// Answer for a single part/sub-question
export interface PracticePartAnswer {
  questionContentId: number;
  answer: string;
}

// Submission for a single question (with multiple parts)
export interface PracticeSubmissionAnswer {
  questionId: number;
  answers: PracticePartAnswer[];
}

export type PracticeGradeRequest = PracticeSubmissionAnswer[];

// Grading result for a single sub-question
export interface PracticeSubResultItem {
  questionContentId: number;
  questionType: number;
  correctAnswer: string;
  studentAnswer: string;
  isCorrect: boolean | null;
  modelAnswer?: string;
}

// Grading result for a single question (with multiple sub-questions)
export interface PracticeResultItem {
  questionId: number;
  subResults: PracticeSubResultItem[];
}

// Overall grading response: aggregated score and per-question results
export interface PracticeGradeResponse {
  score: number;
  total: number;
  results: PracticeResultItem[];
}

export type PracticeGradeApiResponse = ApiResponse<PracticeGradeResponse>;

import type { ApiResponse } from './api.model';
