// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-01-21
// 描述: 比赛相关的业务逻辑服务
// ----------------------------------------------------------------------------
package service

import (
	"errors"
	"log"
	"time"

	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
)

// CreateCompetition 创建比赛/练习场
func CreateCompetition(title, desc string, cType, mode int, start, end time.Time, adminID uint) error {
	comp := model.Competition{
		Title:       title,
		Description: desc,
		Type:        cType,
		StartTime:   start,
		EndTime:     end,
		Mode:        mode,
	}
	err := db.DB.Create(&comp).Error
	if err == nil {
		log.Printf("[审计] 管理员 %d 创建了比赛: %s (ID: %d)", adminID, title, comp.ID)
	}
	return err
}

// GetCompetitions 获取所有比赛列表
func GetCompetitions() ([]model.Competition, error) {
	var list []model.Competition
	err := db.DB.Order("start_time desc").Find(&list).Error
	return list, err
}

// UpdateCompetition 更新比赛信息
func UpdateCompetition(compID uint, title, desc string, cType, mode int,
	start, end time.Time, adminID uint) error {

	var comp model.Competition
	if err := db.DB.First(&comp, compID).Error; err != nil {
		return errors.New("比赛不存在")
	}

	updates := map[string]interface{}{
		"title":       title,
		"description": desc,
		"type":        cType,
		"mode":        mode,
		"start_time":  start,
		"end_time":    end,
	}

	err := db.DB.Model(&comp).Updates(updates).Error
	if err == nil {
		log.Printf("[审计] 管理员 %d 更新了比赛: %s (ID: %d)", adminID, title, compID)
	}
	return err
}

// DeleteCompetition 删除比赛
func DeleteCompetition(compID uint, adminID uint) error {
	var comp model.Competition
	if err := db.DB.First(&comp, compID).Error; err != nil {
		return errors.New("比赛不存在")
	}

	// 检查比赛是否已有题目，如果有则不允许删除
	var count int64
	db.DB.Model(&model.Challenge{}).Where("competition_id = ?", compID).Count(&count)
	if count > 0 {
		return errors.New("比赛已有题目，无法删除")
	}

	err := db.DB.Delete(&comp).Error
	if err == nil {
		log.Printf("[审计] 管理员 %d 删除了比赛: %s (ID: %d)", adminID, comp.Title, compID)
	}
	return err
}
