package service

import (
	_itemManaging "github.com/Rayato159/isekai-shop-api/pkg/itemManaging/model"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/pkg/itemShop/model"
)

type ItemManagingService interface {
	Creating(itemCreatingReq *_itemManaging.ItemCreatingReq) (*_itemShopModel.Item, error)
	Editing(itemID uint64, editItemReq *_itemManaging.ItemEditingReq) (*_itemShopModel.Item, error)
	Archiving(itemID uint64) error
}
