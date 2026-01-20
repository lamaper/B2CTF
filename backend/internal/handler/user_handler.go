// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-17
// 最后修改: 2026-01-17
// 描述: HTTP请求处理函数
// ----------------------------------------------------------------------------
package handler

import (
	"B2CTF/backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRegisterRequest 用户注册请求结构体
// 用于验证和绑定注册请求的JSON数据

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`    // 用户名（必填）
	Password string `json:"password" binding:"required"`    // 密码（必填）
	Email    string `json:"email" binding:"required,email"` // 邮箱（必填，需符合邮箱格式）
}

// UserLoginRequest 用户登录请求结构体
// 用于验证和绑定登录请求的JSON数据

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"` // 用户名（必填）
	Password string `json:"password" binding:"required"` // 密码（必填）
}

// UserRegister 用户注册接口处理函数
//
// 路由: POST /api/register
//
// 请求体:
//
//	{
//	  "username": "user123",
//	  "password": "password123",
//	  "email": "user@example.com"
//	}
//
// 响应:
//
//	成功: {"code": 200, "msg": "注册成功"}
//	失败: {"code": 400, "msg": "错误信息"}
func UserRegister(c *gin.Context) {
	var req UserRegisterRequest

	// 绑定并验证请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用Service层的注册逻辑
	err := service.Register(req.Username, req.Password, req.Email)
	if err != nil {
		// 注册失败
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

// UserLogin 用户登录接口处理函数
//
// 路由: POST /api/login
//
// 请求体:
//
//	{
//	  "username": "user123",
//	  "password": "password123"
//	}
//
// 响应:
//
//	成功: {"code": 200, "msg": "登录成功", "data": {"token": "jwt_token"}}
//	失败: {"error": "错误信息"} (401 状态码)
func UserLogin(c *gin.Context) {
	var req UserLoginRequest

	// 绑定并验证请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用Service层的登录逻辑
	token, err := service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()}) // 401 未授权
		return
	}

	// 返回成功响应和JWT令牌
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

// GetUserProfile 获取个人信息
func GetUserProfile(c *gin.Context) {
	// 1. 从上下文中取出中间件存进去的 userID
	// 这里的 "userID" 必须和 middleware 里的 c.Set("userID", ...) 字符串一致
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户ID"})
		return
	}

	// 2. 假装查数据库（这里直接返回ID演示）
	// 实际开发中，你会用这个 ID 去 db.DB.First(&user, userID)
	
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"user_id": userID,
			"role":    c.GetString("role"), // 也可以取 role
			"info":    "这是一条只有登录用户才能看到的机密信息",
		},
	})
}