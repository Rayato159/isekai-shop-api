package entity

import (
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/domains/playerCoin/model"

	"time"
)

type (
	PlayerCoin struct {
		ID        uint64    `gorm:"primaryKey;autoIncrement;"`
		PlayerID  string    `gorm:"type:varchar(64);not null;"`
		Amount    int64     `gorm:"not null;"`
		CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	}

	PlayerCoinShowingDto struct {
		PlayerID string `json:"playerID"`
		Coin     int64  `json:"coin"`
	}
)

func (p *PlayerCoin) ToPlayerCoinModel() *_playerCoinModel.PlayerCoin {
	return &_playerCoinModel.PlayerCoin{
		ID:        p.ID,
		PlayerID:  p.PlayerID,
		Amount:    p.Amount,
		CreatedAt: p.CreatedAt,
	}
}

func (p *PlayerCoinShowingDto) ToPlayerCoinModel() *_playerCoinModel.PlayerCoinShowing {
	return &_playerCoinModel.PlayerCoinShowing{
		PlayerID: p.PlayerID,
		Balance:  p.Coin,
	}
}
