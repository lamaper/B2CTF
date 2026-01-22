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
│   │   │   ├── challenge_handler.go # 题目接口
│   │   │   ├── competition_handler.go # 比赛接口
│   │   │   ├── rank_handler.go # 排行榜接口
│   │   │   ├── submit_handler.go # Flag提交接口
│   │   │   ├── team_handler.go # 团队接口
│   │   │   ├── upload_handler.go # 文件上传接口
│   │   │   └── user_handler.go # 用户接口
│   │   ├── middleware/     # 中间件
│   │   │   ├── auth.go     # JWT认证中间件
│   │   │   └── logger.go   # 日志中间件
│   │   ├── model/          # 数据模型
│   │   │   ├── challenge.go # 题目模型
│   │   │   ├── competition.go # 比赛模型
│   │   │   ├── rank.go     # 排行榜模型
│   │   │   ├── solve.go    # 解题记录模型
│   │   │   ├── team.go     # 团队模型
│   │   │   └── user.go     # 用户模型
│   │   ├── pkg/            # 通用工具包
│   │   │   └── utils/      # 工具函数
│   │   │       └── jwt.go  # JWT工具
│   │   ├── router/         # 路由配置
│   │   │   └── router.go   # 路由注册
│   │   └── service/        # 业务逻辑层
│   │       ├── challenge_service.go # 题目服务
│   │       ├── competition_service.go # 比赛服务
│   │       ├── rank_service.go # 排行榜服务
│   │       ├── submit_service.go # Flag提交服务
│   │       ├── team_service.go # 团队服务
│   │       └── user_service.go # 用户服务
│   ├── go.mod              # Go模块文件
│   └── go.sum              # Go依赖校验文件
├── frontend/               # 前端代码目录
│   ├── public/             # 静态资源
│   ├── src/                # 源代码
│   │   ├── api/            # API调用
│   │   │   ├── http.js     # Axios封装
│   │   │   └── ping.js     # Ping接口调用
│   │   ├── assets/         # 资源文件
│   │   ├── components/     # 组件
│   │   ├── router/         # 路由配置
│   │   │   └── index.js    # 路由定义
│   │   ├── views/          # 页面组件
│   │   │   ├── ChallengeDetail.vue # 题目详情页
│   │   │   ├── Challenges.vue # 题目列表页
│   │   │   ├── Competitions.vue # 比赛列表页
│   │   │   ├── CreateChallenge.vue # 创建题目页
│   │   │   ├── CreateCompetition.vue # 创建比赛页
│   │   │   ├── Home.vue    # 首页
│   │   │   ├── Login.vue   # 登录页
│   │   │   ├── Profile.vue # 用户信息页
│   │   │   └── Register.vue # 注册页
│   │   ├── App.vue         # 根组件
│   │   ├── main.js         # 入口文件
│   │   └── style.css       # 全局样式
│   ├── index.html          # HTML模板
│   ├── package.json        # npm配置
│   └── vite.config.js      # Vite配置
├── docs/                   # 文档目录
│   ├── api.md              # API文档
│   └── project_structure.md # 项目结构文档
├── deploy/                 # 部署相关
├── .gitignore              # Git忽略文件
├── LICENSE                 # 许可证文件
├── README.md               # 项目说明
├── start_linux.sh          # Linux启动脚本
└── start_windows.ps1       # Windows启动脚本
```

## 核心文件说明

### 后端核心文件

#### 已实现的文件

1. **backend/cmd/server/main.go**
   - 功能：服务器主入口文件
   - 创建者：lamaper
   - 创建日期：2026-01-17
   - 最后修改：2026-01-17
   - 主要函数：
     - `main()` - 程序入口，加载配置并启动服务

2. **backend/internal/bootstrap/app.go**
   - 功能：应用启动逻辑
   - 创建者：lamaper
   - 创建日期：2026-01-17
   - 最后修改：2026-01-17
   - 主要函数：
     - `Run(configPath string) error` - 启动应用服务

3. **backend/internal/config/config.go**
   - 功能：配置加载和管理
   - 创建者：lamaper
   - 创建日期：2026-01-21
   - 最后修改：2026-01-21
   - 主要函数：
     - `Init(configPath string)` - 初始化配置

4. **backend/internal/db/db.go**
   - 功能：数据库初始化和连接管理
   - 创建者：lamaper
   - 创建日期：2026-01-21
   - 最后修改：2026-01-21
   - 主要函数：
     - `Init()` - 初始化数据库连接

5. **backend/internal/handler/challenge_handler.go**
   - 功能：处理题目相关的HTTP请求
   - 创建者：lamaper
   - 创建日期：2026-01-21
   - 最后修改：2026-01-21
   - 主要函数：
     - `CreateChallenge(c *gin.Context)` - 创建题目
     - `ListChallenges(c *gin.Context)` - 获取题目列表

6. **backend/internal/handler/competition_handler.go**
   - 功能：处理比赛相关的HTTP请求
   - 创建者：lamaper
   - 创建日期：2026-01-21
   - 最后修改：2026-01-21
   - 主要函数：
     - `CreateCompetition(c *gin.Context)` - 创建比赛
     - `ListCompetitions(c *gin.Context)` - 获取比赛列表

7. **backend/internal/handler/user_handler.go**
   - 功能：处理用户相关的HTTP请求
   - 创建者：lamaper
   - 创建日期：2026-01-17
   - 最后修改：2026-01-17
   - 主要函数：
     - `UserRegister(c *gin.Context)` - 处理注册请求
     - `UserLogin(c *gin.Context)` - 处理登录请求
     - `GetUserProfile(c *gin.Context)` - 获取用户信息
     - `UploadAvatar(c *gin.Context)` - 修改头像

8. **backend/internal/model/challenge.go**
   - 功能：题目数据模型
   - 创建者：lamaper
   - 创建日期：2026-01-21
   - 最后修改：2026-01-21
   - 主要字段：ID、Title、Category、Description、Score、Flag、Attachment、Tags、SolvedCount

9. **backend/internal/model/competition.go**
   - 功能：比赛数据模型
   - 创建者：lamaper
   - 创建日期：2026-01-21
   - 最后修改：2026-01-21
   - 主要字段：ID、Title、Description、Type、StartTime、EndTime、Mode

10. **backend/internal/model/solve.go**
    - 功能：解题记录数据模型
    - 创建者：lamaper
    - 创建日期：2026-01-21
    - 最后修改：2026-01-21
    - 主要字段：ID、UserID、ChallengeID、CompetitionID、Score

11. **backend/internal/model/team.go**
    - 功能：团队数据模型
    - 创建者：lamaper
    - 创建日期：2026-01-21
    - 最后修改：2026-01-21
    - 主要字段：ID、Name、Description、Token、CaptainID、Members

12. **backend/internal/model/user.go**
    - 功能：用户数据模型
    - 创建者：lamaper
    - 创建日期：2026-01-17
    - 最后修改：2026-01-17
    - 主要字段：ID、Username、PasswordHash、Email、Role、CreatedAt、Score、Avatar、TeamID

13. **backend/internal/router/router.go**
    - 功能：路由配置
    - 创建者：lamaper
    - 创建日期：2026-01-17
    - 最后修改：2026-01-22
    - 主要路由：
      - POST /api/register - 用户注册
      - POST /api/login - 用户登录
      - GET /user/profile - 获取用户信息
      - GET /challenge - 获取题目列表
      - POST /challenge - 创建题目
      - POST /upload - 上传文件
      - POST /competitions - 创建比赛
      - GET /competitions - 获取比赛列表
      - POST /team/create - 创建团队
      - POST /team/join - 加入团队
      - GET /team/my - 获取我的团队

14. **backend/internal/service/user_service.go**
    - 功能：用户注册和登录的业务逻辑
    - 创建者：lamaper
    - 创建日期：2026-01-17
    - 最后修改：2026-01-17
    - 主要函数：
      - `Register(username, password, email) error` - 用户注册
      - `Login(username, password) (string, string, error)` - 用户登录

15. **backend/internal/service/challenge_service.go**
    - 功能：题目管理的业务逻辑
    - 创建者：lamaper
    - 创建日期：2026-01-21
    - 最后修改：2026-01-21
    - 主要函数：
      - `CreateChallenge(title, category, desc, flag, score, compID, attachment, tags) error` - 创建题目
      - `GetAllChallenges() ([]model.Challenge, error)` - 获取所有题目

16. **backend/internal/service/competition_service.go**
    - 功能：比赛管理的业务逻辑
    - 创建者：lamaper
    - 创建日期：2026-01-21
    - 最后修改：2026-01-21
    - 主要函数：
      - `CreateCompetition(title, desc, cType, start, end) error` - 创建比赛
      - `GetCompetitions() ([]model.Competition, error)` - 获取比赛列表

17. **backend/internal/pkg/utils/jwt.go**
    - 功能：JWT令牌生成
    - 创建者：lamaper
    - 创建日期：2026-01-17
    - 最后修改：2026-01-17
    - 主要函数：
      - `GenerateToken(userID uint, role string) (string, error)` - 生成JWT令牌

### 前端核心文件

1. **frontend/src/main.js**
   - 功能：前端应用入口文件
   - 创建者：lamaper
   - 创建日期：2026-01-17
   - 最后修改：2026-01-17
   - 主要功能：初始化Vue应用并挂载

2. **frontend/src/App.vue**
   - 功能：前端应用主组件，包含导航栏和全局状态管理
   - 创建者：lamaper
   - 创建日期：2026-01-17
   - 最后修改：2026-01-22
   - 主要功能：管理登录状态、显示导航栏、处理路由变化

3. **frontend/src/router/index.js**
   - 功能：前端路由配置文件
   - 创建者：lamaper
   - 创建日期：2026-01-17
   - 最后修改：2026-01-17
   - 主要功能：定义路由规则和导航守卫

4. **frontend/src/api/http.js**
   - 功能：Axios HTTP客户端配置
   - 创建者：lamaper
   - 创建日期：2026-01-17
   - 最后修改：2026-01-17
   - 主要功能：配置Axios实例、添加请求/响应拦截器

## 技术栈

- **后端**：Go 1.25+, Gin, GORM, MySQL, JWT
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
