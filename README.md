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