// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-22
// 最后修改: 2026-01-22
// 描述: 比赛积分榜模型
// ----------------------------------------------------------------------------
package model

type RankItem struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"` // 如果你后面做了头像功能
	Score    int    `json:"score"`
	Rank     int    `json:"rank"`   // 名次
}

type TeamRankItem struct {
	TeamID      uint   `json:"team_id"`
	TeamName    string `json:"team_name"`
	TeamToken   string `json:"team_token"` // 可选，展示一下战队Token
	MemberCount int    `json:"member_count"` // 战队人数
	Score       int    `json:"score"`
	Rank        int    `json:"rank"`
}