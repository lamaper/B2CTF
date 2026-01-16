import http from './http';

// 调用后端ping接口
export const ping = () => {
  return http.get('/ping');
};
