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
var _a, _b, _c, _d;
Object.defineProperty(exports, "__esModule", { value: true });
var vue_1 = require("vue");
var vue_router_1 = require("vue-router");
var vue_i18n_1 = require("vue-i18n");
var learningPlanService_1 = require("../../services/learningPlanService");
var classService_1 = require("../../services/classService");
var examNodeService_1 = require("../../services/examNodeService");
var t = (0, vue_i18n_1.useI18n)().t;
var route = (0, vue_router_1.useRoute)();
var classId = Number(route.params.classId);
var activeTab = (0, vue_1.ref)('plans');
var plans = (0, vue_1.ref)([]);
var classInfo = (0, vue_1.ref)(null);
var versions = (0, vue_1.ref)([]);
var loading = (0, vue_1.ref)(false);
var saving = (0, vue_1.ref)(false);
var loadingVersions = (0, vue_1.ref)(false);
var filterUserId = (0, vue_1.ref)('');
var filterPlanType = (0, vue_1.ref)('');
var filterIsPersonal = (0, vue_1.ref)(null);
var showModal = (0, vue_1.ref)(false);
var showVersionsModal = (0, vue_1.ref)(false);
var showDeleteModal = (0, vue_1.ref)(false);
var editingPlan = (0, vue_1.ref)(null);
var versionsPlan = (0, vue_1.ref)(null);
var deletingPlan = (0, vue_1.ref)(null);
var form = (0, vue_1.ref)({ userId: 0, planType: 'long', content: '', comment: '', isPersonal: false });
// Students tab
var students = (0, vue_1.ref)([]);
var loadingStudents = (0, vue_1.ref)(false);
function loadStudents() {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_1;
        var _a;
        return __generator(this, function (_b) {
            switch (_b.label) {
                case 0:
                    loadingStudents.value = true;
                    _b.label = 1;
                case 1:
                    _b.trys.push([1, 3, 4, 5]);
                    return [4 /*yield*/, classService_1.default.listStudents(classId)];
                case 2:
                    res = _b.sent();
                    if (res.code === 0)
                        students.value = (_a = res.data.list) !== null && _a !== void 0 ? _a : [];
                    return [3 /*break*/, 5];
                case 3:
                    e_1 = _b.sent();
                    console.error(e_1);
                    return [3 /*break*/, 5];
                case 4:
                    loadingStudents.value = false;
                    return [7 /*endfinally*/];
                case 5: return [2 /*return*/];
            }
        });
    });
}
function openCreateModalForStudent(student) {
    editingPlan.value = null;
    form.value = { userId: student.id, planType: 'long', content: '', comment: '', isPersonal: true };
    showModal.value = true;
    activeTab.value = 'plans';
}
function loadPlans() {
    return __awaiter(this, void 0, void 0, function () {
        var query, res, e_2;
        var _a;
        return __generator(this, function (_b) {
            switch (_b.label) {
                case 0:
                    loading.value = true;
                    _b.label = 1;
                case 1:
                    _b.trys.push([1, 3, 4, 5]);
                    query = { classId: classId, pageSize: 100, pageIndex: 1 };
                    if (filterUserId.value)
                        query.userId = filterUserId.value;
                    if (filterPlanType.value)
                        query.planType = filterPlanType.value;
                    if (filterIsPersonal.value !== null)
                        query.isPersonal = filterIsPersonal.value;
                    return [4 /*yield*/, learningPlanService_1.default.list(query)];
                case 2:
                    res = _b.sent();
                    if (res.code === 0)
                        plans.value = (_a = res.data.list) !== null && _a !== void 0 ? _a : [];
                    return [3 /*break*/, 5];
                case 3:
                    e_2 = _b.sent();
                    console.error(e_2);
                    return [3 /*break*/, 5];
                case 4:
                    loading.value = false;
                    return [7 /*endfinally*/];
                case 5: return [2 /*return*/];
            }
        });
    });
}
function loadClassInfo() {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_3;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    _a.trys.push([0, 2, , 3]);
                    return [4 /*yield*/, classService_1.default.getById(classId)];
                case 1:
                    res = _a.sent();
                    if (res.code === 0)
                        classInfo.value = res.data;
                    return [3 /*break*/, 3];
                case 2:
                    e_3 = _a.sent();
                    console.error(e_3);
                    return [3 /*break*/, 3];
                case 3: return [2 /*return*/];
            }
        });
    });
}
function openCreateModal() {
    editingPlan.value = null;
    form.value = { userId: 0, planType: 'long', content: '', comment: '', isPersonal: false };
    showModal.value = true;
}
function openEditModal(plan) {
    editingPlan.value = plan;
    form.value = { userId: plan.userId, planType: plan.planType, content: plan.content, comment: '', isPersonal: plan.isPersonal };
    showModal.value = true;
}
function closeModal() {
    showModal.value = false;
}
function submitForm() {
    return __awaiter(this, void 0, void 0, function () {
        var res, res, e_4;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    saving.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 9, 10, 11]);
                    if (!editingPlan.value) return [3 /*break*/, 5];
                    return [4 /*yield*/, learningPlanService_1.default.update({ id: editingPlan.value.id, content: form.value.content, comment: form.value.comment })];
                case 2:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 4];
                    return [4 /*yield*/, loadPlans()];
                case 3:
                    _a.sent();
                    closeModal();
                    _a.label = 4;
                case 4: return [3 /*break*/, 8];
                case 5: return [4 /*yield*/, learningPlanService_1.default.create({ classId: classId, userId: form.value.userId, planType: form.value.planType, content: form.value.content, comment: form.value.comment, isPersonal: form.value.isPersonal })];
                case 6:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 8];
                    return [4 /*yield*/, loadPlans()];
                case 7:
                    _a.sent();
                    closeModal();
                    _a.label = 8;
                case 8: return [3 /*break*/, 11];
                case 9:
                    e_4 = _a.sent();
                    console.error(e_4);
                    return [3 /*break*/, 11];
                case 10:
                    saving.value = false;
                    return [7 /*endfinally*/];
                case 11: return [2 /*return*/];
            }
        });
    });
}
function openVersionsModal(plan) {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_5;
        var _a;
        return __generator(this, function (_b) {
            switch (_b.label) {
                case 0:
                    versionsPlan.value = plan;
                    showVersionsModal.value = true;
                    loadingVersions.value = true;
                    versions.value = [];
                    _b.label = 1;
                case 1:
                    _b.trys.push([1, 3, 4, 5]);
                    return [4 /*yield*/, learningPlanService_1.default.versions(plan.id)];
                case 2:
                    res = _b.sent();
                    if (res.code === 0)
                        versions.value = (_a = res.data.list) !== null && _a !== void 0 ? _a : [];
                    return [3 /*break*/, 5];
                case 3:
                    e_5 = _b.sent();
                    console.error(e_5);
                    return [3 /*break*/, 5];
                case 4:
                    loadingVersions.value = false;
                    return [7 /*endfinally*/];
                case 5: return [2 /*return*/];
            }
        });
    });
}
function doRollback(v) {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_6;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    if (!versionsPlan.value)
                        return [2 /*return*/];
                    saving.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 5, 6, 7]);
                    return [4 /*yield*/, learningPlanService_1.default.rollback(versionsPlan.value.id, v.version, "Rollback to v".concat(v.version))];
                case 2:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 4];
                    return [4 /*yield*/, loadPlans()];
                case 3:
                    _a.sent();
                    showVersionsModal.value = false;
                    _a.label = 4;
                case 4: return [3 /*break*/, 7];
                case 5:
                    e_6 = _a.sent();
                    console.error(e_6);
                    return [3 /*break*/, 7];
                case 6:
                    saving.value = false;
                    return [7 /*endfinally*/];
                case 7: return [2 /*return*/];
            }
        });
    });
}
function confirmDelete(plan) {
    deletingPlan.value = plan;
    showDeleteModal.value = true;
}
function doDelete() {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_7;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    if (!deletingPlan.value)
                        return [2 /*return*/];
                    saving.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 5, 6, 7]);
                    return [4 /*yield*/, learningPlanService_1.default.delete(deletingPlan.value.id)];
                case 2:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 4];
                    return [4 /*yield*/, loadPlans()];
                case 3:
                    _a.sent();
                    showDeleteModal.value = false;
                    _a.label = 4;
                case 4: return [3 /*break*/, 7];
                case 5:
                    e_7 = _a.sent();
                    console.error(e_7);
                    return [3 /*break*/, 7];
                case 6:
                    saving.value = false;
                    return [7 /*endfinally*/];
                case 7: return [2 /*return*/];
            }
        });
    });
}
var showGenerateModal = (0, vue_1.ref)(false);
var generating = (0, vue_1.ref)(false);
var generateResult = (0, vue_1.ref)(null);
var syllabusExamNodes = (0, vue_1.ref)([]);
var genForm = (0, vue_1.ref)({
    phaseRatios: [30, 20, 20, 10],
    examNodes: [],
    comment: ''
});
var phaseLabels = ['新课学习(%)', '一轮复习(%)', '专题综合(%)', '集中刷题(%)'];
function openGenerateModal() {
    return __awaiter(this, void 0, void 0, function () {
        var res, nodes, e_8;
        var _a, _b, _c;
        return __generator(this, function (_d) {
            switch (_d.label) {
                case 0:
                    generateResult.value = null;
                    // Load exam nodes for the bound syllabus.
                    syllabusExamNodes.value = [];
                    if (!((_a = classInfo.value) === null || _a === void 0 ? void 0 : _a.syllabusId)) return [3 /*break*/, 4];
                    _d.label = 1;
                case 1:
                    _d.trys.push([1, 3, , 4]);
                    return [4 /*yield*/, examNodeService_1.default.list(classInfo.value.syllabusId)];
                case 2:
                    res = _d.sent();
                    if (res.code === 0) {
                        nodes = ((_c = (_b = res.data) === null || _b === void 0 ? void 0 : _b.list) !== null && _c !== void 0 ? _c : []).map(function (n) { return ({ id: n.id, name: n.name }); });
                        syllabusExamNodes.value = nodes;
                        // Pre-fill examNodes with empty time ranges.
                        genForm.value.examNodes = nodes.map(function (n) { return ({
                            examNodeId: n.id,
                            name: n.name,
                            startMonth: '',
                            endMonth: ''
                        }); });
                    }
                    return [3 /*break*/, 4];
                case 3:
                    e_8 = _d.sent();
                    console.error(e_8);
                    return [3 /*break*/, 4];
                case 4:
                    showGenerateModal.value = true;
                    return [2 /*return*/];
            }
        });
    });
}
function closeGenerateModal() {
    showGenerateModal.value = false;
    if (generateResult.value)
        loadPlans();
}
function submitGenerate() {
    return __awaiter(this, void 0, void 0, function () {
        var hasAllTimes, res, e_9;
        var _a;
        return __generator(this, function (_b) {
            switch (_b.label) {
                case 0:
                    hasAllTimes = genForm.value.examNodes.length > 0 &&
                        genForm.value.examNodes.every(function (n) { return n.startMonth && n.endMonth; });
                    if (!hasAllTimes) {
                        alert(t('learningPlan.examNodesTimeRequired'));
                        return [2 /*return*/];
                    }
                    if (!((_a = classInfo.value) === null || _a === void 0 ? void 0 : _a.syllabusId)) {
                        alert(t('learningPlan.noSyllabus'));
                        return [2 /*return*/];
                    }
                    generating.value = true;
                    generateResult.value = null;
                    _b.label = 1;
                case 1:
                    _b.trys.push([1, 3, 4, 5]);
                    return [4 /*yield*/, learningPlanService_1.default.generateTemplate({
                            classId: classId,
                            syllabusId: classInfo.value.syllabusId,
                            phaseRatios: genForm.value.phaseRatios,
                            examNodes: genForm.value.examNodes.map(function (n) { return ({
                                examNodeId: n.examNodeId,
                                startMonth: n.startMonth,
                                endMonth: n.endMonth
                            }); }),
                            comment: genForm.value.comment
                        })];
                case 2:
                    res = _b.sent();
                    if (res.code === 0)
                        generateResult.value = res.data;
                    return [3 /*break*/, 5];
                case 3:
                    e_9 = _b.sent();
                    console.error(e_9);
                    return [3 /*break*/, 5];
                case 4:
                    generating.value = false;
                    return [7 /*endfinally*/];
                case 5: return [2 /*return*/];
            }
        });
    });
}
(0, vue_1.onMounted)(function () {
    loadPlans();
    loadClassInfo();
});
var __VLS_ctx = __assign(__assign({}, {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-6" }));
/** @type {__VLS_StyleScopedClasses['p-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.header, __VLS_intrinsics.header)(__assign({ class: "mb-6" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-between items-start" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
/** @type {__VLS_StyleScopedClasses['items-start']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.h1, __VLS_intrinsics.h1)(__assign({ class: "text-3xl font-bold text-gray-900" }));
/** @type {__VLS_StyleScopedClasses['text-3xl']} */ ;
/** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
(__VLS_ctx.t('learningPlan.title'));
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-2 text-sm text-gray-600" }));
/** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
(__VLS_ctx.t('learningPlan.subtitle'));
if (__VLS_ctx.classInfo) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-3 text-sm text-gray-700 font-medium" }));
    /** @type {__VLS_StyleScopedClasses['mt-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    (__VLS_ctx.classInfo.name);
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex gap-2" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
if (__VLS_ctx.activeTab === 'plans') {
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.openGenerateModal) }, { class: "inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-green-600 hover:bg-green-700 focus:outline-none" }));
    /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-transparent']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-green-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-green-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
    (__VLS_ctx.t('learningPlan.generateTemplate'));
}
if (__VLS_ctx.activeTab === 'plans') {
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.openCreateModal) }, { class: "inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none" }));
    /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-transparent']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-indigo-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-indigo-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
    (__VLS_ctx.t('learningPlan.createPlan'));
}
var __VLS_0;
/** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
routerLink;
// @ts-ignore
var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0(__assign({ to: "/admin/classes" }, { class: "inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50" })));
var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([__assign({ to: "/admin/classes" }, { class: "inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50" })], __VLS_functionalComponentArgsRest(__VLS_1), false));
/** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
var __VLS_5 = __VLS_3.slots.default;
(__VLS_ctx.t('learningPlan.backToClasses'));
// @ts-ignore
[t, t, t, t, t, classInfo, classInfo, activeTab, activeTab, openGenerateModal, openCreateModal,];
var __VLS_3;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4 border-b border-gray-200" }));
/** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
/** @type {__VLS_StyleScopedClasses['border-b']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.nav, __VLS_intrinsics.nav)(__assign({ class: "-mb-px flex space-x-6" }));
/** @type {__VLS_StyleScopedClasses['-mb-px']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['space-x-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
        var _a = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            _a[_i] = arguments[_i];
        }
        var $event = _a[0];
        __VLS_ctx.activeTab = 'plans';
        // @ts-ignore
        [activeTab,];
    } }, { class: (['pb-2 text-sm font-medium border-b-2 transition-colors', __VLS_ctx.activeTab === 'plans' ? 'border-indigo-500 text-indigo-600' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300']) }));
/** @type {__VLS_StyleScopedClasses['pb-2']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['border-b-2']} */ ;
/** @type {__VLS_StyleScopedClasses['transition-colors']} */ ;
(__VLS_ctx.t('learningPlan.planList'));
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
        var _a = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            _a[_i] = arguments[_i];
        }
        var $event = _a[0];
        __VLS_ctx.activeTab = 'students';
        __VLS_ctx.loadStudents();
        // @ts-ignore
        [t, activeTab, activeTab, loadStudents,];
    } }, { class: (['pb-2 text-sm font-medium border-b-2 transition-colors', __VLS_ctx.activeTab === 'students' ? 'border-indigo-500 text-indigo-600' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300']) }));
/** @type {__VLS_StyleScopedClasses['pb-2']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['border-b-2']} */ ;
/** @type {__VLS_StyleScopedClasses['transition-colors']} */ ;
(__VLS_ctx.t('learningPlan.studentList'));
if (__VLS_ctx.activeTab === 'plans') {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4 flex gap-4 items-center flex-wrap" }));
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex-wrap']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "text-sm font-medium text-gray-700 mr-2" }));
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['mr-2']} */ ;
    (__VLS_ctx.t('learningPlan.userId'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign(__assign({ onChange: (__VLS_ctx.loadPlans) }, { type: "number", placeholder: "User ID" }), { class: "border border-gray-300 rounded-md px-3 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 w-32" }));
    (__VLS_ctx.filterUserId);
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-32']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex gap-1" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['gap-1']} */ ;
    var _loop_1 = function (type) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(__VLS_ctx.activeTab === 'plans'))
                    return;
                __VLS_ctx.filterPlanType = type;
                __VLS_ctx.loadPlans();
                // @ts-ignore
                [t, t, activeTab, activeTab, loadPlans, loadPlans, filterUserId, filterPlanType,];
            } }, { key: (type) }), { class: (['px-3 py-1 text-sm rounded-md font-medium border', __VLS_ctx.filterPlanType === type ? 'bg-indigo-600 text-white border-indigo-600' : 'bg-white text-gray-700 border-gray-300 hover:bg-gray-50']) }));
        /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        (type === '' ? 'All' : __VLS_ctx.t("learningPlan.type".concat(type.charAt(0).toUpperCase() + type.slice(1))));
        // @ts-ignore
        [t, filterPlanType,];
    };
    for (var _i = 0, _e = __VLS_vFor((['', 'long', 'mid', 'short'])); _i < _e.length; _i++) {
        var type = _e[_i][0];
        _loop_1(type);
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex gap-1" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['gap-1']} */ ;
    var _loop_2 = function (src) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(__VLS_ctx.activeTab === 'plans'))
                    return;
                __VLS_ctx.filterIsPersonal = src;
                __VLS_ctx.loadPlans();
                // @ts-ignore
                [loadPlans, filterIsPersonal,];
            } }, { key: (String(src)) }), { class: (['px-3 py-1 text-sm rounded-md font-medium border', __VLS_ctx.filterIsPersonal === src ? 'bg-indigo-600 text-white border-indigo-600' : 'bg-white text-gray-700 border-gray-300 hover:bg-gray-50']) }));
        /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        (src === null ? 'All' : (src ? __VLS_ctx.t('learningPlan.planSourcePersonal') : __VLS_ctx.t('learningPlan.planSourceBatch')));
        // @ts-ignore
        [t, t, filterIsPersonal,];
    };
    for (var _f = 0, _g = __VLS_vFor(([null, true, false])); _f < _g.length; _f++) {
        var src = _g[_f][0];
        _loop_2(src);
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white shadow rounded-lg overflow-hidden" }));
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['overflow-hidden']} */ ;
    if (__VLS_ctx.loading) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-12 text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-12']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (__VLS_ctx.t('learningPlan.loading'));
    }
    else if (!__VLS_ctx.plans.length) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-12 text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-12']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (__VLS_ctx.t('learningPlan.noPlans'));
    }
    else {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "overflow-x-auto" }));
        /** @type {__VLS_StyleScopedClasses['overflow-x-auto']} */ ;
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
        (__VLS_ctx.t('learningPlan.userId'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
        /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
        (__VLS_ctx.t('learningPlan.planType'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
        /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
        (__VLS_ctx.t('learningPlan.planSource'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
        /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
        (__VLS_ctx.t('learningPlan.version'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
        /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-right']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
        /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.tbody, __VLS_intrinsics.tbody)(__assign({ class: "bg-white divide-y divide-gray-200" }));
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
        /** @type {__VLS_StyleScopedClasses['divide-gray-200']} */ ;
        var _loop_3 = function (plan) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({
                key: (plan.id),
            });
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
            /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            (plan.id);
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-900" }));
            /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
            (plan.userId);
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm" }));
            /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" }, { class: ({
                    'bg-blue-100 text-blue-800': plan.planType === 'long',
                    'bg-yellow-100 text-yellow-800': plan.planType === 'mid',
                    'bg-green-100 text-green-800': plan.planType === 'short'
                }) }));
            /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
            /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
            /** @type {__VLS_StyleScopedClasses['px-2.5']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            /** @type {__VLS_StyleScopedClasses['bg-blue-100']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-blue-800']} */ ;
            /** @type {__VLS_StyleScopedClasses['bg-yellow-100']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-yellow-800']} */ ;
            /** @type {__VLS_StyleScopedClasses['bg-green-100']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-green-800']} */ ;
            (__VLS_ctx.t("learningPlan.type".concat(plan.planType.charAt(0).toUpperCase() + plan.planType.slice(1))));
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm" }));
            /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" }, { class: (plan.isPersonal ? 'bg-purple-100 text-purple-800' : 'bg-gray-100 text-gray-700') }));
            /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
            /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
            /** @type {__VLS_StyleScopedClasses['px-2.5']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            (plan.isPersonal ? __VLS_ctx.t('learningPlan.planSourcePersonal') : __VLS_ctx.t('learningPlan.planSourceBatch'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
            /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            (plan.version);
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
            /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            ((_a = plan.updatedAt) === null || _a === void 0 ? void 0 : _a.slice(0, 10));
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2" }));
            /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-right']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            /** @type {__VLS_StyleScopedClasses['space-x-2']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!(__VLS_ctx.activeTab === 'plans'))
                        return;
                    if (!!(__VLS_ctx.loading))
                        return;
                    if (!!(!__VLS_ctx.plans.length))
                        return;
                    __VLS_ctx.openEditModal(plan);
                    // @ts-ignore
                    [t, t, t, t, t, t, t, t, t, loading, plans, plans, openEditModal,];
                } }, { class: "text-indigo-600 hover:text-indigo-900" }));
            /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
            /** @type {__VLS_StyleScopedClasses['hover:text-indigo-900']} */ ;
            (__VLS_ctx.t('learningPlan.editPlan'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!(__VLS_ctx.activeTab === 'plans'))
                        return;
                    if (!!(__VLS_ctx.loading))
                        return;
                    if (!!(!__VLS_ctx.plans.length))
                        return;
                    __VLS_ctx.openVersionsModal(plan);
                    // @ts-ignore
                    [t, openVersionsModal,];
                } }, { class: "text-gray-600 hover:text-gray-900" }));
            /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
            /** @type {__VLS_StyleScopedClasses['hover:text-gray-900']} */ ;
            (__VLS_ctx.t('learningPlan.viewVersions'));
            var __VLS_6 = void 0;
            /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
            routerLink;
            // @ts-ignore
            var __VLS_7 = __VLS_asFunctionalComponent1(__VLS_6, new __VLS_6(__assign({ to: ("/admin/learning-plans/".concat(plan.id, "/phase-plans")) }, { class: "text-purple-600 hover:text-purple-900" })));
            var __VLS_8 = __VLS_7.apply(void 0, __spreadArray([__assign({ to: ("/admin/learning-plans/".concat(plan.id, "/phase-plans")) }, { class: "text-purple-600 hover:text-purple-900" })], __VLS_functionalComponentArgsRest(__VLS_7), false));
            /** @type {__VLS_StyleScopedClasses['text-purple-600']} */ ;
            /** @type {__VLS_StyleScopedClasses['hover:text-purple-900']} */ ;
            var __VLS_11 = __VLS_9.slots.default;
            (__VLS_ctx.t('learningPlan.phaseManage'));
            // @ts-ignore
            [t, t,];
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!(__VLS_ctx.activeTab === 'plans'))
                        return;
                    if (!!(__VLS_ctx.loading))
                        return;
                    if (!!(!__VLS_ctx.plans.length))
                        return;
                    __VLS_ctx.confirmDelete(plan);
                    // @ts-ignore
                    [confirmDelete,];
                } }, { class: "text-red-600 hover:text-red-900" }));
            /** @type {__VLS_StyleScopedClasses['text-red-600']} */ ;
            /** @type {__VLS_StyleScopedClasses['hover:text-red-900']} */ ;
            (__VLS_ctx.t('common.delete'));
            // @ts-ignore
            [t,];
        };
        var __VLS_9;
        for (var _h = 0, _j = __VLS_vFor((__VLS_ctx.plans)); _h < _j.length; _h++) {
            var plan = _j[_h][0];
            _loop_3(plan);
        }
    }
}
if (__VLS_ctx.activeTab === 'students') {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white shadow rounded-lg overflow-hidden" }));
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['overflow-hidden']} */ ;
    if (__VLS_ctx.loadingStudents) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-12 text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-12']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (__VLS_ctx.t('learningPlan.loading'));
    }
    else if (!__VLS_ctx.students.length) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-12 text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-12']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (__VLS_ctx.t('learningPlan.noStudents'));
    }
    else {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "overflow-x-auto" }));
        /** @type {__VLS_StyleScopedClasses['overflow-x-auto']} */ ;
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
        (__VLS_ctx.t('learningPlan.studentUsername'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
        /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
        (__VLS_ctx.t('learningPlan.studentName'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-right']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
        /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.tbody, __VLS_intrinsics.tbody)(__assign({ class: "bg-white divide-y divide-gray-200" }));
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
        /** @type {__VLS_StyleScopedClasses['divide-gray-200']} */ ;
        var _loop_4 = function (student) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({
                key: (student.id),
            });
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
            /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            (student.id);
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-900" }));
            /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
            (student.username);
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
            /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            (student.nickName || student.name || '—');
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-right text-sm font-medium" }));
            /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-right']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!(__VLS_ctx.activeTab === 'students'))
                        return;
                    if (!!(__VLS_ctx.loadingStudents))
                        return;
                    if (!!(!__VLS_ctx.students.length))
                        return;
                    __VLS_ctx.openCreateModalForStudent(student);
                    // @ts-ignore
                    [t, t, t, t, activeTab, loadingStudents, students, students, openCreateModalForStudent,];
                } }, { class: "text-indigo-600 hover:text-indigo-900" }));
            /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
            /** @type {__VLS_StyleScopedClasses['hover:text-indigo-900']} */ ;
            (__VLS_ctx.t('learningPlan.createPersonalPlan'));
            // @ts-ignore
            [t,];
        };
        for (var _k = 0, _l = __VLS_vFor((__VLS_ctx.students)); _k < _l.length; _k++) {
            var student = _l[_k][0];
            _loop_4(student);
        }
    }
}
if (__VLS_ctx.showModal) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40" }));
    /** @type {__VLS_StyleScopedClasses['fixed']} */ ;
    /** @type {__VLS_StyleScopedClasses['inset-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['z-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-black']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-opacity-40']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow-xl w-full max-w-lg p-6" }));
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-xl']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['max-w-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['p-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-lg font-semibold text-gray-900 mb-4" }));
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    (__VLS_ctx.editingPlan ? __VLS_ctx.t('learningPlan.editPlan') : __VLS_ctx.t('learningPlan.createPlan'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-4" }));
    /** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
    if (!__VLS_ctx.editingPlan) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
        __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
        (__VLS_ctx.t('learningPlan.userId'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "number" }, { class: "w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" }));
        (__VLS_ctx.form.userId);
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
    }
    if (!__VLS_ctx.editingPlan) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
        __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
        (__VLS_ctx.t('learningPlan.planType'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ value: (__VLS_ctx.form.planType) }, { class: "w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" }));
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
            value: "long",
        });
        (__VLS_ctx.t('learningPlan.typeLong'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
            value: "mid",
        });
        (__VLS_ctx.t('learningPlan.typeMid'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
            value: "short",
        });
        (__VLS_ctx.t('learningPlan.typeShort'));
    }
    if (!__VLS_ctx.editingPlan) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-center gap-2" }));
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ id: "isPersonalCheck", type: "checkbox" }, { class: "h-4 w-4 text-indigo-600 border-gray-300 rounded" }));
        (__VLS_ctx.form.isPersonal);
        /** @type {__VLS_StyleScopedClasses['h-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "isPersonalCheck" }, { class: "text-sm font-medium text-gray-700" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        (__VLS_ctx.t('learningPlan.planSourcePersonal'));
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
    (__VLS_ctx.t('learningPlan.content'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.textarea)(__assign({ value: (__VLS_ctx.form.content), rows: "6" }, { class: "w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" }));
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
    (__VLS_ctx.t('learningPlan.comment'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ value: (__VLS_ctx.form.comment), type: "text" }, { class: "w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" }));
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-6 flex justify-end space-x-3" }));
    /** @type {__VLS_StyleScopedClasses['mt-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-end']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-x-3']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.closeModal) }, { class: "px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50" }));
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
    (__VLS_ctx.t('learningPlan.cancel'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: (__VLS_ctx.submitForm) }, { disabled: (__VLS_ctx.saving) }), { class: "px-4 py-2 text-sm font-medium text-white bg-indigo-600 rounded-md hover:bg-indigo-700 disabled:opacity-50" }));
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-indigo-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-indigo-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
    (__VLS_ctx.saving ? __VLS_ctx.t('learningPlan.loading') : (__VLS_ctx.editingPlan ? __VLS_ctx.t('learningPlan.save') : __VLS_ctx.t('learningPlan.create')));
}
if (__VLS_ctx.showVersionsModal) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40" }));
    /** @type {__VLS_StyleScopedClasses['fixed']} */ ;
    /** @type {__VLS_StyleScopedClasses['inset-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['z-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-black']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-opacity-40']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow-xl w-full max-w-2xl p-6" }));
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-xl']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['max-w-2xl']} */ ;
    /** @type {__VLS_StyleScopedClasses['p-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-lg font-semibold text-gray-900 mb-4" }));
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    (__VLS_ctx.t('learningPlan.viewVersions'));
    if (__VLS_ctx.loadingVersions) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-8 text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-8']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (__VLS_ctx.t('learningPlan.loading'));
    }
    else if (!__VLS_ctx.versions.length) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-8 text-gray-400" }));
        /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-8']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
    }
    else {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-3 max-h-96 overflow-y-auto" }));
        /** @type {__VLS_StyleScopedClasses['space-y-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['max-h-96']} */ ;
        /** @type {__VLS_StyleScopedClasses['overflow-y-auto']} */ ;
        var _loop_5 = function (v) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (v.id) }, { class: "border border-gray-200 rounded-lg p-4" }));
            /** @type {__VLS_StyleScopedClasses['border']} */ ;
            /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
            /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-between items-start" }));
            /** @type {__VLS_StyleScopedClasses['flex']} */ ;
            /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
            /** @type {__VLS_StyleScopedClasses['items-start']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
            __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium text-gray-900" }));
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
            (v.version);
            __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "ml-3 text-sm text-gray-500" }));
            /** @type {__VLS_StyleScopedClasses['ml-3']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            ((_b = v.createdAt) === null || _b === void 0 ? void 0 : _b.slice(0, 10));
            if (v.comment) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm text-gray-600 mt-1" }));
                /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
                /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
                (v.comment);
            }
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!(__VLS_ctx.showVersionsModal))
                        return;
                    if (!!(__VLS_ctx.loadingVersions))
                        return;
                    if (!!(!__VLS_ctx.versions.length))
                        return;
                    __VLS_ctx.doRollback(v);
                    // @ts-ignore
                    [t, t, t, t, t, t, t, t, t, t, t, t, t, t, t, t, showModal, editingPlan, editingPlan, editingPlan, editingPlan, editingPlan, form, form, form, form, form, closeModal, submitForm, saving, saving, showVersionsModal, loadingVersions, versions, versions, doRollback,];
                } }, { disabled: (__VLS_ctx.saving) }), { class: "text-sm text-indigo-600 hover:text-indigo-900 disabled:opacity-50" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
            /** @type {__VLS_StyleScopedClasses['hover:text-indigo-900']} */ ;
            /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
            (__VLS_ctx.t('learningPlan.rollback'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-2 text-xs text-gray-400 line-clamp-2" }));
            /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
            /** @type {__VLS_StyleScopedClasses['line-clamp-2']} */ ;
            (v.content);
            // @ts-ignore
            [t, saving,];
        };
        for (var _m = 0, _o = __VLS_vFor((__VLS_ctx.versions)); _m < _o.length; _m++) {
            var v = _o[_m][0];
            _loop_5(v);
        }
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-6 flex justify-end" }));
    /** @type {__VLS_StyleScopedClasses['mt-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-end']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.showVersionsModal))
                return;
            __VLS_ctx.showVersionsModal = false;
            // @ts-ignore
            [showVersionsModal,];
        } }, { class: "px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50" }));
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
    (__VLS_ctx.t('learningPlan.cancel'));
}
if (__VLS_ctx.showGenerateModal) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40" }));
    /** @type {__VLS_StyleScopedClasses['fixed']} */ ;
    /** @type {__VLS_StyleScopedClasses['inset-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['z-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-black']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-opacity-40']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow-xl w-full max-w-lg p-6 max-h-[90vh] overflow-y-auto" }));
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-xl']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['max-w-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['p-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['max-h-[90vh]']} */ ;
    /** @type {__VLS_StyleScopedClasses['overflow-y-auto']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-lg font-semibold text-gray-900 mb-4" }));
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    (__VLS_ctx.t('learningPlan.generateTemplate'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm text-gray-600 mb-4" }));
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    (__VLS_ctx.t('learningPlan.generateTemplateDesc'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-4" }));
    /** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-2" }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
    (__VLS_ctx.t('learningPlan.examNodeSchedules'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-xs text-gray-500 mb-2" }));
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
    (__VLS_ctx.t('learningPlan.examNodeSchedulesDesc'));
    if (__VLS_ctx.genForm.examNodes.length === 0) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-400 italic" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
        /** @type {__VLS_StyleScopedClasses['italic']} */ ;
        (__VLS_ctx.t('learningPlan.noExamNodesFound'));
    }
    for (var _p = 0, _q = __VLS_vFor((__VLS_ctx.genForm.examNodes)); _p < _q.length; _p++) {
        var _r = _q[_p], node = _r[0], idx = _r[1];
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (node.examNodeId) }, { class: "mb-3 p-3 border border-gray-200 rounded-md" }));
        /** @type {__VLS_StyleScopedClasses['mb-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['p-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm font-medium text-gray-700 mb-2" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
        (node.name);
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "grid grid-cols-2 gap-3" }));
        /** @type {__VLS_StyleScopedClasses['grid']} */ ;
        /** @type {__VLS_StyleScopedClasses['grid-cols-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['gap-3']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
        __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-xs text-gray-500 mb-1" }));
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
        (__VLS_ctx.t('learningPlan.startMonth'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "month" }, { class: "w-full border border-gray-300 rounded-md px-2 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" }));
        (__VLS_ctx.genForm.examNodes[idx].startMonth);
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:ring-green-500']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
        __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-xs text-gray-500 mb-1" }));
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
        (__VLS_ctx.t('learningPlan.endMonth'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "month" }, { class: "w-full border border-gray-300 rounded-md px-2 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" }));
        (__VLS_ctx.genForm.examNodes[idx].endMonth);
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:ring-green-500']} */ ;
        // @ts-ignore
        [t, t, t, t, t, t, t, t, showGenerateModal, genForm, genForm, genForm, genForm,];
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
    (__VLS_ctx.t('learningPlan.phaseRatios'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-xs text-gray-500 mb-1" }));
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
    (__VLS_ctx.t('learningPlan.phaseRatiosDesc'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "grid grid-cols-4 gap-2" }));
    /** @type {__VLS_StyleScopedClasses['grid']} */ ;
    /** @type {__VLS_StyleScopedClasses['grid-cols-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
    for (var _s = 0, _t = __VLS_vFor((__VLS_ctx.phaseLabels)); _s < _t.length; _s++) {
        var _u = _t[_s], label = _u[0], i = _u[1];
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({
            key: (i),
        });
        __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-xs text-gray-500 mb-1" }));
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
        (label);
        __VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "number", min: "0", max: "100" }, { class: "w-full border border-gray-300 rounded-md px-2 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" }));
        (__VLS_ctx.genForm.phaseRatios[i]);
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['focus:ring-green-500']} */ ;
        // @ts-ignore
        [t, t, genForm, phaseLabels,];
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-xs text-gray-400 mt-1" }));
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
    /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
    (__VLS_ctx.t('learningPlan.phaseRatiosSum'));
    (__VLS_ctx.genForm.phaseRatios.reduce(function (a, b) { return a + b; }, 0));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
    (__VLS_ctx.t('learningPlan.comment'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ value: (__VLS_ctx.genForm.comment), type: "text" }, { class: "w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" }));
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-green-500']} */ ;
    if (__VLS_ctx.generateResult) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-4 p-3 rounded-md" }, { class: (((_c = __VLS_ctx.generateResult.errors) === null || _c === void 0 ? void 0 : _c.length) ? 'bg-yellow-50 border border-yellow-200' : 'bg-green-50 border border-green-200') }));
        /** @type {__VLS_StyleScopedClasses['mt-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['p-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm font-medium" }));
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        (__VLS_ctx.t('learningPlan.generateSuccess'));
        (__VLS_ctx.generateResult.count);
        (__VLS_ctx.t('learningPlan.plans'));
        (__VLS_ctx.generateResult.studentCount);
        (__VLS_ctx.t('learningPlan.students'));
        if ((_d = __VLS_ctx.generateResult.errors) === null || _d === void 0 ? void 0 : _d.length) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.ul, __VLS_intrinsics.ul)(__assign({ class: "mt-2 text-xs text-yellow-700 space-y-1" }));
            /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-yellow-700']} */ ;
            /** @type {__VLS_StyleScopedClasses['space-y-1']} */ ;
            for (var _v = 0, _w = __VLS_vFor((__VLS_ctx.generateResult.errors)); _v < _w.length; _v++) {
                var e = _w[_v][0];
                __VLS_asFunctionalElement1(__VLS_intrinsics.li, __VLS_intrinsics.li)({
                    key: (e),
                });
                (e);
                // @ts-ignore
                [t, t, t, t, t, genForm, genForm, generateResult, generateResult, generateResult, generateResult, generateResult, generateResult,];
            }
        }
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-6 flex justify-end space-x-3" }));
    /** @type {__VLS_StyleScopedClasses['mt-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-end']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-x-3']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.closeGenerateModal) }, { class: "px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50" }));
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
    (__VLS_ctx.t('learningPlan.cancel'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: (__VLS_ctx.submitGenerate) }, { disabled: (__VLS_ctx.generating) }), { class: "px-4 py-2 text-sm font-medium text-white bg-green-600 rounded-md hover:bg-green-700 disabled:opacity-50" }));
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-green-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-green-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
    (__VLS_ctx.generating ? __VLS_ctx.t('learningPlan.generating') : __VLS_ctx.t('learningPlan.generate'));
}
if (__VLS_ctx.showDeleteModal) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40" }));
    /** @type {__VLS_StyleScopedClasses['fixed']} */ ;
    /** @type {__VLS_StyleScopedClasses['inset-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['z-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-black']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-opacity-40']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow-xl w-full max-w-sm p-6" }));
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-xl']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['max-w-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['p-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-lg font-semibold text-gray-900 mb-2" }));
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
    (__VLS_ctx.t('learningPlan.deletePlan'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm text-gray-600 mb-6" }));
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
    (__VLS_ctx.t('learningPlan.confirmDelete'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-end space-x-3" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-end']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-x-3']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.showDeleteModal))
                return;
            __VLS_ctx.showDeleteModal = false;
            // @ts-ignore
            [t, t, t, t, t, closeGenerateModal, submitGenerate, generating, generating, showDeleteModal, showDeleteModal,];
        } }, { class: "px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50" }));
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
    (__VLS_ctx.t('learningPlan.cancel'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: (__VLS_ctx.doDelete) }, { disabled: (__VLS_ctx.saving) }), { class: "px-4 py-2 text-sm font-medium text-white bg-red-600 rounded-md hover:bg-red-700 disabled:opacity-50" }));
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-red-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-red-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
    (__VLS_ctx.saving ? __VLS_ctx.t('learningPlan.loading') : __VLS_ctx.t('learningPlan.confirm'));
}
// @ts-ignore
[t, t, t, saving, saving, doDelete,];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
