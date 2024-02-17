package service

import (
	_inventoryModel "github.com/Rayato159/isekai-shop-api/modules/inventory/model"
)

type InventoryService interface {
	InventoryListing(playerID string) ([]*_inventoryModel.Inventory, error)
}
