package service

import (
	_itemShopModel "github.com/Rayato159/isekai-shop-api/domains/itemShop/model"
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/domains/playerCoin/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
	Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error)
	Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error)
}
