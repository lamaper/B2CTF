<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import http from '../api/http';

const router = useRouter();

const form = ref({
  username: '',
  password: ''
});

const loading = ref(false);
const error = ref('');

const handleLogin = async () => {
  if (!form.value.username || !form.value.password) {
    error.value = '请填写用户名和密码';
    return;
  }

  loading.value = true;
  error.value = '';

  try {
    const response = await http.post('/login', form.value);
    if (response.code === 200) {
      // 存储token
      localStorage.setItem('token', response.data.token);
      // 跳转到首页
      router.push('/');
    } else {
      error.value = response.msg || '登录失败';
    }
  } catch (err) {
    error.value = err.response?.data?.error || '登录失败，请稍后重试';
    console.error(err);
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div class="login-container">
    <div class="login-card">
      <h1 class="login-title">用户登录</h1>
      
      <form @submit.prevent="handleLogin" class="login-form">
        <div class="form-group">
          <label for="username">用户名</label>
          <input 
            type="text" 
            id="username" 
            v-model="form.username" 
            placeholder="请输入用户名"
            required
          >
        </div>
        
        <div class="form-group">
          <label for="password">密码</label>
          <input 
            type="password" 
            id="password" 
            v-model="form.password" 
            placeholder="请输入密码"
            required
          >
        </div>
        
        <button 
          type="submit" 
          class="login-button"
          :disabled="loading"
        >
          {{ loading ? '登录中...' : '登录' }}
        </button>
      </form>
      
      <div v-if="error" class="message error">
        {{ error }}
      </div>
      
      <div class="register-link">
        没有账号？ <router-link to="/register">立即注册</router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
  padding: 20px;
}

.login-card {
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  padding: 40px;
  width: 100%;
  max-width: 400px;
}

.login-title {
  font-size: 1.8rem;
  color: #2c3e50;
  text-align: center;
  margin-bottom: 30px;
  font-weight: bold;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group label {
  font-size: 0.9rem;
  color: #34495e;
  font-weight: 500;
}

.form-group input {
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

.form-group input:focus {
  outline: none;
  border-color: #e74c3c;
  box-shadow: 0 0 0 3px rgba(231, 76, 60, 0.1);
}

.login-button {
  background-color: #e74c3c;
  color: white;
  border: none;
  padding: 12px;
  border-radius: 4px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-top: 10px;
}

.login-button:hover:not(:disabled) {
  background-color: #c0392b;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(231, 76, 60, 0.3);
}

.login-button:disabled {
  background-color: #bdc3c7;
  cursor: not-allowed;
}

.message {
  padding: 12px;
  border-radius: 4px;
  margin-top: 15px;
  text-align: center;
  font-size: 0.9rem;
}

.message.error {
  background-color: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}

.register-link {
  margin-top: 20px;
  text-align: center;
  font-size: 0.9rem;
  color: #7f8c8d;
}

.register-link a {
  color: #e74c3c;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s ease;
}

.register-link a:hover {
  color: #c0392b;
  text-decoration: underline;
}

@media (max-width: 480px) {
  .login-card {
    padding: 30px 20px;
  }
  
  .login-title {
    font-size: 1.5rem;
  }
}
</style>