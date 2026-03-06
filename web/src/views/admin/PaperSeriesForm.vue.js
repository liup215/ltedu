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
var vue_router_1 = require("vue-router");
var paperSeriesService_1 = require("../../services/paperSeriesService");
var syllabusService_1 = require("../../services/syllabusService");
var organisationService_1 = require("../../services/organisationService");
var qualificationService_1 = require("../../services/qualificationService");
var route = (0, vue_router_1.useRoute)();
var router = (0, vue_router_1.useRouter)();
var form = (0, vue_1.ref)({
    name: '',
    syllabusId: 0,
});
var errors = (0, vue_1.ref)({});
var submitError = (0, vue_1.ref)(null);
var isSubmitting = (0, vue_1.ref)(false);
var isEditMode = (0, vue_1.computed)(function () { return !!route.params.id; });
var paperSeriesId = (0, vue_1.computed)(function () { return Number(route.params.id) || null; });
var organisations = (0, vue_1.ref)([]);
var allQualifications = (0, vue_1.ref)([]); // Store all qualifications
var allSyllabuses = (0, vue_1.ref)([]); // Store all syllabuses
var filterOrganisationId = (0, vue_1.ref)(0);
var filterQualificationId = (0, vue_1.ref)(0);
// Fetch initial data for dropdowns and form (if editing)
var fetchOrganisations = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_1;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                _a.trys.push([0, 2, , 3]);
                return [4 /*yield*/, organisationService_1.default.getAllOrganisations({})];
            case 1:
                response = _a.sent();
                organisations.value = response.data.list;
                return [3 /*break*/, 3];
            case 2:
                error_1 = _a.sent();
                console.error('Failed to fetch organisations:', error_1);
                return [3 /*break*/, 3];
            case 3: return [2 /*return*/];
        }
    });
}); };
var fetchAllQualifications = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_2;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                _a.trys.push([0, 2, , 3]);
                return [4 /*yield*/, qualificationService_1.default.getAllQualifications({ pageIndex: 1, pageSize: 10000 })];
            case 1:
                response = _a.sent();
                allQualifications.value = response.data.list;
                return [3 /*break*/, 3];
            case 2:
                error_2 = _a.sent();
                console.error('Failed to fetch qualifications:', error_2);
                return [3 /*break*/, 3];
            case 3: return [2 /*return*/];
        }
    });
}); };
var fetchAllSyllabuses = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_3;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                _a.trys.push([0, 2, , 3]);
                return [4 /*yield*/, syllabusService_1.default.getAllSyllabuses({ pageIndex: 1, pageSize: 10000 })];
            case 1:
                response = _a.sent();
                allSyllabuses.value = response.data.list;
                return [3 /*break*/, 3];
            case 2:
                error_3 = _a.sent();
                console.error('Failed to fetch syllabuses:', error_3);
                return [3 /*break*/, 3];
            case 3: return [2 /*return*/];
        }
    });
}); };
var fetchPaperSeriesDetails = function (id) { return __awaiter(void 0, void 0, void 0, function () {
    var response, seriesData, error_4;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                _a.trys.push([0, 4, , 5]);
                return [4 /*yield*/, paperSeriesService_1.default.getPaperSeriesById(id)];
            case 1:
                response = _a.sent();
                seriesData = response.data;
                form.value = {
                    id: seriesData.id,
                    name: seriesData.name,
                    syllabusId: seriesData.syllabusId,
                };
                if (!(seriesData.syllabus && seriesData.syllabus.qualification)) return [3 /*break*/, 3];
                filterOrganisationId.value = seriesData.syllabus.qualification.organisationId || 0;
                // Wait for qualifications to be filtered by organisation before setting qualification filter
                return [4 /*yield*/, new Promise(function (resolve) { return setTimeout(resolve, 0); })];
            case 2:
                // Wait for qualifications to be filtered by organisation before setting qualification filter
                _a.sent(); // Allow computed prop to update
                filterQualificationId.value = seriesData.syllabus.qualificationId || 0;
                _a.label = 3;
            case 3: return [3 /*break*/, 5];
            case 4:
                error_4 = _a.sent();
                console.error('Failed to fetch paper series details:', error_4);
                submitError.value = 'Failed to load paper series data.';
                return [3 /*break*/, 5];
            case 5: return [2 /*return*/];
        }
    });
}); };
// Computed properties for filtering dropdowns
var qualificationsForFilter = (0, vue_1.computed)(function () {
    if (!filterOrganisationId.value) {
        return allQualifications.value; // Or an empty array if you prefer no qualifications shown without org
    }
    return allQualifications.value.filter(function (q) { return q.organisationId === filterOrganisationId.value; });
});
var availableSyllabuses = (0, vue_1.computed)(function () {
    var syllabusesToFilter = allSyllabuses.value;
    if (filterOrganisationId.value) {
        syllabusesToFilter = syllabusesToFilter.filter(function (s) { var _a; return ((_a = s.qualification) === null || _a === void 0 ? void 0 : _a.organisationId) === filterOrganisationId.value; });
    }
    if (filterQualificationId.value) {
        syllabusesToFilter = syllabusesToFilter.filter(function (s) { return s.qualificationId === filterQualificationId.value; });
    }
    return syllabusesToFilter;
});
// Watchers to reset dependent filters
(0, vue_1.watch)(filterOrganisationId, function () {
    filterQualificationId.value = 0;
    // If the current form.syllabusId is not in the new availableSyllabuses, reset it
    if (form.value.syllabusId && !availableSyllabuses.value.find(function (s) { return s.id === form.value.syllabusId; })) {
        form.value.syllabusId = 0;
    }
});
(0, vue_1.watch)(filterQualificationId, function () {
    // If the current form.syllabusId is not in the new availableSyllabuses, reset it
    if (form.value.syllabusId && !availableSyllabuses.value.find(function (s) { return s.id === form.value.syllabusId; })) {
        form.value.syllabusId = 0;
    }
});
var validateForm = function () {
    var _a;
    errors.value = {};
    if (!((_a = form.value.name) === null || _a === void 0 ? void 0 : _a.trim())) {
        errors.value.name = 'Name is required.';
    }
    if (!form.value.syllabusId || form.value.syllabusId === 0) {
        errors.value.syllabusId = 'Syllabus is required.';
    }
    return Object.keys(errors.value).length === 0;
};
var handleSubmit = function () { return __awaiter(void 0, void 0, void 0, function () {
    var payload, error_5;
    var _a, _b;
    return __generator(this, function (_c) {
        switch (_c.label) {
            case 0:
                if (!validateForm()) {
                    return [2 /*return*/];
                }
                isSubmitting.value = true;
                submitError.value = null;
                _c.label = 1;
            case 1:
                _c.trys.push([1, 6, 7, 8]);
                payload = {
                    name: form.value.name,
                    syllabusId: form.value.syllabusId,
                };
                if (!(isEditMode.value && paperSeriesId.value)) return [3 /*break*/, 3];
                return [4 /*yield*/, paperSeriesService_1.default.updatePaperSeries(__assign(__assign({}, payload), { id: paperSeriesId.value }))];
            case 2:
                _c.sent();
                return [3 /*break*/, 5];
            case 3: return [4 /*yield*/, paperSeriesService_1.default.createPaperSeries(payload)];
            case 4:
                _c.sent();
                _c.label = 5;
            case 5:
                router.push('/admin/paper-series');
                return [3 /*break*/, 8];
            case 6:
                error_5 = _c.sent();
                console.error('Failed to submit paper series:', error_5);
                submitError.value = ((_b = (_a = error_5.response) === null || _a === void 0 ? void 0 : _a.data) === null || _b === void 0 ? void 0 : _b.message) || 'An unexpected error occurred.';
                return [3 /*break*/, 8];
            case 7:
                isSubmitting.value = false;
                return [7 /*endfinally*/];
            case 8: return [2 /*return*/];
        }
    });
}); };
(0, vue_1.onMounted)(function () { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, Promise.all([
                    fetchOrganisations(),
                    fetchAllQualifications(),
                    fetchAllSyllabuses(),
                ])];
            case 1:
                _a.sent();
                if (!(isEditMode.value && paperSeriesId.value)) return [3 /*break*/, 3];
                return [4 /*yield*/, fetchPaperSeriesDetails(paperSeriesId.value)];
            case 2:
                _a.sent();
                _a.label = 3;
            case 3: return [2 /*return*/];
        }
    });
}); });
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
(__VLS_ctx.isEditMode ? __VLS_ctx.$t('paperSeries.edit') : __VLS_ctx.$t('paperSeries.add'));
__VLS_asFunctionalElement1(__VLS_intrinsics.form, __VLS_intrinsics.form)(__assign({ onSubmit: (__VLS_ctx.handleSubmit) }, { class: "bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4" }));
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-md']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['px-8']} */ ;
/** @type {__VLS_StyleScopedClasses['pt-6']} */ ;
/** @type {__VLS_StyleScopedClasses['pb-8']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4" }));
/** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-gray-700 text-sm font-bold mb-2" }, { for: "name" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
(__VLS_ctx.$t('paperSeries.name'));
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign(__assign(__assign({ id: "name", type: "text", value: (__VLS_ctx.form.name) }, { class: "shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" }), { class: ({ 'border-red-500': __VLS_ctx.errors.name }) }), { required: true }));
/** @type {__VLS_StyleScopedClasses['shadow']} */ ;
/** @type {__VLS_StyleScopedClasses['appearance-none']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['leading-tight']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:shadow-outline']} */ ;
/** @type {__VLS_StyleScopedClasses['border-red-500']} */ ;
if (__VLS_ctx.errors.name) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-red-500 text-xs italic" }));
    /** @type {__VLS_StyleScopedClasses['text-red-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['italic']} */ ;
    (__VLS_ctx.errors.name);
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4 grid grid-cols-1 md:grid-cols-2 gap-4" }));
/** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
/** @type {__VLS_StyleScopedClasses['grid']} */ ;
/** @type {__VLS_StyleScopedClasses['grid-cols-1']} */ ;
/** @type {__VLS_StyleScopedClasses['md:grid-cols-2']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-gray-700 text-sm font-bold mb-2" }, { for: "filter-organisation" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
(__VLS_ctx.$t('syllabusForm.organisation'));
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ id: "filter-organisation", value: (__VLS_ctx.filterOrganisationId) }, { class: "shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" }));
/** @type {__VLS_StyleScopedClasses['shadow']} */ ;
/** @type {__VLS_StyleScopedClasses['appearance-none']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['leading-tight']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:shadow-outline']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (0),
});
(__VLS_ctx.$t('syllabusManagement.allOrganisations'));
for (var _i = 0, _a = __VLS_vFor((__VLS_ctx.organisations)); _i < _a.length; _i++) {
    var org = _a[_i][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (org.id),
        value: (org.id),
    });
    (org.name);
    // @ts-ignore
    [isEditMode, $t, $t, $t, $t, $t, handleSubmit, form, errors, errors, errors, filterOrganisationId, organisations,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-gray-700 text-sm font-bold mb-2" }, { for: "filter-qualification" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
(__VLS_ctx.$t('syllabusForm.qualification'));
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ id: "filter-qualification", value: (__VLS_ctx.filterQualificationId) }, { class: "shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" }), { disabled: (!__VLS_ctx.filterOrganisationId && !__VLS_ctx.qualificationsForFilter.length) }));
/** @type {__VLS_StyleScopedClasses['shadow']} */ ;
/** @type {__VLS_StyleScopedClasses['appearance-none']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['leading-tight']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:shadow-outline']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (0),
});
(__VLS_ctx.$t('syllabusManagement.allQualifications'));
for (var _b = 0, _c = __VLS_vFor((__VLS_ctx.qualificationsForFilter)); _b < _c.length; _b++) {
    var q = _c[_b][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (q.id),
        value: (q.id),
    });
    (q.name);
    // @ts-ignore
    [$t, $t, filterOrganisationId, filterQualificationId, qualificationsForFilter, qualificationsForFilter,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-6" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "block text-gray-700 text-sm font-bold mb-2" }, { for: "syllabus" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-2']} */ ;
(__VLS_ctx.$t('syllabusManagement.name'));
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign(__assign({ id: "syllabus", value: (__VLS_ctx.form.syllabusId) }, { class: "shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" }), { class: ({ 'border-red-500': __VLS_ctx.errors.syllabusId }) }), { required: true, disabled: (!__VLS_ctx.availableSyllabuses.length) }));
/** @type {__VLS_StyleScopedClasses['shadow']} */ ;
/** @type {__VLS_StyleScopedClasses['appearance-none']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['leading-tight']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:shadow-outline']} */ ;
/** @type {__VLS_StyleScopedClasses['border-red-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (0),
    disabled: true,
});
(__VLS_ctx.availableSyllabuses.length ? __VLS_ctx.$t('paperSeries.selectSyllabus') : __VLS_ctx.$t('paperSeries.noSyllabusMatch'));
for (var _d = 0, _e = __VLS_vFor((__VLS_ctx.availableSyllabuses)); _d < _e.length; _d++) {
    var syllabus = _e[_d][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (syllabus.id),
        value: (syllabus.id),
    });
    (syllabus.name);
    // @ts-ignore
    [$t, $t, $t, form, errors, availableSyllabuses, availableSyllabuses, availableSyllabuses,];
}
if (__VLS_ctx.errors.syllabusId) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-red-500 text-xs italic" }));
    /** @type {__VLS_StyleScopedClasses['text-red-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['italic']} */ ;
    (__VLS_ctx.errors.syllabusId);
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-center justify-between" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ type: "submit" }, { class: "bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" }), { disabled: (__VLS_ctx.isSubmitting) }));
/** @type {__VLS_StyleScopedClasses['bg-indigo-600']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-indigo-700']} */ ;
/** @type {__VLS_StyleScopedClasses['text-white']} */ ;
/** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:shadow-outline']} */ ;
(__VLS_ctx.isEditMode ? __VLS_ctx.$t('paperSeries.edit') : __VLS_ctx.$t('paperSeries.add'));
var __VLS_0;
/** @ts-ignore @type {typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink | typeof __VLS_components.routerLink | typeof __VLS_components.RouterLink} */
routerLink;
// @ts-ignore
var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0(__assign({ to: "/admin/paper-series" }, { class: "inline-block align-baseline font-bold text-sm text-indigo-600 hover:text-indigo-800" })));
var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([__assign({ to: "/admin/paper-series" }, { class: "inline-block align-baseline font-bold text-sm text-indigo-600 hover:text-indigo-800" })], __VLS_functionalComponentArgsRest(__VLS_1), false));
/** @type {__VLS_StyleScopedClasses['inline-block']} */ ;
/** @type {__VLS_StyleScopedClasses['align-baseline']} */ ;
/** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-indigo-800']} */ ;
var __VLS_5 = __VLS_3.slots.default;
(__VLS_ctx.$t('common.cancel'));
// @ts-ignore
[isEditMode, $t, $t, $t, errors, errors, isSubmitting,];
var __VLS_3;
if (__VLS_ctx.submitError) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-red-500 text-xs italic mt-4" }));
    /** @type {__VLS_StyleScopedClasses['text-red-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['italic']} */ ;
    /** @type {__VLS_StyleScopedClasses['mt-4']} */ ;
    (__VLS_ctx.submitError);
}
// @ts-ignore
[submitError, submitError,];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
