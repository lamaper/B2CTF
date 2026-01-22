<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import http from '../api/http';

const router = useRouter();
const loading = ref(false);
const error = ref('');
const success = ref('');

const form = ref({
  title: '',
  description: '',
  type: 0, // 0: 限时比赛, 1: 常驻题库
  startTime: new Date().toISOString().slice(0, 16),
  endTime: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000).toISOString().slice(0, 16)
});

const handleSubmit = async () => {
  if (!form.value.title || !form.value.description) {
    error.value = '请填写标题和描述';
    return;
  }

  loading.value = true;
  error.value = '';
  success.value = '';

  try {
    // 转换时间格式
    const competitionData = {
      title: form.value.title,
      description: form.value.description,
      type: form.value.type,
      start_time: new Date(form.value.startTime).toISOString(),
      end_time: new Date(form.value.endTime).toISOString()
    };

    const response = await http.post('/competitions', competitionData);
    
    if (response.code === 200) {
      success.value = '比赛创建成功';
      // 重置表单
      form.value = {
        title: '',
        description: '',
        type: 0,
        startTime: new Date().toISOString().slice(0, 16),
        endTime: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000).toISOString().slice(0, 16),
        mode: 0
      };
      // 3秒后跳转到比赛列表
      setTimeout(() => {
        router.push('/competitions');
      }, 3000);
    } else {
      error.value = response.msg || '创建比赛失败';
    }
  } catch (err) {
    error.value = '创建比赛失败，请稍后重试';
    console.error(err);
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div class="container">
    <h1>创建比赛</h1>
    
    <div v-if="success" class="message message-success">
      {{ success }}
    </div>
    
    <div v-if="error" class="message message-error">
      {{ error }}
    </div>
    
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="title">比赛标题</label>
        <input type="text" id="title" v-model="form.title" required>
      </div>
      
      <div class="form-group">
        <label for="description">比赛描述</label>
        <textarea id="description" v-model="form.description" required></textarea>
      </div>
      
      <div class="form-group">
        <label for="type">比赛类型</label>
        <select id="type" v-model="form.type">
          <option value="0">限时比赛</option>
          <option value="1">常驻题库</option>
        </select>
      </div>
      
      <div class="form-group">
        <label for="startTime">开始时间</label>
        <input type="datetime-local" id="startTime" v-model="form.startTime" required>
      </div>
      
      <div class="form-group">
        <label for="endTime">结束时间</label>
        <input type="datetime-local" id="endTime" v-model="form.endTime" required>
      </div>
      

      
      <button type="submit" class="btn btn-primary" :disabled="loading">
        {{ loading ? '创建中...' : '创建比赛' }}
      </button>
    </form>
  </div>
</template>