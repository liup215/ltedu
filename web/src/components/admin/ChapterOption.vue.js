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
/**
 * ChapterOption Component
 *
 * A recursive component that renders a chapter tree structure with checkboxes
 * for multiple selection. Supports unlimited nesting levels with proper tree
 * visualization using ASCII characters.
 *
 * Features:
 * - Tree visualization with proper branch lines and indentation
 * - Multiple chapter selection with checkboxes
 * - Two-way binding for selected chapters
 * - Recursive rendering of nested chapter structures
 * - Proper handling of last items in each branch
 *
 * Usage:
 * <ChapterOption
 *   :chapter="chapterData"
 *   :level="0"
 *   :is-last="false"
 *   :selected-chapters="selectedIds"
 *   @update:selected="updateSelection"
 * />
 *
 * The component emits 'update:selected' events when chapters are selected/deselected,
 * providing an array of selected chapter IDs.
 *
 * Tree visualization example:
 * Chapter 1
 * ├── Section 1.1
 * │   ├── Topic 1.1.1
 * │   └── Topic 1.1.2
 * └── Section 1.2
 *     └── Topic 1.2.1
 */
var vue_1 = require("vue");
var props = withDefaults(defineProps(), {
    level: 0,
    isLast: false,
    parentTree: function () { return []; },
    selectedChapters: function () { return []; }
});
var emit = defineEmits();
/**
 * Computed value to check if current chapter is selected
 * A chapter is considered selected if:
 * - It is directly selected OR
 * - All its children are selected (if it has children)
 */
var isSelected = (0, vue_1.computed)(function () {
    var _a, _b, _c;
    if (!((_a = props.chapter.children) === null || _a === void 0 ? void 0 : _a.length)) {
        return (_c = (_b = props.selectedChapters) === null || _b === void 0 ? void 0 : _b.includes(props.chapter.id)) !== null && _c !== void 0 ? _c : false;
    }
    var descendantIds = getAllDescendantIds(props.chapter);
    return descendantIds.every(function (id) { var _a; return (_a = props.selectedChapters) === null || _a === void 0 ? void 0 : _a.includes(id); });
});
/**
 * Computed value to track indeterminate state
 * True when some but not all descendants are selected
 */
var isIndeterminate = (0, vue_1.computed)(function () {
    var _a;
    if (!((_a = props.chapter.children) === null || _a === void 0 ? void 0 : _a.length))
        return false;
    var descendantIds = getAllDescendantIds(props.chapter);
    var selectedCount = descendantIds.filter(function (id) { var _a; return (_a = props.selectedChapters) === null || _a === void 0 ? void 0 : _a.includes(id); }).length;
    return selectedCount > 0 && selectedCount < descendantIds.length;
});
/**
 * Get all descendant chapter IDs including the current chapter
 */
var getAllDescendantIds = function (chapter) {
    var ids = [chapter.id];
    if (chapter.children) {
        chapter.children.forEach(function (child) {
            ids.push.apply(ids, getAllDescendantIds(child));
        });
    }
    return ids;
};
/**
 * Toggle chapter selection and emit changes
 * When a chapter is selected/deselected, all its descendants are also selected/deselected
 */
var toggleChapter = function () {
    var newSelection = __spreadArray([], (props.selectedChapters || []), true);
    var currentSelected = isSelected.value;
    var allIds = getAllDescendantIds(props.chapter);
    if (!currentSelected) {
        // Select this chapter and all its descendants
        allIds.forEach(function (id) {
            if (!newSelection.includes(id)) {
                newSelection.push(id);
            }
        });
    }
    else {
        // Deselect this chapter and all its descendants
        allIds.forEach(function (id) {
            var index = newSelection.indexOf(id);
            if (index !== -1) {
                newSelection.splice(index, 1);
            }
        });
    }
    emit('update:selected', newSelection);
};
/**
 * Handle selection updates from child components
 */
var updateSelected = function (chapters) {
    emit('update:selected', chapters);
};
/**
 * Gets the count of selected descendants for the current chapter
 * Used to display selection count in indeterminate state
 */
var getSelectedCount = function () {
    var descendantIds = getAllDescendantIds(props.chapter);
    return descendantIds.filter(function (id) { var _a; return (_a = props.selectedChapters) === null || _a === void 0 ? void 0 : _a.includes(id); }).length;
};
/**
 * Computed tree context for the current node
 * Manages the state of vertical lines for the current level and its children
 */
var treeContext = (0, vue_1.computed)(function () {
    // Use inherited parent tree or create empty one
    var currentTree = __spreadArray([], props.parentTree, true);
    // Add current level's line state based on isLast
    if (props.level > 0) {
        currentTree.push(!props.isLast);
    }
    return {
        isLast: props.isLast,
        level: props.level,
        parentTree: currentTree
    };
});
// Compute the indentation based on parent context
var indent = (0, vue_1.computed)(function () {
    if (props.level <= 0)
        return '';
    // Create indentation based on parent levels only
    var parentLevels = props.parentTree;
    var indentParts = parentLevels.map(function (hasLine) { return hasLine ? '│' : ''; });
    return indentParts.join('');
});
// Compute the prefix for current node
var prefix = (0, vue_1.computed)(function () {
    if (props.level <= 0)
        return '';
    return props.isLast ? '└──' : '├──';
});
// Method to update parent tree for child nodes
var getChildParentTree = function (index, total) {
    // Get the parent tree state up to the current level
    var childTree = __spreadArray([], treeContext.value.parentTree, true);
    // Add information about siblings for this new level
    // When this child isn't the last one, add a vertical line for its children
    var hasNextSibling = index < total - 1;
    childTree.push(hasNextSibling);
    return childTree;
};
var __VLS_defaults = {
    level: 0,
    isLast: false,
    parentTree: function () { return []; },
    selectedChapters: function () { return []; }
};
var __VLS_ctx = __assign(__assign(__assign(__assign(__assign({}, {}), {}), {}), {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex items-center text-sm text-gray-900" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-mono whitespace-pre text-gray-500" }));
/** @type {__VLS_StyleScopedClasses['font-mono']} */ ;
/** @type {__VLS_StyleScopedClasses['whitespace-pre']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
(__VLS_ctx.indent);
(__VLS_ctx.prefix);
__VLS_asFunctionalElement1(__VLS_intrinsics.label, __VLS_intrinsics.label)(__assign({ class: "flex items-center gap-2 hover:bg-gray-50 px-2 py-1 rounded cursor-pointer group" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['gap-2']} */ ;
/** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
/** @type {__VLS_StyleScopedClasses['px-2']} */ ;
/** @type {__VLS_StyleScopedClasses['py-1']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['cursor-pointer']} */ ;
/** @type {__VLS_StyleScopedClasses['group']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign(__assign({ onChange: (__VLS_ctx.toggleChapter) }, { type: "checkbox", checked: (__VLS_ctx.isSelected), indeterminate: (__VLS_ctx.isIndeterminate) }), { class: "rounded text-indigo-600 focus:ring-indigo-500 group-hover:border-indigo-500 transition-colors" }));
/** @type {__VLS_StyleScopedClasses['rounded']} */ ;
/** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
/** @type {__VLS_StyleScopedClasses['focus:ring-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['group-hover:border-indigo-500']} */ ;
/** @type {__VLS_StyleScopedClasses['transition-colors']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: ([
        'truncate',
        __VLS_ctx.isIndeterminate ? 'text-gray-600' : '',
        'group-hover:text-gray-900 transition-colors'
    ]) }));
/** @type {__VLS_StyleScopedClasses['truncate']} */ ;
/** @type {__VLS_StyleScopedClasses['group-hover:text-gray-900']} */ ;
/** @type {__VLS_StyleScopedClasses['transition-colors']} */ ;
(__VLS_ctx.chapter.name);
if (__VLS_ctx.isIndeterminate) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-xs text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    (__VLS_ctx.$t('chapterOption.selectedCount', { count: __VLS_ctx.getSelectedCount() }));
}
if (__VLS_ctx.chapter.children && __VLS_ctx.chapter.children.length > 0) {
    for (var _i = 0, _a = __VLS_vFor((__VLS_ctx.chapter.children)); _i < _a.length; _i++) {
        var _b = _a[_i], child = _b[0], index = _b[1];
        __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ key: (child.id) }, { class: "ml-6" }));
        /** @type {__VLS_StyleScopedClasses['ml-6']} */ ;
        var __VLS_0 = void 0;
        /** @ts-ignore @type {typeof __VLS_components.ChapterOption} */
        ChapterOption;
        // @ts-ignore
        var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0(__assign({ 'onUpdate:selected': {} }, { chapter: (child), level: (props.level + 1), isLast: (index === __VLS_ctx.chapter.children.length - 1), parentTree: (__VLS_ctx.getChildParentTree(index, __VLS_ctx.chapter.children.length)), selectedChapters: (__VLS_ctx.selectedChapters) })));
        var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([__assign({ 'onUpdate:selected': {} }, { chapter: (child), level: (props.level + 1), isLast: (index === __VLS_ctx.chapter.children.length - 1), parentTree: (__VLS_ctx.getChildParentTree(index, __VLS_ctx.chapter.children.length)), selectedChapters: (__VLS_ctx.selectedChapters) })], __VLS_functionalComponentArgsRest(__VLS_1), false));
        var __VLS_5 = void 0;
        var __VLS_6 = ({ 'update:selected': {} },
            { 'onUpdate:selected': (__VLS_ctx.updateSelected) });
        var __VLS_3;
        var __VLS_4;
        // @ts-ignore
        [indent, prefix, toggleChapter, isSelected, isIndeterminate, isIndeterminate, isIndeterminate, chapter, chapter, chapter, chapter, chapter, chapter, $t, getSelectedCount, getChildParentTree, selectedChapters, updateSelected,];
    }
}
// @ts-ignore
[];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({
    __typeEmits: {},
    __defaults: __VLS_defaults,
    __typeProps: {},
});
exports.default = {};
