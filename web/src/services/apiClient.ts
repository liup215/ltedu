import axios from 'axios';
import type { AxiosInstance, InternalAxiosRequestConfig, AxiosResponse, AxiosError } from 'axios';
import { useUserStore } from '../stores/userStore';
import { useAppStore } from '../stores/appStore';
import type { ApiResponse } from '../models/api.model';
import { API_BASE_URL } from '../const/config';
import { showError } from '../utils/notification'
// import { useRouter } from 'vue-router';
import Vroute from '../router'

 const router = Vroute

const getApiClient = async (): Promise<AxiosInstance> => {

const apiClient: AxiosInstance = axios.create({
  baseURL: await API_BASE_URL(),
  timeout: 100000, // 100 seconds timeout
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request Interceptor
apiClient.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {

    const userStore = useUserStore();
    const appStore = useAppStore();

    // Set loading state to true
    appStore.setLoading(true);

    // Add auth token to headers if available
    if (userStore.currentToken) {
      config.headers.Authorization = `Bearer ${userStore.currentToken}`;
    }
    return config;
  },
  (error: AxiosError) => {
    const appStore = useAppStore();
    appStore.setLoading(false); // Ensure loading is false on error
    return Promise.reject(error);
  }
);

// Helper function to check if response is ApiResponse format
function isApiResponse(data: any): data is ApiResponse {
  return data && typeof data === 'object' && 'code' in data && 'message' in data;
}

// Response Interceptor
apiClient.interceptors.response.use(
  (response: AxiosResponse) => {
    const appStore = useAppStore();
    // Set loading state to false
    appStore.setLoading(false);
    
    // Check if response has our custom ApiResponse format
    if (isApiResponse(response.data)) {
      // Handle custom response format: code 0 = success, others = error
      if (response.data.code === 0) {
        // Success: modify response to contain the ApiResponse
        response.data = response.data;
        return response;
      } else {
        // Error: throw error with custom message
        const errorMessage = response.data.message || 'Operation failed';
        showError(errorMessage);
        const error = new Error(errorMessage);
        (error as any).code = response.data.code;
        (error as any).data = response.data.data;
        return Promise.reject(error);
      }
    }
    
    // For responses without custom format, return as is
    return response;
  },
  (error: AxiosError) => {
    const appStore = useAppStore();
    const userStore = useUserStore();
    appStore.setLoading(false); // Ensure loading is false on error

    if (error.response) {
      // Handle common HTTP error statuses
      if (error.response.status === 401) {
        // const responseData = error.response.data as ApiResponse;
        const message = 'Please log in!';
        showError(message);
        // Unauthorized: e.g., token expired or invalid
        userStore.logout(); // Clear user session
        // Optionally redirect to login page
        router.push('/login'); // Make sure router is accessible here or handle in component
      } else if (error.response.status === 403) {
        // Forbidden
        showError('You do not have permission to perform this action.');
      }
      
      // Check if error response has our custom ApiResponse format
      if (isApiResponse(error.response.data)) {
        const errorMessage = error.response.data.message || 'Operation failed';
        // showError(errorMessage);
        const customError = new Error(errorMessage);
        (customError as any).code = error.response.data.code;
        (customError as any).data = error.response.data.data;
        return Promise.reject(customError);
      }
      
      // You can add more specific error handling here
    } else if (error.request) {
      // The request was made but no response was received
      showError('Network error or no response from server');
    } else {
      // Something happened in setting up the request that triggered an Error
      showError('Error setting up request');
    }
    return Promise.reject(error);
  }
);

return apiClient;

}

export default getApiClient;
