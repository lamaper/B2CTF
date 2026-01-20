package model

import (
	"time"
	"gorm.io/gorm"
)

type Competition struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	// 0: 限时比赛 (需要校验时间)
	// 1: 常驻题库/练习区 (不校验时间)
	Type        int       `gorm:"default:0" json:"type"` 
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}