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
var vue_i18n_1 = require("vue-i18n");
var classService_1 = require("../../services/classService");
var syllabusService_1 = require("../../services/syllabusService");
var class_model_1 = require("../../models/class.model");
var t = (0, vue_i18n_1.useI18n)().t;
var classes = (0, vue_1.ref)([]);
var syllabuses = (0, vue_1.ref)([]);
var loading = (0, vue_1.ref)(false);
var saving = (0, vue_1.ref)(false);
var showModal = (0, vue_1.ref)(false);
var showBindModal = (0, vue_1.ref)(false);
var showDeleteModal = (0, vue_1.ref)(false);
var editingClass = (0, vue_1.ref)(null);
var bindingClass = (0, vue_1.ref)(null);
var deletingClass = (0, vue_1.ref)(null);
var selectedSyllabusId = (0, vue_1.ref)('');
var form = (0, vue_1.ref)({ name: '', classType: 1 });
function loadClasses() {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_1;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    loading.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 3, 4, 5]);
                    return [4 /*yield*/, classService_1.default.list({ pageSize: 100, pageIndex: 1 })];
                case 2:
                    res = _a.sent();
                    if (res.code === 0)
                        classes.value = res.data.list;
                    return [3 /*break*/, 5];
                case 3:
                    e_1 = _a.sent();
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
function loadSyllabuses() {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_2;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    _a.trys.push([0, 2, , 3]);
                    return [4 /*yield*/, syllabusService_1.default.getSyllabuses({ pageSize: 100, pageIndex: 1 })];
                case 1:
                    res = _a.sent();
                    if (res.code === 0)
                        syllabuses.value = res.data.list;
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
function openCreateModal() {
    editingClass.value = null;
    form.value = { name: '', classType: 1 };
    showModal.value = true;
}
function openEditModal(cls) {
    editingClass.value = cls;
    form.value = { name: cls.name, classType: cls.classType };
    showModal.value = true;
}
function closeModal() {
    showModal.value = false;
}
function submitForm() {
    return __awaiter(this, void 0, void 0, function () {
        var res, res, e_3;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    saving.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 9, 10, 11]);
                    if (!editingClass.value) return [3 /*break*/, 5];
                    return [4 /*yield*/, classService_1.default.update({ id: editingClass.value.id, name: form.value.name })];
                case 2:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 4];
                    return [4 /*yield*/, loadClasses()];
                case 3:
                    _a.sent();
                    closeModal();
                    _a.label = 4;
                case 4: return [3 /*break*/, 8];
                case 5: return [4 /*yield*/, classService_1.default.create({ name: form.value.name, classType: form.value.classType })];
                case 6:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 8];
                    return [4 /*yield*/, loadClasses()];
                case 7:
                    _a.sent();
                    closeModal();
                    _a.label = 8;
                case 8: return [3 /*break*/, 11];
                case 9:
                    e_3 = _a.sent();
                    console.error(e_3);
                    return [3 /*break*/, 11];
                case 10:
                    saving.value = false;
                    return [7 /*endfinally*/];
                case 11: return [2 /*return*/];
            }
        });
    });
}
function openBindSyllabusModal(cls) {
    bindingClass.value = cls;
    selectedSyllabusId.value = '';
    showBindModal.value = true;
}
function doBindSyllabus() {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_4;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    if (!bindingClass.value || !selectedSyllabusId.value)
                        return [2 /*return*/];
                    saving.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 5, 6, 7]);
                    return [4 /*yield*/, classService_1.default.bindSyllabus({ classId: bindingClass.value.id, syllabusId: Number(selectedSyllabusId.value) })];
                case 2:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 4];
                    return [4 /*yield*/, loadClasses()];
                case 3:
                    _a.sent();
                    showBindModal.value = false;
                    _a.label = 4;
                case 4: return [3 /*break*/, 7];
                case 5:
                    e_4 = _a.sent();
                    console.error(e_4);
                    return [3 /*break*/, 7];
                case 6:
                    saving.value = false;
                    return [7 /*endfinally*/];
                case 7: return [2 /*return*/];
            }
        });
    });
}
function doUnbindSyllabus(cls) {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_5;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    saving.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 5, 6, 7]);
                    return [4 /*yield*/, classService_1.default.unbindSyllabus({ classId: cls.id })];
                case 2:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 4];
                    return [4 /*yield*/, loadClasses()];
                case 3:
                    _a.sent();
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
function confirmDelete(cls) {
    deletingClass.value = cls;
    showDeleteModal.value = true;
}
function doDelete() {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_6;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    if (!deletingClass.value)
                        return [2 /*return*/];
                    saving.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 5, 6, 7]);
                    return [4 /*yield*/, classService_1.default.delete(deletingClass.value.id)];
                case 2:
                    res = _a.sent();
                    if (!(res.code === 0)) return [3 /*break*/, 4];
                    return [4 /*yield*/, loadClasses()];
                case 3:
                    _a.sent();
                    showDeleteModal.value = false;
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
// --- Students management ---
var showStudentsModal = (0, vue_1.ref)(false);
var studentsClass = (0, vue_1.ref)(null);
var students = (0, vue_1.ref)([]);
var studentsLoading = (0, vue_1.ref)(false);
function studentStatusLabel(status) {
    switch (status) {
        case class_model_1.CLASS_STUDENT_STATUS_STUDYING: return t('class.statusStudying');
        case class_model_1.CLASS_STUDENT_STATUS_GRADUATED: return t('class.statusGraduated');
        case class_model_1.CLASS_STUDENT_STATUS_TRANSFERRED: return t('class.statusTransferred');
        case class_model_1.CLASS_STUDENT_STATUS_DROPPED: return t('class.statusDropped');
        default: return t('class.statusStudying');
    }
}
function studentStatusClass(status) {
    switch (status) {
        case class_model_1.CLASS_STUDENT_STATUS_STUDYING: return 'bg-green-100 text-green-800';
        case class_model_1.CLASS_STUDENT_STATUS_GRADUATED: return 'bg-blue-100 text-blue-800';
        case class_model_1.CLASS_STUDENT_STATUS_TRANSFERRED: return 'bg-yellow-100 text-yellow-800';
        case class_model_1.CLASS_STUDENT_STATUS_DROPPED: return 'bg-red-100 text-red-800';
        default: return 'bg-gray-100 text-gray-800';
    }
}
function openStudentsModal(cls) {
    return __awaiter(this, void 0, void 0, function () {
        var res, e_7;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    studentsClass.value = cls;
                    showStudentsModal.value = true;
                    studentsLoading.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 3, 4, 5]);
                    return [4 /*yield*/, classService_1.default.getStudents(cls.id)];
                case 2:
                    res = _a.sent();
                    if (res.code === 0)
                        students.value = res.data.list;
                    return [3 /*break*/, 5];
                case 3:
                    e_7 = _a.sent();
                    console.error(e_7);
                    return [3 /*break*/, 5];
                case 4:
                    studentsLoading.value = false;
                    return [7 /*endfinally*/];
                case 5: return [2 /*return*/];
            }
        });
    });
}
function onStatusChange(stu, value) {
    return __awaiter(this, void 0, void 0, function () {
        var newStatus, res, e_8;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    if (!studentsClass.value)
                        return [2 /*return*/];
                    newStatus = Number(value);
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 3, , 4]);
                    return [4 /*yield*/, classService_1.default.updateStudentStatus(studentsClass.value.id, stu.id, newStatus)];
                case 2:
                    res = _a.sent();
                    if (res.code === 0) {
                        stu.studentStatus = newStatus;
                    }
                    else {
                        alert(t('class.updateStatusFailed'));
                    }
                    return [3 /*break*/, 4];
                case 3:
                    e_8 = _a.sent();
                    console.error(e_8);
                    alert(t('class.updateStatusFailed'));
                    return [3 /*break*/, 4];
                case 4: return [2 /*return*/];
            }
        });
    });
}
(0, vue_1.onMounted)(function () {
    loadClasses();
    loadSyllabuses();
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
(__VLS_ctx.t('class.title'));
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-2 text-sm text-gray-600" }));
/** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
(__VLS_ctx.t('class.subtitle'));
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
(__VLS_ctx.t('class.createClass'));
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
    (__VLS_ctx.t('class.loading'));
}
else if (!__VLS_ctx.classes.length) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-12 text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-12']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    (__VLS_ctx.t('class.noClasses'));
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
    (__VLS_ctx.t('class.name'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
    /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
    (__VLS_ctx.t('class.type'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
    /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
    (__VLS_ctx.t('class.syllabus'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
    /** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
    (__VLS_ctx.t('class.inviteCode'));
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
    var _loop_1 = function (cls) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({
            key: (cls.id),
        });
        __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
        (cls.name);
        __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" }, { class: (cls.classType === 1 ? 'bg-blue-100 text-blue-800' : 'bg-purple-100 text-purple-800') }));
        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        (cls.classType === 1 ? __VLS_ctx.t('class.typeTeaching') : __VLS_ctx.t('class.typeAdmin'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        if (cls.syllabus) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
            (cls.syllabus.name);
        }
        else {
            __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-gray-400 italic" }));
            /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
            /** @type {__VLS_StyleScopedClasses['italic']} */ ;
            (__VLS_ctx.t('class.syllabusNotBound'));
        }
        __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500 font-mono" }));
        /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-mono']} */ ;
        (cls.inviteCode);
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
                if (!!(!__VLS_ctx.classes.length))
                    return;
                __VLS_ctx.openEditModal(cls);
                // @ts-ignore
                [t, t, t, t, t, t, t, t, t, t, t, t, openCreateModal, loading, classes, classes, openEditModal,];
            } }, { class: "text-indigo-600 hover:text-indigo-900" }));
        /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-indigo-900']} */ ;
        (__VLS_ctx.t('common.edit'));
        if (!cls.syllabus) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!!(__VLS_ctx.loading))
                        return;
                    if (!!(!__VLS_ctx.classes.length))
                        return;
                    if (!(!cls.syllabus))
                        return;
                    __VLS_ctx.openBindSyllabusModal(cls);
                    // @ts-ignore
                    [t, openBindSyllabusModal,];
                } }, { class: "text-green-600 hover:text-green-900" }));
            /** @type {__VLS_StyleScopedClasses['text-green-600']} */ ;
            /** @type {__VLS_StyleScopedClasses['hover:text-green-900']} */ ;
            (__VLS_ctx.t('class.bindSyllabus'));
        }
        if (cls.syllabus) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!!(__VLS_ctx.loading))
                        return;
                    if (!!(!__VLS_ctx.classes.length))
                        return;
                    if (!(cls.syllabus))
                        return;
                    __VLS_ctx.doUnbindSyllabus(cls);
                    // @ts-ignore
                    [t, doUnbindSyllabus,];
                } }, { class: "text-yellow-600 hover:text-yellow-900" }));
            /** @type {__VLS_StyleScopedClasses['text-yellow-600']} */ ;
            /** @type {__VLS_StyleScopedClasses['hover:text-yellow-900']} */ ;
            (__VLS_ctx.t('class.unbindSyllabus'));
        }
        var __VLS_0 = void 0;
        /** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
        routerLink;
        // @ts-ignore
        var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0(__assign({ to: ("/admin/classes/".concat(cls.id, "/learning-plans")) }, { class: "text-purple-600 hover:text-purple-900" })));
        var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([__assign({ to: ("/admin/classes/".concat(cls.id, "/learning-plans")) }, { class: "text-purple-600 hover:text-purple-900" })], __VLS_functionalComponentArgsRest(__VLS_1), false));
        /** @type {__VLS_StyleScopedClasses['text-purple-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-purple-900']} */ ;
        var __VLS_5 = __VLS_3.slots.default;
        (__VLS_ctx.t('learningPlan.title'));
        // @ts-ignore
        [t, t,];
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!!(__VLS_ctx.loading))
                    return;
                if (!!(!__VLS_ctx.classes.length))
                    return;
                __VLS_ctx.openStudentsModal(cls);
                // @ts-ignore
                [openStudentsModal,];
            } }, { class: "text-teal-600 hover:text-teal-900" }));
        /** @type {__VLS_StyleScopedClasses['text-teal-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-teal-900']} */ ;
        (__VLS_ctx.t('class.students'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!!(__VLS_ctx.loading))
                    return;
                if (!!(!__VLS_ctx.classes.length))
                    return;
                __VLS_ctx.confirmDelete(cls);
                // @ts-ignore
                [t, confirmDelete,];
            } }, { class: "text-red-600 hover:text-red-900" }));
        /** @type {__VLS_StyleScopedClasses['text-red-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-red-900']} */ ;
        (__VLS_ctx.t('common.delete'));
        // @ts-ignore
        [t,];
    };
    var __VLS_3;
    for (var _i = 0, _b = __VLS_vFor((__VLS_ctx.classes)); _i < _b.length; _i++) {
        var cls = _b[_i][0];
        _loop_1(cls);
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
    (__VLS_ctx.editingClass ? __VLS_ctx.t('class.editClass') : __VLS_ctx.t('class.createClass'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-4" }));
    /** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
    (__VLS_ctx.t('class.name'));
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
    if (!__VLS_ctx.editingClass) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
        __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
        (__VLS_ctx.t('class.type'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ value: (__VLS_ctx.form.classType) }, { class: "w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" }));
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
            value: (1),
        });
        (__VLS_ctx.t('class.typeTeaching'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
            value: (2),
        });
        (__VLS_ctx.t('class.typeAdmin'));
    }
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
    (__VLS_ctx.t('class.cancel'));
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
    (__VLS_ctx.saving ? __VLS_ctx.t('class.loading') : (__VLS_ctx.editingClass ? __VLS_ctx.t('class.save') : __VLS_ctx.t('class.create')));
}
if (__VLS_ctx.showBindModal) {
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
    (__VLS_ctx.t('class.bindSyllabus'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-sm font-medium text-gray-700 mb-1" }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
    (__VLS_ctx.t('class.selectSyllabus'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ value: (__VLS_ctx.selectedSyllabusId) }, { class: "w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" }));
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
        value: "",
    });
    (__VLS_ctx.t('class.selectSyllabus'));
    for (var _c = 0, _d = __VLS_vFor((__VLS_ctx.syllabuses)); _c < _d.length; _c++) {
        var s = _d[_c][0];
        __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
            key: (s.id),
            value: (s.id),
        });
        (s.name);
        // @ts-ignore
        [t, t, t, t, t, t, t, t, t, t, t, t, t, showModal, editingClass, editingClass, editingClass, form, form, closeModal, submitForm, saving, saving, showBindModal, selectedSyllabusId, syllabuses,];
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
            if (!(__VLS_ctx.showBindModal))
                return;
            __VLS_ctx.showBindModal = false;
            // @ts-ignore
            [showBindModal,];
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
    (__VLS_ctx.t('class.cancel'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: (__VLS_ctx.doBindSyllabus) }, { disabled: (!__VLS_ctx.selectedSyllabusId || __VLS_ctx.saving) }), { class: "px-4 py-2 text-sm font-medium text-white bg-green-600 rounded-md hover:bg-green-700 disabled:opacity-50" }));
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-green-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-green-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
    (__VLS_ctx.saving ? __VLS_ctx.t('class.loading') : __VLS_ctx.t('class.confirm'));
}
if (__VLS_ctx.showStudentsModal) {
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
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-between items-center mb-4" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-lg font-semibold text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (__VLS_ctx.t('class.students'));
    ((_a = __VLS_ctx.studentsClass) === null || _a === void 0 ? void 0 : _a.name);
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.showStudentsModal))
                return;
            __VLS_ctx.showStudentsModal = false;
            // @ts-ignore
            [t, t, t, t, saving, saving, selectedSyllabusId, doBindSyllabus, showStudentsModal, showStudentsModal, studentsClass,];
        } }, { class: "text-gray-400 hover:text-gray-600 text-xl leading-none" }));
    /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-gray-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xl']} */ ;
    /** @type {__VLS_StyleScopedClasses['leading-none']} */ ;
    if (__VLS_ctx.studentsLoading) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-8 text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-8']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (__VLS_ctx.t('class.loading'));
    }
    else if (!__VLS_ctx.students.length) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-8 text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-8']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (__VLS_ctx.t('class.noStudents'));
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
        __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase" }));
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
        (__VLS_ctx.t('class.name'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase" }));
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
        (__VLS_ctx.t('class.studentStatus'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase" }));
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-right']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
        (__VLS_ctx.t('common.actions'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.tbody, __VLS_intrinsics.tbody)(__assign({ class: "bg-white divide-y divide-gray-200" }));
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
        /** @type {__VLS_StyleScopedClasses['divide-gray-200']} */ ;
        var _loop_2 = function (stu) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({
                key: (stu.id),
            });
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-4 py-2 text-sm text-gray-900" }));
            /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
            (stu.realname || stu.nickname || stu.username);
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-4 py-2 text-sm" }));
            /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium" }, { class: (__VLS_ctx.studentStatusClass(stu.studentStatus)) }));
            /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
            /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
            /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            (__VLS_ctx.studentStatusLabel(stu.studentStatus));
            __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-4 py-2 text-right text-sm" }));
            /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-right']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ onChange: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!(__VLS_ctx.showStudentsModal))
                        return;
                    if (!!(__VLS_ctx.studentsLoading))
                        return;
                    if (!!(!__VLS_ctx.students.length))
                        return;
                    __VLS_ctx.onStatusChange(stu, $event.target.value);
                    // @ts-ignore
                    [t, t, t, t, t, studentsLoading, students, students, studentStatusClass, studentStatusLabel, onStatusChange,];
                } }, { value: (stu.studentStatus) }), { class: "border border-gray-300 rounded px-2 py-1 text-sm focus:outline-none focus:ring-1 focus:ring-indigo-500" }));
            /** @type {__VLS_StyleScopedClasses['border']} */ ;
            /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
            /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
            /** @type {__VLS_StyleScopedClasses['focus:ring-1']} */ ;
            /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
                value: (1),
            });
            (__VLS_ctx.t('class.statusStudying'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
                value: (2),
            });
            (__VLS_ctx.t('class.statusGraduated'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
                value: (3),
            });
            (__VLS_ctx.t('class.statusTransferred'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
                value: (4),
            });
            (__VLS_ctx.t('class.statusDropped'));
            // @ts-ignore
            [t, t, t, t,];
        };
        for (var _e = 0, _f = __VLS_vFor((__VLS_ctx.students)); _e < _f.length; _e++) {
            var stu = _f[_e][0];
            _loop_2(stu);
        }
    }
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
    (__VLS_ctx.t('class.deleteClass'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm text-gray-600 mb-6" }));
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
    (__VLS_ctx.t('class.confirmDelete'));
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
            [t, t, showDeleteModal, showDeleteModal,];
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
    (__VLS_ctx.t('class.cancel'));
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
    (__VLS_ctx.saving ? __VLS_ctx.t('class.loading') : __VLS_ctx.t('class.confirm'));
}
// @ts-ignore
[t, t, t, saving, saving, doDelete,];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
