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
var mcpTokenService_1 = require("../../services/mcpTokenService");
var notification_1 = require("../../utils/notification");
var config_1 = require("../../const/config");
var loading = (0, vue_1.ref)(false);
var creating = (0, vue_1.ref)(false);
var tokens = (0, vue_1.ref)([]);
var total = (0, vue_1.ref)(0);
var baseUrl = (0, vue_1.ref)('');
var showFull = (0, vue_1.ref)({});
var pendingDelete = (0, vue_1.ref)(null);
var newToken = (0, vue_1.ref)({ name: '', expiresAt: undefined });
(0, vue_1.onMounted)(function () { return __awaiter(void 0, void 0, void 0, function () {
    var _a;
    return __generator(this, function (_b) {
        switch (_b.label) {
            case 0:
                _a = baseUrl;
                return [4 /*yield*/, (0, config_1.API_BASE_URL)()];
            case 1:
                _a.value = _b.sent();
                loadTokens();
                return [2 /*return*/];
        }
    });
}); });
function loadTokens() {
    return __awaiter(this, void 0, void 0, function () {
        var response, error_1;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    loading.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 3, 4, 5]);
                    return [4 /*yield*/, mcpTokenService_1.default.listTokens({ pageIndex: 1, pageSize: 100 })];
                case 2:
                    response = _a.sent();
                    if (response.code === 0 && response.data) {
                        tokens.value = response.data.list || [];
                        total.value = response.data.total || 0;
                    }
                    else {
                        (0, notification_1.showError)(response.message || 'Failed to load tokens');
                    }
                    return [3 /*break*/, 5];
                case 3:
                    error_1 = _a.sent();
                    (0, notification_1.showError)(error_1.message || 'Failed to load tokens');
                    return [3 /*break*/, 5];
                case 4:
                    loading.value = false;
                    return [7 /*endfinally*/];
                case 5: return [2 /*return*/];
            }
        });
    });
}
function createToken() {
    return __awaiter(this, void 0, void 0, function () {
        var data, response, error_2;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    if (!newToken.value.name)
                        return [2 /*return*/];
                    creating.value = true;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 6, 7, 8]);
                    data = { name: newToken.value.name };
                    if (newToken.value.expiresAt) {
                        data.expiresAt = new Date(newToken.value.expiresAt).toISOString();
                    }
                    return [4 /*yield*/, mcpTokenService_1.default.createToken(data)];
                case 2:
                    response = _a.sent();
                    if (!(response.code === 0)) return [3 /*break*/, 4];
                    (0, notification_1.showSuccess)('Token created successfully');
                    newToken.value = { name: '', expiresAt: undefined };
                    return [4 /*yield*/, loadTokens()];
                case 3:
                    _a.sent();
                    return [3 /*break*/, 5];
                case 4:
                    (0, notification_1.showError)(response.message || 'Failed to create token');
                    _a.label = 5;
                case 5: return [3 /*break*/, 8];
                case 6:
                    error_2 = _a.sent();
                    (0, notification_1.showError)(error_2.message || 'Failed to create token');
                    return [3 /*break*/, 8];
                case 7:
                    creating.value = false;
                    return [7 /*endfinally*/];
                case 8: return [2 /*return*/];
            }
        });
    });
}
function deleteToken() {
    return __awaiter(this, void 0, void 0, function () {
        var response, error_3;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    if (pendingDelete.value === null)
                        return [2 /*return*/];
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 6, , 7]);
                    return [4 /*yield*/, mcpTokenService_1.default.deleteToken(pendingDelete.value)];
                case 2:
                    response = _a.sent();
                    if (!(response.code === 0)) return [3 /*break*/, 4];
                    (0, notification_1.showSuccess)('Token deleted');
                    pendingDelete.value = null;
                    return [4 /*yield*/, loadTokens()];
                case 3:
                    _a.sent();
                    return [3 /*break*/, 5];
                case 4:
                    (0, notification_1.showError)(response.message || 'Failed to delete token');
                    _a.label = 5;
                case 5: return [3 /*break*/, 7];
                case 6:
                    error_3 = _a.sent();
                    (0, notification_1.showError)(error_3.message || 'Failed to delete token');
                    return [3 /*break*/, 7];
                case 7: return [2 /*return*/];
            }
        });
    });
}
function deactivateToken(id) {
    return __awaiter(this, void 0, void 0, function () {
        var response, error_4;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    _a.trys.push([0, 5, , 6]);
                    return [4 /*yield*/, mcpTokenService_1.default.deactivateToken(id)];
                case 1:
                    response = _a.sent();
                    if (!(response.code === 0)) return [3 /*break*/, 3];
                    (0, notification_1.showSuccess)('Token deactivated');
                    return [4 /*yield*/, loadTokens()];
                case 2:
                    _a.sent();
                    return [3 /*break*/, 4];
                case 3:
                    (0, notification_1.showError)(response.message || 'Failed to deactivate token');
                    _a.label = 4;
                case 4: return [3 /*break*/, 6];
                case 5:
                    error_4 = _a.sent();
                    (0, notification_1.showError)(error_4.message || 'Failed to deactivate token');
                    return [3 /*break*/, 6];
                case 6: return [2 /*return*/];
            }
        });
    });
}
function activateToken(id) {
    return __awaiter(this, void 0, void 0, function () {
        var response, error_5;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    _a.trys.push([0, 5, , 6]);
                    return [4 /*yield*/, mcpTokenService_1.default.activateToken(id)];
                case 1:
                    response = _a.sent();
                    if (!(response.code === 0)) return [3 /*break*/, 3];
                    (0, notification_1.showSuccess)('Token activated');
                    return [4 /*yield*/, loadTokens()];
                case 2:
                    _a.sent();
                    return [3 /*break*/, 4];
                case 3:
                    (0, notification_1.showError)(response.message || 'Failed to activate token');
                    _a.label = 4;
                case 4: return [3 /*break*/, 6];
                case 5:
                    error_5 = _a.sent();
                    (0, notification_1.showError)(error_5.message || 'Failed to activate token');
                    return [3 /*break*/, 6];
                case 6: return [2 /*return*/];
            }
        });
    });
}
function maskToken(token) {
    if (token.length <= 8)
        return token;
    return token.substring(0, 4) + '...' + token.substring(token.length - 4);
}
function copyText(text) {
    return __awaiter(this, void 0, void 0, function () {
        var _a;
        return __generator(this, function (_b) {
            switch (_b.label) {
                case 0:
                    _b.trys.push([0, 2, , 3]);
                    return [4 /*yield*/, navigator.clipboard.writeText(text)];
                case 1:
                    _b.sent();
                    (0, notification_1.showSuccess)('Copied to clipboard');
                    return [3 /*break*/, 3];
                case 2:
                    _a = _b.sent();
                    (0, notification_1.showError)('Failed to copy');
                    return [3 /*break*/, 3];
                case 3: return [2 /*return*/];
            }
        });
    });
}
function copyToken(token) {
    return __awaiter(this, void 0, void 0, function () {
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, copyText(token)];
                case 1:
                    _a.sent();
                    return [2 /*return*/];
            }
        });
    });
}
function copySetTokenCmd(token) {
    return __awaiter(this, void 0, void 0, function () {
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, copyText("edu-cli config set-token ".concat(token))];
                case 1:
                    _a.sent();
                    (0, notification_1.showSuccess)('CLI command copied — paste it in your terminal');
                    return [2 /*return*/];
            }
        });
    });
}
function formatDate(dateString) {
    if (!dateString)
        return 'N/A';
    return new Date(dateString).toLocaleString();
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
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-sm text-gray-500" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.code, __VLS_intrinsics.code)(__assign({ class: "font-mono bg-gray-100 px-1 rounded" }));
/** @type {__VLS_StyleScopedClasses['font-mono']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-100']} */ ;
/** @type {__VLS_StyleScopedClasses['px-1']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-gray-50 border border-gray-200 rounded-lg p-5 mb-6" }));
/** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
/** @type {__VLS_StyleScopedClasses['p-5']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-base font-semibold text-gray-800 mb-3 flex items-center gap-2" }));
/** @type {__VLS_StyleScopedClasses['text-base']} */ ;
/** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-800']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-3']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "h-5 w-5 text-gray-500" }, { xmlns: "http://www.w3.org/2000/svg", fill: "none", viewBox: "0 0 24 24", stroke: "currentColor" }));
/** @type {__VLS_StyleScopedClasses['h-5']} */ ;
/** @type {__VLS_StyleScopedClasses['w-5']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.path)({
    'stroke-linecap': "round",
    'stroke-linejoin': "round",
    'stroke-width': "2",
    d: "M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z",
});
__VLS_asFunctionalElement1(__VLS_intrinsics.ol, __VLS_intrinsics.ol)(__assign({ class: "space-y-3 text-sm text-gray-700" }));
/** @type {__VLS_StyleScopedClasses['space-y-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.li, __VLS_intrinsics.li)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium" }));
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.code, __VLS_intrinsics.code)(__assign({ class: "font-mono bg-gray-200 px-1 rounded" }));
/** @type {__VLS_StyleScopedClasses['font-mono']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-200']} */ ;
/** @type {__VLS_StyleScopedClasses['px-1']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.a, __VLS_intrinsics.a)(__assign({ href: "https://github.com/liup215/ltedu/releases", target: "_blank", rel: "noopener noreferrer" }, { class: "text-indigo-600 hover:text-indigo-800 underline" }));
/** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-indigo-800']} */ ;
/** @type {__VLS_StyleScopedClasses['underline']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.li, __VLS_intrinsics.li)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium" }));
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-1 relative" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['relative']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.code, __VLS_intrinsics.code)(__assign({ class: "block bg-gray-800 text-green-300 text-xs font-mono px-4 py-2 rounded" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-800']} */ ;
/** @type {__VLS_StyleScopedClasses['text-green-300']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-mono']} */ ;
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
(__VLS_ctx.baseUrl);
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
        var _a = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            _a[_i] = arguments[_i];
        }
        var $event = _a[0];
        __VLS_ctx.copyText("edu-cli config set-url ".concat(__VLS_ctx.baseUrl));
        // @ts-ignore
        [baseUrl, baseUrl, copyText,];
    } }, { class: "absolute top-1 right-2 text-gray-400 hover:text-white text-xs px-1" }), { 'aria-label': "Copy set-url command", title: "Copy" }));
/** @type {__VLS_StyleScopedClasses['absolute']} */ ;
/** @type {__VLS_StyleScopedClasses['top-1']} */ ;
/** @type {__VLS_StyleScopedClasses['right-2']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-white']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['px-1']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.li, __VLS_intrinsics.li)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium" }));
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-1 relative" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['relative']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.code, __VLS_intrinsics.code)(__assign({ class: "block bg-gray-800 text-green-300 text-xs font-mono px-4 py-2 rounded" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-800']} */ ;
/** @type {__VLS_StyleScopedClasses['text-green-300']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-mono']} */ ;
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.li, __VLS_intrinsics.li)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium" }));
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-1 relative" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['relative']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.code, __VLS_intrinsics.code)(__assign({ class: "block bg-gray-800 text-green-300 text-xs font-mono px-4 py-2 rounded" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-800']} */ ;
/** @type {__VLS_StyleScopedClasses['text-green-300']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-mono']} */ ;
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-2']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-4 text-xs text-gray-500" }));
/** @type {__VLS_StyleScopedClasses['mt-4']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.code, __VLS_intrinsics.code)(__assign({ class: "font-mono bg-gray-200 px-1 rounded" }));
/** @type {__VLS_StyleScopedClasses['font-mono']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-200']} */ ;
/** @type {__VLS_StyleScopedClasses['px-1']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
(__VLS_ctx.baseUrl);
__VLS_asFunctionalElement1(__VLS_intrinsics.code, __VLS_intrinsics.code)(__assign({ class: "font-mono bg-gray-200 px-1 rounded" }));
/** @type {__VLS_StyleScopedClasses['font-mono']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-200']} */ ;
/** @type {__VLS_StyleScopedClasses['px-1']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white shadow sm:rounded-lg p-6 mb-6" }));
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:rounded-lg']} */ ;
/** @type {__VLS_StyleScopedClasses['p-6']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-lg font-medium text-gray-900 mb-4" }));
/** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
/** @type {__VLS_StyleScopedClasses['mb-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.form, __VLS_intrinsics.form)(__assign({ onSubmit: (__VLS_ctx.createToken) }, { class: "space-y-4" }));
/** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "tokenName" }, { class: "block text-sm font-medium text-gray-700" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ value: (__VLS_ctx.newToken.name), type: "text", id: "tokenName", required: true, placeholder: "e.g., My Laptop CLI" }, { class: "mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" }));
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
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ for: "expiresAt" }, { class: "block text-sm font-medium text-gray-700" }));
/** @type {__VLS_StyleScopedClasses['block']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "datetime-local", id: "expiresAt" }, { class: "mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" }));
(__VLS_ctx.newToken.expiresAt);
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
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:border-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-1 text-xs text-gray-500" }));
/** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ type: "submit", disabled: (__VLS_ctx.creating || !__VLS_ctx.newToken.name) }, { class: "inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed" }));
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
/** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
/** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
(__VLS_ctx.creating ? 'Creating...' : 'Create Token');
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white shadow overflow-hidden sm:rounded-lg" }));
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow']} */ ;
/** @type {__VLS_StyleScopedClasses['overflow-hidden']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:rounded-lg']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "px-4 py-5 sm:px-6 flex justify-between items-center" }));
/** @type {__VLS_StyleScopedClasses['px-4']} */ ;
/** @type {__VLS_StyleScopedClasses['py-5']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.h2, __VLS_intrinsics.h2)(__assign({ class: "text-lg font-medium text-gray-900" }));
/** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: (__VLS_ctx.loadTokens) }, { class: "inline-flex items-center px-3 py-1 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none" }));
/** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['px-3']} */ ;
/** @type {__VLS_StyleScopedClasses['py-1']} */ ;
/** @type {__VLS_StyleScopedClasses['border']} */ ;
/** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
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
__VLS_asFunctionalElement1(__VLS_intrinsics.tbody, __VLS_intrinsics.tbody)(__assign({ class: "bg-white divide-y divide-gray-200" }));
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
/** @type {__VLS_StyleScopedClasses['divide-gray-200']} */ ;
if (__VLS_ctx.loading) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ colspan: "6" }, { class: "px-6 py-4 text-center text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
}
else if (!__VLS_ctx.tokens || __VLS_ctx.tokens.length === 0) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ colspan: "6" }, { class: "px-6 py-4 text-center text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
}
var _loop_1 = function (token) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)(__assign({ key: (token.id) }, { class: "hover:bg-gray-50" }));
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    (token.name);
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-center space-x-2" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-x-2']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.code, __VLS_intrinsics.code)(__assign({ class: "bg-gray-100 px-2 py-1 rounded text-xs font-mono max-w-xs truncate" }));
    /** @type {__VLS_StyleScopedClasses['bg-gray-100']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-mono']} */ ;
    /** @type {__VLS_StyleScopedClasses['max-w-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['truncate']} */ ;
    (__VLS_ctx.showFull[token.id] ? token.token : __VLS_ctx.maskToken(token.token));
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            __VLS_ctx.showFull[token.id] = !__VLS_ctx.showFull[token.id];
            // @ts-ignore
            [baseUrl, createToken, newToken, newToken, newToken, creating, creating, loadTokens, loading, tokens, tokens, tokens, showFull, showFull, showFull, maskToken,];
        } }, { class: "text-indigo-600 hover:text-indigo-900 text-xs whitespace-nowrap" }));
    /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-indigo-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    (__VLS_ctx.showFull[token.id] ? 'Hide' : 'Show');
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            __VLS_ctx.copyToken(token.token);
            // @ts-ignore
            [showFull, copyToken,];
        } }, { class: "text-indigo-600 hover:text-indigo-900 text-xs whitespace-nowrap" }));
    /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-indigo-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            __VLS_ctx.copySetTokenCmd(token.token);
            // @ts-ignore
            [copySetTokenCmd,];
        } }, { class: "text-green-700 hover:text-green-900 text-xs whitespace-nowrap font-medium" }));
    /** @type {__VLS_StyleScopedClasses['text-green-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-green-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: (token.isActive ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800') }, { class: "px-2 inline-flex text-xs leading-5 font-semibold rounded-full" }));
    /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['leading-5']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
    (token.isActive ? 'Active' : 'Inactive');
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    (__VLS_ctx.formatDate(token.createdAt));
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    (__VLS_ctx.formatDate(token.expiresAt));
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-x-2']} */ ;
    if (token.isActive) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(token.isActive))
                    return;
                __VLS_ctx.deactivateToken(token.id);
                // @ts-ignore
                [formatDate, formatDate, deactivateToken,];
            } }, { class: "text-yellow-600 hover:text-yellow-900" }));
        /** @type {__VLS_StyleScopedClasses['text-yellow-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-yellow-900']} */ ;
    }
    else {
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!!(token.isActive))
                    return;
                __VLS_ctx.activateToken(token.id);
                // @ts-ignore
                [activateToken,];
            } }, { class: "text-green-600 hover:text-green-900" }));
        /** @type {__VLS_StyleScopedClasses['text-green-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-green-900']} */ ;
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            __VLS_ctx.pendingDelete = token.id;
            // @ts-ignore
            [pendingDelete,];
        } }, { class: "text-red-600 hover:text-red-900" }));
    /** @type {__VLS_StyleScopedClasses['text-red-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-red-900']} */ ;
    // @ts-ignore
    [];
};
for (var _i = 0, _a = __VLS_vFor((__VLS_ctx.tokens)); _i < _a.length; _i++) {
    var token = _a[_i][0];
    _loop_1(token);
}
if (__VLS_ctx.pendingDelete !== null) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "fixed z-10 inset-0 overflow-y-auto" }, { role: "dialog", 'aria-modal': "true" }));
    /** @type {__VLS_StyleScopedClasses['fixed']} */ ;
    /** @type {__VLS_StyleScopedClasses['z-10']} */ ;
    /** @type {__VLS_StyleScopedClasses['inset-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['overflow-y-auto']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-end']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['min-h-screen']} */ ;
    /** @type {__VLS_StyleScopedClasses['pt-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['pb-20']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:block']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:p-0']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.pendingDelete !== null))
                return;
            __VLS_ctx.pendingDelete = null;
            // @ts-ignore
            [pendingDelete, pendingDelete,];
        } }, { class: "fixed inset-0 bg-gray-500 bg-opacity-75" }));
    /** @type {__VLS_StyleScopedClasses['fixed']} */ ;
    /** @type {__VLS_StyleScopedClasses['inset-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-opacity-75']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "hidden sm:inline-block sm:align-middle sm:h-screen" }));
    /** @type {__VLS_StyleScopedClasses['hidden']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:inline-block']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:align-middle']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:h-screen']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full" }));
    /** @type {__VLS_StyleScopedClasses['inline-block']} */ ;
    /** @type {__VLS_StyleScopedClasses['align-bottom']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
    /** @type {__VLS_StyleScopedClasses['overflow-hidden']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-xl']} */ ;
    /** @type {__VLS_StyleScopedClasses['transform']} */ ;
    /** @type {__VLS_StyleScopedClasses['transition-all']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:my-8']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:align-middle']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:max-w-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:w-full']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4" }));
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['pt-5']} */ ;
    /** @type {__VLS_StyleScopedClasses['pb-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:p-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:pb-4']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "sm:flex sm:items-start" }));
    /** @type {__VLS_StyleScopedClasses['sm:flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:items-start']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10" }));
    /** @type {__VLS_StyleScopedClasses['mx-auto']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex-shrink-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-12']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-12']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-red-100']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:mx-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:h-10']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:w-10']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "h-6 w-6 text-red-600" }, { fill: "none", viewBox: "0 0 24 24", stroke: "currentColor" }));
    /** @type {__VLS_StyleScopedClasses['h-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-red-600']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
        'stroke-linecap': "round",
        'stroke-linejoin': "round",
        'stroke-width': "2",
        d: "M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z",
    });
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left" }));
    /** @type {__VLS_StyleScopedClasses['mt-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:mt-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:ml-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:text-left']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.h3, __VLS_intrinsics.h3)(__assign({ class: "text-lg leading-6 font-medium text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['text-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['leading-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "mt-2 text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse" }));
    /** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:flex-row-reverse']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: (__VLS_ctx.deleteToken) }, { type: "button" }), { class: "w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 sm:ml-3 sm:w-auto sm:text-sm" }));
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-transparent']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-red-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-base']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-red-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:ml-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.pendingDelete !== null))
                return;
            __VLS_ctx.pendingDelete = null;
            // @ts-ignore
            [pendingDelete, deleteToken,];
        } }, { type: "button" }), { class: "mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm" }));
    /** @type {__VLS_StyleScopedClasses['mt-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
    /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-base']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:mt-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:ml-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:text-sm']} */ ;
}
// @ts-ignore
[];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
