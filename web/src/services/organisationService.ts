import apiClient from './apiClient';
import type { Organisation, OrganisationQuery, PaginatedOrganisations } from '../models/organisation.model';
import type { ApiResponse } from '../models/api.model';

class OrganisationService {
  async getOrganisations(query: OrganisationQuery): Promise<ApiResponse<PaginatedOrganisations>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/organisation/list', query);
    return response.data;
  }

  async getOrganisationById(id: number): Promise<ApiResponse<Organisation>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/organisation/byId', { id });
    return response.data;
  }

  async getAllOrganisations(query: OrganisationQuery): Promise<ApiResponse<PaginatedOrganisations>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/organisation/all', query);
    return response.data as ApiResponse<PaginatedOrganisations>;
  }

  async createOrganisation(organisation: Omit<Organisation, 'id'>): Promise<ApiResponse<Organisation>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/organisation/create', organisation);
    return response.data;
  }

  async updateOrganisation(organisation: Organisation): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/organisation/edit', organisation);
    return response.data;
  }

  async deleteOrganisation(id: number): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/organisation/delete', { id });
    return response.data;
  }
}

export default new OrganisationService();
