package repository

import (
	_itemEntity "github.com/Rayato159/isekai-shop-api/modules/item/entity"
)

type ItemRepository interface {
	FindItems(itemFilterDto *_itemEntity.ItemFilterDto) ([]*_itemEntity.Item, error)
	FindItemByID(itemID uint64) (*_itemEntity.Item, error)
	CountItems(itemFilterDto *_itemEntity.ItemFilterDto) (int64, error)
	InsertItem(itemEntity *_itemEntity.Item) (*_itemEntity.Item, error)
	UpdateItem(itemID uint64, updateItemDto *_itemEntity.UpdateItemDto) (uint64, error)
	ArchiveItem(itemID uint64) error // Soft delete
}
