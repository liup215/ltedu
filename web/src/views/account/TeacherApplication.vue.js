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
Object.defineProperty(exports, "__esModule", { value: true });
var vue_1 = require("vue");
var vue_i18n_1 = require("vue-i18n");
var vue_router_1 = require("vue-router");
var teacherApplicationService_1 = require("../../services/teacherApplicationService");
var teacher_application_model_1 = require("../../models/teacher-application.model");
var userStore_1 = require("../../stores/userStore");
var authService_1 = require("../../services/authService");
var router = (0, vue_router_1.useRouter)();
var t = (0, vue_i18n_1.useI18n)().t;
var loading = (0, vue_1.ref)(false);
var existingApplication = (0, vue_1.ref)(null);
var showNewApplicationForm = (0, vue_1.ref)(false);
var form = (0, vue_1.ref)({
    motivation: '',
    experience: '',
});
// Load existing application if any
var loadApplication = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, userStore, res, error_1;
    var _a;
    return __generator(this, function (_b) {
        switch (_b.label) {
            case 0:
                _b.trys.push([0, 4, , 5]);
                return [4 /*yield*/, teacherApplicationService_1.teacherApplicationService.getCurrentApplication()];
            case 1:
                response = _b.sent();
                existingApplication.value = response.data;
                if (!(existingApplication.value && existingApplication.value.status === teacher_application_model_1.TeacherApplicationStatus.Approved)) return [3 /*break*/, 3];
                userStore = (0, userStore_1.useUserStore)();
                if (!(userStore.user && !userStore.user.isTeacher)) return [3 /*break*/, 3];
                return [4 /*yield*/, authService_1.authService.getCurrentUserProfile()];
            case 2:
                res = _b.sent();
                if (res.code === 0) {
                    userStore.setUser(res.data);
                }
                else {
                    console.error('更新用户信息失败:', res.message);
                }
                _b.label = 3;
            case 3:
                showNewApplicationForm.value = false;
                return [3 /*break*/, 5];
            case 4:
                error_1 = _b.sent();
                if (((_a = error_1.response) === null || _a === void 0 ? void 0 : _a.status) !== 404) {
                    // Use browser's native alert since we're not using Element Plus
                    alert('加载申请信息失败');
                }
                return [3 /*break*/, 5];
            case 5: return [2 /*return*/];
        }
    });
}); };
var startNewApplication = function () {
    showNewApplicationForm.value = true;
    form.value = {
        motivation: '',
        experience: ''
    };
};
var onSubmit = function () { return __awaiter(void 0, void 0, void 0, function () {
    var res, error_2;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                loading.value = true;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 4, 5, 6]);
                return [4 /*yield*/, teacherApplicationService_1.teacherApplicationService.apply(form.value)];
            case 2:
                res = _a.sent();
                console.log('申请结果:', res.message);
                if (res.code !== 0) {
                    throw new Error(res.message);
                }
                return [4 /*yield*/, loadApplication()];
            case 3:
                _a.sent();
                return [3 /*break*/, 6];
            case 4:
                error_2 = _a.sent();
                alert(error_2.message || '申请提交失败');
                return [3 /*break*/, 6];
            case 5:
                loading.value = false;
                return [7 /*endfinally*/];
            case 6: return [2 /*return*/];
        }
    });
}); };
var getStatusText = function (status) {
    switch (status) {
        case teacher_application_model_1.TeacherApplicationStatus.Pending:
            return t('teacherApplication.pending');
        case teacher_application_model_1.TeacherApplicationStatus.Approved:
            return t('teacherApplication.approved');
        case teacher_application_model_1.TeacherApplicationStatus.Rejected:
            return t('teacherApplication.rejected');
        default:
            return t('teacherApplication.unknownStatus');
    }
};
var formatDate = function (dateString) {
    return new Date(dateString).toLocaleString('zh-CN', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
    });
};
(0, vue_1.onMounted)(function () {
    loadApplication();
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
(__VLS_ctx.$t('teacherApplication.title'));
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-sm text-gray-500" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
(__VLS_ctx.$t('teacherApplication.subtitle'));
if (__VLS_ctx.existingApplication) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white shadow sm:rounded-lg mb-6" }));
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:rounded-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "px-8 py-6" }));
    /** @type {__VLS_StyleScopedClasses['px-8']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-between items-start mb-4" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-start']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-lg font-medium leading-6 text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['leading-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (__VLS_ctx.$t('teacherApplication.statusTitle'));
    if (__VLS_ctx.existingApplication.status === __VLS_ctx.TeacherApplicationStatus.Rejected) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.startNewApplication) }, { class: "inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500" }));
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
        (__VLS_ctx.$t('teacherApplication.reapply'));
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-4" }));
    /** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-center" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium text-gray-700 w-24" }));
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-24']} */ ;
    (__VLS_ctx.$t('teacherApplication.status'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: ({
            'px-3 py-1 rounded-full text-sm font-medium': true,
            'bg-yellow-100 text-yellow-700': __VLS_ctx.existingApplication.status === __VLS_ctx.TeacherApplicationStatus.Pending,
            'bg-green-100 text-green-700': __VLS_ctx.existingApplication.status === __VLS_ctx.TeacherApplicationStatus.Approved,
            'bg-red-100 text-red-700': __VLS_ctx.existingApplication.status === __VLS_ctx.TeacherApplicationStatus.Rejected,
        }) }));
    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-yellow-100']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-yellow-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-green-100']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-green-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-red-100']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-red-700']} */ ;
    (__VLS_ctx.getStatusText(__VLS_ctx.existingApplication.status));
    if (__VLS_ctx.existingApplication.adminNotes) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex" }));
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium text-gray-700 w-24" }));
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-24']} */ ;
        (__VLS_ctx.$t('teacherApplication.adminNotes'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-gray-900" }));
        /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
        (__VLS_ctx.existingApplication.adminNotes);
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium text-gray-700 w-24" }));
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-24']} */ ;
    (__VLS_ctx.$t('teacherApplication.appliedAt'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (__VLS_ctx.formatDate(__VLS_ctx.existingApplication.appliedAt));
    if (__VLS_ctx.existingApplication.reviewedAt) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex" }));
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium text-gray-700 w-24" }));
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-24']} */ ;
        (__VLS_ctx.$t('teacherApplication.reviewedAt'));
        __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-gray-900" }));
        /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
        (__VLS_ctx.formatDate(__VLS_ctx.existingApplication.reviewedAt));
    }
}
if (!__VLS_ctx.existingApplication || (__VLS_ctx.existingApplication.status === __VLS_ctx.TeacherApplicationStatus.Rejected && __VLS_ctx.showNewApplicationForm)) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white shadow sm:rounded-lg" }));
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:rounded-lg']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.form, __VLS_intrinsics.form)(__assign({ onSubmit: (__VLS_ctx.onSubmit) }, { class: "space-y-8 divide-y divide-gray-200 px-8 py-6" }));
    /** @type {__VLS_StyleScopedClasses['space-y-8']} */ ;
    /** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
    /** @type {__VLS_StyleScopedClasses['divide-gray-200']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-8']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-6" }));
    /** @type {__VLS_StyleScopedClasses['space-y-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.h3, __VLS_intrinsics.h3)(__assign({ class: "text-lg font-medium leading-6 text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['leading-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (__VLS_ctx.$t('teacherApplication.formTitle'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    (__VLS_ctx.$t('teacherApplication.formTip'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-6" }));
    /** @type {__VLS_StyleScopedClasses['space-y-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "sm:w-3/4" }));
    /** @type {__VLS_StyleScopedClasses['sm:w-3/4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "motivation" }, { class: "block text-sm font-medium text-gray-700 mb-1" }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
    (__VLS_ctx.$t('teacherApplication.motivation'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.textarea, __VLS_intrinsics.textarea)(__assign(__assign({ id: "motivation", value: (__VLS_ctx.form.motivation), rows: "4" }, { class: "block w-full rounded-md border-gray-300 shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm resize-none" }), { placeholder: (__VLS_ctx.$t('teacherApplication.motivationPlaceholder')), maxlength: "1000", disabled: (__VLS_ctx.loading), required: true }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:border-indigo-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['resize-none']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-sm text-gray-500 text-right" }));
    /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-right']} */ ;
    (__VLS_ctx.form.motivation.length);
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "sm:w-3/4" }));
    /** @type {__VLS_StyleScopedClasses['sm:w-3/4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "experience" }, { class: "block text-sm font-medium text-gray-700 mb-1" }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-1']} */ ;
    (__VLS_ctx.$t('teacherApplication.experience'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.textarea, __VLS_intrinsics.textarea)(__assign(__assign({ id: "experience", value: (__VLS_ctx.form.experience), rows: "4" }, { class: "block w-full rounded-md border-gray-300 shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm resize-none" }), { placeholder: (__VLS_ctx.$t('teacherApplication.experiencePlaceholder')), maxlength: "1000", disabled: (__VLS_ctx.loading), required: true }));
    /** @type {__VLS_StyleScopedClasses['block']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:border-indigo-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['resize-none']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-sm text-gray-500 text-right" }));
    /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-right']} */ ;
    (__VLS_ctx.form.experience.length);
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "pt-6 flex justify-end space-x-3" }));
    /** @type {__VLS_StyleScopedClasses['pt-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-end']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-x-3']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(!__VLS_ctx.existingApplication || (__VLS_ctx.existingApplication.status === __VLS_ctx.TeacherApplicationStatus.Rejected && __VLS_ctx.showNewApplicationForm)))
                return;
            __VLS_ctx.router.back();
            // @ts-ignore
            [$t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, $t, existingApplication, existingApplication, existingApplication, existingApplication, existingApplication, existingApplication, existingApplication, existingApplication, existingApplication, existingApplication, existingApplication, existingApplication, existingApplication, teacher_application_model_1.TeacherApplicationStatus, teacher_application_model_1.TeacherApplicationStatus, teacher_application_model_1.TeacherApplicationStatus, teacher_application_model_1.TeacherApplicationStatus, teacher_application_model_1.TeacherApplicationStatus, startNewApplication, getStatusText, formatDate, formatDate, showNewApplicationForm, onSubmit, form, form, form, form, loading, loading, router,];
        } }, { type: "button" }), { class: "inline-flex justify-center rounded-md border border-gray-300 bg-white py-2 px-4 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2" }));
    /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-offset-2']} */ ;
    (__VLS_ctx.$t('teacherApplication.cancel'));
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ type: "submit", disabled: (__VLS_ctx.loading) }, { class: "inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 disabled:opacity-50" }));
    /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-transparent']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-indigo-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-indigo-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['focus:ring-offset-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
    (__VLS_ctx.loading ? __VLS_ctx.$t('teacherApplication.submitting') : __VLS_ctx.$t('teacherApplication.submit'));
}
// @ts-ignore
[$t, $t, $t, loading, loading,];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
