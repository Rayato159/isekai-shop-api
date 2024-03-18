package service

import (
	_itemShopModel "github.com/Rayato159/isekai-shop-api/domains/itemShop/model"
)

type ItemService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
}
