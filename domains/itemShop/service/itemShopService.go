package service

import (
	_itemShopModel "github.com/Rayato159/isekai-shop-api/domains/itemShop/model"
	_playerBalancingModel "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
	Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerBalancingModel.PlayerBalancing, error)
	Selling(sellingReq *_itemShopModel.SellingReq) (*_playerBalancingModel.PlayerBalancing, error)
}
