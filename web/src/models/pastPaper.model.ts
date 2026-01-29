import type { ApiResponse } from './api.model'
import type { Syllabus } from './syllabus.model'
import type { PaperCode } from './paperCode.model'
import type { PaperSeries } from './paperSeries.model'

// SyllabusId     uint        `json:"syllabusId"`
// Syllabus       Syllabus    `json:"syllabus"`
// Year           int         `json:"year"`
// PaperCodeId    uint        `json:"paperCodeId" gorm:"index"`
// PaperCode      PaperCode   `json:"paperCode"`
// PaperSeriesId  uint        `json:"paperSeriesId" gorm:"index"`
// PaperSeries    PaperSeries `json:"paperSeries"`
// QuestionNumber int         `json:"questionNumber"`
// Questions      []*Question `json:"questions" gorm:"-"`

export interface PastPaper {
  id: number
  name: string
  year: number
  syllabusId: number
  syllabus: Syllabus
  paperCodeId: number
  paperCode: PaperCode
  paperSeriesId: number
  paperSeries: PaperSeries
  questionNumber: number
  // questions: any[]
}

export interface PastPaperCreateRequest {
  name: string
  year: number
  syllabusId: number
  paperSeriesId: number
  paperCodeId: number
  questionNumber: number
}

export interface PastPaperUpdateRequest extends PastPaperCreateRequest {
  id: number
}

export interface PastPaperQuery {
  id?: number
  name?: string
  syllabusId?: number
  year?: number
  paperCodeId?: number
  paperSeriesId?: number
  pageIndex?: number
  pageSize?: number
}

export interface PaginatedPastPaperResponse {
  list: PastPaper[]
  total: number
}

export type PastPaperListResponse = ApiResponse<PaginatedPastPaperResponse>
export type PastPaperResponse = ApiResponse<PastPaper>
