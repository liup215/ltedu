import type { Qualification } from './qualification.model'
import type { Chapter } from './chapter.model'
import type { ApiResponse } from './api.model'

export interface Syllabus {
  id: number
  name: string
  code: string
  qualificationId: number
  qualification: Qualification
  chapters?: Chapter[] // Optional, as it might not always be loaded
  createdAt: string
  updatedAt: string
  deletedAt?: string | null
}

export interface SyllabusQuery {
  id?: number
  name?: string
  code?: string
  qualificationId?: number
  pageSize?: number
  pageIndex?: number
}

export interface SyllabusCreateRequest {
  name: string
  code: string
  qualificationId: number
}

export interface SyllabusUpdateRequest {
  id: number
  name?: string
  code?: string
  qualificationId?: number
}

export interface PaginatedSyllabuses {
  list: Syllabus[]
  total: number
}

export type SyllabusListResponse = ApiResponse<PaginatedSyllabuses>
export type SyllabusResponse = ApiResponse<Syllabus>
