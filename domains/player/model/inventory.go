package model

import (
	_itemModel "github.com/Rayato159/isekai-shop-api/domains/item/model"
)

type (
	Inventory struct {
		Item     *_itemModel.Item `json:"item"`
		Quantity uint             `json:"quantity"`
	}

	ItemQuantityCounting struct {
		ItemID   uint64
		Quantity uint
	}
)
