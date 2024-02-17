package model

import "time"

type (
	Order struct {
		ID              uint64    `json:"id"`
		PlayerID        string    `json:"playerID"`
		ItemID          uint64    `json:"itemID"`
		ItemName        string    `json:"itemName"`
		ItemDescription string    `json:"itemDescription"`
		ItemPrice       int64     `json:"itemPrice"`
		Quantity        int       `json:"quantity"`
		TotalPrice      int64     `json:"totalPrice"`
		CreatedAt       time.Time `json:"createdAt"`
	}
)
