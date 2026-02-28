import type { ApiResponse } from './api.model'

export const CLASS_STUDENT_STATUS_STUDYING    = 1 // 在读
export const CLASS_STUDENT_STATUS_GRADUATED   = 2 // 结业
export const CLASS_STUDENT_STATUS_TRANSFERRED = 3 // 转走
export const CLASS_STUDENT_STATUS_DROPPED     = 4 // 弃科

export interface ClassStudent {
  id: number
  username: string
  nickname: string
  realname: string
  email: string
  mobile: string
  avatar: string
  studentStatus: number // 1:在读, 2:结业, 3:转走, 4:弃科
}

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

export interface PaginatedClassStudents {
  list: ClassStudent[]
  total: number
}

export type ClassListResponse = ApiResponse<PaginatedClasses>
export type ClassResponse = ApiResponse<Class>
export type ClassStudentListResponse = ApiResponse<PaginatedClassStudents>
