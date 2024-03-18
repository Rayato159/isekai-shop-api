package model

import "time"

type (
	PlayerBalancing struct {
		ID        uint64    `json:"id"`
		PlayerID  string    `json:"playerID"`
		Amount    int64     `json:"amount"`
		CreatedAt time.Time `json:"createdAt"`
	}

	TopUpReq struct {
		PlayerID string
		Amount   int64 `json:"amount" validate:"required,gt=0"`
	}

	PlayerBalanceShowing struct {
		PlayerID string `json:"playerID"`
		Balance  int64  `json:"balance"`
	}
)
