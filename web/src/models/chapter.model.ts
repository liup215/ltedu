import type { ApiResponse } from './api.model';
import type { Syllabus } from './syllabus.model';

export interface Chapter {
  id: number;
  name: string;
  syllabusId: number;
  parentId: number;
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
}

export interface ChapterUpdateRequest {
  id: number;
  name: string;
}

export interface PaginatedChapters {
  list: Chapter[];
  total: number;
}

export type ChapterListResponse = ApiResponse<PaginatedChapters>;
export type ChapterResponse = ApiResponse<Chapter>;
export type ChapterTreeResponse = ApiResponse<Chapter[]>;
