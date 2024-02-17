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

func (r *inventoryRepositoryImpl) InsertInventory(inventoryEntity *_inventoryEntity.Inventory) (*_inventoryEntity.Inventory, error) {
	insertedInventory := new(_inventoryEntity.Inventory)

	if err := r.db.Create(inventoryEntity).Scan(insertedInventory).Error; err != nil {
		r.logger.Error("Failed to insert item", err.Error())
		return nil, &_inventoryException.InsertInventoryException{
			PlayerID: inventoryEntity.PlayerID,
			ItemID:   inventoryEntity.ItemID,
		}
	}

	return insertedInventory, nil
}

func (r *inventoryRepositoryImpl) FindInventories(playerID string) ([]*_inventoryEntity.Inventory, error) {
	inventories := make([]*_inventoryEntity.Inventory, 0)

	if err := r.db.Where("player_id = ?", playerID).Find(&inventories).Error; err != nil {
		r.logger.Error("Failed to find inventories", err.Error())
		return nil, &_inventoryException.FindInventoriesException{
			PlayerID: playerID,
		}
	}

	return inventories, nil
}

func (r *inventoryRepositoryImpl) DeleteItem(playerID string, itemID uint64) error {
	if err := r.db.Where(
		"player_id = ? AND item_id = ?", playerID, itemID,
	).Delete(
		&_inventoryEntity.Inventory{},
	).Error; err != nil {
		r.logger.Error("Failed to delete item", err.Error())
		return &_inventoryException.DeleteInventoryException{
			ItemID: itemID,
		}
	}

	return nil
}
