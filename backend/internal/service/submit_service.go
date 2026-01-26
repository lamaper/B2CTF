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

// SubmitFlag 提交 Flag (支持动态靶机版)
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

	// 3. 查用户
	var user model.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return false, errors.New("用户不存在")
	}

	// 4. 时间校验
	now := time.Now()
	if comp.Type == 0 {
		if now.Before(comp.StartTime) {
			return false, errors.New("比赛未开始")
		}
		if now.After(comp.EndTime) {
			return false, errors.New("比赛已结束")
		}
	}

	// 5. 查重逻辑
	if comp.Mode == 1 {
		// --- 团队赛模式 ---
		if user.TeamID == 0 {
			return false, errors.New("本场是团队赛，请先创建或加入一个战队")
		}
		var count int64
		db.DB.Table("solves").
			Joins("INNER JOIN users ON users.id = solves.user_id").
			Where("users.team_id = ? AND solves.challenge_id = ?", user.TeamID, challengeID).
			Count(&count)
		if count > 0 {
			return false, errors.New("你的队友已经解出过这道题了")
		}
	} else {
		// --- 个人赛模式 ---
		var count int64
		db.DB.Model(&model.Solve{}).
			Where("user_id = ? AND challenge_id = ?", userID, challengeID).
			Count(&count)
		if count > 0 {
			return false, errors.New("你已经解出过这道题了")
		}
	}

	// 6. 获取正确答案
	// 默认使用题目表里的静态 Flag
	correctFlag := chal.Flag

	// 如果这道题标记为动态题 (IsDynamic=true)
	if chal.IsDynamic {
		var container model.Container
		
		// 去查 containers 表：找当前用户、当前题目、且状态为 running 的容器
		// 使用 Last() 获取最新创建的一个，防止有旧的历史残留
		result := db.DB.Where("user_id = ? AND challenge_id = ? AND status = ?", userID, challengeID, "running").
			Last(&container)
		
		if result.Error != nil {
			// 没查到 -> 说明用户没启动靶机，或者靶机已过期销毁
			// 此时无法验证 Flag，直接报错提示用户
			return false, errors.New("请先启动靶机环境")
		}

		// 查到了 -> 用容器里的动态 Flag 覆盖掉静态 Flag
		correctFlag = container.Flag
	}
	
	// 7. 比对 Flag
	// 注意：这里比较的是 correctFlag (可能是静态的，也可能是动态获取的)
	if strings.TrimSpace(submittedFlag) != correctFlag {
		return false, nil // Flag 错误
	}

	// 8. 开启事务：写入记录 + 加分
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

		// B. 个人加分
		if err := tx.Model(&model.User{}).Where("id = ?", userID).
			UpdateColumn("score", gorm.Expr("score + ?", chal.Score)).Error; err != nil {
			return err
		}

		// C. 题目解题数 +1
		if err := tx.Model(&model.Challenge{}).Where("id = ?", challengeID).
			UpdateColumn("solved_count", gorm.Expr("solved_count + 1")).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return true, err
	}

	return true, nil
}
