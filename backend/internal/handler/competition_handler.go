// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-01-21
// 描述: 比赛相关的HTTP处理函数
// ----------------------------------------------------------------------------
package handler

import (
	"net/http"
	"time"

	"B2CTF/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// CreateCompetitionRequest 请求参数
type CreateCompetitionRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Type        int       `json:"type"`       // 0 或 1
	StartTime   time.Time `json:"start_time"` // 格式：2026-01-20T15:04:05Z
	EndTime     time.Time `json:"end_time"`
}

// CreateCompetition 创建接口
func CreateCompetition(c *gin.Context) {
	// 1. 鉴权：只有管理员能建
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "只有管理员可以创建比赛"})
		return
	}

	// 2. 绑定参数
	var req CreateCompetitionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误: " + err.Error()})
		return
	}

	// 3. 调用 Service
	err := service.CreateCompetition(req.Title, req.Description, req.Type, req.StartTime, req.EndTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "比赛创建成功"})
}

// ListCompetitions 列表接口
func ListCompetitions(c *gin.Context) {
	list, err := service.GetCompetitions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": list,
	})
}