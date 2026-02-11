# B2CTF 权限系统改动说明文档

**版本**: v2.0  
**发布日期**: 2026年2月8日  
**作者**: 开发团队  
**状态**: ✅ 发布就绪

---

## 📋 目录

1. [概述](#概述)
2. [改动汇总](#改动汇总)
3. [详细改动说明](#详细改动说明)
4. [API 变更](#api-变更)
5. [前端适配指南](#前端适配指南)
6. [数据库迁移](#数据库迁移)
7. [测试指南](#测试指南)
8. [部署说明](#部署说明)
9. [问题排查](#问题排查)
10. [参考资源](#参考资源)

---

## 概述

本次更新主要完善了 B2CTF 后端的**管理员权限系统**，实现了完整的权限隔离、权限验证和审计日志功能。

### 核心改进

| 改进方向 | 说明 |
|--------|------|
| **权限架构** | 从单一的角色检查升级到完整的中间件 + handler 双重检查 |
| **路由结构** | 将管理员操作隔离到 `/admin` 路由组 |
| **功能完善** | 添加了题目/比赛/用户管理的更新和删除能力 |
| **审计机制** | 为所有管理员操作添加了审计日志记录 |
| **参数验证** | 增强了所有输入参数的验证检查 |
| **错误处理** | 统一了所有错误响应格式 |

---

## 改动汇总

### 👥 新增文件 (7 个)

#### 1. Middleware 层
- `middleware/permission.go` (87 行)
  - `AdminOnly()` 中间件 - 权限检查
  - `GetUserID()`, `GetUserRole()`, `IsAdmin()` 辅助函数

#### 2. Handler 层 - 题目管理
- `handler/challenge_admin_handler.go` (115 行)
  - `UpdateChallenge()` - 更新题目
  - `DeleteChallenge()` - 删除题目

#### 3. Handler 层 - 比赛管理
- `handler/competition_admin_handler.go` (122 行)
  - `UpdateCompetition()` - 更新比赛
  - `DeleteCompetition()` - 删除比赛

#### 4. Handler 层 - 用户管理
- `handler/user_admin_handler.go` (155 行)
  - `ListUsers()` - 列出所有用户
  - `SetUserRole()` - 设置用户角色
  - `DeleteUser()` - 删除用户

#### 5. Handler 层 - 统计
- `handler/statistics_handler.go` (35 行)
  - `GetCompetitionStatistics()` - 获取比赛统计

#### 6. Service 层 - 用户管理
- `service/user_admin_service.go` (102 行)
  - `ListAllUsers()` - 获取用户列表
  - `SetUserRole()` - 修改用户角色
  - `DeleteUser()` - 删除用户及关联数据

#### 7. Service 层 - 统计
- `service/statistics_service.go` (57 行)
  - `GetCompetitionStatistics()` - 计算比赛统计数据

### 📝 修改文件 (6 个)

#### 1. Handler 层

**challenge_handler.go**
```go
// 改动点：
// 1. 添加 middleware 包导入
// 2. 修改 CreateChallenge() 函数
//    - 添加 adminID 参数处理
//    - 添加 isDynamic, imageName, containerPort 等动态题参数
//    - 添加详细的参数验证
//    - 改进错误响应格式
//    - 添加权限检查使用 IsAdmin()

// 变更前: 支持基础题目创建
// 变更后: 支持动态题目，参数验证完整
```

**competition_handler.go**
```go
// 改动点：
// 1. 添加 mode 参数到 CreateCompetitionRequest
// 2. 修改 CreateCompetition() 函数
//    - 添加 mode 参数支持（个人赛/团队赛）
//    - 使用 middleware.IsAdmin() 替代 c.GetString("role")
//    - 添加时间验证（StartTime < EndTime）
//    - 改进错误响应格式
//    - 传递 mode 参数到 service 层

// 变更前: 不支持 mode 参数
// 变更后: 完整支持个人赛和团队赛模式
```

**upload_handler.go**
```go
// 改动点：
// 1. 添加管理员权限检查
// 2. 添加文件大小检查（≤500MB）
// 3. 改进错误处理和目录创建
// 4. 增加返回的文件信息（原始文件名）

// 变更前: 任何认证用户都可以上传
// 变更后: 仅管理员可以上传，有大小限制
```

#### 2. Service 层

**challenge_service.go**
```go
// 改动点：
// 1. CreateChallenge() 新增参数
//    - isDynamic, imageName, containerPort, memoryLimit, cpuLimit
//    - adminID (用于审计)
// 2. 新增 UpdateChallenge() 函数
//    - 更新题目的各个字段
//    - 添加审计日志
// 3. 修改 DeleteChallenge() 函数
//    - 添加 adminID 参数
//    - 添加审计日志
// 4. 添加 log 包导入和审计记录

// 代码变更量: ~40 行新增，~30 行修改
```

**competition_service.go**
```go
// 改动点：
// 1. CreateCompetition() 新增参数
//    - mode 参数支持
//    - adminID (用于审计)
// 2. 新增 UpdateCompetition() 函数
//    - 更新比赛的各个字段
//    - 添加审计日志
// 3. 新增 DeleteCompetition() 函数
//    - 级联检查（比赛有题目时不允许删除）
//    - 添加审计日志
// 4. 添加 errors 和 log 包导入

// 代码变更量: ~50 行新增，~5 行修改
```

#### 3. Router 层

**router.go**
```go
// 改动点：
// 1. 路由结构完全重组
//    - 公共路由: /api (无需认证)
//    - 受保护路由: / (需要 JWT)
//    - 管理员路由: /admin (需要 JWT + admin)
// 
// 2. 新增 /admin 路由组
//    - POST /challenge - 创建题目
//    - PUT /challenge/:id - 更新题目
//    - DELETE /challenge/:id - 删除题目
//    - POST /competition - 创建比赛
//    - PUT /competition/:id - 更新比赛
//    - DELETE /competition/:id - 删除比赛
//    - POST /upload - 上传文件
//    - GET /users - 列出用户
//    - PUT /user/:id/role - 设置用户角色
//    - DELETE /user/:id - 删除用户
//    - GET /statistics/competition/:id - 比赛统计
//
// 3. 将以下接口从 POST 改为管理员交由 /admin
//    - POST /challenge → POST /admin/challenge
//    - POST /competitions → POST /admin/competition
//    - POST /upload → POST /admin/upload
//
// 4. 添加 AdminOnly() 中间件到管理员路由组

// 代码变更量: ~60% 重写
```

### 📚 文档文件 (新增 4 个)

1. **admin_permission_system.md** - 权限系统详细文档 (~400 行)
2. **ADMIN_PERMISSION_CHANGELOG.md** - 改动记录 (~500 行)
3. **BUILD_CHECK_REPORT.md** - 编译检查报告
4. **FINAL_CHECK_REPORT.md** - 完整检查报告

---

## 详细改动说明

### 权限系统架构

#### 旧架构 ❌

```
请求 → JWT 验证 → 业务处理甚至Handler中检查角色
        ↓
     如果 role != "admin" 返回错误
```

**问题**:
- 权限检查散落在各个 handler 中
- 无统一的权限拒绝响应格式
- 易于遗漏权限检查

#### 新架构 ✅

```
请求
  ↓
CORS 中间件
  ↓
路由匹配
  ↓
[认证中间件] JWTAuth
  ├─ 验证 Token
  ├─ 提取 userID, role
  └─ 存入 Context
  ↓
[权限中间件] AdminOnly (仅 /admin 路由)
  ├─ 检查 role == "admin"
  └─ 拒绝非管理员请求
  ↓
Handler 业务处理
  ├─ 再次验证权限（双重检查）
  ├─ 参数验证
  └─ 调用 Service
  ↓
Service 业务逻辑
  ├─ 数据库操作
  ├─ 审计日志记录
  └─ 返回结果
```

**优点**:
- 权限集中管理，易于维护
- 统一的错误响应格式
- 审计日志完善
- 双重检查，更加安全

### 中间件改动

#### 新增: permission.go

```go
// AdminOnly() 中间件
// - 检查用户是否为管理员
// - 非管理员返回 403 Forbidden
// - 返回统一的错误格式

// GetUserID() 辅助函数
// - 安全地从 Context 提取 userID
// - 处理 float64 类型转换

// GetUserRole() 辅助函数
// - 从 Context 提取用户角色

// IsAdmin() 快速检查函数
// - 一行代码检查是否为管理员
```

### Handler 层改动详情

#### CreateChallenge - 增强的参数支持

```go
// 旧签名
CreateChallenge(c *gin.Context)

// 新特性
// 1. 添加权限检查
if !middleware.IsAdmin(c) {
    // 拒绝请求
}

// 2. 支持动态题目参数
type CreateChallengeRequest struct {
    // ... 原有字段
    IsDynamic     bool   // 新增：是否动态题
    ImageName     string // 新增：Docker 镜像
    ContainerPort int    // 新增：容器端口
    MemoryLimit   string // 新增：内存限制
    CPULimit      string // 新增：CPU 限制
}

// 3. 参数验证更严格
if req.Score <= 0 {
    // 检查分数
}
if !req.IsDynamic {
    // 动态题必填检查
}

// 4. 管理员信息传递
adminID, _ := middleware.GetUserID(c)
service.CreateChallenge(..., adminID)
```

#### UpdateChallenge - 新增的更新能力

```go
// 完全新增的函数
func UpdateChallenge(c *gin.Context) {
    // 1. 权限检查
    // 2. 参数验证
    // 3. 调用服务更新
    // 4. 返回统一响应
}
```

#### DeleteChallenge - 新增的删除能力

```go
// 完全新增的函数
func DeleteChallenge(c *gin.Context) {
    // 1. 权限检查
    // 2. ID 验证
    // 3. 调用服务删除
    // 4. 返回统一响应
}
```

### Service 层改动详情

#### Audit Logging - 审计日志

所有创建/更新/删除操作都会记录：

```go
log.Printf("[审计] 管理员 %d 创建了题目: %s (ID: %d)", adminID, title, challenge.ID)
log.Printf("[审计] 管理员 %d 更新了题目: %s (ID: %d)", adminID, title, chalID)
log.Printf("[审计] 管理员 %d 删除了题目: %s (ID: %d)", adminID, chal.Title, chalID)
log.Printf("[审计] 管理员 %d 将用户 %d (%s) 的角色从 %s 修改为 %s", adminID, userID, username, oldRole, newRole)
```

#### 数据保护 - 级联操作

DeleteUser 时会：
1. 检查用户是否存在
2. 如果用户在团队中：
   - 如果团队只有此用户，删除团队
   - 否则将用户移出团队
3. 删除用户的所有解题记录
4. 删除用户本身

DeleteCompetition 时会：
1. 检查比赛是否存在
2. 检查比赛是否已有题目
   - 有题目则拒绝删除
   - 无题目才允许删除

---

## API 变更

### 🚨 Breaking Changes

这些接口的 URL 已改动：

#### 1. 创建题目

```
旧: POST /challenge
新: POST /admin/challenge

权限: 仅管理员
状态码: 403 (权限不足)
```

**请求示例**:
```bash
curl -X POST http://localhost:8080/admin/challenge \
  -H "Authorization: Bearer <admin_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "XSS漏洞",
    "category": "Web",
    "description": "...",
    "flag": "flag{...}",
    "score": 100,
    "competition_id": 1,
    "tags": ["Web"],
    "is_dynamic": false
  }'
```

#### 2. 创建比赛

```
旧: POST /competitions
新: POST /admin/competition

权限: 仅管理员
状态码: 403 (权限不足)
```

**请求示例**:
```bash
curl -X POST http://localhost:8080/admin/competition \
  -H "Authorization: Bearer <admin_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "CTF2026",
    "description": "...",
    "type": 0,
    "mode": 0,
    "start_time": "2026-02-15T10:00:00Z",
    "end_time": "2026-02-16T10:00:00Z"
  }'
```

#### 3. 上传文件

```
旧: POST /upload
新: POST /admin/upload

权限: 仅管理员
状态码: 403 (权限不足)
文件限制: ≤500MB
```

### ✅ 新增接口

#### 题目管理

```
PUT /admin/challenge/:id
- 更新题目信息
- 权限: 管理员
- 参数: 同创建时

DELETE /admin/challenge/:id
- 删除题目
- 权限: 管理员
- 返回: {code: 200, msg: "题目删除成功"}
```

#### 比赛管理

```
PUT /admin/competition/:id
- 更新比赛信息
- 权限: 管理员
- 参数: 包括新增的 mode 参数

DELETE /admin/competition/:id
- 删除比赛
- 权限: 管理员
- 限制: 比赛无题目时才能删除
- 返回: {code: 200, msg: "比赛删除成功"}
```

#### 用户管理

```
GET /admin/users
- 列出所有用户
- 权限: 管理员
- 返回: [{id, username, email, role, score, ...}]

PUT /admin/user/:id/role
- 设置用户角色
- 权限: 管理员
- 请求体: {"role": "user|admin"}
- 限制: 不能修改自己的角色

DELETE /admin/user/:id
- 删除用户
- 权限: 管理员
- 限制: 不能删除自己
- 级联: 删除团队（如果只有此成员）、解题记录
```

#### 统计信息

```
GET /admin/statistics/competition/:id
- 获取比赛统计信息
- 权限: 管理员
- 返回: {
    "competition_id": 1,
    "competition_name": "CTF2026",
    "total_challenges": 15,
    "total_participants": 50,
    "total_solves": 120,
    "average_solve_rate": 16.0
  }
```

### 响应格式统一

所有接口现在使用统一的响应格式：

```json
// 成功
{
  "code": 200,
  "msg": "操作成功",
  "data": {...}
}

// 权限拒绝
{
  "code": 403,
  "msg": "没有权限，仅管理员可以执行此操作"
}

// 參數錯誤
{
  "code": 400,
  "msg": "参数错误: ..."
}

// 服务器错误
{
  "code": 500,
  "msg": "操作失败: ..."
}

// 未认证
{
  "code": 401,
  "msg": "未登录，请提供 Token"
}
```

---

## 前端适配指南

### 必须修改的功能

#### 1. 创建题目页面

**文件**: `frontend/src/views/CreateChallenge.vue`

```javascript
// 旧代码
const response = await http.post('/challenge', challengeData);

// 新代码
const response = await http.post('/admin/challenge', challengeData);
```

#### 2. 创建比赛页面

**文件**: `frontend/src/views/CreateCompetition.vue`

```javascript
// 旧代码
const response = await http.post('/competitions', competitionData);

// 新代码
const response = await http.post('/admin/competition', competitionData);
```

还需要添加 `mode` 字段：
```javascript
const competitionData = {
  title: form.value.title,
  description: form.value.description,
  type: form.value.type,
  mode: form.value.mode,  // 新增：0=个人赛，1=团队赛
  start_time: new Date(form.value.startTime).toISOString(),
  end_time: new Date(form.value.endTime).toISOString()
};
```

#### 3. 文件上传功能

**文件**: `frontend/src/api/*.js` 或 `frontend/src/views/CreateChallenge.vue`

```javascript
// 旧代码
formData.post('/upload', file);

// 新代码
formData.post('/admin/upload', file);
```

### 可选的新功能

#### 用户管理页面（仅管理员）

```javascript
// 获取所有用户
const users = await http.get('/admin/users');

// 设置用户角色
await http.put(`/admin/user/${userId}/role`, {
  role: 'admin'  // 或 'user'
});

// 删除用户
await http.delete(`/admin/user/${userId}`);
```

#### 比赛统计页面（仅管理员）

```javascript
// 获取比赛统计
const stats = await http.get(`/admin/statistics/competition/${compId}`);

// 显示统计信息
console.log(stats.data.total_participants);
console.log(stats.data.average_solve_rate);
```

### 权限检查建议

在前端添加权限检查，避免 403 错误：

```javascript
// 检查是否为管理员
const isAdmin = () => {
  const userRole = localStorage.getItem('userRole');
  return userRole === 'admin';
};

// 在创建题目前检查
if (!isAdmin()) {
  alert('您没有权限创建题目');
  router.push('/');
  return;
}
```

---

## 数据库迁移

### ✅ 无需迁移

本次更改**无需任何数据库迁移**，原因：

1. **无新表**：所有新功能都使用现有表
2. **无表结构变更**：现有字段完全兼容
3. **向后兼容**：旧数据无需转换

### 现有表结构回顾

```
users
├── id (PK)
├── username (UNIQUE)
├── password_hash
├── email
├── role ('user' or 'admin')  ← 使用现有字段
├── score
├── avatar
├── team_id
└── ...

competitions
├── id (PK)
├── title
├── description
├── type
├── mode (0=个人, 1=团队)      ← 使用现有字段
├── start_time
├── end_time
├── container_limit
└── ...

challenges
├── id (PK)
├── title
├── category
├── description
├── flag
├── score
├── competition_id (FK)
├── attachment
├── tags (JSON)
├── is_dynamic                 ← 使用现有字段
├── image_name                 ← 使用现有字段
├── container_port             ← 使用现有字段
├── memory_limit               ← 使用现有字段
├── cpu_limit                  ← 使用现有字段
└── ...

[其他表无变化]
```

### 初始化建议

在部署前，创建初始管理员账户：

```sql
-- 创建初始管理员用户
-- 密码需要通过 bcrypt 加密
-- 可以通过客户端在注册页面创建普通用户后，
-- 然后通过管理员在 /admin/user/:id/role 接口升级权限
```

---

## 测试指南

### 单元测试建议

#### 权限中间件测试

```go
// 测试 AdminOnly 中间件
func TestAdminOnlyMiddleware(t *testing.T) {
    // 1. 普通用户请求 → 403
    // 2. 管理员请求 → 通过
    // 3. 无 Token 请求 → 401
}
```

#### Handler 测试

```go
// 测试 CreateChallenge Handler
func TestCreateChallenge(t *testing.T) {
    // 1. 普通用户创建题目 → 403
    // 2. 管理员创建有效题目 → 200
    // 3. 管理员创建无效题目（分数≤0）→ 400
    // 4. 动态题缺少参数 → 400
}
```

### 集成测试检查清单

- [ ] 普通用户无法访问 `/admin/*` 路由
- [ ] 管理员可以创建题目
- [ ] 管理员可以更新题目
- [ ] 管理员可以删除题目
- [ ] 管理员可以创建比赛
- [ ] 管理员可以更新比赛
- [ ] 管理员可以删除有题目的比赛（应失败）
- [ ] 管理员可以删除无题目的比赛
- [ ] 管理员可以上传文件
- [ ] 普通用户上传文件被拒绝
- [ ] 管理员可以列出所有用户
- [ ] 管理员可以修改用户角色
- [ ] 管理员无法修改自己的角色
- [ ] 管理员可以删除其他用户
- [ ] 管理员无法删除自己
- [ ] 删除用户时级联删除其解题记录
- [ ] 审计日志正确记录所有操作

### 手动测试脚本

```bash
#!/bin/bash

# 获取管理员 Token
ADMIN_TOKEN=$(curl -s -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}' \
  | jq -r '.data.token')

# 测试创建题目
curl -X POST http://localhost:8080/admin/challenge \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "测试题目",
    "category": "Web",
    "description": "测试",
    "flag": "flag{test}",
    "score": 100,
    "competition_id": 1,
    "tags": ["test"]
  }'

# 测试更新题目
curl -X PUT http://localhost:8080/admin/challenge/1 \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "更新后的题目",
    "score": 200
  }'

# 测试删除题目
curl -X DELETE http://localhost:8080/admin/challenge/1 \
  -H "Authorization: Bearer $ADMIN_TOKEN"

# 测试普通用户无权限
USER_TOKEN=$(curl -s -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"user","password":"password"}' \
  | jq -r '.data.token')

curl -X POST http://localhost:8080/admin/challenge \
  -H "Authorization: Bearer $USER_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{...}'
# 预期: 403 Forbidden
```

---

## 部署说明

### 前置条件检查

- [x] 后端代码编译成功
- [x] 所有 handler 函数已实现
- [x] 所有 service 函数已实现
- [x] 路由配置正确
- [x] 无编译警告

### 部署步骤

#### 1. 编译后端

```bash
cd backend
go build -o server ./cmd/server
chmod +x server
```

#### 2. 更新前端（必须）

更新以下 3 个 endpoint：
- `POST /challenge` → `POST /admin/challenge`
- `POST /competitions` → `POST /admin/competition`
- `POST /upload` → `POST /admin/upload`

```bash
# 前端文件需要修改
# frontend/src/views/CreateChallenge.vue
# frontend/src/views/CreateCompetition.vue
# frontend/src/api/*.js
```

#### 3. 初始化管理员账户

```bash
# 可选方法 1：通过注册后升级
# 1. 注册普通账户
# 2. 使用现有的任何管理员账户升级他为管理员

# 可选方法 2：数据库直接插入
# 使用数据库脚本直接创建管理员账户
```

#### 4. 启动服务

```bash
# 开发环境
./server -config configs/config.example.yaml

# 生产环境（后台运行）
nohup ./server -config configs/config.yaml > server.log 2>&1 &

# 使用 systemd 管理
# 见部署文档的 systemd service 配置
```

#### 5. 验证

```bash
# 检查服务是否启动
curl http://localhost:8080/api/login
# 预期: 405 Method Not Allowed (正常，因为 POST 才对)

# 测试权限
curl -X GET http://localhost:8080/admin/users
# 预期: 401 Unauthorized (未提供 Token)
```

### 回滚方案

如果出现问题，可以快速回滚：

```bash
# 1. 停止服务
killall server

# 2. 恢复旧的二进制文件
cp /backup/server_old server

# 3. 恢复前端
git checkout frontend/src/views/CreateChallenge.vue
git checkout frontend/src/views/CreateCompetition.vue
# ...其他前端文件

# 4. 重启服务
./server -config configs/config.yaml

# 注意：数据库无需恢复，因为本次更新无数据库变更
```

---

## 问题排查

### 常见问题

#### Q: 提示 "没有权限，仅管理员可以创建题目"

**原因**: 当前用户不是管理员

**解决**:
1. 使用管理员账户登录
2. 或者让管理员将你的账户升级为管理员

```bash
# 使用管理员 Token 升级用户角色
curl -X PUT http://localhost:8080/admin/user/2/role \
  -H "Authorization: Bearer <admin_token>" \
  -H "Content-Type: application/json" \
  -d '{"role": "admin"}'
```

#### Q: 前端调用 `/challenge` 接口返回 404

**原因**: 接口已改名为 `/admin/challenge`

**解决**:
检查前端代码，将 endpoint 更新为 `/admin/challenge`

```javascript
// 错误
await http.post('/challenge', data);

// 正确
await http.post('/admin/challenge', data);
```

#### Q: 上传文件提示 "文件过大"

**原因**: 文件大于 500MB

**解决**:
- 压缩文件大小
- 或者联系管理员修改 upload_handler.go 中的大小限制

```go
const maxFileSize = 500 * 1024 * 1024  // 改大这个值
```

#### Q: 删除比赛时提示 "比赛已有题目，无法删除"

**原因**: 比赛中还有题目

**解决**:
1. 先删除比赛中的所有题目
2. 然后再删除比赛

#### Q: 权限检查总是失败，即使是管理员

**原因**: 可能以下几个原因
1. Token 已过期
2. 用户角色字段为空
3. 中间件未正确初始化

**排查**:
```bash
# 1. 检查 Token 有效性
curl -X GET http://localhost:8080/user/profile \
  -H "Authorization: Bearer <token>"

# 2. 检查用户角色
# 登录后检查返回的 userRole 字段是否为 "admin"

# 3. 查看服务器日志
tail -f server.log | grep "authenticate\|permission"
```

### 日志查看

所有审计操作都会记录在日志中：

```bash
# 查看审计日志
grep "\[审计\]" server.log

# 输出示例
# [审计] 管理员 1 创建了题目: XSS漏洞 (ID: 5)
# [审计] 管理员 1 更新了比赛: CTF2026 (ID: 1)
# [审计] 管理员 1 删除了用户 3 (alice)
```

---

## 参考资源

### 相关文档

1. **权限系统设计** - `docs/admin_permission_system.md`
   - 详细的权限系统说明
   - API 使用示例
   - 安全建议

2. **API 文档** - `docs/api.md`
   - 所有 API 接口说明
   - 请求/响应格式
   - 错误码列表

3. **编译检查报告** - `docs/BUILD_CHECK_REPORT.md`
   - 编译过程细节
   - 代码质量检查
   - 部署前清单

4. **完整检查报告** - `docs/FINAL_CHECK_REPORT.md`
   - 全面的检查总结
   - 功能验证清单
   - 改进建议

### 配置文件

- `backend/configs/config.example.yaml` - 后端配置示例
- `backend/go.mod` - Go 依赖管理

### 关键代码文件

#### Middleware
- `backend/internal/middleware/auth.go` - JWT 认证
- `backend/internal/middleware/permission.go` - 权限管理（新）

#### Handler
- `backend/internal/handler/challenge_handler.go` - 题目处理（修改）
- `backend/internal/handler/challenge_admin_handler.go` - 题目管理（新）
- `backend/internal/handler/competition_handler.go` - 比赛处理（修改）
- `backend/internal/handler/competition_admin_handler.go` - 比赛管理（新）
- `backend/internal/handler/user_admin_handler.go` - 用户管理（新）

#### Service
- `backend/internal/service/challenge_service.go` - 题目服务（修改）
- `backend/internal/service/competition_service.go` - 比赛服务（修改）
- `backend/internal/service/user_admin_service.go` - 用户管理服务（新）

#### Router
- `backend/internal/router/router.go` - 路由配置（重组）

### 外部链接

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM Documentation](https://gorm.io)
- [JWT 流程说明](https://jwt.io)

---

## 总结

本次更新完成了 B2CTF 后端权限系统的全面升级，包括：

✅ **权限架构**：从无序到有序，从分散到集中  
✅ **功能完善**：从基础到完整，支持 CRUD 全部操作  
✅ **安全加强**：添加了审计日志和双重权限检查  
✅ **代码质量**：统一的错误格式，详细的参数验证  

**下一步行动**:
1. 更新前端 3 个 endpoint （30 分钟）
2. 初始化管理员账户 （10 分钟）
3. 执行集成测试 （1-2 小时）
4. 部署到生产环境 （30 分钟）

**预计上线时间**: 2-3 天

---

**版本历史**

| 版本 | 日期 | 描述 |
|------|------|------|
| v1.0 | 2026-01-17 | B2CTF 初始版本 |
| v2.0 | 2026-02-08 | 完善权限系统 |

**最后更新**: 2026年2月8日

