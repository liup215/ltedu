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
var _a, _b, _c, _d, _e, _f;
Object.defineProperty(exports, "__esModule", { value: true });
var vue_1 = require("vue");
var vue_router_1 = require("vue-router");
var paperSeriesService_1 = require("../../services/paperSeriesService");
var syllabusService_1 = require("../../services/syllabusService");
var organisationService_1 = require("../../services/organisationService");
var qualificationService_1 = require("../../services/qualificationService");
var router = (0, vue_router_1.useRouter)();
var paperSeries = (0, vue_1.ref)([]);
var organisations = (0, vue_1.ref)([]);
var qualifications = (0, vue_1.ref)([]);
var syllabi = (0, vue_1.ref)([]);
var loading = (0, vue_1.ref)(false); // Overall loading state for the table
var dropdownLoading = (0, vue_1.ref)({
    organisations: false,
    qualifications: false,
    syllabi: false,
});
var totalItems = (0, vue_1.ref)(0);
var currentPage = (0, vue_1.ref)(1);
var pageSize = 10;
var searchQuery = (0, vue_1.ref)('');
var selectedOrganisationId = (0, vue_1.ref)(0);
var selectedQualificationId = (0, vue_1.ref)(0);
var selectedSyllabusId = (0, vue_1.ref)(0);
// Fetch data for dropdowns
var fetchOrganisationsForDropdown = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_1;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                dropdownLoading.value.organisations = true;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, 4, 5]);
                return [4 /*yield*/, organisationService_1.default.getAllOrganisations({ pageIndex: 1, pageSize: 1000 })];
            case 2:
                response = _a.sent();
                organisations.value = response.data.list;
                return [3 /*break*/, 5];
            case 3:
                error_1 = _a.sent();
                console.error('Failed to fetch organisations:', error_1);
                organisations.value = [];
                return [3 /*break*/, 5];
            case 4:
                dropdownLoading.value.organisations = false;
                return [7 /*endfinally*/];
            case 5: return [2 /*return*/];
        }
    });
}); };
var fetchQualificationsForDropdown = function (organisationId) { return __awaiter(void 0, void 0, void 0, function () {
    var query, response, error_2;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                dropdownLoading.value.qualifications = true;
                qualifications.value = []; // Clear previous before fetching
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, 4, 5]);
                query = { pageIndex: 1, pageSize: 1000 };
                if (organisationId) {
                    query.organisationId = organisationId;
                }
                return [4 /*yield*/, qualificationService_1.default.getAllQualifications(query)];
            case 2:
                response = _a.sent();
                qualifications.value = response.data.list;
                return [3 /*break*/, 5];
            case 3:
                error_2 = _a.sent();
                console.error('Failed to fetch qualifications:', error_2);
                qualifications.value = [];
                return [3 /*break*/, 5];
            case 4:
                dropdownLoading.value.qualifications = false;
                return [7 /*endfinally*/];
            case 5: return [2 /*return*/];
        }
    });
}); };
var fetchSyllabiForDropdown = function (qualificationId) { return __awaiter(void 0, void 0, void 0, function () {
    var query, response, error_3;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                dropdownLoading.value.syllabi = true;
                syllabi.value = []; // Clear previous before fetching
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, 4, 5]);
                query = { pageIndex: 1, pageSize: 1000 };
                // Removed: if (organisationId) { query.organisationId = organisationId }
                if (qualificationId) { // Filter by qual if provided
                    query.qualificationId = qualificationId;
                }
                return [4 /*yield*/, syllabusService_1.default.getAllSyllabuses(query)];
            case 2:
                response = _a.sent();
                syllabi.value = response.data.list;
                return [3 /*break*/, 5];
            case 3:
                error_3 = _a.sent();
                console.error('Failed to fetch syllabi:', error_3);
                syllabi.value = [];
                return [3 /*break*/, 5];
            case 4:
                dropdownLoading.value.syllabi = false;
                return [7 /*endfinally*/];
            case 5: return [2 /*return*/];
        }
    });
}); };
var fetchPaperSeriesList = function () { return __awaiter(void 0, void 0, void 0, function () {
    var query, response, error_4;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                loading.value = true;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, 4, 5]);
                query = {
                    pageIndex: currentPage.value,
                    pageSize: pageSize,
                };
                // if (selectedOrganisationId.value) { // Removed: This parameter is not supported by PaperSeriesQuery
                //   query.organisationId = selectedOrganisationId.value;
                // }
                // if (selectedQualificationId.value) { // Removed: This parameter is not supported by PaperSeriesQuery
                //   query.qualificationId = selectedQualificationId.value;
                // }
                if (selectedSyllabusId.value) {
                    query.syllabusId = selectedSyllabusId.value;
                }
                if (searchQuery.value.trim()) { // Removed: This parameter is not supported by PaperSeriesQuery
                    query.name = searchQuery.value.trim();
                }
                return [4 /*yield*/, paperSeriesService_1.default.getPaperSeriesList(query)];
            case 2:
                response = _a.sent();
                paperSeries.value = response.data.list;
                totalItems.value = response.data.total;
                return [3 /*break*/, 5];
            case 3:
                error_4 = _a.sent();
                console.error('Failed to fetch paper series:', error_4);
                paperSeries.value = [];
                totalItems.value = 0;
                return [3 /*break*/, 5];
            case 4:
                loading.value = false;
                return [7 /*endfinally*/];
            case 5: return [2 /*return*/];
        }
    });
}); };
// Pagination
var totalPages = (0, vue_1.computed)(function () { return Math.ceil(totalItems.value / pageSize); });
var paginationRange = (0, vue_1.computed)(function () {
    var range = [];
    var maxPagesToShow = 5;
    var start = Math.max(1, currentPage.value - Math.floor(maxPagesToShow / 2));
    var end = Math.min(totalPages.value, start + maxPagesToShow - 1);
    if (totalPages.value > maxPagesToShow && end - start + 1 < maxPagesToShow) {
        start = Math.max(1, end - maxPagesToShow + 1);
    }
    for (var i = start; i <= end; i++) {
        range.push(i);
    }
    return range;
});
var goToPage = function (page) {
    if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
        currentPage.value = page;
        fetchPaperSeriesList();
    }
};
// Actions
var goToCreatePage = function () {
    router.push('/admin/paper-series/create');
};
var deletePaperSeries = function (id) { return __awaiter(void 0, void 0, void 0, function () {
    var error_5;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (!confirm('Are you sure you want to delete this paper series? This action cannot be undone.')) return [3 /*break*/, 4];
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, , 4]);
                return [4 /*yield*/, paperSeriesService_1.default.deletePaperSeries(id)];
            case 2:
                _a.sent();
                if (paperSeries.value.length === 1 && currentPage.value > 1) {
                    currentPage.value--;
                }
                fetchPaperSeriesList();
                return [3 /*break*/, 4];
            case 3:
                error_5 = _a.sent();
                console.error('Failed to delete paper series:', error_5);
                return [3 /*break*/, 4];
            case 4: return [2 /*return*/];
        }
    });
}); };
// Watchers
var searchDebounceTimer;
(0, vue_1.watch)(searchQuery, function () {
    clearTimeout(searchDebounceTimer);
    searchDebounceTimer = window.setTimeout(function () {
        currentPage.value = 1;
        // Client-side filtering if API doesn't support name search for PaperSeries
        // For now, we assume the API handles filtering or we re-fetch and rely on pagination
        // If the API doesn't support search by name for PaperSeries, this will just reload the current view
        // based on dropdowns.
        // To implement client-side search on the current page's data:
        // 1. Fetch all data if feasible, or
        // 2. Filter the 'paperSeries.value' array directly if API doesn't support search.
        // The current setup re-calls fetchPaperSeriesList, which doesn't send 'name' in query.
        // This means the search input currently doesn't filter by name via API.
        // If client-side search is desired on the fetched page:
        // You would need a computed property that filters `paperSeries.value` based on `searchQuery.value`.
        // However, this is often combined with server-side pagination/filtering for consistency.
        // For now, leaving as is, which means search input doesn't filter server-side.
        fetchPaperSeriesList();
    }, 500);
});
(0, vue_1.watch)(selectedOrganisationId, function (newOrgId) { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                selectedQualificationId.value = 0;
                selectedSyllabusId.value = 0;
                qualifications.value = [];
                syllabi.value = [];
                currentPage.value = 1;
                return [4 /*yield*/, fetchQualificationsForDropdown(newOrgId || undefined)
                    // When Org changes, fetch syllabi for that org (or all if "All Orgs")
                    // If "All Orgs", Quals will be all, Syllabi will be all.
                    // If specific Org, Quals for that Org, Syllabi for that Org (indirectly via Quals).
                    // The fetchSyllabiForDropdown needs to be smart or be called after Quals are set.
                    // Current logic: fetchSyllabiForDropdown is called with newOrgId and undefined QualId.
                    // This might not be what we want if we want syllabi filtered by the new org directly.
                    // However, PaperSeriesQuery only takes syllabusId. So dropdowns are for narrowing down syllabusId.
                ];
            case 1:
                _a.sent();
                // When Org changes, fetch syllabi for that org (or all if "All Orgs")
                // If "All Orgs", Quals will be all, Syllabi will be all.
                // If specific Org, Quals for that Org, Syllabi for that Org (indirectly via Quals).
                // The fetchSyllabiForDropdown needs to be smart or be called after Quals are set.
                // Current logic: fetchSyllabiForDropdown is called with newOrgId and undefined QualId.
                // This might not be what we want if we want syllabi filtered by the new org directly.
                // However, PaperSeriesQuery only takes syllabusId. So dropdowns are for narrowing down syllabusId.
                return [4 /*yield*/, fetchSyllabiForDropdown(undefined)];
            case 2:
                // When Org changes, fetch syllabi for that org (or all if "All Orgs")
                // If "All Orgs", Quals will be all, Syllabi will be all.
                // If specific Org, Quals for that Org, Syllabi for that Org (indirectly via Quals).
                // The fetchSyllabiForDropdown needs to be smart or be called after Quals are set.
                // Current logic: fetchSyllabiForDropdown is called with newOrgId and undefined QualId.
                // This might not be what we want if we want syllabi filtered by the new org directly.
                // However, PaperSeriesQuery only takes syllabusId. So dropdowns are for narrowing down syllabusId.
                _a.sent();
                fetchPaperSeriesList();
                return [2 /*return*/];
        }
    });
}); });
(0, vue_1.watch)(selectedQualificationId, function (newQualId) { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                selectedSyllabusId.value = 0;
                syllabi.value = [];
                currentPage.value = 1;
                return [4 /*yield*/, fetchSyllabiForDropdown(newQualId || undefined)];
            case 1:
                _a.sent();
                fetchPaperSeriesList();
                return [2 /*return*/];
        }
    });
}); });
(0, vue_1.watch)(selectedSyllabusId, function () {
    currentPage.value = 1;
    fetchPaperSeriesList();
});
// Lifecycle
var fetchInitialDropdownData = function () { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, fetchOrganisationsForDropdown()
                // Fetch all qualifications and syllabi initially if no org/qual is pre-selected
            ];
            case 1:
                _a.sent();
                // Fetch all qualifications and syllabi initially if no org/qual is pre-selected
                return [4 /*yield*/, fetchQualificationsForDropdown()];
            case 2:
                // Fetch all qualifications and syllabi initially if no org/qual is pre-selected
                _a.sent();
                return [4 /*yield*/, fetchSyllabiForDropdown()];
            case 3:
                _a.sent();
                return [2 /*return*/];
        }
    });
}); };
(0, vue_1.onMounted)(function () { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                loading.value = true;
                dropdownLoading.value.organisations = true;
                dropdownLoading.value.qualifications = true;
                dropdownLoading.value.syllabi = true;
                return [4 /*yield*/, fetchInitialDropdownData()];
            case 1:
                _a.sent();
                return [4 /*yield*/, fetchPaperSeriesList()
                    // Set loading to false after all initial data is fetched
                    // loading.value = false; // This was here, but fetchPaperSeriesList has its own finally block.
                    // Dropdown loading should be false after their respective fetches.
                ];
            case 2:
                _a.sent();
                return [2 /*return*/];
        }
    });
}); });
(0, vue_1.onUnmounted)(function () {
    clearTimeout(searchDebounceTimer);
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
(__VLS_ctx.$t('paperSeries.title'));
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-sm text-gray-500" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
(__VLS_ctx.$t('paperSeries.subtitle') || '');
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
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ value: (__VLS_ctx.selectedOrganisationId) }, { class: "px-3 py-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-base w-full sm:w-auto" }), { disabled: (__VLS_ctx.dropdownLoading.organisations) }));
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['text-base']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (0),
});
(__VLS_ctx.$t('syllabusManagement.allOrganisations'));
for (var _i = 0, _g = __VLS_vFor((__VLS_ctx.organisations)); _i < _g.length; _i++) {
    var org = _g[_i][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (org.id),
        value: (org.id),
    });
    (org.name);
    // @ts-ignore
    [$t, $t, $t, selectedOrganisationId, dropdownLoading, organisations,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ value: (__VLS_ctx.selectedQualificationId) }, { class: "px-3 py-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-base w-full sm:w-auto" }), { disabled: (__VLS_ctx.dropdownLoading.qualifications || !__VLS_ctx.selectedOrganisationId) }));
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['text-base']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (0),
});
(__VLS_ctx.$t('syllabusManagement.allQualifications'));
for (var _h = 0, _j = __VLS_vFor((__VLS_ctx.qualifications)); _h < _j.length; _h++) {
    var q = _j[_h][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (q.id),
        value: (q.id),
    });
    (q.name);
    // @ts-ignore
    [$t, selectedOrganisationId, dropdownLoading, selectedQualificationId, qualifications,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ value: (__VLS_ctx.selectedSyllabusId) }, { class: "px-3 py-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-base w-full sm:w-auto" }), { disabled: (__VLS_ctx.dropdownLoading.syllabi || !__VLS_ctx.selectedQualificationId) }));
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['text-base']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (0),
});
(__VLS_ctx.$t('syllabusManagement.addSyllabus'));
for (var _k = 0, _l = __VLS_vFor((__VLS_ctx.syllabi)); _k < _l.length; _k++) {
    var syllabus = _l[_k][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (syllabus.id),
        value: (syllabus.id),
    });
    (syllabus.name);
    // @ts-ignore
    [$t, dropdownLoading, selectedQualificationId, selectedSyllabusId, syllabi,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "text", value: (__VLS_ctx.searchQuery), placeholder: (__VLS_ctx.$t('paperSeries.searchByName')) }, { class: "px-3 py-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 text-base w-full sm:w-auto" }));
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['text-base']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.goToCreatePage) }, { class: "inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 w-full sm:w-auto" }));
/** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-transparent']} */ ;
/** @type {__VLS_StyleScopedClasses['text-base']} */ ;
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
(__VLS_ctx.$t('paperSeries.add'));
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
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
(__VLS_ctx.$t('paperSeries.name'));
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
(__VLS_ctx.$t('syllabusManagement.name'));
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
(__VLS_ctx.$t('qualificationManagement.name'));
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
(__VLS_ctx.$t('organisationManagement.name'));
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider min-w-[180px]" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[180px]']} */ ;
(__VLS_ctx.$t('paperSeries.actions'));
__VLS_asFunctionalElement1(__VLS_intrinsics.tbody, __VLS_intrinsics.tbody)(__assign({ class: "bg-white divide-y divide-gray-200" }));
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
/** @type {__VLS_StyleScopedClasses['divide-gray-200']} */ ;
if (__VLS_ctx.loading) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ colspan: "6" }, { class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    (__VLS_ctx.$t('paperSeries.loading'));
}
else if (!__VLS_ctx.paperSeries.length) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ colspan: "6" }, { class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    (__VLS_ctx.$t('paperSeries.noData'));
}
var _loop_1 = function (series) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({
        key: (series.id),
    });
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (series.id);
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (series.name);
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (((_a = series.syllabus) === null || _a === void 0 ? void 0 : _a.name) || '-');
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (((_c = (_b = series.syllabus) === null || _b === void 0 ? void 0 : _b.qualification) === null || _c === void 0 ? void 0 : _c.name) || '-');
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (((_f = (_e = (_d = series.syllabus) === null || _d === void 0 ? void 0 : _d.qualification) === null || _e === void 0 ? void 0 : _e.organisation) === null || _f === void 0 ? void 0 : _f.name) || '-');
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm font-medium space-x-3" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-x-3']} */ ;
    var __VLS_0 = void 0;
    /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
    routerLink;
    // @ts-ignore
    var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0(__assign({ to: ("/admin/paper-series/".concat(series.id, "/edit")) }, { class: "text-indigo-600 hover:text-indigo-900" })));
    var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([__assign({ to: ("/admin/paper-series/".concat(series.id, "/edit")) }, { class: "text-indigo-600 hover:text-indigo-900" })], __VLS_functionalComponentArgsRest(__VLS_1), false));
    /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-indigo-900']} */ ;
    var __VLS_5 = __VLS_3.slots.default;
    (__VLS_ctx.$t('paperSeries.edit'));
    // @ts-ignore
    [$t, $t, $t, $t, $t, $t, $t, $t, $t, $t, searchQuery, goToCreatePage, loading, paperSeries, paperSeries,];
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            __VLS_ctx.deletePaperSeries(series.id);
            // @ts-ignore
            [deletePaperSeries,];
        } }, { class: "text-red-600 hover:text-red-900" }));
    /** @type {__VLS_StyleScopedClasses['text-red-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-red-900']} */ ;
    (__VLS_ctx.$t('paperSeries.delete'));
    // @ts-ignore
    [$t,];
};
var __VLS_3;
for (var _m = 0, _o = __VLS_vFor((__VLS_ctx.paperSeries)); _m < _o.length; _m++) {
    var series = _o[_m][0];
    _loop_1(series);
}
if (!__VLS_ctx.loading && __VLS_ctx.totalItems > 0) {
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
    (__VLS_ctx.$t('paperSeries.pageInfo', { from: (__VLS_ctx.currentPage - 1) * __VLS_ctx.pageSize + 1, to: Math.min(__VLS_ctx.currentPage * __VLS_ctx.pageSize, __VLS_ctx.totalItems), total: __VLS_ctx.totalItems }));
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
                if (!(!__VLS_ctx.loading && __VLS_ctx.totalItems > 0))
                    return;
                if (!(__VLS_ctx.totalPages > 1))
                    return;
                __VLS_ctx.goToPage(__VLS_ctx.currentPage - 1);
                // @ts-ignore
                [$t, loading, totalItems, totalItems, totalItems, currentPage, currentPage, currentPage, pageSize, pageSize, totalPages, goToPage,];
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
        (__VLS_ctx.$t('syllabusManagement.previous'));
        var _loop_2 = function (pageNumber) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!(!__VLS_ctx.loading && __VLS_ctx.totalItems > 0))
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
        for (var _p = 0, _q = __VLS_vFor((__VLS_ctx.paginationRange)); _p < _q.length; _p++) {
            var pageNumber = _q[_p][0];
            _loop_2(pageNumber);
        }
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(!__VLS_ctx.loading && __VLS_ctx.totalItems > 0))
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
        (__VLS_ctx.$t('syllabusManagement.next'));
    }
}
// @ts-ignore
[$t, currentPage, totalPages, totalPages,];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
