import apiClient from './apiClient';
import type { ApiResponse } from '../models/api.model';
import type {
  MCPToken,
  MCPTokenCreateRequest,
  PaginatedMCPTokens,
  MCPTokenQueryCriteria
} from '../models/mcpToken.model';

const MCP_TOKEN_API_BASE_PATH = '/api/v1/mcp/token';

/**
 * Create a new MCP token for the current user
 * Corresponds to: POST /api/v1/mcp/token/create
 */
async function createToken(data: MCPTokenCreateRequest): Promise<ApiResponse<MCPToken>> {
  const client = await apiClient();
  const axiosResponse = await client.post(`${MCP_TOKEN_API_BASE_PATH}/create`, data);
  return axiosResponse.data;
}

/**
 * List MCP tokens for the current user
 * Corresponds to: POST /api/v1/mcp/token/list
 */
async function listTokens(query: MCPTokenQueryCriteria): Promise<ApiResponse<PaginatedMCPTokens>> {
  const client = await apiClient();
  const axiosResponse = await client.post(`${MCP_TOKEN_API_BASE_PATH}/list`, query);
  return axiosResponse.data;
}

/**
 * Delete an MCP token
 * Corresponds to: POST /api/v1/mcp/token/delete
 */
async function deleteToken(id: number): Promise<ApiResponse<void>> {
  const client = await apiClient();
  const axiosResponse = await client.post(`${MCP_TOKEN_API_BASE_PATH}/delete`, { id });
  return axiosResponse.data;
}

/**
 * Deactivate an MCP token
 * Corresponds to: POST /api/v1/mcp/token/deactivate
 */
async function deactivateToken(id: number): Promise<ApiResponse<void>> {
  const client = await apiClient();
  const axiosResponse = await client.post(`${MCP_TOKEN_API_BASE_PATH}/deactivate`, { id });
  return axiosResponse.data;
}

/**
 * Activate an MCP token
 * Corresponds to: POST /api/v1/mcp/token/activate
 */
async function activateToken(id: number): Promise<ApiResponse<void>> {
  const client = await apiClient();
  const axiosResponse = await client.post(`${MCP_TOKEN_API_BASE_PATH}/activate`, { id });
  return axiosResponse.data;
}

/**
 * List all MCP tokens (admin only)
 * Corresponds to: POST /api/v1/mcp/token/admin/list
 */
async function listAllTokens(query: MCPTokenQueryCriteria): Promise<ApiResponse<PaginatedMCPTokens>> {
  const client = await apiClient();
  const axiosResponse = await client.post(`${MCP_TOKEN_API_BASE_PATH}/admin/list`, query);
  return axiosResponse.data;
}

/**
 * Delete any user's MCP token (admin only)
 * Corresponds to: POST /api/v1/mcp/token/admin/delete
 */
async function adminDeleteToken(id: number): Promise<ApiResponse<void>> {
  const client = await apiClient();
  const axiosResponse = await client.post(`${MCP_TOKEN_API_BASE_PATH}/admin/delete`, { id });
  return axiosResponse.data;
}

/**
 * Deactivate any user's MCP token (admin only)
 * Corresponds to: POST /api/v1/mcp/token/admin/deactivate
 */
async function adminDeactivateToken(id: number): Promise<ApiResponse<void>> {
  const client = await apiClient();
  const axiosResponse = await client.post(`${MCP_TOKEN_API_BASE_PATH}/admin/deactivate`, { id });
  return axiosResponse.data;
}

/**
 * Activate any user's MCP token (admin only)
 * Corresponds to: POST /api/v1/mcp/token/admin/activate
 */
async function adminActivateToken(id: number): Promise<ApiResponse<void>> {
  const client = await apiClient();
  const axiosResponse = await client.post(`${MCP_TOKEN_API_BASE_PATH}/admin/activate`, { id });
  return axiosResponse.data;
}

export default {
  createToken,
  listTokens,
  deleteToken,
  deactivateToken,
  activateToken,
  listAllTokens,
  adminDeleteToken,
  adminDeactivateToken,
  adminActivateToken
};
