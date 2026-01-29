import apiClient from './apiClient'
import type { ApiResponse } from '../models/api.model'
import type { Syllabus, SyllabusListResponse, SyllabusQuery, SyllabusResponse } from '../models/syllabus.model'

class SyllabusService {
  async getSyllabuses(query: SyllabusQuery = { pageSize: 20, pageIndex: 1 }): Promise<SyllabusListResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/syllabus/list', query)
    return response.data
  }

  async getAllSyllabuses(query: SyllabusQuery): Promise<SyllabusListResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/syllabus/all', query)
    // If the API returns just an array, wrap it in the expected format

    return response.data
  }

  async getSyllabusById(id: number): Promise<SyllabusResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/syllabus/byId', { id })
    return response.data
  }

  async createSyllabus(syllabus: Partial<Syllabus>): Promise<SyllabusResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/syllabus/create', syllabus)
    return response.data
  }

  async updateSyllabus(syllabus: Partial<Syllabus>): Promise<SyllabusResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/syllabus/edit', syllabus)
    return response.data
  }

  async deleteSyllabus(id: number): Promise<ApiResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/syllabus/delete', { id })
    return response.data
  }
}

export default new SyllabusService()
