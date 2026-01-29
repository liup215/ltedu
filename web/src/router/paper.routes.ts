import PaperSeriesManagement from '../views/admin/PaperSeriesManagement.vue'
import PaperSeriesForm from '../views/admin/PaperSeriesForm.vue' // Added import
import PaperCodeManagement from '../views/admin/PaperCodeManagement.vue'

export const paperRoutes = [
  {
    path: '/admin/paper-series',
    name: 'PaperSeriesManagement',
    component: PaperSeriesManagement,
    meta: {
      requiresAuth: true,
      allowedRoles: ['admin']
    }
  },
  { // Added route for creating paper series
    path: '/admin/paper-series/create',
    name: 'AdminPaperSeriesCreate',
    component: PaperSeriesForm,
    meta: {
      requiresAuth: true,
      allowedRoles: ['admin']
    }
  },
  { // Added route for editing paper series
    path: '/admin/paper-series/edit/:id',
    name: 'AdminPaperSeriesEdit',
    component: PaperSeriesForm,
    meta: {
      requiresAuth: true,
      allowedRoles: ['admin']
    }
  },
  {
    path: '/admin/paper-codes',
    name: 'PaperCodeManagement',
    component: PaperCodeManagement,
    meta: {
      requiresAuth: true,
      allowedRoles: ['admin']
    }
  }
]
