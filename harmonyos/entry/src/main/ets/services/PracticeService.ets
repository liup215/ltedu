import apiService from './ApiService';
import type {
  ApiResponse,
  Question,
  PracticeResult,
  QuickPracticeRequest,
  GradeRequest
} from '../models/ApiModels';

/**
 * PracticeService provides methods for quick practice, paper-based practice,
 * and grading submitted answers against the LTEdu backend.
 */
export class PracticeService {
  private static instance: PracticeService;

  static getInstance(): PracticeService {
    if (!PracticeService.instance) {
      PracticeService.instance = new PracticeService();
    }
    return PracticeService.instance;
  }

  async quickPractice(request: QuickPracticeRequest): Promise<ApiResponse<PracticeResult>> {
    return apiService.post<PracticeResult>('/api/v1/practice/quick', request);
  }

  async paperPractice(request: {
    paperId: number;
    pageIndex?: number;
    pageSize?: number;
  }): Promise<ApiResponse<PracticeResult>> {
    return apiService.post<PracticeResult>('/api/v1/practice/paper', request);
  }

  async gradeAnswers(request: GradeRequest): Promise<ApiResponse<Record<string, boolean>>> {
    return apiService.post<Record<string, boolean>>('/api/v1/practice/grade', request);
  }

  async getQuestions(request: {
    syllabusId?: number;
    qualificationId?: number;
    questionType?: number;
    pageIndex?: number;
    pageSize?: number;
  }): Promise<ApiResponse<PracticeResult>> {
    return apiService.post<PracticeResult>('/api/v1/question/list', request);
  }
}

export default PracticeService.getInstance();
