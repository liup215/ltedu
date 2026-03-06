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
var vue_router_1 = require("vue-router");
var pastPaperService_1 = require("../../services/pastPaperService");
var organisationService_1 = require("../../services/organisationService");
var qualificationService_1 = require("../../services/qualificationService");
var syllabusService_1 = require("../../services/syllabusService");
var paperSeriesService_1 = require("../../services/paperSeriesService");
var paperCodeService_1 = require("../../services/paperCodeService");
var route = (0, vue_router_1.useRoute)();
var router = (0, vue_router_1.useRouter)();
var pastPaperId = (0, vue_1.computed)(function () {
    var id = route.params.id;
    return typeof id === 'string' ? parseInt(id, 10) : undefined;
});
var isEditMode = (0, vue_1.computed)(function () { return !!pastPaperId.value; });
var initialDataLoaded = (0, vue_1.ref)(false);
var form = (0, vue_1.ref)({
    name: '',
    year: new Date().getFullYear(),
    syllabusId: 0,
    paperSeriesId: 0,
    paperCodeId: 0,
    questionNumber: 1
});
var isLoading = (0, vue_1.ref)(false);
var errorMessage = (0, vue_1.ref)(null);
var successMessage = (0, vue_1.ref)(null);
// Dropdown state
var organisations = (0, vue_1.ref)([]);
var qualifications = (0, vue_1.ref)([]);
var syllabuses = (0, vue_1.ref)([]);
var paperSeries = (0, vue_1.ref)([]);
var paperCodes = (0, vue_1.ref)([]);
var selectedOrganisationId = (0, vue_1.ref)(null);
var selectedQualificationId = (0, vue_1.ref)(null);
var selectedSyllabusId = (0, vue_1.ref)(null);
var selectedPaperSeriesId = (0, vue_1.ref)(null);
var selectedPaperCodeId = (0, vue_1.ref)(null);
var isLoadingDropdowns = (0, vue_1.ref)({
    organisations: false,
    qualifications: false,
    syllabuses: false,
    paperSeries: false,
    paperCodes: false,
});
// Loading functions
var loadOrganisations = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_1;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                isLoadingDropdowns.value.organisations = true;
                errorMessage.value = null;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, 4, 5]);
                return [4 /*yield*/, organisationService_1.default.getAllOrganisations({})];
            case 2:
                response = _a.sent();
                if (response.code === 0 && response.data) {
                    organisations.value = response.data.list;
                }
                else {
                    errorMessage.value = response.message || 'Failed to load organisations.';
                }
                return [3 /*break*/, 5];
            case 3:
                error_1 = _a.sent();
                errorMessage.value = 'Error loading organisations: ' + error_1.message;
                return [3 /*break*/, 5];
            case 4:
                isLoadingDropdowns.value.organisations = false;
                return [7 /*endfinally*/];
            case 5: return [2 /*return*/];
        }
    });
}); };
var loadQualifications = function (organisationId) { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_2;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (!organisationId) {
                    qualifications.value = [];
                    selectedQualificationId.value = null;
                    return [2 /*return*/];
                }
                isLoadingDropdowns.value.qualifications = true;
                errorMessage.value = null;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, 4, 5]);
                return [4 /*yield*/, qualificationService_1.default.getAllQualifications({ organisationId: organisationId })];
            case 2:
                response = _a.sent();
                if (response.code === 0 && response.data) {
                    qualifications.value = response.data.list;
                }
                else {
                    qualifications.value = [];
                    errorMessage.value = response.message || 'Failed to load qualifications.';
                }
                return [3 /*break*/, 5];
            case 3:
                error_2 = _a.sent();
                qualifications.value = [];
                errorMessage.value = 'Error loading qualifications: ' + error_2.message;
                return [3 /*break*/, 5];
            case 4:
                isLoadingDropdowns.value.qualifications = false;
                return [7 /*endfinally*/];
            case 5: return [2 /*return*/];
        }
    });
}); };
var loadSyllabuses = function (qualificationId) { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_3;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (!qualificationId) {
                    syllabuses.value = [];
                    selectedSyllabusId.value = null;
                    return [2 /*return*/];
                }
                isLoadingDropdowns.value.syllabuses = true;
                errorMessage.value = null;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, 4, 5]);
                return [4 /*yield*/, syllabusService_1.default.getAllSyllabuses({ qualificationId: qualificationId })];
            case 2:
                response = _a.sent();
                if (response.code === 0 && response.data) {
                    syllabuses.value = response.data.list;
                }
                else {
                    syllabuses.value = [];
                    errorMessage.value = response.message || 'Failed to load syllabuses.';
                }
                return [3 /*break*/, 5];
            case 3:
                error_3 = _a.sent();
                syllabuses.value = [];
                errorMessage.value = 'Error loading syllabuses: ' + error_3.message;
                return [3 /*break*/, 5];
            case 4:
                isLoadingDropdowns.value.syllabuses = false;
                return [7 /*endfinally*/];
            case 5: return [2 /*return*/];
        }
    });
}); };
var loadPaperSeries = function (syllabusId) { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_4;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (!syllabusId) {
                    paperSeries.value = [];
                    selectedPaperSeriesId.value = null;
                    return [2 /*return*/];
                }
                isLoadingDropdowns.value.paperSeries = true;
                errorMessage.value = null;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, 4, 5]);
                return [4 /*yield*/, paperSeriesService_1.default.getAllPaperSeries({ syllabusId: syllabusId })];
            case 2:
                response = _a.sent();
                if (response.code === 0 && response.data) {
                    paperSeries.value = response.data.list;
                }
                else {
                    paperSeries.value = [];
                    errorMessage.value = response.message || 'Failed to load paper series.';
                }
                return [3 /*break*/, 5];
            case 3:
                error_4 = _a.sent();
                paperSeries.value = [];
                errorMessage.value = 'Error loading paper series: ' + error_4.message;
                return [3 /*break*/, 5];
            case 4:
                isLoadingDropdowns.value.paperSeries = false;
                return [7 /*endfinally*/];
            case 5: return [2 /*return*/];
        }
    });
}); };
var loadPaperCodes = function (syllabusId) { return __awaiter(void 0, void 0, void 0, function () {
    var response, error_5;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (!syllabusId) {
                    paperCodes.value = [];
                    selectedPaperCodeId.value = null;
                    return [2 /*return*/];
                }
                isLoadingDropdowns.value.paperCodes = true;
                errorMessage.value = null;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, 4, 5]);
                return [4 /*yield*/, paperCodeService_1.default.getAllPaperCodes({ syllabusId: syllabusId })];
            case 2:
                response = _a.sent();
                if (response.code === 0 && response.data) {
                    paperCodes.value = response.data.list;
                }
                else {
                    paperCodes.value = [];
                    errorMessage.value = response.message || 'Failed to load paper codes.';
                }
                return [3 /*break*/, 5];
            case 3:
                error_5 = _a.sent();
                paperCodes.value = [];
                errorMessage.value = 'Error loading paper codes: ' + error_5.message;
                return [3 /*break*/, 5];
            case 4:
                isLoadingDropdowns.value.paperCodes = false;
                return [7 /*endfinally*/];
            case 5: return [2 /*return*/];
        }
    });
}); };
// Watchers
(0, vue_1.watch)(selectedOrganisationId, function (newOrgId) { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                selectedQualificationId.value = null;
                qualifications.value = [];
                selectedSyllabusId.value = null;
                syllabuses.value = [];
                selectedPaperSeriesId.value = null;
                paperSeries.value = [];
                selectedPaperCodeId.value = null;
                paperCodes.value = [];
                form.value.syllabusId = 0;
                form.value.paperSeriesId = 0;
                form.value.paperCodeId = 0;
                if (!newOrgId) return [3 /*break*/, 2];
                return [4 /*yield*/, loadQualifications(newOrgId)];
            case 1:
                _a.sent();
                _a.label = 2;
            case 2: return [2 /*return*/];
        }
    });
}); });
(0, vue_1.watch)(selectedQualificationId, function (newQualId) { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                selectedSyllabusId.value = null;
                syllabuses.value = [];
                selectedPaperSeriesId.value = null;
                paperSeries.value = [];
                selectedPaperCodeId.value = null;
                paperCodes.value = [];
                form.value.syllabusId = 0;
                form.value.paperSeriesId = 0;
                form.value.paperCodeId = 0;
                if (!newQualId) return [3 /*break*/, 2];
                return [4 /*yield*/, loadSyllabuses(newQualId)];
            case 1:
                _a.sent();
                _a.label = 2;
            case 2: return [2 /*return*/];
        }
    });
}); });
(0, vue_1.watch)(selectedSyllabusId, function (newSyllabusId) { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                selectedPaperSeriesId.value = null;
                paperSeries.value = [];
                selectedPaperCodeId.value = null;
                paperCodes.value = [];
                form.value.paperSeriesId = 0;
                form.value.paperCodeId = 0;
                if (!newSyllabusId) return [3 /*break*/, 2];
                form.value.syllabusId = newSyllabusId;
                return [4 /*yield*/, Promise.all([
                        loadPaperSeries(newSyllabusId),
                        loadPaperCodes(newSyllabusId)
                    ])];
            case 1:
                _a.sent();
                return [3 /*break*/, 3];
            case 2:
                form.value.syllabusId = 0;
                _a.label = 3;
            case 3: return [2 /*return*/];
        }
    });
}); });
(0, vue_1.watch)(selectedPaperSeriesId, function (newSeriesId) {
    if (newSeriesId) {
        var selected = paperSeries.value.find(function (s) { return s.id === newSeriesId; });
        if (selected) {
            form.value.paperSeriesId = selected.id || 0;
        }
    }
    else {
        form.value.paperSeriesId = 0;
    }
});
(0, vue_1.watch)(selectedPaperCodeId, function (newCodeId) {
    if (newCodeId) {
        var selected = paperCodes.value.find(function (c) { return c.id === newCodeId; });
        if (selected) {
            form.value.paperCodeId = selected.id || 0;
        }
    }
    else {
        form.value.paperCodeId = 0;
    }
});
var loadPastPaperDetails = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, paperData, initialQualId_1, qualResponse, targetSyllabus, seriesId, codeId, error_6;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (!(isEditMode.value && pastPaperId.value)) return [3 /*break*/, 12];
                isLoading.value = true;
                errorMessage.value = null;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 9, 10, 11]);
                return [4 /*yield*/, pastPaperService_1.default.getPastPaperById(pastPaperId.value)];
            case 2:
                response = _a.sent();
                if (!(response.code === 0 && response.data)) return [3 /*break*/, 8];
                paperData = response.data;
                form.value = {
                    id: paperData.id,
                    name: paperData.name,
                    year: paperData.year,
                    syllabusId: paperData.syllabusId,
                    paperSeriesId: paperData.paperSeriesId,
                    paperCodeId: 0, // Will need to be updated based on your data structure
                    questionNumber: 1 // Will need to be updated based on your data structure
                };
                if (!paperData.syllabus.qualificationId) return [3 /*break*/, 8];
                initialQualId_1 = paperData.syllabus.qualificationId;
                return [4 /*yield*/, loadOrganisations()];
            case 3:
                _a.sent();
                return [4 /*yield*/, qualificationService_1.default.getQualificationById(initialQualId_1)];
            case 4:
                qualResponse = _a.sent();
                if (!(qualResponse.code === 0 && qualResponse.data)) return [3 /*break*/, 8];
                selectedOrganisationId.value = qualResponse.data.organisationId;
                return [4 /*yield*/, loadQualifications(selectedOrganisationId.value)];
            case 5:
                _a.sent();
                selectedQualificationId.value = initialQualId_1;
                return [4 /*yield*/, loadSyllabuses(selectedQualificationId.value)];
            case 6:
                _a.sent();
                targetSyllabus = syllabuses.value.find(function (s) {
                    return s.qualificationId === initialQualId_1;
                });
                if (!targetSyllabus) return [3 /*break*/, 8];
                selectedSyllabusId.value = targetSyllabus.id;
                return [4 /*yield*/, Promise.all([
                        loadPaperSeries(targetSyllabus.id),
                        loadPaperCodes(targetSyllabus.id)
                    ])
                    // Set paper series ID
                ];
            case 7:
                _a.sent();
                // Set paper series ID
                if (paperData.paperSeriesId) {
                    seriesId = paperData.paperSeriesId;
                    if (!isNaN(seriesId)) {
                        selectedPaperSeriesId.value = seriesId;
                    }
                }
                // Set paper code ID
                if (paperData.paperCodeId) {
                    codeId = paperData.paperCodeId;
                    if (!isNaN(codeId)) {
                        selectedPaperCodeId.value = codeId;
                    }
                }
                // Set question number
                if (paperData.questionNumber != null) {
                    form.value.questionNumber = paperData.questionNumber;
                }
                else {
                    form.value.questionNumber = 1; // Default to 1 if not set
                }
                _a.label = 8;
            case 8: return [3 /*break*/, 11];
            case 9:
                error_6 = _a.sent();
                errorMessage.value = 'Error loading past paper details: ' + error_6.message;
                return [3 /*break*/, 11];
            case 10:
                isLoading.value = false;
                initialDataLoaded.value = true;
                return [7 /*endfinally*/];
            case 11: return [3 /*break*/, 14];
            case 12: return [4 /*yield*/, loadOrganisations()];
            case 13:
                _a.sent();
                initialDataLoaded.value = true;
                _a.label = 14;
            case 14: return [2 /*return*/];
        }
    });
}); };
(0, vue_1.onMounted)(function () { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, loadPastPaperDetails()];
            case 1:
                _a.sent();
                return [2 /*return*/];
        }
    });
}); });
var handleSubmit = function () { return __awaiter(void 0, void 0, void 0, function () {
    var response, updatePayload, error_7;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                isLoading.value = true;
                errorMessage.value = null;
                successMessage.value = null;
                if (!selectedSyllabusId.value || !selectedPaperSeriesId.value || !selectedPaperCodeId.value) {
                    errorMessage.value = "Please select all required fields.";
                    isLoading.value = false;
                    return [2 /*return*/];
                }
                if (!form.value.name || !form.value.year || form.value.questionNumber == null) {
                    errorMessage.value = "Please fill in all required fields.";
                    isLoading.value = false;
                    return [2 /*return*/];
                }
                _a.label = 1;
            case 1:
                _a.trys.push([1, 6, 7, 8]);
                response = void 0;
                if (!(isEditMode.value && pastPaperId.value)) return [3 /*break*/, 3];
                updatePayload = __assign(__assign({}, form.value), { id: pastPaperId.value });
                return [4 /*yield*/, pastPaperService_1.default.updatePastPaper(updatePayload)];
            case 2:
                response = _a.sent();
                return [3 /*break*/, 5];
            case 3: return [4 /*yield*/, pastPaperService_1.default.createPastPaper(form.value)];
            case 4:
                response = _a.sent();
                _a.label = 5;
            case 5:
                if (response.code === 0) {
                    successMessage.value = "Past paper ".concat(isEditMode.value ? 'updated' : 'created', " successfully!");
                    setTimeout(function () {
                        router.push({ name: 'AdminPastPaperManagement' });
                    }, 1500);
                }
                else {
                    errorMessage.value = response.message || "Failed to ".concat(isEditMode.value ? 'update' : 'create', " past paper.");
                }
                return [3 /*break*/, 8];
            case 6:
                error_7 = _a.sent();
                errorMessage.value = 'An unexpected error occurred: ' + error_7.message;
                return [3 /*break*/, 8];
            case 7:
                isLoading.value = false;
                return [7 /*endfinally*/];
            case 8: return [2 /*return*/];
        }
    });
}); };
var __VLS_ctx = __assign(__assign({}, {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "p-6 bg-gray-50 min-h-screen" }));
/** @type {__VLS_StyleScopedClasses['p-6']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
/** @type {__VLS_StyleScopedClasses['min-h-screen']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white p-8 rounded-lg shadow-md max-w-2xl mx-auto" }));
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['p-8']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-md']} */ ;
/** @type {__VLS_StyleScopedClasses['max-w-2xl']} */ ;
/** @type {__VLS_StyleScopedClasses['mx-auto']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.h1, __VLS_intrinsics.h1)(__assign({ class: "text-2xl font-semibold text-gray-700 mb-6" }));
/** @type {__VLS_StyleScopedClasses['text-2xl']} */ ;
/** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
(__VLS_ctx.isEditMode ? __VLS_ctx.$t('pastPaper.edit') : __VLS_ctx.$t('pastPaper.add'));
if (__VLS_ctx.isLoading && __VLS_ctx.isEditMode && !__VLS_ctx.initialDataLoaded) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-center py-4" }));
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-gray-600" }));
    /** @type {__VLS_StyleScopedClasses['text-gray-600']} */ ;
    (__VLS_ctx.$t('common.loading'));
}
if (__VLS_ctx.successMessage) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4 p-3 bg-green-100 text-green-700 rounded-md" }));
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['p-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-green-100']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-green-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    (__VLS_ctx.successMessage);
}
if (__VLS_ctx.errorMessage) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-4 p-3 bg-red-100 text-red-700 rounded-md" }));
    /** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['p-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-red-100']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-red-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    (__VLS_ctx.errorMessage);
}
__VLS_asFunctionalElement1(__VLS_intrinsics.form, __VLS_intrinsics.form)(__assign({ onSubmit: (__VLS_ctx.handleSubmit) }, { class: "space-y-6" }));
/** @type {__VLS_StyleScopedClasses['space-y-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "name" }, { class: "block text-sm font-medium text-gray-700" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
(__VLS_ctx.$t('pastPaper.name'));
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-red-500" }));
/** @type {__VLS_StyleScopedClasses['text-red-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "text", id: "name", value: (__VLS_ctx.form.name), required: true }, { class: "mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "organisation" }, { class: "block text-sm font-medium text-gray-700" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
(__VLS_ctx.$t('pastPaper.organisation'));
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-red-500" }));
/** @type {__VLS_StyleScopedClasses['text-red-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign(__assign({ id: "organisation", value: (__VLS_ctx.selectedOrganisationId), required: true }, { class: "mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm" }), { disabled: (__VLS_ctx.isLoadingDropdowns.organisations) }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (null),
    disabled: true,
});
(__VLS_ctx.isLoadingDropdowns.organisations ? __VLS_ctx.$t('common.loading') : __VLS_ctx.$t('pastPaper.organisationPlaceholder'));
for (var _i = 0, _a = __VLS_vFor((__VLS_ctx.organisations)); _i < _a.length; _i++) {
    var org = _a[_i][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (org.id),
        value: (org.id),
    });
    (org.name);
    // @ts-ignore
    [isEditMode, isEditMode, $t, $t, $t, $t, $t, $t, $t, isLoading, initialDataLoaded, successMessage, successMessage, errorMessage, errorMessage, handleSubmit, form, selectedOrganisationId, isLoadingDropdowns, isLoadingDropdowns, organisations,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "qualification" }, { class: "block text-sm font-medium text-gray-700" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
(__VLS_ctx.$t('pastPaper.qualification'));
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-red-500" }));
/** @type {__VLS_StyleScopedClasses['text-red-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ id: "qualification", value: (__VLS_ctx.selectedQualificationId), required: true, disabled: (!__VLS_ctx.selectedOrganisationId || __VLS_ctx.isLoadingDropdowns.qualifications) }, { class: "mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm disabled:bg-gray-100" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['disabled:bg-gray-100']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (null),
    disabled: true,
});
(__VLS_ctx.isLoadingDropdowns.qualifications ? __VLS_ctx.$t('common.loading') : (__VLS_ctx.selectedOrganisationId ? __VLS_ctx.$t('pastPaper.qualificationPlaceholder') : __VLS_ctx.$t('pastPaper.organisationPlaceholder')));
for (var _b = 0, _c = __VLS_vFor((__VLS_ctx.qualifications)); _b < _c.length; _b++) {
    var qual = _c[_b][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (qual.id),
        value: (qual.id),
    });
    (qual.name);
    // @ts-ignore
    [$t, $t, $t, $t, selectedOrganisationId, selectedOrganisationId, isLoadingDropdowns, isLoadingDropdowns, selectedQualificationId, qualifications,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "syllabus" }, { class: "block text-sm font-medium text-gray-700" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
(__VLS_ctx.$t('pastPaper.syllabus'));
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-red-500" }));
/** @type {__VLS_StyleScopedClasses['text-red-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ id: "syllabus", value: (__VLS_ctx.selectedSyllabusId), required: true, disabled: (!__VLS_ctx.selectedQualificationId || __VLS_ctx.isLoadingDropdowns.syllabuses) }, { class: "mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm disabled:bg-gray-100" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['disabled:bg-gray-100']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (null),
    disabled: true,
});
(__VLS_ctx.isLoadingDropdowns.syllabuses ? __VLS_ctx.$t('common.loading') : (__VLS_ctx.selectedQualificationId ? __VLS_ctx.$t('pastPaper.selectSyllabus') : __VLS_ctx.$t('pastPaper.selectQualification')));
for (var _d = 0, _e = __VLS_vFor((__VLS_ctx.syllabuses)); _d < _e.length; _d++) {
    var syl = _e[_d][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (syl.id),
        value: (syl.id),
    });
    (syl.name);
    (syl.code);
    // @ts-ignore
    [$t, $t, $t, $t, isLoadingDropdowns, isLoadingDropdowns, selectedQualificationId, selectedQualificationId, selectedSyllabusId, syllabuses,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "paperSeries" }, { class: "block text-sm font-medium text-gray-700" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
(__VLS_ctx.$t('pastPaper.series'));
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-red-500" }));
/** @type {__VLS_StyleScopedClasses['text-red-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ id: "paperSeries", value: (__VLS_ctx.selectedPaperSeriesId), required: true, disabled: (!__VLS_ctx.selectedSyllabusId || __VLS_ctx.isLoadingDropdowns.paperSeries) }, { class: "mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm disabled:bg-gray-100" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['disabled:bg-gray-100']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (null),
    disabled: true,
});
(__VLS_ctx.isLoadingDropdowns.paperSeries ? __VLS_ctx.$t('common.loading') : (__VLS_ctx.selectedSyllabusId ? __VLS_ctx.$t('pastPaper.selectSeries') : __VLS_ctx.$t('pastPaper.selectSyllabus')));
for (var _f = 0, _g = __VLS_vFor((__VLS_ctx.paperSeries)); _f < _g.length; _f++) {
    var series = _g[_f][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (series.id),
        value: (series.id),
    });
    (series.name);
    // @ts-ignore
    [$t, $t, $t, $t, isLoadingDropdowns, isLoadingDropdowns, selectedSyllabusId, selectedSyllabusId, selectedPaperSeriesId, paperSeries,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "paperCode" }, { class: "block text-sm font-medium text-gray-700" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
(__VLS_ctx.$t('pastPaper.code'));
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-red-500" }));
/** @type {__VLS_StyleScopedClasses['text-red-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ id: "paperCode", value: (__VLS_ctx.selectedPaperCodeId), required: true, disabled: (!__VLS_ctx.selectedSyllabusId || __VLS_ctx.isLoadingDropdowns.paperCodes) }, { class: "mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm disabled:bg-gray-100" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['disabled:bg-gray-100']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (null),
    disabled: true,
});
(__VLS_ctx.isLoadingDropdowns.paperCodes ? __VLS_ctx.$t('common.loading') : (__VLS_ctx.selectedSyllabusId ? __VLS_ctx.$t('pastPaper.selectCode') : __VLS_ctx.$t('pastPaper.selectSyllabus')));
for (var _h = 0, _j = __VLS_vFor((__VLS_ctx.paperCodes)); _h < _j.length; _h++) {
    var code = _j[_h][0];
    __VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
        key: (code.id),
        value: (code.id),
    });
    (code.name);
    // @ts-ignore
    [$t, $t, $t, $t, isLoadingDropdowns, isLoadingDropdowns, selectedSyllabusId, selectedSyllabusId, selectedPaperCodeId, paperCodes,];
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "grid grid-cols-1 md:grid-cols-2 gap-6" }));
/** @type {__VLS_StyleScopedClasses['grid']} */ ;
/** @type {__VLS_StyleScopedClasses['grid-cols-1']} */ ;
/** @type {__VLS_StyleScopedClasses['md:grid-cols-2']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "year" }, { class: "block text-sm font-medium text-gray-700" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
(__VLS_ctx.$t('pastPaper.year'));
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-red-500" }));
/** @type {__VLS_StyleScopedClasses['text-red-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "number", id: "year", required: true }, { class: "mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm" }));
(__VLS_ctx.form.year);
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "questionNumber" }, { class: "block text-sm font-medium text-gray-700" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
(__VLS_ctx.$t('examPaperManagement.questionCount'));
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-red-500" }));
/** @type {__VLS_StyleScopedClasses['text-red-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "number", id: "questionNumber", required: true, min: "0" }, { class: "mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm" }));
(__VLS_ctx.form.questionNumber);
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex justify-end space-x-3 pt-4" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['justify-end']} */ ;
/** @type {__VLS_StyleScopedClasses['space-x-3']} */ ;
/** @type {__VLS_StyleScopedClasses['pt-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign(__assign({ onClick: function () {
        var _a = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            _a[_i] = arguments[_i];
        }
        var $event = _a[0];
        __VLS_ctx.router.push({ name: 'AdminPastPaperManagement' });
        // @ts-ignore
        [$t, $t, form, form, router,];
    } }, { type: "button" }), { class: "px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500" }), { disabled: (__VLS_ctx.isLoading) }));
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-offset-2']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
(__VLS_ctx.$t('common.cancel'));
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ type: "submit" }, { class: "px-4 py-2 bg-blue-600 text-white rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50" }), { disabled: (__VLS_ctx.isLoading || __VLS_ctx.isLoadingDropdowns.organisations || __VLS_ctx.isLoadingDropdowns.qualifications || __VLS_ctx.isLoadingDropdowns.syllabuses) }));
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-blue-600']} */ ;
/** @type {__VLS_StyleScopedClasses['text-white']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-blue-700']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-offset-2']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-blue-500']} */ ;
/** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
if (__VLS_ctx.isLoading) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
    (__VLS_ctx.$t('pastPaper.saving'));
}
else {
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)({});
    (__VLS_ctx.isEditMode ? __VLS_ctx.$t('pastPaper.edit') : __VLS_ctx.$t('pastPaper.add'));
}
// @ts-ignore
[isEditMode, $t, $t, $t, $t, isLoading, isLoading, isLoading, isLoadingDropdowns, isLoadingDropdowns, isLoadingDropdowns,];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
