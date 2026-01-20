package service

import (
	"time"

	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
)

// CreateCompetition 创建比赛/练习场
func CreateCompetition(title, desc string, cType int, start, end time.Time) error {
	comp := model.Competition{
		Title:       title,
		Description: desc,
		Type:        cType, // 0:比赛, 1:练习
		StartTime:   start,
		EndTime:     end,
	}
	return db.DB.Create(&comp).Error
}

// GetCompetitions 获取所有比赛列表
// 实际开发中可能需要分页，或者只显示“进行中”的，这里先查全部
func GetCompetitions() ([]model.Competition, error) {
	var list []model.Competition
	// 按开始时间倒序，最新的在最前
	err := db.DB.Order("start_time desc").Find(&list).Error
	return list, err
}