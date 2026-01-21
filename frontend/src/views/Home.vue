<script setup>
import { ref } from 'vue';
import { ping } from '../api/ping';

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
  <div class="home-container">
    <header class="home-header">
      <h1>B2CTF 平台</h1>
      <p>CTF 竞赛平台 - 挑战你的安全技能</p>
    </header>
    
    <main class="home-main">
      <div class="hero-section">
        <h2>欢迎来到 B2CTF</h2>
        <p>这是一个专为网络安全爱好者打造的CTF竞赛平台，提供各种类型的安全挑战。</p>
        <div class="hero-buttons">
          <router-link to="/register" class="btn btn-primary">立即注册</router-link>
          <router-link to="/login" class="btn btn-secondary">登录</router-link>
        </div>
      </div>
      
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
      
      <div class="features-section">
        <h2>平台特性</h2>
        <div class="features-grid">
          <div class="feature-card">
            <h3>多种题型</h3>
            <p>包含Web、Pwn、Reverse、Crypto等多种类型的题目</p>
          </div>
          <div class="feature-card">
            <h3>实时排行榜</h3>
            <p>实时更新比赛排名，展示玩家实力</p>
          </div>
          <div class="feature-card">
            <h3>文件上传</h3>
            <p>支持题目附件上传和管理</p>
          </div>
          <div class="feature-card">
            <h3>比赛管理</h3>
            <p>创建和管理CTF比赛，设置时间和规则</p>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<style scoped>
.home-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.home-header {
  text-align: center;
  padding: 60px 20px;
  background: linear-gradient(135deg, #2c3e50 0%, #34495e 100%);
  color: white;
}

.home-header h1 {
  font-size: 3rem;
  margin-bottom: 10px;
  font-weight: bold;
}

.home-header p {
  font-size: 1.2rem;
  opacity: 0.9;
}

.home-main {
  flex: 1;
  padding: 40px 20px;
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
}

.hero-section {
  text-align: center;
  margin-bottom: 60px;
  padding: 40px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.hero-section h2 {
  font-size: 2rem;
  margin-bottom: 20px;
  color: #2c3e50;
}

.hero-section p {
  font-size: 1.1rem;
  margin-bottom: 30px;
  color: #7f8c8d;
  max-width: 600px;
  margin-left: auto;
  margin-right: auto;
}

.hero-buttons {
  display: flex;
  gap: 20px;
  justify-content: center;
  flex-wrap: wrap;
}

.btn {
  padding: 12px 24px;
  font-size: 1rem;
  border-radius: 4px;
  text-decoration: none;
  font-weight: bold;
  transition: all 0.3s ease;
  border: none;
  cursor: pointer;
}

.btn-primary {
  background-color: #3498db;
  color: white;
}

.btn-primary:hover {
  background-color: #2980b9;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(52, 152, 219, 0.3);
}

.btn-secondary {
  background-color: #95a5a6;
  color: white;
}

.btn-secondary:hover {
  background-color: #7f8c8d;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(149, 165, 166, 0.3);
}

.test-section {
  margin-bottom: 60px;
  padding: 30px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.test-section h2 {
  font-size: 1.5rem;
  margin-bottom: 20px;
  color: #2c3e50;
}

.ping-button {
  background-color: #27ae60;
  color: white;
  border: none;
  padding: 12px 24px;
  font-size: 1rem;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.ping-button:hover:not(:disabled) {
  background-color: #229954;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(39, 174, 96, 0.3);
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

.features-section {
  margin-bottom: 40px;
}

.features-section h2 {
  font-size: 1.8rem;
  margin-bottom: 30px;
  text-align: center;
  color: #2c3e50;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
}

.feature-card {
  background-color: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}

.feature-card h3 {
  font-size: 1.2rem;
  margin-bottom: 15px;
  color: #3498db;
}

.feature-card p {
  color: #7f8c8d;
  line-height: 1.5;
}

@media (max-width: 768px) {
  .home-header h1 {
    font-size: 2.5rem;
  }
  
  .hero-buttons {
    flex-direction: column;
    align-items: center;
  }
  
  .btn {
    width: 100%;
    max-width: 200px;
    text-align: center;
  }
}
</style>