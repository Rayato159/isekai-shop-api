package service

import (
	_itemModel "github.com/Rayato159/isekai-shop-api/modules/item/model"
)

type AdminService interface {
	ItemCreating(createItemReq *_itemModel.ItemCreatingReq) (*_itemModel.Item, error)
	ItemEditing(itemID uint64, editItemReq *_itemModel.ItemEditingReq) (*_itemModel.Item, error)
	ItemArchiving(itemID uint64) error
}
