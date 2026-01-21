// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-01-21
// 描述: Flag提交相关的HTTP处理函数
// ----------------------------------------------------------------------------
package handler

import (
	"net/http"
	"B2CTF/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type SubmitRequest struct {
	ChallengeID uint   `json:"challenge_id" binding:"required"`
	Flag        string `json:"flag" binding:"required"`
}

// SubmitFlag 提交 Flag 接口
func SubmitFlag(c *gin.Context) {
	var req SubmitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从 Token 获取 UserID (注意类型断言)
	userIDFloat, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}
	userID := uint(userIDFloat.(float64))

	isCorrect, err := service.SubmitFlag(userID, req.ChallengeID, req.Flag)

	// 处理业务逻辑报错 (如: 时间不对、重复提交、题目不存在)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	if isCorrect {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Flag 正确！得分！"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "Flag 错误"})
	}
}