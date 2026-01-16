<script setup>
import { ref } from 'vue';
import { ping } from './api/ping';

const response = ref(null);
const loading = ref(false);
const error = ref(null);

// 调用后端ping接口
const testPing = async () => {
  loading.value = true;
  error.value = null;
  
  try {
    const res = await ping();
    response.value = res;
  } catch (err) {
    error.value = '与后端通信失败，请检查后端服务是否启动';
    console.error(err);
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div class="app-container">
    <header class="app-header">
      <h1>B2CTF 平台</h1>
      <p>CTF 竞赛平台 - 简单起步</p>
    </header>
    
    <main class="app-main">
      <div class="test-section">
        <h2>测试与后端通信</h2>
        <button 
          class="ping-button" 
          @click="testPing"
          :disabled="loading"
        >
          {{ loading ? '正在测试...' : '测试后端连接' }}
        </button>
        
        <div v-if="response" class="response-success">
          <h3>通信成功！</h3>
          <p>后端返回：{{ JSON.stringify(response) }}</p>
        </div>
        
        <div v-if="error" class="response-error">
          <h3>通信失败</h3>
          <p>{{ error }}</p>
        </div>
      </div>
    </main>
    
    <footer class="app-footer">
      <p>© 2026 B2CTF Platform</p>
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
}

.app-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.app-header {
  text-align: center;
  padding: 40px 0;
  background-color: #2c3e50;
  color: white;
  border-radius: 8px;
  margin-bottom: 30px;
}

.app-header h1 {
  font-size: 2.5rem;
  margin-bottom: 10px;
}

.app-header p {
  font-size: 1.1rem;
  opacity: 0.9;
}

.app-main {
  background-color: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.test-section h2 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.ping-button {
  background-color: #3498db;
  color: white;
  border: none;
  padding: 12px 24px;
  font-size: 1rem;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.ping-button:hover:not(:disabled) {
  background-color: #2980b9;
}

.ping-button:disabled {
  background-color: #bdc3c7;
  cursor: not-allowed;
}

.response-success, .response-error {
  margin-top: 20px;
  padding: 15px;
  border-radius: 4px;
}

.response-success {
  background-color: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
}

.response-error {
  background-color: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}

.response-success h3, .response-error h3 {
  margin-bottom: 10px;
}

.app-footer {
  text-align: center;
  margin-top: 30px;
  padding: 20px;
  color: #7f8c8d;
  font-size: 0.9rem;
}
</style>
