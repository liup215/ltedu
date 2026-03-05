"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g = Object.create((typeof Iterator === "function" ? Iterator : Object).prototype);
    return g.next = verb(0), g["throw"] = verb(1), g["return"] = verb(2), typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.useUserStore = void 0;
var pinia_1 = require("pinia");
var authService_1 = require("../services/authService"); // Import authService
exports.useUserStore = (0, pinia_1.defineStore)('user', {
    state: function () { return ({
        user: null,
        token: localStorage.getItem('authToken') || null,
        isAuthenticated: !!localStorage.getItem('authToken'),
    }); },
    getters: {
        currentUser: function (state) { return state.user; },
        currentToken: function (state) { return state.token; },
        isUserAuthenticated: function (state) { return state.isAuthenticated; },
    },
    actions: {
        setUser: function (userData) {
            this.user = userData;
            if (userData) {
                localStorage.setItem('user', JSON.stringify(userData));
            }
            else {
                localStorage.removeItem('user');
            }
        },
        setToken: function (tokenData) {
            this.token = tokenData;
            if (tokenData) {
                localStorage.setItem('authToken', tokenData);
                this.isAuthenticated = true;
            }
            else {
                localStorage.removeItem('authToken');
                this.isAuthenticated = false;
            }
        },
        login: function (token) {
            return __awaiter(this, void 0, void 0, function () {
                var userProfile, error_1;
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0:
                            this.setToken(token);
                            if (!token) return [3 /*break*/, 5];
                            _a.label = 1;
                        case 1:
                            _a.trys.push([1, 3, , 4]);
                            return [4 /*yield*/, authService_1.authService.getCurrentUserProfile()];
                        case 2:
                            userProfile = _a.sent();
                            this.setUser(userProfile.data);
                            console.log('User profile fetched and set:', userProfile);
                            return [3 /*break*/, 4];
                        case 3:
                            error_1 = _a.sent();
                            console.error('Failed to fetch user profile after login:', error_1);
                            this.logout(); // Clear token and user if profile fetch fails
                            return [3 /*break*/, 4];
                        case 4: return [3 /*break*/, 6];
                        case 5:
                            // If no token, ensure logout state
                            this.logout();
                            _a.label = 6;
                        case 6: return [2 /*return*/];
                    }
                });
            });
        },
        logout: function () {
            this.setUser(null);
            this.setToken(null);
            // Potentially call a backend logout endpoint via authService
        },
        // Action to load user from localStorage on app init
        loadUserFromStorage: function () {
            var storedToken = localStorage.getItem('authToken');
            var storedUser = localStorage.getItem('user');
            if (storedToken && storedUser) {
                this.token = storedToken;
                this.user = JSON.parse(storedUser);
                this.isAuthenticated = true;
            }
            else {
                this.logout(); // Ensure clean state if no token/user
            }
        },
        // Action: fetch current user info from backend and update state
        fetchCurrentUser: function () {
            return __awaiter(this, void 0, void 0, function () {
                var userProfile, error_2;
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0:
                            _a.trys.push([0, 2, , 3]);
                            return [4 /*yield*/, authService_1.authService.getCurrentUserProfile()];
                        case 1:
                            userProfile = _a.sent();
                            this.setUser(userProfile.data);
                            return [3 /*break*/, 3];
                        case 2:
                            error_2 = _a.sent();
                            console.error('Failed to fetch current user:', error_2);
                            return [3 /*break*/, 3];
                        case 3: return [2 /*return*/];
                    }
                });
            });
        }
    },
});
