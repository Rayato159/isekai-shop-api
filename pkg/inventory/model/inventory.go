package model

import (
	_itemShopModel "github.com/Rayato159/isekai-shop-api/pkg/itemShop/model"
)

type (
	Inventory struct {
		Item     *_itemShopModel.Item `json:"item"`
		Quantity uint                 `json:"quantity"`
	}

	ItemQuantityCounting struct {
		ItemID   uint64
		Quantity uint
	}
)
