import apiClient from './apiClient';
import type { ApiResponse } from '../models/api.model';
import type {
  ConversationSession,
  ConversationMessage,
  SendMessageResponse,
  StartSessionRequest,
  SendMessageRequest,
  ConversationHistoryRequest,
  ResetContextRequest,
} from '../models/conversation.model';

const CONVERSATION_API_BASE = '/api/v1/ai/conversation';

/**
 * Start a new AI conversation session.
 * POST /api/v1/ai/conversation/start
 */
async function startSession(
  data?: StartSessionRequest,
): Promise<ApiResponse<ConversationSession>> {
  const client = await apiClient();
  const response = await client.post<ApiResponse<ConversationSession>>(
    `${CONVERSATION_API_BASE}/start`,
    data ?? {},
  );
  return response.data;
}

/**
 * Send a message in an existing conversation session.
 * POST /api/v1/ai/conversation/message
 */
async function sendMessage(
  data: SendMessageRequest,
): Promise<ApiResponse<SendMessageResponse>> {
  const client = await apiClient();
  const response = await client.post<ApiResponse<SendMessageResponse>>(
    `${CONVERSATION_API_BASE}/message`,
    data,
  );
  return response.data;
}

/**
 * Retrieve the full message history for a session.
 * POST /api/v1/ai/conversation/history
 */
async function getHistory(
  data: ConversationHistoryRequest,
): Promise<ApiResponse<ConversationMessage[]>> {
  const client = await apiClient();
  const response = await client.post<ApiResponse<ConversationMessage[]>>(
    `${CONVERSATION_API_BASE}/history`,
    data,
  );
  return response.data;
}

/**
 * List all active conversation sessions for the current user.
 * POST /api/v1/ai/conversation/sessions
 */
async function listSessions(): Promise<ApiResponse<ConversationSession[]>> {
  const client = await apiClient();
  const response = await client.post<ApiResponse<ConversationSession[]>>(
    `${CONVERSATION_API_BASE}/sessions`,
    {},
  );
  return response.data;
}

/**
 * Reset the context of a conversation session (clear history and optionally set new context).
 * POST /api/v1/ai/conversation/reset
 */
async function resetContext(
  data: ResetContextRequest,
): Promise<ApiResponse<null>> {
  const client = await apiClient();
  const response = await client.post<ApiResponse<null>>(
    `${CONVERSATION_API_BASE}/reset`,
    data,
  );
  return response.data;
}

/**
 * Close (deactivate) a conversation session.
 * POST /api/v1/ai/conversation/close
 */
async function closeSession(
  sessionKey: string,
): Promise<ApiResponse<null>> {
  const client = await apiClient();
  const response = await client.post<ApiResponse<null>>(
    `${CONVERSATION_API_BASE}/close`,
    { sessionKey },
  );
  return response.data;
}

export const conversationService = {
  startSession,
  sendMessage,
  getHistory,
  listSessions,
  resetContext,
  closeSession,
};
