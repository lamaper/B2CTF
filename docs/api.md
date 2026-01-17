# B2CTF API 文档

## 概述
本文档详细描述B2CTF平台的API接口，包括请求方法、URL、请求体、响应格式等信息。

## 基本信息

- **基础URL**：`http://localhost:8080/api`
- **Content-Type**：`application/json`
- **认证方式**：JWT Token（通过Authorization头传递）

## 错误响应格式

```json
{
  "code": 错误码,
  "msg": "错误信息"
}
```

## 已实现的API

### 1. 用户注册接口

**请求**
- **方法**：`POST`
- **URL**：`/register`
- **请求体**：
  ```json
  {
    "username": "testuser",
    "password": "password123",
    "email": "test@example.com"
  }
  ```

**响应**
- **成功**（201 Created）：
  ```json
  {
    "code": 201,
    "msg": "用户注册成功"
  }
  ```
- **失败**：
  ```json
  {
    "code": 400,
    "msg": "用户名已存在"
  }
  ```
  或
  ```json
  {
    "code": 500,
    "msg": "用户注册失败"
  }
  ```

### 2. 用户登录接口

**请求**
- **方法**：`POST`
- **URL**：`/login`
- **请求体**：
  ```json
  {
    "username": "testuser",
    "password": "password123"
  }
  ```

**响应**
- **成功**（200 OK）：
  ```json
  {
    "code": 200,
    "msg": "登录成功",
    "data": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
  }
  ```
- **失败**：
  ```json
  {
    "code": 401,
    "msg": "用户名或密码错误"
  }
  ```
  或
  ```json
  {
    "code": 500,
    "msg": "登录失败"
  }
  ```

### 3. 测试接口

**请求**
- **方法**：`GET`
- **URL**：`/ping`

**响应**
- **成功**（200 OK）：
  ```json
  {
    "msg": "pong"
  }
  ```

## 待实现的API

### 1. 题目相关API

- **获取所有题目**：`GET /challenges`
- **获取单个题目**：`GET /challenges/:id`
- **提交Flag**：`POST /challenges/:id/submit`
- **获取题目分类**：`GET /challenges/categories`

### 2. 排行榜API

- **获取排行榜**：`GET /scoreboard`

### 3. 用户相关API

- **获取当前用户信息**：`GET /user/me`
- **更新用户信息**：`PUT /user/me`
- **获取用户提交记录**：`GET /user/submissions`

### 4. 管理员API

- **添加题目**：`POST /admin/challenges`
- **更新题目**：`PUT /admin/challenges/:id`
- **删除题目**：`DELETE /admin/challenges/:id`
- **获取所有用户**：`GET /admin/users`
- **管理用户权限**：`PUT /admin/users/:id/role`

## API认证

需要认证的API需要在请求头中添加Authorization字段：

```
Authorization: Bearer <token>
```

其中`<token>`是登录成功后返回的JWT令牌。

## 注意事项

1. 所有API都支持跨域请求（CORS）
2. 敏感操作需要JWT认证
3. 密码在传输和存储时都会进行加密处理
4. 输入数据会进行严格的验证
