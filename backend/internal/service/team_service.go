package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"

	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"

	"gorm.io/gorm"
)

// CreateTeam 创建队伍
func CreateTeam(userID uint, name, desc string) (string, error) {
	// 1. 检查用户是否已有队伍
	var user model.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return "", err
	}
	if user.TeamID != 0 {
		return "", errors.New("你已经在一个队伍中了，请先退出")
	}

	// 2. 生成随机邀请码 (取 UUID 前8位)
	token := "T-" + strings.ToUpper(strings.Split(uuid.New().String(), "-")[0])

	// 3. 开启事务
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// A. 创建队伍记录
		team := model.Team{
			Name:        name,
			Description: desc,
			Token:       token,
			CaptainID:   userID,
		}
		if err := tx.Create(&team).Error; err != nil {
			// 如果队名重复会报错
			return err
		}

		// B. 更新队长的 TeamID
		if err := tx.Model(&model.User{}).Where("id = ?", userID).Update("team_id", team.ID).Error; err != nil {
			return err
		}
		return nil
	})

	return token, err
}

// JoinTeam 加入队伍
func JoinTeam(userID uint, token string) error {
	// 1. 检查用户状态
	var user model.User
	db.DB.First(&user, userID)
	if user.TeamID != 0 {
		return errors.New("你已在队伍中")
	}

	// 2. 检查 Token 是否有效
	var team model.Team
	if err := db.DB.Where("token = ?", token).First(&team).Error; err != nil {
		return errors.New("邀请码无效或队伍不存在")
	}

	// 3. 更新用户
	return db.DB.Model(&user).Update("team_id", team.ID).Error
}

// LeaveTeam 退出队伍 (如果是队长则不能退，必须解散，这里先简化逻辑)
func LeaveTeam(userID uint) error {
	var user model.User
	db.DB.First(&user, userID)

	if user.TeamID == 0 {
		return errors.New("你当前没有队伍")
	}

	// 检查是否是队长
	var team model.Team
	db.DB.First(&team, user.TeamID)
	if team.CaptainID == userID {
		return errors.New("队长无法直接退出，请选择解散队伍")
	}

	return db.DB.Model(&user).Update("team_id", 0).Error
}

// GetMyTeamInfo 获取我的队伍详情（包含成员列表和总分）
func GetMyTeamInfo(userID uint) (*model.Team, error) {
	var user model.User
	db.DB.First(&user, userID)
	if user.TeamID == 0 {
		return nil, errors.New("NO_TEAM") // 特殊错误标记
	}

	var team model.Team
	// Preload("Members") 会自动把队伍里的 User 全查出来
	if err := db.DB.Preload("Members").First(&team, user.TeamID).Error; err != nil {
		return nil, err
	}

	return &team, nil
}

func GetTeamByID(teamID uint) (*model.Team, error) {
	var team model.Team
	if err := db.DB.First(&team, teamID).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

// TODO: GetTeamRank 获取团队赛排行榜

func DeleteTeam(teamID uint) error {
	return db.DB.Delete(&model.Team{}, teamID).Error
}

func SetTeamInfo(teamID uint, name, desc string) error {
	var team model.Team
	if err := db.DB.First(&team, teamID).Error; err != nil {
		return err
	}

	team.Name = name
	team.Description = desc

	return db.DB.Save(&team).Error
}
