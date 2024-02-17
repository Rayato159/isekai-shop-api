package entity

import (
	"time"

	_inventoryEntity "github.com/Rayato159/isekai-shop-api/modules/inventory/entity"
	_paymentEntity "github.com/Rayato159/isekai-shop-api/modules/payment/entity"
	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
)

type (
	Player struct {
		ID          string                       `gorm:"primaryKey;type:varchar(64);"`
		Email       string                       `gorm:"type:varchar(128);unique;not null;"`
		Name        string                       `gorm:"type:varchar(128);not null;"`
		Avatar      string                       `gorm:"type:varchar(256);not null;default:'';"`
		Username    *string                      `gorm:"type:varchar(64);unique;"`
		Inventories []_inventoryEntity.Inventory `gorm:"foreignKey:PlayerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Payments    []_paymentEntity.Payment     `gorm:"foreignKey:PlayerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		CreatedAt   time.Time                    `gorm:"not null;autoCreateTime;"`
		UpdatedAt   time.Time                    `gorm:"not null;autoUpdateTime;"`
	}

	UpdatePlayerDto struct {
		Username string
	}
)

func (p *Player) ToPlayerModel() *_playerModel.Player {
	return &_playerModel.Player{
		ID:        p.ID,
		Email:     p.Email,
		Name:      p.Name,
		Avatar:    p.Avatar,
		Username:  p.Username,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
