import { type Syllabus } from './syllabus.model'
import { type ApiResponse, type Page } from './api.model'

export interface PaperSeries {
  id?: number
  name: string
  syllabusId: number
  syllabus?: Syllabus
  createdAt?: string
  updatedAt?: string
}

export interface PaperSeriesQuery extends Page {
  syllabusId?: number
  name?: string
}

export interface PaginatedPaperSeries {
  list: PaperSeries[]
  total: number
}

export interface PaperSeriesCreateRequest {
  name: string
  syllabusId: number
}

export interface PaperSeriesUpdateRequest {
  id: number
  name: string
  syllabusId: number
}

export type PaperSeriesListResponse = ApiResponse<PaginatedPaperSeries>

export type SinglePaperSeriesResponse = ApiResponse<PaperSeries>
