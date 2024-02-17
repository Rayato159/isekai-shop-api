package entity

import "time"

type (
	Inventory struct {
		ID        uint64    `gorm:"primaryKey;autoIncrement;"`
		PlayerID  string    `gorm:"type:varchar(64);not null;"`
		ItemID    uint64    `gorm:"type:bigint;not null;"`
		CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	}
)
