import { defineStore } from 'pinia';
import { authService } from '../services/authService'; // Import authService
import type { User } from '../models/user.model'; // Import User model

// Local User interface definition removed

interface UserState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
}

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    user: null,
    token: localStorage.getItem('authToken') || null,
    isAuthenticated: !!localStorage.getItem('authToken'),
  }),
  getters: {
    currentUser: (state) => state.user,
    currentToken: (state) => state.token,
    isUserAuthenticated: (state) => state.isAuthenticated,
  },
  actions: {
    setUser(userData: User | null) {
      this.user = userData;
      if (userData) {
        localStorage.setItem('user', JSON.stringify(userData));
      } else {
        localStorage.removeItem('user');
      }
    },
    setToken(tokenData: string | null) {
      this.token = tokenData;
      if (tokenData) {
        localStorage.setItem('authToken', tokenData);
        this.isAuthenticated = true;
      } else {
        localStorage.removeItem('authToken');
        this.isAuthenticated = false;
      }
    },
    async login(token: string) {
      this.setToken(token);
      if (token) {
        try {
          // authService.getCurrentUserProfile() will call GET /v1/admin/user
          // This service method needs to be created in authService.ts
          // It should return a Promise<User> after mapping the backend response.
          const userProfile = await authService.getCurrentUserProfile(); 
          this.setUser(userProfile.data);
          console.log('User profile fetched and set:', userProfile);
        } catch (error) {
          console.error('Failed to fetch user profile after login:', error);
          this.logout(); // Clear token and user if profile fetch fails
        }
      } else {
        // If no token, ensure logout state
        this.logout();
      }
    },
    logout() {
      this.setUser(null);
      this.setToken(null);
      // Potentially call a backend logout endpoint via authService
    },
    // Action to load user from localStorage on app init
    loadUserFromStorage() {
      const storedToken = localStorage.getItem('authToken');
      const storedUser = localStorage.getItem('user');
      if (storedToken && storedUser) {
        this.token = storedToken;
        this.user = JSON.parse(storedUser);
        this.isAuthenticated = true;
      } else {
        this.logout(); // Ensure clean state if no token/user
      }
    },
    // Action: fetch current user info from backend and update state
    async fetchCurrentUser() {
      try {
        const userProfile = await authService.getCurrentUserProfile();
        this.setUser(userProfile.data);
      } catch (error) {
        console.error('Failed to fetch current user:', error);
        // Optionally, logout or show error
      }
    }
  },
});
