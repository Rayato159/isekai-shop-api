package repository

import (
	_inventoryException "github.com/Rayato159/isekai-shop-api/domains/inventory/exception"
	entities "github.com/Rayato159/isekai-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type inventoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db *gorm.DB, logger echo.Logger) InventoryRepository {
	return &inventoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *inventoryImpl) Filling(inventoryEntities []*entities.Inventory) ([]*entities.Inventory, error) {
	insertedInventories := make([]*entities.Inventory, 0)

	if err := r.db.Create(inventoryEntities).Scan(&insertedInventories).Error; err != nil {
		r.logger.Error("Failed to insert items", err.Error())
		return nil, &_inventoryException.InsertInventoryException{}
	}

	return insertedInventories, nil
}

func (r *inventoryImpl) Listing(playerID string) ([]*entities.Inventory, error) {
	inventories := make([]*entities.Inventory, 0)

	if err := r.db.Where(
		"player_id = ? AND is_deleted = ?", playerID, false,
	).Find(&inventories).Error; err != nil {
		r.logger.Error("Failed to find inventories", err.Error())
		return nil, &_inventoryException.FindInventoriesException{
			PlayerID: playerID,
		}
	}

	return inventories, nil
}

func (r *inventoryImpl) Removing(playerID string, itemID uint64, limit int) error {
	inventories, err := r.findPlayerItemInInventoryByID(playerID, itemID, limit)
	if err != nil {
		return err
	}

	for _, inventory := range inventories {
		inventory.IsDeleted = true

		if err := r.db.Model(
			&entities.Inventory{},
		).Where(
			"id = ?", inventory.ID,
		).Updates(
			inventory,
		).Error; err != nil {
			r.logger.Error("Failed to delete items", err.Error())
			return &_inventoryException.DeleteInventoryException{ItemID: itemID}
		}
	}

	return nil
}

func (r *inventoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {
	var count int64

	if err := r.db.Model(
		&entities.Inventory{},
	).Where(
		"player_id = ? AND item_id = ? AND is_deleted = ?", playerID, itemID, false,
	).Count(&count).Error; err != nil {
		r.logger.Error("Failed to count player item", err.Error())
		return -1
	}

	return count
}

func (r *inventoryImpl) findPlayerItemInInventoryByID(
	playerID string,
	itemID uint64,
	limit int,
) ([]*entities.Inventory, error) {
	inventories := make([]*entities.Inventory, 0)

	if err := r.db.Where(
		"player_id = ? AND item_id = ? AND is_deleted = ?", playerID, itemID, false,
	).Limit(
		limit,
	).Find(&inventories).Error; err != nil {
		r.logger.Error("Failed to find inventories", err.Error())
		return nil, &_inventoryException.FindInventoriesException{PlayerID: playerID}
	}

	return inventories, nil
}
