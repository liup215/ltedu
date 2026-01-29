import apiClient from './apiClient';
import type { ApiResponse } from '../models/api.model';
import type {
  User,
  UserQueryCriteria,
  PaginatedUsers,
  UserCreationPayload,
  UserUpdatePayload
} from '../models/user.model';

const USER_API_BASE_PATH = '/api/v1/user';

/**
 * Fetches a paginated list of users.
 * Corresponds to: POST /backend/api/v1/user/list
 */
async function getUsers(query: UserQueryCriteria): Promise<ApiResponse<PaginatedUsers>> {
  const client = await apiClient(); // Ensure we use the async client
  const axiosResponse = await client.post(`${USER_API_BASE_PATH}/list`, query);
  return axiosResponse.data;
}

/**
 * Fetches a single user by their ID.
 * Corresponds to: POST /backend/api/v1/user/byId
 */
async function getUserById(id: number): Promise<ApiResponse<User>> {
  const client = await apiClient(); // Ensure we use the async client
  const axiosResponse = await client.post(`${USER_API_BASE_PATH}/byId`, { id });
  return axiosResponse.data;
}

/**
 * Creates a new user.
 * Corresponds to: POST /backend/api/v1/user/create
 */
async function createUser(userData: UserCreationPayload): Promise<ApiResponse<User>> {
  const client = await apiClient(); // Ensure we use the async client
  const axiosResponse = await client.post(`${USER_API_BASE_PATH}/create`, userData);
  return axiosResponse.data;
}

/**
 * Updates an existing user.
 * The user ID is passed as a parameter.
 * Corresponds to: POST /backend/api/v1/user/edit
 */
async function updateUser(id: number, userData: UserUpdatePayload): Promise<ApiResponse<void>> {
  const client = await apiClient(); // Ensure we use the async client
  const payloadWithId = { ...userData, id }; 
  const axiosResponse = await client.post(`${USER_API_BASE_PATH}/edit`, payloadWithId);
  return axiosResponse.data;
}

/**
 * Deletes a user by their ID.
 * Corresponds to: POST /backend/api/v1/user/delete
 */
async function deleteUser(id: number): Promise<ApiResponse<void>> {
  const client = await apiClient(); // Ensure we use the async client
  const axiosResponse = await client.post(`${USER_API_BASE_PATH}/delete`, { id });
  return axiosResponse.data;
}

/**
 * Sets a user as admin.
 * Corresponds to: POST /backend/api/v1/user/setAdmin
 */
async function setAdmin(id: number): Promise<ApiResponse<void>> {
  const client = await apiClient();
  const axiosResponse = await client.post(`${USER_API_BASE_PATH}/setAdmin`, { id });
  return axiosResponse.data;
}

/**
 * Removes admin privileges from a user.
 * Corresponds to: POST /backend/api/v1/user/removeAdmin
 */
async function removeAdmin(id: number): Promise<ApiResponse<void>> {
  const client = await apiClient();
  const axiosResponse = await client.post(`${USER_API_BASE_PATH}/removeAdmin`, { id });
  return axiosResponse.data;
}

/**
 * Grants one month VIP to a user.
 * Corresponds to: POST /api/v1/user/vip
 */
async function grantVip(id: number): Promise<ApiResponse<void>> {
  const client = await apiClient();
  const axiosResponse = await client.post(`${USER_API_BASE_PATH}/vip`, { id });
  return axiosResponse.data;
}

export default {
  getUsers,
  getUserById,
  createUser,
  updateUser,
  deleteUser,
  setAdmin,
  removeAdmin,
  grantVip,
};
