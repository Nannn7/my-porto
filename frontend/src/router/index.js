import { createRouter, createWebHistory } from 'vue-router'
<<<<<<< HEAD

=======
>>>>>>> origin/main
import Home from '../views/Home.vue'
import About from '../views/About.vue'
import Projects from '../views/Projects.vue'
import Contact from '../views/Contact.vue'

<<<<<<< HEAD
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'home', component: Home },
    { path: '/about', name: 'about', component: About },
    { path: '/projects', name: 'projects', component: Projects },
    { path: '/contact', name: 'contact', component: Contact },
  ],
})

export default router
=======
const routes = [
  { path: '/', name: 'home', component: Home },
  { path: '/about', name: 'about', component: About },
  { path: '/projects', name: 'projects', component: Projects },
  { path: '/contact', name: 'contact', component: Contact },
]

export default createRouter({
  history: createWebHistory(),
  routes,
})
>>>>>>> origin/main
