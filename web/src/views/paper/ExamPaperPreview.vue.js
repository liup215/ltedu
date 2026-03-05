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
var _a, _b, _c, _d, _e, _f, _g, _h, _j, _k, _l, _m, _o, _p, _q, _r, _s;
Object.defineProperty(exports, "__esModule", { value: true });
var vue_1 = require("vue");
var vue_router_1 = require("vue-router");
var examPaperService_1 = require("../../services/examPaperService");
var questionService_1 = require("../../services/questionService");
var question_model_1 = require("../../models/question.model");
var exportDocx_1 = require("../../utils/exportDocx");
var index_vue_1 = require("../../components/QuillEditor/index.vue");
var route = (0, vue_router_1.useRoute)();
var paper = (0, vue_1.ref)(null);
var questions = (0, vue_1.ref)([]);
var isLoading = (0, vue_1.ref)(true);
var errorMessage = (0, vue_1.ref)(null);
var getQuestionTotalScore = function (question) {
    var _a;
    if (!((_a = question.questionContents) === null || _a === void 0 ? void 0 : _a.length))
        return 0;
    return question.questionContents.reduce(function (sum, content) { return sum + (content.score || 0); }, 0);
};
var totalPaperScore = (0, vue_1.computed)(function () {
    return questions.value.reduce(function (sum, q) { return sum + (q.totalScore || getQuestionTotalScore(q)); }, 0);
});
var syllabusInfo = (0, vue_1.computed)(function () {
    var _a, _b, _c, _d;
    var syl = (_a = paper.value) === null || _a === void 0 ? void 0 : _a.syllabus;
    if (!syl)
        return '';
    var parts = [
        (_c = (_b = syl.qualification) === null || _b === void 0 ? void 0 : _b.organisation) === null || _c === void 0 ? void 0 : _c.name,
        (_d = syl.qualification) === null || _d === void 0 ? void 0 : _d.name,
        syl.name
    ].filter(Boolean);
    return parts.join(' - ') + (syl.code ? " (".concat(syl.code, ")") : '');
});
var fetchPaper = function () { return __awaiter(void 0, void 0, void 0, function () {
    var paperId, paperRes, questionsPromises, responses, error_1;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                isLoading.value = true;
                errorMessage.value = null;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 5, 6, 7]);
                paperId = Number(route.params.id);
                if (!paperId) {
                    throw new Error('Invalid paper ID');
                }
                return [4 /*yield*/, examPaperService_1.examPaperService.getExamPaperById({ id: paperId })];
            case 2:
                paperRes = _a.sent();
                if (!paperRes.data) {
                    throw new Error('Failed to load paper');
                }
                paper.value = paperRes.data;
                if (!(paper.value.questionIds && paper.value.questionIds.length > 0)) return [3 /*break*/, 4];
                questionsPromises = paper.value.questionIds.map(function (qId) {
                    return questionService_1.default.getQuestionById(qId);
                });
                return [4 /*yield*/, Promise.all(questionsPromises)];
            case 3:
                responses = _a.sent();
                questions.value = responses
                    .filter(function (res) { return res.data; })
                    .map(function (res) { return res.data; })
                    .sort(function (a, b) {
                    var _a, _b, _c, _d;
                    var indexA = (_b = (_a = paper.value) === null || _a === void 0 ? void 0 : _a.questionIds.indexOf(a.id)) !== null && _b !== void 0 ? _b : 0;
                    var indexB = (_d = (_c = paper.value) === null || _c === void 0 ? void 0 : _c.questionIds.indexOf(b.id)) !== null && _d !== void 0 ? _d : 0;
                    return indexA - indexB;
                });
                _a.label = 4;
            case 4: return [3 /*break*/, 7];
            case 5:
                error_1 = _a.sent();
                errorMessage.value = error_1.message || '加载试卷失败';
                return [3 /*break*/, 7];
            case 6:
                isLoading.value = false;
                return [7 /*endfinally*/];
            case 7: return [2 /*return*/];
        }
    });
}); };
var handleExportPaper = function () { return __awaiter(void 0, void 0, void 0, function () {
    var exportPaper;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (!(paper.value && Array.isArray(questions.value) && questions.value.length > 0)) return [3 /*break*/, 2];
                exportPaper = __assign(__assign({}, paper.value), { questions: questions.value });
                return [4 /*yield*/, (0, exportDocx_1.exportExamPaperToDocx)(exportPaper)];
            case 1:
                _a.sent();
                _a.label = 2;
            case 2: return [2 /*return*/];
        }
    });
}); };
(0, vue_1.onMounted)(fetchPaper);
var __VLS_ctx = __assign(__assign({}, {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-6 bg-gray-50 min-h-screen" }));
/** @type {__VLS_StyleScopedClasses['p-6']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
/** @type {__VLS_StyleScopedClasses['min-h-screen']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.header, __VLS_intrinsics.header)(__assign({ class: "mb-6" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-between items-center" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.h1, __VLS_intrinsics.h1)(__assign({ class: "text-3xl font-bold text-gray-900" }));
/** @type {__VLS_StyleScopedClasses['text-3xl']} */ ;
/** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
(__VLS_ctx.$t('examPaperPreview.title'));
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-sm text-gray-500" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
(__VLS_ctx.$t('examPaperPreview.subtitle'));
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex gap-3" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-3']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.handleExportPaper) }, { class: "px-4 py-2 text-white bg-blue-600 rounded-md hover:bg-blue-700 flex items-center gap-2" }));
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['text-white']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-blue-600']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-blue-700']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-5 h-5" }, { fill: "none", stroke: "currentColor", viewBox: "0 0 24 24" }));
/** @type {__VLS_StyleScopedClasses['w-5']} */ ;
/** @type {__VLS_StyleScopedClasses['h-5']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.path)({
    'stroke-linecap': "round",
    'stroke-linejoin': "round",
    'stroke-width': "2",
    d: "M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4",
});
(__VLS_ctx.$t('examPaperPreview.exportPaper'));
var __VLS_0;
/** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
routerLink;
// @ts-ignore
var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0(__assign({ to: ("/paper/exam/teacher") }, { class: "px-4 py-2 text-gray-700 bg-white border rounded-md hover:bg-gray-50" })));
var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([__assign({ to: ("/paper/exam/teacher") }, { class: "px-4 py-2 text-gray-700 bg-white border rounded-md hover:bg-gray-50" })], __VLS_functionalComponentArgsRest(__VLS_1), false));
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
var __VLS_5 = __VLS_3.slots.default;
(__VLS_ctx.$t('examPaperPreview.backToList'));
// @ts-ignore
[$t, $t, $t, $t, handleExportPaper,];
var __VLS_3;
if (__VLS_ctx.isLoading) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-center items-center py-12" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-12']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600" }));
    /** @type {__VLS_StyleScopedClasses['animate-spin']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-8']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-8']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-b-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-blue-600']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "ml-3 text-gray-600" }));
    /** @type {__VLS_StyleScopedClasses['ml-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
    (__VLS_ctx.$t('examPaperPreview.loading'));
}
else if (__VLS_ctx.errorMessage) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-12" }));
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-12']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-red-600 bg-red-50 p-4 rounded-md" }));
    /** @type {__VLS_StyleScopedClasses['text-red-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-red-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    (__VLS_ctx.errorMessage);
}
else if (__VLS_ctx.paper) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow-md" }));
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-md']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-6 border-b border-gray-200" }));
    /** @type {__VLS_StyleScopedClasses['p-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-b']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-2xl font-bold text-center mb-4" }));
    /** @type {__VLS_StyleScopedClasses['text-2xl']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    (__VLS_ctx.paper.name);
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-center items-center gap-4 text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
    (__VLS_ctx.$t('examPaperPreview.year'));
    (__VLS_ctx.paper.year);
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
    (__VLS_ctx.$t('examPaperPreview.totalScore'));
    (__VLS_ctx.totalPaperScore);
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
    (__VLS_ctx.$t('examPaperPreview.questions'));
    (((_a = __VLS_ctx.paper.questionIds) === null || _a === void 0 ? void 0 : _a.length) || 0);
    if (__VLS_ctx.syllabusInfo) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-3 text-sm text-gray-600 text-center" }));
        /** @type {__VLS_StyleScopedClasses['mt-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
        (__VLS_ctx.syllabusInfo);
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-6" }));
    /** @type {__VLS_StyleScopedClasses['p-6']} */ ;
    for (var _i = 0, _t = __VLS_vFor((__VLS_ctx.questions)); _i < _t.length; _i++) {
        var _u = _t[_i], q = _u[0], index = _u[1];
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (q.id) }, { class: "mb-8" }));
        /** @type {__VLS_StyleScopedClasses['mb-8']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg border border-gray-300 p-4" }));
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-between items-start mb-4" }));
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-start']} */ ;
        /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-lg font-semibold" }));
        /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
        (__VLS_ctx.$t('examPaperPreview.question'));
        (index + 1);
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex gap-2" }));
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "px-2 py-1 bg-blue-100 text-blue-800 rounded-full text-xs font-medium" }));
        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-blue-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-blue-800']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        (__VLS_ctx.$t('examPaperPreview.score'));
        (q.totalScore || __VLS_ctx.getQuestionTotalScore(q) || '-');
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "px-2 py-1 bg-yellow-100 text-yellow-800 rounded-full text-xs font-medium" }));
        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-yellow-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-yellow-800']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        (__VLS_ctx.DIFFICULTY_NAMES[q.difficult] || __VLS_ctx.$t('examPaperPreview.unknown'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "prose max-w-none" }));
        /** @type {__VLS_StyleScopedClasses['prose']} */ ;
        /** @type {__VLS_StyleScopedClasses['max-w-none']} */ ;
        var __VLS_6 = index_vue_1.default || index_vue_1.default;
        // @ts-ignore
        var __VLS_7 = __VLS_asFunctionalComponent1(__VLS_6, new __VLS_6({
            modelValue: (q.stem),
            readOnly: true,
            height: "100%",
        }));
        var __VLS_8 = __VLS_7.apply(void 0, __spreadArray([{
                modelValue: (q.stem),
                readOnly: true,
                height: "100%",
            }], __VLS_functionalComponentArgsRest(__VLS_7), false));
        if (q.questionContents && q.questionContents.length > 0) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-4 space-y-4" }));
            /** @type {__VLS_StyleScopedClasses['mt-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
            for (var _v = 0, _w = __VLS_vFor((q.questionContents)); _v < _w.length; _v++) {
                var _x = _w[_v], content = _x[0], idx = _x[1];
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (idx) }, { class: "ml-4" }));
                /** @type {__VLS_StyleScopedClasses['ml-4']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "font-medium mb-2" }));
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
                (content.partLabel);
                (content.subpartLabel ? '.' + content.subpartLabel : '');
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-sm text-gray-500" }));
                /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                (content.score);
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "prose max-w-none mt-2" }));
                /** @type {__VLS_StyleScopedClasses['prose']} */ ;
                /** @type {__VLS_StyleScopedClasses['max-w-none']} */ ;
                /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
                if (content.questionTypeId === 1 && content.singleChoice) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    for (var _y = 0, _z = __VLS_vFor((content.singleChoice.options)); _y < _z.length; _y++) {
                        var option = _z[_y][0];
                        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (option.prefix) }, { class: ([
                                'pl-4 flex items-start',
                                ((_b = content.singleChoice) === null || _b === void 0 ? void 0 : _b.answer) === option.prefix ? 'text-blue-600' : ''
                            ]) }));
                        /** @type {__VLS_StyleScopedClasses['pl-4']} */ ;
                        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                        /** @type {__VLS_StyleScopedClasses['items-start']} */ ;
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium mr-2" }));
                        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                        /** @type {__VLS_StyleScopedClasses['mr-2']} */ ;
                        (option.prefix);
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
                        __VLS_asFunctionalDirective(__VLS_directives.vHtml, {})(null, __assign(__assign({}, __VLS_directiveBindingRestFields), { value: (option.content) }), null, null);
                        if (((_c = content.singleChoice) === null || _c === void 0 ? void 0 : _c.answer) === option.prefix) {
                            __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-4 h-4 ml-2 mt-1 flex-shrink-0" }, { fill: "currentColor", viewBox: "0 0 20 20" }));
                            /** @type {__VLS_StyleScopedClasses['w-4']} */ ;
                            /** @type {__VLS_StyleScopedClasses['h-4']} */ ;
                            /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
                            /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
                            /** @type {__VLS_StyleScopedClasses['flex-shrink-0']} */ ;
                            __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
                                'fill-rule': "evenodd",
                                d: "M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z",
                                'clip-rule': "evenodd",
                            });
                        }
                        // @ts-ignore
                        [$t, $t, $t, $t, $t, $t, $t, isLoading, errorMessage, errorMessage, paper, paper, paper, paper, totalPaperScore, syllabusInfo, syllabusInfo, questions, getQuestionTotalScore, question_model_1.DIFFICULTY_NAMES, question_model_1.DIFFICULTY_NAMES,];
                    }
                }
                else if (content.questionTypeId === 2 && content.multipleChoice) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    for (var _0 = 0, _1 = __VLS_vFor((content.multipleChoice.options)); _0 < _1.length; _0++) {
                        var option = _1[_0][0];
                        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (option.prefix) }, { class: ([
                                'pl-4 flex items-start',
                                ((_e = (_d = content.multipleChoice) === null || _d === void 0 ? void 0 : _d.answer) === null || _e === void 0 ? void 0 : _e.includes(option.prefix)) ? 'text-blue-600' : ''
                            ]) }));
                        /** @type {__VLS_StyleScopedClasses['pl-4']} */ ;
                        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                        /** @type {__VLS_StyleScopedClasses['items-start']} */ ;
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium mr-2" }));
                        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                        /** @type {__VLS_StyleScopedClasses['mr-2']} */ ;
                        (option.prefix);
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
                        __VLS_asFunctionalDirective(__VLS_directives.vHtml, {})(null, __assign(__assign({}, __VLS_directiveBindingRestFields), { value: (option.content) }), null, null);
                        if ((_g = (_f = content.multipleChoice) === null || _f === void 0 ? void 0 : _f.answer) === null || _g === void 0 ? void 0 : _g.includes(option.prefix)) {
                            __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-4 h-4 ml-2 mt-1 flex-shrink-0" }, { fill: "currentColor", viewBox: "0 0 20 20" }));
                            /** @type {__VLS_StyleScopedClasses['w-4']} */ ;
                            /** @type {__VLS_StyleScopedClasses['h-4']} */ ;
                            /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
                            /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
                            /** @type {__VLS_StyleScopedClasses['flex-shrink-0']} */ ;
                            __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
                                'fill-rule': "evenodd",
                                d: "M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z",
                                'clip-rule': "evenodd",
                            });
                        }
                        // @ts-ignore
                        [];
                    }
                }
                else if (content.questionTypeId === 3 && content.trueOrFalse) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "pl-4" }));
                    /** @type {__VLS_StyleScopedClasses['pl-4']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: ([
                            'pl-4 flex items-start',
                            Boolean((_h = content.trueOrFalse) === null || _h === void 0 ? void 0 : _h.answer) ? 'text-blue-600' : ''
                        ]) }));
                    /** @type {__VLS_StyleScopedClasses['pl-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                    /** @type {__VLS_StyleScopedClasses['items-start']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "mr-2" }));
                    /** @type {__VLS_StyleScopedClasses['mr-2']} */ ;
                    if (Boolean((_j = content.trueOrFalse) === null || _j === void 0 ? void 0 : _j.answer)) {
                        __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-4 h-4 ml-2 flex-shrink-0" }, { fill: "currentColor", viewBox: "0 0 20 20" }));
                        /** @type {__VLS_StyleScopedClasses['w-4']} */ ;
                        /** @type {__VLS_StyleScopedClasses['h-4']} */ ;
                        /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
                        /** @type {__VLS_StyleScopedClasses['flex-shrink-0']} */ ;
                        __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
                            'fill-rule': "evenodd",
                            d: "M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z",
                            'clip-rule': "evenodd",
                        });
                    }
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: ([
                            'pl-4 flex items-start',
                            ((_k = content.trueOrFalse) === null || _k === void 0 ? void 0 : _k.answer) === 0 ? 'text-blue-600' : ''
                        ]) }));
                    /** @type {__VLS_StyleScopedClasses['pl-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                    /** @type {__VLS_StyleScopedClasses['items-start']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "mr-2" }));
                    /** @type {__VLS_StyleScopedClasses['mr-2']} */ ;
                    if (((_l = content.trueOrFalse) === null || _l === void 0 ? void 0 : _l.answer) === 0) {
                        __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-4 h-4 ml-2 flex-shrink-0" }, { fill: "currentColor", viewBox: "0 0 20 20" }));
                        /** @type {__VLS_StyleScopedClasses['w-4']} */ ;
                        /** @type {__VLS_StyleScopedClasses['h-4']} */ ;
                        /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
                        /** @type {__VLS_StyleScopedClasses['flex-shrink-0']} */ ;
                        __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
                            'fill-rule': "evenodd",
                            d: "M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z",
                            'clip-rule': "evenodd",
                        });
                    }
                }
                else if (content.questionTypeId === 4) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "pl-4" }));
                    /** @type {__VLS_StyleScopedClasses['pl-4']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-gray-600" }));
                    /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
                    (__VLS_ctx.$t('examPaperPreview.gapFilling'));
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-2 text-blue-600 font-medium" }));
                    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-blue-600']} */ ;
                    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                    if (Array.isArray((_m = content.gapFilling) === null || _m === void 0 ? void 0 : _m.answer)) {
                        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                        (__VLS_ctx.$t('examPaperPreview.answers'));
                        __VLS_asFunctionalElement1(__VLS_intrinsics.ul, __VLS_intrinsics.ul)(__assign({ class: "list-decimal pl-5 mt-1 space-y-1" }));
                        /** @type {__VLS_StyleScopedClasses['list-decimal']} */ ;
                        /** @type {__VLS_StyleScopedClasses['pl-5']} */ ;
                        /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
                        /** @type {__VLS_StyleScopedClasses['space-y-1']} */ ;
                        for (var _2 = 0, _3 = __VLS_vFor((content.gapFilling.answer)); _2 < _3.length; _2++) {
                            var _4 = _3[_2], ans = _4[0], idx_1 = _4[1];
                            __VLS_asFunctionalElement1(__VLS_intrinsics.li, __VLS_intrinsics.li)({
                                key: (idx_1),
                            });
                            (ans);
                            // @ts-ignore
                            [$t, $t,];
                        }
                    }
                    else {
                        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                        (__VLS_ctx.$t('examPaperPreview.answer'));
                        ((_o = content.gapFilling) === null || _o === void 0 ? void 0 : _o.answer);
                    }
                }
                else if (content.questionTypeId === 5) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "pl-4" }));
                    /** @type {__VLS_StyleScopedClasses['pl-4']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-gray-600" }));
                    /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
                    (__VLS_ctx.$t('examPaperPreview.shortAnswer'));
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-2 space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-blue-600 font-medium" }));
                    /** @type {__VLS_StyleScopedClasses['text-blue-600']} */ ;
                    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                    (__VLS_ctx.$t('examPaperPreview.answer'));
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "pl-4 border-l-2 border-blue-200" }));
                    /** @type {__VLS_StyleScopedClasses['pl-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['border-l-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['border-blue-200']} */ ;
                    var __VLS_11 = index_vue_1.default || index_vue_1.default;
                    // @ts-ignore
                    var __VLS_12 = __VLS_asFunctionalComponent1(__VLS_11, new __VLS_11({
                        modelValue: ((_q = (_p = content.shortAnswer) === null || _p === void 0 ? void 0 : _p.answer) !== null && _q !== void 0 ? _q : ''),
                        readOnly: true,
                        height: "100%",
                    }));
                    var __VLS_13 = __VLS_12.apply(void 0, __spreadArray([{
                            modelValue: ((_s = (_r = content.shortAnswer) === null || _r === void 0 ? void 0 : _r.answer) !== null && _s !== void 0 ? _s : ''),
                            readOnly: true,
                            height: "100%",
                        }], __VLS_functionalComponentArgsRest(__VLS_12), false));
                }
                else {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "pl-4" }));
                    /** @type {__VLS_StyleScopedClasses['pl-4']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-gray-600" }));
                    /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
                    (__VLS_ctx.$t('examPaperPreview.otherType'));
                }
                // @ts-ignore
                [$t, $t, $t, $t,];
            }
        }
        // @ts-ignore
        [];
    }
}
// @ts-ignore
[];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
