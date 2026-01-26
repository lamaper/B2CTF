// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-17
// 最后修改: 2026-01-17
// 描述: JWT工具函数
// ----------------------------------------------------------------------------
package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// jwtSecret JWT签名密钥
// 注意：此密钥仅用于测试环境，生产环境必须使用安全的随机生成密钥
var jwtSecret = []byte("my_super_secret_key_for_ctf_2026")

// GenerateToken 生成JWT令牌
//
// 参数:
//
//	userID: 用户ID
//	role: 用户角色 (user/admin)
//
// 返回值:
//
//	string: 生成的JWT令牌
//	error: 生成失败时返回的错误信息
//
// 令牌有效期: 24小时
//
// 令牌包含的声明(claims):
//   - user_id: 用户ID
//   - role: 用户角色
//   - exp: 过期时间戳
//   - iat: 签发时间戳
func GenerateToken(userID uint, role string) (string, error) {
	// 创建JWT声明
	claims := jwt.MapClaims{
		"user_id": userID,                                // 用户ID
		"role":    role,                                  // 用户角色
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 过期时间：24小时后
		"iat":     time.Now().Unix(),                     // 签发时间
	}

	// 创建令牌实例
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名令牌并获取完整的编码字符串
	return token.SignedString(jwtSecret)
}

// ParseToken 解析 JWT Token
// 输入：tokenString
// 输出：Claims结构体, error
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	// 1. 解析 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否是我们预期的 HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 2. 验证 token 是否有效，并提取 claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
