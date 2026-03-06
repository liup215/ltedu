import { invoke } from '@tauri-apps/api/core'

export interface ApiOptions {
  method?: string
  path: string
  body?: unknown
  query?: Record<string, string>
}

export interface ApiResult {
  status: number
  body: unknown
}

/**
 * Send an authenticated request to the LTEdu cloud API via the Rust backend.
 * The token and api_base_url are managed by the auth Tauri command.
 */
export async function apiRequest(options: ApiOptions): Promise<ApiResult> {
  return invoke<ApiResult>('cmd_api_request', {
    request: {
      method: options.method ?? 'GET',
      path: options.path,
      body: options.body ?? null,
      query: options.query ?? null,
    },
  })
}

export async function cacheGet(key: string): Promise<string | null> {
  return invoke<string | null>('cmd_cache_get', { key })
}

export async function cacheSet(key: string, value: string): Promise<void> {
  return invoke('cmd_cache_set', { key, value })
}

export async function cacheDelete(key: string): Promise<void> {
  return invoke('cmd_cache_delete', { key })
}

export async function cacheClear(): Promise<void> {
  return invoke('cmd_cache_clear')
}
