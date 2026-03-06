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
var _a, _b, _c, _d;
Object.defineProperty(exports, "__esModule", { value: true });
var vue_1 = require("vue");
var userStore_1 = require("../stores/userStore");
var index_vue_1 = require("../components/QuillEditor/index.vue");
var PracticeSidebar_vue_1 = require("../components/PracticeSidebar.vue");
var practiceService_1 = require("../services/practiceService");
var syllabusService_1 = require("../services/syllabusService");
var organisationService_1 = require("../services/organisationService");
var qualificationService_1 = require("../services/qualificationService");
var chapterService_1 = require("../services/chapterService");
var ChapterOption_vue_1 = require("../components/admin/ChapterOption.vue");
var question_model_1 = require("../models/question.model");
var userStore = (0, userStore_1.useUserStore)();
var organisations = (0, vue_1.ref)([]);
var qualifications = (0, vue_1.ref)([]);
var syllabuses = (0, vue_1.ref)([]);
var selectedOrganisationId = (0, vue_1.ref)('');
var selectedQualificationId = (0, vue_1.ref)('');
var syllabusId = (0, vue_1.ref)('');
var questionCount = (0, vue_1.ref)(5);
var questionIds = (0, vue_1.ref)([]);
var questions = (0, vue_1.ref)({});
var currentIndex = (0, vue_1.ref)(0);
var currentQuestion = (0, vue_1.computed)(function () { return questions.value[questionIds.value[currentIndex.value]] || { stem: '', questionContents: [] }; });
var answers = (0, vue_1.reactive)({});
var answersMulti = (0, vue_1.reactive)({});
var result = (0, vue_1.ref)(null);
var sidebarVisible = (0, vue_1.ref)(true);
var resultItem = (0, vue_1.computed)(function () {
    if (!result.value || !result.value.results)
        return null;
    return result.value.results.find(function (r) { return r.questionId === questionIds.value[currentIndex.value]; }) || null;
});
var currentQuestionContents = (0, vue_1.computed)(function () {
    var q = questions.value[questionIds.value[currentIndex.value]];
    return q && Array.isArray(q.questionContents) ? q.questionContents : [];
});
// 新增章节筛选相关
var chapterTree = (0, vue_1.ref)([]);
var selectedChapterIds = (0, vue_1.ref)([]);
var showChapterSelector = (0, vue_1.ref)(false);
var resetAll = function () {
    questions.value = [];
    result.value = null;
    Object.keys(answers).forEach(function (k) { return delete answers[k]; });
    Object.keys(answersMulti).forEach(function (k) { return delete answersMulti[k]; });
    selectedOrganisationId.value = '';
    selectedQualificationId.value = '';
    syllabusId.value = '';
    questionCount.value = 5;
    qualifications.value = [];
    syllabuses.value = [];
    chapterTree.value = [];
    selectedChapterIds.value = [];
    showChapterSelector.value = false;
};
var fetchOrganisations = function () { return __awaiter(void 0, void 0, void 0, function () {
    var res;
    var _a;
    return __generator(this, function (_b) {
        switch (_b.label) {
            case 0: return [4 /*yield*/, organisationService_1.default.getOrganisations({ pageIndex: 1, pageSize: 100 })];
            case 1:
                res = _b.sent();
                organisations.value = ((_a = res.data) === null || _a === void 0 ? void 0 : _a.list) || [];
                return [2 /*return*/];
        }
    });
}); };
var fetchQualifications = function () { return __awaiter(void 0, void 0, void 0, function () {
    var res;
    var _a;
    return __generator(this, function (_b) {
        switch (_b.label) {
            case 0:
                if (!selectedOrganisationId.value) {
                    qualifications.value = [];
                    return [2 /*return*/];
                }
                return [4 /*yield*/, qualificationService_1.default.getQualifications({ organisationId: selectedOrganisationId.value, pageIndex: 1, pageSize: 100 })];
            case 1:
                res = _b.sent();
                qualifications.value = ((_a = res.data) === null || _a === void 0 ? void 0 : _a.list) || [];
                return [2 /*return*/];
        }
    });
}); };
var fetchSyllabuses = function () { return __awaiter(void 0, void 0, void 0, function () {
    var res;
    var _a;
    return __generator(this, function (_b) {
        switch (_b.label) {
            case 0:
                if (!selectedQualificationId.value) {
                    syllabuses.value = [];
                    return [2 /*return*/];
                }
                return [4 /*yield*/, syllabusService_1.default.getSyllabuses({ qualificationId: selectedQualificationId.value, pageIndex: 1, pageSize: 100 })];
            case 1:
                res = _b.sent();
                syllabuses.value = ((_a = res.data) === null || _a === void 0 ? void 0 : _a.list) || [];
                return [2 /*return*/];
        }
    });
}); };
var fetchChapterTree = function () { return __awaiter(void 0, void 0, void 0, function () {
    var res, e_1;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                chapterTree.value = [];
                selectedChapterIds.value = [];
                if (!syllabusId.value)
                    return [2 /*return*/];
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, , 4]);
                return [4 /*yield*/, chapterService_1.default.getChapterTree(Number(syllabusId.value))];
            case 2:
                res = _a.sent();
                chapterTree.value = res.data || [];
                return [3 /*break*/, 4];
            case 3:
                e_1 = _a.sent();
                chapterTree.value = [];
                return [3 /*break*/, 4];
            case 4: return [2 /*return*/];
        }
    });
}); };
var onOrganisationChange = function () { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                selectedQualificationId.value = '';
                syllabusId.value = '';
                qualifications.value = [];
                syllabuses.value = [];
                chapterTree.value = [];
                selectedChapterIds.value = [];
                showChapterSelector.value = false;
                return [4 /*yield*/, fetchQualifications()];
            case 1:
                _a.sent();
                return [2 /*return*/];
        }
    });
}); };
var onQualificationChange = function () { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                syllabusId.value = '';
                syllabuses.value = [];
                chapterTree.value = [];
                selectedChapterIds.value = [];
                showChapterSelector.value = false;
                return [4 /*yield*/, fetchSyllabuses()];
            case 1:
                _a.sent();
                return [2 /*return*/];
        }
    });
}); };
(0, vue_1.watch)(syllabusId, function (newVal) { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                chapterTree.value = [];
                selectedChapterIds.value = [];
                showChapterSelector.value = false;
                if (!newVal) return [3 /*break*/, 2];
                return [4 /*yield*/, fetchChapterTree()];
            case 1:
                _a.sent();
                _a.label = 2;
            case 2: return [2 /*return*/];
        }
    });
}); });
(0, vue_1.onMounted)(function () { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, fetchOrganisations()];
            case 1:
                _a.sent();
                return [2 /*return*/];
        }
    });
}); });
var startPractice = function () { return __awaiter(void 0, void 0, void 0, function () {
    var req, res;
    var _a;
    return __generator(this, function (_b) {
        switch (_b.label) {
            case 0:
                if (!syllabusId.value || !questionCount.value)
                    return [2 /*return*/];
                req = {
                    syllabusId: Number(syllabusId.value),
                    questionCount: questionCount.value,
                    chapterIds: selectedChapterIds.value.length > 0 ? selectedChapterIds.value : undefined
                };
                return [4 /*yield*/, practiceService_1.practiceService.quickPractice(req)];
            case 1:
                res = _b.sent();
                questionIds.value = ((_a = res.data) === null || _a === void 0 ? void 0 : _a.list) || [];
                questions.value = {};
                currentIndex.value = 0;
                result.value = null;
                Object.keys(answers).forEach(function (k) { return delete answers[k]; });
                Object.keys(answersMulti).forEach(function (k) { return delete answersMulti[k]; });
                if (!(questionIds.value.length > 0)) return [3 /*break*/, 3];
                return [4 /*yield*/, loadQuestionByIndex(0)];
            case 2:
                _b.sent();
                _b.label = 3;
            case 3: return [2 /*return*/];
        }
    });
}); };
var loadQuestionByIndex = function (idx) { return __awaiter(void 0, void 0, void 0, function () {
    var id, res;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                id = questionIds.value[idx];
                if (!id)
                    return [2 /*return*/];
                if (!!questions.value[id]) return [3 /*break*/, 3];
                return [4 /*yield*/, Promise.resolve().then(function () { return require('../services/questionService'); })];
            case 1: return [4 /*yield*/, (_a.sent()).default.getQuestionById(id)];
            case 2:
                res = _a.sent();
                questions.value[id] = res.data;
                _a.label = 3;
            case 3:
                currentIndex.value = idx;
                return [2 /*return*/];
        }
    });
}); };
var jumpTo = function (idx) { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, loadQuestionByIndex(idx)];
            case 1:
                _a.sent();
                return [2 /*return*/];
        }
    });
}); };
var submitAnswers = function () { return __awaiter(void 0, void 0, void 0, function () {
    var submissionAnswers, req, res;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                submissionAnswers = [];
                questionIds.value.forEach(function (id) {
                    var q = questions.value[id];
                    if (q && q.questionContents && q.questionContents.length > 0) {
                        var partAnswers_1 = [];
                        q.questionContents.forEach(function (content, cidx) {
                            var answer = '';
                            if (content.questionTypeId === question_model_1.QUESTION_TYPE_MULTIPLE_CHOICE) {
                                answer = answersMulti[id + '-' + cidx] || [];
                            }
                            else {
                                answer = answers[id + '-' + cidx] || '';
                            }
                            partAnswers_1.push({
                                questionContentId: cidx, // If content.id exists, use content.id instead
                                answer: Array.isArray(answer) ? JSON.stringify(answer) : answer
                            });
                        });
                        submissionAnswers.push({
                            questionId: id,
                            answers: partAnswers_1
                        });
                    }
                });
                req = submissionAnswers;
                return [4 /*yield*/, practiceService_1.practiceService.gradePractice(req)];
            case 1:
                res = _a.sent();
                result.value = res.data;
                return [2 /*return*/];
        }
    });
}); };
var __VLS_ctx = __assign(__assign({}, {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-6 w-full" }));
/** @type {__VLS_StyleScopedClasses['p-6']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
var __VLS_0 = PracticeSidebar_vue_1.default;
// @ts-ignore
var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0({
    visible: (__VLS_ctx.sidebarVisible),
    questionIds: (__VLS_ctx.questionIds),
    questions: (__VLS_ctx.questions),
    currentIndex: (__VLS_ctx.currentIndex),
    answers: (__VLS_ctx.answers),
    result: (__VLS_ctx.result),
    jumpTo: (__VLS_ctx.jumpTo),
    onClose: (function () { return __VLS_ctx.sidebarVisible = false; }),
}));
var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([{
        visible: (__VLS_ctx.sidebarVisible),
        questionIds: (__VLS_ctx.questionIds),
        questions: (__VLS_ctx.questions),
        currentIndex: (__VLS_ctx.currentIndex),
        answers: (__VLS_ctx.answers),
        result: (__VLS_ctx.result),
        jumpTo: (__VLS_ctx.jumpTo),
        onClose: (function () { return __VLS_ctx.sidebarVisible = false; }),
    }], __VLS_functionalComponentArgsRest(__VLS_1), false));
if (__VLS_ctx.questionIds.length > 0 && !__VLS_ctx.sidebarVisible) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.questionIds.length > 0 && !__VLS_ctx.sidebarVisible))
                return;
            __VLS_ctx.sidebarVisible = true;
            // @ts-ignore
            [sidebarVisible, sidebarVisible, sidebarVisible, sidebarVisible, questionIds, questionIds, questions, currentIndex, answers, result, jumpTo,];
        } }, { class: "fixed right-4 bottom-4 z-50 bg-indigo-600 text-white px-4 py-2 rounded shadow-lg hover:bg-indigo-700" }));
    /** @type {__VLS_StyleScopedClasses['fixed']} */ ;
    /** @type {__VLS_StyleScopedClasses['right-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['bottom-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['z-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-indigo-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-indigo-700']} */ ;
    (__VLS_ctx.$t('quickPractice.showSidebar'));
}
if (!__VLS_ctx.userStore.isAuthenticated) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-6 bg-yellow-100 border-l-4 border-yellow-500 text-yellow-700 p-4 rounded" }));
    /** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-yellow-100']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-l-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-yellow-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-yellow-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
    (__VLS_ctx.$t('quickPractice.notLoggedIn'));
    var __VLS_5 = void 0;
    /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
    routerLink;
    // @ts-ignore
    var __VLS_6 = __VLS_asFunctionalComponent1(__VLS_5, new __VLS_5(__assign({ to: "/login" }, { class: "ml-2 text-blue-600 underline" })));
    var __VLS_7 = __VLS_6.apply(void 0, __spreadArray([__assign({ to: "/login" }, { class: "ml-2 text-blue-600 underline" })], __VLS_functionalComponentArgsRest(__VLS_6), false));
    /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-blue-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['underline']} */ ;
    var __VLS_10 = __VLS_8.slots.default;
    (__VLS_ctx.$t('quickPractice.goToLogin'));
    // @ts-ignore
    [$t, $t, $t, userStore,];
    var __VLS_8;
}
__VLS_asFunctionalElement1(__VLS_intrinsics.header, __VLS_intrinsics.header)(__assign({ class: "mb-6" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.h1, __VLS_intrinsics.h1)(__assign({ class: "text-2xl font-bold text-gray-900" }));
/** @type {__VLS_StyleScopedClasses['text-2xl']} */ ;
/** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
(__VLS_ctx.$t('quickPractice.title'));
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-sm text-gray-500" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
(__VLS_ctx.$t('quickPractice.subtitle'));
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-2 bg-yellow-50 border-l-4 border-yellow-400 text-yellow-700 p-3 rounded" }));
/** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-yellow-50']} */ ;
/** @type {__VLS_StyleScopedClasses['border-l-4']} */ ;
/** @type {__VLS_StyleScopedClasses['border-yellow-400']} */ ;
/** @type {__VLS_StyleScopedClasses['text-yellow-700']} */ ;
/** @type {__VLS_StyleScopedClasses['p-3']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
(__VLS_ctx.$t('quickPractice.knowledgePointNotice'));
__VLS_asFunctionalElement1(__VLS_intrinsics.form, __VLS_intrinsics.form)(__assign({ onSubmit: (__VLS_ctx.startPractice) }, { class: "mb-6 flex flex-wrap gap-4 items-end w-full" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-wrap']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
/** @type {__VLS_StyleScopedClasses['items-end']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
(__VLS_ctx.$t('quickPractice.organisation'));
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ onChange: (__VLS_ctx.onOrganisationChange) }, { value: (__VLS_ctx.selectedOrganisationId), required: true }), { class: "px-3 py-2 border border-gray-300 rounded-md min-w-[180px] w-full" }));
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[180px]']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "",
    disabled: true,
});
(__VLS_ctx.$t('quickPractice.selectOrganisation'));
for (var _i = 0, _e = __VLS_vFor((__VLS_ctx.organisations)); _i < _e.length; _i++) {
    var org = _e[_i][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (org.id),
        value: (org.id),
    });
    (org.name);
    // @ts-ignore
    [$t, $t, $t, $t, $t, startPractice, onOrganisationChange, selectedOrganisationId, organisations,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
(__VLS_ctx.$t('quickPractice.qualification'));
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ onChange: (__VLS_ctx.onQualificationChange) }, { value: (__VLS_ctx.selectedQualificationId), disabled: (!__VLS_ctx.selectedOrganisationId), required: true }), { class: "px-3 py-2 border border-gray-300 rounded-md min-w-[180px] w-full" }));
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[180px]']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "",
    disabled: true,
});
(__VLS_ctx.$t('quickPractice.selectQualification'));
for (var _f = 0, _g = __VLS_vFor((__VLS_ctx.qualifications)); _f < _g.length; _f++) {
    var q = _g[_f][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (q.id),
        value: (q.id),
    });
    (q.name);
    // @ts-ignore
    [$t, $t, selectedOrganisationId, onQualificationChange, selectedQualificationId, qualifications,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
(__VLS_ctx.$t('quickPractice.syllabus'));
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ value: (__VLS_ctx.syllabusId), disabled: (!__VLS_ctx.selectedQualificationId), required: true }, { class: "px-3 py-2 border border-gray-300 rounded-md min-w-[180px] w-full" }));
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[180px]']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "",
    disabled: true,
});
(__VLS_ctx.$t('quickPractice.selectSyllabus'));
for (var _h = 0, _j = __VLS_vFor((__VLS_ctx.syllabuses)); _h < _j.length; _h++) {
    var syllabus = _j[_h][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (syllabus.id),
        value: (syllabus.id),
    });
    (syllabus.name);
    (syllabus.code);
    // @ts-ignore
    [$t, $t, selectedQualificationId, syllabusId, syllabuses,];
}
if (__VLS_ctx.syllabusId) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
    (__VLS_ctx.$t('quickPractice.chapters'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "relative min-w-[200px] max-w-xl" }));
    /** @type {__VLS_StyleScopedClasses['relative']} */ ;
    /** @type {__VLS_StyleScopedClasses['min-w-[200px]']} */ ;
    /** @type {__VLS_StyleScopedClasses['max-w-xl']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.syllabusId))
                return;
            __VLS_ctx.showChapterSelector = !__VLS_ctx.showChapterSelector;
            // @ts-ignore
            [$t, syllabusId, showChapterSelector, showChapterSelector,];
        } }, { type: "button", disabled: (!__VLS_ctx.chapterTree.length) }), { class: "flex justify-between items-center w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 disabled:bg-gray-100 disabled:cursor-not-allowed" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:border-indigo-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['disabled:bg-gray-100']} */ ;
    /** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-sm truncate" }));
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['truncate']} */ ;
    (__VLS_ctx.selectedChapterIds.length ? __VLS_ctx.$t('quickPractice.chaptersSelected', { count: __VLS_ctx.selectedChapterIds.length }) : __VLS_ctx.$t('quickPractice.allChapters'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-5 h-5 text-gray-400" }, { viewBox: "0 0 20 20", fill: "currentColor" }));
    /** @type {__VLS_StyleScopedClasses['w-5']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-5']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
        'fill-rule': "evenodd",
        d: "M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z",
        'clip-rule': "evenodd",
    });
    if (__VLS_ctx.showChapterSelector) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "absolute z-10 w-full lg:w-96 mt-1 bg-white rounded-md shadow-lg border border-gray-200" }));
        /** @type {__VLS_StyleScopedClasses['absolute']} */ ;
        /** @type {__VLS_StyleScopedClasses['z-10']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['lg:w-96']} */ ;
        /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['shadow-lg']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "max-h-96 overflow-y-auto p-2" }));
        /** @type {__VLS_StyleScopedClasses['max-h-96']} */ ;
        /** @type {__VLS_StyleScopedClasses['overflow-y-auto']} */ ;
        /** @type {__VLS_StyleScopedClasses['p-2']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-center justify-between p-2 border-b" }));
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
        /** @type {__VLS_StyleScopedClasses['p-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-b']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-sm font-medium text-gray-900" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
        (__VLS_ctx.$t('quickPractice.selectChapters'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(__VLS_ctx.syllabusId))
                    return;
                if (!(__VLS_ctx.showChapterSelector))
                    return;
                __VLS_ctx.selectedChapterIds = [];
                __VLS_ctx.showChapterSelector = false;
                // @ts-ignore
                [$t, $t, $t, showChapterSelector, showChapterSelector, chapterTree, selectedChapterIds, selectedChapterIds, selectedChapterIds,];
            } }, { type: "button" }), { class: "text-sm text-gray-500 hover:text-gray-700" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-gray-700']} */ ;
        (__VLS_ctx.$t('quickPractice.clear'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-2" }));
        /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
        for (var _k = 0, _l = __VLS_vFor((__VLS_ctx.chapterTree)); _k < _l.length; _k++) {
            var _m = _l[_k], chapter = _m[0], index = _m[1];
            var __VLS_11 = ChapterOption_vue_1.default;
            // @ts-ignore
            var __VLS_12 = __VLS_asFunctionalComponent1(__VLS_11, new __VLS_11(__assign({ 'onUpdate:selected': {} }, { key: (chapter.id), chapter: (chapter), level: (0), isLast: (index === __VLS_ctx.chapterTree.length - 1), selectedChapters: (__VLS_ctx.selectedChapterIds) })));
            var __VLS_13 = __VLS_12.apply(void 0, __spreadArray([__assign({ 'onUpdate:selected': {} }, { key: (chapter.id), chapter: (chapter), level: (0), isLast: (index === __VLS_ctx.chapterTree.length - 1), selectedChapters: (__VLS_ctx.selectedChapterIds) })], __VLS_functionalComponentArgsRest(__VLS_12), false));
            var __VLS_16 = void 0;
            var __VLS_17 = ({ 'update:selected': {} },
                { 'onUpdate:selected': (function (val) { __VLS_ctx.selectedChapterIds = val; }) });
            var __VLS_14;
            var __VLS_15;
            // @ts-ignore
            [$t, chapterTree, chapterTree, selectedChapterIds, selectedChapterIds,];
        }
    }
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
(__VLS_ctx.$t('quickPractice.questionCount'));
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "number", min: "1", max: "20", required: true }, { class: "px-3 py-2 border border-gray-300 rounded-md min-w-[100px] w-full" }));
(__VLS_ctx.questionCount);
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[100px]']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ type: "submit" }, { class: "px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700" }));
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-blue-600']} */ ;
/** @type {__VLS_StyleScopedClasses['text-white']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-blue-700']} */ ;
(__VLS_ctx.$t('quickPractice.start'));
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: (__VLS_ctx.resetAll) }, { type: "button" }), { class: "px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300" }));
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-200']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-800']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-gray-300']} */ ;
(__VLS_ctx.$t('quickPractice.reset'));
if (__VLS_ctx.questionIds.length > 0) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-6" }));
    /** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.form, __VLS_intrinsics.form)(__assign({ onSubmit: (__VLS_ctx.submitAnswers) }));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-6 bg-white rounded shadow p-4" }));
    /** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
    /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
    if (__VLS_ctx.questions[__VLS_ctx.questionIds[__VLS_ctx.currentIndex]]) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-2 font-semibold text-gray-800" }));
        /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-800']} */ ;
        (__VLS_ctx.currentIndex + 1);
        var __VLS_18 = index_vue_1.default || index_vue_1.default;
        // @ts-ignore
        var __VLS_19 = __VLS_asFunctionalComponent1(__VLS_18, new __VLS_18({
            modelValue: (__VLS_ctx.currentQuestion.stem),
            readOnly: true,
            height: "100%",
        }));
        var __VLS_20 = __VLS_19.apply(void 0, __spreadArray([{
                modelValue: (__VLS_ctx.currentQuestion.stem),
                readOnly: true,
                height: "100%",
            }], __VLS_functionalComponentArgsRest(__VLS_19), false));
        if (__VLS_ctx.currentQuestionContents.length > 0) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
            var _loop_1 = function (content, cidx) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (cidx) }, { class: "mb-4 p-3 bg-white rounded border" }));
                /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
                /** @type {__VLS_StyleScopedClasses['p-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
                /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
                /** @type {__VLS_StyleScopedClasses['border']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-2 flex items-center gap-2" }));
                /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
                /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-xs font-medium text-gray-500" }));
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                (content.partLabel);
                (content.subpartLabel ? '.' + content.subpartLabel : '');
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-xs text-gray-400" }));
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-xs font-medium text-gray-500" }));
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                (__VLS_ctx.QUESTION_TYPE_NAMES[content.questionTypeId] || content.questionTypeId || '');
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-xs text-gray-400" }));
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-xs text-gray-500" }));
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                (content.score);
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-3 text-sm text-gray-600" }));
                /** @type {__VLS_StyleScopedClasses['mt-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
                if (content.questionTypeId === __VLS_ctx.QUESTION_TYPE_SINGLE_CHOICE) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                    for (var _u = 0, _v = __VLS_vFor(((_a = content.singleChoice) === null || _a === void 0 ? void 0 : _a.options)); _u < _v.length; _u++) {
                        var opt = _v[_u][0];
                        __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ key: (opt.prefix) }, { class: "block" }));
                        /** @type {__VLS_StyleScopedClasses['block']} */ ;
                        __VLS_asFunctionalElement1(__VLS_intrinsics.input)({
                            type: "radio",
                            name: ('single-' + __VLS_ctx.questionIds[__VLS_ctx.currentIndex] + '-' + cidx),
                            value: (opt.prefix),
                        });
                        (__VLS_ctx.answers[__VLS_ctx.questionIds[__VLS_ctx.currentIndex] + '-' + cidx]);
                        (opt.prefix);
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
                        __VLS_asFunctionalDirective(__VLS_directives.vHtml, {})(null, __assign(__assign({}, __VLS_directiveBindingRestFields), { value: (opt.content) }), null, null);
                        // @ts-ignore
                        [questionIds, questionIds, questionIds, questionIds, questions, currentIndex, currentIndex, currentIndex, currentIndex, answers, $t, $t, $t, questionCount, resetAll, submitAnswers, currentQuestion, currentQuestionContents, currentQuestionContents, question_model_1.QUESTION_TYPE_NAMES, question_model_1.QUESTION_TYPE_NAMES, question_model_1.QUESTION_TYPE_SINGLE_CHOICE,];
                    }
                }
                else if (content.questionTypeId === __VLS_ctx.QUESTION_TYPE_MULTIPLE_CHOICE) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                    for (var _w = 0, _x = __VLS_vFor(((_b = content.multipleChoice) === null || _b === void 0 ? void 0 : _b.options)); _w < _x.length; _w++) {
                        var opt = _x[_w][0];
                        __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ key: (opt.prefix) }, { class: "block" }));
                        /** @type {__VLS_StyleScopedClasses['block']} */ ;
                        __VLS_asFunctionalElement1(__VLS_intrinsics.input)({
                            type: "checkbox",
                            value: (opt.prefix),
                        });
                        (__VLS_ctx.answersMulti[__VLS_ctx.questionIds[__VLS_ctx.currentIndex] + '-' + cidx]);
                        (opt.prefix);
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
                        __VLS_asFunctionalDirective(__VLS_directives.vHtml, {})(null, __assign(__assign({}, __VLS_directiveBindingRestFields), { value: (opt.content) }), null, null);
                        // @ts-ignore
                        [questionIds, currentIndex, question_model_1.QUESTION_TYPE_MULTIPLE_CHOICE, answersMulti,];
                    }
                }
                else if (content.questionTypeId === __VLS_ctx.QUESTION_TYPE_TRUE_FALSE) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)({});
                    __VLS_asFunctionalElement1(__VLS_intrinsics.input)({
                        type: "radio",
                        name: ('tf-' + __VLS_ctx.questionIds[__VLS_ctx.currentIndex] + '-' + cidx),
                        value: "true",
                    });
                    (__VLS_ctx.answers[__VLS_ctx.questionIds[__VLS_ctx.currentIndex] + '-' + cidx]);
                    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)({});
                    __VLS_asFunctionalElement1(__VLS_intrinsics.input)({
                        type: "radio",
                        name: ('tf-' + __VLS_ctx.questionIds[__VLS_ctx.currentIndex] + '-' + cidx),
                        value: "false",
                    });
                    (__VLS_ctx.answers[__VLS_ctx.questionIds[__VLS_ctx.currentIndex] + '-' + cidx]);
                }
                else if (content.questionTypeId === __VLS_ctx.QUESTION_TYPE_GAP_FILLING) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                    __VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign(__assign({ value: (__VLS_ctx.answers[__VLS_ctx.questionIds[__VLS_ctx.currentIndex] + '-' + cidx]), type: "text" }, { class: "px-3 py-2 border border-gray-300 rounded-md w-full" }), { placeholder: "Fill the gap" }));
                    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['border']} */ ;
                    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
                    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
                    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
                }
                else if (content.questionTypeId === __VLS_ctx.QUESTION_TYPE_SHORT_ANSWER) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                    var __VLS_23 = index_vue_1.default;
                    // @ts-ignore
                    var __VLS_24 = __VLS_asFunctionalComponent1(__VLS_23, new __VLS_23(__assign({ 'onUpdate:modelValue': {} }, { modelValue: ((_c = __VLS_ctx.answers[__VLS_ctx.questionIds[__VLS_ctx.currentIndex] + '-' + cidx]) !== null && _c !== void 0 ? _c : ''), readOnly: (false), height: "120px", placeholder: ('Your answer') })));
                    var __VLS_25 = __VLS_24.apply(void 0, __spreadArray([__assign({ 'onUpdate:modelValue': {} }, { modelValue: ((_d = __VLS_ctx.answers[__VLS_ctx.questionIds[__VLS_ctx.currentIndex] + '-' + cidx]) !== null && _d !== void 0 ? _d : ''), readOnly: (false), height: "120px", placeholder: ('Your answer') })], __VLS_functionalComponentArgsRest(__VLS_24), false));
                    var __VLS_28 = void 0;
                    var __VLS_29 = ({ 'update:modelValue': {} },
                        { 'onUpdate:modelValue': (function (val) { __VLS_ctx.answers[__VLS_ctx.questionIds[__VLS_ctx.currentIndex] + '-' + cidx] = val !== null && val !== void 0 ? val : ''; }) });
                }
                // @ts-ignore
                [questionIds, questionIds, questionIds, questionIds, questionIds, questionIds, questionIds, currentIndex, currentIndex, currentIndex, currentIndex, currentIndex, currentIndex, currentIndex, answers, answers, answers, answers, answers, question_model_1.QUESTION_TYPE_TRUE_FALSE, question_model_1.QUESTION_TYPE_GAP_FILLING, question_model_1.QUESTION_TYPE_SHORT_ANSWER,];
            };
            var __VLS_26, __VLS_27;
            for (var _o = 0, _p = __VLS_vFor((__VLS_ctx.currentQuestionContents)); _o < _p.length; _o++) {
                var _q = _p[_o], content = _q[0], cidx = _q[1];
                _loop_1(content, cidx);
            }
        }
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-center items-center gap-4 mb-4" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.questionIds.length > 0))
                return;
            __VLS_ctx.jumpTo(__VLS_ctx.currentIndex - 1);
            // @ts-ignore
            [currentIndex, jumpTo,];
        } }, { type: "button" }), { class: "px-4 py-2 bg-gray-200 text-gray-800 rounded hover:bg-gray-300" }), { disabled: (__VLS_ctx.currentIndex === 0) }));
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-gray-200']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-800']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-300']} */ ;
    (__VLS_ctx.$t('quickPractice.previous'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-bold text-lg" }));
    /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    (__VLS_ctx.currentIndex + 1);
    (__VLS_ctx.questionIds.length);
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.questionIds.length > 0))
                return;
            __VLS_ctx.jumpTo(__VLS_ctx.currentIndex + 1);
            // @ts-ignore
            [questionIds, currentIndex, currentIndex, currentIndex, jumpTo, $t,];
        } }, { type: "button" }), { class: "px-4 py-2 bg-gray-200 text-gray-800 rounded hover:bg-gray-300" }), { disabled: (__VLS_ctx.currentIndex === __VLS_ctx.questionIds.length - 1) }));
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-gray-200']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-800']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-300']} */ ;
    (__VLS_ctx.$t('quickPractice.next'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ type: "submit" }, { class: "px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700" }));
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-green-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-green-700']} */ ;
    (__VLS_ctx.$t('quickPractice.submitAnswers'));
}
if (__VLS_ctx.result && __VLS_ctx.questionIds.length > 0) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-8 bg-white rounded shadow p-6" }));
    /** @type {__VLS_StyleScopedClasses['mt-8']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
    /** @type {__VLS_StyleScopedClasses['p-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-xl font-bold mb-4" }));
    /** @type {__VLS_StyleScopedClasses['text-xl']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    (__VLS_ctx.$t('quickPractice.practiceResult'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-2" }));
    /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
    (__VLS_ctx.$t('quickPractice.score'));
    (__VLS_ctx.result.score);
    (__VLS_ctx.result.total);
    if (__VLS_ctx.resultItem) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4 border-b pb-2" }));
        /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-b']} */ ;
        /** @type {__VLS_StyleScopedClasses['pb-2']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "font-semibold" }));
        /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
        (__VLS_ctx.currentIndex + 1);
        for (var _r = 0, _s = __VLS_vFor((__VLS_ctx.resultItem.subResults)); _r < _s.length; _r++) {
            var _t = _s[_r], sub = _t[0], pidx = _t[1];
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (pidx) }, { class: "ml-4 mb-2" }));
            /** @type {__VLS_StyleScopedClasses['ml-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
            (__VLS_ctx.$t('quickPractice.part'));
            (pidx + 1);
            if (sub.questionType === __VLS_ctx.QUESTION_TYPE_SHORT_ANSWER) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-1" }));
                /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
                (__VLS_ctx.$t('quickPractice.yourAnswer'));
                var __VLS_30 = index_vue_1.default;
                // @ts-ignore
                var __VLS_31 = __VLS_asFunctionalComponent1(__VLS_30, new __VLS_30({
                    modelValue: (sub.studentAnswer),
                    readOnly: (true),
                    height: "120px",
                }));
                var __VLS_32 = __VLS_31.apply(void 0, __spreadArray([{
                        modelValue: (sub.studentAnswer),
                        readOnly: (true),
                        height: "120px",
                    }], __VLS_functionalComponentArgsRest(__VLS_31), false));
                if (sub.correctAnswer) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-1 mt-2" }));
                    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
                    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
                    (__VLS_ctx.$t('quickPractice.correctAnswer'));
                }
                if (sub.correctAnswer) {
                    var __VLS_35 = index_vue_1.default;
                    // @ts-ignore
                    var __VLS_36 = __VLS_asFunctionalComponent1(__VLS_35, new __VLS_35({
                        modelValue: (sub.correctAnswer),
                        readOnly: (true),
                        height: "120px",
                    }));
                    var __VLS_37 = __VLS_36.apply(void 0, __spreadArray([{
                            modelValue: (sub.correctAnswer),
                            readOnly: (true),
                            height: "120px",
                        }], __VLS_functionalComponentArgsRest(__VLS_36), false));
                }
                if (sub.modelAnswer) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-1 mt-2" }));
                    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
                    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
                    (__VLS_ctx.$t('quickPractice.modelAnswer'));
                }
                if (sub.modelAnswer) {
                    var __VLS_40 = index_vue_1.default;
                    // @ts-ignore
                    var __VLS_41 = __VLS_asFunctionalComponent1(__VLS_40, new __VLS_40({
                        modelValue: (sub.modelAnswer),
                        readOnly: (true),
                        height: "120px",
                    }));
                    var __VLS_42 = __VLS_41.apply(void 0, __spreadArray([{
                            modelValue: (sub.modelAnswer),
                            readOnly: (true),
                            height: "120px",
                        }], __VLS_functionalComponentArgsRest(__VLS_41), false));
                }
            }
            else {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                (__VLS_ctx.$t('quickPractice.yourAnswer'));
                (sub.studentAnswer);
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                (__VLS_ctx.$t('quickPractice.correctAnswer'));
                (sub.correctAnswer);
                if (sub.isCorrect === true) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-green-600" }));
                    /** @type {__VLS_StyleScopedClasses['text-green-600']} */ ;
                    (__VLS_ctx.$t('quickPractice.correct'));
                }
                else if (sub.isCorrect === false) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-red-600" }));
                    /** @type {__VLS_StyleScopedClasses['text-red-600']} */ ;
                    (__VLS_ctx.$t('quickPractice.incorrect'));
                }
                else {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-gray-600" }));
                    /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
                    (__VLS_ctx.$t('quickPractice.subjective'));
                }
                if (sub.modelAnswer) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                    (__VLS_ctx.$t('quickPractice.modelAnswer'));
                    (sub.modelAnswer);
                }
            }
            // @ts-ignore
            [questionIds, questionIds, currentIndex, currentIndex, result, result, result, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, question_model_1.QUESTION_TYPE_SHORT_ANSWER, resultItem, resultItem,];
        }
    }
}
// @ts-ignore
[];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
