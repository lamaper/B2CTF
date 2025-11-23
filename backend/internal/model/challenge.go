package model

import "time"

type Challenge struct {
    ID          uint      `gorm:"primaryKey"`
    Title       string    `gorm:"size:64;not null"`
    Category    string    `gorm:"size:32;not null"`
    Description string    `gorm:"type:text"`
    Score       int       `gorm:"not null"`
    FlagHash    string    `gorm:"size:255;not null"`
    CreatedAt   time.Time
}
