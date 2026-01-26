package model

import (
	"time"
	"gorm.io/gorm"
)

type Container struct {
	gorm.Model
	
	UserID      uint      `gorm:"index;not null" json:"user_id"`      // 属于哪个用户
	ChallengeID uint      `gorm:"index;not null" json:"challenge_id"` // 属于哪道题
	
	// Docker 里的容器 ID (长字符串，用于停止/删除容器)
	ContainerID string    `gorm:"size:64;uniqueIndex" json:"-"` 
	
	// 映射到宿主机的端口 (比如 32001)
	// 用户通过 http://IP:Port 访问
	Port        int       `json:"port"`
	
	// 本次生成的动态 Flag (这是校验答案的唯一标准！)
	Flag        string    `gorm:"size:255" json:"-"`
	
	// 自动销毁时间 (比如 1 小时后)
	ExpireTime  time.Time `json:"expire_time"`
	
	// 状态: starting, running, destroyed
	Status      string    `gorm:"default:'running'" json:"status"`
}