import axios from 'axios';

// 创建axios实例
const http = axios.create({
  baseURL: 'http://localhost:8080/api', // 后端API基础URL
  timeout: 5000, // 请求超时时间
  headers: {
    'Content-Type': 'application/json'
  }
});

// 请求拦截器
http.interceptors.request.use(
  config => {
    // 可以在这里添加token等认证信息
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// 响应拦截器
http.interceptors.response.use(
  response => {
    // 统一处理响应数据
    return response.data;
  },
  error => {
    // 统一处理错误
    console.error('API请求错误:', error);
    return Promise.reject(error);
  }
);

export default http;
