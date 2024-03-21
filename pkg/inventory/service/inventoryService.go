package service

import (
	_inventoryModel "github.com/Rayato159/isekai-shop-api/pkg/inventory/model"
)

type InventoryService interface {
	Listing(playerID string) ([]*_inventoryModel.Inventory, error)
}
