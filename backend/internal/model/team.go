package model

import "gorm.io/gorm"

type Team Struct{
	gorm.Model
	Name        string `gorm:"uniqueIndex;size:64;not null" json:"name"`
	Description string `gorm:"size:255" json:"description"`
	Token       string `gorm:"uniqueIndex;size:32;not null" json:"token"` 
	CaptainID   uint   `json:"captain_id"`
	Members     []User `gorm:"foreignKey:TeamID" json:"members"`
}