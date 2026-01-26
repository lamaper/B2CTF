package handler

import (
	"net/http"
	"B2CTF/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type LaunchReq struct {
	ChallengeID uint `json:"challenge_id" binding:"required"`
	ContainerLimit int `json:"container_limit"` // 选填
}

// LaunchContainer 启动靶机接口
func LaunchContainer(c *gin.Context) {
	// 1. 安全获取当前用户 ID
	// 不要用 userIDFloat, _ := ... 这种忽略错误的写法
	val, exists := c.Get("userID")
	if !exists {
		// 如果取不到，说明中间件没起作用，或者没登录
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录或Token无效"})
		return
	}
	
	// 安全断言
	userIDFloat, ok := val.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户ID解析错误"})
		return
	}
	userID := uint(userIDFloat)

	// 2. 解析参数
	var req LaunchReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. 调用 Service 启动容器
	container, err := service.LaunchContainer(userID, req.ChallengeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "启动失败: " + err.Error()})
		return
	}

	// 4. 返回连接信息
	// 这里的 Host 暂时写 localhost，或者你可以获取服务器公网IP
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "靶机启动成功",
		"data": gin.H{
			"container_id": container.ID,
			"host":         "127.0.0.1", 
			"port":         container.Port,
			"expire_time":  container.ExpireTime,
			"remaining":    3600, // 剩余秒数，暂时写死
		},
	})
}

// TerminateContainer 销毁接口
func TerminateContainer(c *gin.Context) {
    userIDFloat, _ := c.Get("userID")
    userID := uint(userIDFloat.(float64))

    var req LaunchReq // 复用之前的结构体，只需要 challenge_id
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := service.TerminateContainer(userID, req.ChallengeID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "销毁失败: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "靶机已销毁"})
}