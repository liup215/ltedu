"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.useAppStore = void 0;
var pinia_1 = require("pinia");
exports.useAppStore = (0, pinia_1.defineStore)('app', {
    state: function () { return ({
        isLoading: false,
        theme: 'light', // Default theme
    }); },
    getters: {
        currentTheme: function (state) { return state.theme; },
        appIsLoading: function (state) { return state.isLoading; },
    },
    actions: {
        setLoading: function (loadingStatus) {
            this.isLoading = loadingStatus;
        },
        toggleTheme: function () {
            this.theme = this.theme === 'light' ? 'dark' : 'light';
            // You might want to save theme preference to localStorage
            localStorage.setItem('appTheme', this.theme);
            // And apply it to the document body or root element
            document.documentElement.setAttribute('data-theme', this.theme);
            if (this.theme === 'dark') {
                document.documentElement.classList.add('dark');
            }
            else {
                document.documentElement.classList.remove('dark');
            }
        },
        // Action to load theme from localStorage on app init
        loadThemeFromStorage: function () {
            var storedTheme = localStorage.getItem('appTheme');
            if (storedTheme) {
                this.theme = storedTheme;
            }
            // Apply initial theme
            document.documentElement.setAttribute('data-theme', this.theme);
            if (this.theme === 'dark') {
                document.documentElement.classList.add('dark');
            }
            else {
                document.documentElement.classList.remove('dark');
            }
        }
    },
});
