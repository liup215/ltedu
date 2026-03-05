"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var adminRoutes = [
    {
        path: '/admin',
        component: function () { return Promise.resolve().then(function () { return require('../../components/admin/AdminLayout.vue'); }); },
        meta: { requiresAdmin: true },
        children: [
            {
                path: '',
                redirect: '/admin/dashboard'
            },
            {
                path: 'dashboard',
                name: 'AdminDashboard',
                component: function () { return Promise.resolve().then(function () { return require('./AdminDashboard.vue'); }); }
            },
            {
                path: 'users',
                name: 'AdminUserManagement',
                component: function () { return Promise.resolve().then(function () { return require('./UserManagement.vue'); }); }
            },
            {
                path: 'organisations',
                name: 'AdminOrganisationManagement',
                component: function () { return Promise.resolve().then(function () { return require('./OrganisationManagement.vue'); }); }
            },
            {
                path: 'organisations/create',
                name: 'AdminOrganisationCreate',
                component: function () { return Promise.resolve().then(function () { return require('./OrganisationForm.vue'); }); }
            },
            {
                path: 'organisations/:id/edit',
                name: 'AdminOrganisationEdit',
                component: function () { return Promise.resolve().then(function () { return require('./OrganisationForm.vue'); }); }
            },
            {
                path: 'qualifications',
                name: 'AdminQualificationManagement',
                component: function () { return Promise.resolve().then(function () { return require('./QualificationManagement.vue'); }); }
            },
            {
                path: 'qualifications/create',
                name: 'AdminQualificationCreate',
                component: function () { return Promise.resolve().then(function () { return require('./QualificationForm.vue'); }); }
            },
            {
                path: 'qualifications/:id/edit',
                name: 'AdminQualificationEdit',
                component: function () { return Promise.resolve().then(function () { return require('./QualificationForm.vue'); }); }
            },
            {
                path: 'syllabuses',
                name: 'AdminSyllabusManagement',
                component: function () { return Promise.resolve().then(function () { return require('./SyllabusManagement.vue'); }); }
            },
            {
                path: 'syllabuses/create',
                name: 'AdminSyllabusCreate',
                component: function () { return Promise.resolve().then(function () { return require('./SyllabusForm.vue'); }); }
            },
            {
                path: 'syllabuses/:id/edit',
                name: 'AdminSyllabusEdit',
                component: function () { return Promise.resolve().then(function () { return require('./SyllabusForm.vue'); }); }
            },
            {
                path: 'syllabuses/:id/chapters',
                name: 'AdminChapterManagement',
                component: function () { return Promise.resolve().then(function () { return require('./ChapterManagement.vue'); }); }
            },
            {
                path: 'syllabuses/:id/knowledge-points',
                name: 'AdminKnowledgePointManagement',
                component: function () { return Promise.resolve().then(function () { return require('./KnowledgePointManagement.vue'); }); }
            },
            {
                path: 'knowledge-points/migration',
                name: 'AdminKnowledgePointMigration',
                component: function () { return Promise.resolve().then(function () { return require('./KnowledgePointMigration.vue'); }); }
            },
            {
                path: 'system-settings',
                name: 'AdminSystemSettings',
                component: function () { return Promise.resolve().then(function () { return require('./SystemSettingsAdmin.vue'); }); }
            },
            {
                path: 'paper-series',
                name: 'AdminPaperSeriesManagement',
                component: function () { return Promise.resolve().then(function () { return require('./PaperSeriesManagement.vue'); }); }
            },
            {
                path: 'paper-series/create',
                name: 'AdminPaperSeriesCreate',
                component: function () { return Promise.resolve().then(function () { return require('./PaperSeriesForm.vue'); }); }
            },
            {
                path: 'paper-series/:id/edit',
                name: 'AdminPaperSeriesEdit',
                component: function () { return Promise.resolve().then(function () { return require('./PaperSeriesForm.vue'); }); }
            },
            {
                path: 'paper-codes',
                name: 'AdminPaperCodeManagement',
                component: function () { return Promise.resolve().then(function () { return require('./PaperCodeManagement.vue'); }); }
            },
            {
                path: 'paper-code/create',
                name: 'AdminPaperCodeCreate',
                component: function () { return Promise.resolve().then(function () { return require('./PaperCodeForm.vue'); }); }
            },
            {
                path: 'paper-code/:id/edit',
                name: 'AdminPaperCodeEdit',
                component: function () { return Promise.resolve().then(function () { return require('./PaperCodeForm.vue'); }); }
            },
            {
                path: 'past-papers',
                name: 'AdminPastPaperManagement',
                component: function () { return Promise.resolve().then(function () { return require('./PastPaperManagement.vue'); }); }
            },
            {
                path: 'past-paper/create',
                name: 'AdminPastPaperCreate',
                component: function () { return Promise.resolve().then(function () { return require('./PastPaperForm.vue'); }); }
            },
            {
                path: 'past-paper/:id/edit',
                name: 'AdminPastPaperEdit',
                component: function () { return Promise.resolve().then(function () { return require('./PastPaperForm.vue'); }); }
            },
            {
                path: 'exam-papers',
                name: 'AdminExamPaperManagement',
                component: function () { return Promise.resolve().then(function () { return require('./ExamPaperManagement.vue'); }); }
            },
            {
                path: 'questions',
                name: 'AdminQuestionManagement',
                component: function () { return Promise.resolve().then(function () { return require('./QuestionManagement.vue'); }); }
            },
            {
                path: 'questions/create',
                name: 'AdminQuestionCreate',
                component: function () { return Promise.resolve().then(function () { return require('./QuestionForm.vue'); }); }
            },
            {
                path: 'questions/:id',
                name: 'AdminQuestionView',
                component: function () { return Promise.resolve().then(function () { return require('./QuestionView.vue'); }); }
            },
            {
                path: 'questions/:id/edit',
                name: 'AdminQuestionEdit',
                component: function () { return Promise.resolve().then(function () { return require('./QuestionForm.vue'); }); }
            },
            {
                path: 'teacher-applications',
                name: 'AdminTeacherApplications',
                component: function () { return Promise.resolve().then(function () { return require('./teacher-applications/List.vue'); }); }
            },
            {
                path: 'mcp-tokens',
                name: 'AdminMCPTokenManagement',
                component: function () { return Promise.resolve().then(function () { return require('./MCPTokenManagement.vue'); }); }
            },
            {
                path: 'classes',
                name: 'AdminClassManagement',
                component: function () { return Promise.resolve().then(function () { return require('./ClassManagement.vue'); }); }
            },
            {
                path: 'syllabuses/:syllabusId/exam-nodes',
                name: 'AdminExamNodeManagement',
                component: function () { return Promise.resolve().then(function () { return require('./ExamNodeManagement.vue'); }); }
            },
            {
                path: 'classes/:classId/learning-plans',
                name: 'AdminLearningPlanManagement',
                component: function () { return Promise.resolve().then(function () { return require('./LearningPlanManagement.vue'); }); }
            },
            {
                path: 'learning-plans/:planId/phase-plans',
                name: 'AdminPhasePlanManagement',
                component: function () { return Promise.resolve().then(function () { return require('./PhasePlanManagement.vue'); }); }
            },
            {
                path: 'roles',
                name: 'AdminRoleManagement',
                component: function () { return Promise.resolve().then(function () { return require('./RoleManagement.vue'); }); }
            },
            {
                path: 'permissions',
                name: 'AdminPermissionManagement',
                component: function () { return Promise.resolve().then(function () { return require('./PermissionManagement.vue'); }); }
            },
            {
                path: 'analytics',
                name: 'AdminAnalyticsDashboard',
                component: function () { return Promise.resolve().then(function () { return require('./AnalyticsDashboard.vue'); }); }
            }
        ]
    }
];
exports.default = adminRoutes;
