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
Object.defineProperty(exports, "__esModule", { value: true });
var vue_1 = require("vue");
var vue_i18n_1 = require("vue-i18n");
var analyticsService_1 = require("../../services/analyticsService");
var classService_1 = require("../../services/classService");
var t = (0, vue_i18n_1.useI18n)().t;
var classes = (0, vue_1.ref)([]);
var selectedClassId = (0, vue_1.ref)('');
var loading = (0, vue_1.ref)(false);
var activeTab = (0, vue_1.ref)('students');
var summary = (0, vue_1.ref)(null);
var students = (0, vue_1.ref)([]);
var heatmap = (0, vue_1.ref)(null);
var trends = (0, vue_1.ref)([]);
var warnings = (0, vue_1.ref)([]);
var tabs = (0, vue_1.computed)(function () { return [
    { key: 'students', label: t('analytics.studentPerformance') },
    { key: 'heatmap', label: t('analytics.heatmap') },
    { key: 'trends', label: t('analytics.trends') },
    { key: 'warnings', label: t('analytics.earlyWarnings') + (warnings.value.length ? " (".concat(warnings.value.length, ")") : '') }
]; });
var maxTrendValue = (0, vue_1.computed)(function () { return Math.max.apply(Math, __spreadArray([1], trends.value.map(function (p) { return p.totalAttempts; }), false)); });
(0, vue_1.onMounted)(function () { return __awaiter(void 0, void 0, void 0, function () {
    var res, _a;
    return __generator(this, function (_b) {
        switch (_b.label) {
            case 0:
                _b.trys.push([0, 2, , 3]);
                return [4 /*yield*/, classService_1.default.list({ pageSize: 100, pageIndex: 1 })];
            case 1:
                res = _b.sent();
                if (res.code === 0 && res.data) {
                    classes.value = (res.data.list || []).map(function (c) { return ({ id: c.id, name: c.name }); });
                }
                return [3 /*break*/, 3];
            case 2:
                _a = _b.sent();
                return [3 /*break*/, 3];
            case 3: return [2 /*return*/];
        }
    });
}); });
function loadClassData() {
    return __awaiter(this, void 0, void 0, function () {
        var id, _a, sumRes, studentsRes, heatmapRes, trendsRes, warningsRes, _b;
        var _c, _d;
        return __generator(this, function (_e) {
            switch (_e.label) {
                case 0:
                    if (!selectedClassId.value)
                        return [2 /*return*/];
                    loading.value = true;
                    id = Number(selectedClassId.value);
                    _e.label = 1;
                case 1:
                    _e.trys.push([1, 3, 4, 5]);
                    return [4 /*yield*/, Promise.all([
                            analyticsService_1.default.getClassSummary(id),
                            analyticsService_1.default.getStudentPerformanceList(id),
                            analyticsService_1.default.getClassHeatmap(id),
                            analyticsService_1.default.getAttemptTrends({ classId: id }),
                            analyticsService_1.default.getEarlyWarnings(id)
                        ])];
                case 2:
                    _a = _e.sent(), sumRes = _a[0], studentsRes = _a[1], heatmapRes = _a[2], trendsRes = _a[3], warningsRes = _a[4];
                    if (sumRes.code === 0)
                        summary.value = sumRes.data;
                    if (studentsRes.code === 0)
                        students.value = ((_c = studentsRes.data) === null || _c === void 0 ? void 0 : _c.list) || [];
                    if (heatmapRes.code === 0)
                        heatmap.value = heatmapRes.data;
                    if (trendsRes.code === 0)
                        trends.value = trendsRes.data || [];
                    if (warningsRes.code === 0)
                        warnings.value = ((_d = warningsRes.data) === null || _d === void 0 ? void 0 : _d.list) || [];
                    return [3 /*break*/, 5];
                case 3:
                    _b = _e.sent();
                    return [3 /*break*/, 5];
                case 4:
                    loading.value = false;
                    return [7 /*endfinally*/];
                case 5: return [2 /*return*/];
            }
        });
    });
}
function masteryColor(value) {
    if (value >= 70)
        return 'text-green-600';
    if (value >= 40)
        return 'text-yellow-600';
    return 'text-red-600';
}
function barColor(value) {
    if (value >= 70)
        return 'bg-green-500';
    if (value >= 40)
        return 'bg-yellow-400';
    return 'bg-red-400';
}
function heatmapCellClass(value) {
    if (value >= 70)
        return 'bg-green-100 text-green-800';
    if (value >= 30)
        return 'bg-yellow-100 text-yellow-800';
    return 'bg-red-100 text-red-800';
}
function severityBorderClass(severity) {
    if (severity === 'high')
        return 'border-red-500';
    if (severity === 'medium')
        return 'border-yellow-500';
    return 'border-blue-400';
}
function severityBadgeClass(severity) {
    if (severity === 'high')
        return 'bg-red-100 text-red-800';
    if (severity === 'medium')
        return 'bg-yellow-100 text-yellow-800';
    return 'bg-blue-100 text-blue-800';
}
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
(__VLS_ctx.t('analytics.title'));
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-2 text-sm text-gray-600" }));
/** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
(__VLS_ctx.t('analytics.subtitle'));
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-6 flex items-center gap-4" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "text-sm font-medium text-gray-700" }));
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
(__VLS_ctx.t('class.title', 'Class'));
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ onChange: (__VLS_ctx.loadClassData) }, { value: (__VLS_ctx.selectedClassId) }), { class: "border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 w-64" }));
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['w-64']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: "",
});
(__VLS_ctx.t('analytics.selectClass'));
for (var _i = 0, _a = __VLS_vFor((__VLS_ctx.classes)); _i < _a.length; _i++) {
    var cls = _a[_i][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (cls.id),
        value: (cls.id),
    });
    (cls.name);
    // @ts-ignore
    [t, t, t, t, loadClassData, selectedClassId, classes,];
}
if (__VLS_ctx.selectedClassId) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.loadClassData) }, { class: "inline-flex items-center px-3 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 bg-white hover:bg-gray-50" }));
    /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
}
if (!__VLS_ctx.selectedClassId) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-16 text-gray-400" }));
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-16']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "mx-auto h-16 w-16 mb-4" }, { fill: "none", viewBox: "0 0 24 24", stroke: "currentColor" }));
    /** @type {__VLS_StyleScopedClasses['mx-auto']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-16']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-16']} */ ;
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
        'stroke-linecap': "round",
        'stroke-linejoin': "round",
        'stroke-width': "1",
        d: "M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z",
    });
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-lg" }));
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    (__VLS_ctx.t('analytics.selectClass'));
}
else {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
    if (__VLS_ctx.loading) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-8 text-gray-500" }));
        /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-8']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        (__VLS_ctx.t('analytics.loading'));
    }
    else {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
        if (__VLS_ctx.summary) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "grid grid-cols-2 md:grid-cols-4 gap-4 mb-8" }));
            /** @type {__VLS_StyleScopedClasses['grid']} */ ;
            /** @type {__VLS_StyleScopedClasses['grid-cols-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['md:grid-cols-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['mb-8']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow p-4 text-center" }));
            /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
            /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
            /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-3xl font-bold text-indigo-600" }));
            /** @type {__VLS_StyleScopedClasses['text-3xl']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
            (__VLS_ctx.summary.totalStudents);
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-500 mt-1" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
            (__VLS_ctx.t('analytics.totalStudents'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow p-4 text-center" }));
            /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
            /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
            /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-3xl font-bold text-green-600" }));
            /** @type {__VLS_StyleScopedClasses['text-3xl']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-green-600']} */ ;
            (__VLS_ctx.summary.activeStudents);
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-500 mt-1" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
            (__VLS_ctx.t('analytics.activeStudents'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow p-4 text-center" }));
            /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
            /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
            /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-3xl font-bold" }, { class: (__VLS_ctx.masteryColor(__VLS_ctx.summary.avgMastery)) }));
            /** @type {__VLS_StyleScopedClasses['text-3xl']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
            (__VLS_ctx.summary.avgMastery.toFixed(1));
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-500 mt-1" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
            (__VLS_ctx.t('analytics.avgMastery'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow p-4 text-center" }));
            /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
            /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
            /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-3xl font-bold" }, { class: (__VLS_ctx.masteryColor(__VLS_ctx.summary.avgAccuracy)) }));
            /** @type {__VLS_StyleScopedClasses['text-3xl']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
            (__VLS_ctx.summary.avgAccuracy.toFixed(1));
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-500 mt-1" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
            (__VLS_ctx.t('analytics.avgAccuracy'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow p-4 text-center" }));
            /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
            /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
            /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-3xl font-bold text-blue-600" }));
            /** @type {__VLS_StyleScopedClasses['text-3xl']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-blue-600']} */ ;
            (__VLS_ctx.summary.avgCoverage.toFixed(1));
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-500 mt-1" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
            (__VLS_ctx.t('analytics.avgCoverage'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow p-4 text-center" }));
            /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
            /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
            /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-3xl font-bold text-gray-700" }));
            /** @type {__VLS_StyleScopedClasses['text-3xl']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
            (__VLS_ctx.summary.totalAttempts.toLocaleString());
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-500 mt-1" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
            (__VLS_ctx.t('analytics.totalAttempts'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow p-4 text-center" }));
            /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
            /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
            /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-3xl font-bold text-purple-600" }));
            /** @type {__VLS_StyleScopedClasses['text-3xl']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-purple-600']} */ ;
            (__VLS_ctx.summary.weeklyAttempts.toLocaleString());
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-500 mt-1" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
            (__VLS_ctx.t('analytics.weeklyAttempts'));
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white rounded-lg shadow p-4 text-center" }));
            /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
            /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
            /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
            /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-3xl font-bold text-red-600" }));
            /** @type {__VLS_StyleScopedClasses['text-3xl']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-red-600']} */ ;
            (__VLS_ctx.summary.atRiskCount);
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-sm text-gray-500 mt-1" }));
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
            /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
            (__VLS_ctx.t('analytics.atRiskStudents'));
        }
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4 border-b border-gray-200" }));
        /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-b']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.nav, __VLS_intrinsics.nav)(__assign({ class: "-mb-px flex space-x-6 overflow-x-auto" }));
        /** @type {__VLS_StyleScopedClasses['-mb-px']} */ ;
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['space-x-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['overflow-x-auto']} */ ;
        var _loop_1 = function (tab) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!!(!__VLS_ctx.selectedClassId))
                        return;
                    if (!!(__VLS_ctx.loading))
                        return;
                    __VLS_ctx.activeTab = tab.key;
                    // @ts-ignore
                    [t, t, t, t, t, t, t, t, t, t, loadClassData, selectedClassId, selectedClassId, loading, summary, summary, summary, summary, summary, summary, summary, summary, summary, summary, summary, masteryColor, masteryColor, tabs, activeTab,];
                } }, { key: (tab.key) }), { class: (['pb-2 text-sm font-medium border-b-2 transition-colors whitespace-nowrap', __VLS_ctx.activeTab === tab.key ? 'border-indigo-500 text-indigo-600' : 'border-transparent text-gray-500 hover:text-gray-700']) }));
            /** @type {__VLS_StyleScopedClasses['pb-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            /** @type {__VLS_StyleScopedClasses['border-b-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['transition-colors']} */ ;
            /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
            (tab.label);
            // @ts-ignore
            [activeTab,];
        };
        for (var _b = 0, _c = __VLS_vFor((__VLS_ctx.tabs)); _b < _c.length; _b++) {
            var tab = _c[_b][0];
            _loop_1(tab);
        }
        if (__VLS_ctx.activeTab === 'students') {
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
            if (!__VLS_ctx.students.length) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-8 text-gray-400" }));
                /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-8']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
                (__VLS_ctx.t('analytics.noData'));
            }
            else {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white shadow rounded-lg overflow-x-auto" }));
                /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
                /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
                /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
                /** @type {__VLS_StyleScopedClasses['overflow-x-auto']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.table, __VLS_intrinsics.table)(__assign({ class: "min-w-full divide-y divide-gray-200" }));
                /** @type {__VLS_StyleScopedClasses['min-w-full']} */ ;
                /** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
                /** @type {__VLS_StyleScopedClasses['divide-gray-200']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.thead, __VLS_intrinsics.thead)(__assign({ class: "bg-gray-50" }));
                /** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
                __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase" }));
                /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
                (__VLS_ctx.t('analytics.student'));
                __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase" }));
                /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
                (__VLS_ctx.t('analytics.mastery'));
                __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase" }));
                /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
                (__VLS_ctx.t('analytics.coverage'));
                __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase" }));
                /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
                (__VLS_ctx.t('analytics.accuracy'));
                __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase" }));
                /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
                (__VLS_ctx.t('analytics.totalAttempts'));
                __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase" }));
                /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
                (__VLS_ctx.t('analytics.lastActive'));
                __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase" }));
                /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                /** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
                (__VLS_ctx.t('analytics.atRisk'));
                __VLS_asFunctionalElement1(__VLS_intrinsics.tbody, __VLS_intrinsics.tbody)(__assign({ class: "bg-white divide-y divide-gray-200" }));
                /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
                /** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
                /** @type {__VLS_StyleScopedClasses['divide-gray-200']} */ ;
                for (var _d = 0, _e = __VLS_vFor((__VLS_ctx.students)); _d < _e.length; _d++) {
                    var s = _e[_d][0];
                    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({
                        key: (s.userId),
                    });
                    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-4 py-3 text-sm font-medium text-gray-900" }));
                    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
                    (s.userName);
                    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-4 py-3 text-sm" }));
                    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-center gap-2" }));
                    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                    /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "w-24 bg-gray-200 rounded-full h-2" }));
                    /** @type {__VLS_StyleScopedClasses['w-24']} */ ;
                    /** @type {__VLS_StyleScopedClasses['bg-gray-200']} */ ;
                    /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
                    /** @type {__VLS_StyleScopedClasses['h-2']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign(__assign({ class: "h-2 rounded-full" }, { class: (__VLS_ctx.barColor(s.masteryLevel)) }), { style: ({ width: s.masteryLevel + '%' }) }));
                    /** @type {__VLS_StyleScopedClasses['h-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: (__VLS_ctx.masteryColor(s.masteryLevel)) }));
                    (s.masteryLevel.toFixed(1));
                    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-4 py-3 text-sm text-gray-700" }));
                    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
                    (s.coverageLevel.toFixed(1));
                    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-4 py-3 text-sm" }, { class: (__VLS_ctx.masteryColor(s.accuracyRate)) }));
                    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                    (s.accuracyRate.toFixed(1));
                    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-4 py-3 text-sm text-gray-700" }));
                    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
                    (s.totalAttempts);
                    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-4 py-3 text-sm text-gray-500" }));
                    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                    (s.lastActiveAt ? new Date(s.lastActiveAt).toLocaleDateString() : '—');
                    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-4 py-3 text-sm" }));
                    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                    if (s.isAtRisk) {
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-red-100 text-red-800" }));
                        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
                        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
                        /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
                        /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
                        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                        /** @type {__VLS_StyleScopedClasses['bg-red-100']} */ ;
                        /** @type {__VLS_StyleScopedClasses['text-red-800']} */ ;
                        (__VLS_ctx.t('analytics.atRisk'));
                    }
                    else {
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-green-600 text-xs" }));
                        /** @type {__VLS_StyleScopedClasses['text-green-600']} */ ;
                        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                    }
                    // @ts-ignore
                    [t, t, t, t, t, t, t, t, t, masteryColor, masteryColor, activeTab, students, students, barColor,];
                }
            }
        }
        if (__VLS_ctx.activeTab === 'heatmap') {
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
            if (!__VLS_ctx.heatmap || !__VLS_ctx.heatmap.chapters.length) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-8 text-gray-400" }));
                /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-8']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
                (__VLS_ctx.t('analytics.noData'));
            }
            else {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white shadow rounded-lg overflow-x-auto" }));
                /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
                /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
                /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
                /** @type {__VLS_StyleScopedClasses['overflow-x-auto']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.table, __VLS_intrinsics.table)(__assign({ class: "min-w-full text-xs" }));
                /** @type {__VLS_StyleScopedClasses['min-w-full']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.thead, __VLS_intrinsics.thead)(__assign({ class: "bg-gray-50" }));
                /** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
                __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-3 py-2 text-left font-medium text-gray-500 min-w-[160px]" }));
                /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                /** @type {__VLS_StyleScopedClasses['min-w-[160px]']} */ ;
                (__VLS_ctx.t('analytics.chapter'));
                __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ class: "px-3 py-2 text-center font-medium text-gray-500" }));
                /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                for (var _f = 0, _g = __VLS_vFor((__VLS_ctx.heatmap.students)); _f < _g.length; _f++) {
                    var student = _g[_f][0];
                    __VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign(__assign({ key: (student.userId) }, { class: "px-2 py-2 text-center font-medium text-gray-500 min-w-[60px] truncate max-w-[80px]" }), { title: (student.userName) }));
                    /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
                    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                    /** @type {__VLS_StyleScopedClasses['min-w-[60px]']} */ ;
                    /** @type {__VLS_StyleScopedClasses['truncate']} */ ;
                    /** @type {__VLS_StyleScopedClasses['max-w-[80px]']} */ ;
                    (student.userName.slice(0, 6));
                    // @ts-ignore
                    [t, t, activeTab, heatmap, heatmap, heatmap,];
                }
                __VLS_asFunctionalElement1(__VLS_intrinsics.tbody, __VLS_intrinsics.tbody)(__assign({ class: "divide-y divide-gray-100" }));
                /** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
                /** @type {__VLS_StyleScopedClasses['divide-gray-100']} */ ;
                for (var _h = 0, _j = __VLS_vFor((__VLS_ctx.heatmap.chapters)); _h < _j.length; _h++) {
                    var row = _j[_h][0];
                    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({
                        key: (row.chapterId),
                    });
                    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-3 py-2 font-medium text-gray-800 truncate max-w-[160px]" }, { title: (row.chapterName) }));
                    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-gray-800']} */ ;
                    /** @type {__VLS_StyleScopedClasses['truncate']} */ ;
                    /** @type {__VLS_StyleScopedClasses['max-w-[160px]']} */ ;
                    (row.chapterName);
                    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-3 py-2 text-center" }));
                    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "px-1.5 py-0.5 rounded text-xs font-medium" }, { class: (__VLS_ctx.heatmapCellClass(row.avgMastery)) }));
                    /** @type {__VLS_StyleScopedClasses['px-1.5']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
                    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                    (row.avgMastery.toFixed(0));
                    for (var _k = 0, _l = __VLS_vFor((row.studentData)); _k < _l.length; _k++) {
                        var score = _l[_k][0];
                        __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ key: (score.userId) }, { class: "px-2 py-2 text-center" }));
                        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
                        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
                        /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
                        if (score.isCovered) {
                            __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "px-1.5 py-0.5 rounded text-xs font-medium" }, { class: (__VLS_ctx.heatmapCellClass(score.masteryLevel)) }));
                            /** @type {__VLS_StyleScopedClasses['px-1.5']} */ ;
                            /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
                            /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
                            /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                            (score.masteryLevel.toFixed(0));
                        }
                        else {
                            __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-gray-300 text-xs" }));
                            /** @type {__VLS_StyleScopedClasses['text-gray-300']} */ ;
                            /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                        }
                        // @ts-ignore
                        [heatmap, heatmapCellClass, heatmapCellClass,];
                    }
                    // @ts-ignore
                    [];
                }
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-3 border-t border-gray-100 flex items-center gap-4 text-xs text-gray-500" }));
                /** @type {__VLS_StyleScopedClasses['p-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['border-t']} */ ;
                /** @type {__VLS_StyleScopedClasses['border-gray-100']} */ ;
                /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "flex items-center gap-1" }));
                /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['gap-1']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-block w-3 h-3 rounded bg-red-200" }));
                /** @type {__VLS_StyleScopedClasses['inline-block']} */ ;
                /** @type {__VLS_StyleScopedClasses['w-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['h-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
                /** @type {__VLS_StyleScopedClasses['bg-red-200']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "flex items-center gap-1" }));
                /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['gap-1']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-block w-3 h-3 rounded bg-yellow-200" }));
                /** @type {__VLS_StyleScopedClasses['inline-block']} */ ;
                /** @type {__VLS_StyleScopedClasses['w-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['h-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
                /** @type {__VLS_StyleScopedClasses['bg-yellow-200']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "flex items-center gap-1" }));
                /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['gap-1']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-block w-3 h-3 rounded bg-green-200" }));
                /** @type {__VLS_StyleScopedClasses['inline-block']} */ ;
                /** @type {__VLS_StyleScopedClasses['w-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['h-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
                /** @type {__VLS_StyleScopedClasses['bg-green-200']} */ ;
            }
        }
        if (__VLS_ctx.activeTab === 'trends') {
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
            if (!__VLS_ctx.trends.length) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-8 text-gray-400" }));
                /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-8']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
                (__VLS_ctx.t('analytics.noData'));
            }
            else {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white shadow rounded-lg p-4" }));
                /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
                /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
                /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
                /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "overflow-x-auto" }));
                /** @type {__VLS_StyleScopedClasses['overflow-x-auto']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-end gap-1 h-40 min-w-max" }));
                /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                /** @type {__VLS_StyleScopedClasses['items-end']} */ ;
                /** @type {__VLS_StyleScopedClasses['gap-1']} */ ;
                /** @type {__VLS_StyleScopedClasses['h-40']} */ ;
                /** @type {__VLS_StyleScopedClasses['min-w-max']} */ ;
                for (var _m = 0, _o = __VLS_vFor((__VLS_ctx.trends)); _m < _o.length; _m++) {
                    var point = _o[_m][0];
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex flex-col items-center gap-0.5 w-8" }, { key: (point.date) }));
                    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                    /** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
                    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                    /** @type {__VLS_StyleScopedClasses['gap-0.5']} */ ;
                    /** @type {__VLS_StyleScopedClasses['w-8']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign(__assign({ class: "w-4 bg-indigo-400 rounded-t" }, { style: ({ height: Math.max(2, (point.totalAttempts / __VLS_ctx.maxTrendValue) * 120) + 'px' }) }), { title: ("".concat(point.date, ": ").concat(point.totalAttempts, " attempts")) }));
                    /** @type {__VLS_StyleScopedClasses['w-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['bg-indigo-400']} */ ;
                    /** @type {__VLS_StyleScopedClasses['rounded-t']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign(__assign({ class: "w-4 bg-green-400 rounded-t" }, { style: ({ height: Math.max(2, (point.correctAttempts / __VLS_ctx.maxTrendValue) * 120) + 'px' }) }), { title: ("".concat(point.date, ": ").concat(point.correctAttempts, " correct")) }));
                    /** @type {__VLS_StyleScopedClasses['w-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['bg-green-400']} */ ;
                    /** @type {__VLS_StyleScopedClasses['rounded-t']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-gray-400 text-xs transform -rotate-45 origin-top-left mt-1 hidden sm:block" }));
                    /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                    /** @type {__VLS_StyleScopedClasses['transform']} */ ;
                    /** @type {__VLS_StyleScopedClasses['-rotate-45']} */ ;
                    /** @type {__VLS_StyleScopedClasses['origin-top-left']} */ ;
                    /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
                    /** @type {__VLS_StyleScopedClasses['hidden']} */ ;
                    /** @type {__VLS_StyleScopedClasses['sm:block']} */ ;
                    (point.date.slice(5));
                    // @ts-ignore
                    [t, activeTab, trends, trends, maxTrendValue, maxTrendValue,];
                }
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-3 flex items-center gap-4 text-xs text-gray-500" }));
                /** @type {__VLS_StyleScopedClasses['mt-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "flex items-center gap-1" }));
                /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['gap-1']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-block w-3 h-3 rounded bg-indigo-400" }));
                /** @type {__VLS_StyleScopedClasses['inline-block']} */ ;
                /** @type {__VLS_StyleScopedClasses['w-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['h-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
                /** @type {__VLS_StyleScopedClasses['bg-indigo-400']} */ ;
                (__VLS_ctx.t('analytics.total_attempts'));
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "flex items-center gap-1" }));
                /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['gap-1']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "inline-block w-3 h-3 rounded bg-green-400" }));
                /** @type {__VLS_StyleScopedClasses['inline-block']} */ ;
                /** @type {__VLS_StyleScopedClasses['w-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['h-3']} */ ;
                /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
                /** @type {__VLS_StyleScopedClasses['bg-green-400']} */ ;
                (__VLS_ctx.t('analytics.correct_attempts'));
            }
        }
        if (__VLS_ctx.activeTab === 'warnings') {
            __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
            if (!__VLS_ctx.warnings.length) {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-8 text-green-600" }));
                /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
                /** @type {__VLS_StyleScopedClasses['py-8']} */ ;
                /** @type {__VLS_StyleScopedClasses['text-green-600']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "mx-auto h-12 w-12 mb-2" }, { fill: "none", viewBox: "0 0 24 24", stroke: "currentColor" }));
                /** @type {__VLS_StyleScopedClasses['mx-auto']} */ ;
                /** @type {__VLS_StyleScopedClasses['h-12']} */ ;
                /** @type {__VLS_StyleScopedClasses['w-12']} */ ;
                /** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
                __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
                    'stroke-linecap': "round",
                    'stroke-linejoin': "round",
                    'stroke-width': "2",
                    d: "M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z",
                });
                (__VLS_ctx.t('analytics.noWarnings'));
            }
            else {
                __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "space-y-3" }));
                /** @type {__VLS_StyleScopedClasses['space-y-3']} */ ;
                for (var _p = 0, _q = __VLS_vFor((__VLS_ctx.warnings)); _p < _q.length; _p++) {
                    var w = _q[_p][0];
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (w.userId) }, { class: (['bg-white shadow rounded-lg p-4 border-l-4', __VLS_ctx.severityBorderClass(w.severity)]) }));
                    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
                    /** @type {__VLS_StyleScopedClasses['shadow']} */ ;
                    /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
                    /** @type {__VLS_StyleScopedClasses['p-4']} */ ;
                    /** @type {__VLS_StyleScopedClasses['border-l-4']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-start justify-between" }));
                    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                    /** @type {__VLS_StyleScopedClasses['items-start']} */ ;
                    /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
                    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "font-medium text-gray-900 flex items-center gap-2" }));
                    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
                    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                    /** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
                    (w.userName);
                    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: (['text-xs px-2 py-0.5 rounded-full font-medium', __VLS_ctx.severityBadgeClass(w.severity)]) }));
                    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
                    /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
                    /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
                    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
                    (w.severity);
                    __VLS_asFunctionalElement1(__VLS_intrinsics.ul, __VLS_intrinsics.ul)(__assign({ class: "mt-2 space-y-1" }));
                    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
                    /** @type {__VLS_StyleScopedClasses['space-y-1']} */ ;
                    for (var _r = 0, _s = __VLS_vFor((w.reasons)); _r < _s.length; _r++) {
                        var _t = _s[_r], reason = _t[0], i = _t[1];
                        __VLS_asFunctionalElement1(__VLS_intrinsics.li, __VLS_intrinsics.li)(__assign({ key: (i) }, { class: "text-sm text-gray-600 flex items-center gap-1" }));
                        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
                        /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
                        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
                        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
                        /** @type {__VLS_StyleScopedClasses['gap-1']} */ ;
                        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-red-400" }));
                        /** @type {__VLS_StyleScopedClasses['text-red-400']} */ ;
                        (reason);
                        // @ts-ignore
                        [t, t, t, activeTab, warnings, warnings, severityBorderClass, severityBadgeClass,];
                    }
                    // @ts-ignore
                    [];
                }
            }
        }
    }
}
// @ts-ignore
[];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
