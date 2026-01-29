import apiClient from './apiClient';
import type { ApiResponse } from '../models/api.model'; // Updated import path
// Import models from their new locations
import type { User } from '../models/user.model';
import type { LoginCredentials, LoginResponseData, RegistrationPayload, ChangePasswordRequest, ChangePasswordResponse } from '../models/auth.model';



// Local type definitions for PermissionData and RoleData are removed as they are in user.model.ts

export const authService = {
  async login(credentials: LoginCredentials): Promise<ApiResponse<LoginResponseData>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/login', credentials);
    return response.data;
  },

  async getCurrentUserProfile(): Promise<ApiResponse<User>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.get('/api/v1/user');
    return response.data;
  },

  async register(payload: RegistrationPayload): Promise<ApiResponse<string>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post('/api/v1/register', payload);
    return response.data;
  },

  async sendVerificationCode(payload: { email: string; captchaId: string; captchaValue: string }): Promise<ApiResponse<string>> {
    const client = await apiClient();
    const response = await client.post('/api/v1/verification/send-code', payload);
    return response.data;
  },

  async registerWithCode(payload: { username: string; password: string; email: string; code: string }): Promise<ApiResponse<string>> {
    const client = await apiClient();
    const response = await client.post('/api/v1/verification/register-with-code', payload);
    return response.data;
  },

  async changePassword(payload: ChangePasswordRequest): Promise<ApiResponse<ChangePasswordResponse>> {
    const client = await apiClient();
    const response = await client.post('/api/v1/change-password', payload);
    return response.data;
  }
  // ... (other methods like logout, if any, would go here)
};
