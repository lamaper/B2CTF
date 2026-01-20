package handler

import (
	"net/http"
	"B2CTF/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// CreateChallengeRequest 创建题目的参数
type CreateChallengeRequest struct {
	Title       string `json:"title" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description" binding:"required"`
	Flag        string `json:"flag" binding:"required"`
	Score       int    `json:"score" binding:"required"`
	Competition string `json:"competition" binding:"required"`
	Tags        []string `json:"tags" binding:"required"`
	Type        int `json:"type" binding:"required"`
	DockerImage string `json:"docker_image"`
	InternalPort int `json:"internal_port"`
	Attachment string `json:"attachment"`
}

// CreateChallenge 接口
func CreateChallenge(c *gin.Context) {
	// TODO 1. 这里应该加一个权限检查：只有管理员才能创建
	// role := c.GetString("role") // 从 Token 解析出来的角色
	// if role != "admin" {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "只有管理员可以创建题目"})
	// 	return
	// }

	var req CreateChallengeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateChallenge(req.Title, req.Category, 
		req.Description, req.Flag, req.Score, 
		req.Competition, req.Tags, req.Type, 
		req.DockerImage, 
		req.InternalPort,req.Attachment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "题目创建成功"})
}

// ListChallenges 接口
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