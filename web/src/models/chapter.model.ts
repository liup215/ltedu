import type { ApiResponse } from './api.model';
import type { Syllabus } from './syllabus.model';

export const CHAPTER_LEVEL_AS = 'AS';
export const CHAPTER_LEVEL_A2 = 'A2';

export interface Chapter {
  id: number;
  name: string;
  syllabusId: number;
  parentId: number;
  level?: string; // syllabus level: "AS", "A2", or "" for non-A-Level
  syllabus?: Syllabus;
  children?: Chapter[];
  isLeaf?: number;
  createdAt?: string;
  updatedAt?: string;
}

export interface ChapterQuery {
  pageIndex?: number;
  pageSize?: number;
  syllabusId?: number;
  parentId?: number;
}

export interface ChapterCreateRequest {
  name: string;
  syllabusId: number;
  parentId?: number;
  level?: string;
}

export interface ChapterUpdateRequest {
  id: number;
  name: string;
  level?: string;
}

export interface PaginatedChapters {
  list: Chapter[];
  total: number;
}

export type ChapterListResponse = ApiResponse<PaginatedChapters>;
export type ChapterResponse = ApiResponse<Chapter>;
export type ChapterTreeResponse = ApiResponse<Chapter[]>;
