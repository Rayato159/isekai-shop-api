package model

import (
	_itemModel "github.com/Rayato159/isekai-shop-api/modules/item/model"
)

type (
	Inventory struct {
		Item     *_itemModel.Item `json:"item"`
		Quantity uint             `json:"quantity"`
	}

	ItemQuantityCounter struct {
		ItemID   uint64
		Quantity uint
	}
)
