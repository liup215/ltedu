import apiService from './ApiService';
import storageService from './StorageService';
import type { ApiResponse, LoginCredentials, LoginResponseData, User } from '../models/ApiModels';

/**
 * AuthService handles user authentication: login, registration, and logout.
 * Persists the JWT token and user info via StorageService upon successful login.
 */
export class AuthService {
  private static instance: AuthService;

  static getInstance(): AuthService {
    if (!AuthService.instance) {
      AuthService.instance = new AuthService();
    }
    return AuthService.instance;
  }

  async login(credentials: LoginCredentials): Promise<ApiResponse<LoginResponseData>> {
    const response = await apiService.post<LoginResponseData>('/api/v1/login', credentials);
    if (response.code === 0 && response.data?.token) {
      await storageService.setToken(response.data.token);
      await storageService.setUser(JSON.stringify(response.data.user));
    }
    return response;
  }

  async register(payload: {
    username: string;
    password: string;
    email: string;
    code: string;
  }): Promise<ApiResponse<string>> {
    return apiService.post<string>('/api/v1/verification/register-with-code', payload);
  }

  async sendVerificationCode(payload: {
    email: string;
    captchaId: string;
    captchaValue: string;
  }): Promise<ApiResponse<string>> {
    return apiService.post<string>('/api/v1/verification/send-code', payload);
  }

  async getCurrentUser(): Promise<ApiResponse<User>> {
    return apiService.get<User>('/api/v1/user');
  }

  async logout(): Promise<void> {
    await storageService.clearToken();
  }

  async isLoggedIn(): Promise<boolean> {
    return storageService.isLoggedIn();
  }

  async getCachedUser(): Promise<User | null> {
    const userJson = await storageService.getUser();
    if (!userJson) return null;
    try {
      return JSON.parse(userJson) as User;
    } catch {
      return null;
    }
  }
}

export default AuthService.getInstance();
