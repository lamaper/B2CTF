// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-17
// 最后修改: 2026-01-17
// 描述: HTTP请求处理函数
// ----------------------------------------------------------------------------
package handler

import (
	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
	"B2CTF/backend/internal/service"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
//	成功: {"code": 200, "msg": "登录成功", "data": {"token": "jwt_token", "role": "admin"}}
//	失败: {"error": "错误信息"} (401 状态码)
func UserLogin(c *gin.Context) {
	var req UserLoginRequest

	// 绑定并验证请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用Service层的登录逻辑
	token, role, err := service.Login(req.Username, req.Password)
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
			"role":  role,
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

	// 2. 查询数据库获取用户信息
	var user model.User
	result := db.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户信息"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"user_id":    user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"role":       user.Role,
			"created_at": user.CreatedAt,
		},
	})
}

// UploadAvatar 修改头像接口
func UploadAvatar(c *gin.Context) {
	// 1. 获取当前登录用户的 ID
	userIDFloat, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}
	userID := uint(userIDFloat.(float64))

	// 2. 获取上传的文件
	file, err := c.FormFile("file") // 前端 FormData 的 key 必须是 "file"
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传图片文件"})
		return
	}

	// 3. 校验文件类型 (简单的后缀名校验)
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}
	if !allowExts[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "只支持 jpg/png/gif 格式的图片"})
		return
	}

	// 4. 校验文件大小 (例如限制 2MB)
	if file.Size > 2*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "图片大小不能超过 2MB"})
		return
	}

	// 5. 准备存储路径
	// 目录: ./uploads/avatars
	saveDir := "./uploads/avatars"
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		os.MkdirAll(saveDir, os.ModePerm)
	}

	// 生成随机文件名: uuid.png
	newFileName := uuid.New().String() + ext
	dst := filepath.Join(saveDir, newFileName)

	// 6. 保存文件到磁盘
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存图片失败"})
		return
	}

	// 7. 生成对外访问的 URL
	avatarURL := "/uploads/avatars/" + newFileName

	// 8. 调用 Service 更新数据库
	if err := service.UpdateAvatar(userID, avatarURL); err != nil {
		// 如果数据库更新失败，最好把刚才上传的文件删掉，避免垃圾文件堆积
		os.Remove(dst)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新头像数据失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "头像修改成功",
		"data": gin.H{
			"avatar": avatarURL, // 返回新头像给前端立即回显
		},
	})
}
