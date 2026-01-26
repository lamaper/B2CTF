// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-17
// 最后修改: 2026-01-17
// 描述: 服务器主入口文件
// ----------------------------------------------------------------------------
package main

import (
    "flag"
    "log"

    "B2CTF/backend/internal/bootstrap"
)

func main() {
    configPath := flag.String("config", "configs/config.yaml", "config file path")
    flag.Parse()

    if err := bootstrap.Run(*configPath); err != nil {
        log.Fatalf("server exit with error: %v", err)
    }
}
