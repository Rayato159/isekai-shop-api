package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error)
	Listing(playerID string) ([]*entities.Inventory, error)
	Removing(tx *gorm.DB, playerID string, itemID uint64, limit int) error
	PlayerItemCounting(playerID string, itemID uint64) int64
}
