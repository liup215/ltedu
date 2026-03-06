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
var _a, _b;
Object.defineProperty(exports, "__esModule", { value: true });
var __VLS_props = defineProps();
var emit = defineEmits();
var onSelect = function (id) {
    emit('select', id);
};
var onToggleExpand = function (id) {
    emit('toggle-expand', id);
};
var __VLS_ctx = __assign(__assign(__assign(__assign(__assign({}, {}), {}), {}), {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "chapter-tree" }));
/** @type {__VLS_StyleScopedClasses['chapter-tree']} */ ;
if (!__VLS_ctx.chapters || __VLS_ctx.chapters.length === 0) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "text-gray-400 text-sm px-3 py-2" }));
    /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    (__VLS_ctx.$t('chapterTree.noChapters'));
}
var _loop_1 = function (chapter) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (chapter.id) }, { class: "chapter-item" }));
    /** @type {__VLS_StyleScopedClasses['chapter-item']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            __VLS_ctx.onSelect(chapter.id);
            // @ts-ignore
            [chapters, chapters, chapters, $t, onSelect,];
        } }, { class: ([
            'chapter-row flex items-center py-2 px-3 hover:bg-gray-50 cursor-pointer rounded',
            __VLS_ctx.selectedId === chapter.id ? 'bg-indigo-50 text-indigo-700' : 'text-gray-700'
        ]) }));
    /** @type {__VLS_StyleScopedClasses['chapter-row']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['cursor-pointer']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    if ((_a = chapter.children) === null || _a === void 0 ? void 0 : _a.length) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign({ onClick: function () {
                var _a;
                var _b = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _b[_i] = arguments[_i];
                }
                var $event = _b[0];
                if (!((_a = chapter.children) === null || _a === void 0 ? void 0 : _a.length))
                    return;
                __VLS_ctx.onToggleExpand(chapter.id);
                // @ts-ignore
                [selectedId, onToggleExpand,];
            } }, { class: "mr-2 w-4 h-4 flex items-center justify-center text-gray-400 hover:text-gray-600" }));
        /** @type {__VLS_StyleScopedClasses['mr-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['w-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['h-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['justify-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-gray-600']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.svg, __VLS_intrinsics.svg)(__assign({ class: ([
                'w-4 h-4 transform transition-transform',
                __VLS_ctx.expandedChapters.includes(chapter.id) ? 'rotate-90' : ''
            ]) }, { xmlns: "http://www.w3.org/2000/svg", viewBox: "0 0 20 20", fill: "currentColor" }));
        /** @type {__VLS_StyleScopedClasses['w-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['h-4']} */ ;
        /** @type {__VLS_StyleScopedClasses['transform']} */ ;
        /** @type {__VLS_StyleScopedClasses['transition-transform']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.path)({
            'fill-rule': "evenodd",
            d: "M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z",
            'clip-rule': "evenodd",
        });
    }
    else {
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "w-6" }));
        /** @type {__VLS_StyleScopedClasses['w-6']} */ ;
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "truncate" }));
    /** @type {__VLS_StyleScopedClasses['truncate']} */ ;
    (chapter.name);
    if (((_b = chapter.children) === null || _b === void 0 ? void 0 : _b.length) && __VLS_ctx.expandedChapters.includes(chapter.id)) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "ml-6 mt-1" }));
        /** @type {__VLS_StyleScopedClasses['ml-6']} */ ;
        /** @type {__VLS_StyleScopedClasses['mt-1']} */ ;
        var __VLS_0 = void 0;
        /** @ts-ignore @type {typeof __VLS_components.ChapterTree} */
        ChapterTree;
        // @ts-ignore
        var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0(__assign(__assign({ 'onSelect': {} }, { 'onToggleExpand': {} }), { chapters: (chapter.children), selectedId: (__VLS_ctx.selectedId), expandedChapters: (__VLS_ctx.expandedChapters) })));
        var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([__assign(__assign({ 'onSelect': {} }, { 'onToggleExpand': {} }), { chapters: (chapter.children), selectedId: (__VLS_ctx.selectedId), expandedChapters: (__VLS_ctx.expandedChapters) })], __VLS_functionalComponentArgsRest(__VLS_1), false));
        var __VLS_5 = void 0;
        var __VLS_6 = ({ select: {} },
            { onSelect: (function (id) { return __VLS_ctx.$emit('select', id); }) });
        var __VLS_7 = ({ toggleExpand: {} },
            { onToggleExpand: (function (id) { return __VLS_ctx.$emit('toggle-expand', id); }) });
    }
    // @ts-ignore
    [selectedId, expandedChapters, expandedChapters, expandedChapters, $emit, $emit,];
};
var __VLS_3, __VLS_4;
for (var _i = 0, _c = __VLS_vFor((__VLS_ctx.chapters)); _i < _c.length; _i++) {
    var chapter = _c[_i][0];
    _loop_1(chapter);
}
// @ts-ignore
[];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({
    __typeEmits: {},
    __typeProps: {},
});
exports.default = {};
