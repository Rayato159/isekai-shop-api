package repository

import (
	_itemEntity "github.com/Rayato159/isekai-shop-api/modules/item/entity"
)

type ItemRepository interface {
	FindItems(itemFilterDto *_itemEntity.ItemFilterDto) ([]*_itemEntity.Item, error)
	CountItems(itemFilterDto *_itemEntity.ItemFilterDto) (int64, error)
	InsertItem(item *_itemEntity.Item) (*_itemEntity.Item, error)
	UpdateItem(itemID uint64, item *_itemEntity.UpdateItemDto) (*_itemEntity.Item, error)
	ArchiveItem(itemID uint64) error // Soft delete
}
