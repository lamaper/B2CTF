package service

import (
	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
)

// CreateChallenge 创建题目
// 参数列表较长，建议按顺序对应 handler 里的传参
func CreateChallenge(title, category, desc, flag string, score int, compID uint, attachment string, tags []string) error {
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