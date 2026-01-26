// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-01-21
// 描述: 题目相关的业务逻辑服务
// ----------------------------------------------------------------------------
package service

import (
	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
)

// CreateChallenge 创建题目
// 参数列表较长，建议按顺序对应 handler 里的传参
func CreateChallenge(title, category, desc, flag string,
	score int, compID uint, attachment string,
	tags []string) error {
	challenge := model.Challenge{
		Title:         title,
		Category:      category,
		Description:   desc,
		Flag:          flag,
		Score:         score,
		CompetitionID: compID,     // 关联比赛ID
		Attachment:    attachment, // 附件URL
		Tags:          tags,       // 标签数组
	}

	return db.DB.Create(&challenge).Error
}

// GetAllChallenges 获取所有题目
// 实际生产中通常会写 GetChallengesByCompetitionID(compID uint)，这里先简化逻辑
func GetAllChallenges() ([]model.Challenge, error) {
	var challenges []model.Challenge
	// 按照 ID 倒序排列
	result := db.DB.Order("id desc").Find(&challenges)
	return challenges, result.Error
}

func GetChallengeByID(chalID uint) (*model.Challenge, error) {
	var chal model.Challenge
	err := db.DB.First(&chal, chalID).Error
	if err != nil {
		return nil, err
	}
	return &chal, nil
}

func DeleteChallenge(chalID uint) error {
	return db.DB.Delete(&model.Challenge{}, chalID).Error
}

func SetChallengeInfo(chalID uint, title, category, desc, flag string,
	score int, compID uint, attachment string,
	tags []string) error {
	var chal model.Challenge
	if err := db.DB.First(&chal, chalID).Error; err != nil {
		return err
	}

	chal.Title = title
	chal.Category = category
	chal.Description = desc
	chal.Flag = flag
	chal.Score = score
	chal.CompetitionID = compID
	chal.Attachment = attachment
	chal.Tags = tags

	return db.DB.Save(&chal).Error
}
