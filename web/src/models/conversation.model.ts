// Models for the AI multi-turn conversation context management feature

export interface ConversationContext {
  userRole?: string;
  preferences?: Record<string, string>;
  recentActions?: string[];
  currentSelection?: Record<string, unknown>;
}

export interface ConversationSession {
  sessionKey: string;
  userRole: string;
  messageCount: number;
  lastActiveAt: string;
  expiresAt: string;
  isActive: boolean;
  createdAt: string;
}

export interface ConversationMessage {
  id: number;
  sessionId: number;
  role: 'user' | 'assistant' | 'system';
  content: string;
  orderIndex: number;
  createdAt: string;
}

export interface StartSessionRequest {
  context?: ConversationContext;
}

export interface SendMessageRequest {
  sessionKey: string;
  message: string;
}

export interface SendMessageResponse {
  userMessage: ConversationMessage;
  assistantMessage: ConversationMessage;
  sessionKey: string;
  messageCount: number;
}

export interface ConversationHistoryRequest {
  sessionKey: string;
}

export interface ResetContextRequest {
  sessionKey: string;
  context?: ConversationContext;
}
