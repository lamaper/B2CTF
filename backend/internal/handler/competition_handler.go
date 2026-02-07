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

	"B2CTF/backend/internal/middleware"
	"B2CTF/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// CreateCompetitionRequest 请求参数
type CreateCompetitionRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Type        int       `json:"type"`       // 0 或 1
	Mode        int       `json:"mode"`       // 0: 个人赛, 1: 团队赛
	StartTime   time.Time `json:"start_time"` // 格式：2026-01-20T15:04:05Z
	EndTime     time.Time `json:"end_time"`
}

// CreateCompetition 创建比赛接口（管理员专用）
func CreateCompetition(c *gin.Context) {
	// 1. 权限检查：只有管理员能创建
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "没有权限，仅管理员可以创建比赛",
		})
		return
	}

	// 2. 获取管理员信息（用于审计日志）
	adminID, _ := middleware.GetUserID(c)

	// 3. 绑定并验证参数
	var req CreateCompetitionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	// 4. 参数验证
	if req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "比赛标题不能为空",
		})
		return
	}

	if req.StartTime.After(req.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "开始时间不能晚于结束时间",
		})
		return
	}

	// 5. 调用 Service 创建比赛
	err := service.CreateCompetition(req.Title, req.Description, req.Type, req.Mode, req.StartTime, req.EndTime, adminID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建比赛失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "比赛创建成功",
	})
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
