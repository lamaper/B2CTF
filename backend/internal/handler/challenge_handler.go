// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-01-21
// 描述: 题目相关的HTTP处理函数
// ----------------------------------------------------------------------------
package handler

import (
	"B2CTF/backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateChallengeRequest 接收前端参数
type CreateChallengeRequest struct {
	Title         string   `json:"title" binding:"required"`
	Category      string   `json:"category" binding:"required"`
	Description   string   `json:"description"`
	Flag          string   `json:"flag" binding:"required"`
	Score         int      `json:"score" binding:"required"`
	CompetitionID uint     `json:"competition_id" binding:"required"` // 必填：关联比赛
	Attachment    string   `json:"attachment"`                        // 选填：附件路径
	Tags          []string `json:"tags"`                              // 选填：标签
}

// CreateChallenge 创建题目接口
func CreateChallenge(c *gin.Context) {
	// 鉴权：只有 admin 能创建
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "权限不足"})
		return
	}

	var req CreateChallengeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateChallenge(
		req.Title,
		req.Category,
		req.Description,
		req.Flag,
		req.Score,
		req.CompetitionID, // 传入 ID
		req.Attachment,
		req.Tags,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "题目创建成功"})
}

// ListChallenges 获取列表接口
func ListChallenges(c *gin.Context) {
	challenges, err := service.GetAllChallenges()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": challenges,
	})
}
