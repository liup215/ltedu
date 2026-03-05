export interface BlogPost {
  id: number
  title: string
  slug: string
  summary: string
  content: string
  coverImage: string
  category: string
  tags: string
  authorId: number
  status: 'draft' | 'published'
  viewCount: number
  isTop: boolean
  createdAt: string
  updatedAt: string
}

export interface BlogPostCreateRequest {
  title: string
  summary?: string
  content: string
  coverImage?: string
  category?: string
  tags?: string
  status?: 'draft' | 'published'
}

export interface BlogPostUpdateRequest {
  id: number
  title?: string
  summary?: string
  content?: string
  coverImage?: string
  category?: string
  tags?: string
  status?: 'draft' | 'published'
  isTop?: boolean
}

export interface BlogPostQuery {
  pageIndex?: number
  pageSize?: number
  category?: string
  keyword?: string
  status?: string
}

export interface BlogListResponse {
  code: number
  msg: string
  data: {
    list: BlogPost[]
    total: number
  }
}

export interface BlogPostResponse {
  code: number
  msg: string
  data: BlogPost
}
