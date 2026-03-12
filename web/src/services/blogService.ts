import apiClient from './apiClient'
import type { ApiResponse } from '../models/api.model'
import type { BlogPostCreateRequest, BlogPostQuery, BlogListResponse, BlogDetailResponse } from '../models/blog.model'

class BlogService {
  // Public endpoints (no auth required)
  async getPublishedPosts(query: BlogPostQuery = { pageSize: 10, pageIndex: 1 }): Promise<BlogListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/public/list', query)
    return response.data
  }

  async getPostBySlug(slug: string): Promise<BlogDetailResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/public/bySlug', { slug })
    return response.data
  }

  // Admin endpoints (require admin auth)
  async createPost(post: BlogPostCreateRequest): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/create', post)
    return response.data
  }

  async updatePost(post: BlogPostCreateRequest): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/edit', post)
    return response.data
  }

  async deletePost(id: number): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/delete', { id })
    return response.data
  }

  async getPostById(id: number): Promise<BlogDetailResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/byId', { id })
    return response.data
  }

  async listPosts(query: BlogPostQuery = { pageSize: 20, pageIndex: 1 }): Promise<BlogListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/list', query)
    return response.data
  }
}

export default new BlogService()
