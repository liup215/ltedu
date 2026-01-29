export interface Qualification {
  id: number;
  name: string;
  organisationId: number;
  organisation?: {
    id: number;
    name: string;
  };
}

export interface QualificationQuery {
  pageIndex?: number;
  pageSize?: number;
  id?: number;
  organisationId?: number;
}

export interface PaginatedQualifications {
  list: Qualification[];
  total: number;
}

export interface QualificationCreateRequest {
  name: string;
  organisationId: number;
}

export interface QualificationUpdateRequest {
  id: number;
  name: string;
  organisationId: number;
}
