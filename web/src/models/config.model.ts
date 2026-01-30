import type { ApiResponse } from './api.model';

// Configuration keys
export const CONFIG_KEY_IMAGE_UPLOAD = 'ltedu.upload.image.upload';

// Disk types
export const DISK_PUBLIC = 'public';
export const DISK_OSS = 'oss';
export const DISK_COS = 'cos';
export const DISK_QINIU = 'qiniu';

export interface ImageUploadConfig {
  disk: string;
  ossAccessKeyId?: string;
  ossAccessKeySecret?: string;
  ossBucket?: string;
  ossEndpoint?: string;
  ossCDNUrl?: string;

  cosRegion?: string;
  cosAppId?: string;
  cosSecretId?: string;
  cosSecretKey?: string;
  cosBucket?: string;
  cosCDNUrl?: string;

  qiniuAccessKey?: string;
  qiniuSecretKey?: string;
  qiniuBucket?: string;
  qiniuCDNUrl?: string;
}

export type ImageUploadConfigResponse = ApiResponse<ImageUploadConfig>;