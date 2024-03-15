package entity

import (
	_paymentModel "github.com/Rayato159/isekai-shop-api/domains/payment/model"

	"time"
)

type (
	Payment struct {
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

func (p *Payment) ToPaymentModel() *_paymentModel.Payment {
	return &_paymentModel.Payment{
		ID:        p.ID,
		PlayerID:  p.PlayerID,
		Amount:    p.Amount,
		CreatedAt: p.CreatedAt,
	}
}

func (p *PlayerBalanceDto) ToPlayerBalanceModel() *_paymentModel.PlayerBalance {
	return &_paymentModel.PlayerBalance{
		PlayerID: p.PlayerID,
		Balance:  p.Balance,
	}
}
