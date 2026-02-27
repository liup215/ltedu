import type { ApiResponse } from './api.model'

export interface SyllabusExamNode {
  id: number
  syllabusId: number
  name: string
  description?: string
  sortOrder: number
  chapters?: any[]
  paperCodes?: any[]
  createdAt: string
  updatedAt: string
}

export type ExamNodeListResponse = ApiResponse<{ list: SyllabusExamNode[]; total: number }>
export type ExamNodeResponse = ApiResponse<SyllabusExamNode>
