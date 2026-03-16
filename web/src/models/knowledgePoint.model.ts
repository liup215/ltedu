export interface KnowledgePoint {
  id: number
  chapterId: number
  name: string
  description: string
  difficulty: 'basic' | 'medium' | 'hard'
  estimatedMinutes: number
  orderIndex: number
  createdAt: string
  updatedAt: string
  chapter?: {
    id: number
    name: string
    syllabusId: number
  }
}

export interface KnowledgePointCreateRequest {
  chapterId: number
  name: string
  description: string
  difficulty: 'basic' | 'medium' | 'hard'
  estimatedMinutes: number
  orderIndex?: number
}

export interface KnowledgePointUpdateRequest {
  id: number
  name: string
  description: string
  difficulty: 'basic' | 'medium' | 'hard'
  estimatedMinutes: number
  orderIndex?: number
}

export interface KnowledgePointListRequest {
  syllabusId?: number
  chapterId?: number
  page?: number
  pageSize?: number
}

export interface KnowledgePointListResponse {
  code: number
  msg: string
  data: {
    list: KnowledgePoint[]
    total: number
    page: number
    pageSize: number
  }
}

export interface KnowledgePointResponse {
  code: number
  msg: string
  data: KnowledgePoint
}

export interface GenerateKeypointsRequest {
  chapterId: number
  mode?: 'auto' | 'manual'
}

export interface GenerateKeypointsResponse {
  code: number
  msg: string
  data: {
    keypoints: KnowledgePoint[]
    count: number
  }
}

export interface AutoLinkRequest {
  questionId: number
  chapterId?: number
  syllabusId?: number
}

export interface AutoLinkResponse {
  code: number
  msg: string
  data: {
    linkedKeypoints: number[]
    count: number
  }
}

export interface MigrateOptions {
  generateKeypoints: boolean
  linkQuestions: boolean
  batchSize?: number
}

export interface MigrateSyllabusRequest {
  syllabusId: number
  options: MigrateOptions
}

export interface MigrateReport {
  generatedKeypoints: number
  linkedQuestions: number
  totalLinks: number
  errors: string[]
}

export interface MigrateSyllabusResponse {
  code: number
  msg: string
  data: MigrateReport
}

// Migration Job types
export type MigrationJobStatus = 'pending' | 'running' | 'completed' | 'failed'

export interface MigrationJob {
  id: number
  syllabusId: number
  status: MigrationJobStatus
  options: string  // JSON string
  progress: number
  totalItems: number
  doneItems: number
  report: string  // JSON string
  errorMessage: string
  createdBy: number
  startedAt?: string
  completedAt?: string
  createdAt: string
  updatedAt: string
}

export interface MigrationJobCreateRequest {
  syllabusId: number
  options: MigrateOptions
}

export interface MigrationJobListRequest {
  syllabusId?: number
  status?: MigrationJobStatus
  pageIndex?: number
  pageSize?: number
}

export interface MigrationJobResponse {
  code: number
  msg: string
  data: MigrationJob
}

export interface MigrationJobListResponse {
  code: number
  msg: string
  data: {
    list: MigrationJob[]
    total: number
    page: number
    pageSize: number
  }
}
