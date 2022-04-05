import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';

import HomeView from '../views/HomeView.vue';
import LoginView from '../views/LoginView.vue';
import ProfileView from '../views/ProfileView.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/:tab(login|register)',
    name: 'login',
    component: LoginView,
  },
  {
    path: '/profile/:userId',
    name: 'profile',
    component: ProfileView,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
