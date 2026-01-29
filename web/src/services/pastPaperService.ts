// ltedu-web/src/services/pastPaperService.ts
import apiClient from './apiClient'
import type { ApiResponse } from '../models/api.model'
import type {
  PastPaperCreateRequest,
  PastPaperListResponse,
  PastPaperQuery,
  PastPaperResponse,
  PastPaperUpdateRequest,
} from '../models/pastPaper.model'

class PastPaperService {
  private baseUrl = '/api/v1/paper/past'

  async getPastPapers(query: PastPaperQuery = {}): Promise<PastPaperListResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<PastPaperListResponse>(`${this.baseUrl}/list`, query)
    return response.data
  }

  async getPastPaperById(id: number): Promise<PastPaperResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<PastPaperResponse>(`${this.baseUrl}/getById`, { id })
    return response.data
  }

  async getAllPastPapers(query: PastPaperQuery = {}): Promise<PastPaperListResponse> {
    const client = await apiClient(); // Ensure we use the async client
    // Assuming the backend /all endpoint for PastPaper returns PaginatedPastPapers
    // If it returns PastPaper[] directly, adjust PastPaperListResponse or create a new type
    const response = await client.post<PastPaperListResponse>(`${this.baseUrl}/all`, query)
    return response.data
  }

  async createPastPaper(entity: PastPaperCreateRequest): Promise<PastPaperResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<PastPaperResponse>(`${this.baseUrl}/create`, entity)
    return response.data
  }

  async updatePastPaper(entity: PastPaperUpdateRequest): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<ApiResponse<void>>(`${this.baseUrl}/edit`, entity)
    return response.data
  }

  async deletePastPaper(id: string): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<ApiResponse<void>>(`${this.baseUrl}/delete`, { id })
    return response.data
  }
}

export default new PastPaperService()
