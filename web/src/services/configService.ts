import getApiClient from './apiClient';
import type { ApiResponse } from '../models/api.model';
import type { ImageUploadConfig } from '../models/config.model';

class ConfigService {
  private baseUrl = '/api/v1/syssetting';

  // Get image upload config
  async getImageUploadConfig(): Promise<ApiResponse<ImageUploadConfig>> {
    const client = await getApiClient();
    const response = await client.get(`${this.baseUrl}/imageUpload`);
    return response.data;
  }

  // Save image upload config
  async saveImageUploadConfig(config: ImageUploadConfig): Promise<ApiResponse<void>> {
    const client = await getApiClient();
    const response = await client.post(`${this.baseUrl}/imageUpload`, config);
    return response.data;
  }

  // Migrate base64 images to file storage
  async migrateImages(): Promise<ApiResponse<void>> {
    const client = await getApiClient();
    const response = await client.post(`${this.baseUrl}/image/migrate`);
    return response.data;
  }
}

export default new ConfigService();