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
var _a, _b, _c, _d, _e, _f, _g, _h, _j;
Object.defineProperty(exports, "__esModule", { value: true });
var vue_1 = require("vue");
var questionService_1 = require("../../services/questionService");
var organisationService_1 = require("../../services/organisationService");
var qualificationService_1 = require("../../services/qualificationService");
var syllabusService_1 = require("../../services/syllabusService");
var question_model_1 = require("../../models/question.model");
var index_vue_1 = require("../../components/QuillEditor/index.vue");
// Reactive data
var questions = (0, vue_1.ref)([]);
var loading = (0, vue_1.ref)(true);
var activeTabMap = (0, vue_1.ref)(new Map()); // Map to store active tab index for each question
var totalQuestions = (0, vue_1.ref)(0);
var currentPage = (0, vue_1.ref)(1);
var pageSize = 12; // Cards per page
var searchQuery = (0, vue_1.ref)('');
var paperNameQuery = (0, vue_1.ref)('');
var selectedDifficulty = (0, vue_1.ref)('');
var selectedStatus = (0, vue_1.ref)('');
// Filter state
var organisations = (0, vue_1.ref)([]);
var qualifications = (0, vue_1.ref)([]);
var syllabi = (0, vue_1.ref)([]);
var selectedOrganisationId = (0, vue_1.ref)(null);
var selectedQualificationId = (0, vue_1.ref)(null);
var selectedSyllabusId = (0, vue_1.ref)(null);
// Computed properties
var totalPages = (0, vue_1.computed)(function () {
    return Math.ceil(totalQuestions.value / pageSize);
});
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
    for (var i = start; i <= end; i++) {
        if (i > 0)
            range.push(i);
    }
    return range;
});
// Methods for cascading dropdowns
var fetchOrganisations = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_1;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                _a.trys.push([0, 2, , 3]);
                return [4 /*yield*/, organisationService_1.default.getOrganisations({ pageIndex: 1, pageSize: 1000 })];
            case 1:
                response = _a.sent();
                organisations.value = response.data.list;
                return [3 /*break*/, 3];
            case 2:
                error_1 = _a.sent();
                console.error('Failed to fetch organisations:', error_1);
                organisations.value = [];
                return [3 /*break*/, 3];
            case 3: return [2 /*return*/];
        }
    });
}); };
var fetchQualifications = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_2;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (!selectedOrganisationId.value) {
                    qualifications.value = [];
                    syllabi.value = [];
                    selectedQualificationId.value = null;
                    selectedSyllabusId.value = null;
                    return [2 /*return*/];
                }
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, , 4]);
                return [4 /*yield*/, qualificationService_1.default.getQualifications({
                        pageIndex: 1,
                        pageSize: 1000,
                        organisationId: Number(selectedOrganisationId.value)
                    })];
            case 2:
                response = _a.sent();
                qualifications.value = response.data.list;
                return [3 /*break*/, 4];
            case 3:
                error_2 = _a.sent();
                console.error('Failed to fetch qualifications:', error_2);
                qualifications.value = [];
                return [3 /*break*/, 4];
            case 4:
                syllabi.value = [];
                selectedSyllabusId.value = null;
                return [2 /*return*/];
        }
    });
}); };
var fetchSyllabi = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_3;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (!selectedQualificationId.value) {
                    syllabi.value = [];
                    selectedSyllabusId.value = null;
                    return [2 /*return*/];
                }
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, , 4]);
                return [4 /*yield*/, syllabusService_1.default.getSyllabuses({
                        pageIndex: 1,
                        pageSize: 1000,
                        qualificationId: Number(selectedQualificationId.value)
                    })];
            case 2:
                response = _a.sent();
                syllabi.value = response.data.list;
                return [3 /*break*/, 4];
            case 3:
                error_3 = _a.sent();
                console.error('Failed to fetch syllabi:', error_3);
                syllabi.value = [];
                return [3 /*break*/, 4];
            case 4:
                selectedSyllabusId.value = null; // Reset syllabus selection when new list is fetched
                return [2 /*return*/];
        }
    });
}); };
// Methods
var fetchQuestions = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_4;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                loading.value = true;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, 4, 5]);
                return [4 /*yield*/, questionService_1.default.getQuestions({
                        pageIndex: currentPage.value,
                        pageSize: pageSize,
                        syllabusId: selectedSyllabusId.value ? Number(selectedSyllabusId.value) : undefined,
                        difficult: selectedDifficulty.value ? Number(selectedDifficulty.value) : undefined,
                        status: selectedStatus.value ? Number(selectedStatus.value) : undefined,
                        stem: searchQuery.value.trim() || undefined,
                        paperName: paperNameQuery.value.trim() || undefined
                    })];
            case 2:
                response = _a.sent();
                questions.value = response.data.list;
                totalQuestions.value = response.data.total;
                return [3 /*break*/, 5];
            case 3:
                error_4 = _a.sent();
                console.error('Failed to fetch questions:', error_4);
                return [3 /*break*/, 5];
            case 4:
                loading.value = false;
                return [7 /*endfinally*/];
            case 5: return [2 /*return*/];
        }
    });
}); };
var deleteQuestion = function (id) { return __awaiter(void 0, void 0, void 0, function () {
    var error_5;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (!confirm('Are you sure you want to delete this question? This action cannot be undone.')) return [3 /*break*/, 4];
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, , 4]);
                return [4 /*yield*/, questionService_1.default.deleteQuestion(id)];
            case 2:
                _a.sent();
                if (questions.value.length === 1 && currentPage.value > 1) {
                    currentPage.value--;
                }
                fetchQuestions();
                return [3 /*break*/, 4];
            case 3:
                error_5 = _a.sent();
                console.error('Failed to delete question:', error_5);
                return [3 /*break*/, 4];
            case 4: return [2 /*return*/];
        }
    });
}); };
var goToPage = function (page) {
    if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
        currentPage.value = page;
        fetchQuestions();
    }
};
// Utility methods
var getDifficultyName = function (difficult) {
    return question_model_1.DIFFICULTY_NAMES[difficult] || 'Unknown';
};
var getDifficultyClass = function (difficult) {
    var classes = {
        1: 'bg-green-100 text-green-800',
        2: 'bg-yellow-100 text-yellow-800',
        3: 'bg-orange-100 text-orange-800',
        4: 'bg-red-100 text-red-800',
        5: 'bg-purple-100 text-purple-800'
    };
    return classes[difficult] || 'bg-gray-100 text-gray-800';
};
var getStatusName = function (status) {
    return question_model_1.QUESTION_STATUS_NAMES[status] || 'Unknown';
};
var getStatusClass = function (status) {
    if (status === question_model_1.QUESTION_STATE_NORMAL) {
        return 'bg-green-100 text-green-800';
    }
    else if (status === question_model_1.QUESTION_STATE_FORBIDDEN) {
        return 'bg-red-100 text-red-800';
    }
    return 'bg-gray-100 text-gray-800';
};
var formatDate = function (dateString) {
    if (!dateString)
        return 'N/A';
    return new Date(dateString).toLocaleDateString();
};
// Search debounce
var searchDebounceTimer;
(0, vue_1.watch)([searchQuery, paperNameQuery, selectedSyllabusId, selectedDifficulty, selectedStatus], function () {
    clearTimeout(searchDebounceTimer);
    searchDebounceTimer = window.setTimeout(function () {
        currentPage.value = 1;
        fetchQuestions();
    }, 500);
});
(0, vue_1.watch)(selectedSyllabusId, function (newValue) {
    if (!newValue) { // If syllabus is cleared
        questions.value = [];
        totalQuestions.value = 0;
    }
});
(0, vue_1.watch)(selectedOrganisationId, function (newValue) {
    selectedQualificationId.value = null;
    selectedSyllabusId.value = null;
    qualifications.value = [];
    syllabi.value = [];
    if (newValue) {
        fetchQualifications();
        // Don't fetch questions here, let the main filter watcher do it if syllabus changes
        // or if no syllabus is selected, questions should be empty or for the whole org (not implemented yet)
        // For now, if org is selected but no syllabus, clear questions.
        questions.value = []; // Clear questions as syllabus is now reset
        totalQuestions.value = 0;
    }
    else {
        // Organisation filter cleared, fetch all questions (or based on other active filters)
        fetchQuestions();
    }
});
(0, vue_1.watch)(selectedQualificationId, function (newValue) {
    selectedSyllabusId.value = null;
    syllabi.value = [];
    if (newValue) {
        fetchSyllabi();
        // Similar to above, clear questions as syllabus is now reset
        questions.value = [];
        totalQuestions.value = 0;
    }
    else if (selectedOrganisationId.value) {
        // Qualification cleared, but organisation is still selected.
        // Fetch questions for the organisation (if such a feature is desired) or clear.
        // For now, clear questions as no specific syllabus path is selected.
        questions.value = [];
        totalQuestions.value = 0;
        // Potentially fetch questions for the selectedOrganisationId if that's a requirement
        // fetchQuestions(); // This would fetch based on current filters (likely no syllabusId)
    }
    // If selectedQualificationId is cleared and selectedOrganisationId is also cleared,
    // the selectedOrganisationId watcher would have already triggered a fetchQuestions for all.
});
// The main watch for [searchQuery, selectedSyllabusId, ...] handles fetching when selectedSyllabusId is set.
(0, vue_1.onMounted)(function () {
    fetchOrganisations(); // Load organisations first
    fetchQuestions(); // Then load initial questions (likely all, as filters are empty)
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
(__VLS_ctx.$t('question.title'));
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-sm text-gray-500" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
(__VLS_ctx.$t('question.noData'));
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-6 flex flex-col space-y-4" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
/** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex flex-wrap gap-4" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-wrap']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "text", value: (__VLS_ctx.searchQuery), placeholder: (__VLS_ctx.$t('question.searchByStem')) }, { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_260px] min-w-[200px] max-w-lg" }));
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
/** @type {__VLS_StyleScopedClasses['flex-[1_1_260px]']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[200px]']} */ ;
/** @type {__VLS_StyleScopedClasses['max-w-lg']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "text", value: (__VLS_ctx.paperNameQuery), placeholder: (__VLS_ctx.$t('question.searchByPaperName')) }, { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_260px] min-w-[200px] max-w-lg" }));
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
/** @type {__VLS_StyleScopedClasses['flex-[1_1_260px]']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[200px]']} */ ;
/** @type {__VLS_StyleScopedClasses['max-w-lg']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ value: (__VLS_ctx.selectedOrganisationId) }, { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_220px] min-w-[180px] max-w-md" }));
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
/** @type {__VLS_StyleScopedClasses['flex-[1_1_220px]']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[180px]']} */ ;
/** @type {__VLS_StyleScopedClasses['max-w-md']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (null),
});
(__VLS_ctx.$t('question.selectOrganisation'));
for (var _i = 0, _k = __VLS_vFor((__VLS_ctx.organisations)); _i < _k.length; _i++) {
    var org = _k[_i][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (org.id),
        value: (org.id),
    });
    (org.name);
    // @ts-ignore
    [$t, $t, $t, $t, $t, searchQuery, paperNameQuery, selectedOrganisationId, organisations,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ value: (__VLS_ctx.selectedQualificationId) }, { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_220px] min-w-[180px] max-w-md" }));
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
/** @type {__VLS_StyleScopedClasses['flex-[1_1_220px]']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[180px]']} */ ;
/** @type {__VLS_StyleScopedClasses['max-w-md']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (null),
});
(__VLS_ctx.$t('question.selectQualification'));
for (var _l = 0, _m = __VLS_vFor((__VLS_ctx.qualifications)); _l < _m.length; _l++) {
    var qual = _m[_l][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (qual.id),
        value: (qual.id),
    });
    (qual.name);
    // @ts-ignore
    [$t, selectedQualificationId, qualifications,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ value: (__VLS_ctx.selectedSyllabusId) }, { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_220px] min-w-[180px] max-w-md" }));
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
/** @type {__VLS_StyleScopedClasses['flex-[1_1_220px]']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[180px]']} */ ;
/** @type {__VLS_StyleScopedClasses['max-w-md']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (null),
});
(__VLS_ctx.$t('question.selectSyllabus'));
for (var _o = 0, _p = __VLS_vFor((__VLS_ctx.syllabi)); _o < _p.length; _o++) {
    var syllabus = _p[_o][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (syllabus.id),
        value: (syllabus.id),
    });
    (syllabus.name);
    // @ts-ignore
    [$t, selectedSyllabusId, syllabi,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ value: (__VLS_ctx.selectedDifficulty) }, { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_180px] min-w-[140px] max-w-md" }));
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
/** @type {__VLS_StyleScopedClasses['flex-[1_1_180px]']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[140px]']} */ ;
/** @type {__VLS_StyleScopedClasses['max-w-md']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "",
});
(__VLS_ctx.$t('examPaperForm.allDifficulty'));
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "1",
});
(__VLS_ctx.$t('examPaperForm.easy'));
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "2",
});
(__VLS_ctx.$t('examPaperForm.medium'));
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "3",
});
(__VLS_ctx.$t('examPaperForm.hard'));
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "4",
});
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "5",
});
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ value: (__VLS_ctx.selectedStatus) }, { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm flex-[1_1_180px] min-w-[140px] max-w-md" }));
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
/** @type {__VLS_StyleScopedClasses['flex-[1_1_180px]']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[140px]']} */ ;
/** @type {__VLS_StyleScopedClasses['max-w-md']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "",
});
(__VLS_ctx.$t('common.actions'));
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "1",
});
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "2",
});
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-between items-center" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
        var _a = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            _a[_i] = arguments[_i];
        }
        var $event = _a[0];
        __VLS_ctx.$router.push('/admin/questions/create');
        // @ts-ignore
        [$t, $t, $t, $t, $t, selectedDifficulty, selectedStatus, $router,];
    } }, { class: "px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition duration-150" }));
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-blue-600']} */ ;
/** @type {__VLS_StyleScopedClasses['text-white']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-blue-700']} */ ;
/** @type {__VLS_StyleScopedClasses['transition']} */ ;
/** @type {__VLS_StyleScopedClasses['duration-150']} */ ;
(__VLS_ctx.$t('question.add'));
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-600" }));
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
(__VLS_ctx.$t('question.pageInfo', { from: ((__VLS_ctx.currentPage - 1) * __VLS_ctx.pageSize) + 1, to: Math.min(__VLS_ctx.currentPage * __VLS_ctx.pageSize, __VLS_ctx.totalQuestions), total: __VLS_ctx.totalQuestions }));
if (__VLS_ctx.loading) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-center items-center py-12" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-12']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600" }));
    /** @type {__VLS_StyleScopedClasses['animate-spin']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-8']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-8']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-b-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-indigo-600']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "ml-3 text-gray-600" }));
    /** @type {__VLS_StyleScopedClasses['ml-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
}
else if (!__VLS_ctx.questions || __VLS_ctx.questions.length === 0) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-12" }));
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-12']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "mx-auto h-12 w-12 text-gray-400" }, { fill: "none", viewBox: "0 0 24 24", stroke: "currentColor" }));
    /** @type {__VLS_StyleScopedClasses['mx-auto']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-12']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-12']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
        'stroke-linecap': "round",
        'stroke-linejoin': "round",
        'stroke-width': "2",
        d: "M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z",
    });
    __VLS_asFunctionalElement1(__VLS_intrinsics.h3, __VLS_intrinsics.h3)(__assign({ class: "mt-2 text-sm font-medium text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (__VLS_ctx.$t('question.noData'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    (__VLS_ctx.$t('question.add'));
}
else {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "grid grid-cols-1 gap-6" }));
    /** @type {__VLS_StyleScopedClasses['grid']} */ ;
    /** @type {__VLS_StyleScopedClasses['grid-cols-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['gap-6']} */ ;
    var _loop_1 = function (question) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (question.id) }, { class: "bg-white rounded-lg shadow-md border border-gray-200 hover:shadow-lg transition-shadow duration-200" }));
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
        /** @type {__VLS_StyleScopedClasses['shadow-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:shadow-lg']} */ ;
        /** @type {__VLS_StyleScopedClasses['transition-shadow']} */ ;
        /** @type {__VLS_StyleScopedClasses['duration-200']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-4 border-b border-gray-200" }));
        /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-b']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-between items-start" }));
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-start']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex-1" }));
        /** @type {__VLS_StyleScopedClasses['flex-1']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-center space-x-2 mb-2" }));
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['space-x-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800" }));
        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-blue-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-blue-800']} */ ;
        (__VLS_ctx.$t('common.id'));
        (question.id);
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" }, { class: (__VLS_ctx.getDifficultyClass(question.difficult)) }));
        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        (__VLS_ctx.$t('question.difficulty'));
        (__VLS_ctx.getDifficultyName(question.difficult));
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" }, { class: (__VLS_ctx.getStatusClass(question.status)) }));
        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        (__VLS_ctx.getStatusName(question.status));
        __VLS_asFunctionalElement1(__VLS_intrinsics.h3, __VLS_intrinsics.h3)(__assign({ class: "text-sm font-medium text-gray-900 mb-1" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
        /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
        (((_a = question.syllabus) === null || _a === void 0 ? void 0 : _a.name) || __VLS_ctx.$t('examPaperForm.unknown'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-xs text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (__VLS_ctx.$t('examPaperForm.totalScore'));
        (question.totalScore || __VLS_ctx.$t('examPaperForm.unknown'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-4" }));
        /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4" }));
        /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
        var __VLS_0 = index_vue_1.default || index_vue_1.default;
        // @ts-ignore
        var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0({
            modelValue: (question.stem),
            readOnly: true,
            height: "100%",
        }));
        var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([{
                modelValue: (question.stem),
                readOnly: true,
                height: "100%",
            }], __VLS_functionalComponentArgsRest(__VLS_1), false));
        if (question.questionContents && question.questionContents.length > 0) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4" }));
            /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.h4, __VLS_intrinsics.h4)(__assign({ class: "text-sm font-medium text-gray-700 mb-2" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
            /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
            (__VLS_ctx.$t('examPaperForm.parts'));
            (__VLS_ctx.$t('examPaperForm.totalScore'));
            (question.totalScore);
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "border border-gray-300 rounded-lg overflow-hidden" }));
            /** @type {__VLS_StyleScopedClasses['border']} */ ;
            /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
            /** @type {__VLS_StyleScopedClasses['overflow-hidden']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex overflow-x-auto bg-gray-50 border-b border-gray-300" }));
            /** @type {__VLS_StyleScopedClasses['flex']} */ ;
            /** @type {__VLS_StyleScopedClasses['overflow-x-auto']} */ ;
            /** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
            /** @type {__VLS_StyleScopedClasses['border-b']} */ ;
            /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
            var _loop_3 = function (content, index) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: (function () { return __VLS_ctx.activeTabMap.set(question.id, index); }) }, { key: (index) }), { class: ([
                        'px-4 py-2 text-sm font-medium border-r border-gray-200 whitespace-nowrap transition',
                        __VLS_ctx.activeTabMap.get(question.id) === index
                            ? 'text-blue-600 bg-white border-b-2 border-b-blue-500 -mb-px'
                            : 'text-gray-700 hover:text-blue-600 hover:bg-gray-100'
                    ]) }));
                /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['border-r']} */ ;
                /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
                /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
                /** @type {__VLS_StyleScopedClasses['transition']} */ ;
                (content.partLabel);
                (content.subpartLabel ? '.' + content.subpartLabel : '');
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "ml-1 text-xs text-gray-500" }));
                /** @type {__VLS_StyleScopedClasses['ml-1']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                (content.score);
                // @ts-ignore
                [$t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, currentPage, currentPage, pageSize, pageSize, totalQuestions, totalQuestions, loading, questions, questions, questions, getDifficultyClass, getDifficultyName, getStatusClass, getStatusName, activeTabMap, activeTabMap,];
            };
            for (var _u = 0, _v = __VLS_vFor((question.questionContents)); _u < _v.length; _u++) {
                var _w = _v[_u], content = _w[0], index = _w[1];
                _loop_3(content, index);
            }
            if (question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0]) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-3 bg-white" }));
                /** @type {__VLS_StyleScopedClasses['p-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-2 flex items-center gap-2" }));
                /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
                /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-xs font-medium text-gray-500" }));
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                (__VLS_ctx.$t('examPaperForm.type'));
                (__VLS_ctx.QUESTION_TYPE_NAMES[question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].questionTypeId] || __VLS_ctx.$t('examPaperForm.unknown'));
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-xs text-gray-400" }));
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-xs text-gray-500" }));
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                (__VLS_ctx.$t('examPaperForm.score'));
                (question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].score);
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-3 text-sm text-gray-600" }));
                /** @type {__VLS_StyleScopedClasses['mt-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
                if (question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].questionTypeId === __VLS_ctx.QUESTION_TYPE_SINGLE_CHOICE) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    for (var _x = 0, _y = __VLS_vFor(((_b = question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].singleChoice) === null || _b === void 0 ? void 0 : _b.options)); _x < _y.length; _x++) {
                        var option = _y[_x][0];
                        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (option.prefix) }, { class: "flex gap-2" }));
                        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                        /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-gray-500" }));
                        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                        (option.prefix);
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
                        __VLS_asFunctionalDirective(__VLS_directives.vHtml, {})(null, __assign(__assign({}, __VLS_directiveBindingRestFields), { value: (option.content) }), null, null);
                        // @ts-ignore
                        [$t, $t, $t, activeTabMap, activeTabMap, activeTabMap, activeTabMap, activeTabMap, question_model_1.QUESTION_TYPE_NAMES, question_model_1.QUESTION_TYPE_NAMES, question_model_1.QUESTION_TYPE_SINGLE_CHOICE,];
                    }
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-2" }));
                    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-semibold" }));
                    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
                    (__VLS_ctx.$t('examPaperForm.correctAnswer'));
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "ml-2 text-blue-600" }));
                    /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-blue-600']} */ ;
                    ((_c = question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].singleChoice) === null || _c === void 0 ? void 0 : _c.answer);
                }
                else if (question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].questionTypeId === __VLS_ctx.QUESTION_TYPE_MULTIPLE_CHOICE) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    for (var _z = 0, _0 = __VLS_vFor(((_d = question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].multipleChoice) === null || _d === void 0 ? void 0 : _d.options)); _z < _0.length; _z++) {
                        var option = _0[_z][0];
                        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (option.prefix) }, { class: "flex gap-2" }));
                        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                        /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-gray-500" }));
                        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                        (option.prefix);
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
                        __VLS_asFunctionalDirective(__VLS_directives.vHtml, {})(null, __assign(__assign({}, __VLS_directiveBindingRestFields), { value: (option.content) }), null, null);
                        // @ts-ignore
                        [$t, activeTabMap, activeTabMap, activeTabMap, question_model_1.QUESTION_TYPE_MULTIPLE_CHOICE,];
                    }
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-2" }));
                    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-semibold" }));
                    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "ml-2 text-blue-600" }));
                    /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-blue-600']} */ ;
                    ((((_e = question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].multipleChoice) === null || _e === void 0 ? void 0 : _e.answer) || []).join(', '));
                }
                else if (question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].questionTypeId === __VLS_ctx.QUESTION_TYPE_TRUE_FALSE) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-semibold" }));
                    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "ml-2" }));
                    /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
                    (String((_f = question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].trueOrFalse) === null || _f === void 0 ? void 0 : _f.answer).toLowerCase() === 'true'
                        ? 'True'
                        : 'False');
                }
                else if (question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].questionTypeId === __VLS_ctx.QUESTION_TYPE_GAP_FILLING) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    for (var _1 = 0, _2 = __VLS_vFor(((_g = question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].gapFilling) === null || _g === void 0 ? void 0 : _g.answer)); _1 < _2.length; _1++) {
                        var _3 = _2[_1], ans = _3[0], gidx = _3[1];
                        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({
                            key: (gidx),
                        });
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-semibold" }));
                        /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
                        (__VLS_ctx.$t('examPaperForm.gap'));
                        (gidx + 1);
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "ml-2 font-semibold" }));
                        /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
                        /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
                        (__VLS_ctx.$t('examPaperForm.correctAnswer'));
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "ml-2" }));
                        /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
                        (ans);
                        // @ts-ignore
                        [$t, $t, activeTabMap, activeTabMap, activeTabMap, activeTabMap, activeTabMap, question_model_1.QUESTION_TYPE_TRUE_FALSE, question_model_1.QUESTION_TYPE_GAP_FILLING,];
                    }
                }
                else if (question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].questionTypeId === __VLS_ctx.QUESTION_TYPE_SHORT_ANSWER) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-semibold" }));
                    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
                    var __VLS_5 = index_vue_1.default || index_vue_1.default;
                    // @ts-ignore
                    var __VLS_6 = __VLS_asFunctionalComponent1(__VLS_5, new __VLS_5({
                        modelValue: (((_h = question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].shortAnswer) === null || _h === void 0 ? void 0 : _h.answer) || ''),
                        readOnly: true,
                        height: "100%",
                    }));
                    var __VLS_7 = __VLS_6.apply(void 0, __spreadArray([{
                            modelValue: (((_j = question.questionContents[__VLS_ctx.activeTabMap.get(question.id) || 0].shortAnswer) === null || _j === void 0 ? void 0 : _j.answer) || ''),
                            readOnly: true,
                            height: "100%",
                        }], __VLS_functionalComponentArgsRest(__VLS_6), false));
                }
                else {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "italic" }));
                    /** @type {__VLS_StyleScopedClasses['italic']} */ ;
                    (__VLS_ctx.$t('examPaperForm.pleaseViewFull'));
                }
            }
        }
        if (question.pastPaper) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4" }));
            /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.h4, __VLS_intrinsics.h4)(__assign({ class: "text-sm font-medium text-gray-700 mb-2" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
            /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
            (__VLS_ctx.$t('pastPaper.title'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-xs text-gray-600 bg-yellow-50 p-2 rounded" }));
            /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
            /** @type {__VLS_StyleScopedClasses['bg-yellow-50']} */ ;
            /** @type {__VLS_StyleScopedClasses['p-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
            (question.pastPaper.name);
            (question.pastPaper.year);
            if (question.indexInPastPaper) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
                (question.indexInPastPaper);
            }
        }
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "px-4 py-3 bg-gray-50 border-t border-gray-200 flex justify-between items-center" }));
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-t']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-xs text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (__VLS_ctx.formatDate(question.updatedAt));
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex space-x-2" }));
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['space-x-2']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex space-x-2" }));
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['space-x-2']} */ ;
        var __VLS_10 = void 0;
        /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
        routerLink;
        // @ts-ignore
        var __VLS_11 = __VLS_asFunctionalComponent1(__VLS_10, new __VLS_10(__assign({ to: ("/admin/questions/".concat(question.id)) }, { class: "text-indigo-600 hover:text-indigo-900 text-sm font-medium" })));
        var __VLS_12 = __VLS_11.apply(void 0, __spreadArray([__assign({ to: ("/admin/questions/".concat(question.id)) }, { class: "text-indigo-600 hover:text-indigo-900 text-sm font-medium" })], __VLS_functionalComponentArgsRest(__VLS_11), false));
        /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-indigo-900']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        var __VLS_15 = __VLS_13.slots.default;
        (__VLS_ctx.$t('common.view'));
        // @ts-ignore
        [$t, $t, $t, activeTabMap, activeTabMap, question_model_1.QUESTION_TYPE_SHORT_ANSWER, formatDate,];
        var __VLS_16 = void 0;
        /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
        routerLink;
        // @ts-ignore
        var __VLS_17 = __VLS_asFunctionalComponent1(__VLS_16, new __VLS_16(__assign({ to: ("/admin/questions/".concat(question.id, "/edit")) }, { class: "text-indigo-600 hover:text-indigo-900 text-sm font-medium" })));
        var __VLS_18 = __VLS_17.apply(void 0, __spreadArray([__assign({ to: ("/admin/questions/".concat(question.id, "/edit")) }, { class: "text-indigo-600 hover:text-indigo-900 text-sm font-medium" })], __VLS_functionalComponentArgsRest(__VLS_17), false));
        /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-indigo-900']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        var __VLS_21 = __VLS_19.slots.default;
        (__VLS_ctx.$t('common.edit'));
        // @ts-ignore
        [$t,];
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!!(__VLS_ctx.loading))
                    return;
                if (!!(!__VLS_ctx.questions || __VLS_ctx.questions.length === 0))
                    return;
                __VLS_ctx.deleteQuestion(question.id);
                // @ts-ignore
                [deleteQuestion,];
            } }, { class: "text-red-600 hover:text-red-900 text-sm font-medium" }));
        /** @type {__VLS_StyleScopedClasses['text-red-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-red-900']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        (__VLS_ctx.$t('common.delete'));
        // @ts-ignore
        [$t,];
    };
    var __VLS_13, __VLS_19;
    for (var _q = 0, _r = __VLS_vFor((__VLS_ctx.questions)); _q < _r.length; _q++) {
        var question = _r[_q][0];
        _loop_1(question);
    }
}
if (!__VLS_ctx.loading && __VLS_ctx.totalQuestions > 0) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-8 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0" }));
    /** @type {__VLS_StyleScopedClasses['mt-8']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:flex-row']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:space-y-0']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-700" }));
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    (((__VLS_ctx.currentPage - 1) * __VLS_ctx.pageSize) + 1);
    (Math.min(__VLS_ctx.currentPage * __VLS_ctx.pageSize, __VLS_ctx.totalQuestions));
    (__VLS_ctx.totalQuestions);
    if (__VLS_ctx.totalPages > 1) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.nav, __VLS_intrinsics.nav)(__assign({ class: "flex items-center space-x-2" }));
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['space-x-2']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(!__VLS_ctx.loading && __VLS_ctx.totalQuestions > 0))
                    return;
                if (!(__VLS_ctx.totalPages > 1))
                    return;
                __VLS_ctx.goToPage(__VLS_ctx.currentPage - 1);
                // @ts-ignore
                [currentPage, currentPage, currentPage, pageSize, pageSize, totalQuestions, totalQuestions, totalQuestions, loading, totalPages, goToPage,];
            } }, { disabled: (__VLS_ctx.currentPage === 1) }), { class: "px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed" }));
        /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
        var _loop_2 = function (page) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!(!__VLS_ctx.loading && __VLS_ctx.totalQuestions > 0))
                        return;
                    if (!(__VLS_ctx.totalPages > 1))
                        return;
                    __VLS_ctx.goToPage(page);
                    // @ts-ignore
                    [currentPage, goToPage, paginationRange,];
                } }, { key: (page) }), { class: ([
                    'px-3 py-2 text-sm font-medium rounded-md',
                    page === __VLS_ctx.currentPage
                        ? 'text-white bg-indigo-600 border border-indigo-600'
                        : 'text-gray-700 bg-white border border-gray-300 hover:bg-gray-50'
                ]) }));
            /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
            (page);
            // @ts-ignore
            [currentPage,];
        };
        for (var _s = 0, _t = __VLS_vFor((__VLS_ctx.paginationRange)); _s < _t.length; _s++) {
            var page = _t[_s][0];
            _loop_2(page);
        }
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(!__VLS_ctx.loading && __VLS_ctx.totalQuestions > 0))
                    return;
                if (!(__VLS_ctx.totalPages > 1))
                    return;
                __VLS_ctx.goToPage(__VLS_ctx.currentPage + 1);
                // @ts-ignore
                [currentPage, goToPage,];
            } }, { disabled: (__VLS_ctx.currentPage === __VLS_ctx.totalPages) }), { class: "px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed" }));
        /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
    }
}
// @ts-ignore
[currentPage, totalPages,];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
