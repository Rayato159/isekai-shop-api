package service

import (
	_inventoryModel "github.com/Rayato159/isekai-shop-api/domains/inventory/model"
)

type InventoryService interface {
	Listing(playerID string) ([]*_inventoryModel.Inventory, error)
}
