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
  {
    path: "/constructorPageNoSchedule",
    name: "constructorPageNoSchedule",
    component: () => import("../views/constructorPage - no schedule.vue"),
  },
  {
    path: "/constructorPageScheduleChosen",
    name: "constructorPageScheduleChosen",
    component: () => import("../views/constructorPage - schedule chosen.vue"),
  },
  {
    path: "/DataBaseEditorPageNoObj",
    name: "DataBaseEditorPageNoObj",
    component: () => import("../views/DataBaseEditorPage - no obj.vue"),
  },
  {
    path: "/DataBaseEditorPage",
    name: "DataBaseEditorPage",
    component: () => import("../views/DataBaseEditorPage.vue"),
  },
  {
    path: "/schedulePagePageNoSchedule",
    name: "schedulePagePageNoSchedule",
    component: () => import("../views/schedulePage - no schedule.vue"),
  },
  {
    path: "/schedulePage",
    name: "schedulePage",
    component: () => import("../views/SchedulePage.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
