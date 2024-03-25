package entities

import (
	"time"
)

type Player struct {
	ID          string      `gorm:"primaryKey;type:varchar(64);"`
	Email       string      `gorm:"type:varchar(128);unique;not null;"`
	Name        string      `gorm:"type:varchar(128);not null;"`
	Avatar      string      `gorm:"type:varchar(256);not null;default:'';"`
	Inventories []Inventory `gorm:"foreignKey:PlayerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time   `gorm:"not null;autoCreateTime;"`
	UpdatedAt   time.Time   `gorm:"not null;autoUpdateTime;"`
}
