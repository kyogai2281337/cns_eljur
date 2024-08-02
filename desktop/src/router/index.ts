import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import AuthPage from '../views/AuthPage.vue'
import HomePage from '../views/HomePage.vue'
import dbStore from '@/views/dbStore.vue'
import constructorPage from '@/views/constructorPage.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'authPage',
    component: AuthPage
  },
  {
    path: '/home',
    name: 'homePage',
    component: HomePage
  },
  {
    path: '/db',
    name: 'dbStore',
    component: dbStore
  },
  {
    path: '/contra',
    name: 'constructorPage',
    component: constructorPage
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
