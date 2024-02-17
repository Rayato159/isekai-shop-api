package model

import "time"

type (
	Order struct {
		ID              uint64    `json:"id"`
		PlayerID        string    `json:"playerID"`
		ItemID          uint64    `json:"itemID"`
		ItemName        string    `json:"itemName"`
		ItemDescription string    `json:"itemDescription"`
		ItemPicture     string    `json:"itemPicture"`
		ItemPrice       uint      `json:"itemPrice"`
		Quantity        uint      `json:"quantity"`
		TotalPrice      int64     `json:"totalPrice"`
		CreatedAt       time.Time `json:"createdAt"`
	}
)
