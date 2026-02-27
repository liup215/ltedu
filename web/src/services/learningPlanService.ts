import apiClient from './apiClient'
import type {
  LearningPlanListResponse,
  LearningPlanResponse,
  LearningPlanVersionListResponse,
  PhasePlanListResponse,
  PhasePlanResponse
} from '../models/learningPlan.model'
import type { ApiResponse } from '../models/api.model'

class LearningPlanService {
  async create(req: { classId: number; userId: number; planType: string; content: string; comment?: string }): Promise<LearningPlanResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/create', req)
    return response.data
  }

  async update(req: { id: number; content: string; comment?: string }): Promise<LearningPlanResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/edit', req)
    return response.data
  }

  async delete(id: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/delete', { id })
    return response.data
  }

  async getById(id: number): Promise<LearningPlanResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/byId', { id })
    return response.data
  }

  async list(query: { classId?: number; userId?: number; planType?: string; pageSize?: number; pageIndex?: number }): Promise<LearningPlanListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/list', query)
    return response.data
  }

  async versions(planId: number, page = 1, pageSize = 20): Promise<LearningPlanVersionListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/versions', { planId, pageIndex: page, pageSize })
    return response.data
  }

  async rollback(planId: number, version: number, comment: string): Promise<LearningPlanResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/rollback', { planId, version, comment })
    return response.data
  }

  async createPhasePlan(req: { planId: number; examNodeId: number; title: string; startDate?: string; endDate?: string; sortOrder?: number }): Promise<PhasePlanResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/phase/create', req)
    return response.data
  }

  async updatePhasePlan(req: { id: number; title?: string; startDate?: string; endDate?: string; sortOrder?: number }): Promise<PhasePlanResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/phase/edit', req)
    return response.data
  }

  async deletePhasePlan(id: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/phase/delete', { id })
    return response.data
  }

  async getPhasePlanById(id: number): Promise<PhasePlanResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/phase/byId', { id })
    return response.data
  }

  async listPhasePlans(planId: number): Promise<PhasePlanListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/phase/list', { planId })
    return response.data
  }

  async generateTemplate(req: {
    classId: number;
    syllabusId: number;
    startMonth: string;
    endMonth: string;
    phaseRatios: number[];
    comment?: string;
  }): Promise<ApiResponse<{ studentCount: number; count: number; errors?: string[] }>> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/generateTemplate', req)
    return response.data
  }

  async addChapterToPhasePlan(phasePlanId: number, chapterId: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/phase/chapter/add', { phasePlanId, chapterId })
    return response.data
  }

  async removeChapterFromPhasePlan(phasePlanId: number, chapterId: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/learning-plan/phase/chapter/remove', { phasePlanId, chapterId })
    return response.data
  }
}

export default new LearningPlanService()
