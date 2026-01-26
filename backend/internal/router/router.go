// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-17
// 最后修改: 2026-01-22
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

	r.Static("/upload", "./uploads")
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
		
	}

	// 创建一个使用了 JWTAuth 中间件的路由组
	protected := r.Group("/")
	protected.Use(middleware.JWTAuth())
	{
		// 这里面的所有接口，都必须带 Token 才能访问
		protected.GET("/user/profile", handler.GetUserProfile)
		protected.GET("/challenge", handler.ListChallenges)
		protected.POST("/challenge", handler.CreateChallenge)

		// 注意！仅供测试！后续添加管理员策略组后转移这个接口
		protected.POST("/upload", handler.UploadFile)
		protected.POST("/competitions", handler.CreateCompetition)
		protected.GET("/competitions", handler.ListCompetitions)
		protected.POST("/container/launch", handler.LaunchContainer)
		protected.POST("/container/terminate", handler.TerminateContainer)
	}

	teamGroup := protected.Group("/team")
	{
    	teamGroup.POST("/create", handler.CreateTeam)
    	teamGroup.POST("/join", handler.JoinTeam)
    	teamGroup.GET("/my", handler.GetMyTeam) // 前端一进“我的战队”页面就调这个
		
	}

	return r
}
