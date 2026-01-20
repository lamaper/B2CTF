package service

import (
	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
)

// CreateChallenge 创建新题目
func CreateChallenge(title, category, desc, flag string, 
	score int, competition string, tags []string, type_ int, 
	dockerImage string, internalPort int, attachment string) error {
	if type_ == 1 {
		dockerImage = "null" ;
		internalPort = -1 ;
	}

	challenge := model.Challenge{
		Title:       title,
		Category:    category,
		Description: desc,
		Flag:        flag,
		Score:       score,
		Tags:        tags,
		Competition: competition,
		Type:        type_,
		DockerImage: dockerImage,
		InternalPort: internalPort,
		Attachment: attachment,
	}

	return db.DB.Create(&challenge).Error
}

// GetAllChallenges 获取所有题目列表
func GetAllChallenges() ([]model.Challenge, error) {
	var challenges []model.Challenge
	// 当 Handler 层把它转成 JSON 发给前端时，Flag 字段会被自动隐藏。
	result := db.DB.Order("id desc").Find(&challenges)
	return challenges, result.Error
}
