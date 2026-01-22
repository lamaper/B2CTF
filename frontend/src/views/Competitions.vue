<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import http from '../api/http';

const router = useRouter();

const competitions = ref([]);
const loading = ref(true);
const error = ref('');

const fetchCompetitions = async () => {
  try {
    const response = await http.get('/competitions');
    competitions.value = response.data || [];
  } catch (err) {
    error.value = '获取比赛列表失败';
    console.error(err);
  } finally {
    loading.value = false;
  }
};

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleString();
};

const getCompetitionStatus = (startTime, endTime) => {
  const now = new Date();
  const start = new Date(startTime);
  const end = new Date(endTime);
  
  if (now < start) {
    return { status: 'upcoming', text: '即将开始' };
  } else if (now >= start && now <= end) {
    return { status: 'ongoing', text: '进行中' };
  } else {
    return { status: 'ended', text: '已结束' };
  }
};

onMounted(() => {
  fetchCompetitions();
});
</script>

<template>
  <div class="competitions-container">
    <div class="competitions-header">
      <h1>比赛列表</h1>
      <button class="create-button">创建比赛</button>
    </div>
    
    <div class="competitions-content">
      <div v-if="loading" class="loading">
        加载中...
      </div>
      
      <div v-else-if="error" class="error-message">
        {{ error }}
      </div>
      
      <div v-else-if="competitions.length === 0" class="empty-state">
        <h3>暂无比赛</h3>
        <p>系统中还没有创建任何比赛</p>
      </div>
      
      <div v-else class="competitions-grid">
        <div v-for="competition in competitions" :key="competition.id" class="competition-card">
          <div class="competition-header">
            <h3>{{ competition.title }}</h3>
            <span 
              class="competition-status" 
              :class="getCompetitionStatus(competition.start_time, competition.end_time).status"
            >
              {{ getCompetitionStatus(competition.start_time, competition.end_time).text }}
            </span>
          </div>
          
          <div class="competition-body">
            <p class="competition-desc">{{ competition.description }}</p>
            
            <div class="competition-time">
              <div class="time-item">
                <span class="label">开始时间：</span>
                <span class="value">{{ formatDate(competition.start_time) }}</span>
              </div>
              <div class="time-item">
                <span class="label">结束时间：</span>
                <span class="value">{{ formatDate(competition.end_time) }}</span>
              </div>
            </div>
            
            <div class="competition-type">
              <span class="label">类型：</span>
              <span class="value">{{ competition.type === 0 ? '限时比赛' : '常驻题库' }}</span>
            </div>
          </div>
          
          <div class="competition-footer">
            <button class="view-button">查看详情</button>
            <button class="join-button">参加比赛</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.competitions-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 40px 20px;
}

.competitions-header {
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

.competitions-header h1 {
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

.competitions-content {
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

.competitions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
  gap: 25px;
}

.competition-card {
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  padding: 25px;
  transition: all 0.3s ease;
}

.competition-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.15);
}

.competition-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 15px;
}

.competition-header h3 {
  font-size: 1.2rem;
  color: #2c3e50;
  margin: 0;
  flex: 1;
}

.competition-status {
  padding: 4px 10px;
  border-radius: 12px;
  font-size: 0.8rem;
  font-weight: 500;
}

.competition-status.upcoming {
  background-color: #3498db;
  color: white;
}

.competition-status.ongoing {
  background-color: #27ae60;
  color: white;
}

.competition-status.ended {
  background-color: #95a5a6;
  color: white;
}

.competition-body {
  margin-bottom: 20px;
}

.competition-desc {
  color: #666;
  line-height: 1.5;
  margin-bottom: 15px;
  font-size: 0.95rem;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.competition-time {
  margin-bottom: 10px;
}

.time-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 5px;
  font-size: 0.9rem;
}

.time-item .label {
  color: #888;
}

.time-item .value {
  color: #333;
  font-weight: 500;
}

.competition-type {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.9rem;
  margin-top: 10px;
}

.competition-type .label {
  color: #888;
}

.competition-type .value {
  color: #333;
  font-weight: 500;
}

.competition-footer {
  display: flex;
  gap: 10px;
}

.view-button, .join-button {
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

.join-button {
  background-color: #e67e22;
  color: white;
}

.join-button:hover {
  background-color: #d35400;
}

@media (max-width: 768px) {
  .competitions-header {
    flex-direction: column;
    gap: 15px;
    text-align: center;
  }
  
  .competitions-grid {
    grid-template-columns: 1fr;
  }
  
  .competition-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 10px;
  }
  
  .time-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 5px;
  }
  
  .competition-type {
    flex-direction: column;
    align-items: flex-start;
    gap: 5px;
  }
  
  .competition-footer {
    flex-direction: column;
  }
}
</style>