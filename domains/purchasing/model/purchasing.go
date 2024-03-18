package model

type (
	ItemBuyingReq struct {
		PlayerID string
		ItemID   uint64 `json:"itemID" validate:"required,gt=0"`
		Quantity uint   `json:"quantity" validate:"required,gt=0"`
	}

	ItemSellingReq struct {
		PlayerID string
		ItemID   uint64 `json:"itemID" validate:"required,gt=0"`
		Quantity uint   `json:"quantity" validate:"required,gt=0"`
	}
)
