# 管理员权限系统改动总结

## 修改日期
2026年2月8日

## 概述
完善了 B2CTF 后端的管理员权限系统，实现了完整的权限验证、权限隔离和审计日志功能。

---

## 文件修改清单

### 新增文件

#### 中间件
- **permission.go** (`backend/internal/middleware/permission.go`)
  - 添加 `AdminOnly()` 中间件
  - 添加 `GetUserID()`, `GetUserRole()`, `IsAdmin()` 辅助函数
  - 统一的权限拒绝响应

#### 题目管理
- **challenge_admin_handler.go** (`backend/internal/handler/challenge_admin_handler.go`)
  - `UpdateChallenge()` - 更新题目
  - `DeleteChallenge()` - 删除题目

#### 比赛管理
- **competition_admin_handler.go** (`backend/internal/handler/competition_admin_handler.go`)
  - `UpdateCompetition()` - 更新比赛
  - `DeleteCompetition()` - 删除比赛

#### 用户管理
- **user_admin_handler.go** (`backend/internal/handler/user_admin_handler.go`)
  - `ListUsers()` - 列出所有用户
  - `SetUserRole()` - 设置用户角色
  - `DeleteUser()` - 删除用户

- **user_admin_service.go** (`backend/internal/service/user_admin_service.go`)
  - `ListAllUsers()` - 获取用户列表（排除密码）
  - `SetUserRole()` - 修改用户角色
  - `DeleteUser()` - 删除用户及其关联数据

#### 统计
- **statistics_handler.go** (`backend/internal/handler/statistics_handler.go`)
  - `GetCompetitionStatistics()` - 获取比赛统计

- **statistics_service.go** (`backend/internal/service/statistics_service.go`)
  - `GetCompetitionStatistics()` - 计算统计数据

#### 文档
- **admin_permission_system.md** (`docs/admin_permission_system.md`)
  - 完整的权限系统文档和 API 示例

---

## 文件修改清单

### 修改的文件

#### 中间件
**auth.go** (`backend/internal/middleware/auth.go`)
- 无变化，但现在与 `permission.go` 配合使用

#### 处理器

**challenge_handler.go** (`backend/internal/handler/challenge_handler.go`)
- 更新 `CreateChallenge()` 函数
  - 添加对动态题目参数的支持
  - 添加详细的参数验证
  - 添加 `adminID` 传递与审计
  - 改进错误信息格式

**competition_handler.go** (`backend/internal/handler/competition_handler.go`)
- 更新 `CreateCompetition()` 函数
  - 添加 `mode` 参数支持（个人赛/团队赛）
  - 添加详细的参数验证
  - 改进权限检查使用 `middleware.IsAdmin()`
  - 改进错误响应格式

**upload_handler.go** (`backend/internal/handler/upload_handler.go`)
- 更新 `UploadFile()` 函数
  - 添加管理员权限检查
  - 添加文件大小限制（500MB）
  - 改进错误处理
  - 返回更多信息（原始文件名）

#### 服务层

**challenge_service.go** (`backend/internal/service/challenge_service.go`)
- 完全重写
  - `CreateChallenge()` 新增 adminID, isDynamic, imageName 等参数
  - 新增 `UpdateChallenge()` 函数
  - 新增 `DeleteChallenge()` 函数（带权限检查）
  - 所有函数添加审计日志记录

**competition_service.go** (`backend/internal/service/competition_service.go`)
- 主要更新
  - `CreateCompetition()` 新增 mode 和 adminID 参数
  - 新增 `UpdateCompetition()` 函数
  - 新增 `DeleteCompetition()` 函数（带级联检查）
  - 所有函数添加审计日志记录

#### 路由

**router.go** (`backend/internal/router/router.go`)
- 完全重组
  - 分离公共接口 (`/api`)
  - 分离认证接口 (`/`)
  - 新增管理员接口 (`/admin`)
  - 添加 `AdminOnly()` 中间件到管理员路由
  - 脱离保护路由中的管理员操作（`POST /challenge`, `POST /competitions`, `POST /upload`）

---

## 关键改进点

### 1. 权限隔离
```
旧方式：所有认证用户都可以尝试创建题目/比赛
新方式：只有管理员可以访问 /admin/* 路由
```

### 2. 错误响应统一
```json
{
  "code": 403,           // HTTP 语义的数字码
  "msg": "权限说明"      // 详细错误信息
}
```

### 3. 参数验证增强
- 题目分数 > 0
- 起始时间 < 结束时间
- 动态题必须有镜像和端口信息
- 角色值校验

### 4. 审计日志
```
[审计] 管理员 1 创建了题目: XSS漏洞 (ID: 8)
[审计] 管理员 1 更新了比赛: CTF2026 (ID: 1)
[审计] 管理员 1 删除了用户 3 (bob)
```

### 5. 数据保护
- 密码哈希从不返回
- Flag 只在创建/更新时传递，列表接口不返回
- UserInfo 结构体用于返回用户信息

---

## API 路由变更

### 原路由 → 新路由

| 操作 | 旧路由 | 新路由 | 备注 |
|------|-------|--------|------|
| 创建题目 | `POST /challenge` | `POST /admin/challenge` | 需 admin |
| 创建比赛 | `POST /competitions` | `POST /admin/competition` | 需 admin |
| 上传文件 | `POST /upload` | `POST /admin/upload` | 需 admin |
| 查看题目 | `GET /challenge` | `GET /challenge` | 无变化 |
| 查看比赛 | `GET /competitions` | `GET /competitions` | 无变化 |

### 新增路由

| 方法 | 路由 | 功能 | 权限 |
|------|------|------|------|
| PUT | `/admin/challenge/:id` | 更新题目 | admin |
| DELETE | `/admin/challenge/:id` | 删除题目 | admin |
| PUT | `/admin/competition/:id` | 更新比赛 | admin |
| DELETE | `/admin/competition/:id` | 删除比赛 | admin |
| GET | `/admin/users` | 列出用户 | admin |
| PUT | `/admin/user/:id/role` | 设置用户角色 | admin |
| DELETE | `/admin/user/:id` | 删除用户 | admin |
| GET | `/admin/statistics/competition/:id` | 比赛统计 | admin |

---

## 前端需要的接口更新

### 1. 创建题目
从 `POST /challenge` 改为 `POST /admin/challenge`

```javascript
// 旧
http.post('/challenge', challengeData)

// 新
http.post('/admin/challenge', challengeData)
```

### 2. 创建比赛
从 `POST /competitions` 改为 `POST /admin/competition`

```javascript
// 旧
http.post('/competitions', competitionData)

// 新
http.post('/admin/competition', competitionData)
```

### 3. 文件上传
从 `POST /upload` 改为 `POST /admin/upload`

```javascript
// 旧
formData.post('/upload', formData)

// 新
formData.post('/admin/upload', formData)
```

---

## 数据库影响

**无数据库表结构变化**，所有改动都是业务逻辑层的。

现有表结构完全兼容：
- user 表的 role 字段已支持
- Challenge 表的 is_dynamic, image_name 等字段已支持
- Competition 表的 mode, container_limit 字段已支持

---

## 测试建议

### 1. 权限检查
```bash
# 使用普通用户 Token 创建题目（应该失败）
curl -X POST http://localhost:8080/admin/challenge \
  -H "Authorization: Bearer <user_token>" \
  -d {...}
# 预期：403 权限拒绝

# 使用管理员 Token 创建题目（应该成功）
curl -X POST http://localhost:8080/admin/challenge \
  -H "Authorization: Bearer <admin_token>" \
  -d {...}
# 预期：200 成功
```

### 2. 用户管理
```bash
# 管理员可以列出所有用户
curl -X GET http://localhost:8080/admin/users \
  -H "Authorization: Bearer <admin_token>"

# 管理员可以提升用户权限
curl -X PUT http://localhost:8080/admin/user/2/role \
  -H "Authorization: Bearer <admin_token>" \
  -d '{"role": "admin"}'
```

### 3. 审计日志
检查启动日志中是否出现 `[审计]` 开头的记录

---

## 兼容性说明

### 向后兼容
- ✅ 普通用户的操作完全不变
- ✅ 现有的认证流程不变
- ✅ 数据库表结构不变

### 破坏性变更
- ❌ 管理员创建题目/比赛需要改用新路由 `/admin/*`
- ❌ 旧路由 `POST /challenge` 和 `POST /competitions` 在受保护路由中不再可用
- ❌ 前端需要更新创建操作的 endpoint

---

## 部署检查清单

- [ ] 后端代码已编译无错误
- [ ] 数据库连接正常
- [ ] 管理员账户存在且 role=admin
- [ ] 前端已更新创建操作的 endpoint
- [ ] 测试了权限检查（用户无法访问 /admin/*）
- [ ] 测试了管理员操作（创建、更新、删除）
- [ ] 检查了审计日志输出
- [ ] 验证了错误响应格式

---

## 相关文档

- **权限系统详细文档**：`docs/admin_permission_system.md`
- **API 文档**：`docs/api.md`（需要更新）
- **项目结构**：`docs/project_structure.md`（可更新）

