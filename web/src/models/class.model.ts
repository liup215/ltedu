import type { ApiResponse } from './api.model'

export interface Class {
  id: number
  name: string
  classType: number   // 1=teaching, 2=administrative
  syllabusId?: number
  syllabus?: any
  adminUserId: number
  inviteCode: string
  createdAt: string
  updatedAt: string
}

export interface ClassQuery {
  id?: number
  name?: string
  classType?: number
  pageSize?: number
  pageIndex?: number
}

export interface ClassCreateRequest {
  name: string
  classType: number
  syllabusId?: number
}

export interface ClassUpdateRequest {
  id: number
  name?: string
}

export interface BindSyllabusRequest {
  classId: number
  syllabusId: number
}

export interface UnbindSyllabusRequest {
  classId: number
}

export interface PaginatedClasses {
  list: Class[]
  total: number
}

export type ClassListResponse = ApiResponse<PaginatedClasses>
export type ClassResponse = ApiResponse<Class>
