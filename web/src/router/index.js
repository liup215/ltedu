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
var __spreadArray = (this && this.__spreadArray) || function (to, from, pack) {
    if (pack || arguments.length === 2) for (var i = 0, l = from.length, ar; i < l; i++) {
        if (ar || !(i in from)) {
            if (!ar) ar = Array.prototype.slice.call(from, 0, i);
            ar[i] = from[i];
        }
    }
    return to.concat(ar || Array.prototype.slice.call(from));
};
Object.defineProperty(exports, "__esModule", { value: true });
var vue_router_1 = require("vue-router");
var Login_vue_1 = require("../views/Login.vue");
var Home_vue_1 = require("../views/Home.vue");
var Layout_vue_1 = require("../components/Layout.vue");
var Register_vue_1 = require("../views/Register.vue");
// Modular route imports
var routes_1 = require("../views/account/routes");
var routes_2 = require("../views/admin/routes");
// Import ExamPaperForm for teacher paper builder
var ExamPaperForm = function () { return Promise.resolve().then(function () { return require('../views/paper/ExamPaperForm.vue'); }); };
var TeacherExamPaperManagement = function () { return Promise.resolve().then(function () { return require('../views/paper/TeacherExamPaperManagement.vue'); }); };
var ExamPaperPreview = function () { return Promise.resolve().then(function () { return require('../views/paper/ExamPaperPreview.vue'); }); };
var QuickPractice = function () { return Promise.resolve().then(function () { return require('../views/QuickPractice.vue'); }); };
var PaperPractice = function () { return Promise.resolve().then(function () { return require('../views/PaperPractice.vue'); }); };
var routes = __spreadArray(__spreadArray([
    {
        path: '/register',
        name: 'Register',
        component: Register_vue_1.default
    },
    {
        path: '/',
        component: Layout_vue_1.default,
        children: [
            {
                path: '',
                name: 'Home',
                component: Home_vue_1.default
            },
            {
                path: 'login',
                name: 'Login',
                component: Login_vue_1.default
            },
            // Teacher Exam Paper Builder route
            {
                path: '/paper/exam/create',
                name: 'ExamPaperBuilder',
                component: ExamPaperForm,
                meta: { requiresAuth: false, requiresTeacher: false }
            },
            // Teacher Exam Paper Management
            {
                path: '/paper/exam/teacher',
                name: 'TeacherExamPaperManagement',
                component: TeacherExamPaperManagement,
                meta: { requiresAuth: true, requiresTeacher: true }
            },
            // Teacher Exam Paper Edit Route
            {
                path: '/paper/exam/edit/:id',
                name: 'ExamPaperEdit',
                component: ExamPaperForm,
                meta: { requiresAuth: true, requiresTeacher: true }
            },
            // Teacher Exam Paper Preview Route
            {
                path: '/paper/exam/preview/:id',
                name: 'ExamPaperPreview',
                component: ExamPaperPreview,
                meta: { requiresAuth: true, requiresTeacher: true }
            },
            // Quick Practice route for students
            {
                path: '/practice/quick',
                name: 'QuickPractice',
                component: QuickPractice
            },
            // Past Paper Practice route for students
            {
                path: '/practice/paper',
                name: 'PaperPractice',
                component: PaperPractice
            },
            // Donation page
            {
                path: '/donate',
                name: 'Donation',
                component: function () { return Promise.resolve().then(function () { return require('../views/Donation.vue'); }); }
            },
            // Help & Documentation page
            {
                path: '/help',
                name: 'Help',
                component: function () { return Promise.resolve().then(function () { return require('../views/Help.vue'); }); }
            }
        ]
    }
], routes_1.default, true), routes_2.default, true);
var router = (0, vue_router_1.createRouter)({
    history: (0, vue_router_1.createWebHistory)(),
    routes: routes
});
// Import user store for navigation guards
var userStore_1 = require("../stores/userStore");
router.beforeEach(function (to, _, next) { return __awaiter(void 0, void 0, void 0, function () {
    var userStore, token, isAuthenticated, userIsAdmin, publicRouteNames, requiresAuth;
    var _a, _b, _c;
    return __generator(this, function (_d) {
        switch (_d.label) {
            case 0:
                userStore = (0, userStore_1.useUserStore)();
                token = localStorage.getItem('authToken');
                if (!(token && !userStore.user)) return [3 /*break*/, 2];
                return [4 /*yield*/, userStore.fetchCurrentUser()];
            case 1:
                _d.sent();
                _d.label = 2;
            case 2:
                isAuthenticated = userStore.isAuthenticated;
                userIsAdmin = ((_a = userStore.user) === null || _a === void 0 ? void 0 : _a.isAdmin) === true;
                publicRouteNames = ['Login', 'Register', 'Setup', 'Home', 'QuickPractice', 'PaperPractice', 'Donation', 'ExamPaperBuilder', 'Help'];
                requiresAuth = !publicRouteNames.includes(to.name);
                if (requiresAuth && !isAuthenticated) {
                    console.log("Navigation guard: Route ".concat(String(to.name), " requires auth, user not authenticated. Redirecting to Login."));
                    return [2 /*return*/, next({ name: 'Login' })];
                }
                if (to.meta.requiresTeacher && !(((_b = userStore.user) === null || _b === void 0 ? void 0 : _b.isTeacher) || ((_c = userStore.user) === null || _c === void 0 ? void 0 : _c.isAdmin))) {
                    return [2 /*return*/, next({ name: 'Home' })];
                }
                if (to.meta.requiresAdmin && !userIsAdmin) {
                    console.log("Navigation guard: Route ".concat(String(to.name), " requires admin, user is not admin. Redirecting to Home."));
                    return [2 /*return*/, next({ name: 'Home' })];
                }
                next();
                return [2 /*return*/];
        }
    });
}); });
exports.default = router;
