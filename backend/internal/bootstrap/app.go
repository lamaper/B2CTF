package bootstrap

import (
    "fmt"
    "net/http"

    "B2CTF/backend/internal/config"
    "B2CTF/backend/internal/db"
    "B2CTF/backend/internal/router"
)

func Run(configPath string) error {
    // 1. 加载配置
    config.Init(configPath)

    // 2. 初始化 DB
    db.Init()

    // 3. 初始化路由
    r := router.SetupRouter()

    // 4. 启动服务
    addr := config.C.Server.Addr
    fmt.Println("server listening on", addr)
    return http.ListenAndServe(addr, r)
}
