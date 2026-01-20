package model

import (
	"time"
)

type Challenge struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Category    string `gorm:"type:varchar(50);not null" json:"category"` // Web, Pwn, Crypto, Misc...
	Description string `gorm:"type:text" json:"description"`
	Score       int    `gorm:"default:1000" json:"score"`
	Flag        string `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt   time.Time
    Tags      []string `gorm:"type:json;serializer:json" json:"tags"` 
	Competition string `gorm:"type:varchar(255)" json:"competition"`
    Type int `gorm:"default:0" json:"type"` // 1: 静态题, 2: 动态容器题
    // 动态题配置 (静态题留空)
    DockerImage  string `json:"docker_image"`  // 例如: "nginx:alpine"
    InternalPort int    `json:"internal_port"` // 例如: 80
    //附件
    Attachment string `gorm:"type:varchar(255)" json:"attachment"`
}
