// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-02-08
// 描述: 题目相关的业务逻辑服务
// ----------------------------------------------------------------------------
package service

import (
	"errors"
	"log"

	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
)

// CreateChallenge 创建题目
func CreateChallenge(title, category, desc, flag string,
	score int, compID uint, attachment string,
	tags []string, isDynamic bool, imageName string,
	containerPort int, memoryLimit, cpuLimit string, adminID uint) error {

	challenge := model.Challenge{
		Title:         title,
		Category:      category,
		Description:   desc,
		Flag:          flag,
		Score:         score,
		CompetitionID: compID,
		Attachment:    attachment,
		Tags:          tags,
		IsDynamic:     isDynamic,
		ImageName:     imageName,
		ContainerPort: containerPort,
		MemoryLimit:   memoryLimit,
		CPULimit:      cpuLimit,
	}

	err := db.DB.Create(&challenge).Error
	if err == nil {
		log.Printf("[审计] 管理员 %d 创建了题目: %s (ID: %d)", adminID, title, challenge.ID)
	}
	return err
}

// GetAllChallenges 获取所有题目
func GetAllChallenges() ([]model.Challenge, error) {
	var challenges []model.Challenge
	result := db.DB.Order("id desc").Find(&challenges)
	return challenges, result.Error
}

// GetChallengeByID 根据ID获取题目详情
func GetChallengeByID(chalID uint) (*model.Challenge, error) {
	var chal model.Challenge
	err := db.DB.First(&chal, chalID).Error
	if err != nil {
		return nil, err
	}
	return &chal, nil
}

// UpdateChallenge 更新题目信息
func UpdateChallenge(chalID uint, title, category, desc, flag string,
	score int, attachment string, tags []string, isDynamic bool,
	imageName string, containerPort int, memoryLimit, cpuLimit string, adminID uint) error {

	var chal model.Challenge
	if err := db.DB.First(&chal, chalID).Error; err != nil {
		return errors.New("题目不存在")
	}

	// 更新字段
	updates := map[string]interface{}{
		"title":          title,
		"category":       category,
		"description":    desc,
		"flag":           flag,
		"score":          score,
		"attachment":     attachment,
		"tags":           tags,
		"is_dynamic":     isDynamic,
		"image_name":     imageName,
		"container_port": containerPort,
		"memory_limit":   memoryLimit,
		"cpu_limit":      cpuLimit,
	}

	err := db.DB.Model(&chal).Updates(updates).Error
	if err == nil {
		log.Printf("[审计] 管理员 %d 更新了题目: %s (ID: %d)", adminID, title, chalID)
	}
	return err
}

// DeleteChallenge 删除题目
func DeleteChallenge(chalID uint, adminID uint) error {
	var chal model.Challenge
	if err := db.DB.First(&chal, chalID).Error; err != nil {
		return errors.New("题目不存在")
	}

	err := db.DB.Delete(&chal).Error
	if err == nil {
		log.Printf("[审计] 管理员 %d 删除了题目: %s (ID: %d)", adminID, chal.Title, chalID)
	}
	return err
}
