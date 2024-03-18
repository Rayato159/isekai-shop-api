package service

import (
	_itemGettingModel "github.com/Rayato159/isekai-shop-api/domains/itemGetting/model"
	_itemManaging "github.com/Rayato159/isekai-shop-api/domains/itemManaging/model"
)

type ItemManagingService interface {
	ItemCreating(itemCreatingReq *_itemManaging.ItemCreatingReq) (*_itemGettingModel.Item, error)
	ItemEditing(itemID uint64, editItemReq *_itemManaging.ItemEditingReq) (*_itemGettingModel.Item, error)
	ItemArchiving(itemID uint64) error
}
