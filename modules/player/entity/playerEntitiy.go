package entity

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Id        string `gorm:"primaryKey;"`
	Email     string `gorm:"unique;not null;"`
	Name      string `gorm:"not null;"`
	Picture   string
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
