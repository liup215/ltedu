// Represents permission data
export interface PermissionData {
  id: number;
  name: string;
  slug: string;
}

// Represents role data, e.g., model.AdminRole
export interface AdminRoleData {
  id: number;
  displayName: string;
  slug: string;
  description?: string;
  permissions?: PermissionData[];
}

// Represents the data part of ApiResponse for a user profile
export interface UserProfileData {
  id: number;
  username: string;
  email?: string | null;
  mobile?: string | null;
  isAdmin: boolean;
  adminRoleId?: number | null;
  adminRole?: AdminRoleData | null;
}

// Main User interface used in the application (e.g., in userStore)
// This will be moved from userStore.ts here.
export interface User {
  id: number | null;
  username: string | null;
  email?: string | null;
  isAdmin: boolean;
  isTeacher: boolean;
  adminRoleId?: number | null;
  adminRole?: AdminRoleData | null;
  nickname?: string;
  realname?: string;
  engname?: string;
  sex?: number; // e.g., 0: unknown, 1: male, 2: female
  mobile?: string;
  avatar?: string;
  status: number; // e.g., 1 for active, 0 for inactive.
  isActive?: number; // From API response example.
  createdAt: string; // ISO date string (registration date)
  updatedAt: string; // ISO date string
  teacherApplyStatus?: number; // Status of teacher application if any
  teacherApplicationId?: number | null; // ID of teacher application if any
  vipExpireAt?: number | null; // Unix timestamp for VIP expiration
}

// Criteria for querying the user list (admin context)
export interface UserQueryCriteria {
  pageSize?: number;
  pageIndex?: number;
  username?: string;
  realname?: string;
  mobile?: string;
  status?: number; // API specific: e.g., 0 for all, 1 for active
}

// Structure for paginated user records response (admin context)
export interface PaginatedUsers { // Renamed from PaginatedUserRecords
  list: User[]; // Uses the unified User interface
  total: number;
}

// Payload for creating a new user (admin context)
export interface UserCreationPayload {
  username: string;
  password?: string; // Should be required for new user
  nickname?: string;
  realname?: string;
  engname?: string;
  email?: string;
  sex?: number;
  mobile?: string;
  avatar?: string;
  status: number; // e.g., 1 for active
}

// Payload for updating an existing user (admin context)
// The ID for update will typically be passed as a separate parameter to the service function.
export interface UserUpdatePayload {
  username?: string;
  nickname?: string;
  realname?: string;
  engname?: string;
  email?: string;
  sex?: number;
  mobile?: string;
  avatar?: string;
  status?: number;
  isActive?: number;
  roleId?: number; // Corresponds to User.adminRoleId
  password?: string; // For password changes (often handled separately)
}
