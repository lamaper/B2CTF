// ----------------------------------------------------------------------------
// Copyright (c) 2026 West
// 创建日期: 2026-02-08
// 最后修改: 2026-02-08
// 描述: 用户管理相关的业务逻辑服务（管理员专用）
// ----------------------------------------------------------------------------
package service

import (
	"errors"
	"log"

	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
)

// UserInfo 用户信息返回结构（不包含敏感信息）
type UserInfo struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Score     int    `json:"score"`
	Avatar    string `json:"avatar"`
	TeamID    uint   `json:"team_id"`
	CreatedAt string `json:"created_at"`
}

// ListAllUsers 获取所有用户信息（仅管理员可用）
func ListAllUsers(adminID uint) ([]UserInfo, error) {
	var users []model.User
	err := db.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	// 转换为 UserInfo，不包含密码哈希
	var userInfos []UserInfo
	for _, u := range users {
		userInfos = append(userInfos, UserInfo{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			Role:      u.Role,
			Score:     u.Score,
			Avatar:    u.Avatar,
			TeamID:    u.TeamID,
			CreatedAt: u.CreatedAt.String(),
		})
	}

	log.Printf("[审计] 管理员 %d 查看了用户列表", adminID)
	return userInfos, nil
}

// SetUserRole 设置用户角色
func SetUserRole(userID uint, newRole string, adminID uint) error {
	var user model.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return errors.New("用户不存在")
	}

	oldRole := user.Role
	err := db.DB.Model(&user).Update("role", newRole).Error
	if err == nil {
		log.Printf("[审计] 管理员 %d 将用户 %d (%s) 的角色从 %s 修改为 %s",
			adminID, userID, user.Username, oldRole, newRole)
	}
	return err
}

// DeleteUser 删除用户
func DeleteUser(userID uint, adminID uint) error {
	var user model.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return errors.New("用户不存在")
	}

	// 如果用户所在团队只有这一个人，则解散团队
	if user.TeamID != 0 {
		var count int64
		db.DB.Model(&model.User{}).Where("team_id = ?", user.TeamID).Count(&count)
		if count == 1 {
			// 删除团队
			db.DB.Delete(&model.Team{}, user.TeamID)
		} else {
			// 将该用户移出团队
			db.DB.Model(&user).Update("team_id", 0)
		}
	}

	// 删除用户的解题记录
	db.DB.Delete(&model.Solve{}, "user_id = ?", userID)

	// 删除用户
	err := db.DB.Delete(&user).Error
	if err == nil {
		log.Printf("[审计] 管理员 %d 删除了用户 %d (%s)", adminID, userID, user.Username)
	}
	return err
}
