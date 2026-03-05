"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.paperRoutes = void 0;
var PaperSeriesManagement_vue_1 = require("../views/admin/PaperSeriesManagement.vue");
var PaperSeriesForm_vue_1 = require("../views/admin/PaperSeriesForm.vue"); // Added import
var PaperCodeManagement_vue_1 = require("../views/admin/PaperCodeManagement.vue");
exports.paperRoutes = [
    {
        path: '/admin/paper-series',
        name: 'PaperSeriesManagement',
        component: PaperSeriesManagement_vue_1.default,
        meta: {
            requiresAuth: true,
            allowedRoles: ['admin']
        }
    },
    {
        path: '/admin/paper-series/create',
        name: 'AdminPaperSeriesCreate',
        component: PaperSeriesForm_vue_1.default,
        meta: {
            requiresAuth: true,
            allowedRoles: ['admin']
        }
    },
    {
        path: '/admin/paper-series/edit/:id',
        name: 'AdminPaperSeriesEdit',
        component: PaperSeriesForm_vue_1.default,
        meta: {
            requiresAuth: true,
            allowedRoles: ['admin']
        }
    },
    {
        path: '/admin/paper-codes',
        name: 'PaperCodeManagement',
        component: PaperCodeManagement_vue_1.default,
        meta: {
            requiresAuth: true,
            allowedRoles: ['admin']
        }
    }
];
