package service

import (
	_itemModel "github.com/Rayato159/isekai-shop-api/domains/item/model"
	_itemManaging "github.com/Rayato159/isekai-shop-api/domains/itemManaging/model"
)

type ItemManagingService interface {
	ItemCreating(itemCreatingReq *_itemManaging.ItemCreatingReq) (*_itemModel.Item, error)
	ItemEditing(itemID uint64, editItemReq *_itemManaging.ItemEditingReq) (*_itemModel.Item, error)
	ItemArchiving(itemID uint64) error
}
