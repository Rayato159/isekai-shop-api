package repository

import (
	_playerException "github.com/Rayato159/isekai-shop-api/domains/player/exception"
	entities "github.com/Rayato159/isekai-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type inventoryRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db *gorm.DB, logger echo.Logger) InventoryRepository {
	return &inventoryRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *inventoryRepositoryImpl) InventorySearching(playerID string) ([]*entities.Inventory, error) {
	inventories := make([]*entities.Inventory, 0)

	if err := r.db.Where(
		"player_id = ? AND is_deleted = ?", playerID, false,
	).Find(&inventories).Error; err != nil {
		r.logger.Error("Failed to find inventories", err.Error())
		return nil, &_playerException.FindInventoriesException{
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
			&entities.Inventory{},
		).Where(
			"id = ?", inventory.ID,
		).Updates(
			inventory,
		).Error; err != nil {
			r.logger.Error("Failed to delete items", err.Error())
			return &_playerException.DeleteInventoryException{ItemID: itemID}
		}
	}

	return nil
}

func (r *inventoryRepositoryImpl) InventoryFilling(inventoryEntities []*entities.Inventory) ([]*entities.Inventory, error) {
	insertedInventories := make([]*entities.Inventory, 0)

	if err := r.db.Create(inventoryEntities).Scan(&insertedInventories).Error; err != nil {
		r.logger.Error("Failed to insert items", err.Error())
		return nil, &_playerException.InsertInventoryException{}
	}

	return insertedInventories, nil
}

func (r *inventoryRepositoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {
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

func (r *inventoryRepositoryImpl) findInventoriesToDeleteByItemIDAndPlayerIDByLimit(
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
		return nil, &_playerException.FindInventoriesException{PlayerID: playerID}
	}

	for _, inventory := range inventories {
		inventory.IsDeleted = true
	}

	return inventories, nil
}
