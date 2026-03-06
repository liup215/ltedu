export interface BlogPost {
  id: number
  title: string
  slug: string
  summary: string
  content: string
  category: string
  tags: string
  coverImage: string
  status: string
  authorId: number
  authorName: string
  viewCount: number
  publishedAt?: string
  createdAt: string
  updatedAt: string
}

export interface BlogPostCreateRequest {
  id?: number
  title: string
  slug?: string
  summary: string
  content: string
  category: string
  tags: string
  coverImage: string
  status: string
}

export interface BlogPostQuery {
  id?: number
  category?: string
  status?: string
  keyword?: string
  pageSize?: number
  pageIndex?: number
}

export const BLOG_CATEGORIES = [
  { value: 'system_updates', labelEn: 'System Updates', labelZh: '系统改进' },
  { value: 'user_guides', labelEn: 'User Guides', labelZh: '操作说明' },
  { value: 'learning_methods', labelEn: 'Learning Methods', labelZh: '学习方法' },
  { value: 'major_events', labelEn: 'Major Events', labelZh: '大事件' },
]

export const BLOG_STATUSES = [
  { value: 'draft', labelEn: 'Draft', labelZh: '草稿' },
  { value: 'published', labelEn: 'Published', labelZh: '已发布' },
]
