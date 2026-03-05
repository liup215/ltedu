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
var NavBar_vue_1 = require("../NavBar.vue"); // Adjust path if NavBar is moved or aliased
var AdminSidebar_vue_1 = require("./AdminSidebar.vue");
var sidebarOpen = (0, vue_1.ref)(false);
var __VLS_ctx = __assign(__assign({}, {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex flex-col min-h-screen bg-gray-100" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
/** @type {__VLS_StyleScopedClasses['min-h-screen']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-gray-100']} */ ;
var __VLS_0 = NavBar_vue_1.default;
// @ts-ignore
var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0({}));
var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([{}], __VLS_functionalComponentArgsRest(__VLS_1), false));
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex flex-1 pt-16" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-1']} */ ;
/** @type {__VLS_StyleScopedClasses['pt-16']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
        var _a = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            _a[_i] = arguments[_i];
        }
        var $event = _a[0];
        __VLS_ctx.sidebarOpen = !__VLS_ctx.sidebarOpen;
        // @ts-ignore
        [sidebarOpen, sidebarOpen,];
    } }, { class: "md:hidden fixed bottom-4 right-4 z-50 w-14 h-14 bg-indigo-600 text-white rounded-full shadow-lg flex items-center justify-center focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500" }), { 'aria-label': "Toggle sidebar" }));
/** @type {__VLS_StyleScopedClasses['md:hidden']} */ ;
/** @type {__VLS_StyleScopedClasses['fixed']} */ ;
/** @type {__VLS_StyleScopedClasses['bottom-4']} */ ;
/** @type {__VLS_StyleScopedClasses['right-4']} */ ;
/** @type {__VLS_StyleScopedClasses['z-50']} */ ;
/** @type {__VLS_StyleScopedClasses['w-14']} */ ;
/** @type {__VLS_StyleScopedClasses['h-14']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-indigo-600']} */ ;
/** @type {__VLS_StyleScopedClasses['text-white']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-lg']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:outline-none']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-2']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-offset-2']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
if (!__VLS_ctx.sidebarOpen) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-6 h-6" }, { fill: "none", stroke: "currentColor", viewBox: "0 0 24 24" }));
    /** @type {__VLS_StyleScopedClasses['w-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
        'stroke-linecap': "round",
        'stroke-linejoin': "round",
        'stroke-width': "2",
        d: "M4 6h16M4 12h16M4 18h16",
    });
}
else {
    __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: "w-6 h-6" }, { fill: "none", stroke: "currentColor", viewBox: "0 0 24 24" }));
    /** @type {__VLS_StyleScopedClasses['w-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['h-6']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
        'stroke-linecap': "round",
        'stroke-linejoin': "round",
        'stroke-width': "2",
        d: "M6 18L18 6M6 6l12 12",
    });
}
if (__VLS_ctx.sidebarOpen) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            if (!(__VLS_ctx.sidebarOpen))
                return;
            __VLS_ctx.sidebarOpen = false;
            // @ts-ignore
            [sidebarOpen, sidebarOpen, sidebarOpen,];
        } }, { class: "md:hidden fixed inset-0 bg-gray-600 bg-opacity-50 z-30" }));
    /** @type {__VLS_StyleScopedClasses['md:hidden']} */ ;
    /** @type {__VLS_StyleScopedClasses['fixed']} */ ;
    /** @type {__VLS_StyleScopedClasses['inset-0']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-gray-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['bg-opacity-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['z-30']} */ ;
}
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: ([
        'fixed md:static inset-y-0 left-0 z-40 transform transition-transform duration-300 ease-in-out md:transform-none',
        __VLS_ctx.sidebarOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0',
        'pt-16 md:pt-0'
    ]) }));
/** @type {__VLS_StyleScopedClasses['fixed']} */ ;
/** @type {__VLS_StyleScopedClasses['md:static']} */ ;
/** @type {__VLS_StyleScopedClasses['inset-y-0']} */ ;
/** @type {__VLS_StyleScopedClasses['left-0']} */ ;
/** @type {__VLS_StyleScopedClasses['z-40']} */ ;
/** @type {__VLS_StyleScopedClasses['transform']} */ ;
/** @type {__VLS_StyleScopedClasses['transition-transform']} */ ;
/** @type {__VLS_StyleScopedClasses['duration-300']} */ ;
/** @type {__VLS_StyleScopedClasses['ease-in-out']} */ ;
/** @type {__VLS_StyleScopedClasses['md:transform-none']} */ ;
/** @type {__VLS_StyleScopedClasses['pt-16']} */ ;
/** @type {__VLS_StyleScopedClasses['md:pt-0']} */ ;
var __VLS_5 = AdminSidebar_vue_1.default;
// @ts-ignore
var __VLS_6 = __VLS_asFunctionalComponent1(__VLS_5, new __VLS_5(__assign({ 'onNavigate': {} })));
var __VLS_7 = __VLS_6.apply(void 0, __spreadArray([__assign({ 'onNavigate': {} })], __VLS_functionalComponentArgsRest(__VLS_6), false));
var __VLS_10;
var __VLS_11 = ({ navigate: {} },
    { onNavigate: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            __VLS_ctx.sidebarOpen = false;
            // @ts-ignore
            [sidebarOpen, sidebarOpen,];
        } });
var __VLS_8;
var __VLS_9;
__VLS_asFunctionalElement1(__VLS_intrinsics.main, __VLS_intrinsics.main)(__assign({ class: "flex-1 p-4 md:p-6 min-w-0" }));
/** @type {__VLS_StyleScopedClasses['flex-1']} */ ;
/** @type {__VLS_StyleScopedClasses['p-4']} */ ;
/** @type {__VLS_StyleScopedClasses['md:p-6']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-0']} */ ;
var __VLS_12;
/** @ts-ignore @type {typeof __VLS_components.routerView | typeof __VLS_components.RouterView} */
routerView;
// @ts-ignore
var __VLS_13 = __VLS_asFunctionalComponent1(__VLS_12, new __VLS_12({}));
var __VLS_14 = __VLS_13.apply(void 0, __spreadArray([{}], __VLS_functionalComponentArgsRest(__VLS_13), false));
// @ts-ignore
[];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
