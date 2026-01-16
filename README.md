# B2CTF
A CTF platform

## 快速开始

### 环境要求
- **Go**: 1.25+ （后端开发）
- **Node.js**: 18+ （前端开发）

### 安装步骤

1. **克隆项目**
   ```bash
   git clone <项目仓库地址>
   cd B2CTF
   ```

2. **安装后端依赖**
   ```bash
   cd backend
   go env -w GOPROXY="https://goproxy.cn,direct"
   go mod tidy
   cd ..
   ```

3. **安装前端依赖**
   ```bash
   cd frontend
   npm install
   cd ..
   ```

### 启动方式

#### 方式一：使用一键启动脚本（推荐）

##### Windows平台
```powershell
# 使用PowerShell执行
.\start_windows.ps1
```

##### Linux平台（Debian系列）
```bash
# 赋予执行权限（首次使用）
chmod +x start_linux.sh

# 执行脚本
./start_linux.sh
```

#### 方式二：手动启动

1. **启动后端服务**
   ```bash
   cd backend
   go run ./cmd/server -config .\configs\config.example.yaml
   ```

2. **启动前端服务**（新终端）
   ```bash
   cd frontend
   npm run dev
   ```

### 访问地址
- **前端界面**: http://localhost:5173
- **后端API**: http://localhost:8080
- **测试接口**: http://localhost:8080/api/ping

## 项目结构

### 前端 (Vue 3 + Vite)
```
frontend/
├── index.html
├── package.json
├── vite.config.js
└── src/
    ├── main.js
    ├── App.vue                  # 主应用组件
    └── api/
        ├── http.js              # axios 封装
        └── ping.js              # ping 接口调用
```

### 后端 (Go + Gin)
```
backend/
├── cmd/
│   └── server/
│       └── main.go              # 入口
├── internal/
│   ├── config/                  # 配置加载
│   ├── router/                  # 路由注册（包含CORS配置）
│   ├── middleware/              # 中间件
│   ├── model/                   # 数据模型
│   ├── repository/              # 数据访问层
│   ├── service/                 # 业务逻辑层
│   ├── handler/                 # HTTP handler
│   ├── db/                      # 数据库初始化
│   └── pkg/                     # 通用小工具
├── configs/
│   └── config.example.yaml      # 配置示例
└── go.mod
```

## 核心功能

- ✅ 用户认证（注册、登录）
- ✅ 题目管理
- ✅ Flag提交
- ✅ 排行榜
- ✅ 管理后台
- ✅ 跨域支持
- ✅ 一键启动脚本

## 配置说明

### 后端配置
配置文件路径：`backend/configs/config.example.yaml`

```yaml
server:
  addr: ":8080"

database:
  driver: "mysql"
  dsn: "b2ctf_user:b2ctf_test@tcp(127.0.0.1:3306)/b2ctf?charset=utf8mb4&parseTime=True&loc=Local"

jwt:
  secret: "change-me-to-a-random-secret"
  issuer: "ctf-platform"
  expire_hours: 72
```

### 跨域配置
后端已配置CORS，允许所有来源的请求：
- 允许的方法：GET、POST、PUT、DELETE、OPTIONS
- 允许的请求头：Content-Type、Authorization

## 开发说明

### 前端开发
```bash
cd frontend
npm run dev      # 开发模式
npm run build    # 构建生产版本
```

### 后端开发
```bash
cd backend
go run ./cmd/server -config ./configs/config.example.yaml    # 开发模式
go build -o b2ctf-server ./cmd/server/                       # 构建生产版本
```

## 注意事项

1. **首次使用**：请先安装Go和Node.js环境
2. **数据库**：默认使用MySQL，需自行创建数据库
3. **JWT密钥**：生产环境请修改JWT密钥
4. **跨域配置**：生产环境建议限制允许的来源
5. **一键脚本**：Windows需要PowerShell 5.1+，Linux需要bash环境

## License
MIT