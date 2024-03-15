package repository

import (
	_itemEntity "github.com/Rayato159/isekai-shop-api/domains/item/entity"
)

type ItemRepository interface {
	ItemListing(itemFilterDto *_itemEntity.ItemFilterDto) ([]*_itemEntity.Item, error)
	FindItemByID(itemID uint64) (*_itemEntity.Item, error)
	FindItemByIDs(itemIDs []uint64) ([]*_itemEntity.Item, error)
	ItemCounting(itemFilterDto *_itemEntity.ItemFilterDto) (int64, error)
	ItemCreating(itemEntity *_itemEntity.Item) (*_itemEntity.Item, error)
	ItemEditing(itemID uint64, updateItemDto *_itemEntity.ItemEditingDto) (uint64, error)
	ItemArchiving(itemID uint64) error // Soft delete
}
