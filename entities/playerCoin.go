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

	PlayerBalanceShowingDto struct {
		PlayerID string `json:"playerID"`
		Balance  int64  `json:"balance"`
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

func (p *PlayerBalanceShowingDto) ToPlayerBalanceModel() *_playerCoinModel.PlayerBalanceShowing {
	return &_playerCoinModel.PlayerBalanceShowing{
		PlayerID: p.PlayerID,
		Balance:  p.Balance,
	}
}
