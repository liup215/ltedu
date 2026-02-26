import apiClient from './apiClient';
import type { 
  Question, 
  QuestionQuery, 
  PaginatedQuestions,
  QuestionCreateRequest,
  QuestionUpdateRequest,
} from '../models/question.model';
import type { ApiResponse } from '../models/api.model';

class QuestionService {
  private baseUrl = '/api/v1/question';

  // GET LIST - Paginated list with filtering
  async getQuestions(query: QuestionQuery = {}): Promise<ApiResponse<PaginatedQuestions>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post(`${this.baseUrl}/list`, query);
    return response.data;
  }

  // GET BY ID - Single question by ID
  async getQuestionById(id: number): Promise<ApiResponse<Question>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post(`${this.baseUrl}/byId`, { id });
    return response.data;
  }

  // GET ALL - All questions (for dropdowns/selects)
  async getAllQuestions(query: QuestionQuery = {}): Promise<ApiResponse<PaginatedQuestions>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post(`${this.baseUrl}/all`, query);
    return response.data;
  }

  // CREATE - Create new question
  async createQuestion(question: QuestionCreateRequest): Promise<ApiResponse<Question>> {
    const client = await apiClient();
    const response = await client.post(`${this.baseUrl}/create`, question);
    return response.data;
  }

  // UPDATE - Update existing question
  async updateQuestion(question: QuestionUpdateRequest): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post(`${this.baseUrl}/edit`, question);
    return response.data;
  }

  // DELETE - Delete question
  async deleteQuestion(id: number): Promise<ApiResponse<void>> {
    const client = await apiClient(); // Ensure we use the async client
    const response = await client.post(`${this.baseUrl}/delete`, { id });
    return response.data;
  }
}

export default new QuestionService();
