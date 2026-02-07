// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-02-08
// 描述: 题目相关的HTTP处理函数
// ----------------------------------------------------------------------------
package handler

import (
	"B2CTF/backend/internal/middleware"
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
	IsDynamic     bool     `json:"is_dynamic"`                        // 选填：是否为动态题
	ImageName     string   `json:"image_name"`                        // 动态题：Docker 镜像名
	ContainerPort int      `json:"container_port"`                    // 动态题：容器内部端口
	MemoryLimit   string   `json:"memory_limit"`                      // 动态题：内存限制
	CPULimit      string   `json:"cpu_limit"`                         // 动态题：CPU 限制
}

// CreateChallenge 创建题目接口（管理员专用）
func CreateChallenge(c *gin.Context) {
	// 1. 权限检查：只有管理员能创建题目
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "没有权限，仅管理员可以创建题目",
		})
		return
	}

	// 2. 获取管理员信息
	adminID, _ := middleware.GetUserID(c)

	// 3. 绑定并验证参数
	var req CreateChallengeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	// 4. 数据验证
	if req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "题目标题不能为空",
		})
		return
	}

	if req.Score <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "题目分数必须大于 0",
		})
		return
	}

	// 5. 动态题目参数验证
	if req.IsDynamic {
		if req.ImageName == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "动态题目必须提供 Docker 镜像名称",
			})
			return
		}
		if req.ContainerPort <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "动态题目必须提供有效的容器端口",
			})
			return
		}
	}

	// 6. 调用 Service 创建题目
	err := service.CreateChallenge(
		req.Title,
		req.Category,
		req.Description,
		req.Flag,
		req.Score,
		req.CompetitionID,
		req.Attachment,
		req.Tags,
		req.IsDynamic,
		req.ImageName,
		req.ContainerPort,
		req.MemoryLimit,
		req.CPULimit,
		adminID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建题目失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "题目创建成功",
	})
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
