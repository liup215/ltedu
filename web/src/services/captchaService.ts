import apiClient from './apiClient';
import type { CaptchaImageResponse } from '../models/captcha.model';

const captchaService = {
  async getCaptchaImage(): Promise<CaptchaImageResponse> {
    const client = await apiClient();
    const response = await client.post('/api/v1/captcha');
    return response.data;
  }
};

export default captchaService;
