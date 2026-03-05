export interface ConversationSession {
  id: number
  userId: number
  title: string
  isActive: boolean
  messageCount: number
  createdAt: string
  updatedAt: string
}

export interface ConversationMessage {
  id: number
  sessionId: number
  role: 'user' | 'assistant'
  content: string
  createdAt: string
}

export interface ConversationStartRequest {
  title?: string
}

export interface ConversationMessageRequest {
  sessionId: number
  content: string
}

export interface ConversationHistoryRequest {
  sessionId: number
  pageIndex?: number
  pageSize?: number
}

export interface ConversationSessionsRequest {
  pageIndex?: number
  pageSize?: number
}

export interface ConversationStartResponse {
  code: number
  msg: string
  data: ConversationSession
}

export interface ConversationMessageResponse {
  code: number
  msg: string
  data: ConversationMessage
}

export interface ConversationHistoryResponse {
  code: number
  msg: string
  data: {
    list: ConversationMessage[]
    total: number
  }
}

export interface ConversationSessionsResponse {
  code: number
  msg: string
  data: {
    list: ConversationSession[]
    total: number
  }
}
