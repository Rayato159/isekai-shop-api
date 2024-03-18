package entity

import (
	_playerBalancingModel "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/model"

	"time"
)

type (
	PlayerBalancing struct {
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

func (p *PlayerBalancing) ToPlayerBalancingModel() *_playerBalancingModel.PlayerBalancing {
	return &_playerBalancingModel.PlayerBalancing{
		ID:        p.ID,
		PlayerID:  p.PlayerID,
		Amount:    p.Amount,
		CreatedAt: p.CreatedAt,
	}
}

func (p *PlayerBalanceShowingDto) ToPlayerBalanceModel() *_playerBalancingModel.PlayerBalanceShowing {
	return &_playerBalancingModel.PlayerBalanceShowing{
		PlayerID: p.PlayerID,
		Balance:  p.Balance,
	}
}
