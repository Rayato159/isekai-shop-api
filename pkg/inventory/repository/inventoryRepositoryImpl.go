package repository

import (
	"github.com/Rayato159/isekai-shop-api/databases"
	entities "github.com/Rayato159/isekai-shop-api/entities"
	_inventory "github.com/Rayato159/isekai-shop-api/pkg/inventory/exception"
	"github.com/labstack/echo/v4"
)

type inventoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger) InventoryRepository {
	return &inventoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *inventoryImpl) Filling(inventoryEntities []*entities.Inventory) ([]*entities.Inventory, error) {
	inventoryEntitiesResult := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Create(inventoryEntities).Scan(&inventoryEntitiesResult).Error; err != nil {
		r.logger.Error("Item creating failed:", err.Error())
		return nil, &_inventory.InventoryFilling{
			PlayerID: inventoryEntities[0].PlayerID,
			ItemID:   inventoryEntities[0].ItemID,
		}
	}

	return inventoryEntitiesResult, nil
}

func (r *inventoryImpl) Listing(playerID string) ([]*entities.Inventory, error) {
	inventories := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id = ? and is_deleted = ?", playerID, false,
	).Find(&inventories).Error; err != nil {
		r.logger.Error("Listing player's item failed:", err.Error())
		return nil, &_inventory.PlayerItemsFinding{
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

		if err := r.db.Connect().Model(
			&entities.Inventory{},
		).Where(
			"id = ?", inventory.ID,
		).Updates(
			inventory,
		).Error; err != nil {
			r.logger.Error("Removing item failed:", err.Error())
			return &_inventory.PlayerItemRemoving{ItemID: itemID}
		}
	}

	return nil
}

func (r *inventoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {
	var count int64

	if err := r.db.Connect().Model(
		&entities.Inventory{},
	).Where(
		"player_id = ? and item_id = ? and is_deleted = ?", playerID, itemID, false,
	).Count(&count).Error; err != nil {
		r.logger.Error("Player's item counting failed:", err.Error())
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

	if err := r.db.Connect().Where(
		"player_id = ? and item_id = ? and is_deleted = ?", playerID, itemID, false,
	).Limit(
		limit,
	).Find(&inventories).Error; err != nil {
		r.logger.Error("Finding player's item in inventory failed:", err.Error())
		return nil, &_inventory.PlayerItemsFinding{PlayerID: playerID}
	}

	return inventories, nil
}
