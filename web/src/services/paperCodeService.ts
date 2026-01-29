import apiClient from './apiClient'
import type { ApiResponse } from '../models/api.model'
import type { PaperCodeCreateRequest, PaperCodeListResponse, PaperCodeQuery, PaperCodeUpdateRequest, SinglePaperCodeResponse } from '../models/paperCode.model'

class PaperCodeService {
  private paperCodeBaseUrl = '/api/v1/pastPaper/code'

  // --- PaperCode ---
  async getPaperCodeList(query: PaperCodeQuery = {}): Promise<PaperCodeListResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<PaperCodeListResponse>(`${this.paperCodeBaseUrl}/list`, query)
    return response.data
  }

  async getPaperCodeById(id: number): Promise<SinglePaperCodeResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<SinglePaperCodeResponse>(`${this.paperCodeBaseUrl}/getById`, { id })
    return response.data
  }

  async getAllPaperCodes(query: PaperCodeQuery = {}): Promise<PaperCodeListResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<PaperCodeListResponse>(`${this.paperCodeBaseUrl}/all`, query)
    return response.data
  }

  async createPaperCode(entity: PaperCodeCreateRequest): Promise<SinglePaperCodeResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<SinglePaperCodeResponse>(`${this.paperCodeBaseUrl}/create`, entity)
    return response.data
  }

  async updatePaperCode(entity: PaperCodeUpdateRequest): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<ApiResponse<void>>(`${this.paperCodeBaseUrl}/edit`, entity)
    return response.data
  }

  async deletePaperCode(id: number): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<ApiResponse<void>>(`${this.paperCodeBaseUrl}/delete`, { id })
    return response.data
  }
}

export default new PaperCodeService()
