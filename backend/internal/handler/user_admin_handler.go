// ----------------------------------------------------------------------------
// Copyright (c) 2026 West
// 创建日期: 2026-02-08
// 最后修改: 2026-02-08
// 描述: 用户管理相关的HTTP处理函数（管理员专用）
// ----------------------------------------------------------------------------
package handler

import (
	"net/http"
	"strconv"

	"B2CTF/backend/internal/middleware"
	"B2CTF/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// SetUserRoleRequest 设置用户角色的请求参数
type SetUserRoleRequest struct {
	Role string `json:"role" binding:"required"` // "user" 或 "admin"
}

// ListUsers 列出所有用户接口（管理员专用）
func ListUsers(c *gin.Context) {
	// 权限检查由路由中的 AdminOnly 中间件完成

	// 获取管理员信息（用于审计）
	adminID, _ := middleware.GetUserID(c)

	// 调用服务获取用户列表
	users, err := service.ListAllUsers(adminID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "获取用户列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": users,
	})
}

// SetUserRole 设置用户角色接口（管理员专用）
func SetUserRole(c *gin.Context) {
	// 权限检查由路由中的 AdminOnly 中间件完成

	// 获取用户 ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的用户 ID",
		})
		return
	}

	// 获取管理员信息
	adminID, _ := middleware.GetUserID(c)

	// 防止管理员修改自己的角色
	if uint(id) == adminID {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "不能修改自己的角色",
		})
		return
	}

	// 绑定参数
	var req SetUserRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	// 验证角色值
	if req.Role != "user" && req.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的角色，仅支持 'user' 或 'admin'",
		})
		return
	}

	// 调用服务更新用户角色
	err = service.SetUserRole(uint(id), req.Role, adminID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "更新失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户角色更新成功",
	})
}

// DeleteUser 删除用户接口（管理员专用）
func DeleteUser(c *gin.Context) {
	// 权限检查由路由中的 AdminOnly 中间件完成

	// 获取用户 ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的用户 ID",
		})
		return
	}

	// 获取管理员信息
	adminID, _ := middleware.GetUserID(c)

	// 防止管理员删除自己
	if uint(id) == adminID {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "不能删除自己",
		})
		return
	}

	// 调用服务删除用户
	err = service.DeleteUser(uint(id), adminID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "删除失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户删除成功",
	})
}
