import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import Login from '../views/Login.vue'
import Home from '../views/Home.vue'
import Layout from '../components/Layout.vue'
import Register from '../views/Register.vue'

// Modular route imports
import accountRoutes from '../views/account/routes'
import adminRoutes from '../views/admin/routes'

// Import ExamPaperForm for teacher paper builder
const ExamPaperForm = () => import('../views/paper/ExamPaperForm.vue')
const TeacherExamPaperManagement = () => import('../views/paper/TeacherExamPaperManagement.vue')
const ExamPaperPreview = () => import('../views/paper/ExamPaperPreview.vue')
const QuickPractice = () => import('../views/QuickPractice.vue')
const PaperPractice = () => import('../views/PaperPractice.vue')

const routes: RouteRecordRaw[] = [
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/',
    component: Layout,
    children: [
      {
        path: '',
        name: 'Home',
        component: Home
      },
      {
        path: 'login',
        name: 'Login',
        component: Login
      },
      // Teacher Exam Paper Builder route
      {
        path: '/paper/exam/create',
        name: 'ExamPaperBuilder',
        component: ExamPaperForm,
        meta: { requiresAuth: false, requiresTeacher: false }
      },
      // Teacher Exam Paper Management
      {
        path: '/paper/exam/teacher',
        name: 'TeacherExamPaperManagement',
        component: TeacherExamPaperManagement,
        meta: { requiresAuth: true, requiresTeacher: true }
      },
      // Teacher Exam Paper Edit Route
      {
        path: '/paper/exam/edit/:id',
        name: 'ExamPaperEdit',
        component: ExamPaperForm,
        meta: { requiresAuth: true, requiresTeacher: true }
      },
      // Teacher Exam Paper Preview Route
      {
        path: '/paper/exam/preview/:id',
        name: 'ExamPaperPreview',
        component: ExamPaperPreview,
        meta: { requiresAuth: true, requiresTeacher: true }
      },
      // Quick Practice route for students
      {
        path: '/practice/quick',
        name: 'QuickPractice',
        component: QuickPractice
      },
      // Past Paper Practice route for students
      {
        path: '/practice/paper',
        name: 'PaperPractice',
        component: PaperPractice
      },
      // Donation page
      {
        path: '/donate',
        name: 'Donation',
        component: () => import('../views/Donation.vue')
      },
      // Blog routes (public)
      {
        path: '/blog',
        name: 'Blog',
        component: () => import('../views/Blog.vue')
      },
      {
        path: '/blog/:slug',
        name: 'BlogPost',
        component: () => import('../views/BlogPost.vue')
      }
    ]
  },
  ...accountRoutes,
  ...adminRoutes
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Import user store for navigation guards
import { useUserStore } from '../stores/userStore'

router.beforeEach(async (to, _, next) => {
  const userStore = useUserStore();
  const token = localStorage.getItem('authToken');

  // Only fetch user if a token exists
  if (token && !userStore.user) {
    await userStore.fetchCurrentUser();
  }

  const isAuthenticated = userStore.isAuthenticated;
  const userIsAdmin = userStore.user?.isAdmin === true;

  // const guestOnlyRouteNames = ['Login', 'Register'];
  // if (guestOnlyRouteNames.includes(to.name as string) && isAuthenticated) {
  //   console.log('Navigation guard: Guest only route, authenticated. Redirecting to Home.');
  //   return next({ name: 'Home' });
  // }

  const publicRouteNames = ['Login', 'Register', 'Setup', 'Home', 'QuickPractice', 'PaperPractice', 'Donation', 'ExamPaperBuilder', 'Blog', 'BlogPost'];
  const requiresAuth = !publicRouteNames.includes(to.name as string);

  if (requiresAuth && !isAuthenticated) {
    console.log(`Navigation guard: Route ${String(to.name)} requires auth, user not authenticated. Redirecting to Login.`);
    return next({ name: 'Login' });
  }

  if (to.meta.requiresTeacher && !(userStore.user?.isTeacher || userStore.user?.isAdmin)) {
    return next({ name: 'Home' });
  }

  if (to.meta.requiresAdmin && !userIsAdmin) {
    console.log(`Navigation guard: Route ${String(to.name)} requires admin, user is not admin. Redirecting to Home.`);
    return next({ name: 'Home' });
  }

  next();
})

export default router
