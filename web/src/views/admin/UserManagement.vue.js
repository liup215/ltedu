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
var _a, _b, _c, _d;
Object.defineProperty(exports, "__esModule", { value: true });
var vue_1 = require("vue");
var userService_1 = require("../../services/userService");
var users = (0, vue_1.ref)([]);
var loading = (0, vue_1.ref)(true);
var totalUsers = (0, vue_1.ref)(0);
var currentPage = (0, vue_1.ref)(1);
var pageSize = 10;
var searchQuery = (0, vue_1.ref)('');
var statusFilter = (0, vue_1.ref)(0);
var fetchUsers = function () { return __awaiter(void 0, void 0, void 0, function () {
    var query, response, error_1;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                loading.value = true;
                _a.label = 1;
            case 1:
                _a.trys.push([1, 3, 4, 5]);
                query = {
                    pageIndex: currentPage.value,
                    pageSize: pageSize,
                };
                if (searchQuery.value.trim() !== '') {
                    query.username = searchQuery.value.trim();
                }
                if (statusFilter.value !== undefined) {
                    query.status = statusFilter.value;
                }
                return [4 /*yield*/, userService_1.default.getUsers(query)];
            case 2:
                response = _a.sent();
                console.log('Fetched users:', response);
                users.value = response.data.list;
                totalUsers.value = response.data.total;
                return [3 /*break*/, 5];
            case 3:
                error_1 = _a.sent();
                console.error('Failed to fetch users:', error_1);
                users.value = [];
                totalUsers.value = 0;
                return [3 /*break*/, 5];
            case 4:
                loading.value = false;
                return [7 /*endfinally*/];
            case 5: return [2 /*return*/];
        }
    });
}); };
(0, vue_1.onMounted)(function () {
    fetchUsers();
});
// const openAddUserModal = () => {
//   console.log('Open add user modal/navigate to create user page (Placeholder)');
//   // Example: router.push('/admin/users/create');
// };
// Edit user function removed as per request
// const editUser = (userId: number | null) => {
//   if (userId === null) {
//     console.warn('Edit action on user with null ID');
//     return;
//   }
//   console.log('Edit user (placeholder):', userId);
//   // Example: router.push(`/admin/users/edit/${userId}`);
// };
// Delete user function removed as per request
// const deleteUser = async (userId: number | null) => {
//   if (userId === null) {
//     console.warn('Delete action on user with null ID');
//     return;
//   }
//   if (confirm(`Are you sure you want to delete user with ID ${userId}? This action cannot be undone.`)) {
//     try {
//       loading.value = true;
//       await userService.deleteUser(userId);
//       if (users.value.length === 1 && currentPage.value > 1 && totalUsers.value > 1) { // check totalUsers to prevent going to page 0
//         currentPage.value--;
//       }
//       await fetchUsers(); // Refresh list
//     } catch (error) {
//       console.error(`Failed to delete user ${userId}:`, error);
//       // TODO: Show error message to user
//     } finally {
//       loading.value = false;
//     }
//   }
// };
var setUserStatus = function (userId, newStatus) { return __awaiter(void 0, void 0, void 0, function () {
    var actionText, error_2;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (userId === null) {
                    console.warn('Set status action on user with null ID');
                    // TODO: Show error message to user
                    return [2 /*return*/];
                }
                actionText = newStatus === 1 ? 'activate' : 'deactivate';
                if (!confirm("Are you sure you want to ".concat(actionText, " user with ID ").concat(userId, "?"))) return [3 /*break*/, 6];
                _a.label = 1;
            case 1:
                _a.trys.push([1, 4, 5, 6]);
                loading.value = true; // Indicate an operation is in progress
                return [4 /*yield*/, userService_1.default.updateUser(userId, { status: newStatus })];
            case 2:
                _a.sent();
                // Optionally, find the user in the list and update their status locally
                // for a more responsive UI before refetching, or just refetch.
                // const userIndex = users.value.findIndex(u => u.id === userId);
                // if (userIndex !== -1) {
                //   users.value[userIndex].status = newStatus;
                // }
                return [4 /*yield*/, fetchUsers()];
            case 3:
                // Optionally, find the user in the list and update their status locally
                // for a more responsive UI before refetching, or just refetch.
                // const userIndex = users.value.findIndex(u => u.id === userId);
                // if (userIndex !== -1) {
                //   users.value[userIndex].status = newStatus;
                // }
                _a.sent(); // Refresh the entire list to ensure data consistency
                return [3 /*break*/, 6];
            case 4:
                error_2 = _a.sent();
                console.error("Failed to ".concat(actionText, " user ").concat(userId, ":"), error_2);
                return [3 /*break*/, 6];
            case 5:
                loading.value = false;
                return [7 /*endfinally*/];
            case 6: return [2 /*return*/];
        }
    });
}); };
var formatDate = function (dateValue) {
    if (!dateValue)
        return 'N/A';
    try {
        var dateObj = typeof dateValue === 'number' ? new Date(dateValue) : new Date(dateValue);
        return dateObj.toLocaleDateString();
    }
    catch (e) {
        console.error('Invalid date value for formatting:', dateValue, e);
        return 'Invalid Date';
    }
};
var getUserStatusInfo = function (statusValue) {
    switch (statusValue) {
        case 1:
            return { text: 'Active', class: 'bg-green-100 text-green-800' };
        case 2:
            return { text: 'Inactive', class: 'bg-yellow-100 text-yellow-800' };
        default:
            return { text: 'Unknown', class: 'bg-gray-100 text-gray-800' };
    }
};
var totalPages = (0, vue_1.computed)(function () {
    if (totalUsers.value === 0)
        return 0; // Return 0 if no users, to prevent NaN or Infinity
    return Math.ceil(totalUsers.value / pageSize);
});
var goToPage = function (page) {
    if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
        currentPage.value = page;
        fetchUsers();
    }
};
var paginationRange = (0, vue_1.computed)(function () {
    var range = [];
    var maxPagesToShow = 5;
    var start = Math.max(1, currentPage.value - Math.floor(maxPagesToShow / 2));
    var end = Math.min(totalPages.value, start + maxPagesToShow - 1);
    if (totalPages.value > 0 && end - start + 1 < maxPagesToShow) { // Ensure totalPages > 0
        if (currentPage.value <= Math.floor(maxPagesToShow / 2)) { // If near the beginning
            end = Math.min(totalPages.value, maxPagesToShow);
        }
        else { // If near the end
            start = Math.max(1, totalPages.value - maxPagesToShow + 1);
        }
    }
    // Recalculate end if start was adjusted
    if (totalPages.value > 0 && end - start + 1 < maxPagesToShow && start === 1 && totalPages.value < maxPagesToShow) {
        end = totalPages.value;
    }
    for (var i = start; i <= end; i++) {
        if (i > 0)
            range.push(i); // Ensure only positive page numbers
    }
    return range;
});
var searchDebounceTimer;
(0, vue_1.watch)(searchQuery, function () {
    clearTimeout(searchDebounceTimer);
    searchDebounceTimer = window.setTimeout(function () {
        currentPage.value = 1;
        fetchUsers();
    }, 500);
});
(0, vue_1.watch)(statusFilter, function () {
    currentPage.value = 1;
    fetchUsers();
});
var setUserAsAdmin = function (userId) { return __awaiter(void 0, void 0, void 0, function () {
    var error_3;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (userId === null) {
                    console.warn('Set admin action on user with null ID');
                    return [2 /*return*/];
                }
                if (!confirm("Are you sure you want to set user with ID ".concat(userId, " as admin?"))) return [3 /*break*/, 6];
                _a.label = 1;
            case 1:
                _a.trys.push([1, 4, 5, 6]);
                loading.value = true;
                return [4 /*yield*/, userService_1.default.setAdmin(userId)];
            case 2:
                _a.sent();
                return [4 /*yield*/, fetchUsers()];
            case 3:
                _a.sent();
                return [3 /*break*/, 6];
            case 4:
                error_3 = _a.sent();
                console.error("Failed to set user ".concat(userId, " as admin:"), error_3);
                return [3 /*break*/, 6];
            case 5:
                loading.value = false;
                return [7 /*endfinally*/];
            case 6: return [2 /*return*/];
        }
    });
}); };
var removeAdmin = function (userId) { return __awaiter(void 0, void 0, void 0, function () {
    var error_4;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (userId === null) {
                    console.warn('Remove admin action on user with null ID');
                    return [2 /*return*/];
                }
                if (!confirm("Are you sure you want to remove admin privileges from user ID ".concat(userId, "?"))) return [3 /*break*/, 6];
                _a.label = 1;
            case 1:
                _a.trys.push([1, 4, 5, 6]);
                loading.value = true;
                return [4 /*yield*/, userService_1.default.removeAdmin(userId)];
            case 2:
                _a.sent();
                return [4 /*yield*/, fetchUsers()];
            case 3:
                _a.sent();
                return [3 /*break*/, 6];
            case 4:
                error_4 = _a.sent();
                console.error("Failed to remove admin from user ".concat(userId, ":"), error_4);
                return [3 /*break*/, 6];
            case 5:
                loading.value = false;
                return [7 /*endfinally*/];
            case 6: return [2 /*return*/];
        }
    });
}); };
var grantVip = function (userId) { return __awaiter(void 0, void 0, void 0, function () {
    var error_5;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (userId === null) {
                    console.warn('Grant VIP action on user with null ID');
                    return [2 /*return*/];
                }
                if (!confirm("Are you sure you want to grant 1 month VIP to user ID ".concat(userId, "?"))) return [3 /*break*/, 6];
                _a.label = 1;
            case 1:
                _a.trys.push([1, 4, 5, 6]);
                loading.value = true;
                return [4 /*yield*/, userService_1.default.grantVip(userId)];
            case 2:
                _a.sent();
                alert('VIP granted for 1 month!');
                return [4 /*yield*/, fetchUsers()];
            case 3:
                _a.sent();
                return [3 /*break*/, 6];
            case 4:
                error_5 = _a.sent();
                console.error("Failed to grant VIP to user ".concat(userId, ":"), error_5);
                alert('Failed to grant VIP.');
                return [3 /*break*/, 6];
            case 5:
                loading.value = false;
                return [7 /*endfinally*/];
            case 6: return [2 /*return*/];
        }
    });
}); };
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
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mb-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0 sm:space-x-4" }));
/** @type {__VLS_StyleScopedClasses['mb-6']} */ ;
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:flex-row']} */ ;
/** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
/** @type {__VLS_StyleScopedClasses['items-center']} */ ;
/** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:space-y-0']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:space-x-4']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 w-full sm:w-auto" }));
/** @type {__VLS_StyleScopedClasses['flex']} */ ;
/** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:flex-row']} */ ;
/** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:space-y-0']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:space-x-4']} */ ;
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.input)(__assign({ type: "text", value: (__VLS_ctx.searchQuery), placeholder: "Search by username..." }, { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full sm:w-auto" }));
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
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.select, __VLS_intrinsics.select)(__assign({ value: (__VLS_ctx.statusFilter) }, { class: "px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full sm:w-auto" }));
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
/** @type {__VLS_StyleScopedClasses['w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:w-auto']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (0),
});
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (1),
});
__VLS_asFunctionalElement1(__VLS_intrinsics.option, __VLS_intrinsics.option)({
    value: (2),
});
__VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "bg-white shadow overflow-x-auto sm:rounded-lg" }));
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow']} */ ;
/** @type {__VLS_StyleScopedClasses['overflow-x-auto']} */ ;
/** @type {__VLS_StyleScopedClasses['sm:rounded-lg']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.table, __VLS_intrinsics.table)(__assign({ class: "min-w-full divide-y divide-gray-200" }));
/** @type {__VLS_StyleScopedClasses['min-w-full']} */ ;
/** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
/** @type {__VLS_StyleScopedClasses['divide-gray-200']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.thead, __VLS_intrinsics.thead)(__assign({ class: "bg-gray-50" }));
/** @type {__VLS_StyleScopedClasses['bg-gray-50']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ scope: "col" }, { class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ scope: "col" }, { class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ scope: "col" }, { class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ scope: "col" }, { class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ scope: "col" }, { class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ scope: "col" }, { class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ scope: "col" }, { class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.th, __VLS_intrinsics.th)(__assign({ scope: "col" }, { class: "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider min-w-[120px]" }));
/** @type {__VLS_StyleScopedClasses['px-6']} */ ;
/** @type {__VLS_StyleScopedClasses['py-3']} */ ;
/** @type {__VLS_StyleScopedClasses['text-left']} */ ;
/** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
/** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
/** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
/** @type {__VLS_StyleScopedClasses['uppercase']} */ ;
/** @type {__VLS_StyleScopedClasses['tracking-wider']} */ ;
/** @type {__VLS_StyleScopedClasses['min-w-[120px]']} */ ;
__VLS_asFunctionalElement1(__VLS_intrinsics.tbody, __VLS_intrinsics.tbody)(__assign({ class: "bg-white divide-y divide-gray-200" }));
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['divide-y']} */ ;
/** @type {__VLS_StyleScopedClasses['divide-gray-200']} */ ;
if (__VLS_ctx.loading) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ colspan: "7" }, { class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
}
else if (!__VLS_ctx.users || __VLS_ctx.users.length === 0) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({});
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ colspan: "7" }, { class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-center" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-center']} */ ;
}
var _loop_1 = function (user) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.tr, __VLS_intrinsics.tr)({
        key: ((_a = user.id) !== null && _a !== void 0 ? _a : Math.random()),
    });
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    ((_b = user.id) !== null && _b !== void 0 ? _b : 'N/A');
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-900" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-900']} */ ;
    ((_c = user.username) !== null && _c !== void 0 ? _c : 'N/A');
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    ((_d = user.email) !== null && _d !== void 0 ? _d : 'N/A');
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "flex flex-wrap gap-1" }));
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex-wrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['gap-1']} */ ;
    for (var _h = 0, _j = __VLS_vFor(((user.roles || []))); _h < _j.length; _h++) {
        var role = _j[_h][0];
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ key: (role.id) }, { class: "px-2 py-0.5 text-xs rounded-full bg-indigo-100 text-indigo-700" }));
        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-0.5']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-indigo-100']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-indigo-700']} */ ;
        (role.displayName);
        // @ts-ignore
        [searchQuery, statusFilter, loading, users, users, users,];
    }
    if (!user.roles || user.roles.length === 0) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "text-gray-400" }));
        /** @type {__VLS_StyleScopedClasses['text-gray-400']} */ ;
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: ([
            __VLS_ctx.getUserStatusInfo(user.status).class,
            'px-2 inline-flex text-xs leading-5 font-semibold rounded-full'
        ]) }));
    /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
    /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['leading-5']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-semibold']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded-full']} */ ;
    (__VLS_ctx.getUserStatusInfo(user.status).text);
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    (__VLS_ctx.formatDate(user.createdAt));
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-sm text-gray-500" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
    (user.vipExpireAt ? __VLS_ctx.formatDate(user.vipExpireAt) : 'None');
    __VLS_asFunctionalElement1(__VLS_intrinsics.td, __VLS_intrinsics.td)(__assign({ class: "px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2" }));
    /** @type {__VLS_StyleScopedClasses['px-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['whitespace-nowrap']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-right']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-x-2']} */ ;
    if (user.status === 2) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(user.status === 2))
                    return;
                __VLS_ctx.setUserStatus(user.id, 1);
                // @ts-ignore
                [getUserStatusInfo, getUserStatusInfo, formatDate, formatDate, setUserStatus,];
            } }, { disabled: (user.id === null) }), { class: "text-green-600 hover:text-green-900 disabled:opacity-50 disabled:cursor-not-allowed px-3 py-1 text-xs rounded" }));
        /** @type {__VLS_StyleScopedClasses['text-green-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-green-900']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    }
    if (user.status === 1) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(user.status === 1))
                    return;
                __VLS_ctx.setUserStatus(user.id, 2);
                // @ts-ignore
                [setUserStatus,];
            } }, { disabled: (user.id === null) }), { class: "text-yellow-600 hover:text-yellow-900 disabled:opacity-50 disabled:cursor-not-allowed px-3 py-1 text-xs rounded" }));
        /** @type {__VLS_StyleScopedClasses['text-yellow-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-yellow-900']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    }
    if (!user.isAdmin) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(!user.isAdmin))
                    return;
                __VLS_ctx.setUserAsAdmin(user.id);
                // @ts-ignore
                [setUserAsAdmin,];
            } }, { disabled: (user.id === null) }), { class: "text-indigo-600 hover:text-indigo-900 disabled:opacity-50 disabled:cursor-not-allowed px-3 py-1 text-xs rounded" }));
        /** @type {__VLS_StyleScopedClasses['text-indigo-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-indigo-900']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    }
    if (user.isAdmin) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(user.isAdmin))
                    return;
                __VLS_ctx.removeAdmin(user.id);
                // @ts-ignore
                [removeAdmin,];
            } }, { disabled: (user.id === null) }), { class: "text-red-600 hover:text-red-900 disabled:opacity-50 disabled:cursor-not-allowed px-3 py-1 text-xs rounded" }));
        /** @type {__VLS_StyleScopedClasses['text-red-600']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:text-red-900']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    }
    __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
            var _a = [];
            for (var _i = 0; _i < arguments.length; _i++) {
                _a[_i] = arguments[_i];
            }
            var $event = _a[0];
            __VLS_ctx.grantVip(user.id);
            // @ts-ignore
            [grantVip,];
        } }, { disabled: (user.id === null) }), { class: "text-yellow-600 hover:text-yellow-900 disabled:opacity-50 disabled:cursor-not-allowed px-3 py-1 text-xs rounded" }));
    /** @type {__VLS_StyleScopedClasses['text-yellow-600']} */ ;
    /** @type {__VLS_StyleScopedClasses['hover:text-yellow-900']} */ ;
    /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
    /** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
    /** @type {__VLS_StyleScopedClasses['px-3']} */ ;
    /** @type {__VLS_StyleScopedClasses['py-1']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-xs']} */ ;
    /** @type {__VLS_StyleScopedClasses['rounded']} */ ;
    // @ts-ignore
    [];
};
for (var _i = 0, _e = __VLS_vFor((__VLS_ctx.users)); _i < _e.length; _i++) {
    var user = _e[_i][0];
    _loop_1(user);
}
if (!__VLS_ctx.loading && __VLS_ctx.totalUsers > 0) {
    __VLS_asFunctionalElement1(__VLS_intrinsics.div, __VLS_intrinsics.div)(__assign({ class: "mt-6 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0" }));
    /** @type {__VLS_StyleScopedClasses['mt-6']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex']} */ ;
    /** @type {__VLS_StyleScopedClasses['flex-col']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:flex-row']} */ ;
    /** @type {__VLS_StyleScopedClasses['justify-between']} */ ;
    /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
    /** @type {__VLS_StyleScopedClasses['space-y-4']} */ ;
    /** @type {__VLS_StyleScopedClasses['sm:space-y-0']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.p, __VLS_intrinsics.p)(__assign({ class: "text-sm text-gray-700" }));
    /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
    /** @type {__VLS_StyleScopedClasses['text-gray-700']} */ ;
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium" }));
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    ((__VLS_ctx.currentPage - 1) * __VLS_ctx.pageSize + 1);
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium" }));
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    (Math.min(__VLS_ctx.currentPage * __VLS_ctx.pageSize, __VLS_ctx.totalUsers));
    __VLS_asFunctionalElement1(__VLS_intrinsics.span, __VLS_intrinsics.span)(__assign({ class: "font-medium" }));
    /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
    (__VLS_ctx.totalUsers);
    if (__VLS_ctx.totalPages > 1) {
        __VLS_asFunctionalElement1(__VLS_intrinsics.nav, __VLS_intrinsics.nav)(__assign({ class: "relative z-0 inline-flex rounded-md shadow-sm -space-x-px" }, { 'aria-label': "Pagination" }));
        /** @type {__VLS_StyleScopedClasses['relative']} */ ;
        /** @type {__VLS_StyleScopedClasses['z-0']} */ ;
        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['shadow-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['-space-x-px']} */ ;
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(!__VLS_ctx.loading && __VLS_ctx.totalUsers > 0))
                    return;
                if (!(__VLS_ctx.totalPages > 1))
                    return;
                __VLS_ctx.goToPage(__VLS_ctx.currentPage - 1);
                // @ts-ignore
                [loading, totalUsers, totalUsers, totalUsers, currentPage, currentPage, currentPage, pageSize, pageSize, totalPages, goToPage,];
            } }, { disabled: (__VLS_ctx.currentPage === 1) }), { class: "relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed" }));
        /** @type {__VLS_StyleScopedClasses['relative']} */ ;
        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-l-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
        var _loop_2 = function (pageNumber) {
            __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                    var _a = [];
                    for (var _i = 0; _i < arguments.length; _i++) {
                        _a[_i] = arguments[_i];
                    }
                    var $event = _a[0];
                    if (!(!__VLS_ctx.loading && __VLS_ctx.totalUsers > 0))
                        return;
                    if (!(__VLS_ctx.totalPages > 1))
                        return;
                    __VLS_ctx.goToPage(pageNumber);
                    // @ts-ignore
                    [currentPage, goToPage, paginationRange,];
                } }, { key: (pageNumber) }), { class: ([
                    'relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium',
                    __VLS_ctx.currentPage === pageNumber ? 'z-10 bg-indigo-50 border-indigo-500 text-indigo-600' : 'bg-white text-gray-700 hover:bg-gray-50'
                ]) }));
            /** @type {__VLS_StyleScopedClasses['relative']} */ ;
            /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
            /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
            /** @type {__VLS_StyleScopedClasses['px-4']} */ ;
            /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
            /** @type {__VLS_StyleScopedClasses['border']} */ ;
            /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
            /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
            /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
            (pageNumber);
            // @ts-ignore
            [currentPage,];
        };
        for (var _f = 0, _g = __VLS_vFor((__VLS_ctx.paginationRange)); _f < _g.length; _f++) {
            var pageNumber = _g[_f][0];
            _loop_2(pageNumber);
        }
        __VLS_asFunctionalElement1(__VLS_intrinsics.button, __VLS_intrinsics.button)(__assign(__assign({ onClick: function () {
                var _a = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    _a[_i] = arguments[_i];
                }
                var $event = _a[0];
                if (!(!__VLS_ctx.loading && __VLS_ctx.totalUsers > 0))
                    return;
                if (!(__VLS_ctx.totalPages > 1))
                    return;
                __VLS_ctx.goToPage(__VLS_ctx.currentPage + 1);
                // @ts-ignore
                [currentPage, goToPage,];
            } }, { disabled: (__VLS_ctx.currentPage === __VLS_ctx.totalPages || __VLS_ctx.totalPages === 0) }), { class: "relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed" }));
        /** @type {__VLS_StyleScopedClasses['relative']} */ ;
        /** @type {__VLS_StyleScopedClasses['inline-flex']} */ ;
        /** @type {__VLS_StyleScopedClasses['items-center']} */ ;
        /** @type {__VLS_StyleScopedClasses['px-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['py-2']} */ ;
        /** @type {__VLS_StyleScopedClasses['rounded-r-md']} */ ;
        /** @type {__VLS_StyleScopedClasses['border']} */ ;
        /** @type {__VLS_StyleScopedClasses['border-gray-300']} */ ;
        /** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-sm']} */ ;
        /** @type {__VLS_StyleScopedClasses['font-medium']} */ ;
        /** @type {__VLS_StyleScopedClasses['text-gray-500']} */ ;
        /** @type {__VLS_StyleScopedClasses['hover:bg-gray-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:opacity-50']} */ ;
        /** @type {__VLS_StyleScopedClasses['disabled:cursor-not-allowed']} */ ;
    }
}
// @ts-ignore
[currentPage, totalPages, totalPages,];
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
