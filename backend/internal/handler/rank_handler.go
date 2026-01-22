package handler

import (
	"net/http"
	"strconv"

	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
	"B2CTF/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// GetGlobalRank 全局榜单接口
func GetGlobalRank(c *gin.Context) {
	rankList, err := service.GetGlobalRank()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取排名失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": rankList,
	})
}

// GetCompRank 获取比赛排行榜 (智能切换)
func GetCompRank(c *gin.Context) {
	compIDStr := c.Param("id")
	compID, err := strconv.Atoi(compIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的比赛ID"})
		return
	}
	uid := uint(compID)

	// 1. 先查比赛信息，看它是个人赛还是团队赛
	var comp model.Competition
	if err := db.DB.First(&comp, uid).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "比赛不存在"})
		return
	}

	// 2. 根据模式走不同的分支
	if comp.Mode == 1 {
		// --- 分支 A：团队赛 ---
		rankList, err := service.GetTeamRank(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取团队排名失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"mode": "team", // 告诉前端这是团队榜
			"data": rankList,
		})

	} else {
		// --- 分支 B：个人赛 (默认) ---
		rankList, err := service.GetCompetitionRank(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取个人排名失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"mode": "individual", // 告诉前端这是个人榜
			"data": rankList,
		})
	}
}