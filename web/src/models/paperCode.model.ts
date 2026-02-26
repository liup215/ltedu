import { type Syllabus } from './syllabus.model'
import { type ApiResponse, type Page } from './api.model'

export const PAPER_CODE_LEVEL_AS = 'AS';
export const PAPER_CODE_LEVEL_A2 = 'A2';

export interface PaperCode {
  id?: number
  name: string // As per user feedback
  level?: string // syllabus level: "AS", "A2", or "" for non-A-Level
  syllabusId: number
  syllabus?: Syllabus
  createdAt?: string
  updatedAt?: string
}

export interface PaperCodeQuery extends Page {
  syllabusId?: number
  name?: string // For searching by name
}

export interface PaginatedPaperCodes {
  list: PaperCode[]
  total: number
}

export interface PaperCodeCreateRequest {
  name: string
  level?: string
  syllabusId: number
}

export interface PaperCodeUpdateRequest {
  id: number
  name: string
  level?: string
  syllabusId: number
}

export type PaperCodeListResponse = ApiResponse<PaginatedPaperCodes>

export interface SinglePaperCodeResponse extends ApiResponse<PaperCode> {}
