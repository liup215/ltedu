import type { ApiResponse, Page } from './api.model'

export const FEEDBACK_TYPE_GENERAL = 'general'
export const FEEDBACK_TYPE_BUG = 'bug'
export const FEEDBACK_TYPE_FEATURE = 'feature'
export const FEEDBACK_TYPE_PRAISE = 'praise'

export const FEEDBACK_STATUS_NEW = 'new'
export const FEEDBACK_STATUS_REVIEWED = 'reviewed'
export const FEEDBACK_STATUS_RESOLVED = 'resolved'

export const FEEDBACK_SENTIMENT_POSITIVE = 'positive'
export const FEEDBACK_SENTIMENT_NEUTRAL = 'neutral'
export const FEEDBACK_SENTIMENT_NEGATIVE = 'negative'

export interface UserFeedback {
  id: number
  userId: number
  type: string
  content: string
  rating: number
  sentiment: string
  pageContext: string
  consentGiven: boolean
  status: string
  adminNote: string
  userAgent: string
  createdAt: string
  updatedAt: string
}

export interface SubmitFeedbackRequest {
  type?: string
  content: string
  rating?: number
  pageContext?: string
  consentGiven: boolean
  userAgent?: string
}

export interface FeedbackListRequest extends Page {
  status?: string
  type?: string
  userId?: number
}

export interface FeedbackListData {
  list: UserFeedback[]
  total: number
}

export interface FeedbackStats {
  total: number
  byType: Record<string, number>
  bySentiment: Record<string, number>
  byStatus: Record<string, number>
  avgRating: number
}

export type FeedbackSubmitResponse = ApiResponse<UserFeedback>
export type FeedbackListResponse = ApiResponse<FeedbackListData>
export type FeedbackStatsResponse = ApiResponse<FeedbackStats>
