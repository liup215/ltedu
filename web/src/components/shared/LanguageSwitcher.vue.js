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
Object.defineProperty(exports, "__esModule", { value: true });
var vue_1 = require("vue");
var vue_i18n_1 = require("vue-i18n");
var locale = (0, vue_i18n_1.useI18n)().locale;
var languages = [
    { value: 'en', label: 'English' },
    { value: 'zh', label: '中文' },
];
var currentLocale = (0, vue_1.ref)(locale.value);
var dropdownOpen = (0, vue_1.ref)(false);
var currentLabel = (0, vue_1.computed)(function () {
    var found = languages.find(function (l) { return l.value === currentLocale.value; });
    return found ? found.label : '';
});
function toggleDropdown() {
    dropdownOpen.value = !dropdownOpen.value;
}
function selectLanguage(lang) {
    currentLocale.value = lang;
    locale.value = lang;
    localStorage.setItem('locale', lang);
    dropdownOpen.value = false;
    location.reload();
}
var __VLS_ctx = __assign(__assign({}, {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "relative" }));
/** @type {__VLS_StyleScopedClasses['relative']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ onClick: (__VLS_ctx.toggleDropdown) }, { class: "flex items-center cursor-pointer select-none text-gray-700 text-sm font-medium hover:text-indigo-600" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['cursor-pointer']} */ ;
/** @type {__VLS_StyleScopedClasses['select-none']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:text-indigo-600']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "mr-1" }));
/** @type {__VLS_StyleScopedClasses['mr-1']} */ ;
(__VLS_ctx.currentLabel);
__VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-4 h-4" }, { fill: "none", stroke: "currentColor", viewBox: "0 0 24 24" }));
/** @type {__VLS_StyleScopedClasses['w-4']} */ ;
/** @type {__VLS_StyleScopedClasses['h-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.path)({
    'stroke-linecap': "round",
    'stroke-linejoin': "round",
    'stroke-width': "2",
    d: "M19 9l-7 7-7-7",
});
if (__VLS_ctx.dropdownOpen) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.dropdownOpen))
                return;
            __VLS_ctx.dropdownOpen = false;
            // @ts-ignore
            [toggleDropdown, currentLabel, dropdownOpen, dropdownOpen,];
        } }, { class: "absolute right-0 mt-2 w-32 bg-white rounded-md shadow-lg py-1 z-50 border border-gray-200" }));
    /** @type {__VLS_StyleScopedClasses['absolute']} */ ;
    /** @type {__VLS_StyleScopedClasses['right-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['mt-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['w-32']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
    /** @type {__VLS_StyleScopedClasses['shadow-lg']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['z-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['border']} */ ;
    /** @type {__VLS_StyleScopedClasses['border-gray-200']} */ ;
    var _loop_1 = function (lang) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(__VLS_ctx.dropdownOpen))
                    return;
                __VLS_ctx.selectLanguage(lang.value);
                // @ts-ignore
                [languages, selectLanguage,];
            } }, { key: (lang.value) }), { class: "block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" }), { class: ({ 'font-bold text-indigo-600': __VLS_ctx.currentLocale === lang.value }) }));
        /** @type {__VLS_StyleScopedClasses['block']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-left']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-bold']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
        (lang.label);
        // @ts-ignore
        [currentLocale,];
    };
    for (var _i = 0, _a = __VLS_vFor((__VLS_ctx.languages)); _i < _a.length; _i++) {
        var lang = _a[_i][0];
        _loop_1(lang);
    }
}
// @ts-ignore
[];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
