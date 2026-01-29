import type { ApiResponse } from './api.model'

export const TeacherApplicationStatus = {
  Pending: 1,
  Approved: 2,
  Rejected: 3
} as const

export type TeacherApplicationStatusType = (typeof TeacherApplicationStatus)[keyof typeof TeacherApplicationStatus]

export interface TeacherApplicationUser {
  id: number
  username: string
  email?: string
  realname?: string
}

export interface TeacherApplication {
  id: number
  userId: number
  user?: TeacherApplicationUser
  motivation: string
  experience: string
  status: TeacherApplicationStatusType
  adminNotes?: string
  appliedAt: string
  reviewedAt?: string
  reviewedById?: number
  createdAt: string
  updatedAt: string
}

export interface TeacherApplicationCreateRequest {
  motivation: string
  experience: string
}

export interface TeacherApplicationQuery {
  status?: TeacherApplicationStatusType
  startDate?: string
  endDate?: string
  pageIndex: number
  pageSize: number
}

export interface PaginatedTeacherApplications {
  list: TeacherApplication[]
  total: number
}

export type TeacherApplicationListResponse = ApiResponse<PaginatedTeacherApplications>
export type TeacherApplicationResponse = ApiResponse<TeacherApplication>
