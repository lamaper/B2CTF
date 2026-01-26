// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-17
// 最后修改: 2026-01-17
// 描述: 应用启动逻辑
// ----------------------------------------------------------------------------
package bootstrap

import (
	"fmt"
	"net/http"

	"B2CTF/backend/internal/config"
	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/pkg/docker"
	"B2CTF/backend/internal/router"
)

func Run(configPath string) error {
	// 1. 加载配置
	config.Init(configPath)

	// 2. 初始化 DB
	db.Init()
	docker.Init()
	// 3. 初始化路由
	r := router.SetupRouter()


	// 4. 启动服务
	addr := config.C.Server.Addr
	fmt.Println("server listening on", addr)
	return http.ListenAndServe(addr, r)
}
