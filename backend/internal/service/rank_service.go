package service

import (
	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
)

// GetCompetitionRank 获取指定比赛的排行榜
func GetCompetitionRank(compID uint) ([]model.RankItem, error) {
	var rankList []model.RankItem

	// 核心 SQL 逻辑：
	// 1. 联表 (JOIN users)：为了拿到用户名
	// 2. 筛选 (WHERE)：只看这场比赛 (competition_id)
	// 3. 分组 (GROUP BY)：按用户归类
	// 4. 排序 (ORDER BY)：分数高的在前，分数一样的时间早的在前(还没做last_solve_time，先只排分数)
	err := db.DB.Table("solves").
		Select("solves.user_id, users.username, SUM(solves.score) as score").
		Joins("left join users on users.id = solves.user_id").
		Where("solves.competition_id = ?", compID).
		Group("solves.user_id, users.username").
		Order("score desc").
		Scan(&rankList).Error

	if err != nil {
		return nil, err
	}

	// 填入名次 (1, 2, 3...)
	for i := range rankList {
		rankList[i].Rank = i + 1
	}

	return rankList, nil
}

// GetGlobalRank 获取全局排行榜 (直接查 Users 表)
func GetGlobalRank() ([]model.RankItem, error) {
	var rankList []model.RankItem

	// 直接查 User 表的 Score 字段
	err := db.DB.Model(&model.User{}).
		Select("id as user_id, username, score").
		Where("role != ?", "admin"). // 通常管理员不参与排名
		Order("score desc").
		Limit(100). // 只取前100名
		Scan(&rankList).Error

	if err != nil {
		return nil, err
	}

	for i := range rankList {
		rankList[i].Rank = i + 1
	}

	return rankList, nil
}

// GetTeamRank 获取团队排行榜
func GetTeamRank(compID uint) ([]model.TeamRankItem, error) {
	var rankList []model.TeamRankItem

	// 核心 SQL 逻辑：
	// Solves (解题记录)
	//   INNER JOIN Users (为了拿到 team_id)
	//   INNER JOIN Teams (为了拿到 team_name)
	// WHERE competition_id = ?
	// GROUP BY team_id
	// ORDER BY total_score DESC
	
	err := db.DB.Table("solves").
		Select("teams.id as team_id, teams.name as team_name, teams.token as team_token, COUNT(DISTINCT solves.user_id) as member_count, SUM(solves.score) as score").
		Joins("INNER JOIN users ON users.id = solves.user_id").
		Joins("INNER JOIN teams ON teams.id = users.team_id").
		Where("solves.competition_id = ?", compID).
		Group("teams.id, teams.name, teams.token").
		Order("score desc").
		Scan(&rankList).Error

	if err != nil {
		return nil, err
	}

	// 计算排名
	for i := range rankList {
		rankList[i].Rank = i + 1
	}

	return rankList, nil
}