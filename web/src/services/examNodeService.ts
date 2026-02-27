import apiClient from './apiClient'
import type { ExamNodeListResponse, ExamNodeResponse } from '../models/examNode.model'
import type { ApiResponse } from '../models/api.model'

class ExamNodeService {
  async create(req: { syllabusId: number; name: string; description?: string; sortOrder?: number }): Promise<ExamNodeResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/syllabus/examNode/create', req)
    return response.data
  }

  async update(req: { id: number; name?: string; description?: string; sortOrder?: number }): Promise<ExamNodeResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/syllabus/examNode/edit', req)
    return response.data
  }

  async delete(id: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/syllabus/examNode/delete', { id })
    return response.data
  }

  async getById(id: number): Promise<ExamNodeResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/syllabus/examNode/byId', { id })
    return response.data
  }

  async list(syllabusId: number): Promise<ExamNodeListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/syllabus/examNode/list', { syllabusId })
    return response.data
  }

  async addChapter(examNodeId: number, chapterId: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/syllabus/examNode/chapter/add', { examNodeId, chapterId })
    return response.data
  }

  async removeChapter(examNodeId: number, chapterId: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/syllabus/examNode/chapter/remove', { examNodeId, chapterId })
    return response.data
  }

  async addPaperCode(examNodeId: number, paperCodeId: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/syllabus/examNode/paperCode/add', { examNodeId, paperCodeId })
    return response.data
  }

  async removePaperCode(examNodeId: number, paperCodeId: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/syllabus/examNode/paperCode/remove', { examNodeId, paperCodeId })
    return response.data
  }
}

export default new ExamNodeService()
