// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-22
// 最后修改: 2026-01-22
// 描述: 团队数据模型
// ----------------------------------------------------------------------------
package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;size:64;not null" json:"name"`
	Description string `gorm:"size:255" json:"description"`
	Token       string `gorm:"uniqueIndex;size:32;not null" json:"token"`
	CaptainID   uint   `json:"captain_id"`
	Members     []User `gorm:"foreignKey:TeamID" json:"members"`
}
