import { createRouter, createWebHistory } from 'vue-router'

import Home from '../views/Home.vue'
import About from '../views/About.vue'
import Projects from '../views/Projects.vue'
import Contact from '../views/Contact.vue'
import AdminLogin from '../views/admin/AdminLogin.vue'
import AdminProjects from '../views/admin/AdminProjects.vue'

const ADMIN_TOKEN_STORAGE_KEY = 'myporto_admin_token'
const ADMIN_EXPIRES_AT_STORAGE_KEY = 'myporto_admin_expires_at'

function hasAdminSession() {
  if (typeof window === 'undefined') {
    return false
  }

  const token = window.localStorage.getItem(ADMIN_TOKEN_STORAGE_KEY)
  const expiresAt = window.localStorage.getItem(ADMIN_EXPIRES_AT_STORAGE_KEY)

  if (!token || !expiresAt) {
    return false
  }

  const expiresTime = Date.parse(expiresAt)
  return Number.isFinite(expiresTime) && expiresTime > Date.now()
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'home', component: Home },
    { path: '/about', name: 'about', component: About },
    { path: '/projects', name: 'projects', component: Projects },
    { path: '/contact', name: 'contact', component: Contact },
    {
      path: '/admin/login',
      name: 'admin-login',
      component: AdminLogin,
      meta: { layout: 'admin', adminGuestOnly: true },
    },
    {
      path: '/admin/projects',
      name: 'admin-projects',
      component: AdminProjects,
      meta: { layout: 'admin', requiresAdmin: true },
    },
  ],
})

router.beforeEach((to) => {
  const isAuthenticated = hasAdminSession()

  if (to.meta.requiresAdmin && !isAuthenticated) {
    return {
      name: 'admin-login',
      query: { redirect: to.fullPath },
    }
  }

  if (to.meta.adminGuestOnly && isAuthenticated) {
    return { name: 'admin-projects' }
  }

  return true
})

export default router
