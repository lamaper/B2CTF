// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-02-08
// 最后修改: 2026-02-08
// 描述: 权限相关的中间件和工具函数
// ----------------------------------------------------------------------------
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminOnly 中间件：检查用户是否为管理员
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 context 中获取 role（由 JWTAuth 中间件设置）
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未登录",
			})
			c.Abort()
			return
		}

		// 检查是否为管理员
		roleStr, ok := role.(string)
		if !ok || roleStr != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"code": 403,
				"msg":  "没有权限执行这个操作，仅管理员可访问",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// GetUserID 从 context 中安全地获取用户 ID
func GetUserID(c *gin.Context) (uint, bool) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		return 0, false
	}

	// 类型转换处理 float64（来自 JWT claims）
	userIDFloat, ok := userIDValue.(float64)
	if !ok {
		return 0, false
	}

	return uint(userIDFloat), true
}

// GetUserRole 从 context 中安全地获取用户角色
func GetUserRole(c *gin.Context) (string, bool) {
	roleValue, exists := c.Get("role")
	if !exists {
		return "", false
	}

	role, ok := roleValue.(string)
	if !ok {
		return "", false
	}

	return role, true
}

// IsAdmin 检查用户是否为管理员
func IsAdmin(c *gin.Context) bool {
	role, exists := c.Get("role")
	if !exists {
		return false
	}

	roleStr, ok := role.(string)
	return ok && roleStr == "admin"
}
