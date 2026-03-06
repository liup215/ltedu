import apiClient from './apiClient'
import type { ApiResponse } from '../models/api.model'

export interface Goal {
  id: number
  userId: number
  syllabusId: number
  syllabus?: { id: number; name: string }
  status: string
  createdAt: string
  updatedAt: string
}

class GoalService {
  async getActiveGoals(): Promise<ApiResponse<Goal[]>> {
    const client = await apiClient()
    const response = await client.post('/api/v1/goal/active', {})
    return response.data
  }

  async list(): Promise<ApiResponse<{ list: Goal[]; total: number }>> {
    const client = await apiClient()
    const response = await client.post('/api/v1/goal/list', { pageIndex: 1, pageSize: 100 })
    return response.data
  }

  async create(req: { syllabusId: number }): Promise<ApiResponse<Goal>> {
    const client = await apiClient()
    const response = await client.post('/api/v1/goal/create', req)
    return response.data
  }
}

export const goalService = new GoalService()
export default goalService
