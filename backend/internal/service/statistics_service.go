// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-02-08
// 最后修改: 2026-02-08
// 描述: 比赛统计相关的业务逻辑服务（管理员专用）
// ----------------------------------------------------------------------------
package service

import (
	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
)

// CompetitionStatistics 比赛统计信息
type CompetitionStatistics struct {
	CompetitionID     uint    `json:"competition_id"`
	CompetitionName   string  `json:"competition_name"`
	TotalChallenges   int     `json:"total_challenges"`
	TotalParticipants int     `json:"total_participants"`
	TotalSolves       int     `json:"total_solves"`
	AverageSolveRate  float64 `json:"average_solve_rate"`
}

// GetCompetitionStatistics 获取比赛统计信息
func GetCompetitionStatistics(compID uint) (*CompetitionStatistics, error) {
	var comp model.Competition
	if err := db.DB.First(&comp, compID).Error; err != nil {
		return nil, err
	}

	// 获取比赛的题目数
	var totalChallenges int64
	db.DB.Model(&model.Challenge{}).Where("competition_id = ?", compID).Count(&totalChallenges)

	// 获取参赛人数（参与过解题的唯一用户数）
	var totalParticipants int64
	db.DB.Table("solves").
		Select("COUNT(DISTINCT user_id)").
		Where("competition_id = ?", compID).
		Row().
		Scan(&totalParticipants)

	// 获取总解题数
	var totalSolves int64
	db.DB.Model(&model.Solve{}).Where("competition_id = ?", compID).Count(&totalSolves)

	// 计算平均解题率
	var avgSolveRate float64
	if totalChallenges > 0 && totalParticipants > 0 {
		avgSolveRate = float64(totalSolves) / (float64(totalChallenges) * float64(totalParticipants)) * 100
	}

	stats := &CompetitionStatistics{
		CompetitionID:     compID,
		CompetitionName:   comp.Title,
		TotalChallenges:   int(totalChallenges),
		TotalParticipants: int(totalParticipants),
		TotalSolves:       int(totalSolves),
		AverageSolveRate:  avgSolveRate,
	}

	return stats, nil
}
