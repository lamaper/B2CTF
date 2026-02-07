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

	r.Static("/uploads", "./uploads")
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
		// --- 用户相关 ---
		protected.GET("/user/profile", handler.GetUserProfile)

		// --- 题目相关（查看不需权限，但创建需要管理员）---
		protected.GET("/challenge", handler.ListChallenges)

		// --- 比赛相关（查看不需权限，但创建需要管理员）---
		protected.GET("/competitions", handler.ListCompetitions)

		// --- 容器相关 ---
		protected.POST("/container/launch", handler.LaunchContainer)
		protected.POST("/container/terminate", handler.TerminateContainer)

		// --- Flag 提交 ---
		protected.POST("/submit", handler.SubmitFlag)

		// --- 排行榜 ---
		protected.GET("/rank", handler.GetGlobalRank)
		protected.GET("/rank/:id", handler.GetCompRank)

		// --- 团队相关 ---
		teamGroup := protected.Group("/team")
		{
			teamGroup.POST("/create", handler.CreateTeam)
			teamGroup.POST("/join", handler.JoinTeam)
			teamGroup.GET("/my", handler.GetMyTeam)
		}
	}

	// ============ 管理员专用接口 ============
	admin := r.Group("/admin")
	admin.Use(middleware.JWTAuth())
	admin.Use(middleware.AdminOnly())
	{
		// --- 题目管理 ---
		admin.POST("/challenge", handler.CreateChallenge)
		admin.PUT("/challenge/:id", handler.UpdateChallenge)
		admin.DELETE("/challenge/:id", handler.DeleteChallenge)

		// --- 比赛管理 ---
		admin.POST("/competition", handler.CreateCompetition)
		admin.PUT("/competition/:id", handler.UpdateCompetition)
		admin.DELETE("/competition/:id", handler.DeleteCompetition)

		// --- 文件上传 ---
		admin.POST("/upload", handler.UploadFile)

		// --- 用户管理 ---
		admin.GET("/users", handler.ListUsers)
		admin.PUT("/user/:id/role", handler.SetUserRole)
		admin.DELETE("/user/:id", handler.DeleteUser)

		// --- 比赛统计 ---
		admin.GET("/statistics/competition/:id", handler.GetCompetitionStatistics)
	}

	return r
}
