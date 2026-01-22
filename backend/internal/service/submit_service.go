// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-01-21
// 描述: Flag提交相关的业务逻辑服务
// ----------------------------------------------------------------------------
package service

import (
	"errors"
	"strings"
	"time"

	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
	"gorm.io/gorm"
)

// SubmitFlag 提交 Flag (最终完整版)
func SubmitFlag(userID uint, challengeID uint, submittedFlag string) (bool, error) {
	// 1. 查题目
	var chal model.Challenge
	if err := db.DB.First(&chal, challengeID).Error; err != nil {
		return false, errors.New("题目不存在")
	}

	// 2. 查比赛
	var comp model.Competition
	if err := db.DB.First(&comp, chal.CompetitionID).Error; err != nil {
		return false, errors.New("比赛不存在")
	}

	// 3. 查用户 (新增：为了获取 TeamID)
	var user model.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return false, errors.New("用户不存在")
	}

	// 4. 时间校验 (限时赛)
	now := time.Now()
	if comp.Type == 0 {
		if now.Before(comp.StartTime) {
			return false, errors.New("比赛未开始")
		}
		if now.After(comp.EndTime) {
			return false, errors.New("比赛已结束")
		}
	}

	// =========================================================
	// 5. 核心修改：查重逻辑 (区分个人赛/团队赛)
	// =========================================================
	
	if comp.Mode == 1 { 
		// --- 团队赛模式 (Mode=1) ---
		
		// A. 必须先加入战队
		if user.TeamID == 0 {
			return false, errors.New("本场是团队赛，请先创建或加入一个战队")
		}

		// B. 检查“全队”是否做过
		// 逻辑：在 solves 表里找，有没有任何一个记录，它的 user_id 对应的 team_id 是当前用户的 team_id
		var count int64
		// SQL: SELECT count(*) FROM solves JOIN users ON solves.user_id = users.id 
		//      WHERE users.team_id = ? AND solves.challenge_id = ?
		db.DB.Table("solves").
			Joins("INNER JOIN users ON users.id = solves.user_id").
			Where("users.team_id = ? AND solves.challenge_id = ?", user.TeamID, challengeID).
			Count(&count)
		
		if count > 0 {
			return false, errors.New("你的队友已经解出过这道题了")
		}

	} else {
		// --- 个人赛/练习模式 (Mode=0) ---
		
		// 检查“自己”是否做过
		var count int64
		db.DB.Model(&model.Solve{}).
			Where("user_id = ? AND challenge_id = ?", userID, challengeID).
			Count(&count)
		
		if count > 0 {
			return false, errors.New("你已经解出过这道题了")
		}
	}

	// =========================================================

	// 6. 比对 Flag
	if strings.TrimSpace(submittedFlag) != chal.Flag {
		// 简单的防爆破可以在这里加，比如记录错误次数
		return false, nil 
	}

	// 7. 开启事务：写入记录 + 加分
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// A. 创建 Solve 记录
		solve := model.Solve{
			UserID:        userID,
			ChallengeID:   challengeID,
			CompetitionID: chal.CompetitionID,
			Score:         chal.Score,
		}
		if err := tx.Create(&solve).Error; err != nil {
			return err
		}

		// B. 个人加分 (全局总分)
		// 即使是团队赛，个人贡献分也应该涨，这不影响比赛排名(比赛排名是算sum)
		if err := tx.Model(&model.User{}).Where("id = ?", userID).
			UpdateColumn("score", gorm.Expr("score + ?", chal.Score)).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return true, err
	}

	return true, nil
}