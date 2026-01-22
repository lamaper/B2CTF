<!-- ---------------------------------------------------------------------------- -->
<!-- Copyright (c) 2026 lamaper                                                      -->
<!-- 创建日期: 2026-01-17                                                          -->
<!-- 最后修改: 2026-01-22                                                          -->
<!-- 描述: 前端应用主组件，包含导航栏和全局状态管理                                   -->
<!-- ---------------------------------------------------------------------------- -->
<script setup>
import { ref, computed, onMounted, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import http from './api/http';

const route = useRoute();
const router = useRouter();
const isLoggedIn = ref(false);
const isAdmin = ref(false);

const checkLoginStatus = () => {
  const token = localStorage.getItem('token');
  isLoggedIn.value = !!token;
  const userRole = localStorage.getItem('userRole');
  isAdmin.value = userRole === 'admin';
};

const logout = () => {
  localStorage.removeItem('token');
  localStorage.removeItem('userRole');
  isLoggedIn.value = false;
  isAdmin.value = false;
  router.push('/login');
};

const navLinks = computed(() => {
  if (isLoggedIn.value) {
    const links = [
      { name: '首页', path: '/' },
      { name: '题目', path: '/challenges' },
      { name: '比赛', path: '/competitions' },
      { name: '用户信息', path: '/profile' }
    ];
    
    // 管理员链接
    if (isAdmin.value) {
      links.push({ name: '创建比赛', path: '/create-competition' });
      links.push({ name: '创建题目', path: '/create-challenge' });
    }
    
    return links;
  } else {
    return [
      { name: '首页', path: '/' },
      { name: '注册', path: '/register' },
      { name: '登录', path: '/login' }
    ];
  }
});

onMounted(() => {
  checkLoginStatus();
});

// 监听路由变化，每次路由变化时都检查登录状态
watch(() => route.path, () => {
  checkLoginStatus();
});
</script>

<template>
  <div class="app">
    <!-- 导航栏 -->
    <nav class="navbar">
      <div class="navbar-container">
        <div class="navbar-brand">
          <router-link to="/" class="brand-link">
            <h1>B2CTF</h1>
          </router-link>
        </div>
        
        <div class="navbar-links">
          <router-link 
            v-for="link in navLinks" 
            :key="link.path" 
            :to="link.path"
            class="nav-link"
            :class="{ active: route.path === link.path }"
          >
            {{ link.name }}
          </router-link>
          
          <button v-if="isLoggedIn" class="logout-button" @click="logout">
            退出登录
          </button>
        </div>
      </div>
    </nav>
    
    <!-- 主内容区 -->
    <main class="main-content">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>
    
    <!-- 页脚 -->
    <footer class="footer">
      <div class="footer-container">
        <p>&copy; 2026 B2CTF Platform. All rights reserved.</p>
        <div class="footer-links">
          <a href="#">关于我们</a>
          <a href="#">使用帮助</a>
          <a href="#">联系我们</a>
        </div>
      </div>
    </footer>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: Arial, sans-serif;
  background-color: #f5f5f5;
  color: #333;
  line-height: 1.6;
}

.app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

/* 导航栏样式 */
.navbar {
  background-color: #2c3e50;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  position: sticky;
  top: 0;
  z-index: 1000;
}

.navbar-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 70px;
}

.navbar-brand .brand-link {
  text-decoration: none;
  color: white;
}

.navbar-brand h1 {
  font-size: 1.8rem;
  font-weight: bold;
}

.navbar-links {
  display: flex;
  align-items: center;
  gap: 20px;
}

.nav-link {
  color: white;
  text-decoration: none;
  font-size: 1rem;
  font-weight: 500;
  transition: color 0.3s ease;
  padding: 8px 12px;
  border-radius: 4px;
}

.nav-link:hover {
  color: #3498db;
  background-color: rgba(255, 255, 255, 0.1);
}

.nav-link.active {
  color: #3498db;
  background-color: rgba(255, 255, 255, 0.1);
}

.logout-button {
  background-color: #e74c3c;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: all 0.3s ease;
}

.logout-button:hover {
  background-color: #c0392b;
  transform: translateY(-2px);
}

/* 主内容区样式 */
.main-content {
  flex: 1;
  width: 100%;
}

/* 页脚样式 */
.footer {
  background-color: #2c3e50;
  color: white;
  padding: 30px 0;
  margin-top: auto;
}

.footer-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.footer-links {
  display: flex;
  gap: 20px;
}

.footer-links a {
  color: rgba(255, 255, 255, 0.8);
  text-decoration: none;
  transition: color 0.3s ease;
}

.footer-links a:hover {
  color: white;
}

/* 过渡动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .navbar-container {
    flex-direction: column;
    height: auto;
    padding: 15px 20px;
    gap: 10px;
  }
  
  .navbar-links {
    flex-wrap: wrap;
    justify-content: center;
  }
  
  .footer-container {
    flex-direction: column;
    gap: 15px;
    text-align: center;
  }
}
</style>
