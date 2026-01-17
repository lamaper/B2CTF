// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-17
// 最后修改: 2026-01-17
// 描述: 用户数据模型
// ----------------------------------------------------------------------------
package model

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"uniqueIndex;size:32;not null"`
	PasswordHash string    `gorm:"size:255;not null"`
	Role         string    `gorm:"size:16;default:user;not null"` // user/admin
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	Email        string    `gorm:"uniqueIndex;size:32;not null"`
}
