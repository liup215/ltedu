import http from '@ohos.net.http';
import type { ApiResponse } from '../models/ApiModels';
import storageService from './StorageService';

/**
 * ApiService provides HTTP request utilities for communicating with the
 * LTEdu backend. All requests include the JWT token from StorageService,
 * and responses are parsed according to the ApiResponse<T> format where
 * code=0 indicates success and code=1 indicates an error.
 */
export class ApiService {
  private static instance: ApiService;

  static getInstance(): ApiService {
    if (!ApiService.instance) {
      ApiService.instance = new ApiService();
    }
    return ApiService.instance;
  }

  private async getBaseUrl(): Promise<string> {
    return storageService.getApiBaseUrl();
  }

  private async getHeaders(): Promise<Record<string, string>> {
    const token = await storageService.getToken();
    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    };
    if (token) {
      headers['Authorization'] = `Bearer ${token}`;
    }
    return headers;
  }

  async post<T>(path: string, body: object): Promise<ApiResponse<T>> {
    const baseUrl = await this.getBaseUrl();
    const headers = await this.getHeaders();
    const url = `${baseUrl}${path}`;

    const httpRequest = http.createHttp();
    try {
      const response = await httpRequest.request(url, {
        method: http.RequestMethod.POST,
        header: headers,
        extraData: JSON.stringify(body),
        connectTimeout: 10000,
        readTimeout: 30000,
      });
      const result: ApiResponse<T> = JSON.parse(response.result as string);
      return result;
    } finally {
      httpRequest.destroy();
    }
  }

  async get<T>(path: string): Promise<ApiResponse<T>> {
    const baseUrl = await this.getBaseUrl();
    const headers = await this.getHeaders();
    const url = `${baseUrl}${path}`;

    const httpRequest = http.createHttp();
    try {
      const response = await httpRequest.request(url, {
        method: http.RequestMethod.GET,
        header: headers,
        connectTimeout: 10000,
        readTimeout: 30000,
      });
      const result: ApiResponse<T> = JSON.parse(response.result as string);
      return result;
    } finally {
      httpRequest.destroy();
    }
  }
}

export default ApiService.getInstance();
