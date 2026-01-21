<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import http from '../api/http';

const router = useRouter();

const form = ref({
  username: '',
  password: '',
  email: ''
});

const loading = ref(false);
const error = ref('');
const success = ref('');

const handleRegister = async () => {
  if (!form.value.username || !form.value.password || !form.value.email) {
    error.value = '请填写所有必填字段';
    return;
  }

  loading.value = true;
  error.value = '';
  success.value = '';

  try {
    const response = await http.post('/register', form.value);
    success.value = response.msg || '注册成功';
    setTimeout(() => {
      router.push('/login');
    }, 1500);
  } catch (err) {
    error.value = err.response?.data?.error || '注册失败，请稍后重试';
    console.error(err);
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div class="register-container">
    <div class="register-card">
      <h1 class="register-title">用户注册</h1>
      
      <form @submit.prevent="handleRegister" class="register-form">
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
          <label for="email">邮箱</label>
          <input 
            type="email" 
            id="email" 
            v-model="form.email" 
            placeholder="请输入邮箱"
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
          class="register-button"
          :disabled="loading"
        >
          {{ loading ? '注册中...' : '注册' }}
        </button>
      </form>
      
      <div v-if="success" class="message success">
        {{ success }}
      </div>
      
      <div v-if="error" class="message error">
        {{ error }}
      </div>
      
      <div class="login-link">
        已有账号？ <router-link to="/login">立即登录</router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
.register-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  padding: 20px;
}

.register-card {
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  padding: 40px;
  width: 100%;
  max-width: 400px;
}

.register-title {
  font-size: 1.8rem;
  color: #2c3e50;
  text-align: center;
  margin-bottom: 30px;
  font-weight: bold;
}

.register-form {
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
  border-color: #3498db;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.1);
}

.register-button {
  background-color: #3498db;
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

.register-button:hover:not(:disabled) {
  background-color: #2980b9;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(52, 152, 219, 0.3);
}

.register-button:disabled {
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

.message.success {
  background-color: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
}

.message.error {
  background-color: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}

.login-link {
  margin-top: 20px;
  text-align: center;
  font-size: 0.9rem;
  color: #7f8c8d;
}

.login-link a {
  color: #3498db;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s ease;
}

.login-link a:hover {
  color: #2980b9;
  text-decoration: underline;
}

@media (max-width: 480px) {
  .register-card {
    padding: 30px 20px;
  }
  
  .register-title {
    font-size: 1.5rem;
  }
}
</style>