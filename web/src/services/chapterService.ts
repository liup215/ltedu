import apiClient from './apiClient'
import type { 
  ChapterResponse,
  ChapterTreeResponse,
  ChapterCreateRequest,
  ChapterUpdateRequest
} from '../models/chapter.model'

class ChapterService {
  async getChapterTree(syllabusId: number): Promise<ChapterTreeResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/chapter/tree', { syllabusId })
    return response.data
  }

  async getChapterById(id: number): Promise<ChapterResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/chapter/byId', { id })
    return response.data
  }

  async createChapter(chapter: ChapterCreateRequest): Promise<ChapterResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/chapter/create', chapter)
    return response.data
  }

  async updateChapter(chapter: ChapterUpdateRequest): Promise<ChapterResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/chapter/edit', chapter)
    return response.data
  }

  async deleteChapter(id: number): Promise<ChapterResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/chapter/delete', { id })
    return response.data
  }
}

export default new ChapterService()
