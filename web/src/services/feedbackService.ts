import apiClient from './apiClient'
import type {
  SubmitFeedbackRequest,
  FeedbackListRequest,
  FeedbackSubmitResponse,
  FeedbackListResponse,
  FeedbackStatsResponse,
  UserFeedback
} from '../models/feedback.model'
import type { ApiResponse, Page } from '../models/api.model'

class FeedbackService {
  async submit(req: SubmitFeedbackRequest): Promise<FeedbackSubmitResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/feedback/submit', {
      ...req,
      userAgent: navigator.userAgent
    })
    return response.data
  }

  async myFeedback(page: Page = { pageSize: 20, pageIndex: 1 }): Promise<FeedbackListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/feedback/my', page)
    return response.data
  }

  async list(req: FeedbackListRequest = { pageSize: 20, pageIndex: 1 }): Promise<FeedbackListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/feedback/list', req)
    return response.data
  }

  async getById(id: number): Promise<ApiResponse<UserFeedback>> {
    const client = await apiClient()
    const response = await client.post('/api/v1/feedback/byId', { id })
    return response.data
  }

  async updateStatus(id: number, status: string, adminNote = ''): Promise<ApiResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/feedback/updateStatus', { id, status, adminNote })
    return response.data
  }

  async getStats(): Promise<FeedbackStatsResponse> {
    const client = await apiClient()
    const response = await client.get('/api/v1/feedback/stats')
    return response.data
  }
}

export default new FeedbackService()
