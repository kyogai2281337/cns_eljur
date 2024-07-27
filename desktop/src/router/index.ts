import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import AuthPage from '../views/AuthPage.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'authPage',
    component: AuthPage
  },
  {
    path: '/about',
    name: 'about',
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
