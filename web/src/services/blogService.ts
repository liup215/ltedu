import apiClient from './apiClient'
import type {
  BlogPostCreateRequest,
  BlogPostUpdateRequest,
  BlogPostQuery,
  BlogListResponse,
  BlogPostResponse,
} from '../models/blog.model'

class BlogService {
  async listPosts(query: BlogPostQuery = {}): Promise<BlogListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/list', query)
    return response.data
  }

  async getPost(id: number): Promise<BlogPostResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/byId', { id })
    return response.data
  }

  async getPostBySlug(slug: string): Promise<BlogPostResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/bySlug', { slug })
    return response.data
  }

  async createPost(req: BlogPostCreateRequest): Promise<BlogPostResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/create', req)
    return response.data
  }

  async updatePost(req: BlogPostUpdateRequest): Promise<BlogPostResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/edit', req)
    return response.data
  }

  async deletePost(id: number): Promise<void> {
    const client = await apiClient()
    await client.post('/api/v1/blog/delete', { id })
  }

  // Admin methods
  async adminListPosts(query: BlogPostQuery = {}): Promise<BlogListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/admin/list', query)
    return response.data
  }

  async adminUpdatePost(req: BlogPostUpdateRequest): Promise<BlogPostResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/blog/admin/edit', req)
    return response.data
  }

  async adminDeletePost(id: number): Promise<void> {
    const client = await apiClient()
    await client.post('/api/v1/blog/admin/delete', { id })
  }
}

export default new BlogService()
