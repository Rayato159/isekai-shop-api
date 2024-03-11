package repository

import (
	_inventoryEntity "github.com/Rayato159/isekai-shop-api/modules/inventory/entity"
	_inventoryException "github.com/Rayato159/isekai-shop-api/modules/inventory/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type inventoryRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewInventoryRepository(db *gorm.DB, logger echo.Logger) InventoryRepository {
	return &inventoryRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *inventoryRepositoryImpl) InventorySearching(playerID string) ([]*_inventoryEntity.Inventory, error) {
	inventories := make([]*_inventoryEntity.Inventory, 0)

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

func (r *inventoryRepositoryImpl) DeleteItemByLimit(playerID string, itemID uint64, limit int) error {
	inventories, err := r.findInventoriesToDeleteByItemIDAndPlayerIDByLimit(playerID, itemID, limit)
	if err != nil {
		return err
	}

	for _, inventory := range inventories {
		if err := r.db.Model(
			&_inventoryEntity.Inventory{},
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

func (r *inventoryRepositoryImpl) InventoryFilling(inventoryEntities []*_inventoryEntity.Inventory) ([]*_inventoryEntity.Inventory, error) {
	insertedInventories := make([]*_inventoryEntity.Inventory, 0)

	if err := r.db.Create(inventoryEntities).Scan(&insertedInventories).Error; err != nil {
		r.logger.Error("Failed to insert items", err.Error())
		return nil, &_inventoryException.InsertInventoryException{}
	}

	return insertedInventories, nil
}

func (r *inventoryRepositoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {
	var count int64

	if err := r.db.Model(
		&_inventoryEntity.Inventory{},
	).Where(
		"player_id = ? AND item_id = ? AND is_deleted = ?", playerID, itemID, false,
	).Count(&count).Error; err != nil {
		r.logger.Error("Failed to count player item", err.Error())
		return -1
	}

	return count
}

func (r *inventoryRepositoryImpl) findInventoriesToDeleteByItemIDAndPlayerIDByLimit(
	playerID string,
	itemID uint64,
	limit int,
) ([]*_inventoryEntity.Inventory, error) {
	inventories := make([]*_inventoryEntity.Inventory, 0)

	if err := r.db.Where(
		"player_id = ? AND item_id = ? AND is_deleted = ?", playerID, itemID, false,
	).Limit(
		limit,
	).Find(&inventories).Error; err != nil {
		r.logger.Error("Failed to find inventories", err.Error())
		return nil, &_inventoryException.FindInventoriesException{PlayerID: playerID}
	}

	for _, inventory := range inventories {
		inventory.IsDeleted = true
	}

	return inventories, nil
}
