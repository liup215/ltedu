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
var examNodeService_1 = require("../../services/examNodeService");
var syllabusService_1 = require("../../services/syllabusService");
var chapterService_1 = require("../../services/chapterService");
var t = (0, vue_i18n_1.useI18n)().t;
var route = (0, vue_router_1.useRoute)();
var syllabusId = Number(route.params.syllabusId);
var nodes = (0, vue_1.ref)([]);
var syllabus = (0, vue_1.ref)(null);
var chapters = (0, vue_1.ref)([]);
var loading = (0, vue_1.ref)(false);
var saving = (0, vue_1.ref)(false);
var showModal = (0, vue_1.ref)(false);
var showAddChapterModal = (0, vue_1.ref)(false);
var showDeleteModal = (0, vue_1.ref)(false);
var editingNode = (0, vue_1.ref)(null);
var addingChapterNode = (0, vue_1.ref)(null);
var deletingNode = (0, vue_1.ref)(null);
var expandedNodeId = (0, vue_1.ref)(null);
var selectedChapterId = (0, vue_1.ref)(null);
var form = (0, vue_1.ref)({ name: '', description: '', sortOrder: 0 });
function loadNodes() {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_1;
        var _a;
        return __generator(this, function (_b) {
            switch (_b.label) {
                case 0:
                    loading.value = true;
                    _b.label = 1;
                case 1:
                    _b.trys.push([1, 3, 4, 5]);
                    return [4 /*yield*/, examNodeService_1.default.list(syllabusId)];
                case 2:
                    res = _b.sent();
                    if (res.code === 0)
                        nodes.value = (_a = res.data.list) !== null && _a !== void 0 ? _a : [];
                    return [3 /*break*/, 5];
                case 3:
                    e_1 = _b.sent();
                    console.error(e_1);
                    return [3 /*break*/, 5];
                case 4:
                    loading.value = false;
                    return [7 /*endfinally*/];
                case 5: return [2 /*return*/];
            }
        });
    });
}
function loadSyllabus() {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_2;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    _a.trys.push([0, 2, , 3]);
                    return [4 /*yield*/, syllabusService_1.default.getSyllabusById(syllabusId)];
                case 1:
                    res = _a.sent();
                    if (res.code === 0)
                        syllabus.value = res.data;
                    return [3 /*break*/, 3];
                case 2:
                    e_2 = _a.sent();
                    console.error(e_2);
                    return [3 /*break*/, 3];
                case 3: return [2 /*return*/];
            }
        });
    });
}
function loadChapters() {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_3;
        var _a;
        return __generator(this, function (_b) {
            switch (_b.label) {
                case 0:
                    _b.trys.push([0, 2, , 3]);
                    return [4 /*yield*/, chapterService_1.default.getChapterList({ syllabusId: syllabusId, pageSize: 200, pageIndex: 1 })];
                case 1:
                    res = _b.sent();
                    if (res.code === 0)
                        chapters.value = (_a = res.data.list) !== null && _a !== void 0 ? _a : res.data;
                    return [3 /*break*/, 3];
                case 2:
                    e_3 = _b.sent();
                    console.error(e_3);
                    return [3 /*break*/, 3];
                case 3: return [2 /*return*/];
            }
        });
    });
}
function toggleExpand(id) {
    expandedNodeId.value = expandedNodeId.value === id ? null : id;
}
function openCreateModal() {
    editingNode.value = null;
    form.value = { name: '', description: '', sortOrder: 0 };
    showModal.value = true;
}
function openEditModal(node) {
    editingNode.value = node;
    form.value = { name: node.name, description: node.description || '', sortOrder: node.sortOrder };
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
                    if (!editingNode.value) return [3 /*break*/, 5];
                    return [4 /*yield*/, examNodeService_1.default.update({ id: editingNode.value.id, name: form.value.name, description: form.value.description, sortOrder: form.value.sortOrder })];
                case 2:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 4];
                    return [4 /*yield*/, loadNodes()];
                case 3:
                    _a.sent();
                    closeModal();
                    _a.label = 4;
                case 4: return [3 /*break*/, 8];
                case 5: return [4 /*yield*/, examNodeService_1.default.create({ syllabusId: syllabusId, name: form.value.name, description: form.value.description, sortOrder: form.value.sortOrder })];
                case 6:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 8];
                    return [4 /*yield*/, loadNodes()];
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
function openAddChapterModal(node) {
    addingChapterNode.value = node;
    selectedChapterId.value = null;
    showAddChapterModal.value = true;
}
function doAddChapter() {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_5;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    if (!addingChapterNode.value || !selectedChapterId.value)
                        return [2 /*return*/];
                    saving.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 5, 6, 7]);
                    return [4 /*yield*/, examNodeService_1.default.addChapter(addingChapterNode.value.id, selectedChapterId.value)];
                case 2:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 4];
                    return [4 /*yield*/, loadNodes()];
                case 3:
                    _a.sent();
                    showAddChapterModal.value = false;
                    _a.label = 4;
                case 4: return [3 /*break*/, 7];
                case 5:
                    e_5 = _a.sent();
                    console.error(e_5);
                    return [3 /*break*/, 7];
                case 6:
                    saving.value = false;
                    return [7 /*endfinally*/];
                case 7: return [2 /*return*/];
            }
        });
    });
}
function doRemoveChapter(node, chapterId) {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_6;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    _a.trys.push([0, 4, , 5]);
                    return [4 /*yield*/, examNodeService_1.default.removeChapter(node.id, chapterId)];
                case 1:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 3];
                    return [4 /*yield*/, loadNodes()];
                case 2:
                    _a.sent();
                    _a.label = 3;
                case 3: return [3 /*break*/, 5];
                case 4:
                    e_6 = _a.sent();
                    console.error(e_6);
                    return [3 /*break*/, 5];
                case 5: return [2 /*return*/];
            }
        });
    });
}
function confirmDelete(node) {
    deletingNode.value = node;
    showDeleteModal.value = true;
}
function doDelete() {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_7;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    if (!deletingNode.value)
                        return [2 /*return*/];
                    saving.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 5, 6, 7]);
                    return [4 /*yield*/, examNodeService_1.default.delete(deletingNode.value.id)];
                case 2:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 4];
                    return [4 /*yield*/, loadNodes()];
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
(0, vue_1.onMounted)(function () {
    loadNodes();
    loadSyllabus();
    loadChapters();
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
(__VLS_ctx.t('examNode.title'));
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-2 text-sm text-gray-600" }));
/** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
(__VLS_ctx.t('examNode.subtitle'));
if (__VLS_ctx.syllabus) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-3 text-sm text-gray-700 font-medium" }));
    /** @type {__VLS_StyleScopedClasses['mt-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    (__VLS_ctx.syllabus.name);
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex gap-2" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.openCreateModal) }, { class: "inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500" }));
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
/** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-offset-2']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
(__VLS_ctx.t('examNode.createNode'));
var __VLS_0;
/** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
routerLink;
// @ts-ignore
var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0(__assign({ to: "/admin/syllabuses" }, { class: "inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50" })));
var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([__assign({ to: "/admin/syllabuses" }, { class: "inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50" })], __VLS_functionalComponentArgsRest(__VLS_1), false));
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
(__VLS_ctx.t('examNode.backToSyllabuses'));
// @ts-ignore
[t, t, t, t, syllabus, syllabus, openCreateModal,];
var __VLS_3;
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
    (__VLS_ctx.t('examNode.loading'));
}
else if (!__VLS_ctx.nodes.length) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-12 text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-12']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    (__VLS_ctx.t('examNode.noNodes'));
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
    (__VLS_ctx.t('examNode.name'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
    /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
    (__VLS_ctx.t('examNode.description'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
    /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
    (__VLS_ctx.t('examNode.sortOrder'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
    /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
    (__VLS_ctx.t('examNode.chapterCount'));
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
    var _loop_1 = function (node) {
        (node.id);
        __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
        __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
        (node.name);
        __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 text-sm text-gray-500 max-w-xs truncate" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['max-w-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['truncate']} */ ;
        (node.description || '-');
        __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (node.sortOrder);
        __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        ((_b = (_a = node.chapters) === null || _a === void 0 ? void 0 : _a.length) !== null && _b !== void 0 ? _b : 0);
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
                if (!!(__VLS_ctx.loading))
                    return;
                if (!!(!__VLS_ctx.nodes.length))
                    return;
                __VLS_ctx.toggleExpand(node.id);
                // @ts-ignore
                [t, t, t, t, t, t, loading, nodes, nodes, toggleExpand,];
            } }, { class: "text-gray-600 hover:text-gray-900" }));
        /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-gray-900']} */ ;
        (__VLS_ctx.expandedNodeId === node.id ? '▲' : '▼');
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!!(__VLS_ctx.loading))
                    return;
                if (!!(!__VLS_ctx.nodes.length))
                    return;
                __VLS_ctx.openAddChapterModal(node);
                // @ts-ignore
                [expandedNodeId, openAddChapterModal,];
            } }, { class: "text-green-600 hover:text-green-900" }));
        /** @type {__VLS_StyleScopedClasses['text-green-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-green-900']} */ ;
        (__VLS_ctx.t('examNode.addChapter'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!!(__VLS_ctx.loading))
                    return;
                if (!!(!__VLS_ctx.nodes.length))
                    return;
                __VLS_ctx.openEditModal(node);
                // @ts-ignore
                [t, openEditModal,];
            } }, { class: "text-indigo-600 hover:text-indigo-900" }));
        /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-indigo-900']} */ ;
        (__VLS_ctx.t('common.edit'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!!(__VLS_ctx.loading))
                    return;
                if (!!(!__VLS_ctx.nodes.length))
                    return;
                __VLS_ctx.confirmDelete(node);
                // @ts-ignore
                [t, confirmDelete,];
            } }, { class: "text-red-600 hover:text-red-900" }));
        /** @type {__VLS_StyleScopedClasses['text-red-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-red-900']} */ ;
        (__VLS_ctx.t('common.delete'));
        if (__VLS_ctx.expandedNodeId === node.id) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ colspan: "5" }, { class: "px-6 py-4 bg-gray-50" }));
            /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm font-medium text-gray-700 mb-2" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
            /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
            (__VLS_ctx.t('examNode.chapters'));
            if (!((_c = node.chapters) === null || _c === void 0 ? void 0 : _c.length)) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-400 italic mb-3" }));
                /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
                /** @type {__VLS_StyleScopedClasses['italic']} */ ;
                /** @type {__VLS_StyleScopedClasses['mb-3']} */ ;
            }
            else {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex flex-wrap gap-2 mb-3" }));
                /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                /** @type {__VLS_StyleScopedClasses['flex-wrap']} */ ;
                /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
                /** @type {__VLS_StyleScopedClasses['mb-3']} */ ;
                var _loop_3 = function (ch) {
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ key: (ch.id) }, { class: "inline-flex items-center gap-1 px-2 py-1 bg-white border border-gray-200 rounded text-xs text-gray-700" }));
                    /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
                    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                    /** @type {__VLS_StyleScopedClasses['gap-1']} */ ;
                    /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
                    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
                    /** @type {__VLS_StyleScopedClasses['border']} */ ;
                    /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
                    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
                    (ch.name);
                    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                            var _a;
                            var _b = [];
                            for (var _i = 0; _i < arguments.length; _i++) {
                                _b[_i] = arguments[_i];
                            }
                            var $event = _b[0];
                            if (!!(__VLS_ctx.loading))
                                return;
                            if (!!(!__VLS_ctx.nodes.length))
                                return;
                            if (!(__VLS_ctx.expandedNodeId === node.id))
                                return;
                            if (!!(!((_a = node.chapters) === null || _a === void 0 ? void 0 : _a.length)))
                                return;
                            __VLS_ctx.doRemoveChapter(node, ch.id);
                            // @ts-ignore
                            [t, t, expandedNodeId, doRemoveChapter,];
                        } }, { class: "text-red-400 hover:text-red-600 ml-1" }));
                    /** @type {__VLS_StyleScopedClasses['text-red-400']} */ ;
                    /** @type {__VLS_StyleScopedClasses['hover:text-red-600']} */ ;
                    /** @type {__VLS_StyleScopedClasses['ml-1']} */ ;
                    // @ts-ignore
                    [];
                };
                for (var _h = 0, _j = __VLS_vFor((node.chapters)); _h < _j.length; _h++) {
                    var ch = _j[_h][0];
                    _loop_3(ch);
                }
            }
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm font-medium text-gray-700 mb-2" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
            /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
            (__VLS_ctx.t('examNode.paperCodes'));
            if (!((_d = node.paperCodes) === null || _d === void 0 ? void 0 : _d.length)) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-400 italic" }));
                /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
                /** @type {__VLS_StyleScopedClasses['italic']} */ ;
            }
            else {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex flex-wrap gap-2" }));
                /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                /** @type {__VLS_StyleScopedClasses['flex-wrap']} */ ;
                /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
                for (var _k = 0, _l = __VLS_vFor((node.paperCodes)); _k < _l.length; _k++) {
                    var pc = _l[_k][0];
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ key: (pc.id) }, { class: "inline-flex items-center px-2 py-1 bg-blue-50 border border-blue-200 rounded text-xs text-blue-700" }));
                    /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
                    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                    /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
                    /** @type {__VLS_StyleScopedClasses['bg-blue-50']} */ ;
                    /** @type {__VLS_StyleScopedClasses['border']} */ ;
                    /** @type {__VLS_StyleScopedClasses['border-blue-200']} */ ;
                    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-blue-700']} */ ;
                    (pc.code || pc.name);
                    // @ts-ignore
                    [t,];
                }
            }
        }
        // @ts-ignore
        [];
    };
    for (var _i = 0, _e = __VLS_vFor((__VLS_ctx.nodes)); _i < _e.length; _i++) {
        var node = _e[_i][0];
        _loop_1(node);
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
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow-xl w-full max-w-md p-6" }));
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-xl']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['max-w-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['p-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-lg font-semibold text-gray-900 mb-4" }));
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    (__VLS_ctx.editingNode ? __VLS_ctx.t('examNode.editNode') : __VLS_ctx.t('examNode.createNode'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-4" }));
    /** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
    (__VLS_ctx.t('examNode.name'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ value: (__VLS_ctx.form.name), type: "text" }, { class: "w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" }));
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
    (__VLS_ctx.t('examNode.description'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.textarea)(__assign({ value: (__VLS_ctx.form.description), rows: "3" }, { class: "w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" }));
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
    (__VLS_ctx.t('examNode.sortOrder'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "number" }, { class: "w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" }));
    (__VLS_ctx.form.sortOrder);
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
    (__VLS_ctx.t('examNode.cancel'));
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
    (__VLS_ctx.saving ? __VLS_ctx.t('examNode.loading') : (__VLS_ctx.editingNode ? __VLS_ctx.t('examNode.save') : __VLS_ctx.t('examNode.create')));
}
if (__VLS_ctx.showAddChapterModal) {
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
    (__VLS_ctx.t('examNode.addChapter'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm text-gray-500 mb-3" }));
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-3']} */ ;
    (__VLS_ctx.t('examNode.selectChapter'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "max-h-80 overflow-y-auto border border-gray-200 rounded-md" }));
    /** @type {__VLS_StyleScopedClasses['max-h-80']} */ ;
    /** @type {__VLS_StyleScopedClasses['overflow-y-auto']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    if (!__VLS_ctx.chapters.length) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-4 text-sm text-gray-400 text-center" }));
        /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
        (__VLS_ctx.t('examNode.loading'));
    }
    var _loop_2 = function (ch) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(__VLS_ctx.showAddChapterModal))
                    return;
                __VLS_ctx.selectedChapterId = ch.id;
                // @ts-ignore
                [t, t, t, t, t, t, t, t, t, t, t, t, showModal, editingNode, editingNode, form, form, form, closeModal, submitForm, saving, saving, showAddChapterModal, chapters, chapters, selectedChapterId,];
            } }, { key: (ch.id) }), { class: (['px-4 py-2 cursor-pointer text-sm hover:bg-indigo-50', __VLS_ctx.selectedChapterId === ch.id ? 'bg-indigo-100 font-medium text-indigo-900' : 'text-gray-700']) }));
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['cursor-pointer']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-indigo-50']} */ ;
        (ch.name);
        // @ts-ignore
        [selectedChapterId,];
    };
    for (var _f = 0, _g = __VLS_vFor((__VLS_ctx.chapters)); _f < _g.length; _f++) {
        var ch = _g[_f][0];
        _loop_2(ch);
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-6 flex justify-end space-x-3" }));
    /** @type {__VLS_StyleScopedClasses['mt-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-end']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-x-3']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.showAddChapterModal))
                return;
            __VLS_ctx.showAddChapterModal = false;
            // @ts-ignore
            [showAddChapterModal,];
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
    (__VLS_ctx.t('examNode.cancel'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: (__VLS_ctx.doAddChapter) }, { disabled: (!__VLS_ctx.selectedChapterId || __VLS_ctx.saving) }), { class: "px-4 py-2 text-sm font-medium text-white bg-green-600 rounded-md hover:bg-green-700 disabled:opacity-50" }));
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-green-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-green-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
    (__VLS_ctx.saving ? __VLS_ctx.t('examNode.loading') : __VLS_ctx.t('examNode.addChapter'));
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
    (__VLS_ctx.t('examNode.deleteNode'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm text-gray-600 mb-6" }));
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
    (__VLS_ctx.t('examNode.confirmDelete'));
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
            [t, t, t, t, t, saving, saving, selectedChapterId, doAddChapter, showDeleteModal, showDeleteModal,];
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
    (__VLS_ctx.t('examNode.cancel'));
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
    (__VLS_ctx.saving ? __VLS_ctx.t('examNode.loading') : __VLS_ctx.t('examNode.confirm'));
}
// @ts-ignore
[t, t, t, saving, saving, doDelete,];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
