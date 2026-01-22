package handler

import (
	"net/http"
	"B2CTF/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type CreateTeamReq struct {
	Name string `json:"name" binding:"required"`
	Desc string `json:"description"`
}

type JoinTeamReq struct {
	Token string `json:"token" binding:"required"`
}

// CreateTeam 创建
func CreateTeam(c *gin.Context) {
	userID := uint(c.GetFloat64("userID")) // 辅助函数获取ID
	
	var req CreateTeamReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := service.CreateTeam(userID, req.Name, req.Desc)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200, 
		"msg": "队伍创建成功",
		"data": gin.H{"token": token},
	})
}

// JoinTeam 加入
func JoinTeam(c *gin.Context) {
	userID := uint(c.GetFloat64("userID"))
	
	var req JoinTeamReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.JoinTeam(userID, req.Token); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "加入成功"})
}

// GetMyTeam 获取信息
func GetMyTeam(c *gin.Context) {
	userID := uint(c.GetFloat64("userID"))
	
	team, err := service.GetMyTeamInfo(userID)
	if err != nil {
		if err.Error() == "NO_TEAM" {
			c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil}) // 没有队伍返回 null
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": team})
}