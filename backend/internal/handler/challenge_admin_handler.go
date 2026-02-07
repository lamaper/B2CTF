// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-02-08
// 最后修改: 2026-02-08
// 描述: 题目管理相关的HTTP处理函数（管理员专用）
// ----------------------------------------------------------------------------
package handler

import (
	"net/http"
	"strconv"

	"B2CTF/backend/internal/middleware"
	"B2CTF/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// UpdateChallengeRequest 更新题目的请求参数
type UpdateChallengeRequest struct {
	Title         string   `json:"title"`
	Category      string   `json:"category"`
	Description   string   `json:"description"`
	Flag          string   `json:"flag"`
	Score         int      `json:"score"`
	Attachment    string   `json:"attachment"`
	Tags          []string `json:"tags"`
	IsDynamic     bool     `json:"is_dynamic"`
	ImageName     string   `json:"image_name"`
	ContainerPort int      `json:"container_port"`
	MemoryLimit   string   `json:"memory_limit"`
	CPULimit      string   `json:"cpu_limit"`
}

// UpdateChallenge 更新题目接口（管理员专用）
func UpdateChallenge(c *gin.Context) {
	// 权限检查由路由中的 AdminOnly 中间件完成

	// 获取题目 ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的题目 ID",
		})
		return
	}

	// 获取管理员信息
	adminID, _ := middleware.GetUserID(c)

	// 绑定参数
	var req UpdateChallengeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	// 调用服务更新题目
	err = service.UpdateChallenge(uint(id), req.Title, req.Category, req.Description,
		req.Flag, req.Score, req.Attachment, req.Tags, req.IsDynamic,
		req.ImageName, req.ContainerPort, req.MemoryLimit, req.CPULimit, adminID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "更新失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "题目更新成功",
	})
}

// DeleteChallenge 删除题目接口（管理员专用）
func DeleteChallenge(c *gin.Context) {
	// 权限检查由路由中的 AdminOnly 中间件完成

	// 获取题目 ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的题目 ID",
		})
		return
	}

	// 获取管理员信息
	adminID, _ := middleware.GetUserID(c)

	// 调用服务删除题目
	err = service.DeleteChallenge(uint(id), adminID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "删除失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "题目删除成功",
	})
}
