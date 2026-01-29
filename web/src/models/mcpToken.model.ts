/**
 * MCP Token Model Definitions
 */

// MCP Token response from API
export interface MCPToken {
  id: number;
  token: string;
  name: string;
  expiresAt: string;
  isActive: boolean;
  createdAt: string;
  lastUsed: string;
  userId?: number; // Available in admin responses
}

// Request to create a new MCP token
export interface MCPTokenCreateRequest {
  name: string;
  expiresAt?: string; // ISO 8601 format, optional
}

// Paginated MCP tokens response
export interface PaginatedMCPTokens {
  list: MCPToken[];
  total: number;
}

// Query criteria for listing MCP tokens
export interface MCPTokenQueryCriteria {
  pageIndex: number;
  pageSize: number;
  userId?: number; // For admin use only
}
