package model

import (
	_itemGettingModel "github.com/Rayato159/isekai-shop-api/domains/itemGetting/model"
)

type (
	Inventory struct {
		Item     *_itemGettingModel.Item `json:"item"`
		Quantity uint                    `json:"quantity"`
	}

	ItemQuantityCounting struct {
		ItemID   uint64
		Quantity uint
	}
)
