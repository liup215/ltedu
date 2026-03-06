import apiService from './ApiService';
import type { ApiResponse, LearningPlan, LearningPlanListResponse } from '../models/ApiModels';

/**
 * LearningPlanService provides access to user learning plans, supporting
 * list, detail, create, and update operations.
 */
export class LearningPlanService {
  private static instance: LearningPlanService;

  static getInstance(): LearningPlanService {
    if (!LearningPlanService.instance) {
      LearningPlanService.instance = new LearningPlanService();
    }
    return LearningPlanService.instance;
  }

  async listMyPlans(request: {
    planType?: string;
    pageIndex?: number;
    pageSize?: number;
  }): Promise<ApiResponse<LearningPlanListResponse>> {
    return apiService.post<LearningPlanListResponse>('/api/v1/learning-plan/list', {
      isPersonal: true,
      ...request,
    });
  }

  async getPlanById(id: number): Promise<ApiResponse<LearningPlan>> {
    return apiService.post<LearningPlan>('/api/v1/learning-plan/byId', { id });
  }

  async createPlan(request: {
    planType: string;
    content: string;
    comment?: string;
  }): Promise<ApiResponse<LearningPlan>> {
    return apiService.post<LearningPlan>('/api/v1/learning-plan/create', {
      isPersonal: true,
      ...request,
    });
  }

  async updatePlan(request: {
    id: number;
    content: string;
    comment?: string;
  }): Promise<ApiResponse<LearningPlan>> {
    return apiService.post<LearningPlan>('/api/v1/learning-plan/edit', request);
  }
}

export default LearningPlanService.getInstance();
