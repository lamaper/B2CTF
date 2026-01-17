// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-17
// 最后修改: 2026-01-17
// 描述: 路由配置文件
// ----------------------------------------------------------------------------
package router

import (
	"github.com/gin-gonic/gin"

	"B2CTF/backend/internal/handler"
	"B2CTF/backend/internal/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 添加CORS中间件
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 公共接口
	api := r.Group("/api")
	{
		api.POST("/register", handler.UserRegister)
		api.POST("/login", handler.UserLogin)

		// api.GET("/challenges", handler.ListChallenges)
		// api.GET("/challenges/:id", handler.GetChallenge)
		// api.POST("/challenges/:id/submit", middleware.AuthRequired(), handler.SubmitFlag)

		// api.GET("/scoreboard", handler.GetScoreboard)
		api.GET("/ping", handler.Ping) // just for test
	}

	// 管理端
	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	{
		admin.POST("/challenges", handler.AdminCreateChallenge)
		admin.PUT("/challenges/:id", handler.AdminUpdateChallenge)
		admin.DELETE("/challenges/:id", handler.AdminDeleteChallenge)
		admin.GET("/challenges/:id/solves", handler.AdminListSolves)
	}

	return r
}
