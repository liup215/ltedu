import type { ApiResponse } from './api.model'

export interface StudentLearningPlan {
  id: number
  classId: number
  userId: number
  planType: string  // "long" | "mid" | "short"
  content: string
  version: number
  createdBy: number
  isPersonal: boolean
  createdAt: string
  updatedAt: string
  class?: any
  user?: any
}

export interface StudentLearningPlanVersion {
  id: number
  planId: number
  version: number
  content: string
  changedBy: number
  comment: string
  createdAt: string
}

export interface LearningPhasePlan {
  id: number
  planId: number
  examNodeId: number
  title: string
  startDate?: string
  endDate?: string
  sortOrder: number
  chapters?: any[]
  examNode?: any
  createdAt: string
  updatedAt: string
}

export type LearningPlanListResponse = ApiResponse<{ list: StudentLearningPlan[]; total: number }>
export type LearningPlanResponse = ApiResponse<StudentLearningPlan>
export type LearningPlanVersionListResponse = ApiResponse<{ list: StudentLearningPlanVersion[]; total: number }>
export type PhasePlanListResponse = ApiResponse<{ list: LearningPhasePlan[]; total: number }>
export type PhasePlanResponse = ApiResponse<LearningPhasePlan>
