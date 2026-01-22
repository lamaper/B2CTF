package model

import "gorm.io/gorm"

type Challenge struct {
	gorm.Model
	
	// 关联字段：这道题属于哪个比赛
	CompetitionID uint     `json:"competition_id"` 

	Title         string   `gorm:"type:varchar(255);not null" json:"title"`
	Category      string   `gorm:"type:varchar(50);not null" json:"category"`
	Description   string   `gorm:"type:text" json:"description"`
	Score         int      `gorm:"default:100" json:"score"`
	
	// Flag 绝对不能传给前端，所以 json:"-"
	Flag          string   `gorm:"type:varchar(255);not null" json:"-"`
	
	// 附件下载链接
	Attachment    string   `gorm:"type:varchar(255)" json:"attachment"`
	
	// 标签：自动将 Go 的 []string 转为数据库的 JSON 字符串存储
	Tags          []string `gorm:"type:json;serializer:json" json:"tags"`
	SolvedCount int `gorm:"default:0" json:"solved_count"`
}