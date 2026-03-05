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
var _a, _b, _c, _d, _e, _f, _g, _h, _j, _k;
Object.defineProperty(exports, "__esModule", { value: true });
var vue_1 = require("vue");
var vue_router_1 = require("vue-router");
var examPaperService_1 = require("../../services/examPaperService");
var syllabusService_1 = require("../../services/syllabusService");
var questionService_1 = require("../../services/questionService");
var organisationService_1 = require("../../services/organisationService");
var qualificationService_1 = require("../../services/qualificationService");
var chapterService_1 = require("../../services/chapterService");
var question_model_1 = require("../../models/question.model");
var index_vue_1 = require("../../components/QuillEditor/index.vue");
var ChapterOption_vue_1 = require("../../components/admin/ChapterOption.vue");
var router = (0, vue_router_1.useRouter)();
var route = (0, vue_router_1.useRoute)();
var isEdit = !!route.params.id;
var form = (0, vue_1.reactive)({
    name: '',
    syllabusId: 0,
    questionIds: []
});
var organisations = (0, vue_1.ref)([]);
var qualifications = (0, vue_1.ref)([]);
var syllabuses = (0, vue_1.ref)([]);
var chapters = (0, vue_1.ref)([]);
var chapterTree = (0, vue_1.ref)([]);
var showChapterSelector = (0, vue_1.ref)(false);
var selectedOrganisationId = (0, vue_1.ref)('');
var selectedQualificationId = (0, vue_1.ref)('');
var questions = (0, vue_1.ref)([]);
var totalQuestions = (0, vue_1.ref)(0);
var currentPage = (0, vue_1.ref)(1);
var pageSize = 12;
var totalPages = (0, vue_1.computed)(function () { return Math.ceil(totalQuestions.value / pageSize); });
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
// For read-only question part tabs
var questionActiveTabMap = (0, vue_1.ref)(new Map());
var setActiveTab = function (qid, idx) {
    questionActiveTabMap.value.set(qid, idx);
};
var getActiveTab = function (qid) {
    var _a;
    return (_a = questionActiveTabMap.value.get(qid)) !== null && _a !== void 0 ? _a : 0;
};
var questionFilter = (0, vue_1.reactive)({
    stem: '',
    difficult: '',
    chapterId: '',
    paperName: '',
});
var showSelected = (0, vue_1.ref)(false);
var selectedChapterIds = (0, vue_1.ref)([]);
// 存储所有已选qid对应的题目对象，分页切换不丢失
var selectedQuestionsMap = (0, vue_1.reactive)({});
// 监听form.questionIds变化，自动补全selectedQuestionsMap
(0, vue_1.watch)(function () { return __spreadArray([], form.questionIds, true); }, function (qids) { return __awaiter(void 0, void 0, void 0, function () {
    var _i, qids_1, qid, res, e_1;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                _i = 0, qids_1 = qids;
                _a.label = 1;
            case 1:
                if (!(_i < qids_1.length)) return [3 /*break*/, 6];
                qid = qids_1[_i];
                if (!!selectedQuestionsMap[qid]) return [3 /*break*/, 5];
                _a.label = 2;
            case 2:
                _a.trys.push([2, 4, , 5]);
                return [4 /*yield*/, questionService_1.default.getQuestionById(qid)];
            case 3:
                res = _a.sent();
                if (res.data)
                    selectedQuestionsMap[qid] = res.data;
                return [3 /*break*/, 5];
            case 4:
                e_1 = _a.sent();
                // 填充所有Question必需字段，防止类型报错
                selectedQuestionsMap[qid] = {
                    id: qid,
                    stem: 'Load failed',
                    questionContents: [],
                    syllabusId: 0,
                    totalScore: 0,
                    difficult: 1,
                    status: 1,
                    indexInPastPaper: 0
                };
                return [3 /*break*/, 5];
            case 5:
                _i++;
                return [3 /*break*/, 1];
            case 6:
                // 清理已移除的qid
                Object.keys(selectedQuestionsMap).forEach(function (k) {
                    if (!qids.includes(Number(k)))
                        delete selectedQuestionsMap[Number(k)];
                });
                return [2 /*return*/];
        }
    });
}); }, { immediate: true });
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
var fetchChapters = function () { return __awaiter(void 0, void 0, void 0, function () {
    var res;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (!form.syllabusId) {
                    chapters.value = [];
                    return [2 /*return*/];
                }
                return [4 /*yield*/, chapterService_1.default.getChapterTree(form.syllabusId)];
            case 1:
                res = _a.sent();
                chapters.value = res.data || [];
                return [2 /*return*/];
        }
    });
}); };
var fetchQuestions = function () { return __awaiter(void 0, void 0, void 0, function () {
    var res;
    var _a, _b, _c;
    return __generator(this, function (_d) {
        switch (_d.label) {
            case 0: return [4 /*yield*/, questionService_1.default.getQuestions({
                    stem: questionFilter.stem,
                    paperName: ((_a = questionFilter.paperName) === null || _a === void 0 ? void 0 : _a.trim()) || undefined,
                    difficult: questionFilter.difficult ? Number(questionFilter.difficult) : undefined,
                    syllabusId: form.syllabusId || undefined,
                    status: question_model_1.QUESTION_STATE_NORMAL, // Only fetch normal questions
                    pageIndex: currentPage.value,
                    pageSize: pageSize
                })];
            case 1:
                res = _d.sent();
                questions.value = ((_b = res.data) === null || _b === void 0 ? void 0 : _b.list) || [];
                totalQuestions.value = ((_c = res.data) === null || _c === void 0 ? void 0 : _c.total) || 0;
                return [2 /*return*/];
        }
    });
}); };
var addQuestion = function (id) {
    if (!form.questionIds.includes(id)) {
        form.questionIds.push(id);
    }
};
var removeQuestion = function (id) {
    form.questionIds = form.questionIds.filter(function (qid) { return qid !== id; });
};
var onDragStart = function (event, id) {
    if (event.dataTransfer) {
        event.dataTransfer.effectAllowed = 'move';
        event.dataTransfer.setData('text/plain', id.toString());
    }
};
var onDrop = function (event, targetId) {
    if (event.dataTransfer) {
        var sourceId = Number(event.dataTransfer.getData('text/plain'));
        var sourceIndex = form.questionIds.indexOf(sourceId);
        var targetIndex = form.questionIds.indexOf(targetId);
        if (sourceIndex !== -1 && targetIndex !== -1) {
            form.questionIds.splice(sourceIndex, 1);
            form.questionIds.splice(targetIndex, 0, sourceId);
        }
    }
};
var handleSubmit = function () { return __awaiter(void 0, void 0, void 0, function () {
    var req, req;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                console.log('Submitting form:', form);
                if (!form.name || !form.syllabusId || form.questionIds.length === 0) {
                    window.alert('Name, syllabus and at least one question are needed!');
                    return [2 /*return*/];
                }
                if (!(isEdit && form.id)) return [3 /*break*/, 2];
                req = {
                    id: form.id,
                    name: form.name,
                    syllabusId: form.syllabusId,
                    questionIds: form.questionIds
                };
                return [4 /*yield*/, examPaperService_1.examPaperService.updateExamPaper(req)];
            case 1:
                _a.sent();
                window.alert('Exam paper updated!');
                return [3 /*break*/, 4];
            case 2:
                req = {
                    name: form.name,
                    syllabusId: form.syllabusId,
                    questionIds: form.questionIds
                };
                return [4 /*yield*/, examPaperService_1.examPaperService.createExamPaper(req)];
            case 3:
                _a.sent();
                window.alert('Exam paper created!');
                _a.label = 4;
            case 4:
                router.back();
                return [2 /*return*/];
        }
    });
}); };
var handleCancel = function () {
    router.back();
};
var onOrganisationChange = function () { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                selectedQualificationId.value = '';
                form.syllabusId = 0;
                chapters.value = [];
                return [4 /*yield*/, fetchQualifications()];
            case 1:
                _a.sent();
                syllabuses.value = [];
                return [2 /*return*/];
        }
    });
}); };
var onQualificationChange = function () { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                form.syllabusId = 0;
                chapters.value = [];
                return [4 /*yield*/, fetchSyllabuses()];
            case 1:
                _a.sent();
                return [2 /*return*/];
        }
    });
}); };
var onSyllabusChange = function () { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, fetchChapters()
                // 同步章节树
            ];
            case 1:
                _a.sent();
                // 同步章节树
                chapterTree.value = chapters.value;
                return [4 /*yield*/, fetchQuestions()];
            case 2:
                _a.sent();
                return [2 /*return*/];
        }
    });
}); };
var goToPage = function (page) {
    if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
        currentPage.value = page;
        fetchQuestions();
    }
};
var paperNameDebounceTimer;
var onPaperNameInput = function () {
    clearTimeout(paperNameDebounceTimer);
    paperNameDebounceTimer = window.setTimeout(function () {
        currentPage.value = 1;
        fetchQuestions();
    }, 500);
};
(0, vue_1.onMounted)(function () { return __awaiter(void 0, void 0, void 0, function () {
    var res, paper, syllabusRes, syllabus, qualRes;
    var _a, _b, _c;
    return __generator(this, function (_d) {
        switch (_d.label) {
            case 0: return [4 /*yield*/, fetchOrganisations()];
            case 1:
                _d.sent();
                if (!(isEdit && route.params.id)) return [3 /*break*/, 8];
                return [4 /*yield*/, examPaperService_1.examPaperService.getExamPaperById({ id: Number(route.params.id) })];
            case 2:
                res = _d.sent();
                paper = res.data;
                form.id = paper.id;
                form.name = paper.name;
                form.syllabusId = paper.syllabusId;
                form.questionIds = Array.isArray(paper.questionIds) ? __spreadArray([], paper.questionIds, true) : [];
                return [4 /*yield*/, syllabusService_1.default.getSyllabusById(form.syllabusId)];
            case 3:
                syllabusRes = _d.sent();
                syllabus = syllabusRes.data;
                selectedQualificationId.value = syllabus.qualificationId;
                return [4 /*yield*/, qualificationService_1.default.getQualifications({ id: syllabus.qualificationId })];
            case 4:
                qualRes = _d.sent();
                selectedOrganisationId.value = ((_c = (_b = (_a = qualRes.data) === null || _a === void 0 ? void 0 : _a.list) === null || _b === void 0 ? void 0 : _b[0]) === null || _c === void 0 ? void 0 : _c.organisationId) || '';
                return [4 /*yield*/, fetchQualifications()];
            case 5:
                _d.sent();
                return [4 /*yield*/, fetchSyllabuses()];
            case 6:
                _d.sent();
                return [4 /*yield*/, fetchChapters()];
            case 7:
                _d.sent();
                _d.label = 8;
            case 8: return [4 /*yield*/, fetchQuestions()];
            case 9:
                _d.sent();
                return [2 /*return*/];
        }
    });
}); });
var __VLS_ctx = __assign(__assign({}, {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-6 w-full min-h-screen bg-gray-50 flex flex-col md:flex-row gap-6" }));
/** @type {__VLS_StyleScopedClasses['p-6']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['min-h-screen']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
/** @type {__VLS_StyleScopedClasses['md:flex-row']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.aside, __VLS_intrinsics.aside)(__assign({ class: "w-full md:w-72 flex-shrink-0 mb-6 md:mb-0" }));
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['md:w-72']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-shrink-0']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
/** @type {__VLS_StyleScopedClasses['md:mb-0']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-lg font-semibold text-gray-700 mb-4" }));
/** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
/** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
(__VLS_ctx.$t('examPaperForm.paperStructure'));
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4" }));
/** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-gray-600 mb-1" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
(__VLS_ctx.$t('examPaperForm.organisation'));
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ onChange: (__VLS_ctx.onOrganisationChange) }, { value: (__VLS_ctx.selectedOrganisationId) }), { class: "w-full p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500" }));
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['p-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-blue-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "",
    disabled: true,
});
(__VLS_ctx.$t('examPaperForm.selectOrganisation'));
for (var _i = 0, _l = __VLS_vFor((__VLS_ctx.organisations)); _i < _l.length; _i++) {
    var org = _l[_i][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (org.id),
        value: (org.id),
    });
    (org.name);
    // @ts-ignore
    [$t, $t, $t, onOrganisationChange, selectedOrganisationId, organisations,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4" }));
/** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-gray-600 mb-1" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
(__VLS_ctx.$t('examPaperForm.qualification'));
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ onChange: (__VLS_ctx.onQualificationChange) }, { value: (__VLS_ctx.selectedQualificationId) }), { class: "w-full p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500" }));
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['p-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-blue-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "",
    disabled: true,
});
(__VLS_ctx.$t('examPaperForm.selectQualification'));
for (var _m = 0, _o = __VLS_vFor((__VLS_ctx.qualifications)); _m < _o.length; _m++) {
    var q = _o[_m][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (q.id),
        value: (q.id),
    });
    (q.name);
    // @ts-ignore
    [$t, $t, onQualificationChange, selectedQualificationId, qualifications,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4" }));
/** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-gray-600 mb-1" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
(__VLS_ctx.$t('examPaperForm.syllabus'));
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ onChange: (__VLS_ctx.onSyllabusChange) }, { value: (__VLS_ctx.form.syllabusId) }), { class: "w-full p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500" }));
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['p-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-blue-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "",
    disabled: true,
});
(__VLS_ctx.$t('examPaperForm.selectSyllabus'));
for (var _p = 0, _q = __VLS_vFor((__VLS_ctx.syllabuses)); _p < _q.length; _p++) {
    var syllabus = _q[_p][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (syllabus.id),
        value: (syllabus.id),
    });
    (syllabus.name);
    (syllabus.code);
    // @ts-ignore
    [$t, $t, onSyllabusChange, form, syllabuses,];
}
if (__VLS_ctx.form.syllabusId) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-gray-600 mb-1" }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
    (__VLS_ctx.$t('examPaperForm.chapters'));
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
            if (!(__VLS_ctx.form.syllabusId))
                return;
            __VLS_ctx.showChapterSelector = !__VLS_ctx.showChapterSelector;
            // @ts-ignore
            [$t, form, showChapterSelector, showChapterSelector,];
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
    (__VLS_ctx.selectedChapterIds.length ? __VLS_ctx.$t('examPaperForm.chaptersSelected', { count: __VLS_ctx.selectedChapterIds.length }) : __VLS_ctx.$t('examPaperForm.allChapters'));
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
        (__VLS_ctx.$t('examPaperForm.selectChapters'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(__VLS_ctx.form.syllabusId))
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
        (__VLS_ctx.$t('examPaperForm.clear'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-2" }));
        /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
        for (var _r = 0, _s = __VLS_vFor((__VLS_ctx.chapterTree)); _r < _s.length; _r++) {
            var _t = _s[_r], chapter = _t[0], index = _t[1];
            var __VLS_0 = ChapterOption_vue_1.default;
            // @ts-ignore
            var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0(__assign({ 'onUpdate:selected': {} }, { key: (chapter.id), chapter: (chapter), level: (0), isLast: (index === __VLS_ctx.chapterTree.length - 1), selectedChapters: (__VLS_ctx.selectedChapterIds) })));
            var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([__assign({ 'onUpdate:selected': {} }, { key: (chapter.id), chapter: (chapter), level: (0), isLast: (index === __VLS_ctx.chapterTree.length - 1), selectedChapters: (__VLS_ctx.selectedChapterIds) })], __VLS_functionalComponentArgsRest(__VLS_1), false));
            var __VLS_5 = void 0;
            var __VLS_6 = ({ 'update:selected': {} },
                { 'onUpdate:selected': (function (val) { __VLS_ctx.selectedChapterIds = val; __VLS_ctx.fetchQuestions(); }) });
            var __VLS_3;
            var __VLS_4;
            // @ts-ignore
            [$t, chapterTree, chapterTree, selectedChapterIds, selectedChapterIds, fetchQuestions,];
        }
    }
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex-1" }));
/** @type {__VLS_StyleScopedClasses['flex-1']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.header, __VLS_intrinsics.header)(__assign({ class: "mb-6" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.h1, __VLS_intrinsics.h1)(__assign({ class: "text-3xl font-bold text-gray-900" }));
/** @type {__VLS_StyleScopedClasses['text-3xl']} */ ;
/** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
(__VLS_ctx.isEdit ? __VLS_ctx.$t('examPaperForm.editTitle') : __VLS_ctx.$t('examPaperForm.createTitle'));
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-sm text-gray-500" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
(__VLS_ctx.$t('examPaperForm.subtitle'));
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-2 bg-yellow-50 border-l-4 border-yellow-400 text-yellow-700 p-3 rounded" }));
/** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-yellow-50']} */ ;
/** @type {__VLS_StyleScopedClasses['border-l-4']} */ ;
/** @type {__VLS_StyleScopedClasses['border-yellow-400']} */ ;
/** @type {__VLS_StyleScopedClasses['text-yellow-700']} */ ;
/** @type {__VLS_StyleScopedClasses['p-3']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
(__VLS_ctx.$t('examPaperForm.knowledgePointNotice'));
__VLS_asFunctionalElement1(__VLS_intrinsics.form, __VLS_intrinsics.form)(__assign({ onSubmit: (__VLS_ctx.handleSubmit) }, { class: "mb-6 space-y-4 max-w-2xl" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
/** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
/** @type {__VLS_StyleScopedClasses['max-w-2xl']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
(__VLS_ctx.$t('examPaperForm.paperName'));
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ value: (__VLS_ctx.form.name), type: "text", required: true, placeholder: (__VLS_ctx.$t('examPaperForm.paperNamePlaceholder')) }, { class: "w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500 sm:text-sm" }));
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex flex-row gap-2 items-center" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-row']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex-1" }));
/** @type {__VLS_StyleScopedClasses['flex-1']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex gap-2" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ type: "submit" }, { class: "px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition" }));
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-blue-600']} */ ;
/** @type {__VLS_StyleScopedClasses['text-white']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-blue-700']} */ ;
/** @type {__VLS_StyleScopedClasses['transition']} */ ;
(__VLS_ctx.$t('examPaperForm.save'));
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: (__VLS_ctx.handleCancel) }, { type: "button" }), { class: "px-4 py-2 bg-gray-200 text-gray-700 rounded-md hover:bg-gray-300 transition" }));
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-200']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['transition']} */ ;
(__VLS_ctx.$t('examPaperForm.cancel'));
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-6 flex flex-col lg:flex-row gap-4" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
/** @type {__VLS_StyleScopedClasses['lg:flex-row']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign(__assign({ onInput: (__VLS_ctx.fetchQuestions) }, { placeholder: (__VLS_ctx.$t('examPaperForm.searchStem')) }), { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full lg:w-64" }));
(__VLS_ctx.questionFilter.stem);
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
/** @type {__VLS_StyleScopedClasses['lg:w-64']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign(__assign({ onInput: (__VLS_ctx.onPaperNameInput) }, { placeholder: (__VLS_ctx.$t('examPaperForm.searchByPaperName')) }), { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full lg:w-64" }));
(__VLS_ctx.questionFilter.paperName);
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
/** @type {__VLS_StyleScopedClasses['lg:w-64']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ onChange: (__VLS_ctx.fetchQuestions) }, { value: (__VLS_ctx.questionFilter.difficult) }), { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full lg:w-32" }));
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
/** @type {__VLS_StyleScopedClasses['lg:w-32']} */ ;
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
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow-md border border-gray-200 p-4" }));
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-md']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
/** @type {__VLS_StyleScopedClasses['p-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-between items-center mb-4" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-600" }));
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
(__VLS_ctx.$t('examPaperForm.showingQuestions', { count: __VLS_ctx.questions.length }));
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
        var _a = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            _a[_i] = arguments[_i];
        }
        var $event = _a[0];
        __VLS_ctx.showSelected = !__VLS_ctx.showSelected;
        // @ts-ignore
        [$t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, form, fetchQuestions, fetchQuestions, isEdit, handleSubmit, handleCancel, questionFilter, questionFilter, questionFilter, onPaperNameInput, questions, showSelected, showSelected,];
    } }, { class: "relative flex items-center px-3 py-2 bg-blue-600 text-white rounded-full shadow hover:bg-blue-700 transition" }));
/** @type {__VLS_StyleScopedClasses['relative']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-blue-600']} */ ;
/** @type {__VLS_StyleScopedClasses['text-white']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-blue-700']} */ ;
/** @type {__VLS_StyleScopedClasses['transition']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-5 h-5 mr-1" }, { fill: "none", stroke: "currentColor", 'stroke-width': "2", viewBox: "0 0 24 24" }));
/** @type {__VLS_StyleScopedClasses['w-5']} */ ;
/** @type {__VLS_StyleScopedClasses['h-5']} */ ;
/** @type {__VLS_StyleScopedClasses['mr-1']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.path)({
    'stroke-linecap': "round",
    'stroke-linejoin': "round",
    d: "M5 13l4 4L19 7",
});
(__VLS_ctx.$t('examPaperForm.selectedQuestions', { count: __VLS_ctx.form.questionIds.length }));
if (__VLS_ctx.questions.length === 0) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center text-gray-400 py-8" }));
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-8']} */ ;
    (__VLS_ctx.$t('examPaperForm.noQuestions'));
}
else {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "grid grid-cols-1 gap-4" }));
    /** @type {__VLS_StyleScopedClasses['grid']} */ ;
    /** @type {__VLS_StyleScopedClasses['grid-cols-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
    var _loop_1 = function (q) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (q.id) }, { class: "bg-white rounded-lg shadow-md border border-gray-200 hover:shadow-lg transition-shadow duration-200" }));
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
        (__VLS_ctx.$t('examPaperForm.id'));
        (q.id);
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800" }));
        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-green-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-green-800']} */ ;
        (__VLS_ctx.$t('examPaperForm.totalScore'));
        (q.totalScore || 0);
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-yellow-100 text-yellow-800" }));
        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-yellow-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-yellow-800']} */ ;
        (__VLS_ctx.DIFFICULTY_NAMES[q.difficult] || __VLS_ctx.$t('examPaperForm.unknown'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-4" }));
        /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-3" }));
        /** @type {__VLS_StyleScopedClasses['mb-3']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.h4, __VLS_intrinsics.h4)(__assign({ class: "text-sm font-medium text-gray-700 mb-1" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
        (__VLS_ctx.$t('examPaperForm.question'));
        var __VLS_7 = index_vue_1.default || index_vue_1.default;
        // @ts-ignore
        var __VLS_8 = __VLS_asFunctionalComponent1(__VLS_7, new __VLS_7({
            modelValue: (q.stem),
            readOnly: (true),
            height: "100%",
            placeholder: "",
        }));
        var __VLS_9 = __VLS_8.apply(void 0, __spreadArray([{
                modelValue: (q.stem),
                readOnly: (true),
                height: "100%",
                placeholder: "",
            }], __VLS_functionalComponentArgsRest(__VLS_8), false));
        if (q.questionContents && q.questionContents.length > 0) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-3" }));
            /** @type {__VLS_StyleScopedClasses['mb-3']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.h4, __VLS_intrinsics.h4)(__assign({ class: "text-sm font-medium text-gray-700 mb-2" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
            /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
            (__VLS_ctx.$t('examPaperForm.parts'));
            (__VLS_ctx.$t('examPaperForm.totalScore'));
            (q.totalScore);
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
            var _loop_4 = function (content, index) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign(__assign({ onClick: (function () { return __VLS_ctx.setActiveTab(q.id, index); }) }, { key: (index) }), { class: ([
                        'px-4 py-2 text-sm font-medium border-r border-gray-200 whitespace-nowrap transition',
                        __VLS_ctx.getActiveTab(q.id) === index
                            ? 'text-blue-600 bg-white border-b-2 border-b-blue-500 -mb-px'
                            : 'text-gray-700 hover:text-blue-600 hover:bg-gray-100'
                    ]) }), { type: "button" }));
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
                [$t, $t, $t, $t, $t, $t, $t, $t, form, questions, questions, question_model_1.DIFFICULTY_NAMES, question_model_1.DIFFICULTY_NAMES, setActiveTab, getActiveTab,];
            };
            for (var _0 = 0, _1 = __VLS_vFor((q.questionContents)); _0 < _1.length; _0++) {
                var _2 = _1[_0], content = _2[0], index = _2[1];
                _loop_4(content, index);
            }
            if (q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0]) {
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
                (__VLS_ctx.QUESTION_TYPE_NAMES[q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].questionTypeId] || __VLS_ctx.$t('examPaperForm.unknown'));
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-xs text-gray-400" }));
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-xs text-gray-500" }));
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                (__VLS_ctx.$t('examPaperForm.score'));
                (q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].score);
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-3 text-sm text-gray-600" }));
                /** @type {__VLS_StyleScopedClasses['mt-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
                if (q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].questionTypeId === __VLS_ctx.QUESTION_TYPE_SINGLE_CHOICE) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    for (var _3 = 0, _4 = __VLS_vFor(((_a = q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].singleChoice) === null || _a === void 0 ? void 0 : _a.options)); _3 < _4.length; _3++) {
                        var option = _4[_3][0];
                        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (option.prefix) }, { class: "flex gap-2" }));
                        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                        /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-gray-500" }));
                        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                        (option.prefix);
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
                        __VLS_asFunctionalDirective(__VLS_directives.vHtml, {})(null, __assign(__assign({}, __VLS_directiveBindingRestFields), { value: (option.content) }), null, null);
                        // @ts-ignore
                        [$t, $t, $t, getActiveTab, getActiveTab, getActiveTab, getActiveTab, getActiveTab, question_model_1.QUESTION_TYPE_NAMES, question_model_1.QUESTION_TYPE_NAMES, question_model_1.QUESTION_TYPE_SINGLE_CHOICE,];
                    }
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-2" }));
                    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-semibold" }));
                    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
                    (__VLS_ctx.$t('examPaperForm.correctAnswer'));
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "ml-2 text-blue-600" }));
                    /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-blue-600']} */ ;
                    ((_b = q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].singleChoice) === null || _b === void 0 ? void 0 : _b.answer);
                }
                else if (q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].questionTypeId === __VLS_ctx.QUESTION_TYPE_MULTIPLE_CHOICE) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    for (var _5 = 0, _6 = __VLS_vFor(((_c = q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].multipleChoice) === null || _c === void 0 ? void 0 : _c.options)); _5 < _6.length; _5++) {
                        var option = _6[_5][0];
                        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (option.prefix) }, { class: "flex gap-2" }));
                        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                        /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-gray-500" }));
                        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                        (option.prefix);
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
                        __VLS_asFunctionalDirective(__VLS_directives.vHtml, {})(null, __assign(__assign({}, __VLS_directiveBindingRestFields), { value: (option.content) }), null, null);
                        // @ts-ignore
                        [$t, getActiveTab, getActiveTab, getActiveTab, question_model_1.QUESTION_TYPE_MULTIPLE_CHOICE,];
                    }
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-2" }));
                    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-semibold" }));
                    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "ml-2 text-blue-600" }));
                    /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-blue-600']} */ ;
                    ((((_d = q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].multipleChoice) === null || _d === void 0 ? void 0 : _d.answer) || []).join(', '));
                }
                else if (q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].questionTypeId === __VLS_ctx.QUESTION_TYPE_TRUE_FALSE) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-semibold" }));
                    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "ml-2" }));
                    /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
                    (String((_e = q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].trueOrFalse) === null || _e === void 0 ? void 0 : _e.answer).toLowerCase() === 'true'
                        ? 'True'
                        : 'False');
                }
                else if (q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].questionTypeId === __VLS_ctx.QUESTION_TYPE_GAP_FILLING) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    for (var _7 = 0, _8 = __VLS_vFor(((_f = q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].gapFilling) === null || _f === void 0 ? void 0 : _f.answer)); _7 < _8.length; _7++) {
                        var _9 = _8[_7], ans = _9[0], gidx = _9[1];
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
                        [$t, $t, getActiveTab, getActiveTab, getActiveTab, getActiveTab, getActiveTab, question_model_1.QUESTION_TYPE_TRUE_FALSE, question_model_1.QUESTION_TYPE_GAP_FILLING,];
                    }
                }
                else if (q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].questionTypeId === __VLS_ctx.QUESTION_TYPE_SHORT_ANSWER) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-2" }));
                    /** @type {__VLS_StyleScopedClasses['space-y-2']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-semibold" }));
                    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
                    var __VLS_12 = index_vue_1.default || index_vue_1.default;
                    // @ts-ignore
                    var __VLS_13 = __VLS_asFunctionalComponent1(__VLS_12, new __VLS_12({
                        modelValue: (((_g = q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].shortAnswer) === null || _g === void 0 ? void 0 : _g.answer) || ''),
                        readOnly: true,
                        height: "100%",
                    }));
                    var __VLS_14 = __VLS_13.apply(void 0, __spreadArray([{
                            modelValue: (((_h = q.questionContents[__VLS_ctx.getActiveTab(q.id) || 0].shortAnswer) === null || _h === void 0 ? void 0 : _h.answer) || ''),
                            readOnly: true,
                            height: "100%",
                        }], __VLS_functionalComponentArgsRest(__VLS_13), false));
                }
                else {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "italic" }));
                    /** @type {__VLS_StyleScopedClasses['italic']} */ ;
                    (__VLS_ctx.$t('examPaperForm.pleaseViewFull'));
                }
            }
        }
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-3" }));
        /** @type {__VLS_StyleScopedClasses['mb-3']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.h4, __VLS_intrinsics.h4)(__assign({ class: "text-sm font-medium text-gray-700 mb-1" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm text-gray-600" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
        (((_j = q.syllabus) === null || _j === void 0 ? void 0 : _j.name) || '-');
        __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-xs text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (((_k = q.syllabus) === null || _k === void 0 ? void 0 : _k.code) || '');
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "px-4 py-3 bg-gray-50 border-t border-gray-200 flex justify-end" }));
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-t']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['justify-end']} */ ;
        if (!__VLS_ctx.form.questionIds.includes(q.id)) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!!(__VLS_ctx.questions.length === 0))
                        return;
                    if (!(!__VLS_ctx.form.questionIds.includes(q.id)))
                        return;
                    __VLS_ctx.addQuestion(q.id);
                    // @ts-ignore
                    [$t, form, getActiveTab, getActiveTab, question_model_1.QUESTION_TYPE_SHORT_ANSWER, addQuestion,];
                } }, { class: "inline-flex items-center px-3 py-1.5 border border-transparent text-sm font-medium rounded text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500" }));
            /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
            /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
            /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-1.5']} */ ;
            /** @type {__VLS_StyleScopedClasses['border']} */ ;
            /** @type {__VLS_StyleScopedClasses['border-transparent']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
            /** @type {__VLS_StyleScopedClasses['bg-blue-600']} */ ;
            /** @type {__VLS_StyleScopedClasses['hover:bg-blue-700']} */ ;
            /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
            /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['focus:ring-offset-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
            (__VLS_ctx.$t('examPaperForm.addToPaper'));
        }
        else {
            __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-flex items-center px-3 py-1.5 text-sm font-medium text-green-600" }));
            /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
            /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
            /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-1.5']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-green-600']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-5 h-5 mr-1" }, { fill: "none", stroke: "currentColor", viewBox: "0 0 24 24" }));
            /** @type {__VLS_StyleScopedClasses['w-5']} */ ;
            /** @type {__VLS_StyleScopedClasses['h-5']} */ ;
            /** @type {__VLS_StyleScopedClasses['mr-1']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
                'stroke-linecap': "round",
                'stroke-linejoin': "round",
                'stroke-width': "2",
                d: "M5 13l4 4L19 7",
            });
            (__VLS_ctx.$t('examPaperForm.added'));
        }
        // @ts-ignore
        [$t, $t,];
    };
    for (var _u = 0, _v = __VLS_vFor((__VLS_ctx.questions)); _u < _v.length; _u++) {
        var q = _v[_u][0];
        _loop_1(q);
    }
}
if (__VLS_ctx.totalQuestions > 0) {
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
    (__VLS_ctx.$t('examPaperForm.paginationInfo', { from: ((__VLS_ctx.currentPage - 1) * __VLS_ctx.pageSize) + 1, to: Math.min(__VLS_ctx.currentPage * __VLS_ctx.pageSize, __VLS_ctx.totalQuestions), total: __VLS_ctx.totalQuestions }));
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
                if (!(__VLS_ctx.totalQuestions > 0))
                    return;
                if (!(__VLS_ctx.totalPages > 1))
                    return;
                __VLS_ctx.goToPage(__VLS_ctx.currentPage - 1);
                // @ts-ignore
                [$t, totalQuestions, totalQuestions, totalQuestions, currentPage, currentPage, currentPage, pageSize, pageSize, totalPages, goToPage,];
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
        (__VLS_ctx.$t('examPaperForm.previous'));
        var _loop_2 = function (page) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!(__VLS_ctx.totalQuestions > 0))
                        return;
                    if (!(__VLS_ctx.totalPages > 1))
                        return;
                    __VLS_ctx.goToPage(page);
                    // @ts-ignore
                    [$t, currentPage, goToPage, paginationRange,];
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
        for (var _w = 0, _x = __VLS_vFor((__VLS_ctx.paginationRange)); _w < _x.length; _w++) {
            var page = _x[_w][0];
            _loop_2(page);
        }
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(__VLS_ctx.totalQuestions > 0))
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
        (__VLS_ctx.$t('examPaperForm.next'));
    }
}
var __VLS_17;
/** @ts-ignore @type {typeof __VLS_components.transition | typeof __VLS_components.Transition | typeof __VLS_components.transition | typeof __VLS_components.Transition} */
transition;
// @ts-ignore
var __VLS_18 = __VLS_asFunctionalComponent1(__VLS_17, new __VLS_17({
    name: "fade",
}));
var __VLS_19 = __VLS_18.apply(void 0, __spreadArray([{
        name: "fade",
    }], __VLS_functionalComponentArgsRest(__VLS_18), false));
var __VLS_22 = __VLS_20.slots.default;
if (__VLS_ctx.showSelected) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "fixed bottom-6 right-6 z-50 max-w-full bg-white border border-gray-200 rounded-lg shadow-lg p-4" }));
    /** @type {__VLS_StyleScopedClasses['fixed']} */ ;
    /** @type {__VLS_StyleScopedClasses['bottom-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['right-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['z-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['max-w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-between items-center mb-2" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.h3, __VLS_intrinsics.h3)(__assign({ class: "font-semibold text-gray-700" }));
    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    (__VLS_ctx.form.questionIds.length);
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.showSelected))
                return;
            __VLS_ctx.showSelected = false;
            // @ts-ignore
            [$t, form, showSelected, showSelected, currentPage, totalPages,];
        } }, { class: "text-gray-400 hover:text-gray-700" }));
    /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-gray-700']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-5 h-5" }, { fill: "none", stroke: "currentColor", 'stroke-width': "2", viewBox: "0 0 24 24" }));
    /** @type {__VLS_StyleScopedClasses['w-5']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-5']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
        'stroke-linecap': "round",
        'stroke-linejoin': "round",
        d: "M6 18L18 6M6 6l12 12",
    });
    __VLS_asFunctionalElement1(__VLS_intrinsics.ul, __VLS_intrinsics.ul)(__assign({ class: "divide-y divide-gray-100 overflow-y-auto max-h-50%" }));
    /** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
    /** @type {__VLS_StyleScopedClasses['divide-gray-100']} */ ;
    /** @type {__VLS_StyleScopedClasses['overflow-y-auto']} */ ;
    /** @type {__VLS_StyleScopedClasses['max-h-50%']} */ ;
    var _loop_3 = function (qid) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.li, __VLS_intrinsics.li)(__assign(__assign(__assign(__assign(__assign(__assign({ onDragstart: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(__VLS_ctx.showSelected))
                    return;
                __VLS_ctx.onDragStart($event, qid);
                // @ts-ignore
                [form, onDragStart,];
            } }, { onDragover: function () { } }), { onDragenter: function () { } }), { onDrop: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(__VLS_ctx.showSelected))
                    return;
                __VLS_ctx.onDrop($event, qid);
                // @ts-ignore
                [onDrop,];
            } }), { key: (qid) }), { class: "border-b border-gray-100 px-2 py-3 cursor-move" }), { draggable: "true" }));
        /** @type {__VLS_StyleScopedClasses['border-b']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['cursor-move']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex gap-2 items-start" }));
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-start']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex-shrink-0 text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['flex-shrink-0']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (qid);
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex-1" }));
        /** @type {__VLS_StyleScopedClasses['flex-1']} */ ;
        if (__VLS_ctx.selectedQuestionsMap[qid]) {
            var __VLS_23 = index_vue_1.default;
            // @ts-ignore
            var __VLS_24 = __VLS_asFunctionalComponent1(__VLS_23, new __VLS_23({
                modelValue: (__VLS_ctx.selectedQuestionsMap[qid].stem),
                readOnly: (true),
                height: "100%",
                placeholder: "",
            }));
            var __VLS_25 = __VLS_24.apply(void 0, __spreadArray([{
                    modelValue: (__VLS_ctx.selectedQuestionsMap[qid].stem),
                    readOnly: (true),
                    height: "100%",
                    placeholder: "",
                }], __VLS_functionalComponentArgsRest(__VLS_24), false));
        }
        else {
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-gray-400 text-xs" }));
            /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        }
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(__VLS_ctx.showSelected))
                    return;
                __VLS_ctx.removeQuestion(qid);
                // @ts-ignore
                [selectedQuestionsMap, selectedQuestionsMap, removeQuestion,];
            } }, { type: "button" }), { class: "text-red-500 hover:underline ml-2 flex-shrink-0" }));
        /** @type {__VLS_StyleScopedClasses['text-red-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:underline']} */ ;
        /** @type {__VLS_StyleScopedClasses['ml-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['flex-shrink-0']} */ ;
        // @ts-ignore
        [];
    };
    for (var _y = 0, _z = __VLS_vFor((__VLS_ctx.form.questionIds)); _y < _z.length; _y++) {
        var qid = _z[_y][0];
        _loop_3(qid);
    }
}
// @ts-ignore
[];
var __VLS_20;
// @ts-ignore
[];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
