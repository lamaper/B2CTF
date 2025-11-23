package middleware

import "github.com/gin-gonic/gin"

// 先写个啥都不做的 AuthRequired，保证能跑起来
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 这里以后再加 JWT 校验逻辑
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 这里以后再判断 role == admin
		c.Next()
	}
}
