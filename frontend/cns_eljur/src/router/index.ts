import AdminDashBoard from '@/views/adminDashBoard.vue'
import Constructor1 from '@/views/constructor1.vue'
import Constructor2 from '@/views/constructor2.vue'
import StartPage from '@/views/startPage.vue'
import dbStore from '@/views/dbStore.vue'
import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'


const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'startPage',
        component: StartPage
    },
    {
        path: '/admin',
        name: 'adminDashboard',
        component: AdminDashBoard
    },
    {
        path: '/constructor1',
        name: 'constructor1',
        component: Constructor1
    },
    {
        path: '/constructor2',
        name: 'constructor2',
        component: Constructor2
    },
    {
        path: '/db',
        name: 'dbStore',
        component: dbStore
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
