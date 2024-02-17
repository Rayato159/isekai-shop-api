package entity

import (
	_orderModel "github.com/Rayato159/isekai-shop-api/modules/order/model"

	"time"
)

type (
	Order struct {
		ID              uint64    `gorm:"primaryKey;autoIncrement;"`
		PlayerID        string    `gorm:"type:varchar(64);not null;"`
		ItemID          uint64    `gorm:"type:bigint;not null;"`
		ItemName        string    `gorm:"type:varchar(64);not null;"`
		ItemDescription string    `gorm:"type:varchar(128);not null;"`
		ItemPrice       uint      `gorm:"not null;"`
		ItemPicture     string    `gorm:"type:varchar(128);not null;"`
		Quantity        uint      `gorm:"not null;"`
		TotalPrice      int64     `gorm:"not null;"`
		CreatedAt       time.Time `gorm:"not null;autoCreateTime;"`
	}
)

func (o *Order) ToOrderModel() *_orderModel.Order {
	return &_orderModel.Order{
		ID:              o.ID,
		PlayerID:        o.PlayerID,
		ItemID:          o.ItemID,
		ItemName:        o.ItemName,
		ItemDescription: o.ItemDescription,
		ItemPicture:     o.ItemPicture,
		ItemPrice:       o.ItemPrice,
		Quantity:        o.Quantity,
		TotalPrice:      o.TotalPrice,
		CreatedAt:       o.CreatedAt,
	}
}
