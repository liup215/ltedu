import { defineStore } from 'pinia';

interface AppState {
  isLoading: boolean;
  theme: 'light' | 'dark';
  // Add other global app settings as needed
}

export const useAppStore = defineStore('app', {
  state: (): AppState => ({
    isLoading: false,
    theme: 'light', // Default theme
  }),
  getters: {
    currentTheme: (state) => state.theme,
    appIsLoading: (state) => state.isLoading,
  },
  actions: {
    setLoading(loadingStatus: boolean) {
      this.isLoading = loadingStatus;
    },
    toggleTheme() {
      this.theme = this.theme === 'light' ? 'dark' : 'light';
      // You might want to save theme preference to localStorage
      localStorage.setItem('appTheme', this.theme);
      // And apply it to the document body or root element
      document.documentElement.setAttribute('data-theme', this.theme);
      if (this.theme === 'dark') {
        document.documentElement.classList.add('dark');
      } else {
        document.documentElement.classList.remove('dark');
      }
    },
    // Action to load theme from localStorage on app init
    loadThemeFromStorage() {
      const storedTheme = localStorage.getItem('appTheme') as 'light' | 'dark' | null;
      if (storedTheme) {
        this.theme = storedTheme;
      }
      // Apply initial theme
      document.documentElement.setAttribute('data-theme', this.theme);
      if (this.theme === 'dark') {
        document.documentElement.classList.add('dark');
      } else {
        document.documentElement.classList.remove('dark');
      }
    }
  },
});
