package entities

import (
	"time"
)

type PurchaseHistory struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement;"`
	PlayerID        string    `gorm:"type:varchar(64);not null;"`
	ItemID          uint64    `gorm:"type:bigint;not null;"`
	ItemName        string    `gorm:"type:varchar(64);not null;"`
	ItemDescription string    `gorm:"type:varchar(128);not null;"`
	ItemPrice       uint      `gorm:"not null;"`
	ItemPicture     string    `gorm:"type:varchar(128);not null;"`
	Quantity        uint      `gorm:"not null;"`
	IsBuying        bool      `gorm:"type:boolean;not null;"`
	CreatedAt       time.Time `gorm:"not null;autoCreateTime;"`
}
