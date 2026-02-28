import apiClient from './apiClient'
import type {
  ClassListResponse,
  ClassResponse,
  ClassStudentListResponse,
  ClassQuery,
  ClassCreateRequest,
  ClassUpdateRequest,
  BindSyllabusRequest,
  UnbindSyllabusRequest
} from '../models/class.model'
import type { ApiResponse } from '../models/api.model'

class ClassService {
  async list(query: ClassQuery = { pageSize: 20, pageIndex: 1 }): Promise<ClassListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/school/class/list', query)
    return response.data
  }

  async getById(id: number): Promise<ClassResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/school/class/byId', { id })
    return response.data
  }

  async create(req: ClassCreateRequest): Promise<ClassResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/school/class/create', req)
    return response.data
  }

  async update(req: ClassUpdateRequest): Promise<ClassResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/school/class/edit', req)
    return response.data
  }

  async delete(id: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/school/class/delete', { id })
    return response.data
  }

  async bindSyllabus(req: BindSyllabusRequest): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/school/class/bindSyllabus', req)
    return response.data
  }

  async unbindSyllabus(req: UnbindSyllabusRequest): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/school/class/unbindSyllabus', req)
    return response.data
  }

  async joinClass(inviteCode: string, userId: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/school/class/join', { inviteCode, userId })
    return response.data
  }

  async listMembers(classId: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/school/class/members', { classId })
    return response.data
  }

  async getStudents(classId: number): Promise<ClassStudentListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/school/class/studentList', { id: classId })
    return response.data
  }

  async updateStudentStatus(classId: number, userId: number, status: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/school/class/updateStudentStatus', { classId, userId, status })
    return response.data
  }
}

export default new ClassService()
