// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-01-21
// 描述: 比赛数据模型
// ----------------------------------------------------------------------------
package model

import (
	"time"
	"gorm.io/gorm"
)

type Competition struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	
	// Type 比赛模式
	// 0: 限时比赛 (需要校验时间)
	// 1: 常驻题库/练习区 (不校验时间)
	Type        int       `gorm:"default:0" json:"type"` 
	
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	
	// Mode 比赛模式
	// 0: 个人赛 (Individual) - 排行榜显示 User
	// 1: 团队赛 (Team) - 排行榜显示 Team
	Mode        int       `gorm:"default:0" json:"mode"` 

	// 靶机数量限制
	// 允许该比赛/练习场的用户同时开启几个靶机
	// 默认值为 1
	ContainerLimit int    `gorm:"default:1" json:"container_limit"` 
}