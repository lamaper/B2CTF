package docker

import (
	"context"
	"log"

	"github.com/docker/docker/client"
)

// 全局 Docker 客户端
var Cli *client.Client

// Init 初始化 Docker 客户端
func Init() {
	var err error
	// WithAPIVersionNegotiation 自动协商 API 版本，防止服务端客户端版本不一致报错
	Cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("无法连接 Docker 守护进程: %v", err)
	}
	
	// 测试连接
	_, err = Cli.Ping(context.Background())
	if err != nil {
		log.Fatalf("Docker 连接失败 (请确保 Docker Desktop 已启动): %v", err)
	}
	
	log.Println("🐳 Docker 客户端初始化成功")
}