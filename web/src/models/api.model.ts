// Common API response type for all backend responses
export interface ApiResponse<T = any> {
  code: number;
  message: string;
  data: T;
}

// Common pagination parameters
export interface Page {
  pageSize?: number;
  pageIndex?: number;
}
