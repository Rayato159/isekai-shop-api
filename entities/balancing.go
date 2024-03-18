package entity

import (
	_balancingModel "github.com/Rayato159/isekai-shop-api/domains/balancing/model"

	"time"
)

type (
	Balancing struct {
		ID        uint64    `gorm:"primaryKey;autoIncrement;"`
		PlayerID  string    `gorm:"type:varchar(64);not null;"`
		Amount    int64     `gorm:"not null;"`
		CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	}

	PlayerBalanceDto struct {
		PlayerID string `json:"playerID"`
		Balance  int64  `json:"balance"`
	}
)

func (p *Balancing) ToBalancingModel() *_balancingModel.Balancing {
	return &_balancingModel.Balancing{
		ID:        p.ID,
		PlayerID:  p.PlayerID,
		Amount:    p.Amount,
		CreatedAt: p.CreatedAt,
	}
}

func (p *PlayerBalanceDto) ToPlayerBalanceModel() *_balancingModel.PlayerBalance {
	return &_balancingModel.PlayerBalance{
		PlayerID: p.PlayerID,
		Balance:  p.Balance,
	}
}
