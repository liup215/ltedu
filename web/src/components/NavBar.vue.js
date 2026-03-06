"use strict";
/// <reference types="../../../../../../.npm/_npx/2db181330ea4b15b/node_modules/@vue/language-core/types/template-helpers.d.ts" />
/// <reference types="../../../../../../.npm/_npx/2db181330ea4b15b/node_modules/@vue/language-core/types/props-fallback.d.ts" />
var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
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
var _a, _b, _c, _d, _e, _f, _g, _h;
Object.defineProperty(exports, "__esModule", { value: true });
var vue_1 = require("vue");
var userStore_1 = require("../stores/userStore");
var vue_router_1 = require("vue-router");
var config_1 = require("../const/config");
var vue_i18n_1 = require("vue-i18n");
var locale = (0, vue_i18n_1.useI18n)().locale;
var appTitle = (0, vue_1.ref)('LTEDU-Edu'); // Default title, will be updated on mount
var userStore = (0, userStore_1.useUserStore)();
var router = (0, vue_router_1.useRouter)();
var isDropdownOpen = (0, vue_1.ref)(false);
var isMobileMenuOpen = (0, vue_1.ref)(false);
var toggleDropdown = function () {
    isDropdownOpen.value = !isDropdownOpen.value;
};
var closeDropdown = function () {
    isDropdownOpen.value = false;
};
var toggleMobileMenu = function () {
    isMobileMenuOpen.value = !isMobileMenuOpen.value;
    // Close other dropdowns when toggling mobile menu
    isDropdownOpen.value = false;
    langDropdownOpen.value = false;
};
var closeMobileMenu = function () {
    isMobileMenuOpen.value = false;
};
var handleLogout = function () { return __awaiter(void 0, void 0, void 0, function () {
    var error_1;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                _a.trys.push([0, 2, , 3]);
                return [4 /*yield*/, userStore.logout()];
            case 1:
                _a.sent(); // userStore.logout should handle token removal and state reset
                closeDropdown();
                closeMobileMenu();
                router.push('/login');
                return [3 /*break*/, 3];
            case 2:
                error_1 = _a.sent();
                console.error('Logout failed:', error_1);
                return [3 /*break*/, 3];
            case 3: return [2 /*return*/];
        }
    });
}); };
// Exam Paper navigation logic
var handleExamPaperClick = function (type) {
    if (type === 'teacher') {
        router.push('/paper/exam/teacher');
    }
    else {
        router.push('/paper/exam/create');
    }
};
(0, vue_1.onMounted)(function () { return __awaiter(void 0, void 0, void 0, function () {
    var _a;
    return __generator(this, function (_b) {
        switch (_b.label) {
            case 0:
                _a = appTitle;
                return [4 /*yield*/, (0, config_1.APP_TITLE)()];
            case 1:
                _a.value = _b.sent();
                return [2 /*return*/];
        }
    });
}); });
var handleQuickPracticeClick = function () {
    router.push('/practice/quick');
};
var handlePaperPracticeClick = function () {
    router.push('/practice/paper');
};
// Language switcher logic
var languages = [
    { value: 'en', label: 'English' },
    { value: 'zh', label: '中文' },
];
var currentLocale = (0, vue_1.ref)(locale.value);
var langDropdownOpen = (0, vue_1.ref)(false);
var currentLangLabel = (0, vue_1.computed)(function () {
    var found = languages.find(function (l) { return l.value === currentLocale.value; });
    return found ? found.label : '';
});
function toggleLangDropdown() {
    langDropdownOpen.value = !langDropdownOpen.value;
}
function selectLanguage(lang) {
    currentLocale.value = lang;
    locale.value = lang;
    localStorage.setItem('locale', lang);
    langDropdownOpen.value = false;
    location.reload();
}
var __VLS_ctx = __assign(__assign({}, {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
__VLS_asFunctionalElement1(__VLS_intrinsics.nav, __VLS_intrinsics.nav)(__assign({ class: "w-full bg-white border-b border-gray-200 shadow-none py-4 px-8 flex items-center justify-between fixed top-0 left-0 z-50" }));
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['border-b']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-none']} */ ;
/** @type {__VLS_StyleScopedClasses['py-4']} */ ;
/** @type {__VLS_StyleScopedClasses['px-8']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
/** @type {__VLS_StyleScopedClasses['fixed']} */ ;
/** @type {__VLS_StyleScopedClasses['top-0']} */ ;
/** @type {__VLS_StyleScopedClasses['left-0']} */ ;
/** @type {__VLS_StyleScopedClasses['z-50']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-center gap-2" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-xl font-normal text-indigo-700" }));
/** @type {__VLS_StyleScopedClasses['text-xl']} */ ;
/** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
/** @type {__VLS_StyleScopedClasses['text-indigo-700']} */ ;
(__VLS_ctx.appTitle);
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "hidden md:flex gap-4 items-center" }));
/** @type {__VLS_StyleScopedClasses['hidden']} */ ;
/** @type {__VLS_StyleScopedClasses['md:flex']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
var __VLS_0;
/** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
routerLink;
// @ts-ignore
var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0(__assign(__assign({ to: "/" }, { class: "text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" }), { custom: true })));
var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([__assign(__assign({ to: "/" }, { class: "text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" }), { custom: true })], __VLS_functionalComponentArgsRest(__VLS_1), false));
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
/** @type {__VLS_StyleScopedClasses['transition']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-gray-200']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-gray-900']} */ ;
{
    var __VLS_5 = __VLS_3.slots.default;
    var _j = __VLS_vSlot(__VLS_5)[0], navigate = _j.navigate, href = _j.href;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (navigate) }, { href: (href), type: "button" }));
    (__VLS_ctx.$t('navbar.home'));
    // @ts-ignore
    [appTitle, $t,];
    __VLS_3.slots['' /* empty slot name completion */];
}
var __VLS_3;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: (__VLS_ctx.handleQuickPracticeClick) }, { class: "text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-green-100 hover:text-green-900" }), { type: "button" }));
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
/** @type {__VLS_StyleScopedClasses['transition']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-green-100']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-green-900']} */ ;
(__VLS_ctx.$t('navbar.quickPractice'));
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: (__VLS_ctx.handlePaperPracticeClick) }, { class: "text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-purple-100 hover:text-purple-900" }), { type: "button" }));
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
/** @type {__VLS_StyleScopedClasses['transition']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-purple-100']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-purple-900']} */ ;
(__VLS_ctx.$t('navbar.pastPaperPractice'));
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
        var _a = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            _a[_i] = arguments[_i];
        }
        var $event = _a[0];
        __VLS_ctx.handleExamPaperClick('teacher');
        // @ts-ignore
        [$t, $t, handleQuickPracticeClick, handlePaperPracticeClick, handleExamPaperClick,];
    } }, { class: "text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-blue-100 hover:text-blue-900" }), { type: "button" }));
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
/** @type {__VLS_StyleScopedClasses['transition']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-blue-100']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-blue-900']} */ ;
(__VLS_ctx.$t('navbar.myExamPapers'));
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
        var _a = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            _a[_i] = arguments[_i];
        }
        var $event = _a[0];
        __VLS_ctx.handleExamPaperClick('builder');
        // @ts-ignore
        [$t, handleExamPaperClick,];
    } }, { class: "text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-blue-100 hover:text-blue-900" }), { type: "button" }));
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
/** @type {__VLS_StyleScopedClasses['transition']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-blue-100']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-blue-900']} */ ;
(__VLS_ctx.$t('navbar.examPaperBuilder'));
if ((_a = __VLS_ctx.userStore.user) === null || _a === void 0 ? void 0 : _a.isAdmin) {
    var __VLS_6 = void 0;
    /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
    routerLink;
    // @ts-ignore
    var __VLS_7 = __VLS_asFunctionalComponent1(__VLS_6, new __VLS_6(__assign(__assign({ to: "/admin" }, { class: "text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" }), { custom: true })));
    var __VLS_8 = __VLS_7.apply(void 0, __spreadArray([__assign(__assign({ to: "/admin" }, { class: "text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" }), { custom: true })], __VLS_functionalComponentArgsRest(__VLS_7), false));
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
    /** @type {__VLS_StyleScopedClasses['transition']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-200']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-gray-900']} */ ;
    {
        var __VLS_11 = __VLS_9.slots.default;
        var _k = __VLS_vSlot(__VLS_11)[0], navigate = _k.navigate, href = _k.href;
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (navigate) }, { href: (href), type: "button" }));
        (__VLS_ctx.$t('navbar.systemManagement'));
        // @ts-ignore
        [$t, $t, userStore,];
        __VLS_9.slots['' /* empty slot name completion */];
    }
    var __VLS_9;
}
var __VLS_12;
/** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
routerLink;
// @ts-ignore
var __VLS_13 = __VLS_asFunctionalComponent1(__VLS_12, new __VLS_12(__assign(__assign({ to: "/help" }, { class: "text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" }), { custom: true })));
var __VLS_14 = __VLS_13.apply(void 0, __spreadArray([__assign(__assign({ to: "/help" }, { class: "text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" }), { custom: true })], __VLS_functionalComponentArgsRest(__VLS_13), false));
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
/** @type {__VLS_StyleScopedClasses['transition']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-gray-200']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-gray-900']} */ ;
{
    var __VLS_17 = __VLS_15.slots.default;
    var _l = __VLS_vSlot(__VLS_17)[0], navigate = _l.navigate, href = _l.href;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (navigate) }, { href: (href), type: "button" }));
    (__VLS_ctx.$t('navbar.help'));
    // @ts-ignore
    [$t,];
    __VLS_15.slots['' /* empty slot name completion */];
}
var __VLS_15;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
        var _a = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            _a[_i] = arguments[_i];
        }
        var $event = _a[0];
        __VLS_ctx.router.push('/donate');
        // @ts-ignore
        [router,];
    } }, { class: "px-4 py-2 rounded font-normal shadow transition bg-yellow-400 text-white hover:bg-yellow-500 hover:scale-105 border-2 border-yellow-500" }), { type: "button" }));
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow']} */ ;
/** @type {__VLS_StyleScopedClasses['transition']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-yellow-400']} */ ;
/** @type {__VLS_StyleScopedClasses['text-white']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-yellow-500']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:scale-105']} */ ;
/** @type {__VLS_StyleScopedClasses['border-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border-yellow-500']} */ ;
(__VLS_ctx.$t('navbar.donate'));
var __VLS_18;
/** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
routerLink;
// @ts-ignore
var __VLS_19 = __VLS_asFunctionalComponent1(__VLS_18, new __VLS_18(__assign({ to: "/blog" }, { class: "px-4 py-2 rounded font-normal text-gray-700 hover:text-indigo-600 hover:bg-gray-100 transition text-sm" })));
var __VLS_20 = __VLS_19.apply(void 0, __spreadArray([__assign({ to: "/blog" }, { class: "px-4 py-2 rounded font-normal text-gray-700 hover:text-indigo-600 hover:bg-gray-100 transition text-sm" })], __VLS_functionalComponentArgsRest(__VLS_19), false));
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-indigo-600']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
/** @type {__VLS_StyleScopedClasses['transition']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
var __VLS_23 = __VLS_21.slots.default;
(__VLS_ctx.$t('navbar.blog'));
// @ts-ignore
[$t, $t,];
var __VLS_21;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "relative" }));
/** @type {__VLS_StyleScopedClasses['relative']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ onClick: (__VLS_ctx.toggleLangDropdown) }, { class: "flex items-center cursor-pointer select-none text-gray-700 text-sm font-normal hover:text-indigo-600" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['cursor-pointer']} */ ;
/** @type {__VLS_StyleScopedClasses['select-none']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-indigo-600']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "mr-1" }));
/** @type {__VLS_StyleScopedClasses['mr-1']} */ ;
(__VLS_ctx.currentLangLabel);
__VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-4 h-4" }, { fill: "none", stroke: "currentColor", viewBox: "0 0 24 24" }));
/** @type {__VLS_StyleScopedClasses['w-4']} */ ;
/** @type {__VLS_StyleScopedClasses['h-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.path)({
    'stroke-linecap': "round",
    'stroke-linejoin': "round",
    'stroke-width': "2",
    d: "M19 9l-7 7-7-7",
});
if (__VLS_ctx.langDropdownOpen) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.langDropdownOpen))
                return;
            __VLS_ctx.langDropdownOpen = false;
            // @ts-ignore
            [toggleLangDropdown, currentLangLabel, langDropdownOpen, langDropdownOpen,];
        } }, { class: "absolute right-0 mt-2 w-32 bg-white rounded-md shadow-lg py-1 z-50 border border-gray-200" }));
    /** @type {__VLS_StyleScopedClasses['absolute']} */ ;
    /** @type {__VLS_StyleScopedClasses['right-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-32']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['z-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
    var _loop_1 = function (lang) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(__VLS_ctx.langDropdownOpen))
                    return;
                __VLS_ctx.selectLanguage(lang.value);
                // @ts-ignore
                [languages, selectLanguage,];
            } }, { key: (lang.value) }), { class: "block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" }), { class: ({ 'font-bold text-indigo-600': __VLS_ctx.currentLocale === lang.value }) }));
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
        (lang.label);
        // @ts-ignore
        [currentLocale,];
    };
    for (var _i = 0, _m = __VLS_vFor((__VLS_ctx.languages)); _i < _m.length; _i++) {
        var lang = _m[_i][0];
        _loop_1(lang);
    }
}
if (!__VLS_ctx.userStore.isAuthenticated) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-center" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    var __VLS_24 = void 0;
    /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
    routerLink;
    // @ts-ignore
    var __VLS_25 = __VLS_asFunctionalComponent1(__VLS_24, new __VLS_24(__assign(__assign({ to: "/login" }, { class: "text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" }), { custom: true })));
    var __VLS_26 = __VLS_25.apply(void 0, __spreadArray([__assign(__assign({ to: "/login" }, { class: "text-gray-700 px-4 py-2 rounded font-normal transition hover:bg-gray-200 hover:text-gray-900" }), { custom: true })], __VLS_functionalComponentArgsRest(__VLS_25), false));
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
    /** @type {__VLS_StyleScopedClasses['transition']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-200']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-gray-900']} */ ;
    {
        var __VLS_29 = __VLS_27.slots.default;
        var _o = __VLS_vSlot(__VLS_29)[0], navigate = _o.navigate, href = _o.href;
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (navigate) }, { href: (href), type: "button" }));
        (__VLS_ctx.$t('navbar.signIn'));
        // @ts-ignore
        [$t, userStore,];
        __VLS_27.slots['' /* empty slot name completion */];
    }
    var __VLS_27;
}
else {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "relative flex items-center" }));
    /** @type {__VLS_StyleScopedClasses['relative']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.toggleDropdown) }, { class: "flex items-center justify-center w-10 h-10 bg-indigo-600 rounded-full text-white text-lg font-normal focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-10']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-10']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-indigo-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-offset-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
    (((_c = (_b = __VLS_ctx.userStore.user) === null || _b === void 0 ? void 0 : _b.username) === null || _c === void 0 ? void 0 : _c.charAt(0).toUpperCase()) || 'U');
    if (__VLS_ctx.isDropdownOpen) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "absolute right-0 mt-2 top-full w-48 bg-white rounded-md shadow-lg py-1 z-50 border border-gray-200" }));
        /** @type {__VLS_StyleScopedClasses['absolute']} */ ;
        /** @type {__VLS_StyleScopedClasses['right-0']} */ ;
        /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['top-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-48']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['shadow-lg']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['z-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "px-4 py-3" }));
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm text-gray-700" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        (__VLS_ctx.$t('navbar.signedInAs'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm font-normal text-gray-900 truncate" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
        /** @type {__VLS_StyleScopedClasses['truncate']} */ ;
        ((_d = __VLS_ctx.userStore.user) === null || _d === void 0 ? void 0 : _d.username);
        __VLS_asFunctionalElement1(__VLS_intrinsics.hr)(__assign({ class: "border-gray-200" }));
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        var __VLS_30 = void 0;
        /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
        routerLink;
        // @ts-ignore
        var __VLS_31 = __VLS_asFunctionalComponent1(__VLS_30, new __VLS_30(__assign(__assign({ 'onClick': {} }, { to: "/account/profile" }), { class: "block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left" })));
        var __VLS_32 = __VLS_31.apply(void 0, __spreadArray([__assign(__assign({ 'onClick': {} }, { to: "/account/profile" }), { class: "block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left" })], __VLS_functionalComponentArgsRest(__VLS_31), false));
        var __VLS_35 = void 0;
        var __VLS_36 = ({ click: {} },
            { onClick: (__VLS_ctx.closeDropdown) });
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        var __VLS_37 = __VLS_33.slots.default;
        (__VLS_ctx.$t('navbar.yourProfile'));
        // @ts-ignore
        [$t, $t, userStore, userStore, toggleDropdown, isDropdownOpen, closeDropdown,];
        var __VLS_33;
        var __VLS_34;
        var __VLS_38 = void 0;
        /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
        routerLink;
        // @ts-ignore
        var __VLS_39 = __VLS_asFunctionalComponent1(__VLS_38, new __VLS_38(__assign(__assign({ 'onClick': {} }, { to: "/account/settings" }), { class: "block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left" })));
        var __VLS_40 = __VLS_39.apply(void 0, __spreadArray([__assign(__assign({ 'onClick': {} }, { to: "/account/settings" }), { class: "block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left" })], __VLS_functionalComponentArgsRest(__VLS_39), false));
        var __VLS_43 = void 0;
        var __VLS_44 = ({ click: {} },
            { onClick: (__VLS_ctx.closeDropdown) });
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        var __VLS_45 = __VLS_41.slots.default;
        (__VLS_ctx.$t('navbar.settings'));
        // @ts-ignore
        [$t, closeDropdown,];
        var __VLS_41;
        var __VLS_42;
        __VLS_asFunctionalElement1(__VLS_intrinsics.hr)(__assign({ class: "border-gray-200" }));
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.handleLogout) }, { class: "block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" }));
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
        (__VLS_ctx.$t('navbar.signOut'));
    }
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex md:hidden items-center gap-2" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['md:hidden']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "relative" }));
/** @type {__VLS_StyleScopedClasses['relative']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ onClick: (__VLS_ctx.toggleLangDropdown) }, { class: "flex items-center cursor-pointer select-none text-gray-700 text-sm font-normal hover:text-indigo-600" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['cursor-pointer']} */ ;
/** @type {__VLS_StyleScopedClasses['select-none']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-indigo-600']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "mr-1" }));
/** @type {__VLS_StyleScopedClasses['mr-1']} */ ;
(__VLS_ctx.currentLangLabel);
__VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-4 h-4" }, { fill: "none", stroke: "currentColor", viewBox: "0 0 24 24" }));
/** @type {__VLS_StyleScopedClasses['w-4']} */ ;
/** @type {__VLS_StyleScopedClasses['h-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.path)({
    'stroke-linecap': "round",
    'stroke-linejoin': "round",
    'stroke-width': "2",
    d: "M19 9l-7 7-7-7",
});
if (__VLS_ctx.langDropdownOpen) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.langDropdownOpen))
                return;
            __VLS_ctx.langDropdownOpen = false;
            // @ts-ignore
            [$t, toggleLangDropdown, currentLangLabel, langDropdownOpen, langDropdownOpen, handleLogout,];
        } }, { class: "absolute right-0 mt-2 w-32 bg-white rounded-md shadow-lg py-1 z-50 border border-gray-200" }));
    /** @type {__VLS_StyleScopedClasses['absolute']} */ ;
    /** @type {__VLS_StyleScopedClasses['right-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-32']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['z-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
    var _loop_2 = function (lang) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(__VLS_ctx.langDropdownOpen))
                    return;
                __VLS_ctx.selectLanguage(lang.value);
                // @ts-ignore
                [languages, selectLanguage,];
            } }, { key: (lang.value) }), { class: "block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" }), { class: ({ 'font-bold text-indigo-600': __VLS_ctx.currentLocale === lang.value }) }));
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
        (lang.label);
        // @ts-ignore
        [currentLocale,];
    };
    for (var _p = 0, _q = __VLS_vFor((__VLS_ctx.languages)); _p < _q.length; _p++) {
        var lang = _q[_p][0];
        _loop_2(lang);
    }
}
if (__VLS_ctx.userStore.isAuthenticated) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "relative flex items-center" }));
    /** @type {__VLS_StyleScopedClasses['relative']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.toggleDropdown) }, { class: "flex items-center justify-center w-10 h-10 bg-indigo-600 rounded-full text-white text-lg font-normal focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-10']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-10']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-indigo-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-offset-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
    (((_f = (_e = __VLS_ctx.userStore.user) === null || _e === void 0 ? void 0 : _e.username) === null || _f === void 0 ? void 0 : _f.charAt(0).toUpperCase()) || 'U');
    if (__VLS_ctx.isDropdownOpen) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "absolute right-0 mt-2 top-full w-48 bg-white rounded-md shadow-lg py-1 z-50 border border-gray-200" }));
        /** @type {__VLS_StyleScopedClasses['absolute']} */ ;
        /** @type {__VLS_StyleScopedClasses['right-0']} */ ;
        /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['top-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-48']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['shadow-lg']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['z-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "px-4 py-3" }));
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm text-gray-700" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        (__VLS_ctx.$t('navbar.signedInAs'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm font-normal text-gray-900 truncate" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
        /** @type {__VLS_StyleScopedClasses['truncate']} */ ;
        ((_g = __VLS_ctx.userStore.user) === null || _g === void 0 ? void 0 : _g.username);
        __VLS_asFunctionalElement1(__VLS_intrinsics.hr)(__assign({ class: "border-gray-200" }));
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        var __VLS_46 = void 0;
        /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
        routerLink;
        // @ts-ignore
        var __VLS_47 = __VLS_asFunctionalComponent1(__VLS_46, new __VLS_46(__assign(__assign({ 'onClick': {} }, { to: "/account/profile" }), { class: "block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left" })));
        var __VLS_48 = __VLS_47.apply(void 0, __spreadArray([__assign(__assign({ 'onClick': {} }, { to: "/account/profile" }), { class: "block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left" })], __VLS_functionalComponentArgsRest(__VLS_47), false));
        var __VLS_51 = void 0;
        var __VLS_52 = ({ click: {} },
            { onClick: (__VLS_ctx.closeDropdown) });
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        var __VLS_53 = __VLS_49.slots.default;
        (__VLS_ctx.$t('navbar.yourProfile'));
        // @ts-ignore
        [$t, $t, userStore, userStore, userStore, toggleDropdown, isDropdownOpen, closeDropdown,];
        var __VLS_49;
        var __VLS_50;
        var __VLS_54 = void 0;
        /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
        routerLink;
        // @ts-ignore
        var __VLS_55 = __VLS_asFunctionalComponent1(__VLS_54, new __VLS_54(__assign(__assign({ 'onClick': {} }, { to: "/account/settings" }), { class: "block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left" })));
        var __VLS_56 = __VLS_55.apply(void 0, __spreadArray([__assign(__assign({ 'onClick': {} }, { to: "/account/settings" }), { class: "block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left" })], __VLS_functionalComponentArgsRest(__VLS_55), false));
        var __VLS_59 = void 0;
        var __VLS_60 = ({ click: {} },
            { onClick: (__VLS_ctx.closeDropdown) });
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        var __VLS_61 = __VLS_57.slots.default;
        (__VLS_ctx.$t('navbar.settings'));
        // @ts-ignore
        [$t, closeDropdown,];
        var __VLS_57;
        var __VLS_58;
        __VLS_asFunctionalElement1(__VLS_intrinsics.hr)(__assign({ class: "border-gray-200" }));
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.handleLogout) }, { class: "block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" }));
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
        (__VLS_ctx.$t('navbar.signOut'));
    }
}
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: (__VLS_ctx.toggleMobileMenu) }, { class: "inline-flex items-center justify-center w-10 h-10 rounded-md text-gray-700 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-indigo-500" }), { 'aria-label': "Toggle menu" }));
/** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
/** @type {__VLS_StyleScopedClasses['w-10']} */ ;
/** @type {__VLS_StyleScopedClasses['h-10']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
if (!__VLS_ctx.isMobileMenuOpen) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-6 h-6" }, { fill: "none", stroke: "currentColor", viewBox: "0 0 24 24" }));
    /** @type {__VLS_StyleScopedClasses['w-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
        'stroke-linecap': "round",
        'stroke-linejoin': "round",
        'stroke-width': "2",
        d: "M4 6h16M4 12h16M4 18h16",
    });
}
else {
    __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-6 h-6" }, { fill: "none", stroke: "currentColor", viewBox: "0 0 24 24" }));
    /** @type {__VLS_StyleScopedClasses['w-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
        'stroke-linecap': "round",
        'stroke-linejoin': "round",
        'stroke-width': "2",
        d: "M6 18L18 6M6 6l12 12",
    });
}
if (__VLS_ctx.isMobileMenuOpen) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "md:hidden fixed top-16 left-0 right-0 bg-white border-b border-gray-200 shadow-lg z-40 py-2" }));
    /** @type {__VLS_StyleScopedClasses['md:hidden']} */ ;
    /** @type {__VLS_StyleScopedClasses['fixed']} */ ;
    /** @type {__VLS_StyleScopedClasses['top-16']} */ ;
    /** @type {__VLS_StyleScopedClasses['left-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['right-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-b']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['z-40']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex flex-col" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
    var __VLS_62 = void 0;
    /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
    routerLink;
    // @ts-ignore
    var __VLS_63 = __VLS_asFunctionalComponent1(__VLS_62, new __VLS_62(__assign(__assign({ 'onClick': {} }, { to: "/" }), { class: "text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center" })));
    var __VLS_64 = __VLS_63.apply(void 0, __spreadArray([__assign(__assign({ 'onClick': {} }, { to: "/" }), { class: "text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center" })], __VLS_functionalComponentArgsRest(__VLS_63), false));
    var __VLS_67 = void 0;
    var __VLS_68 = ({ click: {} },
        { onClick: (__VLS_ctx.closeMobileMenu) });
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
    /** @type {__VLS_StyleScopedClasses['transition']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
    /** @type {__VLS_StyleScopedClasses['min-h-[48px]']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    var __VLS_69 = __VLS_65.slots.default;
    (__VLS_ctx.$t('navbar.home'));
    // @ts-ignore
    [$t, $t, handleLogout, toggleMobileMenu, isMobileMenuOpen, isMobileMenuOpen, closeMobileMenu,];
    var __VLS_65;
    var __VLS_66;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.isMobileMenuOpen))
                return;
            __VLS_ctx.handleQuickPracticeClick();
            __VLS_ctx.closeMobileMenu();
            // @ts-ignore
            [handleQuickPracticeClick, closeMobileMenu,];
        } }, { class: "text-left text-gray-700 px-6 py-3 font-normal transition hover:bg-green-50 hover:text-green-900 min-h-[48px] flex items-center" }), { type: "button" }));
    /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
    /** @type {__VLS_StyleScopedClasses['transition']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-green-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-green-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['min-h-[48px]']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    (__VLS_ctx.$t('navbar.quickPractice'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.isMobileMenuOpen))
                return;
            __VLS_ctx.handlePaperPracticeClick();
            __VLS_ctx.closeMobileMenu();
            // @ts-ignore
            [$t, handlePaperPracticeClick, closeMobileMenu,];
        } }, { class: "text-left text-gray-700 px-6 py-3 font-normal transition hover:bg-purple-50 hover:text-purple-900 min-h-[48px] flex items-center" }), { type: "button" }));
    /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
    /** @type {__VLS_StyleScopedClasses['transition']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-purple-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-purple-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['min-h-[48px]']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    (__VLS_ctx.$t('navbar.pastPaperPractice'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.isMobileMenuOpen))
                return;
            __VLS_ctx.handleExamPaperClick('teacher');
            __VLS_ctx.closeMobileMenu();
            // @ts-ignore
            [$t, handleExamPaperClick, closeMobileMenu,];
        } }, { class: "text-left text-gray-700 px-6 py-3 font-normal transition hover:bg-blue-50 hover:text-blue-900 min-h-[48px] flex items-center" }), { type: "button" }));
    /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
    /** @type {__VLS_StyleScopedClasses['transition']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-blue-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-blue-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['min-h-[48px]']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    (__VLS_ctx.$t('navbar.myExamPapers'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.isMobileMenuOpen))
                return;
            __VLS_ctx.handleExamPaperClick('builder');
            __VLS_ctx.closeMobileMenu();
            // @ts-ignore
            [$t, handleExamPaperClick, closeMobileMenu,];
        } }, { class: "text-left text-gray-700 px-6 py-3 font-normal transition hover:bg-blue-50 hover:text-blue-900 min-h-[48px] flex items-center" }), { type: "button" }));
    /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
    /** @type {__VLS_StyleScopedClasses['transition']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-blue-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-blue-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['min-h-[48px]']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    (__VLS_ctx.$t('navbar.examPaperBuilder'));
    if ((_h = __VLS_ctx.userStore.user) === null || _h === void 0 ? void 0 : _h.isAdmin) {
        var __VLS_70 = void 0;
        /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
        routerLink;
        // @ts-ignore
        var __VLS_71 = __VLS_asFunctionalComponent1(__VLS_70, new __VLS_70(__assign(__assign({ 'onClick': {} }, { to: "/admin" }), { class: "text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center" })));
        var __VLS_72 = __VLS_71.apply(void 0, __spreadArray([__assign(__assign({ 'onClick': {} }, { to: "/admin" }), { class: "text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center" })], __VLS_functionalComponentArgsRest(__VLS_71), false));
        var __VLS_75 = void 0;
        var __VLS_76 = ({ click: {} },
            { onClick: (__VLS_ctx.closeMobileMenu) });
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
        /** @type {__VLS_StyleScopedClasses['transition']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['min-h-[48px]']} */ ;
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        var __VLS_77 = __VLS_73.slots.default;
        (__VLS_ctx.$t('navbar.systemManagement'));
        // @ts-ignore
        [$t, $t, userStore, closeMobileMenu,];
        var __VLS_73;
        var __VLS_74;
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.isMobileMenuOpen))
                return;
            __VLS_ctx.router.push('/donate');
            __VLS_ctx.closeMobileMenu();
            // @ts-ignore
            [router, closeMobileMenu,];
        } }, { class: "text-left px-6 py-3 font-normal transition text-yellow-700 hover:bg-yellow-50 min-h-[48px] flex items-center" }), { type: "button" }));
    /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
    /** @type {__VLS_StyleScopedClasses['transition']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-yellow-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-yellow-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['min-h-[48px]']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    (__VLS_ctx.$t('navbar.donate'));
    if (!__VLS_ctx.userStore.isAuthenticated) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "border-t border-gray-100 mt-1 pt-1" }));
        /** @type {__VLS_StyleScopedClasses['border-t']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['pt-1']} */ ;
        var __VLS_78 = void 0;
        /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
        routerLink;
        // @ts-ignore
        var __VLS_79 = __VLS_asFunctionalComponent1(__VLS_78, new __VLS_78(__assign(__assign({ 'onClick': {} }, { to: "/login" }), { class: "text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center" })));
        var __VLS_80 = __VLS_79.apply(void 0, __spreadArray([__assign(__assign({ 'onClick': {} }, { to: "/login" }), { class: "text-gray-700 px-6 py-3 font-normal transition hover:bg-gray-100 min-h-[48px] flex items-center" })], __VLS_functionalComponentArgsRest(__VLS_79), false));
        var __VLS_83 = void 0;
        var __VLS_84 = ({ click: {} },
            { onClick: (__VLS_ctx.closeMobileMenu) });
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-normal']} */ ;
        /** @type {__VLS_StyleScopedClasses['transition']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['min-h-[48px]']} */ ;
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        var __VLS_85 = __VLS_81.slots.default;
        (__VLS_ctx.$t('navbar.signIn'));
        // @ts-ignore
        [$t, $t, userStore, closeMobileMenu,];
        var __VLS_81;
        var __VLS_82;
    }
}
// @ts-ignore
[];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
