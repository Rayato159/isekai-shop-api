package repository

import (
	_inventoryEntity "github.com/Rayato159/isekai-shop-api/modules/inventory/entity"
)

type InventoryRepository interface {
	InsertInventoryInBluk(inventoryEntities []*_inventoryEntity.Inventory) ([]*_inventoryEntity.Inventory, error)
	FindPlayerInventories(playerID string) ([]*_inventoryEntity.Inventory, error)
	DeleteItemByLimit(playerID string, itemID uint64, limit int) error
	CountPlayerItem(playerID string, itemID uint64) int64
}
