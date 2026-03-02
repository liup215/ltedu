import type { RouteRecordRaw } from 'vue-router'

const adminRoutes: RouteRecordRaw[] = [
  {
    path: '/admin',
    component: () => import('../../components/admin/AdminLayout.vue'),
    meta: { requiresAdmin: true },
    children: [
      {
        path: '',
        redirect: '/admin/dashboard'
      },
      {
        path: 'dashboard',
        name: 'AdminDashboard',
        component: () => import('./AdminDashboard.vue')
      },
      {
        path: 'users',
        name: 'AdminUserManagement',
        component: () => import('./UserManagement.vue')
      },
      {
        path: 'organisations',
        name: 'AdminOrganisationManagement',
        component: () => import('./OrganisationManagement.vue')
      },
      {
        path: 'organisations/create',
        name: 'AdminOrganisationCreate',
        component: () => import('./OrganisationForm.vue')
      },
      {
        path: 'organisations/:id/edit',
        name: 'AdminOrganisationEdit',
        component: () => import('./OrganisationForm.vue')
      },
      {
        path: 'qualifications',
        name: 'AdminQualificationManagement',
        component: () => import('./QualificationManagement.vue')
      },
      {
        path: 'qualifications/create',
        name: 'AdminQualificationCreate',
        component: () => import('./QualificationForm.vue')
      },
      {
        path: 'qualifications/:id/edit',
        name: 'AdminQualificationEdit',
        component: () => import('./QualificationForm.vue')
      },
      {
        path: 'syllabuses',
        name: 'AdminSyllabusManagement',
        component: () => import('./SyllabusManagement.vue')
      },
      {
        path: 'syllabuses/create',
        name: 'AdminSyllabusCreate',
        component: () => import('./SyllabusForm.vue')
      },
      {
        path: 'syllabuses/:id/edit',
        name: 'AdminSyllabusEdit',
        component: () => import('./SyllabusForm.vue')
      },
      {
        path: 'syllabuses/:id/chapters',
        name: 'AdminChapterManagement',
        component: () => import('./ChapterManagement.vue')
      },
      {
        path: 'syllabuses/:id/knowledge-points',
        name: 'AdminKnowledgePointManagement',
        component: () => import('./KnowledgePointManagement.vue')
      },
      {
        path: 'knowledge-points/migration',
        name: 'AdminKnowledgePointMigration',
        component: () => import('./KnowledgePointMigration.vue')
      },
      {
        path: 'system-settings',
        name: 'AdminSystemSettings',
        component: () => import('./SystemSettingsAdmin.vue')
      },
      {
        path: 'paper-series',
        name: 'AdminPaperSeriesManagement',
        component: () => import('./PaperSeriesManagement.vue')
      },
      {
        path: 'paper-series/create',
        name: 'AdminPaperSeriesCreate',
        component: () => import('./PaperSeriesForm.vue')
      },
      {
        path: 'paper-series/:id/edit',
        name: 'AdminPaperSeriesEdit',
        component: () => import('./PaperSeriesForm.vue')
      },
      {
        path: 'paper-codes',
        name: 'AdminPaperCodeManagement',
        component: () => import('./PaperCodeManagement.vue')
      },
      {
        path: 'paper-code/create',
        name: 'AdminPaperCodeCreate',
        component: () => import('./PaperCodeForm.vue')
      },
      {
        path: 'paper-code/:id/edit',
        name: 'AdminPaperCodeEdit',
        component: () => import('./PaperCodeForm.vue')
      },
      {
        path: 'past-papers',
        name: 'AdminPastPaperManagement',
        component: () => import('./PastPaperManagement.vue')
      },
      {
        path: 'past-paper/create',
        name: 'AdminPastPaperCreate',
        component: () => import('./PastPaperForm.vue')
      },
      {
        path: 'past-paper/:id/edit',
        name: 'AdminPastPaperEdit',
        component: () => import('./PastPaperForm.vue')
      },
      {
        path: 'exam-papers',
        name: 'AdminExamPaperManagement',
        component: () => import('./ExamPaperManagement.vue')
      },
      {
        path: 'questions',
        name: 'AdminQuestionManagement',
        component: () => import('./QuestionManagement.vue')
      },
      {
        path: 'questions/create',
        name: 'AdminQuestionCreate',
        component: () => import('./QuestionForm.vue')
      },
      {
        path: 'questions/:id',
        name: 'AdminQuestionView',
        component: () => import('./QuestionView.vue')
      },
      {
        path: 'questions/:id/edit',
        name: 'AdminQuestionEdit',
        component: () => import('./QuestionForm.vue')
      },
      {
        path: 'teacher-applications',
        name: 'AdminTeacherApplications',
        component: () => import('./teacher-applications/List.vue')
      },
      {
        path: 'mcp-tokens',
        name: 'AdminMCPTokenManagement',
        component: () => import('./MCPTokenManagement.vue')
      },
      {
        path: 'classes',
        name: 'AdminClassManagement',
        component: () => import('./ClassManagement.vue')
      },
      {
        path: 'syllabuses/:syllabusId/exam-nodes',
        name: 'AdminExamNodeManagement',
        component: () => import('./ExamNodeManagement.vue')
      },
      {
        path: 'classes/:classId/learning-plans',
        name: 'AdminLearningPlanManagement',
        component: () => import('./LearningPlanManagement.vue')
      },
      {
        path: 'learning-plans/:planId/phase-plans',
        name: 'AdminPhasePlanManagement',
        component: () => import('./PhasePlanManagement.vue')
      },
      {
        path: 'roles',
        name: 'AdminRoleManagement',
        component: () => import('./RoleManagement.vue')
      },
      {
        path: 'permissions',
        name: 'AdminPermissionManagement',
        component: () => import('./PermissionManagement.vue')
      }
    ]
  }
]

export default adminRoutes
