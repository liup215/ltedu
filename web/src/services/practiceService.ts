import apiClient from './apiClient'
import type { PracticeQuickRequest, PracticeQuickListResponse, PracticePaperRequest, PracticePaperListResponse, PracticeGradeRequest, PracticeGradeApiResponse } from '../models/practice.model'

export const practiceService = {
  async quickPractice(data: PracticeQuickRequest): Promise<PracticeQuickListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/practice/quick', data)
    return response.data as PracticeQuickListResponse
  },

  async paperPractice(data: PracticePaperRequest): Promise<PracticePaperListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/practice/paper', data)
    return response.data as PracticePaperListResponse
  },

  async gradePractice(data: PracticeGradeRequest): Promise<PracticeGradeApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/practice/grade', data)
    return response.data as PracticeGradeApiResponse
  }
}
