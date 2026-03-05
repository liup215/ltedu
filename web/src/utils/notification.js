"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.showSuccess = showSuccess;
exports.showError = showError;
// src/utils/notification.ts
var notivue_1 = require("notivue");
/**
 * Show a global success notification using Notivue.
 * @param message Success message to display
 */
function showSuccess(message) {
    notivue_1.push.success({
        title: 'Success',
        message: message,
        duration: 2000,
    });
}
/**
 * Show a global error notification using Notivue.
 * @param message Error message to display
 */
function showError(message) {
    notivue_1.push.error({
        title: 'Error',
        message: message,
        duration: 2000, // Show for 5 seconds
    });
}
