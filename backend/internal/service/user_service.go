// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-17
// 最后修改: 2026-01-17
// 描述: 用户相关的业务逻辑服务
// ----------------------------------------------------------------------------
package service

import (
	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
	"B2CTF/backend/internal/pkg/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Register 处理用户注册逻辑
//
// 参数:
//
//	username: 用户注册的用户名
//	password: 用户注册的密码
//	email: 用户注册的邮箱地址
//
// 返回值:
//
//	error: 注册成功返回nil, 失败返回具体错误信息
//
// 错误类型:
//
//	"用户名已存在": 用户名已被注册
//	"密码加密失败": 密码哈希生成失败
//	其他: 数据库操作等系统错误
func Register(username string, password string, email string) error {
	// 检查用户名是否存在
	var user model.User
	result := db.DB.Where("username = ?", username).First(&user)

	if result.Error == nil {
		return errors.New("用户名已存在")
	}
	if result.Error != gorm.ErrRecordNotFound {
		return result.Error // 数据库查询出错
	}

	// 密码加密 - 使用bcrypt算法生成哈希值
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), 
	bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败")
	}

	// 创建新用户 - 默认角色为普通用户
	newUser := model.User{
		Username:     username,
		PasswordHash: string(hashedPwd),
		Email:        email,
		Role:         "user", // 默认为普通用户
	}

	// 插入数据库
	if err := db.DB.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

// Login 处理用户登录逻辑
//
// 参数:
//
//	username: 用户登录的用户名
//	password: 用户登录的密码
//
// 返回值:
//
//	string: 登录成功返回JWT令牌
//	error: 登录失败返回具体错误信息
//
// 错误类型:
//
//	"用户名或密码错误": 用户名不存在或密码错误
//	其他: JWT生成失败或数据库操作等系统错误
func Login(username string, password string) (string, error) {
	var user model.User
	// 根据用户名查询用户信息
	result := db.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return "", errors.New("用户名或密码错误") 
	}

	// 比对密码哈希值 - 使用bcrypt算法验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	// 生成JWT令牌 - 包含用户ID和角色信息
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

// UpdateAvatar 更新用户头像
// userID: 用户ID
// newAvatarPath: 新头像的存储路径
func UpdateAvatar(userID uint, newAvatarPath string) error {
	var user model.User
	// 1. 先查出用户当前的信息
	if err := db.DB.First(&user, userID).Error; err != nil {
		return err
	}

	// 2. [优化] 清理旧头像
	// 如果用户原来有头像，且不是默认头像，且文件存在，就删掉它
	if user.Avatar != "" && !strings.Contains(user.Avatar, "default") {
		// user.Avatar 存的是 "/uploads/...", 实际路径是 "./uploads/..."
		// 这里的路径拼接需要根据你实际的存储策略微调
		// 简单的做法是去掉开头的 "/" 然后拼接 "."
		oldPath := "." + user.Avatar
		_ = os.Remove(oldPath) // 忽略错误，删不掉也不影响主流程
	}

	// 3. 更新数据库字段
	// UpdateColumn 仅更新指定字段，效率高且安全
	return db.DB.Model(&user).UpdateColumn("avatar", newAvatarPath).Error
}