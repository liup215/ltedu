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
var quill_1 = require("quill");
var quill_table_better_1 = require("quill-table-better");
var editor_1 = require("../../const/editor");
require("quill/dist/quill.snow.css");
require("quill-table-better/dist/quill-table-better.css");
var apiClient_1 = require("../../services/apiClient");
var notification_1 = require("../../utils/notification");
// Register quill-table-better module
quill_1.default.register({
    'modules/table-better': quill_table_better_1.default
}, true);
var props = withDefaults(defineProps(), {
    height: editor_1.DEFAULT_EDITOR_HEIGHT,
    minHeight: editor_1.DEFAULT_EDITOR_MIN_HEIGHT,
    placeholder: editor_1.DEFAULT_PLACEHOLDER,
    readOnly: false,
    tableWidth: 'auto'
});
var emit = defineEmits();
var editorEl = (0, vue_1.ref)();
var quill = (0, vue_1.shallowRef)();
// Store the paste handler reference for cleanup
var pasteHandlerRef = (0, vue_1.ref)(null);
// Image Handler
var imageHandler = function () {
    var input = document.createElement('input');
    input.setAttribute('type', 'file');
    input.setAttribute('accept', 'image/*');
    input.click();
    input.onchange = function () { return __awaiter(void 0, void 0, void 0, function () {
        var file, formData, client, response, attachment, url_1, q, error_1;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    file = input.files ? input.files[0] : null;
                    if (!file)
                        return [2 /*return*/];
                    formData = new FormData();
                    formData.append('file', file);
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 4, , 5]);
                    return [4 /*yield*/, (0, apiClient_1.default)()];
                case 2:
                    client = _a.sent();
                    return [4 /*yield*/, client.post('/api/v1/upload/image', formData, {
                            headers: {
                                'Content-Type': 'multipart/form-data'
                            }
                        })];
                case 3:
                    response = _a.sent();
                    attachment = response.data.data;
                    url_1 = attachment.path;
                    if (quill.value) {
                        q = (0, vue_1.toRaw)(quill.value);
                        // Focus the editor first to ensure internal selection state is valid
                        q.focus();
                        // Use setTimeout to allow the focus event to process and selection to update
                        setTimeout(function () {
                            if (!quill.value)
                                return;
                            // Re-acquire raw instance inside timeout
                            var q = (0, vue_1.toRaw)(quill.value);
                            var index = 0;
                            try {
                                // getSelection can throw "reading 'offset'" error if DOM selection is invalid
                                // even after focus(). We must try-catch it.
                                var range = q.getSelection();
                                if (range) {
                                    index = range.index;
                                }
                                else {
                                    index = q.getLength();
                                }
                            }
                            catch (e) {
                                // Fallback to inserting at the end if selection retrieval fails
                                index = q.getLength();
                            }
                            q.insertEmbed(index, 'image', url_1);
                            q.setSelection(index + 1, 0);
                        }, 0);
                    }
                    return [3 /*break*/, 5];
                case 4:
                    error_1 = _a.sent();
                    console.error('Image upload failed', error_1);
                    // Show specific error if available to help debugging
                    (0, notification_1.showError)(error_1.message || 'Image upload failed');
                    return [3 /*break*/, 5];
                case 5: return [2 /*return*/];
            }
        });
    }); };
};
// Initialize Quill editor
(0, vue_1.onMounted)(function () {
    if (!editorEl.value)
        return;
    var option = {
        modules: {
            toolbar: false, // Disable toolbar for read-only mode
            table: false // Disable default table module
        },
        theme: 'snow',
        readOnly: props.readOnly,
        placeholder: props.placeholder,
        bounds: 'self'
    };
    // Only configure table-better and keyboard bindings for editable mode
    if (!props.readOnly) {
        option.modules['table-better'] = {
            language: 'en_US',
            menus: ['column', 'row', 'merge', 'table', 'cell', 'wrap', 'delete'],
            toolbarTable: true // Let the module handle toolbar integration
        };
        option.modules.keyboard = {
            bindings: quill_table_better_1.default.keyboardBindings
        };
        option.modules.toolbar = {
            container: [
                ['bold', 'italic', 'underline'],
                [{ 'list': 'ordered' }, { 'list': 'bullet' }],
                [{ 'script': 'sub' }, { 'script': 'super' }],
                [{ 'indent': '-1' }, { 'indent': '+1' }],
                [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
                [{ 'color': [] }, { 'background': [] }],
                [{ 'align': [] }],
                ['link', 'image', 'formula'],
                ['table-better'] // Built-in table button from quill-table-better
            ],
            handlers: {
                image: imageHandler
            }
        };
    }
    quill.value = new quill_1.default(editorEl.value, option);
    // Handle paste event to intercept image paste and upload to server
    var handlePaste = function (e) { return __awaiter(void 0, void 0, void 0, function () {
        var clipboardData, items, i, item, file, q, formData, client, response, url, range, index, error_2;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    clipboardData = e.clipboardData;
                    if (!clipboardData)
                        return [2 /*return*/];
                    items = clipboardData.items;
                    i = 0;
                    _a.label = 1;
                case 1:
                    if (!(i < items.length)) return [3 /*break*/, 8];
                    item = items[i];
                    if (!item.type.startsWith('image/')) return [3 /*break*/, 7];
                    e.preventDefault(); // Prevent default base64 insertion
                    file = item.getAsFile();
                    if (!(file && quill.value)) return [3 /*break*/, 6];
                    q = (0, vue_1.toRaw)(quill.value);
                    _a.label = 2;
                case 2:
                    _a.trys.push([2, 5, , 6]);
                    formData = new FormData();
                    formData.append('file', file, 'pasted-image.png');
                    return [4 /*yield*/, (0, apiClient_1.default)()];
                case 3:
                    client = _a.sent();
                    return [4 /*yield*/, client.post('/api/v1/upload/image', formData, {
                            headers: {
                                'Content-Type': 'multipart/form-data'
                            }
                        })];
                case 4:
                    response = _a.sent();
                    url = response.data.data.path;
                    // Insert the server URL as image
                    q.focus();
                    range = q.getSelection();
                    index = range ? range.index : q.getLength();
                    q.insertEmbed(index, 'image', url);
                    q.setSelection(index + 1, 0);
                    return [3 /*break*/, 6];
                case 5:
                    error_2 = _a.sent();
                    console.error('Pasted image upload failed', error_2);
                    (0, notification_1.showError)(error_2.message || 'Image upload failed');
                    return [3 /*break*/, 6];
                case 6: return [3 /*break*/, 8]; // Only handle the first image
                case 7:
                    i++;
                    return [3 /*break*/, 1];
                case 8: return [2 /*return*/];
            }
        });
    }); };
    // Add paste event listener with capture phase to intercept before Quill
    pasteHandlerRef.value = handlePaste;
    quill.value.root.addEventListener('paste', handlePaste, true);
    // Set initial content
    quill.value.root.innerHTML = props.modelValue;
    // Apply table width class
    updateTableWidthClass();
    // Listen for content changes
    quill.value.on(editor_1.CONTENT_CHANGE_EVENT, function () {
        var _a;
        var html = ((_a = quill.value) === null || _a === void 0 ? void 0 : _a.root.innerHTML) || '';
        if (html !== props.modelValue) {
            emit('update:modelValue', html);
            emit('change', html);
        }
    });
});
// Update table width class on editor root
var updateTableWidthClass = function () {
    if (!quill.value)
        return;
    var root = quill.value.root;
    // Remove existing table width classes
    root.classList.remove('table-width-auto', 'table-width-fixed', 'table-width-full');
    // Add the new class
    root.classList.add("table-width-".concat(props.tableWidth));
};
// Watch for external changes to modelValue
(0, vue_1.watch)(function () { return props.modelValue; }, function (newValue) {
    if (quill.value && quill.value.root.innerHTML !== newValue) {
        quill.value.root.innerHTML = newValue;
    }
});
// Watch for tableWidth prop changes
(0, vue_1.watch)(function () { return props.tableWidth; }, function () {
    updateTableWidthClass();
});
// Clean up when component is destroyed
(0, vue_1.onBeforeUnmount)(function () {
    var _a;
    (_a = quill.value) === null || _a === void 0 ? void 0 : _a.off('text-change');
    // Remove paste event listener
    if (pasteHandlerRef.value && quill.value) {
        quill.value.root.removeEventListener('paste', pasteHandlerRef.value, true);
    }
});
var __VLS_defaults = {
    height: editor_1.DEFAULT_EDITOR_HEIGHT,
    minHeight: editor_1.DEFAULT_EDITOR_MIN_HEIGHT,
    placeholder: editor_1.DEFAULT_PLACEHOLDER,
    readOnly: false,
    tableWidth: 'auto'
};
var __VLS_ctx = __assign(__assign(__assign(__assign({}, {}), {}), {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
/** @type {__VLS_StyleScopedClasses['ql-editor']} */ ;
/** @type {__VLS_StyleScopedClasses['ql-editor']} */ ;
(__VLS_ctx.minHeight);
(__VLS_ctx.height);
// @ts-ignore
[minHeight, height,];
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "quill-editor-container" }));
/** @type {__VLS_StyleScopedClasses['quill-editor-container']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div)(__assign({ ref: "editorEl" }, { style: ({ height: __VLS_ctx.height, minHeight: __VLS_ctx.minHeight }) }));
// @ts-ignore
[height, minHeight,];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({
    __typeEmits: {},
    __defaults: __VLS_defaults,
    __typeProps: {},
});
exports.default = {};
