# B2CTF 项目结构文档

## 项目概述
B2CTF是一个基于Go后端和Vue前端的CTF竞赛平台，提供用户注册、登录、题目管理、Flag提交等功能。

## 项目目录结构

```
B2CTF/
├── backend/                 # 后端代码目录
│   ├── cmd/                # 命令行入口
│   │   └── server/         # 服务器入口
│   │       └── main.go     # 主入口文件
│   ├── configs/            # 配置文件
│   │   └── config.example.yaml # 配置示例
│   ├── internal/           # 内部代码
│   │   ├── bootstrap/      # 应用初始化
│   │   │   └── app.go      # 应用启动逻辑
│   │   ├── config/         # 配置加载
│   │   │   └── config.go   # 配置结构和加载逻辑
│   │   ├── db/             # 数据库初始化
│   │   │   └── db.go       # 数据库连接和迁移
│   │   ├── handler/        # HTTP处理函数
│   │   │   ├── admin_handler.go  # 管理员接口
│   │   │   ├── auth_handler.go   # 认证接口（部分实现）
│   │   │   ├── challenge_handler.go # 题目接口（待实现）
│   │   │   ├── ping_handler.go   # 测试接口
│   │   │   ├── scoreboard_handler.go # 排行榜接口（待实现）
│   │   │   └── user_handler.go    # 用户接口（已实现）
│   │   ├── middleware/     # 中间件
│   │   │   ├── auth.go     # 认证中间件（待实现）
│   │   │   └── logger.go   # 日志中间件
│   │   ├── model/          # 数据模型
│   │   │   ├── challenge.go # 题目模型
│   │   │   ├── solve.go    # 解题记录模型
│   │   │   └── user.go     # 用户模型
│   │   ├── pkg/            # 通用工具包
│   │   │   └── utils/      # 工具函数
│   │   │       └── jwt.go  # JWT工具（已实现）
│   │   ├── repository/     # 数据访问层（待实现）
│   │   │   ├── challenge_repo.go
│   │   │   ├── solve_repo.go
│   │   │   └── user_repo.go
│   │   ├── router/         # 路由配置
│   │   │   └── router.go   # 路由注册
│   │   └── service/        # 业务逻辑层
│   │       └── user_service.go # 用户服务（已实现）
│   ├── go.mod              # Go模块文件
│   └── go.sum              # Go依赖校验文件
├── frontend/               # 前端代码目录
│   ├── public/             # 静态资源
│   ├── src/                # 源代码
│   │   ├── assets/         # 资源文件
│   │   ├── components/     # 组件
│   │   ├── api/            # API调用
│   │   ├── App.vue         # 根组件
│   │   └── main.js         # 入口文件
│   ├── index.html          # HTML模板
│   ├── package.json        # npm配置
│   └── vite.config.js      # Vite配置
├── docs/                   # 文档目录
│   ├── api.md              # API文档
│   └── project_structure.md # 项目结构文档
├── deploy/                 # 部署相关
├── .gitignore              # Git忽略文件
├── README.md               # 项目说明
├── start_linux.sh          # Linux启动脚本
└── start_windows.ps1       # Windows启动脚本
```

## 核心文件说明

### 后端核心文件

#### 已实现的文件

1. **backend/internal/service/user_service.go**
   - 功能：用户注册和登录的业务逻辑
   - 创建者：lamaper
   - 创建日期：2026-01-17
   - 最后修改：2026-01-17
   - 主要函数：
     - `Register(username, password, email) error` - 用户注册
     - `Login(username, password) (string, error)` - 用户登录

2. **backend/internal/handler/user_handler.go**
   - 功能：处理用户相关的HTTP请求
   - 创建者：lamaper
   - 创建日期：2026-01-17
   - 最后修改：2026-01-17
   - 主要函数：
     - `UserRegister(c *gin.Context)` - 处理注册请求
     - `UserLogin(c *gin.Context)` - 处理登录请求

3. **backend/internal/pkg/utils/jwt.go**
   - 功能：JWT令牌生成
   - 创建者：lamaper
   - 创建日期：2026-01-17
   - 最后修改：2026-01-17
   - 主要函数：
     - `GenerateToken(userID uint, role string) (string, error)` - 生成JWT令牌

4. **backend/internal/model/user.go**
   - 功能：用户数据模型
   - 主要字段：ID、Username、PasswordHash、Email、Role、CreatedAt

5. **backend/internal/router/router.go**
   - 功能：路由配置
   - 主要路由：
     - POST /api/register - 用户注册
     - POST /api/login - 用户登录
     - GET /api/ping - 测试接口

#### 待实现的文件

- **challenge相关文件**：题目管理功能
- **scoreboard相关文件**：排行榜功能
- **repository层**：数据访问层
- **auth中间件**：JWT认证中间件

### 前端核心文件

- **App.vue**：主应用组件
- **main.js**：应用入口
- **api/http.js**：axios封装
- **router/index.js**：路由配置

## 技术栈

- **后端**：Go 1.24+, Gin, GORM, MySQL, JWT
- **前端**：Vue 3, Vite, Axios
- **工具**：Git, Docker

## 配置文件

- **backend/configs/config.example.yaml**：包含服务器、数据库、JWT等配置

## 启动脚本

- **start_linux.sh**：Linux平台一键启动脚本
- **start_windows.ps1**：Windows平台一键启动脚本

## 文档

- **docs/api.md**：API文档
- **docs/project_structure.md**：项目结构文档
- **README.md**：项目说明文档
