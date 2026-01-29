import apiClient from './apiClient'
import type { ExamPaperQuery, ExamPaperCreateRequest, ExamPaperUpdateRequest, ExamPaperListResponse, ExamPaperResponse } from '../models/examPaper.model'
import type { ApiResponse } from '../models/api.model'

export const examPaperService = {
  async getExamPaperList(query: ExamPaperQuery): Promise<ExamPaperListResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/paper/exam/list', query)
    return response.data
  },

  async getExamPaperById(query: { id: number }): Promise<ExamPaperResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/paper/exam/byId', query)
    return response.data
  },

  async getAllExamPapers(query: ExamPaperQuery): Promise<ExamPaperListResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/paper/exam/all', query)
    return response.data
  },

  async createExamPaper(data: ExamPaperCreateRequest): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/paper/exam/create', data)
    return response.data
  },

  async updateExamPaper(data: ExamPaperUpdateRequest): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/paper/exam/edit', data)
    return response.data
  },

  async deleteExamPaper(data: { id: number }): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/paper/exam/delete', data)
    return response.data
  },

  async updateExamPaperQuestions(data: { id: number; questionIds: number[] }): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/paper/exam/updateQuestion', data)
    return response.data
  },

  async exportToWord(id: number): Promise<Blob> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/paper/exam/export', { id }, { responseType: 'blob' })
    return response.data
  }
}
