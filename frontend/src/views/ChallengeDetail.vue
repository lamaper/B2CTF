<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import http from '../api/http';

const route = useRoute();
const challengeId = route.params.id;
const challenge = ref(null);
const loading = ref(true);
const error = ref('');
const flag = ref('');
const submitLoading = ref(false);
const submitSuccess = ref('');
const submitError = ref('');

// 获取题目详情
const fetchChallengeDetail = async () => {
  try {
    // 这里假设后端有获取题目详情的接口
    // 暂时使用获取所有题目的接口，然后过滤出当前题目
    const response = await http.get('/challenge');
    if (Array.isArray(response)) {
      const foundChallenge = response.find(c => c.id === parseInt(challengeId));
      if (foundChallenge) {
        challenge.value = foundChallenge;
      } else {
        error.value = '题目不存在';
      }
    }
  } catch (err) {
    error.value = '获取题目详情失败';
    console.error(err);
  } finally {
    loading.value = false;
  }
};

// 提交flag
const submitFlag = async () => {
  if (!flag.value) {
    submitError.value = '请输入flag';
    return;
  }

  submitLoading.value = true;
  submitSuccess.value = '';
  submitError.value = '';

  try {
    // 这里假设后端有提交flag的接口
    // 暂时使用一个模拟的接口
    const response = await http.post('/submit', {
      challenge_id: challengeId,
      flag: flag.value
    });
    
    if (response.code === 200) {
      submitSuccess.value = 'Flag提交成功！';
      flag.value = '';
      // 重新获取题目详情，更新解题状态
      fetchChallengeDetail();
    } else {
      submitError.value = response.msg || 'Flag提交失败';
    }
  } catch (err) {
    submitError.value = 'Flag提交失败，请稍后重试';
    console.error(err);
  } finally {
    submitLoading.value = false;
  }
};

onMounted(() => {
  fetchChallengeDetail();
});
</script>

<template>
  <div class="container">
    <h1>题目详情</h1>
    
    <div v-if="loading" class="loading">
      加载中...
    </div>
    
    <div v-else-if="error" class="message message-error">
      {{ error }}
    </div>
    
    <div v-else-if="challenge" class="challenge-detail">
      <div class="card">
        <div class="card-header">
          <h2>{{ challenge.title }}</h2>
          <div class="challenge-meta">
            <span class="category">{{ challenge.category }}</span>
            <span class="score">{{ challenge.score }} 分</span>
            <span class="solves">{{ challenge.solved_count || 0 }} 人解决</span>
          </div>
        </div>
        
        <div class="card-body">
          <div class="description">
            <h3>题目描述</h3>
            <p>{{ challenge.description }}</p>
          </div>
          
          <div v-if="challenge.attachment" class="attachment">
            <h3>附件</h3>
            <a :href="challenge.attachment" target="_blank">下载附件</a>
          </div>
          
          <div v-if="challenge.tags && challenge.tags.length > 0" class="tags">
            <h3>标签</h3>
            <div class="tag-list">
              <span v-for="tag in challenge.tags" :key="tag" class="tag">
                {{ tag }}
              </span>
            </div>
          </div>
        </div>
        
        <div class="card-footer">
          <h3>提交Flag</h3>
          <div class="flag-submit">
            <div v-if="submitSuccess" class="message message-success">
              {{ submitSuccess }}
            </div>
            <div v-if="submitError" class="message message-error">
              {{ submitError }}
            </div>
            <form @submit.prevent="submitFlag">
              <div class="form-group">
                <input type="text" v-model="flag" placeholder="请输入flag" required>
                <button type="submit" class="btn btn-primary" :disabled="submitLoading">
                  {{ submitLoading ? '提交中...' : '提交' }}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.challenge-detail {
  max-width: 800px;
  margin: 0 auto;
}

.challenge-meta {
  display: flex;
  gap: 15px;
  margin-top: 10px;
}

.category {
  background-color: #3498db;
  color: white;
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 12px;
}

.score {
  font-weight: bold;
  color: #e74c3c;
}

.solves {
  color: #7f8c8d;
}

.description {
  margin-bottom: 20px;
}

.attachment {
  margin-bottom: 20px;
}

.attachment a {
  color: #3498db;
  text-decoration: none;
}

.attachment a:hover {
  text-decoration: underline;
}

.tags {
  margin-bottom: 20px;
}

.tag-list {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.tag {
  background-color: #f1c40f;
  color: #333;
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 12px;
}

.flag-submit {
  margin-top: 20px;
}

.flag-submit .form-group {
  display: flex;
  gap: 10px;
}

.flag-submit input {
  flex: 1;
}
</style>