package service

import (
	_itemModel "github.com/Rayato159/isekai-shop-api/modules/item/model"
)

type AdminService interface {
	CreateItem(createItemReq *_itemModel.CreateItemReq) (*_itemModel.Item, error)
	EditItem(itemID uint64, editItemReq *_itemModel.EditItemReq) (*_itemModel.Item, error)
	ArchiveItem(itemID uint64) error
}
