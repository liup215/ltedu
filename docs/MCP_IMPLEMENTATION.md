# MCP (Model Context Protocol) Implementation Guide

## Overview

This document describes the MCP functionality that has been added to the ltedu-api system. MCP allows AI assistants and other tools to interact with the system programmatically through a standardized protocol.

## Features

### 1. User-Specific MCP Tokens

Users can generate personal MCP tokens for authentication:
- Each token is unique and belongs to a specific user
- Tokens have configurable expiration dates
- Tokens can be activated/deactivated
- Last used timestamp is tracked

### 2. MCP Server

The MCP server implements the Model Context Protocol (JSON-RPC 2.0):
- **Endpoint**: `/api/mcp?token=<user_token>`
- **Transport**: Stateless HTTP (JSON-RPC 2.0 over HTTP POST)
- **Protocol Version**: 2024-11-05

Supported methods:
- `initialize` - Server capability negotiation
- `tools/list` - List all available tools
- `tools/call` - Execute a specific tool

### 3. Available MCP Tools

The system exposes 45 tools across 9 entities:

#### Organisation (考试局)
- `organisation_list` - List all exam organisations
- `organisation_get` - Get organisation by ID
- `organisation_create` - Create new organisation
- `organisation_edit` - Edit existing organisation
- `organisation_delete` - Delete organisation

#### Qualification (考试)
- `qualification_list` - List qualifications
- `qualification_get` - Get qualification by ID
- `qualification_create` - Create new qualification
- `qualification_edit` - Edit existing qualification
- `qualification_delete` - Delete qualification

#### Syllabus (考纲)
- `syllabus_list` - List syllabuses
- `syllabus_get` - Get syllabus by ID
- `syllabus_create` - Create new syllabus
- `syllabus_edit` - Edit existing syllabus
- `syllabus_delete` - Delete syllabus

#### Chapter (章节)
- `chapter_tree` - Get chapter tree for a syllabus
- `chapter_get` - Get chapter by ID
- `chapter_create` - Create new chapter
- `chapter_edit` - Edit existing chapter
- `chapter_delete` - Delete chapter

#### PaperSeries (考试季)
- `paper_series_list` - List paper series
- `paper_series_get` - Get paper series by ID
- `paper_series_create` - Create new paper series
- `paper_series_edit` - Edit existing paper series
- `paper_series_delete` - Delete paper series

#### PaperCode (试卷代码)
- `paper_code_list` - List paper codes
- `paper_code_get` - Get paper code by ID
- `paper_code_create` - Create new paper code
- `paper_code_edit` - Edit existing paper code
- `paper_code_delete` - Delete paper code

#### PastPaper (真题试卷)
- `past_paper_list` - List past papers
- `past_paper_get` - Get past paper by ID
- `past_paper_create` - Create new past paper
- `past_paper_edit` - Edit existing past paper
- `past_paper_delete` - Delete past paper

#### Question (试题)
- `question_list` - List questions
- `question_get` - Get question by ID
- `question_create` - Create new question
- `question_edit` - Edit existing question
- `question_delete` - Delete question

### 4. Token Management API

Authenticated users can manage their MCP tokens through these endpoints:

#### Create Token
```
POST /api/v1/mcp/token/create
Authorization: Bearer <jwt_token>

Body:
{
  "name": "My MCP Token",
  "expiresAt": "2025-12-31T23:59:59Z"  // Optional, defaults to 1 year
}

Response:
{
  "code": 0,
  "message": "MCP令牌创建成功",
  "data": {
    "id": 1,
    "token": "generated_token_string",
    "name": "My MCP Token",
    "expiresAt": "2025-12-31T23:59:59Z",
    "isActive": true,
    "createdAt": "2024-01-29T14:00:00Z",
    "lastUsed": "2024-01-29T14:00:00Z"
  }
}
```

#### List Tokens
```
POST /api/v1/mcp/token/list
Authorization: Bearer <jwt_token>

Body:
{
  "pageIndex": 1,
  "pageSize": 20
}

Response:
{
  "code": 0,
  "message": "获取MCP令牌列表成功",
  "data": {
    "total": 5,
    "records": [...]
  }
}
```

#### Delete Token
```
POST /api/v1/mcp/token/delete
Authorization: Bearer <jwt_token>

Body:
{
  "id": 1
}
```

#### Deactivate Token
```
POST /api/v1/mcp/token/deactivate
Authorization: Bearer <jwt_token>

Body:
{
  "id": 1
}
```

## Usage Example

### 1. Generate an MCP Token

First, log in and create an MCP token through the API:

```bash
curl -X POST http://localhost:9000/api/v1/mcp/token/create \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "My Assistant Token"}'
```

### 2. Connect via MCP (Initialize)

Use the generated token to connect to the MCP server:

```bash
curl -X POST "http://localhost:9000/api/mcp?token=YOUR_MCP_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"initialize"}'
```

### 3. List Available Tools

```bash
curl -X POST "http://localhost:9000/api/mcp?token=YOUR_MCP_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":2,"method":"tools/list"}'
```

### 4. Call a Tool

```bash
curl -X POST "http://localhost:9000/api/mcp?token=YOUR_MCP_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc":"2.0",
    "id":3,
    "method":"tools/call",
    "params":{
      "name":"organisation_list",
      "arguments":{"pageIndex":1,"pageSize":10}
    }
  }'
```

## Security

- All MCP operations are authenticated using user-specific tokens
- Tokens can be revoked at any time by the user
- Tokens have expiration dates
- All operations respect the same user permissions as the regular API
- No security vulnerabilities detected by CodeQL analysis

## Database Schema

The system adds one new table:

### mcp_tokens
- `id` (uint) - Primary key
- `user_id` (uint) - Foreign key to users table
- `token` (string) - Unique token string (64 characters)
- `name` (string) - Friendly name for the token
- `expires_at` (timestamp) - Token expiration date
- `last_used` (timestamp) - Last time the token was used
- `is_active` (boolean) - Whether the token is active
- `created_at` (timestamp)
- `updated_at` (timestamp)
- `deleted_at` (timestamp) - Soft delete

## Technical Implementation

### Files Added/Modified

**New Files:**
- `model/mcp.go` - MCP token data model
- `repository/mcp_token_repository.go` - MCP token data access
- `service/mcp_token.go` - MCP token business logic
- `server/mcp/mcp.go` - MCP server protocol handler
- `server/mcp/tools.go` - MCP tool implementations
- `server/api/v1/mcp_token.go` - MCP token API controller
- `docs/MCP_IMPLEMENTATION.md` - This documentation

**Modified Files:**
- `server/api/controller.go` - Added MCP routes
- `service/service.go` - Added MCPToken auto-migration
- `repository/repository.go` - Added MCPTokenRepo initialization

## Permissions

All MCP tools respect the same permissions as the regular API:
- Users can only access data they have permission to view
- Create/Edit/Delete operations require appropriate permissions
- The MCP token inherits all permissions from the associated user account

## Future Enhancements

Potential future improvements:
- Add resource-level tools (prompts, resources)
- Implement sampling capability for LLM interactions
- Add webhook support for asynchronous operations
- Implement rate limiting per token
- Add audit logging for MCP operations
- Support for custom tool permissions per token
