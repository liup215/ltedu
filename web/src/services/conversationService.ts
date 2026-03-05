import apiClient from './apiClient'
import type {
  ConversationStartRequest,
  ConversationMessageRequest,
  ConversationStartResponse,
  ConversationMessageResponse,
  ConversationHistoryResponse,
  ConversationSessionsResponse,
} from '../models/conversation.model'

class ConversationService {
  async startSession(req: ConversationStartRequest = {}): Promise<ConversationStartResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/ai/conversation/start', req)
    return response.data
  }

  async sendMessage(req: ConversationMessageRequest): Promise<ConversationMessageResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/ai/conversation/message', req)
    return response.data
  }

  async getHistory(
    sessionId: number,
    pageIndex = 1,
    pageSize = 50,
  ): Promise<ConversationHistoryResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/ai/conversation/history', {
      sessionId,
      pageIndex,
      pageSize,
    })
    return response.data
  }

  async listSessions(pageIndex = 1, pageSize = 20): Promise<ConversationSessionsResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/ai/conversation/sessions', {
      pageIndex,
      pageSize,
    })
    return response.data
  }

  async resetSession(sessionId: number): Promise<void> {
    const client = await apiClient()
    await client.post('/api/v1/ai/conversation/reset', { sessionId })
  }

  async closeSession(sessionId: number): Promise<void> {
    const client = await apiClient()
    await client.post('/api/v1/ai/conversation/close', { sessionId })
  }
}

export default new ConversationService()
