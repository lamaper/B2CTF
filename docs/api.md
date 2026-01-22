# B2CTF API 文档

## 概述
本文档详细描述B2CTF平台的API接口，包括请求方法、URL、请求体、响应格式等信息。

## 基本信息

- **基础URL**：`http://localhost:8080`
- **Content-Type**：`application/json`（除文件上传接口外）
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
- **URL**：`/api/register`
- **请求体**：
  ```json
  {
    "username": "testuser",
    "password": "password123",
    "email": "test@example.com"
  }
  ```

**响应**
- **成功**（200 OK）：
  ```json
  {
    "code": 200,
    "msg": "注册成功"
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
    "code": 400,
    "msg": "邮箱已存在"
  }
  ```

### 2. 用户登录接口

**请求**
- **方法**：`POST`
- **URL**：`/api/login`
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
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "role": "user"
    }
  }
  ```
- **失败**：
  ```json
  {
    "error": "用户名或密码错误"
  }
  ```

### 3. 获取用户信息接口

**请求**
- **方法**：`GET`
- **URL**：`/user/profile`
- **认证**：需要JWT Token

**响应**
- **成功**（200 OK）：
  ```json
  {
    "code": 200,
    "msg": "获取成功",
    "data": {
      "user_id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "role": "user",
      "created_at": "2026-01-17T12:00:00Z"
    }
  }
  ```
- **失败**：
  ```json
  {
    "error": "无法获取用户信息"
  }
  ```

### 4. 创建比赛接口

**请求**
- **方法**：`POST`
- **URL**：`/competitions`
- **认证**：需要JWT Token（管理员权限）
- **请求体**：
  ```json
  {
    "title": "CTF比赛2026",
    "description": "新年第一场CTF比赛",
    "type": 0,
    "start_time": "2026-01-20T15:04:05Z",
    "end_time": "2026-01-21T15:04:05Z"
  }
  ```

**响应**
- **成功**（200 OK）：
  ```json
  {
    "code": 200,
    "msg": "比赛创建成功"
  }
  ```
- **失败**：
  ```json
  {
    "error": "只有管理员可以创建比赛"
  }
  ```
  或
  ```json
  {
    "error": "参数错误: invalid character '}' looking for beginning of object key string"
  }
  ```

### 5. 获取比赛列表接口

**请求**
- **方法**：`GET`
- **URL**：`/competitions`
- **认证**：需要JWT Token

**响应**
- **成功**（200 OK）：
  ```json
  {
    "code": 200,
    "msg": "获取成功",
    "data": [
      {
        "id": 1,
        "title": "CTF比赛2026",
        "description": "新年第一场CTF比赛",
        "type": 0,
        "start_time": "2026-01-20T15:04:05Z",
        "end_time": "2026-01-21T15:04:05Z",
        "mode": 0
      }
    ]
  }
  ```
- **失败**：
  ```json
  {
    "error": "获取列表失败"
  }
  ```

### 6. 创建题目接口

**请求**
- **方法**：`POST`
- **URL**：`/challenge`
- **认证**：需要JWT Token（管理员权限）
- **请求体**：
  ```json
  {
    "title": "测试题目",
    "category": "Web",
    "description": "这是一个测试题目",
    "flag": "flag{test123}",
    "score": 100,
    "competition_id": 1,
    "attachment": "/uploads/test.zip",
    "tags": ["Web", "Easy"]
  }
  ```

**响应**
- **成功**（200 OK）：
  ```json
  {
    "code": 200,
    "msg": "题目创建成功"
  }
  ```
- **失败**：
  ```json
  {
    "error": "权限不足"
  }
  ```
  或
  ```json
  {
    "error": "创建失败: 错误信息"
  }
  ```

### 7. 获取题目列表接口

**请求**
- **方法**：`GET`
- **URL**：`/challenge`
- **认证**：需要JWT Token

**响应**
- **成功**（200 OK）：
  ```json
  {
    "code": 200,
    "msg": "获取成功",
    "data": [
      {
        "id": 1,
        "title": "测试题目",
        "category": "Web",
        "description": "这是一个测试题目",
        "score": 100,
        "competition_id": 1,
        "attachment": "/uploads/test.zip",
        "tags": ["Web", "Easy"],
        "solved_count": 0
      }
    ]
  }
  ```
- **失败**：
  ```json
  {
    "error": "获取列表失败"
  }
  ```

### 8. 文件上传接口

**请求**
- **方法**：`POST`
- **URL**：`/upload`
- **Content-Type**：`multipart/form-data`
- **认证**：需要JWT Token
- **表单数据**：
  - `file`：文件（必填）

**响应**
- **成功**（200 OK）：
  ```json
  {
    "code": 200,
    "msg": "上传成功",
    "data": {
      "url": "/uploads/test.jpg"
    }
  }
  ```
- **失败**：
  ```json
  {
    "error": "请上传文件"
  }
  ```
  或
  ```json
  {
    "error": "上传失败"
  }
  ```

### 9. 创建团队接口

**请求**
- **方法**：`POST`
- **URL**：`/team/create`
- **认证**：需要JWT Token
- **请求体**：
  ```json
  {
    "name": "测试团队",
    "description": "这是一个测试团队"
  }
  ```

**响应**
- **成功**（200 OK）：
  ```json
  {
    "code": 200,
    "msg": "团队创建成功",
    "data": {
      "team_id": 1,
      "token": "team_token_123"
    }
  }
  ```
- **失败**：
  ```json
  {
    "error": "创建团队失败"
  }
  ```

### 10. 加入团队接口

**请求**
- **方法**：`POST`
- **URL**：`/team/join`
- **认证**：需要JWT Token
- **请求体**：
  ```json
  {
    "token": "team_token_123"
  }
  ```

**响应**
- **成功**（200 OK）：
  ```json
  {
    "code": 200,
    "msg": "加入团队成功"
  }
  ```
- **失败**：
  ```json
  {
    "error": "加入团队失败"
  }
  ```

### 11. 获取我的团队接口

**请求**
- **方法**：`GET`
- **URL**：`/team/my`
- **认证**：需要JWT Token

**响应**
- **成功**（200 OK）：
  ```json
  {
    "code": 200,
    "msg": "获取成功",
    "data": {
      "id": 1,
      "name": "测试团队",
      "description": "这是一个测试团队",
      "token": "team_token_123",
      "captain_id": 1,
      "members": [
        {
          "id": 1,
          "username": "testuser",
          "email": "test@example.com",
          "role": "user",
          "score": 0,
          "avatar": "",
          "team_id": 1
        }
      ]
    }
  }
  ```
- **失败**：
  ```json
  {
    "error": "获取团队失败"
  }
  ```

## 待实现的API

### 1. 排行榜API

- **获取排行榜**：`GET /rank`

### 2. Flag提交接口

- **提交Flag**：`POST /submit`

### 3. 用户相关API

- **更新用户头像**：`POST /user/avatar`

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
5. 管理员接口需要管理员权限才能访问
