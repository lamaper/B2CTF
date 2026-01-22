// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-17
// 最后修改: 2026-01-17
// 描述: 前端路由配置文件
// ----------------------------------------------------------------------------
import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/Register.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('../views/Profile.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/challenges',
    name: 'Challenges',
    component: () => import('../views/Challenges.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/competitions',
    name: 'Competitions',
    component: () => import('../views/Competitions.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/create-competition',
    name: 'CreateCompetition',
    component: () => import('../views/CreateCompetition.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/create-challenge',
    name: 'CreateChallenge',
    component: () => import('../views/CreateChallenge.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/challenge/:id',
    name: 'ChallengeDetail',
    component: () => import('../views/ChallengeDetail.vue'),
    meta: { requiresAuth: true }
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

// 导航守卫，检查是否需要认证
router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth);
  const token = localStorage.getItem('token');
  
  if (requiresAuth && !token) {
    next({ name: 'Login' });
  } else {
    next();
  }
});

export default router;