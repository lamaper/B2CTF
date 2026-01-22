<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import http from '../api/http';

const router = useRouter();

const userProfile = ref(null);
const loading = ref(true);
const error = ref('');

const fetchProfile = async () => {
  try {
    const response = await http.get('/user/profile');
    userProfile.value = response.data;
  } catch (err) {
    error.value = '获取个人资料失败';
    console.error(err);
  } finally {
    loading.value = false;
  }
};

const logout = () => {
  localStorage.removeItem('token');
  delete http.defaults.headers.common['Authorization'];
  router.push('/login');
};

onMounted(() => {
  fetchProfile();
});
</script>

<template>
  <div class="profile-container">
    <div class="profile-header">
      <h1>个人资料</h1>
      <button class="logout-button" @click="logout">退出登录</button>
    </div>
    
    <div class="profile-content">
      <div v-if="loading" class="loading">
        加载中...
      </div>
      
      <div v-else-if="error" class="error-message">
        {{ error }}
      </div>
      
      <div v-else-if="userProfile" class="profile-card">
        <div class="profile-avatar">
          <div class="avatar">
            {{ userProfile.username?.charAt(0).toUpperCase() || 'U' }}
          </div>
        </div>
        
        <div class="profile-info">
          <h2>{{ userProfile.username }}</h2>
          <div class="info-item">
            <span class="label">邮箱：</span>
            <span class="value">{{ userProfile.email }}</span>
          </div>
          <div class="info-item">
            <span class="label">角色：</span>
            <span class="value">{{ userProfile.role === 'admin' ? '管理员' : '普通用户' }}</span>
          </div>
          <div class="info-item">
            <span class="label">注册时间：</span>
            <span class="value">{{ new Date(userProfile.created_at).toLocaleString() }}</span>
          </div>
        </div>
        
        <div class="profile-stats">
          <div class="stat-card">
            <h3>已解决题目</h3>
            <p class="stat-value">0</p>
          </div>
          <div class="stat-card">
            <h3>积分</h3>
            <p class="stat-value">0</p>
          </div>
          <div class="stat-card">
            <h3>排名</h3>
            <p class="stat-value">--</p>
          </div>
        </div>
      </div>
      
      <div v-else class="empty-state">
        <p>未获取到个人资料</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.profile-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 40px 20px;
}

.profile-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  max-width: 1000px;
  margin: 0 auto 40px;
  background-color: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
}

.profile-header h1 {
  color: white;
  font-size: 1.8rem;
  margin: 0;
}

.logout-button {
  background-color: rgba(255, 255, 255, 0.2);
  color: white;
  border: 1px solid rgba(255, 255, 255, 0.3);
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 0.9rem;
}

.logout-button:hover {
  background-color: rgba(255, 255, 255, 0.3);
  transform: translateY(-2px);
}

.profile-content {
  max-width: 1000px;
  margin: 0 auto;
}

.loading {
  text-align: center;
  color: white;
  font-size: 1.2rem;
  padding: 40px;
  background-color: rgba(255, 255, 255, 0.1);
  border-radius: 8px;
}

.error-message {
  text-align: center;
  color: #ff6b6b;
  font-size: 1.1rem;
  padding: 40px;
  background-color: rgba(255, 255, 255, 0.1);
  border-radius: 8px;
}

.profile-card {
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  padding: 40px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 30px;
}

.profile-avatar {
  margin-bottom: 20px;
}

.avatar {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 48px;
  font-weight: bold;
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.3);
}

.profile-info {
  text-align: center;
  width: 100%;
}

.profile-info h2 {
  font-size: 2rem;
  margin-bottom: 20px;
  color: #333;
}

.info-item {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 10px;
  font-size: 1.1rem;
}

.label {
  font-weight: 500;
  color: #666;
  margin-right: 10px;
}

.value {
  color: #333;
}

.profile-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
  width: 100%;
  margin-top: 20px;
}

.stat-card {
  background-color: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  text-align: center;
  transition: all 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
}

.stat-card h3 {
  font-size: 1rem;
  color: #666;
  margin-bottom: 10px;
}

.stat-value {
  font-size: 2rem;
  font-weight: bold;
  color: #667eea;
  margin: 0;
}

.empty-state {
  text-align: center;
  color: white;
  font-size: 1.2rem;
  padding: 40px;
  background-color: rgba(255, 255, 255, 0.1);
  border-radius: 8px;
}

@media (max-width: 768px) {
  .profile-header {
    flex-direction: column;
    gap: 15px;
    text-align: center;
  }
  
  .profile-card {
    padding: 30px 20px;
  }
  
  .info-item {
    flex-direction: column;
    gap: 5px;
  }
  
  .profile-stats {
    grid-template-columns: 1fr;
  }
}
</style>