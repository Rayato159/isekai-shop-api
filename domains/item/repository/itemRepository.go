package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type ItemRepository interface {
	ItemListing(itemFilterDto *entities.ItemFilterDto) ([]*entities.Item, error)
	FindItemByID(itemID uint64) (*entities.Item, error)
	FindItemByIDs(itemIDs []uint64) ([]*entities.Item, error)
	ItemCounting(itemFilterDto *entities.ItemFilterDto) (int64, error)
}
