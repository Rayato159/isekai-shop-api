package repository

import (
	_playerEntity "github.com/Rayato159/isekai-shop-api/domains/player/entity"
)

type InventoryRepository interface {
	InventoryFilling(inventoryEntities []*_playerEntity.Inventory) ([]*_playerEntity.Inventory, error)
	InventorySearching(playerID string) ([]*_playerEntity.Inventory, error)
	DeleteItemByLimit(playerID string, itemID uint64, limit int) error
	PlayerItemCounting(playerID string, itemID uint64) int64
}
