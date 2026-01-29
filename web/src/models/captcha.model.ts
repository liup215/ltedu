import type { ApiResponse } from './api.model';

export interface CaptchaImage {
  img: string;
  key: string;
}

export type CaptchaImageResponse = ApiResponse<CaptchaImage>;
