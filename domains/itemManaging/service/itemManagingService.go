package service

import (
	_itemManaging "github.com/Rayato159/isekai-shop-api/domains/itemManaging/model"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/domains/itemShop/model"
)

type ItemManagingService interface {
	ItemCreating(itemCreatingReq *_itemManaging.ItemCreatingReq) (*_itemShopModel.Item, error)
	ItemEditing(itemID uint64, editItemReq *_itemManaging.ItemEditingReq) (*_itemShopModel.Item, error)
	ItemArchiving(itemID uint64) error
}
