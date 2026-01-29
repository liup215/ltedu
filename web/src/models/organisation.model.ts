// Base organisation model
export interface Organisation {
  id: number;
  name: string;
}

// Query criteria for fetching organisations
export interface OrganisationQuery {
  pageIndex?: number;
  pageSize?: number;
  id?: number;
  name?: string;
}

// Paginated response for organisation list
export interface PaginatedOrganisations {
  list: Organisation[];
  total: number;
}
