import apiClient from './apiClient';
import type { 
  Qualification, 
  QualificationQuery, 
  PaginatedQualifications,
  QualificationCreateRequest,
  QualificationUpdateRequest 
} from '../models/qualification.model';
import type { ApiResponse } from '../models/api.model';

class QualificationService {
  private baseUrl = '/api/v1/qualification';

  async getQualifications(query: QualificationQuery = {}): Promise<ApiResponse<PaginatedQualifications>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post(`${this.baseUrl}/list`, query);
    return response.data;
  }

  async getQualificationById(id: number): Promise<ApiResponse<Qualification>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post(`${this.baseUrl}/byId`, { id });
    return response.data;
  }

  async getAllQualifications(query: QualificationQuery = {}): Promise<ApiResponse<PaginatedQualifications>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post(`${this.baseUrl}/all`, query);
    return response.data;
  }

  async createQualification(qualification: QualificationCreateRequest): Promise<ApiResponse<Qualification>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post(`${this.baseUrl}/create`, qualification);
    return response.data;
  }

  async updateQualification(qualification: QualificationUpdateRequest): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post(`${this.baseUrl}/edit`, qualification);
    return response.data;
  }

  async deleteQualification(id: number): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post(`${this.baseUrl}/delete`, { id });
    return response.data;
  }
}

export default new QualificationService();
