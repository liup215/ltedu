import apiService from './ApiService';
import type { ApiResponse } from '../models/ApiModels';

export interface AiChatResponse {
  text: string;
  score?: number;
}

export interface AiCheckRequest {
  stem: string;
  answer: string;
  studentAnswer: string;
}

export interface AiQuestionCheckRequest {
  stem: string;
  markScheme: string;
  studentAnswer: string;
  fullMark: number;
}

/**
 * AiService provides AI-powered features: vocabulary checking, question
 * grading, and a general-purpose chat interface backed by the LTEdu AI
 * backend.
 */
export class AiService {
  private static instance: AiService;
  private conversationHistory: Array<{ role: string; content: string }> = [];

  static getInstance(): AiService {
    if (!AiService.instance) {
      AiService.instance = new AiService();
    }
    return AiService.instance;
  }

  /**
   * Checks a vocabulary answer and returns AI feedback and a score.
   */
  async checkVocabulary(request: AiCheckRequest): Promise<ApiResponse<AiChatResponse>> {
    return apiService.post<AiChatResponse>('/api/v1/ai/check-vocabulary', request);
  }

  /**
   * Checks a short/long answer question and returns AI feedback with score.
   */
  async checkQuestion(request: AiQuestionCheckRequest): Promise<ApiResponse<AiChatResponse>> {
    return apiService.post<AiChatResponse>('/api/v1/ai/check-question', request);
  }

  /**
   * Sends a chat message to the MCP-compatible AI endpoint and returns
   * the assistant's response. Maintains conversation history for context.
   */
  async sendChatMessage(userMessage: string): Promise<string> {
    this.conversationHistory.push({ role: 'user', content: userMessage });

    try {
      const response = await apiService.post<AiChatResponse>('/api/v1/ai/chat', {
        messages: this.conversationHistory,
      });

      if (response.code === 0 && response.data?.text) {
        const assistantMessage = response.data.text;
        this.conversationHistory.push({ role: 'assistant', content: assistantMessage });
        return assistantMessage;
      }
      return 'Sorry, I was unable to process your request. Please try again.';
    } catch (err) {
      return 'Network error. Please check your connection and try again.';
    }
  }

  clearHistory(): void {
    this.conversationHistory = [];
  }

  getHistory(): Array<{ role: string; content: string }> {
    return [...this.conversationHistory];
  }
}

export default AiService.getInstance();
