package model

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey"`
    Username     string    `gorm:"uniqueIndex;size:32;not null"`
    PasswordHash string    `gorm:"size:255;not null"`
    Role         string    `gorm:"size:16;default:user;not null"` // user/admin
    CreatedAt    time.Time
}
