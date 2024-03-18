package service

import (
	_balancingModel "github.com/Rayato159/isekai-shop-api/domains/balancing/model"
	_purchasingModel "github.com/Rayato159/isekai-shop-api/domains/purchasing/model"
)

type PurchasingService interface {
	ItemBuying(itemBuyingReq *_purchasingModel.ItemBuyingReq) (*_balancingModel.Balancing, error)
	ItemSelling(itemSellingReq *_purchasingModel.ItemSellingReq) (*_balancingModel.Balancing, error)
}
