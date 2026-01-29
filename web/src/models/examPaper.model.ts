import type { Syllabus } from './syllabus.model'
import type { Question } from './question.model'
import type { ApiResponse } from './api.model'
import type { User } from './user.model'

export interface ExamPaper {
  id: number
  name: string
  syllabusId: number
  syllabus: Syllabus
  year: number
  questions?: Question[]
  questionIds: number[]
  userId: number
  user?: User
  createdAt: string
  updatedAt: string
  deletedAt?: string | null
}

export interface ExamPaperQuery {
  id?: number
  syllabusId?: number
  name?: string
  year?: number
  userId?: number
  pageSize?: number
  pageIndex?: number
}

export interface ExamPaperCreateRequest {
  name: string
  syllabusId: number
  questionIds: number[]
}

export interface ExamPaperUpdateRequest {
  id: number
  name?: string
  syllabusId?: number
  year?: number
  questionIds?: number[]
}

export interface PaginatedExamPapers {
  list: ExamPaper[]
  total: number
}

export type ExamPaperListResponse = ApiResponse<PaginatedExamPapers>
export type ExamPaperResponse = ApiResponse<ExamPaper>
