// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-01-21
// 描述: 数据库初始化和连接管理
// ----------------------------------------------------------------------------
package db

import (
	"log"

	"B2CTF/backend/internal/config"
	"B2CTF/backend/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	cfg := config.C.Database

	var dial gorm.Dialector
	switch cfg.Driver {
	case "mysql":
		dial = mysql.Open(cfg.DSN)
	default:
		log.Fatalf("unsupported db driver: %s", cfg.Driver)
	}

	var err error
	DB, err = gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Fatalf("open db failed: %v", err)
	}

	// 自动迁移
	// ⚠️ 加上 &model.Container{}
	err = DB.AutoMigrate(
		&model.Team{},
		&model.Competition{},
		&model.User{},
		&model.Challenge{},
		&model.Solve{},
		&model.Container{}, // <--- 新增这一行
	)

	if err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}
}