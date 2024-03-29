package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type InventoryRepository interface {
	Filling(playerID string, itemID uint64, qty int) ([]*entities.Inventory, error)
	ReverseFilling(inventoryEntities []*entities.Inventory) error
	Listing(playerID string) ([]*entities.Inventory, error)
	Removing(playerID string, itemID uint64, limit int) error
	ReverseRemoving(playerID string, itemID uint64, limit int) error
	PlayerItemCounting(playerID string, itemID uint64) int64
}
