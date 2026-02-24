import apiClient from './apiClient'
import type {
  KnowledgePointCreateRequest,
  KnowledgePointUpdateRequest,
  KnowledgePointListRequest,
  KnowledgePointListResponse,
  KnowledgePointResponse,
  GenerateKeypointsRequest,
  GenerateKeypointsResponse,
  AutoLinkRequest,
  AutoLinkResponse,
  MigrateSyllabusRequest,
  MigrateSyllabusResponse
} from '../models/knowledgePoint.model'

class KnowledgePointService {
  // CRUD Operations
  async create(request: KnowledgePointCreateRequest): Promise<KnowledgePointResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/knowledge-point/create', request)
    return response.data
  }

  async update(request: KnowledgePointUpdateRequest): Promise<KnowledgePointResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/knowledge-point/edit', request)
    return response.data
  }

  async delete(id: number): Promise<KnowledgePointResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/knowledge-point/delete', { id })
    return response.data
  }

  async getById(id: number): Promise<KnowledgePointResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/knowledge-point/byId', { id })
    return response.data
  }

  async getByChapter(chapterId: number): Promise<KnowledgePointListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/knowledge-point/byChapter', { chapterId })
    return response.data
  }

  async getBySyllabus(syllabusId: number): Promise<KnowledgePointListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/knowledge-point/bySyllabus', { syllabusId })
    return response.data
  }

  async list(request: KnowledgePointListRequest): Promise<KnowledgePointListResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/knowledge-point/list', request)
    return response.data
  }

  // AI Automation
  async generateKeypoints(request: GenerateKeypointsRequest): Promise<GenerateKeypointsResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/chapter/generate-keypoints', request)
    return response.data
  }

  async autoLinkQuestion(request: AutoLinkRequest): Promise<AutoLinkResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/question/auto-link-keypoints', request)
    return response.data
  }

  async migrateSyllabus(request: MigrateSyllabusRequest): Promise<MigrateSyllabusResponse> {
    const client = await apiClient()
    const response = await client.post('/api/v1/syllabus/auto-migrate-keypoints', request)
    return response.data
  }
}

export default new KnowledgePointService()
