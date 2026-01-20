package service

import (
	"errors"
	"strings"
	"time"

	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
	"gorm.io/gorm"
)

// SubmitFlag 提交 Flag
func SubmitFlag(userID uint, challengeID uint, submittedFlag string) (bool, error) {
	// 1. 查题目是否存在
	var chal model.Challenge
	if err := db.DB.First(&chal, challengeID).Error; err != nil {
		return false, errors.New("题目不存在")
	}

	// 2. 查题目所属的比赛信息 (为了检查时间)
	var comp model.Competition
	if err := db.DB.First(&comp, chal.CompetitionID).Error; err != nil {
		return false, errors.New("题目所属的比赛不存在")
	}

	// 3. 校验时间逻辑
	now := time.Now()
	if comp.Type == 0 { // 如果是限时比赛
		if now.Before(comp.StartTime) {
			return false, errors.New("比赛尚未开始")
		}
		if now.After(comp.EndTime) {
			return false, errors.New("比赛已结束")
		}
	}
	// 如果 comp.Type == 1 (练习模式)，则直接跳过时间检查

	// 4. 查重：检查用户是否已经做过这道题
	var count int64
	db.DB.Model(&model.Solve{}).Where("user_id = ? AND challenge_id = ?", userID, challengeID).Count(&count)
	if count > 0 {
		return false, errors.New("你已经解出过这道题了")
	}

	// 5. 比对 Flag (去除首尾空格)
	if strings.TrimSpace(submittedFlag) != chal.Flag {
		return false, nil // Flag 错误，返回 false 但不是系统 error
	}

	// 6. 开启事务：写入解题记录 + 更新用户总分
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// A. 创建 Solve 记录
		solve := model.Solve{
			UserID:        userID,
			ChallengeID:   challengeID,
			CompetitionID: chal.CompetitionID, // 冗余存一份，方便统计
			Score:         chal.Score,
		}
		if err := tx.Create(&solve).Error; err != nil {
			return err
		}

		// B. 用户加分
		if err := tx.Model(&model.User{}).Where("id = ?", userID).
			UpdateColumn("score", gorm.Expr("score + ?", chal.Score)).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return true, err // 数据库写入失败
	}

	return true, nil
}