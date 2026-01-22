<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import http from '../api/http';

const router = useRouter();

const challenges = ref([]);
const loading = ref(true);
const error = ref('');

const fetchChallenges = async () => {
  try {
    const response = await http.get('/challenge');
    challenges.value = response.data || [];
  } catch (err) {
    error.value = '获取题目列表失败';
    console.error(err);
  } finally {
    loading.value = false;
  }
};

const viewChallenge = (challengeId) => {
  router.push(`/challenge/${challengeId}`);
};

onMounted(() => {
  fetchChallenges();
});
</script>

<template>
  <div class="challenges-container">
    <div class="challenges-header">
      <h1>题目列表</h1>
      <button class="create-button">创建题目</button>
    </div>
    
    <div class="challenges-content">
      <div v-if="loading" class="loading">
        加载中...
      </div>
      
      <div v-else-if="error" class="error-message">
        {{ error }}
      </div>
      
      <div v-else-if="challenges.length === 0" class="empty-state">
        <h3>暂无题目</h3>
        <p>系统中还没有添加任何题目</p>
      </div>
      
      <div v-else class="challenges-grid">
        <div v-for="challenge in challenges" :key="challenge.id" class="challenge-card">
          <div class="challenge-header">
            <h3>{{ challenge.title }}</h3>
            <span class="challenge-category">{{ challenge.category }}</span>
          </div>
          
          <div class="challenge-body">
            <p class="challenge-desc">{{ challenge.description }}</p>
            <div class="challenge-meta">
              <span class="challenge-score">{{ challenge.score }} 分</span>
              <span class="challenge-solves">{{ challenge.solves || 0 }} 人解决</span>
            </div>
          </div>
          
          <div class="challenge-footer">
            <button class="view-button" @click="viewChallenge(challenge.id)">查看详情</button>
            <button class="submit-button" @click="viewChallenge(challenge.id)">提交Flag</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.challenges-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
  padding: 40px 20px;
}

.challenges-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  max-width: 1200px;
  margin: 0 auto 40px;
  background-color: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
}

.challenges-header h1 {
  color: white;
  font-size: 1.8rem;
  margin: 0;
}

.create-button {
  background-color: #27ae60;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 0.9rem;
  font-weight: 500;
}

.create-button:hover {
  background-color: #229954;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(39, 174, 96, 0.3);
}

.challenges-content {
  max-width: 1200px;
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

.empty-state {
  text-align: center;
  color: white;
  font-size: 1.2rem;
  padding: 60px;
  background-color: rgba(255, 255, 255, 0.1);
  border-radius: 8px;
}

.empty-state h3 {
  margin-bottom: 10px;
}

.challenges-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
  gap: 25px;
}

.challenge-card {
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  padding: 25px;
  transition: all 0.3s ease;
}

.challenge-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.15);
}

.challenge-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 15px;
}

.challenge-header h3 {
  font-size: 1.2rem;
  color: #2c3e50;
  margin: 0;
  flex: 1;
}

.challenge-category {
  background-color: #3498db;
  color: white;
  padding: 4px 10px;
  border-radius: 12px;
  font-size: 0.8rem;
  font-weight: 500;
}

.challenge-body {
  margin-bottom: 20px;
}

.challenge-desc {
  color: #666;
  line-height: 1.5;
  margin-bottom: 15px;
  font-size: 0.95rem;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.challenge-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.challenge-score {
  font-weight: bold;
  color: #e74c3c;
  font-size: 1rem;
}

.challenge-solves {
  color: #7f8c8d;
  font-size: 0.9rem;
}

.challenge-footer {
  display: flex;
  gap: 10px;
}

.view-button, .submit-button {
  flex: 1;
  padding: 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 0.9rem;
  font-weight: 500;
}

.view-button {
  background-color: #95a5a6;
  color: white;
}

.view-button:hover {
  background-color: #7f8c8d;
}

.submit-button {
  background-color: #3498db;
  color: white;
}

.submit-button:hover {
  background-color: #2980b9;
}

@media (max-width: 768px) {
  .challenges-header {
    flex-direction: column;
    gap: 15px;
    text-align: center;
  }
  
  .challenges-grid {
    grid-template-columns: 1fr;
  }
  
  .challenge-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 10px;
  }
  
  .challenge-footer {
    flex-direction: column;
  }
}
</style>