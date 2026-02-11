// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-02-08
// 最后修改: 2026-02-08
// 描述: 比赛管理相关的HTTP处理函数（管理员专用）
// ----------------------------------------------------------------------------
package handler

import (
	"net/http"
	"strconv"
	"time"

	"B2CTF/backend/internal/middleware"
	"B2CTF/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// UpdateCompetitionRequest 更新比赛的请求参数
type UpdateCompetitionRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Type        int       `json:"type"`
	Mode        int       `json:"mode"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

// UpdateCompetition 更新比赛接口（管理员专用）
func UpdateCompetition(c *gin.Context) {
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

	// 获取管理员信息
	adminID, _ := middleware.GetUserID(c)

	// 绑定参数
	var req UpdateCompetitionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	// 验证时间逻辑
	if req.StartTime.After(req.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "开始时间不能晚于结束时间",
		})
		return
	}

	// 调用服务更新比赛
	err = service.UpdateCompetition(uint(id), req.Title, req.Description, req.Type,
		req.Mode, req.StartTime, req.EndTime, adminID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "更新失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "比赛更新成功",
	})
}

// DeleteCompetition 删除比赛接口（管理员专用）
func DeleteCompetition(c *gin.Context) {
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

	// 获取管理员信息
	adminID, _ := middleware.GetUserID(c)

	// 调用服务删除比赛
	err = service.DeleteCompetition(uint(id), adminID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "删除失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "比赛删除成功",
	})
}
