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
var axios_1 = require("axios");
var userStore_1 = require("../stores/userStore");
var appStore_1 = require("../stores/appStore");
var config_1 = require("../const/config");
var notification_1 = require("../utils/notification");
// import { useRouter } from 'vue-router';
var router_1 = require("../router");
var router = router_1.default;
var getApiClient = function () { return __awaiter(void 0, void 0, void 0, function () {
    // Helper function to check if response is ApiResponse format
    function isApiResponse(data) {
        return data && typeof data === 'object' && 'code' in data && 'message' in data;
    }
    var apiClient, _a, _b;
    var _c;
    return __generator(this, function (_d) {
        switch (_d.label) {
            case 0:
                _b = (_a = axios_1.default).create;
                _c = {};
                return [4 /*yield*/, (0, config_1.API_BASE_URL)()];
            case 1:
                apiClient = _b.apply(_a, [(_c.baseURL = _d.sent(),
                        _c.timeout = 100000,
                        _c.headers = {
                            'Content-Type': 'application/json',
                        },
                        _c)]);
                // Request Interceptor
                apiClient.interceptors.request.use(function (config) {
                    var userStore = (0, userStore_1.useUserStore)();
                    var appStore = (0, appStore_1.useAppStore)();
                    // Set loading state to true
                    appStore.setLoading(true);
                    // Add auth token to headers if available
                    if (userStore.currentToken) {
                        config.headers.Authorization = "Bearer ".concat(userStore.currentToken);
                    }
                    return config;
                }, function (error) {
                    var appStore = (0, appStore_1.useAppStore)();
                    appStore.setLoading(false); // Ensure loading is false on error
                    return Promise.reject(error);
                });
                // Response Interceptor
                apiClient.interceptors.response.use(function (response) {
                    var appStore = (0, appStore_1.useAppStore)();
                    // Set loading state to false
                    appStore.setLoading(false);
                    // Check if response has our custom ApiResponse format
                    if (isApiResponse(response.data)) {
                        // Handle custom response format: code 0 = success, others = error
                        if (response.data.code === 0) {
                            // Success: modify response to contain the ApiResponse
                            response.data = response.data;
                            return response;
                        }
                        else {
                            // Error: throw error with custom message
                            var errorMessage = response.data.message || 'Operation failed';
                            (0, notification_1.showError)(errorMessage);
                            var error = new Error(errorMessage);
                            error.code = response.data.code;
                            error.data = response.data.data;
                            return Promise.reject(error);
                        }
                    }
                    // For responses without custom format, return as is
                    return response;
                }, function (error) {
                    var appStore = (0, appStore_1.useAppStore)();
                    var userStore = (0, userStore_1.useUserStore)();
                    appStore.setLoading(false); // Ensure loading is false on error
                    if (error.response) {
                        // Handle common HTTP error statuses
                        if (error.response.status === 401) {
                            // const responseData = error.response.data as ApiResponse;
                            var message = 'Please log in!';
                            (0, notification_1.showError)(message);
                            // Unauthorized: e.g., token expired or invalid
                            userStore.logout(); // Clear user session
                            // Optionally redirect to login page
                            router.push('/login'); // Make sure router is accessible here or handle in component
                        }
                        else if (error.response.status === 403) {
                            // Forbidden
                            (0, notification_1.showError)('You do not have permission to perform this action.');
                        }
                        // Check if error response has our custom ApiResponse format
                        if (isApiResponse(error.response.data)) {
                            var errorMessage = error.response.data.message || 'Operation failed';
                            // showError(errorMessage);
                            var customError = new Error(errorMessage);
                            customError.code = error.response.data.code;
                            customError.data = error.response.data.data;
                            return Promise.reject(customError);
                        }
                        // You can add more specific error handling here
                    }
                    else if (error.request) {
                        // The request was made but no response was received
                        (0, notification_1.showError)('Network error or no response from server');
                    }
                    else {
                        // Something happened in setting up the request that triggered an Error
                        (0, notification_1.showError)('Error setting up request');
                    }
                    return Promise.reject(error);
                });
                return [2 /*return*/, apiClient];
        }
    });
}); };
exports.default = getApiClient;
