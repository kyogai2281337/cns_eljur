import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "LoginPage",
    component: () => import("../views/LoginPage.vue"),
  },
  {
    path: "/reg",
    name: "RegisterPage",
    component: () => import("../views/RegisterPage.vue"),
  },
  {
    path: "/home",
    name: "HomePage",
    component: () => import("../views/HomePage.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
