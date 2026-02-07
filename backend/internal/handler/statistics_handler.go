// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-02-08
// 最后修改: 2026-02-08
// 描述: 比赛统计相关的HTTP处理函数（管理员专用）
// ----------------------------------------------------------------------------
package handler

import (
	"net/http"
	"strconv"

	"B2CTF/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// GetCompetitionStatistics 获取比赛统计信息接口（管理员专用）
func GetCompetitionStatistics(c *gin.Context) {
	// 权限检查由路由中的 AdminOnly 中间件完成

	// 获取比赛 ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的比赛 ID",
		})
		return
	}

	// 调用服务获取统计信息
	stats, err := service.GetCompetitionStatistics(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "获取统计信息失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": stats,
	})
}
