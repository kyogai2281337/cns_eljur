import { createRouter, RouteRecordRaw, createWebHashHistory } from 'vue-router'
import AuthPage from '../views/AuthPage.vue'
import HomePage from '../views/HomePage.vue'
import dbStore from '@/views/dbStore.vue'
import constructorPage from '@/views/constructorPage.vue'
import filesPage from '@/views/filesPage.vue'

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
  },
  {
    path: '/files',
    name: 'filesPage',
    component: filesPage
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
