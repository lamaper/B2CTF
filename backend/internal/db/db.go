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
    if err := DB.AutoMigrate(&model.User{}, &model.Challenge{}, &model.Solve{}); err != nil {
        log.Fatalf("auto migrate failed: %v", err)
    }
}
