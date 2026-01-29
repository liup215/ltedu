import type { ApiResponse } from '../models/api.model'
import type { TeacherApplication, TeacherApplicationCreateRequest, TeacherApplicationQuery, TeacherApplicationListResponse, TeacherApplicationResponse } from '../models/teacher-application.model'
import apiClient from './apiClient'

class TeacherApplicationService {
  // Submit a teacher application
  async apply(request: TeacherApplicationCreateRequest): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/user/teacher/apply', request)
    return response.data
  }

  // Get current user's application
  async getCurrentApplication(): Promise<TeacherApplicationResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/user/teacher/application')
    return response.data
  }

  // List applications (admin only)
  async list(query: TeacherApplicationQuery): Promise<TeacherApplicationListResponse> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/admin/teacher-applications/list', query)
    return response.data
  }

  // Get application details (admin only)
  async getById(id: number): Promise<TeacherApplication> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/admin/teacher-applications/detail', { id })
    return response.data
  }

  // Approve application (admin only)
  async approve(id: number, adminNotes: string): Promise<void> {
    const client = await apiClient(); // Ensure we use the async client
    await client.post('/api/v1/admin/teacher-applications/approve', { id, adminNotes })
  }

  // Reject application (admin only)
  async reject(id: number, adminNotes: string): Promise<void> {
    const client = await apiClient(); // Ensure we use the async client
    await client.post('/api/v1/admin/teacher-applications/reject', { id, adminNotes })
  }
}

export const teacherApplicationService = new TeacherApplicationService()
