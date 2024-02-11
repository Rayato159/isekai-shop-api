package entity

import (
	"time"

	"gorm.io/gorm"
)

type Credential struct {
	gorm.Model
	Id           string    `gorm:"primaryKey;"`
	PlayerId     string    `gorm:"not null;foreignKey:PlayerRefer"` // Foreign Key
	RefreshToken string    `gorm:"unique;not null;"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt    time.Time `gorm:"not null;autoUpdateTime;"`
}
