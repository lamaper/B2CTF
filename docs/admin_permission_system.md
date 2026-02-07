# B2CTF 管理员权限系统文档

## 概述

为了完善后端的权限管理，我们实现了完整的管理员权限系统。本文档描述了新的权限体系、中间件、路由分离和管理员操作的详细内容。

---

## 1. 权限中间件

### 1.1 JWTAuth 中间件（现有）

位置：`backend/internal/middleware/auth.go`

功能：
- 验证请求头中的 JWT Token
- 提取 userID 和 role，存入 Context
- 拦截无效或过期的 Token

```go
protected.Use(middleware.JWTAuth())
```

### 1.2 AdminOnly 中间件（新增）

位置：`backend/internal/middleware/permission.go`

功能：
- 检查当前用户是否为管理员
- 拒绝非管理员的请求
- 返回统一的权限拒绝响应

```go
admin.Use(middleware.AdminOnly())
```

### 1.3 辅助函数

**GetUserID(c *gin.Context) (uint, bool)**
- 安全地从 Context 获取用户 ID
- 处理 float64 类型转换

**GetUserRole(c *gin.Context) (string, bool)**
- 安全地从 Context 获取用户角色

**IsAdmin(c *gin.Context) bool**
- 快速检查用户是否为管理员

---

## 2. 路由分离

### 2.1 路由结构

```
/api                           (公共接口，无需认证)
  POST /register
  POST /login
  POST /ping

/                              (需要认证的接口)
  GET  /user/profile
  GET  /challenge
  GET  /competitions
  GET  /rank
  GET  /rank/:id
  POST /submit
  POST /container/launch
  POST /container/terminate
  
  /team                        (团队相关，需要认证)
    POST /create
    POST /join
    GET  /my

/admin                         (管理员专用，需认证+admin权限)
  POST   /challenge            创建题目
  PUT    /challenge/:id        更新题目
  DELETE /challenge/:id        删除题目
  
  POST   /competition          创建比赛
  PUT    /competition/:id      更新比赛
  DELETE /competition/:id      删除比赛
  
  POST   /upload               上传文件
  
  GET    /users                列出所有用户
  PUT    /user/:id/role        设置用户角色
  DELETE /user/:id             删除用户
  
  GET    /statistics/competition/:id  获取比赛统计
```

### 2.2 权限验证流程

```
请求 → CORS 中间件 → 路由匹配
       ↓
   [认证类路由]
   ↓
   JWTAuth 中间件 (验证 Token，提取身份)
   ↓
   [管理员类路由]
   ↓
   AdminOnly 中间件 (检查是否为 admin)
   ↓
   业务处理
```

---

## 3. Handler 层改进

### 3.1 统一的错误响应格式

```json
// 权限拒绝
{
  "code": 403,
  "msg": "没有权限，仅管理员可以执行此操作"
}

// 参数错误
{
  "code": 400,
  "msg": "参数错误: ..."
}

// 内部错误
{
  "code": 500,
  "msg": "操作失败: ..."
}

// 成功响应
{
  "code": 200,
  "msg": "操作成功",
  "data": {...}
}
```

### 3.2 新增 Handler 文件

**challenge_admin_handler.go**
- `UpdateChallenge` - 更新题目
- `DeleteChallenge` - 删除题目

**competition_admin_handler.go**
- `UpdateCompetition` - 更新比赛
- `DeleteCompetition` - 删除比赛

**user_admin_handler.go**
- `ListUsers` - 列出所有用户
- `SetUserRole` - 设置用户角色
- `DeleteUser` - 删除用户

**statistics_handler.go**
- `GetCompetitionStatistics` - 获取比赛统计信息

### 3.3 参数验证增强

创建/更新操作中添加了数据验证：
- 标题/描述不能为空
- 分数必须 > 0
- 开始时间 < 结束时间
- 动态题目必须提供镜像和端口
- 用户角色只能是 "user" 或 "admin"
- 防止管理员修改自己的角色或删除自己

### 3.4 审计日志

所有管理员操作都会记录到日志：

```
[审计] 管理员 1 创建了比赛: CTF2026 (ID: 5)
[审计] 管理员 1 更新了题目: Web安全 (ID: 10)
[审计] 管理员 1 删除了用户 3 (alice)
```

---

## 4. Service 层改进

### 4.1 签名更新

所有创建/更新/删除函数添加了 `adminID` 参数用于审计：

```go
// 题目
CreateChallenge(...args, adminID uint) error
UpdateChallenge(...args, adminID uint) error
DeleteChallenge(chalID, adminID uint) error

// 比赛
CreateCompetition(...args, adminID uint) error
UpdateCompetition(...args, adminID uint) error
DeleteCompetition(compID, adminID uint) error

// 用户
ListAllUsers(adminID uint) ([]UserInfo, error)
SetUserRole(userID uint, newRole string, adminID uint) error
DeleteUser(userID uint, adminID uint) error
```

### 4.2 新增 Service 文件

**user_admin_service.go**
- `ListAllUsers` - 获取用户列表
- `SetUserRole` - 修改用户角色
- `DeleteUser` - 删除用户及其相关数据

**statistics_service.go**
- `GetCompetitionStatistics` - 获取组织详细统计

### 4.3 关键数据保护

- 用户密码哈希不会在任何接口返回
- Flag 在列表接口中不返回（仅在创建/更新时需要）
- `UserInfo` 结构体用于返回用户信息（排除敏感字段）

---

## 5. 用户角色说明

### 5.1 user 角色

权限：
- 查看比赛和题目列表
- 提交 Flag
- 创建和加入团队
- 查看排行榜
- 上传文件（可根据需要限制）

限制：
- 不能创建/编辑/删除题目
- 不能创建/编辑/删除比赛
- 不能管理用户
- 不能查看统计信息

### 5.2 admin 角色

权限：
- 拥有所有 user 权限
- 创建/编辑/删除题目
- 创建/编辑/删除比赛
- 管理所有用户和权限
- 查看比赛统计信息
- 上传文件

---

## 6. API 调用示例

### 6.1 创建题目（管理员）

```bash
curl -X POST http://localhost:8080/admin/challenge \
  -H "Authorization: Bearer <admin_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "XSS漏洞利用",
    "category": "Web",
    "description": "发现并利用XSS漏洞",
    "flag": "flag{xss_is_fun}",
    "score": 100,
    "competition_id": 1,
    "tags": ["Web", "XSS"],
    "is_dynamic": false
  }'
```

响应：
```json
{
  "code": 200,
  "msg": "题目创建成功"
}
```

### 6.2 更新题目（管理员）

```bash
curl -X PUT http://localhost:8080/admin/challenge/5 \
  -H "Authorization: Bearer <admin_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "XSS漏洞利用（高级）",
    "score": 150
  }'
```

### 6.3 删除题目（管理员）

```bash
curl -X DELETE http://localhost:8080/admin/challenge/5 \
  -H "Authorization: Bearer <admin_token>"
```

### 6.4 列出所有用户（管理员）

```bash
curl -X GET http://localhost:8080/admin/users \
  -H "Authorization: Bearer <admin_token>"
```

响应：
```json
{
  "code": 200,
  "msg": "获取成功",
  "data": [
    {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "role": "admin",
      "score": 500
    },
    {
      "id": 2,
      "username": "alice",
      "email": "alice@example.com",
      "role": "user",
      "score": 300
    }
  ]
}
```

### 6.5 提升用户为管理员

```bash
curl -X PUT http://localhost:8080/admin/user/2/role \
  -H "Authorization: Bearer <admin_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "role": "admin"
  }'
```

### 6.6 获取比赛统计（管理员）

```bash
curl -X GET http://localhost:8080/admin/statistics/competition/1 \
  -H "Authorization: Bearer <admin_token>"
```

响应：
```json
{
  "code": 200,
  "msg": "获取成功",
  "data": {
    "competition_id": 1,
    "competition_name": "CTF2026",
    "total_challenges": 15,
    "total_participants": 50,
    "total_solves": 120,
    "average_solve_rate": 16.0
  }
}
```

---

## 7. 权限错误处理

### 7.1 常见错误响应

**未登录**
```json
{
  "code": 401,
  "msg": "未登录，请提供 Token"
}
```

**Token 无效**
```json
{
  "code": 401,
  "msg": "Token 无效或已过期"
}
```

**权限不足**
```json
{
  "code": 403,
  "msg": "没有权限，仅管理员可以创建题目"
}
```

**参数错误**
```json
{
  "code": 400,
  "msg": "参数错误: ..."
}
```

### 7.2 客户端处理建议

```javascript
// 前端需要处理的错误情况
if (response.code === 401) {
  // 需要登录或刷新 Token
  redirectToLogin();
} else if (response.code === 403) {
  // 权限不足，显示提示
  showError("您没有权限执行此操作");
} else if (response.code === 400) {
  // 参数错误，显示错误信息
  showError(response.msg);
} else if (response.code === 500) {
  // 服务器错误
  showError("服务器错误，请稍后重试");
}
```

---

## 8. 审计日志

### 8.1 日志位置

所有权限相关操作的日志会输出到标准日志系统，可以通过以下方式查看：

```bash
# 在启动服务器时，所有 [审计] 开头的日志会显示在控制台
# 或者重定向到文件
go run ./cmd/server > server.log 2>&1
```

### 8.2 日志格式

```
[审计] 操作者身份 执行了什么操作 相关信息
示例：
  [审计] 管理员 1 创建了题目: XSS漏洞 (ID: 8)
  [审计] 管理员 1 删除了用户 3 (bob)
  [审计] 管理员 1 将用户 2 (alice) 的角色从 user 修改为 admin
```

---

## 9. 安全建议

### 9.1 部署时的考虑

1. **HTTPS 使用**：始终在生产环境中使用 HTTPS
2. **JWT Secret**：定期更换 JWT 密钥
3. **Token 过期**：设置合理的 Token 过期时间
4. **速率限制**：考虑添加速率限制，防止暴力破解
5. **日志监控**：定期审查审计日志，发现异常操作

### 9.2 初始设置

```go
// 建议在数据库初始化时创建一个管理员账户
// 例如在 bootstrap/app.go 中
admin := model.User{
  Username: "admin",
  Email: "admin@example.com",
  Role: "admin",
  // 密码应该通过 bcrypt 加密
}
db.Create(&admin)
```

### 9.3 权限检查清单

- [ ] 所有创建操作都检查了 admin 权限
- [ ] 所有删除操作都检查了 admin 权限
- [ ] 用户密码从不在响应中返回
- [ ] Flag 从不在列表接口中返回
- [ ] 管理员不能修改自己的权限
- [ ] 管理员不能删除自己
- [ ] 所有关键操作都有审计日志

---

## 10. 后续改进方向

1. **细粒度权限**：实现基于资源的权限检查
2. **角色扩展**：支持更多角色（如出题管理员、比赛管理员）
3. **权限继承**：实现权限组的概念
4. **审计系统**：将审计日志存储到数据库
5. **操作回滚**：实现某些操作的撤销功能
6. **操作历史**：跟踪资源的修改历史

