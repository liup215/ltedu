# HTTP Rules for LTEDU Project

## Overview

This document defines the HTTP communication standards and patterns used in the LTEDU platform. All HTTP interactions should follow these established conventions for consistency, maintainability, and reliability.

## API Architecture

### Unified API Structure

The LTEDU platform uses a unified API architecture:

- **Base URL Pattern**: `{domain}/api/v1/`
- **Environment Configuration**: Base URLs exclude `/api` prefix (handled in service layer)
- **Versioning**: All APIs use `/v1/` for version control
- **Role-Based Access**: Single API serves all user roles (Admin, Teacher, Student)

### Environment Configuration

```bash
# Development
VITE_API_BASE_URL=http://localhost:9000

# Production  
VITE_API_BASE_URL=https://api.ltedu.com
```

## HTTP Methods and Endpoints

### Standard CRUD Operations

All resource endpoints follow consistent CRUD patterns:

```typescript
// Service pattern example
class ResourceService {
  private baseUrl = '/api/v1/resource';

  // LIST - Get paginated resources
  async getResources(query: ResourceQuery = {}): Promise<PaginatedResources> {
    const response = await apiClient.post(`${this.baseUrl}/list`, query);
    return response.data;
  }

  // GET BY ID - Get single resource
  async getResourceById(id: number): Promise<Resource> {
    const response = await apiClient.post(`${this.baseUrl}/byId`, { id });
    return response.data;
  }

  // GET ALL - Get all resources (for dropdowns)
  async getAllResources(query: ResourceQuery = {}): Promise<Resource[]> {
    const response = await apiClient.post(`${this.baseUrl}/all`, query);
    return response.data;
  }

  // CREATE - Create new resource
  async createResource(resource: ResourceCreateRequest): Promise<Resource> {
    const response = await apiClient.post(`${this.baseUrl}/create`, resource);
    return response.data;
  }

  // UPDATE - Update existing resource
  async updateResource(resource: ResourceUpdateRequest): Promise<void> {
    await apiClient.post(`${this.baseUrl}/edit`, resource);
  }

  // DELETE - Delete resource
  async deleteResource(id: number): Promise<void> {
    await apiClient.post(`${this.baseUrl}/delete`, { id });
  }
}
```

### HTTP Method Usage

- **POST**: Used for ALL API operations (list, create, update, delete, single retrieval)
- **GET**: Reserved for static file downloads and special cases
- **Request Body**: All data sent via JSON request body, not URL parameters
- **Consistency**: All endpoints use POST for uniform request/response handling

## Request/Response Patterns

### Standard Request Format

#### Pagination Requests
```json
{
  "pageSize": 20,
  "pageIndex": 1,
  "searchField": "value",
  "filterField": "value"
}
```

#### ID-based Requests
```json
{
  "id": 123
}
```

#### Create/Update Requests
```json
{
  "field1": "value1",
  "field2": "value2",
  "relationId": 456
}
```

### Standard Response Format

#### Success Response
```json
{
  "code": 0,
  "message": "操作成功消息",
  "data": {
    // Response data specific to the endpoint
  }
}
```

#### Error Response
```json
{
  "code": 1,
  "message": "错误消息",
  "data": null
}
```

#### Paginated Response
```json
{
  "code": 0,
  "message": "数据获取成功!",
  "data": {
    "list": [...],
    "total": 100
  }
}
```

## Authentication and Authorization

### Authentication Flow

1. **Login Request**:
   ```json
   POST /api/v1/login
   {
     "username": "string",
     "password": "string"
   }
   ```

2. **Login Response**:
   ```json
   {
     "code": 0,
     "message": "登录成功！",
     "data": {
       "token": "jwt_token_string",
       "expire": "2025-05-19T10:17:30Z"
     }
   }
   ```

3. **Subsequent Requests**:
   ```http
   Authorization: Bearer {jwt_token}
   Content-Type: application/json
   ```

### Role-Based Access Control

- **Admin**: Full access to all `/api/v1/*` endpoints
- **Teacher**: Limited access based on middleware permissions
- **Student**: Read-only access to relevant data
- **Public**: Access to login, register, and setup endpoints only

## API Client Configuration

### Base Client Setup

```typescript
import axios from 'axios';

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor for authentication
apiClient.interceptors.request.use((config) => {
  const token = localStorage.getItem('authToken');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Response interceptor for error handling
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Handle authentication errors
      localStorage.removeItem('authToken');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);
```

## Error Handling Standards

### Custom Response Codes

The LTEDU platform uses a custom response code system in the JSON response body, separate from HTTP status codes:

- **0**: Success - Operation completed successfully
- **1**: General Error - Operation failed (default error code)
- **Other positive integers**: Specific error codes for different error types

**Important**: These response codes are different from HTTP status codes. The HTTP status code will typically be 200 for successful requests, while the JSON response `code` field indicates the actual operation result.

### HTTP Status Codes

Standard HTTP status codes are still used for transport-level issues:

- **200**: Success - HTTP request completed successfully
- **400**: Bad Request - Invalid request data or parameters
- **401**: Unauthorized - Authentication required or token invalid
- **403**: Forbidden - User lacks permission for the operation
- **404**: Not Found - Resource does not exist
- **500**: Internal Server Error - Server-side error occurred

### Error Response Handling

```typescript
// Service method error handling pattern
async serviceMethod(): Promise<ResourceType> {
  try {
    const response = await apiClient.post(endpoint, data);
    return response.data;
  } catch (error) {
    console.error('Service operation failed:', error);
    
    if (error.response) {
      // Server responded with error status
      throw new Error(error.response.data.message || 'Operation failed');
    } else if (error.request) {
      // Network error
      throw new Error('Network error - please check your connection');
    } else {
      // Other error
      throw new Error('An unexpected error occurred');
    }
  }
}
```

## Data Validation Rules

### Request Validation

1. **Required Fields**: All required fields must be validated client-side
2. **Data Types**: Ensure correct TypeScript types for all request data
3. **String Validation**: Trim whitespace, validate length constraints
4. **Number Validation**: Validate numeric ranges and formats
5. **Relationship Validation**: Verify foreign key relationships exist

### Response Validation

1. **Type Safety**: Use TypeScript interfaces for all response data
2. **Null Checks**: Handle potential null/undefined values gracefully
3. **Array Validation**: Check array existence before operations
4. **Data Consistency**: Validate relationships and data integrity

## Performance Optimization

### Request Optimization

1. **Pagination**: Always use pagination for list endpoints
2. **Selective Loading**: Request only necessary data fields
3. **Caching**: Implement appropriate caching strategies
4. **Debouncing**: Debounce search inputs to reduce API calls

### Response Optimization

1. **Efficient Queries**: Backend should optimize database queries
2. **Data Compression**: Use gzip compression for responses
3. **Minimal Payloads**: Return only necessary data in responses
4. **Lazy Loading**: Implement lazy loading for large datasets

## Security Guidelines

### Data Protection

1. **Sensitive Data**: Never log sensitive information
2. **Input Sanitization**: Sanitize all user inputs
3. **HTTPS Only**: All production traffic must use HTTPS
4. **Token Security**: Store JWT tokens securely

### API Security

1. **Authentication**: All protected endpoints require valid JWT
2. **Authorization**: Implement role-based access control
3. **Rate Limiting**: Implement rate limiting on API endpoints
4. **CORS**: Configure CORS policies appropriately

## Testing Standards

### Unit Testing

```typescript
// Service testing pattern
describe('ResourceService', () => {
  it('should fetch resources with pagination', async () => {
    const mockQuery = { pageSize: 10, pageIndex: 1 };
    const mockResponse = { data: { list: [], total: 0 } };
    
    jest.spyOn(apiClient, 'post').mockResolvedValue(mockResponse);
    
    const result = await resourceService.getResources(mockQuery);
    
    expect(apiClient.post).toHaveBeenCalledWith('/api/v1/resource/list', mockQuery);
    expect(result).toEqual(mockResponse.data);
  });
});
```

### Integration Testing

1. **API Contract Testing**: Verify request/response contracts
2. **Authentication Testing**: Test auth flows and token handling
3. **Error Scenario Testing**: Test error handling and edge cases
4. **Performance Testing**: Test response times and load handling

## Documentation Standards

### API Documentation

1. **Endpoint Documentation**: Document all endpoints with examples
2. **Request/Response Schemas**: Provide TypeScript interfaces
3. **Error Codes**: Document all possible error responses
4. **Authentication**: Clear auth requirements for each endpoint

### Code Documentation

1. **Service Methods**: Document all service methods
2. **Type Definitions**: Comprehensive TypeScript interfaces
3. **Error Handling**: Document error handling strategies
4. **Usage Examples**: Provide usage examples for complex operations

## Migration and Versioning

### API Versioning

1. **Version Strategy**: Use `/v1/`, `/v2/` URL versioning
2. **Backward Compatibility**: Maintain compatibility during transitions
3. **Deprecation Policy**: Clear timeline for deprecating old versions
4. **Migration Guide**: Provide migration guides for API changes

### Breaking Changes

1. **Change Management**: Document all breaking changes
2. **Transition Period**: Provide transition period for breaking changes
3. **Client Updates**: Coordinate client updates with API changes
4. **Rollback Strategy**: Have rollback plans for failed deployments

## Monitoring and Logging

### Request Logging

1. **Request Tracking**: Log all API requests with unique identifiers
2. **Performance Metrics**: Track response times and error rates
3. **Error Monitoring**: Monitor and alert on API errors
4. **Usage Analytics**: Track API usage patterns and trends

### Health Monitoring

1. **Health Checks**: Implement API health check endpoints
2. **Uptime Monitoring**: Monitor API availability
3. **Alert Systems**: Set up alerts for critical issues
4. **Performance Dashboards**: Create dashboards for API metrics

## Examples and Best Practices

### Service Implementation Example

```typescript
import apiClient from './apiClient';
import type { 
  Qualification, 
  QualificationQuery, 
  PaginatedQualifications,
  QualificationCreateRequest,
  QualificationUpdateRequest 
} from '../models/qualification.model';

class QualificationService {
  private baseUrl = '/api/v1/qualification';

  async getQualifications(query: QualificationQuery = {}): Promise<PaginatedQualifications> {
    const response = await apiClient.post(`${this.baseUrl}/list`, query);
    return response.data;
  }

  async getQualificationById(id: number): Promise<Qualification> {
    const response = await apiClient.post(`${this.baseUrl}/byId`, { id });
    return response.data;
  }

  async getAllQualifications(query: QualificationQuery = {}): Promise<Qualification[]> {
    const response = await apiClient.post(`${this.baseUrl}/all`, query);
    return response.data;
  }

  async createQualification(qualification: QualificationCreateRequest): Promise<Qualification> {
    const response = await apiClient.post(`${this.baseUrl}/create`, qualification);
    return response.data;
  }

  async updateQualification(qualification: QualificationUpdateRequest): Promise<void> {
    await apiClient.post(`${this.baseUrl}/edit`, qualification);
  }

  async deleteQualification(id: number): Promise<void> {
    await apiClient.post(`${this.baseUrl}/delete`, { id });
  }
}

export default new QualificationService();
```

### Component Usage Example

```vue
<script setup lang="ts">
import { ref, onMounted } from 'vue';
import qualificationService from '../../services/qualificationService';
import type { Qualification, QualificationQuery } from '../../models/qualification.model';

const qualifications = ref<Qualification[]>([]);
const loading = ref(false);
const query = ref<QualificationQuery>({
  pageSize: 20,
  pageIndex: 1,
  name: '',
  organisationId: ''
});

const fetchQualifications = async () => {
  loading.value = true;
  try {
    const response = await qualificationService.getQualifications(query.value);
    qualifications.value = response.list;
  } catch (error) {
    console.error('Failed to fetch qualifications:', error);
    // Handle error appropriately
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchQualifications();
});
</script>
```

This HTTP rules document should be followed by all developers working on the LTEDU project to ensure consistency, reliability, and maintainability of the HTTP communication layer.

## 运行时环境变量读取方式

前端所有环境变量（如 API_BASE_URL、APP_TITLE）在运行时通过 `window.__LTEDU_CONFIG__` 读取。例如：

```typescript
const apiBaseUrl = (window as any).__LTEDU_CONFIG__?.API_BASE_URL;
const appTitle = (window as any).__LTEDU_CONFIG__?.APP_TITLE;
```

如需动态切换环境，无需重建镜像，只需修改 config.json 或通过容器环境变量生成 config.json。
