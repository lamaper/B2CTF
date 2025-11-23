package router

import (
    "github.com/gin-gonic/gin"

    "B2CTF/backend/internal/handler"
    "B2CTF/backend/internal/middleware"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // 公共接口
    api := r.Group("/api")
    {
        api.POST("/auth/register", handler.Register)
        api.POST("/auth/login", handler.Login)

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
