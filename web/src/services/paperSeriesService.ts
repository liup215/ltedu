import apiClient from './apiClient'
import type { ApiResponse } from '../models/api.model'
import type {
  PaperSeriesQuery,
  PaperSeriesListResponse,
  SinglePaperSeriesResponse,
  PaperSeriesCreateRequest,
  PaperSeriesUpdateRequest
} from '../models/paperSeries.model'


class PaperSeriesService {
  private paperSeriesBaseUrl = '/api/v1/pastPaper/series'

  // --- PaperSeries ---
  async getPaperSeriesList(query: PaperSeriesQuery = {}): Promise<PaperSeriesListResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<PaperSeriesListResponse>(`${this.paperSeriesBaseUrl}/list`, query)
    return response.data
  }

  async getPaperSeriesById(id: number): Promise<SinglePaperSeriesResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<SinglePaperSeriesResponse>(`${this.paperSeriesBaseUrl}/getById`, { id })
    return response.data
  }

  async getAllPaperSeries(query: PaperSeriesQuery = {}): Promise<PaperSeriesListResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<PaperSeriesListResponse>(`${this.paperSeriesBaseUrl}/all`, query)
    return response.data
  }

  async createPaperSeries(entity: PaperSeriesCreateRequest): Promise<SinglePaperSeriesResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<SinglePaperSeriesResponse>(`${this.paperSeriesBaseUrl}/create`, entity)
    return response.data
  }

  async updatePaperSeries(entity: PaperSeriesUpdateRequest): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<ApiResponse<void>>(`${this.paperSeriesBaseUrl}/edit`, entity)
    return response.data
  }

  async deletePaperSeries(id: number): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post<ApiResponse<void>>(`${this.paperSeriesBaseUrl}/delete`, { id })
    return response.data
  }

}

export default new PaperSeriesService()
