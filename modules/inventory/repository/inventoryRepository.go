package repository

import (
	_inventoryEntity "github.com/Rayato159/isekai-shop-api/modules/inventory/entity"
)

type InventoryRepository interface {
	InsertInventory(inventoryEntity *_inventoryEntity.Inventory) (*_inventoryEntity.Inventory, error)
	FindInventories(playerID string) ([]*_inventoryEntity.Inventory, error)
	DeleteItem(playerID string, itemID uint64) error
}
