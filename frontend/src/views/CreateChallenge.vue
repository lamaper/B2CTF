<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import http from '../api/http';

const router = useRouter();
const loading = ref(false);
const error = ref('');
const success = ref('');
const competitions = ref([]);

const form = ref({
  competition_id: '',
  title: '',
  category: '',
  description: '',
  score: 100,
  flag: '',
  attachment: '',
  tags: ''
});

// 获取比赛列表
const fetchCompetitions = async () => {
  try {
    const response = await http.get('/competitions');
    competitions.value = response.data || [];
  } catch (err) {
    console.error('获取比赛列表失败:', err);
  }
};

const handleSubmit = async () => {
  if (!form.value.competition_id || !form.value.title || !form.value.category || !form.value.description || !form.value.flag) {
    error.value = '请填写必填字段';
    return;
  }

  loading.value = true;
  error.value = '';
  success.value = '';

  try {
    // 处理标签
    const tags = form.value.tags.split(',').map(tag => tag.trim()).filter(tag => tag);
    
    const challengeData = {
      competition_id: parseInt(form.value.competition_id),
      title: form.value.title,
      category: form.value.category,
      description: form.value.description,
      score: form.value.score,
      flag: form.value.flag,
      attachment: form.value.attachment,
      tags: tags
    };

    const response = await http.post('/challenge', challengeData);
    
    if (response.code === 200 || response === undefined) {
      success.value = '题目创建成功';
      // 重置表单
      form.value = {
        competition_id: '',
        title: '',
        category: '',
        description: '',
        score: 100,
        flag: '',
        attachment: '',
        tags: ''
      };
    } else {
      error.value = response.msg || '创建题目失败';
    }
  } catch (err) {
    error.value = '创建题目失败，请稍后重试';
    console.error(err);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchCompetitions();
});
</script>

<template>
  <div class="container">
    <h1>创建题目</h1>
    
    <div v-if="success" class="message message-success">
      {{ success }}
    </div>
    
    <div v-if="error" class="message message-error">
      {{ error }}
    </div>
    
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="competition_id">所属比赛</label>
        <select id="competition_id" v-model="form.competition_id" required>
          <option value="">请选择比赛</option>
          <option v-for="competition in competitions" :key="competition.id" :value="competition.id">
            {{ competition.title }} ({{ competition.type === 0 ? '限时比赛' : '常驻题库' }})
          </option>
        </select>
      </div>
      
      <div class="form-group">
        <label for="title">题目标题</label>
        <input type="text" id="title" v-model="form.title" required>
      </div>
      
      <div class="form-group">
        <label for="category">题目分类</label>
        <input type="text" id="category" v-model="form.category" required>
      </div>
      
      <div class="form-group">
        <label for="description">题目描述</label>
        <textarea id="description" v-model="form.description" required></textarea>
      </div>
      
      <div class="form-group">
        <label for="score">题目分值</label>
        <input type="number" id="score" v-model.number="form.score" min="1" required>
      </div>
      
      <div class="form-group">
        <label for="flag">Flag</label>
        <input type="text" id="flag" v-model="form.flag" required>
      </div>
      
      <div class="form-group">
        <label for="attachment">附件链接</label>
        <input type="text" id="attachment" v-model="form.attachment">
      </div>
      
      <div class="form-group">
        <label for="tags">标签 (用逗号分隔)</label>
        <input type="text" id="tags" v-model="form.tags">
      </div>
      
      <button type="submit" class="btn btn-primary" :disabled="loading">
        {{ loading ? '创建中...' : '创建题目' }}
      </button>
    </form>
  </div>
</template>