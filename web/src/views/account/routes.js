"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var accountRoutes = [
    {
        path: '/account',
        component: function () { return Promise.resolve().then(function () { return require('../../components/AccountLayout.vue'); }); },
        children: [
            {
                path: '',
                redirect: '/account/profile'
            },
            {
                path: 'profile',
                name: 'AccountProfile',
                component: function () { return Promise.resolve().then(function () { return require('./Profile.vue'); }); }
            },
            {
                path: 'settings',
                name: 'AccountSettings',
                component: function () { return Promise.resolve().then(function () { return require('./Settings.vue'); }); }
            },
            {
                path: 'teacher-application',
                name: 'TeacherApplication',
                component: function () { return Promise.resolve().then(function () { return require('./TeacherApplication.vue'); }); }
            },
            {
                path: 'mcp-tokens',
                name: 'AccountMCPTokens',
                component: function () { return Promise.resolve().then(function () { return require('./MCPTokenManagement.vue'); }); }
            },
            {
                path: 'cli-tokens',
                name: 'AccountCLITokens',
                component: function () { return Promise.resolve().then(function () { return require('./CLITokens.vue'); }); }
            },
            {
                path: 'analytics',
                name: 'StudentAnalytics',
                component: function () { return Promise.resolve().then(function () { return require('./StudentAnalytics.vue'); }); }
            }
        ]
    }
];
exports.default = accountRoutes;
