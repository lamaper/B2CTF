package model

import "time"

type Solve struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      uint      `gorm:"index;not null"`
    ChallengeID uint      `gorm:"index;not null"`
    CreatedAt   time.Time
}
