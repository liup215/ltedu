import type { RouteRecordRaw } from 'vue-router'

const accountRoutes: RouteRecordRaw[] = [
  {
    path: '/account',
    component: () => import('../../components/AccountLayout.vue'),
    children: [
      {
        path: '',
        redirect: '/account/profile'
      },
      {
        path: 'profile',
        name: 'AccountProfile',
        component: () => import('./Profile.vue')
      },
      {
        path: 'settings',
        name: 'AccountSettings',
        component: () => import('./Settings.vue')
      },
      {
        path: 'teacher-application',
        name: 'TeacherApplication',
        component: () => import('./TeacherApplication.vue')
      },
      {
        path: 'mcp-tokens',
        name: 'AccountMCPTokens',
        component: () => import('./MCPTokenManagement.vue')
      },
      {
        path: 'cli-tokens',
        name: 'AccountCLITokens',
        component: () => import('./CLITokens.vue')
      },
      {
        path: 'analytics',
        name: 'StudentAnalytics',
        component: () => import('./StudentAnalytics.vue')
      }
    ]
  }
]

export default accountRoutes
