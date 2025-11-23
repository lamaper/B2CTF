# B2CTF
A CTF platform

## Start
Windows环境下下载[Go1.25.4windows-amd64.msi](https://go.dev/dl/go1.25.4.windows-amd64.msi)并安装

克隆本项目到本地：
```bash
git clone https://github.com/golang/go.git
```

换源并安装依赖：
```bash
cd ./backend

go env -w GOPROXY="https://goproxy.cn,direct"

go mod tidy
```

启动测试环境：
```bash
go run ./cmd/server -config .\configs\config.example.yaml
```

浏览器访问`http://localhost:8080/api/ping`，即可看到返回的json消息`{"msg": "pong"}`

## 项目结构

前端
```
frontend/
├── index.html
├── package.json
├── tsconfig.json
├── vite.config.ts
└── src/
    ├── main.ts
    ├── App.vue
    ├── router/
    │   └── index.ts
    ├── store/
    │   └── user.ts          # Pinia 用户状态，存 token 等
    ├── api/
    │   ├── http.ts          # axios 封装
    │   ├── auth.ts          # 登录注册相关请求
    │   ├── challenge.ts     # 题目列表、详情、提交 flag
    │   └── scoreboard.ts    # 排行榜
    ├── views/
    │   ├── LoginView.vue
    │   ├── ChallengeListView.vue
    │   ├── ChallengeDetailView.vue
    │   ├── ScoreboardView.vue
    │   └── Admin/
    │       ├── AdminLayout.vue
    │       └── ChallengeManageView.vue
    ├── components/
    │   ├── NavBar.vue
    │   └── ChallengeCard.vue
    ├── layouts/
    │   └── DefaultLayout.vue
    ├── styles/
    │   └── global.css
    └── utils/
        └── auth.ts           # token 工具、角色判断等
```

后端
```
backend/
├── cmd/
│   └── server/
│       └── main.go              # 入口
├── internal/
│   ├── config/                  # 配置加载
│   │   └── config.go
│   ├── router/                  # 路由注册
│   │   └── router.go
│   ├── middleware/              # 中间件（JWT、日志、恢复等）
│   │   ├── auth.go
│   │   └── logger.go
│   ├── model/                   # 数据模型（GORM struct）
│   │   ├── user.go
│   │   ├── challenge.go
│   │   └── solve.go
│   ├── repository/              # 数据访问层（对 DB 操作）
│   │   ├── user_repo.go
│   │   ├── challenge_repo.go
│   │   └── solve_repo.go
│   ├── service/                 # 业务逻辑层
│   │   ├── auth_service.go
│   │   ├── challenge_service.go
│   │   └── scoreboard_service.go
│   ├── handler/                 # HTTP handler（Gin 的 handler）
│   │   ├── auth_handler.go
│   │   ├── challenge_handler.go
│   │   ├── scoreboard_handler.go
│   │   └── admin_handler.go
│   ├── db/                      # 数据库初始化
│   │   └── db.go
│   ├── pkg/                     # 通用小工具
│   │   ├── jwt/                 # JWT 相关
│   │   │   └── jwt.go
│   │   ├── hash/                # 密码和 flag hash
│   │   │   └── hash.go
│   │   ├── logger/              # 日志封装
│   │   │   └── logger.go
│   │   └── response/            # 统一响应封装
│   │       └── response.go
│   └── bootstrap/               # 启动初始化逻辑组合
│       └── app.go
├── configs/
│   └── config.example.yaml      # 配置示例（可以拷贝为 config.yaml）
├── go.mod
└── go.sum

```