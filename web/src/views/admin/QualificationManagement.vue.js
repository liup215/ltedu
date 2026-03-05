"use strict";
/// <reference types="../../../../../../../.npm/_npx/2db181330ea4b15b/node_modules/@vue/language-core/types/template-helpers.d.ts" />
/// <reference types="../../../../../../../.npm/_npx/2db181330ea4b15b/node_modules/@vue/language-core/types/props-fallback.d.ts" />
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
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
var vue_1 = require("vue");
var qualificationService_1 = require("../../services/qualificationService");
var organisationService_1 = require("../../services/organisationService");
// Data
var qualifications = (0, vue_1.ref)([]);
var organisations = (0, vue_1.ref)([]);
var loading = (0, vue_1.ref)(true);
var totalQualifications = (0, vue_1.ref)(0);
var currentPage = (0, vue_1.ref)(1);
var pageSize = 10;
var searchQuery = (0, vue_1.ref)('');
var selectedOrganisationId = (0, vue_1.ref)('');
// Fetch qualifications
var fetchQualifications = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_1;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                loading.value = true;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, 4, 5]);
                return [4 /*yield*/, qualificationService_1.default.getQualifications({
                        pageIndex: currentPage.value,
                        pageSize: pageSize,
                        organisationId: selectedOrganisationId.value || undefined,
                    })];
            case 2:
                response = _a.sent();
                qualifications.value = response.data.list;
                totalQualifications.value = response.data.total;
                return [3 /*break*/, 5];
            case 3:
                error_1 = _a.sent();
                console.error('Failed to fetch qualifications:', error_1);
                return [3 /*break*/, 5];
            case 4:
                loading.value = false;
                return [7 /*endfinally*/];
            case 5: return [2 /*return*/];
        }
    });
}); };
// Fetch organisations for filter dropdown
var fetchOrganisations = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_2;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                _a.trys.push([0, 2, , 3]);
                return [4 /*yield*/, organisationService_1.default.getAllOrganisations({})];
            case 1:
                response = _a.sent();
                organisations.value = response.data.list;
                return [3 /*break*/, 3];
            case 2:
                error_2 = _a.sent();
                console.error('Failed to fetch organisations:', error_2);
                return [3 /*break*/, 3];
            case 3: return [2 /*return*/];
        }
    });
}); };
var deleteQualification = function (id) { return __awaiter(void 0, void 0, void 0, function () {
    var error_3;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (!confirm('Are you sure you want to delete this qualification? This action cannot be undone.')) return [3 /*break*/, 4];
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, , 4]);
                return [4 /*yield*/, qualificationService_1.default.deleteQualification(id)];
            case 2:
                _a.sent();
                if (qualifications.value.length === 1 && currentPage.value > 1) {
                    currentPage.value--;
                }
                fetchQualifications(); // Refresh list
                return [3 /*break*/, 4];
            case 3:
                error_3 = _a.sent();
                console.error('Failed to delete qualification:', error_3);
                return [3 /*break*/, 4];
            case 4: return [2 /*return*/];
        }
    });
}); };
// Pagination
var totalPages = (0, vue_1.computed)(function () {
    return Math.ceil(totalQualifications.value / pageSize);
});
var goToPage = function (page) {
    if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
        currentPage.value = page;
        fetchQualifications();
    }
};
var paginationRange = (0, vue_1.computed)(function () {
    var range = [];
    var maxPagesToShow = 5;
    var start = Math.max(1, currentPage.value - Math.floor(maxPagesToShow / 2));
    var end = Math.min(totalPages.value, start + maxPagesToShow - 1);
    if (totalPages.value > 0 && end - start + 1 < maxPagesToShow) {
        if (currentPage.value <= Math.floor(maxPagesToShow / 2)) {
            end = Math.min(totalPages.value, maxPagesToShow);
        }
        else {
            start = Math.max(1, totalPages.value - maxPagesToShow + 1);
        }
    }
    if (totalPages.value > 0 && end - start + 1 < maxPagesToShow && start === 1 && totalPages.value < maxPagesToShow) {
        end = totalPages.value;
    }
    for (var i = start; i <= end; i++) {
        if (i > 0)
            range.push(i);
    }
    return range;
});
// Search debounce
var searchDebounceTimer;
(0, vue_1.watch)([searchQuery, selectedOrganisationId], function () {
    clearTimeout(searchDebounceTimer);
    searchDebounceTimer = window.setTimeout(function () {
        currentPage.value = 1;
        fetchQualifications();
    }, 500);
});
(0, vue_1.onMounted)(function () {
    fetchQualifications();
    fetchOrganisations();
});
var __VLS_ctx = __assign(__assign({}, {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-6" }));
/** @type {__VLS_StyleScopedClasses['p-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.header, __VLS_intrinsics.header)(__assign({ class: "mb-6" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.h1, __VLS_intrinsics.h1)(__assign({ class: "text-3xl font-bold text-gray-900" }));
/** @type {__VLS_StyleScopedClasses['text-3xl']} */ ;
/** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
(__VLS_ctx.$t('qualificationManagement.title'));
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-sm text-gray-500" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
(__VLS_ctx.$t('qualificationManagement.subtitle'));
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0 sm:space-x-4" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:flex-row']} */ ;
/** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:space-y-0']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:space-x-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 w-full sm:w-auto" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:flex-row']} */ ;
/** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:space-y-0']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:space-x-4']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "text", value: (__VLS_ctx.searchQuery), placeholder: (__VLS_ctx.$t('qualificationManagement.searchByName')) }, { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full sm:w-auto" }));
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ value: (__VLS_ctx.selectedOrganisationId) }, { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full sm:w-auto" }));
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "",
});
(__VLS_ctx.$t('qualificationManagement.allOrganisations'));
for (var _i = 0, _b = __VLS_vFor((__VLS_ctx.organisations)); _i < _b.length; _i++) {
    var org = _b[_i][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (org.id),
        value: (org.id),
    });
    (org.name);
    // @ts-ignore
    [$t, $t, $t, $t, searchQuery, selectedOrganisationId, organisations,];
}
var __VLS_0;
/** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
routerLink;
// @ts-ignore
var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0(__assign({ to: "/admin/qualifications/create" }, { class: "inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 w-full sm:w-auto" })));
var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([__assign({ to: "/admin/qualifications/create" }, { class: "inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 w-full sm:w-auto" })], __VLS_functionalComponentArgsRest(__VLS_1), false));
/** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-transparent']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-white']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-indigo-600']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-indigo-700']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-offset-2']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
var __VLS_5 = __VLS_3.slots.default;
(__VLS_ctx.$t('qualificationManagement.addQualification'));
// @ts-ignore
[$t,];
var __VLS_3;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white shadow overflow-x-auto sm:rounded-lg" }));
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow']} */ ;
/** @type {__VLS_StyleScopedClasses['overflow-x-auto']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:rounded-lg']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.table, __VLS_intrinsics.table)(__assign({ class: "min-w-full divide-y divide-gray-200" }));
/** @type {__VLS_StyleScopedClasses['min-w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
/** @type {__VLS_StyleScopedClasses['divide-gray-200']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.thead, __VLS_intrinsics.thead)(__assign({ class: "bg-gray-50" }));
/** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ scope: "col" }, { class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ scope: "col" }, { class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
(__VLS_ctx.$t('qualificationManagement.name'));
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ scope: "col" }, { class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
(__VLS_ctx.$t('qualificationManagement.organisation'));
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ scope: "col" }, { class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider min-w-[120px]" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[120px]']} */ ;
(__VLS_ctx.$t('qualificationManagement.actions'));
__VLS_asFunctionalElement1(__VLS_intrinsics.tbody, __VLS_intrinsics.tbody)(__assign({ class: "bg-white divide-y divide-gray-200" }));
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
/** @type {__VLS_StyleScopedClasses['divide-gray-200']} */ ;
if (__VLS_ctx.loading) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ colspan: "4" }, { class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    (__VLS_ctx.$t('qualificationManagement.loading'));
}
else if (!__VLS_ctx.qualifications || __VLS_ctx.qualifications.length === 0) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ colspan: "4" }, { class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    (__VLS_ctx.$t('qualificationManagement.noData'));
}
var _loop_1 = function (qual) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({
        key: (qual.id),
    });
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (qual.id);
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (qual.name);
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (((_a = qual.organisation) === null || _a === void 0 ? void 0 : _a.name) || '-');
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-right']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-x-2']} */ ;
    var __VLS_6 = void 0;
    /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
    routerLink;
    // @ts-ignore
    var __VLS_7 = __VLS_asFunctionalComponent1(__VLS_6, new __VLS_6(__assign({ to: ("/admin/qualifications/".concat(qual.id, "/edit")) }, { class: "text-indigo-600 hover:text-indigo-900" })));
    var __VLS_8 = __VLS_7.apply(void 0, __spreadArray([__assign({ to: ("/admin/qualifications/".concat(qual.id, "/edit")) }, { class: "text-indigo-600 hover:text-indigo-900" })], __VLS_functionalComponentArgsRest(__VLS_7), false));
    /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-indigo-900']} */ ;
    var __VLS_11 = __VLS_9.slots.default;
    (__VLS_ctx.$t('qualificationManagement.edit'));
    // @ts-ignore
    [$t, $t, $t, $t, $t, $t, loading, qualifications, qualifications, qualifications,];
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            __VLS_ctx.deleteQualification(qual.id);
            // @ts-ignore
            [deleteQualification,];
        } }, { class: "text-red-600 hover:text-red-900" }));
    /** @type {__VLS_StyleScopedClasses['text-red-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-red-900']} */ ;
    (__VLS_ctx.$t('qualificationManagement.delete'));
    // @ts-ignore
    [$t,];
};
var __VLS_9;
for (var _c = 0, _d = __VLS_vFor((__VLS_ctx.qualifications)); _c < _d.length; _c++) {
    var qual = _d[_c][0];
    _loop_1(qual);
}
if (!__VLS_ctx.loading && __VLS_ctx.totalQualifications > 0) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0" }));
    /** @type {__VLS_StyleScopedClasses['mt-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:flex-row']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:space-y-0']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm text-gray-700" }));
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    (__VLS_ctx.$t('qualificationManagement.pageInfo', { from: (__VLS_ctx.currentPage - 1) * __VLS_ctx.pageSize + 1, to: Math.min(__VLS_ctx.currentPage * __VLS_ctx.pageSize, __VLS_ctx.totalQualifications), total: __VLS_ctx.totalQualifications }));
    if (__VLS_ctx.totalPages > 1) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.nav, __VLS_intrinsics.nav)(__assign({ class: "relative z-0 inline-flex rounded-md shadow-sm -space-x-px" }, { 'aria-label': "Pagination" }));
        /** @type {__VLS_StyleScopedClasses['relative']} */ ;
        /** @type {__VLS_StyleScopedClasses['z-0']} */ ;
        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['-space-x-px']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(!__VLS_ctx.loading && __VLS_ctx.totalQualifications > 0))
                    return;
                if (!(__VLS_ctx.totalPages > 1))
                    return;
                __VLS_ctx.goToPage(__VLS_ctx.currentPage - 1);
                // @ts-ignore
                [$t, loading, totalQualifications, totalQualifications, totalQualifications, currentPage, currentPage, currentPage, pageSize, pageSize, totalPages, goToPage,];
            } }, { disabled: (__VLS_ctx.currentPage === 1) }), { class: "relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed" }));
        /** @type {__VLS_StyleScopedClasses['relative']} */ ;
        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-l-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
        (__VLS_ctx.$t('qualificationManagement.previous'));
        var _loop_2 = function (pageNumber) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!(!__VLS_ctx.loading && __VLS_ctx.totalQualifications > 0))
                        return;
                    if (!(__VLS_ctx.totalPages > 1))
                        return;
                    __VLS_ctx.goToPage(pageNumber);
                    // @ts-ignore
                    [$t, currentPage, goToPage, paginationRange,];
                } }, { key: (pageNumber) }), { class: ([
                    'relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium',
                    __VLS_ctx.currentPage === pageNumber ? 'z-10 bg-indigo-50 border-indigo-500 text-indigo-600' : 'bg-white text-gray-700 hover:bg-gray-50'
                ]) }));
            /** @type {__VLS_StyleScopedClasses['relative']} */ ;
            /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
            /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
            /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['border']} */ ;
            /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            (pageNumber);
            // @ts-ignore
            [currentPage,];
        };
        for (var _e = 0, _f = __VLS_vFor((__VLS_ctx.paginationRange)); _e < _f.length; _e++) {
            var pageNumber = _f[_e][0];
            _loop_2(pageNumber);
        }
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(!__VLS_ctx.loading && __VLS_ctx.totalQualifications > 0))
                    return;
                if (!(__VLS_ctx.totalPages > 1))
                    return;
                __VLS_ctx.goToPage(__VLS_ctx.currentPage + 1);
                // @ts-ignore
                [currentPage, goToPage,];
            } }, { disabled: (__VLS_ctx.currentPage === __VLS_ctx.totalPages || __VLS_ctx.totalPages === 0) }), { class: "relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed" }));
        /** @type {__VLS_StyleScopedClasses['relative']} */ ;
        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-r-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
        (__VLS_ctx.$t('qualificationManagement.next'));
    }
}
// @ts-ignore
[$t, currentPage, totalPages, totalPages,];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
