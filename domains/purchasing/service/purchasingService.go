package service

import (
	_playerBalancingModel "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/model"
	_purchasingModel "github.com/Rayato159/isekai-shop-api/domains/purchasing/model"
)

type PurchasingService interface {
	ItemBuying(itemBuyingReq *_purchasingModel.ItemBuyingReq) (*_playerBalancingModel.PlayerBalancing, error)
	ItemSelling(itemSellingReq *_purchasingModel.ItemSellingReq) (*_playerBalancingModel.PlayerBalancing, error)
}
