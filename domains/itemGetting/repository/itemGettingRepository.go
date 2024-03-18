package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type ItemGettingRepository interface {
	Listing(itemFilterDto *entities.ItemFilterDto) ([]*entities.Item, error)
	FindByID(itemID uint64) (*entities.Item, error)
	FindByIDList(itemIDs []uint64) ([]*entities.Item, error)
	Counting(itemFilterDto *entities.ItemFilterDto) (int64, error)
}
