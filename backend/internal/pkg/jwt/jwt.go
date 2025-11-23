package middleware

import "github.com/gin-gonic/gin"

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        // TODO: 从 header 拿 token -> 解析 -> 设置 userID / role 到 context
        c.Next()
    }
}

func AdminOnly() gin.HandlerFunc {
    return func(c *gin.Context) {
        // TODO: 从 context 取 role，判断是否 admin
        c.Next()
    }
}
