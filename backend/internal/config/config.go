// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-01-21
// 描述: 配置加载和管理
// ----------------------------------------------------------------------------
package config

import (
    "log"

    "github.com/spf13/viper"
)

type ServerConfig struct {
    Addr string `mapstructure:"addr"`
}

type DatabaseConfig struct {
    Driver string `mapstructure:"driver"`
    DSN    string `mapstructure:"dsn"`
}

type JWTConfig struct {
    Secret      string `mapstructure:"secret"`
    Issuer      string `mapstructure:"issuer"`
    ExpireHours int    `mapstructure:"expire_hours"`
}

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    JWT      JWTConfig      `mapstructure:"jwt"`
}

var C Config

func Init(path string) {
    viper.SetConfigFile(path)
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("read config failed: %v", err)
    }
    if err := viper.Unmarshal(&C); err != nil {
        log.Fatalf("unmarshal config failed: %v", err)
    }
}
