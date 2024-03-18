package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/domains/entities"
)

type InventoryRepository interface {
	InventoryFilling(inventoryEntities []*entities.Inventory) ([]*entities.Inventory, error)
	InventorySearching(playerID string) ([]*entities.Inventory, error)
	DeleteItemByLimit(playerID string, itemID uint64, limit int) error
	PlayerItemCounting(playerID string, itemID uint64) int64
}
