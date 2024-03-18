package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/domains/entities"
)

type ItemRepository interface {
	ItemListing(itemFilterDto *entities.ItemFilterDto) ([]*entities.Item, error)
	FindItemByID(itemID uint64) (*entities.Item, error)
	FindItemByIDs(itemIDs []uint64) ([]*entities.Item, error)
	ItemCounting(itemFilterDto *entities.ItemFilterDto) (int64, error)
	ItemCreating(itemEntity *entities.Item) (*entities.Item, error)
	ItemEditing(itemID uint64, updateItemDto *entities.ItemEditingDto) (uint64, error)
	ItemArchiving(itemID uint64) error // Soft delete
}
