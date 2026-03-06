import apiClient from './apiClient'
import type { ApiResponse } from '../models/api.model'

export interface ClassAnalyticsQuery {
  classId: number
  goalId?: number
}

export interface StudentAnalyticsQuery {
  goalId: number
}

export interface TrendQuery {
  classId: number
  goalId?: number
  startDate?: string
  endDate?: string
}

export interface ClassPerformanceSummary {
  classId: number
  className: string
  totalStudents: number
  activeStudents: number
  avgMastery: number
  avgCoverage: number
  avgAccuracy: number
  totalAttempts: number
  weeklyAttempts: number
  atRiskCount: number
}

export interface StudentPerformanceSummary {
  userId: number
  userName: string
  masteryLevel: number
  coverageLevel: number
  accuracyRate: number
  totalAttempts: number
  weeklyAttempts: number
  consecWrong: number
  lastActiveAt: string | null
  isAtRisk: boolean
  riskReasons: string[]
}

export interface StudentChapterScore {
  userId: number
  masteryLevel: number
  isCovered: boolean
}

export interface ClassHeatmapRow {
  chapterId: number
  chapterName: string
  avgMastery: number
  studentData: StudentChapterScore[]
}

export interface StudentInfo {
  userId: number
  userName: string
}

export interface ClassHeatmap {
  students: StudentInfo[]
  chapters: ClassHeatmapRow[]
}

export interface AttemptTrendPoint {
  date: string
  totalAttempts: number
  correctAttempts: number
}

export interface EarlyWarningStudent {
  userId: number
  userName: string
  reasons: string[]
  severity: 'low' | 'medium' | 'high'
}

export interface RecommendedChapter {
  questionId: number
  chapterId: number
  chapterName: string
  reason: string
  priority: number
  masteryLevel: number
}

export interface RecommendationResponse {
  goalId: number
  questions: RecommendedChapter[]
  reviewCount: number
  gapCount: number
}

class AnalyticsService {
  async getClassSummary(classId: number): Promise<ApiResponse<ClassPerformanceSummary>> {
    const client = await apiClient()
    const response = await client.post('/api/v1/analytics/class/summary', { classId })
    return response.data
  }

  async getStudentPerformanceList(classId: number): Promise<ApiResponse<{ list: StudentPerformanceSummary[]; total: number }>> {
    const client = await apiClient()
    const response = await client.post('/api/v1/analytics/class/students', { classId })
    return response.data
  }

  async getClassHeatmap(classId: number): Promise<ApiResponse<ClassHeatmap>> {
    const client = await apiClient()
    const response = await client.post('/api/v1/analytics/class/heatmap', { classId })
    return response.data
  }

  async getAttemptTrends(query: TrendQuery): Promise<ApiResponse<AttemptTrendPoint[]>> {
    const client = await apiClient()
    const response = await client.post('/api/v1/analytics/class/trends', query)
    return response.data
  }

  async getEarlyWarnings(classId: number): Promise<ApiResponse<{ list: EarlyWarningStudent[]; total: number }>> {
    const client = await apiClient()
    const response = await client.post('/api/v1/analytics/class/earlyWarning', { classId })
    return response.data
  }

  async getStudentSummary(goalId: number): Promise<ApiResponse<StudentPerformanceSummary>> {
    const client = await apiClient()
    const response = await client.post('/api/v1/analytics/student/summary', { goalId })
    return response.data
  }

  async getRecommendations(goalId: number): Promise<ApiResponse<RecommendationResponse>> {
    const client = await apiClient()
    const response = await client.post('/api/v1/analytics/recommend', { goalId })
    return response.data
  }
}

export const analyticsService = new AnalyticsService()
export default analyticsService
