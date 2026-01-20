package middleware

import (
	"net/http"
	"strings"

	"B2CTF/backend/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取请求头中的 Authorization 字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未登录，请提供 Token",
			})
			c.Abort() // 拦截请求，不再往下执行
			return
		}

		// 2. 格式校验：Token 通常以 "Bearer " 开头
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Token 格式错误",
			})
			c.Abort()
			return
		}

		// 3. 解析 Token
		tokenString := parts[1]
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Token 无效或已过期",
			})
			c.Abort()
			return
		}

		// 4. 将解析出的 UserID 存入上下文 Context 中
		// 这样后续的处理函数（Handler）就能知道是谁在请求了
		c.Set("userID", claims["user_id"])
		c.Set("role", claims["role"])

		// 5. 放行
		c.Next()
	}
}
