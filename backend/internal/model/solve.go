// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-01-21
// 描述: 解题记录数据模型
// ----------------------------------------------------------------------------
package model

import "gorm.io/gorm"

type Solve struct {
	gorm.Model

	// 核心关联：谁 (UserID) 做出了 哪道题 (ChallengeID)
	// gorm:"uniqueIndex:idx_user_challenge" 创建一个复合唯一索引
	// 数据库层面保证：同一个 UserID 和 ChallengeID 的组合只能出现一次
	UserID      uint `gorm:"uniqueIndex:idx_user_challenge;not null" json:"user_id"`
	ChallengeID uint `gorm:"uniqueIndex:idx_user_challenge;not null" json:"challenge_id"`

	// 冗余字段：这道题属于哪个比赛
	// 虽然可以通过 ChallengeID 查到，但这里存一份方便快速生成“比赛排行榜”
	CompetitionID uint `gorm:"not null" json:"competition_id"`

	// 快照分值：记录解题时获得了多少分
	// 防止未来管理员修改了题目分值，导致历史记录对不上
	Score int `gorm:"not null" json:"score"`

	// (可选) 关联对象：方便 GORM 预加载用户信息和题目信息
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	Challenge Challenge `gorm:"foreignKey:ChallengeID" json:"-"`
}
